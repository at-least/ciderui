import { examples } from '../../../test/harness.mjs';

const ex = examples(import.meta.url);
const ROW = '.cui-meter-row:first-child';
const TRACK = `${ROW} .cui-meter-track`;
const FILL = `${ROW} .cui-meter-fill`;
const dark = (html) => `<div class="cui-dark">${html}</div>`;

export const cases = [
  {
    name: 'meter',
    html: ex('meter'),
    widths: [1200, 390],
    checks: [
      { target: TRACK, props: ['height', 'background-color', 'border-bottom-left-radius'], apple: ['appstore-ratings', 'track'] },
      { target: FILL, props: ['height', 'background-color', 'border-top-left-radius', 'border-top-right-radius'], apple: ['appstore-ratings', 'fill'] },
      { target: '.cui-meter', props: ['grid-template-columns', 'align-items', 'row-gap'], golden: 'the labels, the tracks and the values in three columns, each track centred on its row\'s text: the App Store\'s rows are 8px, its labels stars' },
      { target: `${ROW} .cui-meter-label`, props: ['font-size', 'line-height', 'color', 'padding-right'], golden: 'a label and a value in the caption\'s type, the secondary text, 8px from the track' },
      { target: `${ROW} .cui-meter-value`, props: ['padding-left', 'text-align', 'font-variant-numeric'], golden: 'the value 8px after the track, its figures tabular at the row\'s end' },
    ],
  },
  {
    name: 'meter dark',
    html: dark(ex('meter')),
    widths: [1200],
    checks: [
      { target: TRACK, props: ['background-color'], apple: ['appstore-ratings-dark', 'track'] },
      { target: FILL, props: ['background-color'], apple: ['appstore-ratings-dark', 'fill'] },
    ],
  },
  {
    name: 'meter forced colors',
    html: ex('meter'),
    widths: [1200],
    forcedColors: true,
    checks: [
      { target: TRACK, props: ['background-color'], golden: 'the track in GrayText and the filled part in CanvasText, so the share still shows with the page\'s colors forced' },
      { target: FILL, props: ['background-color'], golden: 'the filled part in CanvasText' },
    ],
  },
];
