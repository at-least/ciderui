import { spawnSync } from 'node:child_process';
import { mkdtempSync, readdirSync, readFileSync, rmSync } from 'node:fs';
import { tmpdir } from 'node:os';
import { join, relative } from 'node:path';
import { ROOT, ownSheet } from './compile.mjs';

const fail = (message) => {
  console.log(message);
  process.exit(1);
};
function run(command, ...args) {
  const { status, stdout, stderr, error } = spawnSync(command, args, { cwd: ROOT, encoding: 'utf8' });
  if (error) fail(`go: ${command}: ${error.message}`);
  if (status !== 0) {
    process.stdout.write(stdout + stderr);
    fail(`go: ${[command, ...args].join(' ')} failed (exit ${status})`);
  }
  return stdout;
}

// @invariant: every check that opens no browser runs here, before the image starts
// (the owner, 2026-11-14): a stale dist stops in a second, not after the browser steps
run('node', 'scripts/build-lucide.mjs', '--check');
run('node', 'scripts/build-css.mjs', '--check');
run('node', 'test/tokens.mjs');
// @trap: the module's Go files, not `.` nor a package's directory: gofmt walks a
// directory, the top's node_modules with it, which go.mod's ignore keeps from go's
// own commands
const each = (files) => `{{range .${files}}}{{$.Dir}}/{{.}}{{"\\n"}}{{end}}`;
const sources = run('go', 'list', '-f', each('GoFiles') + each('TestGoFiles') + each('XTestGoFiles'), './...').split('\n').filter(Boolean).map((file) => relative(ROOT, file));
const unformatted = run('gofmt', '-l', ...sources).trim();
if (unformatted) fail(`${unformatted}\ngo: not gofmt'd: gofmt -w <file>`);
run('go', 'vet', './...');
// @invariant: a component's field left out of a struct literal is a default nobody chose
run('go', 'run', './go/cmd/ciderui-check', './...');
// @invariant: a part's every branch has a test: a new panic with no Break, or a new
// branch no example draws, stops here
const tested = run('go', 'test', '-cover', './...').split('\n');
for (const set of ['go/site', 'go/console']) {
  const line = tested.find((l) => l.split('\t')[1] === `github.com/at-least/ciderui/${set}`);
  if (!line.includes('coverage: 100.0% of statements')) fail(`go: a statement of ${set} no test runs (go test -coverprofile): ${line}`);
}
// @invariant: what a site with utilities of its own does: its Go tool writes utilities.css,
// Tailwind builds its second stylesheet from there. Everything else reads the
// checkout, so a stylesheet left out of go:embed fails here and nowhere else.
// Written outside this repository, as a site's target/ is beside its node_modules
// and not under them: a bare @import in a written file would resolve from there.
const written = mkdtempSync(join(tmpdir(), 'written-kit-'));
process.on('exit', () => rmSync(written, { recursive: true }));
run('go', 'run', './go/cmd/ciderui-css', written);
const built = ownSheet(`${written}/utilities.css`, 'px-8 md:px-8 cui-btn cui-headline');
for (const rule of ['.px-8', '.md\\:px-8']) if (!built.includes(rule)) fail(`go: utilities.css written by ciderui-css built with no ${rule}: ${built.slice(0, 200)}`);
for (const kit of ['.cui-btn', '.cui-headline']) if (built.includes(kit)) fail(`go: utilities.css built the kit's ${kit}, which the kit's own sheet ships (utilities.css)`);
// @trap: the button's variants a compile error, not a string: btn.go alone writes the
// class. play_button.go's two glyphs and labels are no <button> the type can hold.
const OWN_BUTTON = /(^|\/)(btn|play_button)\.go$/;
for (const dir of ['go/site', 'go/console']) {
  for (const file of readdirSync(join(ROOT, dir)).filter((f) => f.endsWith('.go') && !f.endsWith('_test.go') && !OWN_BUTTON.test(f))) {
    if (/"cui-(console-)?btn[" ]/.test(readFileSync(join(ROOT, dir, file), 'utf8'))) fail(`go: ${dir}/${file} writes a button's class: use its Btn type`);
  }
}
console.log('go: lucide.go, dist, tokens, gofmt, vet, exhaustruct and tests clean');
