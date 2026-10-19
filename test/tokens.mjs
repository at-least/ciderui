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
  const c = value.match(/^rgb\((\d+) (\d+) (\d+) \/ ([\d.]+)\)$/);
  if (!c) throw new Error(`not a #rrggbb or rgb(r g b / a) color: ${value}`);
  return [+c[1], +c[2], +c[3], +c[4]];
};
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
  // the bag's accordion's rule, and a rule between copy blocks on iPhone 18
  // Pro's dark page
  '--cui-c-divider': [['bag', 'accordion.item', W, 'border-top-color'], ['product', 'divider', W, 'border-top-color']],
  // the footer: the home page's, and the dark product page's, light as well
  '--cui-c-on-footer': [['home', 'footer', W, 'color'], ['product', 'footer', W, 'color']],
  '--cui-c-on-footer-link': [['home', 'footer.link', W, 'color'], ['product', 'footer.link', W, 'color']],
  '--cui-c-on-footer-title': [['home', 'footer.title', W, 'color'], ['product', 'footer.title', W, 'color']],
  '--cui-c-footer-rule': [['home', 'footer.legal-rule', W, 'background-color'], ['home', 'footer.legal-rule', 390, 'background-color']],
};

const colors = Object.keys(vars).filter(k => k.startsWith('--cui-c-'));
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
for (const [on, fills] of [['--cui-c-on-accent', ['--cui-c-accent', '--cui-c-accent-hover', '--cui-c-accent-active']], ['--cui-c-on-neutral', ['--cui-c-neutral', '--cui-c-neutral-hover', '--cui-c-neutral-active']],
  ['--cui-c-on-menu', ['--cui-c-menu']], ['--cui-c-on-bar-hover', ['--cui-c-menu']]]) {
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
// own (the page's ground, light: a field keeps its light look)
for (const fg of ['--cui-c-text', '--cui-c-text-2']) {
  const ground = over(sides('--cui-c-field')[0], sides('--cui-c-bg')[0]);
  check(`${fg} on --cui-c-field over --cui-c-bg, light, at least 4.5:1`, ratio(sides(fg)[0], ground) >= 4.5, true);
}

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
const NAMED = /\b(white|black|red|green|blue|gray|grey|silver|yellow|orange|purple|pink|navy|teal|aqua|fuchsia|lime|maroon|olive)\b/i;
const sheets = ['theme.css', 'base.css', 'type.css', 'index.css',
  ...readdirSync(join(ROOT, 'components'), { recursive: true }).filter(f => f.endsWith('.css')).map(f => join('components', f))];
for (const file of sheets) {
  const text = readFileSync(join(ROOT, file), 'utf8').replace(/\/\*[\s\S]*?\*\//g, '');
  // a value is what follows a colon up to its semicolon or brace
  const values = [...text.matchAll(/:\s*([^;{}]+)[;}]/g)].map(m => m[1]).join(' ; ');
  const literal = values.match(/#[0-9a-f]{3,8}\b|rgba?\(|hsla?\(|oklch\(/i) ?? values.match(NAMED);
  check(`${file} writes no color of its own`, literal?.[0] ?? 'none', 'none');
}

done('tokens');
