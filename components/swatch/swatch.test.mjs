// The swatch against apple.com's buy page (buy:swatch*): the list, the
// ring at rest and checked (a checked copy), the radio over it, the color's
// circle and its ground before its image, the name hidden but read.
// Golden, with their why: the fixed widths, the gaps, an item's height,
// the focus ring.

import { examples } from '../../test/harness.mjs';

const ex = examples(import.meta.url);
const RING = ['padding-top', 'padding-left', 'border-top-width', 'border-top-style', 'border-top-color', 'border-bottom-color', 'border-top-left-radius', 'background-color', 'height'];
const CHECKED = '.cui-swatch:first-child';
const REST = '.cui-swatch:nth-child(2)';

export const cases = [
  {
    name: 'swatches',
    html: ex('swatches'),
    checks: [
      { target: '.cui-swatches', props: ['display', 'flex-wrap'], apple: ['buy', 'swatch.list'] },
      { target: '.cui-swatches', props: ['row-gap', 'column-gap'], golden: 'apple.com\'s items\' margins as the list\'s gaps: 5px apart, 8px between rows (buy:swatch.item: 4px above and below, 5px after each)' },
      { target: `${REST} .cui-swatch-ring`, props: [...RING, 'box-height'], apple: ['buy', 'swatch.link'] },
      { target: `${REST} .cui-swatch-ring`, state: 'hover', props: ['border-top-color', 'background-color'], apple: ['buy', 'swatch.link', 'hover'] },
      { target: `${CHECKED} .cui-swatch-ring`, props: [...RING, 'box-height'], apple: ['buy', 'swatch.checked-link'] },
      { target: `${REST} .cui-swatch-input`, props: ['position', 'height', 'opacity', 'margin-top', 'box-height'], apple: ['buy', 'swatch.input'] },
      { target: '.cui-swatch:nth-child(3) .cui-swatch-color', props: ['border-top-left-radius', 'background-color', 'height', 'box-height'], apple: ['buy', 'swatch'] },
      { target: `${REST} .cui-swatch-name`, props: ['position', 'height', 'clip', 'clip-path', 'overflow'], apple: ['buy', 'swatch.name'] },
      { target: `${REST} .cui-swatch-ring`, props: ['width'], golden: 'apple.com\'s 36px ring (buy:swatch.link), a width no text sets' },
      { target: '.cui-swatch:nth-child(3) .cui-swatch-color', props: ['width'], golden: 'apple.com\'s 28px circle (buy:swatch)' },
      { target: REST, props: ['box-height'], golden: 'the ring a block, so an item is 36px (apple.com\'s inline ring leaves its item 39px)' },
    ],
  },
  {
    // Tab from the page's start: the checked radio takes the focus
    name: 'focus',
    html: ex('swatches'),
    widths: [1200],
    press: 'Tab',
    wait: 100,
    checks: [
      { target: `${CHECKED} .cui-swatch-input`, props: ['focused'], golden: 'the group\'s checked radio takes the keyboard\'s focus' },
      { target: `${CHECKED} .cui-swatch-ring`, props: ['outline-style', 'outline-width', 'outline-color', 'outline-offset'], golden: 'the focus ring around the ring, as on an option (apple.com\'s swatches take no keyboard focus until a model is chosen: not measured)' },
    ],
  },
];
