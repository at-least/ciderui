// Runs every components/<name>/<name>.test.mjs (its `cases`, test/harness.mjs
// says what a case and a check are): each case rendered in headless Chromium
// on the kit and every component's stylesheet, compiled as a site compiles
// them, at each of its widths; each check's values against apple.com's
// measured ones, golden.json's or the element's own at rest. golden.json
// beside a test holds its golden values; UPDATE_GOLDEN=1 rewrites them, and
// an entry no check reads fails the run. CASES=<regex> runs the cases it
// matches.

import { existsSync, readdirSync, readFileSync, writeFileSync } from 'node:fs';
import { join } from 'node:path';
import { pathToFileURL } from 'node:url';
import { ROOT, compile } from './compile.mjs';
import { open } from './browser.mjs';
import { ref, WIDTHS } from './reference.mjs';
import { check, done } from './check.mjs';

const UPDATE = process.env.UPDATE_GOLDEN === '1';
const PICK = process.env.CASES ? new RegExp(process.env.CASES) : null;
// rewriting golden.json from some cases would drop the others' values
if (UPDATE && PICK) throw new Error('UPDATE_GOLDEN=1 runs every case: drop CASES');

const names = readdirSync(join(ROOT, 'components')).filter(n => existsSync(join(ROOT, 'components', n, `${n}.test.mjs`)));
const sheets = names.map(n => `@import "components/${n}/${n}.css";`).join('\n');
// Tailwind scans the components' directories, as a site's @source does: the
// examples and the test files, whose cases write markup of their own (a type
// class around an example)
const css = compile({ css: `@import "index.css";\n${sheets}`, sources: ['components'] });

for (const name of names) {
  const file = join(ROOT, 'components', name, `${name}.test.mjs`);
  const { cases } = await import(pathToFileURL(file));
  const goldenFile = join(ROOT, 'components', name, 'golden.json');
  const golden = existsSync(goldenFile) ? JSON.parse(readFileSync(goldenFile, 'utf8')) : {};
  const used = new Set();
  const written = {};

  for (const c of cases) {
    if (PICK && !PICK.test(c.name)) continue;
    const page = await open(css, c.html, { touch: c.touch, reducedMotion: c.reducedMotion, height: c.height });
    const widths = c.widths ?? WIDTHS;
    try {
      for (const width of widths) {
        await page.resize(width);
        // the case's clicks, once, at its first width (a menu stays open across
        // the widths after it); then the pointer leaves for the window's
        // bottom corner (a click leaves it over what it clicked, hovered: a
        // state the checks force when they mean it), its keys are pressed,
        // and the motion runs to its end, or for `wait` ms (0: its start)
        if ((c.click || c.wait !== undefined) && width === widths[0]) {
          for (const selector of [].concat(c.click ?? [])) await page.page.click(selector);
          if (c.click) await page.page.mouse.move(0, (c.height ?? 800) - 1);
          for (const key of [].concat(c.press ?? [])) await page.page.keyboard.press(key);
          await page.page.waitForTimeout(c.wait ?? (c.reducedMotion ? 100 : 1000));
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
            throw new Error(`${label}: a check is apple, golden or rest`);
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

done('components');
