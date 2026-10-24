// The console's button against App Store Connect's: the analytics sources
// page's gray button at rest and hovered (button:button,
// hover-button:button); the toolkit's button tokens (--tb-component-Button-*)
// for every size, the filled and plain variants and the red, at rest. The
// hovered and pressed colors, the fitted ones and the dark side are golden,
// each with its why.

import { examples } from '../../test/harness.mjs';

const ex = examples(import.meta.url);
const DESKTOP = [1600, 1200, 390];
const B = 'component-Button-variant';
const at = (size) => [`${B}-controlSize-${size}-blockSize`, `${B}-controlSize-${size}-paddingInlineStart`, `${B}-controlSize-${size}-paddingInlineEnd`, `${B}-controlShape-roundedRect-controlSize-${size}-borderRadius`, `${B}-controlSize-${size}-gap`, `${B}-controlSize-${size}-minWidth`];
const SIZE = ['box-height', 'padding-left', 'padding-right', 'border-top-left-radius', 'column-gap', 'min-width'];
const GRAY = '.flex > :nth-child(1)';
const FILLED = '.flex > :nth-child(2)';
const PLAIN = '.flex > :nth-child(3)';
const RED = '.flex > :nth-child(4)';
const RED_FILLED = '.flex > :nth-child(5)';
const RED_PLAIN = '.flex > :nth-child(6)';
const FITTED = 'fitted to 4.5:1 (tokens.css, decision 21)';
const STATE = 'its own hovered or pressed color, the toolkit\'s (tokens.css)';

