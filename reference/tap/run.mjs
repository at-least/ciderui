// BROWSER=chromium|chrome|firefox|ios node run.mjs
// What does a tap leave? Serves page.html on 127.0.0.1 and taps, one after
// another, a button whose hover is a rule of its own (apple.com's), one whose
// hover is behind (hover: hover) (ciderui's, decision 5), somewhere else, the
// menu's button (a popover opened from beside it, ../popover/) and its close
// button; 800ms after each tap the page sends what it left (page.html's
// script: the events, the pseudo-classes and colors, the menu open or not).
// Each tap a touch screen's, never a script's click: in Chromium and Chrome
// (../browsers.mjs) the DevTools protocol's touch events on a touch screen
// emulated from the load, which Chromium's gesture recognizer makes a tap of;
// in Firefox, which has no DevTools protocol in Playwright, Playwright's
// touchscreen taps on a page with a touch screen from the load (hasTouch:
// Gecko's own touch emulation); on the iOS simulator XCUITest's taps (../xcuitest.mjs), touches of the
// simulator's screen through its HID, as a finger's, on an iPhone simulator
// of the run's own, created and deleted (the booted ones are others';
// Safari's tip over the page closed first, recorded as `tip`). The result is
// results/<browser>.json.
import { createServer } from 'node:http';
import { mkdirSync, readFileSync, writeFileSync } from 'node:fs';
import { open } from '../browsers.mjs';
import { taps } from '../xcuitest.mjs';

const BROWSER = process.env.BROWSER;
// what is tapped, in order: its id (the DevTools protocol's taps) and its
// label (XCUITest's)
const STEPS = [['plain', 'Plain'], ['gated', 'Gated'], ['elsewhere', 'Elsewhere'], ['menu-button', 'Menu'], ['menu-close', 'Close']];
const html = readFileSync(new URL('./page.html', import.meta.url));
const reports = [];
const server = createServer((req, res) => {
  if (req.method === 'POST' && req.url === '/report') {
    let body = '';
    req.on('data', (c) => { body += c; });
    req.on('end', () => { reports.push(JSON.parse(body)); res.writeHead(204); res.end(); });
    return;
  }
  res.writeHead(200, { 'content-type': 'text/html; charset=utf-8' });
  res.end(html);
});
await new Promise(r => server.listen(0, '127.0.0.1', r));
const url = `http://127.0.0.1:${server.address().port}/`;

// the page's report of the i-th tap
const reported = async (i, id) => {
  for (let t = 0; reports.length < i + 1; t++) {
    if (t === 100) throw new Error(`no report 10s after tapping ${id}`);
    await new Promise(r => setTimeout(r, 100));
  }
};

async function viaDevTools() {
  const page = await open(BROWSER, url);
  try {
    const pw = page.page;
    await pw.setViewportSize({ width: 390, height: 844 });
    const cdp = await pw.context().newCDPSession(pw);
    await cdp.send('Emulation.setTouchEmulationEnabled', { enabled: true, maxTouchPoints: 5 });
    await pw.reload({ waitUntil: 'load' });
    for (const [i, [id]] of STEPS.entries()) {
      const r = await pw.locator(`#${id}`).boundingBox();
      await cdp.send('Input.dispatchTouchEvent', { type: 'touchStart', touchPoints: [{ x: r.x + r.width / 2, y: r.y + r.height / 2 }] });
      await cdp.send('Input.dispatchTouchEvent', { type: 'touchEnd', touchPoints: [] });
      await reported(i, id);
    }
    return `${page.label}, the DevTools protocol's touch events`;
  } finally {
    await page.close();
  }
}

async function viaTouchscreen() {
  const page = await open(BROWSER, url, { viewport: { width: 390, height: 844 }, hasTouch: true });
  try {
    const pw = page.page;
    for (const [i, [id]] of STEPS.entries()) {
      const r = await pw.locator(`#${id}`).boundingBox();
      await pw.touchscreen.tap(r.x + r.width / 2, r.y + r.height / 2);
      await reported(i, id);
    }
    return `${page.label}, Playwright's touchscreen taps`;
  } finally {
    await page.close();
  }
}

let label;
let tip;
try {
  if (BROWSER === 'ios') ({ label, tip } = await taps(url, STEPS.map(([, label]) => label)));
  else if (BROWSER === 'firefox') label = await viaTouchscreen();
  else label = await viaDevTools();
} finally {
  server.close();
}
mkdirSync(new URL('./results/', import.meta.url), { recursive: true });
writeFileSync(new URL(`./results/${BROWSER}.json`, import.meta.url), JSON.stringify({ browser: label, date: new Date().toISOString(), ...(tip ? { tip } : {}), reports }, null, 1));

const [plain, gated, elsewhere, menu, close] = reports;
const on = (report, id) => report.events.filter(e => e.on === id);
const has = (events, type, pointer) => events.some(e => e.type === type && e.trusted && (pointer === undefined || e.pointer === pointer));
const tapped = on(plain, 'plain');
const claims = [
  ['a tap is a touch: its pointerdown and pointerup `touch` and trusted, then mouseover, mousedown, mouseup and a click on the button', has(tapped, 'pointerdown', 'touch') && has(tapped, 'pointerup', 'touch') && ['mouseover', 'mousedown', 'mouseup', 'click'].every(t => has(tapped, t))],
  ['the screen has no hover: (hover: hover) does not match, (pointer: coarse) does', reports.every(r => !r.media.hover && r.media.coarse)],
  ['the tapped button is left hovered (:hover 800ms on), and a hover rule of its own shows there (apple.com\'s way: #0076df)', plain.plain.hover && plain.plain.background === 'rgb(0, 118, 223)'],
  ['a hover behind (hover: hover) does not: the tapped button hovered, at its rest #0071e3 (ciderui\'s way)', gated.gated.hover && gated.gated.background === 'rgb(0, 113, 227)'],
  ['no focus ring after a tap: :focus-visible on neither button', !plain.plain.focusVisible && !gated.gated.focusVisible],
  ['the menu\'s button opens its popover by a tap', menu.after === 'menu-button' && menu.menuOpen === true],
  ['its close button, tapped, closes it', close.after === 'menu-close' && close.menuOpen === false],
];
console.log(label);
let pass = true;
for (const [claim, ok] of claims) {
  pass &&= ok;
  console.log(`${ok ? 'ok  ' : 'FAIL'} ${claim}`);
}
// what is recorded beside the claims: where the hover goes with the next tap,
// and the focus a tap gives
console.log(`seen: the plain button tapped, focused ${plain.plain.focus}; the gated one tapped, the plain one still hovered ${gated.plain.hover}, the gated one focused ${gated.gated.focus}; somewhere else tapped, the gated one still hovered ${elsewhere.gated.hover}`);
if (!pass) console.log(JSON.stringify(reports, null, 1));
console.log(pass ? 'PASS' : 'FAIL');
process.exit(pass ? 0 : 1);
