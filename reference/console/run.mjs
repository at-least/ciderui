// @external: BROWSER=safari|ios|chrome|chromium|firefox SITE=<dir> node run.mjs
// Do the console parts test/components.mjs holds in Chromium do the same in
// this browser? SITE is build.mjs's pages, served at http://127.0.0.1:8738
// (the browsers: ../browsers.mjs). Each part in a frame of a set width
// (index.html), so its media queries read the frame's width in every
// browser, a phone's below the window's least width in Safari:
// 1. the tabs' phone row (320: six tabs overflow it in every font) is one
//    row that scrolls sideways, none of its height a scrollbar's;
// 2. its edge fades paint, in cui-light and in cui-dark, read off the
//    window's pixels at the bar's top edge (the probe's blue ground, the
//    page white or black behind it): at its start the right edge fades and
//    the left is solid, scrolled to its end the left fades and the right is
//    solid, between them both fade — the fade at the box's edges once
//    scrolled, not moved with its content;
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

// @invariant: how much of the page behind the blue bar shows at points of a picture of
// the window (base64 PNG), each [x, y] in the page's px, 0 the bar to 255 the page:
// light, the red channel (blue 0, white 255); dark, 255 less the blue (blue 255,
// black 0). Read in the page as ../dialog/dialog.js reads its
const grounds = (png, ratio, points, scheme) => page.run(`(async () => {
  const img = new Image();
  img.src = 'data:image/png;base64,${png}';
  await img.decode();
  const c = document.createElement('canvas');
  c.width = img.width;
  c.height = img.height;
  const g = c.getContext('2d');
  g.drawImage(img, 0, 0);
  const points = ${JSON.stringify(points)};
  const dark = ${JSON.stringify(scheme === 'dark')};
  return Object.fromEntries(Object.entries(points).map(([k, [x, y]]) => {
    const d = g.getImageData(Math.round(x * ${ratio}), Math.round(y * ${ratio}), 1, 1).data;
    return [k, dark ? 255 - d[2] : d[0]];
  }));
})()`);

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

  // @external: 2: the ground 1px below the bar's top at 2px and 12px in from each
  // edge: the blue bar 0, the page 255, a fade between
  const solid = (v) => v < 40;
  const faded = (v) => v > 180;
  const between = (v) => v >= 40 && v <= 215;
  const fadesAt = (f, left, right) => solid(f.ground.middle) && (left ? faded(f.ground.left2) && between(f.ground.left12) : solid(f.ground.left2) && solid(f.ground.left12)) && (right ? faded(f.ground.right2) && between(f.ground.right12) : solid(f.ground.right2) && solid(f.ground.right12));
  for (const [scheme, id] of [['light', 'tabs-320'], ['dark', 'tabs-dark-320']]) {
    const fades = {};
    for (const [at, to] of [['start', '0'], ['end', 'bar.scrollWidth'], ['middle', '(bar.scrollWidth - bar.clientWidth) / 2']]) {
      const where = await page.run(`(async () => {
        const w = ${frame(id)};
        const bar = w.document.querySelector('.cui-console-tabs');
        bar.scrollLeft = ${to};
        w.sync();
        await new Promise((r) => requestAnimationFrame(() => requestAnimationFrame(r)));
        const f = document.getElementById(${JSON.stringify(id)}).getBoundingClientRect();
        const b = bar.getBoundingClientRect();
        return { left: f.left + b.left, right: f.left + b.right, top: f.top + b.top, ratio: devicePixelRatio, scrollLeft: bar.scrollLeft,
          attrs: ['data-overflow', 'data-overflow-start', 'data-overflow-end'].filter((a) => bar.hasAttribute(a)), mask: w.getComputedStyle(bar).maskImage,
          scheme: w.getComputedStyle(w.document.documentElement).colorScheme };
      })()`);
      const y = where.top + 1;
      const ground = await grounds(await page.shot(), where.ratio, { left2: [where.left + 2, y], left12: [where.left + 12, y], right12: [where.right - 12, y], right2: [where.right - 2, y], middle: [(where.left + where.right) / 2, y] }, scheme);
      fades[at] = { ...where, ground };
    }
    claim(`2 tabs fade at the start (${scheme}): the right edge`, fadesAt(fades.start, false, true), fades.start);
    claim(`2 tabs fade at the end (${scheme}): the left edge`, fadesAt(fades.end, true, false), fades.end);
    claim(`2 tabs fade between (${scheme}): both edges`, fadesAt(fades.middle, true, true), fades.middle);
  }

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
