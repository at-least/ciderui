// A dark product page modal (ric-modal theme-dark), MacBook Pro's "Explore
// the M5 family of chips", opened (its button clicked once, at the first
// width): its panel's ground and text, a heading's color, its close button's
// circle at rest, hovered, pressed and focused. Its panel is the larger kind
// (l2-modal, 1260px), so its colors alone are read. Read only.
const M = '#performance-modal';

export default {
  url: 'https://www.apple.com/macbook-pro/',
  widths: [1600, 390],
  wait: 2500,
  probes: [
    { name: 'modal', click: `[data-modal-open="performance-modal"]`, select: M, props: ['position', 'color-scheme'] },
    { name: 'content', select: `${M} .ric-modal-content-wrapper`, props: ['background-color', 'color'] },
    { name: 'heading', select: `${M} .ric-modal-content-wrapper h2`, props: ['color'] },
    { name: 'close.icon', select: `${M} .ric-modal-close-icon`, props: ['background-color', 'color', 'border-top-left-radius'] },
    { name: 'close.icon-states', select: `${M} .ric-modal-close-icon`, stateOn: `${M} .ric-modal-close-button`, props: ['background-color', 'color', 'outline-style', 'outline-color', 'outline-width', 'outline-offset'], states: ['hover', 'active', 'focus-visible'] },
  ],
};
