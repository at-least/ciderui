// BROWSER=safari|chrome|chromium|ipad|iphone node run.mjs
// Can the global nav's flyouts open with no script, as globalnav.css opens
// them? Serves page.html on 127.0.0.1. In Safari, Chrome and Chromium
// (../browsers.mjs), at 1440 by 900: the pointer moved as a hand moves it
// (`move`, the browser's own input) to a link, along the bar to the next, off
// the bar to the page, and swept along the bar quicker than the flyouts
// open, each recording which flyouts show every frame (page.html's
// probe.watch()); pictures of the window (the curtain's blur, the bar's
// ground); the link's button by the keyboard in Safari (focused, Return; its
// Tab Option-Tab, which reaches links) and clicked elsewhere, Tab, Escape,
// the curtain clicked; then at 700px, a phone's layout: the menu, a row's
// flyout over it, its back button, its close button. On the iOS simulator
// XCUITest's taps (../xcuitest.mjs), each reported by the page 800ms on: on
// an iPad (`ipad`, 13-inch, a screen wider than 834px with no hover) a link
// tapped, its flyout, then its curtain; on an iPhone (`iphone`) the menu, a
// row, the back button, the row again, the close button. The result is
// results/<browser>.json; it fails if a claim does not hold.
import { createServer } from 'node:http';
import { mkdirSync, readFileSync, writeFileSync } from 'node:fs';
import { open } from '../browsers.mjs';
import { taps } from '../xcuitest.mjs';

const BROWSER = process.env.BROWSER;
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
const got = {};

async function desktop() {
  const page = await open(BROWSER, url);
  // a button by the keyboard in Safari (a click there focuses none), else clicked
  const KEYBOARD = BROWSER === 'safari';
  const run = (expr) => page.run(expr);
  const wait = (ms) => run(`new Promise(r => setTimeout(() => r(true), ${ms}))`);
  const state = () => run('Promise.resolve(window.probe.state())');
  const watch = () => run('Promise.resolve(window.probe.watch())');
  const stop = () => run('Promise.resolve(window.probe.stop())');
  const pixels = async (offset) => run(`window.probe.pixels(${JSON.stringify(await page.shot())}${offset === undefined ? '' : `, ${offset}`})`);
  const mid = (selector) => run(`Promise.resolve((r => [r.left + r.width / 2, r.top + r.height / 2])(document.querySelector(${JSON.stringify(selector)}).getBoundingClientRect()))`);
  const press = async (selector) => {
    if (!KEYBOARD) return page.click(selector);
    await run(`Promise.resolve((document.querySelector(${JSON.stringify(selector)}).focus(), true))`);
    await page.press('Enter');
  };
  try {
    await page.size(1440, 900);
    await wait(800);
    // a point on the page near the window's foot (Safari's window, its frame
    // included, is 900px tall: its page is less)
    const foot = await run('Promise.resolve(innerHeight - 50)');
    await page.move(700, foot, 100);
    await wait(500);
    got.closed = await state();
    got.closedPixels = await pixels();
    const offset = got.closedPixels.offset;
    // the pointer to Store's link, resting there
    await watch();
    await page.move(...await mid('#store'), 150);
    await wait(800);
    got.hovered = await state();
    got.hoveredWatch = await stop();
    got.hoveredPixels = await pixels(offset);
    // on along the bar to Mac's
    await watch();
    await page.move(...await mid('#mac'), 150);
    await wait(900);
    got.switched = await state();
    got.switchedWatch = await stop();
    // off the bar, through Mac's flyout, down to the page
    await watch();
    await page.move(700, foot, 150);
    await wait(900);
    got.left = await state();
    got.leftWatch = await stop();
    // swept along the bar past each link in 90ms, from before the first to
    // after the last, then down to the page
    const [x0, y0] = await mid('#store');
    const [x1] = await mid('#search');
    await page.move(x0 - 40, y0, 100);
    await wait(500);
    await watch();
    await page.move(x1 + 60, y0, 90);
    await page.move(x1 + 60, foot, 50);
    await wait(800);
    got.swept = await state();
    got.sweptWatch = await stop();
    // Mac's button: its flyout a popover
    await press('#mac-button');
    await wait(600);
    got.popover = await state();
    await page.press(KEYBOARD ? 'Alt+Tab' : 'Tab');
    await wait(200);
    got.tabbed = await state();
    await page.press('Escape');
    await wait(400);
    got.escaped = await state();
    await press('#mac-button');
    await wait(600);
    await page.click('#mac-curtain');
    await wait(400);
    got.curtained = await state();
    // a phone's layout
    await page.move(10, foot, 50);
    await page.size(700, 900);
    await wait(600);
    await press('#menu-button');
    await wait(600);
    got.phoneMenu = await state();
    await press('#mac-button');
    await wait(600);
    got.phoneFlyout = await state();
    await press('#mac-back');
    await wait(600);
    got.phoneBack = await state();
    await press('#mac-button');
    await wait(600);
    await press('#mac-close');
    await wait(600);
    got.phoneClosed = await state();
    return page.label;
  } finally {
    await page.close();
  }
}

