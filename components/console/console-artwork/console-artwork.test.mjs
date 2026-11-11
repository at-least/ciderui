import { examples } from '../../../test/harness.mjs';

const ex = examples(import.meta.url);
const A = '.cui-console-artwork';

export const cases = [
  {
    name: 'artwork',
    html: ex('artwork'),
    widths: [1200, 390],
    checks: [
      { target: `img${A}:not([data-aspect])`, props: ['border-top-left-radius'], asc: ['media', 'thumb.img'] },
      { target: `img${A}:not([data-aspect])`, props: ['box-width', 'box-height', 'object-fit', 'background-color'], golden: 'as wide as its column, 16:9, the picture cropped to fill it on the console card\'s gray (App Store Connect\'s frame rgb(242, 242, 242), media.json)' },
      { target: `${A}[data-aspect="square"]`, props: ['box-width', 'box-height'], golden: 'square, as wide as its column' },
      { target: `span${A}`, props: ['display', 'align-items', 'justify-content', 'color', 'font-size', 'line-height', 'text-align', 'box-height'], golden: 'no picture: its words centered in its place, the secondary label at the console\'s xs size, the same 16:9 box' },
    ],
  },
];
