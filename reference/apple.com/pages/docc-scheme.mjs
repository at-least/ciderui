// Apple's developer documentation's (DocC) appearance control in its
// footer: Light, Dark and Auto as one segmented control of radios, the
// group, a choice (hovered), the chosen one. docc-scheme-dark reads the same
// in the dark.
import { TYPE, BOX, FLEX } from './common.mjs';

const CHOICE = [...BOX, ...TYPE, 'display', 'min-width', 'text-align', 'cursor', 'border-left-width', 'border-left-style', 'border-left-color'];
export const probes = [
  { name: 'scheme', select: '.color-scheme-toggle', props: [...BOX, ...TYPE, ...FLEX, 'border-left-width', 'border-left-style', 'border-left-color'], box: true },
  { name: 'scheme.input', select: '.color-scheme-toggle input', props: ['position', 'width', 'height', 'opacity', 'clip-path', 'appearance'] },
  { name: 'scheme.choice', select: '.color-scheme-toggle label:not(:has(input:checked)) .text', props: CHOICE, box: true, states: ['hover'] },
  { name: 'scheme.chosen', select: '.color-scheme-toggle label:has(input:checked) .text', props: CHOICE, box: true },
];
export default {
  url: 'https://developer.apple.com/documentation/xcode/examining-the-fields-in-a-crash-report',
  wait: 4000,
  probes,
};
