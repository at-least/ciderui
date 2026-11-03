import { OUTLINE, examples } from '../../../test/harness.mjs';

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
const RADIO = 'label.cui-console-tab:not(:has(input:checked))';
const CHECKED = 'label.cui-console-tab:has(input:checked)';

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
      { target: '.cui-console-tabs', props: ['scrollbar-width'], golden: 'no scrollbar: the side still hiding content fades instead' },
      { target: '.cui-console-tab[href="#settings"]', props: ['outline-offset'], state: 'focus-visible', golden: 'its focus ring inside the tab, where the scroll box would clip it' },
      { target: '.cui-console-tabs', props: ['scroll-padding-left', 'scroll-padding-right'], golden: 'a focused tab is brought in clear of the 24px fade: a browser\'s focus scrolls only a tab it judges out of view, and left one 3px in view where it was' },
    ],
  },
  ...[
    ['start', 'data-overflow data-overflow-start', [375], 'scrolled to its start, the right edge fades over the content still hidden; the left edge stays clear'],
    ['end', 'data-overflow data-overflow-end', [375], 'scrolled to its end, the left edge fades over the sections clipped behind; the right edge stays clear'],
    ['middle and desktop', 'data-overflow', [375, 1200], 'between the ends, both edges fade; the fade is the phone row\'s — wide, none'],
  ].map(([where, states, widths, golden]) => ({
    name: `tabs overflow fade ${where}`,
    html: ex('many').replace('cui-console-tabs"', `cui-console-tabs" ${states}`),
    widths,
    checks: [{ target: '.cui-console-tabs', props: ['mask-image'], golden }],
  })),
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
    name: 'tabs selected',
    html: '<div class="cui-console-tabs" role="tablist" aria-label="Report"><button class="cui-console-tab" role="tab" type="button" aria-selected="true">Crashes</button><button class="cui-console-tab" role="tab" type="button" aria-selected="false" aria-current="false">Hangs</button></div>',
    widths: [1200],
    checks: [
      { target: '[aria-selected="true"]', props: ['border-bottom-color'], golden: 'aria-selected="true" the current one, on its rule, as aria-current' },
      { target: '[aria-current="false"]', props: ['border-bottom-color'], golden: 'aria-current="false" not the current one: no rule' },
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
    name: 'tabs of radios',
    html: ex('appearance'),
    widths: [1200, 390],
    checks: [
      { target: RADIO, props: SIZE, tb: size('md') },
      { target: RADIO, props: ['background-color', 'padding-left', 'padding-right', 'border-top-left-radius'], tb: [tab('filledGray', 'default', 'false', 'backgroundColor'), `${T}-defaultMdPaddingInline`, `${T}-defaultMdPaddingInline`, `${T}-variant-controlShape-capsule-controlSize-md-borderRadius`] },
      { target: CHECKED, props: ['background-color', 'color'], tb: [tab('filledGray', 'default', 'true', 'backgroundColor'), tab('filledGray', 'default', 'true', 'color')] },
      { target: RADIO, props: ['color', 'cursor'], golden: 'a radio\'s tab as a link\'s: the blue label on the gray pill, the pointer' },
      { target: 'fieldset.cui-console-tabs', props: ['border-top-width', 'border-left-width', 'padding-left', 'padding-top', 'margin-left', 'column-gap'], golden: 'a <fieldset> drawn as the row: no edge, no room, the filled gap' },
      { target: 'fieldset.cui-console-tabs > legend', props: ['box-width', 'position'], golden: 'the legend read, not shown' },
      { target: `${RADIO} > input`, props: ['opacity', 'box-width'], golden: 'the radio read, not shown' },
      { target: CHECKED, props: ['outline-style'], golden: 'no ring at rest' },
    ],
  },
  {
    name: 'tabs of radios focused',
    html: ex('appearance'),
    widths: [1200],
    press: ['Tab'],
    checks: [
      { target: CHECKED, props: OUTLINE, golden: 'the radio hidden, the one focus ring around its tab' },
    ],
  },
  {
    name: 'tabs of radios touch',
    html: ex('appearance'),
    touch: true,
    widths: [390],
    checks: [
      { target: RADIO, props: ['box-height', 'min-width'], golden: 'a radio\'s tab 44px each way at least on a touch screen, as a link\'s' },
    ],
  },
  {
    name: 'tabs of radios forced colors',
    html: ex('appearance'),
    widths: [1200],
    forcedColors: true,
    checks: [
      { target: CHECKED, props: ['background-color', 'color'], golden: 'the checked tab the theme\'s Highlight under HighlightText, as a chosen link\'s' },
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
      { target: PILL, props: ['box-height', 'min-width', 'border-top-left-radius'], golden: 'a tab 44px each way at least on a touch screen, the HIG\'s 44pt (Accessibility, Mobility), a pill\'s corners half its height: a target drawn round it was clipped by a phone\'s sideways-scrolling row' },
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
    name: 'filled forced colors',
    html: ex('filled'),
    widths: [1200],
    forcedColors: true,
    checks: [
      { target: CHOSEN, props: ['background-color', 'color'], golden: 'the chosen pill the theme\'s Highlight under HighlightText: forced colors drop the fill, and it read as the others' },
      { target: PILL, props: ['background-color'], golden: 'the others\' gray fill forced to the ground, at its alpha' },
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
