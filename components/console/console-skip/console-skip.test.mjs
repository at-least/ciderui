// The console's skip link: unseen at rest, the filled button at the
// toolkit's small page margin focused. Golden, with their why: its clip at
// rest, its place over the page.

import { examples } from '../../../test/harness.mjs';

const ex = examples(import.meta.url);
const SKIP = '.cui-console-skip';

export const cases = [
  {
    name: 'skip link',
    html: ex('skip'),
    widths: [1200, 390],
    checks: [
      { target: SKIP, props: ['box-width', 'box-height', 'clip-path'], golden: 'unseen at rest, a 1px box clipped to nothing: still a tab stop and in the accessibility tree' },
      { target: SKIP, props: ['box-top', 'box-left'], state: 'focus', tb: ['space-pageMarginSm', 'space-pageMarginSm'] },
      { target: SKIP, props: ['box-height', 'clip-path', 'position', 'z-index'], state: 'focus', golden: 'focused, the filled button whole, fixed over the page and the sidebar\'s sheet' },
      { target: SKIP, props: ['background-color', 'color'], state: 'focus', golden: 'the filled button\'s brand fill and white label' },
    ],
  },
];
