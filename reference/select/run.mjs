// BROWSER=safari|chrome|chromium|firefox|ios node run.mjs
// What of the global nav can be selected, with the bar's user-select: none
// (globalnav.css's, apple.com's #globalnav's: ../apple.com/, nav-select:*)?
// Serves page.html on 127.0.0.1 (?plain the same page, its bar selectable)
// and reads what each step leaves (the page's cuiState()): where the bar can
// be selected, Select All on the page (the platform's command key and A)
// takes the bar's text with the page's, and a double-click on a flyout's
// heading, the nav's text that is no link, selects it (a link's flyout's,
// then the search's); where it cannot, Select All takes the page's text
// alone and the double-clicks select nothing; the search button opens its
// flyout, the focus on its field; the field clicked, keys typed go into it,
// and Select All in it selects its text whole. Pressed as ../search/
// presses: the buttons by the keyboard in Safari on the Mac (focused,
// Return), clicked elsewhere. On the iOS simulator, XCUITest
// (../xcuitest.mjs), on a simulator of the run's own for each page, ?plain's
// and the page's: the search button tapped, "mac" typed by its keyboard,
// the link's flyout's button tapped (the search closed, its field left),
// then its heading held (a long press: one with the search's field focused
// selected nothing on either page, 2026-10-22), each reported by the page.
// The result is results/<browser>.json.
import { createServer } from 'node:http';
import { mkdirSync, readFileSync, writeFileSync } from 'node:fs';
import { open, wait } from '../browsers.mjs';
import { taps } from '../xcuitest.mjs';

const BROWSER = process.env.BROWSER;
const IOS = BROWSER === 'ios';
const html = readFileSync(new URL('./page.html', import.meta.url), 'utf8');
const reports = [];
const server = createServer((req, res) => {
  if (req.method === 'POST' && req.url === '/report') {
    let body = '';
    req.on('data', (c) => { body += c; });
    req.on('end', () => { reports.push(JSON.parse(body)); res.writeHead(204); res.end(); });
    return;
  }
  res.writeHead(200, { 'content-type': 'text/html; charset=utf-8' });
  res.end(req.url.includes('plain') ? html.replace('<html lang="en">', '<html lang="en" class="plain">') : html);
});
await new Promise(r => server.listen(0, '127.0.0.1', r));
const origin = `http://127.0.0.1:${server.address().port}`;

const got = {};
let label;
let tip;
try {
  if (IOS) {
    // the same steps on each page, each on a simulator of its own
    const steps = ['Search', 'type:mac', 'More', 'press:Explore'];
    const ios = async (path) => {
      const from = reports.length;
      try {
        return [await taps(`${origin}${path}`, steps), reports.slice(from)];
      } catch (e) {
        console.log(`reported before the failure: ${JSON.stringify(reports.slice(from), null, 1)}`);
        throw e;
      }
    };
    [, [got.plainOpened, got.plainTyped, got.plainMore, got.plainExplore]] = await ios('/?plain');
    [{ label, tip }, [got.opened, got.typed, got.more, got.explore]] = await ios('/');
  } else {
    const page = await open(BROWSER, `${origin}/?plain`);
    try {
      label = page.label;
      await page.size(1440, 900);
      const state = () => page.run('Promise.resolve(window.cuiState())');
      // a button: by Return in Safari, else clicked
      const press = (selector) => BROWSER === 'safari'
        ? page.run(`Promise.resolve((document.querySelector('${selector}').focus(), true))`).then(() => page.press('Enter'))
        : page.click(selector);
      const step = async (act, ms = 300) => { await act(); await page.run(wait(ms)); return state(); };
      const clear = () => page.run('Promise.resolve((getSelection().removeAllRanges(), true))');
      await page.run(wait(500));
      // the bar selectable
      got.plainAll = await step(() => page.press('ControlOrMeta+a'));
      await clear();
      await step(() => press('#more'), 600);
      got.plainExplore = await step(() => page.dblclick('#explore'));
      await step(() => press('#search'), 600);
      got.plainHeading = await step(() => page.dblclick('#heading'));
      // the bar not selectable
      await page.go(`${origin}/`);
      await page.run(wait(500));
      got.rest = await state();
      got.all = await step(() => page.press('ControlOrMeta+a'));
      await clear();
      got.more = await step(() => press('#more'), 600);
      got.explore = await step(() => page.dblclick('#explore'));
      got.opened = await step(() => press('#search'), 600);
      got.heading = await step(() => page.dblclick('#heading'));
      got.focused = await step(() => page.click('#input'));
      got.typed = await step(() => page.type('mac'));
      got.selectedAll = await step(() => page.press('ControlOrMeta+a'));
    } finally {
      await page.close();
    }
  }
} finally {
  server.close();
}
mkdirSync(new URL('./results/', import.meta.url), { recursive: true });
writeFileSync(new URL(`./results/${BROWSER}.json`, import.meta.url), JSON.stringify({ browser: label, date: new Date().toISOString(), ...(tip ? { tip } : {}), ...got }, null, 1));

