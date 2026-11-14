import { examples } from '../../../test/harness.mjs';

const ex = examples(import.meta.url);
const STRIP = '.cui-console-strip';
const GROUP = '.cui-console-strip-group:first-child';
const HEADER = `${GROUP} > .cui-console-strip-header`;
const TITLE = `${HEADER} > .cui-console-strip-title`;
const item = (group, n) => `.cui-console-strip-group:nth-child(${group}) > .cui-console-strip-item:nth-of-type(${n})`;
const FIRST = item(1, 1);
const IMG = `${FIRST} > img`;
const CAPTION = `${FIRST} > .cui-console-strip-caption`;
const CURRENT = '.cui-console-strip-item[aria-current="true"]';
const WAITING = '.cui-console-strip-mark[data-state="waiting"]';
const WORKING = '.cui-console-strip-mark[data-state="working"]';
const CURRENT_WORKING = `${CURRENT} > ${WORKING}`;
const TIP = `${FIRST} > .cui-console-strip-tip`;
const MARKS = 'two marks only (the owner, 2026-11-05, for muse\'s storyboard and assets strip): who the shot waits on, the reader or someone at work on it; nothing, no mark';
// @invariant: one scene of one shot, its name longer than the shot is wide, and no caption
const lone = (title) => `<div class="cui-console-page"><nav class="cui-console-strip" aria-label="Shots"><div class="cui-console-strip-group"><p class="cui-console-strip-header"><span class="cui-console-strip-title">${title}</span></p><a class="cui-console-strip-item" href="#a"><img src="data:image/gif;base64,R0lGODlhAQABAAAAACw=" alt="Shot 1"></a></div></nav></div>`;
// @external: tv.apple.com's player's row was read at 1600 and 1200px
const AT = { 390: 1200 };

