// The global nav against apple.com: the bar from 834px (home:nav*: its
// height, material, content, row, links, logo and icons), the phone's bar
// (the same probes at 390px: the logo, the compact icons, the menu's
// button, their places), the phone's menu opened by its button (home:menu*:
// the sheet, its rows, the faded logo and search, the X, the page held) and
// its motion, a dark nav (newsroom:nav*, newsroom:menu*), a clear one
// (product:nav), a touch screen. Golden, with their why: the letter-spacing,
// the images' widths (apple.com's SVG boxes), the X's lines (apple.com's end
// points, reference/icons/), and where ciderui's markup differs from
// apple.com's with the same picture.

import { examples, PADDING } from '../../test/harness.mjs';

const ex = examples(import.meta.url);
const NAV = ex('globalnav');
const dark = NAV.replace('class="cui-globalnav"', 'class="cui-globalnav cui-dark"');
const clear = (cls) => NAV.replace('class="cui-globalnav"', `class="${cls}" data-variant="clear"`);

const DESKTOP = [1600, 1200, 900];
const FLEX = ['display', 'flex-direction', 'flex-wrap', 'align-items', 'justify-content'];
const MOTION = ['transition-property', 'transition-duration', 'transition-timing-function'];
const INLINE = ['padding-left', 'padding-right'];
const LIST = '.cui-globalnav-list';
const LOGO = '.cui-globalnav-logo';
const BUTTON = '.cui-globalnav-menubutton';
// the first and second links (Store, Mac), the search and the bag
const LINK = `${LIST} > li:nth-child(3) > .cui-globalnav-link`;
const LINK2 = `${LIST} > li:nth-child(4) > .cui-globalnav-link`;
const SEARCH = `${LIST} > li:nth-child(7) > .cui-globalnav-icon`;
const BAG = `${LIST} > li:nth-child(8) > .cui-globalnav-icon`;
const LINES = [`${BUTTON} path:first-child`, `${BUTTON} path:last-child`];
// apple.com reads its menu's button at 900 and 390px alone
const TRIGGER_AT = { 1600: 900, 1200: 900 };
// apple.com's phone bar draws its ground on its content, ciderui's on the
// nav at every width: the same ground, apple.com's at 900px
const GROUND_AT = { 390: 900 };
const OPEN = { click: BUTTON, widths: [390], height: 900 };

