// Apple Developer's releases list: the page; an entry (its picture's column
// and the picture, its title and the link around it, its date, its text and
// a link in it, hovered, the share buttons) and the rule over the next
// entry; the RSS link over the list.
import { TYPE, LINK, PAGE } from './common.mjs';

const ENTRY = '.article-list-wrapper article.article';
export default {
  url: 'https://developer.apple.com/news/releases/',
  wait: 3000,
  probes: [
    { name: 'page', select: 'body', props: PAGE },
    { name: 'entry', select: ENTRY, props: ['padding-top', 'padding-bottom', 'margin-top', 'margin-bottom'], box: true },
    { name: 'entry.picture', select: `${ENTRY} .article-image-container`, props: ['width', 'padding-top', 'padding-right', 'padding-bottom', 'padding-left'], box: true },
    { name: 'entry.image', select: `${ENTRY} img.article-image`, props: ['width', 'height', 'border-top-left-radius'], box: true },
    { name: 'entry.content', select: `${ENTRY} .article-content-container`, props: ['padding-top', 'padding-right', 'padding-bottom', 'padding-left', 'border-top-width', 'border-top-style', 'border-top-color'], box: true },
    { name: 'entry.title.link', select: `${ENTRY} a.article-title`, props: [...TYPE, ...LINK], states: ['hover'] },
    { name: 'entry.title', select: `${ENTRY} a.article-title h2`, props: [...TYPE, 'margin-top', 'margin-bottom'], box: true },
    { name: 'entry.date', select: `${ENTRY} p.article-date`, props: [...TYPE, 'margin-top', 'margin-bottom'] },
    { name: 'entry.text', select: `${ENTRY} .article-text`, props: [...TYPE, 'margin-top'] },
    { name: 'entry.link', select: `${ENTRY} .article-text a`, props: [...TYPE, ...LINK], states: ['hover'] },
    { name: 'entry.share', select: `${ENTRY} button.social-icon`, props: [...TYPE, 'width', 'height', 'margin-right'], box: true },
    { name: 'next', select: `${ENTRY} ~ article.article .article-content-container`, props: ['padding-top', 'margin-top', 'border-top-width', 'border-top-style', 'border-top-color'] },
    { name: 'rss', select: 'a.rss-link', props: [...TYPE, 'width', 'height'], box: true },
  ],
};
