// node localnav.mjs: the local nav's chevron, apple.com's SVG read live off
// MacBook Air's page (never stored; its polyline, and the same with the
// points its animation ends on when the menu opens), rasterized at 20 pixels
// to the unit on a blank page (svg.js), against ciderui's own (icon::CHEVRON,
// in the local nav's example) and the same turned upside down about its
// middle, as localnav.css turns it. Fails below 0.85 (intersection over
// union), or if the local nav's example draws another. The result is results/localnav.json.
import { mkdirSync, readFileSync, writeFileSync } from 'node:fs';
import { open } from '../browsers.mjs';

const LINE = 'fill="none" stroke="#000" stroke-width="1.2" stroke-linecap="round" stroke-linejoin="round"';
export const OURS = {
  chevron: `<svg viewBox="0 0 16 9"><path d="M15.265.835 8 8.167.735.835" ${LINE}/></svg>`,
  'chevron-open': `<svg viewBox="0 0 16 9"><path d="M15.265.835 8 8.167.735.835" ${LINE} transform="translate(0 9) scale(1 -1)"/></svg>`,
};

const src = readFileSync(new URL('./svg.js', import.meta.url), 'utf8');
const page = await open('chromium', 'https://www.apple.com/macbook-air/');
let result;
try {
  const apple = await page.run(`(async () => {
    const svg = document.querySelector('#ac-localnav .ac-ln-menucta-chevron svg').cloneNode(true);
    const line = svg.querySelector('polyline');
    const values = (a) => a.getAttribute('values').split(';').map(v => v.trim());
    const opening = [...line.querySelectorAll('animate')].find(a => values(a)[0] === line.getAttribute('points'));
    const open = values(opening).at(-1), dur = opening.getAttribute('dur');
    line.querySelectorAll('animate').forEach(a => a.remove());
    line.setAttribute('stroke-width', '1.2');
    line.removeAttribute('vector-effect');
    const closed = svg.outerHTML;
    line.setAttribute('points', open);
    return { chevron: closed, 'chevron-open': svg.outerHTML, points: line.getAttribute('points'), open, dur };
  })()`);
  console.log(`apple.com's chevron opens from ${apple.points ? '' : ''}to ${apple.open} in ${apple.dur}`);
  await page.page.setContent('<!doctype html><body></body>');
  result = await page.run(`(async () => {
${src}
const apple = ${JSON.stringify(apple)}, ours = ${JSON.stringify(OURS)};
const out = {};
for (const k of Object.keys(ours)) {
  const a = await raster(apple[k], 20), o = await raster(ours[k], 20);
  out[k] = { appleBox: a.box, oursBox: o.box, iou: compare(a, o) };
}
out.open = apple.open; out.dur = apple.dur;
return out;
})()`);
} finally {
  await page.close();
}
mkdirSync(new URL('./results/', import.meta.url), { recursive: true });
writeFileSync(new URL('./results/localnav.json', import.meta.url), JSON.stringify({ browser: page.label, date: new Date().toISOString(), ours: OURS, ...result }, null, 1));
console.log(page.label);
for (const k of Object.keys(OURS)) console.log(`${k}: iou ${result[k].iou}, apple.com's ink ${JSON.stringify(result[k].appleBox)}, ciderui's ${JSON.stringify(result[k].oursBox)}`);
const example = readFileSync(new URL('../../components/localnav/examples/localnav.html', import.meta.url), 'utf8');
const inner = OURS.chevron.replace(/^<svg[^>]*>|<\/svg>$/g, '').replace('stroke="#000"', 'stroke="currentColor"');
const shipped = example.includes(inner);
console.log(`${shipped ? 'ok  ' : 'FAIL'} the chevron is the example's`);
const pass = Object.keys(OURS).every(k => result[k].iou >= 0.85) && shipped;
console.log(pass ? 'PASS: each shares at least 0.85 of its ink with apple.com\'s' : 'FAIL');
process.exit(pass ? 0 : 1);
