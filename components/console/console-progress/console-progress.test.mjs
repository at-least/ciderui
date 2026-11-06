import { examples } from '../../../test/harness.mjs';

const ex = examples(import.meta.url);
const P = 'component-Progress-variant';
const MD = '.cui-console-progress:not([data-size]):not([data-color]):not(:indeterminate)';
const WAITING = '.cui-console-progress:indeterminate';

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
      // the fill read as color, which ::-webkit-progress-value draws: getComputedStyle on that pseudo-element returns the element's own
      { target: MD, props: ['color'], tb: [`${P}-colorRole-brand-anatomy-Fill-backgroundColor`] },
      { target: '[data-color="neutral"]', props: ['color'], tb: [`${P}-colorRole-neutral-anatomy-Fill-backgroundColor`] },
      { target: MD, props: ['width'], golden: 'the column\'s width' },
    ],
  },
  {
    name: 'progress indeterminate',
    html: ex('progress'),
    widths: [1200],
    checks: [
      { target: WAITING, props: ['animation-name', 'animation-duration', 'animation-iteration-count', 'animation-timing-function', 'background-size', 'background-repeat'], golden: 'no value: a small fill loops along the track, the HIG\'s indeterminate progress bar, its one motion; the toolkit has none, and a still track read as nothing happening' },
      { target: WAITING, props: ['color'], golden: 'the fill the brand label\'s blue, as a value\'s' },
    ],
  },
  {
    name: 'progress indeterminate reduced motion',
    html: ex('progress'),
    reducedMotion: 'reduce',
    widths: [1200],
    checks: [
      { target: WAITING, props: ['animation-iteration-count', 'background-position-x'], golden: 'reduced motion stills it (base.css), the fill in the track\'s middle' },
    ],
  },
  {
    name: 'progress indeterminate forced colors',
    html: ex('progress'),
    forcedColors: true,
    widths: [1200],
    checks: [
      { target: WAITING, props: ['background-color', 'outline-style', 'outline-width'], golden: 'forced colors draw neither its track nor its fill: an outline holds its place' },
      { target: WAITING, props: ['background-image'], golden: 'forced colors: the fill the theme\'s Highlight on its Canvas: they drop a gradient, and the bar read as still and empty' },
    ],
  },
  {
    name: 'progress dark',
    html: `<div class="cui-dark">${ex('progress')}</div>`,
    widths: [1200],
    checks: [
      { target: MD, props: ['background-color'], golden: 'the dark side derived (tokens.css, decision 18)' },
      { target: '[data-color="neutral"]', props: ['color'], golden: 'the neutral done part in the dark the label\'s white, derived (tokens.css, decision 18)' },
      { target: MD, props: ['color'], golden: 'the done part in the dark the brand label\'s blue, a step lighter than the accent, 5.66:1 on the dark track: the accent\'s dark side was 2.74:1, under a mark\'s 3:1 (decision 21)' },
    ],
  },
];
