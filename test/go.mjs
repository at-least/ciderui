import { execFileSync } from 'node:child_process';
import { mkdtempSync, rmSync } from 'node:fs';
import { join, relative } from 'node:path';
import { compile, ROOT } from './compile.mjs';

// the module's package directories, not `.`: gofmt would walk node_modules,
// which go.mod's ignore keeps from go's own commands
const dirs = execFileSync('go', ['list', '-f', '{{.Dir}}', './...'], { cwd: ROOT, encoding: 'utf8' }).trim().split('\n');
const unformatted = execFileSync('gofmt', ['-l', ...dirs], { cwd: ROOT, encoding: 'utf8' }).trim();
if (unformatted) throw new Error(`not gofmt'd (gofmt -w <file>):\n${unformatted}`);
for (const args of [
  ['vet', './...'],
  // a component's field left out of a struct literal is a default nobody chose
  ['tool', 'exhaustruct', '-i', 'github.com/at-least/ciderui/.*', './...'],
  ['test', './...'],
]) {
  execFileSync('go', args, { cwd: ROOT, stdio: 'inherit' });
}
// what a site does: its Go tool writes the embedded kit, Tailwind builds each set
// from there. Everything else reads the checkout, so a stylesheet left out of
// go:embed fails here and nowhere else.
const written = mkdtempSync(join(ROOT, 'written-kit-'));
try {
  execFileSync('go', ['run', './go/cmd/ciderui-css', written], { cwd: ROOT, stdio: 'pipe' });
  for (const [set, rule] of [['site', '.cui-btn'], ['console', '.cui-console-btn']]) {
    const built = compile({ css: `@import "${relative(ROOT, join(written, `${set}.css`))}";` });
    if (!built.includes(rule)) throw new Error(`${set}.css written by ciderui-css built with no ${rule}`);
  }
} finally {
  rmSync(written, { recursive: true });
}
console.log('go: gofmt, vet, exhaustruct and tests clean');
