// BROWSER=safari|ios|chrome|chromium|firefox MODE=measure|verify|inter node run.mjs
// The page is https://www.apple.com/ (the browsers: ../browsers.mjs); the
// result is results/<browser>-<mode>.json.
import { mkdirSync, readFileSync, writeFileSync } from 'node:fs';
import { open } from '../browsers.mjs';
import { SCALE, WEIGHTS, SAMPLE, FAMILIES, analyze, analyzeVerify, analyzeInter } from './analyze.mjs';

const BROWSER = process.env.BROWSER;
const MODE = process.env.MODE ?? 'measure';
const src = readFileSync(new URL(`./${MODE}.js`, import.meta.url), 'utf8');
const tracking = () => JSON.parse(readFileSync(new URL('./cui-tracking.json', import.meta.url), 'utf8'));
const args = {
  measure: () => [Object.keys(SCALE).map(Number), WEIGHTS, SAMPLE],
  verify: () => [SCALE, tracking(), WEIGHTS, SAMPLE, FAMILIES],
  // INTER names the font file, by default the kit's Latin subset
  inter: () => [SCALE, tracking(), WEIGHTS, SAMPLE,
    readFileSync(process.env.INTER ?? new URL('../../fonts/inter-roman-latin.woff2', import.meta.url)).toString('base64')],
}[MODE]();
const expr = `(async () => {\n${src}\nreturn ${MODE}(...${JSON.stringify(args)});\n})()`;

const page = await open(BROWSER, 'https://www.apple.com/');
let result;
try {
  result = await page.run(expr);
} finally {
  await page.close();
}
mkdirSync(new URL('./results/', import.meta.url), { recursive: true });
writeFileSync(new URL(`./results/${BROWSER}-${MODE}.json`, import.meta.url), JSON.stringify({ browser: page.label, date: new Date().toISOString(), ...result }, null, 1));
process.exit({ measure: analyze, verify: analyzeVerify, inter: analyzeInter }[MODE](page.label, result) ? 0 : 1);
