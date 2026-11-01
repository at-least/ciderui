// Apple's developer documentation (DocC), the crash report's article below
// 1024px with its navigator opened by the bar's button (`click`, once, at
// the first width): the sheet it opens over the window (its place, size,
// ground, layer and motion), the navigator's card on it, and the button
// that closes it (its size, place in the card, colors). And the bar's place
// as the page scrolls (sticky). docc-sheet-dark reads the same in the dark.
import { BOX } from './common.mjs';

export default {
  url: 'https://developer.apple.com/documentation/xcode/examining-the-fields-in-a-crash-report',
  wait: 4000,
  widths: [1023, 900, 767, 390],
  probes: [
    { name: 'bar.place', select: 'nav.documentation-nav', props: ['position', 'top', 'z-index'] },
    { name: 'sheet', click: 'nav.documentation-nav button.sidenav-toggle', select: '.adjustable-sidebar-width .aside', props: ['position', 'top', 'left', 'width', 'height', 'z-index', 'background-color', 'transform', 'transition-property', 'transition-duration', 'transition-timing-function'], box: true },
    { name: 'sheet.card', select: '.navigator-card', props: ['background-color', 'color'], box: true, position: true },
    { name: 'sheet.close', select: '.navigator button.close-card', props: [...BOX, 'width', 'height', 'color', 'display'], box: true, position: true },
    { name: 'sheet.close.icon', select: '.navigator button.close-card svg', props: ['width', 'height', 'color', 'fill'], box: true },
  ],
};
