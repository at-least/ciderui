// The posts against apple.com's Newsroom: the archive (newsroom-archive:*:
// a month's heading, its rows' room, rule, layout, no hover; their picture's
// room, category, headline and date) and, data-size="sm", the related
// stories under an article (newsroom-article:related.*); at the kit's widths
// and both sides of 735 and 1069. Golden, with their why: the widths and
// places (an apple.com check reads no width), the pictures' sizes, the
// letter-spacing, the colors in a dark block.

import { examples, SIZE } from '../../test/harness.mjs';

const ex = examples(import.meta.url);
const WIDTHS = [1600, 1200, 1069, 1068, 900, 735, 734, 390];
const TRACKING = 'the system font\'s fitted letter-spacing (tokens.css), not apple.com\'s';
const ROW = ['padding-top', 'padding-bottom', 'border-top-width', 'border-top-style', 'border-top-color', 'border-bottom-width', 'display', 'flex-direction', 'align-items', 'cursor'];

const tracking = ['.cui-posts-heading', '.cui-post-category', '.cui-post-headline', '.cui-post-date']
  .map(target => ({ target, props: ['letter-spacing'], golden: TRACKING }));

export const cases = [
  {
    name: 'archive',
    html: ex('archive'),
    widths: WIDTHS,
    checks: [
      { target: '.cui-posts-heading', props: [...SIZE, 'color', 'margin-top', 'margin-bottom'], apple: ['newsroom-archive', 'month.title'] },
      { target: '.cui-post', props: ROW, apple: ['newsroom-archive', 'row'] },
      { target: '.cui-post', state: 'hover', props: ['background-color', 'border-top-color'], apple: ['newsroom-archive', 'row', 'hover'] },
      { target: 'li:last-child .cui-post', props: ['border-bottom-width', 'padding-bottom'], apple: ['newsroom-archive', 'row.last'] },
      { target: '.cui-post-media', props: ['flex-shrink'], apple: ['newsroom-archive', 'row.img'] },
      { target: '.cui-post-media', props: ['margin-right'], as: { 'margin-right': 'margin-left' }, apple: ['newsroom-archive', 'row.copy'] },
      { target: '.cui-post-category', props: [...SIZE, 'color', 'text-transform'], apple: ['newsroom-archive', 'row.category'] },
      { target: '.cui-post-headline', props: [...SIZE, 'color', 'margin-top', 'margin-bottom'], apple: ['newsroom-archive', 'row.headline'] },
      { target: '.cui-post-headline', state: 'hover', on: '.cui-post', props: ['color'], apple: ['newsroom-archive', 'row.headline', 'hover'] },
      { target: '.cui-post-date', props: [...SIZE, 'color', 'margin-top'], apple: ['newsroom-archive', 'row.date'] },
      { target: '.cui-posts', props: ['box-width', 'box-left'], golden: 'apple.com\'s archive: 980px from 1069px, 692px from 735px, 87.5% below (newsroom-archive:results), centered' },
      { target: '.cui-post-media', props: ['box-width', 'box-height', 'object-fit'], golden: 'apple.com\'s picture: 295 by 166 from 1069px, 264 by 148.5 from 735px, 105 by 59 below (newsroom-archive:row.img), covering' },
      ...tracking,
    ],
  },
  {
    name: 'related',
    html: ex('related'),
    widths: WIDTHS,
    checks: [
      { target: '.cui-posts', props: ['padding-top'], apple: ['newsroom-article', 'related.list'] },
      { target: '.cui-posts-heading', props: [...SIZE, 'color', 'margin-top', 'margin-bottom', 'text-align'], apple: ['newsroom-article', 'related.heading'] },
      { target: '.cui-post', props: ['padding-top', 'padding-bottom', 'border-top-width', 'border-bottom-width', 'border-bottom-style', 'border-bottom-color'], apple: ['newsroom-article', 'related.item'] },
      { target: '.cui-post', props: ['display', 'flex-direction', 'align-items', 'cursor'], apple: ['newsroom-article', 'related.link'] },
      { target: '.cui-post', state: 'hover', props: ['text-decoration-line', 'cursor'], apple: ['newsroom-article', 'related.link', 'hover'] },
      { target: '.cui-post-media', props: ['margin-right'], as: { 'margin-right': 'padding-right' }, apple: ['newsroom-article', 'related.img'] },
      { target: '.cui-post-category', props: [...SIZE, 'color', 'text-transform', 'margin-bottom'], apple: ['newsroom-article', 'related.category'] },
      { target: '.cui-post-headline', props: [...SIZE, 'color', 'margin-top', 'margin-bottom'], apple: ['newsroom-article', 'related.title'] },
      { target: '.cui-post-headline', state: 'hover', on: '.cui-post', props: ['color'], apple: ['newsroom-article', 'related.title', 'hover'] },
      { target: '.cui-post-date', props: [...SIZE, 'color', 'margin-top'], apple: ['newsroom-article', 'related.date'] },
      { target: '.cui-posts', props: ['box-width', 'box-left'], golden: 'an article\'s text column: apple.com\'s 653px, 576px, 87.5% of 414px at most (newsroom-article:related.list), centered' },
      { target: '.cui-post-media', props: ['box-width', 'box-height', 'object-fit'], golden: 'apple.com\'s picture: 171 by 96 from 735px, 93 by 93 below (newsroom-article:related.img), covering' },
      ...tracking,
    ],
  },
  {
    name: 'dark',
    html: `<div class="cui-dark">${ex('archive')}</div>`,
    widths: [1600, 390],
    checks: [
      { target: '.cui-post', props: ['border-top-color', 'color'], golden: 'the dark divider (#424245) and text: apple.com\'s archive is light' },
      { target: '.cui-post-category', props: ['color'], golden: 'the dark secondary text' },
    ],
  },
];
