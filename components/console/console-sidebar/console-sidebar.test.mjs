import { examples } from '../../../test/harness.mjs';

const ex = examples(import.meta.url);
const S = 'component-Sidebar';
const D = 'component-DisplayList-variant';
const DS = `${D}-pageScale-desktop-variant-sidebar`;
const MS = `${D}-pageScale-mobile-variant-sidebar`;
const SIDEBAR = '.cui-console-sidebar';
const ITEM = '.cui-console-sidebar .cui-console-menu-item[href="#projects"]';
const CURRENT = '.cui-console-sidebar .cui-console-menu-item[aria-current="page"]';
const FITTED = 'the secondary label fitted to 0.76 (tokens.css, decision 21)';
const PAGE = (html) => `<div class="cui-console-page">${html}</div>`;

export const cases = [
  {
    name: 'sidebar',
    html: PAGE(ex('sidebar')),
    widths: [1600, 1200, 1024],
    checks: [
      { target: '.cui-console-sidebar-account', state: 'hover', props: ['background-color'], golden: 'the pointer finds the account row as it finds a menu row' },
      { target: SIDEBAR, props: ['bottom'], golden: 'in the page from 1024px, not the sheet\'s full height' },
      { target: SIDEBAR, props: ['box-width'], tb: ['size-sidebar-width'] },
      { target: SIDEBAR, props: ['color'], tb: ['color-role-neutral-label-primary'] },
      { target: SIDEBAR, props: ['flex-direction'], golden: 'a column: its header, list and account one under another' },
      { target: SIDEBAR, props: ['display', 'position', 'background-color', 'box-shadow'], golden: 'from 1024px it stands in the page: the toolkit\'s plain sidebar, no ground, no shadow' },
      { target: '.cui-console-sidebar-header', props: ['padding-top', 'padding-bottom', 'padding-left', 'padding-right', 'column-gap', 'min-height'], tb: [`${S}-anatomy-Header-paddingBlockStart`, `${S}-anatomy-Header-paddingBlockEnd`, `${S}-anatomy-Header-paddingInline`, `${S}-anatomy-Header-paddingInline`, `${S}-anatomy-Header-anatomy-Nav-gap`, `${S}-anatomy-Header-anatomy-Nav-minHeight`] },
      { target: '.cui-console-sidebar-wordmark', props: ['margin-left'], tb: [`${S}-anatomy-Header-anatomy-Nav-anatomy-Wordmark-marginInlineStart`] },
      { target: '.cui-console-sidebar-wordmark', props: ['font-size', 'line-height', 'font-weight'], golden: 'the wordmark the product\'s name, 17px on 22px at 600: the toolkit\'s is a picture' },
      { target: '.cui-console-sidebar-close', props: ['display'], golden: 'from 1024px no close button: it stands in the page' },
      { target: '.cui-console-sidebar-body', props: ['padding-left', 'padding-right'], tb: [`${S}-anatomy-Body-paddingInline`, `${S}-anatomy-Body-paddingInline`] },
      { target: ITEM, props: ['min-height', 'padding-top', 'padding-left', 'border-top-left-radius', 'column-gap', 'font-size', 'line-height', 'font-weight'], tb: [`${DS}-anatomy-Item-minHeight`, `${DS}-anatomy-Item-paddingBlock`, `${DS}-anatomy-Item-paddingInline`, `${DS}-anatomy-Item-borderRadius`, `${DS}-anatomy-Item-gap`, `${DS}-anatomy-Item-variant-selected-none-anatomy-Label-textStyle-fontSize`, `${DS}-anatomy-Item-variant-selected-none-anatomy-Label-textStyle-lineHeight`, `${DS}-anatomy-Item-variant-selected-none-anatomy-Label-textStyle-fontWeight`] },
      { target: ITEM, props: ['color'], tb: [`${D.replace('-variant', '')}-anatomy-Item-variant-colorRole-neutral-effect-default-anatomy-Label-color`] },
      { target: ITEM, props: ['background-color'], state: 'hover', tb: [`${D}-variant-sidebar-anatomy-Item-variant-effect-hover-focus-off-selected-none-backgroundColor`] },
      { target: ITEM, props: ['background-color'], state: 'active', golden: 'pressed as hovered: the toolkit\'s pressed fill, oklch(42.6% 0.008 286 / 0.03), is no step of its palette' },
      { target: CURRENT, props: ['background-color', 'font-weight'], tb: [`${D}-variant-sidebar-anatomy-Item-variant-effect-default-focus-off-selected-fill-backgroundColor`, `${DS}-anatomy-Item-variant-selected-fill-anatomy-Label-textStyle-fontWeight`] },
      { target: '.cui-console-sidebar .cui-console-menu-heading', props: ['font-size', 'line-height', 'font-weight', 'padding-top', 'padding-bottom', 'padding-left', 'margin-top'], tb: [`${DS}-anatomy-Heading-textStyle-fontSize`, `${DS}-anatomy-Heading-textStyle-lineHeight`, `${DS}-anatomy-Heading-textStyle-fontWeight`, `${DS}-anatomy-Header-paddingBlockStart`, `${DS}-anatomy-Header-paddingBlockEnd`, `${DS}-anatomy-Header-paddingInline`, `${D}-variant-sidebar-anatomy-groupGap-marginBlockEnd`] },
      { target: '.cui-console-sidebar .cui-console-menu-heading', props: ['color'], golden: FITTED },
      { target: `${ITEM} .cui-console-menu-label`, props: ['white-space', 'text-overflow', 'overflow-wrap'], golden: 'a long label wraps, whole, as a menu\'s' },
      { target: '.cui-console-sidebar .cui-console-menu-rule', props: ['margin-top', 'margin-left', 'background-color', 'box-height'], tb: [`${D}-variant-sidebar-anatomy-HR-marginBlock`, `${DS}-anatomy-HR-marginInline`, 'color-role-neutral-border-secondary', 'border-width-fine'] },
      { target: '.cui-console-sidebar .cui-console-menu-icon', props: ['box-width', 'box-height'], golden: 'its icon Lucide\'s, 20px in 24px (the toolkit\'s art\'s least width): the toolkit draws SF Symbols at its text\'s size' },
      { target: '.cui-console-sidebar-footer', props: ['padding-top', 'padding-bottom'], tb: [`${S}-anatomy-UniversalIDFooter-paddingBlockStart`, `${S}-anatomy-UniversalIDFooter-paddingBlockEnd`] },
      { target: '.cui-console-sidebar-account', props: ['min-height', 'border-top-left-radius', 'column-gap', 'margin-left', 'padding-left', 'padding-right', 'font-size', 'line-height', 'font-weight', 'color'], tb: ['minHeight', 'borderRadius', 'gap', 'marginInline', 'paddingInlineStart', 'paddingInlineEnd'].map((p) => `${S}-anatomy-UniversalIDFooter-anatomy-ID-${p}`).concat(['fontSize', 'lineHeight', 'fontWeight'].map((p) => `${S}-anatomy-UniversalIDFooter-anatomy-ID-anatomy-Name-textStyle-${p}`), [`${S}-anatomy-UniversalIDFooter-anatomy-ID-anatomy-Name-color`]) },
      { target: '.cui-console-sidebar-avatar', props: ['box-height', 'color'], tb: [`${S}-anatomy-UniversalIDFooter-anatomy-ID-anatomy-Art-fontSize`, `${S}-anatomy-UniversalIDFooter-anatomy-ID-anatomy-Art-color`] },
    ],
  },
  {
    name: 'sidebar sheet tall',
    html: PAGE(ex('open')),
    widths: [390],
    height: 1400,
    click: '.cui-console-sidebar-open',
    checks: [
      { target: '.cui-console-sidebar-footer', props: ['box-top', 'box-height'], golden: 'the account at the sheet\'s foot, 11px above the window\'s: the body takes the rest of the column' },
    ],
  },
  {
    name: 'sidebar sheet',
    html: PAGE(ex('open')),
    widths: [390],
    click: '.cui-console-sidebar-open',
    checks: [
      { target: '.cui-console-menu-heading', props: ['padding-left', 'font-size', 'line-height'], golden: 'in the sidebar, a heading lined up with its rows, in their size' },
      { target: SIDEBAR, props: ['box-left', 'box-top', 'border-top-left-radius', 'background-color', 'box-shadow'], tb: [`${S}-inset`, `${S}-inset`, `${S}-borderRadius`, `${S}-variant-variant-opaque-backgroundColor`, `${S}-variant-variant-opaque-boxShadow`] },
      { target: SIDEBAR, props: ['position', 'display'], golden: 'below 1024px, opened, a sheet over the page: a popover, fixed' },
      { target: '.cui-console-sidebar-close', props: ['display'], golden: 'its close button shown in the sheet' },
      { target: '.cui-console-sidebar-close', props: ['box-right'], golden: 'at the header\'s end, 6px in' },
      { target: ITEM, props: ['min-height', 'padding-top', 'padding-left', 'border-top-left-radius', 'font-size', 'line-height'], tb: [`${MS}-anatomy-Item-minHeight`, `${MS}-anatomy-Item-paddingBlock`, `${MS}-anatomy-Item-paddingInline`, `${MS}-anatomy-Item-borderRadius`, `${MS}-anatomy-Item-variant-selected-none-anatomy-Label-textStyle-fontSize`, `${MS}-anatomy-Item-variant-selected-none-anatomy-Label-textStyle-lineHeight`] },
    ],
  },
  {
    name: 'sidebar opened focus',
    html: ex('sidebar').replace('<nav class="cui-console-sidebar"', '<button id="open" popovertarget="sidebar">Open</button><nav class="cui-console-sidebar"'),
    click: '#open',
    widths: [390],
    checks: [
      { target: '.cui-console-sidebar-close', props: ['focused'], golden: 'opened, the focus on its close button, as the global menu\'s (HTML\'s autofocus in a popover): a screen reader is in the sheet, and Escape or Return closes it' },
    ],
  },
  {
    name: 'sidebar load focus',
    html: ex('sidebar'),
    widths: [1200],
    checks: [
      { target: '.cui-console-sidebar-close', props: ['focused'], golden: 'from 1024px, the page loaded, the focus left alone: its close button stands in the page, not shown' },
    ],
  },
  {
    name: 'sidebar closed',
    html: PAGE(ex('open')),
    widths: [1023, 390],
    checks: [
      { target: SIDEBAR, props: ['display'], golden: 'below 1024px, closed, not shown' },
      { target: '.cui-console-sidebar-open', props: ['display'], golden: 'below 1024px, the page\'s button that opens it, shown' },
    ],
  },
  {
    name: 'sidebar open button wide',
    html: PAGE(ex('open')),
    widths: [1024],
    checks: [
      { target: '.cui-console-sidebar-open', props: ['display'], golden: 'from 1024px no button opens it: it stands in the page' },
    ],
  },
  {
    name: 'sidebar forced colors',
    html: PAGE(ex('sidebar')),
    widths: [1200],
    forcedColors: true,
    checks: [
      { target: CURRENT, props: ['background-color', 'color'], golden: 'the current page the theme\'s Highlight under HighlightText: forced colors drop the fill, and it read as the others' },
      { target: ITEM, props: ['background-color'], golden: 'another item no ground' },
      { target: '.cui-console-sidebar .cui-console-menu-rule', props: ['background-color'], golden: 'the rule CanvasText: forced colors drew its fill in the ground\'s color, at its alpha' },
    ],
  },
  {
    name: 'sidebar sheet forced colors',
    html: PAGE(ex('open')),
    widths: [390],
    click: '.cui-console-sidebar-open',
    forcedColors: true,
    checks: [
      { target: SIDEBAR, props: ['border-top-width', 'border-top-style', 'border-top-color', 'box-shadow'], golden: 'the sheet a 1px CanvasText edge, as the popover\'s and the toast\'s: forced colors drop its shadow, and its ground is the page\'s' },
    ],
  },
  {
    name: 'sidebar dark',
    html: `<div class="cui-dark">${PAGE(ex('sidebar'))}</div>`,
    widths: [1200],
    checks: [
      { target: CURRENT, props: ['background-color', 'color'], golden: 'the dark side derived (tokens.css, decision 18)' },
    ],
  },
  {
    name: 'wordmark touch',
    html: PAGE(ex('sidebar')),
    touch: true,
    widths: [1200],
    checks: [
      { target: '.cui-console-sidebar-wordmark', props: ['box-height'], golden: 'the wordmark drawn as before, its 22px line' },
      { target: '.cui-console-sidebar-wordmark', pseudo: '::after', props: ['height'], golden: 'its target a transparent ::after 44px tall at least on a touch screen, the HIG\'s 44pt (22px is under its 28pt minimum), as the console\'s link draws one' },
    ],
  },
];
