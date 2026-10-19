// The store's sign-in, Apple Account's widget in its iframe (idmsa.apple.com):
// the title, a text field and its floating label at rest (the page focuses
// the field on load: blurred first) and focused, the checkbox, the continue
// button, a link.
import { TYPE, BOX, MOTION, LINK, FOCUS } from './common.mjs';
const F = 'idmsa.apple.com';
export default {
  url: 'https://secure.store.apple.com/shop/account/home',
  wait: 5000,
  probes: [
    { name: 'title', frame: F, select: '.si-container-title', props: TYPE },
    { name: 'field', frame: F, select: '.form-textbox', props: [...BOX, 'width', 'height', 'position'], box: true },
    { name: 'field.input', frame: F, blur: true, select: '#account_name_text_field', props: [...TYPE, ...BOX, 'height', ...FOCUS, ...MOTION], box: true, states: ['focus', 'focus-visible', 'hover'] },
    { name: 'field.label', frame: F, blur: true, select: '#apple_id_field_label', stateOn: '#account_name_text_field', props: [...TYPE, 'position', 'top', 'left', 'transform', 'transform-origin', ...MOTION], states: ['focus'] },
    { name: 'checkbox', frame: F, select: '.form-checkbox-indicator', props: [...BOX, 'width', 'height', 'position'], box: true },
    { name: 'checkbox.label', frame: F, select: '#remember-me-label', props: TYPE },
    { name: 'continue', frame: F, blur: true, select: '#sign-in', props: [...TYPE, ...BOX, 'width', 'height', 'opacity', ...FOCUS], box: true, states: ['hover', 'focus-visible'] },
    { name: 'link', frame: F, select: '#iforgot-link', props: [...TYPE, ...LINK], states: ['hover'] },
  ],
};
