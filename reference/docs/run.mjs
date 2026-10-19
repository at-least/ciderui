// BROWSER=safari|ios|chrome|chromium SITE=<dir> node run.mjs
// Does ciderui's documentation (docs/, built with ciderui-zola) work in this
// browser as test/docs.mjs holds it to in Chromium? SITE is the site built
// for http://127.0.0.1:8737 (README.md), served there (the browsers:
// ../browsers.mjs). Every page: nothing scrolls sideways, every picture
// loads, and each example frame loads its stylesheet and its pictures, is no
// wider than the page and as tall as its content. Then what the frames do,
// clicked by the page's script (an element's click(), the activation a
// tap's click runs: safaridriver's taps on the iOS simulator never click):
// the theme's script moves a gallery's row (its Next paddle: the row
// scrolled, the previous paddle enabled, a later dot current) and a shelf's;
// an example's link leaves its frame where it is; in a window narrower than
// 834px the global nav's menu opens over its frame, grown to the window's
// height under the page's local nav and scrolled there, and the local nav's
// menu below its bar, the frame grown to hold it. On a phone at its width;
// elsewhere in a window 1440px wide, then 700px (Safari's window, its frame
// included, so its page a little narrower). The result is
// results/<browser>.json; exits 1 if a claim does not hold.
import { createServer } from 'node:http';
import { existsSync, mkdirSync, readFileSync, readdirSync, statSync, writeFileSync } from 'node:fs';
import { extname, join, relative } from 'node:path';
import { open } from '../browsers.mjs';

const BROWSER = process.env.BROWSER;
const SITE = process.env.SITE;
if (!SITE || !existsSync(join(SITE, 'index.html'))) throw new Error('SITE=<the docs built for http://127.0.0.1:8737> (README.md)');
const ORIGIN = 'http://127.0.0.1:8737';
const TYPES = { '.html': 'text/html; charset=utf-8', '.css': 'text/css', '.js': 'text/javascript', '.woff2': 'font/woff2', '.png': 'image/png', '.svg': 'image/svg+xml', '.xml': 'application/xml', '.txt': 'text/plain' };
const server = createServer((req, res) => {
  let f = join(SITE, decodeURIComponent(new URL(req.url, ORIGIN).pathname));
  if (existsSync(f) && statSync(f).isDirectory()) f = join(f, 'index.html');
  if (!existsSync(f)) { res.writeHead(404); return res.end(); }
  res.writeHead(200, { 'content-type': TYPES[extname(f)] ?? 'application/octet-stream' });
  res.end(readFileSync(f));
});
await new Promise((r) => server.listen(8737, '127.0.0.1', r));
const pages = (dir) => readdirSync(dir).flatMap((f) => {
  const p = join(dir, f);
  return statSync(p).isDirectory() ? pages(p) : f === 'index.html' ? [p] : [];
});
const paths = pages(SITE).map((f) => `/${relative(SITE, f).replace(/index\.html$/, '')}`).sort();

const wait = (ms) => `new Promise((r) => setTimeout(() => r(true), ${ms}))`;
// a page as it settles: the fonts, the frames' sizes a frame after theirs
const INSPECT = `(async () => {
  await document.fonts.ready;
  await new Promise((r) => setTimeout(r, 700));
  const broken = (doc) => [...doc.images].filter((i) => !i.complete || !i.naturalWidth).map((i) => i.getAttribute('src').slice(0, 48));
  return {
    width: innerWidth,
    sideways: document.documentElement.scrollWidth - innerWidth,
    images: broken(document),
    frames: [...document.querySelectorAll('iframe.cz-example')].map((f) => {
      const d = f.contentDocument;
      const sheet = [...d.styleSheets].find((s) => s.href && s.href.includes('/example.css'));
      return { title: f.title, height: f.clientHeight, content: d.documentElement.scrollHeight, css: Boolean(sheet && sheet.cssRules.length), images: broken(d), width: f.getBoundingClientRect().width, page: document.documentElement.clientWidth };
    }),
  };
})()`;
// a frame's element, found by its title, and a call on an element in it
const frame = (title) => `document.querySelector('iframe[title=${JSON.stringify(title)}]')`;
const inFrame = (title, selector, call) => `Promise.resolve((${frame(title)}.contentDocument.querySelector(${JSON.stringify(selector)}).${call}(), true))`;

