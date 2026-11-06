import { examples, PADDING, SIZE } from '../../../test/harness.mjs';

const ex = examples(import.meta.url);
const EDGE = ['border-top-width', 'border-top-style', 'border-top-color', 'border-bottom-width', 'border-bottom-color', 'border-top-left-radius'];
const INPUT = '.cui-field-input';
const LABEL = '.cui-field-label';
const dark = (html) => `<div class="cui-dark">${html}</div>`;
const MESSAGE = '.cui-field-message';
const MARK = `${MESSAGE} svg`;
const PREFIX = '.cui-field-prefix';
const FONT = ['font-size', 'line-height'];
// the stylesheet's :-webkit-autofill rules copied onto a class: no script can make
// a browser autofill. Its reason is the copy's, so a rule dropped fails here
const AUTOFILLED = `<script>
  const copy = (rules) => { for (const rule of [...rules]) {
    if (rule.cssRules) copy(rule.cssRules);
    if (rule.selectorText?.includes(':-webkit-autofill')) (rule.parentRule ?? rule.parentStyleSheet).insertRule(rule.cssText.replaceAll(':-webkit-autofill', '.autofilled'));
  } };
  for (const sheet of document.styleSheets) copy(sheet.cssRules);
  document.querySelector('.cui-field-input').classList.add('autofilled');
</script>`;
const LONG = 'A label far longer than the field it names on a phone, and longer still';

