// The console's description list against App Store Connect's toolkit's
// (--tb-component-DL-*): its gaps, its labels' and values' type and color,
// a vertical value's room, the strong label, each page scale. Golden, with
// their why: the label column, the muted label's fitted alpha, a value
// breaking anywhere, the dark side.

import { examples } from '../../test/harness.mjs';

const ex = examples(import.meta.url);
const D = 'component-DL';
const TYPE = ['font-size', 'line-height', 'font-weight'];
const style = (part, scale, variant) => ['fontSize', 'lineHeight', 'fontWeight'].map((p) => `${D}-anatomy-${part}-variant-pageScale-${scale}${variant ? `-variant-${variant}` : ''}-textStyle-${p}`);

export const cases = [
  {
    name: 'dl',
    html: ex('dl'),
    widths: [1200, 640],
    checks: [
      { target: '.cui-console-dl', props: ['column-gap', 'row-gap'], tb: [`${D}-variant-orientation-horizontal-pageScale-desktop-columnGap`, `${D}-variant-orientation-horizontal-pageScale-desktop-rowGap`] },
      { target: 'dt:first-of-type', props: TYPE, tb: style('DT', 'desktop', 'muted') },
      { target: 'dd:first-of-type', props: [...TYPE, 'color'], tb: [...style('DD', 'desktop'), `${D}-anatomy-DD-color`] },
      { target: 'dt:first-of-type', props: ['color'], golden: 'the muted label, the secondary label at 0.74: the toolkit\'s 0.60 is 3.4:1 on white (tokens.css, decision 21)' },
      { target: '.cui-console-dl', props: ['grid-template-columns'], golden: 'a label column as wide as its longest label, the values in the rest' },
      { target: 'dd:last-of-type', props: ['overflow-wrap'], golden: 'a value breaks anywhere it must: an id, a URL' },
    ],
  },
  {
    name: 'dl phone',
    html: ex('dl'),
    widths: [639, 390],
    checks: [
      { target: '.cui-console-dl', props: ['column-gap', 'row-gap'], tb: [`${D}-variant-orientation-horizontal-pageScale-mobile-columnGap`, `${D}-variant-orientation-horizontal-pageScale-mobile-rowGap`] },
      { target: 'dt:first-of-type', props: TYPE, tb: style('DT', 'mobile', 'muted') },
      { target: 'dd:first-of-type', props: TYPE, tb: style('DD', 'mobile') },
    ],
  },
  {
    name: 'dl vertical strong',
    html: ex('vertical'),
    widths: [1200],
    checks: [
      { target: 'dd:first-of-type', props: ['margin-bottom'], tb: [`${D}-anatomy-DD-variant-orientation-vertical-marginBottom`] },
      { target: 'dt:first-of-type', props: [...TYPE, 'color'], tb: [...style('DT', 'desktop', 'strong'), `${D}-anatomy-DT-variant-variant-strong-color`] },
      { target: '.cui-console-dl', props: ['display'], golden: 'vertical, each label over its value' },
    ],
  },
  {
    name: 'dl dark',
    html: `<div class="cui-dark">${ex('dl')}</div>`,
    widths: [1200],
    checks: [
      { target: 'dt:first-of-type', props: ['color'], golden: 'the dark side derived (tokens.css, decision 18)' },
      { target: 'dd:first-of-type', props: ['color'], golden: 'the dark side derived (tokens.css, decision 18)' },
    ],
  },
];
