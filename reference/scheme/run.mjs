// BROWSER=safari|ios|chrome|chromium node run.mjs
// Does light-dark() follow the color-scheme a section sets, and back again
// inside it? Serves page.html on 127.0.0.1 (the browsers: ../browsers.mjs);
// the result is results/<browser>.json.
import { createServer } from 'node:http';
import { mkdirSync, readFileSync, writeFileSync } from 'node:fs';
import { open } from '../browsers.mjs';

const BROWSER = process.env.BROWSER;
const html = readFileSync(new URL('./page.html', import.meta.url));
const server = createServer((req, res) => { res.writeHead(200, { 'content-type': 'text/html; charset=utf-8' }); res.end(html); });
await new Promise(r => server.listen(0, '127.0.0.1', r));
const src = readFileSync(new URL('./scheme.js', import.meta.url), 'utf8');
let page, result;
try {
  page = await open(BROWSER, `http://127.0.0.1:${server.address().port}/`);
  result = await page.run(`(async () => {\n${src}\nreturn scheme();\n})()`);
} finally {
  await page?.close();
  server.close();
}
mkdirSync(new URL('./results/', import.meta.url), { recursive: true });
writeFileSync(new URL(`./results/${BROWSER}.json`, import.meta.url), JSON.stringify({ browser: page.label, date: new Date().toISOString(), ...result }, null, 1));
const LIGHT = ['rgb(29, 29, 31)', 'rgb(255, 255, 255)'];
const DARK = ['rgb(245, 245, 247)', 'rgb(0, 0, 0)'];
const want = { page: LIGHT, dark: DARK, 'in-dark': DARK, back: LIGHT };
console.log(`${page.label}\n${result.ua}`);
let pass = true;
for (const [id, w] of Object.entries(want)) {
  const ok = JSON.stringify(result.got[id]) === JSON.stringify(w);
  pass &&= ok;
  console.log(`${ok ? 'ok  ' : 'FAIL'} ${id}: ${result.got[id].join(' on ')}${ok ? '' : ` (want ${w.join(' on ')})`}`);
}
console.log(pass ? 'PASS' : 'FAIL');
process.exit(pass ? 0 : 1);
