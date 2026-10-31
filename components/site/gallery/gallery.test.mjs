// The gallery against apple.com's: its scroller (product:gallery.scroller,
// .items, .set, .item: the overflow and snapping, the room, the inset, the
// gap, the first item's place), its paddles' row (product:gallery.nav,
// .nav-list) and a paddle on a light page (iphone:paddle*: its size, ground,
// chevron, motion, hovered, pressed, disabled) and a dark one
// (product:paddle.next, its place at the content's right edge), the store's
// 56px (store:shelf.paddle), its dots (iphone:gallery.dot*,
// product:gallery.dot*: their size, color, gap, the current pill); scrolled
// 250px, the second item snapped to the inset. Golden, with their why: the
// widths of the row, the dark page's disabled opacity, the pill with no
// fill, a large paddle's colors.

import { examples } from '../../../test/harness.mjs';

const ex = examples(import.meta.url);
const SCROLLER = '.cui-gallery-scroller';
const NEXT = '.cui-gallery-paddle:not(:disabled)';
const PREVIOUS = '.cui-gallery-paddle:disabled';
const DOT = '.cui-gallery-dots li:nth-child(2) .cui-gallery-dot';
const CURRENT = '.cui-gallery-dot[aria-current="true"]';
const PADDLE = ['height', 'padding-top', 'padding-left', 'border-top-width', 'border-top-left-radius', 'background-color', 'color', 'opacity', 'cursor', 'transition-property', 'transition-duration', 'transition-timing-function', 'box-height'];
const DOTS = ['height', 'background-color', 'border-top-left-radius', 'padding-top'];
const dark = (html) => html.replace('class="cui-gallery"', 'class="cui-gallery cui-dark"');
const lg = (html) => html.replace('<button class="cui-gallery-paddle" type="button" aria-controls="gallery" aria-label="Next">', '<button class="cui-gallery-paddle" type="button" aria-controls="gallery" aria-label="Next" data-size="lg">');

