// The documentation (docs/): it has a page for each component and no other
// under docs/content/components/, each with its kind and its picture
// (docs/static/pictures/, scripts/docs-pictures.mjs) and showing every
// example file of its component's directory, so no example the crate's
// tests render is missing from it; the frames' stylesheet styles every
// class the examples write. Built with zola, every page at 1440 and 390px
// in Chromium: no page error, nothing the page asks for missing, every
// link within the site to a page that exists, no scrolling sideways; each
// example frame loads its stylesheet, draws every picture it has, is no
// wider than the page and as tall as its content, so it never scrolls.
// And what the frames do: the theme's script moves a gallery's row; a
// modal opens in its frame by its button, the frame grown to the window under
// the page's local nav, and its close button closes it; the global nav's
// search the same, its field focused, Escape closing it; at 390px the global
// nav's menu opens over its frame, grown to the window's height, and the
// local nav's below its bar, the frame grown to hold it; an example's link
// leaves the frame where it is.
//
//   node test/docs.mjs
//
// Needs zola, and ciderui-zola, the theme the site is built with, as a
// sibling checkout: docs/themes/ciderui-zola links to ../ciderui-zola. The
// site is built into a temporary directory and served from it through a
// Playwright route at its base_url.

import { execFileSync } from 'node:child_process';
import { existsSync, mkdtempSync, readFileSync, readdirSync, rmSync, statSync } from 'node:fs';
import { tmpdir } from 'node:os';
import { extname, join, relative } from 'node:path';
import { chromium } from 'playwright-chromium';
import { ROOT } from './compile.mjs';
import { check, done } from './check.mjs';

const DOCS = join(ROOT, 'docs');
const ORIGIN = 'https://ciderui.test';
const TYPES = { '.html': 'text/html', '.css': 'text/css', '.js': 'text/javascript', '.woff2': 'font/woff2', '.png': 'image/png', '.svg': 'image/svg+xml', '.xml': 'application/xml', '.txt': 'text/plain' };

// the pages: one for each component, its examples all shown
const components = readdirSync(join(ROOT, 'components')).sort();
const pages = readdirSync(join(DOCS, 'content/components')).filter((f) => f !== '_index.md').map((f) => f.replace(/\.md$/, ''));
for (const name of components) {
  const file = join(DOCS, 'content/components', `${name}.md`);
  if (!check(`${name} has a page, docs/content/components/${name}.md`, existsSync(file), true)) continue;
  const md = readFileSync(file, 'utf8');
  const shown = [...md.matchAll(/<cui_example file="([^"]+)"/g)].map((m) => m[1]);
  for (const f of readdirSync(join(ROOT, 'components', name, 'examples')).filter((f) => f.endsWith('.html'))) {
    check(`docs/content/components/${name}.md shows ${name}/${f}`, shown.includes(`${name}/${f.replace(/\.html$/, '')}`), true);
  }
  check(`docs/content/components/${name}.md has its kind ([extra] category)`, /^category = "[^"]+"$/m.test(md), true);
  const image = md.match(/^image = "([^"]+)"$/m);
  check(`docs/content/components/${name}.md has its picture, in docs/static/`, Boolean(image && existsSync(join(DOCS, 'static', image[1]))), true);
}
for (const name of pages.filter((p) => !components.includes(p))) check(`docs/content/components/${name}.md is a component's page`, false, true);

