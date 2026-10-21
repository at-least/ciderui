// BROWSER=safari|ios|chrome|chromium node run.mjs
// Can a <dialog> be the store's overlay with no script, as modal.css makes
// it one? Serves page.html on 127.0.0.1 (the browsers: ../browsers.mjs),
// clicks and presses, and reads (dialog.js): the page's button opens it as a
// modal dialog (command="show-modal", no script), the layer over the window,
// the page dimmed under it (::backdrop) and held, the layer scrolling the
// panel taller than the window, a point above the panel hitting the layer,
// the focus on its close button (its first control, as a dialog's opening
// gives it: Chromium 151 does so with autofocus or tabindex on the dialog
// too), fading in (@starting-style,
// sampled each frame from just before the click); its close button, a form's
// (method="dialog"), closes it, the focus back on the button that opened it;
// Escape closes it, the same. Clicked as ../popover/ clicks: by the page's
// script on the iOS simulator (its taps never click there), by the keyboard
// in Safari on the Mac (a button focused, Return), by Playwright elsewhere;
// on the simulator no key, and no focus to give back (a script's click, as a
// tap there, focuses no button). And the marketing pages' modal, as
// modal.css's page variant makes it: its ::backdrop's blur, the page blurred
// and dimmed under it in a picture of the window, its scrim coming in
// (sampled each frame), its close button stuck at the window's top as the
// panel scrolls, closed by it; and its dark one (color-scheme: dark on the
// dialog), its panel and close button from the root's light-dark() pairs, its
// panel black in a picture. The result is results/<browser>.json.
import { createServer } from 'node:http';
import { mkdirSync, readFileSync, writeFileSync } from 'node:fs';
import { open } from '../browsers.mjs';

const BROWSER = process.env.BROWSER;
const html = readFileSync(new URL('./page.html', import.meta.url));
const server = createServer((req, res) => { res.writeHead(200, { 'content-type': 'text/html; charset=utf-8' }); res.end(html); });
await new Promise(r => server.listen(0, '127.0.0.1', r));
const src = readFileSync(new URL('./dialog.js', import.meta.url), 'utf8');
const call = (fn) => `(async () => {\n${src}\nreturn ${fn};\n})()`;
const wait = (ms) => `new Promise(r => setTimeout(() => r(true), ${ms}))`;
let page;
const got = {};
const IOS = BROWSER === 'ios';
const KEYBOARD = BROWSER === 'safari';
const script = (selector, call) => page.run(`Promise.resolve((document.querySelector(${JSON.stringify(selector)}).${call}(), true))`);
const click = (selector) => IOS ? script(selector, 'click') : KEYBOARD ? script(selector, 'focus').then(() => page.press('Enter')) : page.click(selector);
try {
  page = await open(BROWSER, `http://127.0.0.1:${server.address().port}/`);
  if (!IOS) await page.size(1440, 900);
  await page.run(wait(300));
  got.pageClosedPixels = await page.run(call(`pixels(${JSON.stringify(await page.shot())})`));
  got.closed = await page.run(call('state()'));
  await page.run(call('watch()'));
  await click('#button');
  await page.run(wait(700));
  got.opened = await page.run(call('state()'));
  await click('#close');
  await page.run(wait(300));
  got.closedAgain = await page.run(call('state()'));
  if (!IOS) {
    await click('#button');
    await page.run(wait(500));
    await page.press('Escape');
    await page.run(wait(300));
    got.afterEscape = await page.run(call('state()'));
  }
  got.samples = await page.run('Promise.resolve(window.cuiSamples)');
  // the marketing modal: opened, its scrim sampled each frame, a picture;
  // made taller than the window and scrolled 400px; closed by its button
  await page.run(call('watchPage()'));
  await click('#page-button');
  await page.run(wait(1300));
  got.page = await page.run(call('pageState()'));
  got.pageColors = await page.run(call(`colors('page')`));
  got.pageOpenPixels = await page.run(call(`pixels(${JSON.stringify(await page.shot())}, ${got.pageClosedPixels.offset})`));
  got.pageFrames = await page.run('Promise.resolve(window.cuiPage)');
  await page.run(`Promise.resolve((document.querySelector('#page .panel').append(Object.assign(document.createElement('div'), { className: 'tall' })), document.querySelector('#page').scrollTop = 400, true))`);
  await page.run(wait(300));
  got.pageScrolled = await page.run(call('pageState()'));
  await click('#page-close');
  await page.run(wait(300));
  got.pageClosed = await page.run(call('pageState()'));
  // its dark one: opened, its colors, a picture; closed by its button
  await click('#dark-button');
  await page.run(wait(1300));
  got.dark = await page.run(call(`colors('dark')`));
  got.darkPixels = await page.run(call(`pixels(${JSON.stringify(await page.shot())}, ${got.pageClosedPixels.offset})`));
  await click('#dark-close');
  await page.run(wait(300));
  got.darkClosed = await page.run(call(`colors('dark')`));
} finally {
  await page?.close();
  server.close();
}
mkdirSync(new URL('./results/', import.meta.url), { recursive: true });
writeFileSync(new URL(`./results/${BROWSER}.json`, import.meta.url), JSON.stringify({ browser: page.label, date: new Date().toISOString(), ...got }, null, 1));