let label;
let tip;
try {
  if (BROWSER === 'ipad') ({ label, tip } = await taps(url, ['Store menu', 'Store curtain'], { device: 'iPad-Pro-13-inch-M5-12GB', name: 'iPad Pro 13-inch (M5)' }));
  else if (BROWSER === 'iphone') ({ label, tip } = await taps(url, ['Menu', 'Mac menu', 'Back', 'Mac menu', 'Close all']));
  else label = await desktop();
} finally {
  server.close();
}
mkdirSync(new URL('./results/', import.meta.url), { recursive: true });
writeFileSync(new URL(`./results/${BROWSER}.json`, import.meta.url), JSON.stringify({ browser: label, date: new Date().toISOString(), ...(tip ? { tip } : {}), ...got, ...(reports.length && !got.closed ? { reports } : {}) }, null, 1));

const near = (a, b, t = 0.5) => Math.abs(a - b) <= t;
const color = (c, want, t = 2) => c.every((v, i) => Math.abs(v - want[i]) <= t);
// the first frame a flyout shows, from the pointer's entering its item
const shownAfter = (w, id) => {
  const entered = w.entered.find(([e]) => e === id)?.[1];
  const shown = w.frames.find(([, ids]) => ids.includes(id))?.[0];
  return entered === undefined || shown === undefined ? null : shown - entered;
};
let claims;
if (BROWSER === 'ipad') {
  const [opened, curtained] = reports;
  claims = [
    ['the iPad\'s screen has no hover: (hover: hover) does not match', reports.every(r => !r.media.hover)],
    ['a tap on Store\'s link (its button over it) opens its flyout as a popover, from the bar\'s bottom', opened.store.open && near(opened.store.box.top, opened.bar.bottom)],
    ['the link is not followed', opened.hash === '' && opened.clicks.followed === 0],
    ['a tap on its curtain closes it and reaches nothing under it', !curtained.store.open && curtained.clicks.under === 0 && curtained.hash === ''],
  ];
} else if (BROWSER === 'iphone') {
  const [menu, flyout, back, , closed] = reports;
  claims = [
    ['the menu\'s button opens the menu', menu.menu && !menu.mac.open],
    ['a tap on Mac\'s row (its button over it) opens Mac\'s flyout over the open menu, across the window', flyout.menu && flyout.mac.open && flyout.mac.box.top === 0 && near(flyout.mac.box.height, flyout.window.height) && flyout.hash === ''],
    ['its back button closes the flyout, the menu still open', back.menu && !back.mac.open],
    ['its close button closes both', !closed.menu && !closed.mac.open],
  ];
} else {
  const { closed, hovered, switched, left, swept, popover, tabbed, escaped, curtained } = got;
  const p = got.hoveredPixels;
  const after = shownAfter(got.hoveredWatch, 'store');
  claims = [
    ['the browser has anchor positioning and a pointer that hovers', closed.media.anchors && closed.media.hover],
    ['closed, no flyout shows, the curtain is hidden, the bar blurred', !closed.store.shown && !closed.mac.shown && closed.curtain.visibility === 'hidden' && closed.barFilter !== 'none'],
    ['closed, each flyout is the bar\'s height, nothing of it laid out below the bar', near(closed.store.box.height, closed.bar.height) && near(closed.mac.box.height, closed.bar.height)],
    [`the pointer resting on Store's link shows its flyout, no popover opened, 100 to 250ms after it reached the link (${after}ms)`, hovered.store.shown && !hovered.store.open && after >= 100 && after <= 250],
    ['it hangs from the bar\'s top, across the window', near(hovered.store.box.top, hovered.bar.top) && hovered.store.box.left === 0 && near(hovered.store.box.width, hovered.window.page)],
    ['the bar\'s links are over it', hovered.macHit === 'mac'],
    ['the curtain shows, fixed, and the bar\'s blur is off', hovered.curtain.visibility === 'visible' && hovered.curtain.opacity === '1' && hovered.curtain.position === 'fixed' && hovered.barFilter === 'none'],
    ['in the picture the curtain blurs the page under it (the stripes\' spread under a fifth of theirs closed), and the bar shows the flyout\'s ground', got.closedPixels.spread > 60 && p.spread < got.closedPixels.spread / 5 && color(p.bar, [250, 250, 252])],
    ['the pointer moved on to Mac\'s link shows Mac\'s flyout and not Store\'s', switched.mac.shown && !switched.store.shown],
    ['and in every frame between, a flyout shows to its full height (the one left kept until the next is shown)', got.switchedWatch.frames.every(([, ids, reach]) => ids.length > 0 && reach >= switched.mac.panel.bottom - 0.5)],
    ['the pointer off the bar to the page hides it, the curtain hidden, the bar blurred again', !left.store.shown && !left.mac.shown && left.curtain.visibility === 'hidden' && left.barFilter !== 'none'],
    ['and it is the bar\'s height again', near(left.mac.box.height, left.bar.height)],
    [`swept along the bar past its links (${got.sweptWatch.entered.map(([id]) => id).join(', ')} entered), no flyout shows in any frame`, got.sweptWatch.entered.length >= 2 && got.sweptWatch.frames.every(([, ids]) => ids.length === 0) && !swept.store.shown && !swept.mac.shown],
    ['Mac\'s button opens its flyout as a popover, from the bar\'s bottom, its curtain to the window\'s bottom', popover.mac.open && near(popover.mac.box.top, popover.bar.bottom) && near(popover.mac.curtain.bottom, popover.window.height) && popover.mac.curtainOpacity === '1'],
    ['Tab takes the focus to its first link', tabbed.focus === 'mac-first'],
    ['Escape closes it, the focus back on its button', !escaped.mac.open && escaped.focus === 'mac-button'],
    ['a click on its curtain closes it and reaches nothing under it', !curtained.mac.open && curtained.clicks.under === 0],
    ['at a phone\'s width the menu\'s button opens the menu', got.phoneMenu.menu && !got.phoneMenu.mac.open],
    ['Mac\'s row\'s button opens Mac\'s flyout over the open menu, across the window', got.phoneFlyout.menu && got.phoneFlyout.mac.open && got.phoneFlyout.mac.box.top === 0 && near(got.phoneFlyout.mac.box.height, got.phoneFlyout.window.height)],
    ['its back button closes the flyout, the menu still open', got.phoneBack.menu && !got.phoneBack.mac.open],
    ['its close button closes both', !got.phoneClosed.menu && !got.phoneClosed.mac.open],
  ];
}
console.log(label);
let pass = true;
for (const [claim, ok] of claims) {
  pass &&= ok;
  console.log(`${ok ? 'ok  ' : 'FAIL'} ${claim}`);
}
console.log(pass ? 'PASS' : 'FAIL');
process.exit(pass ? 0 : 1);
