// tokens.css against what it is: each color apple.com's measured value
// (reference/apple.com/, the probe named beside it) or the fitted one,
// one light-dark() pair or one value; the contrast minimums of each color's
// job on both sides; the letter-spacing per size the tracking probe's table
// (reference/tracking/cui-tracking.json); the font stack the one the stack
// probe holds (reference/stack/); the focus ring apple.com's; and no other
// stylesheet with a color of its own.

import { readdirSync, readFileSync } from 'node:fs';
import { join } from 'node:path';
import { ROOT } from './compile.mjs';
import { ref } from './reference.mjs';
import { check, done } from './check.mjs';

const css = readFileSync(join(ROOT, 'tokens.css'), 'utf8');
const tbTokens = JSON.parse(readFileSync(join(ROOT, 'reference/asc/results/tokens.json'), 'utf8')).tokens;
const vars = Object.fromEntries([...css.matchAll(/^\s*(--cui-[a-z0-9-]+):\s*([^;]+);/gm)].map(m => [m[1], m[2].trim()]));

// a token's two sides: light-dark(a, b), or one value for both
const sides = (name) => {
  const v = vars[name];
  if (v === undefined) throw new Error(`tokens.css has no ${name}`);
  if (!v.startsWith('light-dark(')) return [v, v];
  // the comma between the two sides is the one outside any parentheses
  const inner = v.slice('light-dark('.length, -1);
  let depth = 0;
  for (let i = 0; i < inner.length; i++) {
    if (inner[i] === '(') depth++;
    else if (inner[i] === ')') depth--;
    else if (inner[i] === ',' && depth === 0) return [inner.slice(0, i).trim(), inner.slice(i + 1).trim()];
  }
  throw new Error(`${name}: no two sides in ${v}`);
};
// a side as its channels, [r, g, b, alpha]: #rrggbb, or rgb(r g b / a)
const channels = (value) => {
  const h = value.match(/^#([0-9a-f]{6})$/i);
  if (h) {
    const n = parseInt(h[1], 16);
    return [n >> 16, (n >> 8) & 255, n & 255, 1];
  }
  const o = value.match(/^oklch\(([\d.]+)% ([\d.]+) ([\d.]+)(?: \/ ([\d.]+))?\)$/);
  if (o) return [...oklchToSrgb(+o[1] / 100, +o[2], +o[3]), o[4] === undefined ? 1 : +o[4]];
  const c = value.match(/^rgb\((\d+) (\d+) (\d+) \/ ([\d.]+)\)$/);
  if (!c) throw new Error(`not a #rrggbb, rgb(r g b / a) or oklch(l% c h / a) color: ${value}`);
  return [+c[1], +c[2], +c[3], +c[4]];
};
// oklch to sRGB's 0-255 channels, unrounded and clipped to the gamut, by
// CSS Color 4's matrices (OKLab to linear sRGB)
function oklchToSrgb(l, c, h) {
  const a = c * Math.cos(h * Math.PI / 180), b = c * Math.sin(h * Math.PI / 180);
  const [l_, m_, s_] = [l + 0.3963377774 * a + 0.2158037573 * b, l - 0.1055613458 * a - 0.0638541728 * b, l - 0.0894841775 * a - 1.2914855480 * b].map(v => v ** 3);
  const lin = [4.0767416621 * l_ - 3.3077115913 * m_ + 0.2309699292 * s_, -1.2684380046 * l_ + 2.6097574011 * m_ - 0.3413193965 * s_, -0.0041960863 * l_ - 0.7034186147 * m_ + 1.7076147010 * s_];
  return lin.map(v => Math.min(1, Math.max(0, v))).map(v => (v <= 0.0031308 ? 12.92 * v : 1.055 * v ** (1 / 2.4) - 0.055) * 255);
}
// a computed rgb()/rgba() as the rgb(r g b / a) channels() reads
const computedToSpace = (value) => {
  const m = value.match(/^rgba?\((\d+), (\d+), (\d+)(?:, ([\d.]+))?\)$/);
  if (!m) throw new Error(`not a computed rgb color: ${value}`);
  return `rgb(${m[1]} ${m[2]} ${m[3]} / ${m[4] ?? 1})`;
};
// as the browser computes it
const rgb = (value) => {
  const [r, g, b, a] = channels(value);
  return a === 1 ? `rgb(${r}, ${g}, ${b})` : `rgba(${r}, ${g}, ${b}, ${a})`;
};

// the contract: each color's light and dark side and where apple.com shows it
// ([page, probe, width, property, state?, opacity?]: with `opacity`, the color
// at the element's opacity, as apple.com draws a title in black at .88), or
// `fitted` with why
const W = 1200;
const CONTRACT = {
  '--cui-c-bg': [['home', 'page.body', W, 'background-color'], ['product', 'page.body', W, 'background-color']],
  '--cui-c-bg-alt': [['home', 'hero.tile', W, 'background-color'], ['product', 'section.dark', W, 'background-color']],
  '--cui-c-text': [['home', 'page.body', W, 'color'], ['home', 'tile.dark-headline', W, 'color']],
  '--cui-c-text-2': [['home', 'hero.callout', W, 'color'], 'fitted: apple.com\'s dark #6e6e73 is below 4.5:1'],
  '--cui-c-text-3': [['buy', 'option', W, 'border-top-color'], ['product', 'localnav.explore', W, 'box-shadow']],
  '--cui-c-border': [['buy', 'option', W, 'border-top-color'], ['product', 'localnav.explore', W, 'box-shadow']],
  '--cui-c-accent': [['home', 'button', W, 'background-color'], ['home', 'tile.dark-button', W, 'background-color']],
  '--cui-c-accent-hover': [['home', 'button', W, 'background-color', 'hover'], ['home', 'tile.dark-button', W, 'background-color', 'hover']],
  '--cui-c-accent-active': [['home', 'button', W, 'background-color', 'active'], ['home', 'tile.dark-button', W, 'background-color', 'active']],
  '--cui-c-on-accent': [['home', 'button', W, 'color'], ['home', 'tile.dark-button', W, 'color']],
  '--cui-c-link': [['home', 'button.secondary', W, 'color'], ['home', 'tile.dark-button-secondary', W, 'color']],
  // the neutral in a dark block: a gallery item's (li.media-gallery-item.theme-dark)
  '--cui-c-neutral': [['store', 'button.neutral', W, 'background-color'], ['home', 'gallery.button-neutral', W, 'background-color']],
  '--cui-c-neutral-hover': [['store', 'button.neutral', W, 'background-color', 'hover'], ['home', 'gallery.button-neutral', W, 'background-color', 'hover']],
  '--cui-c-neutral-active': [['store', 'button.neutral', W, 'background-color', 'active'], ['home', 'gallery.button-neutral', W, 'background-color', 'active']],
  '--cui-c-on-neutral': [['store', 'button.neutral', W, 'color'], ['home', 'gallery.button-neutral', W, 'color']],
  // the global nav's: the home page's, light, and the Newsroom's, a dark nav;
  // its menu opened on a phone
  '--cui-c-bar': [['home', 'nav', W, 'background-color'], ['newsroom', 'nav', W, 'background-color']],
  '--cui-c-on-bar': [['home', 'nav.link', W, 'color'], ['newsroom', 'nav.link', W, 'color']],
  '--cui-c-on-bar-hover': [['home', 'nav.link', W, 'color', 'hover'], ['newsroom', 'nav.link', W, 'color', 'hover']],
  '--cui-c-menu': [['home', 'menu.content', 390, 'background-color'], ['newsroom', 'menu.content', 390, 'background-color']],
  '--cui-c-on-menu': [['home', 'menu.link', 390, 'color'], ['newsroom', 'menu.link', 390, 'color']],
  // the search opened: the home page's, and the Newsroom's dark nav's
  '--cui-c-curtain': [['search', 'curtain', W, 'background-color'], ['search-dark', 'curtain', 1600, 'background-color']],
  // the local nav's: MacBook Air's, light, and MacBook Pro's, dark
  '--cui-c-bar-2': [['localnav', 'localnav.background', W, 'background-color'], ['localnav-dark', 'localnav.background', W, 'background-color']],
  '--cui-c-on-bar-2': [['localnav', 'localnav.title', W, 'color', undefined, 'opacity'], ['localnav-dark', 'localnav.title', W, 'color', undefined, 'opacity']],
  '--cui-c-separator': [['localnav', 'localnav.rule', W, 'background-color'], ['localnav-dark', 'localnav.rule', W, 'background-color']],
  // the home page's tiles: the hero's, light, and a dark one (.tile-wrapper.theme-dark)
  '--cui-c-tile': [['home', 'hero.tile', W, 'background-color'], ['home', 'tile.dark', W, 'background-color']],
  // the store's cards: a light one, a dark one (rf-ccard-darkbg); a dark
  // card's plain eyebrow on a copy the probe adds
  '--cui-c-on-card': [['store', 'card.header', W, 'color'], ['store', 'card.dark-header', W, 'color']],
  '--cui-c-eyebrow': [['store', 'card.eyebrow', W, 'color'], ['store', 'card.dark-eyebrow', W, 'color']],
  '--cui-c-tag': [['store', 'card.tag', W, 'color'], ['store', 'card.dark-tag', W, 'color']],
  // a gallery's: the iPhone page's paddles and dots, light, and iPhone 18
  // Pro's, dark (its first paddle disabled, the next read)
  '--cui-c-paddle': [['iphone', 'paddle', W, 'background-color'], ['product', 'paddle.next', W, 'background-color']],
  '--cui-c-paddle-hover': [['iphone', 'paddle', W, 'background-color', 'hover'], ['product', 'paddle.next', W, 'background-color', 'hover']],
  '--cui-c-paddle-active': [['iphone', 'paddle', W, 'background-color', 'active'], ['product', 'paddle.next', W, 'background-color', 'active']],
  '--cui-c-on-paddle': [['iphone', 'paddle', W, 'color'], ['product', 'paddle.next', W, 'color']],
  '--cui-c-on-paddle-hover': [['iphone', 'paddle', W, 'color', 'hover'], ['product', 'paddle.next', W, 'color', 'hover']],
  '--cui-c-dot': [['iphone', 'gallery.dot', W, 'background-color'], ['product', 'gallery.dot', W, 'background-color']],
  '--cui-c-dot-hover': [['iphone', 'gallery.dot', W, 'background-color', 'hover'], ['product', 'gallery.dot', W, 'background-color', 'hover']],
  // apple.com's own field (the retail page's form-textbox), and in a dark
  // color scheme
  '--cui-c-field': [['retail', 'field.input', W, 'background-color'], ['retail-dark', 'field.input', W, 'background-color']],
  // a field's states: the gift card's page's (an empty required field the
  // focus left, one disabled, the currency field's), and copies on the retail
  // page in a dark color scheme; the currency's ground read on the gift
  // card's page alone (a field is light on any page)
  '--cui-c-field-error': [['giftcard', 'error', W, 'background-color'], ['retail-dark', 'field.error', W, 'background-color']],
  '--cui-c-field-disabled': [['giftcard', 'field.disabled', 1600, 'background-color'], ['retail-dark', 'field.disabled', W, 'background-color']],
  '--cui-c-field-prefix': [['giftcard', 'prefix.currency', 1600, 'background-color'], ['giftcard', 'prefix.currency', 1600, 'background-color']],
  '--cui-c-border-disabled': [['giftcard', 'field.disabled', 1600, 'border-top-color'], ['retail-dark', 'field.disabled', W, 'border-top-color']],
  '--cui-c-error': [['giftcard', 'error.message', W, 'color'], 'fitted: apple.com\'s #e30000 is below 4.5:1 on black'],
  // the bag's accordion's rule, and a rule between copy blocks on iPhone 18
  // Pro's dark page
  '--cui-c-divider': [['bag', 'accordion.item', W, 'border-top-color'], ['product', 'divider', W, 'border-top-color']],
  // the footer: the home page's, and the dark product page's, light as well
  '--cui-c-on-footer': [['home', 'footer', W, 'color'], ['product', 'footer', W, 'color']],
  '--cui-c-on-footer-link': [['home', 'footer.link', W, 'color'], ['product', 'footer.link', W, 'color']],
  '--cui-c-on-footer-title': [['home', 'footer.title', W, 'color'], ['product', 'footer.title', W, 'color']],
  '--cui-c-footer-rule': [['home', 'footer.legal-rule', W, 'background-color'], ['home', 'footer.legal-rule', 390, 'background-color']],
  // the store's overlay (iPhone 18 Pro's buy page, its Show more opened): one
  // value each, read at a wide window and a phone's, a modal light on any page
  '--cui-c-scrim': [['overlay', 'overlay', W, 'background-color'], ['overlay', 'overlay', 390, 'background-color']],
  '--cui-c-close': [['overlay', 'close.circle', W, 'background-color'], ['overlay', 'close.circle', 390, 'background-color']],
  '--cui-c-close-hover': [['overlay', 'close.circle-hover', W, 'background-color', 'hover'], ['overlay', 'close.circle-hover', 390, 'background-color', 'hover']],
  '--cui-c-close-active': [['overlay', 'close.circle-hover', W, 'background-color', 'active'], ['overlay', 'close.circle-hover', 390, 'background-color', 'active']],
  '--cui-c-on-close': [['overlay', 'close.circle', W, 'color'], ['overlay', 'close.circle-hover', 390, 'color', 'active']],
  '--cui-c-on-close-hover': [['overlay', 'close.circle-hover', W, 'color', 'hover'], ['overlay', 'close.circle-hover', 390, 'color', 'hover']],
};

const colors = Object.keys(vars).filter(k => k.startsWith('--cui-c-') && !k.startsWith('--cui-c-console-'));
check('the contract\'s colors are tokens.css\'s', colors.join(' '), Object.keys(CONTRACT).join(' '));

for (const [name, [light, dark]] of Object.entries(CONTRACT)) {
  const [l, d] = sides(name);
  for (const [side, value, source] of [['light', l, light], ['dark', d, dark]]) {
    if (typeof source === 'string') continue; // fitted: held by the contrast checks
    // a ring drawn as a box-shadow reads as its color first
    let want = ref(...source.slice(0, 5)).match(/rgba?\([^)]*\)/)[0];
    if (source[5]) {
      // the color at the element's opacity
      const [r, g, b, a] = channels(computedToSpace(want));
      const alpha = +(a * Number(ref(source[0], source[1], source[2], source[5], source[4]))).toFixed(4);
      want = alpha === 1 ? `rgb(${r}, ${g}, ${b})` : `rgba(${r}, ${g}, ${b}, ${alpha})`;
    }
    check(`${name} ${side} is apple.com's ${source.slice(0, 2).join(':')}${source[4] ? ':' + source[4] : ''} ${source[3]}${source[5] ? ' at its ' + source[5] : ''}`, rgb(value), want);
  }
}

