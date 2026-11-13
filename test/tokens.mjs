import { readdirSync, readFileSync } from 'node:fs';
import { join } from 'node:path';
import { ROOT } from './compile.mjs';
import { ref } from './reference.mjs';
import { check, done } from './check.mjs';

const css = readFileSync(join(ROOT, 'tokens.css'), 'utf8');
const tbTokens = JSON.parse(readFileSync(join(ROOT, 'reference/asc/results/tokens.json'), 'utf8')).tokens;
// @invariant: the first block, :root's: the table a language's block overrides is not this one
const vars = Object.fromEntries([...css.slice(0, css.indexOf('\n}\n')).matchAll(/^\s*(--cui-[a-z0-9-]+):\s*([^;]+);/gm)].map(m => [m[1], m[2].trim()]));

const sides = (name) => {
  const v = vars[name];
  if (v === undefined) throw new Error(`tokens.css has no ${name}`);
  if (!v.startsWith('light-dark(')) return [v, v];
  const inner = v.slice('light-dark('.length, -1);
  let depth = 0;
  for (let i = 0; i < inner.length; i++) {
    if (inner[i] === '(') depth++;
    else if (inner[i] === ')') depth--;
    else if (inner[i] === ',' && depth === 0) return [inner.slice(0, i).trim(), inner.slice(i + 1).trim()];
  }
  throw new Error(`${name}: no two sides in ${v}`);
};
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
// @external: CSS Color 4's matrices (OKLab to linear sRGB)
function oklchToSrgb(l, c, h) {
  const a = c * Math.cos(h * Math.PI / 180), b = c * Math.sin(h * Math.PI / 180);
  const [l_, m_, s_] = [l + 0.3963377774 * a + 0.2158037573 * b, l - 0.1055613458 * a - 0.0638541728 * b, l - 0.0894841775 * a - 1.2914855480 * b].map(v => v ** 3);
  const lin = [4.0767416621 * l_ - 3.3077115913 * m_ + 0.2309699292 * s_, -1.2684380046 * l_ + 2.6097574011 * m_ - 0.3413193965 * s_, -0.0041960863 * l_ - 0.7034186147 * m_ + 1.7076147010 * s_];
  return lin.map(v => Math.min(1, Math.max(0, v))).map(v => (v <= 0.0031308 ? 12.92 * v : 1.055 * v ** (1 / 2.4) - 0.055) * 255);
}
const computedRgb = (value) => {
  const [r, g, b, a] = channels(value);
  return a === 1 ? `rgb(${r}, ${g}, ${b})` : `rgba(${r}, ${g}, ${b}, ${a})`;
};

