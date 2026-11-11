// @external: apple.com's search page for a query nothing matches: its empty state (a
// title and a line under it) and the product tiles it offers instead, a picture
// over its name.
import { TYPE, BOX, FLEX } from './common.mjs';

export default {
  url: 'https://www.apple.com/us/search/xqzvwqk?src=globalnav',
  wait: 3000,
  probes: [
    { name: 'null', select: '.rf-serp-nullsearch', props: [...BOX, 'width'], box: true },
    { name: 'null.heading', select: '.rf-serp-nullsearch-heading', props: [...BOX, ...FLEX, 'text-align'], box: true },
    { name: 'null.title', select: '.rf-serp-nullsearch-title', props: [...TYPE, 'margin-top', 'margin-bottom', 'text-align'], box: true },
    { name: 'null.desc', select: '.rf-serp-nullsearch-desc', props: [...TYPE, 'margin-top', 'margin-bottom', 'text-align'], box: true },
    { name: 'null.results', select: '.rf-serp-nullsearch-results', props: [...BOX], box: true },
    { name: 'section.label', select: '.rf-serp-product-sectionlabel', props: [...TYPE, 'margin-top', 'margin-bottom'], box: true },
    { name: 'tile', select: '.rf-serp-product-tile', props: [...BOX, ...FLEX, 'width', 'text-align', 'overflow'], box: true, states: ['hover'] },
    { name: 'tile.img', select: '.rf-serp-product-tile img', props: ['width', 'height', 'object-fit', 'border-top-left-radius', 'margin-bottom', 'background-color'], box: true },
    { name: 'tile.title', select: '.rf-serp-product-title', props: [...TYPE, 'margin-top', 'text-align', '-webkit-line-clamp', 'overflow'], box: true },
    { name: 'tile.link', select: '.rf-serp-product-tile a, a:has(> .rf-serp-product-tile)', props: [...TYPE, 'text-decoration-line'], states: ['hover'] },
  ],
};
