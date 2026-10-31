// The console's inline code against App Store Connect's toolkit's
// monoBodyRegular text style. Golden, with their why: the font, the fill,
// the room and corners, the dark side.

import { examples } from '../../../test/harness.mjs';

const ex = examples(import.meta.url);
const T = 'typography-textStyle-monoBodyRegular';
const CODE = 'p:first-child .cui-console-code';

export const cases = [
  {
    name: 'code',
    html: ex('code'),
    widths: [1200, 390],
    checks: [
      { target: CODE, props: ['font-size', 'line-height', 'font-weight', 'letter-spacing'], tb: [`${T}-fontSize`, `${T}-lineHeight`, `${T}-fontWeight`, `${T}-letterSpacing`] },
      { target: CODE, props: ['font-family'], golden: 'ciderui\'s mono stack: the toolkit\'s "SF Mono" names a font only Apple\'s platforms have' },
      { target: CODE, props: ['color', 'background-color'], golden: 'the label\'s color on the gray badge\'s fill (the neutral fill at 0.12): the text style has no ground' },
      { target: CODE, props: ['padding-left', 'padding-right', 'border-top-left-radius'], golden: '4px in, 4px corners: the toolkit\'s smallest rounded rectangle, 8px, rounds an 18px line nearly to a capsule' },
    ],
  },
  {
    name: 'code dark',
    html: `<div class="cui-dark">${ex('code')}</div>`,
    widths: [1200],
    checks: [
      { target: CODE, props: ['color', 'background-color'], golden: 'the dark side derived (tokens.css, decision 18)' },
    ],
  },
];
