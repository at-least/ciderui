import { examples } from '../../../test/harness.mjs';

const ex = examples(import.meta.url);
const T = 'component-Tabs';
// @external: the toolkit's neutral filled tab: the picked one black under white, the
// others the neutral label on gray; blue stays for the page's primary action
const V = `${T}-variant-colorRole-neutral-variant`;
const tab = (variant, effect, active, prop) => `${V}-${variant}-anatomy-Tab-variant-effect-${effect}-isActive-${active}-${prop}`;
// @external: App Store Connect's version page draws its underlined tabs' labels gray, rgb(51, 51, 51)
// (asc version: tab.button.text), the current one marked by its rule alone: the toolkit's neutral tab
const underlined = (effect, active, prop) => `${T}-variant-colorRole-neutral-variant-underlined-anatomy-Tab-variant-effect-${effect}-isActive-${active}-${prop}`;
const size = (s) => [`${T}-variant-controlSize-${s}-anatomy-Tab-height`, `${T}-variant-controlSize-${s}-anatomy-Tab-minWidth`, `${T}-variant-controlSize-${s}-anatomy-Tab-gap`, `${T}-variant-controlSize-${s}-textStyle-fontSize`, `${T}-variant-controlSize-${s}-textStyle-lineHeight`, `${T}-variant-controlSize-${s}-textStyle-fontWeight`];
const SIZE = ['box-height', 'min-width', 'column-gap', 'font-size', 'line-height', 'font-weight'];
const REST = '.cui-console-tab[href="#issues"]';
const RADIO = '.cui-console-tabs[aria-label="About"] .cui-console-tab';
const PICKED = `${RADIO}:has(> :checked)`;
const PLAIN = `${RADIO}:not(:has(> :checked))`;

