// Every console part's example at 390px on an emulated touch screen
// (test/browser.mjs's `touch`): each control a finger can press, its box
// and its target (the larger of its box and a ::before or ::after it
// draws), held to the HIG's control sizes (Accessibility, Mobility:
// 44x44pt the default on iOS, 28x28pt the minimum). Prints one line per
// control and exits 1 when a target is under the minimum.
//
//   node scripts/docker.mjs node reference/touch/audit.mjs

import { readdirSync, readFileSync } from 'node:fs';
import { join } from 'node:path';
import { ROOT, compile } from '../../test/compile.mjs';
import { open, quit } from '../../test/browser.mjs';

const parts = readdirSync(join(ROOT, 'components')).filter((d) => d.startsWith('console-')).sort();
const sheets = parts.map((d) => `@import "components/${d}/${d}.css";`).join('\n');
const css = compile({ css: `@import "index.css";\n${sheets}`, sources: ['components'] });
let short = 0;
for (const part of parts) {
  const dir = join(ROOT, 'components', part, 'examples');
  for (const file of readdirSync(dir).filter((f) => f.endsWith('.html')).sort()) {
    const page = await open(css, readFileSync(join(dir, file), 'utf8'), { width: 390, touch: true });
    const rows = await page.page.evaluate(() => {
      const sel = 'a[href], button, input:not([type=hidden]), select, textarea, summary, [role=button], [tabindex]:not([tabindex="-1"])';
      return [...document.querySelectorAll(sel)].filter((e) => e.getClientRects().length).map((e) => {
        const r = e.getBoundingClientRect();
        let w = r.width, h = r.height;
        for (const p of ['::before', '::after']) {
          const cs = getComputedStyle(e, p);
          if (cs.content === 'none' || cs.position !== 'absolute') continue;
          w = Math.max(w, parseFloat(cs.width) || 0);
          h = Math.max(h, parseFloat(cs.height) || 0);
        }
        // a native checkbox or radio inside a <label>: the label is its target
        const label = (e.type === 'checkbox' || e.type === 'radio') && e.closest('label');
        if (label) { const l = label.getBoundingClientRect(); w = Math.max(w, l.width); h = Math.max(h, l.height); }
        const name = e.tagName.toLowerCase() + (e.className && typeof e.className === 'string' ? '.' + e.className.trim().split(/\s+/)[0] : '') + (e.type && e.tagName === 'INPUT' ? `[${e.type}]` : '');
        return { name, box: [+r.width.toFixed(1), +r.height.toFixed(1)], target: [+w.toFixed(1), +h.toFixed(1)] };
      });
    });
    await page.close();
    for (const r of rows) {
      const min = Math.min(...r.target);
      const verdict = min >= 44 ? 'ok   ' : min >= 28 ? 'MIN  ' : 'SHORT';
      if (min < 28) short += 1;
      console.log(`${verdict} ${part}/${file} ${r.name} box ${r.box.join('x')} target ${r.target.join('x')}`);
    }
  }
}
await quit();
console.log(`${short} under 28pt`);
process.exit(short ? 1 : 0);
