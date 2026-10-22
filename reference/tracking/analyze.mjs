// apple.com's scale (measured, ../apple.com/: the store's t- ladder and the
// home page's 56px, 2026-10-19): the size, the cut it sets at that size, and
// its letter-spacing in em. 32px and 48px are the store's and the product
// page's (+0.004em, -0.003em), not the home page's (+0.002em, -0.002em).
import { readFileSync } from 'node:fs';

export const SCALE = {
  12: ['Text', -0.01], 14: ['Text', -0.016], 17: ['Text', -0.022],
  19: ['Display', 0.012], 21: ['Display', 0.011], 24: ['Display', 0.009], 28: ['Display', 0.007],
  32: ['Display', 0.004], 40: ['Display', 0], 48: ['Display', -0.003], 56: ['Display', -0.005],
  64: ['Display', -0.009], 80: ['Display', -0.015], 96: ['Display', -0.015],
};
export const WEIGHTS = [400, 600];
export const SAMPLE = 'Sphinx of black quartz judge my vow 0123456789 AVATAR Wave';
// decided before the run: system-ui matches apple.com's cut within 0.003em per
// character (apple.com's smallest nonzero letter-spacing is 0.002em)
export const TOLERANCE = 0.003;

const em = (a, b, size) => (a - b) / ([...SAMPLE].length * size);
const f = x => (x >= 0 ? '+' : '') + x.toFixed(4);

export function analyze(label, result) {
  console.log(`\n${label}\n${result.ua}`);
  console.log('size weight cut      apple-ls  Δ(sys−cut)em  Δ(sys−Text)  Δ(sys−Display)  ok');
  let pass = true;
  for (const r of result.rows) {
    const [cut, ls] = SCALE[r.size];
    const d = em(r.sys, cut === 'Text' ? r.text : r.display, r.size);
    const ok = Math.abs(d) < TOLERANCE;
    pass &&= ok;
    console.log(`${String(r.size).padStart(4)} ${String(r.weight).padStart(6)} ${cut.padEnd(8)} ${f(ls).padStart(8)}  ${f(d).padStart(12)}  ${f(em(r.sys, r.text, r.size)).padStart(11)}  ${f(em(r.sys, r.display, r.size)).padStart(14)}  ${ok ? 'yes' : 'NO'}`);
  }
  console.log(pass ? `PASS: system-ui within ${TOLERANCE}em of apple.com's cut at every size` : `FAIL: system-ui off apple.com's cut by ${TOLERANCE}em or more somewhere`);
  return pass;
}

export const measureCall = (src) =>
  `(async () => {\n${src}\nreturn measure(${JSON.stringify(Object.keys(SCALE).map(Number))}, ${JSON.stringify(WEIGHTS)}, ${JSON.stringify(SAMPLE)});\n})()`;

// verify: decided before the run, the kit's own font stack (tokens.css's
// --cui-font, whichever of its names the browser has: -apple-system in
// Safari, BlinkMacSystemFont in Chrome, which has no -apple-system) at
// ciderui's letter-spacing within 0.0015em per character of apple.com's text
// (the table averages two weights and two browsers; at 56px 0.0015em is 0.08px
// a character); each name read beside it, system-ui too, which derived it
export const VERIFY_TOLERANCE = 0.0015;
export const STACK = readFileSync(new URL('../../tokens.css', import.meta.url), 'utf8').match(/--cui-font:\s*([^;]+);/)[1];
export const FAMILIES = ['system-ui', '-apple-system', 'BlinkMacSystemFont', STACK];

export function analyzeVerify(label, result) {
  console.log(`\n${label}\n${result.ua}`);
  console.log('Δ(name at ciderui ls − apple.com cut at apple.com ls), em per character');
  console.log('size weight' + FAMILIES.map(x => (x === STACK ? 'the stack' : x).padStart(20)).join(''));
  let pass = true;
  const worst = Object.fromEntries(FAMILIES.map(x => [x, 0]));
  for (const r of result.rows) {
    const cells = FAMILIES.map(fam => {
      const d = em(r.fams[fam], r.web, r.size);
      worst[fam] = Math.max(worst[fam], Math.abs(d));
      return f(d).padStart(20);
    });
    console.log(`${String(r.size).padStart(4)} ${String(r.weight).padStart(6)}${cells.join('')}`);
  }
  for (const fam of FAMILIES) console.log(`worst |Δ| ${fam === STACK ? `the stack (${STACK})` : fam}: ${worst[fam].toFixed(4)}em ${worst[fam] < VERIFY_TOLERANCE ? 'within' : 'OUTSIDE'} ${VERIFY_TOLERANCE}`);
  pass = worst[STACK] < VERIFY_TOLERANCE;
  console.log(pass ? 'PASS: the kit\'s font stack at ciderui\'s letter-spacing lands on apple.com\'s widths' : 'FAIL: the kit\'s font stack at ciderui\'s letter-spacing misses apple.com\'s widths');
  return pass;
}

// inter: no pass mark for the distance, which is what is being measured; the
// run fails only if Inter did not load (inter.js throws) or drew nothing.
export function analyzeInter(label, result) {
  console.log(`\n${label}\n${result.ua}`);
  console.log('size weight  Δ(Inter at ciderui ls − apple.com)  Inter alone: ls to land on apple.com');
  let drew = true;
  for (const r of result.rows) {
    drew &&= r.inter0 > 0;
    const d = em(r.inter, r.web, r.size);
    const own = SCALE[r.size][1] - em(r.inter0, r.web0, r.size);
    console.log(`${String(r.size).padStart(4)} ${String(r.weight).padStart(6)}  ${f(d).padStart(33)}  ${f(own).padStart(36)}`);
  }
  const worst = Math.max(...result.rows.map(r => Math.abs(em(r.inter, r.web, r.size))));
  console.log(`worst |Δ| Inter at ciderui's letter-spacing: ${worst.toFixed(4)}em a character`);
  return drew;
}
