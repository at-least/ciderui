import { examples, SIZE } from '../../../test/harness.mjs';

const ex = examples(import.meta.url);
const L = 'component-Link-variant-colorRole-brand-effect';
const INLINE = '.cui-console-link:not([data-variant]):not([aria-disabled])';
const LABEL = '[data-variant="label"]';
const TEXT = '[data-variant="text"]';

export const cases = [
  {
    name: 'link',
    html: ex('link'),
    widths: [1200, 390],
    checks: [
      { target: INLINE, props: ['color', 'text-decoration-line', 'cursor'], asc: ['analytics-page', 'section.link'] },
      { target: INLINE, props: ['color', 'text-decoration-line', 'cursor'], state: 'hover', asc: ['hover-link', 'section.link'] },
      { target: INLINE, props: ['color'], tb: [`${L}-default-variant-inline-color`] },
      { target: LABEL, props: ['color'], tb: [`${L}-default-variant-label-color`] },
      { target: LABEL, props: ['color'], state: 'hover', tb: [`${L}-hover-variant-label-color`] },
      { target: LABEL, props: ['text-decoration-line'], state: 'hover', golden: 'a name that is a link turns blue hovered, not underlined' },
      { target: '[aria-disabled="true"]', props: ['color'], tb: [`${L}-disabled-variant-inline-color`] },
      { target: '[aria-disabled="true"]', props: ['cursor', 'text-decoration-line'], state: 'hover', golden: 'no pointer, no underline over a disabled link' },
      { target: INLINE, pseudo: '::after', props: ['content'], golden: 'with a mouse, no target drawn: none' },
      { target: TEXT, props: ['color'], tb: [`${L}-default-variant-inline-color`] },
      { target: TEXT, props: ['text-decoration-line'], golden: 'a link inside a sentence underlined at rest: told apart from the words around it by more than its color (WCAG 1.4.1; crashcart\'s links in a secondary caption measured 1.03:1 against it)' },
      { target: TEXT, props: ['text-decoration-line'], state: 'hover', golden: 'hovered, still underlined' },
      { target: TEXT, props: SIZE, golden: 'the type of the text around it' },
    ],
  },
  {
    name: 'link touch',
    html: ex('link'),
    touch: true,
    widths: [390],
    checks: [
      { target: INLINE, pseudo: '::after', props: ['width', 'height'], golden: 'a touch screen\'s target 44px square at least, the HIG\'s 44pt (crashcart\'s HIG pass: its See All and Open issue links, 13px tall, were the only controls without one)' },
    ],
  },
  {
    name: 'link dark',
    html: `<div class="cui-dark">${ex('link')}</div>`,
    widths: [1200],
    checks: [
      { target: INLINE, props: ['color'], golden: 'the dark side derived (tokens.css, decision 18)' },
    ],
  },
];
