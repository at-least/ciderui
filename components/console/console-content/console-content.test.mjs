// The console's content column against App Store Connect's toolkit's
// legacy content: its widest, its side room at each step. Golden, with
// their why: the sidebar's column and the 30px (the analytics overview's
// menu), the sidebar's room dropped, the sheet below 1024px.

import { examples } from '../../../test/harness.mjs';

const ex = examples(import.meta.url);
const C = '.cui-console-content';
const L = 'component-Content-legacy';

export const cases = [
  {
    name: 'content wide',
    html: ex('content'),
    widths: [1920, 1440, 1024],
    checks: [
      { target: C, props: ['max-width', 'padding-left', 'padding-right'], tb: [`${L}-maxWidth`, `${L}-xl-paddingInline`, `${L}-xl-paddingInline`] },
      { target: C, props: ['column-gap'], tb: ['space-spacer-30'] },
      { target: C, props: ['box-left', 'grid-template-columns'], golden: 'centered at its widest; the sidebar\'s 225px (App Store Connect\'s menu: analytics menu.item) beside the page' },
      { target: `${C} .cui-console-sidebar-item[aria-current]`, props: ['box-left', 'box-width'], golden: 'the sidebar\'s items flush with the column, as the analytics menu\'s are' },
      { target: `${C} main`, props: ['box-left'], golden: 'the page 30px past the sidebar\'s column' },
    ],
  },
  {
    name: 'content tablet',
    html: ex('content'),
    widths: [800],
    checks: [
      { target: C, props: ['padding-left', 'padding-right'], tb: [`${L}-md-paddingInline`, `${L}-md-paddingInline`] },
      { target: C, props: ['display'], golden: 'one column below 1024px: the sidebar is a sheet there' },
      { target: `${C} .cui-console-sidebar`, props: ['display'], golden: 'the closed sheet' },
    ],
  },
  {
    name: 'content phone',
    html: ex('content'),
    widths: [390],
    checks: [
      { target: C, props: ['padding-left', 'padding-right'], tb: [`${L}-sm-paddingInline`, `${L}-sm-paddingInline`] },
    ],
  },
];
