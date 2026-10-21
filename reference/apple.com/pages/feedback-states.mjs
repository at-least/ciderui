// The iPhone feedback form's multi-line field focused (its textarea clicked
// once, at the first width: the class apple.com's script sets on its box),
// and left empty (the focus taken off: its script marks it in error, with
// its message). Read only: nothing typed, nothing sent.
import { TYPE, BOX } from './common.mjs';
import { BOXES, WRAP, AREA, LABEL, EDGE, SPACE } from './feedback.mjs';

export default {
  url: 'https://www.apple.com/feedback/iphone/',
  widths: [1600, 390],
  wait: 2500,
  probes: [
    { name: 'focused', click: AREA, select: WRAP, props: [...SPACE, ...EDGE, 'background-color', 'outline-style', 'height'], box: true, widths: [1600] },
    { name: 'focused.input', select: AREA, props: [...SPACE, 'height'], box: true, position: true, widths: [1600] },
    { name: 'focused.label', select: LABEL, props: [...TYPE, 'top', 'left'], box: true, position: true, widths: [1600] },
    { name: 'error', blur: true, select: WRAP, props: [...EDGE, 'background-color', 'margin-bottom', 'height'], box: true },
    { name: 'error.label', select: LABEL, props: [...TYPE, 'top', 'left'] },
    { name: 'error.wrapper', select: `${BOXES} > .form-message-wrapper`, props: [...BOX, 'display'], box: true, position: true },
    { name: 'error.message', select: `${BOXES} .form-message`, props: [...TYPE] },
  ],
};
