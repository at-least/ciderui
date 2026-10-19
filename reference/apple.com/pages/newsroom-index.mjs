// apple.com's Newsroom, its index: the latest stories, a grid of white
// rounded tiles on the alternate ground (one across, two, three), each a
// link with its picture, its category, its headline and its date. The
// event banner above (#takeover) and the modules a script draws below are
// not read. The page's script gives the tiles their classes by the window
// (from 735px to 1068px every tile is tile-2up, two across), so a tile is
// picked by its place in the grid: the first (one across from 1069px and
// below 735px), the third (two across), the fourth (three across).
import { TYPE, BOX, FLEX, TEXT, LINK } from './common.mjs';

const MEDIA = ['aspect-ratio', 'overflow', 'border-top-left-radius', 'border-bottom-left-radius'];

export default {
  url: 'https://www.apple.com/newsroom/',
  widths: [1600, 1200, 1069, 1068, 900, 735, 734, 390],
  probes: [
    { name: 'feed', select: '.everydayfeed', props: [...BOX, 'padding-top', 'padding-bottom'] },
    { name: 'feed.content', select: '.everydayfeed .section-content', props: ['width', 'max-width', 'margin-left', 'margin-right', 'padding-left', 'padding-right'], box: true },
    { name: 'feed.head', select: '.everydayfeed .section-head', props: [...TYPE, 'margin-top', 'margin-bottom'] },
    { name: 'tiles', select: '.everydayfeed .section-tiles', props: [...FLEX, 'grid-template-columns', 'list-style-type', 'padding-left', 'margin-top', 'margin-bottom'], box: true },
    { name: 'tiles.item', select: '.everydayfeed .section-tiles > li', props: ['grid-column-start', 'grid-column-end', 'display', 'margin-top', 'margin-bottom', 'flex-basis', 'width'], box: true },
    // the first tile: one across, its picture beside its text
    { name: 'tile.first', select: '.everydayfeed li:nth-child(1) > a', props: [...BOX, ...TEXT, 'display', 'grid-template-columns', 'column-gap', 'transform', 'transition-property', 'transition-duration', 'overflow'], box: true, states: ['hover', 'focus-visible'] },
    { name: 'tile.first-media', select: '.everydayfeed li:nth-child(1) > a .tile__media', props: MEDIA, box: true },
    { name: 'tile.first-description', select: '.everydayfeed li:nth-child(1) > a .tile__description', props: [...BOX, 'padding-left', 'padding-right', 'display', 'flex-direction', 'justify-content'], box: true },
    { name: 'tile.first-eyebrow', select: '.everydayfeed li:nth-child(1) > a :is(.tile__hero--eyebrow, .tile__category)', props: [...TYPE, ...TEXT, 'text-transform'] },
    { name: 'tile.first-headline', select: '.everydayfeed li:nth-child(1) > a .tile__headline', props: [...TYPE, 'margin-top', 'margin-bottom'] },
    { name: 'tile.first-date', select: '.everydayfeed li:nth-child(1) > a .tile__timestamp', props: [...TYPE, 'margin-top', 'margin-bottom'] },
    // the third: two across, its picture over its text
    { name: 'tile.third', select: '.everydayfeed li:nth-child(3) > a', props: [...BOX, 'display', 'flex-direction', 'transform', 'overflow'], box: true, states: ['hover'] },
    { name: 'tile.third-media', select: '.everydayfeed li:nth-child(3) > a .tile__media', props: MEDIA, box: true },
    { name: 'tile.third-description', select: '.everydayfeed li:nth-child(3) > a .tile__description', props: [...BOX, 'padding-left', 'padding-right', 'display', 'flex-direction', 'justify-content', 'flex-grow'], box: true },
    { name: 'tile.third-category', select: '.everydayfeed li:nth-child(3) > a .tile__category', props: [...TYPE, ...TEXT, 'text-transform'] },
    { name: 'tile.third-headline', select: '.everydayfeed li:nth-child(3) > a .tile__headline', props: [...TYPE, 'margin-top', 'margin-bottom'] },
    { name: 'tile.third-date', select: '.everydayfeed li:nth-child(3) > a .tile__timestamp', props: [...TYPE, 'margin-top', 'margin-bottom'] },
    // the fourth: three across
    { name: 'tile.fourth', select: '.everydayfeed li:nth-child(4) > a', props: [...BOX, 'display', 'flex-direction', 'transform', 'overflow'], box: true },
    { name: 'tile.fourth-media', select: '.everydayfeed li:nth-child(4) > a .tile__media', props: MEDIA, box: true },
    { name: 'tile.fourth-description', select: '.everydayfeed li:nth-child(4) > a .tile__description', props: [...BOX, 'padding-left', 'padding-right', 'flex-grow'], box: true },
    { name: 'tile.fourth-category', select: '.everydayfeed li:nth-child(4) > a .tile__category', props: [...TYPE, ...TEXT, 'text-transform'] },
    { name: 'tile.fourth-headline', select: '.everydayfeed li:nth-child(4) > a .tile__headline', props: [...TYPE, 'margin-top', 'margin-bottom'] },
    { name: 'tile.fourth-date', select: '.everydayfeed li:nth-child(4) > a .tile__timestamp', props: [...TYPE, 'margin-top', 'margin-bottom'] },
    { name: 'tile.link', select: '.everydayfeed li:nth-child(4) > a', props: LINK },
  ],
};
