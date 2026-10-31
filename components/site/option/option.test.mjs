// The option against apple.com's buy page (buy:option*): the tile at rest,
// hovered, focused from the keyboard (on its radio) and checked (a checked
// copy, its box apple.com's at its least), the radio over it, the title,
// its detail and the price. Golden, with their why: the letter-spacing, the
// title's and the price's places, its own ground, no margin above, light in
// a dark block.

import { examples, PADDING, SIZE } from '../../../test/harness.mjs';

const ex = examples(import.meta.url);
const EDGE = ['border-top-width', 'border-top-style', 'border-top-color', 'border-bottom-width', 'border-bottom-color', 'border-top-left-radius'];
const INPUT = '.cui-option-input';
const LABEL = '.cui-option-label';
const TRACKING = 'the system font\'s fitted letter-spacing (tokens.css), not apple.com\'s';

export const cases = [
  {
    name: 'option',
    html: ex('option'),
    checks: [
      { target: LABEL, props: [...SIZE, 'color', ...PADDING, ...EDGE, 'background-color', 'min-height', 'transition-property', 'transition-duration'], apple: ['buy', 'option'] },
      { target: LABEL, state: 'hover', props: ['background-color', 'border-top-color', 'color'], apple: ['buy', 'option', 'hover'] },
      { target: LABEL, on: INPUT, state: 'focus-visible', props: ['outline-style', 'outline-width', 'outline-color', 'outline-offset'], apple: ['buy', 'option.focus', 'focus-visible'] },
      { target: INPUT, props: ['position', 'top', 'left', 'opacity', 'margin-top', 'margin-left', 'display'], apple: ['buy', 'option.input'] },
      { target: '.cui-option-title', props: SIZE, apple: ['buy', 'option.title'] },
      { target: '.cui-option-price', props: [...SIZE, 'display', 'margin-top', 'text-align'], apple: ['buy', 'option.price'] },
      { target: '.cui-option-title', props: ['letter-spacing', 'box-x'], golden: 'at the tile\'s start, 16px in (its edge and padding), the system font\'s letter-spacing' },
      { target: '.cui-option-price', props: ['letter-spacing', 'box-width'], golden: '40% of the row at its end, as apple.com\'s price column (buy:option.right)' },
      { target: '.cui-option', props: ['background-color', 'border-top-left-radius', 'margin-top'], golden: 'its own light ground, the label\'s corners, as a field\'s; no margin above (apple.com\'s 14px, the form\'s spacing, buy:option.box)' },
      { target: LABEL, props: ['letter-spacing'], golden: TRACKING },
    ],
  },
  {
    name: 'checked',
    html: ex('checked'),
    checks: [
      { target: LABEL, props: [...SIZE, 'color', ...PADDING, ...EDGE, 'background-color', 'min-height', 'box-height'], apple: ['buy', 'option.checked'] },
      { target: '.cui-option-detail', props: [...SIZE, 'display', 'margin-top'], apple: ['buy', 'option.detail'] },
      { target: '.cui-option-detail', props: ['letter-spacing'], golden: TRACKING },
    ],
  },
  {
    name: 'dark',
    html: `<div class="cui-dark">${ex('option')}</div>`,
    checks: [
      { target: LABEL, props: ['color', 'background-color', 'border-top-color'], golden: 'light in a dark block, as a field (apple.com shows options on light pages alone)' },
    ],
  },
];