export const cases = [
  {
    name: 'gallery',
    html: ex('gallery'),
    checks: [
      { target: SCROLLER, props: ['overflow-x', 'overflow-y', 'scroll-snap-type', 'scroll-padding-left', 'padding-top', 'padding-bottom', 'margin-top', 'margin-bottom'], apple: ['product', 'gallery.scroller'] },
      { target: SCROLLER, props: ['padding-left', 'padding-right'], apple: ['product', 'gallery.items'] },
      { target: SCROLLER, props: ['column-gap'], apple: ['product', 'gallery.set'] },
      { target: `${SCROLLER} > li`, props: ['scroll-snap-align', 'scroll-snap-stop', 'box-left'], apple: ['product', 'gallery.item'] },
      { target: '.cui-gallery-nav', props: ['display', 'justify-content', 'padding-top', 'box-left', 'box-height'], apple: ['product', 'gallery.nav'] },
      { target: '.cui-gallery-nav', props: ['box-width'], golden: 'the content\'s width, apple.com\'s row\'s (product:gallery.nav: 87.5% of the window, 1260px at most, 420px below 735px)' },
      { target: '.cui-gallery-nav', props: ['column-gap'], apple: ['product', 'gallery.nav-list'] },
      { target: NEXT, props: PADDLE, apple: ['iphone', 'paddle'] },
      { target: NEXT, props: ['box-right'], apple: ['product', 'paddle.next'] },
      { target: NEXT, props: ['width', 'box-width'], golden: 'apple.com\'s 36px circle (iphone:paddle), a width no text sets' },
      { target: NEXT, state: 'hover', props: ['background-color', 'color'], apple: ['iphone', 'paddle', 'hover'] },
      { target: NEXT, state: 'active', props: ['background-color', 'color'], apple: ['iphone', 'paddle', 'active'] },
      { target: `${NEXT} svg`, props: ['height', 'box-height'], apple: ['iphone', 'paddle.icon'] },
      { target: PREVIOUS, props: ['opacity', 'background-color', 'color', 'cursor', 'pointer-events'], apple: ['iphone', 'paddle.disabled'] },
      { target: '.cui-gallery-dots', props: ['display', 'flex-direction', 'align-items', 'justify-content', 'box-height'], apple: ['product', 'gallery.dots'] },
      { target: '.cui-gallery-dots li:nth-child(2)', props: ['margin-left', 'margin-right'], apple: ['product', 'gallery.dot-item'] },
      { target: DOT, props: [...DOTS, 'transition-property', 'transition-duration', 'transition-timing-function'], apple: ['iphone', 'gallery.dot'] },
      { target: DOT, state: 'hover', props: ['background-color'], apple: ['iphone', 'gallery.dot', 'hover'] },
      { target: CURRENT, props: DOTS, apple: ['iphone', 'gallery.dot-current'] },
      { target: DOT, props: ['width'], golden: 'apple.com\'s 8px dot (iphone:gallery.dot), a width no text sets' },
      { target: CURRENT, props: ['width'], golden: 'apple.com\'s pill, 48px, 32px below 735px (iphone:gallery.dot-current)' },
      { target: CURRENT, pseudo: '::after', props: ['content'], golden: 'the pill with no timer\'s fill: no script plays the gallery (apple.com\'s fills it as its item shows, iphone:gallery.dot-fill)' },
    ],
  },
  {
    name: 'dark',
    html: dark(ex('gallery')),
    checks: [
      { target: NEXT, props: ['background-color', 'color', 'opacity'], apple: ['product', 'paddle.next'] },
      { target: NEXT, state: 'hover', props: ['background-color', 'color'], apple: ['product', 'paddle.next', 'hover'] },
      { target: NEXT, state: 'active', props: ['background-color', 'color'], apple: ['product', 'paddle.next', 'active'] },
      { target: PREVIOUS, props: ['background-color', 'color'], apple: ['product', 'paddle'] },
      { target: PREVIOUS, props: ['opacity'], golden: 'one disabled opacity, .42, a disabled button\'s (apple.com\'s dark page dims its paddle to .36, product:paddle)' },
      { target: DOT, props: DOTS, apple: ['product', 'gallery.dot'] },
      { target: DOT, state: 'hover', props: ['background-color'], apple: ['product', 'gallery.dot', 'hover'] },
      { target: CURRENT, props: DOTS, apple: ['product', 'gallery.dot-current'] },
    ],
  },
  {
    name: 'touch',
    html: ex('gallery'),
    touch: true,
    checks: [
      { target: NEXT, state: 'hover', props: ['background-color', 'color'], rest: 'a touch screen\'s paddle does not change hovered (rule 5), where apple.com\'s does (touch-iphone:paddle)' },
      { target: DOT, state: 'hover', props: ['background-color'], rest: 'nor a dot, where apple.com\'s does (touch-iphone:gallery.dot)' },
    ],
  },
  {
    // forced colors (a high-contrast theme): the dots drawn, each a ring of
    // the theme's link color, the current one a longer pill
    name: 'forced',
    html: ex('gallery'),
    forcedColors: true,
    checks: [
      { target: DOT, props: ['border-top-width', 'border-top-style', 'border-top-color'], golden: 'in forced colors a dot drawn by a 4px edge in the theme\'s link color, where its ground alone, forced to the window\'s, drew nothing (apple.com\'s the same: not measured there)' },
      { target: CURRENT, props: ['border-top-width', 'border-top-style', 'box-width'], golden: 'in forced colors the current dot a pill of the same edge, longer' },
    ],
  },
  {
    // between the first item and the second, nearer the second (at 1600px
    // four cards scroll less than an item)
    name: 'snap',
    html: ex('gallery'),
    widths: [1200, 900, 390],
    scroll: [SCROLLER, 250],
    checks: [
      { target: `${SCROLLER} > li:nth-child(2)`, props: ['box-left'], apple: ['product', 'gallery.item'] },
    ],
  },
  {
    name: 'lg',
    html: lg(ex('gallery')),
    checks: [
      { target: NEXT, props: ['height', 'border-top-left-radius', 'box-height'], apple: ['store', 'shelf.paddle'] },
      { target: NEXT, props: ['width'], golden: 'apple.com\'s store\'s 56px circle (store:shelf.paddle)' },
      { target: NEXT, props: ['background-color'], apple: ['store', 'shelf.paddle'] },
      { target: NEXT, props: ['color'], golden: 'the 36px paddle\'s chevron, rgba(0,0,0,.56), 4.6:1 on its ground: apple.com\'s store\'s is white, 1.35:1 (store:shelf.paddle)' },
      { target: NEXT, state: 'hover', props: ['background-color', 'color'], golden: 'the 36px paddle\'s hover: apple.com\'s store\'s, rgba(222,222,227,.698), a unit from it, and its chevron white, 1.35:1 on its ground (store:shelf.paddle:hover)' },
      { target: `${NEXT} svg`, props: ['box-width', 'box-height'], golden: 'the 36px paddle\'s chevron at 56px (apple.com\'s store draws a longer one)' },
    ],
  },
];
