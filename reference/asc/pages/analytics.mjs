// An app's analytics overview, the newer look (its classes rt-…): the
// sidebar's menu (the current item, an item of an opened group, indented,
// an item's icon, the rule between groups), the date's control (the previous and next buttons beside
// the period), a metric's card (its title, its info button, its value, its
// change), a sparkline (its area, its line, its gradient's two stops), a
// bar chart (a grid line, the axes' labels, a bar and its gradient's stops)
// and the plain table beside it (a row, a cell, its text and a figure).
import { TYPE } from '../../apple.com/pages/common.mjs';

const ICON = ['width', 'height', 'fill', 'color'];
const ITEM = ['height', 'background-color', 'border-top-left-radius', 'padding-left', 'padding-right', 'column-gap', 'color'];
const LABEL = 'span[class*="rt-Text--trim--cap-alphabetic"]';
const STEP = ['height', 'background-color', 'color', 'border-top-left-radius', 'border-top-right-radius', 'border-bottom-left-radius', 'border-bottom-right-radius', 'padding-left', 'padding-right', 'margin-left', 'border-left-width', 'border-left-style', 'border-left-color'];
const STOP = ['stop-color', 'stop-opacity'];
export default {
  path: '/apps/<app>/analytics/overview',
  probes: [
    { name: 'menu.current', select: 'nav[class*="rt-DisplayList"] a[class*="rt-Item"]', props: ITEM, box: true },
    { name: 'menu.current.label', select: 'nav[class*="rt-DisplayList"] a[class*="rt-Item"] > div > div', props: TYPE, box: true },
    { name: 'menu.item', select: 'nav[class*="rt-DisplayList"] a[class*="rt-Item"]', nth: 1, props: ITEM, box: true },
    { name: 'menu.item.label', select: 'nav[class*="rt-DisplayList"] a[class*="rt-Item"] > div > div', nth: 1, props: TYPE, box: true },
    { name: 'menu.icon', select: 'nav[class*="rt-DisplayList"] a[class*="rt-Item"] svg', props: ICON, box: true },
    { name: 'menu.rule', select: 'nav[class*="rt-DisplayList"] hr', props: ['height', 'background-color', 'border-top-style', 'margin-top', 'margin-bottom', 'margin-left', 'margin-right'], box: true },
    { name: 'date.previous', select: `button:has(+ button > ${LABEL})`, props: STEP, box: true },
    { name: 'date.period', select: `button:has(> ${LABEL})`, props: STEP, box: true },
    { name: 'date.period.label', select: `button > ${LABEL}`, props: TYPE, box: true },
    { name: 'date.next', select: `button:has(> ${LABEL}) + button`, props: STEP, box: true },
    { name: 'date.chevron', select: `button:has(+ button > ${LABEL}) svg`, props: ICON, box: true },
    { name: 'card.title', select: 'p[class*="rt-Text--textStyle--$bodyEmph"]', props: TYPE, box: true },
    { name: 'card.info', select: 'button[class*="rt-InfoSystemButton"]', props: ['width', 'height', 'color', 'background-color', 'border-top-left-radius'], box: true },
    { name: 'card.info.icon', select: 'button[class*="rt-InfoSystemButton"] svg', props: ICON, box: true },
    { name: 'card.value', select: 'h2[class*="rt-Text--textStyle--$h1PlusEmph"]', props: [...TYPE, 'font-variant-numeric'], box: true },
    { name: 'card.change', select: 'span[class*="rt-Text--textStyle--$h4Medium"]', props: TYPE, box: true },
    { name: 'spark.area', select: 'svg g.group.path path', props: ['fill', 'fill-opacity', 'opacity', 'stroke'], box: true },
    { name: 'spark.line', select: 'svg g.group.path path', nth: 1, props: ['fill', 'stroke', 'stroke-width', 'stroke-linecap', 'stroke-linejoin', 'opacity'], box: true },
    { name: 'spark.top', select: 'linearGradient[id^="linear_"] stop', unseen: true, props: STOP },
    { name: 'spark.bottom', select: 'linearGradient[id^="linear_"] stop', nth: 1, unseen: true, props: STOP },
    { name: 'chart.grid', select: 'svg g > line', props: ['stroke', 'stroke-width', 'stroke-dasharray', 'opacity'], box: true },
    { name: 'chart.y', select: 'svg g.axis.y text', props: [...TYPE, 'fill', 'text-anchor'], box: true },
    { name: 'chart.x', select: 'svg g.axis.x text', props: [...TYPE, 'fill', 'text-anchor'], box: true },
    { name: 'chart.bar', select: 'svg path[fill*="lg_"]', props: ['fill', 'stroke', 'opacity'], box: true },
    { name: 'chart.bar.top', select: 'linearGradient[id^="lg_"] stop', unseen: true, props: STOP },
    { name: 'chart.bar.bottom', select: 'linearGradient[id^="lg_"] stop', nth: 1, unseen: true, props: STOP },
    { name: 'table.row', select: 'table[class*="rt-Table"] tr', props: ['height', 'background-color', 'border-bottom-width', 'border-bottom-style', 'border-bottom-color'], box: true },
    { name: 'table.cell', select: 'table[class*="rt-Table"] td', props: ['padding-top', 'padding-right', 'padding-bottom', 'padding-left', 'vertical-align', 'text-align', 'border-bottom-width', 'border-bottom-style', 'border-bottom-color'], box: true },
    { name: 'table.text', select: 'table[class*="rt-Table"] td p', props: [...TYPE, 'text-align'], box: true },
    { name: 'table.figure', select: 'table[class*="rt-Table"] td p', nth: 1, props: [...TYPE, 'text-align', 'font-variant-numeric'], box: true },
  ],
};
