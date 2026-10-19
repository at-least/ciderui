// BROWSER=safari|ios|chrome|chromium node run.mjs
// Does a nav's list work as its own menu, as globalnav.css makes it one?
// Serves page.html on 127.0.0.1 (the browsers: ../browsers.mjs), clicks the
// list's button with the browser's own input and reads (popover.js): shown
// closed by the page's CSS; opened by its button inside it into a sheet over
// the window, growing and its rows coming in (sampled each frame), the page
// held, the lines an X; closed by the same button at once (no transition back);
// opened and closed by Escape. On the iOS simulator the button is clicked by
// the page's script (its click(), the activation a tap's click runs) and no
// key is pressed: safaridriver's taps there reach the button as pointerdown
// and touchstart and never a click, its element click as nothing, and the
// phone has no Escape. The result is results/<browser>.json.
import { createServer } from 'node:http';
import { mkdirSync, readFileSync, writeFileSync } from 'node:fs';
import { open } from '../browsers.mjs';

const BROWSER = process.env.BROWSER;
const html = readFileSync(new URL('./page.html', import.meta.url));
const server = createServer((req, res) => { res.writeHead(200, { 'content-type': 'text/html; charset=utf-8' }); res.end(html); });
await new Promise(r => server.listen(0, '127.0.0.1', r));
const src = readFileSync(new URL('./popover.js', import.meta.url), 'utf8');
const call = (fn) => `(async () => {\n${src}\nreturn ${fn};\n})()`;
const wait = (ms) => `new Promise(r => setTimeout(() => r(true), ${ms}))`;
let page;
const got = {};
const IOS = BROWSER === 'ios';
const click = (selector) => IOS ? page.run(`Promise.resolve((document.querySelector(${JSON.stringify(selector)}).click(), true))`) : page.click(selector);
try {
  page = await open(BROWSER, `http://127.0.0.1:${server.address().port}/`);
  await page.run(call('arm()'));
  got.closed = await page.run(call('state()'));
  await click('#button');
  await page.run(wait(900));
  got.opened = await page.run(call('state()'));
  await click('#button');
  got.closedAgain = await page.run(call('state()'));
  if (!IOS) {
    await page.run(wait(800));
    await page.click('#button');
    await page.run(wait(600));
    await page.press('Escape');
    got.afterEscape = await page.run(call('state()'));
  }
  got.samples = await page.run('Promise.resolve(window.cuiSamples)');
} finally {
  await page?.close();
  server.close();
}
mkdirSync(new URL('./results/', import.meta.url), { recursive: true });
writeFileSync(new URL(`./results/${BROWSER}.json`, import.meta.url), JSON.stringify({ browser: page.label, date: new Date().toISOString(), ...got }, null, 1));

const near = (a, b, t = 0.05) => Math.abs(a - b) <= t;
const X = [[3.5, 3.5, 11.5, 11.5], [3.5, 3.5, 11.5, 11.5]];
const [open1, close1] = got.samples;
const claims = [
  ['closed, the page shows the list as its row', got.closed.open === false && got.closed.display === 'flex' && got.closed.position === 'static' && got.closed.height === 48],
  ['its button opens it over the window', got.opened.open === true && got.opened.position === 'fixed' && got.opened.top === 0 && got.opened.height === got.opened.windowHeight],
  ['it grows from the bar: a frame between 48px and the window', open1?.to === 'open' && open1.frames.some(([, h]) => h > 48.5 && h < got.opened.windowHeight - 0.5)],
  ['its rows come in from transparent', open1.frames[0][2] === 0 && got.opened.rowOpacity === 1],
  ['the page is held while it is open', got.opened.rootOverflow === 'hidden' && got.closedAgain.rootOverflow !== 'hidden'],
  ['the lines are an X, each 11.5 by 11.5 from 3.5, 3.5', got.opened.lines.every((l, i) => l.every((v, j) => near(v, X[i][j])))],
  ['the same button closes it at once', got.closedAgain.open === false && got.closedAgain.position === 'static' && got.closedAgain.height === 48 && close1?.to === 'closed' && close1.frames.every(([, h]) => h === 48)],
  ...(IOS ? [] : [['Escape closes it', got.afterEscape.open === false]]),
];
console.log(page.label + (IOS ? ' (clicked by script; no Escape)' : ''));
let pass = true;
for (const [claim, ok] of claims) {
  pass &&= ok;
  console.log(`${ok ? 'ok  ' : 'FAIL'} ${claim}`);
}
if (!pass) console.log(JSON.stringify({ ...got, samples: got.samples.map(s => ({ to: s.to, frames: s.frames.slice(0, 12) })) }, null, 1));
console.log(pass ? 'PASS' : 'FAIL');
process.exit(pass ? 0 : 1);
