// The retail page, for apple.com's own text field (its form-textbox, the
// store's forms' field): the page's store search (with its icons) and a
// plain field the probe adds with apple.com's classes, at rest, hovered,
// focused; its label floated (the class apple.com's script sets on it as the
// field is focused or filled, on a copy the probe adds) and the field
// filled (form-textbox-entered).
import { TYPE, BOX, MOTION, FOCUS } from './common.mjs';
const FIELD = (id, labelClass = '', inputClass = '') => `<div class="form-textbox"><input type="text" class="form-textbox-input${inputClass}" id="${id}" required value=""><label class="form-textbox-label${labelClass}" for="${id}">Email</label></div>`;
export const probes = [
  { name: 'field.search', select: '#store-search-input', props: ['padding-top', 'padding-right', 'padding-bottom', 'padding-left', 'height'], box: true },
  { name: 'field', inject: FIELD('cui-field'), into: 'main', props: ['position', 'height'], box: true },
  { name: 'field.input', inject: FIELD('cui-field'), into: 'main', select: 'input', props: [...TYPE, ...BOX, 'height', 'box-sizing', ...FOCUS, ...MOTION], box: true, states: ['hover', 'focus', 'focus-visible', 'focus+focus-visible'] },
  { name: 'field.label', inject: FIELD('cui-field'), into: 'main', select: 'label', props: [...TYPE, 'position', 'top', 'left', 'transform', 'pointer-events', 'white-space', ...MOTION] },
  { name: 'field.label-floated', inject: FIELD('cui-field', ' form-textbox-eyebrow-label', ' form-textbox-entered'), into: 'main', select: 'label', props: [...TYPE, 'position', 'top', 'left'] },
  { name: 'field.input-filled', inject: FIELD('cui-field', ' form-textbox-eyebrow-label', ' form-textbox-entered'), into: 'main', select: 'input', props: [...TYPE, ...BOX] },
];
export default {
  url: 'https://www.apple.com/retail/',
  probes,
};
