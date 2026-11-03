import { LINK, PADDING, SIZE, examples, DOCC_TRACKING } from '../../../test/harness.mjs';

const ex = examples(import.meta.url);
const N = '.cui-docs-navigator';
const TOP = `${N} > .cui-docs-navigator-list > .cui-docs-navigator-item`;
const NESTED = `${N} .cui-docs-navigator-list .cui-docs-navigator-list > .cui-docs-navigator-item`;
const CURRENT = '.cui-docs-navigator-row:has(> [aria-current="page"])';
const WIDE = [1600, 1500, 1499, 1251, 1250, 1200, 1024];
const FILTER = ex('navigator').replace('<div class="cui-docs-navigator-filter" role="search" hidden>', '<div class="cui-docs-navigator-filter" role="search">');

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
      ...[N, '.cui-docs-navigator-title', '.cui-docs-navigator-group'].map(target => ({ target, props: ['letter-spacing'], golden: DOCC_TRACKING })),
      { target: '.cui-docs-navigator-chevron', props: ['width', 'box-x'], golden: 'the chevron 20px at the end of its row\'s spacer (DocC\'s tree-toggle, docc-page:navigator.toggle width 20px; an apple check reads no width)' },
      { target: '.cui-docs-navigator-icon', props: ['width'], golden: 'square, DocC\'s 17px (docc-page:navigator.icon)' },
      { target: '.cui-docs-navigator-group', props: ['padding-left'], golden: 'a group\'s title at its rows\' words\' place: the row\'s 20px of room and its level\'s spacer' },
      { target: '.cui-docs-navigator-group', props: ['color'], golden: 'the secondary text, 5.1:1 (dark 5.8:1): DocC\'s is the muted text, rgb(134, 134, 139), 3.5:1 on white (decision 6)' },
      { target: '.cui-docs-navigator-row', props: ['margin-left', 'margin-right'], golden: 'the row 20px from either side, DocC\'s item\'s room (docc-page:navigator.group padding-left 20px), its gray from there' },
      { target: `${N} [aria-current="page"]`, pseudo: '::after', props: ['position', 'top', 'left', 'height'], golden: 'the whole row its link\'s, the row\'s 35px from its start (DocC\'s row follows its link by its script)' },
      { target: '#docs-navigator-alerts ~ .cui-docs-navigator-list', props: ['display'], golden: 'a closed page\'s rows hidden, by its checkbox (DocC\'s by its script)' },
    ],
  },
  {
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
    name: 'navigator focus',
    html: ex('navigator'),
    widths: [1200],
    checks: [
      { target: '#docs-navigator-alerts + .cui-docs-navigator-row .cui-docs-navigator-chevron', on: '#docs-navigator-alerts', state: 'focus-visible', props: ['outline-style', 'outline-width', 'outline-color', 'outline-offset'], golden: 'the kit\'s ring (rule 9) on the chevron, its checkbox out of sight' },
      { target: '#docs-navigator-alerts', props: ['position', 'width', 'height', 'opacity', 'pointer-events'], golden: 'the checkbox out of sight and under no pointer, in the focus order (DocC\'s chevron is its script\'s button)' },
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
    name: 'navigator filter',
    html: FILTER,
    widths: [1200],
    checks: [
      { target: '.cui-docs-navigator-filter', props: ['position', 'bottom', 'background-color'], apple: ['docc-filter', 'wrapper'] },
      { target: '.cui-docs-navigator-filter', props: [...PADDING, 'border-top-width', 'border-top-style', 'border-top-color', 'display', 'box-height'], apple: ['docc-filter', 'bar'] },
      { target: '.cui-docs-navigator-field', props: ['border-top-width', 'border-top-style', 'border-top-color', 'border-top-left-radius', 'background-color', 'box-height'], apple: ['docc-filter', 'field'] },
      { target: '.cui-docs-navigator-filter-icon', props: ['margin-left', 'margin-right', 'display'], apple: ['docc-filter', 'glyph'] },
      { target: '.cui-docs-navigator-filter-icon', props: ['height'], apple: ['docc-filter', 'glyph.svg'] },
      { target: '.cui-docs-navigator-filter-icon', props: ['width'], golden: 'square, DocC\'s 21px glyph (docc-filter:glyph.svg; an apple check reads no width)' },
      { target: '.cui-docs-navigator-filter-icon', props: ['color'], as: { color: 'fill' }, apple: ['docc-filter', 'glyph.svg'] },
      { target: '.cui-docs-navigator-input', props: [...SIZE, 'color', 'height', 'padding-left', 'background-color', 'border-top-width', 'caret-color'], apple: ['docc-filter', 'input'] },
      { target: '.cui-docs-navigator-input', pseudo: '::placeholder', props: ['color'], apple: ['docc-filter', 'input.placeholder'] },
      { target: '.cui-docs-navigator-field', state: 'focus-within', props: ['border-top-color'], apple: ['docc-filter', 'focused'] },
      { target: '.cui-docs-navigator-field', state: 'focus-within', props: ['box-shadow'], golden: 'a text field\'s ring, its edge 2px of the accent (rule 9): DocC\'s 1px edge and a 4px glow, rgba(0, 125, 250, 0.6) (docc-filter:focused)' },
      { target: '.cui-docs-navigator-input', props: ['letter-spacing', 'padding-right'], golden: 'the system font\'s fitted letter-spacing; 10px of room before the field\'s end (DocC\'s holds its clear button there)' },
      { target: '.cui-docs-navigator-filter', props: ['box-y'], golden: 'at the navigator\'s foot, under its rows' },
    ],
  },
  {
    name: 'navigator filtered',
    html: FILTER.replace('>Reading a crash report</a>', '>Reading a <mark class="cui-docs-navigator-match">crash</mark> report</a>').replace('<p class="cui-docs-navigator-empty" hidden>', '<p class="cui-docs-navigator-empty">'),
    widths: [1200],
    checks: [
      { target: '.cui-docs-navigator-match', props: [...SIZE, 'color', 'background-color', 'border-top-left-radius', 'padding-left', 'padding-right'], apple: ['docc-filter', 'match'] },
      { target: '.cui-docs-navigator-empty', props: [...SIZE, 'padding-top', 'padding-left', 'margin-top', 'text-align'], apple: ['docc-filter', 'none'] },
      { target: '.cui-docs-navigator-empty', props: ['color'], golden: 'the secondary text, 5.1:1: DocC\'s is the muted text, rgb(134, 134, 139), 3.5:1 on white (decision 6)' },
      { target: '.cui-docs-navigator-empty', props: ['letter-spacing'], golden: DOCC_TRACKING },
      { target: '.cui-docs-navigator-status', props: ['position', 'width', 'height', 'overflow-x'], golden: 'read aloud, not shown: what the filter found (DocC\'s visually hidden live region)' },
    ],
  },
  {
    name: 'navigator filter dark',
    html: `<div class="cui-dark">${FILTER.replace('>Reading a crash report</a>', '>Reading a <mark class="cui-docs-navigator-match">crash</mark> report</a>')}</div>`,
    widths: [1200],
    checks: [
      { target: '.cui-docs-navigator-filter', props: ['background-color'], apple: ['docc-filter-dark', 'wrapper'] },
      { target: '.cui-docs-navigator-filter', props: ['border-top-color'], apple: ['docc-filter-dark', 'bar'] },
      { target: '.cui-docs-navigator-field', props: ['border-top-color', 'background-color'], apple: ['docc-filter-dark', 'field'] },
      { target: '.cui-docs-navigator-input', props: ['color', 'caret-color'], apple: ['docc-filter-dark', 'input'] },
      { target: '.cui-docs-navigator-input', pseudo: '::placeholder', props: ['color'], apple: ['docc-filter-dark', 'input.placeholder'] },
      { target: '.cui-docs-navigator-match', props: ['color', 'background-color'], apple: ['docc-filter-dark', 'match'] },
    ],
  },
  {
    name: 'navigator filter narrow',
    html: FILTER.replace('<nav class="cui-docs-navigator" aria-label="Navigator">', '<nav class="cui-docs-navigator" aria-label="Navigator" style="width:240px">'),
    widths: [390],
    checks: [
      { target: '.cui-docs-navigator-field', props: ['box-width'], golden: 'the navigator\'s 240px less the filter\'s 20px either side: the field shrinks below its input\'s own width (min-width 0), where a flex item would overflow the navigator' },
    ],
  },
  {
    name: 'navigator no script',
    html: ex('navigator'),
    widths: [1200],
    checks: [
      { target: '.cui-docs-navigator-filter', props: ['display'], golden: 'hidden until the page\'s script shows it and filters (Tailwind\'s preflight\'s [hidden])' },
      { target: '.cui-docs-navigator-empty', props: ['display'], golden: 'hidden until the filter matches nothing' },
    ],
  },
  {
    name: 'navigator long',
    html: `<nav class="cui-docs-navigator" aria-label="Navigator" style="width:320px"><a class="cui-docs-navigator-title" href="#t">Crashes</a><ul class="cui-docs-navigator-list"><li class="cui-docs-navigator-item"><div class="cui-docs-navigator-row"><a class="cui-docs-navigator-link" href="#long">Understanding the exception types in a crash report</a></div></li></ul></nav>`,
    widths: [1200],
    checks: [
      { target: '.cui-docs-navigator-link', props: ['white-space', 'text-overflow', 'overflow-x', 'box-height'], golden: 'one line, an ellipsis where the row has no room (DocC\'s rows are one 25px line each)' },
    ],
  },
  {
    name: 'navigator deep',
    html: `<nav class="cui-docs-navigator" aria-label="Navigator"><ul class="cui-docs-navigator-list"><li class="cui-docs-navigator-item"><div class="cui-docs-navigator-row"><a class="cui-docs-navigator-link" href="#a">Crash reports</a></div><ul class="cui-docs-navigator-list"><li class="cui-docs-navigator-item"><div class="cui-docs-navigator-row"><a class="cui-docs-navigator-link" href="#b">Reading a crash report</a></div><ul class="cui-docs-navigator-list"><li class="cui-docs-navigator-item"><div class="cui-docs-navigator-row"><a class="cui-docs-navigator-link" href="#c">Exception types</a></div></li></ul></li></ul></li></ul></nav>`,
    widths: [1200],
    checks: [
      { target: `${N} .cui-docs-navigator-list .cui-docs-navigator-list .cui-docs-navigator-list .cui-docs-navigator-row`, props: ['padding-left'], golden: 'a row two levels in after a 65px spacer, DocC\'s step of 15px a level (docc-page:navigator.spacer, .spacer.nested)' },
    ],
  },
];
