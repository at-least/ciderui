import { examples } from '../../../test/harness.mjs';

const ex = examples(import.meta.url);
const D = '.cui-dl';

export const cases = [
  {
    name: 'dl',
    html: ex('dl'),
    widths: [1200, 735],
    checks: [
      { target: D, props: ['display', 'grid-template-columns', 'row-gap', 'column-gap'], golden: 'the pairs in two columns from 735px, 12px down and 24px across: totality\'s order details (pages.go\'s dlRow)' },
      { target: `${D} > div:first-child > dt`, props: ['color', 'font-size', 'line-height'], golden: 'the term in the secondary text over its value, the body\'s reduced size' },
      { target: `${D} > div:first-child > dd`, props: ['color', 'margin-left'], golden: 'the value in the text\'s color under its term, no indent' },
    ],
  },
  {
    name: 'dl phone',
    html: ex('dl'),
    widths: [734, 390],
    checks: [
      { target: D, props: ['grid-template-columns'], golden: 'one column below 735px' },
    ],
  },
  {
    name: 'dl zh',
    html: ex('dl').replace('<dl class', '<dl lang="zh-Hant" class'),
    widths: [1200],
    checks: [
      { target: `${D} > div:first-child > dt`, props: ['line-height'], golden: 'Traditional Chinese a pixel taller, as cui-body-reduced is' },
    ],
  },
];
