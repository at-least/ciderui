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
// records, for each picture, one digest of its scene (its document, the
// examples' markup in it, its size), the stylesheet and the fonts, and --check fails
// when an example, the stylesheet, a font or a scene has changed since, or a
// picture is missing:
//
//   node scripts/docs-pictures.mjs           draw them
//   node scripts/docs-pictures.mjs --check   fail if one was drawn from other inputs

import { createHash } from 'node:crypto';
import { existsSync, mkdirSync, readFileSync, readdirSync, writeFileSync } from 'node:fs';
import { join } from 'node:path';
import { ROOT } from '../test/compile.mjs';

const DOCS = join(ROOT, 'docs');
const OUT = join(DOCS, 'static/pictures');
const MANIFEST = join(DOCS, 'pictures.json');
const example = (id) => readFileSync(join(ROOT, 'components', id.replace('/', '/examples/') + '.html'), 'utf8');

// ciderui's own parts as a shelf of the store's cards, each its name, its
// kind and what it is, for the home's tile and the catalogue's story: the
// examples' cards name a store's products
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
  { name: 'console-badge', html: `<div class="cui-console-page" style="width:480px;padding:24px;display:grid;gap:20px">${example('console-badge/badges')}${example('console-badge/plain')}</div>`, zoom: 2, center: true },
  { name: 'console-chip', html: `<div class="cui-console-page" style="width:420px;padding:24px">${example('console-chip/chips')}</div>`, zoom: 2, center: true },
  { name: 'console-dialog', html: `<div class="cui-console-page" style="width:640px;height:300px;padding:24px;display:grid;place-items:center">${example('console-dialog/dialog').replace(' aria-labelledby="ignore-title">', ' aria-labelledby="ignore-title" open>')}</div>`, zoom: 1.4, center: true },
  { name: 'console-disclosure', html: `<div class="cui-console-page" style="width:420px;padding:24px">${example('console-disclosure/disclosure')}</div>`, zoom: 2, center: true },
  { name: 'console-banner', html: `<div class="cui-console-page" style="width:560px;padding:24px">${example('console-banner/banner')}</div>`, zoom: 1.8, center: true },
  { name: 'console-bar', html: `<div class="cui-console-page" style="width:520px;padding:24px">${example('console-bar/bar')}</div>`, zoom: 2, center: true },
  { name: 'console-btn', html: `<div class="cui-console-page" style="padding:24px">${example('console-btn/variants')}</div>`, zoom: 2, width: 560, center: true },
  { name: 'console-card', html: `<div class="cui-console-page" style="width:900px;padding:24px">${example('console-card/metrics')}</div>`, zoom: 1.2, center: true },
  { name: 'console-chart', html: `<div class="cui-console-page" style="width:640px;padding:24px">${example('console-chart/chart')}</div>`, zoom: 2, center: true },
  { name: 'console-checkbox', html: `<div class="cui-console-page" style="padding:24px">${example('console-checkbox/checkbox')}</div>`, zoom: 2.4, center: true },
  { name: 'console-empty', html: `<div class="cui-console-page" style="width:520px;padding:24px">${example('console-empty/empty')}</div>`, zoom: 1.8, center: true },
  { name: 'console-file', html: `<div class="cui-console-page" style="width:420px;padding:24px">${example('console-file/file')}</div>`, zoom: 2, center: true },
  { name: 'console-hint', html: `<div class="cui-console-page" style="width:460px;height:200px;padding:24px">${example('console-hint/hint').replace(' popover>', ' popover style="display:block;position:static;inset:auto;margin:8px 0 0">')}</div>`, zoom: 2, center: true },
  { name: 'console-popover', html: `<div class="cui-console-page" style="width:460px;height:220px;padding:24px">${example('console-popover/popover').replace(' popover>', ' popover style="display:flex;position:static;inset:auto;margin:8px auto 0">')}</div>`, zoom: 2, center: true },
  { name: 'console-list', html: `<div class="cui-console-page" style="width:480px;padding:24px">${example('console-list/list')}</div>`, zoom: 1.8, center: true },
  { name: 'console-progress', html: `<div class="cui-console-page" style="width:420px;padding:24px">${example('console-progress/progress')}</div>`, zoom: 2, center: true },
  { name: 'console-dl', html: `<div class="cui-console-page" style="width:420px;padding:24px">${example('console-dl/dl')}</div>`, zoom: 2, center: true },
  { name: 'console-field', html: `<div class="cui-console-page" style="width:380px;padding:24px">${example('console-field/fields')}</div>`, zoom: 1.4, center: true },
  { name: 'console-status', html: `<div class="cui-console-page" style="width:420px;padding:24px;display:grid;gap:20px">${example('console-status/status')}${example('console-status/detail')}</div>`, zoom: 2, center: true },
  { name: 'console-code', html: `<div class="cui-console-page" style="width:420px;padding:24px">${example('console-code/code')}</div>`, zoom: 2, center: true },
  { name: 'console-skip', html: `<div class="cui-console-page" style="width:420px;height:120px;padding:24px">${example('console-skip/skip').replace('class="cui-console-skip ', 'style="position:static" class="')}</div>`, zoom: 2, center: true },
  { name: 'console-toast', html: `<div class="cui-console-page" style="width:420px;height:160px;padding:24px;display:grid;place-items:center">${example('console-toast/toast')}</div>`, zoom: 2, center: true },
  { name: 'console-pager', html: `<div class="cui-console-page" style="width:520px;padding:24px">${example('console-pager/pager')}</div>`, zoom: 2, center: true },
  { name: 'console-thumb', html: `<div class="cui-console-page" style="width:360px;padding:24px">${example('console-thumb/thumbs')}</div>`, zoom: 1.6, center: true },
  { name: 'console-content', html: `<div style="width:1280px">${example('console-content/content')}</div>`, zoom: 1, center: true, window: 1280 },
  { name: 'console-link', files: ['console-link/link'], zoom: 2, width: 560, center: true },
  { name: 'console-appearance', html: `<div class="cui-console-page" style="width:240px;padding:24px">${example('console-appearance/appearance')}</div>`, zoom: 2, center: true },
  { name: 'console-copy', html: `<div class="cui-console-page" style="width:560px;padding:24px">${example('console-copy/copy')}</div>`, zoom: 2, center: true },
  { name: 'console-head', html: `<div class="cui-console-page" style="width:600px;padding:24px">${example('console-head/head')}</div>`, zoom: 2, center: true },
  { name: 'console-section-head', html: `<div class="cui-console-page" style="width:600px;padding:24px">${example('console-section-head/section-head')}</div>`, zoom: 2, center: true },
  { name: 'console-listing', html: `<div class="cui-console-page" style="width:640px;padding:24px">${example('console-listing/listing')}</div>`, zoom: 2, center: true },
  { name: 'console-menu', html: `<div class="cui-console-page" style="width:260px;padding:16px">${example('console-menu/menu')}</div>`, zoom: 1.6, center: true },
  { name: 'console-tabs', html: `<div class="cui-console-page" style="width:440px;padding:24px;display:grid;gap:32px">${example('console-tabs/tabs')}${example('console-tabs/filled')}</div>`, zoom: 2, center: true },
  { name: 'console-sidebar', html: `<div class="cui-console-page" style="height:640px;display:flex;padding:8px 0">${example('console-sidebar/sidebar')}</div>`, zoom: 1.1, center: true, window: 1280 },
  { name: 'console-table', html: `<div class="cui-console-page" style="width:520px;padding:24px">${example('console-table/table')}</div>`, zoom: 1.6, center: true },
  { name: 'console-page', files: ['console-page/page'], zoom: 2, width: 560, center: true },
  { name: 'console-radio', html: `<div class="cui-console-page" style="width:420px;padding:24px">${example('console-radio/radios')}</div>`, zoom: 2, center: true },
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
// the fonts the pictures are drawn in (Inter, fonts/), each file's digest
const fonts = readdirSync(join(ROOT, 'fonts')).sort().map((f) => `${f} ${digest(readFileSync(join(ROOT, 'fonts', f)))}`).join('\n');

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
const inputs = Object.fromEntries(scenes.map((s) => [s.name, digest(`${s.width}x${s.height}\n${s.html}\n${css}\n${fonts}`)]));

if (process.argv[2] === '--check') {
  const drawn = existsSync(MANIFEST) ? JSON.parse(readFileSync(MANIFEST, 'utf8')) : {};
  const stale = [
    ...Object.keys(inputs).filter((n) => drawn[n] !== inputs[n]).map((n) => `docs/static/pictures/${n}.png was drawn from ${drawn[n] ? 'other examples, another stylesheet or other fonts' : 'nothing recorded'}`),
    ...Object.keys(inputs).filter((n) => !existsSync(join(OUT, `${n}.png`))).map((n) => `docs/static/pictures/${n}.png is missing`),
    ...Object.keys(drawn).filter((n) => !(n in inputs)).map((n) => `docs/pictures.json records ${n}, which no scene draws`),
  ];
  for (const s of stale) console.log(s);
  console.log(stale.length ? 'the pictures are stale: node scripts/docs-pictures.mjs' : `docs/static/pictures are drawn from these examples, this stylesheet and these fonts (${scenes.length})`);
  process.exit(stale.length ? 1 : 0);
}

const { launch } = await import('../test/chromium.mjs');
const browser = await launch();
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
