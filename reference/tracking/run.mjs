// BROWSER=safari|ios|chrome|chromium MODE=measure|verify node run.mjs
// safari and ios through safaridriver (raw WebDriver), chrome through
// playwright-core and the installed Google Chrome, chromium through Playwright's
// own build (PW names the module when it is not playwright-core beside this
// file); the page is https://www.apple.com/; the result is results/<browser>-<mode>.json.
import { spawn } from 'node:child_process';
import { mkdirSync, readFileSync, writeFileSync } from 'node:fs';
import { SCALE, WEIGHTS, SAMPLE, FAMILIES, analyze, analyzeVerify } from './analyze.mjs';

const BROWSER = process.env.BROWSER;
const MODE = process.env.MODE ?? 'measure';
const src = readFileSync(new URL(`./${MODE}.js`, import.meta.url), 'utf8');
const args = MODE === 'measure'
  ? [Object.keys(SCALE).map(Number), WEIGHTS, SAMPLE]
  : [SCALE, JSON.parse(readFileSync(new URL('./cui-tracking.json', import.meta.url), 'utf8')), WEIGHTS, SAMPLE, FAMILIES];
const expr = `(async () => {\n${src}\nreturn ${MODE}(...${JSON.stringify(args)});\n})()`;

async function viaChromium(channel) {
  const { chromium } = await import(process.env.PW ?? 'playwright-core');
  const browser = await chromium.launch(channel ? { channel } : {});
  try {
    const page = await browser.newPage({ viewport: { width: 1440, height: 900 } });
    await page.goto('https://www.apple.com/', { waitUntil: 'load', timeout: 120000 });
    return { label: `${channel ?? 'chromium'} ${browser.version()} on ${process.platform}`, result: await page.evaluate(expr) };
  } finally {
    await browser.close();
  }
}

async function viaSafariDriver(capabilities) {
  const PORT = 4473;
  const base = `http://localhost:${PORT}`;
  const driver = spawn('/usr/bin/safaridriver', ['-p', String(PORT)], { stdio: ['ignore', 'inherit', 'inherit'] });
  const wd = async (method, path, body) => {
    const res = await fetch(base + path, { method, headers: { 'content-type': 'application/json' }, body: body && JSON.stringify(body) });
    const json = await res.json();
    if (!res.ok) throw new Error(`${method} ${path}: ${res.status} ${JSON.stringify(json.value)}`);
    return json.value;
  };
  let session;
  try {
    for (let i = 0; ; i++) {
      try { await fetch(base + '/status'); break; } catch {
        if (i > 100) throw new Error('safaridriver did not start');
        await new Promise(r => setTimeout(r, 100));
      }
    }
    session = await wd('POST', '/session', { capabilities: { alwaysMatch: capabilities } });
    const id = session.sessionId;
    console.log('capabilities:', JSON.stringify(session.capabilities));
    await wd('POST', `/session/${id}/timeouts`, { script: 120000, pageLoad: 180000 });
    await wd('POST', `/session/${id}/url`, { url: 'https://www.apple.com/' });
    const script = `const done = arguments[arguments.length - 1];\n${expr}.then(r => done({ ok: true, r }), e => done({ ok: false, e: String((e && e.stack) || e) }));`;
    const value = await wd('POST', `/session/${id}/execute/async`, { script, args: [] });
    if (!value.ok) throw new Error(value.e);
    const c = session.capabilities;
    return { label: `safari ${c.browserVersion} on ${c.platformName} ${c['safari:platformVersion'] ?? ''}`.trim(), result: value.r };
  } finally {
    if (session) await wd('DELETE', `/session/${session.sessionId}`).catch(e => console.error('delete session:', e.message));
    driver.kill();
  }
}

const run = {
  chrome: () => viaChromium('chrome'),
  chromium: () => viaChromium(),
  safari: () => viaSafariDriver({ browserName: 'safari' }),
  ios: () => viaSafariDriver({ browserName: 'safari', platformName: 'iOS', 'safari:useSimulator': true, 'safari:deviceType': 'iPhone' }),
}[BROWSER];
if (!run) throw new Error(`BROWSER must be safari, ios or chrome, not ${BROWSER}`);
const { label, result } = await run();
mkdirSync(new URL('./results/', import.meta.url), { recursive: true });
writeFileSync(new URL(`./results/${BROWSER}-${MODE}.json`, import.meta.url), JSON.stringify({ browser: label, date: new Date().toISOString(), ...result }, null, 1));
process.exit((MODE === 'measure' ? analyze : analyzeVerify)(label, result) ? 0 : 1);
