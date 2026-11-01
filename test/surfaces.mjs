// Every component's examples on every surface its set's parts sit on, read
// by axe-core's rules for what a combination gets wrong: text against what
// is behind it (color-contrast) and a link inside text told apart by more
// than color (link-in-text-block). A check per example and surface, failing
// with each node axe-core finds. Combinations, not cases: six times a site
// built on ciderui (crashcart) found such a pair the cases had not listed.

import { readdirSync, readFileSync } from 'node:fs';
import { createRequire } from 'node:module';
import { join } from 'node:path';
import { COMPONENTS, ROOT, compile } from './compile.mjs';
import { open, quit } from './browser.mjs';
import { check, done } from './check.mjs';

const AXE = readFileSync(createRequire(import.meta.url).resolve('axe-core/axe.min.js'), 'utf8');
const RULES = ['color-contrast', 'link-in-text-block'];

// each set's surfaces: the markup an example is set into
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
  docs: {
    page: (html) => `<div class="cui-light">${html}</div>`,
    dark: (html) => `<div class="cui-dark">${html}</div>`,
  },
};

const sheets = COMPONENTS.map((c) => `@import "${c.dir}/${c.name}.css";`).join('\n');
const css = compile({ css: `@import "fonts.css";\n@import "index.css";\n${sheets}`, sources: ['components'] });

for (const c of COMPONENTS) {
  const dir = join(ROOT, c.dir, 'examples');
  for (const file of readdirSync(dir).filter((f) => f.endsWith('.html')).sort()) {
    const example = readFileSync(join(dir, file), 'utf8');
    for (const [surface, wrap] of Object.entries(SURFACES[c.set])) {
      const page = await open(css, wrap(example));
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
