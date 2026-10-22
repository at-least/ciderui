// node apple.mjs: where apple.com lands a footnote's note under its sticky
// local nav, and what puts it there. Three pages in Playwright's Chromium
// (../browsers.mjs), at 1440 by 900 and at 390 by 900: iPhone 18 Pro's,
// whose local nav is a newer one (nav#localnav, 72px), and MacBook Air's
// and AirPods Pro's, whose local nav is the one ciderui's is
// (#ac-localnav, 52px, 48px to 833px). Each scrolled 2000px down, its first
// footnote mark with a box (a link to #footnote-…) clicked, and read 2s on:
// the note's top against the nav's foot, the page at its end or not, the
// root's scroll-padding-top, the note's scroll-margin-top before the click
// and after (its inline style too: what a script sets), and the rules in
// the page's stylesheets that set it on the note (their selector, value and
// sheet; a sheet the page cannot read, another origin's, counted). The
// result is results/apple.json; it records what each page does, and fails
// only where a page has no mark, note or nav to read.
import { mkdirSync, writeFileSync } from 'node:fs';
import { open, wait } from '../browsers.mjs';

const PAGES = ['iphone-18-pro', 'macbook-air', 'airpods-pro'];
// the rules that set the note's scroll-margin-top, through grouping rules
// (@media, @supports, @layer)
const RULES = `(note) => {
  const found = [];
  let unreadable = 0;
  const walk = (rules, sheet) => {
    for (const rule of rules) {
      if (rule.selectorText && rule.style.scrollMarginTop && note.matches(rule.selectorText)) found.push({ selector: rule.selectorText, value: rule.style.scrollMarginTop, sheet });
      if (rule.cssRules) walk(rule.cssRules, sheet);
    }
  };
  for (const s of document.styleSheets) {
    let rules;
    try { rules = s.cssRules; } catch { unreadable++; continue; }
    walk(rules, s.href ?? '(inline)');
  }
  return { found, unreadable };
}`;

const got = {};
let ok = true;
const page = await open('chromium', 'https://www.apple.com/');
try {
  got.browser = page.label;
  for (const name of PAGES) {
    for (const width of [1440, 390]) {
      await page.size(width, 900);
      await page.go(`https://www.apple.com/${name}/`);
      await page.run(wait(3000));
      const href = await page.run(`Promise.resolve((() => {
        const a = [...document.querySelectorAll('a[href^="#footnote"]')].find((a) => a.getBoundingClientRect().width > 0);
        if (!a) return null;
        a.setAttribute('data-probe-mark', '');
        return a.getAttribute('href');
      })())`);
      const id = href?.slice(1);
      const before = href && await page.run(`Promise.resolve(getComputedStyle(document.getElementById(${JSON.stringify(id)})).scrollMarginTop)`);
      if (!href) { got[`${name} at ${width}`] = { mark: null }; ok = false; continue; }
      await page.run(`Promise.resolve((scrollTo(0, 2000), true))`);
      await page.run(wait(800));
      await page.click('[data-probe-mark]');
      await page.run(wait(2000));
      const r = await page.run(`Promise.resolve((() => {
        const note = document.getElementById(${JSON.stringify(id)});
        const nav = document.querySelector('#ac-localnav, nav#localnav');
        if (!note || !nav) return { note: Boolean(note), nav: Boolean(nav) };
        return {
          href: ${JSON.stringify(href)},
          hash: location.hash,
          nav: { id: nav.id, position: getComputedStyle(nav).position, height: nav.getBoundingClientRect().height, bottom: +nav.getBoundingClientRect().bottom.toFixed(2) },
          note: +note.getBoundingClientRect().top.toFixed(2),
          end: scrollY >= document.documentElement.scrollHeight - innerHeight - 1,
          rootScrollPaddingTop: getComputedStyle(document.documentElement).scrollPaddingTop,
          noteScrollMarginTop: getComputedStyle(note).scrollMarginTop,
          noteInlineScrollMarginTop: note.style.scrollMarginTop,
          rules: (${RULES})(note),
        };
      })())`);
      if (typeof r.note !== 'number') { got[`${name} at ${width}`] = r; ok = false; continue; }
      r.noteScrollMarginTopBefore = before;
      r.lands = Math.abs(r.note - r.nav.bottom) <= 1 ? 'at the nav\'s foot' : r.note < r.nav.bottom ? 'under the nav' : r.end ? 'below the nav, the page at its end' : 'below the nav';
      got[`${name} at ${width}`] = r;
      console.log(`${name} at ${width}px (${r.nav.id}, ${r.nav.position}, ${r.nav.height}px): the note ${r.lands}, its top ${r.note}px, the nav's foot ${r.nav.bottom}px; scroll-margin-top ${r.noteScrollMarginTopBefore} before the click, ${r.noteScrollMarginTop} after (inline: ${r.noteInlineScrollMarginTop || 'none'}; set by ${r.rules.found.map((f) => `${f.selector} { ${f.value} }`).join(', ') || 'no readable rule'}; ${r.rules.unreadable} sheets unreadable); the root's scroll-padding-top ${r.rootScrollPaddingTop}`);
    }
  }
} finally {
  await page.close();
}
mkdirSync(new URL('./results/', import.meta.url), { recursive: true });
writeFileSync(new URL('./results/apple.json', import.meta.url), JSON.stringify({ date: new Date().toISOString(), ...got }, null, 1));
if (!ok) console.log('FAIL a page had no footnote mark, note or local nav to read (results/apple.json)');
process.exit(ok ? 0 : 1);
