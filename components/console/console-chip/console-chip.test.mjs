import { examples } from '../../../test/harness.mjs';

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
      { target: TAG, props: ['background-color'], state: 'active', tb: [`${C}-colorRole-neutral-effect-press-variant-gray-backgroundColor`] },
      { target: TAG, props: ['text-decoration-line'], golden: 'a tag that is a link: no underline, its fill says it' },
      { target: '.cui-console-chip-key', props: ['color'], golden: 'the key in the secondary label at 0.76, the toolkit\'s art color fitted (tokens.css, decision 21): its 0.60 is 3.4:1 on white' },
      { target: '.cui-console-chip-remove', props: ['color'], golden: 'the remove control in the secondary label' },
      { target: '.cui-console-chip-remove', props: ['color'], state: 'hover', golden: 'hovered, the label\'s own' },
      { target: '.cui-console-chip-remove > svg', props: ['box-width', 'box-height'], golden: 'its ×, Lucide\'s x, 12px: the toolkit\'s art button draws an SF Symbol' },
    ],
  },
  {
    name: 'chip long value',
    html: '<p class="w-40"><span class="cui-console-chip"><span class="cui-console-chip-key">id</span>60674cfe203893d8499985f493f86663</span></p>',
    widths: [390],
    checks: [
      { target: '.cui-console-chip', props: ['box-width', 'box-height', 'max-width', 'overflow-wrap'], golden: 'a long value wraps inside its place, never past it: on one line it ran off a phone\'s page' },
    ],
  },
  {
    name: 'chips touch',
    html: ex('chips'),
    touch: true,
    widths: [390],
    checks: [
      { target: TAG, props: ['background-color'], state: 'hover', rest: 'a tapped chip keeps no hovered fill (rule 5)' },
      { target: TAG, props: ['box-height'], golden: 'a chip that is a link 44px tall on a touch screen, the HIG\'s 44pt (Accessibility, Mobility): chips in rows 8px apart are 30px from row to row, and a target drawn round each covered the next' },
      { target: FILTER, props: ['box-height'], golden: 'a chip with a remove control 44px tall on a touch screen, for its control' },
      { target: '.cui-console-chip-remove', props: ['box-width', 'box-height'], golden: 'the remove control 44px square on a touch screen, its chip\'s end, the HIG\'s 44pt (Accessibility, Mobility)' },
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
