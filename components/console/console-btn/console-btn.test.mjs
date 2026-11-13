import { examples } from '../../../test/harness.mjs';

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

export const cases = [
  {
    name: 'variants',
    html: ex('variants'),
    widths: DESKTOP,
    checks: [
      { target: GRAY, props: ['box-height', 'font-size', 'line-height', 'font-weight', 'letter-spacing', 'color', 'background-color', 'padding-left', 'padding-right', 'border-top-left-radius', 'border-top-width', 'min-width', 'column-gap', 'text-align', 'text-decoration-line', 'cursor', 'transition-duration', 'box-shadow'], asc: ['button', 'button'] },
      { target: GRAY, props: ['background-color', 'color', 'cursor'], state: 'hover', asc: ['hover-button', 'button'] },
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
      { target: RED, props: ['background-color'], state: 'hover', tb: [`${B}-colorRole-error-effect-hover-variant-gray-backgroundColor`] },
      { target: RED, props: ['background-color'], state: 'active', tb: [`${B}-colorRole-error-effect-press-variant-gray-backgroundColor`] },
      { target: RED, props: ['color'], state: 'active', golden: `red-700 pressed too: the toolkit's is lighter, ${FITTED}` },
      { target: RED_FILLED, props: ['background-color', 'color'], golden: `the red fill red-700: white on the toolkit's red is 3.3:1, ${FITTED}` },
      { target: RED_FILLED, props: ['background-color'], state: 'hover', golden: `hovered red-800, a step darker: ${FITTED}` },
      { target: RED_FILLED, props: ['color'], state: 'hover', tb: [`${B}-colorRole-error-effect-hover-variant-filled-color`] },
      { target: RED_FILLED, props: ['background-color'], state: 'active', golden: `pressed red-800 too: the toolkit's is lighter, ${FITTED}` },
      { target: RED_FILLED, props: ['color'], state: 'active', tb: [`${B}-colorRole-error-effect-press-variant-filled-color`] },
      { target: RED_PLAIN, props: ['color', 'background-color'], golden: `the red label red-700, no fill: ${FITTED}` },
      { target: RED_PLAIN, props: ['color', 'background-color'], state: 'hover', golden: `hovered, red-700, no fill: ${FITTED}` },
      { target: RED_PLAIN, props: ['color', 'background-color'], state: 'active', golden: `pressed, red-700, no fill: ${FITTED}` },
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
    name: 'in a row',
    html: '<div style="display: flex; width: 300px"><input style="flex: 1"><button class="cui-console-btn" type="button">Generate Again</button></div>',
    widths: [1200],
    checks: [
      { target: '.cui-console-btn', props: ['flex-shrink', 'box-width', 'box-height', 'white-space'], golden: 'a control keeps its size beside a field that fills the row: as wide as its label, which stays on one line in its 30px (squashed, the label wrapped out of it; muse wrote shrink-0 on seven buttons); its white-space App Store Connect\'s normal' },
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
    name: 'aria-disabled',
    html: '<a class="cui-console-btn" aria-disabled="true" href="#all">See All</a>',
    widths: [1200],
    checks: [
      { target: '.cui-console-btn', props: ['background-color', 'color'], tb: [`${B}-colorRole-neutral-effect-disabled-variant-gray-backgroundColor`, `${B}-colorRole-neutral-effect-disabled-variant-gray-color`] },
      { target: '.cui-console-btn', props: ['cursor'], golden: 'a link marked disabled: no pointer, as a disabled button' },
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
      { target: '[data-size="sm"]', props: ['font-size', 'line-height', 'font-weight'], tb: ['component-Tabs-variant-controlSize-sm-textStyle-fontSize', 'component-Tabs-variant-controlSize-sm-textStyle-lineHeight', 'component-Tabs-variant-controlSize-sm-textStyle-fontWeight'] },
      { target: '[data-size="lg"]', props: ['font-size', 'line-height', 'font-weight'], tb: ['component-Tabs-variant-controlSize-lg-textStyle-fontSize', 'component-Tabs-variant-controlSize-lg-textStyle-lineHeight', 'component-Tabs-variant-controlSize-lg-textStyle-fontWeight'] },
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
    name: 'icon sizes',
    html: '<div class="flex gap-3"><button class="cui-console-btn cui-console-btn-icon" data-size="sm" type="button" aria-label="Copy"></button><button class="cui-console-btn cui-console-btn-icon" data-size="lg" type="button" aria-label="Copy"></button></div>',
    widths: [1200],
    checks: ['sm', 'lg'].map((size) => ({ target: `[data-size="${size}"]`, props: ['font-size'], tb: [`component-Button-anatomy-artStart-withoutLabel-controlSize-${size}-fontSize`] })),
  },
  {
    name: 'capsule',
    html: ex('capsule'),
    widths: [1200],
    checks: [
      ...['sm', 'md', 'lg'].map((size) => ({ target: `.cui-console-btn-icon${size === 'md' ? ':not([data-size])' : `[data-size="${size}"]`}`, props: ['box-height', 'box-width', 'border-top-left-radius', 'border-bottom-right-radius'], tb: [`${B}-controlSize-${size}-blockSize`, `${B}-controlSize-${size}-minWidth`, `${B}-controlShape-capsule-controlSize-${size}-borderRadius`, `${B}-controlShape-capsule-controlSize-${size}-borderRadius`] })),
      { target: '.cui-console-btn:not(.cui-console-btn-icon)', props: ['box-height', 'border-top-left-radius', 'padding-left', 'padding-right'], tb: [`${B}-controlSize-md-blockSize`, `${B}-controlShape-capsule-controlSize-md-borderRadius`, `${B}-controlSize-md-paddingInlineStart`, `${B}-controlSize-md-paddingInlineEnd`] },
    ],
  },
  {
    name: 'capsule touch',
    html: ex('capsule'),
    touch: true,
    widths: [390],
    checks: [
      { target: '.cui-console-btn-icon[data-size="sm"]', props: ['box-width', 'box-height', 'border-top-left-radius'], golden: 'a capsule icon button 44px square on a touch screen, its corners half that: a circle still, as the toolkit\'s capsule is at each size' },
    ],
  },
  {
    name: 'touch target',
    html: ex('icon'),
    touch: true,
    widths: [390],
    checks: [
      { target: '.cui-console-btn-icon', props: ['box-width', 'box-height'], golden: 'an icon button 44px square on a touch screen, the HIG\'s 44pt (Accessibility, Mobility): a target drawn round a 30px button covered its neighbour\'s, 8px away' },
      { target: '.cui-console-btn:not(.cui-console-btn-icon)', props: ['box-height'], golden: 'a labelled button 44px tall on a touch screen, the HIG\'s 44pt (Accessibility, Mobility)' },
    ],
  },
  {
    name: 'touch sizes',
    html: ex('sizes'),
    touch: true,
    widths: [390],
    checks: [
      { target: '[data-size="sm"]', props: ['box-height'], golden: 'the small button 44px tall on a touch screen too, the HIG\'s 44pt (Accessibility, Mobility)' },
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
  {
    name: 'busy',
    html: ex('busy'),
    widths: [1200, 390],
    reducedMotion: 'reduce',
    checks: [
      { target: '.cui-console-btn', props: ['cursor'], golden: 'aria-busy, set by the page\'s script while the request runs: the pointer says so' },
      { target: '.cui-console-btn', pseudo: '::before', props: ['width', 'height', 'border-top-width', 'border-top-color', 'border-right-color', 'border-top-left-radius', 'animation-name', 'animation-duration', 'animation-iteration-count'], golden: 'a spinner before the label, at the label\'s size, a ring open on one side so its turn shows, the label kept so the button keeps its width and its name; under Reduce Motion it keeps turning, as iOS\'s activity indicator does (base.css)' },
    ],
  },
];
