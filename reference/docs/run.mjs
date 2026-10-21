// BROWSER=safari|ios|chrome|chromium SITE=<dir> node run.mjs
// Does ciderui's documentation (docs/, built with ciderui-zola) work in this
// browser as test/docs.mjs holds it to in Chromium? SITE is the site built
// for http://127.0.0.1:8737 (README.md), served there (the browsers:
// ../browsers.mjs). Every page: nothing scrolls sideways, every picture
// loads, and each example frame loads its stylesheet and its pictures, is no
// wider than the page and as tall as its content. Then what the frames do,
// clicked by the page's script (an element's click(), the activation a
// tap's click runs: safaridriver's taps on the iOS simulator never click):
// the theme's script moves a gallery's row (its Next paddle: the row
// scrolled, the previous paddle enabled, a later dot current) and a shelf's;
// an example's link leaves its frame where it is; the global nav's search
// opens in its frame by its button, its field focused, hung from the bar
// where the frame is 834px wide or more, a sheet where it is narrower,
// and its curtain or its X closes it; the dropdown's label the placeholder
// while it is empty and at its top once a choice is made, its text shown
// then, with no script; the textarea's box focused by its textarea's
// focus and its label floated, focused or filled, with no script; the
// site's search page (?q=modal) listing the pages its script finds in
// Zola's index, the modal's page first, its field holding the query; the
// global nav's search suggesting the modal's page first as "mo" is put in
// its field, its row muted and "Mo" marked, fading in; in a
// window narrower than
// 834px the global nav's menu opens over its frame, grown to the window's
// height under the page's local nav and scrolled there, and the local nav's
// menu below its bar, the frame grown to hold it. On a phone at its width;
// elsewhere in a window 1440px wide, then 700px (Safari's window, its frame
// included, so its page a little narrower). The result is
// results/<browser>.json; exits 1 if a claim does not hold.
import { createServer } from 'node:http';
import { existsSync, mkdirSync, readFileSync, readdirSync, statSync, writeFileSync } from 'node:fs';
import { extname, join, relative } from 'node:path';
import { open } from '../browsers.mjs';

const BROWSER = process.env.BROWSER;
const SITE = process.env.SITE;
if (!SITE || !existsSync(join(SITE, 'index.html'))) throw new Error('SITE=<the docs built for http://127.0.0.1:8737> (README.md)');
const ORIGIN = 'http://127.0.0.1:8737';
const TYPES = { '.html': 'text/html; charset=utf-8', '.css': 'text/css', '.js': 'text/javascript', '.woff2': 'font/woff2', '.png': 'image/png', '.svg': 'image/svg+xml', '.xml': 'application/xml', '.txt': 'text/plain', '.json': 'application/json' };
const server = createServer((req, res) => {
  let f = join(SITE, decodeURIComponent(new URL(req.url, ORIGIN).pathname));
  if (existsSync(f) && statSync(f).isDirectory()) f = join(f, 'index.html');
  if (!existsSync(f)) { res.writeHead(404); return res.end(); }
  res.writeHead(200, { 'content-type': TYPES[extname(f)] ?? 'application/octet-stream' });
  res.end(readFileSync(f));
});
await new Promise((r) => server.listen(8737, '127.0.0.1', r));
const pages = (dir) => readdirSync(dir).flatMap((f) => {
  const p = join(dir, f);
  return statSync(p).isDirectory() ? pages(p) : f === 'index.html' ? [p] : [];
});
const paths = pages(SITE).map((f) => `/${relative(SITE, f).replace(/index\.html$/, '')}`).sort();

