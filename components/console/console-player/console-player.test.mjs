import { examples } from '../../../test/harness.mjs';

const ex = examples(import.meta.url);
const P = '.cui-console-player';
const BTN = '.cui-console-player-btn[aria-label="Subtitles"]';
const SCRUB = '.cui-console-player-progress > .cui-console-player-range';
const VOLUME = '.cui-console-player-volume > .cui-console-player-range';
const OPEN = 'the owner\'s (2026-11-05: only the speaker until hovered, the range opening beside it): tv.apple.com\'s range is 100px open beside its button with the pointer elsewhere (tv-player:volume.range, the pointer at 800,200), so none of the closing is measured; it opens over 0.3s on the header\'s own fade curve (tv-player-playing:hidden.header)';
const TV = [1600, 1200];
const LAYOUT = 'tv.apple.com\'s player\'s room: 30px at the sides, 20px at the top and foot (tv-player:scrim\'s grid, 20px 26px 713px 121px 20px by 30px 1540px 30px); its rows the header, the centre and the footer, whatever the video\'s size';
const STAGE = '.cui-console-player > :first-child';
const LINE = '.cui-console-player-subtitles > p';
const CUE = '.cui-console-player-cue';
const MUSE = 'muse\'s console\'s subtitles, not tv.apple.com\'s, whose cues are its <video>\'s own text track (shown, none active in the first 22s of its first trailer, 2026-11-05; ::cue unreadable): 2.4% of the width at the least 15px, so it grows with the player (full screen too), 1.6 lines, 500, white under a soft shadow (--cui-video-cue-shadow) and no box (the owner, 2026-11-05, from Apple\'s view), 42em at its widest, balanced';
const hidden = (html) => html.replace('data-state="playing"', 'data-state="playing" data-controls="hidden"');
// @external: muse's player: its subtitles' Off and eleven languages, and no voices (one)
// or two, the player in a column beside a 288px sidebar under a 50px bar, as
// wide as the window's height leaves (muse's player_page.go)
const LANGUAGES = [['zh-Hant', '繁體中文'], ['zh-Hans', '简体中文'], ['en', 'English'], ['ja', '日本語'], ['ko', '한국어'], ['es', 'Español'], ['pt-BR', 'Português (Brasil)'], ['de', 'Deutsch'], ['fr', 'Français'], ['th', 'ภาษาไทย'], ['vi', 'Tiếng Việt']];
const languages = (html, voices = 1) => {
  const off = html.match(/<button class="cui-console-menu-item"[^>]*value="off"[\s\S]*?<\/button>/)[0];
  const choice = ([lang, name], i) => off.replace('aria-checked="false"', `aria-checked="${i === 0}"`).replace('value="off"', `value="${lang}"`)
    .replace('<span class="cui-console-menu-label">Off</span>', `<span class="cui-console-menu-label" lang="${lang}">${name}</span>`);
  const subtitles = /(<div role="menu" aria-labelledby="subtitles-subtitles-title">)[\s\S]*?(<\/div><\/div><\/div><\/div><\/div>$)/;
  if (!subtitles.test(html)) throw new Error('the stills example has no subtitles column where languages() looks');
  const voicesColumn = /<div class="cui-console-menu" data-column="voices">[\s\S]*?<\/div><\/div>\n/;
  if (!voicesColumn.test(html)) throw new Error('the stills example has no voices column where languages() looks');
  const voiced = (column) => column.replace(/(<div role="menu" aria-labelledby="subtitles-voices-title">)[\s\S]*?(<\/div><\/div>\n)$/,
    (_, start, end) => `${start}\n${LANGUAGES.slice(0, voices).map(choice).join('\n')}${end}`);
  const player = html.replace(subtitles, (_, start, end) => `${start}\n${[off, ...LANGUAGES.map(choice)].join('\n')}${end}`)
    .replace(voicesColumn, (column) => voices > 1 ? voiced(column) : '');
  return `<div style="height:50px"></div><div style="display:flex"><div style="flex:1;min-width:0;padding:0 24px"><div style="margin:16px auto 0;max-width:calc((100vh - 14rem) * 16 / 9)">${player}</div></div><div style="flex:none;width:288px"></div></div>`;
};
const LAST = '#subtitles [value="vi"]';
// @invariant: the preview as the page's script shows it: its picture, its line and the pointer's x
const previewing = (html) => html
  .replace('aria-hidden="true" hidden>', 'aria-hidden="true" style="--cui-console-player-preview-x: 300px">')
  .replace('<img class="cui-console-player-preview-picture" alt="">', `<img class="cui-console-player-preview-picture" alt="" src="${html.match(/<img class="absolute[^>]* src="([^"]+)"/)[1]}">`)
  .replace('<p class="cui-console-player-preview-label"></p>', '<p class="cui-console-player-preview-label">場景 4　大野狼到外婆家 · 2:31</p>');
