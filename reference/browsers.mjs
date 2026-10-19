// The browsers the reference probes run in, behind one interface:
// open(name, url) loads the page and gives { label, run(expr), close() }, where
// run evaluates an expression whose value is a promise and hands back what it
// resolves to.
//   safari    Safari through safaridriver, raw WebDriver over HTTP (Safari's
//             "Allow remote automation" on)
//   ios       the same on an iPhone simulator (Xcode), whichever safaridriver finds
//   chrome    the installed Google Chrome through playwright-core
//   chromium  Playwright's own Chromium (PW names the module when it is not
//             playwright-core), with --font-render-hinting=none: on Linux it
//             hints every glyph to a whole-pixel advance otherwise, and no text
//             width measured there means anything
import { spawn } from 'node:child_process';

export function open(name, url) {
  switch (name) {
    case 'chrome': return viaPlaywright('chrome', url);
    case 'chromium': return viaPlaywright(undefined, url);
    case 'safari': return viaSafariDriver({ browserName: 'safari' }, url);
    case 'ios': return viaSafariDriver({ browserName: 'safari', platformName: 'iOS', 'safari:useSimulator': true, 'safari:deviceType': 'iPhone' }, url);
    default: throw new Error(`no browser ${name}: safari, ios, chrome or chromium`);
  }
}

async function viaPlaywright(channel, url) {
  const { chromium } = await import(process.env.PW ?? 'playwright-core');
  const browser = await chromium.launch(channel ? { channel } : { args: ['--font-render-hinting=none'] });
  try {
    const page = await browser.newPage({ viewport: { width: 1440, height: 900 } });
    await page.goto(url, { waitUntil: 'load', timeout: 120000 });
    return {
      label: `${channel ?? 'chromium'} ${browser.version()} on ${process.platform}`,
      run: (expr) => page.evaluate(expr),
      close: () => browser.close(),
    };
  } catch (e) {
    await browser.close();
    throw e;
  }
}

async function viaSafariDriver(capabilities, url) {
  const port = 4473;
  const base = `http://localhost:${port}`;
  const driver = spawn('/usr/bin/safaridriver', ['-p', String(port)], { stdio: ['ignore', 'inherit', 'inherit'] });
  const wd = async (method, path, body) => {
    const res = await fetch(base + path, { method, headers: { 'content-type': 'application/json' }, body: body && JSON.stringify(body) });
    const json = await res.json();
    if (!res.ok) throw new Error(`${method} ${path}: ${res.status} ${JSON.stringify(json.value)}`);
    return json.value;
  };
  let session;
  const close = async () => {
    if (session) await wd('DELETE', `/session/${session.sessionId}`).catch(e => console.error('delete session:', e.message));
    driver.kill();
  };
  try {
    for (let i = 0; ; i++) {
      try { await fetch(base + '/status'); break; } catch {
        if (i > 100) throw new Error('safaridriver did not start');
        await new Promise(r => setTimeout(r, 100));
      }
    }
    session = await wd('POST', '/session', { capabilities: { alwaysMatch: capabilities } });
    const id = session.sessionId;
    const c = session.capabilities;
    console.log('capabilities:', JSON.stringify(c));
    await wd('POST', `/session/${id}/timeouts`, { script: 120000, pageLoad: 180000 });
    await wd('POST', `/session/${id}/url`, { url });
    return {
      label: `safari ${c.browserVersion} on ${c.platformName} ${c['safari:platformVersion'] ?? ''}`.trim(),
      run: async (expr) => {
        const script = `const done = arguments[arguments.length - 1];\n${expr}.then(r => done({ ok: true, r }), e => done({ ok: false, e: String((e && e.stack) || e) }));`;
        const value = await wd('POST', `/session/${id}/execute/async`, { script, args: [] });
        if (!value.ok) throw new Error(value.e);
        return value.r;
      },
      close,
    };
  } catch (e) {
    await close();
    throw e;
  }
}
