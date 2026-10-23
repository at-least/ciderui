// The list as the site opens, nothing chosen: the page's ground; the bar
// (Apple's dark local nav: its ground and the layer over it, the title, the
// compose and help glyphs, the account's menu); the list's pane: its search
// field, the sort and filter row, an item (an announcement, unread: its
// card, the unread mark, its title, its kind, the line under it, its date);
// and the other pane's empty state.
import { TYPE, BOX, PAGE } from '../../apple.com/pages/common.mjs';

const CARD = '[class*="ListItemWrapper"] a[href^="/news/"] [class^="Card-sc-"]';
const SIDES = ['border-left-width', 'border-left-style', 'border-left-color', 'border-right-width', 'border-right-style', 'border-right-color'];
export default {
  path: '/',
  probes: [
    { name: 'page', select: 'body', props: PAGE },
    { name: 'page.ground', select: 'div:has(> [class*="NavbarStyles__NavbarWrapper"])', props: ['background-color'], box: true },
    { name: 'bar', select: 'nav.localnav', props: ['height', 'position', 'top', 'z-index'], box: true },
    { name: 'bar.ground', select: 'nav.localnav .localnav-background', props: ['background-color', 'backdrop-filter', '-webkit-backdrop-filter', 'border-bottom-width', 'border-bottom-color'], box: true },
    { name: 'bar.layer', select: 'nav.localnav .localnav-background', pseudo: '::after', props: ['content', 'position', 'top', 'bottom', 'height', 'background-color', 'backdrop-filter', '-webkit-backdrop-filter'] },
    { name: 'bar.title', select: 'nav.localnav [class*="NavbarStyles__BrandText"]', props: TYPE, box: true },
    { name: 'bar.compose', select: '[class*="NavbarStyles__ComposeIcon"]', props: [...TYPE, 'width', 'height'], box: true },
    { name: 'bar.help', select: 'a[class*="NavbarStyles__HelpIconWrapper"]', props: [...TYPE, 'width', 'height'], box: true },
    { name: 'bar.menu', select: '[class*="NavbarStyles__UserMenu"]', props: [...TYPE, 'cursor'], box: true },
    { name: 'pane', select: 'aside[class^="LeftPane-sc-"]', props: ['width', 'background-color', ...SIDES], box: true },
    { name: 'search', select: '[class*="ItemSearchField__SearchWrapper"]', props: [...BOX, 'height'], box: true },
    { name: 'search.field', select: '[class*="ItemSearchField__SearchWrapper"] input', props: [...TYPE, ...BOX, 'height'], box: true },
    { name: 'search.placeholder', select: '[class*="ItemSearchField__SearchWrapper"] input', pseudo: '::placeholder', props: TYPE },
    { name: 'sort', select: '[class*="ListContainerStyles__ItemSortChooser"]', props: [...BOX, 'height'], box: true },
    { name: 'sort.label', select: '[class*="ListContainerStyles__ItemSortChooser"] .itemSorter', props: [...TYPE, 'cursor'], box: true },
    { name: 'sort.chevron', select: '[class*="ListContainerStyles__ItemSortChooser"] .icon-chevrondown', props: [...TYPE, 'width', 'height'], box: true },
    { name: 'filter', select: '[class*="ListContainerStyles__ItemFilter"] button.filterIcon', props: [...TYPE, ...BOX, 'width', 'height', 'cursor'], box: true },
    { name: 'filter.icon', select: '[class*="ListContainerStyles__ItemFilter"] button.filterIcon svg', props: ['width', 'height', 'color', 'fill'], box: true },
    { name: 'item', select: CARD, props: [...BOX, ...SIDES, 'height', 'display'], box: true },
    { name: 'item.mark', select: `${CARD} svg.unread`, props: ['width', 'height', 'margin-top', 'margin-left'], box: true },
    { name: 'item.mark.fill', select: `${CARD} svg.unread path`, props: ['fill'] },
    { name: 'item.title', select: `${CARD} [class^="Title-sc-"]`, props: [...TYPE, 'white-space', 'text-overflow', 'overflow'], box: true },
    { name: 'item.kind', select: `${CARD} [class^="Type-sc-"]`, props: TYPE, box: true },
    { name: 'item.line', select: `${CARD} [class^="Subtext-sc-"]`, props: TYPE, box: true },
    { name: 'item.date', select: `${CARD} [class^="Stamp-sc-"]`, props: [...TYPE, 'text-align'], box: true },
    { name: 'empty', select: 'article[class^="RightPane-sc-"] [class*="routes__Center"]', props: [...TYPE, 'display', 'align-items', 'justify-content', 'text-align'], box: true },
  ],
};
