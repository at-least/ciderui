// @invariant: The kit's rules that hold for every part, read off every example, not listed
// per part: a new part is held to them with no line of its own.

import { readFileSync } from 'node:fs';
import { join } from 'node:path';
import { COMPONENTS, ROOT, exampleFiles, exampleHtml, shipped } from './compile.mjs';
import { open, quit } from './browser.mjs';
import { OUTLINE } from './harness.mjs';
import { ref } from './reference.mjs';
import { check, done } from './check.mjs';

// @invariant: rule 5: what (hover: hover) gates, a touch screen never matches
// (test/browser.mjs throws on a touch page where it does)
const hoverOutsideItsQuery = (text) => {
  const enclosing = [];
  const found = [];
  let start = 0;
  for (let i = 0; i < text.length; i++) {
    if (text[i] === '{') {
      const prelude = text.slice(start, i).trim();
      enclosing.push(prelude);
      if (!prelude.startsWith('@') && prelude.includes(':hover') && !enclosing.some((p) => p.includes('(hover:hover)'))) found.push(prelude);
      start = i + 1;
    } else if (text[i] === '}') {
      enclosing.pop();
      start = i + 1;
    } else if (text[i] === ';') {
      start = i + 1;
    }
  }
  return found;
};
for (const set of ['site', 'console']) check(`${set}: every :hover sits inside (hover: hover) (rule 5)`, hoverOutsideItsQuery(shipped(set)).join(' | '), '');

// @invariant: rule 9's exceptions, each read in its part's test: a text field draws its
// edge or halo in the ring's place; the others move the one ring in, where a
// scroller or their edge would clip it
const NO_RING = ['cui-console-field', 'cui-field-input'];
const RING_MOVED = ['cui-console-list-item'];
// @invariant: rule 10 is a pointer's color change; these change with a panel opening, a
// field's label rising or its text typed, each read in its part's test
const OWN_MOTION = ['.cui-field-label'];
const KEYBOARD_REACHABLE = 'a[href], button, input:not([type=hidden]), select, textarea, summary, [tabindex]:not([tabindex="-1"])';

const tracking = JSON.parse(readFileSync(join(ROOT, 'reference/tracking/cui-tracking.json'), 'utf8'));

const readTextAndMotion = ([set, tracking, ownMotion]) => {
  const label = (e) => `${e.tagName.toLowerCase()}${typeof e.className === 'string' && e.className ? `.${e.className.trim().split(/\s+/).join('.')}` : ''}`;
  const motion = [];
  const motionUsed = [];
  const spacing = [];
  let texts = 0;
  for (const e of document.body.querySelectorAll('*')) {
    const cs = getComputedStyle(e);
    const durations = cs.transitionDuration.split(', ');
    cs.transitionProperty.split(', ').forEach((prop, i) => {
      if (!(parseFloat(durations[i % durations.length]) > 0 && /(^|-)color$|^background$|^all$|^fill$|^stroke$/.test(prop))) return;
      const own = ownMotion.filter((s) => e.matches(s));
      if (own.length) motionUsed.push(...own);
      else motion.push(`${label(e)}: ${prop}`);
    });
    if (!e.matches('input:not([type=checkbox], [type=radio], [type=range]), textarea, select') && ![...e.childNodes].some((n) => n.nodeType === Node.TEXT_NODE && n.textContent.trim())) continue;
    texts++;
    if (set === 'console') {
      if (cs.letterSpacing !== 'normal') spacing.push(`${label(e)}: ${cs.letterSpacing}`);
    } else {
      const size = parseFloat(cs.fontSize);
      // @trap: not "more than off": `normal` parses to NaN, which compares false either way
      if (!(Math.abs(parseFloat(cs.letterSpacing) - size * tracking[size]) <= 0.001)) spacing.push(`${label(e)}: ${cs.fontSize} at ${cs.letterSpacing}`);
    }
  }
  return { motion, motionUsed, spacing, texts };
};

const noRingUsed = new Set();
const ringMovedUsed = new Set();
const motionUsed = new Set();
let examples = 0;
let controls = 0;
let texts = 0;
// @trap: one page a set, its body swapped: no example has a script, and a page's open is
// the most of this step's time
const pages = { site: await open(shipped('site'), ''), console: await open(shipped('console'), '') };

for (const part of COMPONENTS) {
  const page = pages[part.set];
  for (const file of exampleFiles(part)) {
    const name = `${part.name}/${file}`;
    await page.page.evaluate((html) => { document.body.innerHTML = html; }, exampleHtml(part, file));
    examples++;
    for (const width of [390, 900, 1200]) {
      await page.resize(width);
      const found = await page.page.evaluate(readTextAndMotion, [part.set, tracking, OWN_MOTION]);
      texts += found.texts;
      found.motionUsed.forEach((s) => motionUsed.add(s));
      check(`${name} at ${width}px: no color moves (rule 10)`, found.motion.join(' | '), '');
      check(`${name} at ${width}px: ${part.set === 'console' ? 'the font tracks itself' : 'each size its fitted letter-spacing (tokens.css)'}`, found.spacing.join(' | '), '');
    }
    const classes = await page.page.evaluate((selector) => [...document.body.querySelectorAll(selector)].map((e, i) => {
      e.setAttribute('data-cui-control', i);
      return [...e.classList];
    }), KEYBOARD_REACHABLE);
    for (const [i, own] of classes.entries()) {
      controls++;
      const noRing = NO_RING.filter((c) => own.includes(c));
      const moved = RING_MOVED.filter((c) => own.includes(c));
      // @trap: both, as a keyboard's focus is: a field's :focus sets the ring aside
      const ring = await page.read(`[data-cui-control="${i}"]`, OUTLINE, { state: ['focus', 'focus-visible'] });
      const off = OUTLINE.filter((prop) => ring[prop] !== ref('home', 'button', 1200, prop, 'focus-visible'));
      if (noRing.length) {
        if (off.length) noRing.forEach((c) => noRingUsed.add(c));
      } else if (moved.length && off.join() === 'outline-offset') {
        moved.forEach((c) => ringMovedUsed.add(c));
      } else {
        check(`${name}: control ${i}'s focus is base.css's one ring, apple.com's (rule 9)`, off.map((prop) => `${prop} ${ring[prop]}`).join(' '), '');
      }
    }
  }
}
await pages.site.close();
await pages.console.close();
// @invariant: an exception nothing needs is a part let off for nothing
check('every part that sets the ring aside still does', NO_RING.filter((c) => !noRingUsed.has(c)).join(' '), '');
check('every part that moves the ring still does', RING_MOVED.filter((c) => !ringMovedUsed.has(c)).join(' '), '');
check('every part with its own color motion still has it', OWN_MOTION.filter((s) => !motionUsed.has(s)).join(' '), '');
// @trap: a rule read off nothing would pass having held nothing
check(`the rules read every example (${examples}), its controls (${controls}) and its text (${texts})`, examples > 0 && controls > 0 && texts > 0, true);

await quit();
done('rules');
