// A product page, iPhone 18 Pro's, a dark page: apple.com's base type scale
// (the classes its stylesheets define, added to the page), the global nav's
// dark side, the content's width, a section,
// its header, a dark and an alternate section, a rounded tile, the gallery's
// paddles, the "more" link, a stat, a footnote, and the floating local nav,
// which shows once the page is scrolled past its hero.
import { TYPE, BOX, MOTION, BLUR, LINK, PAGE } from './common.mjs';
const TYPOGRAPHY = ['headline-super', 'headline-elevated', 'headline', 'headline-reduced', 'eyebrow-super',
  'intro-elevated', 'intro', 'body', 'body-reduced', 'body-reduced-tight', 'callout-super', 'callout-elevated',
  'callout-base', 'callout-cta', 'caption-elevated', 'caption', 'caption-cta', 'caption-sosumi',
  'caption-elevated-sosumi', 'label', 'section-header-eyebrow', 'section-header-headline', 'section-header-link'];
export default {
  url: 'https://www.apple.com/iphone-18-pro/',
  probes: [
    // a class the page does not define reads as this paragraph does
    { name: 'type.none', inject: '<p>Sample text</p>', into: 'main', props: TYPE },
    ...TYPOGRAPHY.map(c => ({ name: `type.${c}`, inject: `<p class="typography-${c}">Sample text</p>`, into: 'main', props: TYPE })),
    { name: 'page.body', select: 'body', props: PAGE },
    { name: 'page.strong', select: 'main strong, main b, .rs-page-content strong, body strong', props: TYPE },
    // the page is dark (body.globalheader-dark): the global nav's dark side
    { name: 'nav', select: '#globalnav', props: ['background-color', ...BLUR] },
    { name: 'nav.link', select: '#globalnav .globalnav-item:not(.globalnav-item-apple) .globalnav-link', props: TYPE, states: ['hover'] },
    { name: 'content', select: 'main .viewport-content', props: ['width', 'max-width', 'margin-left', 'margin-right', 'padding-left', 'padding-right'], box: true },
    { name: 'section', select: 'main section.section:not(.theme-dark):not(.background-alt)', props: [...BOX, 'padding-top', 'padding-bottom'], box: true },
    { name: 'section.dark', select: 'main section.section.theme-dark', props: [...TYPE, 'background-color'] },
    { name: 'section.alt', select: 'main .background-alt', props: [...TYPE, 'background-color'] },
    { name: 'section.header', select: 'main .section-header', props: [...BOX, 'display', 'align-items', 'justify-content'], box: true },
    { name: 'section.headline', select: 'main .section-header-headline', props: [...TYPE, 'margin-bottom'] },
    { name: 'section.header-link', select: 'main .section-header-cta-list a', props: [...TYPE, ...LINK], states: ['hover'] },
    { name: 'tile.rounded', select: 'main .tile.tile-rounded', props: [...BOX, 'overflow'], box: true },
    { name: 'tile.rounded-light', select: 'main .theme-light .tile.tile-rounded, main .tile.tile-rounded:not(.theme-dark *)', props: [...BOX], box: true },
    { name: 'paddle', select: 'main .paddlenav-arrow', props: [...BOX, 'width', 'height', 'color', 'opacity', ...BLUR, ...MOTION], box: true, states: ['hover', 'active'] },
    { name: 'paddle.icon', select: 'main .paddlenav-arrow svg', props: ['width', 'height', 'fill', 'color'], box: true },
    { name: 'link.more', select: 'main a.icon-wrapper', props: [...TYPE, ...LINK], states: ['hover'] },
    { name: 'link.more-icon', select: 'main a.icon-wrapper .icon-after', props: ['font-family', 'font-size', 'color', 'padding-left', 'margin-left', 'top', 'position', 'content'], pseudo: '::after' },
    { name: 'stat.value', select: 'main .stat-value', props: TYPE },
    { name: 'stat.caption', select: 'main .stat-caption', props: TYPE },
    { name: 'footnote', select: 'main sup.footnote', props: [...TYPE, 'vertical-align', 'top', 'position'] },
    { name: 'localnav', select: '#localnav-container', scroll: 2500, props: [...BOX, 'position', 'top', 'bottom', 'left', 'width', 'height', 'z-index', ...BLUR], box: true },
    { name: 'localnav.wrap', select: '#localnav-container .localnav-wrap', scroll: 2500, props: [...BOX, 'height', ...BLUR], box: true },
    { name: 'localnav.product', select: '#localnav-container .product-name', scroll: 2500, props: TYPE },
    { name: 'localnav.explore', select: '#localnav-container .cta.explore', scroll: 2500, props: [...TYPE, ...BOX], box: true, states: ['hover'] },
    { name: 'localnav.buy', select: '#localnav-container .cta.buy', scroll: 2500, props: [...TYPE, ...BOX], box: true, states: ['hover'] },
  ],
};
