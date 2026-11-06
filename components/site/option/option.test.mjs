import { examples, PADDING, SIZE, WITH_1600 } from '../../../test/harness.mjs';

const ex = examples(import.meta.url);
const EDGE = ['border-top-width', 'border-top-style', 'border-top-color', 'border-bottom-width', 'border-bottom-color', 'border-top-left-radius'];
const INPUT = '.cui-option-input';
const LABEL = '.cui-option-label';

export const cases = [
  {
    name: 'zh-hant',
    html: `<div lang="zh-Hant">${ex('option')}</div>`,
    checks: [{ target: '.cui-option-title', props: ['line-height', 'letter-spacing'], apple: ['buy-tw', 'option.title'] }],
  },
  {
    name: 'option',
    widths: WITH_1600,
    html: ex('option'),
    checks: [
      { target: LABEL, props: [...SIZE, 'color', ...PADDING, ...EDGE, 'background-color', 'min-height', 'transition-property', 'transition-duration'], apple: ['buy', 'option'] },
      { target: LABEL, state: 'hover', props: ['background-color', 'border-top-color', 'color'], apple: ['buy', 'option', 'hover'] },
      { target: LABEL, on: INPUT, state: 'focus-visible', props: ['outline-style', 'outline-width', 'outline-color', 'outline-offset'], apple: ['buy', 'option.focus', 'focus-visible'] },
      { target: INPUT, props: ['position', 'top', 'left', 'opacity', 'margin-top', 'margin-left', 'display'], apple: ['buy', 'option.input'] },
      { target: '.cui-option-title', props: SIZE, apple: ['buy', 'option.title'] },
      { target: '.cui-option-title', props: ['letter-spacing', 'box-x'], golden: 'at the tile\'s start, 16px in (its edge and padding), the system font\'s letter-spacing' },
      { target: '.cui-option', props: ['background-color', 'border-top-left-radius', 'margin-top'], golden: 'its own light ground, the label\'s corners, as a field\'s; no margin above (apple.com\'s 14px, the form\'s spacing, buy:option.box)' },
      { target: LABEL, props: ['display', 'flex-direction', 'align-items'], golden: 'its title centered down the tile (apple.com\'s a flex column holding the row)' },
    ],
  },
  {
    name: 'checked',
    widths: WITH_1600,
    html: ex('checked'),
    checks: [
      { target: LABEL, props: [...SIZE, 'color', ...PADDING, ...EDGE, 'background-color', 'min-height', 'box-height'], apple: ['buy', 'option.checked'] },
      { target: '.cui-option-detail', props: [...SIZE, 'display', 'margin-top'], apple: ['buy', 'option.detail'] },
      { target: '.cui-option-title', props: ['box-width'], golden: 'the rest of the row: nothing beside it' },
    ],
  },
  {
    name: 'dark',
    html: `<div class="cui-dark">${ex('option')}</div>`,
    checks: [
      { target: LABEL, props: ['color', 'background-color', 'border-top-color'], golden: 'light in a dark block, as a field (apple.com shows options on light pages alone)' },
    ],
  },
  {
    name: 'disabled',
    html: ex('option').replace('value="256"', 'value="256" disabled'),
    widths: [1200],
    checks: [
      { target: LABEL, props: ['opacity', 'border-top-color'], golden: 'disabled, dimmed to 0.42 with its rest edge: apple.com\'s buy page, read 2026-10-31 (its disabled tiles 0.42, the base rule; a whole disabled step keeps 1, the page\'s choice)' },
      { target: INPUT, props: ['cursor'], golden: 'disabled, the arrow (apple.com\'s disabled tiles\' labels default)' },
    ],
  },
  {
    name: 'disabled checked',
    html: ex('checked').replace(' checked', ' checked disabled'),
    widths: [1200],
    checks: [
      { target: LABEL, props: ['opacity', 'border-top-color', 'border-top-width'], golden: 'chosen and disabled, dimmed, its 2px edge the rest one\'s gray: the disabled rule after the checked one, as apple.com\'s are ordered' },
    ],
  },
  {
    name: 'pointer',
    html: ex('option'),
    widths: [1200],
    checks: [
      { target: '.cui-option', props: ['position'], golden: 'what its input is placed against' },
      { target: INPUT, props: ['cursor'], golden: 'the pointer over a tile, apple.com\'s (its enabled tiles\' labels and inputs pointer, read 2026-10-31); the input over the tile takes it' },
      { target: LABEL, props: ['hit'], golden: 'the radio over the whole tile: the pointer finds it at the tile\'s corners and middle, so it takes the click' },
    ],
  },
];
