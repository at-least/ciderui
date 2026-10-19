// node lines.mjs: the chevrons ciderui draws as apple.com's own lines, read
// live (never stored): the bag's accordion's (its polyline, 17 by 8.85, the
// bag's compact items') and the footer's (11 by 6, a phone's footer
// section's), each rasterized at 20 pixels to the unit (svg.js) against
// ciderui's: the share of ink in common. Fails below 0.85 (intersection over
// union), a mark set before the run, and if a drawing is not its
// component's example's. The result is results/lines.json.
import { mkdirSync, readFileSync, writeFileSync } from 'node:fs';
import { open } from '../browsers.mjs';

const LINE = (w) => `fill="none" stroke="#000" stroke-width="${w}" stroke-linecap="round" stroke-linejoin="round"`;
export const OURS = {
  disclosure: `<svg viewBox="0 0 17 8.85"><path d="M15 1.13 8.5 7.72 2 1.13" ${LINE(2.25)}/></svg>`,
  footer: `<svg viewBox="0 0 11 6"><path d="M10.075.675 5.5 5.323.925.675" ${LINE(1)}/></svg>`,
};
const EXAMPLES = { disclosure: '../../components/accordion/examples/accordion.html', footer: '../../components/footer/examples/footer.html' };
// apple.com's: the polyline with its stroke as its page computes it
const read = (selector) => `(() => {
  const svg = document.querySelector(${JSON.stringify(selector)}).cloneNode(true);
  for (const l of svg.querySelectorAll('polyline')) {
    l.querySelectorAll('animate').forEach(a => a.remove());
    l.setAttribute('stroke-width', getComputedStyle(document.querySelector(${JSON.stringify(selector)} + ' polyline')).strokeWidth);
  }
  return svg.outerHTML;
})()`;
const src = readFileSync(new URL('./svg.js', import.meta.url), 'utf8');
const apple = {};
for (const [key, url, selector, width] of [
  ['disclosure', 'https://www.apple.com/shop/bag', '.rc-accordion-item .accordion-icon-svg', 1200],
  ['footer', 'https://www.apple.com/', '#ac-globalfooter .footer-icon-svg', 390],
]) {
  const page = await open('chromium', url);
  try {
    await page.page.setViewportSize({ width, height: 900 });
    await page.page.waitForTimeout(2500);
    apple[key] = await page.run(`Promise.resolve(${read(selector)})`);
  } finally {
    await page.close();
  }
}
const page = await open('chromium', 'about:blank');
let result;
try {
  result = await page.run(`(async () => {
${src}
const apple = ${JSON.stringify(apple)}, ours = ${JSON.stringify(OURS)};
const out = {};
for (const k of Object.keys(ours)) {
  const a = await raster(apple[k], 20), o = await raster(ours[k], 20);
  out[k] = { appleBox: a.box, oursBox: o.box, iou: compare(a, o) };
}
return out;
})()`);
} finally {
  await page.close();
}
mkdirSync(new URL('./results/', import.meta.url), { recursive: true });
writeFileSync(new URL('./results/lines.json', import.meta.url), JSON.stringify({ browser: page.label, date: new Date().toISOString(), ours: OURS, ...result }, null, 1));
console.log(page.label);
console.log(JSON.stringify(result, null, 1));
const shipped = Object.entries(OURS).map(([k, svg]) => {
  const inner = svg.replace(/^<svg[^>]*>|<\/svg>$/g, '').replaceAll('stroke="#000"', 'stroke="currentColor"');
  const ok = readFileSync(new URL(EXAMPLES[k], import.meta.url), 'utf8').includes(inner);
  console.log(`${ok ? 'ok  ' : 'FAIL'} ${k} is its component's example's`);
  return ok;
});
const pass = Object.keys(OURS).every(k => result[k].iou >= 0.85) && shipped.every(Boolean);
console.log(pass ? 'PASS: each shares at least 0.85 of its ink with apple.com\'s' : 'FAIL: below 0.85');
process.exit(pass ? 0 : 1);
