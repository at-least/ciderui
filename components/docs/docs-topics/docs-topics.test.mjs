import { DOCC_STEPS, LINK, SIZE, examples, DOCC_TRACKING } from '../../../test/harness.mjs';

const ex = examples(import.meta.url);
const S = '.cui-docs-topics';

export const cases = [
  {
    name: 'topics',
    html: ex('topics'),
    widths: DOCC_STEPS,
    checks: [
      { target: S, props: ['padding-bottom'], apple: ['docc-page', 'seealso.container'] },
      { target: '.cui-docs-topics-title', props: [...SIZE, 'color', 'padding-top', 'margin-top', 'margin-bottom'], apple: ['docc-page', 'seealso.title'] },
      { target: '.cui-docs-topics-group', props: ['padding-top'], apple: ['docc-page', 'seealso.section'] },
      { target: '.cui-docs-topics-heading', props: [...SIZE, 'color', 'padding-bottom', 'margin-top', 'margin-bottom'], apple: ['docc-page', 'seealso.section.title'] },
      { target: '.cui-docs-topics-item', props: ['margin-top'], apple: ['docc-page', 'seealso.block'] },
      { target: '.cui-docs-topics-item + .cui-docs-topics-item', props: ['margin-top'], apple: ['docc-page', 'seealso.block.next'] },
      { target: '.cui-docs-topics-link', props: [...SIZE, 'color', ...LINK, 'display', 'padding-top', 'padding-bottom'], apple: ['docc-page', 'seealso.link'] },
      { target: '.cui-docs-topics-link', state: 'hover', props: ['color', ...LINK], apple: ['docc-page', 'seealso.link', 'hover'] },
      { target: '.cui-docs-topics-icon', props: ['height'], apple: ['docc-page', 'seealso.icon.svg'] },
      { target: '.cui-docs-topics-icon', props: ['margin-right'], apple: ['docc-page', 'seealso.icon'] },
      { target: '.cui-docs-topics-abstract', props: [...SIZE, 'color', 'margin-left'], apple: ['docc-page', 'seealso.abstract'] },
      ...['.cui-docs-topics-title', '.cui-docs-topics-heading', '.cui-docs-topics-link'].map(target => ({ target, props: ['letter-spacing'], golden: DOCC_TRACKING })),
      { target: '.cui-docs-topics-icon', props: ['width', 'padding-left', 'padding-right', 'box-width', 'box-y'], golden: 'the 15px glyph in a 22px column, 5px down its 25px line (DocC\'s 22px box centers it: docc-page:seealso.icon)' },
      { target: '.cui-docs-topics-icon', props: ['color'], golden: 'the muted text, DocC\'s glyph\'s fill rgb(134, 134, 139) (docc-page:seealso.icon.svg; its color, the link\'s, paints nothing): a mark, 3.5:1 on white' },
    ],
  },
  {
    name: 'topics dark',
    html: `<div class="cui-dark">${ex('topics')}</div>`,
    widths: [1200, 390],
    checks: [
      { target: '.cui-docs-topics-title', props: ['color'], apple: ['docc-page-dark', 'seealso.title'] },
      { target: '.cui-docs-topics-link', props: ['color'], apple: ['docc-page-dark', 'seealso.link'] },
      { target: '.cui-docs-topics-abstract', props: ['color'], apple: ['docc-page-dark', 'seealso.abstract'] },
      { target: '.cui-docs-topics-icon', props: ['color'], golden: 'the dark muted text, DocC\'s dark fill rgb(110, 110, 115) (docc-page-dark:seealso.icon.svg)' },
    ],
  },
  {
    name: 'topics touch',
    html: ex('topics'),
    widths: [390],
    touch: true,
    checks: [
      { target: '.cui-docs-topics-link', state: 'hover', props: ['text-decoration-line'], rest: 'a tap leaves no hover behind (rule 5)' },
    ],
  },
];
