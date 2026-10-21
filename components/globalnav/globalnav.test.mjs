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
// touch screen's. A link's flyout (flyout:*, flyout-dark:*, flyout-phone:*,
// tap-flyout:*): from 834px shown by the pointer resting on its link (its
// bar, its panel, its groups, headings, rows and links, the curtain, the
// link's button and its chevron), its motion, the pointer off the bar and on
// to another link, opened by its button (the keys), on a touch screen its
// button over the link; to 833px over the open menu, its back button and X,
// a row hovered; a dark nav's (reference/flyout/). What can be selected
// (nav-select:*): none of it, its search's field too. Golden, with their why: the letter-spacing, the images'
// widths (apple.com's SVG boxes), the X's lines (apple.com's end points,
// reference/icons/), the focus, the places apple.com's recorded with its
// locale banner above them, and where ciderui's markup differs from
// apple.com's with the same picture.

import { examples, OUTLINE, PADDING } from '../../test/harness.mjs';

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
// showing suggestions, as a site's script leaves it once ma is typed
// (search-suggestions:*): the field's value, its two sections, each row's
// typed part marked
const SUGGESTING = `${ex('suggestions')}<script>document.querySelector('.cui-globalnav-searchinput').value = 'ma';</script>`;
const SUGGESTING_DARK = SUGGESTING.replace('class="cui-globalnav"', 'class="cui-globalnav cui-dark"');
const LHEAD = `${PANEL} > .cui-globalnav-searchheading:first-of-type`;
const SHEAD = `${PANEL} > .cui-globalnav-searchheading:last-of-type`;
const LROW = `${PANEL} > .cui-globalnav-searchlinks:first-of-type > li:first-child > .cui-globalnav-searchlink`;
const LROW2 = `${PANEL} > .cui-globalnav-searchlinks:first-of-type > li:last-child > .cui-globalnav-searchlink`;
const SROW = `${PANEL} > .cui-globalnav-searchlinks:last-of-type > li:first-child > .cui-globalnav-searchlink`;
const MARK = `${LROW} mark`;
// a row the site's script writes once the search is open, after a frame
const LATER = `<script>document.querySelector('.cui-globalnav-search').addEventListener('toggle', (e) => { if (e.newState === 'open') requestAnimationFrame(() => requestAnimationFrame(() => { const list = document.querySelector('.cui-globalnav-searchlinks'); list.append(list.firstElementChild.cloneNode(true)); })); });</script>`;
const WIDE = [1600, 1200, 900, 834];
const NARROW = [833, 390];
const SEARCHING = { click: SBUTTON, height: 900 };
// the search's motion held at its start, for a case that reads its start
const HOLD = '<style>.cui-globalnav-search, .cui-globalnav-search * { transition-delay: 10s !important; }</style>';
const TYPE = ['font-size', 'line-height', 'font-weight', 'color'];

// a link's flyout: the example with Store's and Mac's
const FLYOUT_NAV = ex('flyout');
const FLYOUT_DARK = FLYOUT_NAV.replace('class="cui-globalnav"', 'class="cui-globalnav cui-dark"');
const STORE_ITEM = `${LIST} > li:nth-child(3)`;
const MAC_ITEM = `${LIST} > li:nth-child(4)`;
const STORE_LINK = `${STORE_ITEM} > .cui-globalnav-link`;
const MAC_LINK = `${MAC_ITEM} > .cui-globalnav-link`;
const IPAD_LINK = `${LIST} > li:nth-child(5) > .cui-globalnav-link`;
const MAC_BUTTON = `${MAC_ITEM} > .cui-globalnav-flyoutbutton`;
const STORE_FLY = `${STORE_ITEM} > .cui-globalnav-flyout`;
const MAC_FLY = `${MAC_ITEM} > .cui-globalnav-flyout`;
const FPANEL = `${MAC_FLY} > .cui-globalnav-flyoutpanel`;
const [GROUP1, GROUP2, GROUP3] = [1, 2, 3].map(n => `${FPANEL} > .cui-globalnav-flyoutgroup:nth-of-type(${n})`);
const [FHEAD1, FHEAD2, FHEAD3] = [GROUP1, GROUP2, GROUP3].map(g => `${g} > .cui-globalnav-flyoutheading`);
const FLIST2 = `${GROUP2} > .cui-globalnav-flyoutlinks`;
// the first group's first two rows (large), its fifth (small), the second's first two
const ROW_LG = `${GROUP1} > .cui-globalnav-flyoutlinks > li:nth-child(1)`;
const ROW_LG2 = `${GROUP1} > .cui-globalnav-flyoutlinks > li:nth-child(2)`;
const ROW_AFTER = `${GROUP1} > .cui-globalnav-flyoutlinks > li:nth-child(5)`;
const ROW = `${FLIST2} > li:nth-child(1)`;
const ROW2 = `${FLIST2} > li:nth-child(2)`;
const [LINK_LG, LINK_AFTER, FLINK] = [ROW_LG, ROW_AFTER, ROW].map(r => `${r} > .cui-globalnav-flyoutlink`);
const MAC_CURTAIN = `${MAC_FLY} > .cui-globalnav-curtain`;
const BACK = `${FPANEL} > .cui-globalnav-backbutton`;
const FCLOSE = `${FPANEL} > .cui-globalnav-closebutton`;
const CHEVRON = `${MAC_BUTTON} > .cui-globalnav-image`;
const CHEVRON_COMPACT = `${MAC_BUTTON} > .cui-globalnav-image-compact`;
const HOVERING = { hover: MAC_LINK, widths: WIDE, height: 900 };
const ITEM_MOTION = ['transition-property', 'transition-duration', 'transition-timing-function', 'transition-delay'];
// apple.com's flyout is read below its locale banner (70px from 834px): a
// place here is the same distance from the nav's top
const FPLACE = 'the same distance from the nav\'s top as apple.com\'s (flyout:*, its locale banner above its nav)';
// apple.com's places in the window are below its locale banner (70px from
// 834px, 135px to 833px): a place here is the same distance from the bar
const PLACE = 'the same distance from the bar as apple.com\'s (search:*, its locale banner above its nav)';

// the bag: the example with its button opening its flyout, an empty bag's
const BAG_NAV = ex('bag');
const BAG_DARK = BAG_NAV.replace('class="cui-globalnav"', 'class="cui-globalnav cui-dark"');
const BBUTTON = `${LIST} > li:nth-child(8) > .cui-globalnav-icon`;
const BFLY = '.cui-globalnav-bag';
const BPANEL = '.cui-globalnav-bagpanel';
const BMESSAGE = '.cui-globalnav-bagmessage';
const BTEXT = '.cui-globalnav-bagtext';
const BTEXTLINK = `${BTEXT} > a`;
const BHEADING = '.cui-globalnav-bagheading';
const BROW = '.cui-globalnav-baglinks > li:first-child';
const BROW2 = '.cui-globalnav-baglinks > li:nth-child(2)';
const BLINK = `${BROW} > .cui-globalnav-baglink`;
const BCURTAIN = `${BFLY} > .cui-globalnav-curtain`;
const BCLOSE = `${BFLY} .cui-globalnav-closebutton`;
const BAGGING = { click: BBUTTON, height: 900 };
// the bag's motion held at its start, for a case that reads its start
const BHOLD = '<style>.cui-globalnav-bag, .cui-globalnav-bag * { transition-delay: 10s !important; }</style>';
// the flyouts' example with the bag's flyout in the bag link's place: what
// the crate writes for both
const FLYOUT_BAG = FLYOUT_NAV
  .replace(/<li><a class="cui-globalnav-icon" href="\/bag"[^\n]*<\/a><\/li>/, BAG_NAV.match(/<li><button class="cui-globalnav-icon" popovertarget="cui-globalnav-bag"[^\n]*<\/button><\/li>/)[0])
  .replace(/(<button class="cui-globalnav-menubutton"[^\n]*<\/button>)/, (m) => `${m}\n${BAG_NAV.slice(BAG_NAV.indexOf('<div class="cui-globalnav-bag"'), BAG_NAV.lastIndexOf('</div>\n</nav>'))}`);
