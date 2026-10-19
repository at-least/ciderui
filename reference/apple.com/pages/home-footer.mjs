// The home page's footer on a phone, a directory section opened (clicked
// once, so a spec of its own): its list shown, its rows, its chevron.
import { TYPE, LINK } from './common.mjs';
const SECTION = '#ac-globalfooter .ac-gf-directory-column-section';
export default {
  url: 'https://www.apple.com/',
  widths: [390],
  probes: [
    { name: 'section.open', click: `${SECTION} .ac-gf-directory-column-section-title-button`, select: SECTION, props: ['padding-bottom'], box: true },
    { name: 'list.open', select: `${SECTION} .ac-gf-directory-column-section-list`, props: ['visibility', 'padding-top', 'padding-bottom', 'transition-property', 'transition-duration', 'transition-timing-function'], box: true },
    { name: 'item.open', select: `${SECTION} .ac-gf-directory-column-section-item`, props: ['margin-bottom', 'padding-top'], box: true },
    { name: 'link.open', select: `${SECTION} .ac-gf-directory-column-section-link`, props: [...TYPE, ...LINK, 'display', 'padding-top', 'padding-right', 'padding-bottom', 'padding-left'], box: true },
    { name: 'icon.open', select: `${SECTION} .ac-gf-directory-column-section-title-icon`, props: ['color', 'transform'] },
  ],
};
