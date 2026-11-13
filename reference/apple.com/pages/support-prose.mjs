// @external: an Apple Support article's text (the owner, 2026-11-13: what consumers write
// by hand, measured on Apple's sites; totality's store and website write markdown's
// elements by hand): its paragraphs, headings, lists and their items, a picture in a
// step, bold. The article with each of them: support's own (`support`) has no heading,
// numbered list or picture. Its rules are by the element before (`p + p`, `h2`), as
// markdown writes them: no class of the page's own reaches a consumer's text.
import { TYPE, LINK } from './common.mjs';

const FLOW = ['margin-top', 'margin-bottom'];
export default {
  url: 'https://support.apple.com/en-us/102631',
  wait: 3000,
  probes: [
    { name: 'p', select: 'p.gb-paragraph + p.gb-paragraph', props: [...TYPE, ...FLOW] },
    { name: 'h2', select: 'h2.gb-header', props: [...TYPE, ...FLOW] },
    { name: 'h2.next', select: 'h2.gb-header + p.gb-paragraph', props: FLOW },
    { name: 'h3', select: 'h3.gb-header', props: [...TYPE, ...FLOW] },
    { name: 'ul', select: 'p.gb-paragraph + ul.gb-list', props: [...FLOW, 'padding-left', 'margin-left', 'list-style-type', 'list-style-position'] },
    { name: 'ol', select: 'ol.gb-list', props: [...FLOW, 'padding-left', 'margin-left', 'list-style-type', 'list-style-position'] },
    { name: 'li', select: 'ul.gb-list > li', props: [...TYPE, ...FLOW, 'padding-left'] },
    { name: 'li.next', select: 'ul.gb-list > li + li', props: FLOW },
    { name: 'ol.li', select: 'ol.gb-list > li', props: [...TYPE, ...FLOW] },
    { name: 'li.marker', select: 'ol.gb-list > li', pseudo: '::marker', props: TYPE },
    { name: 'list.next', select: 'ul.gb-list + p.gb-paragraph, ol.gb-list + p.gb-paragraph', props: FLOW },
    { name: 'img', select: 'img.gb-image', props: [...FLOW, 'width', 'max-width', 'height', 'display', 'border-top-left-radius'], box: true },
    { name: 'b', select: 'p.gb-paragraph b', props: TYPE },
    { name: 'link', select: 'p.gb-paragraph a', props: [...TYPE, ...LINK], states: ['hover'] },
  ],
};
