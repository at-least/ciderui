// The console's tabs against App Store Connect's toolkit's tab bar
// (--tb-component-Tabs-*, brand: underlined and filled-gray), at rest,
// pressed, current and disabled, at each size. Golden, with their why: the
// colors fitted to 4.5:1, the hovered ones, the dark side.

import { examples } from '../../test/harness.mjs';

const ex = examples(import.meta.url);
const T = 'component-Tabs';
const V = `${T}-variant-colorRole-brand-variant`;
const tab = (variant, effect, active, prop) => `${V}-${variant}-anatomy-Tab-variant-effect-${effect}-isActive-${active}-${prop}`;
const size = (s) => [`${T}-variant-controlSize-${s}-anatomy-Tab-height`, `${T}-variant-controlSize-${s}-anatomy-Tab-minWidth`, `${T}-variant-controlSize-${s}-anatomy-Tab-gap`, `${T}-variant-controlSize-${s}-textStyle-fontSize`, `${T}-variant-controlSize-${s}-textStyle-lineHeight`, `${T}-variant-controlSize-${s}-textStyle-fontWeight`];
const SIZE = ['box-height', 'min-width', 'column-gap', 'font-size', 'line-height', 'font-weight'];
const FITTED = 'fitted to 4.5:1 (tokens.css, decision 21)';
const REST = '.cui-console-tab[href="#issues"]';
const PILL = '.cui-console-tab[href="#30d"]';
const CHOSEN = '.cui-console-tab[aria-current="true"]';

