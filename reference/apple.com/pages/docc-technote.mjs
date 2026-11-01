// Apple's developer documentation (DocC), a TechNote, the one page of
// developer.apple.com's documentation read that shows a table (an article's
// tables are rare there; DocC's own formatting article redirects to
// swift.org, its theme not developer.apple.com's): the table's wrapper, the
// table, a header cell and a cell (and the paragraph DocC writes in each), a
// row; a numbered list and its item; a section's heading after the first
// (the crash report's article has none); emphasis. docc-technote-dark reads
// the same in the dark.
import { TYPE } from './common.mjs';

const CELL = ['padding-top', 'padding-right', 'padding-bottom', 'padding-left', 'text-align', 'vertical-align', 'background-color',
  'border-top-width', 'border-top-style', 'border-top-color', 'border-bottom-width', 'border-bottom-style', 'border-bottom-color',
  'border-left-width', 'border-left-style', 'border-left-color', 'border-right-width', 'border-right-style', 'border-right-color'];
const ROW = ['background-color', 'border-top-width', 'border-top-style', 'border-top-color', 'border-bottom-width', 'border-bottom-style', 'border-bottom-color'];
export const probes = [
  { name: 'table.wrapper', select: '.primary-content .content .table-wrapper', props: ['overflow-x', 'margin-top', 'margin-bottom', 'display', 'width'], box: true },
  { name: 'table', select: '.primary-content .content .table-wrapper table', props: [...TYPE, 'border-collapse', 'border-spacing', 'background-color', 'width', 'margin-top', 'margin-bottom'], box: true },
  { name: 'table.head', select: '.primary-content .content .table-wrapper thead th', props: [...TYPE, ...CELL], box: true },
  { name: 'table.head.p', select: '.primary-content .content .table-wrapper thead th > p', props: [...TYPE, 'margin-top', 'margin-bottom'], box: true },
  { name: 'table.head.row', select: '.primary-content .content .table-wrapper thead tr', props: ROW, box: true },
  { name: 'table.cell', select: '.primary-content .content .table-wrapper tbody td', props: [...TYPE, ...CELL], box: true },
  { name: 'table.cell.p', select: '.primary-content .content .table-wrapper tbody td > p', props: [...TYPE, 'margin-top', 'margin-bottom'], box: true },
  { name: 'table.row', select: '.primary-content .content .table-wrapper tbody tr', props: ROW, box: true },
  { name: 'table.row.even', select: '.primary-content .content .table-wrapper tbody tr:nth-child(2)', props: ROW, box: true },
  { name: 'ol', select: '.primary-content .content > ol', props: ['margin-top', 'margin-bottom', 'margin-left', 'padding-left', 'padding-inline-start', 'list-style-type', 'list-style-position'], box: true },
  { name: 'ol.li', select: '.primary-content .content > ol > li', props: [...TYPE, 'margin-top', 'margin-bottom', 'display'], box: true },
  { name: 'ol.li.next', select: '.primary-content .content > ol > li + li', props: ['margin-top'], box: true },
  { name: 'h2.next', select: '.primary-content .content > h2 ~ h2', props: [...TYPE, 'margin-top', 'margin-bottom'], box: true },
  { name: 'em', select: '.primary-content .content > p > em', props: [...TYPE, 'font-style'] },
];
export default {
  url: 'https://developer.apple.com/documentation/technotes/tn3151-choosing-the-right-networking-api',
  wait: 4000,
  probes,
};
