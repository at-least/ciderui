// @external: tv.apple.com's video player's subtitles menu, opened by its button over
// the paused player (tv-player's opening): the menu, a group and its title,
// an item hovered and the chosen one with its check, a rule between groups,
// an item with a line under its name and a chevron (its language). At 1600px
// alone: a resize closes the menu.
import { TYPE, BOX, BLUR, FLEX } from './common.mjs';
import { URL, open } from './tv-player.mjs';

const MENU = 'amp-contextual-menu .contextual-menu--subtitles';
const ITEM = [...TYPE, ...BOX, ...FLEX, 'height', 'cursor'];
export default {
  url: URL,
  browser: 'chrome',
  wait: 2000,
  widths: [1600],
  probes: [
    ...open,
    { name: 'paused', click: 'button.playback-play__pause', clicked: 1500, pointer: [800, 200], select: 'button.playback-play__play', props: ['opacity'] },
    { name: 'menu', click: 'button[aria-label="Subtitles"]', clicked: 1000, pointer: [800, 200], select: MENU, props: [...BOX, ...BLUR, 'overflow-x', 'overflow-y', 'width', 'z-index'], box: true, position: true },
    { name: 'menu.before', select: MENU, pseudo: '::before', props: [...BOX, ...BLUR, 'content', 'position', 'inset'] },
    { name: 'menu.overlay', select: 'amp-contextual-menu .contextual-menu__overlay', props: [...BOX, ...BLUR, 'position'], box: true },
    { name: 'menu.list', select: `${MENU} .contextual-menu__list`, props: [...BOX, ...BLUR, 'list-style-type'], box: true },
    { name: 'group', select: `${MENU} .contextual-menu__group`, props: BOX, box: true },
    { name: 'group.next', select: `${MENU} .contextual-menu__group + .contextual-menu__group`, props: BOX, box: true },
    { name: 'group.title', select: `${MENU} .contextual-menu__group-title`, props: [...TYPE, ...BOX, 'text-transform'], box: true },
    { name: 'item', select: `${MENU} .contextual-menu-item button[title="On"]`, props: ITEM, box: true, states: ['hover', 'focus-visible', 'active'] },
    { name: 'item.row', select: `${MENU} .contextual-menu-item:has(button[title="On"])`, props: BOX, box: true, states: ['hover'] },
    { name: 'item.text', select: `${MENU} button[title="On"] .contextual-menu-item__option-text`, props: TYPE, box: true, position: true },
    { name: 'chosen', select: `${MENU} .contextual-menu-item button[title="Off"]`, props: ITEM, box: true },
    { name: 'chosen.check', select: `${MENU} button[title="Off"] .contextual-menu-item__icon-container`, props: ['color', 'width', 'height'], box: true, position: true },
    { name: 'sub', select: `${MENU} .contextual-menu-item--use-html-select label`, props: ITEM, box: true },
    { name: 'sub.subtext', select: `${MENU} .contextual-menu-item__subtext`, props: TYPE, box: true },
    { name: 'sub.chevron', select: `${MENU} .contextual-menu-item--use-html-select .contextual-menu-item__icon-container`, props: ['color', 'width', 'height'], box: true, position: true },
  ],
};
