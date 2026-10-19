// The store's dropdown (its form-dropdown), on iPhone 18 Pro's buy page,
// where its trade-in asks for the model to trade: the buy page's choices
// picked once, at the first width (its size, color and capacity by their
// labels, which are over their inputs; Add a trade-in by its input, which is
// over its label), and the dropdown it shows, empty (its label the placeholder),
// hovered and focused; as a copy beside it with a model chosen (its label
// small at its top: the classes apple.com's script sets). Read only: nothing
// chosen in the page's own dropdown, nothing sent.
import { TYPE, BOX, MOTION, FOCUS } from './common.mjs';

const SECTION = '.rf-tradeupinline-form-section';
const DD = `${SECTION} .form-dropdown:not([data-cui-injected])`;
const CHOSEN = '<div class="form-dropdown"><label class="form-dropdown-label" for="cui-dropdown">Choose your model</label><select id="cui-dropdown" name="cui-dropdown" class="form-dropdown-select"><option>iPhone 17 Pro</option></select><span class="form-dropdown-chevron" aria-hidden="true"></span></div>';
const EDGE = ['border-top-width', 'border-top-style', 'border-top-color', 'border-top-left-radius'];

export default {
  url: 'https://www.apple.com/shop/buy-iphone/iphone-18-pro',
  wait: 2500,
  probes: [
    { name: 'choose.size', click: 'input[name=dimensionScreensize] + label', select: 'input[name=dimensionScreensize]', props: ['opacity'] },
    { name: 'choose.color', click: 'input[name=dimensionColor] + label', select: 'input[name=dimensionColor]', props: ['opacity'] },
    { name: 'choose.capacity', click: 'input[name=dimensionCapacity] + label', select: 'input[name=dimensionCapacity]', props: ['opacity'] },
    { name: 'choose.tradein', click: 'input[name=tradeupinline][value=tradeIn]', select: 'input[name=tradeupinline][value=tradeIn]', props: ['opacity'] },
    { name: 'dropdown', select: DD, props: ['position', 'display', 'height', 'margin-bottom', 'color-scheme'], box: true, position: true },
    { name: 'dropdown.select', select: `${DD} .form-dropdown-select`, props: [...TYPE, ...BOX, ...EDGE, 'height', 'box-sizing', 'background-color', 'appearance', '-webkit-appearance', 'cursor', 'text-overflow', ...FOCUS, ...MOTION], box: true, position: true, states: ['hover', 'focus', 'focus-visible'] },
    { name: 'dropdown.label', select: `${DD} .form-dropdown-label`, props: [...TYPE, 'position', 'top', 'left', 'right', 'pointer-events', 'white-space', 'overflow-x', 'text-overflow', ...MOTION], box: true, position: true },
    { name: 'dropdown.chevron', select: `${DD} .form-dropdown-chevron`, props: ['position', 'top', 'right', 'width', 'height', 'background-size', 'background-position-x', 'background-position-y', 'background-repeat', 'pointer-events', 'opacity'], box: true, position: true },
    { name: 'dropdown.chosen', inject: CHOSEN, into: SECTION, select: '.form-dropdown-select', props: [...TYPE, ...BOX, 'height'] },
    { name: 'dropdown.chosen-label', inject: CHOSEN, into: SECTION, select: '.form-dropdown-label', props: [...TYPE, 'top', 'left'] },
  ],
};
