// The global nav's search showing its suggestions, opened as search.mjs opens
// it, with nothing typed and no suggestion asked for: the field given a value
// and its buttons as search-text.mjs gives them, the search marked as
// apple.com's script marks it with suggestions (globalnav-search-with-results,
// added with no event: `set`'s classList), the quick links hidden, and in
// their place a copy the probe adds of what its script writes then (its
// _createAMLSearchTemplate, read in globalheader.umd.js on 2026-10-21): two
// sections, Suggested Links and Suggested Searches, each a heading over rows,
// each row a link with its icon (empty SVGs of apple.com's images' sizes,
// 9 by 16 and 13 by 25: its arrow and its magnifier are read live by
// ../icons/searchfield.mjs) and its text, the part the service marks in a
// span (-text-highlight). The copy is outside the results' current box, so
// none of its script's motion runs on it. Read only.
import search from './search.mjs';
import { TYPE, MOTION, FOCUS } from './common.mjs';

const SPACE = ['padding-top', 'padding-right', 'padding-bottom', 'padding-left', 'margin-top', 'margin-right', 'margin-bottom', 'margin-left'];
const FLY = '#globalnav-submenu-search';
const WIDE = [1600, 1200, 900, 834];
const PHONE = [833, 390];
const INPUT = `${FLY} .globalnav-searchfield-input`;
const SET = [
  [INPUT, 'value', 'iph'],
  ...[`${FLY} .globalnav-searchfield-reset`, `${FLY} .globalnav-searchfield-submit`].flatMap(b => [[b, 'tabindex', '0'], [b, 'disabled', null], [b, 'aria-hidden', null]]),
  ['#globalnav .globalnav-search', 'classList', 'globalnav-search-with-results'],
  [`${FLY} .globalnav-searchresults-current`, 'hidden', ''],
];
const images = (w, h, W, H) => `<span class="globalnav-image-regular globalnav-link-image"><svg width="${w}" height="${h}" viewBox="0 0 ${w} ${h}"></svg></span><span class="globalnav-image-compact globalnav-link-image"><svg width="${W}" height="${H}" viewBox="0 0 ${W} ${H}"></svg></span>`;
const row = (href, term, rest) => `<li class="globalnav-searchresults-list-item" role="listitem"><a class="globalnav-searchresults-list-link" href="${href}"><span class="globalnav-searchresults-list-icon">${images(9, 16, 13, 25)}</span><span class="globalnav-searchresults-list-text"><span class="globalnav-searchresults-list-text-highlight">${term}</span>${rest}</span></a></li>`;
const section = (title, rows) => `<div class="globalnav-searchresults-container"><h2 class="globalnav-searchresults-header">${title}</h2><ul class="globalnav-searchresults-list" role="list">${rows.join('')}</ul></div>`;
const COPY = `<div>${section('Suggested Links', [row('#iphone-18-pro', 'iPh', 'one 18 Pro'), row('#iphone-air', 'iPh', 'one Air')])}${section('Suggested Searches', [row('#search-iphone', 'iph', 'one'), row('#search-iphone-case', 'iph', 'one case')])}</div>`;
const inject = { inject: COPY, into: `${FLY} .globalnav-searchresults` };
const LINKS = '.globalnav-searchresults-container:first-child';
const SEARCHES = '.globalnav-searchresults-container:last-child';

export default {
  ...search,
  probes: [
    ...search.probes.filter(p => ['nav', 'nav.phone'].includes(p.name)),
    { name: 'search', set: SET, select: '#globalnav .globalnav-search', props: ['display'] },
    { name: 'current', select: `${FLY} .globalnav-searchresults-current`, props: ['display'] },
    { name: 'field', select: `${FLY} .globalnav-searchfield`, props: [...SPACE, 'height'], box: true, position: true },
    { name: 'results', select: `${FLY} .globalnav-searchresults`, props: SPACE, box: true, position: true },
    { name: 'links', ...inject, select: LINKS, props: SPACE, box: true, position: true },
    { name: 'searches', ...inject, select: SEARCHES, props: SPACE, box: true, position: true },
    { name: 'links.header', ...inject, select: `${LINKS} .globalnav-searchresults-header`, props: [...TYPE, ...SPACE, 'opacity'], box: true, position: true },
    { name: 'searches.header', ...inject, select: `${SEARCHES} .globalnav-searchresults-header`, props: [...TYPE, ...SPACE, 'opacity'], box: true, position: true },
    { name: 'links.list', ...inject, select: `${LINKS} .globalnav-searchresults-list`, props: [...SPACE, 'list-style-type'], box: true, position: true },
    { name: 'links.item', ...inject, select: `${LINKS} .globalnav-searchresults-list-item:first-child`, props: [...SPACE, 'display', 'opacity'], box: true, position: true },
    { name: 'links.item2', ...inject, select: `${LINKS} .globalnav-searchresults-list-item:last-child`, props: SPACE, box: true, position: true },
    { name: 'searches.item', ...inject, select: `${SEARCHES} .globalnav-searchresults-list-item:first-child`, props: SPACE, box: true, position: true },
    { name: 'searches.item2', ...inject, select: `${SEARCHES} .globalnav-searchresults-list-item:last-child`, props: SPACE, box: true, position: true },
    { name: 'link', ...inject, select: `${LINKS} .globalnav-searchresults-list-link`, props: [...TYPE, ...SPACE, 'display', 'align-items', 'text-decoration-line', 'background-color', 'border-top-left-radius', ...MOTION], box: true, position: true },
    { name: 'text', ...inject, select: `${LINKS} .globalnav-searchresults-list-text`, props: TYPE, box: true, position: true },
    { name: 'highlight', ...inject, select: `${LINKS} .globalnav-searchresults-list-text-highlight`, props: [...TYPE, 'background-color', 'display', ...MOTION], box: true, position: true },
    { name: 'links.icon', ...inject, select: `${LINKS} .globalnav-searchresults-list-icon`, props: [...SPACE, 'display', 'fill'], box: true, position: true },
    { name: 'searches.icon', ...inject, select: `${SEARCHES} .globalnav-searchresults-list-icon`, props: [...SPACE, 'display', 'fill'], box: true, position: true },
    { name: 'searches.image', ...inject, select: `${SEARCHES} .globalnav-image-regular svg`, props: ['width', 'height', 'fill', 'display'], box: true, position: true, widths: WIDE },
    { name: 'searches.image-compact', ...inject, select: `${SEARCHES} .globalnav-image-compact svg`, props: ['width', 'height', 'fill', 'display'], box: true, position: true, widths: PHONE },
    // hovered and focused
    { name: 'link.states', ...inject, select: `${LINKS} .globalnav-searchresults-list-link`, props: ['color', 'background-color', ...FOCUS], states: ['hover', 'focus-visible'] },
    { name: 'highlight.states', ...inject, select: `${LINKS} .globalnav-searchresults-list-text-highlight`, stateOn: `${LINKS} .globalnav-searchresults-list-link`, props: ['color'], states: ['hover', 'focus-visible'] },
    { name: 'searches.icon-states', ...inject, select: `${SEARCHES} .globalnav-searchresults-list-icon`, stateOn: `${SEARCHES} .globalnav-searchresults-list-link`, props: ['fill'], states: ['hover', 'focus-visible'] },
  ],
};
