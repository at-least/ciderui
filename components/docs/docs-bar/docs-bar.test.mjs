// The docs bar against Apple's developer documentation (DocC): its height,
// room, hairline and title (docc-page:bar.*) at each of DocC's steps, its
// button and glyph below 1024px, where it shows, its ground (docc:bar.ground);
// light and dark. Golden, with their why: the letter-spacing, the button
// gone from 1024px and the title's place, the button's touch target.

import { DOCC_STEPS, PADDING, SIZE, examples } from '../../../test/harness.mjs';

const ex = examples(import.meta.url);
const B = '.cui-docs-bar';
const T = '.cui-docs-bar-toggle';
const TITLE = '.cui-docs-bar-title';
const NARROW = [1023, 900, 768, 767, 736, 735, 390];

export const cases = [
  {
    name: 'bar',
    html: ex('bar'),
    widths: DOCC_STEPS,
    checks: [
      { target: B, props: ['height', 'padding-left', 'padding-right', 'box-height'], apple: ['docc-page', 'bar.content'] },
      { target: B, pseudo: '::after', props: ['content', 'position', 'height', 'bottom', 'background-color'], apple: ['docc-page', 'bar.rule'] },
      { target: TITLE, props: [...SIZE, 'color', 'text-decoration-line', 'text-decoration-color', 'display', 'align-items', 'box-height'], apple: ['docc-page', 'bar.title.link'] },
      { target: TITLE, props: ['letter-spacing'], golden: 'the system font\'s fitted letter-spacing (tokens.css), not DocC\'s' },
      { target: TITLE, props: ['box-x'], golden: 'the title 30px after the bar\'s room where the button shows (DocC\'s 46px at 390px, 52px from 768px); first in the bar from 1024px, 22px in, where DocC\'s button stays and collapses the navigator by its script' },
      { target: T, props: ['display'], golden: 'the button below 1024px alone: from there the navigator is a column (cui-docs-page)' },
    ],
  },
  {
    name: 'bar narrow',
    html: ex('bar'),
    widths: NARROW,
    checks: [
      { target: T, props: ['height', 'color', 'background-color', 'border-top-left-radius', 'margin-left', ...PADDING, 'box-height'], apple: ['docc-page', 'bar.toggle'] },
      { target: T, state: 'hover', props: ['color', 'background-color'], apple: ['docc-page', 'bar.toggle', 'hover'] },
      { target: T, props: ['margin-top'], apple: ['docc-page', 'bar.toggle.wrapper'] },
      { target: `${T} svg`, props: ['height'], apple: ['docc-page', 'bar.toggle.icon'] },
      { target: T, props: ['width'], golden: 'square, as DocC\'s 29px (docc-page:bar.toggle width; an apple check reads no width)' },
      { target: `${T} svg`, props: ['width'], golden: 'square, as DocC\'s 19px glyph (docc-page:bar.toggle.icon width)' },
      { target: T, props: ['margin-right', 'box-x'], golden: 'the button 5px out before the bar\'s room and 6px before the title: DocC\'s in its 19px column, 11px before the title, its margins -5px' },
    ],
  },
  {
    name: 'bar ground',
    html: ex('bar'),
    checks: [
      { target: B, props: ['background-color'], apple: ['docc', 'bar.ground'] },
    ],
  },
  {
    name: 'bar dark',
    html: `<div class="cui-dark">${ex('bar')}</div>`,
    widths: [900, 390],
    checks: [
      { target: B, props: ['background-color'], apple: ['docc-dark', 'bar.ground'] },
      { target: B, pseudo: '::after', props: ['background-color'], apple: ['docc-page-dark', 'bar.rule'] },
      { target: TITLE, props: ['color'], apple: ['docc-page-dark', 'bar.title.link'] },
      { target: T, props: ['color'], apple: ['docc-page-dark', 'bar.toggle'] },
    ],
  },
  {
    name: 'bar touch',
    html: ex('bar'),
    widths: [390],
    touch: true,
    checks: [
      { target: T, pseudo: '::after', props: ['width', 'height'], golden: 'the 29px button\'s target 44px square on a touch screen, the HIG\'s 44pt (DocC\'s is its own 29px)' },
    ],
  },
];
