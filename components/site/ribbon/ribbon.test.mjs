import { examples, PADDING, SIZE, TRACKING } from '../../../test/harness.mjs';

const ex = examples(import.meta.url);
const LINK = '.cui-ribbon .cui-link';

export const cases = [
  {
    name: 'ribbon',
    html: ex('ribbon'),
    checks: [
      { target: '.cui-ribbon', props: ['overflow', 'box-height'], apple: ['ribbon', 'ribbon'] },
      { target: '.cui-ribbon-content', props: [...SIZE, 'color', 'background-color', 'padding-top', 'padding-bottom', 'text-align'], apple: ['ribbon', 'ribbon.ground'] },
      { target: '.cui-ribbon-content', props: ['background-color', 'padding-top', 'padding-bottom'], apple: ['bag', 'ribbon'] },
      { target: '.cui-ribbon-content', props: ['color'], apple: ['ribbon', 'ribbon.content'] },
      { target: '.cui-ribbon-content', props: ['padding-left', 'padding-right'], golden: 'the message within 83.33% of the window, centered: apple.com\'s content\'s width (ribbon:ribbon.content)' },
      { target: '.cui-ribbon-content', props: ['letter-spacing'], golden: TRACKING },
      { target: '.cui-ribbon-content p', props: [...SIZE, 'margin-top', 'margin-bottom'], apple: ['ribbon', 'ribbon.text'] },
      { target: LINK, props: [...SIZE, 'color', 'text-decoration-line'], apple: ['ribbon', 'ribbon.link'] },
      { target: LINK, state: 'hover', props: ['color'], apple: ['ribbon', 'ribbon.link', 'hover'] },
      { target: `${LINK} .cui-link-text`, on: LINK, state: 'hover', props: ['text-decoration-line'], apple: ['ribbon', 'ribbon.link', 'hover'] },
      { target: '.cui-ribbon-content', props: ['animation-name', 'transform'], golden: 'still as the page shows it: apple.com\'s iPhone page drops it in on the accent and turns it to its ground 1.6s on (ribbon:ribbon.drop, ribbon.ground), an entrance that only draws the eye (DESIGN.md decision 25)' },
    ],
  },
];
