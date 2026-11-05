import { examples, SIZE, TRACKING } from '../../../test/harness.mjs';

const ex = examples(import.meta.url);

export const cases = [
  {
    name: 'stat',
    html: `<div class="cui-dark">${ex('stat')}</div>`,
    checks: [
      { target: '.cui-stat', props: ['padding-top', 'padding-bottom', 'margin-top', 'margin-bottom'], apple: ['product', 'stat.content'] },
      { target: '.cui-stat-caption', props: [...SIZE, 'display'], apple: ['product', 'stat.caption'] },
      { target: '.cui-stat-value', props: [...SIZE, 'color', 'margin-top', 'margin-bottom', 'display'], apple: ['product', 'stat.value'] },
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
      { target: '.cui-stat-caption', props: [...SIZE, 'color', 'display'], apple: ['airpods', 'stat.caption'] },
      { target: '.cui-stat-value', props: ['color', 'margin-top', 'margin-bottom', 'display'], apple: ['airpods', 'stat.value'] },
      { target: '.cui-stat-value', props: SIZE, golden: '48px at every width, where apple.com\'s steps to 64 and 80px (airpods:stat.value): a number that grows past its page\'s headline competes with it (DESIGN.md decision 25)' },
      { target: '.cui-stat', props: ['margin-top', 'margin-bottom'], golden: 'no margin of its own: apple.com\'s 20px and 40px around its large stat are the page\'s (airpods:stat.content)' },
      { target: '.cui-stat-caption', props: ['letter-spacing'], golden: TRACKING },
      { target: '.cui-stat-value', props: ['letter-spacing'], golden: TRACKING },
    ],
  },
];
