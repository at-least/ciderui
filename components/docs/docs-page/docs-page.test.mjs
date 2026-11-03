import { DOCC_STEPS, MOTION, PADDING, examples } from '../../../test/harness.mjs';

const ex = examples(import.meta.url);
const NARROW = [1023, 900, 767, 390];
const COLUMN = '.cui-docs-page-column:has(> .cui-docs-content)';
const SIDEBAR = '.cui-docs-page-sidebar';
const CLOSE = '.cui-docs-page-close';
const BOX = [...PADDING, 'margin-top', 'margin-bottom', 'border-top-width', 'border-top-style', 'border-bottom-width', 'border-bottom-style', 'border-top-left-radius', 'background-color', 'box-shadow'];

export const cases = [
  {
    name: 'page',
    html: ex('page'),
    widths: DOCC_STEPS,
    checks: [
      { target: COLUMN, props: ['padding-left', 'padding-right', 'max-width'], apple: ['docc-page', 'container'] },
      { target: `${COLUMN} > .cui-docs-content`, props: ['margin-top'], apple: ['docc-page', 'content.top'] },
      { target: COLUMN, props: ['box-width', 'box-x'], golden: 'the content\'s column: 87.5% of the window to 735px, centered; from 736px 80px in within 820px (980px from 1251px, 1080px from 1500px), centered in what the navigator\'s 400px leaves from 1024px (docc-page:container, primary; DocC\'s navigator its script\'s width)' },
      { target: SIDEBAR, props: ['display', 'position', 'top', 'box-width'], golden: 'the navigator\'s place: from 1024px a column 400px wide sticking under the bar (DocC\'s script\'s width at a fresh load up to 1300px wide); below, a popover not shown until the bar\'s button opens it' },
      { target: '.cui-docs-page > .cui-docs-bar', props: ['z-index'], golden: 'over the columns as they scroll under it (DocC\'s 9997, over its own layers)' },
    ],
  },
  {
    name: 'page wide',
    html: ex('page'),
    widths: [1024],
    checks: [
      { target: '.cui-docs-page > .cui-docs-bar', props: ['position', 'top', 'box-width'], golden: 'across both columns, sticking at the window\'s top, as below 1024px (docc-sheet:bar.place, read to 1023px)' },
      { target: SIDEBAR, props: ['overflow-y', 'box-height'], golden: 'the navigator\'s column its own scroll, the window\'s height less the bar\'s with no --cui-docs-page-top' },
      { target: CLOSE, props: ['display'], golden: 'none from 1024px: the navigator is a column, with no sheet to close' },
    ],
  },
  {
    name: 'page bar',
    html: ex('page'),
    widths: NARROW,
    checks: [
      { target: '.cui-docs-page > .cui-docs-bar', props: ['position', 'top'], apple: ['docc-sheet', 'bar.place'] },
    ],
  },
  {
    name: 'page sheet',
    html: ex('page'),
    widths: NARROW,
    click: '.cui-docs-bar-toggle',
    checks: [
      { target: SIDEBAR, props: ['position'], apple: ['docc-sheet', 'sheet'] },
      { target: SIDEBAR, props: ['background-color', 'color'], apple: ['docc-sheet', 'sheet.card'] },
      { target: CLOSE, props: [...BOX, 'height', 'color', 'display'], apple: ['docc-sheet', 'sheet.close'] },
      { target: `${CLOSE} svg`, props: ['height', 'color'], apple: ['docc-sheet', 'sheet.close.icon'] },
      { target: SIDEBAR, props: ['top', 'box-height', 'box-width', 'overflow-y'], golden: 'over the whole window, the global nav too (a popover is in the top layer; DocC\'s starts under its global nav, 44px down, docc-sheet:sheet top), its own scroll' },
      { target: CLOSE, props: ['width', 'box-y', 'box-x'], golden: 'DocC\'s 19px button 20px from the sheet\'s end, its middle 24px down (docc-sheet:sheet.close: 14.5px down to 767px, 11.5px and 25px tall from 768px)' },
      { target: `${CLOSE} svg`, props: ['width'], golden: 'square, DocC\'s 19px glyph (docc-sheet:sheet.close.icon)' },
      { target: SIDEBAR, props: MOTION, golden: 'shown at once: DocC\'s slides in by a transform in 0.15s (docc-sheet:sheet), from a place not read' },
    ],
  },
  {
    name: 'page sheet dark',
    html: `<div class="cui-dark">${ex('page')}</div>`,
    widths: [390],
    click: '.cui-docs-bar-toggle',
    checks: [
      { target: SIDEBAR, props: ['background-color', 'color'], apple: ['docc-sheet-dark', 'sheet.card'] },
      { target: CLOSE, props: ['color'], apple: ['docc-sheet-dark', 'sheet.close'] },
    ],
  },
  {
    name: 'page target',
    html: ex('page').replace('<h2>Overview</h2>', '<h2 id="overview">Overview</h2>').replace('<main class="cui-docs-page-main">', '<main class="cui-docs-page-main"><a id="to-overview" href="#overview">Overview</a>'),
    widths: [1200, 767],
    click: '#to-overview',
    checks: [
      { target: '#overview', props: ['scroll-margin-top'], golden: 'a place a link leads to lands under the bar that sticks over it, its 52px (48px below 768px): DocC\'s script scrolls it there' },
    ],
  },
  {
    name: 'page top',
    html: ex('page').replace('<div class="cui-docs-page-sidebar" id="docs-navigator" popover>', '<div class="cui-docs-page-sidebar" id="docs-navigator" popover style="--cui-docs-page-top: 96px">'),
    widths: [1200],
    checks: [
      { target: SIDEBAR, props: ['box-height'], golden: 'the window\'s height (800px) less its top in the window (--cui-docs-page-top, 96px), so the navigator\'s foot, its filter, is at the window\'s foot (DocC\'s script takes its global nav off the same way)' },
    ],
  },
  {
    name: 'page touch',
    html: ex('page'),
    widths: [390],
    touch: true,
    click: '.cui-docs-bar-toggle',
    checks: [
      { target: CLOSE, pseudo: '::after', props: ['width', 'height'], golden: 'the 19px close button\'s target 44px square on a touch screen, the HIG\'s 44pt' },
    ],
  },
];
