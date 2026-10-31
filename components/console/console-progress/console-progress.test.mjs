// The console's progress bar against App Store Connect's toolkit's
// (--tb-component-Progress-*): its heights, its rounding, its track's and
// fills' colors (the fill read as the bar's color, which its value
// draws: a pseudo-element's computed style is not read). Golden, with
// their why: its width, the dark side.

import { examples } from '../../../test/harness.mjs';

const ex = examples(import.meta.url);
const P = 'component-Progress-variant';
const MD = '.cui-console-progress:not([data-size]):not([data-color])';

export const cases = [
  {
    name: 'progress',
    html: ex('progress'),
    widths: [1200, 390],
    checks: [
      { target: MD, props: ['height'], tb: [`${P}-controlSize-md-anatomy-Track-height`] },
      { target: '[data-size="sm"]', props: ['height'], tb: [`${P}-controlSize-sm-anatomy-Track-height`] },
      { target: '[data-size="lg"]', props: ['height'], tb: [`${P}-controlSize-lg-anatomy-Track-height`] },
      { target: MD, props: ['border-top-left-radius'], tb: [`${P}-displayStyle-horizontalRounded-anatomy-Track-borderRadius`] },
      { target: MD, props: ['background-color'], tb: [`${P}-colorRole-brand-anatomy-Track-backgroundColor`] },
      { target: MD, props: ['color'], tb: [`${P}-colorRole-brand-anatomy-Fill-backgroundColor`] },
      { target: '[data-color="neutral"]', props: ['color'], tb: [`${P}-colorRole-neutral-anatomy-Fill-backgroundColor`] },
      { target: MD, props: ['width'], golden: 'the column\'s width' },
    ],
  },
  {
    name: 'progress dark',
    html: `<div class="cui-dark">${ex('progress')}</div>`,
    widths: [1200],
    checks: [
      { target: MD, props: ['background-color', 'color'], golden: 'the dark side derived (tokens.css, decision 18)' },
    ],
  },
];
