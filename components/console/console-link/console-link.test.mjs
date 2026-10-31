// The console's link against App Store Connect's: the analytics overview's
// See All link, at rest and hovered (analytics-page:section.link,
// hover-link:section.link), and the toolkit's link tokens (brand: inline,
// label, disabled). Golden, with their why: the label variant hovered
// without its underline, the dark side.

import { examples } from '../../../test/harness.mjs';

const ex = examples(import.meta.url);
const L = 'component-Link-variant-colorRole-brand-effect';
const INLINE = '.cui-console-link:not([data-variant]):not([aria-disabled])';
const LABEL = '[data-variant="label"]';

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
    ],
  },
  {
    name: 'link touch',
    html: ex('link'),
    touch: true,
    widths: [390],
    checks: [
      { target: INLINE, props: ['text-decoration-line'], state: 'hover', rest: 'a tapped link keeps no underline (rule 5)' },
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
