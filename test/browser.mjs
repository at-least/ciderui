import { readFileSync } from 'node:fs';
import { join } from 'node:path';
import { launch } from './chromium.mjs';
import { ROOT } from './compile.mjs';

// The stylesheet's url(fonts/…) rewritten to an origin whose fonts are served
// from memory: not inlined as data: URLs, which put every font into every
// case's page. The page stays about:blank, set by setContent: navigated to that
// origin, the pointer's place (0, 0) hovered what the page drew there; a
// <base> for it took #fragment links off the page.
const FONTS = 'http://ciderui.test/fonts/';
const fonts = new Map();
const font = (name) => {
  if (!fonts.has(name)) fonts.set(name, readFileSync(join(ROOT, 'fonts', name)));
  return fonts.get(name);
};

let browser;
export const quit = async () => {
  if (browser) await (await browser).close();
  browser = undefined;
};

export async function open(css, html, { width = 1200, height = 800, reducedMotion, colorScheme, touch, forcedColors } = {}) {
  browser ??= launch();
  const context = await (await browser).newContext({ viewport: { width, height }, reducedMotion, colorScheme, forcedColors });
  const page = await context.newPage();
  const cdp = await page.context().newCDPSession(page);
  // the DevTools protocol's touch emulation is what makes (hover: none) match:
  // Chromium ignores an emulated hover media feature
  if (touch) await cdp.send('Emulation.setTouchEmulationEnabled', { enabled: true, maxTouchPoints: 5 });
  // about:blank's origin is opaque: a font from another needs CORS
  await context.route(`${FONTS}*`, (route) => route.fulfill({ contentType: 'font/woff2', headers: { 'access-control-allow-origin': '*' }, body: font(route.request().url().slice(FONTS.length)) }));
  const served = css.replace(/url\(fonts\//g, `url(${FONTS}`);
  // a form not rewritten would load from about:blank, which fails silently:
  // its text in a fallback face, every check read off that
  if (/url\(["']?(\.\.\/)*fonts\//.test(served)) throw new Error('the stylesheet names its fonts in a form not rewritten here');
  await page.setContent(`<!doctype html><html lang="en"><head><meta charset="utf-8"><style>${served}</style></head><body>${html}</body></html>`);
  await page.evaluate(() => document.fonts.ready);
  const hover = await page.evaluate(() => matchMedia('(hover: hover)').matches);
  if (hover === Boolean(touch)) throw new Error(`(hover: hover) is ${hover} on a ${touch ? 'touch' : 'desktop'} page`);
  const forced = await page.evaluate(() => matchMedia('(forced-colors: active)').matches);
  if (forced !== (forcedColors === 'active')) throw new Error(`(forced-colors: active) is ${forced} on a page whose forcedColors is ${forcedColors}`);
  await cdp.send('DOM.enable');
  await cdp.send('CSS.enable');

  const read = async (selector, props, { state, on, pseudo } = {}) => {
    let nodeId;
    if (state) {
      const { root } = await cdp.send('DOM.getDocument', { depth: 0 });
      ({ nodeId } = await cdp.send('DOM.querySelector', { nodeId: root.nodeId, selector: on ?? selector }));
      if (!nodeId) throw new Error(`no element for ${on ?? selector}`);
      // transitions off until styles are flushed after the forced state is
      // dropped: a state's last change must land while they are off
      // (reference/apple.com/)
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
  const settle = () => page.evaluate(async () => {
    const frames = () => new Promise((r) => requestAnimationFrame(() => requestAnimationFrame(r)));
    await frames();
    const running = document.getAnimations().filter((a) => a.playState === 'running' && a.effect.getComputedTiming().endTime !== Infinity);
    await Promise.all(running.map((a) => a.finished.catch((e) => { if (e.name !== 'AbortError') throw e; })));
    await frames();
  });
  return { page, read, resize, settle, close: () => context.close() };
}
