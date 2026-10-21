// The modal against apple.com's store overlay (overlay:*: the layer, the
// page held, the panel and its content's room, the close button at rest,
// hovered and pressed), opened by its button with no script, at the widths
// the overlay was read at; its fade's start, the focus on its close button,
// the close button and Escape closing it with the focus back on its opener
// (reference/dialog/), and a modal in a dark block. Golden, with their why:
// the layer's and the panel's boxes (an apple.com check reads no width), the
// close button's place and its X's size, the store's step at 736/737px where
// ciderui's is at 735, the focus, and the modal light in a dark block. The
// marketing pages' modal, data-variant="page" (ricmodal:*): the layer and
// its blurred scrim, the page held, the panel centered at its widths, its
// corners and room, the close button's circle at rest, hovered, pressed and
// focused, sticking at the top as the panel scrolls; its motion, fitted to
// apple.com's script's frames, and its start; and its dark one, cui-dark on
// the dialog (ricmodal-dark:*): its panel, its heading, its close button at
// rest, hovered, pressed and focused.

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
// the marketing pages' modal, opened at the widths its reference was read at
const PAGE = ex('page');
const PAGE_OPEN = { click: OPENER, height: 900 };
const PAGE_WIDTHS = [1600, 1200, 1069, 1068, 900, 735, 734, 640, 390];
// its dark one, cui-dark on the dialog (ricmodal-dark:*)
const PAGE_DARK = PAGE.replace('class="cui-modal"', 'class="cui-modal cui-dark"');
const CLOSE = '.cui-modal-close';
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
    // its fade's first frame: held there, as a fade with no delay can move
    // before it is read (0.50 of the way once, the suite's run busy)
    name: 'fade',
    html: `<style>.cui-modal { transition-delay: 10s !important; }</style>${MODAL}`,
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
  {
    name: 'page-open',
    html: PAGE,
    ...PAGE_OPEN,
    widths: PAGE_WIDTHS,
    checks: [
      { target: '.cui-modal', props: ['position', 'top', 'left', 'overflow-y'], apple: ['ricmodal', 'modal'] },
      { target: '.cui-modal', props: [...PADDING, 'display'], apple: ['ricmodal', 'container'] },
      { target: '.cui-modal', props: ['align-items'], golden: 'the panel centered by its auto margins in the flex layer, where apple.com\'s container centers it (align-items: center, ricmodal:container): when the panel is taller, it scrolls from its top' },
      { target: '.cui-modal', pseudo: '::backdrop', props: ['backdrop-filter'], apple: ['ricmodal', 'scrim'] },
      { target: '.cui-modal', pseudo: '::backdrop', props: ['background-color'], golden: 'apple.com\'s scrim, rgba(0,0,0,.48) (ricmodal:scrim, drawn there as a gradient), the store\'s too (scrim)' },
      { target: 'html', props: ['overflow-y'], apple: ['ricmodal', 'root'] },
      { target: '.cui-modal-panel', props: ['border-top-left-radius', 'border-bottom-right-radius'], apple: ['ricmodal', 'overlay'] },
      { target: '.cui-modal-panel', props: ['background-color', 'padding-top', 'padding-bottom'], apple: ['ricmodal', 'content'] },
      { target: '.cui-modal-panel', props: ['padding-left', 'padding-right'], apple: ['ricmodal', 'viewport'] },
      { target: '.cui-modal-panel', props: ['box-width', 'box-left'], golden: 'apple.com\'s panel, 816px from 1069px, 692px from 735px, 480px at most below, 10px from the window\'s sides, centered (ricmodal:overlay): an apple.com check reads no width' },
      { target: CLOSE, props: ['background-color', 'color', 'border-top-left-radius', 'transition-property', 'transition-duration', 'transition-timing-function'], apple: ['ricmodal', 'close.icon'] },
      { target: CLOSE, props: ['box-width', 'box-height', 'box-top', 'box-right'], golden: 'apple.com\'s 36px circle 20px in from the panel\'s top right corner (ricmodal:close.icon, its 44px button 20px down and 12px from the end)' },
      { target: `${CLOSE} svg`, props: ['box-width', 'box-height'], golden: 'its X drawn in 20 by 20 in the circle\'s middle, its ink 12px as apple.com\'s (icon::CLOSE_SMALL, reference/icons/close.mjs)' },
      { target: CLOSE, props: ['focused'], golden: 'opened, the focus is on its close button, the dialog\'s first control (apple.com\'s script focuses the panel)' },
      { target: 'form:has(> .cui-modal-close)', props: ['position', 'top', 'box-height'], golden: 'the close button\'s form sticks at the top as the panel scrolls, 0px tall, as apple.com\'s button\'s box (ricmodal:close.wrapper); its top the layer\'s room above the panel taken back, 40px (28px below 735px), so it sticks at the window\'s top' },
    ],
  },
  {
    name: 'page-states',
    html: PAGE,
    ...PAGE_OPEN,
    widths: [1600, 390],
    checks: [
      { target: CLOSE, state: 'hover', props: ['background-color', 'color'], apple: ['ricmodal', 'close.icon-states', 'hover'] },
      { target: CLOSE, state: 'active', props: ['background-color', 'color'], apple: ['ricmodal', 'close.icon-states', 'active'] },
      { target: CLOSE, state: 'focus-visible', props: ['outline-style', 'outline-color', 'outline-width', 'outline-offset'], apple: ['ricmodal', 'close.icon-states', 'focus-visible'] },
    ],
  },
  {
    // a panel taller than the window, scrolled 400px: its close button stays
    // at the top
    name: 'page-scrolled',
    html: `${PAGE.replace('</p>\n  </div>', '</p><div style="height: 1400px"></div>\n  </div>')}<script>document.addEventListener('click', () => setTimeout(() => { document.querySelector('.cui-modal').scrollTop = 400; }, 50));</script>`,
    ...PAGE_OPEN,
    widths: [1200, 390],
    checks: [
      { target: CLOSE, props: ['box-top'], golden: 'the panel scrolled 400px, its close button 20px from the window\'s top, stuck there (apple.com\'s sticky button, ricmodal:close.wrapper)' },
      { target: '.cui-modal-panel', props: ['box-top'], golden: 'the panel taller than the window scrolled from its top: 400px up from its 40px (28px below 735px) of room' },
    ],
  },
  {
    // its motion's start, held there to be read
    name: 'page-start',
    html: `<style>.cui-modal::backdrop, .cui-modal-panel { transition-delay: 10s !important; }</style>${PAGE}`,
    click: OPENER,
    wait: 0,
    widths: [1600],
    checks: [
      { target: '.cui-modal', pseudo: '::backdrop', props: ['background-color', 'backdrop-filter'], golden: 'the scrim from clear and unblurred, apple.com\'s from (its frames sampled: its darkness and blur from 0)' },
      { target: '.cui-modal-panel', props: ['opacity'], golden: 'the panel from transparent, apple.com\'s from (its frames sampled)' },
      { target: '.cui-modal', props: ['opacity'], golden: 'the layer itself not fading: its scrim and its panel do' },
    ],
  },
  {
    name: 'page-motion',
    html: PAGE,
    ...PAGE_OPEN,
    widths: [1600],
    checks: [
      { target: '.cui-modal', pseudo: '::backdrop', props: MOTION, golden: 'its darkness and blur together in 0.75s, cubic-bezier(.4,0,.6,1), fitted to apple.com\'s script\'s frames (from 0.15s to 0.9s after its click, the script\'s start 0.1s in)' },
      { target: '.cui-modal-panel', props: [...MOTION, 'transition-delay'], golden: 'the panel in 0.5s from 0.25s, cubic-bezier(.4,0,.6,1), fitted to apple.com\'s frames (from about 0.35s to 0.85s)' },
    ],
  },
  {
    name: 'page-dark',
    html: `<section class="cui-dark">${PAGE}</section>`,
    ...PAGE_OPEN,
    widths: [1600],
    checks: [
      { target: '.cui-modal-panel', props: ['background-color', 'color'], golden: 'light in a dark block, as the store\'s modal: dark only with cui-dark on itself (page-own-dark)' },
      { target: CLOSE, props: ['background-color', 'color'], golden: 'the close button\'s dark circle in a dark block, the modal light' },
    ],
  },
  {
    name: 'page-own-dark',
    html: PAGE_DARK,
    ...PAGE_OPEN,
    widths: [1600, 390],
    checks: [
      { target: '.cui-modal-panel', props: ['background-color', 'color'], apple: ['ricmodal-dark', 'content'] },
      { target: '.cui-modal-panel h2', props: ['color'], apple: ['ricmodal-dark', 'heading'] },
      { target: CLOSE, props: ['background-color', 'color', 'border-top-left-radius'], apple: ['ricmodal-dark', 'close.icon'] },
      { target: CLOSE, state: 'hover', props: ['background-color', 'color'], apple: ['ricmodal-dark', 'close.icon-states', 'hover'] },
      { target: CLOSE, state: 'active', props: ['background-color', 'color'], apple: ['ricmodal-dark', 'close.icon-states', 'active'] },
      { target: CLOSE, state: 'focus-visible', props: ['outline-style', 'outline-color', 'outline-width', 'outline-offset'], apple: ['ricmodal-dark', 'close.icon-states', 'focus-visible'] },
      { target: '.cui-modal', props: ['background-color'], golden: 'the layer clear, over the scrim: cui-dark\'s ground is the panel\'s' },
      { target: '.cui-modal', pseudo: '::backdrop', props: ['background-color', 'backdrop-filter'], golden: 'the scrim the light one\'s, rgba(0,0,0,.48) blurred 20px: apple.com\'s dark theme sets none (MacBook Pro\'s dark modal\'s #1d1d1f fill its own, ricmodal-dark:modal; AirPods Max\'s dark story\'s black, reference/apple.com/README.md)' },
    ],
  },
  {
    name: 'page-touch',
    html: PAGE,
    ...PAGE_OPEN,
    touch: true,
    widths: [1200, 390],
    checks: [{ target: CLOSE, state: 'hover', props: ['background-color'], rest: 'no hover on a touch screen (rule 5)' }],
  },
];
