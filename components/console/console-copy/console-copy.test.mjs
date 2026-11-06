import { examples, MOTION } from '../../../test/harness.mjs';

const ex = examples(import.meta.url);
const C = '.cui-console-copy';
const IDLE = `${C}:first-child`;
const COPIED = `${C}:last-child`;
const B = '.cui-console-copy-button';

export const cases = [
  {
    name: 'copy',
    html: ex('copy'),
    widths: [1200, 390],
    checks: [
      { target: `${IDLE} ${B}`, props: ['cursor'], apple: ['docc', 'copy'] },
      { target: `${IDLE} ${B}`, props: ['box-width', 'box-height', 'border-top-left-radius', 'opacity'], golden: 'a 28px circle (DocC\'s 33.5px on its 15px listing), always shown: DocC\'s shows only under a pointer, which a touch screen never has' },
      { target: `${IDLE} ${B}`, props: MOTION, golden: 'no motion, where DocC\'s fades in 0.2s: the toolkit\'s controls draw none' },
      { target: `${IDLE} ${B}`, props: ['margin-top', 'margin-right'], golden: '5px in from the top and the end, as much from each as DocC\'s, centering it on the first 22px line (8px + 11px - 14px)' },
      { target: C, props: ['border-top-left-radius', 'background-color'], golden: 'the listing\'s corners and gray, moved to the pair: the listing scrolls beside the button, never under it' },
      { target: `${IDLE} ${B} > .cui-console-copy-idle`, props: ['display', 'box-width'], golden: 'Lucide\'s copy at 16px, the console\'s icon size' },
      { target: `${IDLE} ${B} > .cui-console-copy-done`, props: ['display'], golden: 'the check hidden before a copy' },
      { target: `${IDLE} ${B}`, props: ['color', 'background-color'], golden: 'the label on the console\'s neutral fill-primary, the nearest the contract holds to DocC\'s rgb(210, 210, 215) over its gray' },
      { target: `${IDLE} ${B}`, props: ['color', 'background-color'], state: 'hover', golden: 'hovered, the console\'s gray hover under the same label: DocC\'s full invert to near-black only draws the eye (decision 25)' },
      { target: `${IDLE} .cui-console-listing`, props: ['padding-right', 'background-color'], golden: 'its end 9px, with the button\'s 5px the listing\'s own 14px; no ground of its own' },
      { target: `${IDLE} .cui-console-listing`, props: ['box-width'], golden: 'the scrolling text ends before the button: the pair\'s width less 38px (5px, 28px, 5px)' },
      { target: `${COPIED} ${B} > .cui-console-copy-idle`, props: ['display'], golden: 'copied: the copy icon hidden' },
      { target: `${COPIED} ${B} > .cui-console-copy-done`, props: ['display', 'box-width'], golden: 'copied: Lucide\'s check in its place' },
    ],
  },
  {
    name: 'copy in a card',
    html: ex('card'),
    widths: [1200],
    checks: [
      { target: `${IDLE}`, props: ['background-color'], golden: 'in a console card, whose gray is the listing\'s, the page\'s white: gray on gray drew no panel' },
    ],
  },
  {
    name: 'copy inline',
    html: ex('inline'),
    widths: [1200],
    checks: [
      { target: 'p:first-child .cui-console-copy-inline', props: ['box-width', 'box-height', 'margin-top', 'padding-top', 'color', 'background-color'], golden: 'a 16px icon with 5px around it (26px to hit), the margin taking it back from the line; the secondary label, no ground' },
      { target: 'p:first-child', props: ['box-height'], golden: 'the line as tall as the same line with no button (the last): the button grows nothing' },
      { target: 'p:last-child', props: ['box-height'], golden: 'the same line with no button: its mono value 20px, the button\'s line no taller' },
      { target: 'p:first-child .cui-console-copy-inline', props: ['color'], state: 'hover', golden: 'hovered, the label' },
      { target: 'p:nth-child(2) .cui-console-copy-inline > .cui-console-copy-done', props: ['display'], golden: 'copied: the check' },
      { target: 'p:nth-child(2) .cui-console-copy-inline > .cui-console-copy-idle', props: ['display'], golden: 'copied: the copy icon hidden' },
    ],
  },
  {
    name: 'copy inline phone',
    html: ex('inline'),
    widths: [320],
    checks: [
      { target: '#issue-id', props: ['overflow-wrap', 'box-right'], golden: 'a long id wraps anywhere in its line, staying in the window: it has no space to wrap at, and ran past a phone\'s edge' },
    ],
  },
  {
    name: 'copy inline touch',
    html: ex('inline'),
    touch: true,
    widths: [390],
    checks: [
      { target: 'p:first-child .cui-console-copy-inline', pseudo: '::after', props: ['width', 'height'], golden: 'a touch screen\'s target 44px square, the HIG\'s 44pt' },
    ],
  },
  {
    name: 'copy touch',
    html: ex('copy'),
    touch: true,
    widths: [390],
    checks: [
      { target: `${IDLE} ${B}`, pseudo: '::after', props: ['width', 'height'], golden: 'a touch screen\'s target 44px square, the HIG\'s 44pt' },
    ],
  },
  {
    name: 'copy dark',
    html: `<div class="cui-dark">${ex('copy')}</div>`,
    widths: [1200],
    checks: [
      { target: `${IDLE} ${B}`, props: ['color', 'background-color'], golden: 'the dark side derived (tokens.css, decision 18), as DocC\'s dark rgb(66, 66, 69) on its gray' },
      { target: `${IDLE} ${B}`, props: ['color', 'background-color'], state: 'hover', golden: 'hovered in dark, the console\'s gray hover, no invert (decision 25)' },
    ],
  },
];
