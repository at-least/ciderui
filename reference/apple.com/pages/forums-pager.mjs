// @external: Apple Developer Forums' pager on a middle page (`/forums/latest?page=3`),
// signed out: the nav, its row, the first page's button (a word), the previous and
// next pages' (a chevron each, hovered and focused), the page's number between.
// A direction with no page is not drawn (`../../developer-forums/`, page 1's).
import { TYPE, BOX, FLEX } from './common.mjs';

const NAV = 'nav.pagination';
const BUTTON = [...TYPE, ...BOX, 'height', 'display', 'cursor', 'text-decoration-line'];
export default {
  url: 'https://developer.apple.com/forums/latest?page=3',
  wait: 3000,
  widths: [1200, 390],
  probes: [
    { name: 'nav', select: NAV, props: [...BOX], box: true },
    { name: 'list', select: `${NAV} ol`, props: [...FLEX, 'margin-left', 'margin-right'], box: true },
    { name: 'first', select: `${NAV} a.first-page`, props: BUTTON, box: true, states: ['hover', 'active', 'focus-visible'] },
    { name: 'previous', select: `${NAV} a.previous-page`, props: BUTTON, box: true, states: ['hover', 'active', 'focus-visible'] },
    { name: 'previous.icon', select: `${NAV} a.previous-page sf-symbol`, props: ['width', 'height', 'font-size', 'vertical-align'], box: true },
    { name: 'page', select: `${NAV} .page-number`, props: [...TYPE, ...BOX, 'display'], box: true },
    { name: 'next', select: `${NAV} a.next-page`, props: BUTTON, box: true },
  ],
};
