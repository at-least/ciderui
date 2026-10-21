// The bag's flyout on a dark page (iPhone 18 Pro's: its bar clear, the
// flyout dark), opened as bag-flyout.mjs opens it (the pointer into its
// room, 100, 200), at 1600 and 390px: the
// bar, the flyout's ground, its message, line, link, heading and profile
// links (hovered and focused last), the curtain.
import { TYPE, BLUR, FOCUS, LINK } from './common.mjs';

const FLY = '#globalnav-submenu-bag';
const LINKED = `${FLY} .ac-gn-bagview-nav-link-orders`;

export default {
  url: 'https://www.apple.com/iphone-18-pro/',
  widths: [1600, 390],
  wait: 2500,
  probes: [
    { name: 'nav', click: '#globalnav-menubutton-link-bag', pointer: [100, 200], select: '#globalnav', props: ['background-color', ...BLUR], widths: [1600] },
    { name: 'nav.phone', click: '#globalnav-menubutton-link-bag', pointer: [200, 850], select: '#globalnav', props: ['background-color', ...BLUR], widths: [390] },
    { name: 'nav.content', select: '#globalnav .globalnav-content', props: ['background-color', ...BLUR] },
    { name: 'link.bag', select: '#globalnav-menubutton-link-bag', props: ['color', 'opacity'] },
    { name: 'menutrigger', select: '#globalnav-menutrigger-button', props: ['color', 'opacity'], widths: [390] },
    { name: 'flyout', select: FLY, props: ['background-color', ...BLUR, 'color'] },
    { name: 'bag.header', select: `${FLY} .ac-gn-bagview-header`, props: TYPE },
    { name: 'bag.subheader', select: `${FLY} .ac-gn-bagview-subheader`, props: TYPE },
    { name: 'bag.nav-header', select: `${FLY} .ac-gn-bagview-nav-item-header`, props: TYPE },
    { name: 'curtain', select: '#globalnav-curtain', props: ['background-color', ...BLUR, 'opacity', 'visibility'] },
    { name: 'bag.subheader-link', select: `${FLY} .ac-gn-bagview-subheader a`, props: [...TYPE, ...LINK, ...FOCUS], states: ['hover', 'focus-visible'] },
    { name: 'bag.nav-link', select: LINKED, props: [...TYPE, 'background-color', ...FOCUS], states: ['hover', 'focus-visible'] },
    { name: 'bag.nav-text', select: `${LINKED} .ac-gn-bagview-nav-text`, props: ['color'], states: ['hover'] },
    { name: 'bag.nav-icon', select: `${LINKED} .ac-gn-bagview-nav-image-regular svg`, stateOn: LINKED, props: ['color', 'fill'], states: ['hover'], widths: [1600] },
    { name: 'bag.nav-icon-compact', select: `${LINKED} .ac-gn-bagview-nav-image-compact svg`, stateOn: LINKED, props: ['color', 'fill'], states: ['hover'], widths: [390] },
  ],
};