// the console's colors (--cui-c-console-*, the console parts'): each App
// Store Connect's toolkit's color (reference/asc/results/tokens.json) on
// the light side, the role of the slot's name (--cui-c-console-brand-label-
// on-accent-primary is --tb-color-role-brand-label-onAccent-primary) or, a
// part's own state color, the token SOURCES names; on the dark side, which
// the toolkit does not ship, that color's step of its dark palette at its
// alpha: the light palette's step of the same lightness, chroma and hue, or,
// for a color that is no step, the step nearest its lightness in the ramp
// nearest its hue (the gray ramps for a gray, chroma under 0.05); black and
// white each other's, and white on an accent (on-accent) white still
// (DESIGN.md, decision 18: derived, not measured). A side its job's
// contrast finds short is fitted (decision 21): FITTED's value, a step of
// the same ramp (or the same step at another alpha), held below by the
// contrast checks.
const SOURCES = {
  '--cui-c-console-brand-fill-hover': 'component-Button-variant-colorRole-brand-effect-hover-variant-filled-backgroundColor',
  '--cui-c-console-brand-fill-press': 'component-Button-variant-colorRole-brand-effect-press-variant-filled-backgroundColor',
  '--cui-c-console-brand-label-hover': 'component-Button-variant-colorRole-brand-effect-hover-variant-gray-color',
  '--cui-c-console-brand-label-press': 'component-Button-variant-colorRole-brand-effect-press-variant-plain-color',
  '--cui-c-console-on-accent-press': 'component-Button-variant-colorRole-brand-effect-press-variant-filled-color',
  '--cui-c-console-gray-fill-hover': 'component-Button-variant-colorRole-neutral-effect-hover-variant-gray-backgroundColor',
  '--cui-c-console-gray-fill-press': 'component-Button-variant-colorRole-neutral-effect-press-variant-gray-backgroundColor',
  '--cui-c-console-neutral-label-hover': 'component-Button-variant-colorRole-neutral-effect-hover-variant-gray-color',
  '--cui-c-console-error-fill-hover': 'component-Button-variant-colorRole-error-effect-hover-variant-filled-backgroundColor',
  '--cui-c-console-error-fill-press': 'component-Button-variant-colorRole-error-effect-press-variant-filled-backgroundColor',
  '--cui-c-console-error-label-hover': 'component-Button-variant-colorRole-error-effect-hover-variant-gray-color',
  '--cui-c-console-error-label-press': 'component-Button-variant-colorRole-error-effect-press-variant-gray-color',
  '--cui-c-console-card': 'color-v2-bgPlaceholder',
  '--cui-c-console-field-border': 'component-TextInput-variant-colorRole-neutral-focus-off-variant-stroked-border-color',
  '--cui-c-console-brand-label-on-fill': 'component-Tabs-variant-colorRole-brand-variant-filledGray-anatomy-Tab-variant-effect-default-isActive-false-color',
};
const sourceOf = (name) => `--tb-${SOURCES[name] ?? `color-role-${name.slice('--cui-c-console-'.length).replace('on-accent', 'onAccent')}`}`;
const lch = (v) => {
  const m = v.match(/^oklch\(([\d.]+)% ([\d.]+) ([\d.]+)(?: \/ ([\d.]+))?\)$/);
  if (!m) throw new Error(`not an oklch() color: ${v}`);
  return { l: +m[1], c: +m[2], h: +m[3], alpha: m[4], opaque: `oklch(${m[1]}% ${m[2]} ${m[3]})` };
};
const steps = (side) => Object.entries(tbTokens).filter(([k]) => k.startsWith(`--tb-color-palette-${side}-`));
const sameColor = (a, b) => a.replace(/ 0\.0\)$/, ' 0)') === b.replace(/ 0\.0\)$/, ' 0)');
const rampName = (k) => k.match(/-palette-(?:light|dark)-([a-zA-Z]+)-/)[1];
// a color's step of a palette side: [token, ramp]
const stepOf = (value, side) => {
  const { l, c, h, opaque } = lch(value);
  const exact = steps(side).filter(([, v]) => sameColor(v, opaque));
  if (exact.length) {
    const ramps = new Set(exact.map(([k]) => rampName(k)));
    if (ramps.size !== 1) throw new Error(`${opaque} is a step of more than one ramp: ${[...ramps].join(', ')}`);
    return exact[0];
  }
  const turn = (a, b) => Math.min(Math.abs(a - b), 360 - Math.abs(a - b));
  const at500 = (r) => lch(tbTokens[`--tb-color-palette-${side}-${r}-500`]);
  const gray = c < 0.05;
  const ramps = [...new Set(steps(side).map(([k]) => rampName(k)))].filter((r) => (at500(r).c < 0.05) === gray);
  const ramp = ramps.reduce((a, b) => (turn(at500(b).h, h) < turn(at500(a).h, h) ? b : a));
  return steps(side).filter(([k]) => rampName(k) === ramp).reduce((a, b) => (Math.abs(lch(b[1]).l - l) < Math.abs(lch(a[1]).l - l) ? b : a));
};
const withAlpha = (opaque, alpha) => (alpha === undefined ? opaque : opaque.replace(/\)$/, ` / ${alpha})`));
const darkOf = (name, light) => {
  const { l, c, alpha } = lch(light);
  if (name.includes('on-accent')) return light;
  if (l === 0 && c === 0) return withAlpha('oklch(100.0% 0.000 0)', alpha);
  if (l === 100 && c === 0) return withAlpha('oklch(0.0% 0.000 0)', alpha);
  return withAlpha(tbTokens[stepOf(light, 'light')[0].replace('-palette-light-', '-palette-dark-')], alpha);
};
const FITTED = {
  '--cui-c-console-neutral-label-secondary': { light: 'oklch(35.9% 0.012 285.8 / 0.74)', why: 'the toolkit\'s 0.60 is 3.4:1 on white' },
  '--cui-c-console-brand-accent-primary': { dark: 'oklch(51.2% 0.203 260.5)', why: 'white on the derived blue-600 is 2.7:1' },
  '--cui-c-console-brand-fill-hover': { light: 'oklch(49.7% 0.209 260.5)', dark: 'oklch(43.6% 0.171 260.6)', why: 'white on the toolkit\'s lighter blue is 3.9:1' },
  '--cui-c-console-brand-fill-press': { dark: 'oklch(43.6% 0.171 260.6)', why: 'white on the derived blue-600 is 2.7:1' },
  '--cui-c-console-brand-label-hover': { light: 'oklch(49.7% 0.209 260.5)', dark: 'oklch(62.4% 0.206 255.5)', why: 'the toolkit\'s lighter blue is 3.5:1 on white, its derived dark 3.5:1 on black' },
  '--cui-c-console-error-accent-primary': { light: 'oklch(49.5% 0.190 31.5)', dark: 'oklch(50.0% 0.186 31.4)', why: 'white on the toolkit\'s red is 3.3:1' },
  '--cui-c-console-field-border': { light: 'oklch(35.9% 0.012 285.8 / 0.58)', dark: 'oklch(87.9% 0.015 286.1 / 0.42)', why: 'a field\'s edge, at the toolkit\'s 0.18, is 1.4:1 on white; a control\'s edge is held to 3:1' },
  '--cui-c-console-brand-label-on-fill': { light: 'oklch(49.7% 0.209 260.5)', why: 'the brand\'s label on the gray fill is 4.4:1' },
  '--cui-c-console-info-label-primary': { light: 'oklch(49.7% 0.209 260.5)', why: 'the toolkit\'s info blue is 4.0:1 on white' },
  '--cui-c-console-success-label-primary': { light: 'oklch(50.4% 0.125 152.9)', why: 'the toolkit\'s green is 2.2:1 on white' },
  '--cui-c-console-warning-label-primary': { light: 'oklch(48.6% 0.112 61.8)', why: 'the toolkit\'s yellow is 1.6:1 on white' },
  '--cui-c-console-error-label-primary': { light: 'oklch(49.5% 0.190 31.5)', why: 'the toolkit\'s red is 3.6:1 on white, red-600 4.2:1 on a gray button' },
  '--cui-c-console-error-fill-hover': { light: 'oklch(37.8% 0.147 29.3)', dark: 'oklch(40.9% 0.152 29.5)', why: 'white on the toolkit\'s lighter red is 3.0:1' },
  '--cui-c-console-error-fill-press': { light: 'oklch(37.8% 0.147 29.3)', dark: 'oklch(40.9% 0.152 29.5)', why: 'white on the toolkit\'s red is 3.5:1' },
  '--cui-c-console-error-label-hover': { light: 'oklch(49.5% 0.190 31.5)', dark: 'oklch(70.6% 0.184 29.4)', why: 'the toolkit\'s lighter red is 2.9:1 on white, its derived dark 4.3:1 on a gray button' },
  '--cui-c-console-error-label-press': { light: 'oklch(49.5% 0.190 31.5)', why: 'the toolkit\'s red is 3.8:1 on white' },
};
for (const name of Object.keys(vars).filter(k => k.startsWith('--cui-c-console-'))) {
  const source = sourceOf(name);
  if (!(source in tbTokens)) throw new Error(`${name}: the toolkit has no ${source}`);
  const [l, d] = sides(name);
  const fit = FITTED[name] ?? {};
  const derived = darkOf(name, tbTokens[source]);
  if (fit.light) {
    check(`${name} light, fitted (${fit.why}), is a step of ${source}'s ramp`, rampName(stepOf(fit.light, 'light')[0]), rampName(stepOf(tbTokens[source], 'light')[0]));
    check(`${name} light is its fitted value`, l, fit.light);
  } else check(`${name} light is App Store Connect's ${source}`, l, tbTokens[source]);
  if (fit.dark) {
    check(`${name} dark, fitted (${fit.why}), is a step of its derived step's ramp`, rampName(stepOf(fit.dark, 'dark')[0]), rampName(stepOf(derived, 'dark')[0]));
    check(`${name} dark is its fitted value`, d, fit.dark);
  } else check(`${name} dark is ${source}'s step of the dark palette`, d, derived);
}
for (const name of Object.keys(FITTED)) if (!(name in vars)) throw new Error(`FITTED names ${name}, which tokens.css has not`);

