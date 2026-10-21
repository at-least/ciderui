// The iPhone page on a touch screen, its product gallery's next paddle tapped
// at 390px (`tap`, the DevTools protocol's touch events: a touch screen's
// tap, not a forced :hover): what the tap leaves hovered, against
// touch-iphone:paddle (forced: rgba(210, 210, 215, 0.64) at rest,
// rgba(223, 223, 227, 0.698) hovered). Read only.
import iphone from './iphone.mjs';

// the one touch-iphone:paddle reads, the first enabled
const PADDLE = 'main .section-select .paddlenav-arrow-next';

export default {
  url: iphone.url,
  wait: 2500,
  touch: true,
  widths: [390],
  probes: [
    // the gallery's current item before the tap (its markup names it)
    { name: 'gallery.current-before', select: 'main .section-select .gallery-item.current', props: ['display'] },
    { name: 'paddle', tap: PADDLE, select: PADDLE, props: ['background-color', 'color', 'border-top-color', 'opacity'], matches: [':hover', ':active', ':focus', ':focus-visible'] },
    // the gallery the tap moved: its current item the next
    { name: 'gallery.current', select: 'main .section-select .gallery-item.current', props: ['display'] },
  ],
};
