// The console's radio against App Store Connect's toolkit's radio and
// label (--tb-component-Radio-*: brand, md; --tb-component-Label-*): its
// size, round corners, edge's width, the chosen mark's size and color and
// shadow, the disabled edge, the label's type. Golden, with their why: the
// edge (fitted), the flat blue chosen, the gap, the caption, the ring, the
// dark side.

import { examples } from '../../../test/harness.mjs';

const ex = examples(import.meta.url);
const R = 'component-Radio-variant';
const ON = 'label:nth-of-type(1)';
const OFF = 'label:nth-of-type(2)';
const DISABLED = 'label:nth-of-type(3)';

export const cases = [
  {
    name: 'radios',
    html: ex('radios'),
    widths: [1200, 390],
    checks: [
      { target: `${OFF} .cui-console-radio-input`, props: ['box-width', 'box-height', 'border-top-width', 'border-top-style'], tb: [`${R}-controlSize-md-size`, `${R}-controlSize-md-size`, `${R}-controlSize-md-state-off-border-width`, `${R}-controlSize-md-state-off-border-style`] },
      { target: `${OFF} .cui-console-radio-input`, props: ['border-top-left-radius'], golden: 'round: the toolkit\'s calc(15px / 2)' },
      { target: OFF, props: ['font-size', 'line-height', 'font-weight', 'color'], tb: ['component-Label-variant-controlSize-md-inline-on-textStyle-fontSize', 'component-Label-variant-controlSize-md-inline-on-textStyle-lineHeight', 'component-Label-variant-controlSize-md-inline-on-textStyle-fontWeight', 'component-Label-color'] },
      { target: OFF, props: ['column-gap'], golden: 'its label 9px after it, the checkbox\'s gap (the toolkit\'s label names none for a radio)' },
      { target: `${OFF} .cui-console-radio-input`, props: ['border-top-color', 'background-color'], golden: 'its edge the field\'s, coolGray-800 at 0.58 (3:1; the toolkit\'s 0.18 is 1.4:1: tokens.css, decision 21), on white' },
      { target: `${ON} .cui-console-radio-input`, props: ['box-shadow'], tb: ['component-Radio-variant-colorRole-brand-effect-default-state-on-boxShadow'] },
      { target: `${ON} .cui-console-radio-input`, props: ['background-color', 'border-top-color', 'background-image'], golden: 'chosen, the brand\'s flat blue, its edge gone, a white mark 6.43px across (the toolkit\'s md mark, white at 0.95): the toolkit\'s blue is a gradient' },
      { target: `${DISABLED} .cui-console-radio-input`, props: ['border-top-color'], tb: [`${R}-controlSize-md-effect-disabled-border-color`] },
      { target: `${DISABLED} .cui-console-radio-input`, props: ['background-color'], golden: 'disabled, the gray fill' },
      { target: DISABLED, props: ['color'], golden: 'a disabled radio\'s label the disabled label' },
      { target: `${ON} .cui-console-radio-caption`, props: ['color', 'font-size', 'line-height', 'grid-column-start'], golden: 'a line under the label, the secondary label at 0.74 (fitted), 13px on 18px' },
    ],
  },
  {
    name: 'radios focus',
    html: ex('radios'),
    widths: [1200],
    press: ['Tab'],
    checks: [
      { target: `${ON} .cui-console-radio-input`, props: ['outline-style', 'outline-width', 'outline-color'], golden: 'focused from the keyboard, base.css\'s ring (rule 9; the toolkit\'s is 3px of the blue at 0.15)' },
    ],
  },
  {
    name: 'radios touch',
    html: ex('radios'),
    touch: true,
    widths: [390],
    checks: [
      { target: DISABLED, props: ['box-height'], golden: 'a radio with no caption 44px tall on a touch screen, the HIG\'s 44pt (Accessibility, Mobility); 18px with a pointer' },
      { target: `${DISABLED} .cui-console-radio-input`, props: ['box-width', 'box-height'], golden: 'the radio drawn as before, 15px' },
    ],
  },
  {
    name: 'radios forced colors',
    html: ex('radios'),
    widths: [1200],
    forcedColors: true,
    checks: [
      { target: `${ON} .cui-console-radio-input`, props: ['border-top-color', 'background-color', 'background-image', 'box-shadow'], golden: 'chosen, its edge and mark the theme\'s Highlight: forced colors drop the gradient and the fill, and the chosen radio read as the others' },
      { target: `${OFF} .cui-console-radio-input`, props: ['border-top-color', 'background-image'], golden: 'not chosen, the forced edge and no mark' },
    ],
  },
  {
    name: 'radios dark',
    html: `<div class="cui-dark">${ex('radios')}</div>`,
    widths: [1200],
    checks: [
      { target: `${OFF} .cui-console-radio-input`, props: ['border-top-color', 'background-color'], golden: 'the dark side derived; its edge fitted to 3:1 on black (tokens.css, decisions 18, 21)' },
      { target: `${ON} .cui-console-radio-input`, props: ['background-color'], golden: 'the dark side derived and fitted: white on it 4.5:1 (tokens.css)' },
    ],
  },
];
