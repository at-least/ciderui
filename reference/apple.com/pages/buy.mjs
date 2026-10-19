// A buy page, iPhone 18 Pro's: the title, an option (a form selector) at rest
// and, as a copy added beside it, checked, a color swatch the same way, an
// accordion's question, a price, a button link.
import { TYPE, BOX, MOTION } from './common.mjs';
const OPTION = `<div class="rc-dimension-selector-row form-selector"><input class="form-selector-input" id="cui-option" type="radio" name="cui-option" checked><label for="cui-option" class="form-selector-label"><span class="row"><span class="column form-selector-left-col rf-bfe-selector-left-col"><span class="form-selector-title">iPhone 18 Pro <span class="form-label-small">6.3-inch display</span></span></span></span></label></div>`;
const SWATCH = `<li class="colornav-item"><input type="radio" class="colornav-value rc-dimension-colornav-input rf-bfe-product-dimension-colornav-input" id="cui-swatch" name="cui-swatch" checked><label for="cui-swatch" class="colornav-link rc-dimension-colornav-link rf-bfe-product-dimension-colornav-label"><img alt="" class="colornav-swatch" src="data:image/gif;base64,R0lGODlhAQABAAAAACw="><span class="colornav-label">Burgundy</span></label></li>`;
export default {
  url: 'https://www.apple.com/shop/buy-iphone/iphone-18-pro',
  wait: 2500,
  probes: [
    { name: 'page.body', select: 'body', props: [...TYPE, 'background-color'] },
    { name: 'title', select: 'h1', props: TYPE },
    { name: 'option', select: '.form-selector:not([data-cui-injected]) .form-selector-label', props: [...TYPE, ...BOX, ...MOTION], box: true, states: ['hover'] },
    { name: 'option.focus', select: '.form-selector:not([data-cui-injected]) .form-selector-label', stateOn: '.form-selector:not([data-cui-injected]) .form-selector-input', props: [...BOX, 'outline-color', 'outline-style', 'outline-width', 'outline-offset'], states: ['focus-visible', 'focus'] },
    { name: 'option.title', select: '.form-selector-title', props: TYPE },
    { name: 'option.small', select: '.form-label-small', props: TYPE },
    { name: 'option.checked', inject: OPTION, into: '.rc-dimension-selector-group', select: '.form-selector-label', props: [...TYPE, ...BOX], box: true },
    { name: 'swatch', select: '.colornav-link:not([data-cui-injected] *) .colornav-swatch', props: [...BOX, 'width', 'height'], box: true },
    { name: 'swatch.link', select: '.colornav-link', props: [...BOX, 'width', 'height'], box: true, states: ['hover'] },
    { name: 'swatch.checked', inject: SWATCH, into: '.colornav-items, .rc-dimension-colornav ul, .rc-dimension-colornav', select: '.colornav-swatch', props: [...BOX, 'width', 'height', 'outline-color', 'outline-width', 'outline-offset'], box: true },
    { name: 'swatch.checked-link', inject: SWATCH, into: '.colornav-items, .rc-dimension-colornav ul, .rc-dimension-colornav', select: '.colornav-link', props: [...BOX, 'width', 'height'], box: true },
    { name: 'swatch.label', select: '.colornav-label', props: [...TYPE, 'display'] },
    { name: 'accordion.item', select: '.rc-accordion-item', props: BOX, box: true },
    { name: 'accordion.button', select: '.rc-accordion-button', props: [...TYPE, ...BOX, 'width'], box: true, states: ['hover', 'focus-visible'] },
    { name: 'accordion.title', select: '.rc-accordion-title', props: TYPE },
    { name: 'accordion.icon', select: '.rc-accordion-button .icon-plus, .rc-accordion-button .as-accordion-plusicon', props: ['font-family', 'font-size', 'color', 'width', 'height'], pseudo: '::after' },
    { name: 'accordion.content', select: '.rc-accordion-box-content', props: [...TYPE, ...BOX] },
    { name: 'price', select: '.price-point-fullPrice', props: TYPE },
    { name: 'buttonlink', select: '.as-buttonlink', props: [...TYPE, ...BOX], states: ['hover'] },
  ],
};
