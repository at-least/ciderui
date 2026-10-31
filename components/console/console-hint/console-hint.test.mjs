// The console's hint against App Store Connect's toolkit's info button and
// stroked popover (--tb-component-InfoSystemButton-*, --tb-component-
// Popover-*) and its analytics overview's card info button: the button's
// height, rounding, fills, labels and halo; the popover's ground, shadow,
// corners, narrowest and widest. Golden, with their why: the circle, the
// letter, the popover's room, type and place, the dark side.

import { examples } from '../../../test/harness.mjs';

const ex = examples(import.meta.url);
const I = 'component-InfoSystemButton';
const P = 'component-Popover';
const B = '.cui-console-hint';
const POP = '.cui-console-popover';
// the example set where the popover has room below it, and room either side
const page = (html) => `<div style="padding:24px 160px 240px">${html}</div>`;

export const cases = [
  {
    name: 'hint',
    html: page(ex('hint')),
    widths: [1200, 390],
    checks: [
      { target: B, props: ['height', 'border-top-left-radius', 'background-color'], asc: ['analytics', 'card.info'] },
      { target: B, pseudo: '::before', props: ['background-color'], tb: [`${I}-variant-effect-default-backgroundColor`] },
      { target: B, props: ['color'], golden: 'the "i" the secondary label at 0.74: the toolkit\'s 0.60 is 3.4:1 on white (tokens.css, decision 21)' },
      { target: B, props: ['width', 'font-size', 'font-weight', 'margin-left'], golden: 'the "i" a 12px letter at 600 in a 20px button, where App Store Connect draws an SF Symbol; 4px after the word' },
      { target: B, pseudo: '::before', props: ['width', 'height', 'border-top-left-radius'], golden: 'its circle 17px, as App Store Connect\'s symbol draws it' },
      { target: B, pseudo: '::before', props: ['background-color'], state: 'hover', tb: [`${I}-variant-effect-hover-backgroundColor`] },
      { target: B, props: ['color'], state: 'hover', tb: [`${I}-variant-effect-hover-color`] },
      { target: B, pseudo: '::before', props: ['background-color'], state: 'active', tb: [`${I}-variant-effect-press-backgroundColor`] },
      { target: B, props: ['color'], state: 'active', tb: [`${I}-variant-effect-press-color`] },
      { target: B, props: ['outline-width', 'outline-style', 'outline-color', 'outline-offset'], state: 'focus-visible', golden: 'focused from the keyboard, base.css\'s ring (rule 9): the toolkit\'s halo, 3px of black at 0.15 1px out, is 1.41:1 on white and 1.39:1 dark, a focus that cannot be seen' },
      { target: POP, props: ['display'], golden: 'closed until its button opens it' },
      { target: B, pseudo: '::after', props: ['width', 'height'], golden: 'its target 20px square, the HIG\'s macOS minimum (Accessibility, Mobility), round the 20 by 18px button' },
    ],
  },
  {
    name: 'hint open',
    html: page(ex('hint')),
    click: B,
    widths: [1200],
    checks: [
      { target: POP, props: ['background-color', 'box-shadow', 'border-top-left-radius', 'min-width'], tb: [`${P}-variant-variant-stroked-backgroundColor`, `${P}-variant-variant-stroked-boxShadow`, `${P}-variant-pageScale-desktop-borderRadius`, `${P}-minWidth`] },
      { target: POP, props: ['padding-top', 'padding-left', 'font-size', 'line-height', 'font-weight', 'color', 'max-width'], golden: 'App Store Connect\'s analytics overview\'s popover: 16px in, 13px on 18px, the label, 350px at its widest' },
      { target: POP, props: ['display', 'box-width', 'box-top', 'box-left'], golden: 'open: 350px wide, 8px under the button, centered on it' },
    ],
  },
  {
    name: 'hint open phone',
    html: page(ex('hint')),
    click: B,
    widths: [390],
    checks: [
      { target: POP, props: ['border-top-left-radius'], tb: [`${P}-variant-pageScale-mobile-borderRadius`] },
      { target: POP, props: ['box-width', 'box-left'], golden: 'on a phone, at most 85% of the window, 16px inside its edge' },
    ],
  },
  {
    name: 'hint touch',
    html: page(ex('hint')),
    touch: true,
    widths: [390],
    checks: [
      { target: B, props: ['color'], state: 'hover', rest: 'a tapped hint keeps no hover (rule 5)' },
      { target: B, pseudo: '::after', props: ['width', 'height'], golden: 'its target on a touch screen 32 by 44px: the HIG\'s 44pt height, over its 28pt minimum across, no wider so it leaves the control beside the word' },
    ],
  },
  {
    name: 'hint dark',
    html: `<div class="cui-dark">${page(ex('hint'))}</div>`,
    widths: [1200],
    checks: [
      { target: B, props: ['color'], golden: 'the dark side derived (tokens.css, decision 18)' },
      { target: B, pseudo: '::before', props: ['background-color'], golden: 'the dark side derived (tokens.css, decision 18)' },
      { target: POP, props: ['background-color', 'color'], golden: 'the dark side derived (tokens.css, decision 18)' },
    ],
  },
];
