// The article against apple.com's Newsroom (newsroom-article:*): the header's
// room and type, the copy's type and rhythm, a link (hovered alike), a
// footnote's mark, a subhead, a figure and its caption, the list's marks,
// the footnotes; at the kit's widths and both sides of 735 and 1069, where
// it steps. Golden, with their why: the columns' widths and places (an
// apple.com check reads no width), the letter-spacing, what Newsroom's copy
// has none of (h3, h4, code, a block of it, blockquote, table, hr), a list
// at the copy's size, a footnote's mark that is a link, the copy's room
// where no figure leads it, and the colors in a dark block.

import { examples, SIZE } from '../../test/harness.mjs';

const ex = examples(import.meta.url);
const PAGE = 'newsroom-article';
const WIDTHS = [1600, 1200, 1069, 1068, 900, 735, 734, 390];
const LINK = ['text-decoration-line', 'text-decoration-color'];
const TRACKING = 'the system font\'s fitted letter-spacing (tokens.css), not apple.com\'s';
const TEXT = 'the text\'s column: apple.com\'s 653px from 1069px, 576px from 735px, 87.5% of 414px at most below (newsroom-article:column.text, column.header), centered';
const BAND = 'the figures\' column: apple.com\'s 980px, 692px, 414px at most (newsroom-article:figure, column.figure), centered';
const NONE = 'Newsroom\'s copy has none';

