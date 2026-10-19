// The global nav's search, opened (its button clicked once, at the first
// width of each of the nav's layouts: it stays open as the window narrows
// within one, and closes where the nav changes, at 833px), so its probes are
// a spec of their own: the home page's flyout, its field (the input, its
// placeholder, the search and clear buttons), its quick links (a heading, a
// list of links each after an arrow), the curtain over the page, and the
// bar while it is open. From 834px the flyout drops under the bar; to 833px it
// is a sheet over the bar and the rest of the window. Once it has opened the
// pointer moves into it, below its links (200, 430; 200, 850 on the sheet),
// where it hovers none: leaving it for the curtain closes it.
// The page is read with its locale banner above the nav (#ac-localeswitcher,
// for a reader outside the U.S.: 70px tall from 834px, 135px to 833px), so a
// box's top is the nav's top plus its own.
import { TYPE, BOX, MOTION, BLUR, FOCUS } from './common.mjs';

const SPACE = ['padding-top', 'padding-right', 'padding-bottom', 'padding-left', 'margin-top', 'margin-right', 'margin-bottom', 'margin-left'];
const FLY = '#globalnav-submenu-search';
const WIDE = [1600, 1200, 900, 834];
const PHONE = [833, 390];
const NAV = ['height', 'position', 'background-color', ...BLUR];
const NAV_VARS = ['--globalnav-background', '--globalnav-backdrop-filter'];

export default {
  url: 'https://www.apple.com/',
  widths: [1600, 1200, 900, 834, 833, 390],
  wait: 2500,
  probes: [
    { name: 'nav', click: '#globalnav-menubutton-link-search', pointer: [200, 430], animations: '#globalnav', select: '#globalnav', props: NAV, vars: NAV_VARS, box: true, position: true, widths: WIDE },
    { name: 'nav.phone', click: '#globalnav-menubutton-link-search', pointer: [200, 850], animations: '#globalnav', select: '#globalnav', props: NAV, vars: NAV_VARS, box: true, position: true, widths: PHONE },
    // the bar's content box, the sheet's ground to 833px; the page held
    { name: 'content', select: '#globalnav .globalnav-content', props: ['position', 'top', 'height', 'background-color', ...BLUR, 'z-index', ...MOTION], box: true, position: true },
    { name: 'root', select: 'html', props: ['overflow-y'] },
    { name: 'body', select: 'body', props: ['overflow-y'] },
    { name: 'search', select: '#globalnav-menubutton-link-search', props: ['opacity', 'visibility', 'color', ...MOTION], box: true, position: true },
    { name: 'flyout', select: FLY, props: ['position', 'top', 'left', 'right', 'height', 'max-height', 'background-color', 'color', 'overflow-y', 'z-index', ...MOTION], box: true, position: true },
    { name: 'flyout.scroll', select: `${FLY} .globalnav-flyout-scroll-container`, props: ['height', 'max-height', 'overflow-y', ...SPACE], box: true, position: true },
    { name: 'flyout.content', select: `${FLY} .globalnav-flyout-content`, props: ['max-width', ...SPACE], box: true, position: true },
    { name: 'field', select: `${FLY} .globalnav-searchfield`, props: [...SPACE, 'height'], box: true, position: true },
    { name: 'field.input', select: `${FLY} .globalnav-searchfield-input`, props: [...TYPE, ...SPACE, 'height', 'background-color', 'border-top-width', 'border-top-style', 'border-top-left-radius', 'outline-style', 'caret-color', 'text-overflow', '-webkit-appearance', 'appearance'], box: true, position: true },
    { name: 'field.placeholder', select: `${FLY} .globalnav-searchfield-input`, pseudo: '::placeholder', props: ['color', 'font-weight', 'opacity'] },
    { name: 'field.focus', select: `${FLY} .globalnav-searchfield-input`, props: FOCUS, states: ['focus-visible', 'focus'] },
    { name: 'field.submit', select: `${FLY} .globalnav-searchfield-submit`, props: [...SPACE, 'width', 'height', 'position', 'left', 'top', 'color', 'fill', 'background-color', 'border-top-width', 'cursor', ...MOTION], box: true, position: true },
    { name: 'field.submit-image', select: `${FLY} .globalnav-searchfield-submit .globalnav-image-regular svg`, props: ['width', 'height', 'fill'], box: true, position: true, widths: WIDE },
    { name: 'field.submit-compact', select: `${FLY} .globalnav-searchfield-submit .globalnav-image-compact svg`, props: ['width', 'height', 'fill'], box: true, position: true, widths: PHONE },
    { name: 'field.reset', select: `${FLY} .globalnav-searchfield-reset`, props: ['opacity', 'visibility', 'width', 'height', ...MOTION], box: true, position: true },
    { name: 'results', select: `${FLY} .globalnav-searchresults`, props: [...SPACE], box: true, position: true },
    { name: 'results.header', select: `${FLY} .globalnav-searchresults-header`, props: [...TYPE, ...SPACE], box: true, position: true },
    { name: 'results.list', select: `${FLY} .globalnav-searchresults-list`, props: [...SPACE, 'list-style-type'], box: true, position: true },
    { name: 'results.item', select: `${FLY} .globalnav-searchresults-list-item:first-child`, props: [...SPACE, 'display'], box: true, position: true },
    { name: 'results.item2', select: `${FLY} .globalnav-searchresults-list-item:nth-child(2)`, props: [...SPACE], box: true, position: true },
    { name: 'results.link', select: `${FLY} .globalnav-searchresults-list-link`, props: [...TYPE, ...SPACE, 'display', 'align-items', 'text-decoration-line', 'background-color', ...MOTION], box: true, position: true, states: ['hover', 'focus-visible'] },
    { name: 'results.icon', select: `${FLY} .globalnav-searchresults-list-icon`, props: [...SPACE, 'display', 'color'], box: true, position: true },
    { name: 'results.icon-image', select: `${FLY} .globalnav-searchresults-list-icon .globalnav-image-regular svg`, props: ['width', 'height', 'fill'], box: true, position: true, widths: WIDE },
    { name: 'results.icon-compact', select: `${FLY} .globalnav-searchresults-list-icon .globalnav-image-compact svg`, props: ['width', 'height', 'fill'], box: true, position: true, widths: PHONE },
    { name: 'results.text', select: `${FLY} .globalnav-searchresults-list-text`, props: [...TYPE, ...SPACE], box: true, position: true },
    { name: 'curtain', select: '#globalnav-curtain', props: ['position', 'top', 'z-index', 'background-color', ...BLUR, 'opacity', 'visibility', ...MOTION], box: true, position: true },
  ],
};
