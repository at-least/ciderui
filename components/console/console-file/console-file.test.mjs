// The console's file field against App Store Connect's toolkit's file
// upload (--tb-component-FileUpload-*: stroked, roundedRect, md, neutral):
// its height, room, corners, ground, type, the empty control's blue, the
// disabled ground and label. Golden, with their why: the edge (fitted),
// the button's gap, the focus, the phone's size, the dark side.

import { examples } from '../../../test/harness.mjs';

const ex = examples(import.meta.url);
const F = 'component-FileUpload-variant';
const ON = '#dsym';
const OFF = '#off';

export const cases = [
  {
    name: 'file',
    html: ex('file'),
    widths: [1200],
    checks: [
      { target: ON, props: ['box-height', 'padding-left', 'padding-right', 'border-top-left-radius', 'font-size', 'line-height', 'font-weight', 'background-color', 'color'], tb: [`${F}-controlSize-md-blockSize`, `${F}-controlShape-roundedRect-controlSize-md-paddingInlineStart`, `${F}-controlShape-roundedRect-controlSize-md-paddingInlineEnd`, `${F}-controlShape-roundedRect-controlSize-md-borderRadius`, `${F}-controlSize-md-textStyle-fontSize`, `${F}-controlSize-md-textStyle-lineHeight`, `${F}-controlSize-md-textStyle-fontWeight`, `${F}-colorRole-neutral-variant-stroked-backgroundColor`, `${F}-effect-default-anatomy-ControlText-color`] },
      { target: ON, pseudo: '::file-selector-button', props: ['color'], tb: [`${F}-effect-default-empty-true-anatomy-ControlText-color`] },
      { target: ON, pseudo: '::file-selector-button', props: ['background-color', 'border-top-width', 'margin-right'], golden: 'its button the blue word alone, 8px before the name' },
      { target: ON, props: ['border-top-color', 'border-top-width'], golden: 'its edge the field\'s, coolGray-800 at 0.58 (3:1; the toolkit\'s 0.18 is 1.4:1: tokens.css, decision 21)' },
      { target: OFF, props: ['background-color'], tb: [`${F}-effect-disabled-variant-stroked-backgroundColor`] },
      { target: OFF, props: ['color'], tb: [`${F}-effect-disabled-anatomy-ControlText-color`] },
      { target: OFF, props: ['border-top-color'], golden: 'disabled, the tertiary edge, as a disabled checkbox\'s and radio\'s' },
    ],
  },
  {
    name: 'file focus',
    html: ex('file'),
    widths: [1200],
    press: ['Tab'],
    checks: [
      { target: ON, props: ['border-top-color', 'outline-width', 'outline-style', 'outline-color'], golden: 'focused, the text field\'s dark edge and 3px halo (the toolkit\'s focus outline)' },
    ],
  },
  {
    name: 'file phone',
    html: ex('file'),
    widths: [639, 390],
    checks: [
      { target: ON, props: ['font-size', 'line-height'], tb: ['typography-textStyle-bodyMobileRegular-fontSize', 'typography-textStyle-bodyMobileRegular-lineHeight'] },
    ],
  },
  {
    name: 'file dark',
    html: `<div class="cui-dark">${ex('file')}</div>`,
    widths: [1200],
    checks: [
      { target: ON, props: ['background-color', 'border-top-color', 'color'], golden: 'the dark side derived; its edge fitted (tokens.css, decisions 18, 21)' },
      { target: ON, pseudo: '::file-selector-button', props: ['color'], golden: 'the dark side derived (tokens.css, decision 18)' },
    ],
  },
];
