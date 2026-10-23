// Write rust/src/lucide.rs, the crate's Lucide icons, from lucide-static as
// node_modules holds it, which must be package.json's exact version: every
// icon in its icon-nodes.json, which holds Lucide's names and not their
// aliases (the files under icons/ carry both), as a const named by the name
// in capitals; and the attributes every icons/<name>.svg's <svg> carries,
// which Icon::lucide writes on each.
//
//   node scripts/build-lucide.mjs           write it
//   node scripts/build-lucide.mjs --check   fail if rust/src/lucide.rs differs from what it writes

import { readFileSync, writeFileSync } from 'node:fs';
import { join } from 'node:path';
import { ROOT } from '../test/compile.mjs';

const OUT = join(ROOT, 'rust/src/lucide.rs');
const PACKAGE = join(ROOT, 'node_modules/lucide-static');

// rustfmt's max_width: a const longer than this it breaks after `=` if the
// call then fits its next line, else after `(`
const WIDTH = 100;

// an icons/<name>.svg as its <svg>'s attributes, its own (class) and those
// the crate writes itself (width, height: 1em) aside, and its elements as
// the crate writes them
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

function source() {
  const { version } = JSON.parse(readFileSync(join(PACKAGE, 'package.json'), 'utf8'));
  const pinned = JSON.parse(readFileSync(join(ROOT, 'package.json'), 'utf8')).devDependencies['lucide-static'];
  if (version !== pinned) throw new Error(`node_modules holds lucide-static ${version}, package.json pins ${pinned}: npm ci`);
  const nodes = JSON.parse(readFileSync(join(PACKAGE, 'icon-nodes.json'), 'utf8'));
  const names = Object.keys(nodes);
  const { root, viewBox, paint } = file(names[0]);
  let out = `//! Every Lucide icon (<https://lucide.dev/icons/>; under its license, the
//! repository's LICENSE-Lucide), lucide-static ${version}'s, by its name in
//! capitals: \`shopping-bag\` is [\`SHOPPING_BAG\`]. Each is an [\`Icon\`] of
//! Lucide's own elements in currentColor, 1em square, so it takes the text's
//! color and size; a utility on it sizes it otherwise. In a line of text a
//! page adds \`inline\`: Tailwind's preflight makes an \`<svg>\` a block. A part
//! apple.com draws an icon in draws ciderui's own ([\`crate::icon\`]), not
//! Lucide's.
//!
//! \`\`\`
//! use hypertext::prelude::*;
//! use ciderui::lucide::HEART;
//!
//! let html = rsx! { <a href="/saved">(HEART.class("inline"))" Saved"</a> }.render();
//! assert!(html.as_inner().starts_with(r#"<a href="/saved"><svg class="inline" aria-hidden="true" width="1em""#));
//! \`\`\`

// Written by scripts/build-lucide.mjs from lucide-static's icon-nodes.json
// and icons/, not by hand: npm test fails while this differs from what it
// writes.

use crate::icon::Icon;

/// The view box of every Lucide icon's \`<svg>\`.
pub(crate) const VIEW_BOX: &str = "${viewBox}";

/// The attributes, its view box aside, on every Lucide icon's \`<svg>\`.
pub(crate) const PAINT: &str = r#"${paint}"#;

`;
  for (const [name, elements] of Object.entries(nodes)) {
    if (!/^[a-z0-9]+(-[a-z0-9]+)*$/.test(name) || /^[0-9]/.test(name)) throw new Error(`${name}: not a Lucide name a const can carry`);
    const inner = elements.map(([tag, attributes]) => `<${tag}${Object.entries(attributes).map(([k, v]) => ` ${k}="${v}"`).join('')}/>`).join('');
    const own = file(name);
    if (own.root !== root) throw new Error(`${name}: its <svg> carries ${own.root}, ${names[0]}'s ${root}`);
    if (own.inner !== inner) throw new Error(`${name}: icon-nodes.json's elements are not icons/${name}.svg's:\n${inner}\n${own.inner}`);
    if (inner.includes('"#')) throw new Error(`${name}: its markup would end the raw string`);
    const constant = name.toUpperCase().replaceAll('-', '_');
    const call = `Icon::lucide(r#"${inner}"#);`;
    const line = `pub const ${constant}: Icon = ${call}`;
    out += `/// [\`${name}\`](https://lucide.dev/icons/${name})\n`;
    if (line.length <= WIDTH) out += `${line}\n`;
    else if (4 + call.length <= WIDTH) out += `pub const ${constant}: Icon =\n    ${call}\n`;
    else out += `pub const ${constant}: Icon = Icon::lucide(\n    r#"${inner}"#,\n);\n`;
  }
  return out;
}

const text = source();
if (process.argv[2] === '--check') {
  const stale = readFileSync(OUT, 'utf8') !== text;
  console.log(stale ? 'rust/src/lucide.rs is stale: node scripts/build-lucide.mjs' : 'rust/src/lucide.rs is up to date');
  process.exit(stale ? 1 : 0);
} else {
  writeFileSync(OUT, text);
}
