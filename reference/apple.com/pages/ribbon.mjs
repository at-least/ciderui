// @external: apple.com's ribbon, the message above a page's content (the AirPods page's,
// which drops in blue and turns white): read once its motion has run, its animations
// as its stylesheets name them. The iPhone page's, read until 2026-11-13, turned gray;
// it is now a campaign's own part with no link (`ac-ribbon__upgrade-2026`), and the
// store's ribbon turns white too.
import { TYPE, BOX, LINK, ANIMATION } from './common.mjs';
// @external: its colors run from the start to the end of its keyframes, which name them
const COLORS = ['--ribbon-background-color-initial', '--ribbon-background-color', '--ribbon-text-color-initial', '--ribbon-text-color', '--ribbon-link-color-initial', '--ribbon-link-color'];
export default {
  url: 'https://www.apple.com/airpods/',
  wait: 1000,
  // @external: its colors turn 1.6s after it shows, in 0.8s, and a new width runs it again
  settle: 3500,
  probes: [
    { name: 'ribbon', select: '.ribbon', props: ['overflow', 'height'], box: true },
    { name: 'ribbon.drop', select: '.ribbon-drop-wrapper', props: ANIMATION, keyframes: true },
    { name: 'ribbon.ground', select: '.ribbon .ribbon-content-wrapper', props: [...TYPE, ...BOX, 'text-align', ...ANIMATION], vars: COLORS, box: true, keyframes: true },
    { name: 'ribbon.content', select: '.ribbon .ribbon-content', props: [...TYPE, 'width', 'max-width', 'margin-left', 'margin-right', ...ANIMATION], box: true, keyframes: true },
    { name: 'ribbon.text', select: '.ribbon .ribbon-content p', props: [...TYPE, 'margin-top', 'margin-bottom'], box: true },
    { name: 'ribbon.link', select: '.ribbon .ribbon-link', props: [...TYPE, ...LINK, ...ANIMATION], states: ['hover'], keyframes: true },
    // @invariant: the pointer on it, a hand's hover, against the forced one: last, as the
    // page keeps it hovered
    { name: 'ribbon.link-hovered', hover: '.ribbon .ribbon-link', select: '.ribbon .ribbon-link', props: ['color', ...LINK, ...ANIMATION], matches: [':hover'] },
  ],
};
