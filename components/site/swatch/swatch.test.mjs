import { examples } from '../../../test/harness.mjs';

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
      { target: `${REST} .cui-swatch-color`, props: ['outline-width', 'outline-style', 'outline-color', 'outline-offset'], golden: 'a 1px hairline just inside the circle, --cui-c-separator: a white swatch reads on a white page, where apple.com\'s draws its edge into its picture (buy:swatch)' },
      { target: `${REST} .cui-swatch-input`, props: ['hit'], golden: 'the input lies over the ring: the pointer rests on it, so the color\'s name, its title, shows' },
    ],
  },
  {
    name: 'forced',
    html: ex('swatches'),
    forcedColors: true,
    checks: [
      { target: `${CHECKED} .cui-swatch-ring`, props: ['forced-color-adjust', 'border-top-color'], golden: 'in forced colors the checked ring the accent still, its color unforced, where forced it is drawn as every other (apple.com\'s the same: not measured there)' },
      { target: `${REST} .cui-swatch-ring`, props: ['border-top-color'], golden: 'in forced colors an unchecked ring clear still' },
      { target: `${CHECKED} .cui-swatch-color`, props: ['background-color'], golden: 'in forced colors the swatch\'s color its own, where forced every swatch is the window\'s ground' },
    ],
  },
  {
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
  {
    name: 'disabled',
    html: ex('swatches').replace('value="silver" title="Silver"', 'value="silver" title="Silver" disabled'),
    checks: [
      { target: `${REST} .cui-swatch-ring`, props: ['opacity'], golden: 'disabled, dimmed to 0.42 as a disabled option tile (apple.com\'s swatches of a model not chosen: not measured)' },
      { target: `${REST} .cui-swatch-input`, props: ['cursor'], golden: 'disabled, the arrow, as an option tile' },
    ],
  },
  {
    name: 'swatch touch',
    html: ex('swatches'),
    widths: [390],
    touch: true,
    checks: [
      { target: REST, props: ['padding-top', 'padding-left', 'box-height'], golden: 'on a touch screen 4px of room round each ring, for its radio\'s 44px' },
    ],
  },
];
