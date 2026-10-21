// The global nav's flyouts, a desktop's menus under its links: Mac's,
// opened by the pointer (`hover`: its link reached in steps, apple.com's
// script opening it 120ms on) at each width, then the pointer into the
// flyout's room at the window's right edge, where it hovers nothing: the bar
// while it is open, its links, the flyout, its content, its groups (the
// first elevated, its links large), their headings, lists, rows and links,
// the curtain over the page, the button beside the link that the keyboard
// opens it by and its chevron; then the pointer along the bar to iPad's link
// (the change from one flyout to another) and off the nav to the page (the
// flyout closing), each recording what it starts. From 834px; to 833px the
// links are the menu's (flyout-phone). The page is read with its locale
// banner above the nav (70px), so a box's top is the nav's top plus its own.
import { TYPE, BOX, MOTION, BLUR, FOCUS, FLEX } from './common.mjs';

const FLY = '#globalnav-submenu-link-mac';
const ELEVATED = `${FLY} .globalnav-submenu-group-elevated`;
const GROUP = `${FLY} .globalnav-submenu-group:nth-child(2)`;
const LAST = `${FLY} .globalnav-submenu-group:nth-child(3)`;
const TRIGGER = '#globalnav-menubutton-link-mac';
const SPACE = ['padding-top', 'padding-right', 'padding-bottom', 'padding-left', 'margin-top', 'margin-right', 'margin-bottom', 'margin-left'];
const ITEM = ['opacity', 'transform', 'transition-delay', ...MOTION];
const SIZE = ['flex-grow', 'flex-shrink', 'flex-basis', 'max-width', 'width'];
const MOTIONS = '#globalnav, #globalnav-curtain';
// the flyout's room at the window's right edge, below its links' first rows
const ROOM = (w) => [w - 10, 400];

