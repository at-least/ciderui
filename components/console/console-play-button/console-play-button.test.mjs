import { examples } from '../../../test/harness.mjs';

const ex = examples(import.meta.url);
const B = 'component-Button-variant';
const PAUSED = '.cui-console-play-button:not([data-state])';
const PLAYING = '.cui-console-play-button[data-state="playing"]';
const SIZES = ['sm', 'md', 'lg'];
const RING = 'playing, a ring round the glyph fills with how much has played (the owner, 2026-11-05: a progress ring round the icon, as Apple\'s Voice Memos and Podcasts mini buttons; no reference measured, so golden): 0.1875em out from the glyph, its line 0.125em, in the label\'s color on a track of it at a quarter (color-mix), from the top round clockwise (a conic gradient to --cui-console-play-button-progress, the page\'s script\'s)';
const LABELLED_PAUSED = '.cui-console-play-button:not([data-state])';
const LABELLED_PLAYING = '.cui-console-play-button[data-state="playing"]';
// @invariant: the page's script's progress on the playing one
const halfway = (html) => html.replace('data-state="playing" type="button">', 'data-state="playing" type="button" style="--cui-console-play-button-progress: 40%">');
const sized = (size) => (size === 'md' ? '.cui-console-play-button:not([data-size])' : `.cui-console-play-button[data-size="${size}"]`);

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
      { target: `${PAUSED} .cui-console-play-button-glyph > .cui-console-play-button-paused`, props: ['box-height', 'box-width'], tb: ['component-Button-anatomy-artStart-withoutLabel-controlSize-md-fontSize', 'component-Button-anatomy-artStart-withoutLabel-controlSize-md-fontSize'] },
      { target: `${PAUSED} .cui-console-play-button-glyph > .cui-console-play-button-paused`, props: ['display', 'fill', 'stroke'], golden: 'paused: Lucide\'s play shown, filled in the label\'s color and not stroked (in a translucent color the stroke over the fill drew a darker outline) (the toolkit has no play glyph; tv.apple.com\'s are SF Symbols, Apple\'s)' },
      { target: `${PAUSED} .cui-console-play-button-glyph > .cui-console-play-button-playing`, props: ['display'], golden: 'paused: the stop glyph hidden' },
      { target: `${PAUSED} .cui-console-play-button-glyph`, pseudo: '::before', props: ['content'], golden: 'paused: no ring' },
      { target: `${PLAYING} .cui-console-play-button-glyph > .cui-console-play-button-paused`, props: ['display'], golden: 'playing (data-state, the page\'s script\'s): the play glyph hidden' },
      { target: `${PLAYING} .cui-console-play-button-glyph > .cui-console-play-button-playing`, props: ['display', 'fill', 'stroke', 'box-width'], golden: 'playing: Lucide\'s square, stop (the owner, 2026-11-05: 「■ 停止」, where it was pause: a line replays from its start), filled, not stroked, at 0.625em, inside its ring' },
      { target: `${PLAYING} .cui-console-play-button-glyph`, pseudo: '::before', props: ['content', 'position', 'top', 'left', 'width', 'height', 'border-top-left-radius', 'background-image', 'mask-image'], golden: RING },
    ],
  },
  {
    name: 'play labelled',
    html: halfway(ex('labelled')),
    widths: [1200],
    checks: [
      { target: LABELLED_PAUSED, props: ['box-height', 'padding-left', 'column-gap', 'font-size', 'border-top-left-radius'], golden: 'muse\'s 「▶ 重聽這句」 (the owner, 2026-11-05): a small capsule console button, its glyph and its words, console-btn\'s small size and gap' },
      { target: `${LABELLED_PAUSED} > span.cui-console-play-button-paused`, props: ['display'], golden: 'paused, its words: what a press does' },
      { target: `${LABELLED_PAUSED} > span.cui-console-play-button-playing`, props: ['display'], golden: 'and not the playing ones' },
      { target: `${LABELLED_PAUSED} .cui-console-play-button-glyph`, props: ['box-width', 'box-height'], golden: 'the glyph the label\'s size, 1em' },
      { target: `${LABELLED_PLAYING} > span.cui-console-play-button-playing`, props: ['display'], golden: 'playing, its playing words (「■ 停止」), switched by data-state with the glyph' },
      { target: `${LABELLED_PLAYING} > span.cui-console-play-button-paused`, props: ['display'], golden: 'and not the paused ones' },
      { target: `${LABELLED_PLAYING} .cui-console-play-button-glyph`, pseudo: '::before', props: ['background-image', 'width', 'height'], golden: `${RING}; 40% here` },
    ],
  },
  {
    name: 'play labelled forced colors',
    html: halfway(ex('labelled')),
    widths: [1200],
    forcedColors: true,
    checks: [
      { target: `${LABELLED_PLAYING} .cui-console-play-button-glyph`, pseudo: '::before', props: ['background-image'], golden: 'the ring in the theme\'s ButtonText on GrayText' },
    ],
  },
  {
    name: 'play sizes',
    html: ex('sizes'),
    widths: [1200],
    checks: SIZES.map((size) => ({ target: sized(size), props: ['box-height', 'box-width', 'border-top-left-radius', 'font-size'], tb: [`${B}-controlSize-${size}-blockSize`, `${B}-controlSize-${size}-minWidth`, `${B}-controlShape-capsule-controlSize-${size}-borderRadius`, `component-Button-anatomy-artStart-withoutLabel-controlSize-${size}-fontSize`] })),
  },
  {
    name: 'play touch',
    html: ex('sizes'),
    touch: true,
    widths: [390],
    checks: [
      { target: sized('sm'), props: ['box-width', 'box-height', 'border-top-left-radius'], golden: 'the small one 44px across on a touch screen, the HIG\'s 44pt, and still a circle (console-btn\'s capsule)' },
    ],
  },
  {
    name: 'play dark',
    html: `<div class="cui-dark">${ex('play')}</div>`,
    widths: [1200],
    checks: [
      { target: PAUSED, props: ['background-color', 'color'], golden: 'the dark side derived (tokens.css, decision 18)' },
      { target: `${PAUSED} .cui-console-play-button-glyph > .cui-console-play-button-paused`, props: ['fill'], golden: 'filled in the dark label\'s color' },
    ],
  },
];
