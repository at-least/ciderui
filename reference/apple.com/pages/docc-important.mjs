// DocC's important aside, on an article that has one (common crashes): its
// box, its label and its text. docc-important-dark reads it in the dark.
import { TYPE } from './common.mjs';
import { ASIDE } from './docc.mjs';

export default {
  url: 'https://developer.apple.com/documentation/xcode/identifying-the-cause-of-common-crashes',
  wait: 4000,
  probes: [
    { name: 'important', select: 'aside.important', props: ASIDE, box: true },
    { name: 'important.label', select: 'aside.important p.label', props: [...TYPE, 'margin-bottom'] },
    { name: 'important.text', select: 'aside.important p:not(.label)', props: TYPE },
  ],
};
