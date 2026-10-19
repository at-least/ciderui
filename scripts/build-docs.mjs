// Build what the documentation (docs/) generates:
//
// - docs/examples/<name>/<id>.html: a copy of every component's example
//   (components/<name>/examples/<id>.html), for the pages' cui_example
//   component (docs/templates/cui-example.html), since zola's load_data
//   reads only inside the site. Committed, like the stylesheet, so the site
//   builds with zola alone. docs/examples/pages/ is not generated: the
//   examples a page writes itself, kept there so the same component shows
//   them.
// - docs/static/example.css, the example frames' stylesheet: Tailwind's
//   preflight, ciderui with its fonts and every component, and the
//   utilities the examples' class attributes name (the type's), compiled
//   as ciderui's checks compile it (test/compile.mjs), minified. Its fonts'
//   url()s are the theme's static/fonts/, beside it in the built site.
//
//   node scripts/build-docs.mjs           write them
//   node scripts/build-docs.mjs --check   fail if the committed ones differ

import { existsSync, mkdirSync, mkdtempSync, readFileSync, readdirSync, rmSync, writeFileSync } from 'node:fs';
import { tmpdir } from 'node:os';
import { join, relative } from 'node:path';
import { ROOT, compile } from '../test/compile.mjs';

const DOCS = join(ROOT, 'docs');

// every component's examples, by their path under docs/examples
function exampleFiles() {
  const out = new Map();
  for (const name of readdirSync(join(ROOT, 'components')).sort()) {
    const dir = join(ROOT, 'components', name, 'examples');
    for (const f of readdirSync(dir).filter((f) => f.endsWith('.html')).sort()) out.set(`${name}/${f}`, readFileSync(join(dir, f), 'utf8'));
  }
  return out;
}

// the class names every example's markup writes, the pages' own included
function classes(examplesDir) {
  const names = new Set();
  for (const e of readdirSync(examplesDir, { recursive: true, withFileTypes: true })) {
    if (!e.isFile() || !e.name.endsWith('.html')) continue;
    const html = readFileSync(join(e.parentPath, e.name), 'utf8');
    for (const [, list] of html.matchAll(/\bclass="([^"]*)"/g)) for (const c of list.split(/\s+/).filter(Boolean)) names.add(c);
  }
  return [...names].sort().join(' ');
}

// writes everything under out: examples/ and static/example.css
function build(out) {
  const examplesDir = join(out, 'examples');
  if (existsSync(examplesDir)) {
    for (const e of readdirSync(examplesDir)) if (e !== 'pages') rmSync(join(examplesDir, e), { recursive: true, force: true });
  }
  mkdirSync(join(examplesDir, 'pages'), { recursive: true });
  // the pages' own examples name classes too: when building elsewhere
  // (--check), they are read from docs/
  if (out !== DOCS) {
    for (const f of readdirSync(join(DOCS, 'examples/pages'))) writeFileSync(join(examplesDir, 'pages', f), readFileSync(join(DOCS, 'examples/pages', f)));
  }
  for (const [path, html] of exampleFiles()) {
    mkdirSync(join(examplesDir, path, '..'), { recursive: true });
    writeFileSync(join(examplesDir, path), html);
  }
  const components = readdirSync(join(ROOT, 'components')).sort().map((c) => `@import "components/${c}/${c}.css";`).join('\n');
  mkdirSync(join(out, 'static'), { recursive: true });
  writeFileSync(join(out, 'static/example.css'), compile({ css: `@import "fonts.css";\n@import "index.css";\n${components}`, classes: classes(examplesDir) }));
}

// every generated file under a root, relative
function generated(root) {
  const out = [];
  for (const e of readdirSync(join(root, 'examples'), { recursive: true, withFileTypes: true })) {
    if (e.isFile()) out.push(relative(root, join(e.parentPath, e.name)));
  }
  return [...out.filter((f) => !f.startsWith('examples/pages/')), 'static/example.css'].sort();
}

if (process.argv[2] === '--check') {
  const dir = mkdtempSync(join(tmpdir(), 'ciderui-docs-'));
  let stale;
  try {
    build(dir);
    const want = generated(dir);
    const have = generated(DOCS).filter((f) => existsSync(join(DOCS, f)));
    stale = [
      ...want.filter((f) => !have.includes(f)).map((f) => `missing docs/${f}`),
      ...have.filter((f) => !want.includes(f)).map((f) => `extra docs/${f}`),
      ...want.filter((f) => have.includes(f) && !readFileSync(join(dir, f)).equals(readFileSync(join(DOCS, f)))).map((f) => `stale docs/${f}`),
    ];
  } finally {
    rmSync(dir, { recursive: true, force: true });
  }
  for (const s of stale) console.log(s);
  console.log(stale.length ? 'docs/ is stale: node scripts/build-docs.mjs' : 'docs/examples and docs/static/example.css are up to date');
  process.exit(stale.length ? 1 : 0);
} else {
  build(DOCS);
}
