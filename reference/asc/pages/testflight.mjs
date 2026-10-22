// An app's TestFlight builds: the app's tabs (Distribution, Analytics,
// TestFlight: a link and the current one), the sidebar (a section's heading,
// a link and the current one with its row, the rule between sections), the
// page's title, a group of builds opened (its chevron and title), and its
// table: the header row and a column's label, a row, a cell, the build
// cell's version link, a status (its colored mark, its word, the line under
// it) and a count, right-aligned, and the info button's icon.
import { TYPE, BOX, LINK } from '../../apple.com/pages/common.mjs';

const ICON = ['width', 'height', 'fill', 'color'];
const CELL = ['padding-top', 'padding-right', 'padding-bottom', 'padding-left', 'vertical-align', 'text-align', 'background-color', 'border-top-width', 'border-top-style', 'border-top-color', 'border-bottom-width', 'border-bottom-style', 'border-bottom-color'];
const ROW = ['height', 'background-color', 'border-top-width', 'border-top-style', 'border-top-color', 'border-bottom-width', 'border-bottom-style', 'border-bottom-color'];
export default {
  path: '/teams/<team>/apps/<app>/testflight/ios',
  probes: [
    { name: 'tabs', select: 'nav.sub-nav-container', props: ['height', 'border-bottom-width', 'border-bottom-style', 'border-bottom-color'], box: true },
    { name: 'tabs.link', select: 'nav.sub-nav-container a', props: [...TYPE, ...LINK, 'padding-left', 'padding-right', 'border-bottom-width', 'border-bottom-style', 'border-bottom-color'], box: true },
    { name: 'tabs.current', select: 'nav.sub-nav-container a', nth: 2, props: [...TYPE, ...LINK, 'padding-left', 'padding-right', 'border-bottom-width', 'border-bottom-style', 'border-bottom-color'], box: true },
    { name: 'sidebar', select: '[role=navigation]', props: ['width', 'padding-right'], box: true },
    { name: 'sidebar.heading', select: '[role=navigation] h2 span', props: TYPE, box: true },
    { name: 'sidebar.link', select: '[role=navigation] a:not([aria-current])', props: [...TYPE, ...LINK, 'padding-left', 'padding-right', 'height'], box: true },
    { name: 'sidebar.link.row', select: '[role=navigation] li:has(> a:not([aria-current]))', props: ['background-color', 'border-top-left-radius', 'height', 'margin-bottom'], box: true },
    { name: 'sidebar.current', select: '[role=navigation] a[aria-current]', props: [...TYPE, ...LINK, 'padding-left', 'padding-right', 'height'], box: true },
    { name: 'sidebar.current.row', select: '[role=navigation] li:has(> a[aria-current])', props: ['background-color', 'border-top-left-radius', 'height', 'margin-bottom'], box: true },
    { name: 'sidebar.rule', select: '[role=navigation] hr', props: ['height', 'background-color', 'border-top-style', 'margin-top', 'margin-bottom'], box: true },
    { name: 'title', select: 'h2:not([role=navigation] h2)', props: [...TYPE, 'margin-top', 'margin-bottom'], box: true },
    { name: 'group', select: 'button[type=button][aria-expanded=true]', props: [...BOX, 'height', 'display', 'align-items', 'column-gap'], box: true },
    { name: 'group.chevron', select: 'button[type=button][aria-expanded=true] svg', props: [...ICON, 'transform'], box: true },
    { name: 'group.title', select: 'button[type=button][aria-expanded=true] h3', props: [...TYPE, 'margin-left'], box: true },
    { name: 'table', select: 'table[role=grid]', props: ['border-collapse', 'background-color', 'margin-top', 'margin-left', ...TYPE], box: true },
    { name: 'head', select: 'thead tr', props: ROW, box: true },
    { name: 'head.cell', select: 'thead th', props: [...CELL, 'height'], box: true },
    { name: 'head.label', select: 'thead th span[class*="Text__BaseText"]', props: [...TYPE, 'text-transform'], box: true },
    { name: 'row', select: 'tbody tr', props: ROW, box: true },
    { name: 'cell', select: 'tbody td[role=gridcell]', props: [...CELL, ...TYPE], box: true },
    { name: 'cell.first', select: 'tbody td[role=rowheader]', props: [...CELL, ...TYPE], box: true },
    { name: 'build.link', select: 'tbody td[role=rowheader] span[class*="Link-sc"]', props: [...TYPE, ...LINK] },
    { name: 'status.icon', select: 'tbody td[role=gridcell] svg[class*="SvgIcon"]', props: [...ICON, 'margin-right', 'vertical-align'], box: true },
    { name: 'status.word', select: 'tbody td[role=gridcell] svg[class*="SvgIcon"] + span', props: [...TYPE, 'margin-left'], box: true },
    { name: 'status.line', select: 'tbody td[role=gridcell] span[class*="Text__BaseText"]', nth: 1, props: [...TYPE, 'display', 'margin-top'], box: true },
    { name: 'count', select: 'tbody tr td:nth-child(5) p', props: [...TYPE, 'text-align'], box: true },
    { name: 'info.icon', select: 'tbody button[class*="Link-sc"] svg', props: ICON, box: true },
  ],
};