export const cases = [
  {
    name: 'tabs',
    html: ex('tabs'),
    widths: [1600, 1200, 390],
    checks: [
      { target: REST, props: SIZE, tb: size('md') },
      { target: REST, props: ['color', 'padding-left', 'padding-right'], tb: [tab('underlined', 'default', 'false', 'color'), `${T}-underlinedPaddingInline`, `${T}-underlinedPaddingInline`] },
      { target: REST, props: ['color'], state: 'active', golden: 'pressed, the plain button\'s pressed blue (tokens.css\'s brand label pressed), where the toolkit\'s tab presses to its gray button\'s, oklch(52.2% 0.156 251.8), a hair lighter: one pressed blue for a label' },
      { target: REST, props: ['color'], state: 'hover', golden: `hovered, the brand's label a step darker, blue-600 (the button's hovered blue): the toolkit's lighter one is 4.1:1 on white, ${FITTED}` },
      { target: '.cui-console-tabs', props: ['column-gap', 'border-bottom-width', 'border-bottom-style', 'border-bottom-color'], tb: [`${T}-variant-controlSize-md-variant-underlined-gap`, `${T}-underlinedFullWidthKeyline-width`, `${T}-underlinedFullWidthKeyline-style`, `${T}-underlinedFullWidthKeyline-color`] },
      { target: '[aria-current="page"]', props: ['color', 'border-bottom-width'], tb: [tab('underlined', 'default', 'true', 'color'), `${T}-underlinedActiveBorderBottomWidth`] },
      { target: '[aria-current="page"]', props: ['border-bottom-color', 'border-bottom-style'], golden: 'the current tab\'s rule in its label\'s color' },
      { target: '[aria-current="page"]', props: ['color'], state: 'hover', tb: [tab('underlined', 'hover', 'true', 'color')] },
      { target: '[aria-disabled="true"]', props: ['color'], tb: [tab('underlined', 'disabled', 'false', 'color')] },
      { target: '[aria-disabled="true"]', props: ['cursor'], golden: 'no pointer over a disabled tab' },
    ],
  },
  {
    name: 'tabs many on a phone',
    html: ex('many'),
    widths: [375],
    checks: [
      { target: '.cui-console-tabs', props: ['flex-wrap', 'overflow-x', 'box-height'], golden: 'a phone: one row (no second hairline), scrolling sideways' },
      { target: '.cui-console-tabs', props: ['mask-image'], golden: 'no overflow state from the host (no-JS among it), the row unfaded' },
      { target: '.cui-console-tab[href="#settings"]', props: ['outline-offset'], state: 'focus-visible', golden: 'its focus ring inside the tab, where the scroll box would clip it' },
      { target: '.cui-console-tabs', props: ['scroll-padding-left', 'scroll-padding-right'], golden: 'a focused tab is brought in clear of the 24px fade: a browser\'s focus scrolls only a tab it judges out of view, and Chromium 151 left one 3px in view where it was (reference/console/)' },
    ],
  },
  {
    name: 'tabs overflow fade start',
    html: ex('many').replace('cui-console-tabs"', 'cui-console-tabs" data-overflow data-overflow-start'),
    widths: [375],
    checks: [
      { target: '.cui-console-tabs', props: ['mask-image'], golden: 'scrolled to its start, the right edge fades over the content still hidden; the left edge stays clear' },
    ],
  },
  {
    name: 'tabs overflow fade end',
    html: ex('many').replace('cui-console-tabs"', 'cui-console-tabs" data-overflow data-overflow-end'),
    widths: [375],
    checks: [
      { target: '.cui-console-tabs', props: ['mask-image'], golden: 'scrolled to its end, the left edge fades over the sections clipped behind; the right edge stays clear' },
    ],
  },
  {
    name: 'tabs overflow fade middle and desktop',
    html: ex('many').replace('cui-console-tabs"', 'cui-console-tabs" data-overflow'),
    widths: [375, 1200],
    checks: [
      { target: '.cui-console-tabs', props: ['mask-image'], golden: 'between the ends, both edges fade; the fade is the phone row\'s — wide, none' },
    ],
  },
  {
    name: 'tabs many wide',
    html: ex('many'),
    widths: [1200],
    checks: [
      { target: '.cui-console-tabs', props: ['flex-wrap', 'overflow-x'], golden: 'wide, the bar as the toolkit draws it' },
      { target: '.cui-console-tab[href="#settings"]', props: ['outline-offset'], state: 'focus-visible', golden: 'wide, base.css\'s ring outside the tab' },
    ],
  },
  {
    name: 'filled',
    html: ex('filled'),
    widths: [1200, 390],
    checks: [
      { target: PILL, props: SIZE, tb: size('md') },
      { target: PILL, props: ['background-color', 'padding-left', 'padding-right', 'border-top-left-radius'], tb: [tab('filledGray', 'default', 'false', 'backgroundColor'), `${T}-defaultMdPaddingInline`, `${T}-defaultMdPaddingInline`, `${T}-variant-controlShape-capsule-controlSize-md-borderRadius`] },
      { target: PILL, props: ['color'], golden: `the blue label on the gray pill blue-600: the toolkit's brand blue is 4.4:1 on it, ${FITTED}` },
      { target: PILL, props: ['background-color'], state: 'hover', tb: [tab('filledGray', 'hover', 'false', 'backgroundColor')] },
      { target: PILL, props: ['color'], state: 'hover', golden: `hovered, blue-600 too: the toolkit's lighter blue is 3.5:1 on white, ${FITTED}` },
      { target: PILL, props: ['background-color'], state: 'active', tb: [tab('filledGray', 'press', 'false', 'backgroundColor')] },
      { target: PILL, props: ['color'], state: 'active', golden: 'pressed, one pressed blue for a label (as an underlined tab\'s)' },
      { target: '.cui-console-tabs', props: ['column-gap'], tb: [`${T}-variant-controlSize-md-variant-filledGray-gap`] },
      { target: '.cui-console-tabs', props: ['border-bottom-width'], golden: 'no hairline under pills' },
      { target: CHOSEN, props: ['background-color', 'color'], tb: [tab('filledGray', 'default', 'true', 'backgroundColor'), tab('filledGray', 'default', 'true', 'color')] },
      { target: CHOSEN, props: ['background-color'], state: 'hover', golden: `the chosen pill hovered blue-600, a step darker: white on the toolkit's lighter blue is 3.9:1, ${FITTED}` },
      { target: CHOSEN, props: ['background-color', 'color'], state: 'active', tb: [tab('filledGray', 'press', 'true', 'backgroundColor'), tab('filledGray', 'press', 'true', 'color')] },
    ],
  },
  {
    name: 'large tabs',
    html: ex('large'),
    widths: [1200],
    checks: [
      { target: '[aria-current="page"]', props: SIZE, tb: size('lg') },
      { target: '.cui-console-tabs', props: ['column-gap'], tb: [`${T}-variant-controlSize-lg-variant-underlined-gap`] },
    ],
  },
  {
    name: 'tab sizes',
    html: ['sm', 'lg', 'xl'].map((s) => ex('filled').replace('data-variant="filled"', `data-variant="filled" data-size="${s}"`)).join(''),
    widths: [1200],
    checks: [
      ...['sm', 'lg', 'xl'].map((s) => ({ target: `[data-size="${s}"] > ${PILL}`, props: [...SIZE, 'padding-left', 'border-top-left-radius'], tb: [...size(s), `${T}-default${s[0].toUpperCase()}${s.slice(1)}PaddingInline`, `${T}-variant-controlShape-capsule-controlSize-${s}-borderRadius`] })),
      ...['sm', 'lg', 'xl'].map((s) => ({ target: `[data-size="${s}"]`, props: ['column-gap'], tb: [`${T}-variant-controlSize-${s}-variant-filledGray-gap`] })),
    ],
  },
  {
    name: 'tabs touch',
    html: ex('filled'),
    touch: true,
    widths: [390],
    checks: [
      { target: PILL, props: ['background-color', 'color'], state: 'hover', rest: 'a tapped tab keeps no hover (rule 5)' },
    ],
  },
  {
    name: 'tabs dark',
    html: `<div class="cui-dark">${ex('tabs')}${ex('filled')}</div>`,
    widths: [1200],
    checks: [
      { target: REST, props: ['color'], golden: 'the dark side derived (tokens.css, decision 18)' },
      { target: PILL, props: ['background-color', 'color'], golden: 'the dark side derived (tokens.css, decision 18)' },
      { target: CHOSEN, props: ['background-color', 'color'], golden: `the chosen pill dark blue-400: white on the derived blue-600 is 2.7:1, ${FITTED}` },
    ],
  },
];
