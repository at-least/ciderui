import { examples } from '../../../test/harness.mjs';

const ex = examples(import.meta.url);
const D = 'component-DisplayList';
const L = `${D}-anatomy-Item-anatomy-Label-variant-pageScale`;
const TYPE = ['font-size', 'font-weight', 'line-height'];
const type = (scale) => ['fontSize', 'fontWeight', 'lineHeight'].map((p) => `${L}-${scale}-textStyle-${p}`);
const ROW = '.cui-console-list-row:nth-child(2)';

export const cases = [
  {
    name: 'list',
    html: ex('list'),
    widths: [1200, 640],
    checks: [
      { target: '.cui-console-list-label', props: ['display'], golden: 'its own line, over what follows it' },
      { target: ROW, props: ['padding-left', 'padding-right'], tb: [`${D}-groupedGray-contentPaddingInlineDesktop`, `${D}-groupedGray-contentPaddingInlineDesktop`] },
      { target: ROW, props: ['min-height', 'padding-top', 'padding-bottom'], tb: [`${D}-variant-pageScale-desktop-variant-sidebar-anatomy-Item-minHeight`, `${D}-variant-pageScale-desktop-variant-sidebar-anatomy-Item-paddingBlock`, `${D}-variant-pageScale-desktop-variant-sidebar-anatomy-Item-paddingBlock`] },
      { target: `${ROW} .cui-console-list-label`, props: TYPE, tb: type('desktop') },
      { target: `${ROW} .cui-console-list-label`, props: ['color'], tb: [`${D}-anatomy-Item-variant-colorRole-neutral-effect-default-anatomy-Label-color`] },
      { target: `${ROW} .cui-console-list-detail`, props: ['color', 'font-size', 'line-height'], golden: 'the secondary label at 0.76: the toolkit\'s 0.60 is 3.4:1 on white (tokens.css, decision 21); the body\'s 13px' },
      { target: ROW, pseudo: '::before', props: ['left', 'height', 'background-color'], golden: 'a 0.5px hairline over each row after the first, from 16px in, the secondary border (the table row\'s)' },
      { target: '.cui-console-list-row:first-child', pseudo: '::before', props: ['content'], golden: 'no hairline over the first row' },
      { target: '.cui-console-list', props: ['background-color', 'border-top-left-radius'], golden: 'the group the console card\'s gray and corners' },
      { target: ROW, props: ['background-color'], state: 'hover', golden: 'a row is no link and does not fill' },
    ],
  },
  {
    name: 'list phone',
    html: ex('list'),
    widths: [639, 390],
    checks: [
      { target: `${ROW} .cui-console-list-label`, props: TYPE, tb: type('mobile') },
      { target: ROW, props: ['padding-left'], tb: [`${D}-groupedGray-contentPaddingInlineMobile`] },
    ],
  },
  {
    name: 'list forced colors',
    html: ex('list'),
    widths: [1200],
    forcedColors: true,
    checks: [
      { target: ROW, pseudo: '::before', props: ['background-color', 'height'], golden: 'the hairline between rows CanvasText: forced colors drew its fill in the ground\'s color, at its alpha' },
    ],
  },
  {
    name: 'list content',
    html: ex('content'),
    widths: [1200],
    checks: [
      { target: '.cui-console-list-row', props: ['padding-left', 'padding-top', 'min-height', 'background-color'], golden: 'a row holding the page\'s own content after its label: a row\'s 16px in, 7px above, 44px at least, no fill of its own' },
    ],
  },
  {
    name: 'list dark',
    html: `<div class="cui-dark">${ex('list')}</div>`,
    widths: [1200],
    checks: [
      { target: '.cui-console-list', props: ['background-color'], golden: 'the dark side derived (tokens.css, decision 18)' },
      { target: `${ROW} .cui-console-list-label`, props: ['color'], golden: 'the dark side derived (tokens.css, decision 18)' },
      { target: ROW, pseudo: '::before', props: ['background-color'], golden: 'the dark side derived (tokens.css, decision 18)' },
    ],
  },
];
