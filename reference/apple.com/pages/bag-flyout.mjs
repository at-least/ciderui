// The global nav's bag, opened (its button clicked once, at the first width
// of each of the nav's layouts, as search.mjs opens the search): the bag's
// flyout, a link's flyout's own (flyout.mjs's: globalnav-flyout), its
// markup its script's, fetched (/shop/bag/flyout) and an empty bag's here:
// its message, a line with a sign-in link, the profile's heading and its
// links, each after its icon; the bar while it is open, its button, the
// curtain, each with what the click starts; the profile's links hovered and
// focused (the forced states last, as search.mjs's). From 834px the pointer
// moves into the flyout's room left of its content, 200px down, where it
// hovers nothing and which the flyout holds before its markup has come
// (0.9s after the click: 174px tall at 0.3s) and at every width (one left
// out of the window as it narrows has left the flyout: leaving it closes
// it); at 834px, last, it leaves for the curtain (left: leaving closes it,
// as the search). To 833px it is a sheet over the window, the menu's
// button its X; the pointer moves into the sheet below its links. The page
// is read with its locale banner above the nav (70px from 834px, 135px to
// 833px), so a box's top is the nav's top plus its own.
import { TYPE, MOTION, BLUR, FOCUS, LINK } from './common.mjs';

const FLY = '#globalnav-submenu-bag';
const LINKED = `${FLY} .ac-gn-bagview-nav-link-orders`;
const SPACE = ['padding-top', 'padding-right', 'padding-bottom', 'padding-left', 'margin-top', 'margin-right', 'margin-bottom', 'margin-left'];
const ITEM = ['opacity', 'transform', 'transition-delay', ...MOTION];
const WIDE = [1600, 1200, 900, 834];
const PHONE = [833, 390];
const MOTIONS = '#globalnav, #globalnav-curtain';
// the flyout's room left of its content
const ROOM = [100, 200];

