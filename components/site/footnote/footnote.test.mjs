import { examples } from '../../../test/harness.mjs';

const ex = examples(import.meta.url);

export const cases = [
  {
    name: 'footnote',
    html: ex('footnote'),
    checks: [
      { target: '.cui-footnote', props: ['font-feature-settings', 'vertical-align', 'position', 'top'], apple: ['airpods', 'footnote'] },
      { target: '.cui-footnote a', props: ['font-feature-settings', 'color', 'text-decoration-line', 'vertical-align'], apple: ['airpods', 'footnote.link'] },
      { target: '.cui-footnote', props: ['font-size'], golden: 'the text\'s size, 17px here (apple.com\'s marks take their text\'s: product:footnote, airpods:footnote)' },
      { target: '.cui-footnote', props: ['letter-spacing'], golden: 'the text\'s letter-spacing, inherited' },
    ],
  },
  {
    name: 'dark',
    html: `<div class="cui-dark">${ex('footnote')}</div>`,
    checks: [
      { target: '.cui-footnote a', props: ['color', 'text-decoration-line'], golden: 'the dark text\'s color, underlined (apple.com\'s on its dark page, #f5f5f7: product:footnote.link)' },
    ],
  },
];
