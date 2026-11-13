import { examples } from '../../../test/harness.mjs';

const ex = examples(import.meta.url);
const M = '.cui-console-message';

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
];
