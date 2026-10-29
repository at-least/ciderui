// Runs a command in test/Dockerfile's image, the only place the browser
// starts (test/chromium.mjs): the repository copied to /src/ciderui (the
// image's node_modules kept), ciderui-zola (the docs' theme, a sibling
// checkout) to /src/ciderui-zola, CASES and UPDATE_GOLDEN passed through; then every file the command wrote that git
// does not ignore (golden.json, the docs' copies and pictures, a probe's
// results) copied back into the checkout. Nothing is bind-mounted: the
// Docker daemon may be on another machine (DOCKER_HOST=ssh://…), where a
// bind mount would read that machine's paths.
//
//   node scripts/docker.mjs                          the browser steps of npm test
//   node scripts/docker.mjs node test/components.mjs a command
//   UPDATE_GOLDEN=1 node scripts/docker.mjs node test/components.mjs

import { execFileSync, execSync, spawnSync } from 'node:child_process';
import { mkdtempSync, realpathSync, rmSync } from 'node:fs';
import { tmpdir } from 'node:os';
import { join } from 'node:path';
import { ROOT } from '../test/compile.mjs';

const IMAGE = 'ciderui-test';
execFileSync('docker', ['build', '-q', '-t', IMAGE, '-f', join(ROOT, 'test/Dockerfile'), ROOT], { stdio: ['ignore', 'ignore', 'inherit'] });

const STEPS = 'node test/tokens.mjs && node test/type.mjs && node test/page.mjs && node test/components.mjs && node scripts/build-lucide.mjs --check && node scripts/build-docs.mjs --check && node scripts/docs-pictures.mjs --check && node test/docs.mjs';
const command = process.argv.length > 2 ? process.argv.slice(2) : ['sh', '-c', STEPS];
const theme = realpathSync(join(ROOT, 'docs/themes/ciderui-zola'));
const pass = ['CASES', 'UPDATE_GOLDEN'].filter((k) => process.env[k] !== undefined).flatMap((k) => ['-e', `${k}=${process.env[k]}`]);
// The image's user (root): what it writes comes back through tar as the
// caller's own, so no --user (the repository's top, the image's WORKDIR, is
// root's). The stamp before the command: what is newer after it, the command wrote
// (the copied files keep the checkout's mtimes)
const wrap = 'touch /tmp/stamp && "$@"; status=$?; find . -type f -newer /tmp/stamp -not -path "./node_modules/*" > /tmp/written && tar -cf /tmp/written.tar -T /tmp/written; exit $status';
const id = execFileSync('docker', [
  'create', '--init', '--ipc=host', ...pass,
  IMAGE, 'sh', '-c', wrap, 'run', ...command,
], { encoding: 'utf8' }).trim();
const q = (s) => `'${s.replaceAll("'", "'\\''")}'`;
// no extended attributes (docker cp refuses macOS's com.apple.provenance,
// "operation not supported", on a file that carries it), no AppleDouble
// ._ files
const TAR_ENV = { ...process.env, COPYFILE_DISABLE: '1' };

let status = 1;
try {
  const exclude = '--exclude ./node_modules --exclude ./rust/target --exclude ./.git';
  execSync(`tar -c --no-xattrs -C ${q(ROOT)} ${exclude} . | docker cp - ${id}:/src/ciderui`, { env: TAR_ENV, stdio: ['ignore', 'ignore', 'inherit'] });
  execSync(`tar -c --no-xattrs -C ${q(theme)} --exclude ./.git . | docker cp - ${id}:/src/ciderui-zola`, { env: TAR_ENV, stdio: ['ignore', 'ignore', 'inherit'] });
  try {
    execFileSync('docker', ['start', '-a', id], { stdio: 'inherit' });
    status = 0;
  } catch (e) {
    status = e.status ?? 1;
  }
  const out = mkdtempSync(join(tmpdir(), 'ciderui-written-'));
  try {
    execSync(`docker cp ${id}:/tmp/written.tar - | tar -xOf - written.tar | tar -xf - -C ${q(out)}`, { stdio: ['ignore', 'ignore', 'inherit'] });
    const written = execSync('find . -type f', { cwd: out, encoding: 'utf8' }).split('\n').filter(Boolean).map((f) => f.replace(/^\.\//, ''));
    // git's own ignore list decides what comes back (docs/public, zola's
    // processed images stay in the container)
    // check-ignore exits 1 when it ignores none of them, 0 when some
    const check = spawnSync('git', ['check-ignore', '--no-index', '--stdin'], { cwd: ROOT, input: written.join('\n'), encoding: 'utf8' });
    if (check.status > 1) throw new Error(`git check-ignore: ${check.stderr}`);
    const ignored = new Set(check.stdout.split('\n').filter(Boolean));
    const kept = written.filter((f) => !ignored.has(f));
    if (kept.length) {
      execSync(`tar -c -C ${q(out)} -T - | tar -xf - -C ${q(ROOT)}`, { input: kept.join('\n'), stdio: ['pipe', 'ignore', 'inherit'] });
      console.log(`docker.mjs: ${kept.length} written file(s) copied back`);
    }
  } finally {
    rmSync(out, { recursive: true, force: true });
  }
} finally {
  execFileSync('docker', ['rm', '-f', id], { stdio: 'ignore' });
}
process.exit(status);
