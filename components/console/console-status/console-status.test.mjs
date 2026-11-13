import { examples } from '../../../test/harness.mjs';

const ex = examples(import.meta.url);
const STATUS = '.cui-console-status';
const MARK = '.cui-console-status-mark';

export const cases = [
  {
    name: 'statuses',
    html: ex('status'),
    widths: [1200, 390],
    checks: [
      { target: MARK, props: ['height', 'margin-right'], asc: ['testflight', 'status.icon'] },
      { target: STATUS, props: ['font-size', 'line-height'], asc: ['testflight', 'status.word'] },
      { target: STATUS, props: ['display', 'align-items'], golden: 'the mark and the word on one line, the mark centered on it' },
      { target: STATUS, props: ['color'], golden: 'the word in the console\'s label: App Store Connect\'s older look wrote rgb(51, 51, 51)' },
      { target: MARK, props: ['box-width'], golden: 'as wide as App Store Connect\'s 12px (an asc check reads no width: its text was SF)' },
      { target: MARK, props: ['background-color', 'border-top-left-radius'], golden: 'the mark in the console\'s warning label, the one color a status takes: it says look here (tokens.css, fitted to 4.5:1, decision 21), a circle (App Store Connect\'s glyph\'s shape is unread)' },
    ],
  },
  {
    name: 'status detail',
    html: ex('detail'),
    widths: [1200],
    checks: [
      { target: '.cui-console-status-detail', props: ['font-size', 'line-height', 'margin-top'], asc: ['testflight', 'status.line'] },
      { target: '.cui-console-status-detail', props: ['color', 'padding-left'], golden: 'the secondary label, 18px in: under the word, past the mark and its 6px' },
    ],
  },
  {
    name: 'statuses forced colors',
    html: ex('status'),
    widths: [1200],
    forcedColors: true,
    checks: [
      { target: MARK, props: ['background-color'], golden: 'CanvasText: forced colors would drop the fill and the mark with it' },
    ],
  },
  {
    name: 'statuses dark',
    html: `<div class="cui-dark">${ex('status')}</div>`,
    widths: [1200],
    checks: [{ target: MARK, props: ['background-color'], golden: 'the dark side derived (tokens.css, decision 18)' }],
  },
];
