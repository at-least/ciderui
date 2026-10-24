// apple.com's measured values (reference/apple.com/results/<page>.json), as
// the checks whose source is apple.com read them: ref('home', 'button', 1200,
// 'background-color', 'hover') is the home page's button's background color,
// hovered, at 1200px. A probe, a width or a property the record lacks throws:
// a check never compares with a value nobody measured.

import { readFileSync } from 'node:fs';
import { join } from 'node:path';
import { ROOT } from './compile.mjs';

const pages = new Map();
const page = (name) => {
  if (!pages.has(name)) pages.set(name, JSON.parse(readFileSync(join(ROOT, 'reference/apple.com/results', `${name}.json`), 'utf8')));
  return pages.get(name);
};

export function ref(name, probe, width, prop, state) {
  const at = page(name).widths[width];
  if (!at) throw new Error(`reference ${name} has no width ${width}`);
  const p = at[probe];
  if (!p) throw new Error(`reference ${name}:${probe} has nothing at ${width}px`);
  // the box, as the probe recorded it (at rest), and its place in the window
  // where the probe recorded that (`position`)
  if (['box-height', 'box-width', 'box-top', 'box-left', 'box-right'].includes(prop)) {
    const v = prop === 'box-right' && p.box?.left !== undefined ? p.box.left + p.box.width : p.box?.[prop.slice(4)];
    if (state || v === undefined) throw new Error(`reference ${name}:${probe} has no ${state ? state + ' ' : ''}${prop} at ${width}px`);
    return `${v}px`;
  }
  const styles = state ? p.states?.[state] : p.styles;
  if (!styles || !(prop in styles)) throw new Error(`reference ${name}:${probe}${state ? ':' + state : ''} has no ${prop} at ${width}px`);
  return styles[prop];
}

export const WIDTHS = [1600, 1200, 900, 390];

// App Store Connect's (reference/asc/results/<page>.json, read at one width,
// 1920px): asc('analytics', 'menu.current', 'height') is that probe's height
// there, `box-height` and `box-width` its box. A probe or a property the
// record lacks throws.
const ascPages = new Map();
const ascPage = (name) => {
  if (!ascPages.has(name)) ascPages.set(name, JSON.parse(readFileSync(join(ROOT, 'reference/asc/results', `${name}.json`), 'utf8')));
  return ascPages.get(name);
};
export function asc(name, probe, prop) {
  const p = ascPage(name).probes[probe];
  if (!p) throw new Error(`reference asc ${name} has no probe ${probe}`);
  if (prop === 'box-height' || prop === 'box-width') {
    if (!p.box) throw new Error(`reference asc ${name}:${probe} has no box`);
    return `${p.box[prop.slice(4)]}px`;
  }
  if (!(prop in p.styles)) throw new Error(`reference asc ${name}:${probe} has no ${prop}`);
  return p.styles[prop];
}

// App Store Connect's toolkit's token (reference/asc/results/tokens.json), as
// its page computed it, by its name without --tb-: tb('size-control-md') is
// "1.875rem". A token the record lacks throws.
export function tb(name) {
  const tokens = ascPage('tokens').tokens;
  const v = tokens[`--tb-${name}`];
  if (v === undefined) throw new Error(`reference asc tokens has no --tb-${name}`);
  return v;
}
