// node docc.mjs: Apple's developer documentation's icons the docs set draws
// (the bar's button that opens the navigator, a navigator row's chevron, an
// article's and a collection's glyph), DocC's SVG read live off the crash
// report's article (never stored), each rasterized at 40 pixels to the unit
// on a blank page (svg.js) beside the drawing the kit ships, read out of
// rust/src/icon.rs (DOCS_SIDEBAR, DOCS_CHEVRON, DOCS_ARTICLE,
// DOCS_COLLECTION): the share of ink each pair has in common, and each one's
// ink and its runs along rows and columns, the numbers the drawings are
// drawn from. Fails below 0.85 (intersection over union), the mark the
// global nav's icons are held to. The result is results/docc.json.
import { mkdirSync, readFileSync, writeFileSync } from 'node:fs';
import { open } from '../browsers.mjs';

const icons = readFileSync(new URL('../../rust/src/icon.rs', import.meta.url), 'utf8');
const shipped = (name) => {
  const m = icons.match(new RegExp(`pub const ${name}: Icon = Icon \\{[^}]*?view_box: "([^"]+)"[^}]*?inner: r#"(.*?)"#`, 's'));
  if (!m) throw new Error(`rust/src/icon.rs has no ${name}`);
  return `<svg viewBox="${m[1]}">${m[2]}</svg>`;
};
const OURS = { sidebar: shipped('DOCS_SIDEBAR'), chevron: shipped('DOCS_CHEVRON'), article: shipped('DOCS_ARTICLE'), collection: shipped('DOCS_COLLECTION') };
// where each drawing's numbers were read
const RUNS = {
  sidebar: [['row', 4.635], ['row', 7], ['col', 1.2], ['col', 5.3], ['col', 7]],
  chevron: [['row', 7], ['row', 3], ['col', 6], ['col', 8]],
  article: [['row', 7.9], ['row', 3], ['col', 6], ['col', 8], ['col', 10.5]],
  collection: [['row', 3], ['row', 7], ['col', 0.875], ['col', 7]],
};

const src = readFileSync(new URL('./svg.js', import.meta.url), 'utf8');
const page = await open('chromium', 'https://developer.apple.com/documentation/xcode/examining-the-fields-in-a-crash-report');
let result;
try {
  await page.run('new Promise(r => setTimeout(r, 4000))');
  const apple = await page.run(`(async () => {
    const nav = [...document.querySelectorAll('.navigator-card-item .navigator-icon svg')];
    return {
      sidebar: document.querySelector('nav.documentation-nav button.sidenav-toggle svg').outerHTML,
      chevron: document.querySelector('.navigator-card-item button.tree-toggle svg').outerHTML,
      article: nav.find(s => s.classList.contains('article-icon')).outerHTML,
      collection: nav.find(s => s.classList.contains('collection-icon')).outerHTML,
    };
  })()`);
  await page.page.setContent('<!doctype html><body></body>');
  result = await page.run(`(async () => {
${src}
const apple = ${JSON.stringify(apple)}, ours = ${JSON.stringify(OURS)}, runs = ${JSON.stringify(RUNS)};
const out = {};
for (const k of Object.keys(ours)) {
  const a = await raster(apple[k], 40), o = await raster(ours[k], 40);
  out[k] = { appleBox: a.box, oursBox: o.box, iou: compare(a, o),
    runs: runs[k].map(([axis, at]) => [axis, at, a.runs(axis, at), o.runs(axis, at)]) };
}
return out;
})()`);
} finally {
  await page.close();
}
mkdirSync(new URL('./results/', import.meta.url), { recursive: true });
writeFileSync(new URL('./results/docc.json', import.meta.url), JSON.stringify({ browser: page.label, date: new Date().toISOString(), ours: OURS, ...result }, null, 1));
console.log(page.label);
let pass = true;
for (const [k, v] of Object.entries(result)) {
  const ok = v.iou >= 0.85;
  pass &&= ok;
  console.log(`${ok ? 'ok  ' : 'FAIL'} ${k}: ${v.iou} of DocC's ink in common`);
}
process.exit(pass ? 0 : 1);
