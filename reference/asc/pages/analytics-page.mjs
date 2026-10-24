// An app's analytics overview, the newer look, what analytics.mjs left
// unread: a metric's card (its ground, corners and padding), a section's
// heading, its See All link, the menu's group (its button, open, and an
// item in it, indented; an item the app has no data for, disabled) and the
// menu's own box.
import { TYPE, BOX, LINK } from '../../apple.com/pages/common.mjs';

const ITEM = ['height', 'background-color', 'border-top-left-radius', 'padding-left', 'padding-right', 'column-gap', 'color', 'opacity', 'cursor'];
const LABEL = '> div > div';
export default {
  path: '/apps/<app>/analytics/overview',
  probes: [
    { name: 'card', select: 'a:has(p[class*="rt-Text--textStyle--$bodyEmph"])', props: [...BOX, 'overflow'], box: true },
    { name: 'section.heading', select: 'p[class*="rt-Text--textStyle--$h4Emph"]', props: [...TYPE, 'margin-top', 'margin-bottom'], box: true },
    { name: 'section.link', select: 'a[class*="rt-Link--variant--inline"]', props: [...TYPE, ...LINK, 'cursor'], box: true },
    { name: 'menu', select: 'nav[class*="rt-DisplayList--variant--menu"]', props: [...BOX, 'row-gap', 'display'], box: true },
    { name: 'group.button', select: 'button[class*="rt-Item"]', props: ITEM, box: true },
    { name: 'group.button.label', select: `button[class*="rt-Item"] ${LABEL}`, props: TYPE, box: true },
    { name: 'group.chevron', select: 'button[class*="rt-Item"] svg', props: ['width', 'height', 'color', 'fill', 'transform'], box: true },
    { name: 'group.item', select: '[class*="rt-DisclosureGroup--state--open"] > a[class*="rt-Item--effect--default"]', props: ITEM, box: true },
    { name: 'group.item.label', select: `[class*="rt-DisclosureGroup--state--open"] > a[class*="rt-Item--effect--default"] ${LABEL}`, props: [...TYPE, 'margin-left', 'padding-left'], box: true },
    { name: 'item.disabled', select: 'a[class*="rt-Item--effect--disabled"]', props: ITEM, box: true },
    { name: 'item.disabled.label', select: `a[class*="rt-Item--effect--disabled"] ${LABEL}`, props: TYPE, box: true },
  ],
};