export const cases = [
  {
    name: 'tabs',
    html: ex('tabs'),
    widths: [1600, 1200, 390],
    checks: [
      { target: REST, props: SIZE, tb: size('md') },
      { target: REST, props: ['color', 'padding-left', 'padding-right'], tb: [underlined('default', 'false', 'color'), `${T}-underlinedPaddingInline`, `${T}-underlinedPaddingInline`] },
      { target: REST, props: ['border-bottom-width', 'border-bottom-style', 'border-bottom-color'], asc: ['version', 'tab.button'] },
      { target: REST, props: ['color'], state: 'active', tb: [underlined('press', 'false', 'color')] },
      { target: REST, props: ['color'], state: 'hover', golden: 'hovered, the gray button\'s hovered label (tokens.css\'s neutral label hovered), where the toolkit\'s tab hovers to oklch(23.8% 0 0 / 0.85): one hovered gray for a label' },
      { target: '.cui-console-tabs', props: ['column-gap', 'border-bottom-width', 'border-bottom-style', 'border-bottom-color'], tb: [`${T}-variant-controlSize-md-variant-underlined-gap`, `${T}-underlinedFullWidthKeyline-width`, `${T}-underlinedFullWidthKeyline-style`, `${T}-underlinedFullWidthKeyline-color`] },
      { target: '[aria-current="page"]', props: ['color', 'border-bottom-width'], tb: [underlined('default', 'true', 'color'), `${T}-underlinedActiveBorderBottomWidth`] },
      { target: '[aria-current="page"]', props: ['border-bottom-width', 'border-bottom-style'], asc: ['version', 'tab.button.current'] },
      { target: '[aria-current="page"]', props: ['border-bottom-color'], golden: 'the current tab\'s rule in its label\'s color, as App Store Connect\'s (asc version: tab.button.current, its label\'s rgb(51, 51, 51))' },
      { target: '[aria-current="page"]', props: ['color'], state: 'hover', tb: [underlined('hover', 'true', 'color')] },
      { target: '[aria-current="page"]', props: ['color'], state: 'active', tb: [underlined('press', 'true', 'color')] },
      { target: '[aria-disabled="true"]', props: ['color'], tb: [underlined('disabled', 'false', 'color')] },
      { target: '[aria-disabled="true"]', props: ['cursor'], golden: 'no pointer over a disabled tab' },
    ],
  },
  {
    name: 'tabs many on a phone',
    html: ex('many'),
    widths: [375],
    checks: [
      { target: '.cui-console-tabs', props: ['flex-wrap', 'overflow-x', 'box-height'], golden: 'a phone: one row (no second hairline), scrolling sideways' },
      { target: '.cui-console-tabs', props: ['scrollbar-width'], golden: 'no scrollbar: the tab clipped at the row\'s edge says more is there' },
      { target: '.cui-console-tab[href="#settings"]', props: ['outline-offset'], state: 'focus-visible', golden: 'its focus ring inside the tab, where the scroll box would clip it' },
      { target: '.cui-console-tabs', props: ['scroll-padding-left', 'scroll-padding-right'], golden: 'a focused tab is brought in whole: a browser\'s focus scrolls only a tab it judges out of view, and left one 3px in view where it was (reference/console/)' },
    ],
  },
  {
    name: 'tabs many wide',
    html: ex('many'),
    widths: [1200],
    checks: [
      { target: '.cui-console-tabs', props: ['flex-wrap', 'overflow-x'], golden: 'wide, the bar as the toolkit draws it' },
    ],
  },
  {
    name: 'tabs not current',
    html: '<div class="cui-console-tabs" aria-label="Report"><a class="cui-console-tab" href="#crashes" aria-current="false">Crashes</a></div>',
    widths: [1200],
    checks: [
      { target: '[aria-current="false"]', props: ['border-bottom-color'], golden: 'aria-current="false" not the current one: no rule' },
    ],
  },
  {
    name: 'tabs of radios',
    html: ex('radios'),
    widths: [1200],
    checks: [
      { target: PLAIN, props: SIZE, tb: size('md') },
      { target: PLAIN, props: ['background-color', 'padding-left', 'padding-right', 'border-top-left-radius'], tb: [tab('filledGray', 'default', 'false', 'backgroundColor'), `${T}-defaultMdPaddingInline`, `${T}-defaultMdPaddingInline`, `${T}-variant-controlShape-capsule-controlSize-md-borderRadius`] },
      { target: PLAIN, props: ['color'], tb: [tab('filledGray', 'default', 'false', 'color')] },
      { target: PLAIN, props: ['background-color'], state: 'hover', tb: [tab('filledGray', 'hover', 'false', 'backgroundColor')] },
      { target: PLAIN, props: ['color'], state: 'hover', tb: [tab('filledGray', 'hover', 'false', 'color')] },
      { target: PLAIN, props: ['background-color'], state: 'active', tb: [tab('filledGray', 'press', 'false', 'backgroundColor')] },
      { target: PLAIN, props: ['color'], state: 'active', tb: [tab('filledGray', 'press', 'false', 'color')] },
      { target: '.cui-console-tabs', props: ['column-gap'], tb: [`${T}-variant-controlSize-md-variant-filledGray-gap`] },
      { target: '.cui-console-tabs', props: ['border-bottom-width'], golden: 'no hairline under pills' },
      { target: PLAIN, props: ['border-bottom-width'], golden: 'no underline under a pill: an underlined tab\'s transparent 2px rule lifted its label 1px off its fill\'s middle' },
      { target: PICKED, props: ['background-color', 'color'], tb: [tab('filledGray', 'default', 'true', 'backgroundColor'), tab('filledGray', 'default', 'true', 'color')] },
      { target: PICKED, props: ['background-color', 'color'], state: 'hover', tb: [tab('filledGray', 'hover', 'true', 'backgroundColor'), tab('filledGray', 'hover', 'true', 'color')] },
      { target: PICKED, props: ['background-color', 'color'], state: 'active', tb: [tab('filledGray', 'press', 'true', 'backgroundColor'), tab('filledGray', 'press', 'true', 'color')] },
      { target: `${RADIO} > input`, props: ['position', 'opacity', 'pointer-events'], golden: 'the page\'s own radio out of the way: the pill is the control\'s look, its words the radio\'s label' },
      { target: PICKED, props: ['outline-style', 'outline-width'], on: `${RADIO} > input`, state: 'focus-visible', golden: 'the keyboard\'s ring on the pill, not around the hidden radio (the one ring, moved)' },
    ],
  },
  {
    name: 'tabs of radios forced colors',
    html: ex('radios'),
    widths: [1200],
    forcedColors: true,
    checks: [
      { target: PICKED, props: ['background-color', 'color'], golden: 'the chosen one the theme\'s Highlight under HighlightText: forced colors drop the fill, and it read as the others' },
      { target: PLAIN, props: ['background-color', 'color'], golden: 'the others\' gray fill forced to the ground, at its alpha, under the forced text: not Highlight' },
    ],
  },
  {
    name: 'tabs touch',
    html: ex('radios'),
    touch: true,
    widths: [390],
    checks: [
      { target: '.cui-console-tab', props: ['justify-content'], golden: 'its label centered in the 44px tab' },
      { target: PLAIN, props: ['box-height', 'min-width', 'border-top-left-radius'], golden: 'a tab 44px each way at least on a touch screen, the HIG\'s 44pt (Accessibility, Mobility), a pill\'s corners half its height: a target drawn round it was clipped by a phone\'s sideways-scrolling row' },
    ],
  },
  {
    name: 'tabs underlined touch',
    html: ex('many'),
    touch: true,
    widths: [390],
    checks: [
      { target: REST, props: ['box-height', 'min-width'], golden: 'an underlined tab 44px each way at least on a touch screen, in its scrolling row' },
    ],
  },
  {
    name: 'tabs forced colors',
    html: ex('tabs'),
    widths: [1200],
    forcedColors: true,
    checks: [
      { target: '.cui-console-tab[aria-current="page"]', props: ['border-bottom-color', 'border-bottom-width'], golden: 'the current tab\'s rule, in the forced text color' },
      { target: REST, props: ['border-bottom-color', 'border-bottom-width'], golden: 'the others\' rule in the ground\'s color, its 2px kept: forced colors drew even the transparent rule in the text\'s color, every tab ruled as the current one' },
    ],
  },
  {
    name: 'tabs dark',
    html: `<div class="cui-dark">${ex('tabs')}${ex('radios')}</div>`,
    widths: [1200],
    checks: [
      { target: REST, props: ['color'], golden: 'the dark side derived (tokens.css, decision 18)' },
      { target: PLAIN, props: ['background-color', 'color'], golden: 'the dark side derived (tokens.css, decision 18)' },
      { target: PICKED, props: ['background-color', 'color'], golden: 'the dark side derived (tokens.css, decision 18): the chosen pill white under black' },
    ],
  },
];
