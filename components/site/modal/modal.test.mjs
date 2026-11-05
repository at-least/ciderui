import { DIM, examples, MOTION, OUTLINE, PADDING, POINTER } from '../../../test/harness.mjs';

const ex = examples(import.meta.url);
const MODAL = ex('modal');
const OPENER = '[command="show-modal"]';
// the overlay's widths but 736 and 735, between ciderui's step and the store's (STEP, golden)
const WIDTHS = [1600, 1200, 900, 816, 737, 734, 640, 390];
const STEP = [736, 735];
const OPEN = { click: OPENER, height: 900 };
const PAGE = ex('page');
const PAGE_OPEN = { click: OPENER, height: 900 };
const PAGE_WIDTHS = [1600, 1200, 1069, 1068, 900, 735, 734, 640, 390];
const PAGE_DARK = PAGE.replace('class="cui-modal"', 'class="cui-modal cui-dark"');
const CLOSE = '.cui-modal-close';
const ONE = 'one close button, the store\'s circle and X: apple.com\'s marketing modal draws a black one, a second look for one job (ricmodal:close.icon, DESIGN.md decision 25)';
const BOX = 'the layer the window\'s box, the panel the window\'s width up to its own (816px, 640px below 735px), centered: an apple.com check reads no width';

export const cases = [
  {
    name: 'open',
    html: MODAL,
    ...OPEN,
    widths: WIDTHS,
    checks: [
      { target: '.cui-modal', props: ['position', 'top', 'left', 'overflow-x', 'overflow-y', 'opacity', ...MOTION], apple: ['overlay', 'overlay'] },
      { target: '.cui-modal', pseudo: '::backdrop', props: ['background-color'], apple: ['overlay', 'overlay'] },
      { target: 'html', props: ['overflow-x', 'overflow-y'], apple: ['overlay', 'root'] },
      { target: '.cui-modal', props: ['box-width', 'box-height', 'box-top', 'box-left'], golden: BOX },
      { target: '.cui-modal-panel', props: ['background-color', 'border-top-left-radius', 'border-top-right-radius', 'border-bottom-left-radius', 'border-bottom-right-radius', 'margin-top', 'margin-bottom', 'min-height'], apple: ['overlay', 'panel'] },
      { target: '.cui-modal-panel', props: PADDING, apple: ['overlay', 'content'] },
      { target: '.cui-modal-panel', props: ['box-width', 'box-left', 'box-top'], golden: BOX },
      { target: '.cui-modal-close', props: ['background-color', 'color', 'border-top-left-radius'], apple: ['overlay', 'close.circle'] },
      { target: '.cui-modal-close', props: MOTION, golden: POINTER },
      { target: '.cui-modal-close', state: 'hover', props: ['background-color', 'color'], apple: ['overlay', 'close.circle-hover', 'hover'] },
      { target: '.cui-modal-close', state: 'active', props: ['background-color', 'color'], apple: ['overlay', 'close.circle-hover', 'active'] },
      { target: '.cui-modal-close', props: ['box-width', 'box-height', 'box-top', 'box-left'], golden: 'apple.com\'s circle, 36px, 24px in from the panel\'s top right corner (its 44px button 20px in, overlay:close)' },
      { target: '.cui-modal-close', pseudo: '::before', props: ['width', 'height', 'top', 'left'], golden: 'its target, 44px, apple.com\'s button around the circle (overlay:close), drawn as the circle\'s ::before' },
      { target: '.cui-modal-close svg', props: ['box-width', 'box-height'], golden: 'the X 20px, apple.com\'s (overlay:close.icon), centered in the circle' },
      { target: '.cui-modal-close', props: ['focused'], golden: 'opened, the focus is on its close button, the dialog\'s first control, as a dialog\'s opening gives it (reference/dialog/; apple.com\'s script focuses the dialog itself)' },
    ],
  },
  {
    name: 'step',
    html: MODAL,
    ...OPEN,
    widths: STEP,
    checks: [
      { target: '.cui-modal-panel', props: ['box-width', 'margin-top', 'margin-bottom', ...PADDING], golden: 'the wide panel from 735px, the kit\'s step, where the store steps at 737px (overlay:panel, content: 640px, 20px down, 40px in, at 736 and 735px)' },
    ],
  },
  {
    // the delay holds the first frame
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
    name: 'scrim',
    html: MODAL,
    click: OPENER,
    height: 900,
    widths: [1600, 390],
    checks: [
      { target: '.cui-modal', props: ['hit'], golden: 'beside and below the short panel the pointer finds the curtain, the dim layer' },
      { target: '.cui-modal-curtain', props: ['position', 'box-top', 'box-left', 'box-width', 'box-height', 'background-color', 'cursor'], golden: 'the curtain the window\'s box, clear over the scrim, an arrow over it' },
      { target: '.cui-modal-curtain', props: ['focused'], golden: 'never the opening\'s focus: that stays on the close button' },
    ],
  },
  {
    name: 'scrim-click',
    html: MODAL,
    click: [OPENER, '.cui-modal-curtain'],
    height: 900,
    widths: [1600],
    checks: [
      { target: '.cui-modal', props: ['display'], golden: 'a click on the dim layer closes it, as apple.com\'s script does (a button in a dialog form: no script)' },
      { target: OPENER, props: ['focused'], golden: 'the focus back on the button that opened it (reference/dialog/)' },
    ],
  },
  {
    name: 'scrolled',
    html: `${MODAL.replace('</p>\n  </div>', '</p><div style="height: 1400px"></div>\n  </div>')}<script>document.addEventListener('click', () => setTimeout(() => { document.querySelector('.cui-modal').scrollTop = 400; }, 50));</script>`,
    ...OPEN,
    widths: [1200, 640, 390],
    checks: [
      { target: CLOSE, props: ['box-top'], golden: 'the panel scrolled 400px, its close button 24px from the window\'s top, stuck there, as the page modal\'s (apple.com\'s store modal scrolls its button away: overlay:close; a long sheet closed from anywhere)' },
      { target: 'form:has(> .cui-modal-close)', props: ['position', 'top', 'box-height'], golden: 'the close button\'s form sticks at the top as the panel scrolls, 0px tall' },
      { target: '.cui-modal', props: ['hit'], golden: 'scrolled, the dim layer beside the panel still the curtain: in the panel\'s cell, scrolled with it, as tall as it (not fixed: over a fixed box the wheel does not scroll the dialog, Chromium 151)' },
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
      { target: '.cui-modal', props: PADDING, apple: ['ricmodal', 'container'] },
      { target: '.cui-modal', props: ['display', 'align-items'], golden: 'the panel and the curtain in one grid cell, the panel centered by its auto margins, where apple.com\'s flex container centers it (align-items: center, ricmodal:container): when the panel is taller, it scrolls from its top' },
      { target: '.cui-modal', pseudo: '::backdrop', props: ['backdrop-filter'], golden: DIM },
      { target: '.cui-modal', pseudo: '::backdrop', props: ['background-color'], golden: 'apple.com\'s scrim, rgba(0,0,0,.48) (ricmodal:scrim, drawn there as a gradient), the store\'s too (scrim)' },
      { target: 'html', props: ['overflow-y'], apple: ['ricmodal', 'root'] },
      { target: '.cui-modal-panel', props: ['border-top-left-radius', 'border-bottom-right-radius'], apple: ['ricmodal', 'overlay'] },
      { target: '.cui-modal-panel', props: ['background-color', 'padding-top', 'padding-bottom'], apple: ['ricmodal', 'content'] },
      { target: '.cui-modal-panel', props: ['padding-left', 'padding-right'], apple: ['ricmodal', 'viewport'] },
      { target: '.cui-modal-panel', props: ['box-width', 'box-left'], golden: 'apple.com\'s panel, 816px from 1069px, 692px from 735px, 480px at most below, 10px from the window\'s sides, centered (ricmodal:overlay): an apple.com check reads no width' },
      { target: CLOSE, props: ['border-top-left-radius'], apple: ['ricmodal', 'close.icon'] },
      { target: CLOSE, props: ['background-color', 'color'], golden: ONE },
      { target: CLOSE, props: MOTION, golden: POINTER },
      { target: CLOSE, props: ['box-width', 'box-height', 'box-top', 'box-right'], golden: 'apple.com\'s 36px circle 20px in from the panel\'s top right corner (ricmodal:close.icon, its 44px button 20px down and 12px from the end)' },
      { target: `${CLOSE} svg`, props: ['box-width', 'box-height'], golden: ONE },
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
      { target: CLOSE, state: 'hover', props: ['background-color', 'color'], golden: ONE },
      { target: CLOSE, state: 'active', props: ['background-color', 'color'], golden: ONE },
      { target: CLOSE, state: 'focus-visible', props: OUTLINE, apple: ['ricmodal', 'close.icon-states', 'focus-visible'] },
      { target: CLOSE, pseudo: '::before', props: ['width', 'height', 'top', 'left'], golden: 'its 44px target around its circle, where apple.com\'s button reaches to its right and below it (ricmodal:close)' },
    ],
  },
  {
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
    // the delay holds the first frame, as in fade
    name: 'page-start',
    html: `<style>.cui-modal { transition-delay: 10s !important; }</style>${PAGE}`,
    click: OPENER,
    wait: 0,
    widths: [1600],
    checks: [
      { target: '.cui-modal', pseudo: '::backdrop', props: ['background-color', 'backdrop-filter'], golden: 'the plain scrim from its first frame, where apple.com\'s darkens and blurs from 0 (DESIGN.md decision 25)' },
      { target: '.cui-modal', props: ['opacity'], golden: 'the regular modal\'s fade, from transparent: one modal motion' },
      { target: '.cui-modal-panel', props: ['opacity'], golden: 'the panel with its layer, not after it' },
    ],
  },
  {
    name: 'page-motion',
    html: PAGE,
    ...PAGE_OPEN,
    widths: [1600],
    checks: [
      { target: '.cui-modal', props: MOTION, golden: 'the regular modal\'s 0.1s: one modal motion, where apple.com\'s marketing modal takes 0.75s and its panel 0.5s more, time the reader waits for nothing (DESIGN.md decision 25)' },
      { target: '.cui-modal', pseudo: '::backdrop', props: MOTION, golden: 'the regular modal\'s 0.1s: apple.com\'s 0.75s darkening and blur left out (DESIGN.md decision 25)' },
      { target: '.cui-modal-panel', props: [...MOTION, 'transition-delay'], golden: 'no motion of its own, where apple.com\'s fades in 0.5s from 0.25s (DESIGN.md decision 25)' },
    ],
  },
  {
    name: 'page-dark',
    html: `<section class="cui-dark">${PAGE}</section>`,
    ...PAGE_OPEN,
    widths: [1600],
    checks: [
      { target: '.cui-modal-panel', props: ['background-color', 'color'], golden: 'light in a dark block, as the store\'s modal: dark only with cui-dark on itself (page-own-dark)' },
      { target: CLOSE, props: ['background-color', 'color'], golden: ONE },
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
      { target: CLOSE, props: ['border-top-left-radius'], apple: ['ricmodal-dark', 'close.icon'] },
      { target: CLOSE, props: ['background-color', 'color'], golden: ONE },
      { target: CLOSE, state: 'hover', props: ['background-color', 'color'], golden: ONE },
      { target: CLOSE, state: 'active', props: ['background-color', 'color'], golden: ONE },
      { target: CLOSE, state: 'focus-visible', props: OUTLINE, apple: ['ricmodal-dark', 'close.icon-states', 'focus-visible'] },
      { target: '.cui-modal', props: ['background-color'], golden: 'the layer clear, over the scrim: cui-dark\'s ground is the panel\'s' },
      { target: '.cui-modal', pseudo: '::backdrop', props: ['background-color', 'backdrop-filter'], golden: 'the scrim the light one\'s, rgba(0,0,0,.48), no blur (DESIGN.md decision 25): apple.com\'s dark theme sets none (MacBook Pro\'s dark modal\'s #1d1d1f fill its own, ricmodal-dark:modal; AirPods Max\'s dark story\'s black, reference/apple.com/README.md)' },
    ],
  },
  {
    name: 'page-touch',
    html: PAGE,
    ...PAGE_OPEN,
    touch: true,
    widths: [1200, 390],
    checks: [{ target: CLOSE, state: 'hover', props: ['background-color', 'color'], rest: 'no hover on a touch screen (rule 5), where apple.com\'s keeps it (touch-ricmodal:close.icon-states)' }],
  },
  {
    name: 'touch',
    html: MODAL,
    ...OPEN,
    touch: true,
    widths: [1200, 390],
    checks: [{ target: CLOSE, state: 'hover', props: ['background-color', 'color'], rest: 'no hover on a touch screen (rule 5), where apple.com\'s keeps it (touch-overlay:close.circle-hover)' }],
  },
];
