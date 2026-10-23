// System Status's other kinds, which its stylesheet draws and the page shows
// only while a service has one: the legend's mark with its class set to
// light-issue, light-outage and light-maintenance in turn (its available
// one is system-status's), and, at 390px, the toggle's glyph opened
// (`active`).
import { MARK } from './system-status.mjs';

const kind = (k) => ({ name: k, widths: [1600], set: [['.section-legend li > div', 'class', `light-${k}`]], select: '.section-legend li > div', props: MARK, image: ['background-image'], box: true });
export default {
  url: 'https://developer.apple.com/system-status/',
  wait: 3000,
  widths: [1600, 390],
  probes: [
    ...['issue', 'outage', 'maintenance'].map(kind),
    { name: 'toggle.active', set: [['.section-lights .light-toggle-icon', 'classList', 'active']], select: '.section-lights .light-toggle-icon', props: MARK, image: ['background-image'], box: true, widths: [390] },
  ],
};
