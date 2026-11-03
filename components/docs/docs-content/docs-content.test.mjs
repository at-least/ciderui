import { LINK, PADDING, SIZE, examples, DOCC_TRACKING } from '../../../test/harness.mjs';

const ex = examples(import.meta.url);
const C = '.cui-docs-content';
const LIST = ['margin-top', 'margin-left', 'padding-left', 'padding-inline-start', 'list-style-type', 'list-style-position'];
const CELL = [...PADDING, 'text-align', 'vertical-align', 'background-color',
  'border-top-width', 'border-top-style', 'border-top-color', 'border-bottom-width', 'border-bottom-style', 'border-bottom-color',
  'border-left-width', 'border-left-style', 'border-left-color', 'border-right-width', 'border-right-style', 'border-right-color'];
const STEPS = [1500, 1499, 1251, 1250, 1024, 1023, 768, 767, 736, 735];
const NONE = 'DocC\'s pages read show none';

export const cases = [
  {
    name: 'content',
    html: ex('content'),
    checks: [
      { target: `${C} > h2:first-child`, props: [...SIZE, 'color', 'margin-top', 'margin-bottom'], apple: ['docc', 'heading'] },
      { target: `${C} > h2 + p`, props: [...SIZE, 'color', 'margin-top'], apple: ['docc', 'text'] },
      { target: `${C} > p a`, props: [...SIZE, 'color', ...LINK], apple: ['docc', 'link'] },
      { target: `${C} > p a`, state: 'hover', props: ['color', ...LINK], apple: ['docc', 'link', 'hover'] },
      { target: `${C} > p code`, props: [...SIZE, 'color', 'letter-spacing', 'background-color', 'padding-left', 'padding-right', 'border-top-left-radius'], apple: ['docc', 'code'] },
      { target: `${C} > pre`, props: ['border-top-left-radius', 'background-color', 'margin-top'], apple: ['docc', 'listing'] },
      { target: `${C} > pre`, props: [...PADDING, 'white-space'], apple: ['docc', 'listing.pre'] },
      { target: `${C} > pre code`, props: [...SIZE, 'color', 'white-space', 'tab-size'], apple: ['docc', 'listing.code'] },
      { target: `${C} > h3`, props: [...SIZE, 'color', 'margin-top', 'margin-bottom'], apple: ['docc-page', 'content.h3'] },
      { target: `${C} > ul`, props: LIST, apple: ['docc-page', 'content.ul'] },
      { target: `${C} > ul > li`, props: [...SIZE, 'color', 'display'], apple: ['docc-page', 'content.li'] },
      { target: `${C} > ul > li + li`, props: ['margin-top'], apple: ['docc-page', 'content.li.next'] },
      { target: `${C} > ul > li > p`, props: [...SIZE, 'color', 'margin-top', 'margin-bottom'], apple: ['docc-page', 'content.li.p'] },
      { target: `${C} > p:has(> img)`, props: ['margin-top', 'text-align'], apple: ['docc-page', 'content.image'] },
      { target: `${C} img`, props: ['max-width', 'display'], apple: ['docc-page', 'content.image.img'] },
      { target: `${C} > h2 ~ h2`, props: [...SIZE, 'color', 'margin-top', 'margin-bottom'], apple: ['docc-technote', 'h2.next'] },
      { target: `${C} > ol`, props: LIST, apple: ['docc-technote', 'ol'] },
      { target: `${C} > ol > li`, props: [...SIZE, 'color', 'display'], apple: ['docc-technote', 'ol.li'] },
      { target: `${C} > ol > li + li`, props: ['margin-top'], apple: ['docc-technote', 'ol.li.next'] },
      { target: `${C} > p em`, props: [...SIZE, 'color', 'font-style'], apple: ['docc-technote', 'em'] },
      { target: `${C} > table`, props: [...SIZE, 'color', 'border-collapse'], apple: ['docc-technote', 'table'] },
      { target: `${C} > table`, props: ['margin-top'], apple: ['docc-technote', 'table.wrapper'] },
      { target: `${C} th`, props: [...SIZE, 'color', ...CELL], apple: ['docc-technote', 'table.head'] },
      { target: `${C} td`, props: [...SIZE, 'color', ...CELL], apple: ['docc-technote', 'table.cell'] },
      ...[`${C} > h2:first-child`, `${C} > h3`, `${C} > p`, `${C} > ul > li`]
        .map(target => ({ target, props: ['letter-spacing'], golden: DOCC_TRACKING })),
      { target: `${C} > p code`, props: ['font-family'], golden: 'the mono face (tokens.css, --cui-font-mono): DocC\'s SF Mono where it is the system\'s' },
      { target: `${C} > pre`, props: ['font-family', 'letter-spacing', 'overflow-x'], golden: 'the mono face at its own spacing, as the console\'s listing (DocC sets -0.405px on its SF Mono webfont); a line too long moves sideways inside the listing (DocC\'s scroller inside its box)' },
      { target: `${C} > table`, props: ['display', 'overflow-x', 'max-width'], golden: 'the table scrolls by itself where the column is narrower: DocC\'s wrapper does (docc-technote:table.wrapper), and markdown writes none' },
      { target: `${C} img`, props: ['box-width', 'box-x'], golden: 'the image at its own size, centered in the column (DocC\'s auto margins)' },
    ],
  },
  {
    name: 'steps',
    html: ex('content'),
    widths: STEPS,
    checks: [
      { target: `${C} > h2:first-child`, props: [...SIZE, 'margin-top'], apple: ['docc-page', 'content.h2'] },
      { target: `${C} > h3`, props: [...SIZE, 'margin-top'], apple: ['docc-page', 'content.h3'] },
      { target: `${C} > p`, props: [...SIZE, 'margin-top'], apple: ['docc-page', 'content.p'] },
      { target: `${C} > ul`, props: ['margin-left', 'margin-top'], apple: ['docc-page', 'content.ul'] },
      { target: `${C} > ul > li + li`, props: ['margin-top'], apple: ['docc-page', 'content.li.next'] },
      { target: `${C} > pre`, props: ['margin-top'], apple: ['docc-page', 'content.listing'] },
    ],
  },
  {
    name: 'dark',
    html: `<div class="cui-dark">${ex('content')}</div>`,
    checks: [
      { target: `${C} > h2 + p`, props: ['color'], apple: ['docc-dark', 'text'] },
      { target: `${C} > h2:first-child`, props: ['color'], apple: ['docc-dark', 'heading'] },
      { target: `${C} > h3`, props: ['color'], apple: ['docc-page-dark', 'content.h3'] },
      { target: `${C} > p a`, props: ['color', 'text-decoration-color'], apple: ['docc-dark', 'link'] },
      { target: `${C} > p code`, props: ['color', 'background-color'], apple: ['docc-dark', 'code'] },
      { target: `${C} > pre`, props: ['background-color'], apple: ['docc-dark', 'listing'] },
      { target: `${C} > pre code`, props: ['color'], apple: ['docc-dark', 'listing.code'] },
      { target: `${C} th`, props: ['color', 'border-top-color', 'border-left-color'], apple: ['docc-technote-dark', 'table.head'] },
      { target: `${C} td`, props: ['color', 'border-top-color', 'border-left-color'], apple: ['docc-technote-dark', 'table.cell'] },
    ],
  },
  {
    name: 'more',
    html: `<div class="cui-docs-content">
<p>The first paragraph, at the top.</p>
<aside><p>An aside.</p></aside>
<h4>A smaller heading</h4>
<p>Text with <strong>strong</strong> words.</p>
<blockquote><p>A quotation.</p></blockquote>
<hr>
<figure><p>An example’s frame.</p></figure>
</div>`,
    checks: [
      { target: `${C} > aside`, props: ['margin-top'], apple: ['docc-page', 'content.note'] },
      { target: `${C} > p:first-child`, props: ['margin-top'], golden: 'the first child at the content\'s top (DocC\'s first heading, docc:heading margin-top 0)' },
      { target: `${C} > h4`, props: [...SIZE, 'margin-top', 'letter-spacing'], golden: `${NONE}: the text's size at 600, a paragraph's room above` },
      { target: `${C} strong`, props: ['font-weight'], golden: `${NONE}: 600, DocC's headings' and labels' weight` },
      { target: `${C} > blockquote`, props: ['padding-left', 'color', 'margin-top'], golden: `${NONE}: indented as a list, in the secondary text` },
      { target: `${C} > hr`, props: ['border-top-width', 'border-top-style', 'border-top-color', 'margin-top'], golden: `${NONE}: a rule in the divider's color, a listing's room above` },
      { target: `${C} > figure`, props: ['margin-top'], golden: `${NONE}: an example's frame, an image's room above (docc-page:content.image)` },
    ],
  },
  {
    // as Zola writes a list in a list item and a code block in one
    name: 'nested',
    html: `<div class="cui-docs-content">
<ul><li>Outer<ul><li>Inner</li></ul></li></ul>
<ol><li><p>Add the packages:</p><pre><code>npm install --save-dev @tailwindcss/cli tailwindcss</code></pre></li></ol>
<ul><li><p>Then build:</p><div class="cui-docs-copy"><pre><code>npx @tailwindcss/cli -i app.css -o out.css</code></pre></div></li></ul>
</div>`,
    widths: [1600, 390],
    checks: [
      { target: `${C} > ul ul`, props: ['list-style-type', 'margin-left', 'margin-top'], golden: `a list in a list item the content's list (${NONE})` },
      { target: `${C} > ol pre`, props: ['background-color', 'padding-left', 'overflow-x', 'margin-top'], golden: `a listing in a list item the content's, a paragraph's room above (${NONE})` },
      { target: `${C} > ul .cui-docs-copy`, props: ['margin-top'], golden: `a listing with its copy button in a list item, a paragraph's room above, as a listing's (${NONE})` },
    ],
  },
  {
    name: 'long',
    html: `<div class="cui-docs-content">
<p>Import <code>@import "../target/ciderui/components/docs/docs-content/docs-content.css";</code> first.</p>
<p><a href="#">https://developer.apple.com/documentation/xcode/examining-the-fields-in-a-crash-report</a></p>
<table><thead><tr><th>file</th><th>its path</th></tr></thead><tbody><tr><td>content</td><td><code>components/docs/docs-content/docs-content.css</code></td></tr></tbody></table>
</div>`,
    widths: [1600, 390],
    checks: [
      { target: C, props: ['overflow-wrap'], golden: 'a word too long for its line breaks where it must: DocC writes <wbr> into a long name (docc:link\'s code), markdown does not' },
      ...[`${C} > p code`, `${C} > p a`, `${C} > table`].map(target => ({ target, props: ['box-width', 'box-x'], golden: 'within the column: a long word broken in it, a table scrolling in it' })),
    ],
  },
  {
    name: 'parts',
    html: `<div class="cui-docs-content"><p>Read <a class="cui-console-link" href="#guide">the guide</a>.</p><pre class="cui-console-listing"><code>Thread 0 Crashed:</code></pre></div>`,
    widths: [1200],
    checks: [
      { target: `${C} .cui-console-link`, props: ['color', 'text-decoration-line'], golden: 'a console link keeps its own look, the brand\'s blue not underlined, over the content\'s link (one class each: the content\'s rules match plain elements alone; rule 4)' },
      { target: `${C} .cui-console-listing`, props: ['background-color', 'font-size'], golden: 'a console listing keeps its own look, not the content\'s (rule 4); the content gives it its room alone' },
    ],
  },
];
