// A page in Playwright's Chromium with ciderui's compiled CSS, for the
// checks: launched with --font-render-hinting=none (on Linux it hints every
// glyph to a whole-pixel advance otherwise; reference/tracking/), at a width
// a check sets, its states forced through the DevTools protocol with
// transitions off and styles flushed before they return
// (reference/apple.com/: a state's last change must land while they are off).

import { chromium } from 'playwright-chromium';

export async function open(css, html, { width = 1200, reducedMotion, colorScheme } = {}) {
  const browser = await chromium.launch({ args: ['--font-render-hinting=none'] });
  const page = await browser.newPage({ viewport: { width, height: 800 }, reducedMotion, colorScheme });
  await page.setContent(`<!doctype html><html lang="en"><head><meta charset="utf-8"><style>${css}</style></head><body>${html}</body></html>`);
  const cdp = await page.context().newCDPSession(page);
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
      return Object.fromEntries(props.map(p => [p, cs.getPropertyValue(p)]));
    }, [selector, props, pseudo]);
    if (state) {
      await cdp.send('CSS.forcePseudoState', { nodeId, forcedPseudoClasses: [] });
      await page.evaluate(() => { void document.body.offsetHeight; document.querySelector('[data-cui-still]').remove(); });
    }
    return values;
  };

  const resize = (w) => page.setViewportSize({ width: w, height: 800 });
  return { page, read, resize, close: () => browser.close() };
}
