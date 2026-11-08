import { examples } from '../../../test/harness.mjs';

const ex = examples(import.meta.url);
const S = '.cui-console-section-header';

export const cases = [
  {
    name: 'section header',
    html: ex('section-header'),
    widths: [1200, 390],
    checks: [
      { target: S, props: ['display', 'align-items', 'row-gap', 'column-gap', 'margin-bottom'], asc: ['head', 'section.head'] },
      { target: `${S} > :first-child`, props: ['flex-grow', 'min-width'], golden: 'the title takes the room, the controls at the row\'s end (App Store Connect\'s title 941px of 1105px)' },
      { target: `${S} h1`, props: ['font-size', 'line-height'], golden: 'the title\'s type its own class\'s (cui-console-h2), none from the header: App Store Connect\'s section heading is its h4, 18px on 22px' },
      { target: `${S} > .cui-console-tabs`, props: ['box-y'], golden: 'at 1200px on the title\'s row; at 390px under it (the row wraps; App Store Connect\'s does not)' },
    ],
  },
];
