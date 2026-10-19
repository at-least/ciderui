// The iPhone page, a light one: its links (the "more" link, its copy and its
// chevron; a section header's link) and its ribbon, the message above the
// content.
import { TYPE, BOX, LINK, PAGE, TEXT } from './common.mjs';
export default {
  url: 'https://www.apple.com/iphone/',
  probes: [
    { name: 'page.body', select: 'body', props: PAGE },
    { name: 'link.more', select: 'main a.icon-wrapper:not(.section-header-cta-link)', props: [...TYPE, ...LINK, ...TEXT], states: ['hover', 'focus-visible'] },
    { name: 'link.more-copy', select: 'main a.icon-wrapper:not(.section-header-cta-link) .icon-copy', stateOn: 'main a.icon-wrapper:not(.section-header-cta-link)', props: [...TYPE, ...LINK], states: ['hover'] },
    { name: 'link.more-icon', select: 'main a.icon-wrapper:not(.section-header-cta-link) .icon-after', props: ['display', 'font-size', 'line-height', 'padding-left', 'margin-left', 'vertical-align', 'position', 'top'], box: true },
    { name: 'link.more-glyph', select: 'main a.icon-wrapper:not(.section-header-cta-link) .icon-after', pseudo: '::after', props: ['content', 'font-family', 'font-size', 'line-height', 'font-weight', 'display', 'padding-left', 'margin-left', 'top', 'position', 'vertical-align', 'color', 'text-decoration-line'] },
    // a link in running text: apple.com's own `a`, on markup added to the page
    { name: 'link.inline', inject: '<p>Text with <a href="#cui">a link</a> in it.</p>', into: 'main', select: 'a', props: [...TYPE, ...LINK], states: ['hover', 'focus-visible', 'active'] },
    { name: 'link.section-header', select: 'main a.section-header-cta-link', props: [...TYPE, ...LINK], states: ['hover'] },
    { name: 'ribbon', select: '.ribbon', props: [...TYPE, ...BOX], box: true },
    { name: 'ribbon.content', select: '.ribbon .ribbon-content', props: [...TYPE, ...BOX, 'text-align'], box: true },
    { name: 'ribbon.link', select: '.ribbon .ribbon-link', props: [...TYPE, ...LINK], states: ['hover'] },
  ],
};
