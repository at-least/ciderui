// The product nav against apple.com's store (store:productnav*): the
// shelf's room below, its row's height in sight, a card's size, padding,
// image and name, the first two cards' places. Golden, with their why: the
// letter-spacing, a card's width, its foot on a phone, the card the link.

import { examples } from '../../test/harness.mjs';

const ex = examples(import.meta.url);
const TYPE = ['font-size', 'line-height', 'font-weight'];
const ITEM = '.cui-shelf-scroller > li:first-child .cui-productnav-item';
const card = (n) => `.cui-shelf-scroller > li:nth-child(${n}) .cui-productnav-item`;

export const cases = [
  {
    name: 'productnav',
    html: ex('productnav'),
    checks: [
      { target: '.cui-productnav', props: ['padding-top', 'padding-bottom'], apple: ['store', 'productnav'] },
      { target: '.cui-shelf-scroller', props: ['column-gap'], as: { 'column-gap': 'margin-right' }, apple: ['store', 'productnav.card'] },
      { target: ITEM, props: ['padding-top', 'padding-right', 'padding-bottom', 'padding-left'], apple: ['store', 'productnav.content'] },
      { target: ITEM, props: ['box-height', 'box-left'], apple: ['store', 'productnav.card'] },
      { target: card(2), props: ['box-left'], apple: ['store', 'productnav.card-second'] },
      { target: `${ITEM} .cui-productnav-image`, props: ['height', 'box-height'], apple: ['store', 'productnav.image'] },
      { target: `${ITEM} .cui-productnav-image`, props: ['margin-bottom'], as: { 'margin-bottom': 'padding-bottom' }, apple: ['store', 'productnav.image-box'] },
      { target: `${ITEM} .cui-productnav-name`, props: [...TYPE, 'color', 'display', 'text-align'], apple: ['store', 'productnav.title'] },
      { target: `${ITEM} .cui-productnav-image`, props: ['box-width'], golden: 'its own shape at its height: 120px, 92.3px for a 200 by 130 image (store:productnav.image)' },
      { target: ITEM, props: ['box-width'], golden: 'its image and its padding: apple.com\'s 136px, 100.3px (store:productnav.card)' },
      { target: '.cui-shelf-row', props: ['box-height'], golden: 'its row in sight: apple.com\'s crop, 186px, and 164px on a phone for apple.com\'s 163px, which cuts its row\'s last pixel (store:productnav.crop)' },
      { target: '.cui-shelf-scroller', props: ['padding-bottom'], golden: 'room under the cards for a focused card\'s ring, which the scroller clips (on a phone 5px, pulled back by its margin, none of it in sight: the row\'s crop above)' },
      { target: ITEM, props: ['text-decoration-line', 'display'], golden: 'the card is the link, no underline (apple.com\'s script follows its name\'s link)' },
      { target: `${ITEM} .cui-productnav-name`, props: ['letter-spacing'], golden: 'the system font\'s fitted letter-spacing (tokens.css), not apple.com\'s' },
    ],
  },
];
