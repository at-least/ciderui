// The card against apple.com's store: a light card with a tag (store:card,
// .info, .tag, .header, .link, .desc, .description, .price: its size,
// ground, corners, shadow and motion, its text's padding, type, colors and
// spacing), hovered and with the focus in it (store:card's hover and
// focus-within), a plain eyebrow (store:card.eyebrow), a headline with no
// eyebrow above it (store:card.header-first), a dark card
// (store:card.dark*), its picture (store:card.media); on a touch screen, not
// lifting, hovered or with the focus in it (touch-store:card). Golden, with their
// why: the letter-spacing, the widths, what moves, the whole card its link,
// the layers.

import { examples, PADDING, SIZE } from '../../test/harness.mjs';

const ex = examples(import.meta.url);
const LIFT = ['box-shadow', 'transform'];
const TRACKING = 'the system font\'s fitted letter-spacing (tokens.css), not apple.com\'s';
const dark = (html) => html.replace('class="cui-card"', 'class="cui-card cui-dark"');
const media = (html) => html.replace('<div class="cui-card">', '<div class="cui-card"><img class="cui-card-media" src="data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHdpZHRoPSIxMjgwIiBoZWlnaHQ9IjcyMCIvPg==" alt="">');
const headlineFirst = ex('card').replace(/\s*<p class="cui-card-eyebrow"[^\n]*/, '');