const o = got.opened;
const frames = got.samples[0].frames.filter(f => f[3]);
const claims = [
  ['the browser has the command attribute', got.closed.command === true],
  ['the page\'s button opens it as a modal dialog, with no script', o.open && o.modal],
  ['it is the layer over the window', o.position === 'fixed' && o.top === 0 && o.left === 0 && o.width === o.windowWidth && o.height === o.windowHeight],
  ['the page is dimmed under it', o.backdrop === 'rgba(0, 0, 0, 0.48)'],
  ['the page is held while it is open', o.rootOverflow === 'hidden' && got.closed.rootOverflow !== 'hidden' && got.closedAgain.rootOverflow !== 'hidden'],
  ['the layer scrolls the panel taller than the window', o.scrolls === true],
  ['a point above the panel hits the layer, not the page', o.hit === 'modal'],
  ['the focus is on its close button, the dialog\'s first control', o.focus === 'close'],
  ['it fades in: an open frame under 1, then 1', frames.some(f => f[1] < 1) && o.open && frames.at(-1)[1] === 1],
  ['its close button closes it', got.closedAgain.open === false],
  ...(IOS ? [] : [
    ['the focus is back on the button that opened it', got.closedAgain.focus === 'button'],
    ['Escape closes it, the focus back on the button', got.afterEscape.open === false && got.afterEscape.focus === 'button'],
  ]),
  ['the marketing modal\'s button opens it, its ::backdrop blur(20px)', got.page.modal === true && got.page.filter === 'blur(20px)'],
  ['its panel centered: as far from the window\'s top as from its foot', Math.abs(got.page.panel - got.page.below) <= 1 && got.page.panel > 28],
  ['in the picture the page under it is blurred and dimmed (the stripes\' spread under a fifth of theirs closed, their brightness 30 lower)', got.pageClosedPixels.spread > 60 && got.pageOpenPixels.spread < got.pageClosedPixels.spread / 5 && got.pageOpenPixels.mean < got.pageClosedPixels.mean - 30],
  ['its scrim comes in: a frame whose blur is between 0 and 20px', got.pageFrames.some(([, f]) => { const v = parseFloat(f.match(/blur\(([\d.]+)px\)/)?.[1]); return v > 0 && v < 20; })],
  ['scrolled 400px, its close button stays 20px from the window\'s top (sticky, its top the layer\'s room taken back)', got.pageScrolled.close === 20 && got.pageScrolled.panel < 0],
  ['its close button closes it', got.pageClosed.open === false],
  ['its dark one (color-scheme: dark on the dialog) opens, its panel black, its text #f5f5f7, its close button #f5f5f7 with a black X, from the root\'s light-dark() pairs; the light one\'s panel white, its text #1d1d1f, from the same', got.dark.modal === true && got.dark.scheme === 'dark' && got.dark.panel === 'rgb(0, 0, 0)' && got.dark.text === 'rgb(245, 245, 247)' && got.dark.close === 'rgb(245, 245, 247)' && got.dark.x === 'rgb(0, 0, 0)' && got.pageColors.scheme === 'light' && got.pageColors.panel === 'rgb(255, 255, 255)' && got.pageColors.text === 'rgb(29, 29, 31)'],
  ['its scrim the light one\'s', got.dark.backdrop === 'rgba(0, 0, 0, 0.48)' && got.pageColors.backdrop === got.dark.backdrop],
  ['in the picture its panel is black, the light one\'s white', got.darkPixels.panel.every(v => v <= 2) && got.pageOpenPixels.panel.every(v => v >= 253)],
  ['its close button closes it', got.darkClosed.open === false],
];
console.log(page.label + (IOS ? ' (clicked by script; no Escape)' : KEYBOARD ? ' (by its keyboard: Return)' : ''));
let pass = true;
for (const [claim, ok] of claims) {
  pass &&= ok;
  console.log(`${ok ? 'ok  ' : 'FAIL'} ${claim}`);
}
if (!pass) console.log(JSON.stringify({ ...got, samples: got.samples.map(s => s.frames.slice(0, 10)) }, null, 1));
console.log(pass ? 'PASS' : 'FAIL');
process.exit(pass ? 0 : 1);
