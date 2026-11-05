import { examples, MOTION, SIZE, TRACKING } from '../../../test/harness.mjs';

const ex = examples(import.meta.url);
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
      // apple.com's scroller runs under its crop: the crop is what shows
      { target: SCROLLER, props: ['box-height'], apple: ['store', 'shelf.crop'] },
      { target: SCROLLER, props: ['padding-bottom'], golden: 'the 22px of apple.com\'s foot in sight: its scroller\'s 40px under a crop 38px below its cards (store:shelf.crop, .platter)' },
      { target: SCROLLER, props: ['scrollbar-width'], golden: 'a thin scrollbar with no paddles, where apple.com\'s crop hides it, as it draws paddles: with neither, a mouse with no sideways wheel cannot reach the rest of the row, nor can anyone see there is more' },
      // apple.com's start inset is its header's margin, its end the last item's padding
      { target: SCROLLER, props: ['padding-left', 'scroll-padding-left'], as: { 'padding-left': 'margin-left', 'scroll-padding-left': 'margin-left' }, apple: ['store', 'shelf.header'] },
      { target: SCROLLER, props: ['padding-right'], apple: ['store', 'shelf.card-last'] },
      { target: SCROLLER, props: ['column-gap'], as: { 'column-gap': 'margin-right' }, apple: ['store', 'shelf.card'] },
      { target: `${SCROLLER} > li`, props: ['scroll-snap-align', 'scroll-snap-stop'], apple: ['store', 'shelf.item'] },
      { target: card(1), props: ['box-left'], apple: ['store', 'shelf.card'] },
      { target: card(2), props: ['box-left'], apple: ['store', 'shelf.card-second'] },
    ],
  },
  {
    // not 1600: there four cards scroll less than a card
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
      { target: '.cui-shelf-paddle:not(:disabled)', props: ['position', 'top', 'right', 'opacity'], apple: ['store', 'shelf.paddle'] },
      { target: '.cui-shelf-paddle:disabled', props: ['position', 'top', 'left'], apple: ['store', 'shelf.paddle-previous'] },
      { target: '.cui-shelf-paddle:not(:disabled)', props: ['margin-top', 'box-height', 'box-width'], golden: 'the gallery\'s 36px paddle, centered where apple.com\'s is, not its 56px (store:shelf.paddle): a large circle over the content competes with it (DESIGN.md decision 25)' },
      { target: '.cui-shelf-paddle:not(:disabled)', props: ['transform', ...MOTION, 'transition-delay'], golden: 'no transform; a gallery paddle\'s 0.1s dim: apple.com\'s scale and 0.5s fade from 0.1s (store:shelf.paddle) are motion past what shows the change of state (DESIGN.md decision 25)' },
      { target: '.cui-shelf-paddle:disabled', props: ['opacity', 'margin-top', 'transform', ...MOTION, 'transition-delay'], golden: 'at an end, dimmed to a gallery paddle\'s .42 in its 0.1s and in its place: apple.com\'s vanishes in 0.7s from 0.1s, shrinking to 0.8 (store:shelf.paddle-previous), so the row\'s other paddle jumps out of a pair (DESIGN.md decision 25)' },
      { target: SCROLLER, props: ['scrollbar-width'], golden: 'hidden with paddles beside it, as apple.com\'s crop hides it' },
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
    name: 'touch',
    html: ex('shelf'),
    touch: true,
    checks: [
      { target: SCROLLER, props: ['scrollbar-width'], golden: 'a touch screen swipes: hidden still, as apple.com\'s crop hides it' },
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
