// The section against apple.com: its rhythm, the iPhone page's
// (iphone:section.incentive, .image-accordion) and iPhone 18 Pro's
// (product:section.battery); its grounds, alternate, dark and dark
// alternate (iphone:section.incentive, product:section.battery,
// product:section.dark); its content's width (iphone:content,
// product:content); its header, headline and links (product:section.*);
// golden: the headline's letter-spacing.

import { examples } from '../../../test/harness.mjs';

const ex = examples(import.meta.url);
const PAD = ['padding-top', 'padding-bottom'];
const FLEX = ['display', 'flex-direction', 'flex-wrap', 'align-items', 'justify-content', 'row-gap', 'column-gap'];
const band = (attrs) => `<section class="cui-section${attrs.dark ? ' cui-dark' : ''}"${attrs.variant ? ` data-variant="${attrs.variant}"` : ''}${attrs.size ? ` data-size="${attrs.size}"` : ''}><div class="cui-section-content"><p>Content</p></div></section>`;

export const cases = [
  {
    name: 'rhythm',
    html: ex('section'),
    checks: [
      { target: '.cui-section', props: PAD, apple: ['iphone', 'section.incentive'] },
      { target: '.cui-section', props: PAD, apple: ['iphone', 'section.image-accordion'] },
      { target: '.cui-section', props: ['background-color'], apple: ['iphone', 'section.image-accordion'] },
    ],
  },
  {
    name: 'rhythm-lg',
    html: band({ size: 'lg' }),
    checks: [{ target: '.cui-section', props: PAD, apple: ['product', 'section.battery'] }],
  },
  {
    name: 'alt',
    html: band({ variant: 'alt' }),
    checks: [{ target: '.cui-section', props: [...PAD, 'background-color'], apple: ['iphone', 'section.incentive'] }],
  },
  {
    name: 'dark',
    html: band({ dark: true, size: 'lg' }),
    checks: [{ target: '.cui-section', props: [...PAD, 'background-color'], apple: ['product', 'section.battery'] }],
  },
  {
    // the highlights: a dark block on the alternate ground
    name: 'dark-alt',
    html: band({ dark: true, variant: 'alt', size: 'lg' }),
    checks: [{ target: '.cui-section', props: ['background-color', 'color'], apple: ['product', 'section.dark'] }],
  },
  {
    name: 'content',
    html: ex('section'),
    checks: [
      { target: '.cui-section-content', props: ['max-width', 'margin-left', 'margin-right'], apple: ['iphone', 'content'] },
      { target: '.cui-section-content', props: ['max-width', 'margin-left', 'margin-right'], apple: ['product', 'content'] },
    ],
  },
  {
    name: 'header',
    html: `<div class="cui-section-content">${ex('header')}</div>`,
    checks: [
      { target: '.cui-section-header', props: [...FLEX, 'margin-bottom'], apple: ['product', 'section.header'] },
      { target: '.cui-section-headline', props: ['font-size', 'line-height', 'font-weight'], apple: ['product', 'section.headline'] },
      { target: '.cui-section-headline', props: ['letter-spacing'], golden: 'the system font\'s fitted letter-spacing (tokens.css), not apple.com\'s' },
      { target: '.cui-section-links', props: [...FLEX.filter(p => p !== 'align-items' && p !== 'justify-content'), 'font-size', 'line-height'], apple: ['product', 'section.header-links'] },
    ],
  },
];
