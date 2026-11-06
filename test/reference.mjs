// @invariant: A probe, width or property the record lacks throws: a check never compares
// with a value nobody measured.

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
  // @external: the probe recorded the box at rest only
  if (['box-height', 'box-width', 'box-top', 'box-left', 'box-right'].includes(prop)) {
    const v = prop === 'box-right' && p.box?.left !== undefined ? p.box.left + p.box.width : p.box?.[prop.slice(4)];
    if (state || v === undefined) throw new Error(`reference ${name}:${probe} has no ${state ? state + ' ' : ''}${prop} at ${width}px`);
    return `${v}px`;
  }
  const styles = state ? p.states?.[state] : p.styles;
  if (!styles || !(prop in styles)) throw new Error(`reference ${name}:${probe}${state ? ':' + state : ''} has no ${prop} at ${width}px`);
  return styles[prop];
}

// @trap: not apple.com's fourth, 1600: no part's rule starts above 1200, so only what
// the window sizes, up to its max-width, reads otherwise there, in the cases
// that say WITH_1600
export const WIDTHS = [1200, 900, 390];

// @external: App Store Connect's pages were read at one width, 1920px
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

// @external: a toolkit token as App Store Connect's page computed it
export function tb(name) {
  const tokens = ascPage('tokens').tokens;
  const v = tokens[`--tb-${name}`];
  if (v === undefined) throw new Error(`reference asc tokens has no --tb-${name}`);
  return v;
}
