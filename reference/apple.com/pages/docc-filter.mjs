// Apple's developer documentation (DocC), the crash report's article: its
// navigator's filter, read 2026-11-01 at 1200px, where the navigator is a
// column. At rest: the bar at the navigator's foot (sticky, its hairline,
// its room), the field (its edge, corners, ground, height), the filter's
// glyph, the input (its type, its placeholder's color). Focused (`click`):
// the field's edge and glow (DocC then also shows its tags above the field:
// not built, not read). Typed "crash" (`type`): a row whose name holds it
// (its words, the matched part marked: its weight, its ground). Typed what
// no name holds: the line saying so. The quick navigation's "/" button
// beside the field is not built and not read. docc-filter-dark reads the
// same in the dark.
import { TYPE } from './common.mjs';

const BAR = '.navigator .navigator-filter';
const FIELD = `${BAR} .filter__wrapper`;
const INPUT = `${BAR} .filter__input`;
export default {
  url: 'https://developer.apple.com/documentation/xcode/examining-the-fields-in-a-crash-report',
  wait: 4000,
  widths: [1200],
  probes: [
    { name: 'wrapper', select: '.navigator .filter-wrapper', props: ['position', 'bottom', 'background-color'], box: true },
    { name: 'bar', select: BAR, props: ['display', 'padding-top', 'padding-right', 'padding-bottom', 'padding-left', 'border-top-width', 'border-top-style', 'border-top-color'], box: true },
    { name: 'field', select: FIELD, props: ['border-top-width', 'border-top-style', 'border-top-color', 'border-top-left-radius', 'background-color', 'height', 'box-shadow'], box: true },
    { name: 'glyph', select: `${BAR} .filter__filter-button`, props: ['width', 'height', 'margin-left', 'margin-right', 'display'], box: true },
    { name: 'glyph.svg', select: `${BAR} .filter__filter-button svg`, props: ['width', 'height', 'fill', 'color'], box: true },
    { name: 'input', select: INPUT, props: [...TYPE, 'height', 'padding-left', 'padding-right', 'background-color', 'border-top-width', 'caret-color'], box: true },
    { name: 'input.placeholder', select: INPUT, pseudo: '::placeholder', props: ['color'] },
    { name: 'focused', click: INPUT, select: FIELD, props: ['border-top-color', 'box-shadow', 'border-top-width'] },
    { name: 'match', type: ['c', 'r', 'a', 's', 'h'], typed: 1500, select: '.navigator .highlight .match', props: [...TYPE, 'background-color', 'border-top-left-radius', 'padding-left', 'padding-right'] },
    { name: 'match.row', select: '.navigator .navigator-card-item:not(.is-group):has(.match) .highlight > span:not(.match)', props: TYPE },
    { name: 'none', type: ['z', 'q', 'x', 'j'], typed: 1500, select: '.navigator .no-items', props: [...TYPE, 'padding-top', 'padding-left', 'margin-top', 'text-align'], box: true },
  ],
};