// the frames' stylesheet has ciderui's rules for every class the examples
// write: a component's (its stylesheet names it), a block's (base.css:
// cui-dark, cui-light) or a type utility (type.css);
// a class ciderui has no rule for (a part's box that needs none,
// cui-post-content) is the markup's alone
const css = readFileSync(join(DOCS, 'static/example.css'), 'utf8');
const ruled = new Set([
  ...[...components.map((c) => `components/${c}/${c}.css`), 'base.css'].flatMap((f) => [...readFileSync(join(ROOT, f), 'utf8').matchAll(/\.(cui-[a-z0-9-]+)/g)].map((m) => m[1])),
  ...[...readFileSync(join(ROOT, 'type.css'), 'utf8').matchAll(/@utility ([a-z0-9-]+)/g)].map((m) => m[1]),
]);
for (const e of readdirSync(join(DOCS, 'examples'), { recursive: true, withFileTypes: true }).filter((e) => e.isFile())) {
  const html = readFileSync(join(e.parentPath, e.name), 'utf8');
  for (const c of new Set([...html.matchAll(/\bclass="([^"]*)"/g)].flatMap((m) => m[1].split(/\s+/).filter((c) => ruled.has(c))))) {
    check(`static/example.css styles .${c} (${relative(DOCS, join(e.parentPath, e.name))})`, css.includes(`.${c}`), true);
  }
}

// the site, built
const out = mkdtempSync(join(tmpdir(), 'ciderui-docs-'));
try {
  execFileSync('zola', ['--root', DOCS, 'build', '--output-dir', out, '--force'], { stdio: 'pipe' });
} catch (e) {
  console.log(`zola build docs/: ${e.stderr}`);
  process.exit(1);
}
const htmlPages = (dir) => readdirSync(dir).flatMap((f) => {
  const p = join(dir, f);
  return statSync(p).isDirectory() ? htmlPages(p) : f === 'index.html' ? [p] : [];
});
const urls = htmlPages(out).map((f) => `${ORIGIN}/${relative(out, f).replace(/index\.html$/, '')}`).sort();
const file = (url) => {
  let f = join(out, decodeURIComponent(new URL(url).pathname));
  if (existsSync(f) && statSync(f).isDirectory()) f = join(f, 'index.html');
  return existsSync(f) ? f : null;
};

const browser = await chromium.launch({ args: ['--font-render-hinting=none'] });
for (const width of [1440, 390]) {
  const page = await browser.newPage({ viewport: { width, height: 900 } });
  let missing = [], errors = [];
  page.on('pageerror', (e) => errors.push(String(e)));
  await page.route(`${ORIGIN}/**`, (route) => {
    const f = file(route.request().url());
    if (!f) {
      missing.push(route.request().url());
      return route.fulfill({ status: 404, body: '' });
    }
    route.fulfill({ body: readFileSync(f), contentType: TYPES[extname(f)] ?? 'application/octet-stream' });
  });
  for (const url of urls) {
    missing = [];
    errors = [];
    const where = `${url.slice(ORIGIN.length)} at ${width}px`;
    await page.goto(url, { waitUntil: 'load' });
    await page.evaluate(() => document.fonts.ready);
    // the frames settle their height a frame after their content's
    await page.evaluate(() => new Promise((r) => requestAnimationFrame(() => requestAnimationFrame(() => setTimeout(r, 50)))));
    const got = await page.evaluate((origin) => ({
      sideways: document.documentElement.scrollWidth - innerWidth,
      links: [...document.querySelectorAll('a[href]')].map((a) => a.href.split('#')[0]).filter((h) => h.startsWith(origin)),
      frames: [...document.querySelectorAll('iframe.cz-example')].map((f) => {
        const doc = f.contentDocument;
        const sheet = [...doc.styleSheets].find((s) => s.href?.includes('/example.css'));
        return {
          title: f.title,
          css: Boolean(sheet && sheet.cssRules.length),
          images: [...doc.images].filter((i) => !i.complete || i.naturalWidth === 0).map((i) => i.outerHTML.slice(0, 60)),
          width: f.getBoundingClientRect().width,
          page: document.documentElement.clientWidth,
          height: f.clientHeight,
          content: doc.documentElement.scrollHeight,
          scroll: doc.documentElement.scrollWidth - doc.documentElement.clientWidth,
        };
      }),
    }), ORIGIN);
    check(`${where}: no page error`, errors.join(' | '), '');
    check(`${where}: nothing it asks for is missing`, missing.join(' '), '');
    check(`${where}: no scrolling sideways`, got.sideways, 0);
    for (const link of new Set(got.links)) check(`${where}: its link ${link} is a page of the site`, Boolean(file(link)), true);
    for (const f of got.frames) {
      check(`${where}, "${f.title}": its frame loads static/example.css`, f.css, true);
      check(`${where}, "${f.title}": every picture in its frame loads`, f.images.join(' '), '');
      check(`${where}, "${f.title}": its frame is no wider than the page`, f.width <= f.page, true);
      check(`${where}, "${f.title}": its frame is as tall as its content (${f.height}px, its content ${f.content}px)`, f.height === f.content, true);
      check(`${where}, "${f.title}": its frame does not scroll sideways`, f.scroll, 0);
    }
  }

  // what the frames do
  const frame = async (path, title) => {
    await page.goto(`${ORIGIN}${path}`, { waitUntil: 'load' });
    await page.evaluate(() => new Promise((r) => requestAnimationFrame(() => requestAnimationFrame(() => setTimeout(r, 50)))));
    const handle = await page.$(`iframe[title="${title}"]`);
    return { handle, inner: await handle.contentFrame() };
  };
  const height = (handle) => handle.evaluate((f) => f.clientHeight);
  const settle = (inner) => inner.evaluate(() => new Promise((r) => requestAnimationFrame(() => requestAnimationFrame(() => setTimeout(r, 400)))));

  {
    const { inner } = await frame('/components/gallery/', 'A gallery of cards');
    await inner.evaluate(() => document.querySelector('[aria-controls="gallery"][aria-label="Next"]').click());
    await settle(inner);
    const row = await inner.evaluate(() => ({ scroll: document.getElementById('gallery').scrollLeft, previous: document.querySelector('[aria-controls="gallery"]').disabled }));
    check(`the gallery's Next scrolls its row in its frame at ${width}px, its previous paddle enabled`, row.scroll > 0 && !row.previous, true);
  }
  {
    const { handle, inner } = await frame('/components/link/', 'A more link');
    await inner.click('a');
    await settle(inner);
    check(`a link in an example leaves its frame where it is at ${width}px`, await inner.evaluate(() => location.href), 'about:srcdoc');
  }
  {
    const { handle, inner } = await frame('/components/modal/', 'A modal and the button that opens it');
    const closed = await height(handle);
    await inner.click('[command="show-modal"]');
    await settle(inner);
    const open = await inner.evaluate(() => document.querySelector('.cui-modal').matches(':modal'));
    const below = await page.evaluate(() => innerHeight - document.querySelector('.cui-localnav').getBoundingClientRect().height);
    check(`at ${width}px a modal opens in its frame by its button, with no script`, open, true);
    check(`at ${width}px its frame grows to the window's height below the page's local nav (${closed}px closed)`, await height(handle), below);
    await inner.click('.cui-modal-close');
    await settle(inner);
    check(`at ${width}px its close button closes it, its frame its content's height again`, await height(handle), closed);
  }
  {
    const { handle, inner } = await frame('/components/globalnav/', 'A global nav and its search');
    const closed = await height(handle);
    await inner.click('[popovertarget="cui-globalnav-search"]');
    await settle(inner);
    const open = await inner.evaluate(() => `${document.querySelector('.cui-globalnav-search').matches(':popover-open')}, ${document.activeElement.className}`);
    const below = await page.evaluate(() => innerHeight - document.querySelector('.cui-localnav').getBoundingClientRect().height);
    check(`at ${width}px the global nav's search opens in its frame by its button, its field focused`, open, 'true, cui-globalnav-searchinput');
    check(`at ${width}px its frame grows to the window's height below the page's local nav (${closed}px closed)`, await height(handle), below);
    await inner.press('.cui-globalnav-searchinput', 'Escape');
    await settle(inner);
    check(`at ${width}px Escape closes it, its frame the bar's height again`, await height(handle), closed);
  }
  if (width === 390) {
    {
      const { handle, inner } = await frame('/components/globalnav/', 'A global nav');
      const closed = await height(handle);
      await inner.click('.cui-globalnav-menubutton');
      await settle(inner);
      const open = await inner.evaluate(() => document.querySelector('.cui-globalnav-list').matches(':popover-open'));
      check('at 390px the global nav\'s menu opens in its frame', open, true);
      const below = await page.evaluate(() => innerHeight - document.querySelector('.cui-localnav').getBoundingClientRect().height);
      check(`at 390px its frame grows to the window's height below the page's local nav (${closed}px closed)`, await height(handle), below);
      check('and it is scrolled there, under the local nav', await handle.evaluate((f) => Math.round(f.getBoundingClientRect().top)), 900 - below);
      await inner.click('.cui-globalnav-closebutton');
      await settle(inner);
      check('closed again, its frame is the bar\'s height', await height(handle), closed);
    }
    {
      const { handle, inner } = await frame('/components/localnav/', 'A local nav');
      const closed = await height(handle);
      await inner.click('.cui-localnav-menubutton');
      await settle(inner);
      const menu = await inner.evaluate(() => Math.ceil(document.querySelector('.cui-localnav-menu').getBoundingClientRect().bottom));
      check(`at 390px the local nav's menu opens below its bar (${closed}px), its frame grown to hold it`, menu > closed && (await height(handle)) >= menu, true);
    }
  }
  await page.close();
}
await browser.close();
rmSync(out, { recursive: true, force: true });
done(`docs: ${components.length} components, ${urls.length} pages`);
