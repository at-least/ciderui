// BROWSER=safari|ios|chrome|chromium|firefox SITE=<dir> node run.mjs
// Do the console parts test/components.mjs holds in Chromium do the same in
// this browser? SITE is build.mjs's pages, served at http://127.0.0.1:8738
// (the browsers: ../browsers.mjs). Each part in a frame of a set width
// (index.html), so its media queries read the frame's width in every
// browser, a phone's below the window's least width in Safari:
// 1. the tabs' phone row (320: six tabs overflow it in every font) is one
//    row that scrolls sideways, none of its height a scrollbar's;
// 2. its edge fades paint, read off the window's pixels at the bar's top
//    edge (the probe's blue ground, white behind it): at its start the right
//    edge fades and the left is solid, scrolled to its end the left fades
//    and the right is solid, between them both fade — the fade at the box's
//    edges once scrolled, not moved with its content;
// 3. the row's last tab, its start a sliver inside the row (Linux's Inter:
//    3px), focused by script: brought in whole (within 1px), its ring inside
//    the tab where :focus-visible matches;
// 4. the thumbnails 240px tall at their own aspect, a wider one than a
//    phone's column narrowed to it, still 240px;
// 5. the pager: on one line, its actions 24px at least after its summary,
//    its buttons 8px apart (375 and 1200);
// 6. the page column: one column at 900; at 1200 the sidebar's 225px
//    column 30px before the page, 40px in, its header and body with no room
//    at their sides; at 1920 the column 1440px wide, centered (240px in);
// 7. the toast, in the window: 16px from its right and bottom edges, never
//    what a point finds there, fading in on data-shown (its opacity read
//    each frame: one between 0 and 1, then 1).
// The result is results/<browser>.json; exits 1 if a claim does not hold.
import { createServer } from 'node:http';
import { existsSync, mkdirSync, readFileSync, statSync, writeFileSync } from 'node:fs';
import { extname, join } from 'node:path';
import { inflateSync } from 'node:zlib';
import { open, wait } from '../browsers.mjs';

const BROWSER = process.env.BROWSER;
const SITE = process.env.SITE;
if (!SITE || !existsSync(join(SITE, 'index.html'))) throw new Error('SITE=<build.mjs\'s dir>');
const ORIGIN = 'http://127.0.0.1:8738';
const TYPES = { '.html': 'text/html; charset=utf-8', '.css': 'text/css', '.woff2': 'font/woff2' };
const server = createServer((req, res) => {
  const f = join(SITE, decodeURIComponent(new URL(req.url, ORIGIN).pathname));
  if (!existsSync(f) || statSync(f).isDirectory()) { res.writeHead(404); return res.end(); }
  res.writeHead(200, { 'content-type': TYPES[extname(f)] ?? 'application/octet-stream' });
  res.end(readFileSync(f));
});
await new Promise((r) => server.listen(8738, '127.0.0.1', r));

