// The field against apple.com's own (retail:field*, the retail page's
// form-textbox): the input at rest, hovered and focused (alone, and with
// :focus-visible, as a text field always is), its label as the placeholder,
// floated when the field is focused and when it is filled; in a dark block,
// as apple.com's in a dark color scheme (retail-dark:field*). Golden, with
// their why: the letter-spacing, the field's own ground, no margin below.

import { examples } from '../../test/harness.mjs';

const ex = examples(import.meta.url);
const TYPE = ['font-size', 'line-height', 'font-weight'];
const EDGE = ['border-top-width', 'border-top-style', 'border-top-color', 'border-bottom-width', 'border-bottom-color', 'border-top-left-radius'];
const PAD = ['padding-top', 'padding-right', 'padding-bottom', 'padding-left'];
const INPUT = '.cui-field-input';
const LABEL = '.cui-field-label';
const TRACKING = 'the system font\'s fitted letter-spacing (tokens.css), not apple.com\'s';
const dark = (html) => `<div class="cui-dark">${html}</div>`;

export const cases = [
  {
    name: 'field',
    html: ex('field'),
    checks: [
      { target: INPUT, props: [...TYPE, 'color', ...PAD, ...EDGE, 'background-color', 'height', 'box-sizing', 'box-height', 'transition-property', 'transition-duration', 'outline-style'], apple: ['retail', 'field.input'] },
      { target: INPUT, state: 'hover', props: ['color', 'background-color', 'border-top-color'], apple: ['retail', 'field.input', 'hover'] },
      { target: INPUT, state: 'focus', props: [...PAD, ...EDGE], apple: ['retail', 'field.input', 'focus'] },
      // focused, its box apple.com's at rest: the 2px edge takes the padding's pixel
      { target: INPUT, state: 'focus', props: ['box-height'], apple: ['retail', 'field.input'] },
      { target: INPUT, state: ['focus', 'focus-visible'], props: ['padding-left', 'border-top-width', 'border-top-color', 'outline-style'], apple: ['retail', 'field.input', 'focus+focus-visible'] },
      { target: LABEL, props: [...TYPE, 'color', 'position', 'top', 'left', 'white-space', 'transition-property', 'transition-duration', 'transition-timing-function'], apple: ['retail', 'field.label'] },
      { target: LABEL, on: INPUT, state: 'focus', props: [...TYPE, 'color', 'top', 'left'], apple: ['retail', 'field.label-floated'] },
      { target: INPUT, props: ['letter-spacing'], golden: TRACKING },
      { target: LABEL, props: ['letter-spacing'], golden: TRACKING },
      { target: LABEL, on: INPUT, state: 'focus', props: ['letter-spacing'], golden: TRACKING },
      { target: '.cui-field', props: ['background-color', 'border-top-left-radius'], golden: 'its own light ground under the translucent input, its corners the input\'s: it reads in a dark block too, where apple.com shows no field' },
      { target: INPUT, props: ['margin-bottom'], golden: 'no margin of its own: apple.com\'s 14px below each field is the form\'s spacing, the page\'s (retail:field.input)' },
    ],
  },
  {
    name: 'filled',
    html: ex('filled'),
    checks: [
      { target: INPUT, props: [...PAD, ...EDGE, 'color', 'background-color'], apple: ['retail', 'field.input-filled'] },
      { target: LABEL, props: [...TYPE, 'color', 'top', 'left'], apple: ['retail', 'field.label-floated'] },
    ],
  },
  {
    name: 'dark',
    html: dark(ex('field')),
    checks: [
      { target: INPUT, props: ['color', 'background-color', 'border-top-color'], apple: ['retail-dark', 'field.input'] },
      { target: LABEL, props: ['color'], apple: ['retail-dark', 'field.label'] },
    ],
  },
];
