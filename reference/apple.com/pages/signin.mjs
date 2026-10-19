// The store's sign-in, Apple Account's widget in its iframe (idmsa.apple.com):
// the title, a text field and its floating label at rest (the page focuses
// the field on load: blurred first) and focused, the checkbox (its parts,
// checked and focused on a copy the probe adds), the continue button, a link.
import { TYPE, BOX, MOTION, LINK, FOCUS } from './common.mjs';
const F = 'idmsa.apple.com';
const CHECKED = '<div class="form-checkbox"><input type="checkbox" id="cui-checked" class="form-checkbox-input" checked><label for="cui-checked" class="form-label"><span class="form-checkbox-indicator" aria-hidden="true"></span>Remember me</label></div>';
export default {
  url: 'https://secure.store.apple.com/shop/account/home',
  wait: 5000,
  probes: [
    { name: 'title', frame: F, select: '.si-container-title', props: TYPE },
    { name: 'field', frame: F, select: '.form-textbox', props: [...BOX, 'width', 'height', 'position'], box: true },
    { name: 'field.input', frame: F, blur: true, select: '#account_name_text_field', props: [...TYPE, ...BOX, 'height', ...FOCUS, ...MOTION], box: true, states: ['focus', 'focus-visible', 'hover'] },
    { name: 'field.label', frame: F, blur: true, select: '#apple_id_field_label', stateOn: '#account_name_text_field', props: [...TYPE, 'position', 'top', 'left', 'transform', 'transform-origin', ...MOTION], states: ['focus'] },
    { name: 'checkbox', frame: F, select: '.form-checkbox-indicator', props: [...BOX, 'width', 'height', 'position', 'top', 'left', 'display', ...MOTION], box: true },
    { name: 'checkbox.label', frame: F, select: '#remember-me-label', props: [...TYPE, 'padding-top', 'padding-left', 'position', 'display'] },
    { name: 'checkbox.box', frame: F, select: '.form-checkbox', props: ['position', 'margin-bottom'], box: true },
    { name: 'checkbox.input', frame: F, select: '#remember-me', props: ['position', 'top', 'left', 'width', 'height', 'opacity', 'margin-top', 'margin-left'] },
    // checked, and focused from the keyboard: apple.com's classes on a copy the
    // probe adds beside the widget's own (which is not checked until clicked)
    { name: 'checkbox.checked', frame: F, inject: CHECKED, into: ':has(> .form-checkbox)', select: '.form-checkbox-indicator', stateOn: '#cui-checked', props: [...BOX, 'color', ...FOCUS], states: ['focus-visible', 'hover'] },
    { name: 'checkbox.check', frame: F, inject: CHECKED, into: ':has(> .form-checkbox)', select: '.form-checkbox-indicator', pseudo: '::after', props: ['content', 'font-family', 'font-size', 'line-height', 'position', 'top', 'left', 'width', 'height', 'color'] },
    { name: 'checkbox.unchecked-focus', frame: F, select: '.form-checkbox-indicator', stateOn: '#remember-me', props: FOCUS, states: ['focus-visible'] },
    { name: 'continue', frame: F, blur: true, select: '#sign-in', props: [...TYPE, ...BOX, 'width', 'height', 'opacity', ...FOCUS], box: true, states: ['hover', 'focus-visible'] },
    { name: 'link', frame: F, select: '#iforgot-link', props: [...TYPE, ...LINK], states: ['hover'] },
  ],
};
