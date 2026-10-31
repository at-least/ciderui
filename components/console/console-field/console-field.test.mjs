// The console's fields against App Store Connect's toolkit's text input,
// label, help text, input error and form field (--tb-component-TextInput-*,
// -Label-*, -HelpText-*, -InputError-*, -FormField-*: stroked, md, and each
// size): their size, room, corners, type, ground, focused and in error.
// Golden, with their why: the edge, the placeholder, the help's and the
// error's colors, all fitted; the chevron; the dark side.

import { examples } from '../../../test/harness.mjs';

const ex = examples(import.meta.url);
const T = 'component-TextInput-variant';
const NAME = '#project';
const BAD = '#url';
const OFF = '#dsn';
const size = (s) => [`${T}-controlSize-${s}-blockSize`, `${T}-controlShape-roundedRect-controlSize-${s}-paddingInlineStart`, `${T}-controlShape-roundedRect-controlSize-${s}-paddingInlineEnd`, `${T}-controlShape-roundedRect-controlSize-${s}-borderRadius`, `${T}-controlSize-${s}-textStyle-fontSize`, `${T}-controlSize-${s}-textStyle-lineHeight`, `${T}-controlSize-${s}-textStyle-fontWeight`];
const SIZE = ['box-height', 'padding-left', 'padding-right', 'border-top-left-radius', 'font-size', 'line-height', 'font-weight'];
const FITTED = 'fitted (tokens.css, decision 21)';

export const cases = [
  {
    name: 'fields',
    html: ex('fields'),
    widths: [1200],
    checks: [
      { target: NAME, props: SIZE, tb: size('md') },
      { target: NAME, props: ['background-color', 'border-top-width', 'border-top-style', 'color'], tb: [`${T}-colorRole-neutral-variant-stroked-backgroundColor`, `${T}-colorRole-neutral-focus-off-variant-stroked-border-width`, `${T}-colorRole-neutral-focus-off-variant-stroked-border-style`, 'component-TextInput-anatomy-input-variant-effect-default-color'] },
      { target: NAME, props: ['border-top-color'], golden: `the edge coolGray-800 at 0.58, 3:1 on white: the toolkit's 0.18 is 1.4:1, ${FITTED}` },
      { target: NAME, props: ['border-top-color', 'outline-width', 'outline-style', 'outline-color'], state: 'focus', tb: [`${T}-colorRole-neutral-focus-on-variant-stroked-border-color`, `${T}-colorRole-neutral-focus-on-outline-width`, `${T}-colorRole-neutral-focus-on-outline-style`, `${T}-colorRole-neutral-focus-on-outline-color`] },
      { target: NAME, props: ['outline-offset'], state: 'focus', golden: 'its halo outside its edge, none between' },
      { target: BAD, props: ['background-color', 'border-top-width'], tb: [`${T}-colorRole-error-focus-off-variant-stroked-backgroundColor`, `${T}-colorRole-error-effect-default-variant-stroked-border-width`] },
      { target: BAD, props: ['border-top-color'], golden: `the red edge the red label, red-700: ${FITTED}` },
      { target: BAD, props: ['background-color', 'outline-color'], state: 'focus', tb: [`${T}-colorRole-error-focus-on-variant-stroked-backgroundColor`, `${T}-colorRole-error-focus-on-outline-color`] },
      { target: OFF, props: ['background-color', 'color'], tb: [`${T}-effect-disabled-variant-stroked-backgroundColor`, 'component-TextInput-anatomy-input-variant-effect-disabled-color'] },
      { target: '#search', props: ['color'], pseudo: '::placeholder', golden: `the placeholder the secondary label at 0.74: the toolkit's, at 0.39, is 2.1:1 on white, ${FITTED}` },
      { target: '.cui-console-label', props: ['font-size', 'line-height', 'font-weight', 'color', 'margin-bottom'], tb: ['component-Label-variant-controlSize-md-inline-off-textStyle-fontSize', 'component-Label-variant-controlSize-md-inline-off-textStyle-lineHeight', 'component-Label-variant-controlSize-md-inline-off-textStyle-fontWeight', 'component-Label-color', 'component-FormField-anatomy-label-variant-controlSize-md-marginBlockEnd'] },
      { target: '.cui-console-help', props: ['font-size', 'line-height', 'font-weight', 'margin-top'], tb: ['component-HelpText-variant-controlSize-md-textStyle-fontSize', 'component-HelpText-variant-controlSize-md-textStyle-lineHeight', 'component-HelpText-variant-controlSize-md-textStyle-fontWeight', 'component-FormField-anatomy-field-variant-controlSize-md-marginBlockEnd'] },
      { target: '.cui-console-help', props: ['color'], golden: `the secondary label at 0.74, ${FITTED}` },
      { target: '.cui-console-error', props: ['font-size', 'line-height', 'font-weight'], tb: ['component-InputError-variant-controlSize-md-textStyle-fontSize', 'component-InputError-variant-controlSize-md-textStyle-lineHeight', 'component-InputError-variant-controlSize-md-textStyle-fontWeight'] },
      { target: '.cui-console-error', props: ['color'], golden: `red-700, the red label: the toolkit's red is 3.6:1 on white, ${FITTED}` },
      { target: '#note', props: ['padding-top', 'padding-bottom', 'resize'], golden: 'a multi-line field: its height its rows\', 6px above and below, taller by its handle' },
    ],
  },
  {
    name: 'fields phone',
    html: ex('fields'),
    widths: [639, 390],
    checks: [
      { target: NAME, props: ['font-size', 'line-height'], tb: ['typography-textStyle-bodyMobileRegular-fontSize', 'typography-textStyle-bodyMobileRegular-lineHeight'] },
      { target: NAME, props: ['box-height'], tb: [`${T}-controlSize-md-blockSize`] },
    ],
  },
  {
    name: 'field sizes',
    html: ex('sizes'),
    widths: [1200],
    checks: ['sm', 'lg', 'xl'].map((s) => ({ target: `[data-size="${s}"]`, props: SIZE, tb: size(s) })),
  },
  {
    name: 'select',
    html: ex('select'),
    widths: [1200],
    checks: [
      { target: 'select', props: SIZE.filter((p) => p !== 'padding-right'), tb: size('md').filter((_, i) => i !== 2) },
      { target: 'select', props: ['padding-right'], golden: 'room for its chevron at its end' },
      { target: '.cui-console-select-chevron', props: ['box-width', 'box-height', 'right', 'color'], golden: 'its chevron, Lucide\'s chevrons-up-down, 13px, 8px in, the secondary label' },
    ],
  },
  {
    name: 'fields dark',
    html: `<div class="cui-dark">${ex('fields')}</div>`,
    widths: [1200],
    checks: [
      { target: NAME, props: ['background-color', 'border-top-color', 'color'], golden: 'the dark side derived; its edge fitted to 3:1 on black at 0.42 (tokens.css, decisions 18, 21)' },
      { target: BAD, props: ['border-top-color', 'background-color'], golden: 'the dark side derived (tokens.css, decision 18)' },
    ],
  },
];
