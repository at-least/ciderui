// The console's head against App Store Connect's app head (reference/asc/
// results/head.json): the row, the name's type, the tabs' type, height and
// rule. Golden, with their why: the gaps read from places, the colors, the
// phone's row.

import { examples } from '../../test/harness.mjs';

const ex = examples(import.meta.url);
const H = '.cui-console-head';
const NAME = '.cui-console-head-name';
const TAB = `${H} .cui-console-tab:not([aria-current])`;
const CURRENT = `${H} .cui-console-tab[aria-current]`;

export const cases = [
  {
    name: 'head',
    html: ex('head'),
    widths: [1200, 390],
    checks: [
      { target: H, props: ['display', 'align-items'], asc: ['head', 'head'] },
      { target: NAME, props: ['font-size', 'line-height', 'font-weight'], asc: ['head', 'head.name.text'] },
      { target: TAB, props: ['font-size', 'line-height', 'font-weight'], asc: ['head', 'head.tab.link'] },
      { target: CURRENT, props: ['height', 'border-bottom-width', 'border-bottom-style'], asc: ['head', 'head.tab.current'] },
      { target: `${H} > .cui-console-tabs`, props: ['border-bottom-width'], asc: ['head', 'head.tabs'] },
      { target: H, props: ['column-gap', 'margin-bottom', 'flex-wrap'], golden: '19px between the name and the tabs and 24px to the content, both from their places in App Store Connect; a row that wraps (its phone was not read)' },
      { target: `${H} > .cui-console-tabs`, props: ['column-gap'], golden: 'the tabs 23px apart, App Store Connect\'s li\'s 23px at its end' },
      { target: NAME, props: ['color', 'letter-spacing', 'text-decoration-line', 'text-overflow'], golden: 'the label (App Store Connect\'s older rgb(17, 17, 17)), the font tracking itself, no underline, an ellipsis where the row has no room' },
      { target: CURRENT, props: ['color', 'border-bottom-color'], golden: 'the console blue, its rule the same (the older look\'s rgb(0, 112, 201))' },
      { target: `${H} > .cui-console-tabs`, props: ['box-y'], golden: 'at 1200px beside the name; at 390px a row of its own under it' },
    ],
  },
  {
    name: 'head wide',
    html: ex('head'),
    widths: [1200],
    checks: [
      { target: H, props: ['height'], asc: ['head', 'head'] },
    ],
  },
  {
    name: 'head dark',
    html: `<div class="cui-dark">${ex('head')}</div>`,
    widths: [1200],
    checks: [
      { target: NAME, props: ['color'], golden: 'the dark side derived (tokens.css, decision 18)' },
    ],
  },
];