export const cases = [
  {
    name: 'strip',
    html: ex('strip'),
    widths: [1200, 390],
    checks: [
      { target: TIP, props: ['font-size', 'line-height'], golden: 'the small type on a phone too, where the console page\'s own is larger' },
      { target: STRIP, props: ['display', 'grid-auto-flow', 'overflow-x', 'scroll-snap-type', 'padding-top', 'padding-left', 'margin-top', 'margin-bottom'], apple: ['tv-player-playlist', 'row'], at: AT },
      { target: FIRST, props: ['scroll-snap-align'], apple: ['tv-player-playlist', 'item'], at: AT },
      { target: IMG, props: ['border-top-left-radius', 'border-top-width', 'border-top-style', 'box-shadow', 'cursor'], asc: ['media', 'thumb.img'] },
      { target: IMG, props: ['box-width', 'box-height', 'object-fit'], golden: 'a 16:9 shot 160px wide, cropped to fill: a storyboard\'s frames are a film\'s shape (App Store Connect\'s are a device\'s), more of them in view than tv.apple.com\'s 288px trailers (tv-player-playlist lockup, five at 1600px)' },
      { target: IMG, props: ['background-color'], golden: 'the console card\'s gray while it loads, as a console thumbnail\'s' },
      { target: GROUP, props: ['column-gap'], golden: 'a scene\'s shots 12px apart: App Store Connect\'s 6px either side of each (media.json thumb.item), where tv.apple.com\'s trailers are 20px apart' },
      { target: STRIP, props: ['column-gap', 'overflow-y'], golden: 'one row, scrolled sideways only as tv.apple.com\'s (its grid of columns); a scene\'s shots twice as close to each other as to the next scene\'s, so they read as its group: neither site groups its row' },
      { target: TITLE, props: ['font-size', 'line-height', 'font-weight', 'color', 'text-overflow', 'white-space'], golden: 'the scene\'s name above its shots in 11px on 13px, the secondary label, cut to its group\'s width; neither site labels a group of its row' },
      { target: `${GROUP} > .cui-console-strip-item:first-of-type > img`, props: ['box-top', 'box-left'], golden: 'the first picture 18px under its group\'s top, its name\'s 13px line and 5px (a thumbnail\'s caption\'s 5px from it), 4px in from the page\'s content edge, the items\' room for their ring' },
      { target: TITLE, props: ['box-top', 'box-left', 'padding-left', 'position', 'left'], golden: 'the scene\'s name over its first picture\'s left edge, its text 4px into its box as the picture is into its item; sticky at the scroller\'s edge, so the name of the scene scrolled half out stays in view (muse\'s storyboard, the owner, 2026-11-05)' },
      { target: HEADER, props: ['position', 'box-top', 'box-left', 'box-width', 'overflow-x'], golden: 'the name\'s track: out of the flow so a long name widens no scene, the group\'s whole width (528px: three 168px shots, each its picture and 4px of ring room either side, 12px apart), not clipped (a clipping box would be the sticky name\'s scroller)' },
      { target: CAPTION, props: ['font-size', 'line-height', 'font-weight', 'color', 'text-overflow', 'white-space'], golden: 'the shot\'s code under it in 11px on 13px, the secondary label, cut to the picture\'s width: muse\'s storyboard names each shot under it, not over it (the owner, 2026-11-05)' },
      { target: CAPTION, props: ['box-top', 'box-left', 'box-width'], golden: 'the caption 5px under its picture (the picture\'s foot at 108px: 18 and 90), as a console thumbnail\'s caption, outside the item\'s ring, at the picture\'s left and as wide' },
      { target: GROUP, props: ['padding-bottom'], golden: 'a scene with captions keeps their room under its shots: 1px past the ring and the 13px line, as its name\'s over them' },
      { target: CURRENT_WORKING, props: ['box-width', 'top', 'right', 'background-color'], golden: 'the current shot\'s dot drawn as any shot\'s, inside its ring' },
      { target: FIRST, props: ['border-top-width', 'border-top-color', 'padding-top', 'border-top-left-radius'], golden: 'a 2px edge 2px out from the picture, clear at rest, so the current shot\'s ring takes no room of its own' },
      { target: CURRENT, props: ['border-top-color', 'border-top-width'], golden: 'the current shot ringed in the brand accent, 2px out from its picture (Keynote\'s and Photos\' chosen thumbnail): neither site marks one current (tv-player-playlist, none marked as playing); a ring, not a fill, so its picture shows whole' },
      { target: WAITING, props: ['box-width', 'box-height', 'top', 'right', 'border-top-width', 'border-top-color', 'background-color', 'border-top-left-radius'], golden: `waiting for the reader: a dot in console-status's warning mark's color (its 12px), ringed 2px in the page's ground so it shows on any picture, 8px in from the item (6px from the picture). ${MARKS}` },
      { target: WORKING, props: ['box-width', 'box-height', 'top', 'right', 'border-top-width', 'border-top-color', 'background-color'], golden: `someone at work on it: the same dot in console-status's neutral mark's color (the secondary label). ${MARKS}` },
      { target: TIP, props: ['visibility', 'opacity', 'position'], golden: 'its line hidden until the pointer is on it' },
      { target: CURRENT, pseudo: '::after', props: ['content'], golden: 'the current shot draws no progress line, so its picture shows whole: the player\'s scrubber says how far it has played (muse\'s owner, 2026-11-07)' },
    ],
  },
  {
    name: 'strip hovered',
    html: ex('strip'),
    widths: [1200],
    checks: [
      { target: TIP, on: FIRST, state: 'focus-visible', props: ['visibility', 'opacity'], golden: 'the keyboard reaches what the pointer shows' },
      { target: IMG, props: ['border-top-left-radius', 'box-shadow'], state: 'hover', on: FIRST, asc: ['media', 'thumb.img'] },
      { target: FIRST, props: ['border-top-color'], state: 'hover', golden: 'unchanged under the pointer, as App Store Connect\'s thumbnail (media: nothing changed with the pointer on one)' },
    ],
  },
  {
    name: 'strip tip',
    html: `<div style="padding-top:120px">${ex('strip')}</div>`,
    widths: [1200],
    hover: FIRST,
    checks: [
      { target: TIP, props: ['visibility', 'opacity', 'position', 'font-size', 'line-height', 'color', 'background-color', 'box-shadow', 'padding-top', 'padding-left', 'border-top-left-radius', 'max-width'], golden: 'the shot\'s first spoken line over it while the pointer is on it (the owner, 2026-11-05, muse\'s storyboard; its assets strip, the asset\'s description), after half a second so a pointer passing along the row raises none: the console popover\'s ground, shadow and 13px on 18px (decision 27, its floating surface), 6px by 10px, its narrow 14px corners, 280px at most; neither site has a hover tip on its row' },
      { target: TIP, props: ['box-left', 'box-top', 'box-width', 'box-height'], golden: 'centred over the item, 8px above it (the popover\'s gap), past the scroller\'s clip (fixed to the item by anchor positioning); at the window\'s edge it keeps inside' },
      { target: FIRST, props: ['box-top', 'box-left', 'box-width'], golden: 'the item it is over' },
      { target: TIP, props: ['transition-property', 'transition-duration', 'transition-delay'], golden: 'half a second\'s wait, then shown at once: the delay alone is the device (a pointer passing along the row raises nothing), a fade on top of it an entrance flourish decision 25 leaves out; leaving hides it at once' },
    ],
  },
  {
    name: 'strip dark',
    html: `<div class="cui-dark">${ex('strip')}</div>`,
    widths: [1200],
    checks: [
      { target: TITLE, props: ['color'], golden: 'the dark side derived (tokens.css, decision 18)' },
      { target: CAPTION, props: ['color'], golden: 'the dark side derived (tokens.css, decision 18)' },
      { target: CURRENT, props: ['border-top-color'], golden: 'the dark side derived (tokens.css, decision 18)' },
      { target: WAITING, props: ['background-color', 'border-top-color'], golden: 'the dark side derived (tokens.css, decision 18), the ring the dark page\'s ground' },
      { target: WORKING, props: ['background-color'], golden: 'the dark side derived (tokens.css, decision 18)' },
    ],
  },
  {
    name: 'strip forced colors',
    html: ex('strip'),
    widths: [1200],
    forcedColors: true,
    checks: [
      { target: CURRENT, props: ['border-top-color'], golden: 'the current shot\'s ring the theme\'s Highlight, as a current tab\'s' },
      { target: FIRST, props: ['border-top-color'], golden: 'the others\' clear edge the ground\'s color: forced colors drew it in the text\'s, every shot ringed as the current one' },
      { target: WAITING, props: ['background-color', 'border-top-color'], golden: 'the waiting dot the forced text color on the ground\'s ring, as console-status\'s mark' },
      { target: WORKING, props: ['background-color', 'border-top-color'], golden: 'the working dot the theme\'s GrayText, so the two keep apart where their colors are forced' },
    ],
  },
  {
    name: 'strip snap',
    html: ex('strip'),
    widths: [1200, 390],
    scroll: [STRIP, 300],
    checks: [
      { target: item(1, 3), props: ['box-left'], golden: 'between two shots the row comes to rest with one at its first shot\'s place (scroll-snap-type x mandatory and each shot\'s start, tv.apple.com\'s row)' },
    ],
  },
  {
    name: 'strip end',
    html: ex('strip'),
    widths: [1200, 390],
    scroll: [STRIP, 99999],
    checks: [
      { target: item(3, 3), props: ['box-right'], golden: 'at its end the last shot as far in from the right as the first from the left' },
    ],
  },
  {
    name: 'strip pinned',
    html: ex('strip'),
    widths: [1200, 390],
    scroll: [STRIP, 400],
    checks: [
      { target: STRIP, props: ['box-left'], golden: 'the scroller 15px out of the content edge, its room for the rings (tv.apple.com\'s row)' },
      { target: GROUP, props: ['box-left'], golden: 'scrolled to the first scene\'s third shot (it snaps there from 400px, 2 × 180px): the scene\'s start 360px out of view, its end still in it' },
      { target: TITLE, props: ['box-left'], golden: 'the scene\'s name stays while its scene is still in view (muse\'s storyboard, the owner, 2026-11-05), at the content edge, 15px in from the scroller\'s (Chromium sticks it inside the scroller\'s padding, its room for the rings), so its text is where the first picture starts at rest; left in the group, it would have scrolled out with the scene\'s start' },
      { target: `.cui-console-strip-group:nth-child(2) > .cui-console-strip-header > .cui-console-strip-title`, props: ['box-left'], golden: 'the next scene\'s name over its first shot, where its scene starts (24px past the first scene\'s end)' },
    ],
  },
  {
    name: 'strip long title',
    html: lone('A scene whose name is longer than its one shot is wide'),
    widths: [1200],
    checks: [
      { target: TITLE, props: ['box-width', 'text-overflow'], golden: 'a name longer than its scene cut to the scene\'s width (one shot, 168px), widening nothing' },
      { target: '.cui-console-strip-group', props: ['box-width', 'padding-bottom'], golden: 'the scene its shot\'s width, whatever its name; no caption, so no room under it' },
    ],
  },
  {
    name: 'strip assets',
    html: ex('assets'),
    widths: [1200, 390],
    checks: [
      { target: `${FIRST} > .cui-console-strip-detail`, props: ['font-size', 'line-height', 'color', 'text-overflow', 'white-space', 'box-top', 'box-left', 'box-width'], golden: 'a second line under the caption, in its type, the next line down (126px: the caption\'s 113 and its 13px line), cut to the picture\'s width: muse\'s assets page says under each asset\'s name how many shots use it (the owner, 2026-11-05)' },
      { target: GROUP, props: ['padding-bottom'], golden: 'a strip with a detail line keeps both lines\' room under every group\'s shots: 1px past the ring and two 13px lines' },
      { target: `${FIRST} > img`, props: ['object-position', 'object-fit', 'box-width', 'box-height'], golden: 'a character sheet cropped to its left edge, where its face is (muse\'s assets page), at the row\'s 16:9 and not muse\'s square: a square would narrow its item, its caption and its marks\' places in a row of 16:9 shots' },
      { target: `${item(2, 1)} > img`, props: ['object-position'], golden: 'any other picture cropped round its middle' },
      { target: `${item(1, 2)} > img`, props: ['opacity'], golden: 'an asset no shot uses: its picture dimmed to half, muse\'s; not its lines, which its words already say (dimmed, the caption fell to 2:1 on the page, the surfaces\' contrast check)' },
      { target: `${item(1, 2)} > .cui-console-strip-caption`, props: ['opacity'], golden: 'its caption not dimmed' },
      { target: `${FIRST} > img`, props: ['opacity'], golden: 'one in use not dimmed' },
    ],
  },
  {
    name: 'strip assets hovered',
    html: ex('assets'),
    widths: [1200],
    checks: [
      { target: `${item(1, 2)} > img`, props: ['opacity'], state: 'hover', on: item(1, 2), golden: 'a dimmed asset under the pointer at 0.8, muse\'s: brighter, so it reads as a link, and still apart from those in use' },
    ],
  },
];