export default {
  url: 'https://www.apple.com/',
  widths: [1600, 1200, 900, 834, 833, 390],
  wait: 2500,
  probes: [
    { name: 'nav', click: '#globalnav-menubutton-link-bag', pointer: ROOM, animations: MOTIONS, select: '#globalnav', props: ['height', 'position', 'background-color', ...BLUR, 'z-index'], vars: ['--globalnav-background', '--globalnav-backdrop-filter'], box: true, position: true, widths: WIDE },
    { name: 'nav.phone', click: '#globalnav-menubutton-link-bag', pointer: [200, 850], animations: MOTIONS, select: '#globalnav', props: ['height', 'position', 'background-color', ...BLUR, 'z-index'], vars: ['--globalnav-background', '--globalnav-backdrop-filter'], box: true, position: true, widths: PHONE },
    { name: 'nav.content', select: '#globalnav .globalnav-content', props: ['position', 'z-index', 'height', 'background-color', ...BLUR], box: true, position: true },
    { name: 'root', select: 'html', props: ['overflow-y'] },
    { name: 'body', select: 'body', props: ['overflow-y'] },
    // the bag's button while it is open, another link, the search beside them
    { name: 'link.bag', select: '#globalnav-menubutton-link-bag', props: ['color', 'opacity', 'visibility', ...MOTION], box: true, position: true },
    { name: 'link.other', select: '#globalnav .globalnav-link-mac', props: ['color', 'opacity', 'visibility'] },
    { name: 'link.search', select: '#globalnav .globalnav-link-search', props: ['color', 'opacity', 'visibility'] },
    // to 833px the menu's button is the sheet's X
    { name: 'menutrigger', select: '#globalnav-menutrigger-button', props: ['color', 'opacity', 'visibility', ...MOTION], box: true, position: true, widths: PHONE },
    { name: 'flyout', select: FLY, props: ['position', 'top', 'left', 'right', 'height', 'max-height', ...SPACE, 'background-color', ...BLUR, 'color', 'overflow-y', 'z-index', 'visibility', ...MOTION], vars: ['--r-globalnav-flyout-height', '--r-globalnav-flyout-rate'], box: true, position: true },
    { name: 'flyout.scroll', select: `${FLY} .globalnav-flyout-scroll-container`, props: ['height', 'max-height', 'overflow-y', ...SPACE], box: true, position: true },
    { name: 'flyout.content', select: `${FLY} .globalnav-flyout-content`, props: ['max-width', ...SPACE, 'box-sizing'], box: true, position: true },
    { name: 'bag', select: `${FLY} .ac-gn-bagview-content`, props: [...SPACE], box: true, position: true },
    { name: 'bag.header', select: `${FLY} .ac-gn-bagview-header`, props: [...TYPE, ...SPACE, ...ITEM], box: true, position: true },
    { name: 'bag.subheader', select: `${FLY} .ac-gn-bagview-subheader`, props: [...TYPE, ...SPACE, ...ITEM], box: true, position: true },
    { name: 'bag.nav', select: `${FLY} .ac-gn-bagview-nav-item-wrapper`, props: [...SPACE], box: true, position: true },
    { name: 'bag.nav-header', select: `${FLY} .ac-gn-bagview-nav-item-header`, props: [...TYPE, ...SPACE, ...ITEM], box: true, position: true },
    { name: 'bag.nav-list', select: `${FLY} .ac-gn-bagview-nav-item-wrapper ul`, props: [...SPACE, 'list-style-type'], box: true, position: true },
    { name: 'bag.nav-item', select: `${FLY} .ac-gn-bagview-nav-item-orders`, props: [...SPACE, 'display', ...ITEM], box: true, position: true },
    { name: 'bag.nav-item2', select: `${FLY} .ac-gn-bagview-nav-item-yoursaves`, props: [...ITEM], box: true, position: true },
    { name: 'bag.nav-item-last', select: `${FLY} .ac-gn-bagview-nav-item-signIn`, props: [...ITEM], box: true, position: true },
    { name: 'bag.nav-image', select: `${LINKED} .ac-gn-bagview-nav-image-container`, props: ['position', 'left', 'top', 'display', ...SPACE], box: true, position: true },
    { name: 'bag.nav-image-regular', select: `${LINKED} .ac-gn-bagview-nav-image-regular svg`, props: ['width', 'height', 'fill', 'color'], box: true, position: true, widths: WIDE },
    { name: 'bag.nav-path', select: `${LINKED} .ac-gn-bagview-nav-image-regular path`, props: ['fill'], widths: WIDE },
    { name: 'bag.nav-image-compact', select: `${LINKED} .ac-gn-bagview-nav-image-compact svg`, props: ['width', 'height', 'fill', 'color'], box: true, position: true, widths: PHONE },
    { name: 'bag.nav-path-compact', select: `${LINKED} .ac-gn-bagview-nav-image-compact path`, props: ['fill'], widths: PHONE },
    { name: 'curtain', select: '#globalnav-curtain', props: ['position', 'top', 'z-index', 'background-color', ...BLUR, 'opacity', 'visibility', ...MOTION], box: true, position: true },
    // the forced states last (search.mjs's): the sign-in link, a profile link,
    // its text and its icon
    { name: 'bag.subheader-link', select: `${FLY} .ac-gn-bagview-subheader a`, props: [...TYPE, ...LINK, ...FOCUS], box: true, position: true, states: ['hover', 'focus-visible'] },
    { name: 'bag.nav-link', select: LINKED, props: [...TYPE, ...SPACE, 'display', 'align-items', 'position', 'text-decoration-line', 'background-color', 'border-top-left-radius', ...FOCUS, ...MOTION], box: true, position: true, states: ['hover', 'focus-visible'] },
    // the text hovered itself: apple.com's rule is the link's :hover and the
    // text's (globalheader.css), so the pointer on the text, as a hand's
    // is, turns it #000 where the link's forced alone leaves it
    { name: 'bag.nav-text', select: `${LINKED} .ac-gn-bagview-nav-text`, props: [...TYPE, ...SPACE, 'display', ...MOTION], box: true, position: true, states: ['hover'] },
    { name: 'bag.nav-icon', select: `${LINKED} .ac-gn-bagview-nav-image-regular svg`, stateOn: LINKED, props: ['color', 'fill', ...MOTION], states: ['hover'], widths: WIDE },
    { name: 'bag.nav-icon-compact', select: `${LINKED} .ac-gn-bagview-nav-image-compact svg`, stateOn: LINKED, props: ['color', 'fill', ...MOTION], states: ['hover'], widths: PHONE },
    // at 834px, last: the pointer off the flyout, onto the curtain
    { name: 'left', hover: [400, 800], hovered: 1500, animations: MOTIONS, select: FLY, props: ['height', 'visibility'], box: true, widths: [834] },
    { name: 'left.button', select: '#globalnav-menubutton-link-bag', props: ['color'], widths: [834] },
  ],
};
