// @external: Apple Developer Forums' post list (`/forums/latest`), signed out: a row's
// title, excerpt, tags and stats, the list's rule, and the count over a tag's list.
// The site draws one list for a wide window (`.desktop-view`) and another to its
// phone width (`.mobile-view`), each hidden at the other's widths.
import { TYPE, BOX, FLEX, LINK, PAGE } from './common.mjs';

const D = '.content-list-wrapper.desktop-view article.article-entry';
const M = '.content-list-wrapper.mobile-view article.article-entry';
const CLAMP = ['display', '-webkit-line-clamp', '-webkit-box-orient', 'overflow', 'text-overflow', 'white-space'];
export default {
  url: 'https://developer.apple.com/forums/latest',
  wait: 3000,
  widths: [1600, 1200, 900, 735, 734, 390],
  probes: [
    { name: 'page', select: 'body', props: PAGE },
    { name: 'title', select: 'h1.overview-title', props: [...TYPE, 'margin-top', 'margin-bottom'], box: true },
    { name: 'list', select: '.content-list-wrapper.desktop-view', props: [...BOX, 'display', 'width'], box: true },
    { name: 'row', select: D, props: [...BOX, ...FLEX, 'grid-template-columns'], box: true, states: ['hover'] },
    { name: 'row.second', select: `${D} + article.article-entry`, props: [...BOX], box: true },
    { name: 'row.summary', select: `${D} .article-summary`, props: [...BOX, ...FLEX, 'flex-grow', 'flex-basis', 'width'], box: true },
    { name: 'row.title', select: `${D} a.post-title`, props: [...TYPE, ...LINK], states: ['hover', 'focus-visible'] },
    { name: 'row.title.text', select: `${D} a.post-title span.underline`, props: [...TYPE, ...LINK], states: ['hover'], stateOn: `${D} a.post-title` },
    { name: 'row.excerpt', select: `${D} .excerpt`, props: [...TYPE, ...CLAMP, 'margin-top', 'padding-top'], box: true },
    { name: 'row.chiclets', select: `${D} .chiclets`, props: [...FLEX, 'margin-top', 'padding-top'], box: true },
    { name: 'row.tag', select: `${D} a.tag`, props: [...TYPE, ...BOX, 'text-transform'], box: true, states: ['hover'] },
    { name: 'row.stats', select: `${D} .article-stats`, props: [...BOX, ...FLEX, 'width', 'flex-shrink', 'grid-template-columns'], box: true },
    { name: 'row.stat', select: `${D} .stat-item.replies`, props: [...TYPE, 'text-align', 'width', 'font-variant-numeric'], box: true },
    { name: 'row.time', select: `${D} .stat-item.timestamp`, props: [...TYPE, 'text-align', 'width'], box: true },
    { name: 'mrow', select: M, props: [...BOX, ...FLEX], box: true, widths: [734, 390] },
    { name: 'mrow.title', select: `${M} a.post-title`, props: [...TYPE], widths: [734, 390] },
    { name: 'mrow.excerpt', select: `${M} .excerpt`, props: [...TYPE, ...CLAMP, 'margin-top'], box: true, widths: [734, 390] },
    { name: 'mrow.stats', select: `${M} dl.mobile-stats`, props: [...BOX, ...FLEX], box: true, widths: [734, 390] },
    { name: 'mrow.stat', select: `${M} dd.mobile-icon.replies`, props: [...TYPE, 'margin-left', 'padding-left'], box: true, widths: [734, 390] },
  ],
};