export const cases = [
  {
    name: 'article',
    html: ex('article'),
    widths: WIDTHS,
    checks: [
      // the header
      { target: '.cui-article-header', props: ['padding-top'], as: { 'padding-top': 'margin-top' }, apple: [PAGE, 'header.eyebrow'] },
      { target: '.cui-article-eyebrow', props: ['margin-bottom'], apple: [PAGE, 'header.eyebrow'] },
      { target: '.cui-article-category', props: [...SIZE, 'color', 'display', 'text-transform'], apple: [PAGE, 'eyebrow.category'] },
      { target: '.cui-article-date', props: [...SIZE, 'color', 'display', 'margin-top'], apple: [PAGE, 'eyebrow.date'] },
      { target: '.cui-article-headline', props: [...SIZE, 'color', 'margin-top', 'margin-bottom', 'text-align'], apple: [PAGE, 'headline'] },
      { target: '.cui-article-subhead', props: [...SIZE, 'color'], apple: [PAGE, 'subhead'] },
      { target: '.cui-article-subhead', props: ['margin-top'], apple: [PAGE, 'header.subhead'] },
      // the copy
      { target: '.cui-article-body > p', props: [...SIZE, 'color', 'margin-top', 'margin-bottom'], apple: [PAGE, 'copy'] },
      { target: '.cui-article-body > p a', props: [...SIZE, 'color', ...LINK], apple: [PAGE, 'copy.link'] },
      { target: '.cui-article-body > p a', state: 'hover', props: ['color', ...LINK], apple: [PAGE, 'copy.link', 'hover'] },
      { target: '.cui-article-body sup', props: ['font-size', 'position', 'top', 'vertical-align', 'font-feature-settings'], apple: [PAGE, 'copy.sup'] },
      { target: '.cui-article-body strong', props: ['font-weight'], apple: [PAGE, 'copy.strong'] },
      { target: '.cui-article-body > h2', props: [...SIZE, 'color', 'margin-top', 'margin-bottom'], apple: [PAGE, 'h2'] },
      { target: '.cui-article-body > ul', props: ['list-style-type', 'list-style-position'], apple: [PAGE, 'list'] },
      { target: '.cui-article-body > figure', props: ['margin-top', 'margin-bottom'], apple: [PAGE, 'figure'] },
      { target: '.cui-article-body figcaption', props: [...SIZE, 'color', 'text-align'], apple: [PAGE, 'figure.caption'] },
      { target: '.cui-article-body figcaption', props: ['margin-top'], apple: [PAGE, 'figure.description'] },
      // the footnotes
      { target: '.cui-article-footnotes', props: ['margin-top'], apple: [PAGE, 'footnotes'] },
      { target: '.cui-article-footnotes ol', props: [...SIZE, 'color', 'padding-left', 'list-style-type'], apple: [PAGE, 'footnotes.ol'] },
      { target: '.cui-article-footnotes li', props: [...SIZE, 'color', 'margin-top', 'margin-bottom'], apple: [PAGE, 'footnotes.li'] },
      { target: '.cui-article-footnotes a', props: [...SIZE, 'color', ...LINK], apple: [PAGE, 'footnotes.link'] },
      { target: '.cui-article-footnotes a', state: 'hover', props: ['color', ...LINK], apple: [PAGE, 'footnotes.link', 'hover'] },
      // the columns
      { target: '.cui-article-header', props: ['box-width', 'box-left'], golden: TEXT },
      { target: '.cui-article-body > p', props: ['box-width', 'box-left'], golden: TEXT },
      { target: '.cui-article-footnotes', props: ['box-width', 'box-left'], golden: TEXT },
      { target: '.cui-article-body > figure', props: ['box-width', 'box-left'], golden: BAND },
      { target: '.cui-article-body figcaption', props: ['box-width', 'box-left'], golden: `the caption in ${TEXT}` },
      { target: '.cui-article-body figure img', props: ['box-width', 'box-height'], golden: 'the picture the figure\'s width, at its own shape' },
      // the letter-spacing
      ...['.cui-article-category', '.cui-article-date', '.cui-article-headline', '.cui-article-subhead', '.cui-article-body > p', '.cui-article-body > h2', '.cui-article-body figcaption', '.cui-article-footnotes ol']
        .map(target => ({ target, props: ['letter-spacing'], golden: TRACKING })),
      // what is ciderui's
      { target: '.cui-article-body', props: ['margin-top'], golden: 'the copy 44px under the header (36px from 735px, 32px below), the room a figure takes there (newsroom-article:figure), where no figure leads it' },
      { target: '.cui-article-body > ul', props: ['padding-left', 'margin-bottom'], golden: 'the small print\'s indent, 1.1765em (newsroom-article:list, 16.47px at 14px), at the copy\'s size; a paragraph\'s room under it' },
      { target: '.cui-article-body > ul li', props: [...SIZE, 'margin-top', 'margin-bottom'], golden: 'a list at the copy\'s size: apple.com\'s is the small print\'s 14px (newsroom-article:list.li)' },
      { target: '.cui-article-body > ol', props: ['list-style-type', 'padding-left'], golden: 'numbered, the list\'s indent' },
      { target: '.cui-article-body > h3', props: [...SIZE, 'margin-top', 'margin-bottom'], golden: `${NONE}: a related story's title, 19/23 at 700 (17/21 below 1069px, newsroom-article:related.title), 32px down and 8px over its own` },
      { target: '.cui-article-body > h4', props: [...SIZE, 'margin-top', 'margin-bottom'], golden: `${NONE}: the copy's size at 700, a paragraph's room above` },
      { target: '.cui-article-body sup a', props: ['color', 'text-decoration-line'], golden: 'a footnote\'s mark links to its note, in the copy\'s color, not underlined: apple.com\'s mark is no link' },
      { target: '.cui-article-body > p code', props: ['font-family', 'font-size'], golden: `${NONE}: the mono font (tokens.css), 0.85em` },
      { target: '.cui-article-body > pre', props: ['padding-top', 'padding-left', 'border-top-left-radius', 'background-color', 'font-family', ...SIZE, 'letter-spacing', 'overflow-x', 'margin-bottom'], golden: `${NONE}: a block on the alternate ground, a field's 12px corners, 14/20` },
      { target: '.cui-article-body > blockquote', props: ['padding-left', 'color', 'margin-bottom'], golden: `${NONE} (a feature story's pull quote is a page's part): the copy indented as a list, the secondary text` },
      { target: '.cui-article-body > table', props: [...SIZE, 'border-collapse', 'margin-bottom'], golden: `${NONE}: 14/20, the small print's size` },
      { target: '.cui-article-body th', props: ['padding-top', 'padding-right', 'border-bottom-width', 'border-bottom-color', 'font-weight', 'text-align'], golden: `${NONE}: rows ruled in the divider's color, 12px in` },
      { target: '.cui-article-body > hr', props: ['border-top-width', 'border-top-style', 'border-top-color', 'margin-top', 'margin-bottom'], golden: `${NONE}: a rule in the divider's color, a figure's room around it` },
    ],
  },
  {
    name: 'dark',
    html: `<div class="cui-dark">${ex('article')}</div>`,
    checks: [
      { target: '.cui-article-headline', props: ['color'], golden: 'the dark text (#f5f5f7), the contract\'s: Newsroom\'s press releases are light' },
      { target: '.cui-article-category', props: ['color'], golden: 'the dark secondary text (#86868b)' },
      { target: '.cui-article-body > p a', props: ['color'], golden: 'the dark link (#2997ff)' },
      { target: '.cui-article-body > pre', props: ['background-color'], golden: 'the dark alternate ground' },
      { target: '.cui-article-body th', props: ['border-bottom-color'], golden: 'the dark divider (#424245)' },
    ],
  },
  {
    // words too long for a line, which markdown's copy has and Newsroom's
    // prose does not: an import in inline code, a bare URL as a link, a
    // hash in a table's cell
    name: 'long',
    html: `<article class="cui-article"><header class="cui-article-header"><h1 class="cui-article-headline">ciderui</h1></header><div class="cui-article-body">
<p>Import <code>@import "../target/ciderui/components/button/button.css";</code> first.</p>
<p><a href="#">https://www.apple.com/newsroom/2026/10/apple-unveils-the-new-apple-tv-4k/</a></p>
<table><thead><tr><th>file</th><th>its hash</th></tr></thead><tbody><tr><td>logo.svg</td><td><code>PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHdpZHRoPSIxMjgwIg==</code></td></tr></tbody></table>
</div></article>`,
    widths: [1600, 390],
    checks: [
      { target: '.cui-article', props: ['overflow-wrap'], golden: 'a word too long for its line breaks where it must, in the copy, a table\'s cell and the header (anywhere: break-word keeps a cell\'s word whole and the table runs past the column): apple.com\'s breaks none (newsroom-article:copy, headline, … overflow-wrap normal), its prose having no such word' },
      ...['.cui-article-body > p code', '.cui-article-body > p a', '.cui-article-body > table'].map(target => ({ target, props: ['box-width', 'box-left'], golden: 'within the text\'s column: a long word broken in it' })),
    ],
  },
  {
    // markdown's lists in a list item and in a quotation, and its code block
    // in a list item (a step's command), as Zola writes them
    name: 'nested',
    html: `<article class="cui-article"><header class="cui-article-header"><h1 class="cui-article-headline">ciderui</h1></header><div class="cui-article-body">
<ul><li>Outer<ul><li>Inner</li></ul></li></ul>
<ol><li><p>Add the packages:</p><pre><code>npm install --save-dev @tailwindcss/cli tailwindcss</code></pre></li></ol>
<blockquote><p>A quotation with a list:</p><ul><li>Quoted</li></ul></blockquote>
</div></article>`,
    widths: [1600, 390],
    checks: [
      { target: '.cui-article-body > ul ul', props: ['list-style-type', 'padding-left'], golden: `a list in a list item the copy's, its marker and its room (${NONE})` },
      { target: '.cui-article-body > blockquote ul', props: ['list-style-type', 'padding-left'], golden: `a list in a quotation the copy's (${NONE})` },
      { target: '.cui-article-body > ol pre', props: ['background-color', 'padding-left', 'overflow-x'], golden: `a code block in a list item the copy's, scrolling in its column (${NONE})` },
    ],
  },
  {
    // a link and a button in a footnote keep their own look (rule 4)
    name: 'footnote-parts',
    html: ex('article').replace(/(<section class="cui-article-footnotes">\s*<ol>)/, '$1<li><p>See <a class="cui-link" href="#guide">the guide</a> <a class="cui-btn" href="#buy">Buy</a></p></li>'),
    widths: [1200],
    checks: [
      { target: '.cui-article-footnotes .cui-link', props: ['color', 'text-decoration-line'], golden: 'a link in a footnote the kit\'s link still, not the footnotes\' (rule 4)' },
      { target: '.cui-article-footnotes .cui-btn', props: ['color', 'background-color', 'text-decoration-line'], golden: 'a button in a footnote the kit\'s button still (rule 4)' },
    ],
  },
];
