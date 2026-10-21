// The global nav against apple.com: the bar from 834px (home:nav*: its
// height, material, content, row, links, logo and icons), the phone's bar
// (the same probes at 390px: the logo, the compact icons, the menu's
// button, their places), the phone's menu opened by its button (home:menu*:
// the sheet, its rows, the faded logo and search, the X, the page held), its
// motion and its start, the keys (the focus on the X, Tab to the first link,
// Escape back to the button: reference/popover/), a dark nav
// (newsroom:nav*, newsroom:menu*), a clear one (product:nav), a touch
// screen. Its search opened by its button (search:*, search-dark:*): from
// 834px the flyout, its field, its quick links, the curtain, the bar over
// it; to 833px the sheet, the page held; its motion and its start; the focus
// on its field, Tab to the first link, Escape, the curtain, a click on the
// bar and its button again closing it (reference/search/); a dark nav's; a
// touch screen's. Golden, with their why: the letter-spacing, the images'
// widths (apple.com's SVG boxes), the X's lines (apple.com's end points,
// reference/icons/), the focus, the places apple.com's recorded with its
// locale banner above them, and where ciderui's markup differs from
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
const CLOSE = '.cui-globalnav-closebutton';
// the first and second links (Store, Mac), the search and the bag
const LINK = `${LIST} > li:nth-child(3) > .cui-globalnav-link`;
const LINK2 = `${LIST} > li:nth-child(4) > .cui-globalnav-link`;
const SEARCH = `${LIST} > li:nth-child(7) > .cui-globalnav-icon`;
const BAG = `${LIST} > li:nth-child(8) > .cui-globalnav-icon`;
const BARS = [`${BUTTON} path:first-child`, `${BUTTON} path:last-child`];
const X = [`${CLOSE} path:first-child`, `${CLOSE} path:last-child`];
// apple.com reads its menu's button at 900 and 390px alone
const TRIGGER_AT = { 1600: 900, 1200: 900 };
// apple.com's phone bar draws its ground on its content, ciderui's on the
// nav at every width: the same ground, apple.com's at 900px
const GROUND_AT = { 390: 900 };
const OPEN = { click: BUTTON, widths: [390], height: 900 };

