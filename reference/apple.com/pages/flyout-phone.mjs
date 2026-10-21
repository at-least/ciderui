// The global nav's flyouts on a phone (to 833px): the menu opened by its
// button, its Mac row hovered (a pointer at a phone's width: its chevron
// slides in), then clicked: the row's flyout over the menu, a back button at
// the bar's start and the menu's X at its end, the menu's rows sliding out
// and the flyout's in (`animations`), its groups one under another, their
// headings, rows and links; then its back button clicked: the menu's rows
// back. At 390px; the page read with its locale banner above the nav (135px),
// which the opened menu covers.
import { TYPE, MOTION, FOCUS, ANIMATION } from './common.mjs';

const FLY = '#globalnav-submenu-link-mac';
const ELEVATED = `${FLY} .globalnav-submenu-group-elevated`;
const GROUP = `${FLY} .globalnav-submenu-group:nth-child(2)`;
const ROW = '#globalnav .globalnav-link-mac';
const BACK = '#globalnav .globalnav-menuback';
const SPACE = ['padding-top', 'padding-right', 'padding-bottom', 'padding-left', 'margin-top', 'margin-right', 'margin-bottom', 'margin-left'];
const SHOWN = ['opacity', 'visibility', 'transform'];

export default {
  url: 'https://www.apple.com/',
  widths: [390],
  wait: 2500,
  probes: [
    { name: 'menu', click: '#globalnav .globalnav-menutrigger-button', select: '#globalnav .globalnav-content', props: ['background-color', 'height'], box: true },
    // a row hovered, and one not: its chevron at its end
    { name: 'row.chevron', hover: ROW, hovered: 900, animations: '#globalnav .globalnav-item-mac', select: `${ROW} .globalnav-link-chevron`, props: ['display', 'width', 'height', ...SPACE, 'color', ...SHOWN, ...ANIMATION, ...MOTION], keyframes: true, box: true, position: true },
    { name: 'row.chevron-image', select: `${ROW} .globalnav-link-chevron svg`, props: ['width', 'height', 'transform'], box: true, position: true },
    { name: 'row.text', select: `${ROW} .globalnav-link-text`, props: ['color', ...SHOWN], box: true, position: true },
    { name: 'row.other-chevron', select: '#globalnav .globalnav-link-ipad .globalnav-link-chevron', props: [...SHOWN, ...ANIMATION], keyframes: true },
    // the row clicked: its flyout over the menu
    { name: 'flyout', click: ROW, animations: '#globalnav', select: FLY, props: ['position', 'top', 'left', 'height', 'background-color', 'overflow-y', 'z-index', ...SHOWN, ...ANIMATION], keyframes: true, box: true, position: true },
    { name: 'flyout.scroll', select: `${FLY} .globalnav-flyout-scroll-container`, props: ['height', 'overflow-y', ...SPACE], box: true, position: true },
    { name: 'flyout.content', select: `${FLY} .globalnav-flyout-content`, props: [...SPACE, 'display', 'flex-direction'], box: true, position: true },
    { name: 'content', select: '#globalnav .globalnav-content', props: ['background-color', 'overflow-y', 'height'], box: true },
    { name: 'group.elevated', select: ELEVATED, props: [...SPACE, 'max-width'], box: true, position: true },
    { name: 'group', select: GROUP, props: [...SPACE, 'max-width'], box: true, position: true },
    { name: 'header.elevated', select: `${ELEVATED} .globalnav-submenu-header`, props: ['display'] },
    { name: 'header', select: `${GROUP} .globalnav-submenu-header`, props: [...TYPE, ...SPACE, ...SHOWN], box: true, position: true },
    { name: 'item.elevated', select: `${ELEVATED} .globalnav-submenu-list-item-elevated`, props: [...TYPE, ...SPACE, ...SHOWN, 'transition-delay'], box: true, position: true },
    { name: 'item.elevated2', select: `${ELEVATED} .globalnav-submenu-list-item-elevated:nth-child(2)`, props: SPACE, box: true, position: true },
    { name: 'item.after', select: `${ELEVATED} .globalnav-submenu-list-item`, props: [...TYPE, ...SPACE], box: true, position: true },
    { name: 'item', select: `${GROUP} .globalnav-submenu-list-item`, props: [...TYPE, ...SPACE], box: true, position: true },
    { name: 'item2', select: `${GROUP} .globalnav-submenu-list-item:nth-child(2)`, props: SPACE, box: true, position: true },
    { name: 'link.elevated', select: `${ELEVATED} .globalnav-submenu-list-item-elevated .globalnav-submenu-link`, props: [...TYPE, ...SPACE, 'display', 'max-width', ...FOCUS, ...MOTION], box: true, position: true, states: ['hover', 'focus-visible'] },
    { name: 'link', select: `${GROUP} .globalnav-submenu-link`, props: [...TYPE, ...SPACE, 'display', 'max-width', ...MOTION], box: true, position: true, states: ['hover'] },
    // the bar over it: the back button at its start, the menu's X at its end
    { name: 'back', select: BACK, props: ['position', 'top', 'left', 'right', 'width', 'height', 'z-index', ...SHOWN, 'transition-delay', ...MOTION], box: true, position: true },
    { name: 'back.button', select: `${BACK}-button`, props: ['width', 'height', ...SPACE, 'color', 'opacity', 'background-color', 'cursor', ...FOCUS, ...MOTION], box: true, position: true, states: ['hover', 'focus-visible'] },
    { name: 'back.image', select: `${BACK}-button svg`, props: ['width', 'height', 'transform'], box: true, position: true },
    { name: 'close', select: '#globalnav .globalnav-menutrigger-button', props: ['color', 'opacity'], box: true, position: true },
    { name: 'row.away', select: `${ROW} .globalnav-link-text`, props: SHOWN, box: true, position: true },
    // its back button clicked: the menu's rows back
    { name: 'back.clicked', click: `${BACK}-button`, animations: '#globalnav', select: FLY, props: [...SHOWN, ...ANIMATION], box: true, position: true },
    { name: 'back.row', select: `${ROW} .globalnav-link-text`, props: SHOWN, box: true, position: true },
    { name: 'back.gone', select: BACK, props: SHOWN },
    // the slides' keyframes, as the page's stylesheet defines them (an element
    // the probe adds, naming them)
    { name: 'keyframes', inject: '<div style="animation-name: globalnav-flyout-slide-forward-previous, globalnav-flyout-slide-forward-next, globalnav-flyout-slide-back-previous, globalnav-flyout-slide-back-next"></div>', into: '#globalnav', props: ['animation-name'], keyframes: true },
  ],
};
