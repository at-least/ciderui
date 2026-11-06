// @external: PAGE=<name> node measure.mjs: pages/<name>.mjs's url at each of its widths,
// each probe's computed styles into results/<name>.json. A probe's keys are what
// this file reads. A probe that matches nothing is recorded as null and fails the run.
// What reading it cannot show:
// - browser `chrome`: Google Chrome plays what needs Widevine (chrome.Dockerfile's
//   image, CIDERUI_CHROME=1).
// - select's ` >>> ` steps into a shadow root: tv.apple.com's controls keep their
//   parts behind theirs. inject: for a class the page defines but may not use.
// - states are forced through the DevTools protocol, transitions off, animations left
//   running: ended, a filled animation's value read as its base's, the state's own as
//   it seemed, and it ran again once the still style went (ribbon:ribbon.link,
//   2026-10-22). A frame's element is found in the page's own session, so only a
//   frame in the page's process can have its states forced; another fails the run.
// - set's `value` sets an input's value with no event: the page's script hears nothing.
// - click leaves the page so for later probes and widths, so a spec puts such probes
//   last or in a spec of their own. The pointer then leaves for the window's bottom
//   corner, or `pointer` inside what opened, where leaving it closes it.
// - tap is on a touch spec alone, through the DevTools protocol's touch events, which
//   Chromium's gesture recognizer makes a tap of; a page scrolled after it keeps what
//   it hovered.
// - hover moves in twelve steps, as a hand does: apple.com's script opens a flyout on
//   its own mouseover, 120ms on.
// - image never records the picture, the page's asset: what it draws (viewBox, shapes,
//   paint); any other value than a base64 data: SVG fails the run.
// - touch: the DevTools protocol's emulation before a reload makes (hover: none) match;
//   the run fails if (hover: hover) still matches.
import { mkdirSync, writeFileSync } from 'node:fs';
import { open } from '../browsers.mjs';

const PAGE = process.env.PAGE;
const spec = (await import(`./pages/${PAGE}.mjs`)).default;
const widths = spec.widths ?? [1600, 1200, 900, 390];

// @invariant: in the frame: tag the probe's element and the one its states are forced on;
// its outerHTML's head, or null
const place = async (probe) => {
  const y = probe.scroll ?? 0;
  if (window.scrollY !== y) {
    window.scrollTo(0, y);
    await new Promise(r => requestAnimationFrame(() => requestAnimationFrame(() => setTimeout(r, 900))));
  }
  if (probe.blur) document.activeElement?.blur();
  for (const [selector, name, value] of probe.set ?? []) {
    const target = document.querySelector(selector);
    if (!target) return null;
    if (name === 'value') target.value = value;
    else if (name === 'classList') target.classList.add(value);
    else if (value === null) target.removeAttribute(name);
    else target.setAttribute(name, value);
  }
  // @external: what a blur or a set moves (the page's script checking a field, a
  // label's transition) settled before it is read
  if (probe.blur || probe.set) await new Promise(r => requestAnimationFrame(() => requestAnimationFrame(() => setTimeout(r, 900))));
  const roots = (root) => [root, ...[...root.querySelectorAll('*')].filter(e => e.shadowRoot).flatMap(e => roots(e.shadowRoot))];
  for (const root of roots(document)) {
    root.querySelectorAll('[data-cui-probe]').forEach(e => e.removeAttribute('data-cui-probe'));
    root.querySelectorAll('[data-cui-state]').forEach(e => e.removeAttribute('data-cui-state'));
  }
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
    el = probe.select.split(' >>> ').reduce((root, part) => root && (root === document ? root : root.shadowRoot)?.querySelector(part), document);
  }
  if (!el) return null;
  const on = probe.stateOn ? scope.querySelector(probe.stateOn) ?? document.querySelector(probe.stateOn) : el;
  if (!on) return null;
  el.setAttribute('data-cui-probe', '');
  on.setAttribute('data-cui-state', '');
  // @external: without the probe's tags, and without query strings (the store's carry
  // session ids)
  return el.outerHTML.replace(/ data-cui-[a-z]+=""/g, '').replace(/(href|src|srcset)="([^"?]*)\?[^"]*"/g, '$1="$2?…"').slice(0, 240);
};

