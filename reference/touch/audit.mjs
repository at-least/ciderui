// @invariant: Every component's example at 390px on an emulated touch screen (SETS=…
// for some of the sets) (test/browser.mjs's `touch`): each control a finger can press,
// its box and its target (hit-tested from its center, below), held to the HIG's control
// sizes (Accessibility, Mobility: 44x44pt the default on iOS, 28x28pt the minimum).
// Prints a line for each control under the default (each control with VERBOSE=1) and
// exits 1 when there is one; a step of npm test (scripts/docker.mjs).
//
//   node scripts/docker.mjs node reference/touch/audit.mjs

import { readFileSync } from 'node:fs';
import { join } from 'node:path';
import { COMPONENTS, ROOT, exampleFiles, shipped } from '../../test/compile.mjs';
import { open, quit } from '../../test/browser.mjs';

const parts = process.env.SETS ? COMPONENTS.filter((c) => process.env.SETS.split(',').includes(c.set)) : COMPONENTS;
// @trap: an empty list would print "0 under 44pt" and pass having checked nothing
if (!parts.length) throw new Error('no parts in COMPONENTS for these sets');
let short = 0;
let controls = 0;
for (const partDef of parts) {
  const { name: part, dir: partDir } = partDef;
  const dir = join(ROOT, partDir, 'examples');
  for (const file of exampleFiles(partDef)) {
    // @external: room round the example, so a target at its edge is not read short
    // where it runs out of the window (a tap there lands nowhere)
    const html = `<div style="padding: 48px 24px">${readFileSync(join(dir, file), 'utf8')}</div>`;
    const page = await open(shipped(partDef.set), html, { width: 390, touch: true });
    const rows = await page.page.evaluate(() => {
      const sel = 'a[href], button, input:not([type=hidden]), select, textarea, summary, [role=button], [tabindex]:not([tabindex="-1"])';
      return [...document.querySelectorAll(sel)]
        // @external: what a finger can tap now: shown, not disabled (a gallery's paddle
        // at the row's end takes no pointer; a hidden toggle takes none
        // either, its label taking the tap)
        .filter((e) => e.checkVisibility({ visibilityProperty: true }) && !e.matches(':disabled') && (e.labels?.length || getComputedStyle(e).pointerEvents !== 'none'))
        // @external: a scroll box takes tabindex for the keyboard; a finger drags it and presses
        // nothing
        .filter((e) => e.matches('a[href], button, input, select, textarea, summary, [role=button]') || !/auto|scroll/.test(getComputedStyle(e).overflow)).map((e) => {
        const r = e.getBoundingClientRect();
        // @external: a link inside a sentence: WCAG 2.5.8's inline exception, listed,
        // not counted (its line's height is the text's)
        let block = e.parentElement;
        while (getComputedStyle(block).display === 'inline') block = block.parentElement;
        const inText = e.tagName === 'A' && getComputedStyle(e).display === 'inline' && block.textContent.replace(e.textContent, '').trim() !== '';
        // @external: its labels (wrapping it or naming it by for=) are its target too: a
        // tap on one acts on it (a hidden toggle's chevron, a floating label)
        const labels = [...(e.labels ?? [])].filter((l) => l.checkVisibility());
        const tiny = r.width <= 1 || r.height <= 1;
        const label = labels.length && (tiny || e.type === 'checkbox' || e.type === 'radio')
          ? labels.reduce((big, l) => (l.getBoundingClientRect().width * l.getBoundingClientRect().height > big.getBoundingClientRect().width * big.getBoundingClientRect().height ? l : big))
          : null;
        const anchor = label || e;
        // @external: a control hidden until it is focused (the skip link): read focused,
        // then blurred, as shown it covers the bar after it (the global nav's wordmark)
        const focused = !label && tiny;
        if (focused) e.focus();
        anchor.scrollIntoView({ block: 'center', inline: 'center' });
        const a = anchor.getBoundingClientRect();
        const x = a.left + a.width / 2, y = a.top + a.height / 2;
        // @trap: the target as a finger meets it: from the anchor's center, each
        // way, as far as a tap still lands on the control (a ::after drawn
        // round it counts; one clipped by a scroll box or under a
        // neighbour's does not), every 0.5px up to 40px: each way reads up
        // to 0.5px short, so 43 holds a 44px target
        const ours = (px, py) => {
          if (px < 0 || py < 0 || px >= innerWidth || py >= innerHeight) return false;
          const hit = document.elementFromPoint(px, py);
          return Boolean(hit) && (e.contains(hit) || labels.some((l) => l.contains(hit)));
        };
        const reach = (dx, dy) => {
          let d = 0;
          while (d < 40 && ours(x + dx * (d + 0.5), y + dy * (d + 0.5))) d += 0.5;
          return d;
        };
        const w = ours(x, y) ? reach(-1, 0) + reach(1, 0) : 0;
        const h = ours(x, y) ? reach(0, -1) + reach(0, 1) : 0;
        if (focused) e.blur();
        const name = e.tagName.toLowerCase() + (e.className && typeof e.className === 'string' ? '.' + e.className.trim().split(/\s+/)[0] : '') + (e.type && e.tagName === 'INPUT' ? `[${e.type}]` : '');
        return { name, inText, box: [+r.width.toFixed(1), +r.height.toFixed(1)], target: [w, h] };
      });
    });
    await page.close();
    for (const r of rows) {
      const min = Math.min(...r.target);
      const verdict = r.inText ? 'text ' : min >= 43 ? 'ok   ' : min >= 28 ? 'MIN  ' : 'SHORT';
      if (min < 43 && !r.inText) short += 1;
      if ((min < 43 && !r.inText) || process.env.VERBOSE) console.log(`${verdict} ${part}/${file} ${r.name} box ${r.box.join('x')} target ${r.target.join('x')}`);
      controls += 1;
    }
  }
}
await quit();
console.log(`touch: ${controls} controls, ${short} under 44pt`);
process.exit(short ? 1 : 0);
