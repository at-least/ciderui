import { DOCC_STEPS, LINK, PADDING, SIZE, examples, DOCC_TRACKING } from '../../../test/harness.mjs';

const ex = examples(import.meta.url);
const H = '.cui-docs-hero';
const CRUMBS = '.cui-docs-hero-crumbs';

export const cases = [
  {
    name: 'hero',
    html: ex('hero'),
    widths: DOCC_STEPS,
    checks: [
      { target: H, props: ['padding-top', 'padding-bottom'], apple: ['docc-page', 'hero.content'] },
      { target: CRUMBS, props: [...SIZE, 'color', 'display', 'flex-direction', 'align-items', 'justify-content', 'margin-top', 'margin-bottom', ...PADDING], apple: ['docc-page', 'hierarchy'] },
      { target: `${CRUMBS} a`, props: [...SIZE, 'color', ...LINK], apple: ['docc-page', 'hierarchy.link'] },
      { target: `${CRUMBS} a`, state: 'hover', props: ['color', ...LINK], apple: ['docc-page', 'hierarchy.link', 'hover'] },
      { target: `${CRUMBS} [aria-current="page"]`, props: ['font-size', 'font-weight', 'color', 'text-decoration-line'], apple: ['docc-page', 'hierarchy.current'] },
      { target: `${CRUMBS} li + li`, pseudo: '::before', props: ['font-size', 'color', 'margin-left', 'margin-right'], apple: ['docc-page', 'hierarchy.separator'] },
      { target: '.cui-docs-hero-eyebrow', props: ['display', 'margin-top', 'margin-bottom'], apple: ['docc-page', 'eyebrow.room'] },
      { target: '.cui-docs-hero-title', props: [...SIZE, 'color', 'margin-top'], apple: ['docc-page', 'title.box'] },
      { target: '.cui-docs-hero-title', props: ['margin-bottom'], apple: ['docc-page', 'topictitle'] },
      { target: '.cui-docs-hero-abstract', props: [...SIZE, 'color', 'margin-top', 'margin-bottom'], apple: ['docc-page', 'abstract.box'] },
      ...[CRUMBS, '.cui-docs-hero-eyebrow', '.cui-docs-hero-title', '.cui-docs-hero-abstract']
        .map(target => ({ target, props: ['letter-spacing'], golden: DOCC_TRACKING })),
      { target: CRUMBS, props: ['flex-wrap', 'list-style-type'], golden: 'the crumbs wrap where the row has no room: DocC\'s script hides its middle crumbs there (docc-page:hierarchy.item display none), and there is no script; no marks (DocC\'s ul reads disc, its items flex boxes that draw none)' },
      { target: `${CRUMBS} li + li`, pseudo: '::before', props: ['content', 'display'], golden: 'a "/" read as nothing (alternative text ""), where DocC\'s is read aloud; a block in the crumb\'s row, DocC\'s where the crumb shows (docc-page:hierarchy.separator display block at 1600px)' },
    ],
  },
  {
    name: 'hero eyebrow',
    html: ex('hero'),
    checks: [
      { target: '.cui-docs-hero-eyebrow', props: [...SIZE, 'color'], apple: ['docc', 'eyebrow'] },
      { target: '.cui-docs-hero-title', props: [...SIZE, 'color'], apple: ['docc', 'title'] },
      { target: '.cui-docs-hero-abstract', props: [...SIZE, 'color'], apple: ['docc', 'abstract'] },
    ],
  },
  {
    name: 'hero dark',
    html: `<div class="cui-dark">${ex('hero')}</div>`,
    widths: [1600, 390],
    checks: [
      { target: `${CRUMBS} a`, props: ['color', 'text-decoration-color'], apple: ['docc-page-dark', 'hierarchy.link'] },
      { target: `${CRUMBS} [aria-current="page"]`, props: ['color'], apple: ['docc-page-dark', 'hierarchy.current'] },
      { target: `${CRUMBS} li + li`, pseudo: '::before', props: ['color'], apple: ['docc-page-dark', 'hierarchy.separator'] },
      { target: '.cui-docs-hero-eyebrow', props: ['color'], apple: ['docc-dark', 'eyebrow'] },
      { target: '.cui-docs-hero-title', props: ['color'], apple: ['docc-dark', 'title'] },
      { target: '.cui-docs-hero-abstract', props: ['color'], apple: ['docc-dark', 'abstract'] },
    ],
  },
  {
    // DocC's pages read show all four parts
    name: 'hero parts',
    html: `<header class="cui-docs-hero" id="title-alone"><h1 class="cui-docs-hero-title">Get Started</h1></header>
<header class="cui-docs-hero" id="no-crumbs"><h1 class="cui-docs-hero-title">Get Started</h1><p class="cui-docs-hero-abstract">Add ciderui to a site.</p></header>`,
    widths: [1600, 390],
    checks: [
      { target: '#title-alone .cui-docs-hero-title', props: ['margin-top', 'margin-bottom'], golden: 'a title alone: no room under it but the hero\'s own 40px' },
      { target: '#no-crumbs .cui-docs-hero-title', props: ['margin-top', 'margin-bottom', 'box-y'], golden: 'no crumbs and no eyebrow: the title at the hero\'s top, 40px down, its abstract 12px under it (none to 735px)' },
    ],
  },
];
