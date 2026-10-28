// The console's pager against Apple Developer Forums' (reference/
// developer-forums/results/pager.json; golden: its values are the forums'
// own, which no probe in reference/apple.com/ or reference/asc/ holds): the
// row, its gaps, the summary's room and color, a missing direction left
// out. Its type the toolkit's body text style; its buttons the console's.

import { examples } from '../../test/harness.mjs';

const ex = examples(import.meta.url);
const N = '.cui-console-pager';
const S = '.cui-console-pager-summary';

export const cases = [
  {
    name: 'pager',
    html: ex('pager'),
    widths: [1200, 390],
    checks: [
      { target: S, props: ['font-size', 'line-height', 'font-weight', 'letter-spacing'], tb: ['typography-textStyle-bodyRegular-fontSize', 'typography-textStyle-bodyRegular-lineHeight', 'typography-textStyle-bodyRegular-fontWeight', 'typography-textStyle-bodyRegular-letterSpacing'] },
      { target: S, props: ['color', 'padding-left', 'padding-right'], golden: 'the forums\' page number: the text\'s color, 15px either side (pager.json pager.page), the console\'s label' },
      { target: N, props: ['justify-content', 'align-items', 'column-gap'], golden: 'the forums\' row: centered, 10px between its parts (pager.json pager.list)' },
      { target: `${N} .cui-console-btn[href]`, props: ['box-height', 'box-width'], golden: 'the console\'s 30px icon button, square, where the forums\' is 39 by 28px' },
    ],
  },
  {
    name: 'pager first',
    html: ex('first'),
    widths: [1200],
    checks: [
      { target: S, props: ['box-left'], golden: 'the first page: no previous button (the forums\' page 1 draws none, pager.json page-1), the summary and the next button centered' },
      { target: `${N} .cui-console-btn`, props: ['box-left'], golden: 'the next button 10px after the summary\'s box, the pair centered' },
    ],
  },
  {
    name: 'pager dark',
    html: `<div class="cui-dark">${ex('pager')}</div>`,
    widths: [1200],
    checks: [
      { target: S, props: ['color'], golden: 'the dark side derived (tokens.css, decision 18)' },
    ],
  },
];
