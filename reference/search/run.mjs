// BROWSER=safari|ios|chrome|chromium|firefox node run.mjs
// Can a popover be the global nav's search with no script, as globalnav.css
// makes it one? Serves page.html on 127.0.0.1 (the browsers: ../browsers.mjs;
// ?sheet the same page as a sheet), clicks and presses, and reads (search.js)
// and looks (pictures of the window): the bar's search button, inside the
// bar's closed list, opens the flyout; the focus on its field (autofocus);
// the flyout placed from the bar's bottom by anchor positioning, under a
// banner and with the page scrolled and the bar stuck at the top; under it
// the curtain to the window's bottom, blurring the page, the bar not under
// it; a click on the bar's link closes it and reaches the link, one on the
// curtain closes it and reaches nothing under it; Escape closes it, the
// focus back on the button; the button closes it again; Return in the field
// sends the form, GET, to the site's search page. As a sheet: over the
// window and the bar, the page held, its close button closing it. Its
// field's clear button, a reset button: hidden while the field is empty,
// shown with text in it; clicked, the field empty and the button hidden, the
// flyout still open, Shift+Tab back to the field, Escape closing it. Clicked
// as ../popover/ clicks: by the page's script on the iOS simulator (its taps
// never click there; no key, no click outside, no focus to give back), by
// the keyboard in Safari on the Mac for the buttons (a button focused,
// Return), and by the browser's own click elsewhere (the bar's link and the
// curtain in Safari too). The result is results/<browser>.json.
import { createServer } from 'node:http';
import { mkdirSync, readFileSync, writeFileSync } from 'node:fs';
import { open, wait } from '../browsers.mjs';

const BROWSER = process.env.BROWSER;
const html = readFileSync(new URL('./page.html', import.meta.url), 'utf8');
const server = createServer((req, res) => {
  res.writeHead(200, { 'content-type': 'text/html; charset=utf-8' });
  res.end(req.url.includes('sheet') ? html.replace('<html lang="en">', '<html lang="en" class="sheet">') : html);
});
await new Promise(r => server.listen(0, '127.0.0.1', r));
const origin = `http://127.0.0.1:${server.address().port}`;
const src = readFileSync(new URL('./search.js', import.meta.url), 'utf8');
const call = (fn) => `(async () => {\n${src}\nreturn ${fn};\n})()`;
let page;
const got = {};
const IOS = BROWSER === 'ios';
const KEYBOARD = BROWSER === 'safari';
// a button: by script on the simulator, by Return in Safari, else clicked
const press = (selector) => IOS ? page.script(selector, 'click') : KEYBOARD ? page.script(selector, 'focus').then(() => page.press('Enter')) : page.click(selector);
// a click anywhere else: the browser's own, but on the simulator
const click = (selector) => IOS ? page.script(selector, 'click') : page.click(selector);
const state = () => page.run(call('state()'));
const pixels = async (offset) => page.run(call(`pixels(${JSON.stringify(await page.shot())}, ${offset})`));
const opened = async () => { await press('#search'); await page.run(wait(900)); };
try {
  page = await open(BROWSER, `${origin}/`);
  if (!IOS) await page.size(1440, 900);
  await page.run(wait(500));
  await page.run(call('arm()'));
  got.closed = await state();
  got.closedPixels = await pixels();
  await opened();
  got.opened = await state();
  got.openedPixels = await pixels(got.closedPixels.offset);
  await page.run(`Promise.resolve((scrollTo(0, 200), true))`);
  await page.run(wait(400));
  got.scrolled = await state();
  await page.run(`Promise.resolve((scrollTo(0, 0), true))`);
  await page.run(wait(400));
  if (!IOS) {
    await click('#link');
    await page.run(wait(400));
    got.afterLink = await state();
    await opened();
  }
  await click('#curtain');
  await page.run(wait(400));
  got.afterCurtain = await state();
  if (!IOS) {
    await opened();
    await page.press('Escape');
    await page.run(wait(400));
    got.afterEscape = await state();
  }
  await opened();
  await press('#search');
  await page.run(wait(400));
  got.afterToggle = await state();
  // text in its field (set by the page's script: what the CSS reads is the
  // value), then its clear button clicked
  await opened();
  got.empty = await state();
  await page.run(`Promise.resolve((document.querySelector('#input').value = 'iphone', true))`);
  await page.run(wait(400));
  got.typed = await state();
  await click('#reset');
  await page.run(wait(400));
  got.cleared = await state();
  if (!IOS) {
    // the focus's way back to the field: Shift+Tab, from where it was
    await page.press('Shift+Tab');
    await page.run(wait(200));
    got.clearedBack = await state();
    await page.press('Escape');
    await page.run(wait(400));
    got.clearedEscape = await state();
  }
  if (!IOS) {
    await opened();
    await page.run(`Promise.resolve((document.querySelector('#input').value = 'iphone', true))`);
    await page.press('Enter');
    // the page it loads: waited for here, not in the page it leaves
    await new Promise(r => setTimeout(r, 1500));
    got.sent = await state();
  }
  await page.go(`${origin}/?sheet`);
  await page.run(wait(500));
  got.sheetClosed = await state();
  await opened();
  got.sheet = await state();
  await press('#close');
  await page.run(wait(400));
  got.sheetAfterClose = await state();
} finally {
  await page?.close();
  server.close();
}
mkdirSync(new URL('./results/', import.meta.url), { recursive: true });
writeFileSync(new URL(`./results/${BROWSER}.json`, import.meta.url), JSON.stringify({ browser: page.label, date: new Date().toISOString(), ...got }, null, 1));

