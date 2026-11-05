import { readFileSync } from 'node:fs';
import { join } from 'node:path';
import { ROOT, compile } from './compile.mjs';
import { open, quit } from './browser.mjs';
import { ref, WIDTHS } from './reference.mjs';
import { check, done } from './check.mjs';

const names = [...readFileSync(join(ROOT, 'type.css'), 'utf8').matchAll(/@utility cui-([a-z-]+)/g)].map(m => m[1]);
const tracking = JSON.parse(readFileSync(join(ROOT, 'reference/tracking/cui-tracking.json'), 'utf8'));
const css = compile({ classes: [...names.map(n => `cui-${n}`), 'lg:cui-headline-elevated', 'text-[20px]'].join(' ') });
const html = names.map(n => `<p id="${n}" class="cui-${n}">Sample text</p>`).join('')
  + '<p id="variant" class="cui-headline lg:cui-headline-elevated">v</p><p id="plain" class="cui-headline text-[20px]">p</p>';
const b = await open(css, html);

// the store defines no t-body: cui-body is the page's body text
const apple = (name, width, prop) => name === 'body'
  ? ref('store', 'page.body', width, prop)
  : ref('store', `type.t-${name}`, width, prop);

const STACK = '-apple-system, BlinkMacSystemFont, Inter, "Helvetica Neue", Helvetica, Arial, "PingFang TC", "Noto Sans TC", "Noto Sans CJK TC", "Microsoft JhengHei", sans-serif';
try {
  for (const width of WIDTHS) {
    await b.resize(width);
    for (const name of names) {
      const got = await b.read(`#${name}`, ['font-size', 'line-height', 'font-weight', 'letter-spacing', 'font-family']);
      for (const prop of ['font-size', 'line-height', 'font-weight']) check(`cui-${name} ${prop} at ${width}px`, got[prop], apple(name, width, prop));
      // for apple.com's size, not the read one: a wrong size is its own failure
      const size = parseFloat(apple(name, width, 'font-size'));
      check(`cui-${name} letter-spacing at ${width}px`, got['letter-spacing'], `${size * tracking[size]}px`, { tolerance: 0.001 });
      check(`cui-${name} font-family at ${width}px`, got['font-family'], STACK);
    }
  }
  for (const [width, at] of [[734, 390], [735, 900]]) {
    await b.resize(width);
    for (const name of names) check(`cui-${name} font-size at ${width}px is its ${at}px size`, (await b.read(`#${name}`, ['font-size']))['font-size'], apple(name, at, 'font-size'));
  }
  for (const [width, variant] of [[390, '32px'], [900, '48px'], [1068, '48px'], [1069, '64px'], [1600, '64px']]) {
    await b.resize(width);
    check(`lg:cui-headline-elevated over cui-headline at ${width}px`, (await b.read('#variant', ['font-size']))['font-size'], variant);
    check(`text-[20px] over cui-headline at ${width}px`, (await b.read('#plain', ['font-size']))['font-size'], '20px');
  }
} finally {
  await b.close();
}

await quit();
done('type');