export const cases = [
  {
    name: 'card',
    html: ex('card'),
    checks: [
      { target: '.cui-card', props: ['background-color', 'color', 'border-top-left-radius', 'box-shadow', 'transform', 'overflow', 'position', 'box-height', 'transition-duration', 'transition-timing-function'], apple: ['store', 'card'] },
      { target: '.cui-card', props: ['transition-property'], golden: 'its shadow and its scale: apple.com\'s all would also move its size as the window crosses 735px (apple.com\'s steps with the screen)' },
      { target: '.cui-card', props: ['box-width'], golden: 'apple.com\'s 400px, 309px on a phone (store:card\'s box): a width no text sets' },
      { target: '.cui-card-content', props: [...PADDING, 'box-height'], apple: ['store', 'card.info'] },
      { target: '.cui-card-content', props: ['box-width'], golden: 'apple.com\'s 396px, 340px and its padding, and on a phone the whole card (store:card.info\'s box)' },
      { target: '.cui-card-eyebrow', props: [...SIZE, 'color', 'margin-top', 'margin-bottom', 'padding-top', 'padding-bottom', 'text-transform'], apple: ['store', 'card.tag'] },
      { target: '.cui-card-headline', props: [...SIZE, 'color', 'margin-top', 'margin-bottom', 'padding-top', 'padding-bottom'], apple: ['store', 'card.header'] },
      { target: '.cui-card-link', props: [...SIZE, 'color', 'display', 'text-decoration-line', 'cursor'], apple: ['store', 'card.link'] },
      { target: '.cui-card-link', state: 'hover', props: ['color', 'text-decoration-line'], apple: ['store', 'card.link', 'hover'] },
      { target: '.cui-card-description', props: [...SIZE, 'color', 'display'], apple: ['store', 'card.description'] },
      // apple.com's description is a block in its desc, which is 10px down
      { target: '.cui-card-description', props: ['margin-top'], as: { 'margin-top': 'padding-top' }, apple: ['store', 'card.desc'] },
      { target: '.cui-card-price', props: [...SIZE, 'color'], apple: ['store', 'card.price'] },
      // apple.com's price is after its description's 6px
      { target: '.cui-card-price', props: ['margin-top'], as: { 'margin-top': 'padding-bottom' }, apple: ['store', 'card.description'] },
      { target: '.cui-card-eyebrow', props: ['letter-spacing'], golden: TRACKING },
      { target: '.cui-card-headline', props: ['letter-spacing'], golden: TRACKING },
      { target: '.cui-card-description', props: ['letter-spacing'], golden: TRACKING },
      { target: '.cui-card-price', props: ['letter-spacing'], golden: TRACKING },
      { target: '.cui-card', props: ['hit'], golden: 'the whole card its headline\'s link: the pointer finds the link at its corners and its middle (apple.com\'s script follows its link from anywhere on it)' },
    ],
  },
  {
    name: 'lift',
    html: ex('card'),
    checks: [
      { target: '.cui-card', state: 'hover', props: LIFT, apple: ['store', 'card', 'hover'] },
      { target: '.cui-card', state: 'focus-within', props: LIFT, apple: ['store', 'card', 'focus-within'] },
    ],
  },
  {
    // Tab from the page's start: the link takes the focus, the card lifts
    name: 'focus',
    html: ex('card'),
    widths: [1200],
    press: 'Tab',
    wait: 1000,
    checks: [
      { target: '.cui-card-link', props: ['focused'], golden: 'the headline\'s link, the card\'s one stop in the focus order' },
      { target: '.cui-card', props: LIFT, apple: ['store', 'card', 'focus-within'] },
    ],
  },
  {
    name: 'touch',
    html: ex('card'),
    touch: true,
    checks: [
      // apple.com's lifts only with no touch (html.no-touch, its script's)
      { target: '.cui-card', state: 'hover', props: LIFT, apple: ['touch-store', 'card', 'hover'] },
      { target: '.cui-card', state: 'focus-within', props: LIFT, apple: ['touch-store', 'card', 'focus-within'] },
    ],
  },
  {
    name: 'headline-first',
    html: headlineFirst,
    checks: [
      { target: '.cui-card-headline', props: [...SIZE, 'padding-top'], apple: ['store', 'card.header-first'] },
    ],
  },
  {
    name: 'eyebrow',
    html: ex('eyebrow'),
    checks: [
      { target: '.cui-card-eyebrow', props: [...SIZE, 'color', 'margin-top', 'margin-bottom', 'padding-top', 'padding-bottom', 'text-transform'], apple: ['store', 'card.eyebrow'] },
    ],
  },
  {
    name: 'dark',
    html: dark(ex('card')),
    checks: [
      { target: '.cui-card', props: ['background-color', 'border-top-left-radius', 'box-shadow', 'transform', 'overflow', 'box-height'], apple: ['store', 'card.dark'] },
      { target: '.cui-card', state: 'hover', props: LIFT, apple: ['store', 'card.dark', 'hover'] },
      { target: '.cui-card-eyebrow', props: ['color'], apple: ['store', 'card.dark-tag'] },
      { target: '.cui-card-headline', props: ['color'], apple: ['store', 'card.dark-header'] },
      { target: '.cui-card-link', props: ['color'], apple: ['store', 'card.dark-link'] },
      { target: '.cui-card-description', props: ['color'], apple: ['store', 'card.dark-desc'] },
      { target: '.cui-card-price', props: ['color'], apple: ['store', 'card.dark-price'] },
    ],
  },
  {
    name: 'dark-eyebrow',
    html: dark(ex('eyebrow')),
    checks: [
      { target: '.cui-card-eyebrow', props: ['color'], apple: ['store', 'card.dark-eyebrow'] },
    ],
  },
  {
    // apple.com's phone card had no picture of its own when measured
    name: 'layers',
    html: media(ex('card')),
    widths: [1600, 1200, 900],
    checks: [
      { target: '.cui-card-media', props: ['display', 'object-fit', 'box-height'], apple: ['store', 'card.media'] },
      { target: '.cui-card-media', props: ['position', 'z-index', 'box-width'], golden: 'the picture laid under the text across the whole card (apple.com\'s in the card\'s flow, its text laid over it: store:card.media, .info)' },
      { target: '.cui-card', props: ['isolation'], golden: 'the card its own stacking: its picture under its text, over its ground' },
      { target: '.cui-card', props: ['hit'], golden: 'its picture under the link: the whole card still its link' },
    ],
  },
];
