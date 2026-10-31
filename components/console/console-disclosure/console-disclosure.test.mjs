// The console's disclosure against the HIG's disclosure controls (a
// triangle at the label's leading edge, inward closed and down open; a
// disclosure button down closed and up open) and App Store Connect's menu
// group's chevron (components/console/console-menu). Golden, with their why.

import { examples } from '../../../test/harness.mjs';

const ex = examples(import.meta.url);
const D = '.cui-console-disclosure';
const CLOSED = `${D}:not([open])`;
const OPEN = `${D}[open]`;

export const cases = [
  {
    name: 'disclosure',
    html: ex('disclosure'),
    widths: [1200],
    checks: [
      { target: `${CLOSED} > summary`, props: ['list-style-type', 'display', 'column-gap', 'color', 'cursor'], golden: 'no marker of the browser\'s; its chevron 4px before its label; the secondary label (the menu group\'s chevron\'s)' },
      { target: `${CLOSED} > summary`, props: ['color'], state: 'hover', golden: 'hovered, the label' },
      { target: `${CLOSED} .cui-console-disclosure-chevron`, props: ['rotate', 'box-width', 'box-height', 'box-left', 'color'], golden: 'closed, pointing inward at the label\'s leading edge (the HIG), 1em, Lucide\'s chevron-right and its secondary label as the menu group\'s' },
      { target: `${CLOSED} .cui-console-disclosure-chevron`, props: ['color'], state: 'hover', on: `${CLOSED} > summary`, golden: 'the label hovered, the chevron still the secondary label, as the menu group\'s: its color its own, not the label\'s' },
      { target: `${OPEN} .cui-console-disclosure-chevron`, props: ['rotate'], golden: 'open, pointing down (the HIG): 90deg' },
    ],
  },
  {
    name: 'disclosure touch',
    html: ex('disclosure'),
    touch: true,
    widths: [390],
    checks: [
      { target: `${CLOSED} > summary`, props: ['box-height', 'min-width'], golden: 'on a touch screen the line itself 44px tall and 44px wide at least, the HIG\'s 44pt: a target drawn round it overlapped the next summary\'s and the opened content' },
      { target: `${CLOSED} > summary`, props: ['color'], state: 'hover', rest: 'a tapped disclosure keeps no hover (rule 5)' },
    ],
  },
  {
    name: 'disclosure button',
    html: ex('button'),
    widths: [1200],
    checks: [
      { target: `${CLOSED} > summary`, props: ['list-style-type'], golden: 'a console button as the summary: no marker of the browser\'s' },
      { target: `${CLOSED} .cui-console-disclosure-chevron`, props: ['rotate'], golden: 'closed, pointing down (the HIG\'s disclosure button)' },
      { target: `${OPEN} .cui-console-disclosure-chevron`, props: ['rotate'], golden: 'open, pointing up (the HIG\'s disclosure button): 180deg' },
    ],
  },
  {
    name: 'disclosure dark',
    html: `<div class="cui-dark">${ex('disclosure')}</div>`,
    widths: [1200],
    checks: [
      { target: `${CLOSED} > summary`, props: ['color'], golden: 'the dark side derived (tokens.css, decision 18)' },
    ],
  },
];
