// The quick links against apple.com's store (store:quicklinks*, its place,
// title, list and items; store:quicklink, a pill at rest and hovered).
// Golden, with their why: the letter-spacing, the block's width, the second
// pill's place (no space character before it).

import { examples, PADDING, SIZE } from '../../test/harness.mjs';

const ex = examples(import.meta.url);
const PILL = '.cui-quicklinks-list li:first-child .cui-btn';

export const cases = [
  {
    // from 735px: the block as far in as apple.com's
    name: 'quicklinks-place',
    html: ex('quicklinks'),
    widths: [1600, 1200, 900],
    checks: [
      { target: '.cui-quicklinks', props: ['margin-left', 'margin-right', 'box-left'], apple: ['store', 'quicklinks'] },
    ],
  },
  {
    name: 'quicklinks-phone',
    html: ex('quicklinks'),
    widths: [390],
    checks: [
      { target: '.cui-quicklinks', props: ['margin-left', 'margin-right'], golden: '22px in on each side, the block\'s: apple.com\'s phone sets its title\'s margin and its list\'s padding to the same 22px (store:quicklinks.title, .list)' },
    ],
  },
  {
    name: 'quicklinks',
    html: ex('quicklinks'),
    checks: [
      { target: '.cui-quicklinks', props: ['padding-top', 'padding-bottom'], apple: ['store', 'quicklinks'] },
      { target: '.cui-quicklinks-title', props: [...SIZE, 'color', 'padding-bottom', 'margin-top', 'margin-bottom', 'box-height'], apple: ['store', 'quicklinks.title'] },
      { target: '.cui-quicklinks-list', props: ['padding-top', 'padding-bottom', 'margin-top'], apple: ['store', 'quicklinks.list'] },
      { target: '.cui-quicklinks-list li:first-child', props: ['padding-top', 'padding-right', 'padding-bottom', 'padding-left', 'box-height'], apple: ['store', 'quicklinks.item'] },
      { target: PILL, props: [...SIZE, 'color', ...PADDING, 'border-top-width', 'border-top-color', 'border-top-left-radius', 'background-color', 'min-width', 'box-height'], apple: ['store', 'quicklink'] },
      { target: PILL, state: 'hover', props: ['color', 'border-top-color', 'background-color'], apple: ['store', 'quicklink', 'hover'] },
      { target: '.cui-quicklinks-list li:nth-child(2)', props: ['box-x'], golden: 'right after the first pill\'s item: a flex row with no space character between them (apple.com\'s inline blocks leave one, 4.4px of SF, from 735px)' },
      { target: '.cui-quicklinks', props: ['box-width'], golden: 'apple.com\'s 980px at most, between its inset and 22px (store:quicklinks)' },
      { target: '.cui-quicklinks-title', props: ['letter-spacing'], golden: 'the system font\'s fitted letter-spacing (tokens.css), not apple.com\'s' },
    ],
  },
];
