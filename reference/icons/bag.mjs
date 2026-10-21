// node bag.mjs: the bag's flyout's profile icons, apple.com's SVG read live
// off its home page with the bag opened (never stored; its markup its
// script's, fetched: ../apple.com/, bag-flyout:*), a regular one from 834px
// (11 by 16) and a compact one to 833px (16 by 25) for each of its four
// links, its orders' box, its saved items' bookmark, its account's gear and
// its sign-in's person, rasterized at 20 pixels to the unit on a blank page
// (svg.js), and ciderui's own drawn from their numbers: the share of ink
// each pair has in common. Fails below 0.85 (intersection over union), a
// mark set before the run, or if the bag's example draws others. The result
// is results/bag.json.
import { mkdirSync, readFileSync, writeFileSync } from 'node:fs';
import { open } from '../browsers.mjs';

// ciderui's, in apple.com's view boxes: the numbers are the README's
const LINE = 'fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round"';
export const OURS = {
  'orders-regular': `<svg viewBox="0 0 11 16"><path d="M6.033 2.843 9.198 4.628A1.0625 1.0625 0 0 1 9.738 5.553V9.063A1.0625 1.0625 0 0 1 9.198 9.989L6.033 11.774A1.0625 1.0625 0 0 1 4.989 11.774L1.825 9.989A1.0625 1.0625 0 0 1 1.285 9.063V5.553A1.0625 1.0625 0 0 1 1.825 4.628L4.989 2.843A1.0625 1.0625 0 0 1 6.033 2.843ZM1.43 5.018 5.511 7.375 9.593 5.018M5.511 7.375V11.91" ${LINE} stroke-width=".875"/><path d="M3.849 3.988 7.254 5.878" fill="none" stroke="#000" stroke-width=".875"/></svg>`,
  'orders-compact': `<svg viewBox="0 0 16 25"><path d="M8.835 6.211 13.135 8.637A1.7 1.7 0 0 1 14 10.118V14.882A1.7 1.7 0 0 1 13.135 16.363L8.835 18.788A1.7 1.7 0 0 1 7.165 18.788L2.865 16.363A1.7 1.7 0 0 1 2 14.882V10.118A1.7 1.7 0 0 1 2.865 8.637L7.165 6.211A1.7 1.7 0 0 1 8.835 6.211ZM2.232 9.26 8 12.594 13.768 9.26M8 12.594V19.007" ${LINE} stroke-width="1.2"/><path d="M5.561 7.806 10.456 10.522" fill="none" stroke="#000" stroke-width="1.2"/></svg>`,
  'yoursaves-regular': `<svg viewBox="0 0 11 16"><path d="M2.719 11.97V3.972A1.184 1.184 0 0 1 3.903 2.789H7.097A1.184 1.184 0 0 1 8.281 3.972V11.97L5.5 9.3Z" ${LINE} stroke-width=".82"/></svg>`,
  'yoursaves-compact': `<svg viewBox="0 0 16 25"><path d="M4 19.25V7.45A1.7 1.7 0 0 1 5.7 5.75H10.3A1.7 1.7 0 0 1 12 7.45V19.25L8 15.75Z" ${LINE} stroke-width="1.2"/></svg>`,
  'account-regular': `<svg viewBox="0 0 11 16"><circle cx="5.5" cy="7.25" r="4.05" fill="none" stroke="#000" stroke-width=".79"/><circle cx="5.5" cy="7.25" r=".9" fill="none" stroke="#000" stroke-width=".65"/><path d="M6.7 7.25H9.2M4.9 8.289 3.65 10.454M4.9 6.211 3.65 4.046" ${LINE} stroke-width=".69"/><path d="M9.9 7.25H10.32M9.635 8.755 10.029 8.899M8.871 10.078 9.192 10.348M7.7 11.061 7.91 11.424M6.264 11.583 6.337 11.997M4.736 11.583 4.663 11.997M3.3 11.061 3.09 11.424M2.129 10.078 1.808 10.348M1.365 8.755 .971 8.899M1.1 7.25H.68M1.365 5.745 .971 5.601M2.129 4.422 1.808 4.152M3.3 3.439 3.09 3.076M4.736 2.917 4.663 2.503M6.264 2.917 6.337 2.503M7.7 3.439 7.91 3.076M8.871 4.422 9.192 4.152M9.635 5.745 10.029 5.601" ${LINE} stroke-width=".68"/></svg>`,
  'account-compact': `<svg viewBox="0 0 16 25"><circle cx="8" cy="12.5" r="6.1" fill="none" stroke="#000" stroke-width="1.28"/><circle cx="8" cy="12.5" r="1.42" fill="none" stroke="#000" stroke-width="1.25"/><path d="M10 12.5H13.5M7 14.232 5.25 17.263M7 10.768 5.25 7.737" ${LINE} stroke-width="1.21"/><path d="M14.7 12.5H15.15M14.296 14.792 14.719 14.945M13.132 16.807 13.477 17.096M11.35 18.302 11.575 18.692M9.163 19.098 9.242 19.541M6.837 19.098 6.758 19.541M4.65 18.302 4.425 18.692M2.868 16.807 2.523 17.096M1.704 14.792 1.281 14.945M1.3 12.5H.85M1.704 10.208 1.281 10.055M2.868 8.193 2.523 7.904M4.65 6.698 4.425 6.308M6.837 5.902 6.758 5.459M9.163 5.902 9.242 5.459M11.35 6.698 11.575 6.308M13.132 8.193 13.477 7.904M14.296 10.208 14.719 10.055" ${LINE} stroke-width="1.05"/></svg>`,
  'signin-regular': `<svg viewBox="0 0 11 16"><circle cx="5.5" cy="7.633" r="4.287" fill="none" stroke="#000" stroke-width=".887"/><ellipse cx="5.5" cy="6.71" rx="1.605" ry="1.69" fill="#000"/><path d="M2.62 10.163A3.898 3.898 0 0 1 8.38 10.163 3.844 3.844 0 0 1 2.62 10.163Z" fill="#000"/></svg>`,
  'signin-compact': `<svg viewBox="0 0 16 25"><circle cx="7.99" cy="12.5" r="6.495" fill="none" stroke="#000" stroke-width="1.21"/><ellipse cx="7.99" cy="11" rx="2.475" ry="2.6" fill="#000"/><path d="M3.529 16.347A6.77 6.77 0 0 1 12.452 16.347 5.89 5.89 0 0 1 3.529 16.347Z" fill="#000"/></svg>`,
};

