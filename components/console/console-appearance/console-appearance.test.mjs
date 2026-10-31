// The console's appearance control against Apple's developer
// documentation's (DocC, reference/apple.com/results/docc-scheme.json,
// -dark): its shape, its room, its type's size. Golden, with their why: the
// console's blues, the font tracking itself, the focus ring, the touch
// target.

import { examples } from '../../../test/harness.mjs';

const ex = examples(import.meta.url);
const G = '.cui-console-appearance';
const CHOICE = `${G} .cui-console-appearance-choice:not(:has(input:checked))`;
const CHOSEN = `${G} .cui-console-appearance-choice:has(input:checked)`;
const ROOM = ['padding-top', 'padding-right', 'padding-bottom', 'padding-left'];
const EDGE = ['border-top-width', 'border-top-style', 'border-bottom-width', 'border-bottom-style'];

export const cases = [
  {
    name: 'appearance',
    html: ex('appearance'),
    widths: [1200, 390],
    checks: [
      { target: G, props: ['display', 'border-top-left-radius', ...ROOM, ...EDGE, 'font-size', 'line-height', 'font-weight'], apple: ['docc-scheme', 'scheme'] },
      { target: G, props: ['box-height'], golden: 'DocC\'s 24px: a 20px choice, 1px in, a 1px edge' },
      { target: `${CHOICE} > span`, props: ['display', 'min-width', 'text-align', 'border-top-left-radius', ...ROOM, ...EDGE], apple: ['docc-scheme', 'scheme.choice'] },
      { target: `${CHOICE} > span`, props: ['cursor'], state: 'hover', apple: ['docc-scheme', 'scheme.choice', 'hover'] },
      { target: `${CHOSEN} > span`, props: ['border-top-left-radius', ...ROOM, ...EDGE], apple: ['docc-scheme', 'scheme.chosen'] },
      { target: G, props: ['border-top-color', 'letter-spacing'], golden: 'the console\'s brand accent-primary, DocC\'s rgb(0, 113, 227) the nearest; the font tracking itself (DocC\'s -0.12px is SF\'s at 12px)' },
      { target: `${CHOICE} > span`, props: ['color', 'background-color'], golden: 'an unchosen word the brand\'s label, fitted to 4.5:1 (DocC\'s rgb(0, 113, 227))' },
      { target: `${CHOSEN} > span`, props: ['color', 'background-color', 'border-top-color'], golden: 'chosen, the label on the accent on the brand\'s accent (DocC\'s white on rgb(0, 113, 227))' },
      { target: `${G} input`, props: ['opacity', 'box-width'], golden: 'the radio read, not shown' },
      { target: CHOSEN, props: ['outline-style'], golden: 'no ring at rest' },
      { target: `${G} > legend`, props: ['box-width', 'position'], golden: 'the legend read, not shown' },
    ],
  },
  {
    name: 'appearance focused',
    html: ex('appearance'),
    widths: [1200],
    press: ['Tab'],
    checks: [
      { target: CHOSEN, props: ['outline-style', 'outline-width', 'outline-color', 'outline-offset'], golden: 'the radio hidden, the one focus ring around its choice' },
    ],
  },
  {
    name: 'appearance touch',
    html: ex('appearance'),
    touch: true,
    widths: [390],
    checks: [
      { target: CHOICE, pseudo: '::after', props: ['height'], golden: 'a touch screen\'s target 44px tall, the HIG\'s 44pt' },
    ],
  },
  {
    name: 'appearance dark',
    html: `<div class="cui-dark">${ex('appearance')}</div>`,
    widths: [1200],
    checks: [
      { target: `${CHOICE} > span`, props: ['color'], golden: 'the dark side derived (tokens.css, decision 18); DocC\'s rgb(41, 151, 255)' },
      { target: `${CHOSEN} > span`, props: ['color', 'background-color'], golden: 'chosen in the dark' },
    ],
  },
];