// contrast: WCAG 2, each color's job on both sides. A translucent color is
// what it draws over the color under it: `over` lays each side in turn on
// the next (the nav's text on its bar on the page)
const over = (...layers) => layers.reduceRight((under, top) => {
  const [r, g, b, a] = channels(top);
  const [ur, ug, ub] = channels(under);
  const mix = (t, u) => Math.round(t * a + u * (1 - a));
  return `#${[mix(r, ur), mix(g, ug), mix(b, ub)].map(v => v.toString(16).padStart(2, '0')).join('')}`;
});
const lum = (value) => channels(value).slice(0, 3).map(v => v / 255).map(v => v <= 0.04045 ? v / 12.92 : ((v + 0.055) / 1.055) ** 2.4)
  .reduce((s, v, i) => s + v * [0.2126, 0.7152, 0.0722][i], 0);
const ratio = (a, b) => { const [x, y] = [lum(a), lum(b)].sort((p, q) => q - p); return (x + 0.05) / (y + 0.05); };
const MIN = [
  ['--cui-c-text', ['--cui-c-bg', '--cui-c-bg-alt', '--cui-c-tile'], 4.5],
  ['--cui-c-text-2', ['--cui-c-bg', '--cui-c-bg-alt', '--cui-c-tile'], 4.5],
  ['--cui-c-link', ['--cui-c-tile'], 4.5],
  ['--cui-c-link', ['--cui-c-bg', '--cui-c-bg-alt'], 4.5],
  ['--cui-c-text-3', ['--cui-c-bg', '--cui-c-bg-alt'], 3],
  // a field's message on the page
  ['--cui-c-error', ['--cui-c-bg', '--cui-c-bg-alt'], 4.5],
  ['--cui-c-border', ['--cui-c-bg', '--cui-c-bg-alt'], 3],
  // a card's text on its ground, the page's
  ['--cui-c-on-card', ['--cui-c-bg'], 4.5],
  ['--cui-c-eyebrow', ['--cui-c-bg'], 4.5],
  ['--cui-c-tag', ['--cui-c-bg'], 4.5],
];
for (const [fg, grounds, min] of MIN) {
  for (const bg of grounds) {
    for (const i of [0, 1]) {
      const r = ratio(sides(fg)[i], sides(bg)[i]);
      check(`${fg} on ${bg}, ${['light', 'dark'][i]}, at least ${min}:1`, r >= min, true);
    }
  }
}
// a console part's text on the console's grounds: each ground the layers
// named, over the page (white light, black dark)
const PAGE = '--cui-c-console-neutral-background-primary';
const ON_FILLS = ['--cui-c-console-brand-accent-primary', '--cui-c-console-brand-fill-hover', '--cui-c-console-error-accent-primary', '--cui-c-console-error-fill-hover'];
const GRAY = ['--cui-c-console-neutral-fill-tertiary', '--cui-c-console-gray-fill-hover', '--cui-c-console-gray-fill-press'];
const CONSOLE_MIN = [
  ['--cui-c-console-neutral-label-primary', [[PAGE], ['--cui-c-console-neutral-fill-quaternary', PAGE], ...GRAY.map((g) => [g, PAGE])], 4.5],
  // the secondary label on a fill too: a current menu item's
  ['--cui-c-console-neutral-label-secondary', [[PAGE], ['--cui-c-console-neutral-fill-quaternary', PAGE]], 4.5],
  // a button's label on its fill, at rest, hovered and pressed
  ['--cui-c-console-brand-label-on-accent-primary', [...ON_FILLS.map((f) => [f]), ['--cui-c-console-brand-fill-press'], ['--cui-c-console-error-fill-press']], 4.5],
  ['--cui-c-console-on-accent-press', [['--cui-c-console-brand-fill-press'], ['--cui-c-console-error-fill-press']], 4.5],
  ['--cui-c-console-neutral-label-hover', [['--cui-c-console-gray-fill-hover', PAGE]], 4.5],
  // a badge's label on the page, on its own tint and on the gray fill
  ...['brand', 'info', 'success', 'warning', 'error'].map((role) => [role === 'brand' ? '--cui-c-console-brand-label-on-fill' : `--cui-c-console-${role}-label-primary`, [[PAGE], [`--cui-c-console-${role}-accent-quaternary`, PAGE], ['--cui-c-console-neutral-fill-tertiary', PAGE]], 4.5]),
  ['--cui-c-console-neutral-label-primary', [['--cui-c-console-neutral-accent-quaternary', PAGE]], 4.5],
  // a banner's text, its caption and its icon on its tint
  ...['--cui-c-console-neutral-label-primary', '--cui-c-console-neutral-label-secondary'].map((fg) => [fg, ['info', 'success', 'warning', 'error'].map((role) => [`--cui-c-console-${role}-fill-tertiary`, PAGE]), 4.5]),
  ...['info', 'success', 'warning', 'error'].map((role) => [`--cui-c-console-${role}-label-primary`, [[`--cui-c-console-${role}-fill-tertiary`, PAGE]], 4.5]),
  // a card's text on its ground: its title, value, caption, a change up or down
  ...['--cui-c-console-neutral-label-primary', '--cui-c-console-neutral-label-secondary', '--cui-c-console-success-label-primary', '--cui-c-console-error-label-primary'].map((fg) => [fg, [['--cui-c-console-card', PAGE]], 4.5]),
  // a gray segment's blue label, at rest and hovered, on its gray
  ...['--cui-c-console-brand-label-on-fill', '--cui-c-console-brand-label-hover'].map((fg) => [fg, GRAY.map((g) => [g, PAGE]), 4.5]),
  // a plain button's, a link's: on the page
  ...['--cui-c-console-brand-label-primary', '--cui-c-console-brand-label-hover', '--cui-c-console-brand-label-press'].map((fg) => [fg, [[PAGE]], 4.5]),
  // a gray button's red label: on the page and on its gray
  ...['--cui-c-console-error-label-primary', '--cui-c-console-error-label-hover', '--cui-c-console-error-label-press'].map((fg) => [fg, [[PAGE], ...GRAY.map((g) => [g, PAGE])], 4.5]),
];
CONSOLE_MIN.push(
  // a field's edge and its red edge in error: 3:1 on the page and on a card
  ...['--cui-c-console-field-border', '--cui-c-console-error-label-primary'].map((fg) => [fg, [[PAGE], ['--cui-c-console-card', PAGE]], 3]),
  // its text, its placeholder (the secondary label) on its white
  ['--cui-c-console-neutral-label-secondary', [[PAGE], ['--cui-c-console-card', PAGE]], 4.5],
);
for (const [fg, grounds, min] of CONSOLE_MIN) {
  for (const layers of grounds) {
    for (const i of [0, 1]) {
      const ground = over(...layers.map((l) => sides(l)[i]), ['#ffffff', '#000000'][i]);
      check(`${fg} on ${layers.join(' over ')}, ${['light', 'dark'][i]}, at least ${min}:1`, ratio(over(sides(fg)[i], ground), ground) >= min, true);
    }
  }
}
for (const [on, fills] of [['--cui-c-on-accent', ['--cui-c-accent', '--cui-c-accent-hover', '--cui-c-accent-active']], ['--cui-c-on-neutral', ['--cui-c-neutral', '--cui-c-neutral-hover', '--cui-c-neutral-active']],
  ['--cui-c-on-menu', ['--cui-c-menu']], ['--cui-c-on-bar-hover', ['--cui-c-menu', '--cui-c-bg-alt']], ['--cui-c-text-2', ['--cui-c-menu']]]) {
  for (const fill of fills) {
    for (const i of [0, 1]) check(`${on} on ${fill}, ${['light', 'dark'][i]}, at least 4.5:1`, ratio(sides(on)[i], sides(fill)[i]) >= 4.5, true);
  }
}
// the nav's text on its bar, the bar over the page's grounds (what scrolls
// under it is the page's)
for (const [on, bar] of [['--cui-c-on-bar', '--cui-c-bar'], ['--cui-c-on-bar-hover', '--cui-c-bar'], ['--cui-c-on-bar', '--cui-c-bar-2'], ['--cui-c-on-bar-2', '--cui-c-bar-2']]) {
  for (const bg of ['--cui-c-bg', '--cui-c-bg-alt']) {
    for (const i of [0, 1]) {
      const ground = over(sides(bar)[i], sides(bg)[i]);
      check(`${on} on ${bar} over ${bg}, ${['light', 'dark'][i]}, at least 4.5:1`, ratio(over(sides(on)[i], ground), ground) >= 4.5, true);
    }
  }
}

