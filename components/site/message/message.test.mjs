import { examples } from '../../../test/harness.mjs';

const ex = examples(import.meta.url);
const M = '.cui-message';

export const cases = [
  {
    name: 'message',
    html: ex('message'),
    widths: [1200, 390],
    checks: [
      { target: `${M}:not([data-state])`, props: ['color', 'font-size', 'line-height'], golden: 'what the request did, in the text\'s color at the body\'s size' },
      { target: `${M}[data-state="error"]`, props: ['color'], golden: 'a request that failed in the error color, its words saying what to do next' },
      { target: `${M}:empty`, props: ['display', 'position'], golden: 'empty, it takes no room (no margin or gap of the page\'s is left behind) and stays in the accessibility tree, so the line its script fills is announced' },
    ],
  },
  {
    name: 'message empty in a gap',
    html: '<div style="display: flex; flex-direction: column; gap: 16px"><p id="a" style="height: 20px; margin: 0">A</p><p class="cui-message mt-4" role="status"></p><p id="b" style="height: 20px; margin: 0">B</p></div>',
    widths: [1200],
    checks: [
      { target: '#b', props: ['box-top'], golden: 'one gap between the lines around an empty message, its margin and its own gap gone: 36px down, 20px and the 16px gap' },
    ],
  },
];
