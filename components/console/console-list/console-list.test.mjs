// The console's list against App Store Connect's toolkit's display list
// (--tb-component-DisplayList-*: groupedGray, its item): the 16px in, the
// label's type and color, the pressed fill, the row height (its sidebar
// items'). Golden, with their why: the hairline, the hovered fill, the
// secondary label's fitted alpha, the group's gray and corners, the dark
// side.

import { examples } from '../../../test/harness.mjs';

const ex = examples(import.meta.url);
const D = 'component-DisplayList';
const L = `${D}-anatomy-Item-anatomy-Label-variant-pageScale`;
const TYPE = ['font-size', 'font-weight', 'line-height'];
const type = (scale) => ['fontSize', 'fontWeight', 'lineHeight'].map((p) => `${L}-${scale}-textStyle-${p}`);
const LINK = '.cui-console-list-item[href="#api"]';
const ROW = 'div.cui-console-list-item';

export const cases = [
  {
    name: 'list',
    html: ex('list'),
    widths: [1200, 640],
    checks: [
      { target: LINK, props: ['padding-left', 'padding-right'], tb: [`${D}-groupedGray-contentPaddingInlineDesktop`, `${D}-groupedGray-contentPaddingInlineDesktop`] },
      { target: LINK, props: ['min-height', 'padding-top', 'padding-bottom'], tb: [`${D}-variant-pageScale-desktop-variant-sidebar-anatomy-Item-minHeight`, `${D}-variant-pageScale-desktop-variant-sidebar-anatomy-Item-paddingBlock`, `${D}-variant-pageScale-desktop-variant-sidebar-anatomy-Item-paddingBlock`] },
      { target: `${LINK} .cui-console-list-label`, props: TYPE, tb: type('desktop') },
      { target: `${LINK} .cui-console-list-label`, props: ['color'], tb: [`${D}-anatomy-Item-variant-colorRole-neutral-effect-default-anatomy-Label-color`] },
      { target: LINK, props: ['background-color'], state: 'active', tb: [`${D}-anatomy-Item-activeBackgroundColor`] },
      { target: LINK, props: ['background-color'], state: 'hover', golden: 'hovered, the pressed fill: the toolkit gives a grouped row no hovered one' },
      { target: `${LINK} .cui-console-list-secondary`, props: ['color', 'font-size', 'line-height'], golden: 'the secondary label at 0.74: the toolkit\'s 0.60 is 3.4:1 on white (tokens.css, decision 21); the body\'s 13px' },
      { target: LINK, pseudo: '::before', props: ['left', 'height', 'background-color'], golden: 'a 0.5px hairline over each row after the first, from 16px in, the secondary border (the table row\'s)' },
      { target: '.cui-console-list-item[href="#shop"]', pseudo: '::before', props: ['content'], golden: 'no hairline over the first row' },
      { target: '.cui-console-list', props: ['background-color', 'border-top-left-radius'], golden: 'the group the console card\'s gray and corners' },
      { target: ROW, props: ['background-color'], state: 'hover', golden: 'a row that is no link does not fill' },
    ],
  },
  {
    name: 'list phone',
    html: ex('list'),
    widths: [639, 390],
    checks: [
      { target: `${LINK} .cui-console-list-label`, props: TYPE, tb: type('mobile') },
      { target: LINK, props: ['padding-left'], tb: [`${D}-groupedGray-contentPaddingInlineMobile`] },
    ],
  },
  {
    name: 'list touch',
    html: ex('list'),
    touch: true,
    widths: [390],
    checks: [
      { target: LINK, props: ['background-color'], state: 'hover', rest: 'a tapped row keeps no hover (rule 5)' },
    ],
  },
  {
    name: 'list forced colors',
    html: ex('list'),
    widths: [1200],
    forcedColors: true,
    checks: [
      { target: LINK, pseudo: '::before', props: ['background-color', 'height'], golden: 'the hairline between rows CanvasText: forced colors drew its fill in the ground\'s color, at its alpha' },
    ],
  },
  {
    name: 'list dark',
    html: `<div class="cui-dark">${ex('list')}</div>`,
    widths: [1200],
    checks: [
      { target: '.cui-console-list', props: ['background-color'], golden: 'the dark side derived (tokens.css, decision 18)' },
      { target: `${LINK} .cui-console-list-label`, props: ['color'], golden: 'the dark side derived (tokens.css, decision 18)' },
      { target: LINK, pseudo: '::before', props: ['background-color'], golden: 'the dark side derived (tokens.css, decision 18)' },
    ],
  },
];
