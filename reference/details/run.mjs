// BROWSER=safari|ios|chrome|chromium|firefox node run.mjs: does cui-accordion's
// <details> open with no script, its answer shown and its chevron turned, and
// does its height move as it opens? Serves page.html (the accordion's
// example on the kit's compiled CSS, `node page.mjs` writes it) on
// 127.0.0.1 (the browsers: ../browsers.mjs); clicks the second question (by
// the page's script on the iOS simulator, as ../popover/ does) and reads the
// item's height 120ms on and 700ms on. Then, where there is a keyboard (not
// the simulator), is the focus ring of what is in an open answer shown
// whole, the answer clipping its height as it moves? ring.html, at 1200 and
// 390px: a compact question in the large one's answer and a link beginning
// an answer, each focused by Tab (Option-Tab in Safari, whose Tab skips
// links), the window's picture (../browsers.mjs's shot) read in the page
// (ring.js): the accent's pixels in the 8px beside it, left and right, where
// its ring's sides are drawn 3 to 5px out. And is its height clipped still?
// held.html's answer held 10px tall: none of its text drawn, where
// held-visible.html, clipping nothing, draws it. The click and each ring's
// read wait for the window in front (front()). The result is
// results/<browser>.json.
import { createServer } from 'node:http';
import { mkdirSync, readFileSync, writeFileSync } from 'node:fs';
import { open } from '../browsers.mjs';

const BROWSER = process.env.BROWSER;
const html = readFileSync(new URL('./page.html', import.meta.url));
const PAGES = Object.fromEntries(['ring', 'held', 'held-visible'].map((name) => [`/${name}.html`, readFileSync(new URL(`./${name}.html`, import.meta.url))]));
const server = createServer((req, res) => { res.writeHead(200, { 'content-type': 'text/html; charset=utf-8' }); res.end(PAGES[req.url] ?? html); });
await new Promise(r => server.listen(0, '127.0.0.1', r));
const ITEM = '.cui-accordion-item:nth-child(2)';
// ring.html's: the first compact question in the large one's answer, and the
// link beginning the compact example's first answer
const RINGS = {
  question: '.cui-accordion[data-size="lg"] .cui-accordion:not([data-size]) > .cui-accordion-item:first-child > .cui-accordion-question',
  link: 'main > .cui-accordion:not([data-size]) .cui-accordion-answer a',
};
const read = `Promise.resolve((() => { const d = document.querySelector('${ITEM}'); return { open: d.open, item: d.getBoundingClientRect().height, answer: d.querySelector('.cui-accordion-answer').getBoundingClientRect().height, chevron: getComputedStyle(d.querySelector('svg')).transform }; })())`;
const wait = (ms) => `new Promise(r => setTimeout(() => r(true), ${ms}))`;
let page;
// the window in front, where Safari matches :focus and a click reaches the
// page: an app of the owner's took the Mac's front by itself during a run,
// back and forth with Safari every second or two, no input (2026-10-22), so
// what needs it waits for it, up to 10s, and names how long (null: never)
const front = async () => {
  for (let i = 0; i < 40; i++) {
    if (await page.run('Promise.resolve(document.hasFocus())')) return i * 250;
    await page.run(wait(250));
  }
  return null;
};
const got = { waited: {} };
try {
  page = await open(BROWSER, `http://127.0.0.1:${server.address().port}/`);
  got.closed = await page.run(read);
  got.waited.click = await front();
  if (BROWSER === 'ios') await page.run(`Promise.resolve((document.querySelector('${ITEM} > summary').click(), true))`);
  else await page.click(`${ITEM} > summary`);
  await page.run(wait(120));
  got.opening = await page.run(read);
  await page.run(wait(580));
  got.opened = await page.run(read);
  got.interpolateSize = await page.run(`Promise.resolve(CSS.supports('interpolate-size: allow-keywords'))`);
  const src = readFileSync(new URL('./ring.js', import.meta.url), 'utf8');
  for (const name of ['held', 'held-visible']) {
    await page.go(`http://127.0.0.1:${server.address().port}/${name}.html`);
    await page.run(wait(300));
    got[name] = await page.run(`(async () => {\n${src}\nreturn held(${JSON.stringify(await page.shot())});\n})()`);
  }
  if (BROWSER !== 'ios') {
    const TAB = BROWSER === 'safari' ? 'Alt+Tab' : 'Tab';
    got.rings = {};
    for (const width of [1200, 390]) {
      await page.size(width, 900);
      for (const [name, target] of Object.entries(RINGS)) {
        await page.go(`http://127.0.0.1:${server.address().port}/ring.html`);
        await page.run(wait(300));
        let focused = false;
        for (let i = 0; i < 12 && !focused; i++) {
          await page.press(TAB);
          focused = await page.run(`Promise.resolve(document.activeElement === document.querySelector(${JSON.stringify(target)}))`);
        }
        await page.run(wait(300));
        // read again, up to three times, if the window was behind as it was read
        let r = { focused };
        for (let tries = 0; focused && tries < 3 && !r.hasFocus; tries++) {
          got.waited[`${name} at ${width}, ${tries + 1}`] = await front();
          r = await page.run(`(async () => {\n${src}\nreturn ring(${JSON.stringify(target)}, ${JSON.stringify(await page.shot())});\n})()`);
        }
        got.rings[`${name} at ${width}`] = r;
      }
    }
  }
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
claims.push(['its height clipped as it moves: an answer held 10px tall draws none of its text, which one clipping nothing draws',
  got.held.height === '10px' && got.held.ink === 0 && got['held-visible'].height === '10px' && got['held-visible'].ink > 0]);
// a side shown: at least a column of the screen's pixels the element's
// height in the accent, beside it
const shown = (r, side) => r.focused && r.visible && r[side] >= r.height * r.scale;
for (const [key, r] of Object.entries(got.rings ?? {})) {
  const [name, , width] = key.split(' ');
  claims.push([`the ${name === 'question' ? 'compact question in the large one\'s answer' : 'link beginning an answer'}, focused by the keyboard at ${width}px: its ring's left side shown${name === 'question' ? ', and its right' : ''}`,
    shown(r, 'left') && (name !== 'question' || shown(r, 'right'))]);
}
console.log(page.label);
let pass = true;
for (const [claim, ok] of claims) {
  pass &&= ok;
  console.log(`${ok ? 'ok  ' : 'FAIL'} ${claim}`);
}
process.exit(pass ? 0 : 1);
