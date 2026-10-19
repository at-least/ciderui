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
  if (['box-height', 'box-width', 'box-top', 'box-left'].includes(prop)) {
    const v = p.box?.[prop.slice(4)];
    if (state || v === undefined) throw new Error(`reference ${name}:${probe} has no ${state ? state + ' ' : ''}${prop} at ${width}px`);
    return `${v}px`;
  }
  const styles = state ? p.states?.[state] : p.styles;
  if (!styles || !(prop in styles)) throw new Error(`reference ${name}:${probe}${state ? ':' + state : ''} has no ${prop} at ${width}px`);
  return styles[prop];
}

export const WIDTHS = [1600, 1200, 900, 390];
