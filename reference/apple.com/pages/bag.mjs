// The empty bag: its header and message, its two buttons (the store's form
// buttons), the ribbon above it.
import { TYPE, BOX, MOTION, FOCUS, PAGE } from './common.mjs';
const STATES = ['hover', 'active', 'focus-visible'];
export default {
  url: 'https://www.apple.com/shop/bag',
  wait: 2500,
  probes: [
    { name: 'page.body', select: 'body', props: PAGE },
    { name: 'header', select: '.rs-bag-header', props: [...TYPE, 'margin-bottom'] },
    { name: 'message', select: '.rs-bagempty-message', props: TYPE },
    { name: 'form-button', select: '.rs-bagempty .form-button:not(.form-button-secondary)', props: [...TYPE, ...BOX, 'width', 'display', ...FOCUS, ...MOTION], box: true, states: STATES },
    { name: 'form-button.secondary', select: '.rs-bagempty .form-button-secondary', props: [...TYPE, ...BOX, 'width', 'display', ...FOCUS, ...MOTION], box: true, states: STATES },
    { name: 'accordion.button', select: '.rc-accordion-button', props: [...TYPE, ...BOX, 'width'], box: true, states: ['hover'] },
    { name: 'accordion.item', select: '.rc-accordion-item', props: BOX, box: true },
    { name: 'ribbon', select: '.ribbon-content-wrapper', props: [...TYPE, ...BOX], box: true },
  ],
};
