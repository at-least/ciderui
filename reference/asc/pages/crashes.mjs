// An app's TestFlight crashes, with none: the table's header (a cell, its
// label) over its one row, the empty table's message, centered in its room.
import { TYPE } from '../../apple.com/pages/common.mjs';

const CELL = ['padding-top', 'padding-right', 'padding-bottom', 'padding-left', 'vertical-align', 'text-align', 'background-color', 'border-top-width', 'border-top-style', 'border-top-color', 'border-bottom-width', 'border-bottom-style', 'border-bottom-color'];
export default {
  path: '/teams/<team>/apps/<app>/testflight/crashes',
  probes: [
    { name: 'head.cell', select: 'thead th', props: [...CELL, 'height'], box: true },
    { name: 'head.label', select: 'thead th span[class*="Text__BaseText"]', props: [...TYPE, 'text-transform'], box: true },
    { name: 'row', select: 'tbody tr', props: ['height', 'background-color', 'border-top-width', 'border-top-style', 'border-top-color', 'border-bottom-width', 'border-bottom-style', 'border-bottom-color'], box: true },
    { name: 'empty.cell', select: 'tbody td', props: [...CELL, 'height'], box: true },
    { name: 'empty', select: 'tbody td > div', props: ['padding-top', 'padding-right', 'padding-bottom', 'padding-left', 'text-align', 'background-color'], box: true },
    { name: 'empty.text', select: 'tbody td p', props: [...TYPE, 'text-align'], box: true },
  ],
};
