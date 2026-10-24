// The console's table against App Store Connect's: the analytics overview's
// plain table (analytics:table.*: a large row, its cells, its text and its
// figure; hover-row:table.row hovered) and the toolkit's table tokens
// (--tb-component-Table-*: the header, each row size, the mobile page
// scale). Golden, with their why: no header rule, the last cell flush,
// the subtitle, the selected row, the secondary label's alpha, the dark side.

import { examples } from '../../test/harness.mjs';

const ex = examples(import.meta.url);
const DESKTOP = [1600, 1200, 640];
const T = 'component-Table';
const FITTED = 'the secondary label fitted to 0.74 (tokens.css, decision 21): the toolkit\'s 0.60 is 3.4:1 on white';

export const cases = [
  {
    name: 'table',
    html: ex('table'),
    widths: DESKTOP,
    checks: [
      { target: 'thead th', props: ['box-height', 'padding-top', 'padding-bottom', 'padding-right', 'font-size', 'line-height', 'font-weight'], tb: [`${T}-anatomy-TH-height`, `${T}-anatomy-TH-paddingY`, `${T}-anatomy-TH-paddingY`, `${T}-anatomy-TH-paddingX`, `${T}-anatomy-TH-variant-pageScale-desktop-textStyle-fontSize`, `${T}-anatomy-TH-variant-pageScale-desktop-textStyle-lineHeight`, `${T}-anatomy-TH-variant-pageScale-desktop-textStyle-fontWeight`] },
      { target: 'thead th', props: ['color'], golden: FITTED },
      { target: 'thead th.cui-console-table-figure', props: ['font-weight', 'text-align'], golden: 'a figure\'s header at the header\'s weight, at its end' },
      { target: 'thead th', props: ['border-bottom-width'], golden: 'no rule under the header: the toolkit\'s tokens name none for its header row' },
      { target: 'tbody td:nth-child(2)', props: ['padding-top', 'padding-bottom', 'padding-left', 'padding-right'], tb: [`${T}-anatomy-TD-variant-size-md-paddingTop`, `${T}-anatomy-TD-variant-size-md-paddingBottom`, `${T}-anatomy-TD-paddingX`, `${T}-anatomy-TD-paddingX`] },
      { target: 'tbody td:nth-child(2)', props: ['border-bottom-width', 'border-bottom-style', 'border-bottom-color'], tb: [`${T}-anatomy-TR-borderBottom-width`, `${T}-anatomy-TR-borderBottom-style`, `${T}-anatomy-TR-borderBottom-color`] },
      { target: '.cui-console-table', props: ['color', 'font-size', 'line-height', 'font-weight'], tb: [`${T}-color`, `${T}-variant-pageScale-desktop-textStyle-fontSize`, `${T}-variant-pageScale-desktop-textStyle-lineHeight`, `${T}-variant-pageScale-desktop-textStyle-fontWeight`] },
      { target: 'tbody tr:first-child td:last-child', props: ['padding-right'], golden: 'the last cell flush at its end, as the analytics table\'s first is at its start' },
      { target: 'tbody tr:first-child', props: ['background-color'], state: 'hover', asc: ['hover-row', 'table.row'] },
      { target: 'tbody tr[aria-selected="true"]', props: ['background-color'], golden: 'the brand\'s fill at 0.08, the toolkit\'s selected row\'s (--tb-component-Table-anatomy-TR-…-selected-fill), square: its 12px corners set out 12px are no table row\'s' },
      { target: '.cui-console-table-subtitle', props: ['font-size', 'line-height', 'font-weight', 'color', 'display'], golden: 'a line under a cell\'s text, 12px on 16px, the secondary label: the toolkit\'s display list\'s subtitle at the table\'s small size' },
      { target: 'tbody tr:first-child', props: ['cursor'], golden: 'a row that is a link, all of it' },
      { target: 'tbody tr:first-child td:first-child', props: ['hit'], golden: 'the pointer finds the row\'s link anywhere in its first cell' },
    ],
  },
  {
    name: 'table large',
    html: ex('large'),
    widths: DESKTOP,
    checks: [
      { target: 'tr:first-child', props: ['box-height'], asc: ['analytics', 'table.row'] },
      { target: 'td:first-child', props: ['padding-top', 'padding-bottom', 'padding-left', 'padding-right', 'vertical-align', 'border-bottom-width', 'border-bottom-style', 'border-bottom-color'], asc: ['analytics', 'table.cell'] },
      { target: 'td:first-child', props: ['font-size', 'line-height', 'font-weight', 'letter-spacing', 'color', 'text-align'], asc: ['analytics', 'table.text'] },
      { target: '.cui-console-table-figure', props: ['font-size', 'line-height', 'font-weight', 'letter-spacing', 'color', 'text-align', 'font-variant-numeric'], asc: ['analytics', 'table.figure'] },
    ],
  },
  {
    name: 'table small',
    html: ex('table').replace('<table class="cui-console-table">', '<table class="cui-console-table" data-size="sm">'),
    widths: [1200],
    checks: [
      { target: 'tbody tr:last-child td:nth-child(2)', props: ['padding-top', 'padding-bottom'], tb: [`${T}-anatomy-TD-variant-size-sm-paddingTop`, `${T}-anatomy-TD-variant-size-sm-paddingBottom`] },
    ],
  },
  {
    name: 'table phone',
    html: ex('table'),
    widths: [639, 390],
    checks: [
      { target: 'thead th', props: ['font-size', 'line-height', 'font-weight'], tb: [`${T}-anatomy-TH-variant-pageScale-mobile-textStyle-fontSize`, `${T}-anatomy-TH-variant-pageScale-mobile-textStyle-lineHeight`, `${T}-anatomy-TH-variant-pageScale-mobile-textStyle-fontWeight`] },
      { target: '.cui-console-table', props: ['font-size', 'line-height', 'font-weight'], tb: [`${T}-variant-pageScale-mobile-textStyle-fontSize`, `${T}-variant-pageScale-mobile-textStyle-lineHeight`, `${T}-variant-pageScale-mobile-textStyle-fontWeight`] },
    ],
  },
  {
    name: 'table touch',
    html: ex('table'),
    touch: true,
    widths: [390],
    checks: [
      { target: 'tbody tr:first-child', props: ['background-color'], state: 'hover', rest: 'a tapped row keeps no fill (rule 5)' },
    ],
  },
  {
    name: 'table dark',
    html: `<div class="cui-dark">${ex('table')}</div>`,
    widths: [1200],
    checks: [
      { target: 'tbody td:nth-child(2)', props: ['border-bottom-color', 'color'], golden: 'the dark side derived (tokens.css, decision 18)' },
      { target: 'tbody tr[aria-selected="true"]', props: ['background-color'], golden: 'the dark side derived: the brand blue, no step of the palette, takes the blue ramp\'s step nearest its lightness (tokens.css, decision 18)' },
    ],
  },
];
