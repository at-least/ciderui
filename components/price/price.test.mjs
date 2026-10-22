// The price against apple.com's buy page (buy:price.header, its type and
// color; buy:price.amount, an amount kept whole). Golden, with its why: the
// letter-spacing; at 735px, by the window.

import { examples, SIZE } from '../../test/harness.mjs';

const ex = examples(import.meta.url);

export const cases = [
  {
    name: 'price',
    html: ex('price'),
    widths: [1600, 1200, 900, 390],
    checks: [
      { target: '.cui-price', props: [...SIZE, 'color', 'display'], apple: ['buy', 'price.header'] },
      { target: '.cui-price-amount', props: [...SIZE, 'white-space', 'display'], apple: ['buy', 'price.amount'] },
      { target: '.cui-price', props: ['letter-spacing'], golden: 'the system font\'s fitted letter-spacing (tokens.css), not apple.com\'s' },
    ],
  },
  {
    name: 'price-step',
    html: ex('price'),
    widths: [735],
    checks: [
      { target: '.cui-price', props: ['font-size'], golden: '17px from 735px by the window (apple.com\'s store steps by a phone\'s screen: 12px in a 735px window on a 735px screen)' },
    ],
  },
];