// a field's text, its value and its label, on its ground over the field's
// own (the page's ground, light: a field keeps its light look); in error,
// its label on the error's ground
for (const fg of ['--cui-c-text', '--cui-c-text-2']) {
  const ground = over(sides('--cui-c-field')[0], sides('--cui-c-bg')[0]);
  check(`${fg} on --cui-c-field over --cui-c-bg, light, at least 4.5:1`, ratio(sides(fg)[0], ground) >= 4.5, true);
}
check('--cui-c-error on --cui-c-field-error, light, at least 4.5:1', ratio(sides('--cui-c-error')[0], sides('--cui-c-field-error')[0]) >= 4.5, true);

// the footer's text on its ground, the alternate (light: the footer keeps
// its light look)
for (const fg of ['--cui-c-on-footer', '--cui-c-on-footer-link', '--cui-c-on-footer-title']) {
  const ground = sides('--cui-c-bg-alt')[0];
  check(`${fg} on --cui-c-bg-alt, light, at least 4.5:1`, ratio(over(sides(fg)[0], ground), ground) >= 4.5, true);
}

// a gallery's marks, 3:1 (WCAG 2's for what is not text): a paddle's chevron
// on its paddle, the paddle over the page's grounds; a dot on them
for (const [mark, layers] of [['--cui-c-on-paddle', ['--cui-c-paddle']], ['--cui-c-on-paddle-hover', ['--cui-c-paddle-hover']], ['--cui-c-on-paddle-hover', ['--cui-c-paddle-active']], ['--cui-c-dot', []], ['--cui-c-dot-hover', []]]) {
  for (const bg of ['--cui-c-bg', '--cui-c-bg-alt']) {
    for (const i of [0, 1]) {
      const ground = over(...layers.map(l => sides(l)[i]), sides(bg)[i]);
      check(`${mark}${layers.length ? ' on ' + layers[0] : ''} over ${bg}, ${['light', 'dark'][i]}, at least 3:1`, ratio(over(sides(mark)[i], ground), ground) >= 3, true);
    }
  }
}

