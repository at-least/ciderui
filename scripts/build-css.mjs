// @invariant: dist/<set>.css, the stylesheet a site links: Tailwind's base, Inter, the set's
// parts, every class of the type and every utility an example writes, so a page
// copying an example needs nothing more. Any other utility is the site's own
// stylesheet's (utilities.css).

import { readFileSync, writeFileSync } from 'node:fs';
import { join } from 'node:path';
import { COMPONENTS, ROOT, compile } from '../test/compile.mjs';

const args = process.argv.slice(2);
if (args.length > 1 || (args.length === 1 && args[0] !== '--check')) {
  console.error('usage: node scripts/build-css.mjs [--check]');
  process.exit(2);
}
const checking = args.length === 1;

const read = (path) => readFileSync(join(ROOT, path), 'utf8');

const build = (set) => {
  const parts = COMPONENTS.filter((c) => c.set === set);
  // @invariant: COMPONENTS' name order: a part's rule over another's of the same
  // specificity relies on it
  const partSheets = parts.map((c) => `${c.dir}/${c.name}.css`);
  const classes = ['type.css', ...partSheets].flatMap((sheet) => [...read(sheet).matchAll(/^@utility (cui-[\w-]+)/gm)].map((m) => m[1]));
  const imports = ['fonts.css', 'index.css', ...partSheets, `${set}.css`];
  return compile({ css: imports.map((sheet) => `@import "${sheet}";`).join('\n'), classes: classes.join(' '), sources: parts.map((c) => `${c.dir}/examples`) });
};

let stale = false;
for (const set of ['site', 'console']) {
  const path = `dist/${set}.css`;
  const css = build(set);
  if (!checking) writeFileSync(join(ROOT, path), css);
  else if (read(path) !== css) {
    console.log(`${path} is stale: node scripts/build-css.mjs`);
    stale = true;
  }
}
process.exit(stale ? 1 : 0);