const W = 1200;
// @invariant: each color's light and dark side: [page, probe, width, property, state?]
// where apple.com shows it, or 'fitted: why', held by the contrast checks
const CONTRACT = {
  '--cui-c-bg': [['home', 'page.body', W, 'background-color'], ['product', 'page.body', W, 'background-color']],
  '--cui-c-bg-alt': [['home', 'hero.tile', W, 'background-color'], ['product', 'section.dark', W, 'background-color']],
  '--cui-c-text': [['home', 'page.body', W, 'color'], ['home', 'tile.dark-headline', W, 'color']],
  '--cui-c-text-2': [['product', 'stat.caption', W, 'color'], 'fitted: apple.com\'s dark #6e6e73 is below 4.5:1'],
  '--cui-c-text-3': [['buy', 'option', W, 'border-top-color'], ['product', 'localnav.explore', W, 'box-shadow']],
  '--cui-c-border': [['buy', 'option', W, 'border-top-color'], ['product', 'localnav.explore', W, 'box-shadow']],
  '--cui-c-accent': [['home', 'button', W, 'background-color'], ['home', 'tile.dark-button', W, 'background-color']],
  '--cui-c-accent-hover': [['home', 'button', W, 'background-color', 'hover'], ['home', 'tile.dark-button', W, 'background-color', 'hover']],
  '--cui-c-accent-active': [['home', 'button', W, 'background-color', 'active'], ['home', 'tile.dark-button', W, 'background-color', 'active']],
  '--cui-c-on-accent': [['home', 'button', W, 'color'], ['home', 'tile.dark-button', W, 'color']],
  '--cui-c-link': [['home', 'button.secondary', W, 'color'], ['home', 'tile.dark-button-secondary', W, 'color']],
  // @external: the neutral in a dark block: a gallery item's (li.media-gallery-item.theme-dark)
  '--cui-c-neutral': [['store', 'button.neutral', W, 'background-color'], ['home', 'gallery.button-neutral', W, 'background-color']],
  '--cui-c-neutral-hover': [['store', 'button.neutral', W, 'background-color', 'hover'], ['home', 'gallery.button-neutral', W, 'background-color', 'hover']],
  '--cui-c-neutral-active': [['store', 'button.neutral', W, 'background-color', 'active'], ['home', 'gallery.button-neutral', W, 'background-color', 'active']],
  '--cui-c-on-neutral': [['store', 'button.neutral', W, 'color'], ['home', 'gallery.button-neutral', W, 'color']],
  // @external: the global nav's: the home page's, light, and the Newsroom's, a dark nav;
  // its menu opened on a phone
  '--cui-c-bar': [['home', 'nav', W, 'background-color'], ['newsroom', 'nav', W, 'background-color']],
  '--cui-c-on-bar': [['home', 'nav.link', W, 'color'], ['newsroom', 'nav.link', W, 'color']],
  '--cui-c-on-bar-hover': [['home', 'nav.link', W, 'color', 'hover'], ['newsroom', 'nav.link', W, 'color', 'hover']],
  '--cui-c-menu': [['home', 'menu.content', 390, 'background-color'], ['newsroom', 'menu.content', 390, 'background-color']],
  '--cui-c-on-menu': [['home', 'menu.link', 390, 'color'], ['newsroom', 'menu.link', 390, 'color']],
  // @external: apple.com's own field (the retail page's form-textbox), and in a dark
  // color scheme
  '--cui-c-field': [['retail', 'field.input', W, 'background-color'], ['retail-dark', 'field.input', W, 'background-color']],
  // @external: a field's states: the gift card's page's (an empty required field the
  // focus left, one disabled, the currency field's), and copies on the retail
  // page in a dark color scheme; the currency's ground read on the gift
  // card's page alone (a field is light on any page)
  '--cui-c-field-error': [['giftcard', 'error', W, 'background-color'], ['retail-dark', 'field.error', W, 'background-color']],
  '--cui-c-field-disabled': [['giftcard', 'field.disabled', 1600, 'background-color'], ['retail-dark', 'field.disabled', W, 'background-color']],
  '--cui-c-field-prefix': [['giftcard', 'prefix.currency', 1600, 'background-color'], ['giftcard', 'prefix.currency', 1600, 'background-color']],
  '--cui-c-border-disabled': [['giftcard', 'field.disabled', 1600, 'border-top-color'], ['retail-dark', 'field.disabled', W, 'border-top-color']],
  '--cui-c-error': [['giftcard', 'error.message', W, 'color'], 'fitted: apple.com\'s #e30000 is below 4.5:1 on black'],
  // @external: the bag's accordion's rule, and a rule between copy blocks on iPhone 18
  // Pro's dark page
  '--cui-c-divider': [['bag', 'accordion.item', W, 'border-top-color'], ['product', 'divider', W, 'border-top-color']],
  // @external: the footer: the home page's, and the dark product page's, light as well
  '--cui-c-on-footer': [['home', 'footer', W, 'color'], ['product', 'footer', W, 'color']],
  '--cui-c-on-footer-link': [['home', 'footer.link', W, 'color'], ['product', 'footer.link', W, 'color']],
  '--cui-c-on-footer-title': [['home', 'footer.title', W, 'color'], ['product', 'footer.title', W, 'color']],
  '--cui-c-footer-separator': [['home', 'footer.section-rule', 390, 'background-color'], ['home', 'footer.section-rule', 390, 'background-color']],
  // @external: the store's overlay (iPhone 18 Pro's buy page, its Show more opened): one
  // value each, read at a wide window and a phone's
  '--cui-c-scrim': [['overlay', 'overlay', W, 'background-color'], ['overlay', 'overlay', 390, 'background-color']],
  '--cui-c-meter': [['appstore-ratings', 'fill', W, 'background-color'], ['appstore-ratings-dark', 'fill', W, 'background-color']],
  '--cui-c-meter-track': [['appstore-ratings', 'track', W, 'background-color'], ['appstore-ratings-dark', 'track', W, 'background-color']],
  // @external: tv.apple.com's video player, paused (decision 28): one value each, a
  // video dark on any page
  '--cui-c-video': [['tv-player', 'dialog', W, 'background-color'], ['tv-player', 'dialog', 1600, 'background-color']],
  '--cui-c-on-video': [['tv-player', 'play', W, 'color'], ['tv-player', 'play', 1600, 'color']],
  '--cui-c-on-video-control': [['tv-player', 'subtitles', W, 'color'], ['tv-player', 'fullscreen', W, 'color']],
  '--cui-c-on-video-time': [['tv-player', 'elapsed', W, 'color'], ['tv-player', 'remaining', W, 'color']],
  '--cui-c-video-track': [['tv-player', 'volume.range', W, '--progress-track-color'], ['tv-player', 'volume.range', 1600, '--progress-track-color']],
};