// a modal's close button: its X on its circle, at rest, hovered and pressed,
// 3:1 (a mark, not text); the circle opaque, on the modal's light ground
for (const [mark, fill] of [['--cui-c-on-close', '--cui-c-close'], ['--cui-c-on-close-hover', '--cui-c-close-hover'], ['--cui-c-on-close', '--cui-c-close-active']]) {
  const ground = sides(fill)[0];
  check(`${mark} on ${fill}, at least 3:1`, ratio(over(sides(mark)[0], ground), ground) >= 3, true);
}

// the letter-spacing per size
const table = JSON.parse(readFileSync(join(ROOT, 'reference/tracking/cui-tracking.json'), 'utf8'));
const tracking = Object.keys(vars).filter(k => k.startsWith('--cui-tracking-'));
check('a letter-spacing for each size of the table', tracking.map(k => k.slice(15)).join(' '), Object.keys(table).join(' '));
for (const [size, em] of Object.entries(table)) check(`--cui-tracking-${size}`, parseFloat(vars[`--cui-tracking-${size}`]), em);

// the font stack and the focus ring
check('--cui-font', vars['--cui-font'], '-apple-system, BlinkMacSystemFont, Inter, "Helvetica Neue", Helvetica, Arial, sans-serif');
check('--cui-font-mono', vars['--cui-font-mono'], 'ui-monospace, "SF Mono", Menlo, Consolas, "Liberation Mono", monospace');
check('--cui-focus-ring', vars['--cui-focus-ring'], '2px solid var(--cui-c-accent)');
check('the focus ring is apple.com\'s width', ref('home', 'button', W, 'outline-width', 'focus-visible'), '2px');
check('the focus ring is apple.com\'s color', rgb(sides('--cui-c-accent')[0]), ref('home', 'button', W, 'outline-color', 'focus-visible'));
check('--cui-focus-ring-offset is apple.com\'s', vars['--cui-focus-ring-offset'], ref('home', 'button', W, 'outline-offset', 'focus-visible'));