export const cases = [
  {
    name: 'field',
    html: ex('field'),
    checks: [
      { target: INPUT, props: ['display'], golden: 'no line box under it' },
      { target: '.cui-field', props: ['position'], golden: 'what its label and prefix are placed against, wherever it sits on a page' },
      { target: INPUT, props: [...SIZE, 'color', ...PADDING, ...EDGE, 'background-color', 'height', 'box-sizing', 'box-height', 'transition-property', 'transition-duration', 'outline-style'], apple: ['retail', 'field.input'] },
      { target: INPUT, state: 'hover', props: ['color', 'background-color', 'border-top-color'], apple: ['retail', 'field.input', 'hover'] },
      { target: INPUT, state: 'focus', props: [...PADDING, ...EDGE], apple: ['retail', 'field.input', 'focus'] },
      // against the rest probe: focused, the 2px edge takes the padding's pixel, the box unchanged
      { target: INPUT, state: 'focus', props: ['box-height'], apple: ['retail', 'field.input'] },
      { target: INPUT, state: ['focus', 'focus-visible'], props: ['padding-left', 'border-top-width', 'border-top-color', 'outline-style'], apple: ['retail', 'field.input', 'focus+focus-visible'] },
      { target: LABEL, props: [...SIZE, 'color', 'position', 'top', 'left', 'white-space', 'transition-property', 'transition-duration', 'transition-timing-function'], apple: ['retail', 'field.label'] },
      { target: LABEL, on: INPUT, state: 'focus', props: [...SIZE, 'color', 'top', 'left'], apple: ['retail', 'field.label-floated'] },
      { target: '.cui-field', props: ['background-color', 'border-top-left-radius'], golden: 'its own light ground under the translucent input, its corners the input\'s: it reads in a dark block too, where apple.com shows no field' },
      { target: INPUT, props: ['margin-bottom'], golden: 'no margin of its own: apple.com\'s 14px below each field is the form\'s spacing, the page\'s (retail:field.input)' },
    ],
  },
  {
    name: 'long label',
    html: ex('field').replace('>Email</label>', `>${LONG}</label>`),
    widths: [390],
    checks: [
      { target: LABEL, props: ['box-width', 'overflow-x', 'text-overflow', 'white-space'], golden: 'a long label stops 17px short of the field\'s end, with an ellipsis, where it ran past the field (apple.com\'s labels are short: not measured)' },
    ],
  },
  {
    name: 'autofilled',
    html: `${ex('field')}${AUTOFILLED}`,
    widths: [1200],
    checks: [
      { target: LABEL, props: ['top', 'font-size', 'line-height'], golden: 'an autofilled field\'s label floats, as over any value, so it does not lie over the filled text (a browser\'s real fill not read: no script can make one)' },
      { target: INPUT, props: ['box-shadow'], golden: 'its ground the page\'s, an inset shadow over the browser\'s own autofill color, which is !important' },
    ],
  },
  {
    name: 'filled',
    html: ex('filled'),
    checks: [
      { target: INPUT, props: [...PADDING, ...EDGE, 'color', 'background-color'], apple: ['retail', 'field.input-filled'] },
      { target: LABEL, props: [...SIZE, 'color', 'top', 'left'], apple: ['retail', 'field.label-floated'] },
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
      { target: MARK, props: ['flex-shrink'], golden: 'a long message does not squeeze its icon' },
      { target: INPUT, props: [...SIZE, 'color', 'height', 'border-top-width', 'border-top-color', 'border-bottom-color', 'background-color'], apple: ['giftcard', 'error'] },
      { target: INPUT, state: 'focus', props: [...PADDING, 'border-top-width', 'border-top-color', 'background-color'], apple: ['giftcard', 'error', 'focus'] },
      { target: LABEL, props: [...SIZE, 'color', 'top', 'left'], apple: ['giftcard', 'error.label'] },
      { target: LABEL, on: INPUT, state: 'focus', props: ['color'], golden: 'its label red focused too: apple.com\'s keeps the class that makes it so (giftcard:error.label), not read focused' },
      { target: MESSAGE, props: [...SIZE, 'color'], apple: ['giftcard', 'error.message'] },
      { target: MESSAGE, props: ['margin-top', 'display'], apple: ['giftcard', 'error.wrapper'] },
      { target: MESSAGE, props: ['margin-bottom'], golden: 'no margin below it: apple.com\'s 12px is the form\'s spacing, the page\'s (giftcard:error.wrapper)' },
      { target: MARK, props: ['margin-right'], as: { 'margin-right': 'padding-right' }, apple: ['giftcard', 'error.icon'] },
      { target: MARK, props: ['box-width', 'box-height', 'margin-top', 'color'], golden: 'apple.com\'s mark (U+F312) drawn as ciderui\'s ALERT: its advance by its ink\'s height at 12px, its ink 2.75px down the row as apple.com\'s (reference/icons/form.mjs), in the message\'s red' },
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
    // the gift card's page read at 1600px alone (before its choices)
    name: 'disabled-page',
    html: ex('disabled'),
    widths: [1600],
    checks: [
      { target: INPUT, props: [...SIZE, 'color', 'background-color', 'border-top-color', 'border-top-width', 'cursor', 'height'], apple: ['giftcard', 'field.disabled'] },
      { target: LABEL, props: [...SIZE, 'color'], apple: ['giftcard', 'label.disabled'] },
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
      { target: PREFIX, props: ['pointer-events'], golden: 'a click on it reaches the input under it' },
      { target: INPUT, props: [...PADDING, ...EDGE, 'height', 'background-color'], apple: ['giftcard', 'prefix.input'] },
      { target: INPUT, state: 'focus', props: ['padding-left', 'border-top-width', 'border-top-color'], apple: ['giftcard', 'prefix.input', 'focus'] },
      { target: LABEL, props: [...SIZE, 'color', 'top', 'left', 'padding-left'], apple: ['giftcard', 'prefix.label'] },
      { target: PREFIX, props: [...FONT, 'color', 'padding-top', 'padding-left', 'padding-right', 'background-color', 'border-right-width', 'border-right-style', 'border-right-color', 'border-top-left-radius', 'border-bottom-left-radius', 'border-top-right-radius', 'position', 'top', 'left', 'height'], apple: ['giftcard', 'prefix.currency'] },
      { target: PREFIX, props: ['box-width', 'letter-spacing'], golden: 'its text\'s width and 16px either side, as apple.com\'s (43.34px in SF: giftcard:prefix.currency); the fitted letter-spacing' },
    ],
  },
];
