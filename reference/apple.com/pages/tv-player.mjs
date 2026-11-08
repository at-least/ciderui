// @external: tv.apple.com's video player, a trailer paused with its controls shown (the
// owner, 2026-11-05: 「播放器照 tv.apple.com」): the first trailer opened from Ted Lasso's
// trailers, the video clicked (its controls shown), the centre pause clicked; paused,
// the controls stay. The video is clicked at a point: over it lies the trailers'
// carousel, which takes a click on an element. The overlay's gradient, its rows (the
// header's volume, the centre play, the footer's scrubber, times and buttons), each
// control at rest, hovered, pressed and focused. Read in Google Chrome (`browser`):
// Playwright's Chromium opens the player and never plays (its <video> never made).
// tv-player-playing reads the playing state, tv-player-menu the subtitles menu,
// tv-player-playlist the trailers' row.
import { TYPE, BOX, MOTION, FOCUS, FLEX } from './common.mjs';

export const URL = 'https://tv.apple.com/us/collection/trailers/uts.col.Trailers.umc.cmc.vtoh0mn0xn7t3c643xqonfzy';
const STATES = ['hover', 'active', 'focus-visible'];
const CONTROL = [...BOX, 'color', 'opacity', 'cursor', 'width', 'height', 'transform', ...FOCUS, ...MOTION];
// @external: the custom properties its range inputs are drawn by (a track, the part
// played, the thumb), read on the input: a thumb's computed style is not
// read (getComputedStyle with ::-webkit-slider-thumb gave the input's own)
const RANGE = ['--progress-track-height', '--progress-track-max-height', '--progress-track-color', '--progress-track-color-elapsed', '--progress-track-color-elapsed-active',
  '--progress-thumb-height', '--progress-thumb-width', '--progress-thumb-color', '--progress-thumb-color-active', '--progress-thumb-multiplier', '--progress-thumb-multiplier-active', '--progress-thumb-border-radius', '--progress-thumb-box-shadow', '--progress-thumb-outline-active', '--volume-thumb-box-shadow'];

// @invariant: opening the player, showing its controls and pausing it: each once, in turn
export const open = [
  { name: 'dialog', click: 'a.lockup', clicked: 6000, select: 'dialog.playback-view', props: ['background-color', 'color', 'padding-top', 'padding-left'], box: true },
  { name: 'scrim', click: [700, 400], clicked: 1000, select: '.video-player-scrim', props: [...FLEX, 'grid-template-rows', 'grid-template-columns', 'padding-top', 'padding-left', 'padding-right', 'padding-bottom', 'opacity', ...MOTION], box: true },
];
export const probes = [
  ...open,
  { name: 'play', click: 'button.playback-play__pause', clicked: 1500, pointer: [800, 200], select: 'button.playback-play__play', props: CONTROL, box: true, states: STATES },
  { name: 'play.icon', select: 'button.playback-play__play amp-icon[slot="play"]', props: ['color', 'width', 'height'], box: true },
  { name: 'pause.hidden', select: 'button.playback-play__pause', props: ['opacity', 'visibility', 'pointer-events', ...MOTION] },
  { name: 'scrim.gradient', select: '.scrim__gradient', props: ['background-image', 'opacity', ...MOTION], box: true, position: true },
  { name: 'scrim.header', select: '.scrim__header', props: ['opacity', ...MOTION], box: true, position: true },
  { name: 'scrim.footer', select: '.scrim__footer', props: ['opacity', 'padding-left', 'padding-right', 'padding-bottom', ...MOTION], box: true, position: true },
  { name: 'header', select: '.video-player__header', props: [...FLEX, 'grid-template-columns', 'padding-left', 'padding-right', 'padding-top'], box: true, position: true },
  { name: 'footer', select: '.video-player__footer', props: [...FLEX, 'grid-template-columns', 'grid-template-rows', 'grid-template-areas', 'padding-left', 'padding-right', 'padding-bottom'], box: true, position: true },
  { name: 'title', select: '.video-player__metadata .title', props: TYPE, box: true, position: true },
  { name: 'progress', select: '.video-player__progress', props: ['padding-top', 'padding-bottom', 'margin-top'], box: true, position: true },
  { name: 'progress.bar', select: 'amp-playback-controls-progress', props: [...FLEX, 'grid-template-columns', 'grid-template-rows', 'grid-template-areas'], box: true, position: true },
  { name: 'elapsed', select: 'amp-playback-controls-progress >>> time.elapsed', props: [...TYPE, 'font-variant-numeric', 'font-feature-settings', 'text-align', 'grid-area'], box: true, position: true },
  { name: 'remaining', select: 'amp-playback-controls-progress >>> time.remaining', props: [...TYPE, 'font-variant-numeric', 'font-feature-settings', 'text-align', 'grid-area'], box: true, position: true },
  { name: 'scrubber', select: 'amp-playback-controls-progress >>> input[type="range"]', props: ['height', 'cursor', 'background-color', 'color', ...FOCUS], vars: RANGE, box: true, position: true, states: STATES },
  { name: 'volume', select: '.video-player__volume', props: [...FLEX, 'grid-template-columns'], box: true, position: true },
  { name: 'volume.range', select: 'amp-volume-control >>> input[type="range"]', props: ['height', 'width', 'cursor', 'color', ...FOCUS], vars: RANGE, box: true, position: true, states: STATES },
  { name: 'volume.button', select: 'button.volume-unified__indicator', props: CONTROL, box: true, position: true, states: STATES },
  { name: 'volume.icon', select: 'button.volume-unified__indicator svg', props: ['color', 'fill', 'width', 'height'], box: true },
  { name: 'controls', select: '.video-player__controls', props: [...FLEX, 'grid-area'], box: true, position: true },
  { name: 'subtitles', select: 'amp-captions-control >>> button.contextual-menu__trigger', props: [...CONTROL, 'backdrop-filter'], box: true, position: true, states: STATES },
  { name: 'subtitles.icon', select: 'amp-captions-control >>> amp-icon', props: ['color', 'width', 'height'], box: true },
  { name: 'audio', select: 'amp-audio-tracks-control >>> label', props: CONTROL, box: true, position: true, states: ['hover'] },
  { name: 'audio.select', select: 'amp-audio-tracks-control >>> select', props: ['opacity', 'color', 'appearance', 'cursor', 'position', 'width', 'height', ...FOCUS], box: true, states: ['hover', 'focus-visible'] },
  { name: 'fullscreen', select: 'amp-playback-controls-full-screen >>> button', props: CONTROL, box: true, position: true, states: STATES },
  { name: 'fullscreen.icon', select: 'amp-playback-controls-full-screen >>> amp-icon', props: ['color', 'width', 'height'], box: true },
];
export default {
  url: URL,
  browser: 'chrome',
  wait: 2000,
  widths: [1600, 1200],
  probes,
};
