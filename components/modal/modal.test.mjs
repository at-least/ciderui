// The modal against apple.com's store overlay (overlay:*: the layer, the
// page held, the panel and its content's room, the close button at rest,
// hovered and pressed), opened by its button with no script, at the widths
// the overlay was read at; its fade's start, the focus on its close button,
// the close button and Escape closing it with the focus back on its opener
// (reference/dialog/), and a modal in a dark block. Golden, with their why:
// the layer's and the panel's boxes (an apple.com check reads no width), the
// close button's place and its X's size, the store's step at 736/737px where
// ciderui's is at 735, the focus, and the modal light in a dark block.

import { examples, PADDING } from '../../test/harness.mjs';

const ex = examples(import.meta.url);
const MODAL = ex('modal');
const OPENER = '[command="show-modal"]';
const MOTION = ['transition-property', 'transition-duration', 'transition-timing-function'];
// the overlay's widths but the two between ciderui's step (735) and the
// store's (737), golden below
const WIDTHS = [1600, 1200, 900, 816, 737, 734, 640, 390];
const STEP = [736, 735];
const OPEN = { click: OPENER, height: 900 };
const BOX = 'the layer the window\'s box, the panel the window\'s width up to its own (816px, 640px below 735px), centered: an apple.com check reads no width';

export const cases = [
  {
    name: 'open',
    html: MODAL,
    ...OPEN,
    widths: WIDTHS,
    checks: [
      // the layer, the page dimmed under it and held
      { target: '.cui-modal', props: ['position', 'top', 'left', 'overflow-x', 'overflow-y', 'opacity', ...MOTION], apple: ['overlay', 'overlay'] },
      { target: '.cui-modal', pseudo: '::backdrop', props: ['background-color'], apple: ['overlay', 'overlay'] },
      { target: 'html', props: ['overflow-x', 'overflow-y'], apple: ['overlay', 'root'] },
      { target: '.cui-modal', props: ['box-width', 'box-height', 'box-top', 'box-left'], golden: BOX },
      // the panel and its content's room (apple.com's content box, inside its panel)
      { target: '.cui-modal-panel', props: ['background-color', 'border-top-left-radius', 'border-top-right-radius', 'border-bottom-left-radius', 'border-bottom-right-radius', 'margin-top', 'margin-bottom'], apple: ['overlay', 'panel'] },
      { target: '.cui-modal-panel', props: PADDING, apple: ['overlay', 'content'] },
      { target: '.cui-modal-panel', props: ['box-width', 'box-left', 'box-top'], golden: BOX },
      // the close button: its circle, hovered and pressed
      { target: '.cui-modal-close', props: ['background-color', 'color', 'border-top-left-radius', ...MOTION], apple: ['overlay', 'close.circle'] },
      { target: '.cui-modal-close', state: 'hover', props: ['background-color', 'color'], apple: ['overlay', 'close.circle-hover', 'hover'] },
      { target: '.cui-modal-close', state: 'active', props: ['background-color', 'color'], apple: ['overlay', 'close.circle-hover', 'active'] },
      { target: '.cui-modal-close', props: ['box-width', 'box-height', 'box-top', 'box-left'], golden: 'apple.com\'s circle, 36px, 24px in from the panel\'s top right corner (its 44px button 20px in, overlay:close)' },
      { target: '.cui-modal-close', pseudo: '::before', props: ['width', 'height', 'top', 'left'], golden: 'its target, 44px, apple.com\'s button around the circle (overlay:close), drawn as the circle\'s ::before' },
      { target: '.cui-modal-close svg', props: ['box-width', 'box-height'], golden: 'the X 20px, apple.com\'s (overlay:close.icon), centered in the circle' },
      // opened, the focus on the close button
      { target: '.cui-modal-close', props: ['focused'], golden: 'opened, the focus is on its close button, the dialog\'s first control, as a dialog\'s opening gives it (reference/dialog/; apple.com\'s script focuses the dialog itself)' },
    ],
  },
  {
    // the window between ciderui's step and the store's
    name: 'step',
    html: MODAL,
    ...OPEN,
    widths: STEP,
    checks: [
      { target: '.cui-modal-panel', props: ['box-width', 'margin-top', 'margin-bottom', ...PADDING], golden: 'the wide panel from 735px, the kit\'s step, where the store steps at 737px (overlay:panel, content: 640px, 20px down, 40px in, at 736 and 735px)' },
    ],
  },
  {
    // its fade's first frame
    name: 'fade',
    html: MODAL,
    click: OPENER,
    wait: 0,
    widths: [1600],
    checks: [
      { target: '.cui-modal', props: ['opacity'], golden: 'it fades in from transparent (@starting-style), apple.com\'s layer\'s 0.1s (overlay:overlay\'s animation, 0 to 1)' },
    ],
  },
  {
    name: 'closed',
    html: MODAL,
    click: [OPENER, '.cui-modal-close'],
    widths: [1600],
    checks: [
      { target: '.cui-modal', props: ['display'], golden: 'its close button closes it (a form\'s, method="dialog"), at once' },
      { target: OPENER, props: ['focused'], golden: 'the focus back on the button that opened it (reference/dialog/)' },
      { target: 'html', props: ['overflow-y'], golden: 'the page no longer held' },
    ],
  },
  {
    name: 'escape',
    html: MODAL,
    click: OPENER,
    press: 'Escape',
    widths: [1600],
    checks: [
      { target: '.cui-modal', props: ['display'], golden: 'Escape closes it' },
      { target: OPENER, props: ['focused'], golden: 'the focus back on the button that opened it (reference/dialog/)' },
    ],
  },
  {
    name: 'dark',
    html: `<section class="cui-dark">${MODAL}</section>`,
    ...OPEN,
    widths: [1600],
    checks: [
      { target: '.cui-modal-panel', props: ['background-color', 'color'], golden: 'light in a dark block, as a field: apple.com\'s store has no dark overlay' },
      { target: '.cui-modal-close', props: ['background-color', 'color'], golden: 'the close button\'s light colors in a dark block' },
    ],
  },
];
