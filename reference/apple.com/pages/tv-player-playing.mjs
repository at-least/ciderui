// tv.apple.com's video player playing (tv-player's page and opening): its
// controls shown by a click on the video, the centre pause in the play's
// place; then, the pointer still, the controls hidden by its script
// (`video-player-scrim--hidden`, 4s on), and what that leaves. At 1600px
// alone: a resize, and the time it takes, would move both.
import { MOTION } from './common.mjs';
import { URL, open } from './tv-player.mjs';

const SHOWN = ['opacity', 'visibility', 'pointer-events', ...MOTION];
export default {
  url: URL,
  browser: 'chrome',
  wait: 2000,
  widths: [1600],
  probes: [
    ...open,
    { name: 'pause', select: 'button.playback-play__pause', props: [...SHOWN, 'color', 'width', 'height'], box: true },
    { name: 'pause.icon', select: 'button.playback-play__pause amp-icon[slot="pause"]', props: ['color', 'width', 'height'], box: true },
    { name: 'play.hidden', select: 'button.playback-play__play', props: SHOWN },
    { name: 'hidden', select: '.scrim__gradient', props: SHOWN, box: true, hover: [1599, 200], hovered: 6000 },
    { name: 'hidden.scrim', select: '.video-player-scrim', props: [...SHOWN, 'cursor'], matches: ['.video-player-scrim--hidden'] },
    { name: 'hidden.header', select: '.scrim__header', props: [...SHOWN, 'transform'] },
    { name: 'hidden.body', select: '.scrim__body', props: [...SHOWN, 'transform'] },
    { name: 'hidden.footer', select: '.scrim__footer', props: [...SHOWN, 'transform'] },
  ],
};
