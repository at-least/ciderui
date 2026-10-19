// apple.com's Newsroom, a press release (Apple TV 4K, 2026-10-13: a hero
// image, an inline image, subheads, small print with a list, footnotes, the
// related stories below): the article's column and rhythm, its header, its
// copy and what markdown writes in it (a paragraph, a link, a footnote's
// mark, a subhead, a list, a figure and its caption), its footnotes, its
// related rows, and the local nav it has (Newsroom's, ac-ln), at the steps
// it takes (735 and 1069) and the kit's widths. What is a press release's
// alone (the dateline, the share sheet, the downloads, the press contacts,
// About Apple) is not read.
import { TYPE, BOX, BLUR, LINK, PAGE, FLEX, TEXT } from './common.mjs';

const COLUMN = ['width', 'max-width', 'margin-left', 'margin-right', 'padding-left', 'padding-right'];
const SPACE = ['margin-top', 'margin-bottom', 'padding-top', 'padding-bottom'];
// how a word too long for its line breaks, if it does
const WRAP = ['overflow-wrap', 'word-break', 'hyphens', 'white-space'];

export default {
  url: 'https://www.apple.com/newsroom/2026/10/apple-unveils-the-new-apple-tv-4k/',
  widths: [1600, 1200, 1069, 1068, 900, 735, 734, 390],
  probes: [
    { name: 'page.body', select: 'body', props: PAGE },
    // the column: each component's content box, the copy's and a figure's
    { name: 'column.header', select: '.article-header .component-content', props: COLUMN, box: true },
    { name: 'column.copy', select: '.pagebody .component-content', props: COLUMN, box: true },
    // the text itself: below 735px the copy's content box is the figure's
    // (414px at most) and each paragraph takes its own margins
    { name: 'column.text', select: '.pagebody .component-content > .pagebody-copy', props: COLUMN, box: true },
    { name: 'column.figure', select: '.pagebody', props: COLUMN, box: true },
    // the local nav (Newsroom's own, not a product page's)
    { name: 'localnav', select: '#ac-localnav', props: ['height', 'position', 'top', 'z-index'], box: true },
    { name: 'localnav.background', select: '#ac-localnav .ac-ln-background', props: ['background-color', ...BLUR] },
    { name: 'localnav.content', select: '#ac-localnav .ac-ln-content', props: ['max-width', 'padding-left', 'padding-right'], box: true },
    { name: 'localnav.title', select: '#ac-localnav .ac-ln-title a', props: [...TYPE, ...LINK], states: ['hover'] },
    // the header: its parts' room, the eyebrow (category, date), the
    // headline, the subhead
    { name: 'header', select: '.article-header', props: SPACE, box: true },
    { name: 'header.eyebrow', select: '.article-header .category.component', props: SPACE, box: true },
    { name: 'header.headline', select: '.article-header .pagetitle.component', props: SPACE, box: true },
    { name: 'header.subhead', select: '.article-header .article-subhead.component', props: SPACE, box: true },
    { name: 'eyebrow', select: '.category-eyebrow', props: [...BOX, ...FLEX] },
    { name: 'eyebrow.category', select: '.category-eyebrow__category', props: [...TYPE, ...TEXT, 'text-transform'] },
    { name: 'eyebrow.date', select: '.category-eyebrow__date', props: [...TYPE, ...TEXT, 'margin-top'] },
    { name: 'headline', select: 'h1.hero-headline', props: [...TYPE, 'margin-top', 'margin-bottom', 'text-align', ...WRAP] },
    { name: 'subhead', select: '.article-subhead .component-content', props: [...TYPE, 'margin-top', 'margin-bottom', ...WRAP] },
    // the hero figure, its picture and its caption
    { name: 'figure', select: 'figure.image-fullbleed', props: [...COLUMN, 'margin-top', 'margin-bottom'], box: true },
    { name: 'figure.img', select: 'figure.image-fullbleed img.picture-image', props: ['max-width', 'border-top-left-radius'], box: true },
    { name: 'figure.caption', select: 'figure.image-fullbleed .image-caption', props: [...TYPE, ...SPACE, 'text-align', ...WRAP] },
    // the caption's room: the text's column, under the picture
    { name: 'figure.description', select: 'figure.image-fullbleed .image-description', props: [...COLUMN, 'margin-top'], box: true },
    { name: 'figure.inline', select: 'figure.image-inline', props: [...COLUMN, 'margin-top', 'margin-bottom'], box: true },
    // the copy: a paragraph of it, a paragraph a page adds there, a link, a
    // footnote's mark, bold, a subhead and its strong
    { name: 'copy', select: '.pagebody .component-content > .pagebody-copy', props: [...TYPE, 'margin-top', 'margin-bottom', ...WRAP] },
    { name: 'copy.p', inject: '<p>Sample text</p>', into: '.pagebody .component-content', props: [...TYPE, 'margin-top', 'margin-bottom'] },
    { name: 'copy.link', select: '.pagebody-copy a', props: [...TYPE, ...LINK, ...WRAP], states: ['hover'] },
    { name: 'copy.sup', select: '.pagebody-copy sup', props: [...TYPE, 'font-feature-settings', 'vertical-align', 'position', 'top'] },
    { name: 'copy.strong', select: '.pagebody-small .pagebody-copy li strong', props: TYPE },
    { name: 'h2', select: 'h2.pagebody-header', props: [...TYPE, 'margin-top', 'margin-bottom', ...WRAP] },
    { name: 'h2.strong', select: 'h2.pagebody-header strong', props: TYPE },
    // the small print's list
    { name: 'list', select: '.pagebody-small .pagebody-copy ul', props: ['margin-top', 'margin-bottom', 'padding-left', 'list-style-type', 'list-style-position', 'margin-left'] },
    { name: 'list.li', select: '.pagebody-small .pagebody-copy li', props: [...TYPE, 'margin-top', 'margin-bottom'] },
    { name: 'list.link', select: '.pagebody-small .pagebody-copy li a', props: [...TYPE, ...LINK], states: ['hover'] },
    // the footnotes
    { name: 'footnotes', select: '.sosumi', props: [...SPACE, 'border-top-width', 'border-top-style', 'border-top-color'] },
    { name: 'footnotes.ol', select: '.sosumi ol', props: [...TYPE, 'margin-top', 'margin-bottom', 'padding-left', 'list-style-type', 'margin-left'] },
    { name: 'footnotes.li', select: '.sosumi li', props: [...TYPE, 'margin-top', 'margin-bottom', ...WRAP] },
    { name: 'footnotes.link', select: '.sosumi li a', props: [...TYPE, ...LINK], states: ['hover'] },
    // the related stories: a heading over ruled rows, each a link with a
    // picture, a category, a title and a date
    { name: 'related.heading', select: '.article-list__heading', props: [...TYPE, 'margin-top', 'margin-bottom', 'text-align'] },
    { name: 'related.list', select: '.article-list', props: [...SPACE, ...COLUMN], box: true },
    { name: 'related.item', select: '.article-list__item', props: [...BOX, 'border-bottom-width'], box: true },
    { name: 'related.link', select: '.article-list__item__link', props: [...FLEX, ...LINK, 'cursor'], box: true, states: ['hover'] },
    { name: 'related.img', select: '.article-list__item__img-wrapper', props: [...BOX, 'width', 'height', 'flex-shrink'], box: true },
    { name: 'related.category', select: '.article-list__item__category', props: [...TYPE, ...TEXT, 'text-transform', 'margin-bottom'] },
    { name: 'related.title', select: '.article-list__item__title', props: [...TYPE, 'margin-top', 'margin-bottom', ...LINK], states: ['hover'] },
    { name: 'related.date', select: '.article-list__item__date', props: [...TYPE, 'margin-top'] },
    // the room under the related stories, before the footer's band (the
    // footer tile, Newsroom's own, not read)
    { name: 'footertile', select: '.footertile', props: ['margin-top', 'margin-bottom'] },
  ],
};
