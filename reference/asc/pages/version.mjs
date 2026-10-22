// An app's version page (its form, read, nothing typed or pressed): the
// tabs over a section (links: the current one, another; buttons, the media's:
// the current one, another, and their text), the header's buttons (the
// primary, a neutral one, the save button disabled while nothing has
// changed), a help button, a text field and a disabled one, a field's label,
// a multi-line field, the count under it, and the checkbox and radio (the
// browser's own).
import { TYPE, BOX, LINK } from '../../apple.com/pages/common.mjs';

const TAB = ['height', 'padding-bottom', 'border-bottom-width', 'border-bottom-style', 'border-bottom-color'];
const BUTTON = [...TYPE, ...BOX, 'height', 'cursor', 'opacity'];
const CHOICE = ['width', 'height', 'appearance', 'accent-color', 'margin-top', 'margin-right', 'margin-left', 'vertical-align'];
export default {
  path: '/apps/<app>/distribution/ios/version/inflight',
  probes: [
    { name: 'tab.current', select: '[role=tablist] a[role=tab][aria-selected=true]', props: [...TYPE, ...LINK, ...TAB], box: true },
    { name: 'tab', select: '[role=tablist] a[role=tab][aria-selected=false]', props: [...TYPE, ...LINK, ...TAB], box: true },
    { name: 'tab.button.current', select: '[role=tablist] button[role=tab][aria-selected=true]', props: TAB, box: true },
    { name: 'tab.button.current.text', select: '[role=tablist] button[role=tab][aria-selected=true] > div', props: TYPE, box: true },
    { name: 'tab.button', select: '[role=tablist] button[role=tab][aria-selected=false]', props: TAB, box: true },
    { name: 'tab.button.text', select: '[role=tablist] button[role=tab][aria-selected=false] > div', props: TYPE, box: true },
    { name: 'button.primary', select: 'button:not([type])[class*="Text__BaseText"]', props: BUTTON, box: true },
    { name: 'button.neutral', select: 'button:not([type])[class*="Text__BaseText"]', nth: 1, props: BUTTON, box: true },
    { name: 'button.disabled', select: 'button[type=button][class*="Text__BaseText"]:disabled', props: BUTTON, box: true },
    { name: 'help', select: 'button[role=button][type=button][class*="Text__BaseText"]', props: [...TYPE, 'width', 'height', 'background-color', 'border-top-left-radius'], box: true },
    { name: 'field', select: 'input[class*="Input__BaseInput"]:not([type]):enabled', props: [...TYPE, ...BOX, 'height'], box: true },
    { name: 'field.disabled', select: 'input[class*="Input__BaseInput"]:disabled', props: [...TYPE, ...BOX, 'height'], box: true },
    { name: 'field.label', select: 'label[class*="Text__BaseText"]', props: [...TYPE, 'display', 'margin-bottom'], box: true },
    { name: 'textarea', select: 'textarea[class*="Input__BaseInput"]', props: [...TYPE, ...BOX, 'height', 'resize'], box: true },
    { name: 'field.count', select: 'p[role=status]', props: [...TYPE, 'margin-top', 'text-align'], box: true },
    { name: 'checkbox', select: 'input[type=checkbox]', props: CHOICE, box: true },
    { name: 'radio', select: 'input[type=radio]', props: CHOICE, box: true },
  ],
};