// @invariant: the head as a result records it, without the page's words: its text, and
// the values of the attributes that hold words (one the head cuts too), so a
// result says where a value was read, not what the page says
const strip = (head) => head
  .replace(/>[^<]*/g, '>')
  .replace(/ (aria-label|alt|title|placeholder)="[^"]+("|$)/g, ' $1="…$2');

// @invariant: in the frame: from now on, every 4ms, describe every animation not seen
// yet whose element is inside `within`
const watchAnimations = (within) => {
  const seen = new Set();
  const t0 = performance.now();
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
        kind: a.constructor.name, what, on: name(target), duration: timing.duration, delay: Math.round(timing.delay * 100) / 100, direction: timing.direction,
        // @external: a transition's timing is its effect's, an animation's its keyframes'
        easing: a.transitionProperty ? timing.easing : frames[0]?.easing,
        from: prop ? frames[0][prop] : null, to: prop ? frames.at(-1)[prop] : null,
        seen: Math.round(performance.now() - t0),
      });
    }
  };
  window.cuiWatch = setInterval(look, 4);
};

const read = (probe) => {
  // @external: in the document or a shadow root in it
  const walk = (root, s) => root.querySelector(s) ?? [...root.querySelectorAll('*')].filter(e => e.shadowRoot).map(e => walk(e.shadowRoot, s)).find(Boolean) ?? null;
  const el = walk(document, '[data-cui-probe]');
  const cs = getComputedStyle(el, probe.pseudo ?? null);
  const out = { styles: {} };
  for (const p of probe.props ?? []) out.styles[p] = cs.getPropertyValue(p);
  for (const v of probe.vars ?? []) out.styles[v] = cs.getPropertyValue(v).trim();
  for (const p of probe.image ?? []) {
    const m = cs.getPropertyValue(p).match(/^url\("data:image\/svg\+xml;base64,([A-Za-z0-9+/=]+)"\)$/);
    if (!m) throw new Error(`${probe.name}: ${p} is not a base64 SVG`);
    const svg = new DOMParser().parseFromString(atob(m[1]), 'image/svg+xml').documentElement;
    const paint = (e, a) => e.closest(`[${a}]`)?.getAttribute(a) ?? null;
    out.styles[p] = { viewBox: svg.getAttribute('viewBox'), shapes: [...svg.querySelectorAll('circle, ellipse, rect, path, polygon, polyline, line')].map(e => ({ shape: e.tagName, fill: paint(e, 'fill'), stroke: paint(e, 'stroke') })) };
  }
  if (probe.keyframes) {
    out.keyframes = {};
    for (const name of cs.animationName.split(',').map(n => n.trim()).filter(n => n !== 'none')) {
      let rule = null;
      for (const sheet of document.styleSheets) {
        // @external: a stylesheet from another origin cannot be read: the rule, if it is
        // there, stays not found, and the run fails
        let rules;
        try { rules = sheet.cssRules; } catch { continue; }
        rule = [...rules].find(r => r.type === CSSRule.KEYFRAMES_RULE && r.name === name) ?? rule;
      }
      out.keyframes[name] = rule && [...rule.cssRules].map(k => ({ at: k.keyText, style: k.style.cssText }));
    }
  }
  if (probe.sections) {
    const [section, attribute, row] = probe.sections;
    out.sections = [...el.querySelectorAll(section)].map(s => [s.getAttribute(attribute), s.querySelectorAll(row).length]);
  }
  if (probe.matches) out.matches = Object.fromEntries(probe.matches.map(s => [s, el.matches(s)]));
  if (probe.box) {
    const r = el.getBoundingClientRect();
    out.box = { width: r.width, height: r.height, ...(probe.position ? { top: r.top, left: r.left } : {}) };
  }
  return out;
};

