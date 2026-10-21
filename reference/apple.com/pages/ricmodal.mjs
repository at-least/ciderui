// The marketing pages' modal (ric-modal), the iPhone page's "Read more about
// Apple Trade In", opened (its button clicked once, at the first width; it
// stays open as the window narrows): the layer over the window and its
// scrim (its ::after: a dimming and a blur), the page held, the overlay
// centered in it, its corners (its content clipped round), its content's
// room, its close button (a circle in a 44px target) at rest, hovered,
// pressed and focused; every animation the click starts. It steps
// at the marketing pages' widths, 1069 and 735px. Read only.
import { MOTION, BLUR } from './common.mjs';

const M = '#trade-in-modal';
const SPACE = ['padding-top', 'padding-right', 'padding-bottom', 'padding-left', 'margin-top', 'margin-right', 'margin-bottom', 'margin-left'];

export default {
  url: 'https://www.apple.com/iphone/',
  widths: [1600, 1200, 1069, 1068, 900, 735, 734, 640, 390],
  wait: 2500,
  probes: [
    { name: 'modal', click: `button[data-modal-open="trade-in-modal"]`, animations: M, select: M, props: ['position', 'top', 'left', 'right', 'bottom', 'overflow-y', 'z-index', 'opacity', 'display', 'background-color'], box: true },
    { name: 'scrim', select: M, pseudo: '::after', props: ['content', 'position', 'top', 'left', 'right', 'bottom', 'background-color', 'background-image', ...BLUR, 'opacity', 'z-index', ...MOTION] },
    { name: 'root', select: 'html', props: ['overflow-y'] },
    { name: 'container', select: `${M} .ric-modal-overlay-container`, props: [...SPACE, 'display', 'align-items', 'justify-content', 'min-height', 'box-sizing'], box: true },
    { name: 'overlay', select: `${M} .ric-modal-overlay`, props: [...SPACE, 'width', 'max-width', 'border-top-left-radius', 'border-bottom-right-radius', 'overflow-x', 'overflow-y', 'position', 'z-index', 'opacity', 'transform', 'background-color', ...MOTION], box: true, position: true },
    { name: 'content', select: `${M} .ric-modal-content-wrapper`, props: [...SPACE, 'background-color', 'color', 'border-top-left-radius'], box: true, position: true },
    { name: 'viewport', select: `${M} .ric-modal-viewport-content`, props: [...SPACE], box: true, position: true },
    { name: 'close.wrapper', select: `${M} .ric-modal-close-button-wrapper`, props: ['position', 'top', 'right', 'left', 'width', 'height', 'margin-left', 'margin-right', 'z-index', 'display', 'justify-content', 'align-items', 'order'], box: true, position: true },
    { name: 'close', select: `${M} .ric-modal-close-button`, props: ['width', 'height', ...SPACE, 'display', 'align-items', 'justify-content', 'background-color', 'color', 'cursor', 'border-top-width', ...MOTION], box: true, position: true },
    // its circle (an element between the button and its svg) and the svg
    { name: 'close.icon', select: `${M} .ric-modal-close-icon`, props: ['width', 'height', 'display', 'align-items', 'position', 'background-color', 'color', 'border-top-left-radius', ...MOTION], box: true, position: true },
    { name: 'close.svg', select: `${M} .ric-modal-close-icon svg`, props: ['width', 'height', 'fill', 'position'], box: true, position: true },
    // hovered, pressed and focused last: the still style that holds a state
    // cancels the page's animations, and taking it off starts them again
    { name: 'close.icon-states', select: `${M} .ric-modal-close-icon`, stateOn: `${M} .ric-modal-close-button`, props: ['background-color', 'color', 'outline-style', 'outline-color', 'outline-width', 'outline-offset'], states: ['hover', 'active', 'focus-visible'] },
  ],
};