const colors = Object.keys(vars).filter(k => k.startsWith('--cui-c-') && !k.startsWith('--cui-c-console-'));
check('the contract\'s colors are tokens.css\'s', colors.join(' '), Object.keys(CONTRACT).join(' '));

for (const [name, [light, dark]] of Object.entries(CONTRACT)) {
  const [l, d] = sides(name);
  for (const [side, value, source] of [['light', l, light], ['dark', d, dark]]) {
    if (typeof source === 'string') continue;
    // @external: a ring drawn as a box-shadow reads as its color first
    const want = ref(...source).match(/rgba?\([^)]*\)/)[0];
    check(`${name} ${side} is apple.com's ${source.slice(0, 2).join(':')}${source[4] ? ':' + source[4] : ''} ${source[3]}`, computedRgb(value), want);
  }
}

// @invariant: the console's colors: light the toolkit's, dark derived (decision 18:
// not measured). A side its job's contrast finds short is fitted (decision 21):
// FITTED's value, a step of the same ramp (or the same step at another alpha),
// held by the contrast checks.
const SOURCES = {
  '--cui-c-console-brand-fill-hover': 'component-Button-variant-colorRole-brand-effect-hover-variant-filled-backgroundColor',
  '--cui-c-console-brand-fill-press': 'component-Button-variant-colorRole-brand-effect-press-variant-filled-backgroundColor',
  '--cui-c-console-brand-label-hover': 'component-Button-variant-colorRole-brand-effect-hover-variant-gray-color',
  '--cui-c-console-brand-label-press': 'component-Button-variant-colorRole-brand-effect-press-variant-plain-color',
  '--cui-c-console-on-accent-press': 'component-Button-variant-colorRole-brand-effect-press-variant-filled-color',
  '--cui-c-console-gray-fill-hover': 'component-Button-variant-colorRole-neutral-effect-hover-variant-gray-backgroundColor',
  '--cui-c-console-gray-fill-press': 'component-Button-variant-colorRole-neutral-effect-press-variant-gray-backgroundColor',
  '--cui-c-console-neutral-label-hover': 'component-Button-variant-colorRole-neutral-effect-hover-variant-gray-color',
  '--cui-c-console-neutral-fill-hover': 'component-Button-variant-colorRole-neutral-effect-hover-variant-filled-backgroundColor',
  '--cui-c-console-neutral-on-accent-press': 'component-Button-variant-colorRole-neutral-effect-press-variant-filled-color',
  '--cui-c-console-error-fill-hover': 'component-Button-variant-colorRole-error-effect-hover-variant-filled-backgroundColor',
  '--cui-c-console-error-fill-press': 'component-Button-variant-colorRole-error-effect-press-variant-filled-backgroundColor',
  '--cui-c-console-error-label-hover': 'component-Button-variant-colorRole-error-effect-hover-variant-gray-color',
  '--cui-c-console-error-label-press': 'component-Button-variant-colorRole-error-effect-press-variant-gray-color',
  '--cui-c-console-card': 'color-v2-bgPlaceholder',
  '--cui-c-console-field-border': 'component-TextInput-variant-colorRole-neutral-focus-off-variant-stroked-border-color',
  '--cui-c-console-brand-label-on-tint': 'component-Link-variant-colorRole-brand-effect-default-variant-inline-color',
  '--cui-c-console-dialog-stroke': 'component-Modal-variant-variant-modal-stroke',
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
  // @invariant: a label on a brand or error accent keeps its white, the accent a dark hue on
  // either side; the neutral's accent is black, white in dark, so its label turns with it
  if (name.includes('on-accent') && !name.startsWith('--cui-c-console-neutral-')) return light;
  if (l === 0 && c === 0) return withAlpha('oklch(100.0% 0.000 0)', alpha);
  if (l === 100 && c === 0) return withAlpha('oklch(0.0% 0.000 0)', alpha);
  return withAlpha(tbTokens[stepOf(light, 'light')[0].replace('-palette-light-', '-palette-dark-')], alpha);
};
const FITTED = {
  '--cui-c-console-neutral-label-secondary': { light: 'oklch(35.9% 0.012 285.8 / 0.76)', dark: 'oklch(87.9% 0.015 286.1 / 0.63)', why: 'the toolkit\'s 0.60 is 3.4:1 on white, 0.74 4.4:1 on a chip over a card; its dark 0.60 short of 4.5:1 on a success or warning banner\'s tint over a card (crashcart\'s demo has one)' },
  '--cui-c-console-brand-accent-primary': { dark: 'oklch(51.2% 0.203 260.5)', why: 'white on the derived blue-600 is 2.7:1' },
  '--cui-c-console-brand-fill-hover': { light: 'oklch(49.7% 0.209 260.5)', dark: 'oklch(43.6% 0.171 260.6)', why: 'white on the toolkit\'s lighter blue is 3.9:1' },
  '--cui-c-console-brand-fill-press': { dark: 'oklch(43.6% 0.171 260.6)', why: 'white on the derived blue-600 is 2.7:1' },
  '--cui-c-console-brand-label-hover': { light: 'oklch(49.7% 0.209 260.5)', dark: 'oklch(68.3% 0.167 256.6)', why: 'the toolkit\'s lighter blue is 3.5:1 on white, its derived dark 3.5:1 on black' },
  '--cui-c-console-error-accent-primary': { light: 'oklch(49.5% 0.190 31.5)', dark: 'oklch(50.0% 0.186 31.4)', why: 'white on the toolkit\'s red is 3.3:1' },
  '--cui-c-console-field-border': { light: 'oklch(35.9% 0.012 285.8 / 0.58)', dark: 'oklch(87.9% 0.015 286.1 / 0.42)', why: 'a field\'s edge, at the toolkit\'s 0.18, is 1.4:1 on white; a control\'s edge is held to 3:1' },
  '--cui-c-console-brand-label-on-tint': { light: 'oklch(49.7% 0.209 260.5)', dark: 'oklch(75.6% 0.120 260.5)', why: 'a link in a banner\'s sentence: the toolkit\'s blue 3.9:1 on the info tint over a card, its derived dark blue-600 4.47:1 on the warning tint over a card (crashcart\'s unsymbolicated notice)' },
  '--cui-c-console-info-label-primary': { light: 'oklch(49.7% 0.209 260.5)', dark: 'oklch(68.3% 0.167 256.6)', why: 'the toolkit\'s info blue is 4.0:1 on white; its derived dark blue-500 short of 4.5:1 on its tint over a card' },
  '--cui-c-console-success-label-primary': { light: 'oklch(38.5% 0.094 150.1)', why: 'the toolkit\'s green is 2.2:1 on white, green-800 short of 4.5:1 on the gray fill over a card' },
  '--cui-c-console-warning-label-primary': { light: 'oklch(48.6% 0.112 61.8)', why: 'the toolkit\'s yellow is 1.6:1 on white' },
  '--cui-c-console-error-label-primary': { light: 'oklch(49.5% 0.190 31.5)', dark: 'oklch(70.6% 0.184 29.4)', why: 'the toolkit\'s red is 3.6:1 on white, red-600 4.2:1 on a gray button; its derived dark red-500 short of 4.5:1 on its tint and a gray button over a card (crashcart: 4.09:1, 3.92:1)' },
  '--cui-c-console-error-fill-hover': { light: 'oklch(37.8% 0.147 29.3)', dark: 'oklch(40.9% 0.152 29.5)', why: 'white on the toolkit\'s lighter red is 3.0:1' },
  '--cui-c-console-error-fill-press': { light: 'oklch(37.8% 0.147 29.3)', dark: 'oklch(40.9% 0.152 29.5)', why: 'white on the toolkit\'s red is 3.5:1' },
  '--cui-c-console-error-label-hover': { light: 'oklch(49.5% 0.190 31.5)', dark: 'oklch(70.6% 0.184 29.4)', why: 'the toolkit\'s lighter red is 2.9:1 on white, its derived dark 4.3:1 on a gray button' },
  '--cui-c-console-error-label-press': { light: 'oklch(49.5% 0.190 31.5)', dark: 'oklch(70.6% 0.184 29.4)', why: 'the toolkit\'s red is 3.8:1 on white; its derived dark red-500 short of 4.5:1 on a gray button over a card' },
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

// @external: WCAG 2 contrast. A translucent color is what it draws over the color
// under it.
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
  ['--cui-c-text', ['--cui-c-bg', '--cui-c-bg-alt'], 4.5],
  ['--cui-c-text-2', ['--cui-c-bg', '--cui-c-bg-alt'], 4.5],
  ['--cui-c-link', ['--cui-c-bg', '--cui-c-bg-alt'], 4.5],
  ['--cui-c-text-3', ['--cui-c-bg', '--cui-c-bg-alt'], 3],
  // @invariant: a field's message on the page
  ['--cui-c-error', ['--cui-c-bg', '--cui-c-bg-alt'], 4.5],
  ['--cui-c-border', ['--cui-c-bg', '--cui-c-bg-alt'], 3],
  // @invariant: a share's filled part, a graphic that carries its meaning (WCAG 1.4.11)
  ['--cui-c-meter', ['--cui-c-bg', '--cui-c-bg-alt'], 3],
  // @invariant: the player's text and controls on its video
  ['--cui-c-on-video-control', ['--cui-c-video'], 4.5],
  ['--cui-c-on-video-time', ['--cui-c-video'], 4.5],
];
for (const [fg, grounds, min] of MIN) {
  for (const bg of grounds) {
    for (const i of [0, 1]) {
      const r = ratio(sides(fg)[i], sides(bg)[i]);
      check(`${fg} on ${bg}, ${['light', 'dark'][i]}, at least ${min}:1`, r >= min, true);
    }
  }
}
// @trap: each label on the fills it is drawn on (top first; [] the bare base), each
// over both bases: pairs on a card were missed here and found by a site's
// axe-core run
const PAGE = '--cui-c-console-neutral-background-primary';
const CARD = '--cui-c-console-card';
const BASES = [[PAGE], [CARD, PAGE]];
const ON_FILLS = ['--cui-c-console-brand-accent-primary', '--cui-c-console-brand-fill-hover', '--cui-c-console-error-accent-primary', '--cui-c-console-error-fill-hover'];
const GRAY = ['--cui-c-console-neutral-fill-tertiary', '--cui-c-console-gray-fill-hover', '--cui-c-console-gray-fill-press'];
const ROLES = ['info', 'success', 'warning', 'error'];
// @invariant: a banner's tints: the neutral's and each role's
const TINTS = [['--cui-c-console-neutral-fill-quaternary'], ...ROLES.map((role) => [`--cui-c-console-${role}-fill-tertiary`])];
const CONSOLE_MIN = [
  ['--cui-c-console-neutral-label-primary', [[], ['--cui-c-console-neutral-accent-quaternary'], ...GRAY.map((g) => [g]), ...TINTS], 4.5],
  // @invariant: the secondary label: a current menu item's, a chip's key and its remove
  // control on its gray at rest, hovered and pressed, a banner's caption, a
  // field's placeholder, a card's caption
  ['--cui-c-console-neutral-label-secondary', [[], ...GRAY.map((g) => [g]), ...TINTS], 4.5],
  // @invariant: a button's label on its fill, at rest, hovered and pressed
  ['--cui-c-console-brand-label-on-accent-primary', [...ON_FILLS.map((f) => [f]), ['--cui-c-console-brand-fill-press'], ['--cui-c-console-error-fill-press']], 4.5],
  ['--cui-c-console-on-accent-press', [['--cui-c-console-brand-fill-press'], ['--cui-c-console-error-fill-press']], 4.5],
  ['--cui-c-console-neutral-label-hover', [['--cui-c-console-gray-fill-hover']], 4.5],
  // @invariant: a picked segment's label on its black (white in dark), at rest, hovered and pressed
  ['--cui-c-console-neutral-label-on-accent-primary', [['--cui-c-console-neutral-accent-primary'], ['--cui-c-console-neutral-fill-hover']], 4.5],
  ['--cui-c-console-neutral-on-accent-press', [['--cui-c-console-neutral-accent-primary']], 4.5],
  // @invariant: a badge's label on the base, on its own tint and on the gray fill
  ...ROLES.map((role) => [`--cui-c-console-${role}-label-primary`, [[], [`--cui-c-console-${role}-accent-quaternary`], ['--cui-c-console-neutral-fill-tertiary']], 4.5]),
  // @invariant: a banner's icon on its tint; a card's change up or down
  ...ROLES.map((role) => [`--cui-c-console-${role}-label-primary`, [[`--cui-c-console-${role}-fill-tertiary`]], 4.5]),
  // @invariant: a plain button's, a link's
  ...['--cui-c-console-brand-label-primary', '--cui-c-console-brand-label-hover', '--cui-c-console-brand-label-press'].map((fg) => [fg, [[]], 4.5]),
  // @invariant: a link in a banner's sentence, on the banner's tint
  ['--cui-c-console-brand-label-on-tint', TINTS, 4.5],
  // @invariant: a gray button's red label, on the base and on its gray
  ...['--cui-c-console-error-label-primary', '--cui-c-console-error-label-hover', '--cui-c-console-error-label-press'].map((fg) => [fg, [[], ...GRAY.map((g) => [g])], 4.5]),
  // @invariant: a field's edge and its red edge in error
  ...['--cui-c-console-field-border', '--cui-c-console-error-label-primary'].map((fg) => [fg, [[]], 3]),
];
for (const [fg, fills, min] of CONSOLE_MIN) {
  for (const fill of fills) {
    for (const base of BASES) {
      const layers = [...fill, ...base];
      for (const i of [0, 1]) {
        const ground = over(...layers.map((l) => sides(l)[i]), ['#ffffff', '#000000'][i]);
        check(`${fg} on ${layers.join(' over ')}, ${['light', 'dark'][i]}, at least ${min}:1`, ratio(over(sides(fg)[i], ground), ground) >= min, true);
      }
    }
  }
}
for (const [on, fills] of [['--cui-c-on-accent', ['--cui-c-accent', '--cui-c-accent-hover', '--cui-c-accent-active']], ['--cui-c-on-neutral', ['--cui-c-neutral', '--cui-c-neutral-hover', '--cui-c-neutral-active']],
  ['--cui-c-on-menu', ['--cui-c-menu']], ['--cui-c-on-bar-hover', ['--cui-c-menu', '--cui-c-bg-alt']], ['--cui-c-text-2', ['--cui-c-menu']]]) {
  for (const fill of fills) {
    for (const i of [0, 1]) check(`${on} on ${fill}, ${['light', 'dark'][i]}, at least 4.5:1`, ratio(sides(on)[i], sides(fill)[i]) >= 4.5, true);
  }
}
// @invariant: the nav's text on its bar, the bar over the page's grounds (what scrolls
// under it is the page's)
for (const [on, bar] of [['--cui-c-on-bar', '--cui-c-bar'], ['--cui-c-on-bar-hover', '--cui-c-bar']]) {
  for (const bg of ['--cui-c-bg', '--cui-c-bg-alt']) {
    for (const i of [0, 1]) {
      const ground = over(sides(bar)[i], sides(bg)[i]);
      check(`${on} on ${bar} over ${bg}, ${['light', 'dark'][i]}, at least 4.5:1`, ratio(over(sides(on)[i], ground), ground) >= 4.5, true);
    }
  }
}

