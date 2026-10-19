// node report.mjs <name> [probe-prefix]: each probe's values across the widths
// of results/<name>.json, a line per value that is read.
import { readFileSync } from 'node:fs';
const [name, prefix = ''] = process.argv.slice(2);
const r = JSON.parse(readFileSync(new URL(`./results/${name}.json`, import.meta.url), 'utf8'));
const widths = Object.keys(r.widths);
console.log(`${name}: ${r.finalUrl} ${r.browser} ${r.date}`);
const probes = [...new Set(widths.flatMap(w => Object.keys(r.widths[w])))].filter(p => p.startsWith(prefix));
// a typography class that reads as the unclassed paragraph at every width is
// one the page does not define
const control = (p) => p.startsWith('type.') && p !== 'type.none' && widths.every(w =>
  r.widths[w]['type.none'] && JSON.stringify(r.widths[w][p]?.styles) === JSON.stringify(r.widths[w]['type.none'].styles));
for (const p of probes) {
  if (control(p)) { console.log(`\n${p}  NOT DEFINED HERE: reads as type.none at every width`); continue; }
  console.log(`\n${p}  ${(r.widths[widths[0]][p]?.html ?? 'null').replace(/\s+/g, ' ').slice(0, 150)}`);
  const keys = new Set();
  for (const w of widths) {
    const v = r.widths[w][p];
    if (!v) continue;
    Object.keys(v.styles).forEach(k => keys.add(k));
    if (v.box) ['box.width', 'box.height'].forEach(k => keys.add(k));
    // a state's value only where it differs from rest, and never its transition
    // (the harness turns transitions off while it forces a state)
    for (const [s, st] of Object.entries(v.states ?? {})) {
      Object.keys(st).filter(k => !k.startsWith('transition') && widths.some(w => {
        const x = r.widths[w][p];
        return x && x.states?.[s]?.[k] !== x.styles[k];
      })).forEach(k => keys.add(`:${s} ${k}`));
    }
  }
  for (const k of keys) {
    const vals = widths.map(w => {
      const v = r.widths[w][p];
      if (!v) return 'null';
      if (k.startsWith('box.')) return String(+v.box[k.slice(4)].toFixed(3));
      if (k.startsWith(':')) { const [s, prop] = k.slice(1).split(' '); return v.states?.[s]?.[prop] ?? '-'; }
      return v.styles[k];
    });
    const same = vals.every(v => v === vals[0]);
    console.log(`  ${k.padEnd(28)} ${same ? `all ${vals[0]}` : widths.map((w, i) => `${w}:${vals[i]}`).join('  ')}`);
  }
}
