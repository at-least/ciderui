// The console's checkbox against App Store Connect's toolkit's checkbox and
// label (--tb-component-Checkbox-*, -Label-*: brand, md): its size, corners,
// edge's width, the gap to its label, the label's type, the disabled fill.
// Golden, with their why: the edge (fitted), the flat blue checked, the
// check and the dash, the ring on the box, the dark side.

import { examples } from '../../../test/harness.mjs';

const ex = examples(import.meta.url);
const C = 'component-Checkbox-variant';
const OFF = 'label:nth-child(1)';
const ON = 'label:nth-child(2)';
const DISABLED = 'label:nth-child(3)';

export const cases = [
  {
    name: 'checkbox',
    html: ex('checkbox'),
    widths: [1200, 390],
    checks: [
      { target: `${OFF} .cui-console-checkbox-box`, props: ['box-width', 'box-height', 'border-top-left-radius', 'border-top-width', 'border-top-style'], tb: [`${C}-controlSize-md-size`, `${C}-controlSize-md-size`, `${C}-controlSize-md-borderRadius`, `${C}-controlSize-md-effect-default-state-off-border-width`, `${C}-controlSize-md-effect-default-state-off-border-style`] },
      { target: OFF, props: ['column-gap', 'font-size', 'line-height', 'font-weight', 'color'], tb: ['component-Label-anatomy-checkbox-variant-controlSize-md-inline-on-gap', 'component-Label-variant-controlSize-md-inline-on-textStyle-fontSize', 'component-Label-variant-controlSize-md-inline-on-textStyle-lineHeight', 'component-Label-variant-controlSize-md-inline-on-textStyle-fontWeight', 'component-Label-color'] },
      { target: `${OFF} .cui-console-checkbox-box`, props: ['border-top-color', 'background-color'], golden: 'its edge the field\'s, coolGray-800 at 0.58 (3:1; the toolkit\'s 0.18 is 1.4:1: tokens.css, decision 21), on white' },
      { target: `${ON} .cui-console-checkbox-box`, props: ['background-color', 'border-top-color'], golden: 'checked, the brand\'s flat blue, its edge gone: the toolkit\'s is a gradient around it under a shadow' },
      { target: `${ON} .cui-console-checkbox-box svg:first-child`, props: ['visibility', 'color', 'box-width'], golden: 'checked, Lucide\'s check, white, 11px' },
      { target: `${OFF} .cui-console-checkbox-box svg:first-child`, props: ['visibility'], golden: 'unchecked, no check' },
      { target: `${DISABLED} .cui-console-checkbox-box`, props: ['background-color'], tb: [`${C}-effect-disabled-backgroundColor`] },
      { target: `${DISABLED} .cui-console-checkbox-box`, props: ['border-top-color'], tb: [`${C}-controlSize-md-effect-disabled-state-off-border-color`] },
      { target: DISABLED, props: ['color'], golden: 'a disabled checkbox\'s label the disabled label' },
      { target: `${OFF} .cui-console-checkbox-input`, props: ['opacity', 'box-width', 'box-height', 'box-left', 'box-top'], golden: 'the input transparent, the box\'s size, over it' },
      { target: `${OFF} .cui-console-checkbox-box`, props: ['box-left', 'box-top'], golden: 'the box where its input is' },
    ],
  },
  {
    name: 'checkbox focus',
    html: ex('checkbox'),
    widths: [1200],
    press: ['Tab'],
    checks: [
      { target: `${OFF} .cui-console-checkbox-box`, props: ['outline-style', 'outline-width', 'outline-color', 'outline-offset'], golden: 'focused from the keyboard, base.css\'s ring on the box (rule 9; the toolkit\'s is 3px of the blue at 0.15)' },
    ],
  },
  {
    name: 'checkbox indeterminate',
    html: `${ex('checkbox')}<script>document.querySelector('[name="all"]').indeterminate = true</script>`,
    widths: [1200],
    checks: [
      { target: '[name="all"] + .cui-console-checkbox-box', props: ['background-color'], golden: 'indeterminate, the brand\'s blue' },
      { target: '[name="all"] + .cui-console-checkbox-box svg:last-child', props: ['visibility'], golden: 'indeterminate, Lucide\'s minus' },
      { target: '[name="all"] + .cui-console-checkbox-box svg:first-child', props: ['visibility'], golden: 'indeterminate, no check' },
    ],
  },
  {
    name: 'checkbox touch',
    html: ex('checkbox'),
    touch: true,
    widths: [390],
    checks: [
      { target: `${OFF} .cui-console-checkbox-input`, props: ['box-width', 'box-height'], golden: 'the transparent input 44px square on a touch screen, the HIG\'s 44pt (Accessibility, Mobility)' },
      { target: `${OFF} .cui-console-checkbox-box`, props: ['box-width', 'box-height'], golden: 'the box drawn as before, 15px' },
      { target: OFF, props: ['box-height'], golden: 'a checkbox with its label 44px tall on a touch screen, as the radio, so its input\'s 44px fits its row: stacked 12px apart, the next one\'s reached 14px into it' },
    ],
  },
  {
    name: 'checkbox dark',
    html: `<div class="cui-dark">${ex('checkbox')}</div>`,
    widths: [1200],
    checks: [
      { target: `${OFF} .cui-console-checkbox-box`, props: ['border-top-color', 'background-color'], golden: 'the dark side: its edge at 0.42 (tokens.css, decisions 18, 21)' },
      { target: `${ON} .cui-console-checkbox-box`, props: ['background-color'], golden: 'the dark blue fill, fitted (tokens.css, decision 21)' },
    ],
  },
];
