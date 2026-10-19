// apple.com's Newsroom, its archive: each month a heading over ruled rows,
// each row a link with its picture, its category, its headline and its
// date. The filter bar above (a script's) and the button that loads more
// are not read.
import { TYPE, BOX, FLEX, TEXT, LINK } from './common.mjs';

export default {
  url: 'https://www.apple.com/newsroom/archive/',
  widths: [1600, 1200, 1069, 1068, 900, 735, 734, 390],
  probes: [
    { name: 'results', select: '.results__content', props: ['width', 'max-width', 'margin-left', 'margin-right', 'padding-left', 'padding-right', 'padding-top', 'padding-bottom'], box: true },
    { name: 'month', select: '.result', props: ['margin-top', 'margin-bottom', 'padding-top', 'padding-bottom'] },
    { name: 'month.title', select: '.result__title', props: [...TYPE, 'margin-top', 'margin-bottom'] },
    { name: 'row', select: '.result__item', props: [...BOX, ...FLEX, ...LINK, 'border-bottom-width', 'border-bottom-style', 'border-bottom-color', 'cursor'], box: true, states: ['hover', 'focus-visible'] },
    { name: 'row.last', select: '.result .result__item:last-child', props: ['border-bottom-width', 'border-bottom-style', 'border-bottom-color', 'padding-bottom'] },
    { name: 'row.img', select: '.result__item picture', props: ['width', 'height', 'aspect-ratio', 'border-top-left-radius', 'overflow', 'margin-right', 'flex-shrink'], box: true },
    { name: 'row.copy', select: '.result__item .item__content', props: ['padding-left', 'margin-left', 'flex-grow', 'flex-basis'], box: true },
    { name: 'row.category', select: '.result__item .item__category', props: [...TYPE, ...TEXT, 'text-transform'] },
    { name: 'row.headline', select: '.result__item .item__headline', props: [...TYPE, 'margin-top', 'margin-bottom'], states: ['hover'] },
    { name: 'row.date', select: '.result__item .item__date', props: [...TYPE, 'margin-top', 'margin-bottom'] },
  ],
};
