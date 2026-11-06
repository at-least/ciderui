import { examples, SIZE, WITH_1600 } from '../../../test/harness.mjs';

const ex = examples(import.meta.url);

export const cases = [
  {
    name: 'price',
    html: ex('price'),
    widths: WITH_1600,
    checks: [
      { target: '.cui-price', props: [...SIZE, 'color', 'display'], apple: ['buy', 'price.header'] },
      { target: '.cui-price-amount', props: [...SIZE, 'white-space', 'display'], apple: ['buy', 'price.amount'] },
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