export const cases = [
  {
    name: 'variants',
    html: ex('variants'),
    widths: DESKTOP,
    checks: [
      { target: GRAY, props: ['box-height', 'font-size', 'line-height', 'font-weight', 'letter-spacing', 'color', 'background-color', 'padding-left', 'padding-right', 'border-top-left-radius', 'border-top-width', 'min-width', 'column-gap', 'text-align', 'text-decoration-line', 'cursor', 'transition-duration', 'box-shadow'], asc: ['button', 'button'] },
      { target: GRAY, props: ['background-color', 'color', 'cursor'], state: 'hover', asc: ['hover-button', 'button'] },
      { target: GRAY, props: SIZE, tb: at('md') },
      { target: GRAY, props: ['background-color', 'color'], state: 'active', tb: [`${B}-colorRole-neutral-effect-press-variant-gray-backgroundColor`, `${B}-colorRole-neutral-effect-press-variant-gray-color`] },
      { target: FILLED, props: ['background-color', 'color'], tb: [`${B}-colorRole-brand-effect-default-variant-filled-backgroundColor`, `${B}-colorRole-brand-effect-default-variant-filled-color`] },
      { target: FILLED, props: ['background-color'], state: 'hover', golden: `the filled blue hovered a step darker, blue-600: white on the toolkit's lighter blue is 3.9:1, ${FITTED}` },
      { target: FILLED, props: ['background-color', 'color'], state: 'active', tb: [`${B}-colorRole-brand-effect-press-variant-filled-backgroundColor`, `${B}-colorRole-brand-effect-press-variant-filled-color`] },
      { target: PLAIN, props: ['background-color'], golden: 'no fill: transparent, where the toolkit writes its transparent white, the same nothing' },
      { target: PLAIN, props: ['color', 'padding-left', 'padding-right', 'min-width'], tb: [`${B}-colorRole-brand-effect-default-variant-plain-color`, `${B}-variant-plain-paddingInlineStart`, `${B}-variant-plain-paddingInlineEnd`, 'component-Button-anatomy-label-variant-variant-plain-minWidth'] },
      { target: PLAIN, props: ['color'], state: 'hover', golden: `the brand's label hovered a step darker, blue-600: the toolkit's lighter blue is 3.5:1 on white, ${FITTED}` },
      { target: PLAIN, props: ['color'], state: 'active', tb: [`${B}-colorRole-brand-effect-press-variant-plain-color`] },
      { target: RED, props: ['background-color'], tb: [`${B}-colorRole-error-effect-default-variant-gray-backgroundColor`] },
      { target: RED, props: ['color'], golden: `the red label red-700: the toolkit's is 3.6:1 on white, ${FITTED}` },
      { target: RED, props: ['color'], state: 'hover', golden: `red-700 hovered too: the toolkit's lighter red is 2.9:1 on white, ${FITTED}` },
      { target: RED_FILLED, props: ['background-color', 'color'], golden: `the red fill red-700: white on the toolkit's red is 3.3:1, ${FITTED}` },
      { target: RED_FILLED, props: ['background-color'], state: 'hover', golden: `hovered red-800, a step darker: ${FITTED}` },
      { target: RED_PLAIN, props: ['color', 'background-color'], golden: `the red label red-700, no fill: ${FITTED}` },
    ],
  },
  {
    name: 'alone',
    html: '<p><button class="cui-console-btn" type="button">Compare</button></p>',
    widths: [1200],
    checks: [
      { target: '.cui-console-btn', props: ['display', 'vertical-align', 'white-space'], asc: ['button', 'button'] },
    ],
  },
  {
    name: 'disabled',
    html: ex('disabled'),
    widths: [1200],
    checks: [
      { target: GRAY, props: ['background-color', 'color'], tb: [`${B}-colorRole-neutral-effect-disabled-variant-gray-backgroundColor`, `${B}-colorRole-neutral-effect-disabled-variant-gray-color`] },
      { target: FILLED, props: ['background-color', 'color'], tb: [`${B}-colorRole-brand-effect-disabled-variant-filled-backgroundColor`, `${B}-colorRole-brand-effect-disabled-variant-filled-color`] },
      { target: PLAIN, props: ['color'], tb: [`${B}-colorRole-brand-effect-disabled-variant-plain-color`] },
      { target: RED, props: ['color'], tb: [`${B}-colorRole-error-effect-disabled-variant-gray-color`] },
      { target: GRAY, props: ['cursor'], golden: 'no pointer over a disabled button' },
      { target: GRAY, props: ['background-color', 'color'], state: 'hover', golden: 'hovered, a disabled button stays as it is' },
    ],
  },
  {
    name: 'sizes',
    html: ex('sizes'),
    widths: [1200],
    checks: [
      { target: '[data-size="sm"]', props: SIZE, tb: at('sm') },
      { target: '.flex > :nth-child(2)', props: SIZE, tb: at('md') },
      { target: '[data-size="lg"]', props: SIZE, tb: at('lg') },
      { target: '[data-size="xl"]', props: SIZE, tb: at('xl') },
      { target: '[data-size="sm"]', props: ['font-size', 'line-height', 'font-weight'], tb: ['component-Tabs-variant-controlSize-sm-textStyle-fontSize', 'component-Tabs-variant-controlSize-sm-textStyle-lineHeight', 'component-Tabs-variant-controlSize-sm-textStyle-fontWeight'] },
      { target: '[data-size="lg"]', props: ['font-size', 'line-height', 'font-weight'], tb: ['component-Tabs-variant-controlSize-lg-textStyle-fontSize', 'component-Tabs-variant-controlSize-lg-textStyle-lineHeight', 'component-Tabs-variant-controlSize-lg-textStyle-fontWeight'] },
      { target: '[data-size="xl"]', props: ['font-size', 'line-height', 'font-weight'], tb: ['component-Tabs-variant-controlSize-xl-textStyle-fontSize', 'component-Tabs-variant-controlSize-xl-textStyle-lineHeight', 'component-Tabs-variant-controlSize-xl-textStyle-fontWeight'] },
    ],
  },
  {
    name: 'icon',
    html: ex('icon'),
    widths: [1200],
    checks: [
      { target: '.cui-console-btn-icon', props: ['box-height', 'box-width', 'font-size'], tb: [`${B}-controlSize-md-blockSize`, `${B}-controlSize-md-minWidth`, 'component-Button-anatomy-artStart-withoutLabel-controlSize-md-fontSize'] },
      { target: '.cui-console-btn-icon svg', props: ['box-height', 'box-width'], tb: ['component-Button-anatomy-artStart-withoutLabel-controlSize-md-fontSize', 'component-Button-anatomy-artStart-withoutLabel-controlSize-md-fontSize'] },
      { target: '.cui-console-btn:not(.cui-console-btn-icon) svg', props: ['box-height', 'box-width'], golden: 'an icon beside a label at the label\'s size, 1em' },
    ],
  },
  {
    name: 'touch',
    html: ex('variants'),
    touch: true,
    widths: [390],
    checks: [
      { target: GRAY, props: ['background-color', 'color'], state: 'hover', rest: 'a tapped button keeps no hover (rule 5)' },
    ],
  },
  {
    name: 'dark',
    html: `<div class="cui-dark">${ex('variants')}</div>`,
    widths: [1200],
    checks: [
      { target: GRAY, props: ['background-color', 'color'], golden: 'the dark side derived (tokens.css, decision 18)' },
      { target: FILLED, props: ['background-color', 'color'], golden: `the dark blue fill dark blue-400: white on the derived blue-600 is 2.7:1, ${FITTED}` },
      { target: PLAIN, props: ['color'], golden: 'the dark side derived (tokens.css, decision 18)' },
      { target: RED, props: ['color'], golden: 'the dark side derived (tokens.css, decision 18)' },
      { target: RED_FILLED, props: ['background-color'], golden: `the dark red fill dark red-300: ${FITTED}` },
    ],
  },
];
