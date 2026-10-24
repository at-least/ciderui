// The console's status against App Store Connect's: TestFlight's build
// status (its mark, its word, the line under it). Golden, with their why:
// the mark's colors (the console's fitted labels), its circle, the line's
// indent, a key's own color, the dark side.

import { examples } from '../../test/harness.mjs';

const ex = examples(import.meta.url);
const COLORS = ['neutral', 'brand', 'info', 'success', 'warning', 'error'];
const at = (c) => `.cui-console-status${c === 'neutral' ? ':not([data-color])' : `[data-color="${c}"]`}`;
const MARK = (c) => `${at(c)} .cui-console-status-mark`;

export const cases = [
  {
    name: 'statuses',
    html: ex('status'),
    widths: [1200, 390],
    checks: [
      { target: MARK('success'), props: ['height', 'margin-right'], asc: ['testflight', 'status.icon'] },
      { target: at('success'), props: ['font-size', 'line-height'], asc: ['testflight', 'status.word'] },
      { target: at('success'), props: ['color'], golden: 'the word in the console\'s label: App Store Connect\'s older look wrote rgb(51, 51, 51)' },
      { target: MARK('success'), props: ['box-width'], golden: 'as wide as App Store Connect\'s 12px (an asc check reads no width: its text was SF)' },
      ...COLORS.map((c) => ({ target: MARK(c), props: ['background-color', 'border-top-left-radius'], golden: `${c}: the mark in the console's ${c} label (tokens.css, fitted to 4.5:1, decision 21), a circle (App Store Connect's glyph's shape is unread)` })),
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
    name: 'status key',
    html: ex('key'),
    widths: [1200],
    checks: [
      { target: '.cui-console-status-mark', props: ['background-color'], golden: 'a series\' own color through --cui-console-status-color (a chart\'s key)' },
    ],
  },
  {
    name: 'statuses forced colors',
    html: ex('status'),
    widths: [1200],
    forcedColors: true,
    checks: [
      { target: MARK('error'), props: ['background-color'], golden: 'CanvasText: forced colors would drop the fill and the mark with it' },
    ],
  },
  {
    name: 'statuses dark',
    html: `<div class="cui-dark">${ex('status')}</div>`,
    widths: [1200],
    checks: COLORS.map((c) => ({ target: MARK(c), props: ['background-color'], golden: 'the dark side derived (tokens.css, decision 18)' })),
  },
];