// @external: `options`: Playwright's page options for the run (a permission it needs)
const page = await open(spec.browser ?? 'chromium', spec.url, spec.options);
const pw = page.page;
// @external: `colorScheme`: the page read in that scheme (prefers-color-scheme), from its load
if (spec.colorScheme) {
  await pw.emulateMedia({ colorScheme: spec.colorScheme });
  await pw.reload({ waitUntil: 'load' });
}
const cdp = await pw.context().newCDPSession(pw);
// @external: `touch`: the page read on a touch screen, from its load
if (spec.touch) {
  await cdp.send('Emulation.setTouchEmulationEnabled', { enabled: true, maxTouchPoints: 5 });
  await pw.reload({ waitUntil: 'load' });
  if (await pw.evaluate(() => matchMedia('(hover: hover)').matches)) throw new Error('(hover: hover) matches on the emulated touch screen');
}
await cdp.send('DOM.enable');
await cdp.send('CSS.enable');
// @external: the DevTools node of the element tagged data-cui-state in this frame: a
// frame's document is reached through its iframe in the page's whole tree
const stateNode = async (frame) => {
  if (frame === pw.mainFrame()) {
    const { root } = await cdp.send('DOM.getDocument', { depth: 0 });
    const { nodeId } = await cdp.send('DOM.querySelector', { nodeId: root.nodeId, selector: '[data-cui-state]' });
    if (nodeId) return nodeId;
    // @external: in a shadow root, which DOM.querySelector does not enter
    const whole = await cdp.send('DOM.getDocument', { depth: -1, pierce: true });
    const find = (n) => n.attributes?.includes('data-cui-state') ? n
      : [...(n.children ?? []), ...(n.shadowRoots ?? [])].map(find).find(Boolean);
    return find(whole.root)?.nodeId;
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
const clicked = new Set();
// @invariant: a point a probe names: [x, y], or a function of the window's width giving one
const point = (p, width) => typeof p === 'function' ? p(width) : p;
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
      // @external: a click once, at the first width the probe is read: what it opens
      // stays so as the window narrows
      const clicking = (probe.click || probe.type || probe.tap || probe.hover) && !clicked.has(probe.name);
      const settle = () => pw.evaluate(() => new Promise(r => requestAnimationFrame(() => requestAnimationFrame(() => setTimeout(r, 900)))));
      if (probe.hover) {
        // @external: at every width: what the pointer opens a window's resize may close
        if (spec.touch) throw new Error(`${probe.name}: a hover on a touch spec`);
        clicked.add(probe.name);
        if (clicking && probe.animations) await frame.evaluate(watchAnimations, probe.animations);
        // @invariant: a probe's scroll before the pointer moves, so what it hovers is in
        // the window (the read's own scroll then finds it there)
        if (probe.scroll !== undefined) {
          await frame.evaluate((y) => window.scrollTo(0, y), probe.scroll);
          await settle();
        }
        let to = point(probe.hover, width);
        if (!Array.isArray(to)) {
          const r = await frame.locator(probe.hover).first().boundingBox();
          if (!r?.width || !r.height) throw new Error(`${probe.name}: ${probe.hover} has no box to hover`);
          to = [r.x + r.width / 2, r.y + r.height / 2];
        }
        await pw.mouse.move(...to, { steps: 12 });
        await pw.waitForTimeout(probe.hovered ?? 1500);
        if (probe.pointer) {
          await pw.mouse.move(...point(probe.pointer, width), { steps: 6 });
          await settle();
        }
      } else if (clicking && probe.tap) {
        if (!spec.touch) throw new Error(`${probe.name}: a tap on a spec that is not touch`);
        clicked.add(probe.name);
        if (probe.animations) await frame.evaluate(watchAnimations, probe.animations);
        const target = frame.locator(probe.tap).first();
        await target.evaluate(e => e.scrollIntoView({ block: 'center' }));
        await settle();
        const r = await target.boundingBox();
        if (!r?.width || !r.height) throw new Error(`${probe.name}: ${probe.tap} has no box to tap`);
        await cdp.send('Input.dispatchTouchEvent', { type: 'touchStart', touchPoints: [{ x: r.x + r.width / 2, y: r.y + r.height / 2 }] });
        await cdp.send('Input.dispatchTouchEvent', { type: 'touchEnd', touchPoints: [] });
        await settle();
      } else if (clicking && probe.type) {
        clicked.add(probe.name);
        if (probe.animations) await frame.evaluate(watchAnimations, probe.animations);
        for (const key of probe.type) {
          await pw.keyboard.press(key);
          await pw.waitForTimeout(150);
        }
        await pw.waitForTimeout(probe.typed ?? 2500);
      } else if (clicking) {
        clicked.add(probe.name);
        if (probe.animations) await frame.evaluate(watchAnimations, probe.animations);
        // @invariant: a point, [x, y]: the window clicked there (what a layer over the
        // element would take a click on it for)
        if (Array.isArray(probe.click)) await pw.mouse.click(...probe.click);
        else await frame.click(probe.click);
        // @external: the pointer off what it clicked, which it would leave hovered, to the
        // window's bottom corner; or, once what the click opened has opened, to
        // `pointer`, a point inside it, which the pointer leaving would close
        // (moved at once, it lands beyond what is still growing)
        if (probe.pointer) await settle();
        await pw.mouse.move(...(point(probe.pointer, width) ?? [0, 899]));
        if (probe.clicked) await pw.waitForTimeout(probe.clicked);
        else await settle();
      }
      used.add(frame);
      // @invariant: what the page is handed of the probe: its data, not its functions (a
      // point given by the width)
      const data = JSON.parse(JSON.stringify(probe));
      const html = await frame.evaluate(place, data);
      if (html === null) {
        at[probe.name] = null;
        nulls.push(`${width} ${probe.name}`);
        console.log(`NULL ${width}px ${probe.name}: ${probe.select ?? ''} ${probe.inject ?? ''} ${probe.into ?? ''} ${probe.stateOn ?? ''}`);
        continue;
      }
      at[probe.name] = { html: strip(html), ...(await frame.evaluate(read, data)) };
      for (const [k, v] of Object.entries(at[probe.name].keyframes ?? {})) if (!v) nulls.push(`${width} ${probe.name} @keyframes ${k}`);
      if (clicking && probe.animations) at[probe.name].animations = await frame.evaluate(() => { clearInterval(window.cuiWatch); return window.cuiAnimations; });
      if (probe.states?.length) {
        const nodeId = await stateNode(frame);
        if (!nodeId) throw new Error(`${probe.name}: the DevTools session does not see the tagged element`);
        await frame.addStyleTag({ content: '*, *::before, *::after { transition: none !important; }' })
          .then(h => h.evaluate(e => e.setAttribute('data-cui-still', '')));
        // @external: a document's style does not reach into a shadow root: one there too
        await frame.evaluate(() => {
          const walk = (root, s) => root.querySelector(s) ?? [...root.querySelectorAll('*')].filter(e => e.shadowRoot).map(e => walk(e.shadowRoot, s)).find(Boolean) ?? null;
          const root = walk(document, '[data-cui-state]').getRootNode();
          if (root === document) return;
          const still = document.createElement('style');
          still.textContent = '*, *::before, *::after { transition: none !important; }';
          still.setAttribute('data-cui-still', '');
          root.append(still);
        });
        at[probe.name].states = {};
        for (const state of probe.states) {
          // @invariant: a state may be several at once: 'focus+focus-visible'
          await cdp.send('CSS.forcePseudoState', { nodeId, forcedPseudoClasses: state.split('+') });
          at[probe.name].states[state] = (await frame.evaluate(read, data)).styles;
          await cdp.send('CSS.forcePseudoState', { nodeId, forcedPseudoClasses: [] });
        }
        // @invariant: the state's last change must land while transitions are still off, or
        // it starts one that a later probe reads in flight: flush, then remove
        await frame.evaluate(() => {
          const walk = (root, s) => root.querySelector(s) ?? [...root.querySelectorAll('*')].filter(e => e.shadowRoot).map(e => walk(e.shadowRoot, s)).find(Boolean) ?? null;
          void document.body.offsetHeight;
          for (let still; (still = walk(document, '[data-cui-still]'));) still.remove();
        });
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
