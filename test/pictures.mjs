import { mkdirSync, readFileSync, writeFileSync } from 'node:fs';
import { createHash } from 'node:crypto';
import { join } from 'node:path';
import { COMPONENTS, ROOT, exampleFiles, exampleHtml, shipped } from './compile.mjs';
import { open, quit } from './browser.mjs';
import { WIDTHS } from './reference.mjs';
import { check, done } from './check.mjs';

const UPDATE = process.env.UPDATE_GOLDEN === '1';
const GOLDEN = join(ROOT, 'test/pictures.json');
const FAILED = 'test/pictures-failed';
const golden = UPDATE ? {} : JSON.parse(readFileSync(GOLDEN, 'utf8'));
const seen = {};

const pages = {};
for (const set of ['site', 'console']) {
  pages[set] = await open(shipped(set), '');
  // every face now: one first asked for after a resize would be drawn in its fallback
  await pages[set].page.evaluate(() => Promise.all([...document.fonts].map((face) => face.load())));
}

for (const part of COMPONENTS) {
  const page = pages[part.set];
  for (const file of exampleFiles(part)) {
    await page.page.evaluate((html) => { document.body.innerHTML = html; }, exampleHtml(part, file));
    for (const width of WIDTHS) {
      await page.resize(width);
      const picture = await page.page.screenshot({ fullPage: true, animations: 'disabled', caret: 'hide' });
      const hash = createHash('sha256').update(picture).digest('hex');
      const name = `${part.name}/${file} at ${width}`;
      seen[name] = hash;
      const saved = `${FAILED}/${part.name}-${file}-${width}.png`;
      if (!UPDATE && !check(`${name}px: the picture changed, see ${saved} (UPDATE_GOLDEN=1 if meant)`, hash, golden[name])) {
        mkdirSync(join(ROOT, FAILED), { recursive: true });
        writeFileSync(join(ROOT, saved), picture);
      }
    }
  }
}
await quit();

if (UPDATE) writeFileSync(GOLDEN, JSON.stringify(seen, null, 1) + '\n');
check('every name in pictures.json has an example', Object.keys(golden).filter((name) => !(name in seen)).join(' '), '');
done('pictures');
