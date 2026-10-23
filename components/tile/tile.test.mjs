// The tile against apple.com's home page: a promo tile (home:tile.light,
// .promo-*: its ground, size, padding and layout, its headline, subhead and
// buttons), a full-width one (home:hero.*: the same a step larger, its
// callout, its call-to-action buttons), a dark one (home:tile.dark*), the
// type through the widths (home:type.tile-*, .promo-*), the layers (its
// picture, its link, its buttons above it), the product page's rounded
// corners. Golden, with their why: the letter-spacing, a promo tile's
// callout, the buttons' row, the layers' order.

import { examples, PADDING, SIZE } from '../../test/harness.mjs';

const ex = examples(import.meta.url);
const PAD = ['padding-top', 'padding-bottom'];
const media = (html) => html.replace('<a class="cui-tile-link"', '<img class="cui-tile-media" src="data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHdpZHRoPSIxMjgwIiBoZWlnaHQ9IjcyMCIvPg==" alt=""><a class="cui-tile-link"');
const dark = (html) => html.replace('class="cui-tile"', 'class="cui-tile cui-dark"');
const promoCallout = ex('tile').replace('<div class="cui-tile-ctas">', '<p class="cui-tile-callout">Trade in your phone.</p><div class="cui-tile-ctas">');

export const cases = [
  {
    name: 'promo',
    html: ex('tile'),
    checks: [
      { target: '.cui-tile', props: ['background-color', 'min-height', ...PAD, 'display', 'justify-content', 'overflow', 'text-align', 'position'], apple: ['home', 'tile.light'] },
      { target: '.cui-tile-content', props: ['display', 'flex-direction', 'align-items'], apple: ['home', 'tile.promo-content'] },
      { target: '.cui-tile-headline', props: [...SIZE, 'color', 'margin-top', 'margin-bottom'], apple: ['home', 'tile.promo-headline'] },
      { target: '.cui-tile-headline', props: SIZE, apple: ['home', 'type.promo-headline'] },
      { target: '.cui-tile-subhead', props: [...SIZE, 'color', 'margin-top'], apple: ['home', 'tile.promo-subhead'] },
      { target: '.cui-tile-subhead', props: SIZE, apple: ['home', 'type.promo-subhead'] },
      { target: '.cui-tile-ctas', props: ['column-gap', 'justify-content', 'margin-top'], apple: ['home', 'tile.promo-ctas'] },
      { target: '.cui-tile-ctas .cui-btn', props: [...SIZE, 'padding-top', 'padding-left', 'min-width', 'box-height'], apple: ['home', 'tile.promo-button'] },
      { target: '.cui-tile-link', props: ['position'], apple: ['home', 'tile.link'] },
      { target: '.cui-tile-link', props: ['box-height'], apple: ['home', 'tile.light'] },
      { target: '.cui-tile-headline', props: ['letter-spacing'], golden: 'the system font\'s fitted letter-spacing (tokens.css), not apple.com\'s' },
      { target: '.cui-tile-ctas', props: ['display'], golden: 'a flex row of the gaps apple.com\'s grid has (home:tile.promo-ctas)' },
    ],
  },
  {
    name: 'promo-callout',
    html: promoCallout,
    checks: [
      { target: '.cui-tile-callout', props: SIZE, apple: ['home', 'type.promo-callout'] },
      { target: '.cui-tile-callout', props: ['color', 'margin-top'], golden: 'the secondary text, 8px below the subhead: apple.com\'s promo tiles measured carry no callout; a full tile\'s, below 735px (home:hero.callout)' },
    ],
  },
  {
    name: 'hero',
    html: ex('hero'),
    checks: [
      { target: '.cui-tile', props: ['background-color', 'min-height', ...PAD, 'box-height'], apple: ['home', 'hero.tile'] },
      { target: '.cui-tile-content', props: ['display', 'flex-direction', 'align-items', 'text-align'], apple: ['home', 'hero.content'] },
      { target: '.cui-tile-headline', props: [...SIZE, 'color', 'margin-bottom'], apple: ['home', 'hero.headline'] },
      { target: '.cui-tile-headline', props: SIZE, apple: ['home', 'type.tile-headline'] },
      { target: '.cui-tile-subhead', props: [...SIZE, 'color', 'margin-top'], apple: ['home', 'hero.subhead'] },
      { target: '.cui-tile-callout', props: [...SIZE, 'color', 'margin-top'], apple: ['home', 'hero.callout'] },
      { target: '.cui-tile-ctas', props: ['column-gap', 'justify-content', 'margin-top'], apple: ['home', 'hero.ctas'] },
      { target: '.cui-tile-ctas .cui-btn', props: [...SIZE, ...PADDING, 'background-color', 'color', 'border-top-left-radius', 'box-height'], apple: ['home', 'button'] },
      { target: '.cui-tile-ctas .cui-btn[data-variant="secondary"]', props: [...SIZE, ...PADDING, 'color', 'border-top-color', 'box-height'], apple: ['home', 'button.secondary'] },
    ],
  },
  {
    name: 'dark',
    html: dark(ex('hero')),
    checks: [
      { target: '.cui-tile', props: ['background-color', 'min-height', ...PAD], apple: ['home', 'tile.dark'] },
      { target: '.cui-tile-headline', props: [...SIZE, 'color'], apple: ['home', 'tile.dark-headline'] },
      { target: '.cui-tile-subhead', props: [...SIZE, 'color'], apple: ['home', 'tile.dark-subhead'] },
      { target: '.cui-tile-ctas .cui-btn', props: ['background-color', 'color'], apple: ['home', 'tile.dark-button'] },
      { target: '.cui-tile-ctas .cui-btn[data-variant="secondary"]', props: ['color', 'border-top-color'], apple: ['home', 'tile.dark-button-secondary'] },
    ],
  },
  {
    name: 'layers',
    html: media(ex('tile')),
    widths: [1200, 390],
    checks: [
      { target: '.cui-tile-media', props: ['position', 'object-fit', 'box-height'], apple: ['home', 'tile.media-image'] },
      { target: '.cui-tile-media', props: ['position', 'top', 'left', 'box-height'], apple: ['home', 'tile.media'] },
      { target: '.cui-tile-media', props: ['box-width'], golden: 'the picture across the whole tile, its own cut (apple.com\'s, home:tile.media)' },
      { target: '.cui-tile-link', props: ['z-index', 'box-width'], golden: 'the link over the picture and the copy, the whole tile (apple.com\'s at 3, its copy at 2, home:tile.link)' },
      { target: '.cui-tile-link', state: 'focus-visible', props: ['outline-offset'], golden: 'the one ring moved in, 3px inside the tile\'s edge: the link, the whole tile, takes the focus (apple.com\'s none, tabindex="-1": home:tile.link), and the tile clips what is outside it (overflow: clip, apple.com\'s)' },
      { target: '.cui-tile-ctas', props: ['position', 'z-index'], golden: 'the buttons over the link (apple.com\'s at 3 and 4)' },
    ],
  },
  {
    name: 'rounded',
    html: ex('tile').replace('class="cui-tile"', 'class="cui-tile cui-tile-rounded"'),
    widths: [1200, 390],
    checks: [
      { target: '.cui-tile', props: ['border-top-left-radius', 'overflow'], apple: ['product', 'tile.rounded'] },
      { target: '.cui-tile-link', props: ['border-top-left-radius'], golden: 'the tile\'s corners, which its ring, moved in, follows' },
    ],
  },
];
