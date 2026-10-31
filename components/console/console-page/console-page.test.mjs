// The console's page against App Store Connect's newer look: its ground,
// its text and its type each the toolkit's token (reference/asc/results/
// tokens.json), at the desktop's page scale from 640px and the mobile's
// below; a section's heading and a metric's value as the analytics page drew
// them (analytics-page:section.heading, analytics:card.value). Golden, with
// their why: the font stacks, the dark side.

import { examples } from '../../../test/harness.mjs';

const ex = examples(import.meta.url);
const TYPE = ['font-size', 'line-height', 'font-weight', 'letter-spacing'];
const style = (s) => ['fontSize', 'lineHeight', 'fontWeight', 'letterSpacing'].map(p => `typography-textStyle-${s}-${p}`);
const DESKTOP = [1600, 1200, 900, 640];
const STACK = 'tokens.css\'s stack, SF on Apple\'s platforms and Inter elsewhere, where the toolkit names -system-ui first (decision 4)';

// each utility and the toolkit's text style it is
const STYLES = [
  ['h1-plus', 'h1PlusRegular'],
  ['h1-plus-emph', 'h1PlusEmph'],
  ['h1', 'h1Regular'],
  ['h1-emph', 'h1Emph'],
  ['h2', 'h2Emph'],
  ['h3', 'h3Regular'],
  ['h3-medium', 'h3Medium'],
  ['h3-emph', 'h3Emph'],
  ['h4', 'h4Regular'],
  ['h4-medium', 'h4Medium'],
  ['h4-emph', 'h4Emph'],
  ['h5', 'h5Regular'],
  ['h5-medium', 'h5Medium'],
  ['h5-emph', 'h5Emph'],
  ['h6', 'h6Regular'],
  ['h6-medium', 'h6Medium'],
  ['h6-emph', 'h6Emph'],
  ['body', 'bodyRegular'],
  ['body-medium', 'bodyMedium'],
  ['body-emph', 'bodyEmph'],
  ['small', 'bodySmallRegular'],
  ['small-medium', 'bodySmallMedium'],
  ['small-emph', 'bodySmallEmph'],
  ['xs', 'bodyXsRegular'],
  ['xs-medium', 'bodyXsMedium'],
  ['xs-emph', 'bodyXsEmph'],
];

export const cases = [
  {
    name: 'page',
    html: ex('page'),
    widths: DESKTOP,
    checks: [
      { target: '.cui-console-page', props: ['color', 'background-color'], tb: ['color-role-neutral-label-primary', 'color-role-neutral-background-primary'] },
      { target: '.cui-console-page', props: TYPE, tb: style('bodyRegular') },
      { target: '.cui-console-page', props: ['font-family'], golden: STACK },
      { target: '.cui-console-h4-emph', props: ['font-size', 'line-height', 'font-weight', 'letter-spacing', 'color'], asc: ['analytics-page', 'section.heading'] },
    ],
  },
  {
    name: 'page phone',
    html: ex('page'),
    widths: [639, 390],
    checks: [
      { target: '.cui-console-page', props: TYPE, tb: style('bodyMobileRegular') },
    ],
  },
  {
    name: 'page dark',
    html: `<div class="cui-dark">${ex('page')}</div>`,
    widths: [1200],
    checks: [
      { target: '.cui-console-page', props: ['color', 'background-color'], golden: 'the dark side: black and white each other\'s (tokens.css, decision 18)' },
    ],
  },
  {
    name: 'type',
    html: ex('type'),
    widths: [1200, 390],
    checks: [
      ...STYLES.map(([name, s]) => ({ target: `.cui-console-${name}`, props: TYPE, tb: style(s) })),
      { target: '.cui-console-h1-plus-emph', props: ['font-size', 'line-height', 'font-weight', 'letter-spacing'], asc: ['analytics', 'card.value'] },
      { target: '.cui-console-mono', props: ['font-size', 'line-height', 'font-weight', 'letter-spacing'], tb: style('monoBodyRegular') },
      { target: '.cui-console-mono', props: ['font-family'], golden: 'tokens.css\'s code stack, SF Mono on Apple\'s platforms and the platform\'s own elsewhere, where the toolkit names SF Mono alone' },
    ],
  },
];
