import { examples } from '../../../test/harness.mjs';

const ex = examples(import.meta.url);
const CARD = '.cui-console-card';
const TYPE = ['font-size', 'line-height', 'font-weight', 'letter-spacing'];

export const cases = [
  {
    name: 'card',
    html: ex('list'),
    widths: [1600, 1200],
    checks: [
      { target: CARD, props: ['padding-top', 'border-top-left-radius', 'background-color', 'overflow', 'box-shadow', 'border-top-width'], asc: ['analytics-page', 'card'] },
      { target: '.cui-console-card-title', props: [...TYPE, 'color'], asc: ['analytics', 'card.title'] },
      { target: CARD, props: ['padding-left', 'padding-right', 'padding-bottom'], golden: 'its text 15px in at its sides and foot, as at its top (the analytics card\'s chart runs to its edges)' },
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
    html: `<div class="cui-dark">${ex('list')}</div>`,
    widths: [1200],
    checks: [
      { target: CARD, props: ['background-color'], golden: 'the dark side derived: the toolkit\'s gray coolGray-50\'s dark step (tokens.css, decision 18)' },
    ],
  },
];
