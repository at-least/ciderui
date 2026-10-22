// A page in Playwright's Chromium with ciderui's compiled CSS, for the
// checks: launched with --font-render-hinting=none (on Linux it hints every
// glyph to a whole-pixel advance otherwise; reference/tracking/), at a width
// a check sets, its states forced through the DevTools protocol with
// transitions off and styles flushed before they return
// (reference/apple.com/: a state's last change must land while they are off).
// `touch` emulates a touch screen, as vpkit's runner does: the DevTools
// protocol's touch emulation is what makes (hover: none) match (Chromium
// ignores an emulated hover feature); the page refuses to open otherwise.
// `forcedColors` ('active') emulates forced colors (a high-contrast theme),
// the page refusing to open unless (forced-colors: active) matches.
// Besides a property, a check reads `box-height` and `box-width`, the
// element's box, `box-top`, `box-left` and `box-right`, its place in the window,
// `box-x` and `box-y`, its place in its parent's box (a line in its SVG),
// `focused`, whether it has the focus, and `hit`, what the pointer finds at
// five points of its box (a tenth in from each corner, and its middle). One
// Chromium serves a run, each page in a context of its own (its close closes
// that); quit() closes the browser, as a runner ends.

import { chromium } from 'playwright-chromium';

let browser;
export const quit = async () => {
  if (browser) await (await browser).close();
  browser = undefined;
};

export async function open(css, html, { width = 1200, height = 800, reducedMotion, colorScheme, touch, forcedColors } = {}) {
  browser ??= chromium.launch({ args: ['--font-render-hinting=none'] });
  const context = await (await browser).newContext({ viewport: { width, height }, reducedMotion, colorScheme, forcedColors });
  const page = await context.newPage();
  const cdp = await page.context().newCDPSession(page);
  if (touch) await cdp.send('Emulation.setTouchEmulationEnabled', { enabled: true, maxTouchPoints: 5 });
  await page.setContent(`<!doctype html><html lang="en"><head><meta charset="utf-8"><style>${css}</style></head><body>${html}</body></html>`);
  // the page's fonts loaded before anything is read (a case's Inter)
  await page.evaluate(() => document.fonts.ready);
  const hover = await page.evaluate(() => matchMedia('(hover: hover)').matches);
  if (hover === Boolean(touch)) throw new Error(`(hover: hover) is ${hover} on a ${touch ? 'touch' : 'desktop'} page`);
  const forced = await page.evaluate(() => matchMedia('(forced-colors: active)').matches);
  if (forced !== (forcedColors === 'active')) throw new Error(`(forced-colors: active) is ${forced} on a page whose forcedColors is ${forcedColors}`);
  await cdp.send('DOM.enable');
  await cdp.send('CSS.enable');

  // the computed values of props on the element a selector picks, with the
  // pseudo-classes `state` forced on it (or on `on`)
  const read = async (selector, props, { state, on, pseudo } = {}) => {
    let nodeId;
    if (state) {
      const { root } = await cdp.send('DOM.getDocument', { depth: 0 });
      ({ nodeId } = await cdp.send('DOM.querySelector', { nodeId: root.nodeId, selector: on ?? selector }));
      if (!nodeId) throw new Error(`no element for ${on ?? selector}`);
      await page.addStyleTag({ content: '*, *::before, *::after { transition: none !important; animation: none !important; }' })
        .then(h => h.evaluate(e => e.setAttribute('data-cui-still', '')));
      await cdp.send('CSS.forcePseudoState', { nodeId, forcedPseudoClasses: [].concat(state) });
    }
    const values = await page.evaluate(([selector, props, pseudo]) => {
      const el = document.querySelector(selector);
      if (!el) throw new Error(`no element for ${selector}`);
      const cs = getComputedStyle(el, pseudo ?? null);
      const box = el.getBoundingClientRect();
      const parent = () => el.parentElement.getBoundingClientRect();
      const BOX = { 'box-height': () => box.height, 'box-width': () => box.width, 'box-top': () => box.top, 'box-left': () => box.left, 'box-right': () => box.right,
        'box-x': () => box.left - parent().left, 'box-y': () => box.top - parent().top };
      // each point's element, by its tag and classes
      const hit = () => [[0.1, 0.1], [0.9, 0.1], [0.1, 0.9], [0.9, 0.9], [0.5, 0.5]].map(([x, y]) => {
        const e = document.elementFromPoint(box.left + box.width * x, box.top + box.height * y);
        return e ? [e.tagName.toLowerCase(), ...e.classList].join('.') : 'none';
      }).join(' ');
      return Object.fromEntries(props.map(p => [p, p in BOX ? `${BOX[p]()}px` : p === 'focused' ? String(document.activeElement === el) : p === 'hit' ? hit() : cs.getPropertyValue(p)]));
    }, [selector, props, pseudo]);
    if (state) {
      await cdp.send('CSS.forcePseudoState', { nodeId, forcedPseudoClasses: [] });
      await page.evaluate(() => { void document.body.offsetHeight; document.querySelector('[data-cui-still]').remove(); });
    }
    return values;
  };

  const resize = (w) => page.setViewportSize({ width: w, height });
  // the page's motion run to its end: every animation and transition running
  // (or waiting out its delay) finished, an endless one aside, two frames on;
  // one cancelled on the way is done too
  const settle = () => page.evaluate(async () => {
    const frames = () => new Promise((r) => requestAnimationFrame(() => requestAnimationFrame(r)));
    await frames();
    const running = document.getAnimations().filter((a) => a.playState === 'running' && a.effect.getComputedTiming().endTime !== Infinity);
    await Promise.all(running.map((a) => a.finished.catch(() => {})));
    await frames();
  });
  return { page, read, resize, settle, close: () => context.close() };
}