// apple.com's bag is read below its locale banner (70px from 834px, 135px
// to 833px): a place here is the same distance from the bar
const BPLACE = 'the same distance from the bar as apple.com\'s (bag-flyout:*, its locale banner above its nav)';

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
  // nothing in the nav can be selected, as apple.com's (nav-select:*): the
  // bar, its list, logo, links, icons and button, a closed flyout, and the
  // search's field (reference/select/: it still takes text)
  {
    name: 'select',
    html: FLYOUT_NAV,
    widths: [1600, 390],
    checks: [
      { target: '.cui-globalnav', props: ['user-select'], apple: ['nav-select', 'nav'] },
      { target: LIST, props: ['user-select'], apple: ['nav-select', 'list'] },
      { target: LOGO, props: ['user-select'], apple: ['nav-select', 'logo'] },
      { target: MAC_LINK, props: ['user-select'], apple: ['nav-select', 'link'] },
      { target: BAG, props: ['user-select'], apple: ['nav-select', 'bag'] },
      { target: BUTTON, props: ['user-select'], apple: ['nav-select', 'menutrigger'], at: { 1600: 390 } },
      { target: MAC_FLY, props: ['user-select'], apple: ['nav-select', 'flyout'] },
    ],
  },
  {
    name: 'select-search',
    html: SEARCH_NAV,
    widths: [1600],
    ...SEARCHING,
    checks: [
      { target: INPUT, props: ['user-select'], apple: ['nav-select', 'search.field'] },
    ],
  },
  {
    name: 'select-search-phone',
    html: SEARCH_NAV,
    widths: [390],
    ...SEARCHING,
    checks: [
      { target: INPUT, props: ['user-select'], apple: ['nav-select', 'search.field-phone'] },
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
      { target: '.cui-globalnav', props: ['background-color'], apple: ['search', 'flyout'] },
      { target: '.cui-globalnav', props: ['background-image'], golden: 'the flyout\'s ground alone: apple.com\'s flyout runs from the nav\'s top (search:flyout, margin-top -44px) over the nav\'s own ground, its #fafafc at the bar (a pixel read 2026-10-21)' },
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
    // its suggestions: two sections, the quick links' type and rows, each
    // row's typed part marked
    name: 'search-suggestions',
    html: SUGGESTING,
    ...SEARCHING,
    widths: WIDE,
    reducedMotion: 'reduce',
    checks: [
      { target: LHEAD, props: [...TYPE, 'box-height'], apple: ['search-suggestions', 'links.header'] },
      { target: SHEAD, props: [...TYPE, 'box-height'], apple: ['search-suggestions', 'searches.header'] },
      { target: LHEAD, props: ['box-top'], golden: `32px below the field, ${PLACE}` },
      { target: SHEAD, props: ['box-top'], golden: `34px below the first section's last row, ${PLACE}` },
      { target: LROW, props: [...TYPE, ...PADDING, 'display', 'align-items', 'text-decoration-line', 'background-color', 'border-top-left-radius', 'box-height'], apple: ['search-suggestions', 'link'] },
      { target: `${LROW} > span`, props: TYPE, apple: ['search-suggestions', 'text'] },
      { target: MARK, props: [...TYPE, 'background-color', 'display'], apple: ['search-suggestions', 'highlight'] },
      { target: LROW, state: 'hover', props: ['color', 'background-color'], apple: ['search-suggestions', 'link.states', 'hover'] },
      { target: LROW, state: 'focus-visible', props: ['color', 'background-color', 'outline-style', 'outline-color', 'outline-width'], apple: ['search-suggestions', 'link.states', 'focus-visible'] },
      { target: MARK, state: 'hover', on: LROW, props: ['color'], apple: ['search-suggestions', 'highlight.states', 'hover'] },
      { target: MARK, state: 'focus-visible', on: LROW, props: ['color'], apple: ['search-suggestions', 'highlight.states', 'focus-visible'] },
      { target: LROW, props: ['box-top', 'box-left'], golden: `8px below its heading, 8px out from the field's start, ${PLACE}` },
      { target: LROW2, props: ['box-top'], golden: `4px below the first, ${PLACE}` },
      { target: SROW, props: ['box-top'], golden: `8px below its heading, ${PLACE}` },
      { target: `${LROW} > .cui-globalnav-image`, props: ['color'], as: { color: 'fill' }, apple: ['search-suggestions', 'links.icon'] },
      { target: `${SROW} > .cui-globalnav-image`, props: ['box-height', 'color'], as: { color: 'fill' }, apple: ['search-suggestions', 'searches.image'] },
      { target: `${SROW} > .cui-globalnav-image`, props: ['margin-right'], apple: ['search-suggestions', 'searches.icon'] },
      { target: `${SROW} > .cui-globalnav-image`, state: 'hover', on: SROW, props: ['color'], as: { color: 'fill' }, apple: ['search-suggestions', 'searches.icon-states', 'hover'] },
      { target: `${SROW} > .cui-globalnav-image`, props: ['box-width'], golden: 'apple.com\'s magnifier, 9 by 16 (search-suggestions:searches.image)' },
      { target: `${SROW} > .cui-globalnav-image-compact`, props: ['box-height'], golden: 'the phone\'s magnifier, not drawn from 834px' },
    ],
  },
  {
    name: 'search-suggestions-phone',
    html: SUGGESTING,
    ...SEARCHING,
    widths: NARROW,
    reducedMotion: 'reduce',
    checks: [
      { target: LHEAD, props: [...TYPE, 'box-height'], apple: ['search-suggestions', 'links.header'] },
      { target: SHEAD, props: [...TYPE, 'box-height'], apple: ['search-suggestions', 'searches.header'] },
      { target: SHEAD, props: ['box-top'], golden: `48px below the first section's last row, ${PLACE}` },
      { target: LROW, props: [...TYPE, ...PADDING, 'background-color', 'border-top-left-radius', 'box-height'], apple: ['search-suggestions', 'link'] },
      { target: MARK, props: [...TYPE, 'background-color'], apple: ['search-suggestions', 'highlight'] },
      { target: LROW, state: 'hover', props: ['color', 'background-color'], apple: ['search-suggestions', 'link.states', 'hover'] },
      { target: MARK, state: 'hover', on: LROW, props: ['color'], apple: ['search-suggestions', 'highlight.states', 'hover'] },
      { target: SROW, props: ['box-top'], golden: `9px below its heading, ${PLACE}` },
      { target: `${SROW} > .cui-globalnav-image-compact`, props: ['box-height', 'color'], as: { color: 'fill' }, apple: ['search-suggestions', 'searches.image-compact'] },
      { target: `${SROW} > .cui-globalnav-image-compact`, props: ['margin-right'], apple: ['search-suggestions', 'searches.icon'] },
      { target: `${SROW} > .cui-globalnav-image-compact`, props: ['box-width'], golden: 'apple.com\'s magnifier, 13 by 25 (search-suggestions:searches.image-compact)' },
    ],
  },
  {
    name: 'search-suggestions-dark',
    html: SUGGESTING_DARK,
    ...SEARCHING,
    widths: [1600, 390],
    reducedMotion: 'reduce',
    checks: [
      { target: LHEAD, props: ['color'], apple: ['search-suggestions-dark', 'links.header'] },
      { target: LROW, props: ['color', 'background-color'], apple: ['search-suggestions-dark', 'link'] },
      { target: MARK, props: ['color', 'background-color'], apple: ['search-suggestions-dark', 'highlight'] },
      { target: LROW, state: 'hover', props: ['color', 'background-color'], apple: ['search-suggestions-dark', 'link.states', 'hover'] },
      { target: MARK, state: 'hover', on: LROW, props: ['color'], apple: ['search-suggestions-dark', 'highlight.states', 'hover'] },
    ],
  },
  {
    name: 'search-suggestions-dark-flyout',
    html: SUGGESTING_DARK,
    ...SEARCHING,
    widths: [1600],
    reducedMotion: 'reduce',
    checks: [
      { target: `${SROW} > .cui-globalnav-image`, props: ['color'], as: { color: 'fill' }, apple: ['search-suggestions-dark', 'searches.image'] },
      { target: `${SROW} > .cui-globalnav-image`, state: 'hover', on: SROW, props: ['color'], as: { color: 'fill' }, apple: ['search-suggestions-dark', 'searches.icon-states', 'hover'] },
    ],
  },
  {
    // with text in the field, a row the site's script writes fades in where
    // it is, as apple.com's new results do (search-typed:ip, a section to
    // a section: its rows before fade out under them, here gone); read 150ms
    // after the click, inside the opening's 200ms delay
    name: 'search-suggestions-rows',
    html: `${SUGGESTING}${LATER}`,
    ...SEARCHING,
    wait: 150,
    widths: [1600, 390],
    checks: [
      { target: SHEAD, props: ['transition-property', 'transition-duration', 'transition-timing-function', 'transition-delay'], golden: 'apple.com\'s new heading in 320ms from 40ms, its opacity alone (search-typed:ip, globalnav-search-fade: in place)' },
      ...[[1, 60], [2, 80], [3, 100]].map(([n, ms]) => ({ target: `.cui-globalnav-searchlinks > li:nth-child(${n})`, props: ['transition-property', 'transition-duration', 'transition-delay'], golden: `apple.com's new rows in 320ms, each 20ms after the one above: the ${['first', 'second', 'third'][n - 1]} from ${ms}ms (search-typed:ip)${n === 3 ? ', written once the search was open' : ''}` })),
    ],
  },
  {
    // read at its start, any motion held there: from transparent, in place
    name: 'search-suggestions-start',
    html: `${HOLD}${SUGGESTING}${LATER}`,
    ...SEARCHING,
    widths: [1600],
    checks: [
      { target: '.cui-globalnav-searchlinks > li:nth-child(3)', props: ['opacity', 'transform'], golden: 'a row written once something is typed from transparent, where it is (search-typed:ip\'s keyframes: opacity 0 to 1, translateY(0px))' },
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
  {
    // shown by the pointer resting on Mac's link
    name: 'flyout',
    html: FLYOUT_NAV,
    ...HOVERING,
    checks: [
      { target: '.cui-globalnav', props: ['background-color', 'backdrop-filter'], apple: ['flyout', 'nav'] },
      { target: '.cui-globalnav', props: ['position', 'isolation'], golden: 'the containing block and the stacking context of a flyout the pointer shows (a site\'s sticky replaces the position)' },
      { target: MAC_LINK, props: ['color', ...MOTION], apple: ['flyout', 'link.open'] },
      { target: IPAD_LINK, props: ['color'], apple: ['flyout', 'link.other'] },
      { target: MAC_FLY, props: ['position', 'visibility'], apple: ['flyout', 'flyout'] },
      { target: MAC_FLY, props: ['box-top', 'box-left', 'box-width', 'height', 'overflow-y'], golden: 'from the nav\'s top across the window, as apple.com\'s (flyout:flyout, 44px down less 44px: at its nav\'s top, below its banner), its panel\'s height (apple.com\'s its own, its overflow hidden)' },
      { target: FPANEL, props: ['background-color'], apple: ['flyout', 'flyout'] },
      { target: FPANEL, props: ['padding-bottom'], apple: ['flyout', 'flyout.content'] },
      { target: FPANEL, props: ['padding-top', 'box-height', 'clip-path'], golden: 'its content 84px down from the nav\'s top, apple.com\'s 44px of bar (flyout:flyout.scroll\'s margin) and 40px (flyout.content\'s padding); revealed whole, its height its groups\' (apple.com\'s 552px for Mac\'s, its groups 388px)' },
      { target: GROUP1, props: ['flex-grow', 'flex-shrink', 'flex-basis', 'max-width', 'padding-right', 'margin-bottom', 'box-left'], apple: ['flyout', 'group.elevated'] },
      { target: GROUP1, props: ['box-top'], golden: `84px down, ${FPLACE}` },
      { target: GROUP2, props: ['flex-grow', 'flex-shrink', 'flex-basis', 'max-width', 'padding-right', 'margin-bottom'], apple: ['flyout', 'group'] },
      { target: GROUP3, props: ['max-width', 'padding-right', 'margin-bottom'], apple: ['flyout', 'group.last'] },
      { target: FHEAD1, props: [...TYPE, 'margin-top', 'margin-bottom', 'display', 'box-height'], apple: ['flyout', 'header.elevated'] },
      { target: FHEAD2, props: [...TYPE, 'margin-top', 'margin-bottom', 'box-height'], apple: ['flyout', 'header'] },
      { target: FHEAD1, props: ['letter-spacing'], golden: 'the system font\'s fitted letter-spacing (tokens.css), not apple.com\'s' },
      { target: FLIST2, props: ['list-style-type', 'margin-top', 'margin-bottom', 'padding-top', 'padding-left'], apple: ['flyout', 'list'] },
      { target: ROW_LG, props: ['margin-top', 'margin-left', 'margin-right', 'box-height'], apple: ['flyout', 'item.elevated'] },
      { target: ROW_LG2, props: ['margin-top', 'box-height'], apple: ['flyout', 'item.elevated2'] },
      { target: ROW_AFTER, props: ['margin-top', 'margin-left', 'box-height'], apple: ['flyout', 'item.after'] },
      { target: ROW, props: ['margin-top', 'margin-left', 'box-height'], apple: ['flyout', 'item'] },
      { target: ROW2, props: ['margin-top', 'box-height'], apple: ['flyout', 'item2'] },
      { target: ROW_LG, props: ['box-top', 'box-left'], golden: `6px below its heading, 11px out, ${FPLACE}` },
      { target: LINK_LG, props: [...TYPE, ...PADDING, 'margin-bottom', 'display', 'text-decoration-line', 'background-color', ...MOTION, 'box-height'], apple: ['flyout', 'link.elevated'] },
      { target: LINK_LG, state: 'hover', props: ['color'], apple: ['flyout', 'link.elevated', 'hover'] },
      { target: LINK_LG, state: 'focus-visible', props: ['color', ...OUTLINE], apple: ['flyout', 'link.elevated', 'focus-visible'] },
      { target: LINK_LG, props: ['letter-spacing'], golden: 'the system font\'s fitted letter-spacing (tokens.css), not apple.com\'s' },
      { target: LINK_AFTER, props: [...TYPE, ...PADDING, 'margin-bottom', 'display', 'box-height'], apple: ['flyout', 'link.after'] },
      { target: LINK_AFTER, state: 'hover', props: ['color'], apple: ['flyout', 'link.after', 'hover'] },
      { target: FLINK, props: [...TYPE, ...PADDING, 'margin-bottom', 'display', 'text-decoration-line', ...MOTION, 'box-height'], apple: ['flyout', 'link'] },
      { target: FLINK, state: 'hover', props: ['color'], apple: ['flyout', 'link', 'hover'] },
      { target: FLINK, state: 'focus-visible', props: ['color', ...OUTLINE], apple: ['flyout', 'link', 'focus-visible'] },
      // the curtain, the list's ::after
      { target: LIST, pseudo: '::after', props: ['position', 'top', 'background-color', 'backdrop-filter', 'opacity', 'visibility'], apple: ['flyout', 'curtain'] },
      { target: LIST, pseudo: '::after', props: ['z-index', 'pointer-events'], golden: 'under the bar\'s links and the flyout, in the nav (apple.com\'s under the nav, 9998 to its 9999); taking no pointer, so that the pointer on it is off the link (apple.com\'s takes it, its script closing the flyout as the pointer leaves the nav: reference/flyout/)' },
      // the link's button and its chevron
      { target: MAC_BUTTON, props: ['box-height', 'cursor', 'color', 'display', 'align-items', 'justify-content', 'background-color', 'border-top-width'], apple: ['flyout', 'trigger'] },
      { target: MAC_BUTTON, state: 'focus-visible', props: OUTLINE, apple: ['flyout', 'trigger', 'focus-visible'] },
      { target: MAC_BUTTON, props: ['box-width', 'margin-right'], golden: 'apple.com\'s 22 by 44 (flyout:trigger), from 8px before the link\'s end to 14px past it, over its room, not its text (apple.com\'s from 15px before it, over its text\'s end, which a click there opens the flyout from)' },
      { target: CHEVRON, props: ['opacity', 'margin-top', 'color', 'box-height', ...MOTION], apple: ['flyout', 'trigger.icon'] },
      { target: CHEVRON, state: 'focus-visible', on: MAC_BUTTON, props: ['opacity', 'margin-top'], apple: ['flyout', 'trigger.icon', 'focus-visible'] },
      { target: CHEVRON, props: ['box-width'], golden: 'apple.com\'s chevron, 10 by 10 (flyout:trigger.icon)' },
    ],
  },
  {
    name: 'flyout-motion',
    html: FLYOUT_NAV,
    ...HOVERING,
    widths: [1600],
    checks: [
      { target: MAC_FLY, props: ['transition-property', 'transition-duration', 'transition-delay'], golden: 'shown once the pointer has rested 120ms on its link (apple.com\'s script\'s 120ms, read in globalheader.umd.js: flyout:*)' },
      { target: FPANEL, props: ITEM_MOTION, golden: 'revealed down from the bar and from its ground in 240ms after those 120ms: apple.com\'s flyout\'s height and ground in half its content\'s height in ms, 240 at least (flyout:flyout, 254ms for Mac\'s; 240ms for seven of the eleven)' },
      { target: FHEAD1, props: ['transition-duration', 'transition-timing-function'], apple: ['flyout', 'header.elevated'] },
      { target: FHEAD1, props: ['transition-delay'], golden: 'apple.com\'s first heading from 80ms (flyout:header.elevated), after the 120ms' },
      { target: FHEAD2, props: ['transition-delay'], golden: 'apple.com\'s second from 160ms (flyout:header), after the 120ms' },
      { target: FHEAD3, props: ['transition-delay'], golden: 'apple.com\'s third from 200ms (flyout:header.last), after the 120ms' },
      { target: ROW_LG, props: ['transition-duration', 'transition-timing-function'], apple: ['flyout', 'item.elevated'] },
      { target: ROW_LG, props: ['transition-delay'], golden: 'apple.com\'s rows 40ms after their heading (flyout:item.elevated, 120ms), after the 120ms' },
      { target: ROW_LG2, props: ['transition-delay'], golden: 'each 20ms after the one above (flyout:item.elevated2, 140ms), after the 120ms' },
      { target: ROW, props: ['transition-delay'], golden: 'the second group\'s from 200ms (flyout:item), after the 120ms' },
      { target: LIST, pseudo: '::after', props: ['transition-property', 'transition-duration', 'transition-timing-function', 'transition-delay'], golden: 'apple.com\'s curtain in 0.32s from 80ms (flyout:curtain), after the 120ms' },
      { target: '.cui-globalnav', props: ['transition-property', 'transition-duration', 'transition-delay'], golden: 'its blur off as the flyout shows, after the 120ms' },
    ],
  },
  {
    // the pointer off the bar, down to the page
    name: 'flyout-closed',
    html: FLYOUT_NAV,
    ...HOVERING,
    hover: [MAC_LINK, [700, 890]],
    reducedMotion: 'reduce',
    checks: [
      { target: MAC_FLY, props: ['visibility'], apple: ['flyout', 'close'] },
      { target: FPANEL, props: ['background-color'], apple: ['flyout', 'close'] },
      { target: FPANEL, props: ['clip-path'], golden: 'back to the bar: apple.com\'s flyout the bar\'s 44px again (flyout:close)' },
      { target: MAC_FLY, props: ['height', 'box-height'], apple: ['flyout', 'close'] },
      { target: LIST, pseudo: '::after', props: ['opacity'], apple: ['flyout', 'close.curtain'] },
      { target: LIST, pseudo: '::after', props: ['visibility'], golden: 'the curtain hidden' },
      { target: '.cui-globalnav', props: ['background-color', 'backdrop-filter'], apple: ['flyout', 'close.nav'] },
    ],
  },
  {
    name: 'flyout-closing',
    html: FLYOUT_NAV,
    ...HOVERING,
    hover: [MAC_LINK, [700, 890]],
    widths: [1600],
    checks: [
      { target: MAC_FLY, props: ['transition-property', 'transition-duration', 'transition-delay'], golden: 'hidden at the end of its going (apple.com\'s flyout hidden at its height\'s end: flyout:close\'s animations, 240ms from 120ms)' },
      { target: MAC_FLY, props: ['transition-behavior'], golden: 'its height, which does not interpolate, flipped with its visibility at that end: allow-discrete in the compiled stylesheet too (written inside the shorthand: the minifier moves a transition-behavior written after the shorthand before it, where the shorthand resets it)' },
      { target: FPANEL, props: ITEM_MOTION, golden: 'back to the bar and its ground in 240ms from 120ms (flyout:close\'s animations)' },
      { target: FHEAD1, props: ITEM_MOTION, golden: 'its headings back to 4px up and transparent at once in 240ms (flyout:close\'s animations)' },
      { target: ROW_LG, props: ['transition-duration', 'transition-delay'], golden: 'its rows the same, a group\'s first five (its heading the first) in 240ms (flyout:close\'s animations)' },
      { target: ROW_AFTER, props: ['transition-duration'], golden: 'and each after them 20ms shorter, its fifth row 220ms (flyout:close\'s animations)' },
      { target: LIST, pseudo: '::after', props: ['transition-property', 'transition-duration', 'transition-delay'], golden: 'the curtain out in 0.32s from 80ms, hidden at its end (flyout:close\'s animations)' },
      { target: '.cui-globalnav', props: ['transition-property', 'transition-duration', 'transition-delay'], golden: 'its blur back 0.45s on, the curtain gone (apple.com\'s back 0.4 to 0.53s after a flyout or the search closed: scouted, each frame, 2026-10-21; kept on, the curtain would not reach the page: reference/flyout/)' },
    ],
  },
  {
    // the pointer on along the bar from Store's link to Mac's
    name: 'flyout-change',
    html: FLYOUT_NAV,
    ...HOVERING,
    hover: [STORE_LINK, MAC_LINK],
    reducedMotion: 'reduce',
    checks: [
      { target: MAC_FLY, props: ['visibility'], apple: ['flyout', 'change'] },
      { target: `${MAC_FLY} > .cui-globalnav-flyoutpanel`, props: ['background-color'], apple: ['flyout', 'change'] },
      { target: STORE_FLY, props: ['visibility'], apple: ['flyout', 'change.previous'] },
      { target: STORE_LINK, props: ['color'], apple: ['flyout', 'change.link'] },
    ],
  },
  {
    name: 'flyout-change-motion',
    html: FLYOUT_NAV,
    ...HOVERING,
    hover: [STORE_LINK, MAC_LINK],
    widths: [1600],
    checks: [
      { target: STORE_FLY, props: ['transition-delay'], golden: 'the one left shown 0.6s while the pointer is on the bar, until the next has grown (apple.com\'s keeps it until the next opens, then crossfades them in 120ms: flyout:change\'s animations); in every frame between, one shown to its full height (reference/flyout/)' },
      { target: `${STORE_FLY} > .cui-globalnav-flyoutpanel`, props: ['transition-delay'], golden: 'its panel back to the bar from 0.36s, the next grown by then' },
    ],
  },
  {
    // opened by its button: the keyboard's (a click here)
    name: 'flyout-button',
    html: FLYOUT_NAV,
    click: MAC_BUTTON,
    widths: WIDE,
    height: 900,
    reducedMotion: 'reduce',
    checks: [
      { target: MAC_FLY, props: ['position', 'box-top', 'box-left', 'box-height'], golden: 'a popover from the bar\'s bottom (anchor positioning) to the window\'s, across it (reference/flyout/), where the one the pointer shows is from the nav\'s top: the same content in the same place' },
      { target: FPANEL, props: ['background-color'], apple: ['flyout', 'flyout'] },
      { target: FPANEL, props: ['padding-top', 'padding-bottom'], apple: ['flyout', 'flyout.content'] },
      { target: GROUP1, props: ['max-width', 'padding-right', 'box-left'], apple: ['flyout', 'group.elevated'] },
      { target: LINK_LG, props: [...TYPE, 'box-height'], apple: ['flyout', 'link.elevated'] },
      { target: '.cui-globalnav', props: ['background-color'], apple: ['flyout', 'flyout'] },
      { target: '.cui-globalnav', props: ['backdrop-filter'], apple: ['flyout', 'nav'] },
      { target: MAC_LINK, props: ['color'], apple: ['flyout', 'link.open'] },
      { target: MAC_CURTAIN, props: ['background-color', 'backdrop-filter', 'opacity'], apple: ['flyout', 'curtain'] },
      { target: MAC_CURTAIN, props: ['box-top', 'box-height', 'hit'], golden: 'behind the panel from the bar\'s bottom to the window\'s, a button: what the pointer finds under the panel (as the search\'s)' },
      { target: BACK, props: ['display'], golden: 'the sheet\'s back button, none from 834px' },
      { target: FCLOSE, props: ['display'], golden: 'the sheet\'s X, none from 834px' },
    ],
  },
  {
    name: 'flyout-button-motion',
    html: FLYOUT_NAV,
    click: MAC_BUTTON,
    widths: [1600],
    height: 900,
    checks: [
      { target: FPANEL, props: ITEM_MOTION, golden: 'revealed down from the bar and from its ground in 240ms, as the search\'s (apple.com\'s flyout\'s height and ground: flyout:flyout)' },
      { target: FHEAD1, props: ['transition-delay'], golden: 'apple.com\'s first heading from 80ms (flyout:header.elevated)' },
      { target: ROW_LG2, props: ['transition-delay'], golden: 'apple.com\'s second row from 140ms (flyout:item.elevated2)' },
    ],
  },
  {
    name: 'flyout-tab',
    html: FLYOUT_NAV,
    click: MAC_BUTTON,
    press: 'Tab',
    widths: [1200],
    height: 900,
    reducedMotion: 'reduce',
    checks: [{ target: LINK_LG, props: ['focused'], golden: 'Tab from its button to its first link (reference/flyout/)' }],
  },
  {
    name: 'flyout-escape',
    html: FLYOUT_NAV,
    click: MAC_BUTTON,
    press: 'Escape',
    widths: [1200],
    height: 900,
    reducedMotion: 'reduce',
    checks: [
      { target: MAC_BUTTON, props: ['focused'], golden: 'Escape closes it, the focus back on its button (reference/flyout/)' },
      { target: MAC_FLY, props: ['position', 'visibility'], golden: 'closed, in the bar and hidden again' },
    ],
  },
  {
    // a screen with no hover: its button over the link, which a tap opens the
    // flyout by, as apple.com's script does there (tap-flyout)
    name: 'flyout-touch',
    html: FLYOUT_NAV,
    touch: true,
    widths: WIDE,
    reducedMotion: 'reduce',
    checks: [
      { target: MAC_BUTTON, props: ['box-left', 'box-width', 'box-height'], golden: 'over the whole link' },
      { target: MAC_LINK, props: ['hit'], golden: 'what a tap on the link finds is its button' },
      { target: MAC_FLY, props: ['display'], golden: 'not shown by the pointer: a touch screen has none that hovers' },
      { target: MAC_BUTTON, state: 'hover', props: ['color'], rest: 'no hover on a touch screen (rule 5)' },
      { target: CHEVRON, props: ['opacity'], golden: 'its chevron not drawn' },
    ],
  },
  {
    // a tap on the link lands on its button (flyout-touch's hit)
    name: 'flyout-tap',
    html: FLYOUT_NAV,
    touch: true,
    click: MAC_BUTTON,
    widths: [1200],
    height: 900,
    reducedMotion: 'reduce',
    checks: [
      { target: FPANEL, props: ['background-color'], apple: ['tap-flyout', 'flyout'], at: { 1200: 1024 } },
      { target: MAC_FLY, props: ['position', 'box-top'], golden: 'opened by the tap on its link as a popover from the bar\'s bottom, the link not followed (tap-flyout: apple.com\'s script opens it so)' },
      { target: FLINK, state: 'hover', props: ['color'], rest: 'no hover on a touch screen (rule 5), where apple.com\'s keeps it' },
    ],
  },
  {
    name: 'flyout-dark',
    html: FLYOUT_DARK,
    ...HOVERING,
    widths: [1600],
    reducedMotion: 'reduce',
    checks: [
      { target: FPANEL, props: ['background-color'], apple: ['flyout-dark', 'flyout'] },
      { target: FHEAD2, props: ['color'], apple: ['flyout-dark', 'header'] },
      { target: LINK_LG, props: ['color'], apple: ['flyout-dark', 'link.elevated'] },
      { target: FLINK, props: ['color'], apple: ['flyout-dark', 'link'] },
      { target: FLINK, state: 'hover', props: ['color'], apple: ['flyout-dark', 'link', 'hover'] },
      { target: FLINK, state: 'focus-visible', props: ['color'], apple: ['flyout-dark', 'link', 'focus-visible'] },
      { target: MAC_LINK, props: ['color'], apple: ['flyout-dark', 'link.open'] },
      { target: LIST, pseudo: '::after', props: ['background-color', 'backdrop-filter'], apple: ['flyout-dark', 'curtain'] },
    ],
  },
  {
    // to 833px: the menu opened, then Mac's row
    name: 'flyout-phone',
    html: FLYOUT_NAV,
    click: [BUTTON, MAC_BUTTON],
    widths: [390],
    height: 900,
    checks: [
      { target: MAC_FLY, props: ['position', 'box-top', 'box-left', 'box-height', 'opacity', 'transform'], golden: 'a sheet over the open menu, the window\'s (flyout-phone:flyout over the menu\'s content)' },
      { target: MAC_FLY, props: ['background-color'], apple: ['flyout-phone', 'content'] },
      { target: FPANEL, props: ['padding-top', 'padding-left', 'padding-right'], apple: ['flyout-phone', 'flyout.content'] },
      { target: GROUP1, props: ['padding-bottom', 'max-width'], apple: ['flyout-phone', 'group.elevated'] },
      { target: GROUP2, props: ['padding-bottom', 'max-width'], apple: ['flyout-phone', 'group'] },
      { target: FHEAD1, props: ['display'], apple: ['flyout-phone', 'header.elevated'] },
      { target: FHEAD2, props: [...TYPE, 'box-height'], apple: ['flyout-phone', 'header'] },
      { target: ROW_LG, props: ['margin-top', 'margin-left', 'margin-right', 'box-height', 'box-top', 'box-left'], apple: ['flyout-phone', 'item.elevated'] },
      { target: ROW_LG2, props: ['margin-top', 'box-height'], apple: ['flyout-phone', 'item.elevated2'] },
      { target: ROW_AFTER, props: ['margin-top', 'box-height'], apple: ['flyout-phone', 'item.after'] },
      { target: ROW, props: ['margin-top', 'box-height'], apple: ['flyout-phone', 'item'] },
      { target: ROW2, props: ['margin-top', 'box-height'], apple: ['flyout-phone', 'item2'] },
      { target: LINK_LG, props: [...TYPE, ...PADDING, 'margin-bottom', 'display', 'max-width', ...MOTION, 'box-height'], apple: ['flyout-phone', 'link.elevated'] },
      { target: LINK_LG, state: 'hover', props: ['color'], apple: ['flyout-phone', 'link.elevated', 'hover'] },
      { target: LINK_LG, state: 'focus-visible', props: ['color', ...OUTLINE], apple: ['flyout-phone', 'link.elevated', 'focus-visible'] },
      { target: FLINK, props: [...TYPE, ...PADDING, 'margin-bottom', 'display', 'max-width', 'box-height'], apple: ['flyout-phone', 'link'] },
      { target: FLINK, state: 'hover', props: ['color'], apple: ['flyout-phone', 'link', 'hover'] },
      { target: BACK, props: ['box-top', 'box-left', 'box-height'], apple: ['flyout-phone', 'back'] },
      { target: BACK, props: ['color', 'cursor', 'background-color'], apple: ['flyout-phone', 'back.button'] },
      { target: BACK, state: 'hover', props: ['color'], apple: ['flyout-phone', 'back.button', 'hover'] },
      { target: BACK, state: 'focus-visible', props: OUTLINE, apple: ['flyout-phone', 'back.button', 'focus-visible'] },
      { target: BACK, props: ['box-width'], golden: 'apple.com\'s 48px square (flyout-phone:back)' },
      { target: `${BACK} > svg`, props: ['box-height', 'box-top', 'box-left'], apple: ['flyout-phone', 'back.image'] },
      { target: `${BACK} > svg`, props: ['box-width'], golden: 'apple.com\'s chevron, 9 by 48 (flyout-phone:back.image)' },
      { target: FCLOSE, props: ['color', 'opacity', 'box-top', 'box-left', 'box-height'], apple: ['flyout-phone', 'close'] },
      { target: STORE_ITEM, props: ['opacity', 'transform'], golden: 'the menu\'s rows 8px back and out behind it (flyout-phone:flyout\'s animations, its slides\' keyframes)' },
    ],
  },
  {
    name: 'flyout-phone-motion',
    html: FLYOUT_NAV,
    click: [BUTTON, MAC_BUTTON],
    widths: [390],
    height: 900,
    checks: [
      { target: MAC_FLY, props: ITEM_MOTION, golden: 'in from 8px on and transparent in 240ms from 160ms (flyout-phone:flyout\'s animations, globalnav-flyout-slide-forward-next)' },
      { target: BACK, props: ITEM_MOTION, golden: 'its color in 0.32s (flyout-phone:back.button), in from 4px on and transparent in 240ms from 160ms (flyout-phone:back)' },
      { target: STORE_ITEM, props: ['transition-duration', 'transition-delay'], golden: 'the menu\'s rows out at once in 240ms (flyout-phone:flyout\'s animations, globalnav-flyout-slide-forward-previous)' },
    ],
  },
  {
    // read at its start, its motion held there
    name: 'flyout-phone-start',
    html: `<style>.cui-globalnav-flyout, .cui-globalnav-flyout * { transition-delay: 10s !important; }</style>${FLYOUT_NAV}`,
    click: [BUTTON, MAC_BUTTON],
    widths: [390],
    height: 900,
    wait: 0,
    checks: [
      { target: MAC_FLY, props: ['opacity', 'transform'], golden: 'from transparent, 8px on (flyout-phone:keyframes, globalnav-flyout-slide-forward-next)' },
      { target: BACK, props: ['opacity', 'transform'], golden: 'from transparent, 4px on (flyout-phone:back\'s transition)' },
    ],
  },
  {
    name: 'flyout-phone-back',
    html: FLYOUT_NAV,
    click: [BUTTON, MAC_BUTTON, BACK],
    widths: [390],
    height: 900,
    reducedMotion: 'reduce',
    checks: [
      { target: MAC_FLY, props: ['display'], golden: 'its back button closes it (reference/flyout/)' },
      { target: LIST, props: ['position'], golden: 'the menu still open (reference/flyout/)' },
      { target: STORE_ITEM, props: ['opacity', 'transform'], golden: 'its rows back' },
    ],
  },
  {
    name: 'flyout-phone-close',
    html: FLYOUT_NAV,
    click: [BUTTON, MAC_BUTTON, FCLOSE],
    widths: [390],
    height: 900,
    reducedMotion: 'reduce',
    checks: [
      { target: MAC_FLY, props: ['display'], golden: 'its X closes it (reference/flyout/)' },
      { target: LIST, props: ['position'], golden: 'and the menu (reference/flyout/)' },
    ],
  },
  {
    // a phone's width with a pointer: a row hovered shows its chevron
    name: 'flyout-phone-row',
    html: FLYOUT_NAV,
    click: BUTTON,
    hover: MAC_BUTTON,
    widths: [390],
    height: 900,
    reducedMotion: 'reduce',
    checks: [
      { target: CHEVRON_COMPACT, props: ['opacity', 'color'], apple: ['flyout-phone', 'row.chevron'] },
      { target: CHEVRON_COMPACT, props: ['box-height'], apple: ['flyout-phone', 'row.chevron-image'] },
      { target: CHEVRON_COMPACT, props: ['box-top', 'box-left', 'box-width', 'transform'], golden: 'apple.com\'s 9 by 48 chevron 19px in from the row\'s end, 2px down it (flyout-phone:row.chevron-image, its row 93px down)' },
      { target: MAC_LINK, props: ['color'], apple: ['flyout-phone', 'row.text'] },
    ],
  },
  {
    name: 'flyout-phone-row-rest',
    html: FLYOUT_NAV,
    click: BUTTON,
    widths: [390],
    height: 900,
    checks: [
      { target: `${STORE_ITEM} > .cui-globalnav-flyoutbutton > .cui-globalnav-image-compact`, props: ['opacity'], apple: ['flyout-phone', 'row.other-chevron'] },
      { target: `${STORE_ITEM} > .cui-globalnav-flyoutbutton > .cui-globalnav-image-compact`, props: ['transform', ...MOTION], golden: 'apple.com\'s chevron slides in from 4px back in 0.24s (flyout-phone:row.chevron\'s keyframes) and goes as it came, where apple.com\'s scales to 0.8 (row.other-chevron)' },
    ],
  },
  {
    name: 'flyout-phone-touch',
    html: FLYOUT_NAV,
    touch: true,
    click: BUTTON,
    widths: [390],
    height: 900,
    reducedMotion: 'reduce',
    checks: [
      { target: MAC_BUTTON, state: 'hover', props: ['color'], rest: 'no hover on a touch screen (rule 5)' },
      { target: CHEVRON_COMPACT, state: 'hover', on: MAC_BUTTON, props: ['opacity'], rest: 'its chevron not shown by a tap (rule 5), where apple.com\'s shows on a hovered row' },
    ],
  },
  // the bag (bag-flyout:*): opened by its button, from 834px a flyout from
  // the bar's bottom over its curtain, to 833px a sheet over the window
  {
    name: 'bag',
    html: BAG_NAV,
    ...BAGGING,
    widths: WIDE,
    reducedMotion: 'reduce',
    checks: [
      // the bar over it, its button, the page not held
      { target: '.cui-globalnav', props: ['backdrop-filter'], apple: ['bag-flyout', 'nav'] },
      { target: '.cui-globalnav', props: ['background-color'], apple: ['bag-flyout', 'flyout'] },
      { target: '.cui-globalnav', props: ['background-image'], golden: 'the flyout\'s ground alone: apple.com\'s flyout runs from the nav\'s top (bag-flyout:flyout, margin-top -44px) over the nav\'s own ground' },
      { target: BBUTTON, props: ['color'], apple: ['bag-flyout', 'link.bag'] },
      { target: BBUTTON, props: ['focused'], golden: 'opened by its button, the focus stays on it: its X (autofocus) is not shown from 834px' },
      { target: 'html', props: ['overflow-y'], apple: ['bag-flyout', 'root'] },
      // the flyout from the bar's bottom, its panel, the curtain behind it
      { target: BFLY, props: ['box-top', 'box-left', 'box-height'], golden: 'from the bar\'s bottom (anchor positioning) to the window\'s, across it, as the search\'s (reference/search/)' },
      { target: BPANEL, props: ['background-color'], apple: ['bag-flyout', 'flyout'] },
      { target: BPANEL, props: ['padding-top', 'padding-bottom'], apple: ['bag-flyout', 'flyout.content'] },
      { target: BPANEL, props: ['box-height'], apple: ['bag-flyout', 'flyout.scroll'] },
      { target: BPANEL, props: ['box-top'], golden: 'from the bar\'s bottom, where apple.com\'s flyout\'s room starts (bag-flyout:flyout.scroll, 44px into its nav)' },
      { target: BCURTAIN, props: ['background-color', 'backdrop-filter', 'opacity'], apple: ['bag-flyout', 'curtain'] },
      { target: BCURTAIN, props: ['box-top', 'box-height', 'hit'], golden: 'behind the panel from the bar\'s bottom to the window\'s, a button: what the pointer finds under the panel (apple.com\'s curtain is the window\'s, under its flyout)' },
      { target: BCLOSE, props: ['display'], golden: 'the sheet\'s X, not shown from 834px' },
      // its message, its line, its heading
      { target: BMESSAGE, props: [...TYPE, 'margin-bottom', 'box-height', 'box-left'], apple: ['bag-flyout', 'bag.header'] },
      { target: BMESSAGE, props: ['letter-spacing'], golden: 'the system font\'s fitted letter-spacing (tokens.css), not apple.com\'s' },
      { target: BMESSAGE, props: ['box-top'], golden: `40px below the bar, ${BPLACE}` },
      { target: BTEXT, props: [...TYPE, 'box-height', 'box-left'], apple: ['bag-flyout', 'bag.subheader'] },
      { target: BTEXT, props: ['letter-spacing'], golden: 'the system font\'s fitted letter-spacing (tokens.css), not apple.com\'s' },
      { target: BTEXT, props: ['box-top'], golden: `24px below the message, ${BPLACE}` },
      { target: BTEXTLINK, props: ['color', 'text-decoration-line'], apple: ['bag-flyout', 'bag.subheader-link'] },
      { target: BTEXTLINK, state: 'focus-visible', props: ['outline-style', 'outline-color', 'outline-width'], apple: ['bag-flyout', 'bag.subheader-link', 'focus-visible'] },
      { target: BTEXTLINK, state: 'focus-visible', props: ['outline-offset'], golden: 'the kit\'s one ring, 3px out (base.css), where apple.com\'s is 1px out (bag-flyout:bag.subheader-link)' },
      { target: BHEADING, props: [...TYPE, 'margin-top', 'padding-bottom', 'box-height', 'box-left'], apple: ['bag-flyout', 'bag.nav-header'] },
      { target: BHEADING, props: ['letter-spacing'], golden: 'the system font\'s fitted letter-spacing (tokens.css), not apple.com\'s' },
      { target: BHEADING, props: ['box-top'], golden: `36px below the line, ${BPLACE}` },
      // its links: a 28px row each, its icon before its text
      { target: BROW, props: ['box-height'], apple: ['bag-flyout', 'bag.nav-item'] },
      { target: BROW2, props: ['box-top'], golden: `28px below the first, ${BPLACE}` },
      { target: BLINK, props: ['font-size', 'line-height', 'font-weight', 'margin-top', 'margin-bottom', 'display', 'text-decoration-line', 'background-color', 'box-height', 'box-left'], apple: ['bag-flyout', 'bag.nav-link'] },
      { target: BLINK, props: ['color'], apple: ['bag-flyout', 'bag.nav-text'] },
      { target: BLINK, state: 'hover', props: ['color'], apple: ['bag-flyout', 'bag.nav-text', 'hover'] },
      { target: BLINK, props: ['letter-spacing'], golden: 'the system font\'s fitted letter-spacing (tokens.css), not apple.com\'s' },
      { target: BLINK, props: ['transition-property', 'transition-duration'], apple: ['bag-flyout', 'bag.nav-link'] },
      { target: BLINK, state: 'focus-visible', props: ['outline-style', 'outline-color', 'outline-width'], apple: ['bag-flyout', 'bag.nav-link', 'focus-visible'] },
      { target: BLINK, state: 'focus-visible', props: ['color', 'outline-offset'], golden: 'the kit\'s one ring, 3px out (base.css), where apple.com\'s is 1px out, its color the link\'s (bag-flyout:bag.nav-link)' },
      { target: `${BLINK} .cui-globalnav-image`, props: ['box-height', 'box-left', 'color'], apple: ['bag-flyout', 'bag.nav-image-regular'] },
      { target: `${BLINK} .cui-globalnav-image`, state: 'hover', on: BLINK, props: ['color'], apple: ['bag-flyout', 'bag.nav-icon', 'hover'] },
      { target: `${BLINK} .cui-globalnav-image`, props: ['box-width', 'margin-right'], golden: 'apple.com\'s 11 by 16 icon, 1px in, the text 24px in after it (bag-flyout:bag.nav-image-regular, bag.nav-text\'s padding-left)' },
      { target: `${BLINK} .cui-globalnav-image-compact`, props: ['display'], golden: 'the phone\'s icon, not drawn from 834px' },
    ],
  },
  {
    name: 'bag-phone',
    html: BAG_NAV,
    ...BAGGING,
    widths: NARROW,
    reducedMotion: 'reduce',
    checks: [
      { target: BFLY, props: ['background-color'], apple: ['bag-flyout', 'nav.content'] },
      { target: BFLY, props: ['box-top', 'box-left', 'box-height'], golden: 'a sheet over the window, the bar too (apple.com\'s, bag-flyout:nav.content, from its nav\'s top below its locale banner)' },
      { target: 'html', props: ['overflow-y'], apple: ['bag-flyout', 'root'] },
      { target: BPANEL, props: ['padding-top', 'padding-left', 'padding-right'], apple: ['bag-flyout', 'flyout.content'] },
      { target: BPANEL, props: ['padding-bottom'], apple: ['bag-flyout', 'flyout.scroll'] },
      { target: BCURTAIN, props: ['display'], golden: 'no curtain: the sheet is over the window (apple.com\'s, under it, is not seen)' },
      { target: BMESSAGE, props: [...TYPE, 'margin-bottom', 'box-height', 'box-left'], apple: ['bag-flyout', 'bag.header'] },
      { target: BMESSAGE, props: ['letter-spacing'], golden: 'the system font\'s fitted letter-spacing (tokens.css), not apple.com\'s' },
      { target: BMESSAGE, props: ['box-top'], golden: `50px down the sheet, ${BPLACE}` },
      { target: BTEXT, props: [...TYPE, 'box-left'], apple: ['bag-flyout', 'bag.subheader'] },
      { target: BTEXTLINK, props: ['color', 'text-decoration-line'], apple: ['bag-flyout', 'bag.subheader-link'] },
      { target: BHEADING, props: [...TYPE, 'margin-top', 'padding-bottom', 'box-height', 'box-left'], apple: ['bag-flyout', 'bag.nav-header'] },
      { target: BROW, props: ['box-height'], apple: ['bag-flyout', 'bag.nav-item'] },
      { target: BLINK, props: ['font-size', 'line-height', 'font-weight', 'margin-top', 'margin-bottom', 'display', 'box-height', 'box-left'], apple: ['bag-flyout', 'bag.nav-link'] },
      { target: BLINK, props: ['color'], apple: ['bag-flyout', 'bag.nav-text'] },
      { target: BLINK, state: 'hover', props: ['color'], apple: ['bag-flyout', 'bag.nav-text', 'hover'] },
      { target: `${BLINK} .cui-globalnav-image-compact`, props: ['box-height', 'box-left', 'color'], apple: ['bag-flyout', 'bag.nav-image-compact'] },
      { target: `${BLINK} .cui-globalnav-image-compact`, state: 'hover', on: BLINK, props: ['color'], apple: ['bag-flyout', 'bag.nav-icon-compact', 'hover'] },
      { target: `${BLINK} .cui-globalnav-image-compact`, props: ['box-width', 'margin-right'], golden: 'apple.com\'s 16 by 25 icon, the text 30px in after it (bag-flyout:bag.nav-image-compact, bag.nav-text\'s padding-left)' },
      { target: `${BLINK} .cui-globalnav-image`, props: ['display'], golden: 'the wide icon, not drawn to 833px' },
      { target: BCLOSE, props: ['color', 'opacity'], apple: ['bag-flyout', 'menutrigger'] },
      { target: BCLOSE, props: ['box-top', 'box-left', 'box-width', 'box-height'], golden: 'apple.com\'s X where the menu\'s button is, its Close (bag-flyout:menutrigger: the menu\'s button turned, 48px at the right)' },
      { target: BCLOSE, props: ['focused'], golden: 'opened, the focus on its X (autofocus), as the menu\'s' },
    ],
  },
  {
    name: 'bag-motion',
    html: BAG_NAV,
    ...BAGGING,
    widths: [1600, 390],
    checks: [
      { target: BMESSAGE, props: [...MOTION, 'transition-delay'], golden: 'apple.com\'s rows in 0.32s from 80ms (bag-flyout:bag.header: opacity and transform, cubic-bezier(.4,0,.6,1))' },
      { target: BTEXT, props: ['transition-delay'], golden: 'apple.com\'s line 20ms after the message (bag-flyout:bag.subheader)' },
      { target: BHEADING, props: ['transition-delay'], golden: 'apple.com\'s heading 20ms after the line (bag-flyout:bag.nav-header)' },
      { target: BROW, props: ['transition-delay'], golden: 'apple.com\'s first row 20ms after the heading (bag-flyout:bag.nav-item)' },
      { target: BROW2, props: ['transition-delay'], golden: 'apple.com\'s rows 20ms after the one above (bag-flyout:bag.nav-item2)' },
      { target: BFLY, props: MOTION, golden: 'the sheet\'s height and ground in 450ms to 833px (bag-flyout:nav.phone\'s animations on its content); no motion of its own from 834px' },
      { target: BPANEL, props: MOTION, golden: 'from 834px revealed down and from the bar\'s ground in 0.24s, apple.com\'s flyout\'s height and ground (bag-flyout:nav\'s animations)' },
      { target: BCURTAIN, props: MOTION, golden: 'apple.com\'s curtain in 0.32s (bag-flyout:curtain), from 834px' },
    ],
  },
  {
    // read at its start, held there: apple.com's from (the link flyouts'
    // rule, globalnav-flyout-item, on its rows)
    name: 'bag-start',
    html: `${BHOLD}${BAG_NAV}`,
    ...BAGGING,
    widths: [1600],
    wait: 0,
    checks: [
      { target: BMESSAGE, props: ['opacity', 'transform'], golden: 'each row from transparent, 4px up (globalheader.css\'s globalnav-flyout-item, bag-flyout:bag.header\'s transition)' },
      { target: BROW, props: ['opacity', 'transform'], golden: 'each row from transparent, 4px up' },
      { target: BPANEL, props: ['clip-path', 'background-color'], golden: 'the panel hidden down from the bar, on its ground: apple.com\'s flyout 44px tall, the bar\'s ground (bag-flyout:nav\'s animations)' },
      { target: BCURTAIN, props: ['opacity'], golden: 'the curtain from transparent (bag-flyout:curtain\'s transition)' },
    ],
  },
  {
    name: 'bag-start-phone',
    html: `${BHOLD}${BAG_NAV}`,
    ...BAGGING,
    widths: [390],
    wait: 0,
    checks: [
      { target: BFLY, props: ['height', 'background-color'], golden: 'the sheet from the bar\'s height and ground (bag-flyout:nav.phone\'s animations on its content)' },
      { target: BMESSAGE, props: ['opacity', 'transform'], golden: 'each row from transparent, 8px up (globalheader.css\'s globalnav-flyout-item to 833px)' },
    ],
  },
  {
    name: 'bag-tab',
    html: BAG_NAV,
    ...BAGGING,
    widths: [1600],
    press: 'Tab',
    reducedMotion: 'reduce',
    checks: [{ target: BTEXTLINK, props: ['focused'], golden: 'Tab goes from its button into it, to its line\'s link' }],
  },
  {
    name: 'bag-escape',
    html: BAG_NAV,
    ...BAGGING,
    widths: [1600, 390],
    press: 'Escape',
    reducedMotion: 'reduce',
    checks: [
      { target: BFLY, props: ['display'], golden: 'Escape closes it' },
      { target: BBUTTON, props: ['focused'], golden: 'the focus back on its button' },
      { target: 'html', props: ['overflow-y'], golden: 'the page scrolls again' },
    ],
  },
  {
    name: 'bag-curtain',
    html: BAG_NAV,
    click: [BBUTTON, BCURTAIN],
    height: 900,
    widths: [1600],
    reducedMotion: 'reduce',
    checks: [
      { target: BFLY, props: ['display'], golden: 'a click on the curtain closes it, as the search\'s (reference/search/)' },
      { target: '.cui-globalnav', props: ['background-color', 'backdrop-filter'], apple: ['home', 'nav'] },
    ],
  },
  {
    name: 'bag-again',
    html: BAG_NAV,
    click: [BBUTTON, BBUTTON],
    height: 900,
    widths: [1600],
    reducedMotion: 'reduce',
    checks: [{ target: BFLY, props: ['display'], golden: 'its button closes it again' }],
  },
  {
    // its links to fragments here, as a click on the bar's middle is on one
    name: 'bag-bar',
    html: BAG_NAV.replaceAll('href="/', 'href="#/'),
    click: [BBUTTON, '.cui-globalnav'],
    height: 900,
    widths: [1600],
    reducedMotion: 'reduce',
    checks: [{ target: BFLY, props: ['display'], golden: 'a click on the bar, outside it, closes it, as the search\'s (reference/search/)' }],
  },
  {
    name: 'bag-close',
    html: BAG_NAV,
    click: [BBUTTON, BCLOSE],
    height: 900,
    widths: [390],
    reducedMotion: 'reduce',
    checks: [
      { target: BFLY, props: ['display'], golden: 'the sheet\'s X closes it' },
      { target: 'html', props: ['overflow-y'], golden: 'the page scrolls again' },
    ],
  },
  {
    // the pointer on a link with a flyout while the bag is open shows none
    name: 'bag-pointer',
    html: FLYOUT_BAG,
    ...BAGGING,
    hover: MAC_LINK,
    widths: [1600],
    reducedMotion: 'reduce',
    checks: [
      { target: MAC_FLY, props: ['display'], golden: 'no flyout by the pointer while the bag is open, its popover closed and not shown by the page (apple.com\'s opens the link\'s in the bag\'s place)' },
      { target: BFLY, props: ['display'], golden: 'the bag still open' },
    ],
  },
  {
    name: 'bag-dark',
    html: BAG_DARK,
    ...BAGGING,
    widths: [1600, 390],
    reducedMotion: 'reduce',
    checks: [
      { target: BMESSAGE, props: ['color'], apple: ['bag-flyout-dark', 'bag.header'] },
      { target: BTEXT, props: ['color'], apple: ['bag-flyout-dark', 'bag.subheader'] },
      { target: BTEXTLINK, props: ['color', 'text-decoration-line'], apple: ['bag-flyout-dark', 'bag.subheader-link'] },
      { target: BHEADING, props: ['color'], apple: ['bag-flyout-dark', 'bag.nav-header'] },
      { target: BLINK, props: ['color'], apple: ['bag-flyout-dark', 'bag.nav-text'] },
      { target: BLINK, state: 'hover', props: ['color'], apple: ['bag-flyout-dark', 'bag.nav-text', 'hover'] },
    ],
  },
  {
    name: 'bag-dark-flyout',
    html: BAG_DARK,
    ...BAGGING,
    widths: [1600],
    reducedMotion: 'reduce',
    checks: [
      { target: BPANEL, props: ['background-color'], apple: ['bag-flyout-dark', 'flyout'] },
      { target: BCURTAIN, props: ['background-color', 'backdrop-filter'], apple: ['bag-flyout-dark', 'curtain'] },
      { target: BBUTTON, props: ['color'], apple: ['bag-flyout-dark', 'link.bag'] },
      { target: `${BLINK} .cui-globalnav-image`, props: ['color'], apple: ['bag-flyout-dark', 'bag.nav-icon'] },
      { target: `${BLINK} .cui-globalnav-image`, state: 'hover', on: BLINK, props: ['color'], apple: ['bag-flyout-dark', 'bag.nav-icon', 'hover'] },
    ],
  },
  {
    name: 'bag-dark-sheet',
    html: BAG_DARK,
    ...BAGGING,
    widths: [390],
    reducedMotion: 'reduce',
    checks: [
      { target: BFLY, props: ['background-color'], apple: ['bag-flyout-dark', 'nav.content'] },
      { target: BCLOSE, props: ['color'], apple: ['bag-flyout-dark', 'menutrigger'] },
      { target: `${BLINK} .cui-globalnav-image-compact`, props: ['color'], apple: ['bag-flyout-dark', 'bag.nav-icon-compact'] },
      { target: `${BLINK} .cui-globalnav-image-compact`, state: 'hover', on: BLINK, props: ['color'], apple: ['bag-flyout-dark', 'bag.nav-icon-compact', 'hover'] },
    ],
  },
  {
    name: 'bag-touch',
    html: BAG_NAV,
    ...BAGGING,
    touch: true,
    widths: [1200, 390],
    reducedMotion: 'reduce',
    checks: [
      { target: BLINK, state: 'hover', props: ['color'], rest: 'no hover on a touch screen (rule 5), where apple.com\'s keeps its links\' (touch-*:*)' },
      { target: `${BLINK} .cui-globalnav-image`, state: 'hover', on: BLINK, props: ['color'], rest: 'nor its icon\'s (rule 5)' },
    ],
  },

];
