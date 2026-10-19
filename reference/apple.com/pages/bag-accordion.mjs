// The bag's accordion, opened (clicked once, at the first width: what it
// opens stays so, so these probes are a spec of their own): the outer item
// (the bag's questions, a large title) and its chevron, then, opened, its
// compact items (a question, its chevron), the first of those opened (its
// answer, the chevron's color).
import { TYPE, BOX, MOTION } from './common.mjs';
export default {
  url: 'https://www.apple.com/shop/bag',
  wait: 2500,
  probes: [
    // the accordion, its parts: the outer item (the bag's questions, a large
    // title) and, opened, its compact items (a question, its chevron), the
    // first of those opened (its answer, the chevron turned)
    { name: 'accordion.title', select: '.rc-accordion-item > h2 .rc-accordion-title', props: TYPE },
    { name: 'accordion.icon', select: '.rc-accordion-item > h2 .rc-accordion-chevrondown', props: ['color', 'margin-left', ...MOTION], box: true },
    { name: 'accordion.svg', select: '.rc-accordion-item > h2 .accordion-icon-svg', props: ['width', 'height'], box: true },
    { name: 'accordion.line', select: '.rc-accordion-item > h2 .accordion-icon-svg polyline', props: ['stroke-width', 'stroke-linecap', 'stroke-linejoin'] },
    { name: 'accordion.open', click: '.rc-accordion-item > h2 .rc-accordion-button', select: '.rc-accordion-item > h2 .rc-accordion-chevrondown', props: ['color'] },
    { name: 'accordion.panel', select: '.rc-accordion-item > .r-height-transition-enter-done', props: [...MOTION, 'overflow'], box: true },
    { name: 'compact.item', select: '.rc-accordion-compact > .rc-accordion-item', props: BOX, box: true },
    { name: 'compact.button', select: '.rc-accordion-compact > .rc-accordion-item .rc-accordion-button', props: ['padding-top', 'padding-bottom', 'display', 'text-align', 'cursor'], box: true },
    { name: 'compact.title', select: '.rc-accordion-compact .rc-accordion-title', props: TYPE },
    { name: 'compact.icon', select: '.rc-accordion-compact .rc-accordion-chevrondown', props: ['color', ...MOTION], box: true, states: ['hover'] },
    // its button hovered: the chevron's color (the button's hover, not the icon's)
    { name: 'compact.icon-hover', select: '.rc-accordion-compact .rc-accordion-chevrondown', stateOn: '.rc-accordion-compact > .rc-accordion-item .rc-accordion-button', props: ['color'], states: ['hover'] },
    { name: 'compact.svg', select: '.rc-accordion-compact .accordion-icon-svg', props: ['width', 'height'], box: true },
    { name: 'compact.line', select: '.rc-accordion-compact .accordion-icon-svg polyline', props: ['stroke-width', 'stroke-linecap', 'stroke-linejoin'] },
    { name: 'compact.open', click: '.rc-accordion-compact > .rc-accordion-item:first-child .rc-accordion-button', select: '.rc-accordion-compact > .rc-accordion-item:first-child .rc-accordion-chevrondown', props: ['color'] },
    { name: 'compact.panel', select: '.rc-accordion-compact > .rc-accordion-item:first-child > div', props: ['margin-top', ...MOTION], box: true },
    { name: 'compact.answer', select: '.rc-accordion-compact > .rc-accordion-item:first-child .rc-accordion-content', props: [...TYPE, 'padding-top', 'padding-right', 'padding-bottom', 'padding-left'], box: true },
  ],
};
