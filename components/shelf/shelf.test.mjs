// The shelf against apple.com's store (store:shelf*): its room below, its
// header's inset and room, its headline's and secondary text's type and
// colors, its scroller (its overflow and snapping, the room above and
// below its cards in sight, its inset at each end, its gap), its items'
// snapping, the first two cards' places; scrolled 300px, the second card
// snapped to the inset; scrolled to its end, the last card as far in; its
// paddles (store:shelf.paddle*: their places over the scroller, the next
// shown, the previous hidden, their motion). Golden, with their why: the
// letter-spacing, the header's box, the foot in sight, the scrollbar hidden,
// the last card's place at the end, no paddles on a touch screen.

import { examples, SIZE } from '../../test/harness.mjs';

const ex = examples(import.meta.url);
const TRACKING = 'the system font\'s fitted letter-spacing (tokens.css), not apple.com\'s';
const SCROLLER = '.cui-shelf-scroller';
const card = (n) => `${SCROLLER} > li:nth-child(${n}) .cui-card`;

export const cases = [
  {
    name: 'shelf',
    html: ex('shelf'),
    checks: [
      { target: '.cui-shelf', props: ['padding-top', 'padding-bottom', 'background-color'], apple: ['store', 'shelf'] },
      { target: '.cui-shelf-header', props: ['margin-left', 'padding-top', 'padding-bottom', 'box-left'], apple: ['store', 'shelf.header'] },
      { target: '.cui-shelf-header', props: ['box-width', 'box-height'], golden: 'apple.com\'s header\'s box (store:shelf.header: to the window\'s right, 87.5% of a phone\'s window), its lines Inter\'s here' },
      { target: '.cui-shelf-headline', props: [...SIZE, 'color', 'display', 'margin-top', 'margin-bottom'], apple: ['store', 'shelf.mainheader'] },
      { target: '.cui-shelf-subhead', props: [...SIZE, 'color', 'display'], apple: ['store', 'shelf.subheader'] },
      { target: '.cui-shelf-headline', props: ['letter-spacing'], golden: TRACKING },
      { target: '.cui-shelf-subhead', props: ['letter-spacing'], golden: TRACKING },
      { target: SCROLLER, props: ['overflow-x', 'overflow-y', 'scroll-snap-type', 'overscroll-behavior-x'], apple: ['store', 'shelf.scroller'] },
      { target: SCROLLER, props: ['padding-top'], apple: ['store', 'shelf.platter'] },
      // in sight: apple.com's crop, its cards and the room above and below them
      { target: SCROLLER, props: ['box-height'], apple: ['store', 'shelf.crop'] },
      { target: SCROLLER, props: ['padding-bottom'], golden: 'the 22px of apple.com\'s foot in sight: its scroller\'s 40px under a crop 38px below its cards (store:shelf.crop, .platter)' },
      { target: SCROLLER, props: ['scrollbar-width'], golden: 'hidden, as apple.com\'s crop hides it' },
      // the inset at each end: the header's, the last item's padding
      { target: SCROLLER, props: ['padding-left', 'scroll-padding-left'], as: { 'padding-left': 'margin-left', 'scroll-padding-left': 'margin-left' }, apple: ['store', 'shelf.header'] },
      { target: SCROLLER, props: ['padding-right'], apple: ['store', 'shelf.card-last'] },
      { target: SCROLLER, props: ['column-gap'], as: { 'column-gap': 'margin-right' }, apple: ['store', 'shelf.card'] },
      { target: `${SCROLLER} > li`, props: ['scroll-snap-align', 'scroll-snap-stop'], apple: ['store', 'shelf.item'] },
      { target: card(1), props: ['box-left'], apple: ['store', 'shelf.card'] },
      { target: card(2), props: ['box-left'], apple: ['store', 'shelf.card-second'] },
    ],
  },
  {
    // between the first card and the second, nearer the second: it snaps
    // there, at the inset (at 1600px four cards scroll less than a card)
    name: 'snap',
    html: ex('shelf'),
    widths: [1200, 900, 390],
    scroll: [SCROLLER, 300],
    checks: [
      { target: card(2), props: ['box-left'], apple: ['store', 'shelf.card'] },
    ],
  },
  {
    name: 'paddles',
    html: ex('paddles'),
    checks: [
      { target: '.cui-shelf-paddle:not(:disabled)', props: ['position', 'top', 'right', 'margin-top', 'opacity', 'transform', 'transition-property', 'transition-duration', 'transition-timing-function', 'box-height'], apple: ['store', 'shelf.paddle'] },
      { target: '.cui-shelf-paddle:disabled', props: ['position', 'top', 'left', 'margin-top', 'opacity', 'transform', 'transition-property', 'transition-duration', 'transition-timing-function'], apple: ['store', 'shelf.paddle-previous'] },
      { target: '.cui-shelf-row', props: ['position', 'box-height'], apple: ['store', 'shelf.paddles'] },
    ],
  },
  {
    name: 'paddles-touch',
    html: ex('paddles'),
    touch: true,
    checks: [
      { target: '.cui-shelf-paddle:not(:disabled)', props: ['display'], golden: 'none on a touch screen: apple.com\'s draws no paddles there (an emulated iPhone 13)' },
    ],
  },
  {
    name: 'end',
    html: ex('shelf'),
    scroll: [SCROLLER, 99999],
    checks: [
      { target: card(4), props: ['box-right'], golden: 'at its end the last card as far in from the window\'s right as the first from its left (apple.com\'s last item padded so far, store:shelf.card-last)' },
    ],
  },
];