const PREVIEW = '.cui-console-player-preview';
const PREVIEW_WHY = 'muse\'s, tv.apple.com\'s trailers scrubbing with no preview (the owner, 2026-11-05: the shot\'s thumbnail, its scene and the time over the pointer): the storyboard\'s 320px thumbnail at 160px, sharp at 2x; 12px above the bar, clear of the thumb; the line under it 12px semibold in the times\' tabular figures, white under the cue\'s shadow, on one line however long the scene\'s name';
const TEXT_BTN = '.cui-console-player-text-btn';
const FULL = 'muse\'s, for its 回報問題 and 下一個要看 while the player fills the screen (the owner, 2026-11-05); tv.apple.com\'s bar has no control of words: a 24px capsule, the bar\'s buttons\' height (tv-player:audio), white on a 16% white fill, 24% hovered or pressed, 13px medium, 8px apart; half opacity disabled';
const HEADING = '.cui-console-player-heading';
// @invariant: the page's script's heading, written into the player at rest
const withHeading = (html) => html.replace('data-control="heading"></p>', 'data-control="heading">The Wedding · Scene 2　The Vows · S02-03</p>');
const PAUSED = 'muse\'s (the owner, 2026-11-05: paused, a small line at the picture\'s top left saying the story, the scene and the shot; gone on play): in tv.apple.com\'s header\'s first column, 711px and empty on its trailers (tv-player:header, grid-template-columns 711px 829px), on the header\'s 26px row; 15px semibold white under the cue\'s shadow (--cui-video-cue-shadow) so it reads on a bright picture with no box, cut short with an ellipsis before the volume; tv.apple.com shows no such line';
const INTRO = 'muse\'s pre-play layer (the owner, 2026-11-05: over the cover the title large, the demo\'s length, its description\'s first sentence, one clear Play; tv.apple.com\'s pre-play page was not measured, so golden): at the picture\'s foot in the overlay\'s room (30px at the sides, tv-player:scrim), on a soft gradient in that corner alone (--cui-video-title-scrim; the owner, 2026-11-08: tv.apple.com\'s, --cui-video-scrim, darkened the cover\'s lower half to black), the lines white under the cue\'s shadow; the length 15px medium and the sentence 17px on 24px in the controls\' 90% white, two lines at most, 36em wide';
const END = 'muse\'s end layer (the owner, 2026-11-05: on end a 「看完了」 layer fades in over the last picture, the user\'s next step under it and a small 「重新播放」; tv.apple.com\'s end of a film was not measured, so golden): the intro\'s layer, its place, gradient and title, in the other direction; the next step in the length\'s type (15px medium, the controls\' 90% white), its buttons and links one under another 8px apart, 14px under the title as the intro\'s Play is';
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
      { target: '.cui-console-player-volume', props: ['column-gap'], apple: ['tv-player', 'volume'] },
      { target: '.cui-console-player-header', props: ['display', 'justify-content', 'align-items'], golden: 'the volume at the header\'s end, as tv.apple.com\'s (tv-player:volume at the right)' },
      { target: '.cui-console-player-buttons', props: ['column-gap', 'display'], apple: ['tv-player', 'controls'] },
      { target: '.cui-console-player-footer', props: ['row-gap', 'column-gap', 'grid-template-areas'], golden: 'the buttons at the row\'s end with nothing before them (no title: the heading at the top says what plays), the scrubber under them, 12px apart (tv-player:footer\'s 12px rows and gap)' },
      { target: BTN, props: ['color', 'cursor', 'background-color', 'border-top-width'], apple: ['tv-player', 'subtitles'] },
      { target: '.cui-console-player-btn[aria-label="Full Screen"]', props: ['color'], state: 'hover', apple: ['tv-player', 'fullscreen', 'hover'] },
      { target: BTN, props: ['color'], state: 'hover', golden: 'hovered white, as tv.apple.com\'s full screen button (tv-player:fullscreen:hover); its subtitles button reads no change, one look for the row' },
      { target: BTN, props: ['box-width', 'box-height'], golden: 'tv.apple.com\'s 24px (tv-player:audio, its controls\' row 24px)' },
      { target: `${BTN} > svg`, props: ['box-width', 'box-height'], golden: 'its glyph 20px, tv.apple.com\'s (tv-player:subtitles.icon); Lucide\'s (decision 17)' },
      { target: '.cui-console-player-volume > .cui-console-player-btn', props: ['box-width', 'color'], golden: 'the mute button 26px in white, tv.apple.com\'s (tv-player:volume.button, its glyph filled white); its Lucide glyph 22px, a stroke glyph drawn smaller than a filled one' },
      { target: '.cui-console-player-time', props: ['font-size', 'line-height', 'font-weight', 'color', 'font-variant-numeric'], apple: ['tv-player', 'elapsed'] },
      { target: '.cui-console-player-time:last-of-type', props: ['color'], apple: ['tv-player', 'remaining'] },
      { target: SCRUB, props: ['--cui-console-player-track', '--cui-console-player-thumb'], apple: ['tv-player', 'scrubber'], as: { '--cui-console-player-track': '--progress-track-height', '--cui-console-player-thumb': '--progress-thumb-height' } },
      { target: SCRUB, props: ['cursor'], apple: ['tv-player', 'scrubber'] },
      { target: SCRUB, props: ['box-height', 'appearance', '--cui-console-player-value'], golden: 'a 44px target for any pointer, the HIG\'s default (reference/touch/), where tv.apple.com\'s is 25px (tv-player:scrubber\'s box); the share played the page\'s, written as a percentage' },
      { target: SCRUB, props: ['position'], golden: 'positioned, over the times\' text its target reaches down into: unpositioned, a point on the time played hit the <time>, not the range' },
      { target: SCRUB, props: ['margin-top', 'margin-bottom'], golden: 'its margins give back all but the thumb\'s 11px, tv.apple.com\'s row (tv-player:progress.bar 11px 14px): the bar sits where it did, the line as thin' },
      { target: '.cui-console-player-progress', props: ['grid-template-rows'], apple: ['tv-player', 'progress.bar'] },
      { target: SCRUB, props: ['--cui-console-player-track'], state: 'active', golden: 'held, the line 8px, thicker under the thumb (the owner, 2026-11-05: 「thicker while dragged」); tv.apple.com\'s forced :active reads 5px (tv-player:scrubber:active), its drag drawn by its script' },
      { target: '.cui-console-player-buttons', props: ['z-index', 'position'], golden: 'over the scrubber\'s target, which reaches 4.5px up into their row: a click on a button\'s foot is the button\'s' },
      { target: PREVIEW, props: ['display', 'pointer-events'], golden: 'the preview hidden until the page\'s script shows it, and never in the pointer\'s way' },
      { target: VOLUME, props: ['--cui-console-player-track', '--cui-console-player-thumb'], apple: ['tv-player', 'volume.range'], as: { '--cui-console-player-track': '--progress-track-height', '--cui-console-player-thumb': '--progress-thumb-height' } },
      { target: VOLUME, props: ['box-width', 'opacity', 'margin-right'], golden: 'closed, the speaker alone: no range and none of the 8px gap before the button; ' + OPEN },
      { target: VOLUME, props: ['pointer-events'], golden: 'closed, it takes no pointer: its 11px thumb overflows its 0px over the button, and took 6px of the button\'s target (the touch audit read the button 38px wide)' },
      { target: VOLUME, props: ['transition-property', 'transition-duration', 'transition-timing-function'], golden: OPEN },
      { target: '.cui-console-player-glyph-volume', props: ['display'], golden: 'not muted, the speaker' },
      { target: '.cui-console-player-glyph-muted', props: ['display'], golden: 'not muted, no struck speaker' },
      { target: CUE, props: ['display'], golden: 'no cue, no box: the slot is written empty until the page\'s script fills it' },
      { target: '#subtitles', props: ['display'], golden: 'its menu written inside it, closed until its button opens it' },
    ],
  },
  {
    name: 'player stills',
    html: ex('stills'),
    widths: TV,
    height: 900,
    checks: [
      { target: STAGE, props: ['box-width', 'box-height', 'box-top', 'box-left'], golden: 'an element of stills in the video\'s place, the player\'s first child: the player\'s whole box, under the overlay' },
      { target: `${STAGE} > img`, props: ['box-width', 'box-height'], golden: 'the page\'s stills fill it (its own utilities: the player sizes the element, not what is in it)' },
      { target: LINE, props: ['font-size', 'line-height', 'font-weight', 'color', 'text-align', 'text-wrap-style', 'max-width'], golden: MUSE },
      { target: CUE, props: ['text-shadow', 'background-color', 'padding-top', 'padding-left', 'border-top-left-radius'], golden: MUSE },
      { target: LINE, props: ['box-top', 'box-height'], golden: 'over the footer while the controls show: its foot 12px above the footer\'s top, the footer\'s own gap (tv-player:footer\'s 12px)' },
      { target: '.cui-console-player-footer', props: ['box-top'], golden: 'the footer, under the subtitles, where it is with none' },
      { target: '.cui-console-player-subtitles', props: ['pointer-events', 'transition-property', 'transition-duration', 'transition-timing-function'], golden: 'the controls under it take the pointer; it moves as the footer fades, in the footer\'s 0.4s (tv-player-playing:hidden.footer)' },
      { target: CUE, props: ['hit'], golden: 'a click on the cue reaches the overlay under it, which the page\'s script hears (a click on the picture plays or pauses)' },
    ],
  },
  {
    name: 'player scrub preview',
    html: previewing(ex('stills')),
    widths: TV,
    height: 900,
    checks: [
      { target: PREVIEW, props: ['box-width', 'box-height', 'box-left', 'box-top', 'display', 'row-gap', 'justify-items'], golden: PREVIEW_WHY },
      { target: '.cui-console-player-preview-picture', props: ['box-width', 'box-height', 'border-top-left-radius', 'object-fit', 'background-color', 'box-shadow'], golden: PREVIEW_WHY },
      { target: '.cui-console-player-preview-label', props: ['box-width', 'box-left'], golden: 'a line longer than the picture centred under it, over both its sides' },
      { target: '.cui-console-player-preview-label', props: ['font-size', 'line-height', 'font-weight', 'color', 'white-space', 'font-variant-numeric', 'text-shadow'], golden: PREVIEW_WHY },
      { target: PREVIEW, props: ['hit'], golden: 'it takes no pointer: what is under it does (the overlay), so it never flickers under the pointer that shows it' },
    ],
  },
  {
    name: 'player scrub preview at the start',
    html: previewing(ex('stills')).replace('--cui-console-player-preview-x: 300px', '--cui-console-player-preview-x: 4px'),
    widths: [1600],
    height: 900,
    checks: [
      { target: PREVIEW, props: ['box-left'], golden: 'the pointer near the bar\'s start: the preview stops at it, inside the bar (the bar at 30px)' },
    ],
  },
  {
    name: 'player volume open',
    html: ex('paused'),
    widths: TV,
    height: 900,
    hover: '[data-control="mute"]',
    checks: [
      { target: VOLUME, props: ['pointer-events'], golden: 'open, it takes the pointer' },
      { target: VOLUME, props: ['box-width', 'opacity', 'margin-right'], golden: 'the pointer over its button, the range open: tv.apple.com\'s 100px (tv-player:volume.range), its thumb white as the scrubber\'s, where tv.apple.com\'s is #d0d0d0 until hovered: one thumb; ' + OPEN },
      { target: '.cui-console-player-volume > .cui-console-player-btn', props: ['box-left'], golden: 'the button where it was: the range opens to its left (tv-player:volume.button at the header\'s end)' },
    ],
  },
  {
    name: 'player muted',
    html: ex('paused').replace('<div class="cui-console-player aspect-video">', '<div class="cui-console-player aspect-video" data-muted>'),
    widths: [1600],
    height: 900,
    checks: [
      { target: '.cui-console-player-glyph-muted', props: ['display', 'box-width'], golden: 'muted, the speaker struck through, Lucide\'s volume-x at the speaker\'s 22px (tv.apple.com\'s muted glyph not read)' },
      { target: '.cui-console-player-glyph-volume', props: ['display'], golden: 'muted, no plain speaker' },
    ],
  },
  {
    name: 'player heading paused',
    html: withHeading(ex('paused')),
    widths: TV,
    height: 900,
    checks: [
      { target: HEADING, props: ['display', 'opacity', 'visibility', 'font-size', 'line-height', 'font-weight', 'color', 'text-shadow', 'white-space', 'text-overflow', 'box-left', 'box-top', 'box-height'], golden: PAUSED },
      { target: HEADING, props: ['transition-property', 'transition-duration', 'transition-timing-function'], golden: 'it comes and goes in the header\'s fade (tv-player-playing:hidden.header, 0.3s)' },
    ],
  },
  {
    name: 'player heading playing',
    html: ex('playing'),
    widths: [1600],
    height: 900,
    checks: [
      { target: VOLUME, on: '.cui-console-player-volume > .cui-console-player-btn', state: 'active', props: ['margin-right'], golden: 'its range stays open while a control in it is pressed' },
      { target: BTN, state: 'active', props: ['color'], golden: 'pressed, full white, as the pointer over it shows on a screen that has one' },
      { target: HEADING, props: ['opacity', 'visibility'], golden: 'gone while playing (the owner, 2026-11-05), the controls shown or not' },
    ],
  },
  {
    name: 'player heading empty',
    html: ex('paused'),
    widths: [1600],
    height: 900,
    checks: [
      { target: '.cui-console-player-volume > .cui-console-player-btn', props: ['box-height'], golden: 'the speaker a little taller than the bar\'s other buttons' },
      { target: HEADING, props: ['display'], golden: 'empty, nothing drawn: the volume keeps the header\'s end' },
    ],
  },
  {
    name: 'player full screen glyph',
    html: ex('full-screen'),
    widths: [1600],
    height: 900,
    checks: [
      { target: '[data-control="full-screen"] .cui-console-player-glyph-exit-full-screen', props: ['display', 'box-width'], golden: 'filling the screen, the button shows the way out (Lucide\'s minimize-2 for its maximize-2), as tv.apple.com\'s full-screen button turns to exit (its glyph not read: a page cannot be put in full screen by the probe)' },
      { target: '[data-control="full-screen"] .cui-console-player-glyph-full-screen', props: ['display'], golden: 'and not the way in' },
    ],
  },
  {
    name: 'player full screen glyph not full screen',
    html: ex('paused'),
    widths: [1600],
    height: 900,
    checks: [
      { target: '[data-control="full-screen"] .cui-console-player-glyph-full-screen', props: ['display', 'box-width'], golden: 'not full screen, the way in' },
      { target: '[data-control="full-screen"] .cui-console-player-glyph-exit-full-screen', props: ['display'], golden: 'and not the way out' },
    ],
  },
  {
    name: 'player intro',
    html: ex('intro'),
    widths: TV,
    height: 900,
    checks: [
      { target: '.cui-console-player-intro-play', state: 'active', props: ['transform', 'background-color'], golden: 'pressed, the hovered fill, the controls\' 90% white (--cui-c-on-video-control), and no scale: the owner, 2026-11-09 (the centre play\'s scale is tv.apple.com\'s, measured; this button\'s had no measure)' },
      { target: '.cui-console-player-intro', props: ['display', 'flex-direction', 'justify-content', 'align-items', 'padding-left', 'padding-bottom', 'background-image', 'opacity', 'visibility'], golden: INTRO },
      { target: '.cui-console-player-overlay', props: ['visibility'], golden: 'its controls hidden under it: one Play, nothing else to press (the owner, 2026-11-05)' },
      { target: '.cui-console-player-intro-title', props: ['font-size', 'line-height', 'font-weight', 'color', 'box-left'], golden: `${INTRO}; the title as large as 4% of the player, at the least tv.apple.com's 26px bold title (tv-player:title)` },
      { target: '.cui-console-player-intro-detail', props: ['font-size', 'line-height', 'font-weight', 'color'], golden: INTRO },
      { target: '.cui-console-player-intro-summary', props: ['font-size', 'line-height', 'color', 'max-width', '-webkit-line-clamp'], golden: INTRO },
      { target: '.cui-console-player-intro-play', props: ['box-height', 'padding-left', 'border-top-left-radius', 'background-color', 'color', 'font-size', 'font-weight', 'column-gap'], golden: `${INTRO}; Play a white 44px capsule in the picture's black, the HIG's 44pt target` },
      { target: '.cui-console-player-intro-play', props: ['background-color'], state: 'hover', golden: 'hovered, the controls\' 90% white (--cui-c-on-video-control)' },
      { target: '.cui-console-player-intro [data-control="restart"]', props: ['display'], golden: 'the smaller start-over hidden until the page\'s script finds a saved place' },
      { target: '.cui-console-player-intro', props: ['transition-property', 'transition-duration'], golden: 'it fades out in the subtitles\' 0.4s curve (tv-player-playing:hidden.footer)' },
    ],
  },
  {
    name: 'player intro resume',
    html: ex('intro').replace('data-control="restart" hidden', 'data-control="restart"'),
    widths: [1600],
    height: 900,
    checks: [
      { target: '.cui-console-player-intro-actions', props: ['align-items', 'column-gap'], golden: 'its buttons on one center line, apart' },
      { target: '.cui-console-player-intro [data-control="restart"]', props: ['display', 'box-height', 'background-color'], golden: 'a saved place: start over beside Play, the bar\'s capsule of words, smaller' },
    ],
  },
  {
    name: 'player intro gone',
    html: ex('intro').replace(' data-intro', ''),
    widths: [1600],
    height: 900,
    checks: [
      { target: '.cui-console-player-intro', props: ['opacity', 'visibility', 'pointer-events'], golden: 'playing (data-intro dropped by the page\'s script), the layer gone, taking no click while it fades (muse: a click on Play 0.1s after the intro went landed on the fading intro)' },
      { target: '.cui-console-player-overlay', props: ['visibility'], golden: 'and the controls back' },
    ],
  },
  {
    name: 'player end',
    html: ex('end'),
    widths: TV,
    height: 900,
    checks: [
      { target: '.cui-console-player-end', props: ['display', 'flex-direction', 'justify-content', 'align-items', 'padding-left', 'padding-bottom', 'background-image', 'opacity', 'visibility'], golden: END },
      { target: '.cui-console-player-overlay', props: ['visibility'], golden: 'its controls hidden under it: the next step and Play Again, nothing else to press' },
      { target: '.cui-console-player-end-title', props: ['font-size', 'line-height', 'font-weight', 'color', 'box-left'], golden: `${END}; the title the intro's, 4% of the player and 26px at the least` },
      { target: '.cui-console-player-end-body', props: ['display', 'flex-direction', 'align-items', 'row-gap', 'margin-top', 'max-width', 'font-size', 'line-height', 'font-weight', 'color'], golden: END },
      { target: '.cui-console-player-end-body > p', props: ['margin-top', 'margin-bottom'], golden: 'its words without a paragraph\'s margin: the 8px gap spaces them' },
      { target: '.cui-console-player-end [data-control="replay"]', props: ['display', 'box-height', 'background-color', 'margin-top'], golden: 'Play Again, the bar\'s capsule of words, smaller than the next step, 14px under it' },
      { target: '.cui-console-player-end', props: ['transition-property', 'transition-duration'], golden: 'it fades in in the subtitles\' 0.4s curve, as the intro fades out' },
    ],
  },
  {
    name: 'player end not ended',
    html: ex('end').replace(' data-end', ''),
    widths: [1600],
    height: 900,
    checks: [
      { target: '.cui-console-player-end', props: ['opacity', 'visibility', 'pointer-events'], golden: 'before the end (data-end not set, or dropped as it plays again), the layer gone, taking no click while it fades (muse: a click on Play 0.1s after the intro went landed on the fading intro)' },
      { target: '.cui-console-player-overlay', props: ['visibility'], golden: 'and the controls there' },
    ],
  },
  {
    name: 'player end empty',
    html: ex('end').replace(/<div class="cui-console-player-end-body">[\s\S]*?<\/div>\n?<button class="cui-console-player-text-btn" type="button" data-control="replay">/, '<div class="cui-console-player-end-body"></div><button class="cui-console-player-text-btn" type="button" data-control="replay">'),
    widths: [1600],
    height: 900,
    checks: [
      { target: '.cui-console-player-end-body', props: ['display'], golden: 'no next step yet (the page fills it as the film ends): no room kept for it' },
    ],
  },
  {
    name: 'player full screen controls',
    html: ex('full-screen'),
    widths: TV,
    height: 900,
    checks: [
      { target: TEXT_BTN, state: 'active', props: ['background-color'], golden: 'pressed, the hovered fill' },
      { target: '.cui-console-player-full-screen-controls', props: ['display', 'column-gap', 'box-height'], golden: FULL },
      { target: TEXT_BTN, props: ['box-height', 'padding-left', 'border-top-left-radius', 'background-color', 'color', 'font-size', 'line-height', 'font-weight', 'white-space', 'cursor'], golden: FULL },
      { target: TEXT_BTN, props: ['background-color'], state: 'hover', golden: FULL },
      { target: '.cui-console-player-full-screen-controls', props: ['box-left'], golden: 'before the bar\'s buttons, in their row: the page\'s controls first, the player\'s own at the end' },
    ],
  },
  {
    // @external: no example has a disabled control of words: the page disables one
    // (muse's 下一個要看 with nothing next)
    name: 'player full screen control disabled',
    html: ex('full-screen').replace('type="button">Next to Watch', 'type="button" disabled>Next to Watch'),
    widths: [1600],
    height: 900,
    checks: [
      { target: `${TEXT_BTN}:disabled`, props: ['opacity', 'cursor', 'background-color'], golden: FULL },
      { target: `${TEXT_BTN}:disabled`, state: 'hover', props: ['background-color'], golden: 'disabled, the fill at rest under the pointer' },
    ],
  },
  {
    name: 'player full screen controls not full screen',
    html: ex('full-screen').replace(' data-fullscreen', ''),
    widths: [1600],
    height: 900,
    checks: [
      { target: '.cui-console-player-full-screen-controls', props: ['display'], golden: 'not full screen, the page\'s controls hidden: the page shows them elsewhere (muse\'s under the player)' },
    ],
  },
  {
    name: 'player full screen controls touch',
    html: ex('full-screen'),
    touch: true,
    widths: [390],
    checks: [
      { target: TEXT_BTN, props: ['box-height'], golden: 'a control of words 44px tall on a touch screen (the HIG\'s default, reference/touch/)' },
      { target: '.cui-console-player-buttons', props: ['box-right'], golden: 'a narrow player\'s bar inside it: the controls of words give way, cut short with an ellipsis, the player\'s own buttons whole (its right edge the overlay\'s 30px in)' },
      { target: TEXT_BTN, props: ['text-overflow', 'overflow-x'], golden: 'cut short with an ellipsis where the bar runs out of room' },
    ],
  },
  {
    name: 'player stills hidden',
    html: hidden(ex('stills')),
    widths: [1600],
    height: 900,
    checks: [
      { target: LINE, props: ['box-top', 'box-height'], golden: 'the controls hidden, the subtitles at the foot: 20px up, the overlay\'s room (the footer\'s row\'s foot)' },
      { target: '.cui-console-player-subtitles', props: ['opacity'], golden: 'shown while the controls are hidden' },
      { target: '.cui-console-player-footer', props: ['opacity'], apple: ['tv-player-playing', 'hidden.footer'] },
    ],
  },
  {
    name: 'player stills hidden focused',
    html: hidden(ex('stills')),
    widths: [1600],
    height: 900,
    press: ['Tab'],
    checks: [
      { target: LINE, props: ['box-top'], golden: 'the focus inside shows the controls, and the subtitles go back over the footer' },
    ],
  },
  {
    name: 'player stills phone',
    html: ex('stills'),
    widths: [390],
    checks: [
      { target: LINE, props: ['font-size'], golden: 'at the least 15px, where 2.4% of a narrow player would be smaller' },
    ],
  },
  {
    name: 'player menu open',
    html: hidden(ex('stills')),
    widths: [1600],
    height: 900,
    click: '[data-control="subtitles"]',
    checks: [
      { target: '#subtitles', props: ['display'], golden: 'the subtitles button opens its menu, a console popover (decision 27)' },
      { target: '#subtitles', props: ['box-top', 'box-height'], golden: 'above its button, its foot on the button\'s centre (the player\'s menus open upward, player menu languages)' },
      { target: '.cui-console-player-footer', props: ['opacity'], golden: 'the controls hidden before, shown while its menu is open: a click leaves no keyboard focus, the open menu keeps them' },
    ],
  },
  {
    name: 'player menu languages',
    html: languages(ex('stills')),
    widths: [1440],
    height: 900,
    click: '[data-control="subtitles"]',
    checks: [
      { target: '.cui-console-player-subtitles', props: ['align-items'], golden: 'a cue centered over the picture' },
      { target: '.cui-console-player', props: ['box-left', 'box-top', 'box-width', 'box-height'], golden: 'muse\'s player at 1440 by 900: 1104px wide at 24px, 66px down, beside its sidebar' },
      { target: BTN, props: ['box-left', 'box-top'], golden: 'the subtitles button, 24px: its centre the menu\'s corner' },
      { target: '#subtitles', props: ['box-x', 'box-y', 'box-width', 'box-height'], golden: 'inside the player (its parent: box-x and box-y from its corner): opening upward, its foot and its right edge on its button\'s centre, as tv.apple.com\'s (tv-player-menu:menu at 1258, 637, 185 by 142, its corner 1443, 779; tv-player:subtitles at 1434, 769, 20px: centred on 1444, 779); every choice in it, none scrolled' },
      { target: '#subtitles', props: ['max-height'], golden: 'as tall as the player leaves above its button, 20px from its top (the overlay\'s room), not the toolkit\'s 435px: tv.apple.com\'s menu has none (tv-player-menu:menu, overflow visible)' },
    ],
  },
  {
    name: 'player menu voices and languages narrow',
    html: languages(ex('stills'), 2),
    widths: [735],
    height: 900,
    click: '[data-control="subtitles"]',
    checks: [
      { target: '#subtitles', props: ['width'], golden: 'as wide as its columns, on a window too narrow to hold them in the browser\'s own width' },
    ],
  },
  {
    name: 'player menu voices and languages',
    html: languages(ex('stills'), 2),
    widths: [1440],
    height: 900,
    click: '[data-control="subtitles"]',
    checks: [
      { target: '#subtitles', props: ['box-x', 'box-y', 'box-width', 'box-height'], golden: 'one menu of two columns, muse\'s voices beside its subtitles (the owner, 2026-11-05: one button, one panel, as the Apple TV app\'s Audio and Subtitles; tv.apple.com\'s own menu is one column of groups, its audio a platform <select>, tv-player-menu): inside the player, opening upward, its foot and right edge on its button\'s centre as one column\'s' },
      { target: '#subtitles [data-column="voices"]', props: ['box-x', 'box-width', 'box-height'], golden: 'the voices first, at the start, as wide as its longest name asks' },
      { target: '#subtitles [data-column="subtitles"]', props: ['box-x', 'box-width', 'box-height', 'border-left-width', 'border-left-style', 'padding-left', 'margin-left'], golden: 'the subtitles beside them, after a rule in the popover\'s 8px gap, each column as tall as the menu' },
    ],
  },
  {
    name: 'player menu languages short',
    html: languages(ex('stills')),
    widths: [1440],
    height: 700,
    click: '[data-control="subtitles"]',
    press: Array(12).fill('Tab'),
    checks: [
      { target: '.cui-console-player', props: ['box-width', 'box-height'], golden: 'muse\'s player at 1440 by 700: 846px wide, too short for twelve choices above its button' },
      { target: '#subtitles', props: ['box-x', 'box-y', 'box-width', 'box-height', 'overflow-y'], golden: 'still inside the player: 20px from its top (the overlay\'s room) to its button\'s centre, the choices scrolled inside it' },
      { target: LAST, props: ['focused', 'hit'], golden: 'the last choice reached: the focus Tabbed to it scrolls it into the menu, the pointer finds it there' },
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
      { target: VOLUME, props: ['height'], golden: 'its 44px target, whatever the line it draws' },
      { target: '.cui-console-player-volume', props: ['align-items'], golden: 'the speaker and its range on one center line' },
      { target: VOLUME, props: ['focused'], golden: 'Tab reaches the hidden controls' },
      { target: VOLUME, props: ['box-width', 'opacity'], golden: 'focused from the keyboard, the range opens: a range that took the focus closed would be adjusted unseen' },
      { target: '.cui-console-player-footer', props: ['opacity'], golden: 'shown while the focus is inside, for the keyboard (not read on tv.apple.com, whose script shows them on a key)' },
    ],
  },
  {
    name: 'player hidden clicked',
    html: ex('hidden'),
    widths: [1600],
    height: 900,
    click: '.cui-console-player-play',
    checks: [
      { target: '.cui-console-player-play', props: ['focused'], golden: 'the click leaves the focus on Play, as Chrome does' },
      { target: '.cui-console-player-overlay', pseudo: '::before', props: ['opacity'], golden: 'a click\'s focus keeps no gradient: one click on Play kept it over the picture for the whole film (muse, the owner 2026-11-05: 「似乎太頻繁了，而不是只有滑鼠hover」)' },
      { target: '.cui-console-player-footer', props: ['opacity'], golden: 'nor the controls: only the keyboard\'s focus keeps them (player hidden focused)' },
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
      { target: SCRUB, props: ['box-height'], golden: 'a range 44px tall on a touch screen as on a desktop, its track and thumb the same' },
      { target: SCRUB, props: ['margin-top', 'margin-bottom'], golden: 'on a touch screen its row its whole target: the buttons above are 44px too, and a target reaching up into theirs would leave each short (the touch audit read 39px)' },
      { target: '.cui-console-player-progress', props: ['grid-template-rows'], golden: 'the range\'s row 44px on a touch screen, over the times\' 14px' },
    ],
  },
];
