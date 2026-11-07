// @trap: Every example on every surface, not listed cases: six times a site built
// on ciderui (crashcart) found a failing pair the cases had not listed.

import { readFileSync } from 'node:fs';
import { createRequire } from 'node:module';
import { join } from 'node:path';
import { COMPONENTS, ROOT, exampleFiles, shipped } from './compile.mjs';
import { open, quit } from './browser.mjs';
import { check, done } from './check.mjs';

const AXE = readFileSync(createRequire(import.meta.url).resolve('axe-core/axe.min.js'), 'utf8');
const RULES = ['color-contrast', 'link-in-text-block', 'button-name', 'link-name', 'aria-valid-attr-value'];

const SURFACES = {
  site: {
    page: (html) => `<div class="cui-light">${html}</div>`,
    dark: (html) => `<div class="cui-dark">${html}</div>`,
  },
  console: {
    page: (html) => `<div class="cui-console-page">${html}</div>`,
    dark: (html) => `<div class="cui-console-page cui-dark">${html}</div>`,
    card: (html) => `<div class="cui-console-page"><div class="cui-console-card">${html}</div></div>`,
    'card dark': (html) => `<div class="cui-console-page cui-dark"><div class="cui-console-card">${html}</div></div>`,
  },
};

for (const c of COMPONENTS) {
  for (const file of exampleFiles(c)) {
    const example = readFileSync(join(ROOT, c.dir, 'examples', file), 'utf8');
    for (const [surface, wrap] of Object.entries(SURFACES[c.set])) {
      const page = await open(shipped(c.set), wrap(example));
      await page.settle();
      await page.page.addScriptTag({ content: AXE });
      const found = await page.page.evaluate(async (rules) => {
        const r = await axe.run(document, { runOnly: { type: 'rule', values: rules }, resultTypes: ['violations'] });
        return r.violations.flatMap((v) => v.nodes.map((n) => `${v.id} ${n.target.join(' ')}: ${n.failureSummary.replace(/\s+/g, ' ')}`));
      }, RULES);
      await page.close();
      check(`${c.name}/${file} on the ${surface}`, found.join(' | '), '');
    }
  }
}
await quit();
done('surfaces');
