// The console's badge against App Store Connect's toolkit's badge
// (--tb-component-Badge-*): each size, the tinted fill of each color, the
// gray and the plain. Golden, with their why: the labels fitted to 4.5:1,
// the dot, the dark side.

import { examples } from '../../../test/harness.mjs';

const ex = examples(import.meta.url);
const B = 'component-Badge-variant';
const COLORS = ['neutral', 'brand', 'info', 'success', 'warning', 'error'];
const at = (c) => `.cui-console-badge${c === 'neutral' ? ':not([data-color])' : `[data-color="${c}"]`}`;
const md = [`${B}-size-md-height`, `${B}-size-md-paddingInline`, `${B}-size-md-paddingInline`, `${B}-controlShape-capsule-size-md-borderRadius`, `${B}-size-md-textStyle-fontSize`, `${B}-size-md-textStyle-lineHeight`, `${B}-size-md-textStyle-fontWeight`];
const sm = md.map((t) => t.replace('-md-', '-sm-'));
const BOX = ['box-height', 'padding-left', 'padding-right', 'border-top-left-radius', 'font-size', 'line-height', 'font-weight'];
const FITTED = 'its label fitted to 4.5:1 on the page, its tint and the gray (tokens.css, decision 21)';

export const cases = [
  {
    name: 'badges',
    html: ex('badges'),
    widths: [1200, 390],
    checks: [
      { target: at('neutral'), props: BOX, tb: md },
      ...COLORS.map((c) => ({ target: at(c), props: ['background-color'], tb: [`${B}-colorRole-${c}-variant-tinted-backgroundColor`] })),
      { target: at('neutral'), props: ['color'], tb: [`${B}-colorRole-neutral-variant-tinted-color`] },
      { target: at('error'), props: ['color'], golden: `red-700: ${FITTED}` },
      ...['brand', 'info', 'success', 'warning'].map((c) => ({ target: at(c), props: ['color'], golden: `${c}: ${FITTED}` })),
    ],
  },
  {
    name: 'gray badges',
    html: ex('gray'),
    widths: [1200],
    checks: [
      ...COLORS.map((c) => ({ target: at(c), props: ['background-color'], tb: [`${B}-colorRole-${c}-variant-gray-backgroundColor`] })),
      { target: at('neutral'), props: ['color'], tb: [`${B}-colorRole-neutral-variant-gray-color`] },
    ],
  },
  {
    name: 'plain badges',
    html: ex('plain'),
    widths: [1200],
    checks: [
      ...COLORS.map((c) => ({ target: at(c), props: ['background-color'], golden: 'no fill: transparent, where the toolkit writes its transparent white, the same nothing' })),
      { target: at('neutral'), props: ['color'], tb: [`${B}-colorRole-neutral-variant-plain-color`] },
      { target: at('neutral'), props: ['padding-left', 'padding-right'], golden: 'plain, no room at its sides' },
      { target: `${at('error')} .cui-console-badge-dot`, props: ['box-width', 'box-height', 'border-top-left-radius', 'background-color'], golden: 'a dot before the label, 0.5em, in its label\'s color: the toolkit\'s badge has none' },
    ],
  },
  {
    name: 'small badges',
    html: ex('small'),
    widths: [1200],
    checks: [
      { target: at('neutral'), props: BOX, tb: sm },
    ],
  },
  {
    name: 'badges dark',
    html: `<div class="cui-dark">${ex('badges')}</div>`,
    widths: [1200],
    checks: COLORS.map((c) => ({ target: at(c), props: ['background-color', 'color'], golden: 'the dark side derived (tokens.css, decision 18)' })),
  },
];
