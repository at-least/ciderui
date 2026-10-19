// The gift card's buy page, for the store's form parts a bag is not needed
// to reach (its checkout's are not: a headless browser's Add to Bag is
// refused, 2026-10-20): its Select button disabled before its choices, its
// delivery fields disabled before the design and the amount; its sections; the choices that enable the fields
// (picked once, at the first width: Email, the first design, Other Amount,
// then $50); the amount's field with its currency before it, while Other
// Amount is chosen; a plain field enabled; an
// empty required field the focus has left (clicked, then the focus taken
// off: nothing typed), its error and its message. Read only: nothing typed,
// nothing sent.
import { TYPE, BOX, MOTION, FOCUS } from './common.mjs';

const FIELDS = '.rf-giftcard-delivery-section-fields';
const NAME = `${FIELDS} .form-textbox:first-child`;
const EMAIL = `${FIELDS} .form-textbox:nth-child(2)`;
const OTHER = '.rf-choose-amount-dimensions > div:nth-child(4) .form-selector-input';
const PREFIX = '.rf-custom-input';
const AMOUNT = '.rf-choose-amount-dimensions > div:nth-child(2) .form-selector-input';
const EDGE = ['border-top-width', 'border-top-style', 'border-top-color', 'border-right-width', 'border-right-style', 'border-right-color', 'border-top-left-radius', 'border-bottom-left-radius', 'border-top-right-radius'];

export default {
  url: 'https://www.apple.com/shop/buy-giftcard/giftcard',
  wait: 2500,
  probes: [
    // before a choice: its button disabled (the first width alone, as the
    // fields below: the choices after them enable them)
    { name: 'button.disabled', select: 'button.button-block[disabled]', props: ['background-color', 'color', 'opacity', 'cursor', 'border-top-color'], widths: [1600] },
    // a section, its rule and its title
    { name: 'section', select: '.rf-giftcard-selection-group', props: [...BOX, 'border-top-width', 'border-top-style', 'border-top-color'], box: true },
    { name: 'section.title', select: '.rf-choose-amount-section-title', props: [...TYPE, 'margin-top', 'margin-bottom'] },
    // the choices that enable the fields, each clicked once
    { name: 'choose.type', click: 'input[name=dimensionGiftCardType][value=electronic]', select: 'input[name=dimensionGiftCardType][value=electronic]', props: ['opacity'] },
    // its delivery fields shown, disabled until a design and an amount are
    // chosen
    { name: 'field.disabled', select: `${NAME} .form-textbox-input`, props: [...TYPE, 'background-color', 'border-top-color', 'border-top-width', 'opacity', 'cursor', 'height'], widths: [1600] },
    { name: 'label.disabled', select: `${NAME} .form-textbox-label`, props: [...TYPE, 'opacity'], widths: [1600] },
    { name: 'choose.design', click: 'input[name=dimensionGiftCardTheme]', select: 'input[name=dimensionGiftCardTheme]', props: ['opacity'] },
    { name: 'choose.other', click: OTHER, select: OTHER, props: ['opacity'] },
    // the amount's field, its currency before it (enabled while Other
    // Amount is chosen: at the first width alone, $50 chosen after it)
    { name: 'prefix', select: PREFIX, props: ['position', 'display', 'height'], box: true, position: true, widths: [1600] },
    { name: 'prefix.currency', select: `${PREFIX} .form-textbox-countrycode`, props: [...TYPE, ...BOX, ...EDGE, 'position', 'top', 'left', 'width', 'height', 'display', 'align-items', 'justify-content', 'background-color', 'z-index'], box: true, position: true, widths: [1600] },
    { name: 'prefix.input', select: `${PREFIX} .form-textbox-input`, props: [...TYPE, ...BOX, ...EDGE, 'height', 'background-color'], box: true, position: true, states: ['focus'], widths: [1600] },
    { name: 'prefix.label', select: `${PREFIX} .form-textbox-label`, props: [...TYPE, 'position', 'top', 'left', 'padding-left'], box: true, position: true, widths: [1600] },
    // an amount, which enables the delivery fields
    { name: 'choose.amount', click: AMOUNT, select: AMOUNT, props: ['opacity'] },
    // a plain field, enabled
    { name: 'field', select: `${NAME} .form-textbox-input`, props: [...TYPE, ...BOX, 'height', 'background-color', 'border-top-color', 'border-top-width'], box: true },
    // an empty required field the focus left: clicked, then the focus taken off
    { name: 'error.focus', click: `${EMAIL} .form-textbox-input`, select: `${EMAIL} .form-textbox-input`, props: ['opacity'] },
    { name: 'error', blur: true, select: `${EMAIL} .form-textbox-input`, props: [...TYPE, ...BOX, 'height', 'background-color', 'border-top-color', 'border-top-width', ...FOCUS, ...MOTION], box: true, position: true, states: ['focus'] },
    { name: 'error.label', select: `${EMAIL} .form-textbox-label`, props: [...TYPE, 'top', 'left'], box: true, position: true },
    { name: 'error.wrapper', select: `${EMAIL} .form-message-wrapper`, props: [...BOX, 'display', 'align-items'], box: true, position: true },
    { name: 'error.message', select: `${EMAIL} .form-message`, props: [...TYPE, ...BOX, 'display', 'align-items'], box: true, position: true },
    // its icon, the wrapper's ::before: an SF Pro Icons glyph (U+F312)
    { name: 'error.icon', select: `${EMAIL} .form-message-wrapper`, pseudo: '::before', props: ['content', 'font-family', 'font-size', 'line-height', 'width', 'height', 'padding-right', 'margin-top', 'color', 'display', 'position', 'top'] },
  ],
};
