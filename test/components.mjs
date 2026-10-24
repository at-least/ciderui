// Runs every components/<name>/<name>.test.mjs (its `cases`, test/harness.mjs
// says what a case and a check are): each case rendered in headless Chromium
// on the kit, its fonts and every component's stylesheet, compiled as a site
// compiles them, at each of its widths, its text in Inter as a site that
// imports fonts.css draws it where the system font is not SF (here, Linux:
// without them the stack falls through to this host's Arial, Liberation
// Sans); each check's values against apple.com's
// measured ones, golden.json's or the element's own at rest. golden.json
// beside a test holds its golden values; UPDATE_GOLDEN=1 rewrites them, and
// an entry no check reads fails the run. CASES=<regex> runs the cases it
// matches.

import { existsSync, readdirSync, readFileSync, writeFileSync } from 'node:fs';
import { join } from 'node:path';
import { pathToFileURL } from 'node:url';
import { ROOT, compile } from './compile.mjs';
import { open, quit } from './browser.mjs';
import { ref, asc, tb, WIDTHS } from './reference.mjs';
import { check, done } from './check.mjs';
import { CASE_KEYS, CHECK_KEYS } from './harness.mjs';

const UPDATE = process.env.UPDATE_GOLDEN === '1';
const PICK = process.env.CASES ? new RegExp(process.env.CASES) : null;
// rewriting golden.json from some cases would drop the others' values
if (UPDATE && PICK) throw new Error('UPDATE_GOLDEN=1 runs every case: drop CASES');

