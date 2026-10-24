// The console's card against App Store Connect's analytics overview: its
// ground, corners, room above and clip (analytics-page:card), its title,
// value and change's type (analytics:card.*); the ground the toolkit's
// v2-bgPlaceholder. Golden, with their why: its room at its sides and
// foot, the change's colors, the caption, the figure's row, the dark side.

import { examples } from '../../test/harness.mjs';

const ex = examples(import.meta.url);
const CARD = '.cui-console-card';
const TYPE = ['font-size', 'line-height', 'font-weight', 'letter-spacing'];

export const cases = [
  {
    name: 'metrics',
    html: ex('metrics'),
    widths: [1600, 1200],
    checks: [
      { target: CARD, props: ['padding-top', 'border-top-left-radius', 'background-color', 'overflow', 'box-shadow', 'border-top-width'], asc: ['analytics-page', 'card'] },
      { target: CARD, props: ['background-color'], tb: ['color-v2-bgPlaceholder'] },
      { target: '.cui-console-card-title', props: [...TYPE, 'color'], asc: ['analytics', 'card.title'] },
      { target: '.cui-console-card-value', props: [...TYPE, 'color', 'font-variant-numeric'], asc: ['analytics', 'card.value'] },
      { target: '.cui-console-card-change[data-color="error"]', props: TYPE, asc: ['analytics', 'card.change'] },
      { target: '.cui-console-card-change[data-color="error"]', props: ['color'], golden: 'down, the red label red-700: the toolkit\'s red is 3.6:1 on white (tokens.css, decision 21)' },
      { target: '.cui-console-card-change[data-color="success"]', props: ['color'], golden: 'up, the green label green-800, fitted (tokens.css, decision 21)' },
      { target: '.cui-console-card-change:not([data-color])', props: ['color'], golden: 'no change, the secondary label' },
      { target: CARD, props: ['padding-left', 'padding-right', 'padding-bottom'], golden: 'its text 15px in at its sides and foot, as at its top (the analytics card\'s chart runs to its edges)' },
      { target: '.cui-console-card-caption', props: ['font-size', 'line-height', 'font-weight', 'color'], golden: 'a caption, 12px on 16px, the secondary label: the toolkit\'s bodySmall' },
      { target: '.cui-console-card-figure', props: ['display', 'align-items', 'column-gap'], golden: 'a value and its change on the baseline, 6px apart' },
    ],
  },
  {
    name: 'card list',
    html: ex('list'),
    widths: [1200],
    checks: [
      { target: '.cui-console-card-header', props: ['display', 'justify-content', 'align-items'], golden: 'a header: its title at its start, its link at its end, on the baseline' },
    ],
  },
  {
    name: 'card dark',
    html: `<div class="cui-dark">${ex('metrics')}</div>`,
    widths: [1200],
    checks: [
      { target: CARD, props: ['background-color'], golden: 'the dark side derived: the toolkit\'s gray coolGray-50\'s dark step (tokens.css, decision 18)' },
      { target: '.cui-console-card-change[data-color="error"]', props: ['color'], golden: 'the dark side derived (tokens.css, decision 18)' },
    ],
  },
];
