// BROWSER=safari|ios|chrome|chromium EXPECT=sf|inter node run.mjs
// Serves the repository on 127.0.0.1 and opens page.html (the browsers:
// ../browsers.mjs). EXPECT=sf: the stack draws the system font and no Inter
// face was fetched; EXPECT=inter: the stack draws Inter, and its Latin face
// was fetched. The result is results/<browser>.json.
import { createServer } from 'node:http';
import { mkdirSync, readFileSync, writeFileSync } from 'node:fs';
import { readFile } from 'node:fs/promises';
import { extname } from 'node:path';
import { open } from '../browsers.mjs';

const BROWSER = process.env.BROWSER;
const EXPECT = process.env.EXPECT;
if (!['sf', 'inter'].includes(EXPECT)) throw new Error(`EXPECT must be sf or inter, not ${EXPECT}`);

const root = new URL('../../', import.meta.url);
const types = { '.html': 'text/html; charset=utf-8', '.css': 'text/css', '.woff2': 'font/woff2' };
const server = createServer((req, res) => {
  const path = decodeURIComponent(new URL(req.url, 'http://127.0.0.1').pathname);
  if (path.includes('..')) { res.writeHead(400); res.end(); return; }
  readFile(new URL('.' + path, root)).then(
    body => { res.writeHead(200, { 'content-type': types[extname(path)] ?? 'application/octet-stream' }); res.end(body); },
    () => { res.writeHead(404); res.end(); },
  );
});
await new Promise(r => server.listen(0, '127.0.0.1', r));
const url = `http://127.0.0.1:${server.address().port}/reference/stack/page.html`;

const src = readFileSync(new URL('./stack.js', import.meta.url), 'utf8');
let page, result;
try {
  page = await open(BROWSER, url);
  result = await page.run(`(async () => {\n${src}\nreturn stack();\n})()`);
} finally {
  await page?.close();
  server.close();
}
mkdirSync(new URL('./results/', import.meta.url), { recursive: true });
writeFileSync(new URL(`./results/${BROWSER}.json`, import.meta.url), JSON.stringify({ browser: page.label, date: new Date().toISOString(), expect: EXPECT, ...result }, null, 1));

console.log(`${page.label}\n${result.ua}`);
console.log('fetched by first paint:', JSON.stringify(result.before.fetched));
console.log("Inter's faces then:", JSON.stringify(result.before.inter));
console.log(`width: stack ${result.stackWidth.toFixed(3)}px, system-ui ${result.systemWidth.toFixed(3)}px, Inter ${result.interWidth.toFixed(3)}px`);
const same = (a, b) => Math.abs(a - b) < 0.01;
const interFetched = result.before.fetched.some(n => n.startsWith('inter-'));
const checks = EXPECT === 'sf'
  ? [['the stack draws the system font', same(result.stackWidth, result.systemWidth)],
     ['the system font is not Inter', !same(result.systemWidth, result.interWidth)],
     ['no Inter face fetched', !interFetched],
     ["every Inter face unloaded", result.before.inter.length > 0 && result.before.inter.every(s => s === 'unloaded')]]
  : [['the stack draws Inter', same(result.stackWidth, result.interWidth)],
     ['Inter is not the system font here', !same(result.systemWidth, result.interWidth)],
     ["Inter's Latin face fetched", result.before.fetched.includes('inter-roman-latin.woff2')]];
for (const [name, ok] of checks) console.log(`${ok ? 'ok  ' : 'FAIL'} ${name}`);
const pass = checks.every(([, ok]) => ok);
console.log(pass ? `PASS (EXPECT=${EXPECT})` : `FAIL (EXPECT=${EXPECT})`);
process.exit(pass ? 0 : 1);
