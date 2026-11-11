// @external: the App Store's web page for an app (Pages): its ratings' bars, each a
// track with a filled part as wide as its share (an inline width).
import { TYPE, BOX, FLEX } from './common.mjs';

const ROW = '.numbers__star-graph__row';
export default {
  url: 'https://apps.apple.com/us/app/pages/id361309726',
  wait: 3000,
  probes: [
    { name: 'graph', select: '.numbers__star-graph', props: [...FLEX, 'grid-template-columns', 'width'], box: true },
    { name: 'row', select: ROW, props: [...BOX, ...FLEX, 'grid-template-columns', 'height'], box: true },
    { name: 'track', select: `${ROW}__bar`, props: [...BOX, 'height', 'width', 'border-bottom-left-radius', 'overflow', 'margin-left'], box: true },
    { name: 'fill', select: `${ROW}__bar__foreground`, props: ['background-color', 'height', 'width', 'border-top-left-radius', 'border-top-right-radius'], box: true },
  ],
};
