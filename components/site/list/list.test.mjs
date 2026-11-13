import { examples } from '../../../test/harness.mjs';

const ex = examples(import.meta.url);
const ROW = '.cui-list-row';
const SECOND = `${ROW}:nth-child(2)`;
const TITLE = `${ROW}:first-child .cui-list-title`;
const TEXT = `${ROW}:first-child .cui-list-detail`;
const ASIDE = `${ROW}:first-child .cui-list-aside`;
const SIZE = ['font-size', 'font-weight'];
const CLAMP = ['-webkit-line-clamp', '-webkit-box-orient', 'overflow-x', 'overflow-y'];

export const cases = [
  {
    name: 'list',
    html: ex('list'),
    widths: [1600, 1200, 900, 735],
    checks: [
      { target: SECOND, props: ['padding-top', 'padding-bottom', 'border-top-width', 'border-top-style'], apple: ['forums-list', 'row.second'] },
      { target: ROW, props: ['column-gap'], apple: ['forums-list', 'row'] },
      { target: SECOND, props: ['padding-left', 'padding-right', 'border-top-color'], golden: 'flush with the page\'s column, the kit\'s divider: the forums draw their rows inset on a white card (20px by 40px, a #d3d3d3 rule), a site\'s page is white already' },
      { target: '.cui-list-row:first-child', props: ['border-top-width'], golden: 'no rule over the first row (forums-list:row has one, the card\'s top edge under it)' },
      { target: TITLE, pseudo: '::after', props: ['position', 'top', 'bottom', 'left', 'right'], golden: 'the whole row is the title\'s target, as a list row is on iOS: the forums\' title alone is, 21px tall' },
      { target: ASIDE, props: ['position'], golden: 'over the title\'s target, so a link in the aside is its own' },
      { target: ROW, props: ['grid-template-columns'], golden: 'from 735px the row\'s text, then its aside as wide as it needs (forums-list:row, its stats a fixed 274px of four columns)' },
      { target: TITLE, props: [...SIZE], apple: ['forums-list', 'row.title'] },
      { target: TITLE, props: ['color', 'line-height'], golden: 'the kit\'s text, where the forums\' title is black; the ladder\'s 21px (cui-body-tight), where theirs computes 21.0012px' },
      { target: TITLE, state: 'hover', props: ['text-decoration-line'], apple: ['forums-list', 'row.title', 'hover'] },
      { target: TEXT, props: [...SIZE, 'color', 'margin-top', ...CLAMP], apple: ['forums-list', 'row.excerpt'], as: { 'overflow-x': 'overflow', 'overflow-y': 'overflow' } },
      { target: TEXT, props: ['line-height'], golden: 'the ladder\'s 18px (cui-body-reduced-tight), where the forums\' excerpt computes 18.0008px' },
      { target: ASIDE, props: ['color', 'line-height'], apple: ['forums-list', 'row.stat'] },
      { target: ASIDE, props: ['font-size', 'text-align', 'margin-top', 'font-variant-numeric'], golden: 'the forums\' stats at 14px, the ladder\'s (they compute 14.875px), at the row\'s end, its figures tabular' },
    ],
  },
  {
    name: 'list phone',
    html: ex('list'),
    widths: [734, 390],
    checks: [
      { target: SECOND, props: ['padding-top', 'padding-bottom'], apple: ['forums-list', 'mrow'] },
      { target: ROW, props: ['grid-template-columns'], golden: 'one column: the aside under the text (forums-list:mrow)' },
      { target: TITLE, props: [...SIZE], apple: ['forums-list', 'mrow.title'] },
      { target: TEXT, props: ['margin-top', '-webkit-line-clamp'], apple: ['forums-list', 'mrow.excerpt'] },
      { target: ASIDE, props: ['margin-top'], apple: ['forums-list', 'mrow.stats'] },
      { target: ASIDE, props: ['font-size', 'color', 'text-align'], golden: 'the wide row\'s aside, under the text: the forums\' phone row draws its stats as icons at 16px, the kit draws no icon the page did not choose' },
    ],
  },
  {
    name: 'list zh',
    html: ex('list').replace('<ul class', '<ul lang="zh-Hant" class'),
    widths: [1200],
    checks: [
      { target: TITLE, props: ['line-height'], golden: 'Traditional Chinese a step taller, as cui-body-tight is' },
      { target: TEXT, props: ['line-height'], golden: 'Traditional Chinese a pixel taller, as cui-body-reduced-tight is' },
      { target: ASIDE, props: ['line-height'], golden: 'Traditional Chinese a pixel taller, as cui-body-reduced is' },
    ],
  },
];
