// node close.mjs: the store overlay's close button's X, apple.com's SVG read
// live off iPhone 18 Pro's buy page with its first "Show more" opened (never
// stored; .rc-overlay-close svg, drawn at 20px without a view box, so its
// units are the button's pixels: read here in a 20 by 20 view box, as it is
// drawn), rasterized at 20 pixels to the unit on a blank page (svg.js), and
// ciderui's own drawn from its numbers: the share of ink the two have in
// common; and the marketing pages' modal's X, the iPhone page's trade-in
// modal's (in a 36 by 36 SVG, its circle's), against ciderui's small X in
// the middle of the same box. Fails below 0.85 (intersection over union), a mark set
// before the run, or if the modal's examples draw other X's. The result is
// results/close.json.
import { mkdirSync, readFileSync, writeFileSync } from 'node:fs';
import { open } from '../browsers.mjs';

// ciderui's: two lines 3 thick, round at their ends, between the round
// ends' middles, (4.88, 4.88) to (15.12, 15.12) and (15.12, 4.88) to
// (4.88, 15.12) (the README's numbers)
export const OURS = '<svg viewBox="0 0 20 20"><path d="M4.88 4.88 15.12 15.12M15.12 4.88 4.88 15.12" fill="none" stroke="#000" stroke-width="3" stroke-linecap="round"/></svg>';

// the marketing pages' modal's: the same strokes, each arm 0.62 shorter at
// its ends, (5.5, 5.5) to (14.5, 14.5) and (14.5, 5.5) to (5.5, 14.5)
export const SMALL = '<svg viewBox="0 0 20 20"><path d="M5.5 5.5 14.5 14.5M14.5 5.5 5.5 14.5" fill="none" stroke="#000" stroke-width="3" stroke-linecap="round"/></svg>';

const src = readFileSync(new URL('./svg.js', import.meta.url), 'utf8');
const page = await open('chromium', 'https://www.apple.com/shop/buy-iphone/iphone-18-pro');
let result;
try {
  await page.page.waitForTimeout(2500);
  await page.page.click('a.rc-decisionsection-button');
  await page.page.waitForSelector('.rc-overlay-close svg', { state: 'visible' });
  const apple = await page.run(`Promise.resolve(document.querySelector('.rc-overlay-close svg').outerHTML)`);
  await page.page.setContent('<!doctype html><body></body>');
  result = await page.run(`(async () => {
${src}
const apple = ${JSON.stringify(apple)}.replace(/<svg([^>]*?)>/, (m, a) => '<svg' + a.replace(/ (width|height|viewBox)="[^"]*"/g, '') + ' viewBox="0 0 20 20">');
const a = await raster(apple, 20), o = await raster(${JSON.stringify(OURS)}, 20);
// the rows through the X's middle and an arm: the stroke's run across them
return { appleBox: a.box, oursBox: o.box, iou: compare(a, o), rows: [10, 6, 14].map(y => [y, a.runs('row', y), o.runs('row', y)]) };
})()`);
} finally {
  await page.close();
}
// the marketing pages' modal's X: the iPhone page's trade-in modal opened,
// its SVG (36 by 36, the X in its middle) against ciderui's drawn in the
// middle of the same 36 by 36, as its circle holds it
const marketing = await open('chromium', 'https://www.apple.com/iphone/');
try {
  await marketing.page.waitForTimeout(2500);
  const opener = marketing.page.locator('button[data-modal-open="trade-in-modal"]').first();
  await opener.scrollIntoViewIfNeeded();
  await opener.click();
  await marketing.page.waitForSelector('#trade-in-modal .ric-modal-close-icon svg', { state: 'visible' });
  const apple = await marketing.run(`Promise.resolve(document.querySelector('#trade-in-modal .ric-modal-close-icon svg').outerHTML)`);
  await marketing.page.setContent('<!doctype html><body></body>');
  result.marketing = await marketing.run(`(async () => {
${src}
const ours = ${JSON.stringify(SMALL)}.replace('viewBox="0 0 20 20">', 'viewBox="0 0 36 36"><g transform="translate(8 8)">').replace('</svg>', '</g></svg>');
const a = await raster(${JSON.stringify(apple)}, 20), o = await raster(ours, 20);
return { appleBox: a.box, oursBox: o.box, iou: compare(a, o), rows: [18, 14].map(y => [y, a.runs('row', y), o.runs('row', y)]) };
})()`);
} finally {
  await marketing.close();
}
mkdirSync(new URL('./results/', import.meta.url), { recursive: true });
writeFileSync(new URL('./results/close.json', import.meta.url), JSON.stringify({ browser: page.label, date: new Date().toISOString(), ours: OURS, small: SMALL, ...result }, null, 1));
console.log(page.label);
console.log(JSON.stringify(result, null, 1));
// the X compared is the one the kit ships: the modal's example's, in currentColor
const example = readFileSync(new URL('../../components/modal/examples/modal.html', import.meta.url), 'utf8');
const inner = OURS.replace(/^<svg[^>]*>|<\/svg>$/g, '').replaceAll('stroke="#000"', 'stroke="currentColor"');
const shipped = example.includes(inner);
console.log(`${shipped ? 'ok  ' : 'FAIL'} the X is the modal example's`);
// and the marketing X, the marketing modal's example's
const page2 = readFileSync(new URL('../../components/modal/examples/page.html', import.meta.url), 'utf8');
const shipped2 = page2.includes(SMALL.replace(/^<svg[^>]*>|<\/svg>$/g, '').replaceAll('stroke="#000"', 'stroke="currentColor"'));
console.log(`${shipped2 ? 'ok  ' : 'FAIL'} the marketing X is the marketing modal example's`);
const pass = result.iou >= 0.85 && result.marketing.iou >= 0.85 && shipped && shipped2;
console.log(pass ? 'PASS: it shares at least 0.85 of its ink with apple.com\'s' : 'FAIL');
process.exit(pass ? 0 : 1);
