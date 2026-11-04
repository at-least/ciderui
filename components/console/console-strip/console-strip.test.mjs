import { examples } from '../../../test/harness.mjs';

const ex = examples(import.meta.url);
const STRIP = '.cui-console-strip';
const GROUP = '.cui-console-strip-group:first-child';
const LABEL = `${GROUP} .cui-console-strip-label`;
const item = (group, n) => `.cui-console-strip-group:nth-child(${group}) > .cui-console-strip-item:nth-of-type(${n})`;
const FIRST = item(1, 1);
const IMG = `${FIRST} > img`;
const CURRENT = '.cui-console-strip-item[aria-current="true"]';
const DOUBT = '.cui-console-strip-mark[data-state="doubt"]';
const KEPT = '.cui-console-strip-mark[data-state="kept"]';
const REDRAW = '.cui-console-strip-mark[data-state="redraw"]';
// tv.apple.com's player's row was read at 1600 and 1200px
const AT = { 390: 1200 };

export const cases = [
  {
    name: 'strip',
    html: ex('strip'),
    widths: [1200, 390],
    checks: [
      { target: STRIP, props: ['display', 'grid-auto-flow', 'overflow-x', 'scroll-snap-type', 'padding-top', 'padding-left', 'margin-top', 'margin-bottom'], apple: ['tv-player-playlist', 'row'], at: AT },
      { target: FIRST, props: ['scroll-snap-align'], apple: ['tv-player-playlist', 'item'], at: AT },
      { target: IMG, props: ['border-top-left-radius', 'border-top-width', 'border-top-style', 'box-shadow', 'cursor'], asc: ['media', 'thumb.img'] },
      { target: IMG, props: ['box-width', 'box-height', 'object-fit'], golden: 'a 16:9 shot 160px wide, cropped to fill: a storyboard\'s frames are a film\'s shape (App Store Connect\'s are a device\'s), more of them in view than tv.apple.com\'s 288px trailers (tv-player-playlist lockup, five at 1600px)' },
      { target: IMG, props: ['background-color'], golden: 'the console card\'s gray while it loads, as a console thumbnail\'s' },
      { target: GROUP, props: ['column-gap'], golden: 'a scene\'s shots 12px apart: App Store Connect\'s 6px either side of each (media.json thumb.item), where tv.apple.com\'s trailers are 20px apart' },
      { target: STRIP, props: ['column-gap', 'overflow-y'], golden: 'one row, scrolled sideways only as tv.apple.com\'s (its grid of columns); a scene\'s shots twice as close to each other as to the next scene\'s, so they read as its group: neither site groups its row' },
      { target: LABEL, props: ['font-size', 'line-height', 'font-weight', 'color', 'text-overflow', 'white-space'], golden: 'the scene\'s name above its shots in a console thumbnail\'s caption\'s type (cui-console-thumb-caption: 11px on 13px, the secondary label), cut to its group\'s width; neither site labels a group of its row' },
      { target: `${GROUP} > .cui-console-strip-item:first-of-type > img`, props: ['box-top', 'box-left'], golden: 'the first picture 18px under its group\'s top, its name\'s 13px line and 5px (a thumbnail\'s caption\'s 5px from it), 4px in from the page\'s content edge, the items\' room for their ring' },
      { target: LABEL, props: ['box-top', 'box-left'], golden: 'the scene\'s name over its first picture\'s left edge' },
      { target: FIRST, props: ['border-top-width', 'border-top-color', 'padding-top', 'border-top-left-radius'], golden: 'a 2px edge 2px out from the picture, clear at rest, so the current shot\'s ring takes no room of its own' },
      { target: CURRENT, props: ['border-top-color', 'border-top-width'], golden: 'the current shot ringed in the brand accent, 2px out from its picture (Keynote\'s and Photos\' chosen thumbnail): neither site marks one current (tv-player-playlist, none marked as playing); a ring, not a fill, so its picture shows whole' },
      { target: DOUBT, props: ['box-width', 'box-height', 'top', 'right', 'border-top-width', 'border-top-color', 'background-color', 'border-top-left-radius'], golden: 'a doubt waiting: a dot in console-status\'s warning mark\'s color (its 12px), ringed 2px in the page\'s ground so it shows on any picture, 6px in from the picture\'s top right' },
      { target: KEPT, props: ['box-width', 'box-height', 'top', 'right', 'color', 'background-color'], golden: 'kept: a check (lucide.Check, 3px of its 24, so 1.5px drawn) in the success label\'s color on the page\'s ground, the dot\'s size and place: a glyph, not a color, tells it from the dot' },
      { target: `${KEPT} > svg`, props: ['box-width', 'box-height'], golden: 'the check 12px, inside the ground\'s 2px ring' },
      { target: REDRAW, props: ['top', 'left', 'box-width', 'box-height', 'border-top-width', 'border-top-color', 'border-top-left-radius'], golden: 'to redraw: a 2px edge in the error accent round the picture\'s inside, its corners the picture\'s, where the current ring is outside it: a shot can be both' },
    ],
  },
  {
    name: 'strip hovered',
    html: ex('strip'),
    widths: [1200],
    checks: [
      { target: IMG, props: ['border-top-left-radius', 'box-shadow'], state: 'hover', on: FIRST, asc: ['media', 'thumb.img'] },
      { target: FIRST, props: ['border-top-color'], state: 'hover', golden: 'unchanged under the pointer, as App Store Connect\'s thumbnail (media: nothing changed with the pointer on one)' },
    ],
  },
  {
    name: 'strip dark',
    html: `<div class="cui-dark">${ex('strip')}</div>`,
    widths: [1200],
    checks: [
      { target: LABEL, props: ['color'], golden: 'the dark side derived (tokens.css, decision 18)' },
      { target: CURRENT, props: ['border-top-color'], golden: 'the dark side derived (tokens.css, decision 18)' },
      { target: DOUBT, props: ['background-color', 'border-top-color'], golden: 'the dark side derived (tokens.css, decision 18), the ring the dark page\'s ground' },
      { target: KEPT, props: ['color', 'background-color'], golden: 'the dark side derived (tokens.css, decision 18)' },
      { target: REDRAW, props: ['border-top-color'], golden: 'the dark side derived (tokens.css, decision 18)' },
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
      { target: DOUBT, props: ['background-color', 'border-top-color'], golden: 'the dot the forced text color on the ground\'s ring, as console-status\'s mark' },
      { target: KEPT, props: ['color', 'background-color'], golden: 'the check the forced text color on the ground' },
      { target: REDRAW, props: ['border-top-color'], golden: 'the redraw edge the forced text color, inside the picture where the current ring is out of it' },
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
];
