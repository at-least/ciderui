// Draw the documentation's pictures (docs/static/pictures/<name>.png): ciderui
// drawn by itself, for the places apple.com's pages put a picture. Each is a
// scene of one or more of the components' examples on the alternate ground,
// drawn in Chromium (--font-render-hinting=none, as the checks draw) with
// the example frames' stylesheet (docs/static/example.css, so run
// scripts/build-docs.mjs first) and Inter, at 1280 by 720 (16:9: a
// component's row, a story), or the home's tile at 1440 by 692. A small
// part is drawn larger (zoom) and centered; a band of the page at its
// width, from the top.
//
// A picture is what this Chromium drew, and another version draws other
// bytes, so what is held is what it was drawn from: docs/pictures.json
// records each picture's scene (its document, the examples' markup in it,
// its size) and the stylesheet's digest, and --check fails when an example,
// the stylesheet or a scene has changed since, or a picture is missing:
//
//   node scripts/docs-pictures.mjs           draw them
//   node scripts/docs-pictures.mjs --check   fail if one was drawn from other inputs

import { createHash } from 'node:crypto';
import { existsSync, mkdirSync, readFileSync, writeFileSync } from 'node:fs';
import { join } from 'node:path';
import { ROOT } from '../test/compile.mjs';

const DOCS = join(ROOT, 'docs');
const OUT = join(DOCS, 'static/pictures');
const MANIFEST = join(DOCS, 'pictures.json');
const example = (id) => readFileSync(join(ROOT, 'components', id.replace('/', '/examples/') + '.html'), 'utf8');

// ciderui's own parts as a shelf of the store's cards, each its name, its
// kind and what it is, for the home's tile and the catalogue's story: the
// examples' cards name apple.com's products
const card = (kind, name, what) => `<li><div class="cui-card"><div class="cui-card-content"><p class="cui-card-eyebrow">${kind}</p><h3 class="cui-card-headline"><a class="cui-card-link" href="#">${name}</a></h3><p class="cui-card-description">${what}</p></div></div></li>`;
const PARTS = `<div class="cui-shelf"><div class="cui-shelf-row"><ul class="cui-shelf-scroller">${[
  card('Navigation', 'Global Nav', 'The bar across the top, its menu a popover.'),
  card('Controls', 'Button', 'The pill, in four colors and five sizes.'),
  card('Page', 'Tile', 'The home page’s tile, its picture filling it.'),
  card('Store', 'Card', 'The store’s card, on a shelf like this one.'),
  card('Newsroom', 'Article', 'The Newsroom’s story, for a page’s markdown.'),
].join('')}</ul></div></div>`;

// name: the examples drawn, one after another (or html, a scene's own: the
// modal's opens it, then blurs its close button, whose ring a script's
// opening shows and a click's would not);
// zoom: how much larger;
// width: the scene's own width before the zoom (the window's, for a band;
// none, a small part's own);
// center: a small part in the middle of the picture, else from the top;
// window, height: the picture's size (1280 by 720); top: room above
const SCENES = [
  { name: 'accordion', files: ['accordion/accordion'], zoom: 1.5, width: 640, center: true },
  { name: 'article', files: ['article/article'], zoom: 1, width: 1280 },
  { name: 'button', files: ['button/variants'], zoom: 2.5, center: true },
  { name: 'card', files: ['card/card'], zoom: 1.2, width: 400, center: true },
  { name: 'checkbox', files: ['checkbox/checked'], zoom: 3, center: true },
  { name: 'field', files: ['field/filled'], zoom: 2, width: 400, center: true },
  { name: 'footer', files: ['footer/footer'], zoom: 1.4, width: 914 },
  { name: 'footnote', files: ['footnote/footnote'], zoom: 3, center: true },
  { name: 'gallery', files: ['gallery/gallery'], zoom: 1, width: 1280 },
  { name: 'globalnav', files: ['globalnav/globalnav', 'tile/hero'], zoom: 1, width: 1280 },
  { name: 'link', files: ['link/more'], zoom: 3.5, center: true },
  { name: 'localnav', files: ['localnav/localnav', 'tile/tile'], zoom: 1, width: 1280 },
  { name: 'modal', html: `${example('modal/modal')}<script>document.querySelector('.cui-modal').showModal(); document.activeElement.blur()</script>`, zoom: 1, width: 1280 },
  { name: 'option', files: ['option/checked', 'option/option'], zoom: 1.6, width: 500, center: true },
  { name: 'posts', files: ['posts/archive'], zoom: 1, width: 1280 },
  { name: 'price', files: ['price/price'], zoom: 2.5, center: true },
  { name: 'productnav', files: ['productnav/productnav'], zoom: 2, width: 640, center: true },
  { name: 'quicklinks', files: ['quicklinks/quicklinks'], zoom: 2, width: 640, center: true },
  { name: 'ribbon', files: ['ribbon/ribbon', 'tile/tile'], zoom: 1, width: 1280 },
  { name: 'section', files: ['section/header'], zoom: 1.4, width: 914, center: true },
  { name: 'shelf', files: ['shelf/shelf'], zoom: 1, width: 1280, center: true },
  { name: 'stat', files: ['stat/large'], zoom: 1.5, center: true },
  { name: 'stories', files: ['stories/stories'], zoom: 1, width: 1280 },
  { name: 'swatch', files: ['swatch/swatches'], zoom: 4, center: true },
  { name: 'tile', files: ['tile/hero'], zoom: 1, width: 1280 },
  { name: 'parts', html: PARTS, zoom: 1, width: 1280, center: true },
  { name: 'home', html: PARTS, zoom: 1, width: 1440, top: 300, window: 1440, height: 692 },
];

