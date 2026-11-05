import { examples } from '../../../test/harness.mjs';

const ex = examples(import.meta.url);
const SKIP = '.cui-skip';

export const cases = [
  {
    name: 'skip link',
    html: ex('skip'),
    widths: [1200, 390],
    checks: [
      { target: SKIP, props: ['box-width', 'box-height', 'clip-path'], golden: 'unseen at rest, a 1px box clipped to nothing: still a tab stop and in the accessibility tree' },
      { target: SKIP, state: 'focus', props: ['box-top', 'box-left', 'box-height', 'clip-path', 'position', 'z-index'], golden: 'focused, the 36px button whole, fixed 16px in from the window\'s corner, over the global nav' },
      { target: SKIP, state: 'focus', props: ['background-color', 'color'], golden: 'the default button\'s fill and label' },
    ],
  },
  {
    name: 'skip tab',
    html: `${ex('skip')}<main id="main" tabindex="-1">The page</main>`,
    widths: [1200],
    press: 'Tab',
    wait: 100,
    checks: [
      { target: SKIP, props: ['focused'], golden: 'the page\'s first tab stop' },
      { target: SKIP, props: ['box-height', 'clip-path'], golden: 'the keyboard\'s focus shows it, whole' },
    ],
  },
];
