// The console's chart against App Store Connect's analytics charts
// (reference/asc/results/analytics.json: chart.grid, chart.x, chart.y,
// spark.line). Golden, with their why: the fitted colors, the scale's
// place, the separators, the empty note, the legend, the dark side.

import { examples } from '../../test/harness.mjs';

const ex = examples(import.meta.url);
const C = '.cui-console-chart';

export const cases = [
  {
    name: 'chart',
    html: ex('chart'),
    widths: [1200, 390],
    checks: [
      { target: `${C}-grid`, props: ['stroke-width'], asc: ['analytics', 'chart.grid'] },
      { target: `${C}-scale`, props: ['font-size', 'font-weight'], asc: ['analytics', 'chart.y'] },
      { target: `${C}-axis`, props: ['font-size', 'font-weight'], asc: ['analytics', 'chart.x'] },
      { target: `${C}-grid`, props: ['stroke'], golden: 'the tertiary hairline, where App Store Connect\'s is rgb(214, 214, 214)' },
      { target: `${C}-scale`, props: ['color', 'background-color', 'right', 'top'], golden: 'the fitted secondary label (App Store Connect\'s rgb(136, 136, 136) is 3.5:1 on white, decision 21) on the ground, at the plot\'s trailing top (the HIG keeps the leading edge clean)' },
      { target: `${C}-scale[data-at="half"]`, props: ['box-top'], golden: 'the half scale on the middle gridline' },
      { target: `${C}-axis`, props: ['color', 'margin-top', 'justify-content'], golden: 'the first and last labels at the plot\'s ends, 6px under it, the fitted secondary label' },
      { target: `${C}-mark`, props: ['stroke', 'stroke-width'], golden: 'a 1px line of the ground round each mark: stacked marks stand apart (the HIG\'s Color)' },
      { target: `${C}-legend`, props: ['margin-top', 'row-gap', 'column-gap'], golden: 'the series under the axis, 8px, 4px by 12px apart' },
      { target: `${C}-svg`, props: ['box-height', 'display'], golden: 'its height the page\'s, the svg\'s 140' },
    ],
  },
  {
    name: 'chart hover',
    html: ex('chart'),
    widths: [1200],
    checks: [
      { target: `${C}-bar:last-of-type ${C}-mark`, props: ['opacity'], state: 'hover', on: `${C}-bar:last-of-type`, golden: 'its bar hovered, its marks at 0.85' },
    ],
  },
  {
    name: 'chart in a card',
    html: `<div class="cui-console-page"><div class="cui-console-card">${ex('chart')}</div></div>`,
    widths: [1200],
    checks: [
      { target: `${C}-mark`, props: ['stroke'], golden: 'in a card, the line round a mark is the card\'s ground' },
      { target: `${C}-scale`, props: ['background-color'], golden: 'in a card, the scale on the card\'s ground' },
    ],
  },
  {
    name: 'chart empty',
    html: ex('empty'),
    widths: [1200],
    checks: [
      { target: `${C}-empty`, props: ['font-size', 'color', 'box-top', 'box-height'], golden: 'centered on the plot, 13px, the secondary label' },
    ],
  },
  {
    name: 'sparkline',
    html: ex('sparkline'),
    widths: [1200],
    checks: [
      { target: `${C}-sparkline-line`, props: ['stroke', 'stroke-width', 'stroke-linecap', 'stroke-linejoin'], asc: ['analytics', 'spark.line'] },
      { target: `${C}-sparkline-area`, props: ['fill'], golden: 'the blue\'s tint, flat where App Store Connect\'s is a gradient of it' },
    ],
  },
  {
    name: 'chart dark',
    html: `<div class="cui-dark">${ex('chart')}${ex('sparkline')}</div>`,
    widths: [1200],
    checks: [
      { target: `${C}-grid`, props: ['stroke'], golden: 'the dark side derived (tokens.css, decision 18)' },
      { target: `${C}-scale`, props: ['color', 'background-color'], golden: 'the dark side derived (tokens.css, decision 18)' },
      { target: `${C}-sparkline-line`, props: ['stroke'], golden: 'the dark side derived (tokens.css, decision 18)' },
    ],
  },
];
