import { examples } from '../../../test/harness.mjs';

const ex = examples(import.meta.url);
const PREVIOUS = '.cui-pager-button[aria-label="Previous page"]';
const FIRST = '.cui-pager-button:first-child';
const PAGE = '.cui-pager-page';

export const cases = [
  {
    name: 'pager',
    html: ex('pager'),
    widths: [1200, 390],
    checks: [
      { target: '.cui-pager', props: ['padding-top', 'padding-bottom'], apple: ['forums-pager', 'nav'] },
      { target: '.cui-pager', props: ['align-items', 'justify-content', 'row-gap', 'column-gap'], apple: ['forums-pager', 'list'] },
      { target: PREVIOUS, props: ['color', 'background-color', 'padding-top', 'padding-right', 'padding-bottom', 'padding-left', 'border-top-left-radius', 'min-width', 'height', 'box-height'], apple: ['forums-pager', 'previous'] },
      { target: PREVIOUS, state: 'hover', props: ['background-color'], apple: ['forums-pager', 'previous', 'hover'] },
      { target: PREVIOUS, state: 'active', props: ['background-color'], apple: ['forums-pager', 'previous', 'active'] },
      { target: `${PREVIOUS} svg`, props: ['box-height'], golden: 'the chevron\'s ink, 0.9em at the button\'s size: SF Symbols\' chevron.left drawn past its 0.70459em box (reference/icons/pager.mjs)' },
      { target: FIRST, props: ['height', 'background-color', 'border-top-left-radius'], golden: 'a word\'s button the chevrons\' height, 28px, where the forums\' First is 23px (forums-pager:first): one height in the row' },
      { target: PAGE, props: ['padding-left', 'padding-right'], apple: ['forums-pager', 'page'] },
      { target: PAGE, props: ['color'], golden: 'the kit\'s text, where the forums\' is rgb(51, 51, 51)' },
      { target: `${PAGE}, ${FIRST}`, props: ['font-size', 'line-height'], golden: 'the ladder\'s 14/20 (cui-body-reduced), where the forums draw 15px, a size the kit has no tracking for' },
      { target: PREVIOUS, pseudo: '::before', props: ['width', 'height'], golden: 'a 44 by 44 target about the 28px button, the HIG\'s' },
    ],
  },
  {
    name: 'pager first',
    html: ex('first'),
    widths: [1200],
    checks: [
      { target: '.cui-pager > :first-child', props: ['display'], golden: 'the page\'s number first: no First or Previous button drawn on the first page, as the forums draw none (../developer-forums/: page-1)' },
    ],
  },
];
