// @external: Runs a probe on air (the owner's Mac that nobody sits at, `ssh air`), whose
// Safari is driven through safaridriver: never the Safari of the Mac in front of
// someone. The repository goes to the same path on air (no node_modules, .git or
// .claude: the probes use node's own modules and air's Safari), SITE and INTER too
// where they name a path; the command runs in this directory with BROWSER, SITE,
// MODE, EXPECT and INTER passed through; the directory's results/ comes back.
// air's Safari takes automation (Develop > Allow Remote Automation, on).
//
//   cd reference/console && node build.mjs /tmp/ciderui-console &&
//     BROWSER=safari SITE=/tmp/ciderui-console node ../air.mjs node run.mjs
//
// air has no Xcode: BROWSER=ios and the XCUITest probes stay on a Mac with it.

import { execFileSync, spawnSync } from 'node:child_process';
import { existsSync, mkdirSync, statSync } from 'node:fs';
import { hostname } from 'node:os';
import { dirname, isAbsolute, join, resolve } from 'node:path';

const HOST = 'air';
const ROOT = resolve(new URL('..', import.meta.url).pathname);
const command = process.argv.slice(2);
if (!command.length) throw new Error('usage: node reference/air.mjs <command> [args…]');
// @external: air has Safari alone: Playwright's browsers are not synced, and the
// owner's call is that Chrome runs in Docker (scripts/docker.mjs)
if (process.env.BROWSER !== undefined && process.env.BROWSER !== 'safari') throw new Error(`BROWSER=${process.env.BROWSER} does not run on air, which has Safari alone: Chrome, Chromium and Firefox run in Docker (scripts/docker.mjs), the iOS simulator on a Mac with Xcode`);
const cwd = process.cwd();
if (!cwd.startsWith(ROOT)) throw new Error(`run it from inside the repository (${ROOT}), not ${cwd}`);

const quote = (s) => `'${String(s).replaceAll("'", "'\\''")}'`;
const ssh = (script) => execFileSync('ssh', [HOST, script], { stdio: 'inherit' });
// @trap: a directory's trailing slash: rsync copies its contents into the same path,
// not a copy of it inside; the parent made first, openrsync having no --mkpath
const sync = (from, to, dir, excludes = []) => {
  execFileSync('rsync', ['-a', '--delete', ...excludes.flatMap((e) => ['--exclude', e]), dir ? `${from}/` : from, dir ? `${to}/` : to], { stdio: 'inherit' });
};

const pass = ['BROWSER', 'SITE', 'MODE', 'EXPECT', 'INTER']
  .filter((k) => process.env[k] !== undefined)
  .map((k) => `${k}=${quote(process.env[k])}`);
const paths = [];
for (const k of ['SITE', 'INTER']) {
  const p = process.env[k];
  if (p === undefined) continue;
  if (!isAbsolute(p) || !existsSync(p)) throw new Error(`${k} must name an existing absolute path: ${p}`);
  paths.push(p);
}

// @trap: one run at a time: safaridriver drives one Safari, and each run quits
// Safari first, which would end another run's session. A lock left by a run cut
// off (its ssh dropped) names the run that made it; remove it with
// ssh air rmdir /tmp/ciderui-air.lock
const LOCK = '/tmp/ciderui-air.lock';
ssh(`mkdir ${LOCK} 2>/dev/null || { echo "air is busy with another probe (${LOCK}: $(cat ${LOCK}/owner 2>/dev/null))" >&2; exit 1; }; echo ${quote(`${hostname()} pid ${process.pid} ${cwd}`)} > ${LOCK}/owner`);
let status = 1;
try {
  ssh(`mkdir -p ${[ROOT, ...paths.map((p) => dirname(p))].map(quote).join(' ')}`);
  sync(ROOT, `${HOST}:${ROOT}`, true, ['node_modules', '.git', '.claude']);
  for (const p of paths) sync(p, `${HOST}:${p}`, statSync(p).isDirectory());
  // @external: Safari 27 takes no session in a Safari it did not launch ("Safari was
  // not launched for automation"): quit it first. caffeinate -du: with air's display
  // asleep, Safari's picture came back all black. ssh's PATH lacks Homebrew's node
  const run = spawnSync('ssh', [HOST, [
    'pkill -x Safari; while pgrep -x Safari >/dev/null; do sleep 0.2; done;',
    `cd ${quote(cwd)} &&`,
    `exec env PATH=/opt/homebrew/bin:$PATH ${pass.join(' ')} caffeinate -du ${command.map(quote).join(' ')}`,
  ].join(' ')], { stdio: 'inherit' });
  status = run.status ?? 1;
  const results = join(cwd, 'results');
  if (spawnSync('ssh', [HOST, `test -d ${quote(results)}`]).status === 0) {
    mkdirSync(results, { recursive: true });
    sync(`${HOST}:${results}`, results, true);
  }
} finally {
  ssh(`rm -rf ${LOCK}`);
}
process.exit(status);
