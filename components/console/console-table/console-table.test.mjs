import { examples } from '../../../test/harness.mjs';

const ex = examples(import.meta.url);
const DESKTOP = [1600, 1200, 640];
const T = 'component-Table';
const FITTED = 'the secondary label fitted to 0.76 (tokens.css, decision 21): the toolkit\'s 0.60 is 3.4:1 on white';

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
      { target: '.cui-console-table', props: ['color', 'font-size', 'line-height', 'font-weight'], tb: [`${T}-color`, `${T}-variant-pageScale-desktop-textStyle-fontSize`, `${T}-variant-pageScale-desktop-textStyle-lineHeight`, `${T}-variant-pageScale-desktop-textStyle-fontWeight`] },
      { target: 'tbody tr:first-child td:last-child', props: ['padding-right'], golden: 'the last cell flush at its end, as the analytics table\'s first is at its start' },
      { target: 'tbody tr:first-child', props: ['background-color'], state: 'hover', asc: ['hover-row', 'table.row'] },
      { target: '.cui-console-table-figure:not(th)', props: ['font-size', 'line-height', 'font-weight', 'letter-spacing', 'color', 'text-align', 'font-variant-numeric'], asc: ['analytics', 'table.figure'] },
      { target: 'tbody td:nth-child(2)', props: ['vertical-align', 'border-bottom-width', 'border-bottom-style', 'border-bottom-color'], asc: ['analytics', 'table.cell'] },
      { target: 'tbody td:nth-child(2)', props: ['font-size', 'line-height', 'font-weight', 'letter-spacing', 'color', 'text-align'], asc: ['analytics', 'table.text'] },
      { target: '.cui-console-table-subtitle', props: ['font-size', 'line-height', 'font-weight', 'color', 'display'], golden: 'a line under a cell\'s text, 12px on 16px, the secondary label: the toolkit\'s display list\'s subtitle at the table\'s small size' },
      { target: 'tbody tr:first-child', props: ['cursor'], golden: 'a row that is a link, all of it' },
      { target: '.cui-console-table', props: ['box-width'], golden: 'the full width of its place' },
      { target: 'tbody tr:first-child td:first-child', props: ['hit'], golden: 'the pointer finds the row\'s link anywhere in its first cell' },
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
    name: 'table scroll',
    html: ex('scroll'),
    widths: [1200, 390],
    checks: [
      { target: '.cui-console-table-scroll', props: ['overflow-x'], golden: 'a table wider than its place scrolls sideways in its box' },
      { target: 'tbody tr:first-child', props: ['background-color'], golden: 'a row whose checkbox is checked is selected: the brand\'s fill at 0.08, the toolkit\'s selected row\'s (--tb-component-Table-anatomy-TR-…-selected-fill), square: its 12px corners set out 12px are no table row\'s' },
      { target: 'tbody tr:first-child > :first-child', props: ['border-left-width', 'border-left-style', 'border-left-color', 'padding-left'], golden: 'a selected row its leading bar in every mode, 3px of the brand label\'s blue, its text 8px in as a cell\'s: a fill at 0.08 alone is too faint to find by, and forced colors drop it (the HIG\'s Color: not by color alone)' },
      { target: 'tbody tr:last-child > :first-child', props: ['border-left-width', 'padding-left'], golden: 'another row no bar, its first cell flush' },
      { target: 'tbody tr:first-child .cui-console-checkbox-input', props: ['hit'], golden: 'the pointer finds a linked row\'s checkbox, not the row\'s link: the link\'s ::after covers the row and comes later in it, and took the checkbox\'s click' },
      { target: 'tbody tr:first-child td:nth-child(3)', props: ['hit'], golden: 'elsewhere in the row the pointer still finds its link' },
    ],
  },
  {
    name: 'table checkbox in a cell\'s dialog',
    html: `<div class="cui-console-page"><table class="cui-console-table"><thead><tr><th scope="col">Project</th><th scope="col"><span class="sr-only">Actions</span></th></tr></thead><tbody>
      <tr id="routed"><td>Shop App</td><td><form><button type="button" class="cui-console-btn" data-variant="plain">Edit…</button><dialog class="cui-console-dialog"><label class="cui-console-checkbox"><input class="cui-console-checkbox-input" type="checkbox" checked><span class="cui-console-checkbox-box" aria-hidden="true"></span>ops</label></dialog></form></td></tr>
      <tr id="own"><td><span class="cui-console-checkbox"><input class="cui-console-checkbox-input" type="checkbox" checked aria-label="Select"><span class="cui-console-checkbox-box" aria-hidden="true"></span></span></td><td>Orders API</td></tr>
    </tbody></table></div>`,
    widths: [1200],
    checks: [
      { target: '#routed', props: ['background-color'], golden: 'a checkbox inside a cell\'s form or dialog is that control\'s, not the row\'s: no fill' },
      { target: '#own', props: ['background-color'], golden: 'a checkbox that is the cell\'s own control selects its row: the brand\'s fill at 0.08' },
    ],
  },
  {
    name: 'table row controls',
    html: '<table class="cui-console-table"><tbody><tr><td><a class="cui-console-table-link" href="#shop">Shop App</a></td><td><button type="button" class="cui-console-btn" data-variant="plain">Edit…</button></td><td><a href="#docs">Docs</a></td></tr></tbody></table>',
    widths: [1200],
    checks: [
      { target: 'button', props: ['hit'], golden: 'the pointer finds a linked row\'s button, not the row\'s link' },
      { target: 'a[href="#docs"]', props: ['hit'], golden: 'and another link in the row' },
    ],
  },
  {
    name: 'table checkbox touch',
    html: ex('scroll'),
    widths: [390],
    touch: true,
    checks: [
      { target: 'tbody tr:first-child .cui-console-checkbox', props: ['hit'], golden: 'on a touch screen too a tap on a linked row\'s checkbox finds the checkbox, not the row\'s link' },
      { target: 'tbody tr:first-child > td:first-child', props: ['padding-left', 'padding-right'], golden: 'a cell with a checkbox 14.5px in each way on a touch screen, room for its input\'s 44px: flush at the table\'s start the scroll box clipped it to 29px' },
    ],
  },
  {
    name: 'table stacked',
    html: ex('stack'),
    widths: [639, 390],
    checks: [
      { target: 'tbody tr:first-child', props: ['display', 'padding-top', 'padding-bottom', 'border-bottom-width'], golden: 'below 640px a row is a block, 12px above and below, its hairline under it' },
      { target: 'thead', props: ['display'], golden: 'the header hidden: each value says its own label' },
      { target: 'tbody tr:first-child td:nth-child(2)', props: ['display', 'justify-content', 'padding-top', 'border-bottom-width'], golden: 'a cell a line, its label at its start, its value at its end, 2px above and below' },
      { target: 'tbody tr:first-child td:nth-child(2)', pseudo: '::before', props: ['content', 'color', 'font-size', 'font-weight'], golden: 'its label, the cell\'s data-label, 13px at 500, the secondary label' },
      { target: 'tbody tr:first-child td:first-child', pseudo: '::before', props: ['content'], golden: 'a cell with no data-label, its value alone' },
      { target: 'tbody tr:first-child td[data-empty]', props: ['display'], golden: 'a cell that says nothing (data-empty) left out of the block' },
      { target: 'tbody tr:last-child td:last-child', props: ['display'], golden: 'the same column with a value, a line' },
    ],
  },
  {
    name: 'table scroll box',
    html: ex('scroll'),
    widths: [1200, 390],
    press: ['Tab'],
    checks: [
      { target: '.cui-console-table-scroll', props: ['focused', 'outline-style', 'overflow-x'], golden: 'the first Tab reaches the scroll box, ringed: the keyboard scrolls what overflows (axe\'s scrollable-region-focusable)' },
    ],
  },
  {
    name: 'table stacked desktop',
    html: ex('stack'),
    widths: [1200],
    checks: [
      { target: 'tbody tr:first-child td:nth-child(2)', props: ['display'], golden: 'from 640px a stacking table is a table' },
      { target: 'tbody tr:first-child td[data-empty]', props: ['display'], golden: 'from 640px an empty cell keeps its column' },
    ],
  },
  {
    name: 'table forced colors',
    html: ex('scroll'),
    widths: [1200],
    forcedColors: true,
    checks: [
      { target: 'tbody tr:first-child > :first-child', props: ['border-left-width', 'border-left-style', 'border-left-color'], golden: 'a selected row its leading edge, 3px of the theme\'s Highlight on its first cell: forced colors drop the fill, and the row read as the others' },
      { target: 'tbody tr:last-child > :first-child', props: ['border-left-width'], golden: 'another row no edge' },
    ],
  },
  {
    name: 'table dark',
    html: `<div class="cui-dark">${ex('scroll')}</div>`,
    widths: [1200],
    checks: [
      { target: 'tbody td:nth-child(2)', props: ['border-bottom-color', 'color'], golden: 'the dark side derived (tokens.css, decision 18)' },
      { target: 'tbody tr:first-child > :first-child', props: ['border-left-color'], golden: 'the bar in the dark, the brand label\'s blue, a step lighter than the accent (decision 21)' },
      { target: 'tbody tr:first-child', props: ['background-color'], golden: 'the dark side derived: the brand blue, no step of the palette, takes the blue ramp\'s step nearest its lightness (tokens.css, decision 18)' },
    ],
  },
];
