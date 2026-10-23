// Certificates, Identifiers & Profiles' list of keys: the list's text; the
// side nav (an item and its word, the current one and its word); the list's
// header (its rule, the title, the add glyph, the filter's button and its
// chevron); the header row (a column's label, a sortable one) and a row (a
// cell, the glyph in one).
import { TYPE } from '../../apple.com/pages/common.mjs';

const PAD = ['padding-top', 'padding-right', 'padding-bottom', 'padding-left'];
const RULE = ['border-bottom-width', 'border-bottom-style', 'border-bottom-color'];
const ROW = '[role=row]:not(.th)';
export default {
  path: '/account/resources/authkeys/list',
  probes: [
    { name: 'page', select: '.main-wrapper', props: TYPE },
    { name: 'nav', select: 'nav.app-nav', props: ['width', 'margin-right'], box: true },
    { name: 'nav.item', select: 'nav.app-nav li:not(.active)', props: ['height', 'margin-bottom', 'background-color', 'border-top-left-radius'], box: true },
    { name: 'nav.word', select: 'nav.app-nav li:not(.active) > span', props: [...TYPE, ...PAD], box: true },
    { name: 'nav.current', select: 'nav.app-nav li.active', props: ['background-color', 'border-top-left-radius'], box: true },
    { name: 'nav.current.word', select: 'nav.app-nav li.active > span', props: [...TYPE, ...PAD] },
    { name: 'header', select: '.list-header', props: [...PAD, ...RULE, 'height', 'margin-bottom'], box: true },
    { name: 'title', select: '.list-header h2', props: [...TYPE, 'margin-top', 'margin-bottom'], box: true },
    { name: 'add', select: '.list-header h2 ~ span a svg', props: ['width', 'height', 'color', 'fill', 'margin-left'], box: true },
    { name: 'filter', select: '.list-header button[class*="Link-sc"]', props: [...TYPE, 'cursor'], box: true },
    { name: 'filter.chevron', select: '.list-header button[class*="Link-sc"] svg', props: ['width', 'height', 'color', 'fill', 'margin-left'], box: true },
    { name: 'column', select: '[role=columnheader]:not(.sortable)', props: [...TYPE, ...PAD, ...RULE, 'background-color', 'text-align'], box: true },
    { name: 'column.sortable', select: '[role=columnheader].sortable', props: [...TYPE, 'cursor'], box: true },
    { name: 'row', select: ROW, props: ['height', 'background-color', 'cursor'], box: true },
    { name: 'cell', select: `${ROW} > [role=cell]`, props: [...TYPE, ...PAD, ...RULE, 'vertical-align'], box: true },
    { name: 'cell.glyph', select: `${ROW} [role=cell] img`, props: ['width', 'height', 'vertical-align'], box: true },
  ],
};
