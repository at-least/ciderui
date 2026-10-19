// PAGE=<name> node measure.mjs
// Opens pages/<name>.mjs's url in Playwright's Chromium (../browsers.mjs) and
// at each of its widths (after `settle` ms, 800 unless the spec says; `wait`
// ms more once, after the load) reads every probe's computed styles. A probe either
// selects an element of the page (`select`) or adds markup to it (`inject`,
// into `into` or the body; `select` then picks the element inside it) for a
// class the page defines but may not use, and reads `props`, `vars` (custom
// properties) and `pseudo`'s, its box with `box`, and with `states` the same
// again with each pseudo-class forced through the DevTools protocol on the
// element itself or on `stateOn` (an input whose label is read), transitions
// off while it is forced; a frame's element is found in the page's own
// session, so only a frame in the page's process can have its states forced
// (another fails the run). `scroll` scrolls the page to that many pixels first
// (for what shows only once scrolled), every other probe reading at the top;
// `blur` takes the focus off the element that has it (a field the page
// focuses on load) before reading; `frame` reads in the frame whose URL holds
// that text; `widths` reads at those of the spec's widths alone; `click`
// clicks that element first (a menu's button: the page stays so for the
// probes after it at that width, so a spec puts such probes last);
// `position` adds the box's top and left in the window (where the window is
// what it is placed in: an open menu's rows); `animations`, with `click`,
// records every animation and transition the click starts on an element
// inside that selector, each frame for the wait after it (what it animates,
// on what, its duration, delay and timing, its first and last values);
// `keyframes` records the @keyframes its animation-name names (read from
// the page's own stylesheets; one the page's CSS does not hold is null and
// fails the run). A probe that matches nothing is recorded as null, printed,
// and fails the run. The result is results/<name>.json.
import { mkdirSync, writeFileSync } from 'node:fs';
import { open } from '../browsers.mjs';

const PAGE = process.env.PAGE;
const spec = (await import(`./pages/${PAGE}.mjs`)).default;
const widths = spec.widths ?? [1600, 1200, 900, 390];

// in the frame: tag the probe's element and the one its states are forced on;
// its outerHTML's head, or null
const place = async (probe) => {
  const y = probe.scroll ?? 0;
  if (window.scrollY !== y) {
    window.scrollTo(0, y);
    await new Promise(r => requestAnimationFrame(() => requestAnimationFrame(() => setTimeout(r, 900))));
  }
  if (probe.blur) document.activeElement?.blur();
  document.querySelectorAll('[data-cui-probe]').forEach(e => e.removeAttribute('data-cui-probe'));
  document.querySelectorAll('[data-cui-state]').forEach(e => e.removeAttribute('data-cui-state'));
  document.querySelectorAll('[data-cui-injected]').forEach(e => e.remove());
  let scope = document;
  let el;
  if (probe.inject) {
    const host = probe.into ? document.querySelector(probe.into) : document.body;
    if (!host) return null;
    const t = document.createElement('template');
    t.innerHTML = probe.inject.trim();
    const node = t.content.firstElementChild;
    node.setAttribute('data-cui-injected', '');
    host.append(node);
    scope = node;
    el = probe.select ? (node.matches(probe.select) ? node : node.querySelector(probe.select)) : node;
  } else {
    el = document.querySelector(probe.select);
  }
  if (!el) return null;
  const on = probe.stateOn ? scope.querySelector(probe.stateOn) ?? document.querySelector(probe.stateOn) : el;
  if (!on) return null;
  el.setAttribute('data-cui-probe', '');
  on.setAttribute('data-cui-state', '');
  // without the probe's tags, and without query strings (the store's carry
  // session ids)
  return el.outerHTML.replace(/ data-cui-[a-z]+=""/g, '').replace(/(href|src|srcset)="([^"?]*)\?[^"]*"/g, '$1="$2?…"').slice(0, 240);
};

// in the frame: from now on, each frame, describe every animation not seen
// yet whose element is inside `within`
const watchAnimations = (within) => {
  const seen = new Set();
  window.cuiAnimations = [];
  const name = (e) => `${e.tagName.toLowerCase()}${e.id ? '#' + e.id : ''}${[...e.classList].slice(0, 3).map(c => '.' + c).join('')}`;
  const look = () => {
    for (const a of document.getAnimations()) {
      const target = a.effect?.target;
      if (seen.has(a) || !target?.closest?.(within)) continue;
      seen.add(a);
      const timing = a.effect.getTiming();
      const frames = a.effect.getKeyframes();
      const what = a.transitionProperty ?? a.animationName;
      const prop = Object.keys(frames[0] ?? {}).find(k => !['offset', 'computedOffset', 'easing', 'composite'].includes(k));
      window.cuiAnimations.push({
        kind: a.constructor.name, what, on: name(target), duration: timing.duration, delay: Math.round(timing.delay * 100) / 100,
        // a transition's timing is its effect's, an animation's its keyframes'
        easing: a.transitionProperty ? timing.easing : frames[0]?.easing,
        from: prop ? frames[0][prop] : null, to: prop ? frames.at(-1)[prop] : null,
      });
    }
  };
  window.cuiWatch = setInterval(look, 4);
};

