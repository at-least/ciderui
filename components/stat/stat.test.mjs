// The stat against apple.com's: iPhone 18 Pro's (product:stat.*: its
// captions' and value's type, the gaps; on its dark page the value's color)
// and AirPods Pro's large one (airpods:stat.*, on a light page: its three
// tiers, its colors). Golden, with their why: the letter-spacing, the dark
// captions' fitted color.

import { examples } from '../../test/harness.mjs';

const ex = examples(import.meta.url);
const TYPE = ['font-size', 'line-height', 'font-weight'];
const TRACKING = 'the system font\'s fitted letter-spacing (tokens.css), not apple.com\'s';

export const cases = [
  {
    name: 'stat',
    html: `<div class="cui-dark">${ex('stat')}</div>`,
    checks: [
      { target: '.cui-stat', props: ['padding-top', 'padding-bottom', 'margin-top', 'margin-bottom'], apple: ['product', 'stat.content'] },
      { target: '.cui-stat-caption', props: [...TYPE, 'display'], apple: ['product', 'stat.caption'] },
      { target: '.cui-stat-value', props: [...TYPE, 'color', 'margin-top', 'margin-bottom', 'display'], apple: ['product', 'stat.value'] },
      { target: '.cui-stat-caption', props: ['color'], golden: 'the dark secondary text, #86868b, 5.8:1 on black: apple.com\'s dark page\'s #6e6e73 is 4.1:1 (decision 6)' },
      { target: '.cui-stat-caption', props: ['letter-spacing'], golden: TRACKING },
      { target: '.cui-stat-value', props: ['letter-spacing'], golden: TRACKING },
    ],
  },
  {
    name: 'large',
    html: ex('large'),
    checks: [
      { target: '.cui-stat', props: ['padding-top', 'padding-bottom'], apple: ['airpods', 'stat.content'] },
      { target: '.cui-stat-caption', props: [...TYPE, 'color', 'display'], apple: ['airpods', 'stat.caption'] },
      { target: '.cui-stat-value', props: [...TYPE, 'color', 'margin-top', 'margin-bottom', 'display'], apple: ['airpods', 'stat.value'] },
      { target: '.cui-stat', props: ['margin-top', 'margin-bottom'], golden: 'no margin of its own: apple.com\'s 20px and 40px around its large stat are the page\'s (airpods:stat.content)' },
      { target: '.cui-stat-caption', props: ['letter-spacing'], golden: TRACKING },
      { target: '.cui-stat-value', props: ['letter-spacing'], golden: TRACKING },
    ],
  },
];
