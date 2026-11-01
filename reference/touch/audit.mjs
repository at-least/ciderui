// Every console part's example at 390px on an emulated touch screen
// (test/browser.mjs's `touch`): each control a finger can press, its box
// and its target (hit-tested from its center, below), held to the HIG's
// control sizes (Accessibility, Mobility: 44x44pt the default on iOS,
// 28x28pt the minimum). Prints one line per control and exits 1 when a
// target is under the default.
//
//   node scripts/docker.mjs node reference/touch/audit.mjs

import { readdirSync, readFileSync } from 'node:fs';
import { join } from 'node:path';
import { COMPONENTS, ROOT, compile } from '../../test/compile.mjs';
import { open, quit } from '../../test/browser.mjs';

const parts = COMPONENTS.filter((c) => c.set === 'console');
// an empty list would print "0 under 28pt" and pass having checked nothing
if (!parts.length) throw new Error('no console parts in COMPONENTS');
const sheets = parts.map((c) => `@import "${c.dir}/${c.name}.css";`).join('\n');
const css = compile({ css: `@import "index.css";\n${sheets}`, sources: ['components'] });
let short = 0;
for (const { name: part, dir: partDir } of parts) {
  const dir = join(ROOT, partDir, 'examples');
  for (const file of readdirSync(dir).filter((f) => f.endsWith('.html')).sort()) {
    // room round the example, so a target at its edge is not read short
    // where it runs out of the window (a tap there lands nowhere)
    const html = `<div style="padding: 48px 24px">${readFileSync(join(dir, file), 'utf8')}</div>`;
    const page = await open(css, html, { width: 390, touch: true });
    const rows = await page.page.evaluate(() => {
      const sel = 'a[href], button, input:not([type=hidden]), select, textarea, summary, [role=button], [tabindex]:not([tabindex="-1"])';
      return [...document.querySelectorAll(sel)].filter((e) => e.checkVisibility()).map((e) => {
        const r = e.getBoundingClientRect();
        // a native checkbox or radio inside a <label>: the label is its target
        const label = (e.type === 'checkbox' || e.type === 'radio') && e.closest('label');
        const anchor = label || e;
        // a control hidden until it is focused (the skip link): read focused
        if (!label && r.width <= 1 && r.height <= 1) e.focus();
        anchor.scrollIntoView({ block: 'center', inline: 'center' });
        const a = anchor.getBoundingClientRect();
        const x = a.left + a.width / 2, y = a.top + a.height / 2;
        // the target as a finger meets it: from the anchor's center, each
        // way, as far as a tap still lands on the control (a ::after drawn
        // round it counts; one clipped by a scroll box or under a
        // neighbour's does not), every 0.5px up to 40px: each way reads up
        // to 0.5px short, so 43 holds a 44px target
        const ours = (px, py) => {
          if (px < 0 || py < 0 || px >= innerWidth || py >= innerHeight) return false;
          const hit = document.elementFromPoint(px, py);
          return Boolean(hit) && (hit === e || e.contains(hit) || anchor.contains(hit));
        };
        const reach = (dx, dy) => {
          let d = 0;
          while (d < 40 && ours(x + dx * (d + 0.5), y + dy * (d + 0.5))) d += 0.5;
          return d;
        };
        const w = ours(x, y) ? reach(-1, 0) + reach(1, 0) : 0;
        const h = ours(x, y) ? reach(0, -1) + reach(0, 1) : 0;
        const name = e.tagName.toLowerCase() + (e.className && typeof e.className === 'string' ? '.' + e.className.trim().split(/\s+/)[0] : '') + (e.type && e.tagName === 'INPUT' ? `[${e.type}]` : '');
        return { name, box: [+r.width.toFixed(1), +r.height.toFixed(1)], target: [w, h] };
      });
    });
    await page.close();
    for (const r of rows) {
      const min = Math.min(...r.target);
      const verdict = min >= 43 ? 'ok   ' : min >= 28 ? 'MIN  ' : 'SHORT';
      if (min < 43) short += 1;
      console.log(`${verdict} ${part}/${file} ${r.name} box ${r.box.join('x')} target ${r.target.join('x')}`);
    }
  }
}
await quit();
console.log(`${short} under 44pt`);
process.exit(short ? 1 : 0);
