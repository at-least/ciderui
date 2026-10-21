// The retail page, for apple.com's own text field (its form-textbox, the
// store's forms' field): the page's store search (with its icons) and a
// plain field the probe adds with apple.com's classes, at rest, hovered,
// focused; its label floated (the class apple.com's script sets on it as the
// field is focused or filled, on a copy the probe adds) and the field
// filled (form-textbox-entered).
import { TYPE, BOX, MOTION, FOCUS } from './common.mjs';
const FIELD = (id, labelClass = '', inputClass = '') => `<div class="form-textbox"><input type="text" class="form-textbox-input${inputClass}" id="${id}" required value=""><label class="form-textbox-label${labelClass}" for="${id}">Email</label></div>`;
// a field in error, as the gift card's page marks an empty required one the
// focus has left (giftcard:error*): its classes, its message and its icon
const ERROR = `<div class="form-textbox is-error"><input type="text" class="form-textbox-input" id="cui-error" aria-invalid="true" value=""><label class="form-textbox-label" for="cui-error">Email</label><div><div class="is-error"><div class="form-message-wrapper" id="cui-error-message"><span class="form-message">Please enter a valid email.</span></div></div></div></div>`;
const DISABLED = `<div class="form-textbox"><input type="text" class="form-textbox-input" id="cui-disabled" disabled value=""><label class="form-textbox-label" for="cui-disabled">Email</label></div>`;
// the multi-line field, the feedback form's markup (feedback:textarea*: a
// bordered box, its textarea in it, its label after it), with the store's
// stylesheet here: at rest, focused and in error (the classes apple.com's
// script sets: the box's focused, the label's eyebrow, the field's is-error)
const AREA = (cls = '', box = '', label = '') => `<div class="form-textbox${cls}"><div class="form-textarea${box}"><textarea id="cui-area" aria-labelledby="cui-area-label"></textarea></div><span class="form-textbox-label${label}" id="cui-area-label">Comments</span></div>`;
// the store's dropdown (tradein:dropdown*), in error and disabled: the
// trade-in's markup, nothing chosen, with the class apple.com's script sets
// (is-error, as the feedback form's empty required dropdowns get it)
const DROPDOWN = (cls = '', attrs = '') => `<div class="form-dropdown${cls}"><label class="form-dropdown-label form-dropdown-label-placeholder" for="cui-dd">Choose your model</label><select id="cui-dd" class="form-dropdown-select form-dropdown-selectnone"${attrs}><option value="">Select</option></select><span class="form-dropdown-chevron" aria-hidden="true"></span></div>`;
export const probes = [
  { name: 'field.search', select: '#store-search-input', props: ['padding-top', 'padding-right', 'padding-bottom', 'padding-left', 'height'], box: true },
  { name: 'field', inject: FIELD('cui-field'), into: 'main', props: ['position', 'height'], box: true },
  { name: 'field.input', inject: FIELD('cui-field'), into: 'main', select: 'input', props: [...TYPE, ...BOX, 'height', 'box-sizing', ...FOCUS, ...MOTION], box: true, states: ['hover', 'focus', 'focus-visible', 'focus+focus-visible'] },
  { name: 'field.label', inject: FIELD('cui-field'), into: 'main', select: 'label', props: [...TYPE, 'position', 'top', 'left', 'transform', 'pointer-events', 'white-space', ...MOTION] },
  { name: 'field.label-floated', inject: FIELD('cui-field', ' form-textbox-eyebrow-label', ' form-textbox-entered'), into: 'main', select: 'label', props: [...TYPE, 'position', 'top', 'left'] },
  { name: 'field.error', inject: ERROR, into: 'main', select: 'input', props: ['color', 'background-color', 'border-top-color', 'border-top-width'] },
  { name: 'field.error-label', inject: ERROR, into: 'main', select: 'label', props: ['color'] },
  { name: 'field.error-message', inject: ERROR, into: 'main', select: '.form-message', props: [...TYPE] },
  { name: 'field.error-icon', inject: ERROR, into: 'main', select: '.form-message-wrapper', pseudo: '::before', props: ['content', 'font-family', 'font-size', 'line-height', 'color', 'padding-right', 'display'] },
  { name: 'field.disabled', inject: DISABLED, into: 'main', select: 'input', props: ['color', 'background-color', 'border-top-color', 'opacity', 'cursor'] },
  { name: 'field.disabled-label', inject: DISABLED, into: 'main', select: 'label', props: ['color', 'opacity'] },
  { name: 'textarea', inject: AREA(), into: 'main', select: '.form-textarea', props: [...TYPE, ...BOX, 'height', 'box-sizing', 'display', 'border-top-width', 'border-top-style', 'border-top-color', 'border-top-left-radius', 'outline-style', ...MOTION], box: true },
  { name: 'textarea.input', inject: AREA(), into: 'main', select: 'textarea', props: [...TYPE, ...BOX, 'height', 'min-height', 'box-sizing', 'resize', 'outline-style', 'border-top-width', 'border-top-style'], box: true },
  { name: 'textarea.label', inject: AREA(), into: 'main', select: '.form-textbox-label', props: [...TYPE, 'position', 'top', 'left', 'pointer-events', 'white-space', ...MOTION] },
  { name: 'textarea.focused', inject: AREA('', ' focused', ' form-textbox-eyebrow-label'), into: 'main', select: '.form-textarea', props: [...BOX, 'height', 'border-top-width', 'border-top-color', 'outline-style'], box: true },
  { name: 'textarea.label-floated', inject: AREA('', ' focused', ' form-textbox-eyebrow-label'), into: 'main', select: '.form-textbox-label', props: [...TYPE, 'top', 'left'] },
  { name: 'textarea.error', inject: AREA(' is-error'), into: 'main', select: '.form-textarea', props: ['color', 'background-color', 'border-top-color', 'border-top-width'] },
  { name: 'textarea.error-label', inject: AREA(' is-error'), into: 'main', select: '.form-textbox-label', props: ['color'] },
  { name: 'textarea.disabled', inject: AREA('', ' disabled'), into: 'main', select: '.form-textarea', props: ['color', 'background-color', 'border-top-color'] },
  { name: 'textarea.disabled-input', inject: AREA('', ' disabled'), into: 'main', select: 'textarea', props: ['color', '-webkit-text-fill-color'] },
  { name: 'dropdown.error', inject: DROPDOWN(' is-error'), into: 'main', select: 'select', props: ['background-color', 'border-top-color', 'border-top-width', 'cursor'], states: ['focus'] },
  { name: 'dropdown.error-label', inject: DROPDOWN(' is-error'), into: 'main', select: 'label', props: ['color'] },
  { name: 'dropdown.disabled', inject: DROPDOWN('', ' disabled'), into: 'main', select: 'select', props: ['color', 'background-color', 'border-top-color', 'cursor', 'opacity'] },
  { name: 'dropdown.disabled-label', inject: DROPDOWN('', ' disabled'), into: 'main', select: 'label', props: ['color'] },
  { name: 'dropdown.disabled-chevron', inject: DROPDOWN('', ' disabled'), into: 'main', select: '.form-dropdown-chevron', props: ['opacity'] },
  { name: 'field.input-filled', inject: FIELD('cui-field', ' form-textbox-eyebrow-label', ' form-textbox-entered'), into: 'main', select: 'input', props: [...TYPE, ...BOX] },
];
export default {
  url: 'https://www.apple.com/retail/',
  probes,
};
