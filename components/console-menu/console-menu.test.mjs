// The console's menu against App Store Connect's analytics menu: an item, the
// current one, an item in a group, a group's button and chevron, a disabled
// item, the rule, and an item hovered as its page drew them
// (analytics:menu.*, analytics-page:group.*, item.disabled.*,
// hover-menu:menu.item); the heading and the mobile page scale the
// toolkit's tokens (--tb-component-DisplayList-*menu*). Golden, with their
// why: the secondary label's fitted alpha, the pressed fill, the dark side.

import { examples } from '../../test/harness.mjs';

const ex = examples(import.meta.url);
const DESKTOP = [1600, 1200, 640];
const PHONE = [639, 390];
const ITEM = ['box-height', 'background-color', 'border-top-left-radius', 'padding-left', 'padding-right'];
const LABEL = ['font-size', 'line-height', 'font-weight', 'letter-spacing', 'color'];
const MENU = 'component-DisplayList-variant-pageScale';
const FITTED = 'the secondary label fitted to 0.74 (tokens.css, decision 21): the toolkit\'s 0.60 is 3.4:1 on white';
const CURRENT = '.cui-console-menu > [aria-current="page"]';
const SUB = '.cui-console-menu-group[open] > a[href]';
const PLAIN = '.cui-console-menu > a[href="#shop"]';

export const cases = [
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
      { target: SUB, props: ['padding-left'], golden: 'a group\'s items 48px in on a phone too: the toolkit names no mobile value, the analytics page\'s is read at 1920px' },
    ],
  },
  {
    name: 'menu touch',
    html: ex('menu'),
    touch: true,
    widths: [390],
    checks: [
      { target: PLAIN, props: ['background-color'], state: 'hover', rest: 'a tapped item keeps no fill (rule 5)' },
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
