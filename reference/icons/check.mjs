// node check.mjs: the checkbox's check mark, apple.com's read live off the
// sign-in widget (never stored: an SF Pro Icons glyph, white on the checked
// box), the box captured at 20 pixels to the unit, its white ink read inside
// its 14 by 14 content box; ciderui's own drawn from its numbers, rasterized
// at the same scale (svg.js): the share of ink the two have in common. Fails
// below 0.85 (intersection over union), a mark set before the run, and if
// ciderui's drawing is not the checkbox's example's. The result is
// results/check.json.
import { mkdirSync, readFileSync, writeFileSync } from 'node:fs';
import { open } from '../browsers.mjs';

// ciderui's, in the content box: a line 1.35 thick, round at its ends and
// its corner, through the README's points
export const OURS = '<svg viewBox="0 0 14 14"><path d="M3.3 7.65 6.25 11.35 11.15 3.7" fill="none" stroke="#000" stroke-width="1.35" stroke-linecap="round" stroke-linejoin="round"/></svg>';

const S = 20;
const src = readFileSync(new URL('./svg.js', import.meta.url), 'utf8');
const page = await open('chromium', 'https://secure.store.apple.com/shop/account/home');
let result;
try {
  const pw = page.page;
  await pw.waitForTimeout(6000);
  const widget = pw.frames().find(f => f.url().includes('idmsa.apple.com'));
  await widget.click('#remember-me-label');
  await pw.waitForTimeout(400);
  const box = await (await widget.$('#remember-me-label .form-checkbox-indicator')).boundingBox();
  const cdp = await pw.context().newCDPSession(pw);
  const { data } = await cdp.send('Page.captureScreenshot', { format: 'png', clip: { x: box.x, y: box.y, width: 16, height: 16, scale: S } });
  await pw.setContent('<!doctype html><body></body>');
  result = await page.run(`(async () => {
${src}
const img = new Image();
img.src = 'data:image/png;base64,${data}';
await img.decode();
const c = document.createElement('canvas');
c.width = img.width; c.height = img.height;
const ctx = c.getContext('2d');
ctx.drawImage(img, 0, 0);
const d = ctx.getImageData(0, 0, c.width, c.height).data;
const W = c.width;
// apple.com's ink: white over the accent (its red channel past half), in the
// content box (the 1px border off each side)
const apple = { W: W - 2 * ${S}, H: c.height - 2 * ${S} };
apple.mask = new Uint8Array(apple.W * apple.H);
for (let y = 0; y < apple.H; y++) for (let x = 0; x < apple.W; x++) apple.mask[y * apple.W + x] = d[((y + ${S}) * W + x + ${S}) * 4] >= 128 ? 1 : 0;
let left = apple.W, right = -1, top = apple.H, bottom = -1;
for (let y = 0; y < apple.H; y++) for (let x = 0; x < apple.W; x++) if (apple.mask[y * apple.W + x]) { left = Math.min(left, x); right = Math.max(right, x); top = Math.min(top, y); bottom = Math.max(bottom, y); }
const u = v => +(v / ${S}).toFixed(3);
const ours = await raster(${JSON.stringify(OURS)}, ${S});
return { appleBox: { left: u(left), right: u(right + 1), top: u(top), bottom: u(bottom + 1) }, oursBox: ours.box, iou: compare(apple, ours) };
})()`);
} finally {
  await page.close();
}
mkdirSync(new URL('./results/', import.meta.url), { recursive: true });
writeFileSync(new URL('./results/check.json', import.meta.url), JSON.stringify({ browser: page.label, date: new Date().toISOString(), ours: OURS, ...result }, null, 1));
console.log(page.label);
console.log(JSON.stringify(result, null, 1));
const example = readFileSync(new URL('../../components/checkbox/examples/checkbox.html', import.meta.url), 'utf8');
const inner = OURS.replace(/^<svg[^>]*>|<\/svg>$/g, '').replaceAll('stroke="#000"', 'stroke="currentColor"');
const shipped = example.includes(inner);
console.log(`${shipped ? 'ok  ' : 'FAIL'} the check is the example's`);
const pass = result.iou >= 0.85 && shipped;
console.log(pass ? 'PASS: it shares at least 0.85 of its ink with apple.com\'s' : 'FAIL: below 0.85');
process.exit(pass ? 0 : 1);
