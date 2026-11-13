import { examples } from '../../../test/harness.mjs';

const ex = examples(import.meta.url);
const T = 'component-TextInput-variant';
const NAME = '#project';
const BAD = '#url';
const OFF = '#dsn';
const size = (s) => [`${T}-controlSize-${s}-blockSize`, `${T}-controlShape-roundedRect-controlSize-${s}-paddingInlineStart`, `${T}-controlShape-roundedRect-controlSize-${s}-paddingInlineEnd`, `${T}-controlShape-roundedRect-controlSize-${s}-borderRadius`, `${T}-controlSize-${s}-textStyle-fontSize`, `${T}-controlSize-${s}-textStyle-lineHeight`, `${T}-controlSize-${s}-textStyle-fontWeight`];
const SIZE = ['box-height', 'padding-left', 'padding-right', 'border-top-left-radius', 'font-size', 'line-height', 'font-weight'];
const FITTED = 'fitted (tokens.css, decision 21)';
const FILE = 'component-FileUpload-variant';
const FILE_ON = '#dsym';
const FILE_OFF = '#off';

export const cases = [
  {
    name: 'fields',
    html: ex('fields'),
    widths: [1200],
    checks: [
      { target: NAME, state: ['focus', 'focus-visible'], props: ['outline-offset'], golden: 'its halo starts at its edge, not the ring\'s gap away' },
      { target: NAME, props: SIZE, tb: size('md') },
      { target: NAME, props: ['background-color', 'border-top-width', 'border-top-style', 'color'], tb: [`${T}-colorRole-neutral-variant-stroked-backgroundColor`, `${T}-colorRole-neutral-focus-off-variant-stroked-border-width`, `${T}-colorRole-neutral-focus-off-variant-stroked-border-style`, 'component-TextInput-anatomy-input-variant-effect-default-color'] },
      { target: NAME, props: ['border-top-color'], golden: `the edge coolGray-800 at 0.58, 3:1 on white: the toolkit's 0.18 is 1.4:1, ${FITTED}` },
      { target: NAME, props: ['border-top-color', 'outline-width', 'outline-style', 'outline-color'], state: 'focus', tb: [`${T}-colorRole-neutral-focus-on-variant-stroked-border-color`, `${T}-colorRole-neutral-focus-on-outline-width`, `${T}-colorRole-neutral-focus-on-outline-style`, `${T}-colorRole-neutral-focus-on-outline-color`] },
      { target: NAME, props: ['outline-offset'], state: 'focus', golden: 'its halo outside its edge, none between' },
      { target: BAD, props: ['background-color', 'border-top-width'], tb: [`${T}-colorRole-error-focus-off-variant-stroked-backgroundColor`, `${T}-colorRole-error-effect-default-variant-stroked-border-width`] },
      { target: BAD, props: ['border-top-color'], golden: `the red edge the red label, red-700: ${FITTED}` },
      { target: BAD, props: ['background-color', 'outline-color'], state: 'focus', tb: [`${T}-colorRole-error-focus-on-variant-stroked-backgroundColor`, `${T}-colorRole-error-focus-on-outline-color`] },
      { target: OFF, props: ['background-color', 'color'], tb: [`${T}-effect-disabled-variant-stroked-backgroundColor`, 'component-TextInput-anatomy-input-variant-effect-disabled-color'] },
      { target: '#search', props: ['color'], pseudo: '::placeholder', golden: `the placeholder the secondary label at 0.76: the toolkit's, at 0.39, is 2.1:1 on white, ${FITTED}` },
      { target: '.cui-console-field-label', props: ['font-size', 'line-height', 'font-weight', 'color', 'margin-bottom'], tb: ['component-Label-variant-controlSize-md-inline-off-textStyle-fontSize', 'component-Label-variant-controlSize-md-inline-off-textStyle-lineHeight', 'component-Label-variant-controlSize-md-inline-off-textStyle-fontWeight', 'component-Label-color', 'component-FormField-anatomy-label-variant-controlSize-md-marginBlockEnd'] },
      { target: '.cui-console-field-error', props: ['font-size', 'line-height', 'font-weight'], tb: ['component-InputError-variant-controlSize-md-textStyle-fontSize', 'component-InputError-variant-controlSize-md-textStyle-lineHeight', 'component-InputError-variant-controlSize-md-textStyle-fontWeight'] },
      { target: '.cui-console-field-error', props: ['margin-top'], tb: ['component-FormField-anatomy-field-variant-controlSize-md-marginBlockEnd'] },
      { target: '.cui-console-field-error', props: ['color'], golden: `red-700, the red label: the toolkit's red is 3.6:1 on white, ${FITTED}` },
      { target: '#note', props: ['padding-top', 'padding-bottom', 'resize', 'box-height'], golden: 'a multi-line field: its height its 3 rows\' and its padding, not the 30px field\'s, 6px above and below, taller by its handle' },
      { target: NAME, props: ['box-width'], golden: 'as wide as its box, the example\'s 320px column' },
      { target: '.cui-console-field', props: ['display', 'flex-direction'], golden: 'its label, field and lines in a column' },
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
    name: 'fields touch',
    html: ex('fields'),
    touch: true,
    widths: [390],
    checks: [
      { target: NAME, props: ['box-height'], golden: 'a field 44px tall on a touch screen, the HIG\'s 44pt (Accessibility, Mobility)' },
    ],
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
  {
    name: 'file touch',
    html: ex('file'),
    touch: true,
    widths: [390],
    checks: [
      { target: FILE_ON, props: ['box-height', 'padding-top', 'padding-bottom'], golden: 'the field 44px tall on a touch screen, the HIG\'s 44pt (Accessibility, Mobility), its line centered in it: a file input\'s text sits at its padding' },
    ],
  },
  {
    name: 'file',
    html: ex('file'),
    widths: [1200],
    checks: [
      { target: FILE_ON, props: ['box-height', 'padding-left', 'padding-right', 'border-top-left-radius', 'font-size', 'line-height', 'font-weight', 'background-color', 'color'], tb: [`${FILE}-controlSize-md-blockSize`, `${FILE}-controlShape-roundedRect-controlSize-md-paddingInlineStart`, `${FILE}-controlShape-roundedRect-controlSize-md-paddingInlineEnd`, `${FILE}-controlShape-roundedRect-controlSize-md-borderRadius`, `${FILE}-controlSize-md-textStyle-fontSize`, `${FILE}-controlSize-md-textStyle-lineHeight`, `${FILE}-controlSize-md-textStyle-fontWeight`, `${FILE}-colorRole-neutral-variant-stroked-backgroundColor`, `${FILE}-effect-default-anatomy-ControlText-color`] },
      { target: FILE_ON, pseudo: '::file-selector-button', props: ['color'], tb: [`${FILE}-effect-default-empty-true-anatomy-ControlText-color`] },
      { target: FILE_ON, pseudo: '::file-selector-button', props: ['background-color', 'border-top-width', 'margin-right'], golden: 'its button the blue word alone, 8px before the name' },
      { target: FILE_ON, props: ['border-top-color', 'border-top-width'], golden: 'its edge the field\'s, coolGray-800 at 0.58 (3:1; the toolkit\'s 0.18 is 1.4:1: tokens.css, decision 21)' },
      { target: FILE_OFF, props: ['background-color'], tb: [`${FILE}-effect-disabled-variant-stroked-backgroundColor`] },
      { target: FILE_OFF, props: ['color'], tb: [`${FILE}-effect-disabled-anatomy-ControlText-color`] },
      { target: FILE_OFF, pseudo: '::file-selector-button', props: ['color'], tb: [`${FILE}-effect-disabled-anatomy-ControlText-color`] },
      { target: FILE_OFF, props: ['border-top-color'], golden: 'disabled, the tertiary edge, as a disabled checkbox\'s and radio\'s' },
      { target: FILE_OFF, props: ['cursor'], golden: 'disabled, no pointer over it, as over a disabled text field' },
    ],
  },
  {
    name: 'file focus',
    html: ex('file'),
    widths: [1200],
    press: ['Tab'],
    checks: [
      { target: FILE_OFF, pseudo: '::file-selector-button', props: ['cursor'], golden: 'disabled, its button does nothing and does not say it does' },
      { target: FILE_ON, pseudo: '::file-selector-button', props: ['cursor'], golden: 'its button opens the file picker' },
      { target: FILE_ON, props: ['cursor'], golden: 'all of it opens the file picker' },
      { target: FILE_ON, props: ['border-top-color', 'outline-width', 'outline-style', 'outline-color'], golden: 'focused, the text field\'s dark edge and 3px halo (the toolkit\'s focus outline)' },
    ],
  },
  {
    name: 'file phone',
    html: ex('file'),
    widths: [639, 390],
    checks: [
      { target: FILE_ON, props: ['font-size', 'line-height'], tb: ['typography-textStyle-bodyMobileRegular-fontSize', 'typography-textStyle-bodyMobileRegular-lineHeight'] },
    ],
  },
  {
    name: 'file dark',
    html: `<div class="cui-dark">${ex('file')}</div>`,
    widths: [1200],
    checks: [
      { target: FILE_ON, props: ['background-color', 'border-top-color', 'color'], golden: 'the dark side derived; its edge fitted (tokens.css, decisions 18, 21)' },
      { target: FILE_ON, pseudo: '::file-selector-button', props: ['color'], golden: 'the dark side derived (tokens.css, decision 18)' },
    ],
  },
];
