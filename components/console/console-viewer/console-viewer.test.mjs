import { examples } from '../../../test/harness.mjs';

const ex = examples(import.meta.url);
const V = '.cui-console-viewer';
// @invariant: the example renders the viewer empty; a page's script sets the picture and the caption before
// it opens it, as this does (a 1600x900 gray picture)
const PICTURE = 'data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHdpZHRoPSIxNjAwIiBoZWlnaHQ9IjkwMCI+PHJlY3Qgd2lkdGg9IjE2MDAiIGhlaWdodD0iOTAwIiBmaWxsPSIjOGU4ZTkzIi8+PC9zdmc+';
const filled = (caption) => ex('viewer')
  .replace('<img class="cui-console-viewer-image" alt="">', `<img class="cui-console-viewer-image" src="${PICTURE}" alt="The cover">`)
  .replace('<p class="cui-console-viewer-caption"></p>', `<p class="cui-console-viewer-caption">${caption}</p>`);

export const cases = [
  {
    name: 'viewer',
    html: filled('Drawn 2 November'),
    click: '[command="show-modal"]',
    widths: [1200, 390],
    checks: [
      { target: V, props: ['display', 'flex-direction', 'row-gap', 'padding-top', 'border-top-left-radius', 'background-color'], golden: 'Quick Look\'s: the picture on the console\'s ground, 8px in, 20px corners, its caption 8px under it' },
      { target: V, props: ['box-width', 'box-height'], golden: 'the picture whole, fit to the window less 16px at each side, never past it' },
      { target: `${V}-image`, props: ['border-top-left-radius'], asc: ['media', 'thumb.img'] },
      { target: `${V}-image`, props: ['object-fit'], golden: 'contain, not App Store Connect\'s thumbnail\'s fill: the picture whole, at its own shape' },
      { target: `${V}-caption`, props: ['color', 'font-size', 'line-height', 'text-align'], golden: 'the secondary label at the body\'s size, centered under the picture' },
      { target: `${V}-close`, props: ['position', 'align-self', 'box-top'], golden: 'the close button in its own row at the trailing end above the picture, as Quick Look\'s title bar: over the picture\'s corner a white picture hid it' },
      { target: `${V}-image`, props: ['box-top'], golden: 'the picture under the close button\'s row, 8px below it' },
      { target: V, pseudo: '::backdrop', props: ['background-color'], golden: 'the dialog\'s scrim behind it (--cui-c-scrim)' },
    ],
  },
  {
    name: 'viewer no caption',
    html: filled(''),
    click: '[command="show-modal"]',
    widths: [1200],
    checks: [
      { target: `${V}-caption`, props: ['display'], golden: 'no caption, no room for one' },
    ],
  },
];
