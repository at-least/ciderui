import { examples } from '../../../test/harness.mjs';

const ex = examples(import.meta.url);
const V = '.cui-console-viewer';

export const cases = [
  {
    name: 'viewer',
    html: ex('viewer'),
    click: '[command="show-modal"]',
    widths: [1200, 390],
    checks: [
      { target: V, props: ['display', 'flex-direction', 'row-gap', 'padding-top', 'border-top-left-radius', 'background-color'], golden: 'Quick Look\'s: the picture on the console\'s ground, 8px in, 20px corners, its caption 8px under it' },
      { target: V, props: ['box-width', 'box-height'], golden: 'the picture whole, fit to the window less 16px at each side, never past it' },
      { target: `${V}-image`, props: ['border-top-left-radius'], asc: ['media', 'thumb.img'] },
      { target: `${V}-image`, props: ['object-fit'], golden: 'contain, not App Store Connect\'s thumbnail\'s fill: the picture whole, at its own shape' },
      { target: `${V}-caption`, props: ['color', 'font-size', 'line-height', 'text-align'], golden: 'the secondary label at the body\'s size, centered under the picture' },
      { target: `${V}-close`, props: ['position', 'top', 'right'], golden: 'the close button over the picture\'s corner, 16px in' },
      { target: V, pseudo: '::backdrop', props: ['background-color'], golden: 'the dialog\'s scrim behind it (--cui-c-scrim)' },
    ],
  },
  {
    name: 'viewer no caption',
    html: ex('viewer').replace('Drawn 2 November', ''),
    click: '[command="show-modal"]',
    widths: [1200],
    checks: [
      { target: `${V}-caption`, props: ['display'], golden: 'no caption, no room for one' },
    ],
  },
];
