// PAGE=<name> node measure.mjs
// Opens pages/<name>.mjs's url in Playwright's Chromium (../browsers.mjs) and
// at each of its widths (after `settle` ms, 800 unless the spec says; `wait`
// ms more once, after the load) reads every probe's computed styles. A probe either
// selects an element of the page (`select`) or adds markup to it (`inject`,
// into `into` or the body; `select` then picks the element inside it) for a
// class the page defines but may not use, and reads `props`, `vars` (custom
// properties) and `pseudo`'s, its text with `text` (its spaces collapsed),
// its box with `box`, and with `states` the same
// again with each pseudo-class (or several, 'focus+focus-visible') forced
// through the DevTools protocol on the
// element itself or on `stateOn` (an input whose label is read), transitions
// off while it is forced, its animations left as they run (ended, a filled
// animation's value read as its base's, the state's own as it seemed, and
// run again once the still style went: ribbon:ribbon.link, 2026-10-22); a
// frame's element is found in the page's own
// session, so only a frame in the page's process can have its states forced
// (another fails the run). `scroll` scrolls the page to that many pixels first
// (for what shows only once scrolled), every other probe reading at the top;
// `blur` takes the focus off the element that has it (a field the page
// focuses on load) before reading; `set`, [[selector, name, value], …], gives
// the page's elements what its script would before reading (an attribute
// set, or removed with null; `value` an input's value, set with no event, so
// the page's script hears nothing; `classList` a class added), then waits for what that moves (its
// transitions), and they keep it for the probes after; `frame` reads in the frame whose URL holds
// that text; `widths` reads at those of the spec's widths alone; `click`
// clicks that element first, once, at the first width the probe is read (a
// menu's button, an accordion's: the page stays so for the probes after it
// and at the widths after, so a spec puts such probes last, or in a spec of
// their own; the pointer then leaves for the window's bottom corner, or, once
// what the click opened has opened, for `pointer`, [x, y], inside it, when
// leaving it closes it); `type` presses those keys, once, at the first
// width the probe is read, into what has the focus (a field: typed one key
// after another, 150ms apart, then `typed` ms, 2500 unless it says, for what
// the page answers), and, with `animations`, records what they start; `tap`
// taps that element, once, at the first width the probe is read, on a touch
// spec alone: scrolled to the window's middle, its middle touched and let go
// through the DevTools protocol's touch events, which Chromium's gesture
// recognizer makes a tap of, as a touch screen's (its pointer events `touch`
// and trusted, the mouse events and the click after them; a page scrolled
// after it keeps what it hovered); `hover` moves the pointer there, at each
// width the probe is read, on a spec that is not touch: to that element's
// middle (or a point, [x, y], or what a function of the width gives) in
// twelve steps, as a hand does (apple.com's script opens a flyout on its own
// mouseover, 120ms on), then `hovered` ms, 1500 unless it says, and then, with
// `pointer`, on to that point in six steps (a menu's own room, where the
// pointer hovers nothing); with `animations` it records what they start at
// the first width; `matches` records whether the element
// matches each of those selectors (':hover', what a tap left);
// `position` adds the box's top and left in the window (where the window is
// what it is placed in: an open menu's rows); `animations`, with `click`,
// records every animation and transition the click starts on an element
// inside that selector, each frame for the wait after it (what it animates,
// on what, its duration, delay, direction and timing, its first and last values, and
// when it was first seen, in ms from the click or the first key);
// `keyframes` records the @keyframes its animation-name names (read from
// the page's own stylesheets; one the page's CSS does not hold is null and
// fails the run). A spec's `colorScheme` reads the page in that color scheme
// (`prefers-color-scheme`, emulated before a reload); its `touch` on a touch
// screen (the DevTools protocol's touch emulation before a reload, which
// makes (hover: none) match, the page's script finding touch; the run fails
// if (hover: hover) still matches). A probe that matches nothing is recorded as null, printed,
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
  for (const [selector, name, value] of probe.set ?? []) {
    const target = document.querySelector(selector);
    if (!target) return null;
    if (name === 'value') target.value = value;
    else if (name === 'classList') target.classList.add(value);
    else if (value === null) target.removeAttribute(name);
    else target.setAttribute(name, value);
  }
  if (probe.set) await new Promise(r => requestAnimationFrame(() => requestAnimationFrame(() => setTimeout(r, 900))));
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
        // a transition's timing is its effect's, an animation's its keyframes'
        easing: a.transitionProperty ? timing.easing : frames[0]?.easing,
        from: prop ? frames[0][prop] : null, to: prop ? frames.at(-1)[prop] : null,
        seen: Math.round(performance.now() - t0),
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
  if (probe.text) out.text = el.textContent.replace(/\s+/g, ' ').trim();
  if (probe.matches) out.matches = Object.fromEntries(probe.matches.map(s => [s, el.matches(s)]));
  if (probe.box) {
    const r = el.getBoundingClientRect();
    out.box = { width: r.width, height: r.height, ...(probe.position ? { top: r.top, left: r.left } : {}) };
  }
  return out;
};

