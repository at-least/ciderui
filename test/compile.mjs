import { execFileSync } from 'node:child_process';
import { mkdtempSync, readFileSync, readdirSync, rmSync, writeFileSync } from 'node:fs';
import { tmpdir } from 'node:os';
import { dirname, join } from 'node:path';
import { fileURLToPath } from 'node:url';

export const ROOT = join(dirname(fileURLToPath(import.meta.url)), '..');

// @invariant: sorted by name across the sets: the order the tests' stylesheets import
// them in
export const COMPONENTS = ['site', 'console']
  .flatMap((set) => readdirSync(join(ROOT, 'components', set)).map((name) => ({ set, name, dir: `components/${set}/${name}` })))
  .sort((a, b) => (a.name < b.name ? -1 : 1));
export const componentDir = (name) => COMPONENTS.find((c) => c.name === name).dir;

// @invariant: the entry README.md tells a site to write, read from there so the two cannot
// part: its own path to utilities.css, no @source of its pages
export const UTILITIES_ENTRY = (utilities) => {
  const [, block] = readFileSync(join(ROOT, 'README.md'), 'utf8').match(/```css\n([^]*?)```/);
  if (!block.includes('../target/ciderui/utilities.css')) throw new Error('README.md writes no utilities.css');
  return block.replace('../target/ciderui/utilities.css', utilities).replace(/^@source .*\n/m, '');
};

// @external: entry, a file of its own: `@import` resolves from where it sits. Minified, as a
// site ships it.
export function tailwind(entry) {
  const dir = mkdtempSync(join(tmpdir(), 'ciderui-out-'));
  try {
    const out = join(dir, 'out.css');
    execFileSync(join(ROOT, 'node_modules/.bin/tailwindcss'), ['-i', entry, '-o', out, '--minify'], { stdio: 'pipe' });
    return readFileSync(out, 'utf8');
  } finally {
    rmSync(dir, { recursive: true });
  }
}

// @invariant: css: the entry after Tailwind's import, its @import paths relative to the
// root; classes: class strings for Tailwind to find; sources: paths relative
// to the root Tailwind scans, as a site's @source does
export function compile({ css = '@import "index.css";', classes = '', sources = [] } = {}) {
  const dir = mkdtempSync(join(tmpdir(), 'ciderui-test-'));
  try {
    const entry = join(dir, 'entry.css');
    const imports = css.replace(/@import "([^"]+)"/g, (_, p) => `@import "${join(ROOT, p)}"`);
    const scanned = sources.map(s => `@source "${join(ROOT, s)}";\n`).join('');
    writeFileSync(entry, `@import "${join(ROOT, 'node_modules/tailwindcss/index.css')}" source(none);\n@source inline(${JSON.stringify(classes)});\n${scanned}${imports}\n`);
    return tailwind(entry);
  } finally {
    rmSync(dir, { recursive: true });
  }
}

// @invariant: what a site's page links: scripts/build-css.mjs --check, a step before the
// tests that read it, holds it to the sources
export const shipped = (set) => readFileSync(join(ROOT, 'dist', `${set}.css`), 'utf8');

export const exampleFiles = (part) => readdirSync(join(ROOT, part.dir, 'examples')).filter((f) => f.endsWith('.html')).sort();

// @invariant: as its set's page holds it
export const exampleHtml = (part, file) => {
  const example = readFileSync(join(ROOT, part.dir, 'examples', file), 'utf8');
  return part.set === 'console' ? `<div class="cui-console-page">${example}</div>` : example;
};