// a PNG's pixels (8-bit RGB or RGBA, not interlaced: what the browsers' shots are)
function png(base64) {
  const buf = Buffer.from(base64, 'base64');
  let at = 8;
  let width, height, type;
  const idat = [];
  while (at < buf.length) {
    const len = buf.readUInt32BE(at);
    const name = buf.toString('latin1', at + 4, at + 8);
    const data = buf.subarray(at + 8, at + 8 + len);
    if (name === 'IHDR') {
      width = data.readUInt32BE(0);
      height = data.readUInt32BE(4);
      type = data[9];
      if (data[8] !== 8 || ![2, 6].includes(type) || data[12] !== 0) throw new Error(`a PNG this reader does not read: depth ${data[8]}, type ${type}, interlace ${data[12]}`);
    }
    if (name === 'IDAT') idat.push(data);
    at += 12 + len;
  }
  const bpp = type === 6 ? 4 : 3;
  const raw = inflateSync(Buffer.concat(idat));
  const stride = width * bpp;
  const px = Buffer.alloc(height * stride);
  for (let y = 0; y < height; y++) {
    const f = raw[y * (stride + 1)];
    for (let x = 0; x < stride; x++) {
      const v = raw[y * (stride + 1) + 1 + x];
      const a = x >= bpp ? px[y * stride + x - bpp] : 0;
      const b = y ? px[(y - 1) * stride + x] : 0;
      const c = x >= bpp && y ? px[(y - 1) * stride + x - bpp] : 0;
      const p = a + b - c;
      const pred = [0, a, b, (a + b) >> 1, Math.abs(p - a) <= Math.abs(p - b) && Math.abs(p - a) <= Math.abs(p - c) ? a : Math.abs(p - b) <= Math.abs(p - c) ? b : c][f];
      px[y * stride + x] = (v + pred) & 255;
    }
  }
  return { width, height, red: (x, y) => px[y * stride + x * bpp] };
}

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
  // every frame loaded, its fonts and images too
  await page.run(`(async () => {
    for (const f of document.querySelectorAll('iframe')) {
      if (f.contentDocument?.readyState !== 'complete') await new Promise((r) => f.addEventListener('load', r, { once: true }));
      await f.contentDocument.fonts.ready;
      await Promise.all([...f.contentDocument.images].map((i) => i.decode()));
    }
    return true;
  })()`);
  const frame = (id) => `document.getElementById(${JSON.stringify(id)}).contentWindow`;

  // 1
  const row = await page.run(`(async () => {
    const w = ${frame('tabs-320')};
    const bar = w.document.querySelector('.cui-console-tabs');
    return { wrap: w.getComputedStyle(bar).flexWrap, scrollHeight: bar.scrollHeight, clientHeight: bar.clientHeight, offsetHeight: bar.offsetHeight, scrollWidth: bar.scrollWidth, clientWidth: bar.clientWidth };
  })()`);
  claim('1 tabs one row, scrolling sideways', row.wrap === 'nowrap' && row.scrollHeight === row.clientHeight && row.offsetHeight - row.clientHeight <= 1 && row.scrollWidth > row.clientWidth, row);

  // 2: the red channel 1px below the bar's top at 2px and 12px in from each
  // edge: the blue ground 0, white 255, a fade between
  const fades = {};
  for (const [at, to] of [['start', '0'], ['end', 'bar.scrollWidth'], ['middle', '(bar.scrollWidth - bar.clientWidth) / 2']]) {
    const where = await page.run(`(async () => {
      const w = ${frame('tabs-320')};
      const bar = w.document.querySelector('.cui-console-tabs');
      bar.scrollLeft = ${to};
      w.sync();
      await new Promise((r) => requestAnimationFrame(() => requestAnimationFrame(r)));
      const f = document.getElementById('tabs-320').getBoundingClientRect();
      const b = bar.getBoundingClientRect();
      return { left: f.left + b.left, right: f.left + b.right, top: f.top + b.top, ratio: devicePixelRatio, scrollLeft: bar.scrollLeft,
        attrs: ['data-overflow', 'data-overflow-start', 'data-overflow-end'].filter((a) => bar.hasAttribute(a)), mask: w.getComputedStyle(bar).maskImage };
    })()`);
    const shot = png(await page.shot());
    const r = (x) => shot.red(Math.round(x * where.ratio), Math.round((where.top + 1) * where.ratio));
    fades[at] = { ...where, red: { left2: r(where.left + 2), left12: r(where.left + 12), right12: r(where.right - 12), right2: r(where.right - 2), middle: r((where.left + where.right) / 2) } };
  }
  const solid = (v) => v < 40;
  const faded = (v) => v > 180;
  const between = (v) => v >= 40 && v <= 215;
  const fadesAt = (f, left, right) => solid(f.red.middle) && (left ? faded(f.red.left2) && between(f.red.left12) : solid(f.red.left2) && solid(f.red.left12)) && (right ? faded(f.red.right2) && between(f.red.right12) : solid(f.red.right2) && solid(f.red.right12));
  claim('2 tabs fade at the start: the right edge', fadesAt(fades.start, false, true), fades.start);
  claim('2 tabs fade at the end: the left edge', fadesAt(fades.end, true, false), fades.end);
  claim('2 tabs fade between: both edges', fadesAt(fades.middle, true, true), fades.middle);

  // 3
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

  // 4
  for (const id of ['thumbs-375', 'thumbs-1200']) {
    const imgs = await page.run(`(async () => {
      const w = ${frame(id)};
      const col = w.document.querySelector('.cui-console-thumbs').getBoundingClientRect();
      return { column: col.width, imgs: [...w.document.querySelectorAll('.cui-console-thumb img')].map((i) => ({ natural: [i.naturalWidth, i.naturalHeight], box: [i.getBoundingClientRect().width, i.getBoundingClientRect().height] })) };
    })()`);
    const ok = imgs.imgs.every(({ natural: [nw, nh], box: [bw, bh] }) => bh === 240 && Math.abs(bw - Math.min(240 * nw / nh, imgs.column)) < 0.5);
    claim(`4 thumbnails at ${id.split('-')[1]}: 240px tall, their aspect, no wider than the column`, ok, imgs);
  }

  // 5
  for (const id of ['pager-375', 'pager-1200']) {
    const p = await page.run(`(async () => {
      const d = ${frame(id)}.document;
      const s = d.querySelector('.cui-console-pager-summary').getBoundingClientRect();
      const [a, b] = [...d.querySelectorAll('.cui-console-pager-actions > *')].map((e) => e.getBoundingClientRect());
      return { summaryRight: s.right, summaryTop: s.top, firstLeft: a.left, firstTop: a.top, between: b.left - a.right };
    })()`);
    claim(`5 pager at ${id.split('-')[1]}: one line, 24px at least after the summary, buttons 8px apart`, Math.abs(p.firstTop - p.summaryTop) < 20 && p.firstLeft - p.summaryRight >= 24 && Math.abs(p.between - 8) < 0.01, p);
  }

  // 6
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
  claim('6 page column at 900: one column', c9.display !== 'grid', c9);
  const c12 = content['content-1200'];
  claim('6 page column at 1200: the sidebar 225px, 40px in, the page 30px past it, no room at its sides', c12.display === 'grid' && c12.sidebarLeft === 40 && c12.sidebarWidth === 225 && c12.mainLeft === 40 + 225 + 30 && c12.header === '0px' && c12.body === '0px', c12);
  const c19 = content['content-1920'];
  claim('6 page column at 1920: 1440px, centered, the sidebar 240px in', c19.contentLeft === 240 && c19.contentWidth === 1440 && c19.sidebarLeft === 280, c19);

  // 7
  await page.go(`${ORIGIN}/toast.html`);
  const toast = await page.run(`(async () => {
    await document.fonts.ready;
    const t = document.querySelector('.cui-console-toast');
    const r = t.getBoundingClientRect();
    const root = document.documentElement;
    const at = () => document.elementFromPoint(r.left + r.width / 2, r.top + r.height / 2);
    const rest = { opacity: getComputedStyle(t).opacity, found: at() === t };
    const frames = [];
    const start = performance.now();
    t.setAttribute('data-shown', '');
    while (performance.now() - start < 400) {
      await new Promise((f) => requestAnimationFrame(f));
      frames.push([Math.round(performance.now() - start), Number(getComputedStyle(t).opacity)]);
    }
    return { right: root.clientWidth - r.right, bottom: root.clientHeight - r.bottom, innerHeight, visualHeight: visualViewport.height, rest, shownFound: at() === t, frames };
  })()`);
  const mid = toast.frames.filter(([, o]) => o > 0 && o < 1);
  claim('7 toast: 16px from the right and bottom, never hit, fading in', Math.abs(toast.right - 16) < 0.01 && Math.abs(toast.bottom - 16) < 0.01 && toast.rest.opacity === '0' && !toast.rest.found && !toast.shownFound && mid.length > 0 && toast.frames.at(-1)[1] === 1, toast);
} finally {
  await page.close();
  server.close();
}
mkdirSync(new URL('./results/', import.meta.url), { recursive: true });
writeFileSync(new URL(`./results/${BROWSER}.json`, import.meta.url), JSON.stringify(got, null, 1));
console.log(fails.length ? `${fails.length} failed: ${fails.join('; ')}` : 'all held');
process.exitCode = fails.length ? 1 : 0;
