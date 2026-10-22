// node form.mjs: the store's form icons, apple.com's read live (never
// stored) and ciderui's own drawn from their numbers: the share of ink each
// pair has in common. The error's mark, SF Pro Icons U+F312 before a
// field's message (the gift card's page, giftcard:error.icon), drawn at 400px
// against ciderui's ALERT drawn on the same canvas, and where its ink sits in
// the message's row (the page at four device pixels to the CSS pixel, an
// empty required field the focus has left: nothing typed); the dropdown's
// chevron, the SVG its background draws (iPhone 18 Pro's buy page, its
// trade-in, tradein:dropdown.chevron), rasterized at 20 pixels to the unit
// (svg.js) against ciderui's DROPDOWN. Fails below 0.85 (intersection over
// union), a mark set before the run, or if the field's examples draw others.
// The result is results/form.json.
import { mkdirSync, readFileSync, writeFileSync } from 'node:fs';
import { open } from '../browsers.mjs';

// ciderui's, the README's numbers. ALERT in thousandths of an em, its box
// the glyph's advance (1.171em) by its ink's height (0.995em, from 0.85em
// above the baseline): a ring at (585, 497.5), radius 456.25, 82.5 thick; the
// stem from (584, 280) to (584, 550), 70 thick, round; the dot at (585,
// 692.5), radius 50. DROPDOWN in apple.com's 21 by 21: a line through
// (6.075, 9.275), (10.5, 13.7) and (14.925, 9.275), 1.45 thick, round.
export const ALERT = '<circle cx="585" cy="497.5" r="456.25" fill="none" stroke="#000" stroke-width="82.5"/><path d="M584 280V550" fill="none" stroke="#000" stroke-width="70" stroke-linecap="round"/><circle cx="585" cy="692.5" r="50" fill="#000"/>';
export const DROPDOWN = '<svg viewBox="0 0 21 21"><path d="M6.075 9.275 10.5 13.7 14.925 9.275" fill="none" stroke="#000" stroke-width="1.45" stroke-linecap="round" stroke-linejoin="round"/></svg>';

const src = readFileSync(new URL('./svg.js', import.meta.url), 'utf8');
const result = {};
let page = await open('chromium', 'https://www.apple.com/shop/buy-giftcard/giftcard');
try {
  // the mark, drawn on a canvas at 400px, and ciderui's on the same one
  result.alert = await page.run(`(async () => {
    const size = 400, ch = String.fromCodePoint(0xF312), font = size + 'px "SF Pro Icons"';
    if (!(await document.fonts.load(font, ch)).length) throw new Error('no SF Pro Icons face on the page');
    const W = size * 2, H = size * 2, X = size / 2, BASE = size * 1.4;
    const mask = async (draw) => {
      const c = document.createElement('canvas'); c.width = W; c.height = H;
      const ctx = c.getContext('2d'); await draw(ctx);
      const d = ctx.getImageData(0, 0, W, H).data, m = new Uint8Array(W * H);
      for (let i = 0; i < W * H; i++) m[i] = d[i * 4 + 3] >= 128 ? 1 : 0;
      return m;
    };
    const apple = await mask((ctx) => { ctx.font = font; ctx.fillStyle = '#000'; ctx.fillText(ch, X, BASE); });
    const ours = await mask(async (ctx) => {
      const img = new Image();
      img.src = 'data:image/svg+xml,' + encodeURIComponent('<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 1171 995" width="' + 1.171 * size + '" height="' + 0.995 * size + '">' + ${JSON.stringify(ALERT)} + '</svg>');
      await img.decode();
      ctx.drawImage(img, X, BASE - 0.85 * size, 1.171 * size, 0.995 * size);
    });
    let l = W, r = -1, t = H, b = -1, both = 0, either = 0;
    for (let y = 0; y < H; y++) for (let x = 0; x < W; x++) if (apple[y * W + x]) { l = Math.min(l, x); r = Math.max(r, x); t = Math.min(t, y); b = Math.max(b, y); }
    for (let i = 0; i < W * H; i++) { both += apple[i] & ours[i]; either += apple[i] | ours[i]; }
    const em = v => +(v / size).toFixed(4);
    const ctx = document.createElement('canvas').getContext('2d'); ctx.font = font;
    return { char: 'U+F312', size, advance: em(ctx.measureText(ch).width), ink: { left: em(l - X), right: em(r + 1 - X), top: em(BASE - t), bottom: em(BASE - b - 1) }, iou: +(both / either).toFixed(4) };
  })()`);
} finally {
  await page.close();
}

