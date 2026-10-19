// The Rust crate (rust/), every feature on: formatted, clippy-clean with
// warnings as errors, its tests passing (nextest, every failure shown, then
// the documentation's examples, which nextest does not run), and its
// documentation building with no warning (a broken intra-doc link, an
// undocumented public item). The target directory is the crate's own: the
// shared one an environment may export could hold another checkout's build.

import { execFileSync } from 'node:child_process';
import { join } from 'node:path';
import { ROOT } from './compile.mjs';

const MANIFEST = join(ROOT, 'rust/Cargo.toml');
const env = { ...process.env, CARGO_TARGET_DIR: join(ROOT, 'rust/target'), RUSTDOCFLAGS: '-D warnings' };

for (const args of [
  ['fmt', '--manifest-path', MANIFEST, '--check'],
  ['clippy', '--manifest-path', MANIFEST, '--all-targets', '--all-features', '--', '-D', 'warnings'],
  ['nextest', 'run', '--manifest-path', MANIFEST, '--all-features', '--no-fail-fast'],
  ['test', '--manifest-path', MANIFEST, '--all-features', '--doc'],
  ['doc', '--manifest-path', MANIFEST, '--no-deps', '--all-features'],
]) {
  execFileSync('cargo', args, { env, stdio: 'inherit' });
}
console.log('rust: fmt, clippy, tests, documentation examples and documentation clean');
