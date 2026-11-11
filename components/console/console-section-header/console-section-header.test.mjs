import { examples } from '../../../test/harness.mjs';

const ex = examples(import.meta.url);
const S = '.cui-console-section-header';

export const cases = [
  {
    name: 'section header',
    html: ex('section-header'),
    widths: [1200, 390],
    checks: [
      { target: S, props: ['display', 'align-items', 'row-gap', 'column-gap'], asc: ['head', 'section.head'] },
      { target: S, props: ['margin-bottom'], golden: 'none of its own: the room under it is its container\'s (rule 6), App Store Connect\'s 20px a section\'s (console-section)' },
      { target: `${S} > :first-child`, props: ['flex-grow', 'min-width'], golden: 'the title takes the room, the controls at the row\'s end (App Store Connect\'s title 941px of 1105px)' },
      { target: `${S} h1`, props: ['font-size', 'line-height'], golden: 'the title\'s type its own class\'s (cui-console-h2), none from the header: App Store Connect\'s section heading is its h4, 18px on 22px' },
      { target: `${S} > .cui-console-tabs`, props: ['box-y'], golden: 'at 1200px on the title\'s row; at 390px under it (the row wraps; App Store Connect\'s does not)' },
      { target: `${S}-detail`, props: ['color', 'margin-top', 'max-width'], golden: 'a page\'s line under its title: the secondary label, 8px under the title, a readable measure (the shape crashcart and totality wrote with utilities)' },
    ],
  },
  {
    name: 'section header section',
    html: ex('section'),
    widths: [1200],
    checks: [
      { target: `${S} h2`, props: ['font-size', 'line-height', 'font-weight'], golden: 'a section\'s title its own class\'s (cui-console-h4-emph): App Store Connect\'s section heading, 18px on 22px' },
      { target: `${S}-detail`, props: ['color', 'margin-top'], golden: 'a section\'s line under its title: the secondary label, 4px under its smaller title' },
    ],
  },
  {
    name: 'section header back',
    html: ex('back'),
    widths: [1200, 390],
    checks: [
      { target: `${S}-back`, props: ['display', 'color', 'font-size', 'margin-bottom', 'box-width'], golden: 'the way back over the title, named by the page it goes to with a chevron before it (iOS\'s back button), a console link at the body\'s size, 8px over the title, as wide as its words' },
      { target: `${S}-back > svg`, props: ['box-width'], golden: 'the chevron at the text\'s size' },
    ],
  },
];