// @invariant: a field's text, its value and its label, on its ground over the field's
// own (the page's ground, light: a field keeps its light look); in error,
// its label on the error's ground
for (const fg of ['--cui-c-text', '--cui-c-text-2']) {
  const ground = over(sides('--cui-c-field')[0], sides('--cui-c-bg')[0]);
  check(`${fg} on --cui-c-field over --cui-c-bg, light, at least 4.5:1`, ratio(sides(fg)[0], ground) >= 4.5, true);
}
check('--cui-c-error on --cui-c-field-error, light, at least 4.5:1', ratio(sides('--cui-c-error')[0], sides('--cui-c-field-error')[0]) >= 4.5, true);

// @invariant: the footer's text on its ground, the alternate (light: the footer keeps
// its light look)
for (const fg of ['--cui-c-on-footer', '--cui-c-on-footer-link', '--cui-c-on-footer-title']) {
  const ground = sides('--cui-c-bg-alt')[0];
  check(`${fg} on --cui-c-bg-alt, light, at least 4.5:1`, ratio(over(sides(fg)[0], ground), ground) >= 4.5, true);
}

const table = JSON.parse(readFileSync(join(ROOT, 'reference/tracking/cui-tracking.json'), 'utf8'));
const tracking = Object.keys(vars).filter(k => k.startsWith('--cui-tracking-'));
// @invariant: the table is what was measured; a size no stylesheet reads has no token
const SHEETS = ['theme.css', 'base.css', 'type.css', 'index.css', 'site.css', 'console.css',
  ...readdirSync(join(ROOT, 'components'), { recursive: true }).filter(f => f.endsWith('.css')).map(f => join('components', f))];
