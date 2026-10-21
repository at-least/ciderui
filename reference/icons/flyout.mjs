// node flyout.mjs: the global nav's flyouts' chevrons, apple.com's SVG read
// live off its home page (never stored; the flyouts' markup is in the page
// closed), rasterized at 20 pixels to the unit on a blank page (svg.js), and
// ciderui's own drawn from their numbers: the share of ink each pair has in
// common. The chevron of the button beside a link from 834px (10 by 10, down,
// shown to the keyboard), the same button's on a phone (9 by 48, a row's
// chevron pointing on), the phone's back button's (9 by 48, back), and the
// chevron a phone's row shows hovered (9 by 48: apple.com's drawn pointing
// back and turned by its stylesheet at a phone's width, `scaleX(-1)`, so
// turned here the same).
// Fails below 0.85 (intersection over union), a mark set before the run, or
// if the flyout's example draws others. The result is results/flyout.json.
import { mkdirSync, readFileSync, writeFileSync } from 'node:fs';
import { open } from '../browsers.mjs';

// ciderui's, in apple.com's view boxes: the numbers are the README's
const LINE = 'fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round"';
export const OURS = {
  flyout: `<svg viewBox="0 0 10 10"><path d="M2.33 3.66 5 6.34 7.67 3.66" ${LINE} stroke-width=".88"/></svg>`,
  'flyout-compact': `<svg viewBox="0 0 9 48"><path d="M1.3 17.21 8.3 24.07 1.3 30.79" ${LINE} stroke-width="1.2"/></svg>`,
  back: `<svg viewBox="0 0 9 48"><path d="M7.7 17.21 .7 24.07 7.7 30.79" ${LINE} stroke-width="1.2"/></svg>`,
};
// what each of apple.com's is held to: the row's chevron is the compact one
const PAIRS = { flyout: 'flyout', 'flyout-compact': 'flyout-compact', back: 'back', row: 'flyout-compact' };

const src = readFileSync(new URL('./svg.js', import.meta.url), 'utf8');
const page = await open('chromium', 'https://www.apple.com/');
let result;
try {
  // a phone's width, where the row's chevron is drawn
  await page.size(390, 844);
  const apple = await page.run(`(async () => {
    const svg = (sel) => document.querySelector('#globalnav ' + sel).outerHTML;
    // the row's chevron as its stylesheet turns it
    const row = document.querySelector('#globalnav .globalnav-link-mac .globalnav-link-chevron svg').cloneNode(true);
    row.innerHTML = '<g transform="translate(9 0) scale(-1 1)">' + row.innerHTML + '</g>';
    return {
      flyout: svg('#globalnav-menubutton-link-mac .globalnav-image-regular svg'),
      'flyout-compact': svg('#globalnav-menubutton-link-mac .globalnav-image-compact svg'),
      back: svg('.globalnav-menuback-button svg'),
      row: row.outerHTML,
      rowTransform: getComputedStyle(document.querySelector('#globalnav .globalnav-link-mac .globalnav-link-chevron svg')).transform,
    };
  })()`);
  if (apple.rowTransform !== 'matrix(-1, 0, 0, 1, 0, 0)') throw new Error(`the row's chevron is turned ${apple.rowTransform}, not scaleX(-1)`);
  await page.page.setContent('<!doctype html><body></body>');
  result = await page.run(`(async () => {
${src}
const apple = ${JSON.stringify(apple)}, ours = ${JSON.stringify(OURS)}, pairs = ${JSON.stringify(PAIRS)};
const out = {};
for (const [k, mine] of Object.entries(pairs)) {
  const a = await raster(apple[k], 20), o = await raster(ours[mine], 20);
  out[k] = { ours: mine, appleBox: a.box, oursBox: o.box, iou: compare(a, o),
    // the runs across each leg's middle and the tip's row
    rows: k === 'flyout' ? [3.66, 5, 6.34].map(y => [y, a.runs('row', y), o.runs('row', y)]) : [20.6, 24.07, 27.4].map(y => [y, a.runs('row', y), o.runs('row', y)]) };
}
return out;
})()`);
} finally {
  await page.close();
}
mkdirSync(new URL('./results/', import.meta.url), { recursive: true });
writeFileSync(new URL('./results/flyout.json', import.meta.url), JSON.stringify({ browser: page.label, date: new Date().toISOString(), ours: OURS, ...result }, null, 1));
console.log(page.label);
console.log(JSON.stringify(result, null, 1));
// the icons compared are the ones the kit ships: each drawing is in the
// flyout's example, in currentColor
const example = readFileSync(new URL('../../components/globalnav/examples/flyout.html', import.meta.url), 'utf8');
const shipped = Object.entries(OURS).map(([k, svg]) => {
  const inner = svg.replace(/^<svg[^>]*>|<\/svg>$/g, '').replaceAll('stroke="#000"', 'stroke="currentColor"');
  const ok = example.includes(inner);
  console.log(`${ok ? 'ok  ' : 'FAIL'} ${k} is the example's`);
  return ok;
});
const pass = Object.keys(PAIRS).every(k => result[k].iou >= 0.85) && shipped.every(Boolean);
console.log(pass ? 'PASS: each shares at least 0.85 of its ink with apple.com\'s' : 'FAIL: below 0.85');
process.exit(pass ? 0 : 1);
