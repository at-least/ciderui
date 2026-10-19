// A buy page, iPhone 18 Pro's: the title, an option (a form selector, its
// parts) at rest and, as a copy added beside it, checked, a color swatch the same way, an
// accordion's question, a price, a button link.
import { TYPE, BOX, MOTION, PAGE, FLEX } from './common.mjs';
const OPTION = `<div class="rc-dimension-selector-row form-selector"><input class="form-selector-input" id="cui-option" type="radio" name="cui-option" checked><label for="cui-option" class="form-selector-label"><span class="row"><span class="column form-selector-left-col rf-bfe-selector-left-col"><span class="form-selector-title">iPhone 18 Pro <span class="form-label-small">6.3-inch display</span></span></span></span></label></div>`;
const SWATCH = `<li class="colornav-item"><input type="radio" class="colornav-value rc-dimension-colornav-input rf-bfe-product-dimension-colornav-input" id="cui-swatch" name="cui-swatch" checked><label for="cui-swatch" class="colornav-link rc-dimension-colornav-link rf-bfe-product-dimension-colornav-label"><img alt="" class="colornav-swatch" src="data:image/gif;base64,R0lGODlhAQABAAAAACw="><span class="colornav-label">Burgundy</span></label></li>`;
export default {
  url: 'https://www.apple.com/shop/buy-iphone/iphone-18-pro',
  wait: 2500,
  probes: [
    { name: 'page.body', select: 'body', props: PAGE },
    { name: 'title', select: 'h1', props: TYPE },
    { name: 'option', select: '.form-selector:not([data-cui-injected]) .form-selector-label', props: [...TYPE, ...BOX, ...MOTION], box: true, states: ['hover'] },
    { name: 'option.focus', select: '.form-selector:not([data-cui-injected]) .form-selector-label', stateOn: '.form-selector:not([data-cui-injected]) .form-selector-input', props: [...BOX, 'outline-color', 'outline-style', 'outline-width', 'outline-offset'], states: ['focus-visible', 'focus'] },
    { name: 'option.title', select: '.form-selector-title', props: TYPE },
    { name: 'option.small', select: '.form-label-small', props: TYPE },
    // its parts: the radio over the tile, the row, its columns (the title's
    // centered, the price's at the top), the title's detail, the price
    { name: 'option.box', select: '.form-selector:not([data-cui-injected])', props: ['display', 'position', 'margin-top'], box: true },
    { name: 'option.input', select: '.form-selector:not([data-cui-injected]) .form-selector-input', props: ['position', 'top', 'left', 'opacity', 'margin-top', 'margin-left', 'display'], box: true },
    { name: 'option.label', select: '.form-selector:not([data-cui-injected]) .form-selector-label', props: [...FLEX, 'text-align'], box: true },
    { name: 'option.left', select: '.form-selector:not([data-cui-injected]) .form-selector-left-col', props: ['align-self', 'padding-right', 'text-align'], box: true, position: true },
    { name: 'option.detail', select: '.form-selector:not([data-cui-injected]) .form-selector-left-col .form-label-small', props: [...TYPE, 'display', 'margin-top'] },
    { name: 'option.right', select: '.form-selector:not([data-cui-injected]) .form-selector-right-col', props: ['align-self', 'text-align'], box: true, position: true },
    { name: 'option.price', select: '.form-selector:not([data-cui-injected]) .form-selector-right-col .form-label-small', props: [...TYPE, 'display', 'margin-top', 'text-align'], box: true, position: true },
    { name: 'option.checked', inject: OPTION, into: '.rc-dimension-selector-group', select: '.form-selector-label', props: [...TYPE, ...BOX], box: true },
    { name: 'swatch', select: '.colornav-link:not([data-cui-injected] *) .colornav-swatch', props: [...BOX, 'width', 'height'], box: true },
    { name: 'swatch.link', select: '.colornav-link', props: [...BOX, 'width', 'height'], box: true, states: ['hover'] },
    { name: 'swatch.checked', inject: SWATCH, into: '.colornav-items, .rc-dimension-colornav ul, .rc-dimension-colornav', select: '.colornav-swatch', props: [...BOX, 'width', 'height', 'outline-color', 'outline-width', 'outline-offset'], box: true },
    { name: 'swatch.checked-link', inject: SWATCH, into: '.colornav-items, .rc-dimension-colornav ul, .rc-dimension-colornav', select: '.colornav-link', props: [...BOX, 'width', 'height'], box: true },
    { name: 'swatch.label', select: '.colornav-label', props: [...TYPE, 'display'] },
    // the swatches' list and an item's parts: the radio over the ring, the
    // color's name (hidden but read), the ring focused (on the copy)
    { name: 'swatch.list', select: '.colornav-items', props: ['display', 'flex-wrap', 'padding-top', 'padding-bottom', 'margin-left', 'margin-right'], box: true },
    { name: 'swatch.item', select: '.colornav-item:not([data-cui-injected])', props: ['margin-top', 'margin-right', 'margin-bottom', 'margin-left'], box: true },
    { name: 'swatch.input', select: '.colornav-item:not([data-cui-injected]) .colornav-value', props: ['position', 'top', 'left', 'width', 'height', 'opacity', 'margin-top'], box: true },
    { name: 'swatch.name', select: '.colornav-item:not([data-cui-injected]) .colornav-label', props: ['position', 'width', 'height', 'clip', 'clip-path', 'overflow'] },
    { name: 'swatch.focus', inject: SWATCH, into: '.colornav-items, .rc-dimension-colornav ul, .rc-dimension-colornav', select: '.colornav-link', stateOn: '#cui-swatch', props: ['outline-style', 'outline-width', 'outline-color', 'outline-offset', 'border-top-color'], states: ['focus-visible'] },
    { name: 'accordion.item', select: '.rc-accordion-item', props: BOX, box: true },
    { name: 'accordion.button', select: '.rc-accordion-button', props: [...TYPE, ...BOX, 'width'], box: true, states: ['hover', 'focus-visible'] },
    { name: 'accordion.title', select: '.rc-accordion-title', props: TYPE },
    { name: 'accordion.icon', select: '.rc-accordion-button .icon-plus, .rc-accordion-button .as-accordion-plusicon', props: ['font-family', 'font-size', 'color', 'width', 'height'], pseudo: '::after' },
    { name: 'accordion.content', select: '.rc-accordion-box-content', props: [...TYPE, ...BOX] },
    { name: 'price', select: '.price-point-fullPrice', props: TYPE },
    // the page's own price, under its title: its amounts kept whole
    { name: 'price.header', select: '.rf-bfe-header-price .rc-prices-currentprice', props: [...TYPE, 'display'] },
    { name: 'price.amount', select: '.rf-bfe-header-price .nowrap', props: [...TYPE, 'white-space', 'display'] },
    { name: 'buttonlink', select: '.as-buttonlink', props: [...TYPE, ...BOX], states: ['hover'] },
  ],
};
