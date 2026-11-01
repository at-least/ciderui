// The docs aside against Apple's developer documentation (DocC): a note
// (docc:note*), its label's and its text's room (docc-page:content.note.*),
// an important aside (docc-important:*), and the note's class set to each
// other kind (docc-asides:*), light and dark (their -dark). Golden, with
// their why: the letter-spacing.

import { examples, SIZE, PADDING } from '../../../test/harness.mjs';

const ex = examples(import.meta.url);
const A = '.cui-docs-aside';
const L = '.cui-docs-aside-label';
// the edge as DocC's asides were read (docc.mjs's ASIDE: its right side's
// color alone)
const EDGE = [...['top', 'bottom', 'left'].flatMap(s => [`border-${s}-width`, `border-${s}-style`, `border-${s}-color`]), 'border-right-color', 'border-top-left-radius', 'box-shadow', 'background-color'];
// each kind's aside in the kinds example, and where DocC's was read
const KINDS = [['important', 'docc-important'], ['warning', 'docc-asides'], ['tip', 'docc-asides'], ['deprecated', 'docc-asides'], ['experiment', 'docc-asides']];

export const cases = [
  {
    name: 'note',
    html: ex('note'),
    checks: [
      { target: A, props: [...PADDING, ...EDGE, ...SIZE, 'color'], apple: ['docc', 'note'] },
      { target: L, props: [...SIZE, 'color', 'margin-bottom'], apple: ['docc', 'note.label'] },
      { target: L, props: ['margin-top', 'margin-bottom'], apple: ['docc-page', 'content.note.label'] },
      { target: `${A} > p:not(${L})`, props: [...SIZE, 'color'], apple: ['docc', 'note.text'] },
      { target: `${A} > p:not(${L})`, props: ['margin-top', 'margin-bottom'], apple: ['docc-page', 'content.note.text'] },
      ...[A, L].map(target => ({ target, props: ['letter-spacing'], golden: 'the system font\'s fitted letter-spacing (tokens.css), not DocC\'s' })),
    ],
  },
  {
    name: 'kinds',
    html: ex('kinds'),
    widths: [1200, 390],
    checks: KINDS.flatMap(([kind, page]) => [
      { target: `${A}[data-variant="${kind}"]`, props: [...EDGE, 'color'], apple: [page, kind] },
      { target: `${A}[data-variant="${kind}"] ${L}`, props: [...SIZE, 'color'], apple: [page, `${kind}.label`] },
    ]),
  },
  {
    name: 'note dark',
    html: `<div class="cui-dark">${ex('note')}</div>`,
    widths: [1200],
    checks: [
      { target: A, props: [...EDGE, 'color'], apple: ['docc-dark', 'note'] },
      { target: L, props: ['color'], apple: ['docc-dark', 'note.label'] },
    ],
  },
  {
    name: 'kinds dark',
    html: `<div class="cui-dark">${ex('kinds')}</div>`,
    widths: [1200],
    checks: KINDS.flatMap(([kind, page]) => [
      { target: `${A}[data-variant="${kind}"]`, props: [...EDGE, 'color'], apple: [`${page}-dark`, kind] },
      { target: `${A}[data-variant="${kind}"] ${L}`, props: ['color'], apple: [`${page}-dark`, `${kind}.label`] },
    ]),
  },
  {
    // a second paragraph, which DocC's pages read show none of
    name: 'paragraphs',
    html: `<aside class="cui-docs-aside" aria-label="Note"><p class="cui-docs-aside-label">Note</p><p>The first paragraph.</p><p>The second paragraph.</p></aside>`,
    widths: [1200],
    checks: [
      { target: `${A} > p:last-child`, props: ['margin-top'], golden: 'a second paragraph 6.8px under the first, as the text under the label (docc-page:content.note.text): DocC\'s pages read show none' },
    ],
  },
];
