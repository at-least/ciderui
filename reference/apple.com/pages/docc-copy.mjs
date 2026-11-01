// Apple's developer documentation (DocC), the crash report's article: its
// code listing's copy button as the pointer and a click show it, read
// 2026-11-01. At rest it is transparent (docc:copy); with the pointer on
// the listing (`hover`) it shows, its glyph (the CopyIcon); with the
// pointer on the button its circle turns the text's color under a glyph of
// the same color, which hides the glyph (seen in a picture of it, not
// kept); clicked (`click`, last, at 1200px alone: back to the CopyIcon
// within 3.3s), its glyph the CheckmarkIcon in the link's color, its class
// success (the clipboard granted: denied, DocC draws a red failure). docc-copy-dark reads the same in the dark.
const BUTTON = '.code-listing button.copy-button';
export default {
  url: 'https://developer.apple.com/documentation/xcode/examining-the-fields-in-a-crash-report',
  wait: 4000,
  widths: [1200, 390],
  // its click writes the clipboard; denied, DocC shows a red failure
  options: { permissions: ['clipboard-read', 'clipboard-write'] },
  probes: [
    { name: 'listing.hovered', scroll: 900, hover: '.code-listing', select: BUTTON, props: ['opacity', 'background-color', 'color', 'transition-property', 'transition-duration', 'transition-timing-function'] },
    { name: 'listing.hovered.icon', scroll: 900, hover: '.code-listing', select: `${BUTTON} svg`, props: ['width', 'height', 'fill', 'color'], box: true },
    { name: 'button.hovered', scroll: 900, hover: BUTTON, select: BUTTON, props: ['opacity', 'background-color', 'color'] },
    { name: 'button.hovered.icon', scroll: 900, hover: BUTTON, select: `${BUTTON} svg`, props: ['fill', 'color'] },
    { name: 'copied', widths: [1200], scroll: 900, click: BUTTON, select: `${BUTTON} svg`, props: ['width', 'height', 'fill', 'color'], box: true },
  ],
};
