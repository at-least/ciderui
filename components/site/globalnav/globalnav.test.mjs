import { examples, MOTION, PADDING, SIZE } from '../../../test/harness.mjs';

const ex = examples(import.meta.url);
const NAV = ex('globalnav');
const dark = NAV.replace('class="cui-globalnav"', 'class="cui-globalnav cui-dark"');

const DESKTOP = [1600, 1200, 900];
const FLEX = ['display', 'flex-direction', 'flex-wrap', 'align-items', 'justify-content'];
const INLINE = ['padding-left', 'padding-right'];
const LIST = '.cui-globalnav-list';
const LOGO = '.cui-globalnav-logo';
const BUTTON = '.cui-globalnav-menubutton';
const CLOSE = '.cui-globalnav-closebutton';
const LINK = `${LIST} > li:nth-child(3) > .cui-globalnav-link`;
const LINK2 = `${LIST} > li:nth-child(4) > .cui-globalnav-link`;
const BARS = [`${BUTTON} path:first-child`, `${BUTTON} path:last-child`];
const X = [`${CLOSE} path:first-child`, `${CLOSE} path:last-child`];
// @external: apple.com reads its menu's button at 900 and 390px alone
const TRIGGER_AT = { 1600: 900, 1200: 900 };
// @external: apple.com's phone bar draws its ground on its content, ciderui's on the
// nav at every width: the same ground, apple.com's at 900px
const GROUND_AT = { 390: 900 };
const OPEN = { click: BUTTON, widths: [390], height: 900 };
const ROWS = 'shown with the sheet, where apple.com\'s fade in 240ms from 200ms (home:menu.row, home:menu.open\'s animations): a phone\'s menu is read at once (DESIGN.md decision 25)';




const POST_NAV = ex('post');
const POST = `${LIST} > li:nth-child(6) > .cui-globalnav-link`;
// @external: what a <button> brings of its own, a link not
const BUTTONS_OWN = ['font-family', 'font-style', 'text-align', 'vertical-align', 'border-top-width', 'border-left-width', 'background-color', 'cursor', 'margin-top', 'margin-left', 'box-sizing', 'box-top', 'opacity'];
const AS_LINK = 'a POST item (sign out) is a <button>, drawn as the link: each value equal to the link\'s beside it';