const got = { pages: {}, frames: {} };
const fails = [];
const claim = (what, ok, value) => { if (!ok) fails.push(`${what}: ${JSON.stringify(value)}`); return ok; };
let page;
const read = async (width) => {
  for (const path of paths) {
    await page.go(ORIGIN + path);
    const r = await page.run(INSPECT);
    got.pages[`${path} at ${width}`] = r;
    const where = `${path} at ${r.width}px`;
    claim(`${where}: nothing scrolls sideways`, r.sideways <= 0, r.sideways);
    claim(`${where}: every picture loads`, r.images.length === 0, r.images);
    for (const f of r.frames) {
      claim(`${where}, "${f.title}": its frame loads its stylesheet and its pictures`, f.css && f.images.length === 0, f);
      claim(`${where}, "${f.title}": its frame is no wider than the page`, f.width <= f.page, f);
      claim(`${where}, "${f.title}": its frame is as tall as its content`, f.height === f.content, f);
    }
  }
};
const behave = async (width) => {
  // a row's Next paddle, clicked: the row scrolled on, the previous paddle enabled
  for (const [path, title, row] of [['/components/gallery/', 'A gallery of cards', 'gallery'], ['/components/shelf/', 'A shelf with paddles', 'shelf']]) {
    await page.go(ORIGIN + path);
    await page.run(wait(700));
    await page.run(inFrame(title, `[aria-controls="${row}"][aria-label="Next"]`, 'click'));
    await page.run(wait(1200));
    const r = await page.run(`Promise.resolve((() => {
      const d = ${frame(title)}.contentDocument;
      const paddles = [...d.querySelectorAll('[aria-controls="${row}"]')];
      return { scroll: d.getElementById('${row}').scrollLeft, previous: paddles[0].disabled, shown: getComputedStyle(paddles[1]).display !== 'none', current: [...d.querySelectorAll('.cui-gallery-dot')].findIndex((a) => a.getAttribute('aria-current') === 'true') };
    })())`);
    got.frames[`${row} at ${width}`] = r;
    claim(`${row}'s Next scrolls its row on in its frame at ${width}px, the previous paddle enabled`, r.scroll > 0 && !r.previous && (row !== 'gallery' || r.current > 0), r);
  }
  // an example's link, clicked
  await page.go(`${ORIGIN}/components/link/`);
  await page.run(wait(700));
  await page.run(inFrame('A more link', 'a', 'click'));
  await page.run(wait(700));
  const at = await page.run(`Promise.resolve(${frame('A more link')}.contentWindow.location.href)`);
  claim(`a link in an example leaves its frame where it is at ${width}px`, at === 'about:srcdoc', at);
};
const menus = async (width) => {
  await page.go(`${ORIGIN}/components/globalnav/`);
  await page.run(wait(700));
  const height = (title) => page.run(`Promise.resolve(${frame(title)}.clientHeight)`);
  const closed = await height('A global nav');
  await page.run(inFrame('A global nav', '.cui-globalnav-menubutton', 'click'));
  await page.run(wait(1000));
  const opened = await page.run(`Promise.resolve((() => {
    const f = ${frame('A global nav')};
    const bar = document.querySelector('.cui-localnav').getBoundingClientRect().height;
    return { open: f.contentDocument.querySelector('.cui-globalnav-list').matches(':popover-open'), height: f.clientHeight, top: Math.round(f.getBoundingClientRect().top), window: innerHeight, bar };
  })())`);
  got.frames[`globalnav menu at ${width}`] = { closed, ...opened };
  claim(`the global nav's menu opens in its frame at ${width}px, the frame the window's height under the local nav, scrolled there`, opened.open && Math.abs(opened.height - (opened.window - opened.bar)) <= 1 && Math.abs(opened.top - opened.bar) <= 1, { closed, ...opened });
  await page.run(inFrame('A global nav', '.cui-globalnav-closebutton', 'click'));
  await page.run(wait(700));
  const again = await height('A global nav');
  claim(`closed again, its frame is the bar's height at ${width}px`, again === closed, { closed, again });

  await page.go(`${ORIGIN}/components/localnav/`);
  await page.run(wait(700));
  const bar = await height('A local nav');
  await page.run(inFrame('A local nav', '.cui-localnav-menubutton', 'click'));
  await page.run(wait(1000));
  const menu = await page.run(`Promise.resolve((() => {
    const f = ${frame('A local nav')};
    return { checked: f.contentDocument.querySelector('.cui-localnav-toggle').checked, bottom: Math.ceil(f.contentDocument.querySelector('.cui-localnav-menu').getBoundingClientRect().bottom), height: f.clientHeight };
  })())`);
  got.frames[`localnav menu at ${width}`] = { bar, ...menu };
  claim(`the local nav's menu opens below its bar at ${width}px, its frame grown to hold it`, menu.checked && menu.bottom > bar && menu.height >= menu.bottom, { bar, ...menu });
};

try {
  page = await open(BROWSER, `${ORIGIN}/`);
  got.browser = page.label;
  // a phone's width as it is; elsewhere a wide window, then a narrow one
  for (const size of BROWSER === 'ios' ? [null] : [[1440, 900], [700, 900]]) {
    if (size) await page.size(...size);
    const width = await page.run('Promise.resolve(innerWidth)');
    await read(width);
    await behave(width);
    if (width < 834) await menus(width);
  }
} finally {
  if (page) await page.close();
  server.close();
}
got.fails = fails;
mkdirSync(new URL('./results/', import.meta.url), { recursive: true });
writeFileSync(new URL(`./results/${BROWSER}.json`, import.meta.url), JSON.stringify(got, null, 1));
for (const f of fails) console.log(`FAIL ${f}`);
console.log(`${got.browser}: ${paths.length} pages, ${fails.length} failures`);
process.exitCode = fails.length ? 1 : 0;
