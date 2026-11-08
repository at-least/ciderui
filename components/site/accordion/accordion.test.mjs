import { examples, MOTION, SIZE } from '../../../test/harness.mjs';

const ex = examples(import.meta.url);
const ITEM = '.cui-accordion-item';
const OPEN = '.cui-accordion-item[open]';
const CLOSED = '.cui-accordion-item:not([open])';
const Q = '.cui-accordion-question';
// @invariant: an item renders closed; a case reads an open one after a click opens the first
const OPEN_FIRST = `${ITEM}:nth-child(1) > ${Q}`;

export const cases = [
  {
    name: 'zh-hant',
    html: `<div lang="zh-Hant">${ex('accordion')}</div>`,
    click: OPEN_FIRST,
    checks: [{ target: `${OPEN} .cui-accordion-answer`, props: ['line-height', 'letter-spacing'], apple: ['buy-tw', 'accordion.content'] }],
  },
  {
    name: 'accordion',
    html: ex('accordion'),
    click: OPEN_FIRST,
    checks: [
      { target: ITEM, props: ['border-top-width', 'border-top-style', 'border-top-color', 'padding-top', 'margin-top'], apple: ['bag', 'accordion.item'] },
      { target: `${CLOSED} > ${Q}`, props: ['padding-top', 'padding-bottom', 'display', 'cursor', 'box-height'], apple: ['bag-accordion', 'compact.button'] },
      { target: `${CLOSED} > ${Q}`, props: [...SIZE, 'color'], apple: ['bag-accordion', 'compact.title'] },
      { target: `${CLOSED} > ${Q} svg`, props: ['color'], apple: ['bag-accordion', 'compact.icon'] },
      { target: `${CLOSED} > ${Q} svg`, props: ['box-height'], apple: ['bag-accordion', 'compact.svg'] },
      { target: `${CLOSED} > ${Q} svg`, props: ['transition-property', 'transition-duration', 'transition-timing-function'], golden: 'its turn in 0.32s, its color at once, where apple.com\'s eases in 0.1s linear (bag-accordion:compact.icon): motion that adds nothing to a pointer\'s color change (DESIGN.md decision 25)' },
      { target: `${CLOSED} > ${Q} svg`, on: `${CLOSED} > ${Q}`, state: 'hover', props: ['color'], apple: ['bag-accordion', 'compact.icon-hover', 'hover'] },
      { target: `${OPEN} > ${Q} svg`, props: ['color'], apple: ['bag-accordion', 'compact.open'] },
      { target: `${OPEN} .cui-accordion-answer`, props: [...SIZE, 'color', 'padding-right', 'padding-bottom'], apple: ['bag-accordion', 'compact.answer'] },
      { target: `${OPEN} .cui-accordion-answer`, props: ['padding-top'], as: { 'padding-top': 'margin-top' }, apple: ['bag-accordion', 'compact.panel'] },
      { target: `${OPEN} > ${Q} svg`, props: ['transform'], golden: 'open, turned up: apple.com\'s line moved to the same points, through a flat one, in 320ms' },
      { target: `${CLOSED} > ${Q} svg`, props: ['box-width', 'margin-right'], golden: 'apple.com\'s 17px chevron (14px) in the middle of its 32px column at the end (bag-accordion:compact.svg, .icon)' },
      { target: ITEM, pseudo: '::details-content', props: MOTION, golden: 'its answer opens at once, where apple.com\'s height slides in 0.3s ease-in-out (bag-accordion:compact.panel): the chevron\'s turn says it opened, and the words are there to read (DESIGN.md decision 25)' },
      { target: Q, props: ['list-style-type'], golden: 'no marker: the chevron is the question\'s' },
    ],
  },
  {
    name: 'click',
    html: ex('accordion'),
    widths: [1200],
    click: '.cui-accordion-item:nth-child(2) > .cui-accordion-question',
    checks: [
      { target: '.cui-accordion-item:nth-child(2) > .cui-accordion-question svg', props: ['transform'], golden: 'clicked, its question\'s chevron turned up: <details> opens with no script' },
      { target: '.cui-accordion-item:nth-child(2) .cui-accordion-answer', props: ['box-height'], golden: 'its answer shown, at its height (13.6px above, 28px below its line)' },
    ],
  },
];
