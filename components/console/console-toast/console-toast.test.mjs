// The console's toast against App Store Connect's toolkit's stroked
// popover: its ground, shadow and corners (desktop and phone), a popover
// footer's room, the small page margin. Golden, with their why: its place,
// unseen until shown, the fade, the dark side.

import { examples } from '../../../test/harness.mjs';

const ex = examples(import.meta.url);
const P = 'component-Popover';
const T = '.cui-console-toast';
const F = `${P}-anatomy-Content-anatomy-Footer`;

export const cases = [
  {
    name: 'toast',
    html: ex('toast'),
    widths: [1200, 390],
    checks: [
      { target: T, props: ['background-color', 'box-shadow'], tb: [`${P}-variant-variant-stroked-backgroundColor`, `${P}-variant-variant-stroked-boxShadow`] },
      { target: T, props: ['padding-top', 'padding-bottom', 'padding-left', 'padding-right'], tb: [`${F}-paddingBlock`, `${F}-paddingBlock`, `${F}-paddingInline`, `${F}-paddingInline`] },
      { target: T, props: ['font-size', 'line-height', 'font-weight'], tb: ['typography-textStyle-bodyRegular-fontSize', 'typography-textStyle-bodyRegular-lineHeight', 'typography-textStyle-bodyRegular-fontWeight'] },
      { target: T, props: ['bottom', 'right'], tb: ['space-pageMarginSm', 'space-pageMarginSm'] },
      { target: T, props: ['position', 'opacity', 'color', 'z-index'], golden: 'shown: fixed at the bottom right corner, the label, under the skip link' },
    ],
  },
  {
    name: 'toast desktop corners',
    html: ex('toast'),
    widths: [1200],
    checks: [
      { target: T, props: ['border-top-left-radius'], tb: [`${P}-variant-pageScale-desktop-borderRadius`] },
    ],
  },
  {
    name: 'toast phone corners',
    html: ex('toast'),
    widths: [390],
    checks: [
      { target: T, props: ['border-top-left-radius'], tb: [`${P}-variant-pageScale-mobile-borderRadius`] },
    ],
  },
  {
    name: 'toast hidden',
    html: ex('toast').replace(' data-shown', ''),
    widths: [1200],
    checks: [
      { target: T, props: ['opacity', 'pointer-events', 'transition-duration'], golden: 'unseen and untouchable until data-shown; it fades in over 200ms' },
    ],
  },
  {
    name: 'toast reduced motion',
    html: ex('toast'),
    widths: [1200],
    reducedMotion: 'reduce',
    checks: [
      { target: T, props: ['transition-duration'], golden: 'no fade with reduced motion' },
    ],
  },
  {
    name: 'toast dark',
    html: `<div class="cui-dark">${ex('toast')}</div>`,
    widths: [1200],
    checks: [
      { target: T, props: ['background-color', 'color'], golden: 'the dark side derived (tokens.css, decision 18)' },
    ],
  },
];
