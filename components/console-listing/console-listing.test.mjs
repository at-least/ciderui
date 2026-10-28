// The console's code listing against Apple's developer documentation's
// (DocC, reference/apple.com/results/docc.json): its corners, its room, its
// scroll. Golden, with their why: the console's type and gray, the label,
// the dark side.

import { examples } from '../../test/harness.mjs';

const ex = examples(import.meta.url);
const L = '.cui-console-listing';

export const cases = [
  {
    name: 'listing',
    html: ex('listing'),
    widths: [1200, 390],
    checks: [
      { target: L, props: ['border-top-left-radius'], apple: ['docc', 'listing'] },
      { target: L, props: ['padding-top', 'padding-right', 'padding-bottom', 'padding-left'], apple: ['docc', 'listing.pre'] },
      { target: L, props: ['white-space'], apple: ['docc', 'listing.pre'] },
      { target: L, props: ['font-size', 'line-height', 'font-weight'], golden: 'the console\'s 13px mono body (monoBodyRegular), its line in DocC\'s proportion, 25px on 15px' },
      { target: L, props: ['background-color', 'color', 'overflow-x'], golden: 'the console card\'s gray, the nearest the contract holds to DocC\'s rgb(245, 245, 247); the label; a long line scrolls sideways inside it, as DocC\'s scroller does' },
    ],
  },
  {
    name: 'listing dark',
    html: `<div class="cui-dark">${ex('listing')}</div>`,
    widths: [1200],
    checks: [
      { target: L, props: ['background-color', 'color'], golden: 'the dark side derived (tokens.css, decision 18)' },
    ],
  },
];
