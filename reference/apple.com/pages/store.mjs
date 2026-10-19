// The store's home: the store's own type scale (its t- classes, added to the
// page), the buttons of its shared button styles (common-css), the header, the
// quick links, the product nav, a shelf's header, a card and its hover, the
// shelf's paddles.
import { TYPE, BOX, MOTION, FOCUS } from './common.mjs';
const T = ['headline-super', 'headline-elevated', 'headline', 'headline-reduced', 'headline-standalone', 'eyebrow-super',
  'eyebrow-elevated', 'eyebrow', 'eyebrow-reduced', 'intro-elevated', 'intro', 'body', 'body-tight', 'body-reduced',
  'body-reduced-tight', 'caption', 'label', 'callout', 'tout', 'quote', 'quote-reduced', 'manifesto'];
const BUTTONS = ['', 'button-secondary', 'button-neutral', 'button-secondary-neutral', 'button-super', 'button-elevated',
  'button-reduced', 'button-block', 'button-light', 'button-dark'];
const STATES = ['hover', 'active', 'focus-visible'];
export default {
  url: 'https://www.apple.com/store',
  probes: [
    { name: 'type.none', inject: '<p>Sample text</p>', into: '.rs-page-content, main, body', props: TYPE },
    ...T.map(c => ({ name: `type.t-${c}`, inject: `<p class="t-${c}">Sample text</p>`, into: '.rs-page-content, main, body', props: TYPE })),
    ...BUTTONS.map(c => ({ name: `button${c ? '.' + c.replace('button-', '') : ''}`, inject: `<a href="#" class="button ${c}">Buy</a>`, into: '.rs-page-content, main, body', props: [...TYPE, ...BOX, 'display', 'width', ...FOCUS, ...MOTION], box: true, states: STATES })),
    { name: 'page.body', select: 'body', props: [...TYPE, 'background-color'] },
    { name: 'header', select: 'h1.rs-shop-header', props: [...TYPE, 'margin-top', 'margin-bottom'] },
    { name: 'header.tagline', select: '.rs-shop-header + *, .rs-shop-subheader', props: TYPE },
    { name: 'quicklink', select: '.rs-quicklinks-button', props: [...TYPE, ...BOX], box: true, states: ['hover'] },
    { name: 'productnav.title', select: '.rf-productnav-card-title', props: TYPE },
    { name: 'productnav.image', select: '.rf-productnav-card-image img', props: ['width', 'height'], box: true },
    { name: 'shelf.header', select: '.rs-cards-shelf-mainheader', props: TYPE },
    { name: 'shelf.subheader', select: '.rs-cards-shelf-secondaryheader', props: TYPE },
    { name: 'card', select: '.rf-ccard:not(.rf-ccard-withvideo) .rf-ccard-content', props: [...BOX, 'transform', 'overflow', ...MOTION], box: true, states: ['hover'] },
    { name: 'card.eyebrow', select: '.rf-ccard-content-header-eyebrow', props: TYPE },
    { name: 'card.header', select: '.rf-ccard-content-header', props: TYPE },
    { name: 'card.desc', select: '.rf-ccard-content-desc', props: TYPE },
    { name: 'shelf.paddle', select: '.rf-cards-scroller .paddlenav-arrow-next', props: [...BOX, 'width', 'height', 'color', 'opacity'], box: true, states: ['hover', 'active'] },
  ],
};
