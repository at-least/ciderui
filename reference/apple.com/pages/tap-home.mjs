// The home page on a touch screen, its menu's button tapped at 390px (`tap`,
// the DevTools protocol's touch events: a touch screen's tap, not a forced
// :hover): what the tap leaves hovered, against touch-home:nav.menutrigger
// (forced: 0.8 at rest, 1 hovered) and home:menu.trigger (the menu opened by
// a click, the pointer gone: 0.8). Read only.
import home from './home.mjs';

const TRIGGER = '#globalnav .globalnav-menutrigger-button';

export default {
  url: home.url,
  wait: 2500,
  touch: true,
  widths: [390],
  probes: [
    { name: 'menutrigger', tap: TRIGGER, select: TRIGGER, props: ['color', 'opacity'], matches: [':hover', ':active', ':focus', ':focus-visible'] },
    // the menu the tap opened: the page held
    { name: 'root', select: 'html', props: ['overflow-y'] },
  ],
};