const o = got.opened, p = got.openedPixels, s = got.sheet;
const near = (a, b, t = 0.5) => Math.abs(a - b) <= t;
const color = (c, want, t = 2) => c.every((v, i) => Math.abs(v - want[i]) <= t);
const claims = [
  ['the browser has anchor positioning', got.closed.anchors === true],
  ['closed, the bar\'s list is its row and the flyout is not shown', got.closed.menuDisplay === 'flex' && got.closed.open === false && got.closed.display === 'none'],
  ['the search button, in the bar\'s closed list, opens the flyout, with no script', o.open === true],
  ['the focus is on its field (autofocus)', o.focus === 'input'],
  ['it hangs from the bar\'s bottom, under the banner, across the window', near(o.flyout.top, o.bar.bottom) && near(o.bar.bottom, 114) && o.flyout.left === 0 && o.flyout.width === o.pageWidth],
  ['its curtain fills the window under its panel', near(o.curtain.top, o.panel.bottom) && near(o.curtain.bottom, o.windowHeight)],
  ['the page scrolled, the bar stuck at the top, the flyout follows it', near(got.scrolled.bar.top, 0) && near(got.scrolled.flyout.top, got.scrolled.bar.bottom) && got.scrolled.open === true],
  ['in the picture the bar is its own color, not under the curtain, and the panel #fafafc', color(p.bar, [200, 220, 240]) && color(got.closedPixels.bar, [200, 220, 240]) && color(p.panel, [250, 250, 252])],
  ['in the picture the curtain blurs the page under it (the stripes\' spread under a fifth of theirs closed)', got.closedPixels.spread > 60 && p.spread < got.closedPixels.spread / 5],
  ['the curtain\'s middle hits the curtain, not the page; the bar hits its link', o.curtainHit === 'curtain' && o.barHit === 'link'],
  ['a popover\'s ::backdrop takes no click (pointer-events: none), so the curtain is a button of its own', o.backdropPointer === 'none'],
  ...(IOS ? [] : [['a click on the bar\'s link closes it and reaches the link', got.afterLink.open === false && got.afterLink.clicks.link === 1]]),
  [`a click on the curtain closes it${IOS ? ' (by script)' : ''} and reaches nothing under it`, got.afterCurtain.open === false && got.afterCurtain.clicks.under === 0],
  ...(IOS ? [] : [['Escape closes it, the focus back on the search button', got.afterEscape.open === false && got.afterEscape.focus === 'search']]),
  ['its button closes it again', got.afterToggle.open === false],
  ...(IOS ? [] : [['Return in its field sends the form to the site\'s search page (GET, ?q=)', got.sent.at === '/search?q=iphone']]),
  ['as a sheet its button opens it over the window and the bar', s.open === true && s.flyout.top === 0 && s.flyout.left === 0 && s.flyout.width === s.pageWidth && near(s.flyout.height, s.windowHeight) && s.barHit !== 'link'],
  ['as a sheet the focus is on its field', s.focus === 'input'],
  ['as a sheet the page is held while it is open, and not after', s.rootOverflow === 'hidden' && got.sheetClosed.rootOverflow !== 'hidden' && got.sheetAfterClose.rootOverflow !== 'hidden'],
  ['as a sheet its close button closes it', got.sheetAfterClose.open === false],
  ...(IOS ? [] : [['as a sheet the focus is back on the search button', got.sheetAfterClose.focus === 'search']]),
  ['empty, its field\'s clear button is hidden', got.empty.reset === 'hidden' && got.empty.value === ''],
  ['with text in the field, its clear button shows, with no script', got.typed.reset === 'visible'],
  [`its clear button (a reset button) clicked${IOS ? ' (by script)' : ''} empties the field and hides itself, the flyout still open`, got.cleared.value === '' && got.cleared.reset === 'hidden' && got.cleared.open === true],
  ...(IOS ? [] : [
    ['then Shift+Tab takes the focus back to the field', got.clearedBack.focus === 'input'],
    ['and Escape still closes it', got.clearedEscape.open === false],
  ]),
];
console.log(page.label + (IOS ? ' (clicked by script; no key, no click outside)' : KEYBOARD ? ' (its buttons by its keyboard: Return)' : ''));
let pass = true;
for (const [claim, ok] of claims) {
  pass &&= ok;
  console.log(`${ok ? 'ok  ' : 'FAIL'} ${claim}`);
}
if (!pass) console.log(JSON.stringify(got, null, 1));
console.log(pass ? 'PASS' : 'FAIL');
process.exit(pass ? 0 : 1);
