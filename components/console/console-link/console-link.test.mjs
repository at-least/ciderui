import { examples } from '../../../test/harness.mjs';

const ex = examples(import.meta.url);
const LINK = '.cui-console-link';

export const cases = [
  {
    name: 'link',
    html: ex('link'),
    widths: [1200, 390],
    checks: [
      { target: LINK, props: ['color', 'text-decoration-line', 'cursor'], asc: ['analytics-page', 'section.link'] },
      { target: LINK, props: ['color', 'text-decoration-line', 'cursor'], state: 'hover', asc: ['hover-link', 'section.link'] },
      { target: LINK, pseudo: '::after', props: ['content'], golden: 'with a mouse, no target drawn: none' },
    ],
  },
  {
    name: 'link touch',
    html: ex('link'),
    touch: true,
    widths: [390],
    checks: [
      { target: LINK, pseudo: '::after', props: ['top', 'left', 'transform'], golden: 'the 44px target centered on the link, wider than it' },
      { target: LINK, pseudo: '::after', props: ['width', 'height'], golden: 'a touch screen\'s target 44px square at least, the HIG\'s 44pt (crashcart\'s HIG pass: its See All and Open issue links, 13px tall, were the only controls without one)' },
    ],
  },
  {
    name: 'link dark',
    html: `<div class="cui-dark">${ex('link')}</div>`,
    widths: [1200],
    checks: [
      { target: LINK, props: ['color'], golden: 'the dark side derived (tokens.css, decision 18)' },
    ],
  },
];
