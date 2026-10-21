// The iPhone feedback form, for apple.com's multi-line field (its
// form-textarea: a bordered box, its textarea inside under its label), at
// rest; its text field's label beside it, for the page's colors. Read only:
// nothing typed, nothing sent. Its states are feedback-states.mjs.
import { TYPE, BOX, MOTION } from './common.mjs';

export const BOXES = '.form-textbox:has(> .form-textarea)';
export const WRAP = `${BOXES} > .form-textarea`;
export const AREA = `${WRAP} textarea`;
export const LABEL = `${BOXES} > .form-textbox-label`;
export const EDGE = ['border-top-width', 'border-top-style', 'border-top-color', 'border-bottom-width', 'border-top-left-radius'];
export const SPACE = ['padding-top', 'padding-right', 'padding-bottom', 'padding-left', 'margin-top', 'margin-bottom'];

export default {
  url: 'https://www.apple.com/feedback/iphone/',
  wait: 2500,
  probes: [
    { name: 'textarea.box', select: BOXES, props: ['position', 'height', 'margin-bottom'], box: true },
    { name: 'textarea', select: WRAP, props: [...TYPE, ...SPACE, ...EDGE, 'display', 'height', 'box-sizing', 'background-color', 'outline-style', ...MOTION], box: true, position: true },
    { name: 'textarea.input', select: AREA, props: [...TYPE, ...SPACE, ...EDGE, 'height', 'min-height', 'width', 'box-sizing', 'background-color', 'resize', 'outline-style', 'overflow-y'], box: true, position: true },
    { name: 'textarea.label', select: LABEL, props: [...TYPE, 'position', 'top', 'left', 'pointer-events', 'white-space', ...MOTION], box: true, position: true },
    { name: 'field.label', select: '.form-textbox:has(> .form-textbox-input) > .form-textbox-label', props: [...TYPE] },
    { name: 'field.input', select: '.form-textbox > .form-textbox-input', props: [...TYPE, 'border-top-color', 'background-color'] },
  ],
};