const read = (probe) => {
  const el = document.querySelector('[data-cui-probe]');
  const cs = getComputedStyle(el, probe.pseudo ?? null);
  const out = { styles: {} };
  for (const p of probe.props ?? []) out.styles[p] = cs.getPropertyValue(p);
  for (const v of probe.vars ?? []) out.styles[v] = cs.getPropertyValue(v).trim();
  if (probe.keyframes) {
    out.keyframes = {};
    for (const name of cs.animationName.split(',').map(n => n.trim()).filter(n => n !== 'none')) {
      let rule = null;
      for (const sheet of document.styleSheets) {
        // a stylesheet from another origin cannot be read: the rule, if it is
        // there, stays not found, and the run fails
        let rules;
        try { rules = sheet.cssRules; } catch { continue; }
        rule = [...rules].find(r => r.type === CSSRule.KEYFRAMES_RULE && r.name === name) ?? rule;
      }
      out.keyframes[name] = rule && [...rule.cssRules].map(k => ({ at: k.keyText, style: k.style.cssText }));
    }
  }
  if (probe.box) {
    const r = el.getBoundingClientRect();
    out.box = { width: r.width, height: r.height, ...(probe.position ? { top: r.top, left: r.left } : {}) };
  }
  return out;
};

const page = await open('chromium', spec.url);
const pw = page.page;
const cdp = await pw.context().newCDPSession(pw);
await cdp.send('DOM.enable');
await cdp.send('CSS.enable');
// the DevTools node of the element tagged data-cui-state in this frame: a
// frame's document is reached through its iframe in the page's whole tree
const stateNode = async (frame) => {
  if (frame === pw.mainFrame()) {
    const { root } = await cdp.send('DOM.getDocument', { depth: 0 });
    return (await cdp.send('DOM.querySelector', { nodeId: root.nodeId, selector: '[data-cui-state]' })).nodeId;
  }
  const iframe = await frame.frameElement();
  await iframe.evaluate(e => e.setAttribute('data-cui-frame', ''));
  const { root } = await cdp.send('DOM.getDocument', { depth: -1, pierce: true });
  await iframe.evaluate(e => e.removeAttribute('data-cui-frame'));
  const find = (n) => n.attributes?.includes('data-cui-frame') ? n
    : [...(n.children ?? []), ...(n.contentDocument ? [n.contentDocument] : []), ...(n.shadowRoots ?? [])].map(find).find(Boolean);
  const doc = find(root)?.contentDocument;
  if (!doc) throw new Error(`the page's session does not reach the document of ${frame.url()}`);
  return (await cdp.send('DOM.querySelector', { nodeId: doc.nodeId, selector: '[data-cui-state]' })).nodeId;
};
const frameOf = (probe) => {
  if (!probe.frame) return pw.mainFrame();
  const f = pw.frames().find(f => f.url().includes(probe.frame));
  if (!f) throw new Error(`no frame whose URL holds ${probe.frame}`);
  return f;
};

const final = new URL(pw.url());
const result = { url: spec.url, finalUrl: final.origin + final.pathname, browser: page.label, date: new Date().toISOString(), widths: {} };
const nulls = [];
const used = new Set();
try {
  if (spec.wait) await pw.waitForTimeout(spec.wait);
  for (const width of widths) {
    await pw.setViewportSize({ width, height: 900 });
    await pw.evaluate((ms) => new Promise(r => requestAnimationFrame(() => requestAnimationFrame(() => setTimeout(r, ms)))), spec.settle ?? 800);
    const at = (result.widths[width] = {});
    for (const probe of spec.probes) {
      if (probe.widths && !probe.widths.includes(width)) continue;
      const frame = frameOf(probe);
      if (probe.click) {
        if (probe.animations) await frame.evaluate(watchAnimations, probe.animations);
        await frame.click(probe.click);
        await pw.evaluate(() => new Promise(r => requestAnimationFrame(() => requestAnimationFrame(() => setTimeout(r, 900)))));
      }
      used.add(frame);
      const html = await frame.evaluate(place, probe);
      if (html === null) {
        at[probe.name] = null;
        nulls.push(`${width} ${probe.name}`);
        console.log(`NULL ${width}px ${probe.name}: ${probe.select ?? ''} ${probe.inject ?? ''} ${probe.into ?? ''} ${probe.stateOn ?? ''}`);
        continue;
      }
      at[probe.name] = { html, ...(await frame.evaluate(read, probe)) };
      for (const [k, v] of Object.entries(at[probe.name].keyframes ?? {})) if (!v) nulls.push(`${width} ${probe.name} @keyframes ${k}`);
      if (probe.click && probe.animations) at[probe.name].animations = await frame.evaluate(() => { clearInterval(window.cuiWatch); return window.cuiAnimations; });
      if (probe.states?.length) {
        const nodeId = await stateNode(frame);
        if (!nodeId) throw new Error(`${probe.name}: the DevTools session does not see the tagged element`);
        await frame.addStyleTag({ content: '*, *::before, *::after { transition: none !important; animation: none !important; }' })
          .then(h => h.evaluate(e => e.setAttribute('data-cui-still', '')));
        at[probe.name].states = {};
        for (const state of probe.states) {
          await cdp.send('CSS.forcePseudoState', { nodeId, forcedPseudoClasses: [state] });
          at[probe.name].states[state] = (await frame.evaluate(read, probe)).styles;
          await cdp.send('CSS.forcePseudoState', { nodeId, forcedPseudoClasses: [] });
        }
        // the state's last change must land while transitions are still off, or
        // it starts one that a later probe reads in flight: flush, then remove
        await frame.evaluate(() => { void document.body.offsetHeight; document.querySelector('[data-cui-still]').remove(); });
      }
    }
    for (const frame of used) await frame.evaluate(() => document.querySelectorAll('[data-cui-injected]').forEach(e => e.remove()));
  }
} finally {
  await page.close();
}
mkdirSync(new URL('./results/', import.meta.url), { recursive: true });
writeFileSync(new URL(`./results/${PAGE}.json`, import.meta.url), JSON.stringify(result, null, 1));
console.log(`${PAGE}: ${spec.probes.length} probes at ${widths.join(', ')}px, ${nulls.length} null; ${result.browser}; ${result.finalUrl}`);
process.exit(nulls.length ? 1 : 0);
