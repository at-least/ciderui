// The ribbon against apple.com's (the iPhone page's, ribbon:*, read once
// its motion has run; the bag's, bag:ribbon): its ground, padding, type and
// width, its message's height, its link; the announcing one's motion (its
// start, on the accent, and its end, on the ground, and its animations'
// timing), its link hovered. Golden, with their why: the letter-spacing, the
// content's width (apple.com's 83.33%), the text on the accent.

import { examples, PADDING } from '../../test/harness.mjs';

const ex = examples(import.meta.url);
const TYPE = ['font-size', 'line-height', 'font-weight'];
const ANIMATION = ['animation-duration', 'animation-delay', 'animation-timing-function', 'animation-fill-mode', 'animation-iteration-count'];
const LINK = '.cui-ribbon .cui-link';

export const cases = [
  {
    name: 'ribbon',
    html: ex('ribbon'),
    checks: [
      { target: '.cui-ribbon', props: ['overflow', 'box-height'], apple: ['ribbon', 'ribbon'] },
      { target: '.cui-ribbon-content', props: [...TYPE, 'color', 'background-color', 'padding-top', 'padding-bottom', 'text-align'], apple: ['ribbon', 'ribbon.ground'] },
      { target: '.cui-ribbon-content', props: ['background-color', 'padding-top', 'padding-bottom'], apple: ['bag', 'ribbon'] },
      { target: '.cui-ribbon-content', props: ['padding-left', 'padding-right'], golden: 'the message within 83.33% of the window, centered: apple.com\'s content\'s width (ribbon:ribbon.content)' },
      { target: '.cui-ribbon-content', props: ['letter-spacing'], golden: 'the system font\'s fitted letter-spacing (tokens.css), not apple.com\'s' },
      { target: '.cui-ribbon-content p', props: [...TYPE, 'margin-top', 'margin-bottom'], apple: ['ribbon', 'ribbon.text'] },
      { target: LINK, props: [...TYPE, 'color', 'text-decoration-line'], apple: ['ribbon', 'ribbon.link'] },
      { target: LINK, state: 'hover', props: ['color'], apple: ['ribbon', 'ribbon.link', 'hover'] },
      { target: `${LINK} .cui-link-text`, on: LINK, state: 'hover', props: ['text-decoration-line'], apple: ['ribbon', 'ribbon.link', 'hover'] },
    ],
  },
  {
    // as the page shows it: on the accent until 1.6s on
    name: 'announce-start',
    html: ex('announce'),
    widths: [1200],
    checks: [
      { target: '.cui-ribbon-content', props: ['background-color'], golden: 'apple.com\'s accent before its colors turn (ribbon:ribbon.ground\'s --ribbon-background-color-initial, #0071e3)' },
      { target: '.cui-ribbon-content', props: ['color'], golden: 'the accent\'s own white on it: apple.com\'s #f5f5f7 (--ribbon-text-color-initial)' },
      { target: LINK, props: ['color'], golden: 'the link as the text on the accent: apple.com\'s #f5f5f7 (--ribbon-link-color-initial)' },
      { target: '.cui-ribbon-content', props: ['animation-name'], golden: 'its drop and its colors\' turn, ciderui\'s names for apple.com\'s ribbon-drop and ribbon-animate-*' },
      { target: LINK, props: ANIMATION.filter(p => p !== 'animation-fill-mode'), apple: ['ribbon', 'ribbon.link'] },
      { target: LINK, props: ['animation-fill-mode'], golden: 'both: its start shown through the delay and its end its own color; apple.com\'s forwards from a start that is its own color' },
    ],
  },
  {
    name: 'announce-motion',
    html: ex('announce'),
    widths: [1200],
    checks: [
      { target: '.cui-ribbon-content', props: ANIMATION, golden: 'apple.com\'s drop (ribbon:ribbon.drop: 0.8s from 0s) and its ground\'s turn (ribbon:ribbon.ground: 0.8s from 1.6s), both ease-in-out, held' },
    ],
  },
  {
    // once it has run
    name: 'announce-end',
    html: ex('announce'),
    widths: [1200],
    wait: 2800,
    checks: [
      { target: '.cui-ribbon-content', props: ['color', 'background-color'], apple: ['ribbon', 'ribbon.ground'] },
      { target: '.cui-ribbon-content', props: ['transform'], golden: 'dropped in, in its place (apple.com\'s ribbon-drop ends at translateY(0))' },
      { target: '.cui-ribbon-content', props: ['color'], apple: ['ribbon', 'ribbon.content'] },
      { target: LINK, props: ['color'], apple: ['ribbon', 'ribbon.link'] },
    ],
  },
];
