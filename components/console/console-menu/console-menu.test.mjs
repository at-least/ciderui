import { examples } from '../../../test/harness.mjs';

const ex = examples(import.meta.url);
const DESKTOP = [1600, 1200, 640];
const PHONE = [639, 390];
const ITEM = ['box-height', 'background-color', 'border-top-left-radius', 'padding-left', 'padding-right'];
const LABEL = ['font-size', 'line-height', 'font-weight', 'letter-spacing', 'color'];
const MENU = 'component-DisplayList-variant-pageScale';
const FITTED = 'the secondary label fitted to 0.76 (tokens.css, decision 21): the toolkit\'s 0.60 is 3.4:1 on white';
const CURRENT = '.cui-console-menu > [aria-current="page"]';
const SUB = '.cui-console-menu-group[open] > a[href]';
const PLAIN = '.cui-console-menu > a[href="#shop"]';
const OPEN = '.cui-console-btn[popovertarget]';
const CHECKED = '[role="menuitemradio"][aria-checked="true"]';
const UNCHECKED = '[role="menuitemradio"][aria-checked="false"]';
const page = (html) => `<div style="padding:24px 24px 320px">${html}</div>`;

export const cases = [
  {
    name: 'menu long label',
    html: '<div class="w-48"><nav class="cui-console-menu" aria-label="Analytics"><a class="cui-console-menu-item" href="#a"><span class="cui-console-menu-label">App Store Product Page Optimization</span></a></nav></div>',
    widths: [1200],
    checks: [
      { target: '.cui-console-menu-label', props: ['box-height'], golden: 'a long label on two lines, whole, not cut with an ellipsis' },
    ],
  },
  {
    name: 'menu choices',
    html: page(ex('choices')),
    widths: [1200],
    click: OPEN,
    checks: [
      { target: '#subtitles', props: ['display'], golden: 'its button opens it, a console popover (decision 27)' },
      { target: UNCHECKED, props: ['box-height', 'border-top-left-radius', 'padding-left', 'padding-right'], asc: ['hover-menu', 'menu.item'] },
      { target: `${UNCHECKED} .cui-console-menu-label`, props: LABEL, asc: ['analytics', 'menu.item.label'] },
      { target: UNCHECKED, props: ['background-color', 'border-top-width', 'text-align', 'font-family'], golden: 'a <button> drawn as the menu\'s <a>: no ground, no edge, its label at its start in the page\'s face' },
      { target: CHECKED, props: ['background-color'], golden: 'the chosen one unfilled, its check alone marking it, as tv.apple.com\'s chosen subtitles (tv-player-menu:chosen, no fill): a fill is the current page\'s' },
      { target: UNCHECKED, props: ['background-color'], state: 'hover', asc: ['hover-menu', 'menu.item'] },
      { target: UNCHECKED, props: ['box-width'], golden: 'a row the menu\'s width, as an <a> item is' },
      { target: `${CHECKED} .cui-console-menu-icon`, props: ['box-width', 'margin-left'], tb: ['component-DisplayList-anatomy-Item-anatomy-Indicator-variant-inlinePosition-start-width', 'component-DisplayList-anatomy-Item-anatomy-Indicator-variant-inlinePosition-start-marginInlineStart'] },
      { target: `${CHECKED} .cui-console-menu-icon`, props: ['visibility', 'color'], golden: `the check in the toolkit's indicator's place at its start, where tv.apple.com draws its own (tv-player-menu:chosen.check, 10px in); Lucide's check (decision 17); ${FITTED}` },
      { target: `${UNCHECKED} .cui-console-menu-icon`, props: ['visibility', 'box-width'], golden: 'unchecked, its check hidden in its place: every label lines up, as tv.apple.com\'s (its text 36px in on every row)' },
      { target: `${UNCHECKED} .cui-console-menu-label`, props: ['box-left'], golden: 'each label at the same start' },
      { target: `${CHECKED} .cui-console-menu-label`, props: ['box-left'], golden: 'each label at the same start' },
      { target: '#subtitles .cui-console-menu-heading', props: ['font-size', 'line-height', 'font-weight', 'padding-top', 'padding-left'], tb: [`${MENU}-desktop-variant-menu-anatomy-Heading-textStyle-fontSize`, `${MENU}-desktop-variant-menu-anatomy-Heading-textStyle-lineHeight`, `${MENU}-desktop-variant-menu-anatomy-Heading-textStyle-fontWeight`, `${MENU}-desktop-variant-menu-anatomy-Header-paddingBlockStart`, `${MENU}-desktop-variant-menu-anatomy-Header-paddingInline`] },
    ],
  },
  {
    name: 'menu choices columns',
    html: page(ex('columns')),
    widths: [1200],
    click: OPEN,
    checks: [
      { target: '#subtitles .cui-console-menu-columns', props: ['display', 'align-items'], golden: 'the columns side by side, the Apple TV app\'s Audio beside its Subtitles (the owner, 2026-11-05, for muse\'s player: one button, one panel), each as tall as the menu so the rule runs its height' },
      { target: '#subtitles [data-column="voices"]', props: ['box-x', 'box-width'], golden: 'the first column at the popover\'s 16px side, as wide as its choices ask' },
      { target: '#subtitles [data-column="subtitles"]', props: ['box-x', 'border-left-width', 'border-left-style', 'border-left-color', 'margin-left', 'padding-left'], golden: 'the next after a rule, the menu\'s own 0.5px in its border color (cui-console-menu-rule), in the middle of the popover\'s 8px gap (cui-console-popover)' },
      { target: '#subtitles', props: ['box-width', 'max-width'], golden: 'the popover as wide as its columns ask, past the one column\'s 350px' },
      { target: `[data-column="subtitles"] ${UNCHECKED} .cui-console-menu-label`, props: LABEL, asc: ['analytics', 'menu.item.label'] },
      { target: `[data-column="voices"] ${CHECKED} .cui-console-menu-icon`, props: ['visibility'], golden: 'each column its own check' },
    ],
  },
  {
    name: 'menu choices columns phone',
    html: page(ex('columns')),
    widths: [390],
    click: OPEN,
    checks: [
      { target: '#subtitles', props: ['box-width'], golden: 'no wider than 85% of a phone\'s window, the popover\'s own cap' },
      { target: '#subtitles [data-column="subtitles"]', props: ['box-right'], golden: 'both columns inside it' },
    ],
  },
  {
    name: 'menu choices chosen',
    html: page(ex('choices')),
    widths: [1200],
    click: [OPEN, '[role="menuitemradio"][value="ja"]'],
    checks: [
      { target: '#subtitles', props: ['display'], golden: 'a choice closes its popover with no script (popovertargetaction="hide"); its check is the page\'s script\'s to move' },
    ],
  },
  {
    name: 'menu choices phone',
    html: page(ex('choices')),
    widths: [390],
    click: OPEN,
    checks: [
      { target: UNCHECKED, props: ['padding-top', 'padding-left', 'border-top-left-radius', 'column-gap'], tb: [`${MENU}-mobile-variant-menu-anatomy-Item-paddingBlock`, `${MENU}-mobile-variant-menu-anatomy-Item-paddingInline`, `${MENU}-mobile-variant-menu-anatomy-Item-borderRadius`, `${MENU}-mobile-variant-menu-anatomy-Item-gap`] },
    ],
  },
  {
    name: 'menu choices dark',
    html: `<div class="cui-dark">${page(ex('choices'))}</div>`,
    widths: [1200],
    click: OPEN,
    checks: [
      { target: '#subtitles', props: ['background-color', 'color'], golden: 'dark over a dark page: the console popover\'s derived dark (decision 18), a menu over the video in muse\'s dark console' },
      { target: `${CHECKED} .cui-console-menu-icon`, props: ['color'], golden: 'the check in the dark secondary label' },
    ],
  },
  {
    name: 'menu',
    html: ex('menu'),
    widths: DESKTOP,
    checks: [
      { target: CURRENT, props: ITEM, asc: ['analytics', 'menu.current'] },
      { target: `${CURRENT} .cui-console-menu-label`, props: LABEL, asc: ['analytics', 'menu.current.label'] },
      { target: SUB, props: ITEM.filter(p => p !== 'background-color'), asc: ['analytics', 'menu.item'] },
      { target: SUB, props: ['background-color'], golden: 'no ground at rest: transparent, where the toolkit writes its transparent white, oklch(1 0 0 / 0), the same nothing' },
      { target: `${SUB} .cui-console-menu-label`, props: LABEL, asc: ['analytics', 'menu.item.label'] },
      { target: '.cui-console-menu-group[open] > summary', props: ['box-height', 'border-top-left-radius', 'padding-left', 'padding-right', 'cursor'], asc: ['analytics-page', 'group.button'] },
      { target: '.cui-console-menu-group[open] > summary .cui-console-menu-label', props: LABEL, asc: ['analytics-page', 'group.button.label'] },
      { target: '[aria-disabled="true"]', props: ['box-height', 'padding-left', 'cursor'], asc: ['analytics-page', 'item.disabled'] },
      { target: '[aria-disabled="true"] .cui-console-menu-label', props: LABEL, asc: ['analytics-page', 'item.disabled.label'] },
      { target: '.cui-console-menu-rule', props: ['box-height', 'background-color', 'margin-top', 'margin-bottom', 'margin-left', 'margin-right'], asc: ['analytics', 'menu.rule'] },
      { target: PLAIN, props: ['background-color'], state: 'hover', asc: ['hover-menu', 'menu.item'] },
      { target: `${PLAIN} .cui-console-menu-label`, props: ['color', 'font-weight'], state: 'hover', on: PLAIN, asc: ['hover-menu', 'menu.item.label'] },
      { target: '.cui-console-menu-heading', props: ['font-size', 'line-height', 'font-weight', 'padding-top', 'padding-bottom', 'padding-left'], tb: [`${MENU}-desktop-variant-menu-anatomy-Heading-textStyle-fontSize`, `${MENU}-desktop-variant-menu-anatomy-Heading-textStyle-lineHeight`, `${MENU}-desktop-variant-menu-anatomy-Heading-textStyle-fontWeight`, `${MENU}-desktop-variant-menu-anatomy-Header-paddingBlockStart`, `${MENU}-desktop-variant-menu-anatomy-Header-paddingBlockEnd`, `${MENU}-desktop-variant-menu-anatomy-Header-paddingInline`] },
      { target: CURRENT, props: ['column-gap'], tb: [`${MENU}-desktop-variant-menu-anatomy-Item-gap`] },
      { target: `${CURRENT} .cui-console-menu-icon`, props: ['box-width', 'margin-left'], tb: ['component-DisplayList-variant-variant-menu-anatomy-artStackWithSFSymbol-width', 'component-DisplayList-variant-variant-menu-anatomy-ArtStack-anatomy-SFSymbol-marginInlineOffset'] },
      { target: `${CURRENT} .cui-console-menu-icon`, props: ['color'], golden: FITTED },
      { target: '.cui-console-menu-heading', props: ['color'], golden: FITTED },
      { target: '.cui-console-menu-secondary', props: ['color'], golden: FITTED },
      { target: '.cui-console-menu-chevron', props: ['color'], golden: FITTED },
      { target: '.cui-console-menu-group:not([open]) .cui-console-menu-chevron', props: ['rotate'], golden: 'closed, the chevron points along the row: Lucide\'s chevron-right, where the toolkit draws an SF Symbol' },
      { target: '.cui-console-menu-group[open] .cui-console-menu-chevron', props: ['rotate'], golden: 'open, it turns down, as the toolkit\'s does' },
      { target: PLAIN, props: ['background-color'], state: 'active', golden: 'pressed as hovered: the toolkit\'s pressed fill, oklch(55.4% 0.012 286 / 0.08), is no step of its palette, so it has no dark side to derive' },
      { target: '.cui-console-menu', props: ['display'], asc: ['analytics-page', 'menu'] },
      { target: '.cui-console-menu', props: ['flex-direction'], golden: 'a column of its items' },
      { target: `${SUB} .cui-console-menu-label`, props: ['white-space', 'text-overflow', 'overflow-wrap'], golden: 'a long label wraps, whole: an ellipsis left two pages that begin alike unreadable apart' },
      { target: '.cui-console-menu-secondary', props: ['box-right'], golden: 'at the item\'s end' },
      { target: '.cui-console-menu-chevron', props: ['box-right'], golden: 'at the group button\'s end' },
      { target: '[aria-disabled="true"]', props: ['background-color'], state: 'hover', golden: 'a disabled item no fill hovered' },
      { target: '.cui-console-menu-secondary', props: ['font-size', 'line-height', 'font-weight'], tb: [`${MENU}-desktop-variant-menu-anatomy-Secondary-textStyle-fontSize`, `${MENU}-desktop-variant-menu-anatomy-Secondary-textStyle-lineHeight`, `${MENU}-desktop-variant-menu-anatomy-Secondary-textStyle-fontWeight`] },
    ],
  },
  {
    name: 'menu phone',
    html: ex('menu'),
    widths: PHONE,
    checks: [
      { target: CURRENT, props: ['padding-top', 'padding-left', 'border-top-left-radius', 'column-gap'], tb: [`${MENU}-mobile-variant-menu-anatomy-Item-paddingBlock`, `${MENU}-mobile-variant-menu-anatomy-Item-paddingInline`, `${MENU}-mobile-variant-menu-anatomy-Item-borderRadius`, `${MENU}-mobile-variant-menu-anatomy-Item-gap`] },
      { target: `${CURRENT} .cui-console-menu-label`, props: ['font-size', 'line-height', 'font-weight'], tb: [`${MENU}-mobile-variant-menu-anatomy-Label-textStyle-fontSize`, `${MENU}-mobile-variant-menu-anatomy-Label-textStyle-lineHeight`, `${MENU}-mobile-variant-menu-anatomy-Label-textStyle-fontWeight`] },
      { target: '.cui-console-menu-heading', props: ['font-size', 'line-height', 'font-weight', 'padding-top', 'padding-bottom', 'padding-left'], tb: [`${MENU}-mobile-variant-menu-anatomy-Heading-textStyle-fontSize`, `${MENU}-mobile-variant-menu-anatomy-Heading-textStyle-lineHeight`, `${MENU}-mobile-variant-menu-anatomy-Heading-textStyle-fontWeight`, `${MENU}-mobile-variant-menu-anatomy-Header-paddingBlockStart`, `${MENU}-mobile-variant-menu-anatomy-Header-paddingBlockEnd`, `${MENU}-mobile-variant-menu-anatomy-Header-paddingInline`] },
      { target: '.cui-console-menu-rule', props: ['margin-left', 'margin-right'], tb: [`${MENU}-mobile-variant-menu-anatomy-HR-marginInline`, `${MENU}-mobile-variant-menu-anatomy-HR-marginInline`] },
      { target: CURRENT, props: ['box-height'], golden: 'the toolkit\'s mobile item, its padding round its line' },
      { target: SUB, props: ['padding-left'], golden: 'a group\'s items 48px in on a phone too: the toolkit names no mobile value, the analytics page\'s is read at 1920px' },
    ],
  },
  {
    name: 'menu touch',
    html: ex('menu'),
    touch: true,
    widths: [390],
    checks: [
      { target: PLAIN, props: ['box-height', 'border-top-left-radius'], golden: 'an item 44px tall on a touch screen, the HIG\'s 44pt (Accessibility, Mobility), a pill\'s corners half of it' },
    ],
  },
  {
    name: 'menu forced colors',
    html: ex('menu'),
    widths: [1200],
    forcedColors: true,
    checks: [
      { target: CURRENT, props: ['background-color', 'color'], golden: 'the current page the theme\'s Highlight under HighlightText: forced colors drop the fill, and it read as the others' },
      { target: PLAIN, props: ['background-color'], golden: 'another item no ground' },
      { target: '.cui-console-menu-rule', props: ['background-color'], golden: 'the rule CanvasText: forced colors drew its fill in the ground\'s color, at its alpha' },
    ],
  },
  {
    name: 'menu dark',
    html: `<div class="cui-dark">${ex('menu')}</div>`,
    widths: [1200],
    checks: [
      { target: CURRENT, props: ['background-color', 'color'], golden: 'the dark side derived (tokens.css, decision 18)' },
      { target: '[aria-disabled="true"]', props: ['color'], golden: 'the dark side derived (tokens.css, decision 18)' },
      { target: '.cui-console-menu-rule', props: ['background-color'], golden: 'the dark side derived (tokens.css, decision 18)' },
    ],
  },
];
