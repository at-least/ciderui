// BROWSER=safari SITE=<dir> node inactive.mjs
// Does Safari match :focus, and run a shelf's smooth scroll, in a window
// that is not its key window, or in a page hidden, as when the owner's app
// comes in front of it during a run (claim 9's textarea not focused, twice in
// five runs; the shelf's row read 37px on, once: README.md)? On the site
// run.mjs reads (served the same), in a window 700px wide: the textarea's
// frame, its textarea focused by the page's script as run.mjs focuses it,
// its box's edge and label read with :focus, the active element and
// document.hasFocus(), the frame's and the page's; the shelf's frame,
// brought into the window, its Next paddle clicked by script and its row
// read 1.2s on, with the page's visibilityState. Each read with Safari's
// window in front; behind a second window of Safari's own (../browsers.mjs's
// behind); in front again. Then the shelf's paddle clicked in front and a tab
// of Safari's own selected at once in the page's window, the page hidden, its
// row read as the tab is closed 1.2s on and again 1.2s after, with the page's
// own log of its row and its visibility every 50ms from the click (its
// timers run on hidden, if less often). The result is
// results/inactive-<browser>.json.
import { createServer } from 'node:http';
import { existsSync, mkdirSync, readFileSync, statSync, writeFileSync } from 'node:fs';
import { extname, join } from 'node:path';
import { open } from '../browsers.mjs';

const BROWSER = process.env.BROWSER;
const SITE = process.env.SITE;
if (!SITE || !existsSync(join(SITE, 'index.html'))) throw new Error('SITE=<the docs built for http://127.0.0.1:8737> (README.md)');
const ORIGIN = 'http://127.0.0.1:8737';
const TYPES = { '.html': 'text/html; charset=utf-8', '.css': 'text/css', '.js': 'text/javascript', '.woff2': 'font/woff2', '.png': 'image/png', '.svg': 'image/svg+xml', '.json': 'application/json' };
const server = createServer((req, res) => {
  let f = join(SITE, decodeURIComponent(new URL(req.url, ORIGIN).pathname));
  if (existsSync(f) && statSync(f).isDirectory()) f = join(f, 'index.html');
  if (!existsSync(f)) { res.writeHead(404); return res.end(); }
  res.writeHead(200, { 'content-type': TYPES[extname(f)] ?? 'application/octet-stream' });
  res.end(readFileSync(f));
});
await new Promise((r) => server.listen(8737, '127.0.0.1', r));
let page;
// the waits are the run's own, no command to the page, which might bring it
// in front again
const sleep = (ms) => new Promise((r) => setTimeout(r, ms));
const frame = (title) => `document.querySelector('iframe[title=${JSON.stringify(title)}]')`;
// what the page saw: its visibility's changes and its window's blur and
// focus, each with its time on the page's clock
const go = async (path) => {
  await page.go(ORIGIN + path);
  await page.run(`Promise.resolve((() => {
    window.seen = [];
    const note = (what) => () => seen.push([what, document.visibilityState, Math.round(performance.now())]);
    document.addEventListener('visibilitychange', note('visibilitychange'));
    addEventListener('blur', note('blur'));
    addEventListener('focus', note('focus'));
    return true;
  })())`);
  await sleep(700);
};
const textarea = async () => {
  const title = 'A message of several lines';
  await go('/components/field/');
  await page.run(`Promise.resolve((${frame(title)}.contentDocument.querySelector('textarea').focus(), true))`);
  await sleep(400);
  return page.run(`Promise.resolve((() => {
    const f = ${frame(title)};
    const d = f.contentDocument, t = d.querySelector('textarea');
    const box = getComputedStyle(d.querySelector('.cui-field-box'));
    return { edge: box.borderTopWidth, label: getComputedStyle(d.querySelector('.cui-field-label')).top, focus: t.matches(':focus'), active: d.activeElement === t, frameHasFocus: d.hasFocus(), pageHasFocus: document.hasFocus(), visibility: document.visibilityState, seen, now: Math.round(performance.now()) };
  })())`);
};
// hidden: a tab selected in front of the page with the click, closed 1.2s on,
// the row read then and 1.2s after
const shelf = async (hidden) => {
  const title = 'A shelf with paddles';
  await go('/components/shelf/');
  await page.run(`Promise.resolve((${frame(title)}.scrollIntoView({ block: 'center' }), true))`);
  await sleep(300);
  await page.run(`Promise.resolve((() => {
    const row = ${frame(title)}.contentDocument.getElementById('shelf');
    row.ownerDocument.querySelector('[aria-controls="shelf"][aria-label="Next"]').click();
    window.clicked = Math.round(performance.now());
    window.rows = [];
    const every = setInterval(() => rows.push([Math.round(performance.now()), row.scrollLeft, document.visibilityState]), 50);
    setTimeout(() => clearInterval(every), 4500);
    return true;
  })())`);
  const made = hidden && await page.behind('tab');
  await sleep(1200);
  if (hidden) await page.behind(false);
  const read = `Promise.resolve({ scroll: ${frame(title)}.contentDocument.getElementById('shelf').scrollLeft, pageHasFocus: document.hasFocus(), visibility: document.visibilityState, now: Math.round(performance.now()) })`;
  const r = await page.run(read);
  if (hidden) {
    await sleep(1200);
    r.after = await page.run(read);
  }
  return { ...(hidden && { made }), ...r, ...await page.run('Promise.resolve({ seen, clicked, rows })') };
};
const got = {};
try {
  page = await open(BROWSER, `${ORIGIN}/`);
  got.browser = page.label;
  await page.size(700, 900);
  got.front = { textarea: await textarea(), shelf: await shelf() };
  const made = await page.behind('window');
  await sleep(1000);
  got.window = { made, textarea: await textarea(), shelf: await shelf() };
  await page.behind(false);
  await sleep(1000);
  got.again = { textarea: await textarea(), shelf: await shelf() };
  got.hidden = { shelf: await shelf(true) };
} finally {
  if (page) await page.close();
  server.close();
}
mkdirSync(new URL('./results/', import.meta.url), { recursive: true });
writeFileSync(new URL(`./results/inactive-${BROWSER}.json`, import.meta.url), JSON.stringify({ date: new Date().toISOString(), ...got }, null, 1));
console.log(got.browser);
for (const k of ['front', 'window', 'again', 'hidden']) console.log(k.padEnd(7), JSON.stringify(got[k]));
