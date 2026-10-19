// The field against apple.com's own (retail:field*, the retail page's
// form-textbox): the input at rest, hovered and focused (alone, and with
// :focus-visible, as a text field always is), its label as the placeholder,
// floated when the field is focused and when it is filled; in a dark block,
// as apple.com's in a dark color scheme (retail-dark:field*). Its states:
// in error, its message under it (giftcard:error*, an empty required field
// the focus left; retail:field.error*, retail-dark:field.error*, copies),
// focused in error; disabled (giftcard:field.disabled, retail:field.disabled*,
// retail-dark:field.disabled*); with a currency before it
// (giftcard:prefix*); the dropdown, empty, hovered, focused and chosen
// (tradein:dropdown*). Golden, with their why: the letter-spacing, the
// field's own ground, no margin below, the message's mark and its dark red,
// the prefix's width, the chevron's place and color, and the states
// apple.com was not seen in.

import { examples } from '../../test/harness.mjs';

const ex = examples(import.meta.url);
const TYPE = ['font-size', 'line-height', 'font-weight'];
const EDGE = ['border-top-width', 'border-top-style', 'border-top-color', 'border-bottom-width', 'border-bottom-color', 'border-top-left-radius'];
const PAD = ['padding-top', 'padding-right', 'padding-bottom', 'padding-left'];
const INPUT = '.cui-field-input';
const LABEL = '.cui-field-label';
const TRACKING = 'the system font\'s fitted letter-spacing (tokens.css), not apple.com\'s';
const dark = (html) => `<div class="cui-dark">${html}</div>`;
const MESSAGE = '.cui-field-message';
const MARK = `${MESSAGE} svg`;
const PREFIX = '.cui-field-prefix';
const CHEVRON = '.cui-field-chevron';
const FONT = ['font-size', 'line-height'];

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
  {
    name: 'error',
    html: ex('error'),
    checks: [
      { target: INPUT, props: [...TYPE, 'color', 'height', 'border-top-width', 'border-top-color', 'border-bottom-color', 'background-color'], apple: ['giftcard', 'error'] },
      { target: INPUT, state: 'focus', props: [...PAD, 'border-top-width', 'border-top-color', 'background-color'], apple: ['giftcard', 'error', 'focus'] },
      { target: LABEL, props: [...TYPE, 'color', 'top', 'left'], apple: ['giftcard', 'error.label'] },
      { target: LABEL, on: INPUT, state: 'focus', props: ['color'], golden: 'its label red focused too: apple.com\'s keeps the class that makes it so (giftcard:error.label), not read focused' },
      { target: MESSAGE, props: [...TYPE, 'color'], apple: ['giftcard', 'error.message'] },
      { target: MESSAGE, props: ['margin-top', 'display'], apple: ['giftcard', 'error.wrapper'] },
      { target: MESSAGE, props: ['margin-bottom'], golden: 'no margin below it: apple.com\'s 12px is the form\'s spacing, the page\'s (giftcard:error.wrapper)' },
      { target: MESSAGE, props: ['letter-spacing'], golden: TRACKING },
      { target: MARK, props: ['margin-right'], as: { 'margin-right': 'padding-right' }, apple: ['giftcard', 'error.icon'] },
      { target: MARK, props: ['box-width', 'box-height', 'margin-top', 'color'], golden: 'apple.com\'s mark (U+F312) drawn as ciderui\'s ALERT: its advance by its ink\'s height at 12px, its ink 2.75px down the row as apple.com\'s (reference/icons/form.mjs), in the message\'s red' },
    ],
  },
  {
    // copies with apple.com's classes on the retail page, light and in a dark
    // color scheme: the field keeps its light look in error too
    name: 'error-copies',
    html: ex('error'),
    checks: [
      { target: INPUT, props: ['color', 'background-color', 'border-top-color', 'border-top-width'], apple: ['retail', 'field.error'] },
      { target: LABEL, props: ['color'], apple: ['retail', 'field.error-label'] },
    ],
  },
  {
    name: 'error-dark',
    html: dark(ex('error')),
    checks: [
      { target: INPUT, props: ['color', 'background-color', 'border-top-color', 'border-top-width'], apple: ['retail-dark', 'field.error'] },
      { target: LABEL, props: ['color'], apple: ['retail-dark', 'field.error-label'] },
      { target: MESSAGE, props: ['color'], golden: 'on the dark block, the fitted dark red (tokens.css: apple.com\'s #e30000 is below 4.5:1 on black; its forms show no message on a dark page)' },
    ],
  },
  {
    name: 'disabled',
    html: ex('disabled'),
    checks: [
      { target: INPUT, props: ['color', 'background-color', 'border-top-color', 'cursor'], apple: ['retail', 'field.disabled'] },
      { target: LABEL, props: ['color'], apple: ['retail', 'field.disabled-label'] },
    ],
  },
  {
    // the gift card's page's own, read at 1600px alone (before its choices)
    name: 'disabled-page',
    html: ex('disabled'),
    widths: [1600],
    checks: [
      { target: INPUT, props: [...TYPE, 'color', 'background-color', 'border-top-color', 'border-top-width', 'cursor', 'height'], apple: ['giftcard', 'field.disabled'] },
      { target: LABEL, props: [...TYPE, 'color'], apple: ['giftcard', 'label.disabled'] },
    ],
  },
  {
    name: 'disabled-dark',
    html: dark(ex('disabled')),
    checks: [
      { target: INPUT, props: ['color', 'background-color', 'border-top-color', 'cursor'], apple: ['retail-dark', 'field.disabled'] },
      { target: LABEL, props: ['color'], apple: ['retail-dark', 'field.disabled-label'] },
    ],
  },
  {
    // read at 1600px alone (the gift card's page enables it at its first width)
    name: 'prefix',
    html: ex('prefix'),
    widths: [1600],
    checks: [
      { target: INPUT, props: [...PAD, ...EDGE, 'height', 'background-color'], apple: ['giftcard', 'prefix.input'] },
      { target: INPUT, state: 'focus', props: ['padding-left', 'border-top-width', 'border-top-color'], apple: ['giftcard', 'prefix.input', 'focus'] },
      { target: LABEL, props: [...TYPE, 'color', 'top', 'left', 'padding-left'], apple: ['giftcard', 'prefix.label'] },
      { target: PREFIX, props: [...FONT, 'color', 'padding-top', 'padding-left', 'padding-right', 'background-color', 'border-right-width', 'border-right-style', 'border-right-color', 'border-top-left-radius', 'border-bottom-left-radius', 'border-top-right-radius', 'position', 'top', 'left', 'height'], apple: ['giftcard', 'prefix.currency'] },
      { target: PREFIX, props: ['box-width', 'letter-spacing'], golden: 'its text\'s width and 16px either side, as apple.com\'s (43.34px in SF: giftcard:prefix.currency); the fitted letter-spacing' },
    ],
  },
  {
    name: 'select',
    html: ex('select'),
    checks: [
      { target: INPUT, props: [...TYPE, 'color', ...PAD, ...EDGE, 'background-color', 'height', 'box-sizing', 'appearance', '-webkit-appearance', 'cursor', 'text-overflow', 'outline-style'], apple: ['tradein', 'dropdown.select'] },
      { target: INPUT, state: 'hover', props: ['color', 'background-color', 'border-top-color', 'border-top-width'], apple: ['tradein', 'dropdown.select', 'hover'] },
      { target: INPUT, state: 'focus', props: [...PAD, 'border-top-width', 'border-top-color', 'background-color', 'outline-style'], apple: ['tradein', 'dropdown.select', 'focus'] },
      { target: LABEL, props: [...TYPE, 'color', 'position', 'top', 'left', 'white-space', 'transition-property', 'transition-duration', 'transition-timing-function'], apple: ['tradein', 'dropdown.label'] },
      { target: LABEL, on: INPUT, state: 'focus', props: [...TYPE, 'top'], golden: 'empty, its label the placeholder focused too: the select\'s text is not shown until a choice is made (apple.com\'s class, set by its script on a choice)' },
      { target: CHEVRON, props: ['position', 'top', 'right', 'pointer-events'], apple: ['tradein', 'dropdown.chevron'] },
      { target: CHEVRON, props: ['box-width', 'box-height', 'color'], golden: 'apple.com\'s 25px chevron in #6e6e73 (its SVG\'s fill, tradein:dropdown.chevron), drawn as ciderui\'s DROPDOWN (reference/icons/form.mjs)' },
      { target: INPUT, props: ['letter-spacing'], golden: TRACKING },
    ],
  },
  {
    name: 'selected',
    html: ex('selected'),
    checks: [
      { target: INPUT, props: [...TYPE, 'color', ...PAD, 'height'], apple: ['tradein', 'dropdown.chosen'] },
      { target: LABEL, props: [...TYPE, 'color', 'top', 'left'], apple: ['tradein', 'dropdown.chosen-label'] },
    ],
  },
  {
    name: 'select-dark',
    html: dark(ex('selected')),
    widths: [1200],
    checks: [
      { target: INPUT, props: ['color', 'background-color', 'border-top-color'], golden: 'light in a dark block, as a field (retail-dark:field.input)' },
      { target: CHEVRON, props: ['color'], golden: 'light in a dark block, as a field' },
    ],
  },
];
