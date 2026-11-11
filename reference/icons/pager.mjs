// @external: node pager.mjs: Apple Developer Forums' pager's chevrons, read live (never
// stored): SF Symbols' chevron.left and chevron.right at medium weight, each a path
// in a 74.561 by 70.459 view box drawn past it (overflow visible), 0.74561em wide
// at the button's 15px (../apple.com/'s forums-pager:previous.icon). ciderui's own,
// drawn from their numbers, rasterized beside them over the ink's whole box
// (0 −10 74.561 90) at 4 pixels to the unit (svg.js): the share of ink each pair
// has in common. Fails if either shares below 0.85 (intersection over union), a mark
// set before the run.
// The result is results/pager.json.
import { mkdirSync, readFileSync, writeFileSync } from 'node:fs';
import { open } from '../browsers.mjs';

// @invariant: ciderui's, in SF Symbols' units: two strokes 11.5 thick, round, meeting at
// the point 11.78 from the advance's side, 35.2 down, their ends 49.9 in and 37.5
// above and below it; the right one the left mirrored in the 74.561 advance
export const LEFT = '<svg viewBox="0 -10 74.561 90"><path d="M49.9 -2.3 11.78 35.2 49.9 72.7" fill="none" stroke="#000" stroke-width="11.5" stroke-linecap="round" stroke-linejoin="round"/></svg>';
export const RIGHT = '<svg viewBox="0 -10 74.561 90"><path d="M24.661 -2.3 62.781 35.2 24.661 72.7" fill="none" stroke="#000" stroke-width="11.5" stroke-linecap="round" stroke-linejoin="round"/></svg>';
const MARK = 0.85;

const src = readFileSync(new URL('./svg.js', import.meta.url), 'utf8');
const page = await open('chromium', 'https://developer.apple.com/forums/latest?page=3');
let result;
try {
  await page.page.waitForTimeout(3000);
  const apple = await page.run(`Promise.resolve(['previous-page', 'next-page'].map((c) => {
    const d = document.querySelector('nav.pagination a.' + c + ' sf-symbol svg path').getAttribute('d');
    return '<svg viewBox="0 -10 74.561 90"><path d="' + d + '"/></svg>';
  }))`);
  // @external: about:blank, not setContent: the forums' Content-Security-Policy stays with
  // the page's origin and refuses the blob: images svg.js draws
  await page.page.goto('about:blank');
  result = await page.run(`(async () => {
${src}
const pair = async (a, o) => { const A = await raster(a, 4), O = await raster(o, 4); return { appleBox: A.box, oursBox: O.box, iou: compare(A, O) }; };
return { left: await pair(${JSON.stringify(apple[0])}, ${JSON.stringify(LEFT)}), right: await pair(${JSON.stringify(apple[1])}, ${JSON.stringify(RIGHT)}) };
})()`);
} finally {
  await page.close();
}

mkdirSync(new URL('./results/', import.meta.url), { recursive: true });
writeFileSync(new URL('./results/pager.json', import.meta.url), JSON.stringify({ browser: page.label, date: new Date().toISOString(), ours: { left: LEFT, right: RIGHT }, ...result }, null, 1));
console.log(page.label);
console.log(JSON.stringify(result, null, 1));
for (const [name, { iou }] of Object.entries(result)) if (iou < MARK) { console.error(`${name} shares ${iou} of its ink with Apple's, under ${MARK}`); process.exit(1); }
