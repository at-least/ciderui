// The console's dialog against App Store Connect's toolkit's modal
// (--tb-component-Modal-*: its modal variant): its width, corners, ground,
// edge, shadow and label. Golden, with their why: the room inside, the
// title, the actions, the window's margin, the dark side.

import { examples } from '../../test/harness.mjs';

const ex = examples(import.meta.url);
const M = 'component-Modal';
// the example's dialog, shown: the page opens it with its button
const open = () => ex('dialog').replace(' aria-labelledby="ignore-title">', ' aria-labelledby="ignore-title" open>');

export const cases = [
  {
    name: 'dialog',
    html: open(),
    widths: [1200],
    checks: [
      { target: '.cui-console-dialog', props: ['box-width', 'border-top-left-radius', 'background-color', 'border-top-color', 'box-shadow', 'color'], tb: [`${M}-variant-variant-modal-minWidth`, `${M}-variant-variant-modal-borderRadius`, `${M}-backgroundColor`, `${M}-variant-variant-modal-stroke`, `${M}-variant-variant-modal-boxShadow`, `${M}-color`] },
      { target: '.cui-console-dialog', props: ['max-width'], tb: [`${M}-variant-variant-modal-maxWidth`] },
      { target: '.cui-console-dialog', props: ['padding-top', 'padding-left', 'border-top-width'], golden: 'ciderui\'s room inside, 24px (the toolkit\'s modal names none); a 1px edge' },
      { target: '.cui-console-dialog-title', props: ['font-size', 'line-height', 'font-weight', 'margin-bottom'], golden: 'the console\'s h4 at 600, 12px over what follows' },
      { target: '.cui-console-dialog-actions', props: ['justify-content', 'column-gap', 'margin-top'], golden: 'its buttons at its end, 9px apart (the banner\'s), 24px under what it asks' },
    ],
  },
  {
    name: 'dialog phone',
    html: open(),
    widths: [390],
    checks: [
      { target: '.cui-console-dialog', props: ['box-width'], golden: 'never wider than the window less 22px, 11px at each side, as the sidebar\'s sheet' },
    ],
  },
  {
    name: 'dialog dark',
    html: `<div class="cui-dark">${open()}</div>`,
    widths: [1200],
    checks: [
      { target: '.cui-console-dialog', props: ['background-color', 'border-top-color', 'color'], golden: 'the dark side derived (tokens.css, decision 18)' },
    ],
  },
];