const src = readFileSync(new URL('./svg.js', import.meta.url), 'utf8');
const page = await open('chromium', 'https://www.apple.com/');
let result;
try {
  await page.page.waitForTimeout(2500);
  await page.page.click('#globalnav-menubutton-link-bag');
  await page.page.waitForSelector('#globalnav-submenu-bag .ac-gn-bagview-nav-link-signin', { state: 'attached', timeout: 15000 });
  // each link's two SVGs; their fill (6E6E73, without its #) is no color,
  // so each is drawn black here, as ours
  const apple = await page.run(`Promise.resolve(Object.fromEntries([...document.querySelectorAll('#globalnav-submenu-bag .ac-gn-bagview-nav-link')].flatMap((a) =>
    ['regular', 'compact'].map((k) => [a.className.match(/nav-link-(\\w+)/)[1] + '-' + k, a.querySelector('.ac-gn-bagview-nav-image-' + k + ' svg').outerHTML.replace(/fill="6E6E73"/g, 'fill="#000"')]))))`);
  await page.page.setContent('<!doctype html><body></body>');
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
writeFileSync(new URL('./results/bag.json', import.meta.url), JSON.stringify({ browser: page.label, date: new Date().toISOString(), ours: OURS, ...result }, null, 1));
console.log(page.label);
for (const [k, v] of Object.entries(result)) console.log(`${k.padEnd(18)} ${v.iou}  apple.com's ink ${JSON.stringify(v.appleBox)}, ciderui's ${JSON.stringify(v.oursBox)}`);
// the icons compared are the ones the kit ships: each drawing is in the
// bag's example, in currentColor
const example = readFileSync(new URL('../../components/globalnav/examples/bag.html', import.meta.url), 'utf8');
const shipped = Object.entries(OURS).map(([k, svg]) => {
  const inner = svg.replace(/^<svg[^>]*>|<\/svg>$/g, '').replaceAll('stroke="#000"', 'stroke="currentColor"').replaceAll('fill="#000"', 'fill="currentColor"');
  const ok = example.includes(inner);
  console.log(`${ok ? 'ok  ' : 'FAIL'} ${k} is the example's`);
  return ok;
});
const pass = Object.values(result).every((v) => v.iou >= 0.85) && shipped.every(Boolean);
console.log(pass ? 'PASS: each shares at least 0.85 of its ink with apple.com\'s' : 'FAIL: below 0.85');
process.exit(pass ? 0 : 1);
