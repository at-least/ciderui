// The bag's accordion on a touch screen at 390px, its outer item tapped open,
// then its first compact item (`tap`, the DevTools protocol's touch events: a
// touch screen's tap, not a forced :hover): what the tap leaves hovered, its
// chevron's color against bag-accordion:compact.open (opened by a click, the
// pointer gone: #86868b) and compact.icon-hover (its button's hover forced:
// #6e6e73). Read only.
import bag from './bag-accordion.mjs';

const OUTER = '.rc-accordion-item > h2 .rc-accordion-button';
const FIRST = '.rc-accordion-compact > .rc-accordion-item:first-child';

export default {
  url: bag.url,
  wait: bag.wait,
  touch: true,
  widths: [390],
  probes: [
    { name: 'accordion', tap: OUTER, select: OUTER, props: ['cursor'], matches: [':hover', ':focus-visible'] },
    { name: 'compact', tap: `${FIRST} .rc-accordion-button`, select: `${FIRST} .rc-accordion-button`, props: ['cursor'], matches: [':hover', ':active', ':focus', ':focus-visible'] },
    { name: 'compact.icon', select: `${FIRST} .rc-accordion-chevrondown`, props: ['color'] },
    // the item the tap opened: its answer's panel
    { name: 'compact.panel', select: `${FIRST} > div`, props: ['height'], box: true },
  ],
};
