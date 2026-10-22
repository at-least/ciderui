// DocC's other asides, which DocC's stylesheets style and no page read
// carries: the crash report article's first note, its class (and its
// aria-label) set to warning, tip, deprecated and experiment in turn, as DocC
// writes each (its label's text stays the note's); its box and its label.
// The article's real note is docc's. docc-asides-dark reads them in the dark.
import { TYPE } from './common.mjs';
import { ASIDE } from './docc.mjs';

const kind = (k) => [
  { name: k, set: [['aside', 'class', k], ['aside', 'aria-label', k]], select: 'aside', props: ASIDE, box: true },
  { name: `${k}.label`, select: 'aside p.label', props: [...TYPE, 'margin-bottom'] },
];
export default {
  url: 'https://developer.apple.com/documentation/xcode/examining-the-fields-in-a-crash-report',
  wait: 4000,
  probes: ['warning', 'tip', 'deprecated', 'experiment'].flatMap(kind),
};
