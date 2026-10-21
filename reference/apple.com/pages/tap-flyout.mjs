// The home page on a touch screen 1024px wide (a tablet's: the bar's links in
// a row), Mac's link in the bar tapped (`tap`, the DevTools protocol's touch
// events): apple.com's script, having seen a touch on the bar, opens its
// flyout in place of following the link. Read only.
export default {
  url: 'https://www.apple.com/',
  wait: 2500,
  touch: true,
  widths: [1024],
  probes: [
    { name: 'link', tap: '#globalnav .globalnav-link-mac', select: '#globalnav .globalnav-link-mac', props: ['color'], matches: [':hover', ':focus', ':focus-visible'] },
    { name: 'flyout', select: '#globalnav-submenu-link-mac', props: ['height', 'visibility', 'background-color'], box: true },
    { name: 'trigger', select: '#globalnav-menubutton-link-mac', props: ['display'], vars: [], box: true },
    { name: 'curtain', select: '#globalnav-curtain', props: ['opacity', 'visibility'] },
    { name: 'page', select: 'link[rel="canonical"]', props: [] },
  ],
};
