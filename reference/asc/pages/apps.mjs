// The apps: the page's ground; the bar (amp-nav's header, in its shadow
// root: its ground, content, logo, links and the current one); the page's
// sticky header (its title, the add button's icon, the status filter, the
// rule under it); the two banners over the apps (a plain one with its icon,
// an informational one); and an app's tile: its icon, its name, its
// version's status (a colored mark before the line).
import { TYPE, BOX, LINK } from '../../apple.com/pages/common.mjs';

const BAR = ['height', 'position', 'top', 'z-index', 'background-color', 'border-bottom-width', 'border-bottom-style', 'border-bottom-color', 'box-shadow', 'backdrop-filter'];
const ICON = ['width', 'height', 'fill', 'color'];
const BANNER = 'div:has(> button[class*="Banner__DismissButtonWrapper"])';
export default {
  path: '/apps',
  probes: [
    { name: 'page', select: 'html', props: ['background-color', 'color-scheme'] },
    { name: 'nav', shadow: 'amp-nav', select: 'header.primary-header', props: [...BAR, 'padding-left', 'padding-right'], box: true },
    { name: 'nav.content', shadow: 'amp-nav', select: '.primary-header-content', props: ['max-width', 'padding-left', 'padding-right', 'display', 'align-items', 'justify-content'], box: true },
    { name: 'nav.logo', shadow: 'amp-nav', select: '.logo', props: ['width', 'height', 'color'], box: true },
    { name: 'nav.link', shadow: 'amp-nav', select: '.primary-nav a', nth: 1, props: [...TYPE, ...LINK, 'padding-bottom', 'border-bottom-width', 'border-bottom-color'], box: true },
    { name: 'nav.current', shadow: 'amp-nav', select: '.primary-nav a.active-item', props: [...TYPE, ...LINK, 'padding-bottom', 'border-bottom-width', 'border-bottom-color'], box: true },
    { name: 'header', select: '[class*="Sticky__StickyWrapper"]', props: ['position', 'top', 'z-index', 'background-color', 'padding-top', 'padding-bottom'], box: true },
    { name: 'title', select: '[class*="Sticky__StickyWrapper"] h2', props: [...TYPE, 'margin-top', 'margin-bottom'], box: true },
    { name: 'add.icon', select: '[class*="Sticky__StickyWrapper"] button[class*="Link-sc"] svg', props: ICON, box: true },
    { name: 'filter', select: 'button[class*="Link-sc"]', text: 'All Statuses', props: [...TYPE, ...LINK, 'padding-top', 'padding-right', 'padding-bottom', 'padding-left', 'cursor'], box: true },
    { name: 'rule', select: '[class*="Sticky__StickyWrapper"] hr', props: ['height', 'border-top-width', 'border-top-style', 'border-top-color', 'background-color', 'margin-top', 'margin-bottom'], box: true },
    { name: 'banner', select: BANNER, props: [...BOX, 'display', 'align-items', 'column-gap'], box: true },
    { name: 'banner.icon', select: `${BANNER} svg.glyph-box`, props: ICON, box: true },
    { name: 'banner.title', select: `${BANNER} p`, props: [...TYPE, 'margin-top', 'margin-bottom'], box: true },
    { name: 'banner.text', select: `${BANNER} p`, nth: 1, props: [...TYPE, 'margin-top', 'margin-bottom'], box: true },
    { name: 'banner.link', select: `${BANNER} p a`, props: [...TYPE, ...LINK] },
    { name: 'banner.dismiss', select: 'button[class*="Banner__DismissButtonWrapper"]', props: ['position', 'top', 'right', 'width', 'height', 'cursor'], box: true },
    { name: 'banner.dismiss.icon', select: 'button[class*="Banner__DismissButtonWrapper"] svg', props: ICON, box: true },
    { name: 'banner.info', select: BANNER, nth: 1, props: [...BOX, 'display', 'align-items', 'column-gap'], box: true },
    { name: 'banner.info.icon', select: `${BANNER} > svg`, props: ICON, box: true },
    { name: 'banner.info.dismiss.icon', select: 'button[class*="Banner__DismissButtonWrapper"] svg', nth: 1, props: ICON },
    { name: 'app.icon', select: '[role=grid] a[class*="Link-sc"]', props: ['width', 'height', 'border-top-left-radius', 'box-shadow', 'overflow'], box: true },
    { name: 'app.name', select: '[role=grid] p[class*="Text__BaseText"]', props: [...TYPE, 'margin-top', 'margin-bottom'], box: true },
    { name: 'app.status', select: '[role=grid] ul[class*="List__BaseList"] a', props: [...TYPE, ...LINK], box: true },
    { name: 'app.status.icon', select: '[role=grid] ul[class*="List__BaseList"] svg', props: [...ICON, 'margin-right', 'vertical-align'], box: true },
  ],
};
