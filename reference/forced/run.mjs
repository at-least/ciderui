// node run.mjs: MEASURED.md holds its claims and method.
import { mkdirSync, readFileSync, writeFileSync } from 'node:fs';
import { chromium, firefox } from 'playwright-core';

const html = readFileSync(new URL('./page.html', import.meta.url), 'utf8');
const MODES = [
  { name: 'chromium', engine: chromium, forced: false, launch: { args: ['--font-render-hinting=none'] }, page: { forcedColors: 'none' } },
  { name: 'chromium-forced', engine: chromium, forced: true, launch: { args: ['--font-render-hinting=none'] }, page: { forcedColors: 'active' } },
  { name: 'chromium-forced-dark', engine: chromium, forced: true, launch: { args: ['--font-render-hinting=none'] }, page: { forcedColors: 'active', colorScheme: 'dark' } },
  { name: 'firefox', engine: firefox, forced: false, launch: {}, page: {} },
  { name: 'firefox-forced', engine: firefox, forced: true, launch: { firefoxUserPrefs: { 'browser.display.document_color_use': 2 } }, page: {} },
];

// a PNG's pixels, decoded in an unforced Chromium page with no color
// conversion
const reader = await chromium.launch();
const decoder = await reader.newPage();
const decode = (png) => decoder.evaluate(async (b64) => {
  const bytes = Uint8Array.from(atob(b64), (c) => c.charCodeAt(0));
  const bmp = await createImageBitmap(new Blob([bytes], { type: 'image/png' }), { colorSpaceConversion: 'none', premultiplyAlpha: 'none' });
  const c = new OffscreenCanvas(bmp.width, bmp.height);
  const g = c.getContext('2d');
  g.drawImage(bmp, 0, 0);
  return { w: bmp.width, h: bmp.height, d: Array.from(g.getImageData(0, 0, bmp.width, bmp.height).data) };
}, png.toString('base64'));
const px = (im, x, y) => { const i = (y * im.w + x) * 4; return `rgb(${im.d[i]},${im.d[i + 1]},${im.d[i + 2]})`; };
// the pixels where two pictures differ: their count and their rows
const diff = (a, b) => {
  if (a.w !== b.w || a.h !== b.h) throw new Error('the pictures differ in size');
  let n = 0, y0 = null, y1 = null;
  for (let p = 0; p < a.w * a.h; p++) {
    const i = p * 4;
    if (a.d[i] !== b.d[i] || a.d[i + 1] !== b.d[i + 1] || a.d[i + 2] !== b.d[i + 2]) { n++; const y = Math.floor(p / a.w); y0 ??= y; y1 = y; }
  }
  return { n, y0, y1 };
};

let pass = true;
mkdirSync(new URL('./results/', import.meta.url), { recursive: true });
for (const m of MODES) {
  const browser = await m.engine.launch(m.launch);
  const page = await browser.newPage({ viewport: { width: 1200, height: 900 }, ...m.page });
  await page.setContent(html);
  await page.evaluate(() => document.fonts.ready);
  const shot = async (selector, i = 0) => decode(await page.locator(selector).nth(i).screenshot());
  const got = await page.evaluate(() => ({
    forced: matchMedia('(forced-colors: active)').matches,
    browser: navigator.userAgent.match(/(Chrome|Firefox)\/[\d.]+/)[0],
  }));
  // an option's edge: the pixel 1px in from its left side at its middle
  const [o0, o1] = [await shot('#opt .cui-option-label', 0), await shot('#opt .cui-option-label', 1)];
  got.options = [px(o0, 1, Math.floor(o0.h / 2)), px(o1, 1, Math.floor(o1.h / 2))];
  await browser.close();
  writeFileSync(new URL(`./results/${m.name}.json`, import.meta.url), JSON.stringify({ date: new Date().toISOString(), ...got }, null, 1));

  const claims = [
    [`forced colors ${m.forced ? '' : 'not '}matched (the control)`, got.forced === m.forced],
    ['the checked option drawn unlike the unchecked one (the control)', got.options[0] !== got.options[1]],
  ];
  console.log(`${m.name}: ${got.browser}`);
  for (const [claim, ok] of claims) {
    pass &&= ok;
    console.log(`  ${ok ? 'ok  ' : 'FAIL'} ${claim}`);
  }
}
await reader.close();
process.exit(pass ? 0 : 1);
