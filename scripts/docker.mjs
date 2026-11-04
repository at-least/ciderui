// Nothing is bind-mounted: the Docker daemon may be on another machine
// (DOCKER_HOST=ssh://…), where a bind mount would read that machine's paths.

import { execFileSync, execSync, spawnSync } from 'node:child_process';
import { createHash } from 'node:crypto';
import { mkdtempSync, readFileSync, realpathSync, rmSync } from 'node:fs';
import { tmpdir } from 'node:os';
import { join } from 'node:path';
import { ROOT } from '../test/compile.mjs';

const IMAGE = 'ciderui-test';
// The image is made from these alone (.dockerignore): one built from the same
// bytes carries their hash as a label and is run by its ID, with no build
// (a cached build still took 2s over ssh). Not by the tag: lab's daemon
// serves every Mac and checkout, and another's build can move the tag
// meanwhile.
const inputs = createHash('sha256');
for (const f of ['test/Dockerfile', 'package.json', 'package-lock.json']) inputs.update(f).update(readFileSync(join(ROOT, f)));
const label = `ciderui.inputs=${inputs.digest('hex')}`;
let image = execFileSync('docker', ['images', '-q', '--no-trunc', '--filter', `label=${label}`], { encoding: 'utf8' }).split('\n')[0].trim();
if (!image) image = execFileSync('docker', ['build', '-q', '-t', IMAGE, '--label', label, '-f', join(ROOT, 'test/Dockerfile'), ROOT], { encoding: 'utf8', stdio: ['ignore', 'pipe', 'inherit'] }).trim();
if (!/^sha256:[0-9a-f]{64}$/.test(image)) throw new Error(`docker gave no image ID: ${image}`);
// CIDERUI_CHROME=1: the image with Google Chrome too (reference/chrome.Dockerfile),
// for a probe of a page Playwright's Chromium cannot play; the tests never run in it
if (process.env.CIDERUI_CHROME === '1') {
  // FROM takes a name, not an ID: a tag of this image's own
  const base = `${IMAGE}:${image.slice('sha256:'.length, 'sha256:'.length + 12)}`;
  execFileSync('docker', ['tag', image, base]);
  // no context: the Dockerfile on stdin is all it needs
  image = execFileSync('docker', ['build', '-q', '--build-arg', `BASE=${base}`, '-'], { input: readFileSync(join(ROOT, 'reference/chrome.Dockerfile')), encoding: 'utf8', stdio: ['pipe', 'pipe', 'inherit'] }).trim();
  if (!/^sha256:[0-9a-f]{64}$/.test(image)) throw new Error(`docker gave no image ID for Chrome's: ${image}`);
}

const STEPS = 'node test/tokens.mjs && node test/type.mjs && node test/page.mjs && node test/components.mjs && node test/surfaces.mjs && node reference/touch/audit.mjs && node test/probes.mjs && node scripts/build-lucide.mjs --check && node scripts/build-docs.mjs --check && node scripts/docs-pictures.mjs --check && node test/docs.mjs';
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
  image, 'sh', '-c', wrap, 'run', ...command,
], { encoding: 'utf8' }).trim();
const q = (s) => `'${s.replaceAll("'", "'\\''")}'`;
// no extended attributes (docker cp refuses macOS's com.apple.provenance,
// "operation not supported", on a file that carries it), no AppleDouble
// ._ files
const TAR_ENV = { ...process.env, COPYFILE_DISABLE: '1' };

let status = 1;
try {
  const exclude = '--exclude ./node_modules --exclude ./.git';
  execSync(`tar -c --no-xattrs -C ${q(ROOT)} ${exclude} . | docker cp - ${id}:/src/ciderui`, { env: TAR_ENV, stdio: ['ignore', 'ignore', 'inherit'] });
  execSync(`tar -c --no-xattrs -C ${q(theme)} --exclude ./.git . | docker cp - ${id}:/src/ciderui-zola`, { env: TAR_ENV, stdio: ['ignore', 'ignore', 'inherit'] });
  status = spawnSync('docker', ['start', '-a', id], { stdio: 'inherit' }).status ?? 1;
  const out = mkdtempSync(join(tmpdir(), 'ciderui-written-'));
  try {
    execSync(`docker cp ${id}:/tmp/written.tar - | tar -xOf - written.tar | tar -xf - -C ${q(out)}`, { stdio: ['ignore', 'ignore', 'inherit'] });
    const written = execSync('find . -type f', { cwd: out, encoding: 'utf8' }).split('\n').filter(Boolean).map((f) => f.replace(/^\.\//, ''));
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
