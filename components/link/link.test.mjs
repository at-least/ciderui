// The link against apple.com: in running text, light and dark, its own `a`
// (iphone:link.inline, product:link.inline); the more link at 17px and 14px
// and dark, its copy underlined hovered, its chevron's place and color
// (iphone:link.more*, product:link.more*); golden: the chevron's box, from
// apple.com's glyph's numbers (reference/icons/); a touch screen's hover is
// rest.

import { examples } from '../../test/harness.mjs';

const ex = examples(import.meta.url);
const LINK = ['color', 'text-decoration-line', 'text-decoration-color'];
const more = (page, glyph) => [
  // the size is the case's text's, built as apple.com's is around its link
  { target: '.cui-link', props: ['font-size', 'color', 'text-decoration-line'], apple: [page, 'link.more'] },
  { target: '.cui-link-text', props: ['text-decoration-line'], apple: [page, 'link.more-copy'] },
  { target: '.cui-link-text', on: '.cui-link', state: 'hover', props: ['text-decoration-line'], apple: [page, 'link.more-copy', 'hover'] },
  { target: '.cui-link-icon', props: ['margin-left', 'color'], apple: [page, glyph], as: { 'margin-left': 'padding-left' } },
];

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
    // the iPhone page's more link is 17px from 1069px, 14px below
    name: 'more',
    html: `<p>${ex('more')}</p>`,
    widths: [1600, 1200],
    checks: [
      ...more('iphone', 'link.more-glyph'),
      { target: '.cui-link-icon', props: ['box-width', 'box-height', 'vertical-align'], golden: 'the chevron\'s box: apple.com\'s glyph\'s advance, 0.4419em, by 0.54em, 0.01em below the baseline (reference/icons/)' },
      { target: '.cui-link-icon', props: ['display'], golden: 'inline, on its text\'s line: Tailwind\'s preflight draws an svg as a block' },
      { target: 'p', props: ['box-height'], golden: 'one line, the chevron after the text: the paragraph\'s 25px line' },
    ],
  },
  {
    name: 'more-small',
    html: `<p class="cui-body-reduced">${ex('more')}</p>`,
    widths: [900, 390],
    checks: more('iphone', 'link.more-glyph'),
  },
  {
    // the product page's more links are 14px at every width
    name: 'more-dark',
    html: `<section class="cui-dark"><p class="cui-body-reduced">${ex('more')}</p></section>`,
    checks: [
      { target: '.cui-link', props: ['font-size', 'color', 'text-decoration-line'], apple: ['product', 'link.more'] },
      { target: '.cui-link-icon', props: ['margin-left', 'color'], apple: ['product', 'link.more-icon'], as: { 'margin-left': 'padding-left' } },
    ],
  },
  {
    name: 'touch',
    html: `<p>${ex('more')}</p>`,
    touch: true,
    widths: [390],
    checks: [{ target: '.cui-link-text', on: '.cui-link', state: 'hover', props: ['text-decoration-line'], rest: 'a touch screen keeps no hover (rule 5), where apple.com\'s more link underlines (touch-iphone:link.more-copy)' }],
  },
];
