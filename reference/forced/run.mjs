// node run.mjs: does each of the kit's parts that draws a state or a color in
// a color of its own still show it in forced colors (a high-contrast theme)?
// page.html (`node page.mjs` writes it) in Playwright's Chromium, its forced
// colors emulated (forcedColors: active, as its DevTools protocol emulates
// them: black on white, and with a dark color scheme white on black) and
// not, and Playwright's Firefox, its own override of a page's
// colors on (Settings > Website contrast > Override colors, the pref
// browser.display.document_color_use 2) and not. Each part's picture
// (Playwright's of the element) is read in another page, unforced:
//   the checked swatch's ring drawn unlike each unchecked one, the
//     swatches' colors set by an element's style and by <img>;
//   the swatches' own colors drawn, three apart;
//   the gallery's dots drawn, the current one's pill more ink than a dot;
//   the empty select's own text ("Select") not drawn where its label lies:
//     not drawn, or drawn below the label's box;
// and the controls, which show the picture can tell a state: forced colors
// match as the mode says, and the checked option and checkbox draw unlike
// the unchecked ones. Not in Safari: WebKit has no forced colors. The result
// is results/<mode>.json, each mode's; exits 1 if a claim does not hold.
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
// pixels unlike the picture's top-left one (its ground)
const ink = (im) => { let n = 0; for (let i = 0; i < im.d.length; i += 4) if (im.d[i] !== im.d[0] || im.d[i + 1] !== im.d[1] || im.d[i + 2] !== im.d[2]) n++; return n; };
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
    checked: [...document.querySelectorAll('#sw .cui-swatch-input')].map((i) => i.checked),
  }));
  // a ring's edge: the four pixels in from its left side at its middle; a
  // color's, the circle's middle
  const edge = (im) => [0, 1, 2, 3].map((x) => px(im, x, 18)).join(' ');
  got.rings = [];
  for (let i = 0; i < 3; i++) { const im = await shot('#sw .cui-swatch-ring', i); got.rings.push({ checked: got.checked[i], edge: edge(im), color: px(im, 18, 18) }); }
  got.imgRings = [];
  for (let i = 0; i < 3; i++) { const im = await shot('#swi .cui-swatch-ring', i); got.imgRings.push({ checked: got.checked[i], edge: edge(im), color: px(im, 18, 18) }); }
  // an option's edge: the pixel 1px in from its left side at its middle
  const [o0, o1] = [await shot('#opt .cui-option-label', 0), await shot('#opt .cui-option-label', 1)];
  got.options = [px(o0, 1, Math.floor(o0.h / 2)), px(o1, 1, Math.floor(o1.h / 2))];
  const [c0, c1] = [await shot('#cb .cui-checkbox-indicator', 0), await shot('#cb .cui-checkbox-indicator', 1)];
  got.checkboxesDiffer = diff(c0, c1).n;
  got.dotsInk = ink(await shot('.cui-gallery-dots'));
  got.dots = [];
  for (let i = 0; i < 4; i++) got.dots.push({ current: await page.locator('.cui-gallery-dot').nth(i).getAttribute('aria-current') === 'true', ink: ink(await shot('.cui-gallery-dot', i)) });
  const [f1, f2, f3, f4] = [await shot('#f1 .cui-field'), await shot('#f2 .cui-field'), await shot('#f3 .cui-field'), await shot('#f4 .cui-field')];
  got.select = diff(f3, f4);
  got.selectShown = diff(f1, f2).n;
  got.label = await page.evaluate(() => { const f = document.querySelector('#f1 .cui-field').getBoundingClientRect(), l = document.querySelector('#f1 .cui-field-label').getBoundingClientRect(); return { y0: Math.round(l.top - f.top), y1: Math.round(l.bottom - f.top) - 1 }; });
  await browser.close();
  writeFileSync(new URL(`./results/${m.name}.json`, import.meta.url), JSON.stringify({ date: new Date().toISOString(), ...got }, null, 1));

  const on = got.rings.find((r) => r.checked), off = got.rings.filter((r) => !r.checked);
  const imgOn = got.imgRings.find((r) => r.checked), imgOff = got.imgRings.filter((r) => !r.checked);
  const current = got.dots.find((d) => d.current), rest = got.dots.filter((d) => !d.current);
  const claims = [
    [`forced colors ${m.forced ? '' : 'not '}matched (the control)`, got.forced === m.forced],
    ['the checked option and checkbox drawn unlike the unchecked ones (the control)', got.options[0] !== got.options[1] && got.checkboxesDiffer > 0],
    ['the checked swatch\'s ring drawn unlike each unchecked one, its colors an element\'s style and <img>', off.every((r) => r.edge !== on.edge) && imgOff.every((r) => r.edge !== imgOn.edge)],
    ['the swatches\' own colors drawn, three apart', new Set(got.rings.map((r) => r.color)).size === 3],
    ['the gallery\'s dots drawn, the current one\'s pill more ink than a dot', got.dotsInk > 0 && rest.every((d) => d.ink > 0 && current.ink > d.ink)],
    ['the empty select\'s own text not drawn where its label lies', got.select.n === 0 || got.select.y0 > got.label.y1],
  ];
  console.log(`${m.name}: ${got.browser}`);
  for (const [claim, ok] of claims) {
    pass &&= ok;
    console.log(`  ${ok ? 'ok  ' : 'FAIL'} ${claim}`);
  }
}
await reader.close();
process.exit(pass ? 0 : 1);
