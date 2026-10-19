// node tables.mjs: the README's two tables, from results/ and cui-tracking.json.
import { readFileSync } from 'node:fs';
import { SCALE, SAMPLE } from './analyze.mjs';
const load = name => JSON.parse(readFileSync(new URL(`./results/${name}.json`, import.meta.url), 'utf8'));
const cui = JSON.parse(readFileSync(new URL('./cui-tracking.json', import.meta.url), 'utf8'));
const n = [...SAMPLE].length;
const em = (a, b, size) => (a - b) / (n * size);
const f = x => (x >= 0 ? '+' : '−') + Math.abs(x).toFixed(4);
const sizes = Object.keys(SCALE).map(Number);

const runs = { safari: load('safari-measure'), chrome: load('chrome-measure'), ios: load('ios-measure') };
const delta = (run, size, w) => {
  const r = run.rows.find(r => r.size === size && r.weight === w);
  return em(r.sys, SCALE[size][0] === 'Text' ? r.text : r.display, size);
};
console.log('| size | apple.com cut | apple.com letter-spacing | Δ at 400 (Safari / Chrome / iOS) | Δ at 600 (Safari / Chrome / iOS) | ciderui letter-spacing |');
console.log('| ---: | --- | ---: | ---: | ---: | ---: |');
for (const s of sizes) {
  const cells = [400, 600].map(w => ['safari', 'chrome', 'ios'].map(b => f(delta(runs[b], s, w))).join(' / '));
  console.log(`| ${s}px | ${SCALE[s][0]} | ${f(SCALE[s][1])} | ${cells[0]} | ${cells[1]} | ${f(cui[s])} |`);
}
console.log();
const inter = load('safari-inter');
console.log('| size | Δ at 400 | Δ at 600 | Inter alone at 400 | Inter alone at 600 |');
console.log('| ---: | ---: | ---: | ---: | ---: |');
for (const s of sizes) {
  const rows = [400, 600].map(w => inter.rows.find(r => r.size === s && r.weight === w));
  console.log(`| ${s}px | ${rows.map(r => f(em(r.inter, r.web, s))).join(' | ')} | ${rows.map(r => f(SCALE[s][1] - em(r.inter0, r.web0, s))).join(' | ')} |`);
}
