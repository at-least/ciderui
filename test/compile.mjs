// Compiles a ciderui entry stylesheet through the Tailwind CLI, the way a
// site compiles it, minified as a site ships it, into a temp dir nothing
// else sees. The entry is written there too, from the imports and sources
// a test names, so a test needs no file of its own.

import { execFileSync } from 'node:child_process';
import { mkdtempSync, readFileSync, rmSync, writeFileSync } from 'node:fs';
import { tmpdir } from 'node:os';
import { dirname, join } from 'node:path';
import { fileURLToPath } from 'node:url';

export const ROOT = join(dirname(fileURLToPath(import.meta.url)), '..');

// css: the entry's text after Tailwind's import (paths relative to the
// root); classes: the class strings Tailwind is to find; sources: files or
// directories (relative to the root) it scans for more, as a site's
// @source does
export function compile({ css = '@import "index.css";', classes = '', sources = [] } = {}) {
  const dir = mkdtempSync(join(tmpdir(), 'ciderui-test-'));
  try {
    const entry = join(dir, 'entry.css');
    const imports = css.replace(/@import "([^"]+)"/g, (_, p) => `@import "${join(ROOT, p)}"`);
    const scanned = sources.map(s => `@source "${join(ROOT, s)}";\n`).join('');
    writeFileSync(entry, `@import "${join(ROOT, 'node_modules/tailwindcss/index.css')}" source(none);\n@source inline(${JSON.stringify(classes)});\n${scanned}${imports}\n`);
    const out = join(dir, 'out.css');
    execFileSync(join(ROOT, 'node_modules/.bin/tailwindcss'), ['-i', entry, '-o', out, '--minify'], { stdio: 'pipe' });
    return readFileSync(out, 'utf8');
  } finally {
    rmSync(dir, { recursive: true });
  }
}
