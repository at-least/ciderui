// The console's chip against App Store Connect's toolkit's chip
// (--tb-component-Chip-*: md, gray, neutral): its room, corners, gap, type,
// fill and label, a link's hovered fill. Golden, with their why: the key's
// fitted label, the remove control, the dark side.

import { examples } from '../../test/harness.mjs';

const ex = examples(import.meta.url);
const C = 'component-Chip-variant';
const FILTER = '.cui-console-chip:first-child';
const TAG = 'a.cui-console-chip';

export const cases = [
  {
    name: 'chips',
    html: ex('chips'),
    widths: [1200, 390],
    checks: [
      { target: FILTER, props: ['padding-top', 'padding-bottom', 'padding-left', 'padding-right', 'border-top-left-radius', 'column-gap', 'font-size', 'line-height', 'font-weight'], tb: [`${C}-controlSize-md-paddingBlockStart`, `${C}-controlSize-md-paddingBlockEnd`, `${C}-controlSize-md-paddingInline`, `${C}-controlSize-md-paddingInline`, `${C}-controlSize-md-borderRadius`, `${C}-controlSize-md-gap`, `${C}-controlSize-md-textStyle-fontSize`, `${C}-controlSize-md-textStyle-lineHeight`, `${C}-controlSize-md-textStyle-fontWeight`] },
      { target: FILTER, props: ['background-color', 'color'], tb: [`${C}-colorRole-neutral-effect-default-variant-gray-backgroundColor`, `${C}-colorRole-neutral-effect-default-variant-gray-color`] },
      { target: TAG, props: ['background-color', 'color'], state: 'hover', tb: [`${C}-colorRole-neutral-effect-hover-variant-gray-backgroundColor`, `${C}-colorRole-neutral-effect-default-variant-gray-color`] },
      { target: TAG, props: ['text-decoration-line'], golden: 'a tag that is a link: no underline, its fill says it' },
      { target: '.cui-console-chip-key', props: ['color'], golden: 'the key in the secondary label at 0.74, the toolkit\'s art color fitted (tokens.css, decision 21): its 0.60 is 3.4:1 on white' },
      { target: '.cui-console-chip-remove', props: ['color'], golden: 'the remove control in the secondary label' },
      { target: '.cui-console-chip-remove', props: ['color'], state: 'hover', golden: 'hovered, the label\'s own' },
      { target: '.cui-console-chip-remove > svg', props: ['box-width', 'box-height'], golden: 'its ×, Lucide\'s x, 12px: the toolkit\'s art button draws an SF Symbol' },
    ],
  },
  {
    name: 'chips touch',
    html: ex('chips'),
    touch: true,
    widths: [390],
    checks: [
      { target: TAG, props: ['background-color'], state: 'hover', rest: 'a tapped chip keeps no hovered fill (rule 5)' },
    ],
  },
  {
    name: 'chips dark',
    html: `<div class="cui-dark">${ex('chips')}</div>`,
    widths: [1200],
    checks: [
      { target: FILTER, props: ['background-color', 'color'], golden: 'the dark side derived (tokens.css, decision 18)' },
      { target: '.cui-console-chip-key', props: ['color'], golden: 'the dark side derived (tokens.css, decision 18)' },
    ],
  },
];
