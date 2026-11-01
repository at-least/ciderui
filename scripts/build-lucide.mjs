// Write go/lucide/lucide.go, the Go module's Lucide icons, from lucide-static
// as node_modules holds it, which must be package.json's exact version: every
// icon in its icon-nodes.json, which holds Lucide's names and not their
// aliases (the files under icons/ carry both), as a var named by the name in
// Go's capitals; and the attributes every icons/<name>.svg's <svg> carries,
// which icon.Lucide writes on each.
//
//   node scripts/build-lucide.mjs           write it
//   node scripts/build-lucide.mjs --check   fail if it differs from what it writes

import { mkdirSync, readFileSync, writeFileSync } from 'node:fs';
import { dirname, join } from 'node:path';
import { ROOT } from '../test/compile.mjs';

const OUT = join(ROOT, 'go/lucide/lucide.go');
const PACKAGE = join(ROOT, 'node_modules/lucide-static');

// an icons/<name>.svg as its <svg>'s attributes, its own (class) and those
// icon.Icon writes itself (width, height: 1em) aside, and its elements as
// icon.Icon writes them
function file(name) {
  const svg = readFileSync(join(PACKAGE, 'icons', `${name}.svg`), 'utf8');
  const [, open, body] = svg.match(/<svg([^>]*)>([^]*)<\/svg>/);
  const attributes = [...open.matchAll(/([a-zA-Z:-]+)="([^"]*)"/g)].filter(([, k]) => !['class', 'xmlns', 'width', 'height'].includes(k));
  return {
    root: attributes.map(([, k, v]) => `${k}="${v}"`).join(' '),
    viewBox: attributes.find(([, k]) => k === 'viewBox')[2],
    paint: attributes.filter(([, k]) => k !== 'viewBox').map(([, k, v]) => `${k}="${v}"`).join(' '),
    inner: body.replace(/\s*\/>/g, '/>').replace(/>\s+</g, '><').trim(),
  };
}

// `shopping-bag` as ShoppingBag
const goName = (name) => name.split('-').map((p) => p[0].toUpperCase() + p.slice(1)).join('');

function source() {
  const { version } = JSON.parse(readFileSync(join(PACKAGE, 'package.json'), 'utf8'));
  const pinned = JSON.parse(readFileSync(join(ROOT, 'package.json'), 'utf8')).devDependencies['lucide-static'];
  if (version !== pinned) throw new Error(`node_modules holds lucide-static ${version}, package.json pins ${pinned}: npm ci`);
  const nodes = JSON.parse(readFileSync(join(PACKAGE, 'icon-nodes.json'), 'utf8'));
  const names = Object.keys(nodes);
  const { root, viewBox, paint } = file(names[0]);
  let out = `// Package lucide is every Lucide icon (https://lucide.dev/icons/; under its
// license, the repository's LICENSE-Lucide), lucide-static ${version}'s, by its
// name in Go's capitals: shopping-bag is [ShoppingBag]. Each is an [icon.Icon]
// of Lucide's own elements in currentColor, 1em square, so it takes the text's
// color and size; a utility on it sizes it otherwise. In a line of text a page
// adds inline: Tailwind's preflight makes an <svg> a block. A part apple.com
// draws an icon in draws ciderui's own (package icon), not Lucide's.
//
// Written by scripts/build-lucide.mjs from lucide-static's icon-nodes.json and
// icons/, not by hand: npm test fails while this differs from what it writes.
package lucide

import "github.com/at-least/ciderui/go/icon"

const (
	viewBox = "${viewBox}"
	paint   = \`${paint}\`
)
`;
  const seen = new Map();
  for (const [name, elements] of Object.entries(nodes)) {
    if (!/^[a-z0-9]+(-[a-z0-9]+)*$/.test(name) || /^[0-9]/.test(name)) throw new Error(`${name}: not a Lucide name a Go identifier can carry`);
    const inner = elements.map(([tag, attributes]) => `<${tag}${Object.entries(attributes).map(([k, v]) => ` ${k}="${v}"`).join('')}/>`).join('');
    const own = file(name);
    if (own.root !== root) throw new Error(`${name}: its <svg> carries ${own.root}, ${names[0]}'s ${root}`);
    if (own.inner !== inner) throw new Error(`${name}: icon-nodes.json's elements are not icons/${name}.svg's:\n${inner}\n${own.inner}`);
    if (inner.includes('`')) throw new Error(`${name}: its markup would end the raw string`);
    const constant = goName(name);
    if (seen.has(constant)) throw new Error(`${name} and ${seen.get(constant)} are both ${constant}`);
    seen.set(constant, name);
    out += `\n// ${constant} is [${name}](https://lucide.dev/icons/${name}).\nvar ${constant} = icon.Lucide(viewBox, paint, \`${inner}\`)\n`;
  }
  return out;
}

const text = source();
if (process.argv[2] === '--check') {
  if (readFileSync(OUT, 'utf8') !== text) {
    console.log('go/lucide/lucide.go is stale: node scripts/build-lucide.mjs');
    process.exit(1);
  }
  console.log('go/lucide/lucide.go is up to date');
} else {
  mkdirSync(dirname(OUT), { recursive: true });
  writeFileSync(OUT, text);
}
