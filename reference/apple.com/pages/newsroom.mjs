// apple.com's Newsroom, a light page under a dark global nav: the bar's dark
// material (the home page's is light), its links and icons, its phone menu
// opened.
import { TYPE, MOTION, BLUR } from './common.mjs';
export default {
  url: 'https://www.apple.com/newsroom/',
  probes: [
    { name: 'nav', select: '#globalnav', props: ['height', 'position', 'background-color', ...BLUR], vars: ['--globalnav-background', '--globalnav-backdrop-filter'], box: true },
    { name: 'nav.content', select: '#globalnav .globalnav-content', props: ['background-color', ...BLUR] },
    { name: 'nav.link', select: '#globalnav .globalnav-item:not(.globalnav-item-apple) .globalnav-link', props: [...TYPE, ...MOTION], states: ['hover'] },
    { name: 'nav.search', select: '#globalnav .globalnav-link-search', props: ['color'], states: ['hover'] },
    { name: 'nav.bag', select: '#globalnav .globalnav-link-bag', props: ['color'], states: ['hover'] },
    { name: 'nav.menutrigger', select: '#globalnav .globalnav-menutrigger-button', props: ['color', 'opacity'], widths: [900, 390], states: ['hover'] },
    // the phone's menu, opened: last, as the page stays open after the click
    { name: 'menu.open', click: '#globalnav-menutrigger-button', select: '#globalnav', props: ['background-color', ...BLUR], widths: [390] },
    { name: 'menu.content', select: '#globalnav .globalnav-content', props: ['background-color', ...BLUR], widths: [390] },
    { name: 'menu.link', select: '#globalnav .globalnav-menu-list .globalnav-link', props: TYPE, widths: [390], states: ['hover'] },
    { name: 'menu.trigger', select: '#globalnav .globalnav-menutrigger-button', props: ['color', 'opacity'], widths: [390] },
  ],
};
