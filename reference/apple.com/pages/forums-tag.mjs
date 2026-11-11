// @external: Apple Developer Forums' page for one tag (`/forums/tags/swiftui`), signed
// out: its header (the tag's name, its line, its links) and the count over its list.
import { TYPE, BOX, FLEX, LINK } from './common.mjs';

export default {
  url: 'https://developer.apple.com/forums/tags/swiftui',
  wait: 3000,
  widths: [1600, 1200, 900, 735, 734, 390],
  probes: [
    { name: 'header', select: 'section.tag-overview', props: [...BOX, 'width'], box: true },
    { name: 'header.content', select: 'section.tag-overview .tag-content', props: [...BOX, ...FLEX], box: true },
    { name: 'header.title', select: 'section.tag-overview .tag-title-and-rss h1, section.tag-overview .tag-title-and-rss h2', props: [...TYPE, 'margin-top', 'margin-bottom'], box: true },
    { name: 'header.info', select: 'section.tag-overview p.tag-info', props: [...TYPE, 'margin-top', 'margin-bottom', 'max-width'], box: true },
    { name: 'header.links', select: 'section.tag-overview p.tag-links', props: [...TYPE, 'margin-top', 'margin-bottom'], box: true },
    { name: 'header.link', select: 'section.tag-overview p.tag-links a', props: [...TYPE, ...LINK], states: ['hover'] },
    { name: 'count.bar', select: 'main section div.header', props: [...BOX, ...FLEX], box: true },
    { name: 'count', select: 'main section div.header span.search-results', props: [...TYPE, 'margin-top', 'margin-bottom'], box: true },
  ],
};
