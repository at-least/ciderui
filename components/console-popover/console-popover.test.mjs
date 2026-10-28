// The console's popover against App Store Connect's toolkit's stroked
// popover (--tb-component-Popover-*) and its analytics overview's info
// popover: the ground, shadow, corners, narrowest and tallest. Golden, with
// their why: the room, the type, the widest, the stack's gap, the place,
// the dark side.

import { examples } from '../../test/harness.mjs';

const ex = examples(import.meta.url);
const P = 'component-Popover';
const POP = '.cui-console-popover';
const BUTTON = '[popovertarget]';
const page = (html) => `<div style="padding:24px 24px 240px">${html}</div>`;

export const cases = [
  {
    name: 'popover',
    html: page(ex('popover')),
    widths: [1200],
    checks: [
      { target: POP, props: ['display'], golden: 'closed until its button opens it' },
    ],
  },
  {
    name: 'popover open',
    html: page(ex('popover')),
    click: BUTTON,
    widths: [1200],
    checks: [
      { target: POP, props: ['background-color', 'box-shadow', 'border-top-left-radius', 'min-width', 'max-height'], tb: [`${P}-variant-variant-stroked-backgroundColor`, `${P}-variant-variant-stroked-boxShadow`, `${P}-variant-pageScale-desktop-borderRadius`, `${P}-minWidth`, `${P}-anatomy-Content-maxHeight`] },
      { target: POP, props: ['padding-top', 'padding-left', 'font-size', 'line-height', 'font-weight', 'color', 'max-width'], golden: 'App Store Connect\'s analytics overview\'s info popover: 16px in, 13px on 18px, the label, 350px at its widest' },
      { target: POP, props: ['display', 'flex-direction', 'row-gap'], golden: 'its parts stacked 8px apart: a control row\'s md button\'s (controlrow.json); the toolkit names no gap for a stack' },
      { target: POP, props: ['box-top', 'box-left', 'box-width'], golden: 'open: 8px under its button, centered on it' },
    ],
  },
  {
    name: 'popover open phone',
    html: page(ex('popover')),
    click: BUTTON,
    widths: [390],
    checks: [
      { target: POP, props: ['border-top-left-radius'], tb: [`${P}-variant-pageScale-mobile-borderRadius`] },
      { target: POP, props: ['box-left', 'box-width'], golden: 'on a phone, inside the window, 16px from its edges at the least' },
    ],
  },
  {
    name: 'popover dark',
    html: `<div class="cui-dark">${page(ex('popover'))}</div>`,
    click: BUTTON,
    widths: [1200],
    checks: [
      { target: POP, props: ['background-color', 'color'], golden: 'the dark side derived (tokens.css, decision 18)' },
    ],
  },
];
