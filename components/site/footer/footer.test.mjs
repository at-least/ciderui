import { examples, SIZE } from '../../../test/harness.mjs';

const ex = examples(import.meta.url);
const WIDE = [1600, 1200, 900];

export const cases = [
  {
    name: 'footer',
    html: ex('footer'),
    checks: [
      { target: '.cui-footer', props: [...SIZE, 'color', 'background-color'], apple: ['home', 'footer'] },
      { target: '.cui-footer-content', props: ['max-width', 'padding-left', 'padding-right'], apple: ['home', 'footer.content'] },
      { target: '.cui-footer-legal', props: ['padding-top', 'padding-bottom'], apple: ['home', 'footer.footer'] },
      { target: '.cui-footer-copyright', props: [...SIZE, 'color', 'margin-top'], apple: ['home', 'footer.copyright'] },
      { target: '.cui-footer-copyright a', props: ['color', 'text-decoration-line'], apple: ['home', 'footer.legal-link'] },
      { target: '.cui-footer', props: ['box-top'], golden: 'the body\'s last child at the window\'s foot on a page shorter than the window, not under the last line of a short page' },
      { target: '.cui-footer-copyright a', state: 'hover', props: ['text-decoration-line'], golden: 'a link in the copyright underlined under the pointer, as apple.com\'s footer links' },
    ],
  },
  {
    name: 'touch target',
    html: ex('footer'),
    touch: true,
    widths: [390],
    checks: [
      { target: '.cui-footer-copyright a', pseudo: '::after', props: ['height'], golden: 'the company link\'s target 44pt on a touch screen, the HIG\'s default, past its 16px line' },
    ],
  },
  {
    name: 'dark',
    html: `<div class="cui-dark">${ex('footer')}</div>`,
    widths: WIDE,
    checks: [
      { target: '.cui-footer', props: ['color', 'background-color'], apple: ['product', 'footer'] },
    ],
  },
];
