// BROWSER=safari|ios|chrome|chromium|firefox node run.mjs: MEASURED.md holds its claims and method.
import { createServer } from 'node:http';
import { mkdirSync, readFileSync, writeFileSync } from 'node:fs';
import { open, wait } from '../browsers.mjs';

const BROWSER = process.env.BROWSER;
const html = readFileSync(new URL('./page.html', import.meta.url));
const server = createServer((req, res) => { res.writeHead(200, { 'content-type': 'text/html; charset=utf-8' }); res.end(html); });
await new Promise(r => server.listen(0, '127.0.0.1', r));
const src = readFileSync(new URL('./dialog.js', import.meta.url), 'utf8');
const call = (fn) => `(async () => {\n${src}\nreturn ${fn};\n})()`;
let page;
const got = {};
const IOS = BROWSER === 'ios';
const KEYBOARD = BROWSER === 'safari';
const click = (selector) => IOS ? page.script(selector, 'click') : KEYBOARD ? page.script(selector, 'focus').then(() => page.press('Enter')) : page.click(selector);
try {
  page = await open(BROWSER, `http://127.0.0.1:${server.address().port}/`);
  if (!IOS) await page.size(1440, 900);
  await page.run(wait(300));
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
  const opened = `Promise.resolve({ open: document.querySelector('#modal').open, hash: location.hash })`;
  await click('#form-button');
  await page.run(wait(500));
  got.formButton = await page.run(opened);
  await click('#close');
  await page.run(wait(300));
  await click('#form-submit');
  await page.run(wait(1000));
  got.formSubmit = await page.run(opened);
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
  ['a button in a form, its type button, opens it, the form not sent', got.formButton.open === true && got.formButton.hash === ''],
  ['a button in a form with no type does nothing, the dialog closed and the form not sent (the control)', got.formSubmit.open === false && got.formSubmit.hash === ''],
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
