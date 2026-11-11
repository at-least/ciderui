import { examples, SIZE, PADDING, BORDER, RADIUS, COLORS, TEXT } from '../../../test/harness.mjs';

const ex = examples(import.meta.url);
const LOOK = [...SIZE, ...PADDING, ...BORDER, ...RADIUS, 'min-width', 'box-height', ...TEXT.filter((p) => p !== 'display')];
const FLEX = { props: ['display', 'align-items', 'justify-content', 'column-gap'], golden: 'inline-flex, not apple.com\'s inline-block: an icon half an em before its label, on its middle (apple.com\'s pill holds text alone)' };
// @external: what the home page's probes read (no text-align, wrapping or opacity; its
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
      { target: '.cui-btn', ...FLEX },
    ],
  },
  {
    name: 'sizes',
    html: ex('sizes'),
    checks: [
      { target: '.cui-btn:not([data-size])', props: LOOK, apple: ['store', 'button'] },
      { target: '.cui-btn[data-size="lg"]', props: LOOK, apple: ['store', 'button.elevated'] },
      { target: '.cui-btn[data-size="lg"]', ...OWN },
      { target: '.cui-btn[data-size="lg"]', ...FLEX },
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
          // @external: a <button>: the class over the preflight's reset of the element
          { target, props: [...LOOK, ...COLORS], apple: ['store', probe] },
          { target, state: 'hover', props: COLORS, golden: 'a disabled button keeps its ground under the pointer, where apple.com\'s changes it' },
        ];
      }),
    ],
  },
  ...[['block', 'a.cui-btn-block'], ['block-button', 'button.cui-btn-block']].map(([name, target]) => ({
    name,
    html: ex(name),
    checks: [
      { target, props: [...SIZE, ...PADDING, ...RADIUS, 'box-height', ...COLORS], apple: ['store', 'button.block'] },
      { target, props: ['display', 'box-width'], golden: 'flex, not apple.com\'s block, so an icon keeps its place; the row\'s whole width, a <button> as a link' },
    ],
  })),
  {
    name: 'icon',
    html: ex('icon'),
    widths: [1200],
    checks: [
      { target: '.cui-btn > svg', props: ['box-width', 'box-height'], golden: 'the icon 1em square, the label\'s 14px' },
      { target: '.cui-btn', props: ['box-height'], golden: 'the icon adds no height: the 36px pill' },
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
      { target: '.cui-btn:not([data-size])', pseudo: '::after', props: ['content'], golden: 'no invisible overlay where no touch screen is: the pointer finds the label, not a sheet over it' },
    ],
  },
  {
    name: 'busy',
    html: ex('busy'),
    widths: [1200, 390],
    reducedMotion: 'reduce',
    checks: [
      { target: '.cui-btn', props: ['cursor'], golden: 'aria-busy, set by the page\'s script while the request runs: the pointer says so' },
      { target: '.cui-btn', pseudo: '::before', props: ['width', 'height', 'border-top-width', 'border-top-left-radius', 'animation-name', 'animation-duration', 'animation-iteration-count'], golden: 'a spinner before the label, at the label\'s size, the label kept so the button keeps its width and its name; under Reduce Motion it keeps turning, as iOS\'s activity indicator does (base.css)' },
    ],
  },
];