const page = await open('chromium', spec.url);
const pw = page.page;
// `colorScheme`: the page read in that scheme (prefers-color-scheme), from its load
if (spec.colorScheme) {
  await pw.emulateMedia({ colorScheme: spec.colorScheme });
  await pw.reload({ waitUntil: 'load' });
}
const cdp = await pw.context().newCDPSession(pw);
// `touch`: the page read on a touch screen, from its load
if (spec.touch) {
  await cdp.send('Emulation.setTouchEmulationEnabled', { enabled: true, maxTouchPoints: 5 });
  await pw.reload({ waitUntil: 'load' });
  if (await pw.evaluate(() => matchMedia('(hover: hover)').matches)) throw new Error('(hover: hover) matches on the emulated touch screen');
}
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
const clicked = new Set();
// a point a probe names: [x, y], or a function of the window's width giving one
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
      // a click once, at the first width the probe is read: what it opens
      // stays so as the window narrows
      const clicking = (probe.click || probe.type || probe.tap || probe.hover) && !clicked.has(probe.name);
      const settle = () => pw.evaluate(() => new Promise(r => requestAnimationFrame(() => requestAnimationFrame(() => setTimeout(r, 900)))));
      if (probe.hover) {
        // at every width: what the pointer opens a window's resize may close
        if (spec.touch) throw new Error(`${probe.name}: a hover on a touch spec`);
        clicked.add(probe.name);
        if (clicking && probe.animations) await frame.evaluate(watchAnimations, probe.animations);
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
        await frame.click(probe.click);
        // the pointer off what it clicked, which it would leave hovered, to the
        // window's bottom corner; or, once what the click opened has opened, to
        // `pointer`, a point inside it, which the pointer leaving would close
        // (moved at once, it lands beyond what is still growing)
        if (probe.pointer) await settle();
        await pw.mouse.move(...(point(probe.pointer, width) ?? [0, 899]));
        await settle();
      }
      used.add(frame);
      // what the page is handed of the probe: its data, not its functions (a
      // point given by the width)
      const data = JSON.parse(JSON.stringify(probe));
      const html = await frame.evaluate(place, data);
      if (html === null) {
        at[probe.name] = null;
        nulls.push(`${width} ${probe.name}`);
        console.log(`NULL ${width}px ${probe.name}: ${probe.select ?? ''} ${probe.inject ?? ''} ${probe.into ?? ''} ${probe.stateOn ?? ''}`);
        continue;
      }
      at[probe.name] = { html, ...(await frame.evaluate(read, data)) };
      for (const [k, v] of Object.entries(at[probe.name].keyframes ?? {})) if (!v) nulls.push(`${width} ${probe.name} @keyframes ${k}`);
      if (clicking && probe.animations) at[probe.name].animations = await frame.evaluate(() => { clearInterval(window.cuiWatch); return window.cuiAnimations; });
      if (probe.states?.length) {
        const nodeId = await stateNode(frame);
        if (!nodeId) throw new Error(`${probe.name}: the DevTools session does not see the tagged element`);
        await frame.addStyleTag({ content: '*, *::before, *::after { transition: none !important; }' })
          .then(h => h.evaluate(e => e.setAttribute('data-cui-still', '')));
        at[probe.name].states = {};
        for (const state of probe.states) {
          // a state may be several at once: 'focus+focus-visible'
          await cdp.send('CSS.forcePseudoState', { nodeId, forcedPseudoClasses: state.split('+') });
          at[probe.name].states[state] = (await frame.evaluate(read, data)).styles;
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
