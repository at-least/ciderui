// The local nav against apple.com's (MacBook Air's, localnav:*; dark,
// MacBook Pro's, localnav-dark:*): the bar from 834px (its height, its
// stickiness, its material and rule, the content's width, the name, the
// links, the current page's mark, the button), the phone's bar (the same at
// 833 and 390px, the chevron), the menu opened by the chevron (the ground,
// the rows, the mark, the chevron turned up) and its motion, the keys (the
// checkbox's focus ring on the chevron, Tab on to the first link), dark, a
// touch screen. Golden, with their why: the letter-spacing, the name's
// color (its opacity in it), the rows' delays (apple.com's recorded
// animations), the chevron turned (reference/icons/).

import { examples, PADDING } from '../../test/harness.mjs';

const ex = examples(import.meta.url);
const NAV = ex('localnav');
const dark = NAV.replace('class="cui-localnav"', 'class="cui-localnav cui-dark"');

const DESKTOP = [1600, 1200, 900, 834];
const PHONE = [833, 390];
const MOTION = ['transition-property', 'transition-duration', 'transition-timing-function'];
const TYPE = ['font-size', 'line-height', 'font-weight'];
const TITLE = '.cui-localnav-title';
const CURRENT = '.cui-localnav-list > li:nth-child(1) > .cui-localnav-link';
const LINK = '.cui-localnav-list > li:nth-child(2) > .cui-localnav-link';
const BUTTON = '.cui-localnav-actions > .cui-btn';
const CHEVRON = '.cui-localnav-menubutton';
const OPEN = { click: CHEVRON, widths: [390] };