const claims = IOS ? [
  ['the search button, tapped, opens its flyout, the focus on its field', got.opened.open === true && got.opened.focus === 'input'],
  ['XCUITest\'s keyboard types into the field ("mac"), the bar\'s user-select: none over it', got.typed.value === 'mac'],
  ['the link\'s flyout\'s button, tapped, closes the search and opens its flyout, the field left', [got.plainMore, got.more].every(r => r.open === false && r.moreOpen === true && r.focus !== 'input')],
  ['where the bar can be selected, a long press on the flyout\'s heading selects its text', got.plainExplore.selection.trim() === 'Explore'],
  ['where it cannot, the same long press selects nothing', got.explore.selection === ''],
] : [
  ['the bar computes user-select: none (-webkit-user-select where the browser has no other)', (got.rest.userSelect.bar[0] || got.rest.userSelect.bar[1]) === 'none'],
  ['where the bar can be selected, Select All on the page takes the bar\'s text with the page\'s', got.plainAll.selection.includes('Store') && got.plainAll.selection.includes('Selectable')],
  ['where it cannot, Select All takes the page\'s text alone', !got.all.selection.includes('Store') && got.all.selection.includes('Selectable')],
  ['where the bar can be selected, a double-click on a link\'s flyout\'s heading selects its text', got.plainExplore.selection.trim() === 'Explore'],
  ['where it cannot, the same double-click selects nothing', got.more.moreOpen === true && got.explore.selection === ''],
  ['the search button opens its flyout, the focus on its field', got.opened.open === true && got.opened.moreOpen === false && got.opened.focus === 'input'],
  ['where the bar can be selected, a double-click on the search\'s heading selects its text', got.plainHeading.selection.trim() === 'Shortcuts'],
  ['where it cannot, the same double-click selects nothing', got.heading.selection === ''],
  ['the field clicked, keys typed go into it ("mac")', got.focused.focus === 'input' && got.typed.value === 'mac'],
  ['Select All (the command key and A) in the field selects its text whole', got.selectedAll.fieldSelection[0] === 0 && got.selectedAll.fieldSelection[1] === 3],
];
console.log(label + (IOS ? '' : BROWSER === 'safari' ? ' (the search button by its keyboard: Return)' : ''));
let pass = true;
for (const [claim, ok] of claims) {
  pass &&= ok;
  console.log(`${ok ? 'ok  ' : 'FAIL'} ${claim}`);
}
const seen = got.rest ?? got.opened;
console.log(`seen: [user-select, -webkit-user-select] ${JSON.stringify(seen.userSelect)}; supported ${JSON.stringify(seen.supports)}`);
if (!pass) console.log(JSON.stringify(got, null, 1));
console.log(pass ? 'PASS' : 'FAIL');
process.exit(pass ? 0 : 1);
