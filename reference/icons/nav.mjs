// node nav.mjs: the global nav's icons, apple.com's SVG read live off its
// home page (never stored), rasterized at 20 pixels to the unit on a blank
// page (svg.js), and ciderui's own drawn from their numbers: the share of ink
// each pair has in common. The search and the bag in the bar (15 and 14 by
// 44) and on a phone (17 by 48, apple.com's compact images), the menu's
// button closed and open (its two lines' points, open, are the end of the
// animation apple.com runs on them). Fails below 0.85 (intersection over
// union), a mark set before the run. The result is results/nav.json.
import { mkdirSync, readFileSync, writeFileSync } from 'node:fs';
import { open } from '../browsers.mjs';

// ciderui's, in apple.com's view boxes: the numbers are the README's. The
// menu's open lines are its closed ones turned about their middles, as
// globalnav.css turns them (the same numbers)
const BAR = 'fill="none" stroke="#000" stroke-width="1.2" stroke-linecap="round"';
export const OURS = {
  search: '<svg viewBox="0 0 15 44"><circle cx="5.95" cy="19.95" r="4.96" fill="none" stroke="#000" stroke-width="1.08"/><path d="M9.46 23.46 13.74 27.74" fill="none" stroke="#000" stroke-width="1.52" stroke-linecap="round"/></svg>',
  bag: '<svg viewBox="0 0 14 44"><rect x="1.1" y="16.55" width="11.8" height="10.9" rx="1.5" fill="none" stroke="#000" stroke-width="1.1"/><path d="M4.55 16.55V16.05a2.45 2.45 0 0 1 4.9 0v.5" fill="none" stroke="#000" stroke-width="1.1"/></svg>',
  'search-compact': '<svg viewBox="0 0 17 48"><circle cx="6.94" cy="21.94" r="5.86" fill="none" stroke="#000" stroke-width="1.22"/><path d="M11.08 26.08 15.6 30.6" fill="none" stroke="#000" stroke-width="1.8" stroke-linecap="round"/></svg>',
  'bag-compact': '<svg viewBox="0 0 17 48"><rect x="1.55" y="17.5" width="13.9" height="12.95" rx="2" fill="none" stroke="#000" stroke-width="1.2"/><path d="M5.375 17.5V17.175a3.125 3.125 0 0 1 6.25 0v.325" fill="none" stroke="#000" stroke-width="1.2"/></svg>',
  menu: `<svg viewBox="0 0 18 18"><path d="M2 5H16" ${BAR}/><path d="M2 12H16" ${BAR}/></svg>`,
  'menu-open': `<svg viewBox="0 0 18 18"><path d="M2 5H16" ${BAR} transform="translate(9.25 9.25) rotate(45) scale(1.16168 1) translate(-9 -5)"/><path d="M2 12H16" ${BAR} transform="translate(9.25 9.25) rotate(-45) scale(1.16168 1) translate(-9 -12)"/></svg>`,
};

const src = readFileSync(new URL('./svg.js', import.meta.url), 'utf8');
const page = await open('chromium', 'https://www.apple.com/');
let result;
try {
  const apple = await page.run(`(async () => {
    const svg = (sel) => document.querySelector('#globalnav ' + sel).outerHTML;
    // the menu's lines open: each polyline's points at the end of the animation
    // that starts from its closed points
    const menu = document.querySelector('#globalnav-menutrigger-button svg').cloneNode(true);
    for (const line of menu.querySelectorAll('polyline')) {
      const opening = [...line.querySelectorAll('animate')].map(a => a.getAttribute('values').split(';').map(v => v.trim()))
        .find(v => v[0] === line.getAttribute('points'));
      line.dataset.open = opening.at(-1);
      line.querySelectorAll('animate').forEach(a => a.remove());
    }
    const closed = menu.outerHTML;
    for (const line of menu.querySelectorAll('polyline')) line.setAttribute('points', line.dataset.open);
    return {
      search: svg('.globalnav-link-search .globalnav-image-regular svg'), bag: svg('.globalnav-link-bag .globalnav-image-regular svg'),
      'search-compact': svg('.globalnav-link-search .globalnav-image-compact svg'), 'bag-compact': svg('.globalnav-link-bag .globalnav-image-compact svg'),
      menu: closed, 'menu-open': menu.outerHTML,
      openPoints: [...menu.querySelectorAll('polyline')].map(l => l.id + ': ' + l.getAttribute('points')),
    };
  })()`);
  console.log('apple.com\'s menu lines, open:', apple.openPoints.join('; '));
  await page.page.setContent('<!doctype html><body></body>');
  result = await page.run(`(async () => {
${src}
const apple = ${JSON.stringify(apple)}, ours = ${JSON.stringify(OURS)};
const out = {};
for (const k of Object.keys(ours)) {
  const a = await raster(apple[k], 20), o = await raster(ours[k], 20);
  out[k] = { appleBox: a.box, oursBox: o.box, iou: compare(a, o),
    handleRows: k === 'search' ? [24.5, 25.2, 25.9, 26.6].map(y => [y, a.runs('row', y), o.runs('row', y)]) : undefined };
}
out.openPoints = apple.openPoints;
return out;
})()`);
} finally {
  await page.close();
}
mkdirSync(new URL('./results/', import.meta.url), { recursive: true });
writeFileSync(new URL('./results/nav.json', import.meta.url), JSON.stringify({ browser: page.label, date: new Date().toISOString(), ours: OURS, ...result }, null, 1));
console.log(page.label);
console.log(JSON.stringify(result, null, 1));
const pass = Object.keys(OURS).every(k => result[k].iou >= 0.85);
console.log(pass ? 'PASS: each shares at least 0.85 of its ink with apple.com\'s' : 'FAIL: below 0.85');
process.exit(pass ? 0 : 1);
