// The console's code listing against Apple's developer documentation's
// (DocC, reference/apple.com/results/docc.json): its corners, its room, its
// scroll. Golden, with their why: the console's type and gray, the label,
// the dark side.

import { examples } from '../../../test/harness.mjs';

const ex = examples(import.meta.url);
const L = '.cui-console-listing';
const N = '.cui-console-listing-line:not([data-highlighted])';
const H = '.cui-console-listing-line[data-highlighted]';

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
    name: 'listing lines',
    html: ex('lines'),
    widths: [1200, 390],
    checks: [
      { target: N, props: ['display', 'box-sizing'], apple: ['docc-lines', 'line'] },
      { target: N, props: ['width'], golden: 'the listing\'s width, a highlight across it (DocC\'s 100%)' },
      { target: N, pseudo: '::before', props: ['display', 'text-align', 'user-select', 'padding-right', 'padding-left'], apple: ['docc-lines', 'line.number'] },
      { target: H, pseudo: '::before', props: ['padding-left', 'border-left-width', 'border-left-style'], apple: ['docc-lines', 'highlighted.number'] },
      { target: N, pseudo: '::before', props: ['content', 'min-width', 'color', 'font-size'], golden: 'the number from data-line, drawn, so a selection or a copy leaves it out; 2em at least (DocC\'s); the secondary label, fitted to 4.5:1 (DocC\'s rgb(110, 110, 115)); the listing\'s 13px' },
      { target: H, props: ['background-color'], golden: 'highlighted, the brand\'s fill-quaternary (DocC\'s rgb(234, 243, 255))' },
      { target: H, pseudo: '::before', props: ['border-left-color'], golden: 'its edge the brand label\'s blue, the accent\'s in light (DocC\'s rgb(76, 169, 255))' },
      { target: L, props: ['padding-left', 'padding-right'], golden: 'numbered lines: no room at its sides, the numbers give it (DocC\'s pre)' },
      { target: N, props: ['box-x'], golden: 'a line from the listing\'s edge, a highlight across it' },
    ],
  },
  {
    name: 'listing nested',
    html: ex('nested'),
    widths: [1200],
    checks: [
      { target: L, props: ['background-color'], golden: 'in a console card, whose gray is its own, the page\'s white' },
    ],
  },
  {
    name: 'listing lines dark',
    html: `<div class="cui-dark">${ex('lines')}</div>`,
    widths: [1200],
    checks: [
      { target: N, pseudo: '::before', props: ['color'], golden: 'the dark side derived (tokens.css, decision 18)' },
      { target: H, props: ['background-color'], golden: 'highlighted in the dark (DocC\'s rgb(66, 66, 69))' },
      { target: H, pseudo: '::before', props: ['border-left-color'], golden: 'its edge in the dark a step lighter than the accent, the brand label\'s, 5.31:1 on the dark highlight: the accent\'s dark side was 2.57:1, under a mark\'s 3:1 (decision 21; DocC\'s rgb(125, 193, 255))' },
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