const css = readFileSync(join(DOCS, 'static/example.css'), 'utf8');
const digest = (text) => createHash('sha256').update(text).digest('hex');

// each scene's size and the document drawn
const scenes = SCENES.map((s) => {
  const height = s.height ?? 720;
  const box = `zoom:${s.zoom};width:${s.width ? `${s.width}px` : 'max-content'};${s.top ? `margin-top:${s.top}px;` : ''}${s.center ? '' : 'margin-inline:auto;'}`;
  return {
    name: s.name,
    width: s.window ?? 1280,
    height,
    html: `<!doctype html><html lang="en"><head><meta charset="utf-8"><link rel="stylesheet" href="/example.css"></head><body style="margin:0;height:${height}px;overflow:hidden;background:var(--cui-c-bg-alt);${s.center ? 'display:flex;align-items:center;justify-content:center;' : ''}"><div style="${box}">${s.html ?? s.files.map(example).join('\n')}</div></body></html>`,
  };
});
// what each picture is drawn from
const inputs = Object.fromEntries(scenes.map((s) => [s.name, digest(`${s.width}x${s.height}\n${s.html}\n${css}`)]));

if (process.argv[2] === '--check') {
  const drawn = existsSync(MANIFEST) ? JSON.parse(readFileSync(MANIFEST, 'utf8')) : {};
  const stale = [
    ...Object.keys(inputs).filter((n) => drawn[n] !== inputs[n]).map((n) => `docs/static/pictures/${n}.png was drawn from ${drawn[n] ? 'other examples or another stylesheet' : 'nothing recorded'}`),
    ...Object.keys(inputs).filter((n) => !existsSync(join(OUT, `${n}.png`))).map((n) => `docs/static/pictures/${n}.png is missing`),
    ...Object.keys(drawn).filter((n) => !(n in inputs)).map((n) => `docs/pictures.json records ${n}, which no scene draws`),
  ];
  for (const s of stale) console.log(s);
  console.log(stale.length ? 'the pictures are stale: node scripts/docs-pictures.mjs' : `docs/static/pictures are drawn from these examples and this stylesheet (${scenes.length})`);
  process.exit(stale.length ? 1 : 0);
}

const { chromium } = await import('playwright-chromium');
const browser = await chromium.launch({ args: ['--font-render-hinting=none'] });
const page = await browser.newPage();
let current = '';
await page.route('https://pictures.test/**', (route) => {
  const path = new URL(route.request().url()).pathname;
  if (path.startsWith('/fonts/')) return route.fulfill({ body: readFileSync(join(ROOT, path)), contentType: 'font/woff2' });
  if (path === '/example.css') return route.fulfill({ body: css, contentType: 'text/css' });
  return route.fulfill({ body: current, contentType: 'text/html' });
});
mkdirSync(OUT, { recursive: true });
for (const s of scenes) {
  await page.setViewportSize({ width: s.width, height: s.height });
  current = s.html;
  await page.goto('https://pictures.test/', { waitUntil: 'load' });
  await page.evaluate(() => document.fonts.ready);
  // what moves as it shows (a modal fading in) at its end
  await page.waitForTimeout(300);
  await page.screenshot({ path: join(OUT, `${s.name}.png`) });
  console.log(`docs/static/pictures/${s.name}.png`);
}
await browser.close();
writeFileSync(MANIFEST, `${JSON.stringify(inputs, null, 1)}\n`);