export const cases = [
  {
    name: 'bar',
    html: NAV,
    widths: DESKTOP,
    checks: [
      { target: '.cui-globalnav', props: ['box-height', 'background-color', 'backdrop-filter'], apple: ['home', 'nav'] },
      { target: '.cui-globalnav-content', props: ['max-width', ...INLINE], apple: ['home', 'nav.content'] },
      { target: LIST, props: [...FLEX, 'margin-left', 'margin-right', 'box-height'], apple: ['home', 'nav.list'] },
      { target: LINK, props: ['font-size', 'line-height', 'font-weight', 'color', ...MOTION], apple: ['home', 'nav.link'] },
      { target: LINK, state: 'hover', props: ['color'], apple: ['home', 'nav.link', 'hover'] },
      { target: LINK, props: [...FLEX, ...INLINE, 'box-height', 'text-decoration-line'], apple: ['home', 'nav.link-box'] },
      { target: LINK, props: ['letter-spacing'], golden: 'the system font\'s fitted letter-spacing (tokens.css), not apple.com\'s' },
      { target: LOGO, props: [...FLEX, ...INLINE, 'box-height', 'box-left', 'color', ...MOTION], apple: ['home', 'nav.logo'] },
      { target: LOGO, state: 'hover', props: ['color'], apple: ['home', 'nav.logo', 'hover'] },
      { target: SEARCH, props: [...FLEX, ...INLINE, 'box-height', 'color', ...MOTION], apple: ['home', 'nav.search'] },
      { target: SEARCH, state: 'hover', props: ['color'], apple: ['home', 'nav.search', 'hover'] },
      { target: BAG, props: [...FLEX, ...INLINE, 'box-height', 'box-left', 'color'], apple: ['home', 'nav.bag'] },
      { target: BAG, state: 'hover', props: ['color'], apple: ['home', 'nav.bag', 'hover'] },
      { target: `${SEARCH} .cui-globalnav-image`, props: ['box-height'], apple: ['home', 'nav.search-image'] },
      { target: `${BAG} .cui-globalnav-image`, props: ['box-height'], apple: ['home', 'nav.bag-image'] },
      { target: `${SEARCH} .cui-globalnav-image`, props: ['box-width'], golden: 'apple.com\'s search image, 15 by 44 (home:nav.search-image)' },
      { target: `${BAG} .cui-globalnav-image`, props: ['box-width'], golden: 'apple.com\'s bag image, 14 by 44 (home:nav.bag-image)' },
      { target: `${SEARCH} .cui-globalnav-image-compact`, props: ['box-height'], golden: 'the phone\'s image, not drawn from 834px' },
      { target: BUTTON, props: ['box-height'], apple: ['home', 'nav.menutrigger'], at: TRIGGER_AT },
    ],
  },
  {
    name: 'phone',
    html: NAV,
    widths: [390],
    checks: [
      { target: '.cui-globalnav', props: ['box-height'], apple: ['home', 'nav'] },
      { target: '.cui-globalnav', props: ['background-color', 'backdrop-filter'], apple: ['home', 'nav'], at: GROUND_AT },
      { target: '.cui-globalnav-content', props: ['max-width', ...INLINE], apple: ['home', 'nav.content'] },
      { target: LIST, props: ['box-height'], apple: ['home', 'nav.list'] },
      { target: LOGO, props: [...FLEX, ...INLINE, 'box-height', 'box-left', 'color'], apple: ['home', 'nav.logo'] },
      { target: SEARCH, props: [...FLEX, ...INLINE, 'box-height', 'box-left', 'color'], apple: ['home', 'nav.search'] },
      { target: BAG, props: [...FLEX, ...INLINE, 'box-height', 'box-left', 'color'], apple: ['home', 'nav.bag'] },
      { target: SEARCH, props: ['box-width'], golden: 'apple.com\'s 48px square (home:nav.search at 390px)' },
      { target: `${SEARCH} .cui-globalnav-image-compact`, props: ['box-height'], apple: ['home', 'nav.search-compact'] },
      { target: `${BAG} .cui-globalnav-image-compact`, props: ['box-height'], apple: ['home', 'nav.bag-compact'] },
      { target: `${SEARCH} .cui-globalnav-image-compact`, props: ['box-width'], golden: 'apple.com\'s compact image, 17 by 48 (home:nav.search-compact)' },
      { target: `${SEARCH} .cui-globalnav-image`, props: ['box-height'], apple: ['home', 'nav.search-image'] },
      { target: BUTTON, props: [...FLEX, 'box-height', 'box-left', 'color', 'opacity', 'cursor', ...MOTION], apple: ['home', 'nav.menutrigger'] },
      { target: BUTTON, state: 'hover', props: ['opacity'], apple: ['home', 'nav.menutrigger', 'hover'] },
      { target: BUTTON, props: ['box-width'], golden: 'apple.com\'s 48px square (home:nav.menutrigger at 390px)' },
      { target: LINES[0], props: ['stroke-width'], apple: ['home', 'nav.menutrigger-bread'] },
      ...LINES.map(target => ({ target, props: ['box-x', 'box-y', 'box-width', 'box-height'], golden: 'apple.com\'s two lines, 2 to 16 across at 5 and 12 down in 18 by 18 (reference/icons/)' })),
      { target: LINK, props: ['box-height'], golden: 'to 833px the links are the menu\'s rows, not the bar\'s: apple.com\'s bar holds the logo, the search, the bag and the button' },
    ],
  },
  {
    name: 'menu',
    html: NAV,
    ...OPEN,
    reducedMotion: 'reduce',
    checks: [
      { target: '.cui-globalnav', props: ['box-height'], apple: ['home', 'menu.open'] },
      { target: LIST, props: ['background-color', 'box-top', 'box-left', 'box-height'], apple: ['home', 'menu.content'] },
      { target: LIST, props: ['padding-bottom'], apple: ['home', 'menu.list'] },
      { target: LINK, props: ['padding-top', 'box-top', 'box-height'], apple: ['home', 'menu.link-first'] },
      { target: LINK2, props: ['font-size', 'line-height', 'font-weight', 'color', 'display', ...PADDING, 'box-top', 'box-height'], apple: ['home', 'menu.link'] },
      { target: LINK2, state: 'hover', props: ['color'], apple: ['home', 'menu.link', 'hover'] },
      { target: LINK2, props: ['letter-spacing'], golden: 'the system font\'s fitted letter-spacing (tokens.css), not apple.com\'s' },
      // apple.com fades the links themselves; ciderui their items
      { target: `${LIST} > li:nth-child(1)`, props: ['opacity', 'visibility'], apple: ['home', 'menu.logo'] },
      { target: `${LIST} > li:nth-child(7)`, props: ['opacity', 'visibility'], apple: ['home', 'menu.search'] },
      { target: BUTTON, props: ['color', 'opacity'], apple: ['home', 'menu.trigger'] },
      { target: LINES[0], props: ['box-height'], apple: ['home', 'menu.trigger-bread-top'] },
      ...LINES.map(target => ({ target, props: ['box-x', 'box-y', 'box-width', 'box-height'], golden: 'apple.com\'s X: its lines from (3.5, 3.5) to (15, 15) and from (3.5, 15) to (15, 3.5), the end of its animation (reference/icons/)' })),
      { target: 'html', props: ['overflow-y'], apple: ['home', 'menu.root'] },
    ],
  },
  {
    name: 'menu-motion',
    html: NAV,
    ...OPEN,
    checks: [
      { target: `${LIST} > li:nth-child(1)`, props: [...MOTION, 'transition-delay'], apple: ['home', 'menu.logo'] },
      { target: `${LIST} > li:nth-child(3)`, props: ['transition-duration', 'transition-delay'], apple: ['home', 'menu.row'] },
      { target: `${LIST} > li:nth-child(4)`, props: ['transition-duration', 'transition-delay'], apple: ['home', 'menu.row-second'] },
      { target: `${LIST} > li:nth-child(3)`, props: ['transition-property', 'transition-timing-function'], golden: 'apple.com\'s rows also turn visible (its rows are hidden closed, ciderui\'s not displayed); their opacity and transform as apple.com\'s' },
      { target: LIST, props: MOTION, golden: 'apple.com\'s ground and height in 450ms (home:menu.open\'s animations; its computed transition names its ground alone, its height\'s declared in a state between)' },
      { target: LINES[0], props: MOTION, golden: 'apple.com\'s lines turn in 0.24s (their SVG animation, reference/icons/)' },
    ],
  },
  {
    name: 'menu-closed',
    html: NAV,
    ...OPEN,
    click: [BUTTON, BUTTON],
    reducedMotion: 'reduce',
    checks: [
      { target: LIST, props: ['box-height'], apple: ['home', 'nav.list'] },
      { target: LIST, props: ['position', 'background-color'], golden: 'closed, the bar\'s row again (reference/popover/)' },
      { target: 'html', props: ['overflow-y'], golden: 'the page scrolls again' },
      { target: LINK, props: ['box-height'], golden: 'the links are the menu\'s again' },
    ],
  },
  {
    // a menu opened on a phone and the window widened (a tablet turned) stays a sheet
    name: 'menu-widened',
    html: NAV,
    click: BUTTON,
    widths: [390, 900],
    reducedMotion: 'reduce',
    checks: [
      { target: LINK2, props: ['font-size', 'box-top'], golden: 'opened, the menu is a sheet at any width: its rows' },
      { target: BUTTON, props: ['box-height', 'opacity'], golden: 'opened, the menu is a sheet at any width: its button, the X' },
    ],
  },
  {
    name: 'dark',
    html: dark,
    checks: [
      { target: '.cui-globalnav', props: ['background-color', 'backdrop-filter'], apple: ['newsroom', 'nav'], at: GROUND_AT },
      { target: SEARCH, props: ['color'], apple: ['newsroom', 'nav.search'] },
      { target: SEARCH, state: 'hover', props: ['color'], apple: ['newsroom', 'nav.search', 'hover'] },
      { target: BAG, props: ['color'], apple: ['newsroom', 'nav.bag'] },
      { target: BUTTON, props: ['color', 'opacity'], apple: ['newsroom', 'nav.menutrigger'], at: TRIGGER_AT },
      { target: BUTTON, state: 'hover', props: ['opacity'], apple: ['newsroom', 'nav.menutrigger', 'hover'], at: TRIGGER_AT },
    ],
  },
  {
    name: 'dark-links',
    html: dark,
    widths: DESKTOP,
    checks: [
      { target: LINK, props: ['color'], apple: ['newsroom', 'nav.link'] },
      { target: LINK, state: 'hover', props: ['color'], apple: ['newsroom', 'nav.link', 'hover'] },
    ],
  },
  {
    name: 'dark-menu',
    html: dark,
    ...OPEN,
    reducedMotion: 'reduce',
    checks: [
      { target: LIST, props: ['background-color'], apple: ['newsroom', 'menu.content'] },
      { target: LINK2, props: ['color'], apple: ['newsroom', 'menu.link'] },
      { target: LINK2, state: 'hover', props: ['color'], apple: ['newsroom', 'menu.link', 'hover'] },
      { target: BUTTON, props: ['color', 'opacity'], apple: ['newsroom', 'menu.trigger'] },
    ],
  },
  {
    // the product page's: a clear bar over its dark hero
    name: 'clear',
    html: clear('cui-globalnav cui-dark'),
    widths: DESKTOP,
    checks: [
      { target: '.cui-globalnav', props: ['background-color', 'backdrop-filter'], apple: ['product', 'nav'] },
      { target: LINK, props: ['color'], apple: ['product', 'nav.link'] },
      { target: LINK, state: 'hover', props: ['color'], apple: ['product', 'nav.link', 'hover'] },
    ],
  },
  {
    name: 'clear-light',
    html: clear('cui-globalnav'),
    widths: [1200],
    checks: [
      { target: '.cui-globalnav', props: ['background-color', 'backdrop-filter'], golden: 'clear on a light page as on a dark one (product:nav)' },
      { target: LINK, props: ['color'], apple: ['home', 'nav.link'] },
    ],
  },
  {
    name: 'touch',
    html: NAV,
    touch: true,
    widths: [1200, 390],
    checks: [
      { target: SEARCH, state: 'hover', props: ['color'], rest: 'no hover on a touch screen (rule 5)' },
      { target: LOGO, state: 'hover', props: ['color'], rest: 'no hover on a touch screen (rule 5)' },
      { target: BUTTON, state: 'hover', props: ['opacity'], rest: 'no hover on a touch screen (rule 5)' },
    ],
  },
  {
    name: 'touch-links',
    html: NAV,
    touch: true,
    widths: [1200],
    checks: [{ target: LINK, state: 'hover', props: ['color'], rest: 'no hover on a touch screen (rule 5)' }],
  },
  {
    name: 'touch-menu',
    html: NAV,
    touch: true,
    ...OPEN,
    reducedMotion: 'reduce',
    checks: [{ target: LINK2, state: 'hover', props: ['color'], rest: 'no hover on a touch screen (rule 5)' }],
  },
];
