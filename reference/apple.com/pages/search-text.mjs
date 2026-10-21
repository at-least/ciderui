// The global nav's search with text in it, opened as search.mjs opens it: the
// field given a value and its two buttons what apple.com's script gives them
// while it holds text (its _updateSearchFieldButtons: tabindex 0, no
// disabled, no aria-hidden), set by the probe (`set`), so nothing is typed
// and no suggestion is asked for. The input's text; the clear button shown
// (its place, size, fill and motion, its images), the search button's fill;
// both hovered and focused, last. Read only.
import search from './search.mjs';
import { TYPE, MOTION, FOCUS } from './common.mjs';

const SPACE = ['padding-top', 'padding-right', 'padding-bottom', 'padding-left', 'margin-top', 'margin-right', 'margin-bottom', 'margin-left'];
const FLY = '#globalnav-submenu-search';
const WIDE = [1600, 1200, 900, 834];
const PHONE = [833, 390];
const INPUT = `${FLY} .globalnav-searchfield-input`;
const RESET = `${FLY} .globalnav-searchfield-reset`;
const SUBMIT = `${FLY} .globalnav-searchfield-submit`;
const FILLED = [
  [INPUT, 'value', 'iPhone'],
  ...[RESET, SUBMIT].flatMap(b => [[b, 'tabindex', '0'], [b, 'disabled', null], [b, 'aria-hidden', null]]),
];
const BUTTON = [...SPACE, 'width', 'height', 'position', 'display', 'align-items', 'justify-content', 'fill', 'color', 'background-color', 'border-top-width', 'cursor', 'opacity', 'visibility', 'z-index', 'order', ...MOTION];

export default {
  ...search,
  probes: [
    ...search.probes.filter(p => ['nav', 'nav.phone'].includes(p.name)),
    { name: 'input', set: FILLED, select: INPUT, props: [...TYPE, ...SPACE, 'width', 'height', 'caret-color'], box: true, position: true },
    { name: 'reset', select: RESET, props: BUTTON, box: true, position: true },
    { name: 'reset.image', select: `${RESET} .globalnav-image-regular svg`, props: ['width', 'height', 'fill', 'display'], box: true, position: true, widths: WIDE },
    { name: 'reset.image-compact', select: `${RESET} .globalnav-image-compact svg`, props: ['width', 'height', 'fill', 'display'], box: true, position: true, widths: PHONE },
    { name: 'submit', select: SUBMIT, props: BUTTON, box: true, position: true },
    // hovered and focused last: the still style that holds a state cancels
    // the page's animations
    { name: 'reset.states', select: RESET, props: ['fill', ...FOCUS], states: ['hover', 'focus-visible'] },
    { name: 'submit.states', select: SUBMIT, props: ['fill', ...FOCUS], states: ['hover', 'focus-visible'] },
  ],
};
