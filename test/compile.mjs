import { execFileSync } from 'node:child_process';
import { mkdtempSync, readFileSync, readdirSync, rmSync, writeFileSync } from 'node:fs';
import { tmpdir } from 'node:os';
import { dirname, join } from 'node:path';
import { fileURLToPath } from 'node:url';

export const ROOT = join(dirname(fileURLToPath(import.meta.url)), '..');

// sorted by name across the sets: the order the tests' and the docs'
// stylesheets import them in
export const COMPONENTS = ['site', 'console', 'docs']
  .flatMap((set) => readdirSync(join(ROOT, 'components', set)).map((name) => ({ set, name, dir: `components/${set}/${name}` })))
  .sort((a, b) => (a.name < b.name ? -1 : 1));
export const componentDir = (name) => COMPONENTS.find((c) => c.name === name).dir;

// css: the entry after Tailwind's import, its @import paths relative to the
// root; classes: class strings for Tailwind to find; sources: paths relative
// to the root Tailwind scans, as a site's @source does
export function compile({ css = '@import "index.css";', classes = '', sources = [] } = {}) {
  const dir = mkdtempSync(join(tmpdir(), 'ciderui-test-'));
  try {
    const entry = join(dir, 'entry.css');
    const imports = css.replace(/@import "([^"]+)"/g, (_, p) => `@import "${join(ROOT, p)}"`);
    const scanned = sources.map(s => `@source "${join(ROOT, s)}";\n`).join('');
    writeFileSync(entry, `@import "${join(ROOT, 'node_modules/tailwindcss/index.css')}" source(none);\n@source inline(${JSON.stringify(classes)});\n${scanned}${imports}\n`);
    const out = join(dir, 'out.css');
    // minified, as a site ships it
    execFileSync(join(ROOT, 'node_modules/.bin/tailwindcss'), ['-i', entry, '-o', out, '--minify'], { stdio: 'pipe' });
    return readFileSync(out, 'utf8');
  } finally {
    rmSync(dir, { recursive: true });
  }
}
