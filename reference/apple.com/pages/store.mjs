// The store's home: the store's own type scale (its t- classes, added to the
// page), the buttons of its shared button styles (common-css), the header, the
// quick links, the product nav, a shelf's header, its cards (a light one with
// a tag, a dark one, one with a plain eyebrow; their parts and hover), the
// shelf's paddles.
import { TYPE, BOX, MOTION, FOCUS, PAGE, TEXT, LINK } from './common.mjs';
const T = ['headline-super', 'headline-elevated', 'headline', 'headline-reduced', 'headline-standalone', 'eyebrow-super',
  'eyebrow-elevated', 'eyebrow', 'eyebrow-reduced', 'intro-elevated', 'intro', 'body', 'body-tight', 'body-reduced',
  'body-reduced-tight', 'caption', 'label', 'callout', 'tout', 'quote', 'quote-reduced', 'manifesto'];
const BUTTONS = ['', 'button-secondary', 'button-neutral', 'button-secondary-neutral', 'button-super', 'button-elevated',
  'button-reduced', 'button-block', 'button-light', 'button-dark'];
const STATES = ['hover', 'active', 'focus-visible'];
// a card by its classes: light with a tag and its picture, dark with a tag,
// light with a plain eyebrow
const LIGHT = '.rf-ccard-40.rf-card-msgtag-orange:not(.rf-ccard-darkbg):not(.rf-ccard-withvideo)';
const DARK = '.rf-ccard-40.rf-ccard-darkbg.rf-card-msgtag-orange:not(.rf-ccard-withvideo)';
const PLAIN = '.rf-ccard:not([class*="msgtag"]):not(.rf-ccard-darkbg):not(.rf-ccard-withvideo)';
// the shelf of those cards
const SHELF = '.rs-cardsshelf:has(.rf-ccard-40)';
export default {
  url: 'https://www.apple.com/store',
  probes: [
    { name: 'type.none', inject: '<p>Sample text</p>', into: '.rs-page-content, main, body', props: TYPE },
    ...T.map(c => ({ name: `type.t-${c}`, inject: `<p class="t-${c}">Sample text</p>`, into: '.rs-page-content, main, body', props: TYPE })),
    ...BUTTONS.map(c => ({ name: `button${c ? '.' + c.replace('button-', '') : ''}`, inject: `<a href="#" class="button ${c}">Buy</a>`, into: '.rs-page-content, main, body', props: [...TYPE, ...BOX, ...TEXT, 'width', ...FOCUS, ...MOTION], box: true, states: STATES })),
    // disabled: a <button> with the attribute, the shared styles' own
    ...['', 'button-secondary', 'button-neutral'].map(c => ({ name: `button${c ? '.' + c.replace('button-', '') : ''}.disabled`, inject: `<button class="button ${c}" disabled>Buy</button>`, into: '.rs-page-content, main, body', props: [...TYPE, ...BOX, ...TEXT], box: true, states: ['hover'] })),
    { name: 'page.body', select: 'body', props: PAGE },
    { name: 'page.strong', select: 'main strong, main b, .rs-page-content strong, body strong', props: TYPE },
    { name: 'header', select: 'h1.rs-shop-header', props: [...TYPE, 'margin-top', 'margin-bottom'] },
    { name: 'header.tagline', select: '.rs-shop-header + *, .rs-shop-subheader', props: TYPE },
    { name: 'quicklink', select: '.rs-quicklinks-button', props: [...TYPE, ...BOX], box: true, states: ['hover'] },
    { name: 'productnav.title', select: '.rf-productnav-card-title', props: TYPE },
    { name: 'productnav.image', select: '.rf-productnav-card-image img', props: ['width', 'height'], box: true },
    // a shelf of cards, the one holding the 400px cards: its header (its
    // headline and the secondary text after it), its scroller (in a crop that
    // hides its scrollbar), its items, the room around them
    { name: 'shelf', select: SHELF, props: ['padding-top', 'padding-bottom', 'position', 'background-color'], box: true },
    { name: 'shelf.header', select: `${SHELF} .rs-cards-shelf-header`, props: ['margin-left', 'margin-right', 'padding-top', 'padding-bottom'], box: true, position: true },
    { name: 'shelf.mainheader', select: `${SHELF} .rs-cards-shelf-mainheader`, props: [...TYPE, 'display', 'margin-top', 'margin-bottom'] },
    { name: 'shelf.subheader', select: `${SHELF} .rs-cards-shelf-secondaryheader`, props: [...TYPE, 'display'] },
    { name: 'shelf.crop', select: `${SHELF} .rf-cards-scroller-crop`, props: ['overflow-x', 'overflow-y', 'padding-bottom', 'height'], box: true },
    { name: 'shelf.scroller', select: `${SHELF} .rf-cards-scroller-content`, props: ['overflow-x', 'overflow-y', 'scroll-snap-type', 'overscroll-behavior-x'], box: true },
    { name: 'shelf.platter', select: `${SHELF} .rf-cards-scroller-platter`, props: ['display', 'padding-top', 'padding-bottom'], box: true },
    { name: 'shelf.item', select: `${SHELF} [data-core-scroller-item]`, props: ['scroll-snap-align', 'scroll-snap-stop'], box: true },
    { name: 'shelf.card', select: `${SHELF} .rf-cards-scroller-itemview`, props: ['margin-right', 'transform'], box: true, position: true },
    { name: 'shelf.card-second', select: `${SHELF} [data-core-scroller-item]:nth-child(2) .rf-cards-scroller-itemview`, props: ['transform'], box: true, position: true },
    { name: 'shelf.card-last', select: `${SHELF} [data-core-scroller-item]:last-child .rf-cards-scroller-itemview`, props: ['margin-right', 'padding-right', 'transform'], box: true },
    // the cards: a light one with a tag (rf-card-msgtag-orange: "NEW"), its
    // picture filling it, its info over the picture; a dark one
    // (rf-ccard-darkbg); a light one with a plain eyebrow ("APPLE
    // SPECIALIST"); a header first in its info (no eyebrow above it); a dark
    // card's plain eyebrow on a copy the probe adds, apple.com's classes (no
    // dark card on the page has one)
    { name: 'card', select: `${LIGHT} .rf-ccard-content`, props: [...BOX, 'transform', 'overflow', 'display', 'flex-direction', 'position', 'color', 'cursor', ...MOTION], box: true, states: ['hover', 'focus-within'] },
    { name: 'card.dark', select: `${DARK} .rf-ccard-content`, props: [...BOX, 'transform', 'overflow'], box: true, states: ['hover'] },
    { name: 'card.media', select: `${LIGHT} .rf-ccard-img-full`, props: ['display', 'object-fit', 'position'], box: true, widths: [1600, 1200, 900] },
    { name: 'card.info', select: `${LIGHT} .rf-ccard-content-info`, props: ['padding-top', 'padding-right', 'padding-bottom', 'padding-left', 'position'], box: true, position: true },
    { name: 'card.tag', select: `${LIGHT} .rf-ccard-content-header-eyebrow`, props: [...TYPE, 'margin-top', 'margin-bottom', 'padding-top', 'padding-bottom', 'text-transform'] },
    { name: 'card.eyebrow', select: `${PLAIN} .rf-ccard-content-header-eyebrow`, props: [...TYPE, 'margin-top', 'margin-bottom', 'padding-top', 'padding-bottom', 'text-transform'] },
    { name: 'card.header', select: `${LIGHT} .rf-ccard-content-header`, props: [...TYPE, 'margin-top', 'margin-bottom', 'padding-top', 'padding-bottom'] },
    { name: 'card.header-first', select: '.rf-ccard-content-info > .rf-ccard-content-header:first-child', props: [...TYPE, 'padding-top'] },
    { name: 'card.link', select: `${LIGHT} .rf-ccard-content-headerlink`, props: [...TYPE, 'display', ...LINK, 'cursor'], states: ['hover'] },
    { name: 'card.desc', select: `${LIGHT} .rf-ccard-content-desc`, props: [...TYPE, 'margin-top', 'padding-top', 'padding-bottom'] },
    { name: 'card.description', select: `${LIGHT} .rf-ccard-content-desccontent`, props: [...TYPE, 'display', 'padding-top', 'padding-bottom'] },
    { name: 'card.price', select: `${LIGHT} .rf-ccard-content-descprice`, props: TYPE },
    { name: 'card.dark-tag', select: `${DARK} .rf-ccard-content-header-eyebrow`, props: ['color'] },
    { name: 'card.dark-eyebrow', inject: '<div class="rf-ccard rf-ccard-40 rf-ccard-darkbg rf-card-bgblack"><div class="rf-ccard-content"><div class="rf-ccard-content-info"><h3 class="rf-ccard-content-header-eyebrow">Apple Specialist</h3></div></div></div>', into: '.rs-page-content, main, body', select: '.rf-ccard-content-header-eyebrow', props: ['color'] },
    { name: 'card.dark-header', select: `${DARK} .rf-ccard-content-header`, props: ['color'] },
    { name: 'card.dark-link', select: `${DARK} .rf-ccard-content-headerlink`, props: ['color'] },
    { name: 'card.dark-desc', select: `${DARK} .rf-ccard-content-desc`, props: ['color'] },
    { name: 'card.dark-price', select: `${DARK} .rf-ccard-content-descprice`, props: ['color'] },
    // the shelf's paddles, at its sides over its scroller: the next one (the
    // previous disabled at the start), its place and motion; the previous
    // hidden
    { name: 'shelf.paddle', select: `${SHELF} .paddlenav-arrow-next`, props: [...BOX, 'width', 'height', 'color', 'opacity', 'transform', 'position', 'top', 'right', ...MOTION], box: true, position: true, states: ['hover', 'active'] },
    { name: 'shelf.paddle-previous', select: `${SHELF} .paddlenav-arrow-previous`, props: ['opacity', 'transform', 'position', 'left', 'top', 'margin-top', ...MOTION], box: true, position: true },
    { name: 'shelf.paddles', select: `${SHELF} .rf-cards-scroller`, props: ['position'], box: true, position: true },
  ],
};
