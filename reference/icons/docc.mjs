// node docc.mjs: Apple's developer documentation's icons the docs set draws
// (the bar's button that opens the navigator, a navigator row's chevron, an
// article's and a collection's glyph, the navigator sheet's close button, a
// listing's copy button and its check once it has copied: the button
// clicked, the clipboard granted, the navigator's filter's glyph),
// DocC's SVG read live off the crash report's article (never stored), each
// rasterized at 40 pixels to the unit on a blank page (svg.js) beside the
// drawing the kit ships, read out of go/icon/icon.go (DocsSidebar,
// DocsChevron, DocsArticle, DocsCollection, DocsClose, DocsCopy,
// DocsCopied, DocsFilter): the share of ink
// each pair has in common, and each one's ink and its runs along rows and
// columns, the numbers the drawings are drawn from. Fails below 0.85
// (intersection over union), the mark the global nav's icons are held to.
// The result is results/docc.json.
import { mkdirSync, readFileSync, writeFileSync } from 'node:fs';
import { open } from '../browsers.mjs';

const icons = readFileSync(new URL('../../go/icon/icon.go', import.meta.url), 'utf8');
const shipped = (name) => {
  const m = icons.match(new RegExp(`^var ${name} = Icon\\{[^\\n]*?viewBox: "([^"]+)"[^\\n]*?inner: \`([^\`]*)\`\\}$`, 'm'));
  if (!m) throw new Error(`go/icon/icon.go has no ${name}`);
  return `<svg viewBox="${m[1]}">${m[2]}</svg>`;
};
const OURS = { sidebar: shipped('DocsSidebar'), chevron: shipped('DocsChevron'), article: shipped('DocsArticle'), collection: shipped('DocsCollection'), close: shipped('DocsClose'), copy: shipped('DocsCopy'), copied: shipped('DocsCopied'), filter: shipped('DocsFilter') };
// where each drawing's numbers were read
const RUNS = {
  sidebar: [['row', 4.635], ['row', 7], ['col', 1.2], ['col', 5.3], ['col', 7]],
  chevron: [['row', 7], ['row', 3], ['col', 6], ['col', 8]],
  article: [['row', 7.9], ['row', 3], ['col', 6], ['col', 8], ['col', 10.5]],
  collection: [['row', 3], ['row', 7], ['col', 0.875], ['col', 7]],
  close: [['row', 3.5], ['row', 7], ['col', 7]],
  copy: [['row', 5], ['row', 10], ['col', 6], ['col', 12]],
  copied: [['row', 8], ['row', 12], ['row', 16], ['col', 10]],
  filter: [['row', 6.3], ['row', 10.5], ['row', 14.7], ['col', 7]],
};

const src = readFileSync(new URL('./svg.js', import.meta.url), 'utf8');
const page = await open('chromium', 'https://developer.apple.com/documentation/xcode/examining-the-fields-in-a-crash-report', { permissions: ['clipboard-read', 'clipboard-write'] });
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
      close: document.querySelector('.navigator button.close-card svg').outerHTML,
      filter: document.querySelector('.navigator .navigator-filter .filter__filter-button svg').outerHTML,
      copy: document.querySelector('.code-listing button.copy-button svg').outerHTML,
      copied: await new Promise((done) => {
        const b = document.querySelector('.code-listing button.copy-button');
        b.click();
        setTimeout(() => done(b.querySelector('svg').outerHTML), 300);
      }),
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
