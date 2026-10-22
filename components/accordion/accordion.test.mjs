// The accordion against apple.com's bag (bag:accordion.item, its rule;
// bag-accordion:*: a compact item's question, its chevron at rest and with
// the question hovered, opened its answer; the outer item, a large
// question). Golden, with their why: the letter-spacing, the chevron's box
// and turn, the height's motion, no marker.

import { examples } from '../../test/harness.mjs';

const ex = examples(import.meta.url);
const TYPE = ['font-size', 'line-height', 'font-weight'];
const ITEM = '.cui-accordion-item';
const OPEN = '.cui-accordion-item[open]';
const CLOSED = '.cui-accordion-item:not([open])';
const Q = '.cui-accordion-question';
const TRACKING = 'the system font\'s fitted letter-spacing (tokens.css), not apple.com\'s';

export const cases = [
  {
    name: 'accordion',
    html: ex('accordion'),
    checks: [
      { target: ITEM, props: ['border-top-width', 'border-top-style', 'border-top-color', 'padding-top', 'margin-top'], apple: ['bag', 'accordion.item'] },
      { target: `${CLOSED} > ${Q}`, props: ['padding-top', 'padding-bottom', 'display', 'cursor', 'box-height'], apple: ['bag-accordion', 'compact.button'] },
      { target: `${CLOSED} > ${Q}`, props: [...TYPE, 'color'], apple: ['bag-accordion', 'compact.title'] },
      { target: `${CLOSED} > ${Q} svg`, props: ['color'], apple: ['bag-accordion', 'compact.icon'] },
      { target: `${CLOSED} > ${Q} svg`, props: ['box-height'], apple: ['bag-accordion', 'compact.svg'] },
      { target: `${CLOSED} > ${Q} svg`, props: ['transition-property', 'transition-duration', 'transition-timing-function'], golden: 'its color in apple.com\'s 0.1s linear (bag-accordion:compact.icon), and its turn in 0.32s' },
      { target: `${CLOSED} > ${Q} svg`, on: `${CLOSED} > ${Q}`, state: 'hover', props: ['color'], apple: ['bag-accordion', 'compact.icon-hover', 'hover'] },
      { target: `${OPEN} > ${Q} svg`, props: ['color'], apple: ['bag-accordion', 'compact.open'] },
      { target: `${OPEN} .cui-accordion-answer`, props: [...TYPE, 'color', 'padding-right', 'padding-bottom'], apple: ['bag-accordion', 'compact.answer'] },
      { target: `${OPEN} .cui-accordion-answer`, props: ['padding-top'], as: { 'padding-top': 'margin-top' }, apple: ['bag-accordion', 'compact.panel'] },
      { target: `${OPEN} > ${Q} svg`, props: ['transform'], golden: 'open, turned up: apple.com\'s line moved to the same points, through a flat one, in 320ms' },
      { target: `${CLOSED} > ${Q} svg`, props: ['box-width', 'margin-right'], golden: 'apple.com\'s 17px chevron (14px) in the middle of its 32px column at the end (bag-accordion:compact.svg, .icon)' },
      { target: ITEM, pseudo: '::details-content', props: ['transition-property', 'transition-duration', 'transition-timing-function'], golden: 'its answer\'s height in apple.com\'s 0.3s ease-in-out (bag-accordion:compact.panel), where the browser can move a height to auto' },
      { target: ITEM, pseudo: '::details-content', props: ['overflow-x', 'overflow-y'], golden: 'its answer\'s height clipped as it moves, not its width: the ring of what begins a line in it, a link or a question in the large one\'s answer, is drawn 3px outside it (reference/details/)' },
      { target: Q, props: ['list-style-type'], golden: 'no marker: the chevron is the question\'s' },
      { target: Q, props: ['letter-spacing'], golden: TRACKING },
      { target: `${OPEN} .cui-accordion-answer`, props: ['letter-spacing'], golden: TRACKING },
    ],
  },
  {
    name: 'large',
    html: ex('large'),
    checks: [
      { target: Q, props: TYPE, apple: ['bag-accordion', 'accordion.title'] },
      { target: `${Q} svg`, props: ['color'], apple: ['bag-accordion', 'accordion.icon'] },
      { target: `${Q} svg`, props: ['box-height'], apple: ['bag-accordion', 'accordion.svg'] },
      { target: Q, props: ['padding-top', 'padding-bottom'], apple: ['bag', 'accordion.button'] },
      { target: `${Q} svg`, props: ['box-width', 'margin-right'], golden: 'apple.com\'s chevron, 32px (24px in a 32px column below 735px: bag-accordion:accordion.svg, .icon)' },
      { target: Q, props: ['letter-spacing'], golden: TRACKING },
    ],
  },
  {
    // the bag's structure: the compact FAQ as the large question's answer
    name: 'nested',
    html: ex('large').replace('<details class="cui-accordion-item">', '<details class="cui-accordion-item" open>').replace('<p>Answers to the most asked questions.</p>', ex('accordion')),
    checks: [
      { target: `.cui-accordion[data-size="lg"] > ${ITEM} > ${Q}`, props: TYPE, apple: ['bag-accordion', 'accordion.title'] },
      { target: `.cui-accordion:not([data-size]) > ${CLOSED} > ${Q}`, props: ['padding-top', 'padding-bottom'], apple: ['bag-accordion', 'compact.button'] },
      { target: `.cui-accordion:not([data-size]) > ${CLOSED} > ${Q}`, props: TYPE, apple: ['bag-accordion', 'compact.title'] },
      { target: `.cui-accordion:not([data-size]) > ${CLOSED} > ${Q} svg`, props: ['box-height'], apple: ['bag-accordion', 'compact.svg'] },
    ],
  },
  {
    // the second question clicked: its answer opens, its chevron turns
    name: 'click',
    html: ex('accordion'),
    widths: [1200],
    click: '.cui-accordion-item:nth-child(2) > .cui-accordion-question',
    checks: [
      { target: '.cui-accordion-item:nth-child(2) > .cui-accordion-question svg', props: ['transform'], golden: 'clicked, its question\'s chevron turned up: <details> opens with no script' },
      { target: '.cui-accordion-item:nth-child(2) .cui-accordion-answer', props: ['box-height'], golden: 'its answer shown, at its height (13.6px above, 28px below its line)' },
    ],
  },
  {
    name: 'touch',
    html: ex('accordion'),
    touch: true,
    checks: [
      { target: `${CLOSED} > ${Q} svg`, on: `${CLOSED} > ${Q}`, state: 'hover', props: ['color'], rest: 'a touch screen\'s chevron does not change hovered (rule 5), where apple.com\'s does (touch-bag-accordion:compact.icon-hover)' },
    ],
  },
];
