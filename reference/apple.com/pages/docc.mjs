import { TYPE, BOX, LINK, FOCUS, MOTION, PAGE } from './common.mjs';

export const ASIDE = [...BOX, ...TYPE, 'border-left-width', 'border-left-style', 'border-left-color', 'border-right-color'];
export const probes = [
  { name: 'page', select: 'body', props: PAGE },
  { name: 'eyebrow', select: '.eyebrow', props: TYPE },
  { name: 'title', select: 'h1.title', props: [...TYPE, 'margin-top', 'margin-bottom'] },
  { name: 'abstract', select: '.abstract', props: TYPE },
  { name: 'heading', select: 'h2:has(+ p)', props: [...TYPE, 'margin-top', 'margin-bottom'] },
  { name: 'text', select: 'h2 + p', props: [...TYPE, 'margin-top'] },
  { name: 'link', select: 'p a.inline-link', props: [...TYPE, ...LINK], states: ['hover'] },
  { name: 'code', select: 'p code.inline-code', props: [...TYPE, 'background-color', 'padding-left', 'padding-right', 'border-top-left-radius'] },
  { name: 'listing', select: '.code-listing', props: [...BOX, 'display', 'overflow-x'], box: true },
  { name: 'listing.scroller', select: '.code-listing .container-general', props: ['display', 'overflow-x'], box: true },
  { name: 'listing.pre', select: '.code-listing pre', props: ['padding-top', 'padding-right', 'padding-bottom', 'padding-left', 'margin-top', 'margin-bottom', 'white-space'], box: true },
  { name: 'listing.code', select: '.code-listing pre code', props: [...TYPE, 'white-space', 'tab-size'] },
  { name: 'listing.line', select: '.code-listing .code-line-container', props: ['display', 'min-height', 'padding-left', 'padding-right', 'background-color'], box: true },
  { name: 'copy', select: '.code-listing button.copy-button', props: [...BOX, 'width', 'height', 'position', 'top', 'right', 'color', 'opacity', 'cursor', ...FOCUS, ...MOTION], box: true, states: ['hover', 'focus-visible'] },
  { name: 'copy.icon', select: '.code-listing button.copy-button svg', props: ['width', 'height', 'color', 'fill'], box: true },
  { name: 'bar', select: 'nav.documentation-nav .nav__wrapper', props: ['height', 'background-color', 'border-bottom-width', 'border-bottom-style', 'border-bottom-color', 'backdrop-filter', '-webkit-backdrop-filter'], box: true },
  { name: 'bar.ground', select: 'nav.documentation-nav .nav__background', props: ['background-color', 'border-bottom-width', 'border-bottom-style', 'border-bottom-color', 'box-shadow', 'opacity', 'backdrop-filter', '-webkit-backdrop-filter'], box: true },
  { name: 'bar.title', select: 'nav.documentation-nav .nav-title', props: TYPE, box: true },
  { name: 'bar.toggle', select: 'nav.documentation-nav button.sidenav-toggle', widths: [390], props: ['width', 'height', 'color', 'background-color'], box: true },
  { name: 'navigator', select: '.navigator', props: ['display', 'width', 'border-right-width', 'border-right-style', 'border-right-color', 'background-color'], box: true },
  { name: 'navigator.item', select: '.navigator-card-item:not(.active)', props: [...TYPE, 'height', 'padding-left', 'padding-right', 'background-color', 'border-top-left-radius'], box: true, states: ['hover'] },
  { name: 'navigator.link', select: '.navigator-card-item:not(.active) a.leaf-link', props: [...TYPE, ...LINK], states: ['hover'] },
  { name: 'navigator.current', select: '.navigator-card-item.active', props: [...TYPE, 'height', 'padding-left', 'background-color', 'border-top-left-radius'], box: true },
  { name: 'navigator.current.link', select: '.navigator-card-item.active a.leaf-link', props: [...TYPE, ...LINK] },
  { name: 'note', select: 'aside.note', props: ASIDE, box: true },
  { name: 'note.label', select: 'aside.note p.label', props: [...TYPE, 'margin-bottom'] },
  { name: 'note.text', select: 'aside.note p:not(.label)', props: TYPE },
];
export default {
  url: 'https://developer.apple.com/documentation/xcode/examining-the-fields-in-a-crash-report',
  wait: 4000,
  probes,
};
