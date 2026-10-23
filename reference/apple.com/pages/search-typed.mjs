// The global nav's search typed into, read 2026-10-21 with the owner's leave
// to type on apple.com: the home page's search opened as search.mjs opens
// it, then one key at a time into its field (`type`), nothing sent: "i"
// (one character, the quick links), "p" (its service asked: one section,
// Suggested Searches), "h" (two sections), "o"; Backspace back to "iph",
// "ip", "i" and the field empty; "x" and "q" (two characters no name of its
// starts with); and the clear button clicked. At each step, every animation
// and transition its script starts in the nav (`animations`, each first seen
// in ms from the key: one seen within 15ms is still held from a step before,
// a filled animation staying among the document's), the search's classes
// (its shift-<sections before>-<after> and with-results), the flyout's
// height, the results now (each section by its analytics region, and the
// number of its rows), and a row's rest once they have run, with its
// animation's keyframes. At 1600px: the rules of this motion have no width
// of their own (globalheader.css), its slide's distance aside.
// Read only.
import search from './search.mjs';

const FLY = '#globalnav-submenu-search';
const STEPS = [
  ['i', { type: ['i'] }],
  ['ip', { type: ['p'] }],
  ['iph', { type: ['h'] }],
  ['ipho', { type: ['o'] }],
  ['iph.back', { type: ['Backspace'] }],
  ['ip.back', { type: ['Backspace'] }],
  ['i.back', { type: ['Backspace'] }],
  ['empty', { type: ['Backspace'] }],
  ['x', { type: ['x'] }],
  ['xq', { type: ['q'] }],
  ['cleared', { click: `${FLY} .globalnav-searchfield-reset`, pointer: [200, 430] }],
];

export default {
  ...search,
  widths: [1600],
  probes: [
    ...search.probes.filter(p => p.name === 'nav'),
    ...STEPS.flatMap(([name, act]) => [
      { name, ...act, animations: '#globalnav', select: '#globalnav .globalnav-search', props: ['display'] },
      { name: `${name}.flyout`, select: FLY, props: ['height'], vars: ['--r-globalnav-flyout-height', '--r-globalnav-suggested-links-height'], box: true },
      { name: `${name}.results`, select: `${FLY} .globalnav-searchresults`, props: ['height'], sections: ['.globalnav-searchresults-container', 'data-analytics-region', '.globalnav-searchresults-list-item'], box: true },
      // a row at rest once the motion has run (none to read where nothing
      // may be suggested)
      ...(name === 'xq' ? [] : [{ name: `${name}.row`, select: `${FLY} .globalnav-searchresults-current .globalnav-searchresults-list-item`, props: ['opacity', 'transform', 'animation-name'], keyframes: true }]),
    ]),
  ],
};
