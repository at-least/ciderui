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
      { target: '.cui-section-content', props: ['box-width', 'box-left'], golden: '87.5% of the window, centered (apple.com\'s, a width no apple check reads)' },
    ],
  },
  {
    name: 'header',
    html: `<div class="cui-section-content">${ex('header')}</div>`,
    checks: [
      { target: '.cui-section-header', props: [...FLEX, 'margin-bottom'], apple: ['product', 'section.header'] },
      { target: '.cui-section-headline', props: ['font-size', 'line-height', 'font-weight'], apple: ['product', 'section.headline'] },
      { target: '.cui-section-links', props: [...FLEX.filter(p => p !== 'align-items' && p !== 'justify-content'), 'font-size', 'line-height'], apple: ['product', 'section.header-links'] },
    ],
  },
  {
    name: 'section joined',
    html: band({}) + band({}),
    widths: [1200],
    checks: [
      { target: '.cui-section + .cui-section', props: PAD, golden: 'each section keeps its own padding beside another: a page joining two takes one away with a utility (pt-0)' },
    ],
  },
];
