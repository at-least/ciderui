import { examples } from '../../../test/harness.mjs';

const ex = examples(import.meta.url);
const S = '.cui-console-section';

export const cases = [
  {
    name: 'section',
    html: ex('section'),
    widths: [1200, 390],
    checks: [
      { target: S, props: ['display', 'row-gap'], golden: 'its header, then each block, 20px apart: the room App Store Connect\'s section head leaves under it (head.json\'s section.head, margin-bottom 20px), one value for every block so no page writes its own' },
      { target: `${S} > p:last-child`, props: ['box-y'], golden: 'the second block 20px under the first' },
    ],
  },
];
