// node run.mjs: the "more" chevron's numbers off apple.com's /iphone/ page,
// in Playwright's Chromium (../browsers.mjs), at 400px; the result is
// results/more.json. Fails if the chevron the kit ships (the more link's
// example's, icon.rs's MORE: its box from 0.01em below the baseline to 0.53em
// above it, its view box in thousandths of an em) shares less than 0.9 of its
// ink with apple.com's (intersection over union); a chevron drawn from the
// numbers alone is read beside it (iou).
import { mkdirSync, readFileSync, writeFileSync } from 'node:fs';
import { open } from '../browsers.mjs';
const src = readFileSync(new URL('./glyph.js', import.meta.url), 'utf8');
// the shipped chevron, read off the example: its path's points and stroke
const example = readFileSync(new URL('../../components/link/examples/more.html', import.meta.url), 'utf8');
const [, d, width] = example.match(/<svg class="cui-link-icon"[^>]*>\s*<path d="M([\d. ]+)"[^>]*stroke-width="([\d.]+)"/) ?? [];
if (!d) throw new Error('no chevron in the more link\'s example');
const numbers = d.trim().split(/\s+/).map(Number);
const shipped = { points: Array.from({ length: numbers.length / 2 }, (_, i) => [numbers[2 * i] / 1000, 0.53 - numbers[2 * i + 1] / 1000]), stroke: Number(width) / 1000 };
const page = await open(process.env.BROWSER ?? 'chromium', 'https://www.apple.com/iphone/');
let result;
try {
  result = await page.run(`(async () => {\n${src}\nreturn glyph(400, ${JSON.stringify(shipped)});\n})()`);
} finally {
  await page.close();
}
mkdirSync(new URL('./results/', import.meta.url), { recursive: true });
writeFileSync(new URL('./results/more.json', import.meta.url), JSON.stringify({ browser: page.label, date: new Date().toISOString(), ...result }, null, 1));
console.log(page.label);
console.log(JSON.stringify(result, null, 1));
const pass = result.shippedIou >= 0.9;
console.log(pass ? `PASS: the chevron the kit ships shares ${result.shippedIou} of apple.com's ink (drawn from the numbers alone, ${result.iou})` : `FAIL: the shipped chevron's ${result.shippedIou} of the ink in common, below 0.9`);
process.exit(pass ? 0 : 1);
