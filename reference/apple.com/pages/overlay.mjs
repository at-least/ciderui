// The store's overlay, opened (clicked once, at the first width: it stays
// open as the window narrows, so its probes are a spec of their own): iPhone
// 18 Pro's buy page, its first "Show more" (Which model is right for you?),
// a modal dialog over the page: the window-wide layer that dims the page and
// scrolls the dialog (rc-overlay), the dialog (role=dialog, aria-modal), its
// white panel and its content's room, its close button (a 44px target, a
// 36px circle, a 20px X), hovered and focused; the page under it held. The
// store steps it at 736/737px, where its own media queries are (the panel
// 816px at most above, 640px below), not at 735; and the panel is the
// window's width below 816 and 640px.
import { BOX, MOTION, FOCUS, BLUR } from './common.mjs';

const SPACE = ['padding-top', 'padding-right', 'padding-bottom', 'padding-left', 'margin-top', 'margin-right', 'margin-bottom', 'margin-left'];

export default {
  url: 'https://www.apple.com/shop/buy-iphone/iphone-18-pro',
  widths: [1600, 1200, 900, 816, 737, 736, 735, 734, 640, 390],
  wait: 2500,
  probes: [
    // the layer: fixed over the window, the page dimmed, fading in; it is
    // what scrolls when the dialog is taller than the window
    { name: 'overlay', click: 'a.rc-decisionsection-button', animations: '.rc-overlay', select: '.rc-overlay', props: ['position', 'top', 'left', 'z-index', 'display', 'align-items', 'justify-content', 'overflow-x', 'overflow-y', 'overscroll-behavior', 'background-color', 'opacity', ...BLUR, ...MOTION], box: true },
    // the page under it, held still
    { name: 'root', select: 'html', props: ['overflow-x', 'overflow-y'] },
    { name: 'dialog', select: '.rc-overlay [role="dialog"][aria-modal="true"]', props: ['position', 'width', 'max-width', 'margin-left', 'margin-right', 'outline-style'], box: true, position: true },
    { name: 'panel', select: '.rc-overlay-popup-outer', props: [...BOX, 'margin-right', 'margin-left', 'border-top-right-radius', 'border-bottom-left-radius', 'border-bottom-right-radius', 'overflow-x', 'overflow-y', 'display', 'flex-direction', 'color'], box: true, position: true },
    { name: 'content', select: '.rc-overlay-popup-content', props: [...SPACE], box: true, position: true },
    // the close button: its target, its circle, its X
    { name: 'close', select: '.rc-overlay-close', props: ['width', 'height', ...SPACE, 'display', 'align-items', 'justify-content', 'background-color', 'border-top-width', 'cursor', 'z-index', 'position'], box: true, position: true, states: ['focus-visible'] },
    { name: 'close.focus', select: '.rc-overlay-close', props: FOCUS, states: ['focus-visible', 'focus+focus-visible'] },
    { name: 'close.circle', select: '.rc-overlay-close .rc-overlay-closesvg', props: ['width', 'height', 'background-color', 'color', 'border-top-left-radius', ...MOTION], box: true, states: ['hover'] },
    { name: 'close.circle-hover', select: '.rc-overlay-close .rc-overlay-closesvg', stateOn: '.rc-overlay-close', props: ['background-color', 'color'], states: ['hover', 'active'] },
    { name: 'close.icon', select: '.rc-overlay-close svg', props: ['width', 'height', 'fill', 'top', 'right', 'position'], box: true },
  ],
};
