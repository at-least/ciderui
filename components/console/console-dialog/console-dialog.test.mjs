import { examples } from '../../../test/harness.mjs';

const ex = examples(import.meta.url);
const M = 'component-Modal';
// @external: shown by its open attribute, not showModal(): no ::backdrop
const open = () => ex('dialog').replace(' aria-labelledby="ignore-title">', ' aria-labelledby="ignore-title" open>');

export const cases = [
  {
    name: 'dialog',
    html: open(),
    widths: [1200],
    checks: [
      { target: '.cui-console-dialog-actions', props: ['display'], golden: 'its actions in a row' },
      { target: '.cui-console-dialog', props: ['margin-left'], golden: 'centered: the kit\'s reset takes the browser\'s margin away' },
      { target: '.cui-console-dialog', props: ['box-width', 'border-top-left-radius', 'background-color', 'border-top-color', 'box-shadow', 'color'], tb: [`${M}-variant-variant-modal-minWidth`, `${M}-variant-variant-modal-borderRadius`, `${M}-backgroundColor`, `${M}-variant-variant-modal-stroke`, `${M}-variant-variant-modal-boxShadow`, `${M}-color`] },
      { target: '.cui-console-dialog', props: ['max-width'], tb: [`${M}-variant-variant-modal-maxWidth`] },
      { target: '.cui-console-dialog', props: ['padding-top', 'padding-left', 'border-top-width'], golden: 'ciderui\'s room inside, 24px (the toolkit\'s modal names none); a 1px edge' },
      { target: '.cui-console-dialog', props: ['font-size', 'line-height'], golden: 'the console\'s body, 13px on 18px' },
      { target: '.cui-console-dialog-title', props: ['font-size', 'line-height', 'font-weight', 'margin-bottom'], golden: 'the console\'s h4 at 600, 12px over what follows' },
      { target: '.cui-console-dialog-actions', props: ['justify-content', 'column-gap', 'margin-top'], golden: 'its buttons at its end, 9px apart (the banner\'s), 24px under what it asks' },
    ],
  },
  {
    name: 'dialog modal',
    html: ex('dialog'),
    click: '[command="show-modal"]',
    widths: [1200],
    checks: [
      { target: '.cui-console-dialog', pseudo: '::backdrop', props: ['background-color'], golden: 'ciderui\'s scrim behind it (--cui-c-scrim): the toolkit names no backdrop' },
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