const wait = (ms) => `new Promise((r) => setTimeout(() => r(true), ${ms}))`;
// a page as it settles: the fonts, the frames' sizes a frame after theirs
const INSPECT = `(async () => {
  await document.fonts.ready;
  await new Promise((r) => setTimeout(r, 700));
  const broken = (doc) => [...doc.images].filter((i) => !i.complete || !i.naturalWidth).map((i) => i.getAttribute('src').slice(0, 48));
  return {
    width: innerWidth,
    sideways: document.documentElement.scrollWidth - innerWidth,
    images: broken(document),
    frames: [...document.querySelectorAll('iframe.cz-example')].map((f) => {
      const d = f.contentDocument;
      const sheet = [...d.styleSheets].find((s) => s.href && s.href.includes('/example.css'));
      return { title: f.title, height: f.clientHeight, content: d.documentElement.scrollHeight, css: Boolean(sheet && sheet.cssRules.length), images: broken(d), width: f.getBoundingClientRect().width, page: document.documentElement.clientWidth };
    }),
  };
})()`;
// a frame's element, found by its title, and a call on an element in it
const frame = (title) => `document.querySelector('iframe[title=${JSON.stringify(title)}]')`;
const inFrame = (title, selector, call) => `Promise.resolve((${frame(title)}.contentDocument.querySelector(${JSON.stringify(selector)}).${call}(), true))`;

