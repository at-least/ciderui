import { examples } from '../../../test/harness.mjs';

const ex = examples(import.meta.url);
const LINK = ['color', 'text-decoration-line', 'text-decoration-color'];
export const cases = [
  {
    name: 'inline',
    html: ex('inline'),
    checks: [
      { target: '.cui-link', props: ['font-size', 'line-height', 'font-weight', ...LINK], apple: ['iphone', 'link.inline'] },
      { target: '.cui-link', state: 'hover', props: LINK, apple: ['iphone', 'link.inline', 'hover'] },
      { target: '.cui-link', state: 'active', props: LINK, apple: ['iphone', 'link.inline', 'active'] },
    ],
  },
  {
    name: 'inline-dark',
    html: `<section class="cui-dark">${ex('inline')}</section>`,
    checks: [
      { target: '.cui-link', props: LINK, apple: ['product', 'link.inline'] },
      { target: '.cui-link', state: 'hover', props: LINK, apple: ['product', 'link.inline', 'hover'] },
      { target: '.cui-link', state: 'active', props: LINK, apple: ['product', 'link.inline', 'active'] },
    ],
  },
  {
    name: 'inline link no touch target',
    html: ex('inline'),
    touch: true,
    widths: [390],
    checks: [
      { target: '.cui-link', pseudo: '::after', props: ['content'], golden: 'a link in running text no ::after on a touch screen (text, exempt from the HIG\'s control sizes)' },
    ],
  },
];
