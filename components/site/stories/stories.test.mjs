// The stories against apple.com's Newsroom index (newsroom-index:*): the
// feed's ground and room, its heading, the grid's gaps, the tiles one, two
// and three across (the example's first, third and fourth, as the probes'):
// their corners, ground, height, no hover; their text's room, category,
// headline and date; at the kit's widths and both sides of 735 and 1069.
// Golden, with their why: the widths and places (an apple.com check reads
// no width), the grid's columns and spans, the picture's shape, the
// letter-spacing, the category uppercase, the colors in a dark block.

import { examples, SIZE } from '../../../test/harness.mjs';

const ex = examples(import.meta.url);
const PAGE = 'newsroom-index';
const WIDTHS = [1600, 1200, 1069, 1068, 900, 735, 734, 390];
const TRACKING = 'the system font\'s fitted letter-spacing (tokens.css), not apple.com\'s';
// the example's stories in the places the probes read: the first (one
// across), the third (two), the fourth (three)
const TILES = [['first', 1], ['third', 3], ['fourth', 4]];
const tile = (n) => `.cui-stories-list > li:nth-child(${n}) > .cui-story`;

export const cases = [
  {
    name: 'stories',
    html: ex('stories'),
    widths: WIDTHS,
    checks: [
      { target: '.cui-stories', props: ['padding-top', 'padding-bottom', 'background-color'], apple: [PAGE, 'feed'] },
      { target: '.cui-stories-heading', props: [...SIZE, 'color', 'margin-top', 'margin-bottom'], apple: [PAGE, 'feed.head'] },
      { target: '.cui-stories-list', props: ['display', 'row-gap', 'column-gap', 'list-style-type', 'padding-left'], apple: [PAGE, 'tiles'] },
      ...TILES.flatMap(([probe, n]) => [
        { target: tile(n), props: ['border-top-left-radius', 'background-color', 'box-shadow', 'min-height', 'overflow', 'padding-top', 'padding-left'], apple: [PAGE, `tile.${probe}`] },
        { target: tile(n), state: 'hover', props: ['background-color', 'box-shadow'], apple: [PAGE, `tile.${probe}`, 'hover'] },
        { target: `${tile(n)} .cui-story-content`, props: ['padding-top', 'padding-right', 'padding-bottom', 'padding-left', 'display', 'flex-direction', 'justify-content'], apple: [PAGE, `tile.${probe}-description`] },
        { target: `${tile(n)} .cui-story-category`, props: [...SIZE, 'color'], apple: [PAGE, probe === 'first' ? 'tile.first-eyebrow' : `tile.${probe}-category`] },
        { target: `${tile(n)} .cui-story-headline`, props: [...SIZE, 'color', 'margin-top', 'margin-bottom'], apple: [PAGE, `tile.${probe}-headline`] },
        { target: `${tile(n)} .cui-story-date`, props: [...SIZE, 'color', 'margin-top'], apple: [PAGE, `tile.${probe}-date`] },
        { target: tile(n), props: ['box-width', 'box-left', 'box-height'], golden: 'apple.com\'s tile\'s place: 980px one across, 470px two across (1px in from its columns), its four columns three across (302.66px; apple.com\'s a fixed 303px) from 1069px, 328px from 735px (3px in), its column below; its height its picture\'s and its text\'s, or its least (newsroom-index:tile.*)' },
        { target: `${tile(n)} .cui-story-media`, props: ['box-width', 'box-height', 'object-fit'], golden: 'the picture 16:9 at its width (apple.com\'s 643 by 362 beside the text one across, newsroom-index:tile.first-media), covering' },
        { target: `${tile(n)} .cui-story-category`, props: ['letter-spacing', 'text-transform'], golden: `${TRACKING}; uppercase: apple.com writes its categories in capitals` },
        { target: `${tile(n)} .cui-story-headline`, props: ['letter-spacing'], golden: TRACKING },
        { target: `${tile(n)} .cui-story-date`, props: ['letter-spacing'], golden: TRACKING },
      ]),
      { target: tile(1), props: ['display', 'grid-template-columns'], golden: 'one across, its picture 643px beside its text from 1069px (newsroom-index:tile.first, grid 643px 337px), its picture over its text below' },
      { target: '.cui-stories-list', props: ['grid-template-columns', 'flex-wrap'], golden: 'twelve columns from 735px (newsroom-index:tiles), a wrapping row below' },
      ...[1, 3, 4].map(n => ({ target: `.cui-stories-list > li:nth-child(${n})`, props: ['grid-column-start', 'grid-column-end', 'box-width'], golden: 'its columns: one across all twelve, two across six, three across four from 1069px; six of them, two across, from 735px; the whole row below' })),
      { target: '.cui-stories-content', props: ['box-width', 'box-left'], golden: 'apple.com\'s feed: 980px from 1069px, 692px from 735px, 87.5% at most 366px below (newsroom-index:feed.content), centered' },
      { target: '.cui-stories-heading', props: ['letter-spacing'], golden: TRACKING },
    ],
  },
  {
    name: 'dark',
    html: `<div class="cui-dark">${ex('stories')}</div>`,
    widths: [1600, 390],
    checks: [
      { target: '.cui-stories', props: ['background-color', 'color'], golden: 'the dark alternate ground and text: apple.com\'s feed is light' },
      { target: tile(1), props: ['background-color'], golden: 'the dark page\'s ground, a tile on the dark alternate' },
      { target: `${tile(1)} .cui-story-category`, props: ['color'], golden: 'the dark secondary text' },
    ],
  },
];
