// The iPhone page, a light one: its links (the "more" link, its copy and its
// chevron; a section header's link) and its ribbon, the message above the
// content.
import { TYPE, BOX, LINK, PAGE, TEXT, FLEX, MOTION } from './common.mjs';
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
    // the content's width; a section's padding (light, alternate), its header
    { name: 'content', select: 'main .viewport-content', props: ['width', 'max-width', 'margin-left', 'margin-right', 'padding-left', 'padding-right'], box: true },
    { name: 'section.incentive', select: 'main section.section-incentive', props: BOX },
    { name: 'section.image-accordion', select: 'main section.section-image-accordion', props: BOX },
    { name: 'section.header', select: 'main .section-header', props: [...BOX, ...FLEX], box: true },
    { name: 'section.headline', select: 'main .section-header-headline', props: [...TYPE, 'margin-bottom'] },
    { name: 'ribbon', select: '.ribbon', props: [...TYPE, ...BOX], box: true },
    { name: 'ribbon.content', select: '.ribbon .ribbon-content', props: [...TYPE, ...BOX, 'text-align'], box: true },
    { name: 'ribbon.link', select: '.ribbon .ribbon-link', props: [...TYPE, ...LINK], states: ['hover'] },
    // a light page's gallery paddles (the first enabled one, and a disabled
    // one) and a media card gallery's dots
    { name: 'paddle', select: 'main .paddlenav-arrow-next:not([disabled])', props: [...BOX, 'width', 'height', 'color', 'opacity', 'cursor', ...MOTION], box: true, states: ['hover', 'active', 'focus-visible'] },
    { name: 'paddle.icon', select: 'main .paddlenav-arrow-next:not([disabled]) svg', props: ['width', 'height', 'fill', 'color'], box: true },
    { name: 'paddle.disabled', select: 'main .paddlenav-arrow[disabled]', props: ['opacity', 'background-color', 'color', 'cursor', 'pointer-events'] },
    { name: 'gallery.dot', select: 'main .dotnav-link:not(.current)', props: ['width', 'height', 'background-color', 'border-top-left-radius', 'padding-top', ...MOTION], box: true, states: ['hover'] },
    { name: 'gallery.dot-current', select: 'main .dotnav-link.current', props: ['width', 'height', 'background-color', 'border-top-left-radius', 'padding-top'], box: true },
    { name: 'gallery.dot-fill', select: 'main .dotnav-link.current', pseudo: '::after', props: ['width', 'height', 'background-color', 'border-top-left-radius', 'position', 'content'] },
  ],
};
