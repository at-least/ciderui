import { examples } from '../../../test/harness.mjs';

const ex = examples(import.meta.url);
const IMG = '.cui-console-thumb:first-child img';

export const cases = [
  {
    name: 'thumbnails',
    html: ex('thumbs'),
    widths: [1200, 390],
    checks: [
      { target: IMG, props: ['border-top-left-radius', 'border-top-width', 'border-top-style', 'box-shadow', 'cursor'], asc: ['media', 'thumb.img'] },
      { target: IMG, props: ['box-height'], golden: 'App Store Connect\'s 240px (media.json thumb.img; an asc check reads no box)' },
      { target: IMG, props: ['box-width'], golden: 'as wide as its image\'s shape, the example\'s 111 by 240: App Store Connect\'s frames are a device\'s shape' },
      { target: '.cui-console-thumbs', props: ['display', 'flex-wrap', 'column-gap', 'row-gap'], golden: 'a row that wraps, 12px between: App Store Connect\'s 6px either side of each (media.json thumb.item)' },
      { target: IMG, props: ['background-color'], golden: 'the console card\'s gray while it loads, where App Store Connect\'s older look wrote rgb(242, 242, 242)' },
      { target: '.cui-console-thumb:first-child .cui-console-thumb-caption', props: ['font-size', 'line-height', 'margin-top', 'color'], golden: 'its name under it in the status\'s line\'s type (11px on 13px, 5px under), the secondary label' },
    ],
  },
  {
    name: 'thumbnails hovered',
    html: ex('thumbs'),
    widths: [1200],
    checks: [
      { target: IMG, props: ['border-top-left-radius', 'box-shadow'], state: 'hover', on: '.cui-console-thumb:first-child > a', asc: ['media', 'thumb.img'] },
    ],
  },
  {
    name: 'thumbnails dark',
    html: `<div class="cui-dark">${ex('thumbs')}</div>`,
    widths: [1200],
    checks: [
      { target: '.cui-console-thumb:first-child .cui-console-thumb-caption', props: ['color'], golden: 'the dark side derived (tokens.css, decision 18)' },
    ],
  },
];
