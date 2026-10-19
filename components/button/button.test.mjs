// The button against apple.com: each size against its shared button
// (store:button*), the call to action against the home page's tile buttons
// across 735px (home:button*), the variants on light and in a dark block
// (home:tile.dark-button*, home:gallery.button-neutral), disabled
// (store:button*.disabled), the rectangle (bag:form-button*), the block
// (store:button.block); golden: the letter-spacing, the font, a disabled
// button under the pointer; a touch screen's hover is rest.

import { examples, SIZE, PADDING, BORDER, RADIUS, COLORS, TEXT, OUTLINE } from '../../test/harness.mjs';

const ex = examples(import.meta.url);
const LOOK = [...SIZE, ...PADDING, ...BORDER, ...RADIUS, 'min-width', 'box-height', ...TEXT];
// what the home page's probes read (no text-align, wrapping or opacity; its
// primary button's cursor alone); not its display, which is its place's: its
// buttons are grid items, blockified (the case cta-grid)
const LOOK_HOME = [...SIZE, ...PADDING, ...BORDER, ...RADIUS, 'min-width', 'box-height'];
const OWN = { props: ['letter-spacing', 'font-family'], golden: 'the system font\'s fitted letter-spacing (tokens.css) and the kit\'s stack, not apple.com\'s' };
const states = (target, probe, props = COLORS) => [
  { target, state: 'hover', props, apple: [...probe, 'hover'] },
  { target, state: 'active', props, apple: [...probe, 'active'] },
];
const STEP = { 734: 390, 735: 900 };

export const cases = [
  {
    name: 'default',
    html: ex('default'),
    checks: [
      { target: '.cui-btn', props: [...LOOK, ...COLORS], apple: ['store', 'button'] },
      ...states('.cui-btn', ['store', 'button']),
      { target: '.cui-btn', state: 'focus-visible', props: OUTLINE, apple: ['store', 'button', 'focus-visible'] },
      { target: '.cui-btn', ...OWN },
    ],
  },
  {
    name: 'sizes',
    html: ex('sizes'),
    checks: [
      { target: '.cui-btn[data-size="sm"]', props: LOOK, apple: ['store', 'button.reduced'] },
      { target: '.cui-btn:not([data-size])', props: LOOK, apple: ['store', 'button'] },
      { target: '.cui-btn[data-size="lg"]', props: LOOK, apple: ['store', 'button.elevated'] },
      { target: '.cui-btn[data-size="xl"]', props: LOOK, apple: ['store', 'button.super'] },
      ...['sm', 'lg', 'xl'].map(s => ({ target: `.cui-btn[data-size="${s}"]`, ...OWN })),
    ],
  },
  {
    name: 'cta',
    html: ex('cta'),
    widths: [390, 734, 735, 900, 1200, 1600],
    checks: [
      { target: '.cui-btn:not([data-variant])', props: [...LOOK_HOME, 'cursor', ...COLORS], apple: ['home', 'button'], at: STEP },
      { target: '.cui-btn[data-variant="secondary"]', props: [...LOOK_HOME, ...COLORS], apple: ['home', 'button.secondary'], at: STEP },
      ...states('.cui-btn[data-variant="secondary"]', ['home', 'button.secondary']).map(k => ({ ...k, at: STEP })),
      { target: '.cui-btn:not([data-variant])', ...OWN },
    ],
  },
  {
    // as the home page places them, in a grid (home:hero.ctas): a grid item is
    // blockified, which is apple.com's display there
    name: 'cta-grid',
    html: `<div style="display: grid; grid-auto-flow: column; justify-content: center; column-gap: 17px">${ex('cta')}</div>`,
    checks: [
      { target: '.cui-btn:not([data-variant])', props: ['display'], apple: ['home', 'button'] },
      { target: '.cui-btn[data-variant="secondary"]', props: ['display'], apple: ['home', 'button.secondary'] },
    ],
  },
  {
    name: 'variants',
    html: ex('variants'),
    checks: [
      { target: '.cui-btn[data-variant="secondary"]', props: [...LOOK, ...COLORS], apple: ['store', 'button.secondary'] },
      ...states('.cui-btn[data-variant="secondary"]', ['store', 'button.secondary']),
      { target: '.cui-btn[data-variant="neutral"]', props: [...LOOK, ...COLORS], apple: ['store', 'button.neutral'] },
      ...states('.cui-btn[data-variant="neutral"]', ['store', 'button.neutral']),
    ],
  },
  {
    name: 'dark',
    html: ex('dark'),
    checks: [
      { target: '.cui-btn:not([data-variant])', props: [...SIZE, ...PADDING, ...RADIUS, 'min-width', ...COLORS], apple: ['home', 'tile.dark-button'] },
      ...states('.cui-btn:not([data-variant])', ['home', 'tile.dark-button']),
      { target: '.cui-btn[data-variant="secondary"]', props: [...SIZE, ...PADDING, ...RADIUS, 'min-width', ...COLORS], apple: ['home', 'tile.dark-button-secondary'] },
      ...states('.cui-btn[data-variant="secondary"]', ['home', 'tile.dark-button-secondary']),
      // the gallery's button steps at 1069px, a size of its own: its colors
      { target: '.cui-btn[data-variant="neutral"]', props: COLORS, apple: ['home', 'gallery.button-neutral'] },
      ...states('.cui-btn[data-variant="neutral"]', ['home', 'gallery.button-neutral']),
    ],
  },
  {
    name: 'disabled',
    html: ex('disabled'),
    checks: [
      ...[['', 'button.disabled'], ['secondary', 'button.secondary.disabled'], ['neutral', 'button.neutral.disabled']].flatMap(([v, probe]) => {
        const target = v ? `.cui-btn[data-variant="${v}"]` : '.cui-btn:not([data-variant])';
        return [
          // a <button>: the class over the preflight's reset of the element
          { target, props: [...LOOK, ...COLORS], apple: ['store', probe] },
          { target, state: 'hover', props: COLORS, golden: 'a disabled button keeps its ground under the pointer, where apple.com\'s changes it' },
        ];
      }),
    ],
  },
  {
    name: 'rect',
    html: ex('rect'),
    checks: [
      { target: '.cui-btn-rect:not([data-variant])', props: [...SIZE, ...PADDING, ...BORDER, ...RADIUS, 'min-width', 'box-height', ...COLORS, 'white-space', 'text-align', 'cursor'], apple: ['bag', 'form-button'] },
      ...states('.cui-btn-rect:not([data-variant])', ['bag', 'form-button']),
      { target: '.cui-btn-rect[data-variant="secondary"]', props: [...SIZE, ...PADDING, ...RADIUS, ...COLORS], apple: ['bag', 'form-button.secondary'] },
      ...states('.cui-btn-rect[data-variant="secondary"]', ['bag', 'form-button.secondary']),
    ],
  },
  {
    name: 'block',
    html: ex('block'),
    checks: [
      // its width is its container's: a block's, not checked
      { target: '.cui-btn-block', props: [...SIZE, ...PADDING, ...RADIUS, 'display', 'box-height', ...COLORS], apple: ['store', 'button.block'] },
    ],
  },
  {
    name: 'touch',
    html: ex('variants'),
    touch: true,
    widths: [390],
    checks: ['.cui-btn:not([data-variant])', '.cui-btn[data-variant="secondary"]', '.cui-btn[data-variant="neutral"]']
      .map(target => ({ target, state: 'hover', props: COLORS, rest: 'a touch screen keeps no hover' })),
  },
];
