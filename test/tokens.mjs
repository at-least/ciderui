// tokens.css against what it is: each color apple.com's measured value
// (reference/apple.com/, the probe named beside it) or the fitted one,
// one light-dark() pair or one value; the contrast minimums of each color's
// job on both sides; the letter-spacing per size the tracking probe's table
// (reference/tracking/cui-tracking.json); the font stack the one the stack
// probe holds (reference/stack/); the focus ring apple.com's; and no other
// stylesheet with a color of its own.

import { readFileSync } from 'node:fs';
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
  const m = v.match(/^light-dark\(\s*([^,]+?)\s*,\s*([^)]+?)\s*\)$/);
  return m ? [m[1], m[2]] : [v, v];
};
const rgb = (hex) => {
  const h = hex.match(/^#([0-9a-f]{6})$/i);
  if (!h) throw new Error(`not a #rrggbb color: ${hex}`);
  const n = parseInt(h[1], 16);
  return `rgb(${n >> 16}, ${(n >> 8) & 255}, ${n & 255})`;
};

// the contract: each color's light and dark side and where apple.com shows it
// ([page, probe, width, property, state?]), or `fitted` with why
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
  '--cui-c-accent-active': [['home', 'button', W, 'background-color', 'active'], ['home', 'button', W, 'background-color', 'active']],
  '--cui-c-link': [['home', 'button.secondary', W, 'color'], ['home', 'tile.dark-button-secondary', W, 'color']],
};

const colors = Object.keys(vars).filter(k => k.startsWith('--cui-c-'));
check('the contract\'s colors are tokens.css\'s', colors.join(' '), Object.keys(CONTRACT).join(' '));

for (const [name, [light, dark]] of Object.entries(CONTRACT)) {
  const [l, d] = sides(name);
  for (const [side, value, source] of [['light', l, light], ['dark', d, dark]]) {
    if (typeof source === 'string') continue; // fitted: held by the contrast checks
    // a ring drawn as a box-shadow reads as its color first
    const want = ref(...source).match(/rgba?\([^)]*\)/)[0];
    check(`${name} ${side} is apple.com's ${source.slice(0, 2).join(':')}${source[4] ? ':' + source[4] : ''} ${source[3]}`, rgb(value), want);
  }
}

// contrast: WCAG 2, each color's job on both sides
const lum = (hex) => {
  const n = parseInt(hex.slice(1), 16);
  return [n >> 16, (n >> 8) & 255, n & 255].map(v => v / 255).map(v => v <= 0.04045 ? v / 12.92 : ((v + 0.055) / 1.055) ** 2.4)
    .reduce((s, v, i) => s + v * [0.2126, 0.7152, 0.0722][i], 0);
};
const ratio = (a, b) => { const [x, y] = [lum(a), lum(b)].sort((p, q) => q - p); return (x + 0.05) / (y + 0.05); };
const MIN = [
  ['--cui-c-text', ['--cui-c-bg', '--cui-c-bg-alt'], 4.5],
  ['--cui-c-text-2', ['--cui-c-bg', '--cui-c-bg-alt'], 4.5],
  ['--cui-c-link', ['--cui-c-bg', '--cui-c-bg-alt'], 4.5],
  ['--cui-c-text-3', ['--cui-c-bg', '--cui-c-bg-alt'], 3],
  ['--cui-c-border', ['--cui-c-bg', '--cui-c-bg-alt'], 3],
];
for (const [fg, grounds, min] of MIN) {
  for (const bg of grounds) {
    for (const i of [0, 1]) {
      const r = ratio(sides(fg)[i], sides(bg)[i]);
      check(`${fg} on ${bg}, ${['light', 'dark'][i]}, at least ${min}:1`, r >= min, true);
    }
  }
}
for (const fill of ['--cui-c-accent', '--cui-c-accent-hover', '--cui-c-accent-active']) {
  check(`white on ${fill} at least 4.5:1`, ratio('#ffffff', sides(fill)[0]) >= 4.5, true);
}

// the letter-spacing per size
const table = JSON.parse(readFileSync(join(ROOT, 'reference/tracking/cui-tracking.json'), 'utf8'));
const tracking = Object.keys(vars).filter(k => k.startsWith('--cui-tracking-'));
check('a letter-spacing for each size of the table', tracking.map(k => k.slice(15)).join(' '), Object.keys(table).join(' '));
for (const [size, em] of Object.entries(table)) check(`--cui-tracking-${size}`, parseFloat(vars[`--cui-tracking-${size}`]), em);

// the font stack and the focus ring
check('--cui-font', vars['--cui-font'], '-apple-system, BlinkMacSystemFont, Inter, "Helvetica Neue", Helvetica, Arial, sans-serif');
check('--cui-focus-ring', vars['--cui-focus-ring'], '2px solid var(--cui-c-accent)');
check('the focus ring is apple.com\'s width', ref('home', 'button', W, 'outline-width', 'focus-visible'), '2px');
check('the focus ring is apple.com\'s color', rgb(sides('--cui-c-accent')[0]), ref('home', 'button', W, 'outline-color', 'focus-visible'));
check('--cui-focus-ring-offset is apple.com\'s', vars['--cui-focus-ring-offset'], ref('home', 'button', W, 'outline-offset', 'focus-visible'));
check('--cui-focus-ring-offset-field is apple.com\'s', vars['--cui-focus-ring-offset-field'], ref('signin', 'field.input', W, 'outline-offset', 'focus-visible'));

// no stylesheet but tokens.css writes a color
for (const file of ['theme.css', 'base.css', 'type.css', 'index.css']) {
  const text = readFileSync(join(ROOT, file), 'utf8').replace(/\/\*[\s\S]*?\*\//g, '');
  const literal = text.match(/#[0-9a-f]{3,8}\b|rgba?\(|hsla?\(|oklch\(/i);
  check(`${file} writes no color of its own`, literal?.[0] ?? 'none', 'none');
}

done('tokens');
