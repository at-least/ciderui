// Apple's developer documentation (DocC), the crash report's article again,
// as a page: what ciderui's docs set builds that docc did not read. Its
// layout at each of DocC's own steps (its stylesheets' media queries: 735,
// 767, 1023, 1250 and 1499px, each read on both sides): the navigator's
// column beside the content's, the content's column and its room; the
// documentation's bar (its hairline, its content's room, the button that
// opens and closes the navigator, at every width); the hero (its room, the
// breadcrumbs, a crumb's separator, a crumb's link hovered, the current
// one, the eyebrow and the title's room); the content's top and its
// elements docc did not read (h3, a list, its item, an inline image, a
// listing's and a note's room, its label's and its text's); the See Also
// section (its ground, its title, a group's title, a link with its icon,
// hovered, its abstract); the navigator, read where it shows, from 1024px
// (the technology's title, a group's title, an item nested, its depth's
// spacer, its toggle, its icon, the current item).
// docc-page-dark reads the same in the dark.
import { TYPE, BOX, LINK, FLEX } from './common.mjs';

const ROOM = ['padding-top', 'padding-right', 'padding-bottom', 'padding-left', 'margin-top', 'margin-bottom', 'margin-left', 'margin-right', 'max-width'];
const NAVIGATOR = [1600, 1500, 1499, 1251, 1250, 1200, 1024];
export const probes = [
  { name: 'sidebar', select: '.adjustable-sidebar-width > .sidebar', props: ['display', 'width', 'max-width', 'min-width'], box: true },
  { name: 'content', select: '.adjustable-sidebar-width > .content', props: ['display', 'min-width'], box: true },
  { name: 'container', select: '.doc-content > .container', props: ROOM, box: true },
  { name: 'primary', select: '.primary-content', props: ROOM, box: true },
  { name: 'bar.rule', select: 'nav.documentation-nav .nav__background', pseudo: '::after', props: ['content', 'position', 'height', 'bottom', 'width', 'background-color'] },
  { name: 'bar.content', select: 'nav.documentation-nav .nav-content', props: [...ROOM, ...FLEX, 'height'], box: true },
  { name: 'bar.title.link', select: 'nav.documentation-nav a.nav-title', props: [...TYPE, ...LINK, 'display', 'align-items'], box: true, states: ['hover'] },
  { name: 'bar.toggle.wrapper', select: 'nav.documentation-nav .sidenav-toggle-wrapper', props: [...ROOM, 'display', 'width', 'height'], box: true },
  { name: 'bar.toggle', select: 'nav.documentation-nav button.sidenav-toggle', props: [...ROOM, 'display', 'width', 'height', 'color', 'background-color', 'border-top-left-radius'], box: true, states: ['hover'] },
  { name: 'bar.toggle.icon', select: 'nav.documentation-nav button.sidenav-toggle svg', props: ['width', 'height', 'color', 'fill', 'stroke'], box: true },
  { name: 'hero', select: '.documentation-hero', props: [...BOX, 'color'], box: true },
  { name: 'hero.content', select: '.documentation-hero__content', props: ROOM, box: true },
  { name: 'hierarchy', select: '.documentation-hero ul.hierarchy', props: [...TYPE, ...FLEX, ...ROOM, 'list-style-type'], box: true },
  { name: 'hierarchy.item', select: '.documentation-hero .hierarchy-item:nth-child(2)', props: [...TYPE, ...ROOM, 'display', 'align-items'], box: true },
  { name: 'hierarchy.separator', select: '.documentation-hero .hierarchy-item:nth-child(2)', pseudo: '::before', props: [...TYPE, 'content', 'display', 'width', 'margin-left', 'margin-right'] },
  { name: 'hierarchy.link', select: '.documentation-hero .hierarchy a.nav-menu-link', props: [...TYPE, ...LINK], box: true, states: ['hover'] },
  { name: 'hierarchy.current', select: '.documentation-hero .hierarchy .current', props: [...TYPE, ...LINK], box: true },
  { name: 'topictitle', select: '.documentation-hero .topictitle', props: ROOM, box: true },
  { name: 'eyebrow.room', select: '.documentation-hero .topictitle .eyebrow', props: ['display', 'margin-top', 'margin-bottom'], box: true },
  { name: 'title.box', select: '.documentation-hero h1.title', props: [...TYPE, 'margin-top', 'margin-bottom'], box: true },
  { name: 'abstract.box', select: '.documentation-hero .abstract', props: [...TYPE, 'margin-top', 'margin-bottom'], box: true },
  { name: 'content.top', select: '.primary-content > .content', props: ['margin-top', 'padding-top'], box: true },
  { name: 'content.h2', select: '.primary-content .content > h2', props: [...TYPE, 'margin-top', 'margin-bottom'], box: true },
  { name: 'content.h3', select: '.primary-content .content > h3', props: [...TYPE, 'margin-top', 'margin-bottom'], box: true },
  { name: 'content.anchor', select: '.primary-content .content > h3 a.header-anchor', props: [...TYPE, ...LINK, 'cursor'], states: ['hover'] },
  { name: 'content.p', select: '.primary-content .content > p:not(.inline-image-container)', props: [...TYPE, 'margin-top', 'margin-bottom'], box: true },
  { name: 'content.ul', select: '.primary-content .content > ul', props: [...ROOM, 'padding-inline-start', 'list-style-type', 'list-style-position'], box: true },
  { name: 'content.li', select: '.primary-content .content > ul > li', props: [...TYPE, ...ROOM, 'display'], box: true },
  { name: 'content.li.next', select: '.primary-content .content > ul > li + li', props: ['margin-top'], box: true },
  { name: 'content.li.p', select: '.primary-content .content > ul > li > p', props: [...TYPE, 'margin-top', 'margin-bottom'], box: true },
  { name: 'content.image', select: '.primary-content .content > p.inline-image-container', props: [...ROOM, 'text-align'], box: true },
  { name: 'content.image.img', select: '.primary-content .content > p.inline-image-container img', props: ['max-width', 'width', 'height', 'display', 'margin-left', 'margin-right'], box: true },
  { name: 'content.listing', select: '.primary-content .content > .code-listing', props: ['margin-top', 'margin-bottom'], box: true },
  { name: 'content.note', select: '.primary-content .content > aside.note', props: ['margin-top', 'margin-bottom'], box: true },
  { name: 'content.note.label', select: '.primary-content .content > aside.note p.label', props: ['margin-top', 'margin-bottom'], box: true },
  { name: 'content.note.text', select: '.primary-content .content > aside.note p:not(.label)', props: ['margin-top', 'margin-bottom'], box: true },
  { name: 'seealso', select: 'section.contenttable', props: [...BOX, 'color'], box: true },
  { name: 'seealso.container', select: 'section.contenttable > .container', props: ROOM, box: true },
  { name: 'seealso.title', select: 'section.contenttable h2.title', props: [...TYPE, ...ROOM, 'border-bottom-width', 'border-bottom-style', 'border-bottom-color'], box: true },
  { name: 'seealso.section', select: 'section.contenttable .contenttable-section', props: [...ROOM, ...FLEX, 'border-top-width', 'border-top-style', 'border-top-color'], box: true },
  { name: 'seealso.section.head', select: 'section.contenttable .contenttable-section > .section-title', props: [...ROOM, 'width', 'flex-basis', 'max-width'], box: true },
  { name: 'seealso.section.title', select: 'section.contenttable .contenttable-title', props: [...TYPE, ...ROOM], box: true },
  { name: 'seealso.section.content', select: 'section.contenttable .contenttable-section > .section-content', props: [...ROOM, 'width', 'flex-basis'], box: true },
  { name: 'seealso.block', select: 'section.contenttable .link-block', props: ROOM, box: true },
  { name: 'seealso.block.next', select: 'section.contenttable .link-block + .link-block', props: ['margin-top', 'padding-top'], box: true },
  { name: 'seealso.link', select: 'section.contenttable .link-block a.link', props: [...TYPE, ...LINK, 'display', 'align-items'], box: true, states: ['hover'] },
  { name: 'seealso.icon', select: 'section.contenttable .link-block .topic-icon-wrapper', props: [...ROOM, 'display', 'width', 'height', 'flex-basis', 'align-items', 'justify-content'], box: true },
  { name: 'seealso.icon.svg', select: 'section.contenttable .link-block .topic-icon-wrapper svg', props: ['width', 'height', 'color', 'fill'], box: true },
  { name: 'seealso.abstract', select: 'section.contenttable .link-block .abstract', props: [...TYPE, ...ROOM], box: true },
  { name: 'navigator.inner', widths: NAVIGATOR, select: '.navigator-card-inner', props: [...ROOM, 'display', 'flex-direction'], box: true },
  { name: 'navigator.technology', widths: NAVIGATOR, select: '.navigator a.technology-title', props: [...TYPE, ...LINK, ...ROOM, 'display', 'align-items', 'border-top-left-radius'], box: true, states: ['hover'] },
  { name: 'navigator.technology.title', widths: NAVIGATOR, select: '.navigator a.technology-title h2', props: [...TYPE, ...ROOM], box: true },
  { name: 'navigator.group', widths: NAVIGATOR, select: '.navigator-card-item.is-group', props: [...TYPE, ...ROOM, 'height', 'border-top-width', 'border-top-style', 'border-top-color'], box: true },
  { name: 'navigator.group.title', widths: NAVIGATOR, select: '.navigator-card-item.is-group h3', props: [...TYPE, ...ROOM], box: true },
  { name: 'navigator.row', widths: NAVIGATOR, select: '.navigator-card-item:not(.is-group):not(.active) .head-wrapper', props: [...ROOM, ...FLEX, 'height'], box: true },
  { name: 'navigator.spacer', widths: NAVIGATOR, select: '.navigator-card-item[data-nesting-index="1"]:not(.is-group) .depth-spacer', props: ['width', 'height', 'flex-basis', 'margin-left', 'margin-right'], box: true },
  { name: 'navigator.spacer.nested', widths: NAVIGATOR, select: '.navigator-card-item[data-nesting-index="2"] .depth-spacer', props: ['width', 'height', 'flex-basis', 'margin-left', 'margin-right'], box: true },
  { name: 'navigator.toggle', widths: NAVIGATOR, select: '.navigator-card-item button.tree-toggle', props: [...ROOM, 'display', 'width', 'height', 'color', 'align-items', 'justify-content'], box: true, states: ['hover'] },
  { name: 'navigator.toggle.icon', widths: NAVIGATOR, select: '.navigator-card-item button.tree-toggle[aria-expanded="false"] svg', props: ['width', 'height', 'color', 'fill', 'transform', 'transition-property', 'transition-duration'], box: true },
  { name: 'navigator.toggle.open', widths: NAVIGATOR, select: '.navigator-card-item button.tree-toggle[aria-expanded="true"] svg', props: ['width', 'height', 'color', 'fill', 'transform'], box: true },
  { name: 'navigator.icon', widths: NAVIGATOR, select: '.navigator-card-item .navigator-icon-wrapper', props: [...ROOM, 'width', 'height', 'display'], box: true },
  { name: 'navigator.icon.svg', widths: NAVIGATOR, select: '.navigator-card-item .navigator-icon', props: ['width', 'height', 'color', 'fill'], box: true },
  { name: 'navigator.current.row', widths: NAVIGATOR, select: '.navigator-card-item.active .head-wrapper', props: [...ROOM, 'height', 'background-color', 'border-top-left-radius'], box: true },
];
export default {
  url: 'https://developer.apple.com/documentation/xcode/examining-the-fields-in-a-crash-report',
  wait: 4000,
  widths: [1600, 1500, 1499, 1251, 1250, 1200, 1024, 1023, 900, 768, 767, 736, 735, 390],
  probes,
};
