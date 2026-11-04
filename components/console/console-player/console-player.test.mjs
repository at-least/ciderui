import { examples } from '../../../test/harness.mjs';

const ex = examples(import.meta.url);
const P = '.cui-console-player';
const BTN = '.cui-console-player-btn[aria-label="Subtitles"]';
const SCRUB = '.cui-console-player-progress > .cui-console-player-range';
const VOLUME = '.cui-console-player-header > .cui-console-player-range';
const TV = [1600, 1200];
const LAYOUT = 'tv.apple.com\'s player\'s room: 30px at the sides, 20px at the top and foot (tv-player:scrim\'s grid, 20px 26px 713px 121px 20px by 30px 1540px 30px); its rows the header, the centre and the footer, whatever the video\'s size';
const NO_SLIDE = 'faded where they are: tv.apple.com\'s header goes 20px up and its footer 20px down as they hide (tv-player-playing:hidden.*), a slide that only draws the eye (DESIGN.md decision 25); the fades\' durations and curves its own';

export const cases = [
  {
    name: 'player paused',
    html: ex('paused'),
    widths: TV,
    height: 900,
    checks: [
      { target: P, props: ['background-color'], apple: ['tv-player', 'dialog'] },
      { target: P, props: ['position', 'overflow-x'], golden: 'the containing block its overlay is laid over its video in; not clipped: a window too small for its controls grows it to them (its aspect-ratio\'s automatic minimum, which a clip would drop)' },
      { target: '.cui-console-player-overlay', props: ['display', 'grid-template-columns', 'grid-template-rows'], golden: LAYOUT },
      { target: '.cui-console-player-overlay', pseudo: '::before', props: ['background-image', 'opacity'], apple: ['tv-player', 'scrim.gradient'] },
      { target: '.cui-console-player-overlay', pseudo: '::before', props: ['transition-property', 'transition-duration'], apple: ['tv-player', 'scrim.gradient'] },
      { target: '.cui-console-player-play', props: ['color', 'cursor', 'opacity', 'background-color', 'border-top-width', 'padding-top'], apple: ['tv-player', 'play'] },
      { target: '.cui-console-player-play', props: ['box-width', 'box-height'], golden: 'tv.apple.com\'s centre button, 66px (tv-player:play\'s box: an apple check reads no width)' },
      { target: '.cui-console-player-play', props: ['transform'], state: 'active', apple: ['tv-player', 'play', 'active'] },
      { target: '.cui-console-player-play', props: ['transition-property', 'transition-duration', 'transition-timing-function'], golden: 'its press\'s scale in tv.apple.com\'s 0.2s ease-out (tv-player:play), and its fade as the controls hide, the centre\'s 0.4s (tv-player-playing:hidden.body)' },
      { target: '.cui-console-player-glyph-play', props: ['box-width', 'box-height', 'display'], golden: 'paused, the play glyph, tv.apple.com\'s 50px (tv-player:play.icon), Lucide\'s filled (decision 17: tv.apple.com\'s is an SF Symbol)' },
      { target: '.cui-console-player-glyph-pause', props: ['display'], golden: 'paused, no pause glyph: tv.apple.com\'s pause is hidden then (tv-player:pause.hidden)' },
      { target: '.cui-console-player-header', props: ['column-gap'], apple: ['tv-player', 'volume'] },
      { target: '.cui-console-player-header', props: ['display', 'justify-content', 'align-items'], golden: 'the volume at the header\'s end, as tv.apple.com\'s (tv-player:volume at the right)' },
      { target: '.cui-console-player-buttons', props: ['column-gap', 'display'], apple: ['tv-player', 'controls'] },
      { target: '.cui-console-player-footer', props: ['row-gap', 'column-gap', 'grid-template-areas'], golden: 'the title beside the buttons, the scrubber under them, 12px apart (tv-player:footer\'s 12px rows and gap)' },
      { target: BTN, props: ['color', 'cursor', 'background-color', 'border-top-width'], apple: ['tv-player', 'subtitles'] },
      { target: '.cui-console-player-btn[aria-label="Full Screen"]', props: ['color'], state: 'hover', apple: ['tv-player', 'fullscreen', 'hover'] },
      { target: BTN, props: ['color'], state: 'hover', golden: 'hovered white, as tv.apple.com\'s full screen button (tv-player:fullscreen:hover); its subtitles button reads no change, one look for the row' },
      { target: BTN, props: ['box-width', 'box-height'], golden: 'tv.apple.com\'s 24px (tv-player:audio, its controls\' row 24px)' },
      { target: `${BTN} > svg`, props: ['box-width', 'box-height'], golden: 'its glyph 20px, tv.apple.com\'s (tv-player:subtitles.icon); Lucide\'s (decision 17)' },
      { target: '.cui-console-player-header > .cui-console-player-btn', props: ['box-width', 'color'], golden: 'the mute button 26px in white, tv.apple.com\'s (tv-player:volume.button, its glyph filled white); its Lucide glyph 22px, a stroke glyph drawn smaller than a filled one' },
      { target: '.cui-console-player-title', props: ['font-size', 'line-height', 'font-weight', 'color'], apple: ['tv-player', 'title'] },
      { target: '.cui-console-player-time', props: ['font-size', 'line-height', 'font-weight', 'color', 'font-variant-numeric'], apple: ['tv-player', 'elapsed'] },
      { target: '.cui-console-player-time:last-of-type', props: ['color'], apple: ['tv-player', 'remaining'] },
      { target: SCRUB, props: ['--cui-console-player-track', '--cui-console-player-thumb'], apple: ['tv-player', 'scrubber'], as: { '--cui-console-player-track': '--progress-track-height', '--cui-console-player-thumb': '--progress-thumb-height' } },
      { target: SCRUB, props: ['cursor'], apple: ['tv-player', 'scrubber'] },
      { target: SCRUB, props: ['box-height', 'appearance', '--cui-console-player-value'], golden: 'the thumb\'s 11px, tv.apple.com\'s (tv-player:scrubber --progress-thumb-height); the share played the page\'s, written as a percentage' },
      { target: VOLUME, props: ['--cui-console-player-track', '--cui-console-player-thumb'], apple: ['tv-player', 'volume.range'], as: { '--cui-console-player-track': '--progress-track-height', '--cui-console-player-thumb': '--progress-thumb-height' } },
      { target: VOLUME, props: ['box-width'], golden: 'tv.apple.com\'s volume, 100px (tv-player:volume.range); its thumb white as the scrubber\'s, where tv.apple.com\'s is #d0d0d0 until hovered: one thumb' },
      { target: SCRUB, props: ['outline-style', 'outline-width', 'outline-offset'], state: 'focus-visible', golden: 'base.css\'s one focus ring (rule 9), where tv.apple.com draws a 4px blue one' },
    ],
  },
  {
    name: 'player playing',
    html: ex('playing'),
    widths: [1600],
    height: 900,
    checks: [
      { target: '.cui-console-player-glyph-pause', props: ['display', 'box-width'], golden: 'playing, the pause glyph in the play\'s place (tv-player-playing:pause, 50px)' },
      { target: '.cui-console-player-glyph-play', props: ['display'], golden: 'playing, no play glyph (tv-player-playing:play.hidden)' },
      { target: '.cui-console-player-play', props: ['color', 'opacity'], apple: ['tv-player-playing', 'pause'] },
      { target: '.cui-console-player-header', props: ['opacity'], golden: 'shown: data-controls is the page\'s to set' },
    ],
  },
  {
    name: 'player hidden',
    html: ex('hidden'),
    widths: [1600],
    height: 900,
    checks: [
      { target: '.cui-console-player-overlay', pseudo: '::before', props: ['opacity'], apple: ['tv-player-playing', 'hidden'] },
      { target: '.cui-console-player-header', props: ['opacity'], apple: ['tv-player-playing', 'hidden.header'] },
      { target: '.cui-console-player-play', props: ['opacity'], apple: ['tv-player-playing', 'hidden.body'] },
      { target: '.cui-console-player-footer', props: ['opacity'], apple: ['tv-player-playing', 'hidden.footer'] },
      { target: '.cui-console-player-header', props: ['transform', 'transition-duration', 'transition-timing-function'], golden: NO_SLIDE },
      { target: '.cui-console-player-footer', props: ['transform', 'transition-duration', 'transition-timing-function'], golden: NO_SLIDE },
      { target: '.cui-console-player-overlay', props: ['cursor'], golden: 'the pointer kept, as tv.apple.com keeps it (tv-player-playing:hidden.scrim)' },
    ],
  },
  {
    name: 'player hidden focused',
    html: ex('hidden'),
    widths: [1600],
    height: 900,
    press: ['Tab'],
    checks: [
      { target: VOLUME, props: ['focused'], golden: 'Tab reaches the hidden controls' },
      { target: '.cui-console-player-footer', props: ['opacity'], golden: 'shown while the focus is inside, for the keyboard (not read on tv.apple.com, whose script shows them on a key)' },
    ],
  },
  {
    name: 'player dark',
    html: `<div class="cui-dark">${ex('paused')}</div>`,
    widths: [1200],
    height: 900,
    checks: [
      { target: P, props: ['background-color'], apple: ['tv-player', 'dialog'] },
      { target: BTN, props: ['color'], apple: ['tv-player', 'subtitles'] },
      { target: '.cui-console-player-time', props: ['color'], apple: ['tv-player', 'elapsed'] },
    ],
  },
  {
    name: 'player touch',
    html: ex('paused'),
    touch: true,
    widths: [390],
    checks: [
      { target: BTN, props: ['box-width', 'box-height'], golden: 'a button 44px each way at least on a touch screen (the HIG\'s default, reference/touch/)' },
      { target: BTN, props: ['color'], state: 'hover', rest: 'a tapped button keeps no hover (rule 5)' },
      { target: SCRUB, props: ['box-height'], golden: 'a range 44px tall on a touch screen, its track and thumb as on a desktop\'s' },
    ],
  },
];
