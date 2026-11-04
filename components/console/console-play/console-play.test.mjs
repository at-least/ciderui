import { examples } from '../../../test/harness.mjs';

const ex = examples(import.meta.url);
const B = 'component-Button-variant';
const PAUSED = '.cui-console-play:not([data-state])';
const PLAYING = '.cui-console-play[data-state="playing"]';
const SIZES = ['sm', 'md', 'lg', 'xl'];
const sized = (size) => (size === 'md' ? '.cui-console-play:not([data-size])' : `.cui-console-play[data-size="${size}"]`);

export const cases = [
  {
    name: 'play',
    html: ex('play'),
    widths: [1600, 1200, 390],
    checks: [
      { target: PAUSED, props: ['box-height', 'box-width', 'border-top-left-radius', 'font-size'], tb: [`${B}-controlSize-md-blockSize`, `${B}-controlSize-md-minWidth`, `${B}-controlShape-capsule-controlSize-md-borderRadius`, 'component-Button-anatomy-artStart-withoutLabel-controlSize-md-fontSize'] },
      { target: PAUSED, props: ['background-color'], tb: [`${B}-colorRole-neutral-effect-default-variant-gray-backgroundColor`] },
      { target: PAUSED, props: ['background-color'], state: 'hover', tb: [`${B}-colorRole-neutral-effect-hover-variant-gray-backgroundColor`] },
      { target: PAUSED, props: ['background-color'], state: 'active', tb: [`${B}-colorRole-neutral-effect-press-variant-gray-backgroundColor`] },
      { target: PAUSED, props: ['color'], golden: 'the gray button\'s label color, as console-btn\'s icon buttons beside it: the toolkit\'s icon-only gray (its artOnlyColor, 60% gray) is not taken here' },
      { target: `${PAUSED} > .cui-console-play-paused`, props: ['box-height', 'box-width'], tb: ['component-Button-anatomy-artStart-withoutLabel-controlSize-md-fontSize', 'component-Button-anatomy-artStart-withoutLabel-controlSize-md-fontSize'] },
      { target: `${PAUSED} > .cui-console-play-paused`, props: ['display', 'fill', 'stroke'], golden: 'paused: Lucide\'s play shown, filled in the label\'s color and not stroked (in a translucent color the stroke over the fill drew a darker outline) (the toolkit has no play glyph; tv.apple.com\'s are SF Symbols, Apple\'s)' },
      { target: `${PAUSED} > .cui-console-play-playing`, props: ['display'], golden: 'paused: the pause glyph hidden' },
      { target: `${PLAYING} > .cui-console-play-paused`, props: ['display'], golden: 'playing (data-state, the page\'s script\'s): the play glyph hidden' },
      { target: `${PLAYING} > .cui-console-play-playing`, props: ['display', 'fill', 'stroke', 'box-width'], golden: 'playing: Lucide\'s pause in its place, filled, not stroked, at the play\'s size' },
    ],
  },
  {
    name: 'play sizes',
    html: ex('sizes'),
    widths: [1200],
    checks: SIZES.map((size) => ({ target: sized(size), props: ['box-height', 'box-width', 'border-top-left-radius', 'font-size'], tb: [`${B}-controlSize-${size}-blockSize`, `${B}-controlSize-${size}-minWidth`, `${B}-controlShape-capsule-controlSize-${size}-borderRadius`, `component-Button-anatomy-artStart-withoutLabel-controlSize-${size}-fontSize`] })),
  },
  {
    name: 'play focused',
    html: ex('play'),
    widths: [1200],
    checks: [
      { target: PAUSED, props: ['outline-style', 'outline-width', 'outline-offset'], state: 'focus-visible', golden: 'base.css\'s one ring, round the circle' },
    ],
  },
  {
    name: 'play touch',
    html: ex('sizes'),
    touch: true,
    widths: [390],
    checks: [
      { target: sized('sm'), props: ['box-width', 'box-height', 'border-top-left-radius'], golden: 'the small one 44px across on a touch screen, the HIG\'s 44pt, and still a circle (console-btn\'s capsule)' },
      { target: PAUSED.replace(':not([data-state])', '[data-size="sm"]'), props: ['background-color', 'color'], state: 'hover', rest: 'a tapped button keeps no hover (rule 5)' },
    ],
  },
  {
    name: 'play dark',
    html: `<div class="cui-dark">${ex('play')}</div>`,
    widths: [1200],
    checks: [
      { target: PAUSED, props: ['background-color', 'color'], golden: 'the dark side derived (tokens.css, decision 18)' },
      { target: `${PAUSED} > .cui-console-play-paused`, props: ['fill'], golden: 'filled in the dark label\'s color' },
    ],
  },
];