const names = readdirSync(join(ROOT, 'components')).filter(n => existsSync(join(ROOT, 'components', n, `${n}.test.mjs`)));
const sheets = names.map(n => `@import "components/${n}/${n}.css";`).join('\n');
// the fonts' files inlined, as each case's page is set from no URL
const inline = (css) => {
  const inlined = css.replace(/url\(fonts\/([^)]+)\)/g, (_, f) => `url(data:font/woff2;base64,${readFileSync(join(ROOT, 'fonts', f)).toString('base64')})`);
  if (inlined === css || /url\(["']?fonts\//.test(inlined)) throw new Error('fonts.css\'s url()s are not the form inlined here');
  return inlined;
};
// Tailwind scans the components' directories, as a site's @source does: the
// examples and the test files, whose cases write markup of their own (a type
// class around an example)
const css = inline(compile({ css: `@import "fonts.css";\n@import "index.css";\n${sheets}`, sources: ['components'] }));

for (const name of names) {
  const file = join(ROOT, 'components', name, `${name}.test.mjs`);
  const { cases } = await import(pathToFileURL(file));
  const goldenFile = join(ROOT, 'components', name, 'golden.json');
  const golden = existsSync(goldenFile) ? JSON.parse(readFileSync(goldenFile, 'utf8')) : {};
  const used = new Set();
  const written = {};

  for (const c of cases) {
    for (const key of Object.keys(c)) if (!CASE_KEYS.includes(key)) throw new Error(`${name} ${c.name}: a case has no key ${key}`);
    for (const k of c.checks) for (const key of Object.keys(k)) if (!CHECK_KEYS.includes(key)) throw new Error(`${name} ${c.name}: ${k.target}: a check has no key ${key}`);
    for (const k of c.checks) if (k.on && !k.state) throw new Error(`${name} ${c.name}: ${k.target}: on needs a state (on is where the state is forced)`);
    if (PICK && !PICK.test(c.name)) continue;
    const page = await open(css, c.html, { touch: c.touch, reducedMotion: c.reducedMotion, forcedColors: c.forcedColors ? 'active' : undefined, height: c.height });
    const widths = c.widths ?? WIDTHS;
    try {
      for (const width of widths) {
        await page.resize(width);
        // the case's clicks, once, at its first width (a menu stays open across
        // the widths after it); then the pointer leaves for the window's
        // bottom corner (a click leaves it over what it clicked, hovered: a
        // state the checks force when they mean it), its keys are pressed,
        // and the motion runs to its end, or for `wait` ms (0: its start)
        if ((c.click || c.press || c.wait !== undefined) && width === widths[0]) {
          for (const selector of [].concat(c.click ?? [])) await page.page.click(selector);
          if (c.click) await page.page.mouse.move(0, (c.height ?? 800) - 1);
          for (const key of [].concat(c.press ?? [])) await page.page.keyboard.press(key);
          if (c.wait !== undefined) await page.page.waitForTimeout(c.wait);
          else await page.settle();
        }
        // `hover`: the pointer off to the window's bottom corner, then moved
        // onto each element (its middle) or point ([x, y]) in turn in steps,
        // as a hand moves it, at every width (what it shows a resize may
        // hide), and the motion run to its end, or for `wait` ms
        if (c.hover) {
          await page.page.mouse.move(0, (c.height ?? 800) - 1);
          await page.settle();
          for (const to of [].concat(c.hover.length === 2 && typeof c.hover[0] === 'number' ? [c.hover] : c.hover)) {
            const r = Array.isArray(to) ? null : await page.page.locator(to).boundingBox();
            if (!Array.isArray(to) && !r) throw new Error(`${name} ${c.name}: ${to} has no box to hover`);
            await page.page.mouse.move(...(r ? [r.x + r.width / 2, r.y + r.height / 2] : to), { steps: 8 });
            if (c.wait !== undefined) await page.page.waitForTimeout(c.wait);
            else await page.settle();
          }
        }
        // `scroll`: [selector, x], that element scrolled to x (snapping where
        // it snaps) at every width, before the checks
        if (c.scroll) {
          await page.page.evaluate(([selector, x]) => { document.querySelector(selector).scrollLeft = x; }, c.scroll);
          await page.page.waitForTimeout(100);
        }
        for (const k of c.checks) {
          const label = `${name} ${c.name}: ${k.target}${k.pseudo ?? ''}${k.state ? ':' + k.state : ''} at ${width}px`;
          const got = await page.read(k.target, k.props, { state: k.state, on: k.on, pseudo: k.pseudo });
          if (k.apple) {
            const [p, probe, state] = k.apple;
            const at = k.at?.[width] ?? width;
            for (const prop of k.props) {
              // apple.com drew SF, ciderui draws Inter here: a width is the text's
              if (prop === 'width' || prop === 'box-width') throw new Error(`${label}: an apple.com check reads no width (${prop})`);
              // `as`: apple.com's property where it does the same with another (a
              // pseudo-element's padding for an icon's margin)
              const theirs = k.as?.[prop] ?? prop;
              check(`${label} ${prop} is apple.com's ${p}:${probe}${state ? ':' + state : ''} ${theirs} at ${at}px`, got[prop], ref(p, probe, at, theirs, state));
            }
          } else if (k.asc) {
            // App Store Connect's, read at 1920px: a value that does not
            // move with the window's width (a console part's own sizes)
            const [p, probe] = k.asc;
            for (const prop of k.props) {
              if (prop === 'width' || prop === 'box-width') throw new Error(`${label}: an asc check reads no width (${prop})`);
              check(`${label} ${prop} is App Store Connect's ${p}:${probe} ${prop}`, got[prop], asc(p, probe, prop));
            }
          } else if (k.tb) {
            // a token of App Store Connect's toolkit, one per property, as
            // this page computes it where the target is (beside it, so a
            // length in em is its parent's)
            if (k.tb.length !== k.props.length) throw new Error(`${label}: tb names a token for each of props`);
            const want = await page.page.evaluate(([selector, pairs]) => {
              const el = document.querySelector(selector);
              // an element no component's selector matches (a div would be
              // a banner's > div), a block where nothing sets its display
              const probe = document.createElement('cui-tb-probe');
              probe.style.display = 'block';
              el.parentElement.append(probe);
              const out = pairs.map(([prop, value]) => {
                // the harness's box reads (test/browser.mjs) are lengths: a
                // token for one is read as a height or a width
                const css = prop === 'box-height' ? 'height' : ['box-width', 'box-top', 'box-left', 'box-right', 'box-x', 'box-y'].includes(prop) ? 'width' : prop;
                probe.style.setProperty(css, value);
                const v = getComputedStyle(probe).getPropertyValue(css);
                probe.style.removeProperty(css);
                return v;
              });
              probe.remove();
              return out;
            }, [k.target, k.props.map((prop, i) => [prop, tb(k.tb[i])])]);
            k.props.forEach((prop, i) => check(`${label} ${prop} is App Store Connect's --tb-${k.tb[i]}`, got[prop], want[i]));
          } else if (k.golden) {
            for (const prop of k.props) {
              const key = `${c.name} | ${k.target}${k.pseudo ?? ''}${k.state ? ':' + k.state : ''} | ${width} | ${prop}`;
              used.add(key);
              if (UPDATE) written[key] = got[prop];
              else if (!(key in golden)) check(`${label} ${prop} has a golden value (UPDATE_GOLDEN=1)`, 'none', 'recorded');
              else check(`${label} ${prop} is golden (${k.golden})`, got[prop], golden[key]);
            }
          } else if (k.rest) {
            const rest = await page.read(k.target, k.props, { pseudo: k.pseudo });
            for (const prop of k.props) check(`${label} ${prop} is its own at rest (${k.rest})`, got[prop], rest[prop]);
          } else {
            throw new Error(`${label}: a check is apple, asc, tb, golden or rest`);
          }
        }
      }
    } finally {
      await page.close();
    }
  }

  if (UPDATE) writeFileSync(goldenFile, JSON.stringify(written, null, 1) + '\n');
  else if (!PICK) for (const key of Object.keys(golden)) check(`${name} golden.json's "${key}" is read by a check`, used.has(key), true);
}

await quit();
done('components');
