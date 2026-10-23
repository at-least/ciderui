// node paddle.mjs: a gallery's paddles' chevrons, apple.com's SVG read live
// off the iPhone page (never stored; the 36px paddles' icon-control-chevronleft
// and -right), rasterized at 20 pixels to the unit on a blank page (svg.js),
// and ciderui's own drawn from their numbers: the share of ink each pair has
// in common. Fails below 0.85 (intersection over union), a mark set before the
// run, or if the gallery's example draws others. The result is results/paddle.json.
import { mkdirSync, readFileSync, writeFileSync } from 'node:fs';
import { open } from '../browsers.mjs';

// ciderui's, in apple.com's view box: a line 3 thick, round at its ends and
// its corner, from each end's middle to the corner's (the README's numbers)
const LINE = 'fill="none" stroke="#000" stroke-width="3" stroke-linecap="round" stroke-linejoin="round"';
export const OURS = {
  previous: `<svg viewBox="0 0 36 36"><path d="M20 12.5 14.5 18 20 23.5" ${LINE}/></svg>`,
  next: `<svg viewBox="0 0 36 36"><path d="M16 12.5 21.5 18 16 23.5" ${LINE}/></svg>`,
};

const src = readFileSync(new URL('./svg.js', import.meta.url), 'utf8');
const page = await open('chromium', 'https://www.apple.com/iphone/');
let result;
try {
  const apple = await page.run(`Promise.resolve({
    previous: document.querySelector('main .paddlenav-arrow-previous svg').outerHTML,
    next: document.querySelector('main .paddlenav-arrow-next svg').outerHTML,
  })`);
  await page.page.setContent('<!doctype html><body></body>');
  result = await page.run(`(async () => {
${src}
const apple = ${JSON.stringify(apple)}, ours = ${JSON.stringify(OURS)};
const out = {};
for (const k of Object.keys(ours)) {
  const a = await raster(apple[k], 20), o = await raster(ours[k], 20);
  // the rows through each arm and the corner: the stroke's run across them
  out[k] = { appleBox: a.box, oursBox: o.box, iou: compare(a, o), rows: [14, 18, 22].map(y => [y, a.runs('row', y), o.runs('row', y)]) };
}
return out;
})()`);
} finally {
  await page.close();
}
mkdirSync(new URL('./results/', import.meta.url), { recursive: true });
writeFileSync(new URL('./results/paddle.json', import.meta.url), JSON.stringify({ browser: page.label, date: new Date().toISOString(), ours: OURS, ...result }, null, 1));
console.log(page.label);
console.log(JSON.stringify(result, null, 1));
// the chevrons compared are the ones the kit ships: each in the gallery's
// example, in currentColor
const example = readFileSync(new URL('../../components/gallery/examples/gallery.html', import.meta.url), 'utf8');
const shipped = Object.entries(OURS).map(([k, svg]) => {
  const inner = svg.replace(/^<svg[^>]*>|<\/svg>$/g, '').replaceAll('stroke="#000"', 'stroke="currentColor"');
  const ok = example.includes(inner);
  console.log(`${ok ? 'ok  ' : 'FAIL'} ${k} is the example's`);
  return ok;
});
const pass = Object.keys(OURS).every(k => result[k].iou >= 0.85) && shipped.every(Boolean);
console.log(pass ? 'PASS: each shares at least 0.85 of its ink with apple.com\'s' : 'FAIL: below 0.85');
process.exit(pass ? 0 : 1);