// a console sheet's shadow: the toolkit's opaque sidebar's
check('--cui-console-shadow-sidebar is App Store Connect\'s --tb-component-Sidebar-variant-variant-opaque-boxShadow', vars['--cui-console-shadow-sidebar'], tbTokens['--tb-component-Sidebar-variant-variant-opaque-boxShadow']);

// the card's shadows, as the browser computes one: its color first, then its
// lengths (the spread 0px where none is written)
const shadow = (value) => {
  const color = value.match(/rgba?\([^)]*\)/)[0];
  const lengths = value.replace(color, '').trim().split(/\s+/);
  while (lengths.length < 4) lengths.push('0px');
  return `${color.includes(',') ? color : rgb(color)} ${lengths.join(' ')}`;
};
for (const [name, state] of [['--cui-shadow-card'], ['--cui-shadow-card-hover', 'hover']]) {
  for (const probe of ['card', 'card.dark']) {
    check(`${name} is apple.com's store:${probe}${state ? ':' + state : ''} box-shadow`, shadow(vars[name]), shadow(ref('store', probe, W, 'box-shadow', state)));
  }
}

// no stylesheet but tokens.css writes a color: not by value, not by name
// (transparent, currentColor and inherit are not a color of its own)
// every named color CSS has
const NAMED = /\b(aliceblue|antiquewhite|aqua|aquamarine|azure|beige|bisque|black|blanchedalmond|blue|blueviolet|brown|burlywood|cadetblue|chartreuse|chocolate|coral|cornflowerblue|cornsilk|crimson|cyan|darkblue|darkcyan|darkgoldenrod|darkgray|darkgreen|darkgrey|darkkhaki|darkmagenta|darkolivegreen|darkorange|darkorchid|darkred|darksalmon|darkseagreen|darkslateblue|darkslategray|darkslategrey|darkturquoise|darkviolet|deeppink|deepskyblue|dimgray|dimgrey|dodgerblue|firebrick|floralwhite|forestgreen|fuchsia|gainsboro|ghostwhite|gold|goldenrod|gray|green|greenyellow|grey|honeydew|hotpink|indianred|indigo|ivory|khaki|lavender|lavenderblush|lawngreen|lemonchiffon|lightblue|lightcoral|lightcyan|lightgoldenrodyellow|lightgray|lightgreen|lightgrey|lightpink|lightsalmon|lightseagreen|lightskyblue|lightslategray|lightslategrey|lightsteelblue|lightyellow|lime|limegreen|linen|magenta|maroon|mediumaquamarine|mediumblue|mediumorchid|mediumpurple|mediumseagreen|mediumslateblue|mediumspringgreen|mediumturquoise|mediumvioletred|midnightblue|mintcream|mistyrose|moccasin|navajowhite|navy|oldlace|olive|olivedrab|orange|orangered|orchid|palegoldenrod|palegreen|paleturquoise|palevioletred|papayawhip|peachpuff|peru|pink|plum|powderblue|purple|rebeccapurple|red|rosybrown|royalblue|saddlebrown|salmon|sandybrown|seagreen|seashell|sienna|silver|skyblue|slateblue|slategray|slategrey|snow|springgreen|steelblue|tan|teal|thistle|tomato|turquoise|violet|wheat|white|whitesmoke|yellow|yellowgreen)\b/i;
const sheets = ['theme.css', 'base.css', 'type.css', 'index.css',
  ...readdirSync(join(ROOT, 'components'), { recursive: true }).filter(f => f.endsWith('.css')).map(f => join('components', f))];
for (const file of sheets) {
  const text = readFileSync(join(ROOT, file), 'utf8').replace(/\/\*[\s\S]*?\*\//g, '');
  // a value is what follows a colon up to its semicolon or brace
  // a custom property's name is no color (--cui-c-console-gray-fill-hover)
  const values = [...text.matchAll(/:\s*([^;{}]+)[;}]/g)].map(m => m[1]).join(' ; ').replace(/--[\w-]+/g, '');
  const literal = values.match(/#[0-9a-f]{3,8}\b|\b(rgba?|hsla?|hwb|lab|lch|oklab|oklch|color)\(/i) ?? values.match(NAMED);
  check(`${file} writes no color of its own`, literal?.[0] ?? 'none', 'none');
}

done('tokens');
