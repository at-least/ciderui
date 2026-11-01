// The docs navigator against Apple's developer documentation (DocC), where
// its navigator shows (from 1024px): its top, its title, a group's title, a
// row, the spacer of a row at the top and one in, the chevron closed and
// open, a row's glyph, the current row (docc-page:navigator.*), a row's
// words and the current one's (docc:navigator.*); light and dark. Golden,
// with their why: the letter-spacing, the widths an apple check does not
// read, a group's title fitted, the whole row its link's, a closed page's
// rows hidden, the chevron's focus ring and touch target, a long name.

import { examples, SIZE, PADDING } from '../../../test/harness.mjs';

const ex = examples(import.meta.url);
const N = '.cui-docs-navigator';
const TOP = `${N} > .cui-docs-navigator-list > .cui-docs-navigator-item`;
const NESTED = `${N} .cui-docs-navigator-list .cui-docs-navigator-list > .cui-docs-navigator-item`;
const CURRENT = '.cui-docs-navigator-row:has(> [aria-current="page"])';
const WIDE = [1600, 1500, 1499, 1251, 1250, 1200, 1024];
const LINK = ['text-decoration-line', 'text-decoration-color'];

export const cases = [
  {
    name: 'navigator',
    html: ex('navigator'),
    widths: WIDE,
    checks: [
      { target: N, props: ['padding-top'], apple: ['docc-page', 'navigator.inner'] },
      { target: '.cui-docs-navigator-title', props: [...PADDING, 'margin-left', 'margin-right', 'border-top-left-radius', 'display', 'box-height'], apple: ['docc-page', 'navigator.technology'] },
      { target: '.cui-docs-navigator-title', props: [...SIZE, 'color'], apple: ['docc-page', 'navigator.technology.title'] },
      { target: '.cui-docs-navigator-group', props: ['padding-right', 'box-height', 'font-size', 'line-height'], apple: ['docc-page', 'navigator.group'] },
      { target: '.cui-docs-navigator-group', props: [...SIZE, 'font-weight'], apple: ['docc-page', 'navigator.group.title'] },
      { target: `${TOP} > .cui-docs-navigator-row`, props: ['height', 'padding-top', 'padding-bottom', 'padding-right', 'display', 'align-items', 'box-height'], apple: ['docc-page', 'navigator.row'] },
      { target: `${TOP} > .cui-docs-navigator-row`, props: ['padding-left'], as: { 'padding-left': 'width' }, apple: ['docc-page', 'navigator.spacer'] },
      { target: `${NESTED} > .cui-docs-navigator-row`, props: ['padding-left'], as: { 'padding-left': 'width' }, apple: ['docc-page', 'navigator.spacer.nested'] },
      { target: '.cui-docs-navigator-chevron', props: ['height', 'padding-right', 'display', 'align-items', 'justify-content', 'color'], apple: ['docc-page', 'navigator.toggle'] },
      { target: '#docs-navigator-alerts + .cui-docs-navigator-row svg', props: ['height', 'color', 'transform'], apple: ['docc-page', 'navigator.toggle.icon'] },
      { target: '#docs-navigator-crash-reports + .cui-docs-navigator-row .cui-docs-navigator-chevron svg', props: ['height', 'color', 'transform'], apple: ['docc-page', 'navigator.toggle.open'] },
      { target: '.cui-docs-navigator-icon', props: ['height', 'margin-right', 'display'], apple: ['docc-page', 'navigator.icon'] },
      { target: '.cui-docs-navigator-icon', props: ['height', 'color'], apple: ['docc-page', 'navigator.icon.svg'] },
      { target: CURRENT, props: ['height', 'padding-right', 'background-color', 'border-top-left-radius', 'box-height'], apple: ['docc-page', 'navigator.current.row'] },
      // golden
      ...[N, '.cui-docs-navigator-title', '.cui-docs-navigator-group'].map(target => ({ target, props: ['letter-spacing'], golden: 'the system font\'s fitted letter-spacing (tokens.css), not DocC\'s' })),
      { target: '.cui-docs-navigator-chevron', props: ['width', 'box-x'], golden: 'the chevron 20px at the end of its row\'s spacer (DocC\'s tree-toggle, docc-page:navigator.toggle width 20px; an apple check reads no width)' },
      { target: '.cui-docs-navigator-icon', props: ['width'], golden: 'square, DocC\'s 17px (docc-page:navigator.icon)' },
      { target: '.cui-docs-navigator-group', props: ['color'], golden: 'the secondary text, 5.1:1 (dark 5.8:1): DocC\'s is the muted text, rgb(134, 134, 139), 3.5:1 on white (decision 6)' },
      { target: '.cui-docs-navigator-row', props: ['margin-left', 'margin-right'], golden: 'the row 20px from either side, DocC\'s item\'s room (docc-page:navigator.group padding-left 20px), its gray from there' },
      { target: `${N} [aria-current="page"]`, pseudo: '::after', props: ['position', 'top', 'left', 'height'], golden: 'the whole row its link\'s, the row\'s 35px from its start (DocC\'s row follows its link by its script)' },
      { target: '#docs-navigator-alerts ~ .cui-docs-navigator-list', props: ['display'], golden: 'a closed page\'s rows hidden, by its checkbox (DocC\'s by its script)' },
    ],
  },
  {
    // a row's words, which docc read at 1600 and 1200px
    name: 'navigator words',
    html: ex('navigator'),
    widths: [1600, 1200],
    checks: [
      { target: `${N} a.cui-docs-navigator-link:not([aria-current])`, props: [...SIZE, 'color', ...LINK], apple: ['docc', 'navigator.link'] },
      { target: `${N} [aria-current="page"]`, props: [...SIZE, 'color', ...LINK], apple: ['docc', 'navigator.current.link'] },
    ],
  },
  {
    name: 'navigator dark',
    html: `<div class="cui-dark">${ex('navigator')}</div>`,
    widths: [1200],
    checks: [
      { target: '.cui-docs-navigator-title', props: ['color'], apple: ['docc-page-dark', 'navigator.technology.title'] },
      { target: CURRENT, props: ['background-color'], apple: ['docc-page-dark', 'navigator.current.row'] },
      { target: '.cui-docs-navigator-icon', props: ['color'], apple: ['docc-page-dark', 'navigator.icon.svg'] },
      { target: '.cui-docs-navigator-chevron', props: ['color'], apple: ['docc-page-dark', 'navigator.toggle'] },
      { target: `${N} a.cui-docs-navigator-link:not([aria-current])`, props: ['color'], apple: ['docc-dark', 'navigator.link'] },
      { target: '.cui-docs-navigator-group', props: ['color'], golden: 'the dark secondary text, fitted as in light (DocC\'s dark rgb(110, 110, 115), 4.1:1 on black)' },
    ],
  },
  {
    // the chevron's focus ring, its checkbox having the keyboard's focus
    name: 'navigator focus',
    html: ex('navigator'),
    widths: [1200],
    checks: [
      { target: '#docs-navigator-alerts + .cui-docs-navigator-row .cui-docs-navigator-chevron', on: '#docs-navigator-alerts', state: 'focus-visible', props: ['outline-style', 'outline-width', 'outline-color', 'outline-offset'], golden: 'the kit\'s ring (rule 9) on the chevron, its checkbox out of sight' },
    ],
  },
  {
    name: 'navigator touch',
    html: ex('navigator'),
    widths: [1024],
    touch: true,
    checks: [
      { target: '.cui-docs-navigator-chevron', pseudo: '::after', props: ['width', 'height'], golden: 'the 20px chevron\'s target 44px square on a touch screen, the HIG\'s 44pt' },
    ],
  },
  {
    // a page's name longer than its row
    name: 'navigator long',
    html: `<nav class="cui-docs-navigator" aria-label="Navigator" style="width:320px"><a class="cui-docs-navigator-title" href="#t">Crashes</a><ul class="cui-docs-navigator-list"><li class="cui-docs-navigator-item"><div class="cui-docs-navigator-row"><a class="cui-docs-navigator-link" href="#long">Understanding the exception types in a crash report</a></div></li></ul></nav>`,
    widths: [1200],
    checks: [
      { target: '.cui-docs-navigator-link', props: ['white-space', 'text-overflow', 'overflow-x', 'box-height'], golden: 'one line, an ellipsis where the row has no room (DocC\'s rows are one 25px line each)' },
    ],
  },
];
