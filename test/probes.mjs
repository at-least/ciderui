// @invariant: Run in a copy of the repository: a probe writes results/<browser>.json, and
// the committed results stay the record of the runs each MEASURED.md names.

import { execFileSync } from 'node:child_process';
import { cpSync, existsSync, mkdtempSync, readFileSync, readdirSync, rmSync, symlinkSync } from 'node:fs';
import { tmpdir } from 'node:os';
import { dirname, join, relative } from 'node:path';
import { ROOT } from './compile.mjs';

const copy = mkdtempSync(join(tmpdir(), 'ciderui-probes-'));
process.on('exit', () => rmSync(copy, { recursive: true, force: true }));
const tree = join(copy, 'ciderui');
const CONSOLE_SITE = join(copy, 'console-site');

const RUN = [['node', 'run.mjs']];
const PROBES = {
  console: { env: { SITE: CONSOLE_SITE }, commands: [['node', 'build.mjs', CONSOLE_SITE], ...RUN] },
  details: { env: {}, commands: [['node', 'page.mjs'], ...RUN] },
  dialog: { env: {}, commands: RUN },
  flyout: { env: {}, commands: RUN },
  forced: { env: {}, commands: [['node', 'page.mjs'], ...RUN] },
  popover: { env: {}, commands: RUN },
  scheme: { env: {}, commands: RUN },
  search: { env: {}, commands: RUN },
  select: { env: {}, commands: RUN },
  // @external: the image has no system font of Apple's: the stack draws Inter
  stack: { env: { EXPECT: 'inter' }, commands: RUN },
  tap: { env: {}, commands: RUN },
};

const NOT_RUN = {
  'reference/anchor/': 'reads apple.com live',
  'reference/apple-account/': 'reads Apple Account live, signed in',
  'reference/apple.com/': 'reads apple.com live (report.mjs prints what it read)',
  'reference/asc/': 'reads App Store Connect live, signed in',
  'reference/developer-account/': 'reads the developer account live, signed in',
  'reference/developer-forums/': 'reads the developer forums live',
  'reference/feedback-assistant/': 'reads Feedback Assistant live, signed in',
  'reference/icons/': 'reads apple.com live',
  'reference/tracking/': 'reads apple.com live (derive.mjs and tables.mjs read what it wrote)',
  'reference/vm.mjs': 'Safari in a tester VM',
  'reference/xcuitest.mjs': 'the iOS simulator',
  'reference/touch/audit.mjs': 'a step of npm test of its own',
};
// @invariant: read by the probes above, not run on their own
const READ = ['reference/browsers.mjs', ...Object.keys(PROBES).map((d) => `reference/${d}/`)];

const scripts = readdirSync(join(ROOT, 'reference'), { recursive: true })
  .map((f) => `reference/${f}`)
  .filter((f) => /\.(mjs|js|sh)$/.test(f) && !f.includes('/results/'));
const unlisted = scripts.filter((f) => ![...Object.keys(NOT_RUN), ...READ].some((p) => f === p || (p.endsWith('/') && f.startsWith(p))));
if (unlisted.length) throw new Error(`not run here and not in NOT_RUN: ${unlisted.join(', ')}`);
const stale = Object.keys(NOT_RUN).filter((p) => !scripts.some((f) => f === p || f.startsWith(p)));
if (stale.length) throw new Error(`in NOT_RUN, no such script: ${stale.join(', ')}`);
// @invariant: a script not run here still has its relative paths (an import, a page, an
// example) checked: up to its first ${…}, the directory it names
const missing = scripts.flatMap((f) => [...readFileSync(join(ROOT, f), 'utf8').matchAll(/['`](\.{1,2}\/[^'`\n]*)['`]/g)]
  .map(([, path]) => (path.includes('${') ? path.slice(0, path.lastIndexOf('/', path.indexOf('${')) + 1) : path))
  .filter((path) => !path.includes('results/') && !existsSync(join(ROOT, dirname(f), path)))
  .map((path) => `${f}: ${path}`));
if (missing.length) throw new Error(`named by a script, no such file: ${missing.join(', ')}`);

cpSync(ROOT, tree, { recursive: true, filter: (f) => !['node_modules', '.git'].includes(relative(ROOT, f)) });
symlinkSync(join(ROOT, 'node_modules'), join(tree, 'node_modules'));

const failed = [];
for (const [dir, { env, commands }] of Object.entries(PROBES)) {
  const start = Date.now();
  try {
    for (const [file, ...args] of commands) {
      execFileSync(file, args, { cwd: join(tree, 'reference', dir), env: { ...process.env, BROWSER: 'chromium', ...env }, stdio: ['ignore', 'pipe', 'pipe'] });
    }
    console.log(`ok   ${dir} (${((Date.now() - start) / 1000).toFixed(1)}s)`);
  } catch (e) {
    failed.push(dir);
    console.log(`FAIL ${dir} (${((Date.now() - start) / 1000).toFixed(1)}s)\n${e.stdout ?? ''}${e.stderr ?? ''}${e.stdout === undefined ? e.message : ''}`);
  }
}
console.log(`probes: ${Object.keys(PROBES).length - failed.length} passed, ${failed.length} failed; ${Object.keys(NOT_RUN).length} not run here`);
process.exit(failed.length ? 1 : 0);
