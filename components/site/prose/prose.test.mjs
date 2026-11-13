import { examples, SIZE, WITH_1600 } from '../../../test/harness.mjs';

const ex = examples(import.meta.url);
const P = '.cui-prose';
const FLOW = ['margin-top', 'margin-bottom'];

export const cases = [
  {
    name: 'prose',
    html: ex('prose'),
    widths: WITH_1600,
    checks: [
      { target: `${P} > p + p`, props: [...SIZE, 'color', ...FLOW], apple: ['support-prose', 'p'] },
      { target: `${P} > h2`, props: [...SIZE, 'color', ...FLOW], apple: ['support-prose', 'h2'] },
      { target: `${P} > h2 + p`, props: FLOW, apple: ['support-prose', 'h2.next'] },
      { target: `${P} > h3`, props: [...SIZE, 'color', ...FLOW], apple: ['support-prose', 'h3'] },
      { target: `${P} > ul`, props: [...FLOW, 'padding-left', 'margin-left', 'list-style-type', 'list-style-position'], apple: ['support-prose', 'ul'] },
      { target: `${P} > ol`, props: [...FLOW, 'padding-left', 'margin-left', 'list-style-type', 'list-style-position'], apple: ['support-prose', 'ol'] },
      { target: `${P} > ul > li`, props: [...SIZE, ...FLOW, 'padding-left'], apple: ['support-prose', 'li'] },
      { target: `${P} > ul > li + li`, props: FLOW, apple: ['support-prose', 'li.next'] },
      { target: `${P} > ol + h3`, props: ['margin-top'], apple: ['support-prose', 'h3'] },
      { target: `${P} > ul + p`, props: FLOW, apple: ['support-prose', 'list.next'] },
      { target: `${P} b`, props: SIZE, apple: ['support-prose', 'b'] },
      { target: `${P} a`, props: [...SIZE, 'color', 'text-decoration-line'], apple: ['support-prose', 'link'] },
      { target: `${P} > p:first-child`, props: ['margin-top'], golden: 'the first block flush with the container: the space above the text is the page\'s' },
      { target: `${P} > table`, props: ['margin-top', 'overflow-x', 'border-collapse'], golden: 'a table 14px under the text before it, as Apple Support\'s wrapper (support:table.wrapper); it scrolls itself, markdown writing no wrapper' },
      { target: `${P} th`, props: [...SIZE, 'padding-top', 'padding-left', 'text-align', 'vertical-align', 'border-bottom-width', 'border-bottom-style'], apple: ['support', 'table.head'] },
      { target: `${P} td`, props: [...SIZE, 'padding-top', 'padding-left', 'text-align', 'vertical-align', 'border-bottom-width', 'border-bottom-style'], apple: ['support', 'table.cell'] },
      { target: `${P} td`, props: ['border-bottom-color'], golden: 'the kit\'s divider, one rule color: Apple Support\'s #ddd is a gray of its own' },
    ],
  },
  {
    name: 'prose code, quote and notes',
    html: ex('code'),
    widths: WITH_1600,
    checks: [
      { target: `${P} p code`, props: SIZE, apple: ['docc', 'code'] },
      { target: `${P} p code`, props: ['background-color', 'padding-left'], apple: ['docc', 'code'] },
      { target: `${P} > pre`, props: ['margin-top', 'border-top-left-radius', 'background-color'], apple: ['docc', 'listing'] },
      { target: `${P} > pre`, props: ['padding-top', 'padding-right', 'padding-bottom', 'padding-left'], apple: ['docc', 'listing.pre'] },
      { target: `${P} > pre code`, props: SIZE, apple: ['docc', 'listing.code'] },
      { target: `${P} > pre`, props: ['overflow-x', 'tab-size'], golden: 'a long line scrolls inside the box, tabs four wide: DocC\'s listing (docc:listing.scroller)' },
      { target: `${P} > blockquote`, props: ['margin-top', 'border-top-left-radius', 'background-color', 'padding-top', 'padding-left'], apple: ['docc', 'note'] },
      { target: `${P} > blockquote`, props: ['border-top-width'], golden: 'no edge: DocC\'s note edges its box in #696969, a gray of its own; the alternate ground alone sets it apart' },
      { target: `${P} > blockquote > p + p`, props: ['margin-top'], golden: 'a quote\'s paragraphs 14px apart, as the text\'s' },
      { target: `${P} > hr`, props: ['margin-top', 'border-top-width', 'border-top-style', 'border-top-color'], golden: 'a rule in the kit\'s divider, 27.2px under the text as a listing is: no Apple article measured draws one' },
      { target: `${P} > .footnotes li`, props: [...SIZE, 'color'], apple: ['newsroom-article', 'footnotes.li'] },
      { target: `${P} > .footnotes ol`, props: ['padding-left', 'margin-left', 'list-style-type'], apple: ['newsroom-article', 'footnotes.ol'] },
      { target: `${P} > p + .footnotes`, props: ['margin-top'], apple: ['newsroom-article', 'footnotes'] },
      { target: `${P} > .footnotes li + li`, props: ['margin-top'], apple: ['newsroom-article', 'footnotes.li'] },
    ],
  },
  {
    name: 'prose with a part',
    html: '<div class="cui-prose"><p>Ready?</p><p><a class="cui-btn" href="#">Buy</a></p></div>',
    widths: [1200],
    checks: [
      { target: '.cui-prose .cui-btn', props: ['color', 'text-decoration-line'], golden: 'a button in the text keeps its own white and no underline: the text\'s link rule leaves a part\'s link alone' },
    ],
  },
  {
    name: 'prose picture',
    html: '<div class="cui-prose" style="width: 300px"><p>Step one.</p><img src="data:image/svg+xml,%3Csvg xmlns=%22http://www.w3.org/2000/svg%22 width=%22640%22 height=%22320%22/%3E" alt=""><p>Step two.</p></div>',
    widths: [1200, 390],
    checks: [
      { target: '.cui-prose > img', props: ['margin-top', 'margin-bottom', 'display'], apple: ['support-prose', 'img'] },
      { target: '.cui-prose > img', props: ['box-width', 'box-height'], golden: 'a picture no wider than the text, its height kept: Apple Support\'s are 320px screenshots (min(100%, 320px)), a consumer\'s any size' },
    ],
  },
];