// where the mark's ink sits in its message's row: the page at four device
// pixels to the CSS pixel, an empty required field the focus left
{
  const { chromium } = await import(process.env.PW ?? 'playwright-core');
  const browser = await chromium.launch({ args: ['--font-render-hinting=none'] });
  try {
    const p = await browser.newPage({ viewport: { width: 1200, height: 900 }, deviceScaleFactor: 4 });
    await p.goto('https://www.apple.com/shop/buy-giftcard/giftcard', { waitUntil: 'load', timeout: 120000 });
    await p.waitForTimeout(2500);
    for (const s of ['input[name=dimensionGiftCardType][value=electronic]', 'input[name=dimensionGiftCardTheme]', '.rf-choose-amount-dimensions > div:nth-child(2) .form-selector-input']) {
      await p.click(s);
      await p.waitForTimeout(1200);
    }
    const field = '.rf-giftcard-delivery-section-fields .form-textbox:nth-child(2)';
    await p.click(`${field} .form-textbox-input`);
    await p.evaluate(() => document.activeElement.blur());
    await p.waitForTimeout(800);
    const row = p.locator(`${field} .form-message-wrapper`).first();
    const box = await row.boundingBox();
    const text = await p.locator(`${field} .form-message`).first().boundingBox();
    const png = (await row.screenshot()).toString('base64');
    // the ink left of the text: its box in CSS pixels from the row's corner
    result.alertPlace = await p.evaluate(async ([png, textLeft]) => {
      const img = new Image(); img.src = `data:image/png;base64,${png}`; await img.decode();
      const c = document.createElement('canvas'); c.width = img.width; c.height = img.height;
      const g = c.getContext('2d'); g.drawImage(img, 0, 0);
      const d = g.getImageData(0, 0, img.width, img.height).data, k = img.width / textLeft.row;
      let l = img.width, r = -1, t = img.height, b = -1;
      for (let y = 0; y < img.height; y++) for (let x = 0; x < Math.floor(textLeft.at * k); x++) {
        const i = (y * img.width + x) * 4;
        if (d[i] - d[i + 1] > 60) { l = Math.min(l, x); r = Math.max(r, x); t = Math.min(t, y); b = Math.max(b, y); }
      }
      // no red ink left of the text is no mark found, not a box: a box read
      // so would be inside out (its left past its right), which the check
      // below took as a place
      if (r < l || b < t) throw new Error('no red ink left of the message\'s text: the mark is not where it was read');
      return { scale: k, ink: { left: l / k, right: (r + 1) / k, top: t / k, bottom: (b + 1) / k }, textLeft: textLeft.at };
    }, [png, { row: box.width, at: text.x - box.x }]);
    result.alertPlace.rowHeight = box.height;
  } finally {
    await browser.close();
  }
}

// the dropdown's chevron, its background's SVG
page = await open('chromium', 'https://www.apple.com/shop/buy-iphone/iphone-18-pro');
try {
  await page.page.waitForTimeout(2500);
  const apple = await page.run(`Promise.resolve((() => { const u = getComputedStyle(document.querySelector('.form-dropdown-chevron')).backgroundImage; return decodeURIComponent(u.slice(u.indexOf(',') + 1, -2)); })())`);
  await page.page.setContent('<!doctype html><body></body>');
  result.dropdown = await page.run(`(async () => {
${src}
const a = await raster(${JSON.stringify(apple)}, 20), o = await raster(${JSON.stringify(DROPDOWN)}, 20);
return { appleBox: a.box, oursBox: o.box, iou: compare(a, o), rows: [10, 12, 13.5].map(y => [y, a.runs('row', y), o.runs('row', y)]) };
})()`);
} finally {
  await page.close();
}

mkdirSync(new URL('./results/', import.meta.url), { recursive: true });
writeFileSync(new URL('./results/form.json', import.meta.url), JSON.stringify({ browser: page.label, date: new Date().toISOString(), ours: { alert: ALERT, dropdown: DROPDOWN }, ...result }, null, 1));
console.log(page.label);
console.log(JSON.stringify(result, null, 1));
// the icons compared are the ones the kit ships, in the field's examples, in
// currentColor
const examples = ['error', 'select'].map(n => readFileSync(new URL(`../../components/field/examples/${n}.html`, import.meta.url), 'utf8')).join('\n');
const shipped = [['alert', ALERT], ['dropdown', DROPDOWN.replace(/^<svg[^>]*>|<\/svg>$/g, '')]].map(([k, inner]) => {
  const ok = examples.includes(inner.replaceAll('"#000"', '"currentColor"'));
  console.log(`${ok ? 'ok  ' : 'FAIL'} ${k} is the examples'`);
  return ok;
});
const pass = result.alert.iou >= 0.85 && result.dropdown.iou >= 0.85 && shipped.every(Boolean);
console.log(pass ? 'PASS: each shares at least 0.85 of its ink with apple.com\'s' : 'FAIL');
process.exit(pass ? 0 : 1);
