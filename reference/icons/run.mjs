// node run.mjs: the "more" chevron's numbers off apple.com's /iphone/ page,
// in Playwright's Chromium (../browsers.mjs), at 400px; the result is
// results/more.json. Fails if ciderui's chevron drawn from them shares less
// than 0.9 of its ink with apple.com's (intersection over union).
import { mkdirSync, readFileSync, writeFileSync } from 'node:fs';
import { open } from '../browsers.mjs';
const src = readFileSync(new URL('./glyph.js', import.meta.url), 'utf8');
const page = await open(process.env.BROWSER ?? 'chromium', 'https://www.apple.com/iphone/');
let result;
try {
  result = await page.run(`(async () => {\n${src}\nreturn glyph(400);\n})()`);
} finally {
  await page.close();
}
mkdirSync(new URL('./results/', import.meta.url), { recursive: true });
writeFileSync(new URL('./results/more.json', import.meta.url), JSON.stringify({ browser: page.label, date: new Date().toISOString(), ...result }, null, 1));
console.log(page.label);
console.log(JSON.stringify(result, null, 1));
const pass = result.iou >= 0.9;
console.log(pass ? `PASS: ciderui's chevron shares ${result.iou} of apple.com's ink` : `FAIL: ${result.iou} of the ink in common, below 0.9`);
process.exit(pass ? 0 : 1);
