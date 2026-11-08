// @external: BROWSER=safari|ios|chrome|chromium|firefox SITE=<dir> node run.mjs
// Do the console parts test/components.mjs holds in Chromium do the same in
// this browser? SITE is build.mjs's pages, served at http://127.0.0.1:8738
// (the browsers: ../browsers.mjs). Each part in a frame of a set width
// (index.html), so its media queries read the frame's width in every
// browser, a phone's below the window's least width in Safari:
// 1. the tabs' phone row (320: six tabs overflow it in every font) is one
//    row that scrolls sideways, none of its height a scrollbar's;
// 2. (the tabs' edge fade, measured to 2026-11-07: the fade is gone)
// 3. the row's last tab, its start a sliver inside the row (Linux's Inter:
//    3px), focused by script: brought in whole (within 1px), its ring inside
//    the tab where :focus-visible matches;
// 4. the page column: one column at 900; at 1200 the sidebar's 225px
//    column 30px before the page, 40px in, its header and body with no room
//    at their sides; at 1920 the column 1440px wide, centered (240px in);
// The result is results/<browser>.json; exits 1 if a claim does not hold.
import { existsSync, mkdirSync, writeFileSync } from 'node:fs';
import { join } from 'node:path';
import { open, serve, wait } from '../browsers.mjs';

const BROWSER = process.env.BROWSER;
const SITE = process.env.SITE;
if (!SITE || !existsSync(join(SITE, 'index.html'))) throw new Error('SITE=<build.mjs\'s dir>');
const ORIGIN = 'http://127.0.0.1:8738';
const server = await serve(SITE, 8738);

const got = { browser: BROWSER, claims: {} };
const fails = [];
const claim = (name, ok, value) => {
  got.claims[name] = { ok, value };
  console.log(`${ok ? 'ok  ' : 'FAIL'} ${name} ${JSON.stringify(value)}`);
  if (!ok) fails.push(name);
};

const page = await open(BROWSER, `${ORIGIN}/index.html`);
try {
  got.label = page.label;
  console.log(page.label);
  // @external: every frame loaded, its fonts and images too
  await page.run(`(async () => {
    for (const f of document.querySelectorAll('iframe')) {
      if (f.contentDocument?.readyState !== 'complete') await new Promise((r) => f.addEventListener('load', r, { once: true }));
      await f.contentDocument.fonts.ready;
      await Promise.all([...f.contentDocument.images].map((i) => i.decode()));
    }
    return true;
  })()`);
  const frame = (id) => `document.getElementById(${JSON.stringify(id)}).contentWindow`;

  // @invariant: 1
  const row = await page.run(`(async () => {
    const w = ${frame('tabs-320')};
    const bar = w.document.querySelector('.cui-console-tabs');
    return { wrap: w.getComputedStyle(bar).flexWrap, scrollHeight: bar.scrollHeight, clientHeight: bar.clientHeight, offsetHeight: bar.offsetHeight, scrollWidth: bar.scrollWidth, clientWidth: bar.clientWidth };
  })()`);
  claim('1 tabs one row, scrolling sideways', row.wrap === 'nowrap' && row.scrollHeight === row.clientHeight && row.offsetHeight - row.clientHeight <= 1 && row.scrollWidth > row.clientWidth, row);

  // @invariant: 3
  const focus = await page.run(`(async () => {
    const w = ${frame('tabs-320')};
    const bar = w.document.querySelector('.cui-console-tabs');
    bar.scrollLeft = 0;
    const tab = bar.querySelector('[href="#settings"]');
    tab.focus();
    // read once the row has not moved for 200ms: a focus's scroll may land
    // after the next frame
    for (let last = -1, still = performance.now(); performance.now() - still < 200;) {
      await new Promise((r) => requestAnimationFrame(r));
      if (bar.scrollLeft !== last) [last, still] = [bar.scrollLeft, performance.now()];
    }
    const t = tab.getBoundingClientRect();
    const b = bar.getBoundingClientRect();
    return { focused: w.document.activeElement === tab, focusVisible: tab.matches(':focus-visible'), outlineOffset: w.getComputedStyle(tab).outlineOffset, scrollLeft: bar.scrollLeft, tab: [t.left - b.left, t.right - b.left], row: b.width };
  })()`);
  focus.shown = Math.min(focus.tab[1], focus.row) - Math.max(focus.tab[0], 0);
  claim('3 a tab focused by script: brought in whole, its ring inside where :focus-visible matches', focus.focused && focus.shown >= focus.tab[1] - focus.tab[0] - 1 && (!focus.focusVisible || focus.outlineOffset === '-2px'), focus);

  // @invariant: 4
  const content = {};
  for (const id of ['content-900', 'content-1200', 'content-1920']) {
    content[id] = await page.run(`(async () => {
      const w = ${frame(id)};
      const c = w.document.querySelector('.cui-console-content');
      const s = w.document.querySelector('.cui-console-sidebar');
      const m = w.document.querySelector('.cui-console-content > main');
      const cs = w.getComputedStyle(c);
      const r = (e) => e.getBoundingClientRect();
      return { display: cs.display, columns: cs.gridTemplateColumns, contentLeft: r(c).left, contentWidth: r(c).width, sidebarLeft: r(s).left, sidebarWidth: r(s).width, mainLeft: r(m).left,
        header: w.getComputedStyle(s.querySelector('.cui-console-sidebar-header')).paddingInlineStart, body: w.getComputedStyle(s.querySelector('.cui-console-sidebar-body')).paddingInlineStart };
    })()`);
  }
  const c9 = content['content-900'];
  claim('4 page column at 900: one column', c9.display !== 'grid', c9);
  const c12 = content['content-1200'];
  claim('4 page column at 1200: the sidebar 225px, 40px in, the page 30px past it, no room at its sides', c12.display === 'grid' && c12.sidebarLeft === 40 && c12.sidebarWidth === 225 && c12.mainLeft === 40 + 225 + 30 && c12.header === '0px' && c12.body === '0px', c12);
  const c19 = content['content-1920'];
  claim('4 page column at 1920: 1440px, centered, the sidebar 240px in', c19.contentLeft === 240 && c19.contentWidth === 1440 && c19.sidebarLeft === 280, c19);
} finally {
  await page.close();
  server.close();
}
mkdirSync(new URL('./results/', import.meta.url), { recursive: true });
writeFileSync(new URL(`./results/${BROWSER}.json`, import.meta.url), JSON.stringify(got, null, 1));
console.log(fails.length ? `${fails.length} failed: ${fails.join('; ')}` : 'all held');
process.exitCode = fails.length ? 1 : 0;
