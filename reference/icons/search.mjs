// @external: node search.mjs: apple.com's search page's two field icons, read live (never
// stored), each an SVG its button's background draws (search-serp:icon.left.button,
// :icon.right): the magnifier before the text and the reset after it, both 25 by 25.
// ciderui's own, drawn from their numbers, rasterized beside them at 20 pixels to the
// unit (svg.js): the share of ink each pair has in common, and the magnifier's ring
// read along its middle row. Fails if either shares below 0.85 (intersection over
// union), a mark set before the run.
// The result is results/search.json.
import { mkdirSync, readFileSync, writeFileSync } from 'node:fs';
import { open } from '../browsers.mjs';

// @invariant: ciderui's, in apple.com's 25 by 25: the magnifier a ring about (10.69, 10.87)
// of radius 6.28, 1.1 thick, its handle from the ring to (20.15, 20.05), 1.7 thick,
// round; the reset a disc of radius 9.5 about the middle, an X cut from it to 3.5 each
// way of the middle, 1.1 thick, round
export const MAGNIFIER = '<svg viewBox="0 0 25 25"><circle cx="10.69" cy="10.87" r="6.28" fill="none" stroke="#000" stroke-width="1.1"/><path d="M15.25 15.35 20.15 20.05" fill="none" stroke="#000" stroke-width="1.7" stroke-linecap="round"/></svg>';
export const RESET = '<svg viewBox="0 0 25 25"><path d="M12.5 3A9.5 9.5 0 1 0 12.5 22 9.5 9.5 0 1 0 12.5 3ZM13.2778 12.5 16.3889 15.6111A.55.55 0 0 1 15.6111 16.3889L12.5 13.2778 9.3889 16.3889A.55.55 0 0 1 8.6111 15.6111L11.7222 12.5 8.6111 9.3889A.55.55 0 0 1 9.3889 8.6111L12.5 11.7222 15.6111 8.6111A.55.55 0 0 1 16.3889 9.3889Z" fill="#000" fill-rule="evenodd"/></svg>';
const MARK = 0.85;

const src = readFileSync(new URL('./svg.js', import.meta.url), 'utf8');
const page = await open('chromium', 'https://www.apple.com/us/search/airpods?src=globalnav');
let result;
try {
  await page.page.waitForTimeout(2500);
  const apple = await page.run(`Promise.resolve(['.form-icons-search15', '.form-icons-reset'].map((s) => {
    const u = getComputedStyle(document.querySelector('.rf-serp-searchbox ' + s)).backgroundImage;
    return decodeURIComponent(u.slice(u.indexOf(',') + 1, -2));
  }))`);
  await page.page.setContent('<!doctype html><body></body>');
  result = await page.run(`(async () => {
${src}
const pair = async (a, o) => { const A = await raster(a, 20), O = await raster(o, 20); return { appleBox: A.box, oursBox: O.box, iou: compare(A, O), rows: [10.87].map(y => [y, A.runs('row', y), O.runs('row', y)]) }; };
return { magnifier: await pair(${JSON.stringify(apple[0])}, ${JSON.stringify(MAGNIFIER)}), reset: await pair(${JSON.stringify(apple[1])}, ${JSON.stringify(RESET)}) };
})()`);
} finally {
  await page.close();
}

mkdirSync(new URL('./results/', import.meta.url), { recursive: true });
writeFileSync(new URL('./results/search.json', import.meta.url), JSON.stringify({ browser: page.label, date: new Date().toISOString(), ours: { magnifier: MAGNIFIER, reset: RESET }, ...result }, null, 1));
console.log(page.label);
console.log(JSON.stringify(result, null, 1));
for (const [name, { iou }] of Object.entries(result)) if (iou < MARK) { console.error(`${name} shares ${iou} of its ink with apple.com's, under ${MARK}`); process.exit(1); }
