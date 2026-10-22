// Apple Support's article with a long table (Apple security releases): the
// page; its title, subheader, a paragraph, a link (hovered), a list and the
// small note; the table's wrapper, which moves it sideways where the window
// is narrower, the table, a header cell, a cell and a link in it (hovered).
// support-dark reads the same in the dark, which Apple Support follows.
import { TYPE, LINK, PAGE } from './common.mjs';

const CELL = ['padding-top', 'padding-right', 'padding-bottom', 'padding-left', 'text-align', 'vertical-align', 'background-color', 'border-top-width', 'border-top-style', 'border-top-color', 'border-bottom-width', 'border-bottom-style', 'border-bottom-color'];
export default {
  url: 'https://support.apple.com/en-us/100100',
  wait: 3000,
  probes: [
    { name: 'page', select: 'body', props: PAGE },
    { name: 'title', select: 'h1.gb-header', props: [...TYPE, 'margin-bottom'] },
    { name: 'subheader', select: 'p.gb-subheader', props: TYPE },
    { name: 'text', select: 'p.gb-paragraph', props: [...TYPE, 'margin-top'] },
    { name: 'link', select: 'a.gb-anchor', props: [...TYPE, ...LINK], states: ['hover'] },
    { name: 'list', select: 'ul.gb-list', props: ['padding-left', 'margin-left', 'list-style-type'] },
    { name: 'note', select: '.gb-note', props: TYPE },
    { name: 'table.wrapper', select: '.table-wrapper', props: ['overflow-x', 'margin-top', 'margin-bottom'], box: true },
    { name: 'table', select: 'table.gb-table', props: ['border-collapse', 'background-color'], box: true },
    { name: 'table.head', select: 'table.gb-table th', props: [...TYPE, ...CELL], box: true },
    { name: 'table.cell', select: 'table.gb-table td', props: [...TYPE, ...CELL], box: true },
    { name: 'table.link', select: 'table.gb-table td a.gb-anchor', props: [...TYPE, ...LINK], states: ['hover'] },
  ],
};
