// @invariant: The gate inside the test image: the browser steps only once the checks
// that open no browser pass (the owner, 2026-11-14: a failure a fast check finds wastes
// no browser run). test/go.mjs runs those on the host before the image starts; the one
// held here too is the one whose result the host's node_modules can change: the image's
// Tailwind is package-lock.json's by its build, the host's whatever was last installed,
// and the browser steps read the dist that check holds. Then every browser step at
// once, each one's output held: a failed step's is printed whole in this order, so two
// steps' lines never interleave, and a passed step's is dropped. The last line names
// every failed step: it is the one a reader of the tail sees. Run one after another
// they took 359s on lab (surfaces 116s, probes 69s, components 64s); at once, about
// the slowest.
import { spawn, spawnSync } from 'node:child_process';

const STATIC = [
  'node scripts/build-css.mjs --check',
];
const BROWSER = [
  'node test/type.mjs',
  'node test/page.mjs',
  'node test/components.mjs',
  'node test/rules.mjs',
  'node test/pictures.mjs',
  'node test/surfaces.mjs',
  'node reference/touch/audit.mjs',
  'node test/probes.mjs',
  'node test/assets.mjs',
  'node test/settings.mjs',
];

for (const step of STATIC) {
  const { status } = spawnSync('sh', ['-c', step], { stdio: 'inherit' });
  if (status !== 0) {
    console.log(`gate: ${step} failed (exit ${status}); no browser step run`);
    process.exit(1);
  }
}

const started = Date.now();
const runs = BROWSER.map((step) => new Promise((resolve) => {
  const child = spawn('sh', ['-c', step], { stdio: ['ignore', 'pipe', 'pipe'] });
  const out = [];
  child.stdout.on('data', (d) => out.push(d));
  child.stderr.on('data', (d) => out.push(d));
  const t = Date.now();
  child.on('close', (code) => resolve({ step, code, out: Buffer.concat(out), seconds: Math.round((Date.now() - t) / 1000) }));
}));
const failed = [];
for (const run of runs) {
  const { step, code, out, seconds } = await run;
  if (code === 0) continue;
  failed.push(`${step} failed`);
  process.stdout.write(out);
  console.log(`gate: ${step} failed (exit ${code}) in ${seconds}s`);
}
const seconds = Math.round((Date.now() - started) / 1000);
if (failed.length) console.log(`gate: ${failed.join(', ')}; ${BROWSER.length - failed.length} of ${BROWSER.length} browser steps passed in ${seconds}s`);
else console.log(`gate: ${BROWSER.length} browser steps passed in ${seconds}s`);
process.exit(failed.length ? 1 : 0);
