// The console's section head against App Store Connect's (reference/asc/
// results/head.json's section.head): the baseline row, its gaps, the room
// under it. Golden, with their why: the title taking the room, the wrap.

import { examples } from '../../../test/harness.mjs';

const ex = examples(import.meta.url);
const S = '.cui-console-section-head';

export const cases = [
  {
    name: 'section head',
    html: ex('section-head'),
    widths: [1200, 390],
    checks: [
      { target: S, props: ['display', 'align-items', 'row-gap', 'column-gap', 'margin-bottom'], asc: ['head', 'section.head'] },
      { target: `${S} > :first-child`, props: ['flex-grow', 'min-width'], golden: 'the title takes the room, the controls at the row\'s end (App Store Connect\'s title 941px of 1105px)' },
      { target: `${S} > .cui-console-tabs`, props: ['box-y'], golden: 'at 1200px on the title\'s row; at 390px under it (the row wraps; App Store Connect\'s does not)' },
    ],
  },
];
