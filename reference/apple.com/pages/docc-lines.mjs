// Apple's developer documentation's (DocC) code listing with line numbers
// and highlighted lines, as a tutorial's code step shows them: the
// listing, its pre, a line, its number, a highlighted line and its
// number. docc-lines-dark reads the same in the dark.
import { TYPE, BOX } from './common.mjs';

const LINE = ['display', 'width', 'box-sizing', 'padding-left', 'padding-right', 'background-color', 'border-left-width', 'border-left-style', 'border-left-color'];
const NUMBER = [...TYPE, 'display', 'min-width', 'text-align', 'user-select', 'padding-top', 'padding-right', 'padding-bottom', 'padding-left', 'border-left-width', 'border-left-style', 'border-left-color'];
export const probes = [
  { name: 'listing', select: '.code-listing:has(.code-number)', props: [...BOX, 'display', 'overflow-x'], box: true },
  { name: 'listing.pre', select: '.code-listing:has(.code-number) pre', props: ['padding-top', 'padding-right', 'padding-bottom', 'padding-left', 'white-space'], box: true },
  { name: 'listing.code', select: '.code-listing:has(.code-number) pre code', props: TYPE },
  { name: 'line', select: '.code-listing:has(.code-number) .code-line-container:not(.highlighted)', props: LINE, box: true },
  { name: 'line.number', select: '.code-listing:has(.code-number) .code-line-container:not(.highlighted) .code-number', props: NUMBER, box: true },
  { name: 'highlighted', select: '.code-listing:has(.code-number) .code-line-container.highlighted', props: LINE, box: true },
  { name: 'highlighted.number', select: '.code-listing:has(.code-number) .code-line-container.highlighted .code-number', props: NUMBER, box: true },
];
export default {
  url: 'https://developer.apple.com/tutorials/swiftui/creating-and-combining-views',
  wait: 6000,
  probes,
};