export default {
  url: 'https://www.apple.com/',
  widths: [1600, 1200, 900, 834],
  wait: 2500,
  probes: [
    { name: 'nav', hover: '#globalnav .globalnav-link-mac', pointer: ROOM, animations: MOTIONS, select: '#globalnav', props: ['height', 'position', 'background-color', ...BLUR, 'z-index'], vars: ['--globalnav-background', '--globalnav-backdrop-filter'], box: true, position: true },
    { name: 'nav.content', select: '#globalnav .globalnav-content', props: ['position', 'z-index', 'background-color', ...BLUR], box: true, position: true },
    // the open flyout's link, another, the search beside them
    { name: 'link.open', select: '#globalnav .globalnav-link-mac', props: ['color', 'opacity', ...MOTION] },
    { name: 'link.other', select: '#globalnav .globalnav-link-ipad', props: ['color', 'opacity'] },
    { name: 'link.search', select: '#globalnav .globalnav-link-search', props: ['color', 'opacity'] },
    { name: 'flyout', select: FLY, props: ['position', 'top', 'left', 'right', 'height', 'max-height', ...SPACE, 'background-color', ...BLUR, 'color', 'overflow-y', 'z-index', 'visibility', ...MOTION], vars: ['--r-globalnav-flyout-height', '--r-globalnav-flyout-rate'], box: true, position: true },
    { name: 'flyout.scroll', select: `${FLY} .globalnav-flyout-scroll-container`, props: ['height', 'max-height', 'overflow-y', ...SPACE], box: true, position: true },
    { name: 'flyout.content', select: `${FLY} .globalnav-flyout-content`, props: ['max-width', ...SPACE, ...FLEX, 'box-sizing'], box: true, position: true },
    { name: 'group.elevated', select: ELEVATED, props: [...SIZE, ...SPACE, 'display'], box: true, position: true },
    { name: 'group', select: GROUP, props: [...SIZE, ...SPACE, 'display'], box: true, position: true },
    { name: 'group.last', select: LAST, props: [...SIZE, ...SPACE], box: true, position: true },
    { name: 'header.elevated', select: `${ELEVATED} .globalnav-submenu-header`, props: [...TYPE, ...SPACE, 'display', ...ITEM], box: true, position: true },
    { name: 'header', select: `${GROUP} .globalnav-submenu-header`, props: [...TYPE, ...SPACE, ...ITEM], box: true, position: true },
    { name: 'header.last', select: `${LAST} .globalnav-submenu-header`, props: ITEM },
    { name: 'list', select: `${GROUP} .globalnav-submenu-list`, props: [...SPACE, 'list-style-type', 'display'], box: true, position: true },
    { name: 'item.elevated', select: `${ELEVATED} .globalnav-submenu-list-item-elevated`, props: [...TYPE, ...SPACE, 'display', ...ITEM], box: true, position: true },
    { name: 'item.elevated2', select: `${ELEVATED} .globalnav-submenu-list-item-elevated:nth-child(2)`, props: [...SPACE, ...ITEM], box: true, position: true },
    { name: 'item.after', select: `${ELEVATED} .globalnav-submenu-list-item`, props: [...TYPE, ...SPACE, ...ITEM], box: true, position: true },
    { name: 'item.after2', select: `${ELEVATED} .globalnav-submenu-list-item ~ .globalnav-submenu-list-item`, props: [...SPACE, ...ITEM], box: true, position: true },
    { name: 'item', select: `${GROUP} .globalnav-submenu-list-item`, props: [...TYPE, ...SPACE, ...ITEM], box: true, position: true },
    { name: 'item2', select: `${GROUP} .globalnav-submenu-list-item:nth-child(2)`, props: [...SPACE, ...ITEM], box: true, position: true },
    { name: 'item.last', select: `${LAST} .globalnav-submenu-list-item`, props: ITEM },
    { name: 'link.elevated', select: `${ELEVATED} .globalnav-submenu-list-item-elevated .globalnav-submenu-link`, props: [...TYPE, ...SPACE, 'display', 'text-decoration-line', 'background-color', 'border-top-left-radius', ...FOCUS, ...MOTION], box: true, position: true, states: ['hover', 'focus-visible'] },
    { name: 'link.after', select: `${ELEVATED} .globalnav-submenu-list-item .globalnav-submenu-link`, props: [...TYPE, ...SPACE, 'display', ...MOTION], box: true, position: true, states: ['hover'] },
    { name: 'link', select: `${GROUP} .globalnav-submenu-link`, props: [...TYPE, ...SPACE, 'display', 'text-decoration-line', 'background-color', ...FOCUS, ...MOTION], box: true, position: true, states: ['hover', 'focus-visible'] },
    { name: 'curtain', select: '#globalnav-curtain', props: ['position', 'top', 'z-index', 'background-color', ...BLUR, 'opacity', 'visibility', 'pointer-events', 'transition-delay', ...MOTION], box: true, position: true },
    // the button beside the link the keyboard opens the flyout by: its box, its
    // chevron shown focused
    { name: 'trigger', select: TRIGGER, props: ['position', 'top', 'left', 'right', 'width', 'height', ...BOX, 'color', 'opacity', 'cursor', ...FLEX, ...FOCUS, ...MOTION], box: true, position: true, states: ['focus-visible', 'hover'] },
    { name: 'trigger.icon', select: `${TRIGGER} .globalnav-image-regular`, stateOn: TRIGGER, props: ['display', 'width', 'height', ...SPACE, 'opacity', 'color', ...MOTION], box: true, position: true, states: ['focus-visible', 'hover'] },
    { name: 'trigger.compact', select: `${TRIGGER} .globalnav-image-compact`, props: ['display'] },
    { name: 'trigger.group', select: '#globalnav .globalnav-item-mac .globalnav-submenu-trigger-group', props: ['position', 'display', 'width', 'height', ...SPACE, 'list-style-type'], box: true, position: true },
    { name: 'trigger.item', select: '#globalnav .globalnav-item-mac .globalnav-submenu-trigger-item', props: ['display', 'position'], box: true, position: true },
    // the change: the pointer along the bar to iPad's link
    { name: 'change', hover: '#globalnav .globalnav-link-ipad', pointer: ROOM, animations: MOTIONS, select: '#globalnav-submenu-link-ipad', props: ['height', 'background-color', 'visibility', ...MOTION], vars: ['--r-globalnav-flyout-height', '--r-globalnav-flyout-rate'], box: true },
    { name: 'change.previous', select: FLY, props: ['height', 'visibility'], box: true },
    { name: 'change.link', select: '#globalnav .globalnav-link-mac', props: ['color'] },
    // and off the nav, down to the page: it closes
    { name: 'close', hover: (w) => [w / 2, 890], hovered: 1200, animations: MOTIONS, select: '#globalnav-submenu-link-ipad', props: ['height', 'visibility', 'background-color'], box: true },
    { name: 'close.curtain', select: '#globalnav-curtain', props: ['opacity', 'visibility'] },
    { name: 'close.nav', select: '#globalnav', props: ['background-color', ...BLUR] },
  ],
};
