// What a component's test file (components/<name>/<name>.test.mjs) builds
// its cases from: the property lists the checks share, and the component's
// examples (components/<name>/examples/<id>.html), the markup the crate's
// tests render and the cases measure.
//
// A case:
//   name     what it is; CASES=<regex> picks cases by it
//   html     ciderui's markup (an example, or markup written here)
//   widths   the viewport widths it is read at (default: the reference's,
//            1600, 1200, 900 and 390px); height, the window's (800px);
//            touch, reducedMotion as the page's
//   click    a selector (or several) clicked in turn once, at the first
//            width, before its checks (a menu's button), then the keys in
//            `press` pressed in turn, its motion then run to its end, or for
//            `wait` ms (0 reads its start); `wait` alone waits so long once
//            the page shows (a motion that runs as it loads); scroll,
//            [selector, x]: that element scrolled to x at every width
//            before its checks (a scroller, snapping where it snaps);
//            hover, a selector or a point ([x, y]), or several: the pointer
//            moved onto each in turn in steps at every width, its motion run
//            to its end (a flyout the pointer shows)
// A check:
//   target   a selector in the case's markup; props, the values read (a
//            property, the box: box-height, box-width, box-top, box-left, box-right,
//            box-x, box-y, `focused` or `hit`, test/browser.mjs), of its `pseudo`
//            element (::after) if it names one; state, the pseudo-classes
//            forced on it (or on `on`)
//   apple    [page, probe, state?]: each value equals apple.com's measured
//            one (reference/apple.com/results/<page>.json) at the same
//            width, or at `at[width]` for a width the reference was not read
//            at; never a box's width or a `width` (apple.com's text was SF,
//            ciderui's here is Inter); `as` maps a property to apple.com's
//            that does the same (an icon's margin-left to its glyph's
//            padding-left)
//   golden   why it is ciderui's: the value equals golden.json's, which
//            UPDATE_GOLDEN=1 rewrites
//   rest     a touch screen's hovered or pressed element equals itself at
//            rest (rule 5)

import { readFileSync } from 'node:fs';

export const SIDES = ['top', 'right', 'bottom', 'left'];
export const PADDING = SIDES.map(s => `padding-${s}`);
export const BORDER = ['border-top-width', 'border-top-style', 'border-bottom-width', 'border-bottom-style'];
export const BORDER_COLOR = ['border-top-color', 'border-bottom-color'];
export const COLORS = ['color', 'background-color', ...BORDER_COLOR];
export const SIZE = ['font-size', 'line-height', 'font-weight'];
export const RADIUS = ['border-top-left-radius'];
export const TEXT = ['display', 'text-align', 'white-space', 'text-decoration-line', 'cursor', 'vertical-align', 'opacity'];
export const OUTLINE = ['outline-style', 'outline-color', 'outline-width', 'outline-offset'];

// a component's examples, read by its test file's URL (`const ex =
// examples(import.meta.url)`): ex(id) is examples/<id>.html's markup
export const examples = (url) => (id) => readFileSync(new URL(`examples/${id}.html`, url), 'utf8').trim();
