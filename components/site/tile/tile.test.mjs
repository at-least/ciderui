import { examples } from '../../../test/harness.mjs';

const ex = examples(import.meta.url);
const TILE = '.cui-tile:first-child';
const TITLE = `${TILE} .cui-tile-title`;
const DETAIL = `${TILE} .cui-tile-detail`;
const MEDIA = `${TILE} .cui-tile-media`;

export const cases = [
  {
    name: 'tiles',
    html: ex('tiles'),
    widths: [1600, 1200, 900, 390],
    checks: [
      { target: TILE, props: ['padding-top', 'padding-bottom', 'border-bottom-width', 'border-bottom-style'], apple: ['search-serp', 'tile'] },
      { target: TILE, props: ['border-bottom-color'], golden: 'the kit\'s divider at every width: apple.com\'s is it from 1069px and #d6d6d6 below (search-serp:tile)' },
      { target: TILE, props: ['text-align'], apple: ['search-serp', 'tile.info'] },
      { target: TITLE, props: ['font-size', 'font-weight', 'color'], apple: ['search-serp', 'tile.name'] },
      { target: TITLE, state: 'hover', props: ['color', 'text-decoration-line'], apple: ['search-serp', 'tile.name.link', 'hover'] },
      { target: DETAIL, props: ['font-size', 'color', 'margin-top'], apple: ['search-serp', 'tile.price'] },
      { target: TILE, props: ['padding-left', 'padding-right'], golden: 'the text\'s column, apple.com\'s 26px each side (search-serp:tile.info), the picture in it too' },
      { target: '.cui-tiles', props: ['grid-template-columns'], golden: 'three across from 735px, two below (search-serp:tile, a third at 900 and 1200px, a half at 390px; the step between not read)' },
      { target: MEDIA, props: ['aspect-ratio', 'object-fit', 'margin-bottom'], golden: 'a square picture, as apple.com\'s (search-serp:tile.img), filled where a page\'s picture is not square (apple.com\'s pictures are), 16px over the title' },
      { target: TITLE, pseudo: '::after', props: ['position', 'top', 'bottom', 'left', 'right'], golden: 'the whole tile, picture and all, is the title\'s target: apple.com links its picture and its name apart' },
    ],
  },
  {
    name: 'tiles type',
    html: ex('tiles'),
    widths: [1600, 1200, 900, 390],
    checks: [
      { target: TITLE, props: ['line-height'], golden: 'the ladder\'s 25px (cui-body) from 1069px and 20px (cui-body-reduced) below, where apple.com\'s 14px computes 20.0003px (search-serp:tile.name)' },
    ],
  },
  {
    name: 'tiles zh',
    html: ex('tiles').replace('<ul class', '<ul lang="zh-Hant" class'),
    widths: [900],
    checks: [
      { target: TITLE, props: ['line-height'], golden: 'Traditional Chinese a pixel taller, as cui-body-reduced is' },
    ],
  },
];
