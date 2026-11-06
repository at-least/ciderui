import { examples, MOTION, WITH_1600 } from '../../../test/harness.mjs';

const ex = examples(import.meta.url);
const SCROLLER = '.cui-gallery-scroller';
const NEXT = '.cui-gallery-paddle:not(:disabled)';
const PREVIOUS = '.cui-gallery-paddle:disabled';
const DOT = '.cui-gallery-dots li:nth-child(2) .cui-gallery-dot';
const CURRENT = '.cui-gallery-dot[aria-current="true"]';
const PADDLE = ['height', 'padding-top', 'padding-left', 'border-top-width', 'border-top-left-radius', 'background-color', 'color', 'opacity', 'cursor', 'box-height'];
const DOTS = ['height', 'background-color', 'border-top-left-radius', 'padding-top'];
const dark = (html) => html.replace('class="cui-gallery"', 'class="cui-gallery cui-dark"');
const noPaddles = (html) => html.replace(/<div class="cui-gallery-nav">[\s\S]*?<\/div>\n/, '');
const lg = (html) => html.replace('<button class="cui-gallery-paddle" type="button" aria-controls="gallery" aria-label="Next">', '<button class="cui-gallery-paddle" type="button" aria-controls="gallery" aria-label="Next" data-size="lg">');

export const cases = [
  {
    name: 'gallery',
    widths: WITH_1600,
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
      { target: NEXT, props: MOTION, golden: 'its dim as it is disabled in apple.com\'s 0.1s (iphone:paddle), its colors at once: motion that adds nothing to a pointer\'s color change (DESIGN.md decision 25)' },
      { target: NEXT, props: ['width', 'box-width'], golden: 'apple.com\'s 36px circle (iphone:paddle), a width no text sets' },
      { target: NEXT, state: 'hover', props: ['background-color', 'color'], apple: ['iphone', 'paddle', 'hover'] },
      { target: NEXT, state: 'active', props: ['background-color', 'color'], apple: ['iphone', 'paddle', 'active'] },
      { target: `${NEXT} svg`, props: ['height', 'box-height'], apple: ['iphone', 'paddle.icon'] },
      { target: PREVIOUS, props: ['opacity', 'background-color', 'color', 'cursor', 'pointer-events'], apple: ['iphone', 'paddle.disabled'] },
      { target: '.cui-gallery-dots', props: ['display', 'flex-direction', 'align-items', 'justify-content', 'box-height'], apple: ['product', 'gallery.dots'] },
      { target: '.cui-gallery-dots li:nth-child(2)', props: ['margin-left', 'margin-right'], apple: ['product', 'gallery.dot-item'] },
      { target: DOT, props: DOTS, apple: ['iphone', 'gallery.dot'] },
      { target: DOT, state: 'hover', props: ['background-color'], apple: ['iphone', 'gallery.dot', 'hover'] },
      { target: CURRENT, props: DOTS, apple: ['iphone', 'gallery.dot-current'] },
      { target: DOT, props: ['width'], golden: 'apple.com\'s 8px dot (iphone:gallery.dot), a width no text sets' },
      { target: CURRENT, props: ['width'], golden: 'apple.com\'s pill, 48px, 32px below 735px (iphone:gallery.dot-current)' },
      { target: CURRENT, pseudo: '::after', props: ['content'], golden: 'the pill with no timer\'s fill: no script plays the gallery (apple.com\'s fills it as its item shows, iphone:gallery.dot-fill)' },
      { target: SCROLLER, props: ['scrollbar-width'], golden: 'hidden with paddles beside it, as apple.com\'s is' },
      { target: '.cui-gallery-nav', props: ['position'], golden: 'the paddles\' row in the gallery\'s flow (apple.com\'s placed below its scroller)' },
    ],
  },
  {
    name: 'no paddles',
    html: noPaddles(ex('gallery')),
    checks: [
      { target: SCROLLER, props: ['scrollbar-width'], golden: 'a thin scrollbar with no paddles, where apple.com\'s hides it, as it draws paddles: with neither, a mouse with no sideways wheel cannot reach the rest of the row, nor can anyone see there is more' },
    ],
  },
  {
    name: 'no paddles touch',
    html: noPaddles(ex('gallery')),
    touch: true,
    checks: [
      { target: SCROLLER, props: ['scrollbar-width'], golden: 'a touch screen swipes: hidden still, as apple.com\'s' },
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
    name: 'forced',
    html: ex('gallery'),
    forcedColors: true,
    checks: [
      { target: DOT, props: ['border-top-width', 'border-top-style', 'border-top-color'], golden: 'in forced colors a dot drawn by a 4px edge in the theme\'s link color, where its ground alone, forced to the window\'s, drew nothing (apple.com\'s the same: not measured there)' },
      { target: CURRENT, props: ['border-top-width', 'border-top-style', 'box-width'], golden: 'in forced colors the current dot a pill of the same edge, longer' },
    ],
  },
  {
    name: 'touch targets',
    html: lg(ex('gallery')),
    widths: [390],
    touch: true,
    checks: [
      { target: NEXT, props: ['width', 'height'], golden: 'a large paddle keeps its 56px on a touch screen' },
      { target: '.cui-gallery-dots li:nth-child(2)', props: ['margin-left', 'margin-right'], golden: 'a touch screen\'s dots 44px apart at least: each dot\'s 44px target (::after) with no neighbour\'s over it' },
    ],
  },
  {
    name: 'snap',
    html: ex('gallery'),
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
