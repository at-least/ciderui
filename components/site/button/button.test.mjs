import { examples, SIZE, PADDING, BORDER, RADIUS, COLORS, TEXT } from '../../../test/harness.mjs';

const ex = examples(import.meta.url);
const LOOK = [...SIZE, ...PADDING, ...BORDER, ...RADIUS, 'min-width', 'box-height', ...TEXT];
// what the home page's probes read (no text-align, wrapping or opacity; its
// primary button's cursor alone); not its display, which is its place's
const LOOK_HOME = [...SIZE, ...PADDING, ...BORDER, ...RADIUS, 'min-width', 'box-height'];
const OWN = { props: ['font-family'], golden: 'the kit\'s stack, not apple.com\'s' };
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
      { target: '.cui-btn', ...OWN },
    ],
  },
  {
    name: 'sizes',
    html: ex('sizes'),
    checks: [
      { target: '.cui-btn:not([data-size])', props: LOOK, apple: ['store', 'button'] },
      { target: '.cui-btn[data-size="lg"]', props: LOOK, apple: ['store', 'button.elevated'] },
      { target: '.cui-btn[data-size="lg"]', ...OWN },
    ],
  },
  {
    name: 'variants',
    html: ex('variants'),
    checks: [
      { target: '.cui-btn[data-variant="secondary"]', props: [...LOOK, ...COLORS], apple: ['store', 'button.secondary'] },
      ...states('.cui-btn[data-variant="secondary"]', ['store', 'button.secondary']),
      { target: '.cui-btn[data-variant="secondary-neutral"]', props: [...LOOK, ...COLORS], apple: ['store', 'button.secondary-neutral'] },
      ...states('.cui-btn[data-variant="secondary-neutral"]', ['store', 'button.secondary-neutral']),
    ],
  },
  {
    name: 'dark secondary-neutral',
    html: '<section class="cui-dark"><a class="cui-btn" data-variant="secondary-neutral" href="#buy">Buy</a></section>',
    widths: [1200],
    checks: [
      { target: '.cui-btn', props: COLORS, golden: 'the neutral\'s dark side, near-white on no ground: apple.com\'s in a dark block not measured' },
      { target: '.cui-btn', state: 'hover', props: COLORS, golden: 'filled with it hovered' },
    ],
  },
  {
    name: 'dark',
    html: ex('dark'),
    widths: [1600, 1200, 900],
    checks: [
      { target: '.cui-btn:not([data-variant])', props: [...SIZE, ...PADDING, ...RADIUS, 'min-width', ...COLORS], apple: ['home', 'tile.dark-button'] },
      ...states('.cui-btn:not([data-variant])', ['home', 'tile.dark-button']),
      { target: '.cui-btn[data-variant="secondary"]', props: [...SIZE, ...PADDING, ...RADIUS, 'min-width', ...COLORS], apple: ['home', 'tile.dark-button-secondary'] },
      ...states('.cui-btn[data-variant="secondary"]', ['home', 'tile.dark-button-secondary']),
    ],
  },
  {
    name: 'disabled',
    html: ex('disabled'),
    checks: [
      ...[['', 'button.disabled'], ['secondary', 'button.secondary.disabled']].flatMap(([v, probe]) => {
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
    name: 'block',
    html: ex('block'),
    checks: [
      // its width is its container's: a block's, not checked
      { target: '.cui-btn-block', props: [...SIZE, ...PADDING, ...RADIUS, 'display', 'box-height', ...COLORS], apple: ['store', 'button.block'] },
    ],
  },
  {
    name: 'touch target',
    html: ex('sizes'),
    touch: true,
    widths: [390],
    checks: [
      { target: '.cui-btn:not([data-size])', pseudo: '::after', props: ['height'], golden: 'the 36px pill\'s target 44px tall' },
    ],
  },
  {
    name: 'no touch target with a mouse',
    html: ex('sizes'),
    widths: [1200],
    checks: [
    ],
  },
];