const got = { pages: {}, frames: {} };
const fails = [];
const claim = (what, ok, value) => { if (!ok) fails.push(`${what}: ${JSON.stringify(value)}`); return ok; };
let page;
const read = async (width) => {
  for (const path of paths) {
    await page.go(ORIGIN + path);
    const r = await page.run(INSPECT);
    got.pages[`${path} at ${width}`] = r;
    const where = `${path} at ${r.width}px`;
    claim(`${where}: nothing scrolls sideways`, r.sideways <= 0, r.sideways);
    claim(`${where}: every picture loads`, r.images.length === 0, r.images);
    for (const f of r.frames) {
      claim(`${where}, "${f.title}": its frame loads its stylesheet and its pictures`, f.css && f.images.length === 0, f);
      claim(`${where}, "${f.title}": its frame is no wider than the page`, f.width <= f.page, f);
      claim(`${where}, "${f.title}": its frame is as tall as its content`, f.height === f.content, f);
    }
  }
};
const behave = async (width) => {
  // a row's Next paddle, clicked: the row scrolled on, the previous paddle enabled
  for (const [path, title, row] of [['/components/gallery/', 'A gallery of cards', 'gallery'], ['/components/shelf/', 'A shelf with paddles', 'shelf']]) {
    await page.go(ORIGIN + path);
    await page.run(wait(700));
    await page.run(inFrame(title, `[aria-controls="${row}"][aria-label="Next"]`, 'click'));
    await page.run(wait(1200));
    const r = await page.run(`Promise.resolve((() => {
      const d = ${frame(title)}.contentDocument;
      const paddles = [...d.querySelectorAll('[aria-controls="${row}"]')];
      return { scroll: d.getElementById('${row}').scrollLeft, previous: paddles[0].disabled, shown: getComputedStyle(paddles[1]).display !== 'none', current: [...d.querySelectorAll('.cui-gallery-dot')].findIndex((a) => a.getAttribute('aria-current') === 'true') };
    })())`);
    got.frames[`${row} at ${width}`] = r;
    claim(`${row}'s Next scrolls its row on in its frame at ${width}px, the previous paddle enabled`, r.scroll > 0 && !r.previous && (row !== 'gallery' || r.current > 0), r);
  }
  // the global nav's search, opened by its button: in a frame 834px wide or
  // more the flyout hung from the bar's bottom, the curtain to the frame's
  // bottom, else the sheet over the frame; its field focused; the frame
  // grown to the window under the local nav; closed by the curtain or the
  // sheet's X, the frame the bar's height again
  {
    const title = 'A global nav and its search';
    await page.go(`${ORIGIN}/components/globalnav/`);
    await page.run(wait(700));
    const closed = await page.run(`Promise.resolve(${frame(title)}.clientHeight)`);
    await page.run(inFrame(title, '[popovertarget="cui-globalnav-search"]', 'click'));
    await page.run(wait(1200));
    const r = await page.run(`Promise.resolve((() => {
      const f = ${frame(title)};
      const d = f.contentDocument, w = f.contentWindow;
      const box = (s) => d.querySelector(s).getBoundingClientRect();
      return {
        open: d.querySelector('.cui-globalnav-search').matches(':popover-open'), focus: d.activeElement.className,
        inner: w.innerWidth, frameHeight: w.innerHeight, nav: box('.cui-globalnav').bottom, search: box('.cui-globalnav-search').top, searchHeight: box('.cui-globalnav-search').height,
        curtain: box('.cui-globalnav-curtain').bottom, height: f.clientHeight, window: innerHeight, bar: document.querySelector('.cui-localnav').getBoundingClientRect().height,
      };
    })())`);
    const flyout = r.inner >= 834;
    got.frames[`globalnav search at ${width}`] = { closed, ...r };
    claim(`the global nav's search opens in its frame at ${width}px, its field focused, the frame the window's height under the local nav`, r.open && r.focus === 'cui-globalnav-searchinput' && Math.abs(r.height - (r.window - r.bar)) <= 1, r);
    claim(flyout ? `in its ${r.inner}px frame it hangs from the bar's bottom, the curtain to the frame's bottom` : `in its ${r.inner}px frame it is a sheet over the frame`,
      flyout ? Math.abs(r.search - r.nav) <= 0.5 && Math.abs(r.curtain - r.frameHeight) <= 0.5 : r.search === 0 && Math.abs(r.searchHeight - r.frameHeight) <= 0.5, r);
    await page.run(inFrame(title, flyout ? '.cui-globalnav-curtain' : '.cui-globalnav-search .cui-globalnav-closebutton', 'click'));
    await page.run(wait(700));
    const again = await page.run(`Promise.resolve(${frame(title)}.clientHeight)`);
    claim(`${flyout ? 'its curtain' : 'its X'} closes it at ${width}px, its frame the bar's height again`, again === closed, { closed, again });
  }
  // the site's own search page, its query in its address as the nav's search
  // form sends it: its field holding it, the rows the theme's script finds
  {
    await page.go(`${ORIGIN}/search/?q=modal`);
    await page.run(wait(1200));
    const r = await page.run(`Promise.resolve({
      q: document.querySelector('.cz-search input').value,
      heading: document.querySelector('.cz-main > .cui-posts .cui-posts-heading')?.textContent ?? null,
      rows: [...document.querySelectorAll('.cz-main > .cui-posts .cui-post')].map((a) => new URL(a.href).pathname),
    })`);
    got.frames[`search at ${width}`] = r;
    claim(`the site's search page at ${width}px lists the pages holding "modal", the modal's page first, its field holding the query`, r.q === 'modal' && r.heading === 'Results for “modal”' && r.rows[0] === '/components/modal/', r);
  }
  // the global nav's suggestions (the theme's script): on the home page its
  // search opened by the page's script, "mo" put in its field and an input
  // event sent, as typing sends one; read in the frame its status line first
  // says what they hold: the modal's page the first suggested link, its row
  // muted and what was typed marked in the rows' color, fading in (still
  // transparent, in 0.32s from 60ms), and 600ms on in full
  {
    await page.go(`${ORIGIN}/`);
    await page.run(wait(700));
    await page.run(`Promise.resolve((document.querySelector('body > .cui-globalnav [popovertarget="cui-globalnav-search"]').click(), true))`);
    await page.run(wait(1200));
    const r = await page.run(`new Promise((done) => {
      const input = document.querySelector('body > .cui-globalnav .cui-globalnav-searchinput');
      input.value = 'mo';
      input.dispatchEvent(new Event('input', { bubbles: true }));
      const t0 = performance.now();
      const read = () => {
        const status = document.querySelector('.cz-status')?.textContent;
        if (!status && performance.now() - t0 < 5000) return requestAnimationFrame(read);
        const row = document.querySelector('body > .cui-globalnav .cui-globalnav-searchlinks .cui-globalnav-searchlink');
        const mark = row?.querySelector('mark');
        done({
          status, ms: Math.round(performance.now() - t0),
          heading: document.querySelector('body > .cui-globalnav .cui-globalnav-searchheading')?.textContent ?? null,
          link: row ? new URL(row.href).pathname : null, marked: mark?.textContent ?? null,
          color: row ? getComputedStyle(row).color : null, markColor: mark ? getComputedStyle(mark).color : null, markGround: mark ? getComputedStyle(mark).backgroundColor : null,
          opacity: row ? +getComputedStyle(row.parentElement).opacity : null,
          transition: row ? [getComputedStyle(row.parentElement).transitionProperty, getComputedStyle(row.parentElement).transitionDuration, getComputedStyle(row.parentElement).transitionDelay].join(' ') : null,
        });
      };
      read();
    })`);
    await page.run(wait(600));
    r.later = await page.run(`Promise.resolve(+getComputedStyle(document.querySelector('body > .cui-globalnav .cui-globalnav-searchlinks > li')).opacity)`);
    got.frames[`suggestions at ${width}`] = r;
    claim(`the global nav's search at ${width}px suggests the modal's page first as "mo" is typed, its row #6e6e73 and "Mo" marked #333336 on no ground, fading in (0.32s from 60ms), 600ms on in full`,
      r.heading === 'Suggested Links' && r.link === '/components/modal/' && r.marked === 'Mo' && r.color === 'rgb(110, 110, 115)' && r.markColor === 'rgb(51, 51, 54)' && r.markGround === 'rgba(0, 0, 0, 0)'
      && r.opacity < 1 && r.transition === 'opacity 0.32s 0.06s' && r.later === 1, r);
  }
  // the dropdown, with no script: empty, its label the placeholder and its
  // own text not shown; a choice made (its value set, as a pick sets it),
  // its label at its top and its text shown
  {
    const title = 'A dropdown, nothing chosen';
    await page.go(`${ORIGIN}/components/field/`);
    await page.run(wait(700));
    const look = `Promise.resolve((() => {
      const d = ${frame(title)}.contentDocument;
      const label = getComputedStyle(d.querySelector('.cui-field-label'));
      return { top: label.top, size: label.fontSize, text: getComputedStyle(d.querySelector('select')).color };
    })())`;
    const empty = await page.run(look);
    await page.run(`Promise.resolve((() => { const s = ${frame(title)}.contentDocument.querySelector('select'); s.value = 'iphone-air'; s.dispatchEvent(new Event('change', { bubbles: true })); return true; })())`);
    await page.run(wait(500));
    const chosen = await page.run(look);
    got.frames[`dropdown at ${width}`] = { empty, chosen };
    claim(`the dropdown at ${width}px: empty, its label the placeholder and its text not shown; chosen, its label at its top and its text shown`,
      empty.top === '18px' && empty.size === '17px' && empty.text === 'rgba(0, 0, 0, 0)' && chosen.top === '10px' && chosen.size === '12px' && chosen.text === 'rgb(29, 29, 31)', { empty, chosen });
  }
  // the textarea, with no script: at rest its label the placeholder; its
  // textarea focused (by the page's script), its box's edge 2px and its
  // label at its top; then blurred and filled (its value set), the label at
  // its top still and the edge 1px
  {
    const title = 'A message of several lines';
    await page.go(`${ORIGIN}/components/field/`);
    await page.run(wait(700));
    const look = `Promise.resolve((() => {
      const d = ${frame(title)}.contentDocument;
      const label = getComputedStyle(d.querySelector('.cui-field-label')), box = getComputedStyle(d.querySelector('.cui-field-box'));
      return { top: label.top, size: label.fontSize, edge: box.borderTopWidth, color: box.borderTopColor };
    })())`;
    const area = (call) => page.run(`Promise.resolve((() => { const t = ${frame(title)}.contentDocument.querySelector('textarea'); ${call}; return true; })())`);
    const rest = await page.run(look);
    await area('t.focus()');
    await page.run(wait(400));
    const focused = await page.run(look);
    await area("t.blur(); t.value = 'A message.'; t.dispatchEvent(new Event('input', { bubbles: true }))");
    await page.run(wait(400));
    const filled = await page.run(look);
    got.frames[`textarea at ${width}`] = { rest, focused, filled };
    claim(`the textarea at ${width}px: at rest its label the placeholder; focused, its box's edge 2px of the accent and its label at its top; filled, its label at its top and its edge 1px`,
      rest.top === '18px' && rest.size === '17px' && rest.edge === '1px' && focused.top === '10px' && focused.size === '12px' && focused.edge === '2px' && focused.color === 'rgb(0, 113, 227)' && filled.top === '10px' && filled.edge === '1px', { rest, focused, filled });
  }
  // an example's link, clicked
  await page.go(`${ORIGIN}/components/link/`);
  await page.run(wait(700));
  await page.run(inFrame('A more link', 'a', 'click'));
  await page.run(wait(700));
  const at = await page.run(`Promise.resolve(${frame('A more link')}.contentWindow.location.href)`);
  claim(`a link in an example leaves its frame where it is at ${width}px`, at === 'about:srcdoc', at);
};
const menus = async (width) => {
  await page.go(`${ORIGIN}/components/globalnav/`);
  await page.run(wait(700));
  const height = (title) => page.run(`Promise.resolve(${frame(title)}.clientHeight)`);
  const closed = await height('A global nav');
  await page.run(inFrame('A global nav', '.cui-globalnav-menubutton', 'click'));
  await page.run(wait(1000));
  const opened = await page.run(`Promise.resolve((() => {
    const f = ${frame('A global nav')};
    const bar = document.querySelector('.cui-localnav').getBoundingClientRect().height;
    return { open: f.contentDocument.querySelector('.cui-globalnav-list').matches(':popover-open'), height: f.clientHeight, top: Math.round(f.getBoundingClientRect().top), window: innerHeight, bar };
  })())`);
  got.frames[`globalnav menu at ${width}`] = { closed, ...opened };
  claim(`the global nav's menu opens in its frame at ${width}px, the frame the window's height under the local nav, scrolled there`, opened.open && Math.abs(opened.height - (opened.window - opened.bar)) <= 1 && Math.abs(opened.top - opened.bar) <= 1, { closed, ...opened });
  await page.run(inFrame('A global nav', '.cui-globalnav-closebutton', 'click'));
  await page.run(wait(700));
  const again = await height('A global nav');
  claim(`closed again, its frame is the bar's height at ${width}px`, again === closed, { closed, again });

  await page.go(`${ORIGIN}/components/localnav/`);
  await page.run(wait(700));
  const bar = await height('A local nav');
  await page.run(inFrame('A local nav', '.cui-localnav-menubutton', 'click'));
  await page.run(wait(1000));
  const menu = await page.run(`Promise.resolve((() => {
    const f = ${frame('A local nav')};
    return { checked: f.contentDocument.querySelector('.cui-localnav-toggle').checked, bottom: Math.ceil(f.contentDocument.querySelector('.cui-localnav-menu').getBoundingClientRect().bottom), height: f.clientHeight };
  })())`);
  got.frames[`localnav menu at ${width}`] = { bar, ...menu };
  claim(`the local nav's menu opens below its bar at ${width}px, its frame grown to hold it`, menu.checked && menu.bottom > bar && menu.height >= menu.bottom, { bar, ...menu });
};

try {
  page = await open(BROWSER, `${ORIGIN}/`);
  got.browser = page.label;
  // a phone's width as it is; elsewhere a wide window, then a narrow one
  for (const size of BROWSER === 'ios' ? [null] : [[1440, 900], [700, 900]]) {
    if (size) await page.size(...size);
    const width = await page.run('Promise.resolve(innerWidth)');
    await read(width);
    await behave(width);
    if (width < 834) await menus(width);
  }
} finally {
  if (page) await page.close();
  server.close();
}
got.fails = fails;
mkdirSync(new URL('./results/', import.meta.url), { recursive: true });
writeFileSync(new URL(`./results/${BROWSER}.json`, import.meta.url), JSON.stringify(got, null, 1));
for (const f of fails) console.log(`FAIL ${f}`);
console.log(`${got.browser}: ${paths.length} pages, ${fails.length} failures`);
process.exitCode = fails.length ? 1 : 0;
