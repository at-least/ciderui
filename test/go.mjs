import { execFileSync } from 'node:child_process';
import { mkdtempSync, readdirSync, readFileSync, rmSync, writeFileSync } from 'node:fs';
import { tmpdir } from 'node:os';
import { join } from 'node:path';
import { ROOT, UTILITIES_ENTRY, tailwind } from './compile.mjs';

// @trap: the module's package directories, not `.`: gofmt would walk node_modules,
// which go.mod's ignore keeps from go's own commands
const dirs = execFileSync('go', ['list', '-f', '{{.Dir}}', './...'], { cwd: ROOT, encoding: 'utf8' }).trim().split('\n');
const unformatted = execFileSync('gofmt', ['-l', ...dirs], { cwd: ROOT, encoding: 'utf8' }).trim();
if (unformatted) throw new Error(`not gofmt'd (gofmt -w <file>):\n${unformatted}`);
for (const args of [
  ['vet', './...'],
  // @invariant: a component's field left out of a struct literal is a default nobody chose
  ['tool', 'exhaustruct', '-i', 'github.com/at-least/ciderui/.*', './...'],
  ['test', './...'],
]) {
  execFileSync('go', args, { cwd: ROOT, stdio: 'inherit' });
}
// @invariant: what a site with utilities of its own does: its Go tool writes utilities.css,
// Tailwind builds its second stylesheet from there. Everything else reads the
// checkout, so a stylesheet left out of go:embed fails here and nowhere else.
// Written outside this repository, as a site's target/ is beside its node_modules
// and not under them: a bare @import in a written file would resolve from there.
const written = mkdtempSync(join(tmpdir(), 'written-kit-'));
const site = mkdtempSync(join(ROOT, 'site-entry-'));
try {
  execFileSync('go', ['run', './go/cmd/ciderui-css', written], { cwd: ROOT, stdio: 'pipe' });
  const entry = join(site, 'entry.css');
  writeFileSync(entry, `${UTILITIES_ENTRY(`${written}/utilities.css`)}\n@source inline("px-8 md:cui-headline");\n`);
  const built = tailwind(entry);
  for (const rule of ['.px-8', '.md\\:cui-headline']) if (!built.includes(rule)) throw new Error(`utilities.css written by ciderui-css built with no ${rule}: ${built.slice(0, 200)}`);
  if (built.includes('.cui-btn')) throw new Error('utilities.css built with a part');
} finally {
  rmSync(written, { recursive: true });
  rmSync(site, { recursive: true });
}
// @invariant: a part's every branch has a test: a new panic with no Break, or a new
// branch no example draws, stops here
const covered = execFileSync('go', ['test', '-cover', './go/site', './go/console'], { cwd: ROOT, encoding: 'utf8' });
for (const line of covered.trim().split('\n')) {
  if (!line.includes('coverage: 100.0% of statements')) throw new Error(`a statement no test runs (go test -coverprofile): ${line}`);
}
// @trap: the button's variants a compile error, not a string: btn.go alone writes the
// class. play.go's two glyphs and labels and disclosure's <summary> are no <button>
// the type can hold.
const OWN_BUTTON = /(^|\/)(btn|play|disclosure)\.go$/;
for (const dir of ['go/site', 'go/console']) {
  for (const file of readdirSync(join(ROOT, dir)).filter((f) => f.endsWith('.go') && !f.endsWith('_test.go') && !OWN_BUTTON.test(f))) {
    if (/"cui-(console-)?btn[" ]/.test(readFileSync(join(ROOT, dir, file), 'utf8'))) throw new Error(`${dir}/${file} writes a button's class: use its Btn type`);
  }
}
console.log('go: gofmt, vet, exhaustruct and tests clean');