export const cases = [
  {
    name: 'post-bar',
    html: POST_NAV,
    widths: DESKTOP,
    checks: [
      { target: POST, props: ['font-size', 'line-height', 'font-weight', 'color'], apple: ['home', 'nav.link'] },
      { target: POST, state: 'hover', props: ['color'], apple: ['home', 'nav.link', 'hover'] },
      { target: POST, props: [...FLEX, ...INLINE, 'box-height', 'text-decoration-line'], apple: ['home', 'nav.link-box'] },
      { target: LINK, props: BUTTONS_OWN, golden: AS_LINK },
      { target: POST, props: BUTTONS_OWN, golden: AS_LINK },
    ],
  },
  {
    name: 'post-phone',
    html: POST_NAV,
    widths: [390],
    checks: [
      { target: POST, props: ['box-height'], golden: 'to 833px the links are the menu\'s rows, not the bar\'s: apple.com\'s bar holds the logo, the search, the bag and the button' },
    ],
  },
  {
    name: 'post-menu',
    html: POST_NAV,
    ...OPEN,
    reducedMotion: 'reduce',
    checks: [
      { target: POST, props: ['font-size', 'line-height', 'font-weight', 'color', 'display', ...PADDING, 'box-height'], apple: ['home', 'menu.link'] },
      { target: POST, state: 'hover', props: ['color'], apple: ['home', 'menu.link', 'hover'] },
      { target: LINK2, props: [...BUTTONS_OWN.filter(p => p !== 'box-top'), 'box-left', 'box-width'], golden: AS_LINK },
      { target: POST, props: [...BUTTONS_OWN.filter(p => p !== 'box-top'), 'box-left', 'box-width'], golden: AS_LINK },
    ],
  },
  {
    name: 'bar',
    html: NAV,
    widths: DESKTOP,
    checks: [
      { target: '.cui-globalnav', props: ['box-height', 'background-color', 'backdrop-filter'], apple: ['home', 'nav'] },
      { target: '.cui-globalnav-content', props: ['max-width', ...INLINE], apple: ['home', 'nav.content'] },
      { target: LIST, props: [...FLEX, 'margin-left', 'margin-right', 'box-height'], apple: ['home', 'nav.list'] },
      { target: LINK, props: ['font-size', 'line-height', 'font-weight', 'color'], apple: ['home', 'nav.link'] },
      { target: LINK, state: 'hover', props: ['color'], apple: ['home', 'nav.link', 'hover'] },
      { target: LINK, props: [...FLEX, ...INLINE, 'box-height', 'text-decoration-line'], apple: ['home', 'nav.link-box'] },
      { target: LOGO, props: [...FLEX, ...INLINE, 'box-height', 'box-left', 'color'], apple: ['home', 'nav.logo'] },
      { target: LOGO, state: 'hover', props: ['color'], apple: ['home', 'nav.logo', 'hover'] },
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
      { target: BUTTON, props: [...FLEX, 'box-height', 'box-left', 'color', 'opacity', 'cursor'], apple: ['home', 'nav.menutrigger'] },
      { target: BUTTON, state: 'hover', props: ['opacity'], apple: ['home', 'nav.menutrigger', 'hover'] },
      { target: BUTTON, props: ['box-width'], golden: 'apple.com\'s 48px square (home:nav.menutrigger at 390px)' },
      { target: BARS[0], props: ['stroke-width'], apple: ['home', 'nav.menutrigger-bread'] },
      ...BARS.map(target => ({ target, props: ['box-x', 'box-y', 'box-width', 'box-height'], golden: 'apple.com\'s two lines, 2 to 16 across at 5 and 12 down in 18 by 18 (reference/icons/)' })),
      { target: `${LIST} > li:nth-child(2)`, props: ['box-height'], golden: 'the close button is the opened menu\'s' },
      { target: LINK, props: ['box-height'], golden: 'to 833px the links are the menu\'s rows, not the bar\'s: apple.com\'s bar holds the logo, the search, the bag and the button' },
    ],
  },
  {
    name: 'select',
    html: NAV,
    widths: [1600, 390],
    checks: [
      { target: '.cui-globalnav', props: ['user-select'], apple: ['nav-select', 'nav'] },
      { target: LIST, props: ['user-select'], apple: ['nav-select', 'list'] },
      { target: LOGO, props: ['user-select'], apple: ['nav-select', 'logo'] },
      { target: BUTTON, props: ['user-select'], apple: ['nav-select', 'menutrigger'], at: { 1600: 390 } },
    ],
  },
  {
    name: 'menu-zh-hant',
    html: `<div lang="zh-Hant">${NAV}</div>`,
    ...OPEN,
    reducedMotion: 'reduce',
    checks: [{ target: LINK2, props: ['line-height', 'letter-spacing'], apple: ['home-tw', 'menu.link'] }],
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
      // @external: apple.com fades the links themselves; ciderui their items
      { target: `${LIST} > li:nth-child(1)`, props: ['opacity', 'visibility'], apple: ['home', 'menu.logo'] },
      { target: CLOSE, props: ['color', 'opacity', 'box-height', 'box-left', 'box-top'], apple: ['home', 'menu.trigger'] },
      { target: CLOSE, props: ['box-width'], golden: 'apple.com\'s 48px square, where the menu\'s button is (home:nav.menutrigger at 390px)' },
      ...X.map(target => ({ target, props: ['box-x', 'box-y', 'box-width', 'box-height'], golden: 'apple.com\'s X: its lines from (3.5, 3.5) to (15, 15) and from (3.5, 15) to (15, 3.5), the end of its animation (reference/icons/)' })),
      { target: 'html', props: ['overflow-y'], apple: ['home', 'menu.root'] },
    ],
  },
  {
    name: 'menu-motion',
    html: NAV,
    ...OPEN,
    checks: [
      { target: `${LIST} > li:nth-child(1)`, props: [...MOTION, 'transition-delay'], apple: ['home', 'menu.logo'] },
      { target: `${LIST} > li:nth-child(3)`, props: ['transition-duration', 'animation-name', 'animation-duration', 'animation-delay'], golden: ROWS },
      { target: LIST, props: MOTION, golden: 'apple.com\'s ground and height in 450ms (home:menu.open\'s animations; its computed transition names its ground alone, its height\'s declared in a state between)' },
      { target: X[0], props: MOTION, golden: 'apple.com\'s lines turn in 0.24s (their SVG animation, reference/icons/)' },
    ],
  },
  {
    name: 'menu-closed',
    html: NAV,
    ...OPEN,
    click: [BUTTON, CLOSE],
    reducedMotion: 'reduce',
    checks: [
      { target: LIST, props: ['box-height'], apple: ['home', 'nav.list'] },
      { target: LIST, props: ['position', 'background-color'], golden: 'closed, the bar\'s row again (reference/popover/)' },
      { target: 'html', props: ['overflow-y'], golden: 'the page scrolls again' },
      { target: LINK, props: ['box-height'], golden: 'the links are the menu\'s again' },
      { target: BUTTON, props: ['focused'], golden: 'the X closes it, the focus back on the menu\'s button (reference/popover/)' },
    ],
  },
  {
    name: 'menu-closed-at-once',
    html: NAV,
    ...OPEN,
    click: [BUTTON, CLOSE],
    wait: 0,
    checks: [
      { target: LIST, props: ['position', 'box-height'], golden: 'closed at once, out of the top layer (apple.com\'s runs its motion back)' },
    ],
  },
  {
    name: 'menu-tab-out',
    html: `${NAV}<a class="after" href="#after">After</a>`,
    ...OPEN,
    press: ['Tab', 'Tab', 'Tab', 'Tab', 'Tab'],
    reducedMotion: 'reduce',
    checks: [
      { target: '.after', props: ['focused'], golden: 'Tab past its last link leaves it for the page: no script holds the focus in it (apple.com\'s does)' },
    ],
  },
  {
    name: 'menu-widened',
    html: NAV,
    click: BUTTON,
    widths: [390, 900, 1600],
    reducedMotion: 'reduce',
    checks: [
      { target: LIST, props: ['width', 'left', 'right'], golden: 'the menu fills the window: the browser places a popover at its content\'s width' },
      { target: LINK2, props: ['font-size', 'box-top'], golden: 'opened, the menu is a sheet at any width: its rows' },
      { target: CLOSE, props: ['box-height', 'opacity'], golden: 'opened, the menu is a sheet at any width: its close button' },
    ],
  },
  {
    name: 'menu-start',
    html: NAV,
    ...OPEN,
    wait: 0,
    checks: [
      { target: `${LIST} > li:nth-child(3)`, props: ['opacity', 'transform'], golden: ROWS },
      { target: LIST, props: ['box-height', 'background-color'], golden: 'a departure at the reveal\'s start: apple.com\'s sheet grows from the bar\'s 48px and its white ground (home:menu.open\'s animations); ciderui\'s grows from its content\'s height, transparent — this Chromium ignores a popover\'s @starting-style where the element renders closed (the bar\'s logo lives in its list), and the bar is not restructured for the first frames of a reveal (measured in the browser, 2026-11-07)' },
    ],
  },
  {
    name: 'menu-focus',
    html: NAV,
    ...OPEN,
    reducedMotion: 'reduce',
    checks: [
      { target: CLOSE, props: ['focused'], golden: 'opened, the focus is on the X (autofocus; reference/popover/): the menu\'s button beside the list, as a popover opened by a button inside it holds the focus on that button' },
    ],
  },
  {
    name: 'menu-tab',
    html: NAV,
    ...OPEN,
    press: 'Tab',
    reducedMotion: 'reduce',
    checks: [
      { target: LINK, props: ['focused'], golden: 'Tab goes from the X to the first link, not back to a button (reference/popover/): the X second in the list, after the logo' },
    ],
  },
  {
    name: 'menu-escape',
    html: NAV,
    ...OPEN,
    press: 'Escape',
    reducedMotion: 'reduce',
    checks: [
      { target: LIST, props: ['position'], golden: 'Escape closes it' },
      { target: BUTTON, props: ['focused'], golden: 'the focus back on the menu\'s button (reference/popover/)' },
    ],
  },
  {
    name: 'dark',
    html: dark,
    checks: [
      { target: '.cui-globalnav', props: ['background-color', 'backdrop-filter'], apple: ['newsroom', 'nav'], at: GROUND_AT },
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
      { target: CLOSE, props: ['color', 'opacity'], apple: ['newsroom', 'menu.trigger'] },
    ],
  },
  {
    name: 'sticky',
    html: NAV.replace('class="cui-globalnav"', 'class="cui-globalnav sticky top-0 z-0"'),
    widths: [1200],
    checks: [
      { target: '.cui-globalnav', props: ['position', 'z-index'], golden: 'a site\'s sticky and z-index, utilities, win over its position and 9999: its place the page\'s' },
    ],
  },

];