const READ = SHEETS.map(f => readFileSync(join(ROOT, f), 'utf8')).join('\n');
const UNCOMMENTED = Object.fromEntries(SHEETS.map(f => [f, readFileSync(join(ROOT, f), 'utf8').replace(/\/\*[\s\S]*?\*\//g, '')]));
check('a letter-spacing for each size of the table a stylesheet reads', tracking.map(k => k.slice(15)).join(' '), Object.keys(table).filter(size => READ.includes(`--cui-tracking-${size})`)).join(' '));
// @invariant: a part's removal leaves no slot behind
check('every token a stylesheet reads', Object.keys(vars).filter(name => !(READ + css).match(new RegExp(`var\\(${name}[,)]`))).join(' '), '');
for (const k of tracking) check(k, parseFloat(vars[k]), table[k.slice(15)]);
const setSlots = (files) => [...new Set(files.flatMap(f => UNCOMMENTED[f].match(/--cui-set-[\w-]+/g)))].sort().join(' ');
const SLOTS_READ = setSlots(['theme.css', 'base.css']);
for (const sheet of ['site.css', 'console.css']) check(`${sheet} declares the --cui-set-* slots theme.css and base.css read, no more`, setSlots([sheet]), SLOTS_READ);

check('--cui-font', vars['--cui-font'], '-apple-system, BlinkMacSystemFont, Inter, "Helvetica Neue", Helvetica, Arial, "PingFang TC", "Noto Sans TC", "Noto Sans CJK TC", "Microsoft JhengHei", sans-serif');
check('--cui-font-mono', vars['--cui-font-mono'], 'ui-monospace, "SF Mono", Menlo, Consolas, "Liberation Mono", monospace');
check('--cui-focus-ring', vars['--cui-focus-ring'], '2px solid var(--cui-c-accent)');
check('the focus ring is apple.com\'s width', ref('home', 'button', W, 'outline-width', 'focus-visible'), '2px');
check('the focus ring is apple.com\'s color', computedRgb(sides('--cui-c-accent')[0]), ref('home', 'button', W, 'outline-color', 'focus-visible'));
check('--cui-focus-ring-offset is apple.com\'s', vars['--cui-focus-ring-offset'], ref('home', 'button', W, 'outline-offset', 'focus-visible'));

check('--cui-console-shadow-sidebar is App Store Connect\'s --tb-component-Sidebar-variant-variant-opaque-boxShadow', vars['--cui-console-shadow-sidebar'], tbTokens['--tb-component-Sidebar-variant-variant-opaque-boxShadow']);
check('--cui-console-shadow-dialog is App Store Connect\'s --tb-component-Modal-variant-variant-modal-boxShadow', vars['--cui-console-shadow-dialog'], tbTokens['--tb-component-Modal-variant-variant-modal-boxShadow']);
check('--cui-console-shadow-popover is App Store Connect\'s --tb-component-Popover-variant-variant-stroked-boxShadow', vars['--cui-console-shadow-popover'], tbTokens['--tb-component-Popover-variant-variant-stroked-boxShadow']);

check('--cui-video-scrim is tv.apple.com\'s tv-player:scrim.gradient background-color', vars['--cui-video-scrim'], ref('tv-player', 'scrim.gradient', W, 'background-color'));
// @trap: golden: muse's subtitles, white under a soft shadow with no box (the owner, 2026-11-05, from Apple's view); tv.apple.com's
// cues are its <video>'s own text track, none read. Two layers: a tight one that edges each letter on a white picture, a wide
// faint one that darkens round the line without a box's edge
check('--cui-video-cue-shadow is muse\'s golden', vars['--cui-video-cue-shadow'], '0 0.05em 0.1em rgb(0 0 0 / 0.6), 0 0 0.5em rgb(0 0 0 / 0.4)');
// @trap: golden: muse's intro and end (the owner, 2026-11-08: 「修播放器還沒有開始播放的預設狀態是暗的」): tv.apple.com's scrim went
// to black at the foot and dimmed the lower half of the cover (2026-11-04; a flat 40% dim over all of it from 2026-11-13). Dark only behind the words in the bottom-left corner, the cue's
// shadow under them
check('--cui-video-title-scrim is muse\'s golden', vars['--cui-video-title-scrim'], 'radial-gradient(ellipse 75% 60% at 0% 100%, rgb(0 0 0 / 0.6), rgb(0 0 0 / 0.35) 40%, rgb(0 0 0 / 0))');
// @trap: golden: muse's scrubber preview (the owner, 2026-11-05: a thumbnail over the time under the pointer); tv.apple.com's
// trailers scrub with no preview. A 1px hairline edges it over a dark picture, a soft drop lifts it off a bright one
check('--cui-video-preview-shadow is muse\'s golden', vars['--cui-video-preview-shadow'], '0 0 0 1px rgb(255 255 255 / 0.15), 0 4px 12px rgb(0 0 0 / 0.4)');
// @trap: golden: muse's full-screen controls (回報問題, 下一個要看), a capsule of words on the picture; tv.apple.com's bar has none.
// A faint white fill that reads on a dark or a bright picture, a little more hovered or pressed
check('--cui-video-control-fill is muse\'s golden', vars['--cui-video-control-fill'], 'rgb(255 255 255 / 0.16)');
check('--cui-video-control-fill-hover is muse\'s golden', vars['--cui-video-control-fill-hover'], 'rgb(255 255 255 / 0.24)');
check('--cui-video-thumb-shadow is tv.apple.com\'s tv-player:volume.range --volume-thumb-box-shadow', vars['--cui-video-thumb-shadow'], ref('tv-player', 'volume.range', W, '--volume-thumb-box-shadow'));

// @external: every named color CSS has; not transparent, currentColor or inherit, which
// are no color of a stylesheet's own
const NAMED = /\b(aliceblue|antiquewhite|aqua|aquamarine|azure|beige|bisque|black|blanchedalmond|blue|blueviolet|brown|burlywood|cadetblue|chartreuse|chocolate|coral|cornflowerblue|cornsilk|crimson|cyan|darkblue|darkcyan|darkgoldenrod|darkgray|darkgreen|darkgrey|darkkhaki|darkmagenta|darkolivegreen|darkorange|darkorchid|darkred|darksalmon|darkseagreen|darkslateblue|darkslategray|darkslategrey|darkturquoise|darkviolet|deeppink|deepskyblue|dimgray|dimgrey|dodgerblue|firebrick|floralwhite|forestgreen|fuchsia|gainsboro|ghostwhite|gold|goldenrod|gray|green|greenyellow|grey|honeydew|hotpink|indianred|indigo|ivory|khaki|lavender|lavenderblush|lawngreen|lemonchiffon|lightblue|lightcoral|lightcyan|lightgoldenrodyellow|lightgray|lightgreen|lightgrey|lightpink|lightsalmon|lightseagreen|lightskyblue|lightslategray|lightslategrey|lightsteelblue|lightyellow|lime|limegreen|linen|magenta|maroon|mediumaquamarine|mediumblue|mediumorchid|mediumpurple|mediumseagreen|mediumslateblue|mediumspringgreen|mediumturquoise|mediumvioletred|midnightblue|mintcream|mistyrose|moccasin|navajowhite|navy|oldlace|olive|olivedrab|orange|orangered|orchid|palegoldenrod|palegreen|paleturquoise|palevioletred|papayawhip|peachpuff|peru|pink|plum|powderblue|purple|rebeccapurple|red|rosybrown|royalblue|saddlebrown|salmon|sandybrown|seagreen|seashell|sienna|silver|skyblue|slateblue|slategray|slategrey|snow|springgreen|steelblue|tan|teal|thistle|tomato|turquoise|violet|wheat|white|whitesmoke|yellow|yellowgreen)\b/i;
for (const file of SHEETS) {
  const text = UNCOMMENTED[file];
  // @trap: a custom property's name is no color (--cui-c-console-gray-fill-hover)
  const values = [...text.matchAll(/:\s*([^;{}]+)[;}]/g)].map(m => m[1]).join(' ; ').replace(/--[\w-]+/g, '');
  const literal = values.match(/#[0-9a-f]{3,8}\b|\b(rgba?|hsla?|hwb|lab|lch|oklab|oklch|color)\(/i) ?? values.match(NAMED);
  check(`${file} writes no color of its own`, literal?.[0] ?? 'none', 'none');
  // @invariant: rule 1: nothing !important, or a utility on the same element
  // no longer wins — the one exception cut out below, which must override every
  // transition a part declares, whatever its layer or specificity
  const still = text.replace(/@media \(prefers-reduced-motion: reduce\) \{[\s\S]*?\n\}/, '');
  check(`${file} writes nothing !important (rule 1)`, /!\s*important/i.test(still) ? 'has !important' : 'none', 'none');
  // @invariant: a console part reads the console's slots alone; the dialog's backdrop the
  // scrim, which the toolkit names none of, and the player the video's
  // (decision 28)
  if (file.startsWith(join('components', 'console'))) {
    const other = [...text.matchAll(/var\((--cui-c-[\w-]+)/g)].map(m => m[1]).filter(name => !name.startsWith('--cui-c-console-') && name !== '--cui-c-scrim' && !/^--cui-c-(on-)?video/.test(name));
    check(`${file} reads the console's colors alone`, other.join(' ') || 'none', 'none');
  }
}

done('tokens');
