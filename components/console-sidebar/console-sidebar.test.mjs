// The console's sidebar against App Store Connect's toolkit's sidebar and
// its display list's sidebar variant (--tb-component-Sidebar-*,
// --tb-component-DisplayList-*sidebar*): its width, header, list's room,
// items at rest, hovered and current, headings, rules, the account at its
// foot; below 1024px a sheet, opened by the page's button, its inset,
// corners, ground and shadow. Golden, with their why: the wordmark, the
// icons, the pressed fill, the secondary label's alpha, the dark side.

import { examples } from '../../test/harness.mjs';

const ex = examples(import.meta.url);
const S = 'component-Sidebar';
const D = 'component-DisplayList-variant';
const DS = `${D}-pageScale-desktop-variant-sidebar`;
const MS = `${D}-pageScale-mobile-variant-sidebar`;
const SIDEBAR = '.cui-console-sidebar';
const ITEM = '.cui-console-sidebar-item[href="#projects"]';
const CURRENT = '.cui-console-sidebar-item[aria-current="page"]';
const PAGE = (html) => `<div class="cui-console-page">${html}</div>`;

export const cases = [
  {
    name: 'sidebar',
    html: PAGE(ex('sidebar')),
    widths: [1600, 1200, 1024],
    checks: [
      { target: SIDEBAR, props: ['box-width'], tb: ['size-sidebar-width'] },
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
      { target: '.cui-console-sidebar-heading', props: ['font-size', 'line-height', 'font-weight', 'padding-top', 'padding-bottom', 'padding-left', 'margin-top'], tb: [`${DS}-anatomy-Heading-textStyle-fontSize`, `${DS}-anatomy-Heading-textStyle-lineHeight`, `${DS}-anatomy-Heading-textStyle-fontWeight`, `${DS}-anatomy-Header-paddingBlockStart`, `${DS}-anatomy-Header-paddingBlockEnd`, `${DS}-anatomy-Header-paddingInline`, `${D}-variant-sidebar-anatomy-groupGap-marginBlockEnd`] },
      { target: '.cui-console-sidebar-heading', props: ['color'], golden: 'the secondary label fitted to 0.74 (tokens.css, decision 21)' },
      { target: '.cui-console-sidebar-secondary', props: ['font-size', 'line-height', 'font-weight'], tb: [`${DS}-anatomy-Secondary-textStyle-fontSize`, `${DS}-anatomy-Secondary-textStyle-lineHeight`, `${DS}-anatomy-Secondary-textStyle-fontWeight`] },
      { target: '.cui-console-sidebar-rule', props: ['margin-top', 'margin-left', 'background-color', 'box-height'], tb: [`${D}-variant-sidebar-anatomy-HR-marginBlock`, `${DS}-anatomy-HR-marginInline`, 'color-role-neutral-border-secondary', 'border-width-fine'] },
      { target: '.cui-console-sidebar-icon', props: ['box-width', 'box-height'], golden: 'its icon Lucide\'s, 20px in 24px (the toolkit\'s art\'s least width): the toolkit draws SF Symbols at its text\'s size' },
      { target: '.cui-console-sidebar-footer', props: ['padding-top', 'padding-bottom'], tb: [`${S}-anatomy-UniversalIDFooter-paddingBlockStart`, `${S}-anatomy-UniversalIDFooter-paddingBlockEnd`] },
      { target: '.cui-console-sidebar-account', props: ['min-height', 'border-top-left-radius', 'column-gap', 'margin-left', 'padding-left', 'padding-right', 'font-size', 'line-height', 'font-weight', 'color'], tb: ['minHeight', 'borderRadius', 'gap', 'marginInline', 'paddingInlineStart', 'paddingInlineEnd'].map((p) => `${S}-anatomy-UniversalIDFooter-anatomy-ID-${p}`).concat(['fontSize', 'lineHeight', 'fontWeight'].map((p) => `${S}-anatomy-UniversalIDFooter-anatomy-ID-anatomy-Name-textStyle-${p}`), [`${S}-anatomy-UniversalIDFooter-anatomy-ID-anatomy-Name-color`]) },
      { target: '.cui-console-sidebar-avatar', props: ['box-height', 'color'], tb: [`${S}-anatomy-UniversalIDFooter-anatomy-ID-anatomy-Art-fontSize`, `${S}-anatomy-UniversalIDFooter-anatomy-ID-anatomy-Art-color`] },
    ],
  },
  {
    name: 'sidebar sheet',
    html: PAGE(ex('open')),
    widths: [390],
    click: '.cui-console-sidebar-open',
    checks: [
      { target: SIDEBAR, props: ['box-left', 'box-top', 'border-top-left-radius', 'background-color', 'box-shadow'], tb: [`${S}-inset`, `${S}-inset`, `${S}-borderRadius`, `${S}-variant-variant-opaque-backgroundColor`, `${S}-variant-variant-opaque-boxShadow`] },
      { target: SIDEBAR, props: ['position', 'display'], golden: 'below 1024px, opened, a sheet over the page: a popover, fixed' },
      { target: '.cui-console-sidebar-close', props: ['display'], golden: 'its close button shown in the sheet' },
      { target: ITEM, props: ['min-height', 'padding-top', 'padding-left', 'border-top-left-radius', 'font-size', 'line-height'], tb: [`${MS}-anatomy-Item-minHeight`, `${MS}-anatomy-Item-paddingBlock`, `${MS}-anatomy-Item-paddingInline`, `${MS}-anatomy-Item-borderRadius`, `${MS}-anatomy-Item-variant-selected-none-anatomy-Label-textStyle-fontSize`, `${MS}-anatomy-Item-variant-selected-none-anatomy-Label-textStyle-lineHeight`] },
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
    name: 'sidebar dark',
    html: `<div class="cui-dark">${PAGE(ex('sidebar'))}</div>`,
    widths: [1200],
    checks: [
      { target: CURRENT, props: ['background-color', 'color'], golden: 'the dark side derived (tokens.css, decision 18)' },
    ],
  },
];
