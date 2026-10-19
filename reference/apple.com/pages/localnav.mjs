// apple.com's local nav, the one most product pages carry under the global
// nav (MacBook Air's, light): the bar, sticky, its material and its rule,
// the product's name, its links (the current page's marked), its Buy
// button, and to 734px its menu, opened by its chevron. Everything is read
// with the page scrolled (the bar stuck to the top), but `*.top`, at the top;
// the menu is the bar to 833px, from 834px (the global nav's step).
import { TYPE, BOX, MOTION, BLUR, FLEX } from './common.mjs';

const S = 1500;
export const probes = [
  { name: 'localnav', select: '#ac-localnav', scroll: S, props: ['position', 'top', 'z-index', 'min-width'], box: true },
  { name: 'localnav.top', select: '#ac-localnav .ac-ln-background', props: ['background-color', ...BLUR] },
  { name: 'localnav.background', select: '#ac-localnav .ac-ln-background', scroll: S, props: ['background-color', ...BLUR, ...MOTION] },
  { name: 'localnav.rule', select: '#ac-localnav .ac-ln-background', pseudo: '::after', scroll: S, props: ['background-color', 'height', 'position', 'bottom', 'z-index'] },
  { name: 'localnav.content', select: '#ac-localnav .ac-ln-content', scroll: S, props: ['max-width', 'padding-left', 'padding-right', 'margin-left', 'margin-right', 'height'], box: true },
  { name: 'localnav.title', select: '#ac-localnav .ac-ln-title a', scroll: S, props: [...TYPE, 'opacity', 'text-decoration-line', 'cursor', ...MOTION], box: true, position: true, states: ['hover'] },
  { name: 'localnav.item', select: '#ac-localnav .ac-ln-menu-item:nth-child(2)', scroll: S, props: ['margin-left', 'margin-right', 'display'] },
  { name: 'localnav.link', select: '#ac-localnav a.ac-ln-menu-link', scroll: S, props: [...TYPE, 'opacity', 'text-decoration-line', 'display', 'cursor', ...MOTION], box: true, position: true, states: ['hover'] },
  { name: 'localnav.current', select: '#ac-localnav .ac-ln-menu-link.current', scroll: S, props: [...TYPE, 'opacity', 'cursor', 'display'], box: true, position: true },
  { name: 'localnav.current-mark', select: '#ac-localnav .ac-ln-menu-link.current', pseudo: '::after', scroll: S, props: ['background-color', 'opacity', 'width', 'height', 'top', 'left', 'bottom', 'border-top-left-radius', 'position'] },
  { name: 'localnav.action', select: '#ac-localnav .ac-ln-action-button', scroll: S, props: ['margin-left', 'margin-top'] },
  { name: 'localnav.button', select: '#ac-localnav .ac-ln-button', scroll: S, props: [...TYPE, ...BOX, 'margin-top', 'cursor'], box: true, position: true, states: ['hover', 'active'] },
  { name: 'localnav.chevron', select: '#ac-localnav .ac-ln-menucta', scroll: S, props: ['width', 'height', 'margin-right', 'display', 'cursor'], box: true, position: true, widths: [833, 390] },
  { name: 'localnav.chevron-image', select: '#ac-localnav .ac-ln-menucta-chevron', scroll: S, props: ['color', 'opacity', ...MOTION, ...FLEX], box: true, states: ['hover'], widths: [833, 390] },
  { name: 'localnav.chevron-svg', select: '#ac-localnav .ac-ln-menucta-chevron svg', scroll: S, props: ['width', 'height', 'transform'], box: true, position: true, widths: [833, 390] },
  { name: 'localnav.tray', select: '#ac-localnav .ac-ln-menu-tray', scroll: S, props: ['visibility', 'max-height', 'overflow-y', ...MOTION], box: true, widths: [833, 390] },
  // its menu opened by the chevron: last, as the page stays so
  { name: 'tray.open', click: '#ac-localnav .ac-ln-menucta', animations: '#ac-localnav', select: '#ac-localnav .ac-ln-menu-tray', scroll: S, props: ['visibility', 'max-height', ...MOTION], box: true, position: true, widths: [390] },
  { name: 'tray.background', select: '#ac-localnav .ac-ln-background', scroll: S, props: ['background-color', ...BLUR, ...MOTION], box: true, widths: [390] },
  { name: 'tray.rule', select: '#ac-localnav .ac-ln-background', pseudo: '::after', scroll: S, props: ['background-color', 'height', 'top'], widths: [390] },
  { name: 'tray.items', select: '#ac-localnav .ac-ln-menu-items', scroll: S, props: [...BOX, 'padding-left', 'padding-right'], box: true, position: true, widths: [390] },
  { name: 'tray.link', select: '#ac-localnav a.ac-ln-menu-link', scroll: S, props: [...TYPE, 'opacity', 'display', 'align-items', 'transform', ...MOTION], box: true, position: true, widths: [390], states: ['hover'] },
  { name: 'tray.current', select: '#ac-localnav .ac-ln-menu-link.current', scroll: S, props: [...TYPE, 'opacity'], box: true, position: true, widths: [390] },
  { name: 'tray.current-mark', select: '#ac-localnav .ac-ln-menu-link.current', pseudo: '::after', scroll: S, props: ['background-color', 'opacity', 'width', 'height', 'top', 'left', 'border-top-left-radius'], widths: [390] },
  { name: 'tray.chevron-svg', select: '#ac-localnav .ac-ln-menucta-chevron svg', scroll: S, props: ['transform'], widths: [390] },
];

export default { url: 'https://www.apple.com/macbook-air/', widths: [1600, 1200, 900, 834, 833, 390], probes };
