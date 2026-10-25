// Runs a command in test/Dockerfile's image, the only place the browser
// starts (test/chromium.mjs): the repository mounted at /src/ciderui as the
// caller's user (golden.json and the pictures it writes are theirs), the
// image's node_modules over the checkout's, ciderui-zola (the docs' theme,
// a sibling checkout) beside it, CASES and UPDATE_GOLDEN passed through.
//
//   node scripts/docker.mjs                          the browser steps of npm test
//   node scripts/docker.mjs node test/components.mjs a command
//   UPDATE_GOLDEN=1 node scripts/docker.mjs node test/components.mjs

import { execFileSync } from 'node:child_process';
import { realpathSync } from 'node:fs';
import { join } from 'node:path';
import { userInfo } from 'node:os';
import { ROOT } from '../test/compile.mjs';

const IMAGE = 'ciderui-test';
execFileSync('docker', ['build', '-q', '-t', IMAGE, '-f', join(ROOT, 'test/Dockerfile'), ROOT], { stdio: ['ignore', 'ignore', 'inherit'] });

const STEPS = 'node test/tokens.mjs && node test/type.mjs && node test/page.mjs && node test/components.mjs && node scripts/build-lucide.mjs --check && node scripts/build-docs.mjs --check && node scripts/docs-pictures.mjs --check && node test/docs.mjs';
const command = process.argv.length > 2 ? process.argv.slice(2) : ['sh', '-c', STEPS];
const theme = realpathSync(join(ROOT, 'docs/themes/ciderui-zola'));
const { uid, gid } = userInfo();
const pass = ['CASES', 'UPDATE_GOLDEN'].filter((k) => process.env[k] !== undefined).flatMap((k) => ['-e', `${k}=${process.env[k]}`]);
execFileSync('docker', [
  'run', '--rm', '--init', '--ipc=host',
  '--user', `${uid}:${gid}`, '-e', 'HOME=/tmp', ...pass,
  '-v', `${ROOT}:/src/ciderui`, '-v', '/src/ciderui/node_modules',
  '-v', `${theme}:/src/ciderui-zola:ro`,
  IMAGE, ...command,
], { stdio: 'inherit' });
