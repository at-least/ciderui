// The console's pager: the summary on the toolkit's body text style; the
// gaps App Store Connect's toolkit's control row's (controlrow.json, golden:
// tb reads tokens.json alone). Golden, with their why: the gaps, the
// summary's fitted label, the line's layout, the dark side.

import { examples } from '../../test/harness.mjs';

const ex = examples(import.meta.url);
const S = '.cui-console-pager-summary';

export const cases = [
  {
    name: 'pager',
    html: ex('pager'),
    widths: [1200, 390],
    checks: [
      { target: S, props: ['font-size', 'line-height', 'font-weight', 'letter-spacing'], tb: ['typography-textStyle-bodyRegular-fontSize', 'typography-textStyle-bodyRegular-lineHeight', 'typography-textStyle-bodyRegular-fontWeight', 'typography-textStyle-bodyRegular-letterSpacing'] },
      { target: S, props: ['color'], golden: 'the secondary label at 0.74: the toolkit\'s 0.60 is 3.4:1 on white (tokens.css, decision 21)' },
      { target: '.cui-console-pager-actions', props: ['column-gap'], golden: '8px: the toolkit\'s control row\'s md button (controlrow.json, ControlRow-anatomy-Button-variant-controlSize-md-marginRight)' },
      { target: '.cui-console-pager', props: ['column-gap', 'justify-content', 'flex-wrap'], golden: '24px at least between the summary and the buttons: a control row\'s md section (controlrow.json); the summary at the start, the buttons at the end, wrapping' },
      { target: '.cui-console-pager-actions .cui-console-btn[href]', props: ['box-height'], golden: 'the console\'s 30px button' },
    ],
  },
  {
    name: 'pager dark',
    html: `<div class="cui-dark">${ex('pager')}</div>`,
    widths: [1200],
    checks: [
      { target: S, props: ['color'], golden: 'the dark side derived (tokens.css, decision 18)' },
    ],
  },
];
