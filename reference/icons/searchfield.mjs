// node searchfield.mjs: the global nav's search flyout's icons, apple.com's
// SVG read live off its home page (never stored; the flyout's markup is in
// the page closed), rasterized at 20 pixels to the unit on a blank page
// (svg.js), and ciderui's own drawn from their numbers: the share of ink each
// pair has in common. The field's magnifier from 834px (30 by 32) and on a
// phone (38 by 40), the quick links' arrow (9 by 16, 13 by 25), the clear
// button's mark, a disc with an X cut out of it (14 by 14, 16 by 16), and a
// suggested search's magnifier (9 by 16, 13 by 25: not in the page's markup,
// which writes it only as one types, but in its settings, the script's
// suggestedSearches images, read off the page's own script). Fails below
// 0.85 (intersection over union), a mark set before the run, or if the
// search's examples draw others. The result is results/searchfield.json.
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
  // a disc and, cut out of it (evenodd), the outline of an X: two strokes
  // 1.1 thick (1.34 compact) between round ends whose middles are 2.409
  // (2.81) from the disc's middle along each diagonal
  clear: '<svg viewBox="0 0 14 14"><path d="M7 .0339A6.9661 6.9661 0 1 0 7 13.9661 6.9661 6.9661 0 1 0 7 .0339ZM7.7778 7 9.798 9.0201A.55.55 0 0 1 9.0201 9.798L7 7.7778 4.9799 9.798A.55.55 0 0 1 4.2021 9.0201L6.2222 7 4.2021 4.9799A.55.55 0 0 1 4.9799 4.2021L7 6.2222 9.0201 4.2021A.55.55 0 0 1 9.798 4.9799Z" fill="#000" fill-rule="evenodd"/></svg>',
  'clear-compact': '<svg viewBox="0 0 16 16"><path d="M8 0A8 8 0 1 0 8 16 8 8 0 1 0 8 0ZM8.9475 8 11.2838 10.3362A.67.67 0 0 1 10.3362 11.2838L8 8.9475 5.6638 11.2838A.67.67 0 0 1 4.7162 10.3362L7.0525 8 4.7162 5.6638A.67.67 0 0 1 5.6638 4.7162L8 7.0525 10.3362 4.7162A.67.67 0 0 1 11.2838 5.6638Z" fill="#000" fill-rule="evenodd"/></svg>',
  // the ring's handle at 45° from just inside its outer edge, where its round
  // end stays out of the ring's hole
  suggestion: `<svg viewBox="0 0 9 16"><circle cx="3.675" cy="6.9" r="3.125" ${LINE} stroke-width="0.7"/><path d="M6 9.381 8.2 11.581" ${LINE} stroke-width="1.1" stroke-linecap="round"/></svg>`,
  'suggestion-compact': `<svg viewBox="0 0 13 25"><circle cx="5.425" cy="11.725" r="4.425" ${LINE} stroke-width="0.9"/><path d="M8.692 15.242 11.8 18.35" ${LINE} stroke-width="1.6" stroke-linecap="round"/></svg>`,
};

const src = readFileSync(new URL('./svg.js', import.meta.url), 'utf8');
const page = await open('chromium', 'https://www.apple.com/');
let result;
try {
  const apple = await page.run(`(async () => {
    const svg = (sel) => document.querySelector('#globalnav-submenu-search ' + sel).outerHTML;
    // the settings' suggestedSearches images, in a script's string (its
    // quotes escaped): the first two SVGs after its name
    const settings = [...document.scripts].map(s => s.textContent).find(t => t.includes('suggestedSearches'));
    const at = settings.indexOf('suggestedSearches');
    const [suggestion, suggestionCompact] = [...settings.slice(at, at + 4000).replaceAll('\\\\"', '"').matchAll(/<svg[^]*?<\\/svg>/g)].map(m => m[0]);
    return {
      suggestion, 'suggestion-compact': suggestionCompact,
      field: svg('.globalnav-searchfield-submit .globalnav-image-regular svg'),
      'field-compact': svg('.globalnav-searchfield-submit .globalnav-image-compact svg'),
      arrow: svg('.globalnav-searchresults-list-icon .globalnav-image-regular svg'),
      'arrow-compact': svg('.globalnav-searchresults-list-icon .globalnav-image-compact svg'),
      clear: svg('.globalnav-searchfield-reset .globalnav-image-regular svg'),
      'clear-compact': svg('.globalnav-searchfield-reset .globalnav-image-compact svg'),
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
  // shaft and head; the clear mark's middle and an arm)
  const rows = { field: [15, 22.5], 'field-compact': [18, 27], arrow: [6, 7.75], 'arrow-compact': [10, 13], clear: [7, 9], 'clear-compact': [8, 10.5], suggestion: [6.9, 11], 'suggestion-compact': [11.725, 17] }[k];
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
// search's examples (its suggestions' magnifier in the suggestions'), in
// currentColor
const example = ['search', 'suggestions'].map(n => readFileSync(new URL(`../../components/globalnav/examples/${n}.html`, import.meta.url), 'utf8')).join('\n');
const shipped = Object.entries(OURS).map(([k, svg]) => {
  const inner = svg.replace(/^<svg[^>]*>|<\/svg>$/g, '').replaceAll('stroke="#000"', 'stroke="currentColor"').replaceAll('fill="#000"', 'fill="currentColor"');
  const ok = example.includes(inner);
  console.log(`${ok ? 'ok  ' : 'FAIL'} ${k} is the example's`);
  return ok;
});
const pass = Object.keys(OURS).every(k => result[k].iou >= 0.85) && shipped.every(Boolean);
console.log(pass ? 'PASS: each shares at least 0.85 of its ink with apple.com\'s' : 'FAIL');
process.exit(pass ? 0 : 1);
