// The console's empty state against App Store Connect's toolkit's
// (--tb-component-EmptyState-*): each size's room, icon, heading, caption
// and their widths and gaps, the backed fill and corners. Golden, with
// their why: the secondary label's alpha, the actions' gap, the dark side.

import { examples } from '../../../test/harness.mjs';

const ex = examples(import.meta.url);
const E = 'component-EmptyState';
const sized = (s, M) => [
  { target: '.cui-console-empty', props: ['padding-top', 'padding-bottom'], tb: [`${E}-variant-size-${s}-paddingBlock`, `${E}-variant-size-${s}-paddingBlock`] },
  { target: '.cui-console-empty-icon', props: ['box-height', 'margin-bottom'], tb: [`${E}-variant-size-${s}-anatomy-Art-fontSize`, `${E}-variant-size-${s}-anatomy-Art-marginBlockEnd`] },
  { target: '.cui-console-empty-heading', props: ['font-size', 'line-height', 'font-weight', 'max-width', 'margin-bottom', 'color'], tb: ['fontSize', 'lineHeight', 'fontWeight'].map((p) => `${E}-variant-size-${s}-anatomy-Heading-textStyle-${p}`).concat([`${E}-maxInlineSize${M}`, `${E}-variant-size-${s}-anatomy-Heading-marginBlockEnd`, `${E}-anatomy-Heading-color`]) },
  { target: '.cui-console-empty-caption', props: ['font-size', 'line-height', 'font-weight', 'max-width', 'margin-bottom'], tb: ['fontSize', 'lineHeight', 'fontWeight'].map((p) => `${E}-variant-size-${s}-anatomy-Caption-textStyle-${p}`).concat([`${E}-maxInlineSize${M}`, `${E}-variant-size-${s}-anatomy-Caption-marginBlockEnd`]) },
  { target: '.cui-console-empty-caption', props: ['color'], golden: 'the secondary label fitted to 0.74 (tokens.css, decision 21)' },
  { target: '.cui-console-empty-icon', props: ['color'], golden: 'the secondary label fitted to 0.74 (tokens.css, decision 21)' },
];

export const cases = [
  {
    name: 'empty',
    html: ex('empty'),
    widths: [1200, 390],
    checks: [
      ...sized('md', 'Md'),
      { target: '.cui-console-empty', props: ['background-color'], golden: 'not backed: no fill' },
      { target: '.cui-console-empty-actions', props: ['column-gap'], golden: 'its buttons 9px apart, as a banner\'s' },
      { target: '.cui-console-empty', props: ['text-align', 'align-items'], golden: 'centered' },
    ],
  },
  {
    name: 'empty small',
    html: ex('small'),
    widths: [1200],
    checks: [
      ...sized('sm', 'Sm'),
      { target: '.cui-console-empty', props: ['background-color', 'border-top-left-radius'], tb: [`${E}-variant-backed-on-backgroundColor`, `${E}-borderRadius`] },
    ],
  },
  {
    name: 'empty large',
    html: ex('large'),
    widths: [1200],
    checks: sized('lg', 'Lg'),
  },
  {
    name: 'empty dark',
    html: `<div class="cui-dark">${ex('small')}</div>`,
    widths: [1200],
    checks: [
      { target: '.cui-console-empty', props: ['background-color', 'color'], golden: 'the dark side derived (tokens.css, decision 18)' },
    ],
  },
];
