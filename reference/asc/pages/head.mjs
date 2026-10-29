// An app's analytics overview, its heads: the app's head (the app's name
// menu, its icon, name and chevron, beside the app's tabs, the current one
// Analytics) and its wrappers' room, the content column, and a section's
// head (its h4 title, the date control at its end). The app's icon is a
// div drawn by its background image, inside a box of its size.
import { TYPE, BOX, FLEX } from '../../apple.com/pages/common.mjs';

const ROOM = [...BOX, 'height', 'margin-left', 'margin-right'];
export default {
  path: '/apps/<app>/analytics/overview',
  probes: [
    { name: 'head', select: 'nav.sub-nav-container', up: 2, props: [...ROOM, ...FLEX], box: true },
    { name: 'head.wrap', select: 'nav.sub-nav-container', up: 3, props: [...ROOM, ...FLEX], box: true },
    { name: 'head.wrap.outer', select: 'nav.sub-nav-container', up: 4, props: [...ROOM, ...FLEX], box: true },
    { name: 'head.name', select: 'div.dd-container', props: [...ROOM, ...FLEX], box: true },
    { name: 'head.name.button', select: 'div.dd-container [role=button]', props: [...ROOM, ...FLEX, 'cursor'], box: true },
    { name: 'head.name.text', select: 'div.dd-container [role=button] span', props: [...TYPE, 'margin-left', 'margin-right', 'margin-top', 'margin-bottom'], box: true },
    { name: 'head.name.icon', select: 'div.dd-container [role=button] div[class*="Box-sc"]', nth: 1, props: ['width', 'height', 'border-top-left-radius', 'margin-left', 'margin-right'], box: true },
    { name: 'head.name.chevron', select: 'div.dd-container svg', props: ['width', 'height', 'color', 'fill', 'margin-left', 'margin-right'], box: true },
    { name: 'head.tabs', select: 'nav.sub-nav-container', props: [...ROOM, ...FLEX], box: true },
    { name: 'head.tab', select: 'ul.sub-nav-links > li', props: ROOM, box: true },
    { name: 'head.tab.link', select: 'ul.sub-nav-links a', props: [...TYPE, ...ROOM], box: true },
    { name: 'head.tab.current', select: 'ul.sub-nav-links a', nth: 1, props: [...TYPE, ...ROOM], box: true },
    { name: 'content', select: 'article[class*="rt-Content"]', props: ROOM, box: true },
    { name: 'section.head', select: 'div:has(> p[class*="rt-Text--textStyle--$h4Emph"])', props: [...ROOM, ...FLEX], box: true },
    { name: 'section.title', select: 'p[class*="rt-Text--textStyle--$h4Emph"]', props: [...TYPE, 'margin-top', 'margin-bottom'], box: true },
  ],
};
