// The checkbox against the sign-in widget's (signin:checkbox*): the input's
// place, the label's type and place, the box's size, edge, corners and
// ground; checked, its accent (signin:checkbox.checked); focused from the
// keyboard, its ring (Tab from the page's start). Golden, with their why:
// the letter-spacing, the box's width, the check shown and hidden, the ring,
// no margin below.

import { examples } from '../../test/harness.mjs';

const ex = examples(import.meta.url);
const TYPE = ['font-size', 'line-height', 'font-weight'];
const BOX = ['padding-top', 'padding-left', 'border-top-width', 'border-top-style', 'border-top-color', 'border-bottom-color', 'border-top-left-radius', 'background-color'];
const INDICATOR = '.cui-checkbox-indicator';

export const cases = [
  {
    name: 'checkbox',
    html: ex('checkbox'),
    checks: [
      { target: '.cui-checkbox', props: ['position', 'box-height'], apple: ['signin', 'checkbox.box'] },
      { target: '.cui-checkbox-input', props: ['position', 'top', 'left', 'height', 'opacity', 'margin-top', 'margin-left'], apple: ['signin', 'checkbox.input'] },
      { target: '.cui-checkbox-label', props: [...TYPE, 'color', 'padding-top', 'padding-left', 'position', 'display'], apple: ['signin', 'checkbox.label'] },
      { target: INDICATOR, props: [...BOX, 'position', 'top', 'left', 'display', 'box-height'], apple: ['signin', 'checkbox'] },
      { target: INDICATOR, props: ['box-width'], golden: 'apple.com\'s 16px box (signin:checkbox), a width no text sets' },
      { target: `${INDICATOR} svg`, props: ['visibility'], golden: 'its check hidden until it is checked' },
      { target: '.cui-checkbox-label', props: ['letter-spacing'], golden: 'the system font\'s fitted letter-spacing (tokens.css), not apple.com\'s' },
      { target: '.cui-checkbox', props: ['margin-bottom'], golden: 'no margin of its own: apple.com\'s 11px below it is the form\'s spacing, the page\'s (signin:checkbox.box)' },
    ],
  },
  {
    name: 'checked',
    html: ex('checked'),
    checks: [
      { target: INDICATOR, props: [...BOX, 'color'], apple: ['signin', 'checkbox.checked'] },
      { target: `${INDICATOR} svg`, props: ['visibility', 'color'], golden: 'its check shown, white (on-accent: apple.com\'s glyph\'s, signin:checkbox.check)' },
    ],
  },
  {
    // Tab from the page's start: the checkbox takes the focus, its box the ring
    name: 'focus',
    html: ex('checkbox'),
    widths: [1200],
    press: 'Tab',
    wait: 100,
    checks: [
      { target: '.cui-checkbox-input', props: ['focused'], golden: 'the input itself takes the focus from the keyboard' },
      { target: INDICATOR, props: ['outline-style', 'outline-width', 'outline-color', 'outline-offset'], golden: 'the focus ring on its box, apple.com\'s as its script draws it for the keyboard (2px #0071e3 3px out: a probe of a Tab, reference/apple.com/README.md)' },
    ],
  },
  {
    // clicked: no ring (a click is not the keyboard's)
    name: 'click',
    html: ex('checkbox'),
    widths: [1200],
    click: '.cui-checkbox-label',
    checks: [
      { target: '.cui-checkbox-input', props: ['focused'], golden: 'a click on the label checks the input and focuses it' },
      { target: INDICATOR, props: ['outline-style', 'background-color'], golden: 'checked by the click, and no ring: a click is not the keyboard\'s' },
    ],
  },
];
