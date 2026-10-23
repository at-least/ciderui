// The Rust crate (rust/), every feature on: formatted, clippy-clean with
// warnings as errors, its tests passing with every failure shown (nextest,
// or, where cargo-nextest is not installed, cargo test --no-fail-fast, the
// run saying so), then the documentation's examples, which nextest does not
// run, and its documentation, private items too, building with no warning
// (a broken intra-doc link, an undocumented public item). The target directory is the crate's
// own: the shared one an environment may export could hold another
// checkout's build.

import { execFileSync, spawnSync } from 'node:child_process';
import { join } from 'node:path';
import { ROOT } from './compile.mjs';

const MANIFEST = join(ROOT, 'rust/Cargo.toml');
const env = { ...process.env, CARGO_TARGET_DIR: join(ROOT, 'rust/target'), RUSTDOCFLAGS: '-D warnings' };

const nextest = spawnSync('cargo', ['nextest', '--version'], { env, stdio: 'ignore' }).status === 0;
if (!nextest) console.log('rust: cargo-nextest is not installed (cargo nextest --version failed): the tests run with cargo test --no-fail-fast');
const tests = nextest
  ? ['nextest', 'run', '--manifest-path', MANIFEST, '--all-features', '--no-fail-fast']
  // --lib --bins --tests: not the doc tests, which the next step runs
  : ['test', '--manifest-path', MANIFEST, '--all-features', '--no-fail-fast', '--lib', '--bins', '--tests'];

for (const args of [
  ['fmt', '--manifest-path', MANIFEST, '--check'],
  ['clippy', '--manifest-path', MANIFEST, '--all-targets', '--all-features', '--', '-D', 'warnings'],
  tests,
  ['test', '--manifest-path', MANIFEST, '--all-features', '--doc'],
  // --document-private-items: without it rustdoc skips a private item's
  // intra-doc links, and a broken one passes
  ['doc', '--manifest-path', MANIFEST, '--no-deps', '--all-features', '--document-private-items'],
]) {
  execFileSync('cargo', args, { env, stdio: 'inherit' });
}
console.log(`rust: fmt, clippy, tests (${nextest ? 'nextest' : 'cargo test'}), documentation examples and documentation clean`);
