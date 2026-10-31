// The console's bar against App Store Connect's (reference/asc/results/
// apps.json: nav, read at 1920px, so held at 1200 alone). Golden, with
// their why: the gray, the narrower screens' room, the gap, the dark side.

import { examples } from '../../../test/harness.mjs';

const ex = examples(import.meta.url);
const B = '.cui-console-bar';

export const cases = [
  {
    name: 'bar',
    html: ex('bar'),
    widths: [1200],
    checks: [
      { target: B, props: ['height', 'padding-left', 'padding-right', 'border-bottom-width', 'box-shadow'], asc: ['apps', 'nav'] },
      { target: B, props: ['background-color', 'column-gap', 'align-items'], golden: 'the console card\'s gray, the nearest the contract holds to App Store Connect\'s rgb(241, 241, 241); its parts 8px apart (a control row\'s md button\'s), centered' },
    ],
  },
  {
    name: 'bar narrow',
    html: ex('bar'),
    widths: [900, 390],
    checks: [
      { target: B, props: ['padding-left', 'box-height'], golden: 'the console column\'s room, 24px from 640px, 16px below (App Store Connect read at 1920px alone), 50px still' },
    ],
  },
  {
    name: 'bar dark',
    html: `<div class="cui-dark">${ex('bar')}</div>`,
    widths: [1200],
    checks: [
      { target: B, props: ['background-color', 'color'], golden: 'the dark side derived (tokens.css, decision 18)' },
    ],
  },
];
