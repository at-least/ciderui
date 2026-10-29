// Runs a probe in a macOS VM of its own: `tester` (the owner's ~/bin/tester,
// its VMs orchard's, one fresh VM per run, deleted after), where Safari is
// driven through safaridriver as on a Mac but never on the Mac in front of
// someone: its Safari, its windows and its focus are left alone. The
// repository goes to the same path in the VM (no node_modules, .git or
// rust: the probes use node's own modules and the VM's Safari), SITE and
// INTER too where they name a path; the command runs in this directory
// with BROWSER, SITE, MODE, EXPECT and INTER passed through; the
// directory's results/ comes back. The VM's Safari takes automation
// (safaridriver --enable, which sudo needs and the VM's admin has).
//
//   cd reference/console && node build.mjs /tmp/ciderui-console &&
//     BROWSER=safari SITE=/tmp/ciderui-console node ../vm.mjs node run.mjs
//
// The VM has no Xcode: BROWSER=ios and the XCUITest probes stay on a Mac.

import { execFileSync, spawnSync } from 'node:child_process';
import { existsSync, readdirSync } from 'node:fs';
import { isAbsolute, join, resolve } from 'node:path';

const ROOT = resolve(new URL('..', import.meta.url).pathname);
const command = process.argv.slice(2);
if (!command.length) throw new Error('usage: node reference/vm.mjs <command> [args…]');
if (process.env.BROWSER === 'ios') throw new Error('BROWSER=ios needs Xcode, which the VM has not: run it on a Mac');
const cwd = process.cwd();
if (!cwd.startsWith(ROOT)) throw new Error(`run it from inside the repository (${ROOT}), not ${cwd}`);

const tester = (...args) => execFileSync('tester', args, { stdio: 'inherit' });
const name = `ciderui-probe-${process.pid}`;
const pass = ['BROWSER', 'SITE', 'MODE', 'EXPECT', 'INTER']
  .filter((k) => process.env[k] !== undefined)
  .map((k) => `${k}=${process.env[k]}`);
const paths = readdirSync(ROOT)
  .filter((e) => !['node_modules', '.git', 'rust'].includes(e))
  .map((e) => join(ROOT, e));
for (const k of ['SITE', 'INTER']) {
  const p = process.env[k];
  if (p === undefined) continue;
  if (!isAbsolute(p) || !existsSync(p)) throw new Error(`${k} must name an existing absolute path: ${p}`);
  paths.push(p);
}

tester('up', name);
let status = 1;
try {
  tester('push', name, ...paths);
  tester('exec', name, 'sudo', '-n', '/usr/bin/safaridriver', '--enable');
  const run = spawnSync('tester', ['exec', name, 'env', ...pass, 'sh', '-c', 'cd "$1" && shift && exec "$@"', 'run', cwd, ...command], { stdio: 'inherit' });
  status = run.status ?? 1;
  if (existsSync(join(cwd, 'results'))) tester('pull', name, join(cwd, 'results'));
} finally {
  tester('down', name);
}
process.exit(status);
