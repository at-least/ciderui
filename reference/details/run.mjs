// BROWSER=safari|ios|chrome|chromium node run.mjs: does cui-accordion's
// <details> open with no script, its answer shown and its chevron turned, and
// does its height move as it opens? Serves page.html (the accordion's
// example on the kit's compiled CSS, `node page.mjs` writes it) on
// 127.0.0.1 (the browsers: ../browsers.mjs); clicks the second question (by
// the page's script on the iOS simulator, as ../popover/ does) and reads the
// item's height 120ms on and 700ms on. The result is results/<browser>.json.
import { createServer } from 'node:http';
import { mkdirSync, readFileSync, writeFileSync } from 'node:fs';
import { open } from '../browsers.mjs';

const BROWSER = process.env.BROWSER;
const html = readFileSync(new URL('./page.html', import.meta.url));
const server = createServer((req, res) => { res.writeHead(200, { 'content-type': 'text/html; charset=utf-8' }); res.end(html); });
await new Promise(r => server.listen(0, '127.0.0.1', r));
const ITEM = '.cui-accordion-item:nth-child(2)';
const read = `Promise.resolve((() => { const d = document.querySelector('${ITEM}'); return { open: d.open, item: d.getBoundingClientRect().height, answer: d.querySelector('.cui-accordion-answer').getBoundingClientRect().height, chevron: getComputedStyle(d.querySelector('svg')).transform }; })())`;
const wait = (ms) => `new Promise(r => setTimeout(() => r(true), ${ms}))`;
let page;
const got = {};
try {
  page = await open(BROWSER, `http://127.0.0.1:${server.address().port}/`);
  got.closed = await page.run(read);
  if (BROWSER === 'ios') await page.run(`Promise.resolve((document.querySelector('${ITEM} > summary').click(), true))`);
  else await page.click(`${ITEM} > summary`);
  await page.run(wait(120));
  got.opening = await page.run(read);
  await page.run(wait(580));
  got.opened = await page.run(read);
  got.interpolateSize = await page.run(`Promise.resolve(CSS.supports('interpolate-size: allow-keywords'))`);
} finally {
  await page?.close();
  server.close();
}
mkdirSync(new URL('./results/', import.meta.url), { recursive: true });
writeFileSync(new URL(`./results/${BROWSER}.json`, import.meta.url), JSON.stringify({ browser: page.label, date: new Date().toISOString(), ...got }, null, 1));
const { closed, opening, opened } = got;
const claims = [
  ['it opens with no script, its answer shown', !closed.open && opened.open && opened.item > closed.item + 30 && opened.answer > 30],
  ['its chevron turned up', opened.chevron === 'matrix(1, 0, 0, -1, 0, 0)'],
  [got.interpolateSize ? 'its height moves as it opens (a height between, 120ms on)' : 'it opens at once (no interpolate-size: its full height 120ms on)',
    got.interpolateSize ? opening.item > closed.item + 1 && opening.item < opened.item - 1 : opening.item === opened.item],
];
console.log(page.label);
let pass = true;
for (const [claim, ok] of claims) {
  pass &&= ok;
  console.log(`${ok ? 'ok  ' : 'FAIL'} ${claim}`);
}
process.exit(pass ? 0 : 1);
