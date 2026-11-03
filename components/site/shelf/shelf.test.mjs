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
      { target: SCROLLER, props: ['scrollbar-width'], golden: 'hidden, as apple.com\'s crop hides it' },
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
      { target: '.cui-shelf-paddle:not(:disabled)', props: ['position', 'top', 'right', 'margin-top', 'opacity', 'box-height'], apple: ['store', 'shelf.paddle'] },
      { target: '.cui-shelf-paddle:disabled', props: ['position', 'top', 'left', 'margin-top', 'opacity'], apple: ['store', 'shelf.paddle-previous'] },
      { target: '.cui-shelf-paddle:not(:disabled)', props: ['transform', ...MOTION, 'transition-delay'], golden: 'no transform; its colors as a gallery paddle\'s, in 0.1s (gallery.css: apple.com\'s shelf paddle changes them at once), its opacity in apple.com\'s 0.5s from 0.1s (store:shelf.paddle): apple.com\'s scale moves too, motion past what shows the change of state (DESIGN.md decision 25)' },
      { target: '.cui-shelf-paddle:disabled', props: ['transform', ...MOTION, 'transition-delay'], golden: 'faded out at its size, its opacity in apple.com\'s 0.7s from 0.1s, its colors as a gallery paddle\'s in 0.1s: apple.com\'s shrinks it to 0.8 as well (store:shelf.paddle-previous, DESIGN.md decision 25)' },
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
