// The docs copy against Apple's developer documentation (DocC): the button
// at rest (docc:copy*), with the pointer on the listing and once it has
// copied (docc-copy:*), light and dark. Golden, with their why: hidden
// until the page's script shows it, shown with the keyboard's focus and on
// a screen with no hover, unchanged with the pointer on it, its fade alone
// transitioned, its width, its touch target.

import { examples, PADDING } from '../../../test/harness.mjs';

const ex = examples(import.meta.url);
// the button as the page's script leaves it: shown
const shown = ex('copy').replace(' hidden>', '>');
const copied = shown.replace('aria-label="Copy">', 'aria-label="Copy" data-state="copied">');
const B = '.cui-docs-copy-button';

export const cases = [
  {
    name: 'copy',
    html: shown,
    widths: [1200, 390],
    checks: [
      { target: B, props: [...PADDING, 'border-top-left-radius', 'background-color', 'height', 'position', 'top', 'right', 'opacity', 'cursor', 'transition-duration', 'transition-timing-function', 'box-height'], apple: ['docc', 'copy'] },
      { target: B, props: ['color'], as: { color: 'fill' }, apple: ['docc', 'copy.icon'] },
      { target: '.cui-docs-copy-idle', props: ['height'], apple: ['docc', 'copy.icon'] },
      { target: B, props: ['width', 'transition-property'], golden: 'a circle, DocC\'s 25.5px glyph 4px in (docc:copy width; an apple check reads no width); its fade alone transitioned (DocC\'s all)' },
      { target: '.cui-docs-copy', props: ['position'], golden: 'the button\'s containing block (rule 6\'s exception)' },
      { target: '.cui-docs-copy-done', props: ['visibility'], golden: 'the check hidden until it has copied' },
    ],
  },
  {
    // the pointer on the listing shows it
    name: 'copy hovered',
    html: shown,
    widths: [1200, 390],
    hover: '.cui-docs-copy pre',
    checks: [
      { target: B, props: ['opacity', 'background-color'], apple: ['docc-copy', 'listing.hovered'] },
      { target: B, props: ['color'], as: { color: 'fill' }, apple: ['docc-copy', 'listing.hovered.icon'] },
    ],
  },
  {
    // the pointer on the button: DocC's hides its glyph
    name: 'copy button hovered',
    html: shown,
    widths: [1200],
    hover: B,
    checks: [
      { target: B, props: ['opacity'], apple: ['docc-copy', 'button.hovered'] },
      { target: B, props: ['background-color', 'color'], golden: 'unchanged with the pointer on it: DocC\'s circle turns the text\'s color under a glyph of the same color (docc-copy:button.hovered rgb(29, 29, 31), its glyph\'s fill rgb(29, 29, 31)), which hides the glyph' },
    ],
  },
  {
    name: 'copy copied',
    html: copied,
    widths: [1200],
    checks: [
      { target: '.cui-docs-copy-done', props: ['color'], apple: ['docc-copy', 'copied'] },
      { target: B, props: ['opacity'], apple: ['docc-copy', 'listing.hovered'] },
      { target: '.cui-docs-copy-idle', props: ['visibility'], golden: 'its glyph hidden while the check shows, as DocC swaps them' },
      { target: '.cui-docs-copy-done', props: ['visibility'], golden: 'the check shown' },
    ],
  },
  {
    name: 'copy dark',
    html: `<div class="cui-dark">${shown}</div>`,
    widths: [1200],
    checks: [
      { target: B, props: ['background-color'], apple: ['docc-dark', 'copy'] },
      { target: B, props: ['color'], as: { color: 'fill' }, apple: ['docc-dark', 'copy.icon'] },
    ],
  },
  {
    name: 'copy copied dark',
    html: `<div class="cui-dark">${copied}</div>`,
    widths: [1200],
    checks: [
      { target: '.cui-docs-copy-done', props: ['color'], apple: ['docc-copy-dark', 'copied'] },
    ],
  },
  {
    name: 'copy hidden',
    html: ex('copy'),
    widths: [1200],
    checks: [
      { target: B, props: ['display'], golden: 'hidden until the page\'s script shows it, with no script copying nothing: Tailwind\'s preflight holds [hidden] to display none (!important), over the button\'s grid' },
    ],
  },
  {
    name: 'copy focus',
    html: shown,
    widths: [1200],
    checks: [
      { target: B, state: 'focus-visible', props: ['opacity', 'outline-style', 'outline-color'], golden: 'shown with the keyboard\'s focus and the kit\'s ring (rule 9): DocC\'s stays transparent focused (read by script, 2026-11-01)' },
    ],
  },
  {
    name: 'copy touch',
    html: shown,
    widths: [390],
    touch: true,
    checks: [
      { target: B, props: ['opacity'], golden: 'shown on a screen with no hover, where no pointer comes onto the listing (DocC\'s shows on a tap\'s hover alone)' },
      { target: B, pseudo: '::after', props: ['width', 'height'], golden: 'the 33.5px button\'s target 44px square on a touch screen, the HIG\'s 44pt' },
    ],
  },
];