// the search: the example with the button opening it, opened at the widths
// apple.com's was read at, a flyout from 834px and a sheet to 833px
const SEARCH_NAV = ex('search');
const SEARCH_DARK = SEARCH_NAV.replace('class="cui-globalnav"', 'class="cui-globalnav cui-dark"');
const SBUTTON = `${LIST} > li:nth-child(7) > .cui-globalnav-icon`;
const FLY = '.cui-globalnav-search';
const PANEL = '.cui-globalnav-searchpanel';
const FIELD = '.cui-globalnav-searchfield';
const INPUT = '.cui-globalnav-searchinput';
const HEADING = '.cui-globalnav-searchheading';
const QLINK = '.cui-globalnav-searchlinks > li:first-child > .cui-globalnav-searchlink';
const QLINK2 = '.cui-globalnav-searchlinks > li:nth-child(2) > .cui-globalnav-searchlink';
const CURTAIN = '.cui-globalnav-curtain';
const SCLOSE = `${FLY} .cui-globalnav-closebutton`;
const RESET = '.cui-globalnav-searchreset';
const MAGNIFIER = `${FIELD} > .cui-globalnav-image`;
// with text in its field, as apple.com's was read (search-text:*): typed,
// its value, not its value attribute, which is what a reset gives it back
const SEARCH_TEXT = `${SEARCH_NAV}<script>document.querySelector('.cui-globalnav-searchinput').value = 'iPhone';</script>`;
const WIDE = [1600, 1200, 900, 834];
const NARROW = [833, 390];
const SEARCHING = { click: SBUTTON, height: 900 };
// the search's motion held at its start, for a case that reads its start
const HOLD = '<style>.cui-globalnav-search, .cui-globalnav-search * { transition-delay: 10s !important; }</style>';
const TYPE = ['font-size', 'line-height', 'font-weight', 'color'];
// apple.com's places in the window are below its locale banner (70px from
// 834px, 135px to 833px): a place here is the same distance from the bar
const PLACE = 'the same distance from the bar as apple.com\'s (search:*, its locale banner above its nav)';

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
      { target: BARS[0], props: ['stroke-width'], apple: ['home', 'nav.menutrigger-bread'] },
      ...BARS.map(target => ({ target, props: ['box-x', 'box-y', 'box-width', 'box-height'], golden: 'apple.com\'s two lines, 2 to 16 across at 5 and 12 down in 18 by 18 (reference/icons/)' })),
      { target: `${LIST} > li:nth-child(2)`, props: ['box-height'], golden: 'the close button is the opened menu\'s' },
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
      { target: CLOSE, props: ['color', 'opacity', 'box-height', 'box-left', 'box-top'], apple: ['home', 'menu.trigger'] },
      { target: CLOSE, props: ['box-width'], golden: 'apple.com\'s 48px square, where the menu\'s button is (home:nav.menutrigger at 390px)' },
      { target: X[0], props: ['box-height'], apple: ['home', 'menu.trigger-bread-top'] },
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
      { target: `${LIST} > li:nth-child(3)`, props: ['transition-duration', 'transition-delay'], apple: ['home', 'menu.row'] },
      { target: `${LIST} > li:nth-child(4)`, props: ['transition-duration', 'transition-delay'], apple: ['home', 'menu.row-second'] },
      { target: `${LIST} > li:nth-child(3)`, props: ['transition-property', 'transition-timing-function'], golden: 'apple.com\'s rows also turn visible (its rows are hidden closed, ciderui\'s not displayed); their opacity and transform as apple.com\'s' },
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
      { target: CLOSE, props: ['box-height', 'opacity'], golden: 'opened, the menu is a sheet at any width: its close button' },
    ],
  },
  {
    // read at its start: the rows still transparent and 8px up (@starting-style
    // through the compiled stylesheet), apple.com's from (home:menu.open's
    // animations), 200ms before they move
    name: 'menu-start',
    html: NAV,
    ...OPEN,
    wait: 0,
    checks: [
      { target: `${LIST} > li:nth-child(3)`, props: ['opacity', 'transform'], golden: 'apple.com\'s rows start transparent, 8px up (home:menu.open\'s animations)' },
    ],
  },
  {
    name: 'menu-focus',
    html: NAV,
    ...OPEN,
    reducedMotion: 'reduce',
    checks: [
      { target: CLOSE, props: ['focused'], golden: 'opened, the focus is on the X (autofocus; reference/popover/)' },
    ],
  },
  {
    name: 'menu-tab',
    html: NAV,
    ...OPEN,
    press: 'Tab',
    reducedMotion: 'reduce',
    checks: [
      { target: LINK, props: ['focused'], golden: 'Tab goes from the X to the first link, not back to a button (reference/popover/)' },
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
      { target: CLOSE, props: ['color', 'opacity'], apple: ['newsroom', 'menu.trigger'] },
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
      { target: SEARCH, state: 'hover', props: ['color'], rest: 'no hover on a touch screen (rule 5), where apple.com\'s keeps it (touch-home:nav.search)' },
      { target: LOGO, state: 'hover', props: ['color'], rest: 'no hover on a touch screen (rule 5), where apple.com\'s keeps it (touch-home:nav.logo)' },
      { target: BAG, state: 'hover', props: ['color'], rest: 'no hover on a touch screen (rule 5), where apple.com\'s keeps it (touch-home:nav.bag)' },
      { target: BUTTON, state: 'hover', props: ['opacity'], rest: 'no hover on a touch screen (rule 5), where apple.com\'s keeps it (touch-home:nav.menutrigger)' },
    ],
  },
  {
    name: 'touch-links',
    html: NAV,
    touch: true,
    widths: [1200],
    checks: [{ target: LINK, state: 'hover', props: ['color'], rest: 'no hover on a touch screen (rule 5), where apple.com\'s keeps it (touch-home:nav.link)' }],
  },
  {
    name: 'search',
    html: SEARCH_NAV,
    ...SEARCHING,
    widths: WIDE,
    reducedMotion: 'reduce',
    checks: [
      // the bar over it, its button, the page not held
      { target: '.cui-globalnav', props: ['backdrop-filter'], apple: ['search', 'nav'] },
      { target: '.cui-globalnav', props: ['background-color', 'background-image'], golden: 'its ground over the flyout\'s, as apple.com\'s bar (search:nav, rgba(255,255,255,.8)) is over its flyout\'s #fafafc' },
      { target: SBUTTON, props: ['color'], apple: ['search', 'search'] },
      { target: 'html', props: ['overflow-y'], apple: ['search', 'root'] },
      // the flyout from the bar's bottom, its panel, the curtain behind it
      { target: FLY, props: ['box-top', 'box-left', 'box-height'], golden: 'from the bar\'s bottom (anchor positioning) to the window\'s, across it (reference/search/)' },
      { target: PANEL, props: ['background-color'], apple: ['search', 'flyout'] },
      { target: PANEL, props: ['box-height', 'box-top'], golden: 'apple.com\'s flyout below the bar, 344px (search:flyout.scroll), its content\'s' },
      { target: PANEL, props: ['padding-top', 'padding-bottom'], apple: ['search', 'flyout.content'] },
      { target: CURTAIN, props: ['background-color', 'backdrop-filter', 'opacity'], apple: ['search', 'curtain'] },
      { target: CURTAIN, props: ['box-top', 'box-height', 'hit'], golden: 'behind the panel from the bar\'s bottom to the window\'s, a button: what the pointer finds under the panel (apple.com\'s curtain is the window\'s, under its flyout)' },
      // the field: its row, its input, its placeholder, its magnifier
      { target: FIELD, props: ['height'], apple: ['search', 'field'] },
      { target: FIELD, props: ['box-top', 'box-left'], golden: PLACE },
      { target: INPUT, props: [...TYPE, 'height', 'background-color', 'border-top-width', 'border-top-style', 'border-top-left-radius', 'outline-style', 'caret-color'], apple: ['search', 'field.input'] },
      { target: INPUT, props: ['letter-spacing'], golden: 'the system font\'s fitted letter-spacing (tokens.css), not apple.com\'s' },
      { target: INPUT, props: ['padding-left', 'padding-top'], golden: 'its text 28px in from the field\'s start, where apple.com\'s is (its input 6px out, 34px in: search:field.input)' },
      { target: INPUT, pseudo: '::placeholder', props: ['color', 'font-weight', 'opacity'], apple: ['search', 'field.placeholder'] },
      { target: `${FIELD} .cui-globalnav-image`, props: ['box-height', 'color'], as: { 'box-height': 'box-height', color: 'fill' }, apple: ['search', 'field.submit-image'] },
      { target: `${FIELD} .cui-globalnav-image`, props: ['box-width', 'box-top', 'box-left'], golden: 'apple.com\'s 30 by 32 magnifier, 6px out from the field\'s start, 0.5px up (search:field.submit-image)' },
      { target: `${FIELD} .cui-globalnav-image-compact`, props: ['box-height'], golden: 'the phone\'s magnifier, not drawn from 834px' },
      { target: INPUT, props: ['focused'], golden: 'opened, the focus is on its field (autofocus; reference/search/)' },
      // empty, its clear button hidden
      { target: RESET, props: ['opacity', 'visibility'], apple: ['search', 'field.reset'] },
      // the quick links
      { target: HEADING, props: [...TYPE], apple: ['search', 'results.header'] },
      { target: HEADING, props: ['box-height'], apple: ['search', 'results.header'] },
      { target: HEADING, props: ['letter-spacing'], golden: 'the system font\'s fitted letter-spacing (tokens.css), not apple.com\'s' },
      { target: HEADING, props: ['box-top'], golden: `32px below the field, ${PLACE}` },
      { target: QLINK, props: [...TYPE, ...PADDING, 'display', 'align-items', 'text-decoration-line', 'background-color', 'border-top-left-radius', 'box-height'], apple: ['search', 'results.link'] },
      { target: QLINK, state: 'hover', props: ['color', 'background-color'], apple: ['search', 'results.link', 'hover'] },
      { target: QLINK, state: 'focus-visible', props: ['color', 'background-color'], apple: ['search', 'results.link', 'focus-visible'] },
      { target: QLINK, state: 'focus-visible', props: ['outline-style'], golden: 'the focus ring with its hovered look (base.css; apple.com forced into focus-visible draws none)' },
      { target: QLINK, props: ['letter-spacing'], golden: 'the system font\'s fitted letter-spacing (tokens.css), not apple.com\'s' },
      { target: QLINK, props: ['box-top', 'box-left'], golden: `8px below the heading, 8px out from the field's start, ${PLACE}` },
      { target: QLINK2, props: ['box-top'], golden: `4px below the first, ${PLACE}` },
      { target: `${QLINK} .cui-globalnav-image`, props: ['box-height', 'color'], as: { color: 'fill' }, apple: ['search', 'results.icon-image'] },
      { target: `${QLINK} .cui-globalnav-image`, state: 'hover', on: QLINK, props: ['color'], as: { color: 'fill' }, apple: ['search', 'results.icon-image', 'hover'] },
      { target: `${QLINK} .cui-globalnav-image`, props: ['margin-right'], apple: ['search', 'results.icon'] },
      { target: `${QLINK} .cui-globalnav-image`, props: ['box-width'], golden: 'apple.com\'s arrow, 9 by 16 (search:results.icon-image)' },
    ],
  },
  {
    name: 'search-phone',
    html: SEARCH_NAV,
    ...SEARCHING,
    widths: NARROW,
    reducedMotion: 'reduce',
    checks: [
      { target: FLY, props: ['background-color'], apple: ['search', 'content'] },
      { target: FLY, props: ['box-top', 'box-left', 'box-height'], golden: 'a sheet over the window, the bar too (apple.com\'s, search:content, from its nav\'s top below its locale banner)' },
      { target: 'html', props: ['overflow-y'], apple: ['search', 'root'] },
      { target: PANEL, props: ['padding-top', 'padding-left', 'padding-right'], apple: ['search', 'flyout.content'] },
      { target: PANEL, props: ['padding-bottom'], as: { 'padding-bottom': 'padding-bottom' }, apple: ['search', 'flyout.scroll'] },
      { target: CURTAIN, props: ['display'], golden: 'no curtain: the sheet is over the window (apple.com\'s, under it, is not seen)' },
      { target: FIELD, props: ['height'], apple: ['search', 'field'] },
      { target: FIELD, props: ['box-top', 'box-left'], golden: PLACE },
      { target: INPUT, props: [...TYPE, 'height', 'background-color', 'outline-style'], apple: ['search', 'field.input'] },
      { target: INPUT, props: ['letter-spacing', 'padding-left'], golden: 'the fitted letter-spacing; its text 34px in from the field\'s start, where apple.com\'s is (its input 8px out, 42px in)' },
      { target: `${FIELD} .cui-globalnav-image-compact`, props: ['box-height', 'color'], as: { color: 'fill' }, apple: ['search', 'field.submit-compact'] },
      { target: `${FIELD} .cui-globalnav-image-compact`, props: ['box-width', 'box-top', 'box-left'], golden: 'apple.com\'s 38 by 40 magnifier, 8px out and 8px up from the field (search:field.submit-compact)' },
      { target: `${FIELD} .cui-globalnav-image`, props: ['box-height'], golden: 'the wide magnifier, not drawn to 833px' },
      { target: INPUT, props: ['focused'], golden: 'opened, the focus is on its field (autofocus; reference/search/)' },
      { target: HEADING, props: [...TYPE, 'box-height'], apple: ['search', 'results.header'] },
      { target: HEADING, props: ['box-top'], golden: `49px below the field, ${PLACE}` },
      { target: QLINK, props: [...TYPE, ...PADDING, 'background-color', 'border-top-left-radius', 'box-height'], apple: ['search', 'results.link'] },
      { target: QLINK, state: 'hover', props: ['color', 'background-color'], apple: ['search', 'results.link', 'hover'] },
      { target: QLINK, props: ['box-top', 'box-left'], golden: `9px below the heading, 12px out, ${PLACE}` },
      { target: QLINK2, props: ['box-top'], golden: `6px below the first, ${PLACE}` },
      { target: `${QLINK} .cui-globalnav-image-compact`, props: ['box-height', 'color'], as: { color: 'fill' }, apple: ['search', 'results.icon-compact'] },
      { target: `${QLINK} .cui-globalnav-image-compact`, props: ['margin-right'], apple: ['search', 'results.icon'] },
      { target: `${QLINK} .cui-globalnav-image-compact`, props: ['box-width'], golden: 'apple.com\'s arrow, 13 by 25 (search:results.icon-compact)' },
      { target: SCLOSE, props: ['box-top', 'box-left', 'box-width', 'box-height', 'color', 'opacity'], golden: 'apple.com\'s X where the menu\'s button is, its Close (scouted: the menu\'s button turned, 48px at the right)' },
      ...[`${SCLOSE} path:first-child`, `${SCLOSE} path:last-child`].map(target => ({ target, props: ['box-x', 'box-y', 'box-width', 'box-height'], golden: 'the menu\'s X: its lines from (3.5, 3.5) to (15, 15) and from (3.5, 15) to (15, 3.5) (reference/icons/)' })),
    ],
  },
  {
    name: 'search-motion',
    html: SEARCH_NAV,
    ...SEARCHING,
    widths: [1600, 390],
    checks: [
      { target: FIELD, props: [...MOTION, 'transition-delay'], golden: 'apple.com\'s field in 240ms from 200ms (search:field.wrapper\'s keyframes, an animation there)' },
      { target: HEADING, props: ['transition-duration', 'transition-delay'], golden: 'apple.com\'s heading in 320ms from 200ms (search:results.header)' },
      { target: '.cui-globalnav-searchlinks > li:first-child', props: ['transition-duration', 'transition-delay'], golden: 'apple.com\'s rows 20ms after the one above (search:nav\'s animations)' },
      { target: '.cui-globalnav-searchlinks > li:nth-child(2)', props: ['transition-delay'], golden: 'apple.com\'s rows 20ms after the one above (search:nav\'s animations)' },
      { target: FLY, props: MOTION, golden: 'the sheet\'s height and ground in 450ms to 833px (search:nav.phone\'s animations on its content); no motion of its own from 834px' },
      { target: PANEL, props: MOTION, golden: 'from 834px revealed down and from the bar\'s ground in 0.24s, apple.com\'s flyout\'s height and ground (search:nav\'s animations)' },
      { target: CURTAIN, props: MOTION, golden: 'apple.com\'s curtain in 0.32s (search:curtain), from 834px' },
      { target: QLINK, props: MOTION, golden: 'apple.com\'s color and ground in 0.02s (search:results.link; its fill too, where the arrow here is drawn in currentColor)' },
    ],
  },
  {
    // read at its start: apple.com's from (search:nav's animations); held
    // there, as what moves can move before it is read (the panel and the
    // curtain at once, the field and its rows after 200ms, which a busy run
    // has passed)
    name: 'search-start',
    html: `${HOLD}${SEARCH_NAV}`,
    ...SEARCHING,
    widths: [1600],
    wait: 0,
    checks: [
      { target: FIELD, props: ['opacity', 'transform'], golden: 'the field from transparent, 4px up (search:field.wrapper\'s keyframes)' },
      { target: '.cui-globalnav-searchlinks > li:first-child', props: ['opacity', 'transform'], golden: 'each row from transparent, 4px up (search:results.header\'s keyframes)' },
      { target: PANEL, props: ['clip-path', 'background-color'], golden: 'the panel hidden down from the bar, on its ground: apple.com\'s flyout 44px tall, the bar\'s ground (search:nav\'s animations)' },
      { target: CURTAIN, props: ['opacity'], golden: 'the curtain from transparent (search:curtain\'s transition)' },
    ],
  },
  {
    name: 'search-start-phone',
    html: `${HOLD}${SEARCH_NAV}`,
    ...SEARCHING,
    widths: [390],
    wait: 0,
    checks: [
      { target: FLY, props: ['height', 'background-color'], golden: 'the sheet from the bar\'s height and ground (search:nav.phone\'s animations on its content)' },
      { target: FIELD, props: ['opacity', 'transform'], golden: 'the field from transparent, 8px up (search:nav.phone\'s animations)' },
    ],
  },
  {
    name: 'search-tab',
    html: SEARCH_NAV,
    ...SEARCHING,
    widths: [1600],
    press: 'Tab',
    reducedMotion: 'reduce',
    checks: [{ target: QLINK, props: ['focused'], golden: 'Tab goes from its field to the first quick link' }],
  },
  {
    name: 'search-escape',
    html: SEARCH_NAV,
    ...SEARCHING,
    widths: [1600, 390],
    press: 'Escape',
    reducedMotion: 'reduce',
    checks: [
      { target: FLY, props: ['display'], golden: 'Escape closes it (a text field: reference/search/)' },
      { target: SBUTTON, props: ['focused'], golden: 'the focus back on its button (reference/search/)' },
      { target: 'html', props: ['overflow-y'], golden: 'the page scrolls again' },
    ],
  },
  {
    name: 'search-curtain',
    html: SEARCH_NAV,
    click: [SBUTTON, CURTAIN],
    height: 900,
    widths: [1600],
    reducedMotion: 'reduce',
    checks: [
      { target: FLY, props: ['display'], golden: 'a click on the curtain closes it (reference/search/)' },
      { target: '.cui-globalnav', props: ['background-color', 'backdrop-filter'], apple: ['home', 'nav'] },
    ],
  },
  {
    // its links to fragments here, as a click on the bar's middle is on one
    name: 'search-bar',
    html: SEARCH_NAV.replaceAll('href="/', 'href="#/'),
    click: [SBUTTON, '.cui-globalnav'],
    height: 900,
    widths: [1600],
    reducedMotion: 'reduce',
    checks: [{ target: FLY, props: ['display'], golden: 'a click on the bar, outside it, closes it (reference/search/)' }],
  },
  {
    name: 'search-again',
    html: SEARCH_NAV,
    click: [SBUTTON, SBUTTON],
    height: 900,
    widths: [1600],
    reducedMotion: 'reduce',
    checks: [{ target: FLY, props: ['display'], golden: 'its button closes it again (reference/search/)' }],
  },
  {
    name: 'search-close',
    html: SEARCH_NAV,
    click: [SBUTTON, SCLOSE],
    height: 900,
    widths: [390],
    reducedMotion: 'reduce',
    checks: [
      { target: FLY, props: ['display'], golden: 'the sheet\'s X closes it (reference/search/)' },
      { target: 'html', props: ['overflow-y'], golden: 'the page scrolls again' },
    ],
  },
  {
    name: 'search-dark',
    html: SEARCH_DARK,
    ...SEARCHING,
    widths: [1600, 390],
    reducedMotion: 'reduce',
    checks: [
      { target: INPUT, props: ['color', 'caret-color'], apple: ['search-dark', 'field.input'] },
      { target: INPUT, pseudo: '::placeholder', props: ['color'], apple: ['search-dark', 'field.placeholder'] },
      { target: HEADING, props: ['color'], apple: ['search-dark', 'results.header'] },
      { target: QLINK, props: ['color', 'background-color'], apple: ['search-dark', 'results.link'] },
      { target: QLINK, state: 'hover', props: ['color', 'background-color'], apple: ['search-dark', 'results.link', 'hover'] },
    ],
  },
  {
    name: 'search-dark-flyout',
    html: SEARCH_DARK,
    ...SEARCHING,
    widths: [1600],
    reducedMotion: 'reduce',
    checks: [
      { target: PANEL, props: ['background-color'], apple: ['search-dark', 'flyout'] },
      { target: CURTAIN, props: ['background-color', 'backdrop-filter'], apple: ['search-dark', 'curtain'] },
      { target: `${FIELD} .cui-globalnav-image`, props: ['color'], as: { color: 'fill' }, apple: ['search-dark', 'field.submit-image'] },
      { target: `${QLINK} .cui-globalnav-image`, props: ['color'], as: { color: 'fill' }, apple: ['search-dark', 'results.icon-image'] },
      { target: `${QLINK} .cui-globalnav-image`, state: 'hover', on: QLINK, props: ['color'], as: { color: 'fill' }, apple: ['search-dark', 'results.icon-image', 'hover'] },
    ],
  },
  {
    name: 'search-dark-sheet',
    html: SEARCH_DARK,
    ...SEARCHING,
    widths: [390],
    reducedMotion: 'reduce',
    checks: [{ target: FLY, props: ['background-color'], apple: ['search-dark', 'content'] }],
  },
  {
    // text in its field: the clear button shown, the magnifier darker
    name: 'search-text',
    html: SEARCH_TEXT,
    ...SEARCHING,
    widths: [...WIDE, ...NARROW],
    checks: [
      { target: INPUT, props: ['color'], apple: ['search-text', 'input'] },
      { target: INPUT, props: ['padding-right'], golden: 'its text ending 31px before the field\'s end (25px from 834px), where apple.com\'s does, 3px before its clear button (its input 7px out and 38px in; 8px out, 34px in: search-text:input)' },
      { target: RESET, props: ['opacity', 'visibility', 'display', 'align-items', 'justify-content', 'cursor', 'background-color', 'border-top-width', 'transition-duration', 'transition-timing-function'], apple: ['search-text', 'reset'] },
      { target: RESET, props: ['color'], as: { color: 'fill' }, apple: ['search-text', 'reset'] },
      { target: RESET, props: ['transition-property'], golden: 'apple.com\'s opacity and visibility (search-text:reset), and its mark\'s color where apple.com\'s fill moves' },
      { target: RESET, props: ['box-width', 'box-height', 'box-top', 'box-right'], golden: 'apple.com\'s 30 by 36 button from 834px, 32 by 38 below, its mark centered on the field and ending at its end (1px in below 834px), the button 8px (7px) past it (search-text:reset, reset.image)' },
      { target: `${RESET} .cui-globalnav-image`, props: ['box-width', 'box-height', 'display'], golden: 'apple.com\'s 14px mark from 834px (search-text:reset.image), not drawn below' },
      { target: `${RESET} .cui-globalnav-image-compact`, props: ['box-width', 'box-height', 'display'], golden: 'apple.com\'s 16px mark to 833px (search-text:reset.image-compact), not drawn from 834px' },
      { target: RESET, state: 'hover', props: ['color'], as: { color: 'fill' }, apple: ['search-text', 'reset.states', 'hover'] },
      { target: RESET, state: 'focus-visible', props: ['color'], as: { color: 'fill' }, apple: ['search-text', 'reset.states', 'focus-visible'] },
      { target: RESET, state: 'focus-visible', props: ['outline-style', 'outline-color', 'outline-width'], apple: ['search-text', 'reset.states', 'focus-visible'] },
      { target: RESET, state: 'focus-visible', props: ['outline-offset'], golden: 'the kit\'s one ring, 3px out (base.css), where apple.com\'s is 1px out (search-text:reset.states)' },
      { target: MAGNIFIER, props: ['color'], as: { color: 'fill' }, apple: ['search-text', 'submit'] },
      { target: MAGNIFIER, props: ['transition-duration', 'transition-timing-function'], apple: ['search-text', 'submit'] },
      { target: MAGNIFIER, state: 'hover', on: FIELD, props: ['color'], golden: 'an image, not a button: no hover (apple.com\'s magnifier is a submit button while the field holds text, #000 hovered: search-text:submit.states); Return sends the form' },
    ],
  },
  {
    // its clear button clicked: the field empty, the button hidden again
    name: 'search-clear',
    html: SEARCH_TEXT,
    click: [SBUTTON, RESET],
    height: 900,
    widths: [1600],
    checks: [
      { target: RESET, props: ['opacity', 'visibility', 'transition-duration', 'transition-timing-function'], apple: ['search', 'field.reset'] },
      { target: MAGNIFIER, props: ['color'], as: { color: 'fill' }, apple: ['search', 'field.submit-image'] },
      { target: FLY, props: ['display'], golden: 'still open: the clear button is in it (reference/search/)' },
      { target: INPUT, props: ['focused'], golden: 'the focus on the page, the button it was on hidden (Safari, Chrome, Chromium: reference/search/), where apple.com\'s script puts it back in the field; Shift+Tab takes it there' },
    ],
  },
  {
    // a field the page fills (value=): the reset gives that back
    name: 'search-clear-value',
    html: SEARCH_NAV.replace('<input class="cui-globalnav-searchinput"', '<input class="cui-globalnav-searchinput" value="iPhone"'),
    click: [SBUTTON, RESET],
    height: 900,
    widths: [1600],
    checks: [
      { target: RESET, props: ['visibility'], golden: 'a reset button gives the field back its value attribute, not an empty one: a page that fills the field with its query has the clear button bring the query back, and leaves it out' },
    ],
  },
  {
    name: 'search-clear-phone',
    html: SEARCH_TEXT,
    click: [SBUTTON, RESET],
    height: 900,
    widths: [390],
    checks: [
      { target: RESET, props: ['opacity', 'visibility'], apple: ['search', 'field.reset'] },
      { target: `${FIELD} > .cui-globalnav-image-compact`, props: ['color'], as: { color: 'fill' }, apple: ['search', 'field.submit-compact'] },
    ],
  },
  {
    name: 'search-touch',
    html: SEARCH_TEXT,
    ...SEARCHING,
    touch: true,
    widths: [1200, 390],
    reducedMotion: 'reduce',
    checks: [
      { target: QLINK, state: 'hover', props: ['color', 'background-color'], rest: 'no hover on a touch screen (rule 5), where apple.com\'s keeps it (touch-search:results.link)' },
      { target: RESET, state: 'hover', props: ['color'], rest: 'nor its clear button (rule 5)' },
    ],
  },
  {
    name: 'touch-menu',
    html: NAV,
    touch: true,
    ...OPEN,
    reducedMotion: 'reduce',
    checks: [{ target: LINK2, state: 'hover', props: ['color'], rest: 'no hover on a touch screen (rule 5), where apple.com\'s keeps it (touch-home:menu.link)' }],
  },
];
