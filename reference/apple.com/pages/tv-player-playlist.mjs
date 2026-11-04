// tv.apple.com's video player's playlist: the paused player's Trailers tab
// clicked (tv-player's opening), the row of trailers it shows over the
// video's foot: the tab chosen and not, the row, a trailer's frame, its
// picture, its title and line, the paddle. No trailer is marked as the one
// playing (the second was, in the reading before this spec).
import { TYPE, BOX, FLEX, FOCUS, MOTION } from './common.mjs';
import { URL, open } from './tv-player.mjs';

const ITEM = '#playlist-tab-pane .shelf-grid__list-item:first-child';
export default {
  url: URL,
  browser: 'chrome',
  wait: 2000,
  widths: [1600, 1200],
  probes: [
    ...open,
    { name: 'tab', select: '#playlist-tab-trigger', props: [...TYPE, ...BOX, 'height', 'cursor', ...FOCUS], box: true, states: ['hover', 'focus-visible'] },
    { name: 'paused', click: 'button.playback-play__pause', clicked: 1500, pointer: [800, 200], select: 'button.playback-play__play', props: ['opacity'] },
    { name: 'tab.chosen', click: '#playlist-tab-trigger', clicked: 1500, pointer: [800, 200], select: '#playlist-tab-trigger', props: [...TYPE, ...BOX, 'height', 'cursor', 'aria-selected'], box: true, position: true, states: ['hover'] },
    { name: 'pane', select: '#playlist-tab-pane', props: [...BOX, ...MOTION, 'opacity'], box: true, position: true },
    { name: 'row', select: '#playlist-tab-pane .shelf-grid__list', props: [...FLEX, ...BOX, 'grid-auto-flow', 'grid-auto-columns', 'grid-template-columns', 'overflow-x', 'scroll-snap-type', 'scroll-padding-left', 'padding-left', 'padding-right'], box: true, position: true },
    { name: 'item', select: ITEM, props: [...BOX, 'padding-left', 'padding-right', 'scroll-snap-align'], box: true, position: true },
    { name: 'lockup', select: `${ITEM} .lockup`, props: [...BOX, 'overflow-x', 'cursor', 'aspect-ratio', ...FOCUS, ...MOTION, 'transform'], vars: ['--lockup-aspect-ratio', '--lockup-border-radius'], box: true, position: true, states: ['hover', 'focus-visible'] },
    { name: 'lockup.after', select: `${ITEM} .lockup`, pseudo: '::after', props: [...BOX, 'content', 'position', 'inset', 'opacity'] },
    { name: 'lockup.grid', select: `${ITEM} .lockup .grid`, props: [...BOX, 'overflow-x'], box: true, position: true },
    { name: 'artwork', select: `${ITEM} .artwork`, props: [...BOX, 'overflow-x'], box: true },
    { name: 'title', select: `${ITEM} .metadata .title`, props: [...TYPE, 'padding-left', 'padding-right'], box: true, position: true },
    { name: 'line', select: `${ITEM} .progress-and-attribution`, props: TYPE, box: true, position: true },
    { name: 'gradient', select: `${ITEM} .legibility-gradient`, props: ['background-image', 'opacity'], box: true, position: true },
    { name: 'paddle', select: '#playlist-tab-pane .shelf-grid-nav__arrow--right', props: [...BOX, 'color', 'opacity', 'width', 'height'], box: true, position: true, states: ['hover'] },
  ],
};
