// node searchfield.mjs: the global nav's search flyout's icons, apple.com's
// SVG read live off its home page (never stored; the flyout's markup is in
// the page closed), rasterized at 20 pixels to the unit on a blank page
// (svg.js), and ciderui's own drawn from their numbers: the share of ink each
// pair has in common. The field's magnifier from 834px (30 by 32) and on a
// phone (38 by 40), and the quick links' arrow (9 by 16, 13 by 25). Fails
// below 0.85 (intersection over union), a mark set before the run, or if the
// search's example draws others. The result is results/searchfield.json.
import { mkdirSync, readFileSync, writeFileSync } from 'node:fs';
import { open } from '../browsers.mjs';

// ciderui's, in apple.com's view boxes: the numbers are the README's. Each
// handle starts where its round end meets the ring's inner edge
const LINE = 'fill="none" stroke="#000"';
export const OURS = {
  field: `<svg viewBox="0 0 30 32"><circle cx="13.475" cy="14.775" r="6.225" ${LINE} stroke-width="1.3"/><path d="M18.11 19.61 22.57 24.07" ${LINE} stroke-width="2.2" stroke-linecap="round"/></svg>`,
  'field-compact': `<svg viewBox="0 0 38 40"><circle cx="17.225" cy="17.9" r="7.25" ${LINE} stroke-width="1.45"/><path d="M22.6 23.6 27.74 28.74" ${LINE} stroke-width="2.62" stroke-linecap="round"/></svg>`,
  arrow: `<svg viewBox="0 0 9 16"><path d="M.75 7.75H8.25M5.25 4.75 8.25 7.75 5.25 10.75" ${LINE} stroke-width="1" stroke-linecap="round" stroke-linejoin="round"/></svg>`,
  'arrow-compact': `<svg viewBox="0 0 13 25"><path d="M1.1 13H11.93" ${LINE} stroke-width="1.3" stroke-linecap="round"/><path d="M7.48 8.55 11.93 13 7.48 17.45" ${LINE} stroke-width="1.2" stroke-linecap="round" stroke-linejoin="round"/></svg>`,
};

const src = readFileSync(new URL('./svg.js', import.meta.url), 'utf8');
const page = await open('chromium', 'https://www.apple.com/');
let result;
try {
  const apple = await page.run(`(async () => {
    const svg = (sel) => document.querySelector('#globalnav-submenu-search ' + sel).outerHTML;
    return {
      field: svg('.globalnav-searchfield-submit .globalnav-image-regular svg'),
      'field-compact': svg('.globalnav-searchfield-submit .globalnav-image-compact svg'),
      arrow: svg('.globalnav-searchresults-list-icon .globalnav-image-regular svg'),
      'arrow-compact': svg('.globalnav-searchresults-list-icon .globalnav-image-compact svg'),
    };
  })()`);
  await page.page.setContent('<!doctype html><body></body>');
  result = await page.run(`(async () => {
${src}
const apple = ${JSON.stringify(apple)}, ours = ${JSON.stringify(OURS)};
const out = {};
for (const k of Object.keys(ours)) {
  const a = await raster(apple[k], 20), o = await raster(ours[k], 20);
  // a row through the ring's middle and one through the handle (the arrow's
  // shaft and head): the strokes' runs across them
  const rows = { field: [15, 22.5], 'field-compact': [18, 27], arrow: [6, 7.75], 'arrow-compact': [10, 13] }[k];
  out[k] = { appleBox: a.box, oursBox: o.box, iou: compare(a, o), rows: rows.map(y => [y, a.runs('row', y), o.runs('row', y)]) };
}
return out;
})()`);
} finally {
  await page.close();
}
mkdirSync(new URL('./results/', import.meta.url), { recursive: true });
writeFileSync(new URL('./results/searchfield.json', import.meta.url), JSON.stringify({ browser: page.label, date: new Date().toISOString(), ours: OURS, ...result }, null, 1));
console.log(page.label);
console.log(JSON.stringify(result, null, 1));
// the icons compared are the ones the kit ships: each drawing is in the
// search's example, in currentColor
const example = readFileSync(new URL('../../components/globalnav/examples/search.html', import.meta.url), 'utf8');
const shipped = Object.entries(OURS).map(([k, svg]) => {
  const inner = svg.replace(/^<svg[^>]*>|<\/svg>$/g, '').replaceAll('stroke="#000"', 'stroke="currentColor"');
  const ok = example.includes(inner);
  console.log(`${ok ? 'ok  ' : 'FAIL'} ${k} is the example's`);
  return ok;
});
const pass = Object.keys(OURS).every(k => result[k].iou >= 0.85) && shipped.every(Boolean);
console.log(pass ? 'PASS: each shares at least 0.85 of its ink with apple.com\'s' : 'FAIL');
process.exit(pass ? 0 : 1);
