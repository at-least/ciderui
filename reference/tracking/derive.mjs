// From the Safari and Chrome runs: how far apart the two browsers are, and the
// letter-spacing system-ui needs per size to land on apple.com's widths
// (apple.com's value minus the measured Δ, averaged over both browsers and both weights).
import { readFileSync, writeFileSync } from 'node:fs';
import { SCALE, SAMPLE } from './analyze.mjs';
const runs = ['safari', 'chrome'].map(b => JSON.parse(readFileSync(new URL(`./results/${b}-measure.json`, import.meta.url), 'utf8')));
const n = [...SAMPLE].length;
const delta = (r) => (r.sys - (SCALE[r.size][0] === 'Text' ? r.text : r.display)) / (n * r.size);
let maxGap = 0;
const table = {};
for (const size of Object.keys(SCALE).map(Number)) {
  const ds = [];
  for (const weight of [400, 600]) {
    const [s, c] = runs.map(run => delta(run.rows.find(r => r.size === size && r.weight === weight)));
    maxGap = Math.max(maxGap, Math.abs(s - c));
    ds.push(s, c);
  }
  const mean = ds.reduce((a, b) => a + b) / ds.length;
  table[size] = Number((SCALE[size][1] - mean).toFixed(4));
}
console.log(runs.map(r => r.browser).join(' | '));
console.log('max |Δ safari − Δ chrome| over all sizes and weights:', maxGap.toFixed(5), 'em');
console.log('letter-spacing for system-ui (em):', JSON.stringify(table));
writeFileSync(new URL('./cui-tracking.json', import.meta.url), JSON.stringify(table) + '\n');