export const cases = [
  {
    name: 'bar',
    html: NAV,
    widths: DESKTOP,
    checks: [
      { target: '.cui-localnav', props: ['position', 'top', 'z-index', 'box-height'], apple: ['localnav', 'localnav'] },
      { target: '.cui-localnav', props: ['background-color', 'backdrop-filter', 'transition-duration', 'transition-timing-function'], apple: ['localnav', 'localnav.background'] },
      { target: '.cui-localnav', props: ['transition-property'], golden: 'apple.com\'s background and backdrop-filter (localnav:localnav.background); Tailwind\'s build names -webkit-backdrop-filter too' },
      { target: '.cui-localnav', pseudo: '::after', props: ['background-color', 'height', 'position', 'bottom'], apple: ['localnav', 'localnav.rule'] },
      { target: '.cui-localnav-content', props: ['max-width', 'padding-left', 'padding-right', 'margin-left', 'margin-right'], apple: ['localnav', 'localnav.content'] },
      { target: TITLE, props: [...TYPE, 'text-decoration-line', ...MOTION, 'box-top', 'box-left', 'box-height'], apple: ['localnav', 'localnav.title'] },
      { target: TITLE, props: ['letter-spacing'], golden: 'the system font\'s fitted letter-spacing (tokens.css), not apple.com\'s' },
      { target: TITLE, props: ['color', 'opacity'], golden: 'apple.com\'s black at opacity .88 (localnav:localnav.title), as one color, on-bar-2' },
      { target: '.cui-localnav-list > li:nth-child(2)', props: ['margin-left', 'margin-right'], apple: ['localnav', 'localnav.item'] },
      { target: LINK, props: [...TYPE, 'color', 'opacity', 'text-decoration-line', 'box-top', 'box-height'], apple: ['localnav', 'localnav.link'] },
      { target: LINK, props: ['display'], golden: 'a flex item, so block: apple.com\'s inline-block sits in a 12px line, ciderui\'s in its item (the same place, localnav:localnav.link)' },
      { target: LINK, state: 'hover', props: ['opacity'], apple: ['localnav', 'localnav.link', 'hover'] },
      { target: LINK, props: ['letter-spacing'], golden: 'the system font\'s fitted letter-spacing (tokens.css), not apple.com\'s' },
      { target: CURRENT, props: [...TYPE, 'color', 'opacity', 'cursor', 'box-top'], apple: ['localnav', 'localnav.current'] },
      { target: CURRENT, pseudo: '::after', props: ['background-color', 'opacity', 'height', 'top', 'left', 'border-top-left-radius', 'position'], apple: ['localnav', 'localnav.current-mark'] },
      { target: BUTTON, props: [...TYPE, 'color', ...PADDING, 'border-top-left-radius', 'background-color', 'min-width', 'box-top', 'box-height'], apple: ['localnav', 'localnav.button'] },
      { target: BUTTON, state: 'hover', props: ['background-color'], apple: ['localnav', 'localnav.button', 'hover'] },
      { target: BUTTON, state: 'active', props: ['background-color'], apple: ['localnav', 'localnav.button', 'active'] },
      { target: '.cui-localnav-actions', props: ['margin-left'], golden: 'the button\'s own 24px before it, apple.com\'s action\'s (localnav:localnav.action)' },
      { target: BUTTON, props: ['margin-left'], apple: ['localnav', 'localnav.action'] },
      { target: CHEVRON, props: ['display'], golden: 'from 834px the links are in the bar: no chevron (apple.com\'s is not displayed)' },
    ],
  },
  {
    name: 'phone',
    html: NAV,
    widths: PHONE,
    checks: [
      { target: '.cui-localnav', props: ['position', 'top', 'z-index', 'box-height'], apple: ['localnav', 'localnav'] },
      { target: '.cui-localnav', props: ['background-color', 'backdrop-filter'], apple: ['localnav', 'localnav.background'] },
      { target: '.cui-localnav', pseudo: '::after', props: ['background-color', 'height'], apple: ['localnav', 'localnav.rule'] },
      { target: '.cui-localnav-content', props: ['padding-left', 'padding-right'], apple: ['localnav', 'localnav.content'] },
      { target: TITLE, props: [...TYPE, 'box-top', 'box-left', 'box-height'], apple: ['localnav', 'localnav.title'] },
      { target: CHEVRON, props: ['display', 'cursor', 'margin-right', 'box-top', 'box-height'], apple: ['localnav', 'localnav.chevron'] },
      { target: CHEVRON, props: ['box-width', 'box-left'], golden: 'apple.com\'s 40 by 30 chevron, 4px into the 16px before the button (localnav:localnav.chevron; its place follows the button\'s text)' },
      { target: CHEVRON, props: ['color', 'opacity'], apple: ['localnav', 'localnav.chevron-image'] },
      { target: CHEVRON, state: 'hover', props: ['opacity'], apple: ['localnav', 'localnav.chevron-image', 'hover'] },
      { target: `${CHEVRON} svg`, props: ['box-top', 'box-height'], apple: ['localnav', 'localnav.chevron-svg'] },
      { target: `${CHEVRON} svg`, props: ['box-width', 'box-x'], golden: 'apple.com\'s 16 by 9 chevron, centered in its 40px (localnav:localnav.chevron-svg)' },
      { target: '.cui-localnav-menu', props: ['visibility', 'box-height'], apple: ['localnav', 'localnav.tray'] },
      { target: BUTTON, props: ['box-top', 'box-right', 'box-height'], apple: ['localnav', 'localnav.button'] },
      { target: '.cui-localnav-toggle', props: ['opacity', 'box-height'], golden: 'the checkbox out of sight, in the focus order (1px, transparent)' },
    ],
  },
  {
    name: 'menu',
    html: NAV,
    ...OPEN,
    reducedMotion: 'reduce',
    checks: [
      { target: '.cui-localnav', props: ['background-color', 'backdrop-filter'], apple: ['localnav', 'tray.background'] },
      { target: '.cui-localnav-menu', props: ['visibility'], apple: ['localnav', 'tray.open'] },
      { target: LINK, props: [...TYPE, 'color', 'opacity', 'display', 'align-items', 'box-top', 'box-left', 'box-height'], apple: ['localnav', 'tray.link'] },
      { target: LINK, state: 'hover', props: ['opacity'], apple: ['localnav', 'tray.link', 'hover'] },
      { target: CURRENT, props: [...TYPE, 'color', 'opacity', 'box-top', 'box-left', 'box-height'], apple: ['localnav', 'tray.current'] },
      { target: CURRENT, pseudo: '::after', props: ['background-color', 'opacity', 'height', 'top', 'left', 'border-top-left-radius'], apple: ['localnav', 'tray.current-mark'] },
      { target: CURRENT, pseudo: '::after', props: ['width'], golden: 'apple.com\'s mark, 1.2px wide (localnav:tray.current-mark, 1.1875px as drawn)' },
      { target: '.cui-localnav-menu', props: ['box-top', 'box-height'], golden: 'the menu below the bar, its last row\'s foot 27px above its own at 221px, apple.com\'s (localnav:tray.open: from 24px, 197px)' },
      { target: `${CHEVRON} svg`, props: ['transform'], golden: 'the chevron turned up: apple.com\'s points flipped about its middle (reference/icons/)' },
      { target: '.cui-localnav', pseudo: '::after', props: ['content'], golden: 'the rule moves to the menu\'s foot' },
    ],
  },
  {
    name: 'menu-motion',
    html: NAV,
    ...OPEN,
    checks: [
      { target: '.cui-localnav', props: ['transition-duration', 'transition-timing-function'], apple: ['localnav', 'tray.background'] },
      { target: LINK, props: ['transition-duration', 'transition-timing-function'], golden: 'apple.com\'s rows in 0.32s with its curve (localnav:tray.link); its list names -webkit-transform as well' },
      { target: CURRENT, props: ['transition-delay'], golden: 'apple.com\'s first row from 260ms (localnav:tray.open\'s animations)' },
      { target: LINK, props: ['transition-delay'], golden: 'apple.com\'s second row from 280ms, 20ms after the first' },
      { target: '.cui-localnav-menu', props: ['transition-duration', 'transition-timing-function'], golden: 'apple.com\'s menu opens in 0.32s with its curve (localnav:tray.open, its max-height)' },
    ],
  },
  {
    // the focus from the bar's empty middle: the name, the checkbox, its space bar, the next stop
    name: 'menu-keys',
    html: NAV,
    click: '.cui-localnav-content',
    press: ['Tab', 'Tab', 'Space', 'Tab'],
    widths: [390],
    reducedMotion: 'reduce',
    checks: [
      { target: LINK.replace('nth-child(2)', 'nth-child(1)'), props: ['focused'], golden: 'the checkbox, then its space bar, opens the menu: the next stop is the first link' },
      { target: '.cui-localnav-menu', props: ['visibility'], golden: 'opened by the keys' },
    ],
  },
  {
    name: 'menu-focus',
    html: NAV,
    click: '.cui-localnav-content',
    press: ['Tab', 'Tab'],
    widths: [390],
    reducedMotion: 'reduce',
    checks: [
      { target: CHEVRON, props: ['outline-style', 'outline-color', 'outline-width', 'outline-offset'], golden: 'the checkbox\'s focus ring drawn on the chevron: the one ring (rule 9), moved' },
      { target: '.cui-localnav-menu', props: ['visibility'], golden: 'closed, its links out of the focus order' },
    ],
  },
  {
    name: 'dark',
    html: dark,
    widths: [1200],
    checks: [
      { target: '.cui-localnav', props: ['background-color', 'backdrop-filter'], apple: ['localnav-dark', 'localnav.background'] },
      { target: '.cui-localnav', pseudo: '::after', props: ['background-color'], apple: ['localnav-dark', 'localnav.rule'] },
      { target: TITLE, props: ['color', 'opacity'], golden: 'apple.com\'s white at opacity .92 (localnav-dark:localnav.title), as one color, on-bar-2' },
      { target: LINK, props: ['color', 'opacity'], apple: ['localnav-dark', 'localnav.link'] },
      { target: CURRENT, props: ['color', 'opacity'], apple: ['localnav-dark', 'localnav.current'] },
      { target: CURRENT, pseudo: '::after', props: ['background-color', 'opacity'], apple: ['localnav-dark', 'localnav.current-mark'] },
    ],
  },
  {
    name: 'dark-menu',
    html: dark,
    ...OPEN,
    reducedMotion: 'reduce',
    checks: [
      { target: '.cui-localnav', props: ['background-color'], apple: ['localnav-dark', 'tray.background'] },
      { target: LINK, props: ['color', 'opacity'], apple: ['localnav-dark', 'tray.link'] },
      { target: CHEVRON, props: ['color', 'opacity'], apple: ['localnav-dark', 'localnav.chevron-image'] },
    ],
  },
  {
    name: 'touch',
    html: NAV,
    touch: true,
    widths: [1200, 390],
    checks: [
      { target: LINK, state: 'hover', props: ['opacity'], rest: 'no hover on a touch screen (rule 5)' },
      { target: CHEVRON, state: 'hover', props: ['opacity'], rest: 'no hover on a touch screen (rule 5)' },
    ],
  },
];
