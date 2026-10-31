// The console's banner against App Store Connect's toolkit's banner
// (--tb-component-Banner-*): its corners, room and least height, its title,
// body and caption, its tints, its icon's column, its close button's place,
// at the desktop's page scale and the mobile's. Golden, with their why: the
// icon's size and color, the actions' room, the dark side.

import { examples } from '../../../test/harness.mjs';

const ex = examples(import.meta.url);
const B = 'component-Banner';
const D = `${B}-anatomy`;
const BANNER = '.cui-console-banner';
const COLORS = [['info'], ['success'], ['error']];

export const cases = [
  {
    name: 'banner',
    html: ex('banner'),
    widths: [1600, 1200, 640],
    checks: [
      { target: BANNER, props: ['border-top-left-radius', 'padding-top', 'padding-bottom', 'padding-left', 'padding-right', 'min-height'], tb: [`${B}-borderRadius`, `${B}-paddingBlock`, `${B}-paddingBlock`, `${B}-paddingInlineStart`, `${B}-variant-pageScale-desktop-paddingInlineEnd`, `${B}-variant-pageScale-desktop-minHeight`] },
      { target: BANNER, props: ['background-color'], tb: [`${B}-variant-colorRole-warning-variant-tinted-backgroundColor`] },
      { target: '.cui-console-banner-title', props: ['font-size', 'line-height', 'font-weight', 'color'], tb: [`${D}-title-variant-pageScale-desktop-textStyle-fontSize`, `${D}-title-variant-pageScale-desktop-textStyle-lineHeight`, `${D}-title-variant-pageScale-desktop-textStyle-fontWeight`, `${D}-title-color`] },
      { target: '.cui-console-banner-body', props: ['margin-top', 'color'], tb: [`${D}-body-variant-pageScale-desktop-marginBlockStart`, `${D}-body-color`] },
      { target: '.cui-console-banner-icon', props: ['box-width'], tb: [`${D}-art-variant-layout-inline-pageScale-desktop-inlineSize`] },
      { target: '.cui-console-banner-icon', props: ['box-height', 'color'], golden: 'its icon 20px tall, the toolkit\'s art size at desktop, in the fitted label of its color (tokens.css, decision 21)' },
      { target: '.cui-console-banner-actions', props: ['column-gap'], tb: [`${D}-actions-variant-pageScale-desktop-gap`] },
      { target: '.cui-console-banner-actions', props: ['margin-top'], golden: 'its actions 8px down, as its body' },
      { target: '.cui-console-banner-dismiss', props: ['top', 'right'], tb: [`${D}-dismiss-insetBlockStart`, `${D}-dismiss-insetInlineEnd`] },
    ],
  },
  {
    name: 'banner phone',
    html: ex('banner'),
    widths: [639, 390],
    checks: [
      { target: BANNER, props: ['padding-right', 'min-height'], tb: [`${B}-variant-pageScale-mobile-paddingInlineEnd`, `${B}-variant-pageScale-mobile-minHeight`] },
      { target: '.cui-console-banner-title', props: ['font-size', 'line-height', 'font-weight'], tb: [`${D}-title-variant-pageScale-mobile-textStyle-fontSize`, `${D}-title-variant-pageScale-mobile-textStyle-lineHeight`, `${D}-title-variant-pageScale-mobile-textStyle-fontWeight`] },
      { target: '.cui-console-banner-body', props: ['margin-top'], tb: [`${D}-body-variant-pageScale-mobile-marginBlockStart`] },
      { target: '.cui-console-banner-icon', props: ['box-height'], golden: 'its icon 20px tall below 640px too: the toolkit\'s 30px is SF Symbols\' size, where a Lucide icon at 30px outweighs its 16px title' },
    ],
  },
  {
    name: 'banner colors',
    html: ex('colors'),
    widths: [1200, 390],
    checks: [
      { target: `${BANNER}:not([data-color])`, props: ['background-color'], tb: [`${B}-variant-colorRole-neutral-variant-tinted-backgroundColor`] },
      ...COLORS.map(([c]) => ({ target: `${BANNER}[data-color="${c}"]`, props: ['background-color'], tb: [`${B}-variant-colorRole-${c}-variant-tinted-backgroundColor`] })),
      ...COLORS.map(([c]) => ({ target: `${BANNER}[data-color="${c}"] .cui-console-banner-icon`, props: ['color'], golden: `its icon in the fitted ${c} label (tokens.css, decision 21)` })),
    ],
  },
  ...['desktop', 'mobile'].map((scale) => ({
    name: `banner caption ${scale}`,
    html: '<div class="cui-console-banner"><div><p class="cui-console-banner-caption">Updated 9:41</p></div></div>',
    widths: scale === 'desktop' ? [1200] : [390],
    checks: [
      { target: '.cui-console-banner-caption', props: ['font-size', 'line-height', 'font-weight'], tb: [`${D}-caption-variant-pageScale-${scale}-textStyle-fontSize`, `${D}-caption-variant-pageScale-${scale}-textStyle-lineHeight`, `${D}-caption-variant-pageScale-${scale}-textStyle-fontWeight`] },
      { target: '.cui-console-banner-caption', props: ['color'], golden: 'the secondary label fitted to 0.74 (tokens.css, decision 21)' },
    ],
  })),
  {
    name: 'banner dark',
    html: `<div class="cui-dark">${ex('banner')}${ex('colors')}</div>`,
    widths: [1200],
    checks: [
      { target: '[data-color="warning"]', props: ['background-color', 'color'], golden: 'the dark side derived (tokens.css, decision 18)' },
      { target: '[data-color="error"]', props: ['background-color'], golden: 'the dark side derived (tokens.css, decision 18)' },
    ],
  },
];
