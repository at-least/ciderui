// The browsers the reference probes run in, behind one interface:
// open(name, url) loads the page and gives { label, run(expr), click(selector),
// press(key), move(x, y, ms), go(url), size(width, height), shot(), close() },
// where run evaluates an expression whose value is a promise and hands back
// what it resolves to, click, press and move are the browser's own input (a
// trusted click on the element; a key pressed: the WebDriver key's name, as
// Escape, or keys held together, as Alt+Tab; the pointer moved to a point of
// the window over ms milliseconds, as a hand moves it: WebDriver's pointer
// move of that duration, Playwright's in a step a frame, none on the
// simulator, which has no pointer), go loads another page in the same window,
// waiting for its load, size sets the window's size (Playwright's viewport;
// Safari's window, its frame included), and shot takes a picture of what the
// window shows, a PNG in base64 (at the screen's pixels: the window's width
// times its pixel ratio).
//   safari    Safari through safaridriver, raw WebDriver over HTTP (Safari's
//             "Allow remote automation" on), in a Safari the run launches and
//             quits when it closes: Safari 27 takes no session in one it did
//             not launch ("Safari was not launched for automation"), so quit
//             Safari before the run
//   ios       the same on an iPhone simulator (Xcode), whichever safaridriver finds
//   chrome    the installed Google Chrome through playwright-core
//   chromium  Playwright's own Chromium (PW names the module when it is not
//             playwright-core), with --font-render-hinting=none: on Linux it
//             hints every glyph to a whole-pixel advance otherwise, and no text
//             width measured there means anything
//   firefox   Playwright's own Firefox (`playwright-core install firefox`):
//             Gecko with Playwright's automation patched in, not Mozilla's
//             release build; with its window, on a virtual X display of the
//             run's own (Xvfb), its pictures that display's pixels
//             (ImageMagick's import): Playwright's picture of a Firefox page
//             leaves backdrop-filter out, headless or not, which Firefox
//             paints on the display (2026-10-22)
import { execFileSync, spawn, spawnSync } from 'node:child_process';

// options: Playwright's page options (a touch screen: hasTouch), for its
// browsers only
export function open(name, url, options) {
  if (options && ['safari', 'ios'].includes(name)) throw new Error(`no page options in ${name}`);
  switch (name) {
    case 'chrome': return viaPlaywright('chromium', 'chrome', url, options);
    case 'chromium': return viaPlaywright('chromium', undefined, url, options);
    case 'firefox': return viaPlaywright('firefox', undefined, url, options);
    case 'safari': return viaSafariDriver({ browserName: 'safari' }, url);
    case 'ios': return viaSafariDriver({ browserName: 'safari', platformName: 'iOS', 'safari:useSimulator': true, 'safari:deviceType': 'iPhone' }, url);
    default: throw new Error(`no browser ${name}: safari, ios, chrome, chromium or firefox`);
  }
}

async function viaPlaywright(engine, channel, url, options) {
  const playwright = await import(process.env.PW ?? 'playwright-core');
  const screen = engine === 'firefox' ? await xvfb() : undefined;
  let browser;
  try {
    browser = await playwright[engine].launch(screen ? { headless: false, env: screen.env } : channel ? { channel } : { args: ['--font-render-hinting=none'] });
  } catch (e) {
    screen?.close();
    throw e;
  }
  try {
    const page = await browser.newPage({ viewport: { width: 1440, height: 900 }, ...options });
    await page.goto(url, { waitUntil: 'load', timeout: 120000 });
    return {
      label: `${channel ?? engine} ${browser.version()} on ${process.platform}`,
      run: (expr) => page.evaluate(expr),
      click: (selector) => page.click(selector),
      press: (key) => page.keyboard.press(key),
      move: (x, y, ms) => page.mouse.move(x, y, { steps: Math.max(1, Math.round(ms / 16)) }),
      go: (to) => page.goto(to, { waitUntil: 'load', timeout: 120000 }),
      size: (width, height) => page.setViewportSize({ width, height }),
      shot: screen ? () => screen.shot(page) : async () => (await page.screenshot()).toString('base64'),
      close: async () => {
        await browser.close();
        screen?.close();
      },
      page, // Playwright's, for a probe that resizes the viewport or forces states
    };
  } catch (e) {
    await browser.close();
    screen?.close();
    throw e;
  }
}

// a virtual X display of the run's own, on a number Xvfb finds free
// (-displayfd); shot(page) crops the page out of the display's pixels, where
// Firefox says its page is (mozInnerScreenX, mozInnerScreenY)
async function xvfb() {
  const server = spawn('Xvfb', ['-displayfd', '3', '-screen', '0', '2560x1600x24', '-nolisten', 'tcp'], { stdio: ['ignore', 'ignore', 'pipe', 'pipe'] });
  let log = '';
  server.stderr.on('data', (c) => { log += c; });
  const display = await new Promise((resolve, reject) => {
    let out = '';
    server.stdio[3].on('data', (c) => {
      out += c;
      if (out.includes('\n')) resolve(`:${out.trim()}`);
    });
    server.on('error', reject);
    server.on('exit', (code) => reject(new Error(`Xvfb exited ${code}: ${log}`)));
  });
  // Firefox on X, not on the desktop's Wayland
  const env = { ...process.env, DISPLAY: display, GDK_BACKEND: 'x11', MOZ_ENABLE_WAYLAND: '0' };
  delete env.WAYLAND_DISPLAY;
  return {
    env,
    shot: async (page) => {
      const [x, y, w, h] = await page.evaluate(() => [mozInnerScreenX, mozInnerScreenY, innerWidth, innerHeight].map((v) => Math.round(v * devicePixelRatio)));
      return execFileSync('import', ['-display', display, '-window', 'root', '-crop', `${w}x${h}+${x}+${y}`, '+repage', 'png:-'], { maxBuffer: 1 << 28 }).toString('base64');
    },
    close: () => server.kill(),
  };
}

// WebDriver's code points for the keys a probe presses
const KEYS = { Tab: '\uE004', Enter: '\uE006', Escape: '\uE00C', Alt: '\uE00A', Shift: '\uE008' };

async function viaSafariDriver(capabilities, url) {
  const port = 4473;
  const base = `http://localhost:${port}`;
  // On the Mac a Safari the run launches is quit when it closes. Left with no
  // window, macOS quits it under memory pressure and still counts it as running
  // (automatic termination): safaridriver then waits for it and launches no
  // other ("waiting for its RWIApplication"). A Safari running before the run
  // is left as it is, and takes no session unless safaridriver launched it:
  // one opened by LaunchServices rejects it ("Safari was not launched for
  // automation"), and the one macOS reopens after quitting it does not even
  // connect (2026-10-19).
  const mac = !capabilities['safari:useSimulator'];
  const running = () => spawnSync('pgrep', ['-x', 'Safari']).status === 0;
  const launched = mac && !running();
  const driver = spawn('/usr/bin/safaridriver', ['-p', String(port)], { stdio: ['ignore', 'inherit', 'inherit'] });
  const wd = async (method, path, body) => {
    const res = await fetch(base + path, { method, headers: { 'content-type': 'application/json' }, body: body && JSON.stringify(body) });
    const json = await res.json();
    if (!res.ok) throw new Error(`${method} ${path}: ${res.status} ${JSON.stringify(json.value)}`);
    return json.value;
  };
  let session;
  // where the pointer is: WebDriver's starts at the window's corner
  let pointer = [0, 0];
  const close = async () => {
    if (session) await wd('DELETE', `/session/${session.sessionId}`).catch(e => console.error('delete session:', e.message));
    driver.kill();
    if (launched && running()) {
      execFileSync('osascript', ['-e', 'quit app "Safari"']);
      for (let i = 0; running(); i++) {
        if (i > 100) throw new Error('Safari did not quit');
        await new Promise(r => setTimeout(r, 100));
      }
    }
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
      click: async (selector) => {
        const el = await wd('POST', `/session/${id}/element`, { using: 'css selector', value: selector });
        await wd('POST', `/session/${id}/element/${Object.values(el)[0]}/click`, {});
      },
      // a key, or keys held together as Playwright writes them (Alt+Tab)
      press: async (key) => {
        const values = key.split('+').map(k => {
          if (!KEYS[k]) throw new Error(`no WebDriver key for ${k}`);
          return KEYS[k];
        });
        const actions = [...values.map(value => ({ type: 'keyDown', value })), ...values.toReversed().map(value => ({ type: 'keyUp', value }))];
        await wd('POST', `/session/${id}/actions`, { actions: [{ type: 'key', id: 'keyboard', actions }] });
      },
      // a move of a duration is one jump at its end in Safari 27 (one
      // pointermove, 2026-10-21): it is moved a step a frame, from where it was
      move: async (x, y, ms) => {
        if (capabilities['safari:useSimulator']) throw new Error('the simulator has no pointer to move');
        const steps = Math.max(1, Math.round(ms / 16));
        const [x0, y0] = pointer;
        const actions = Array.from({ length: steps }, (_, i) => ({ type: 'pointerMove', duration: Math.round(ms / steps), origin: 'viewport',
          x: Math.round(x0 + (x - x0) * (i + 1) / steps), y: Math.round(y0 + (y - y0) * (i + 1) / steps) }));
        await wd('POST', `/session/${id}/actions`, { actions: [{ type: 'pointer', id: 'mouse', parameters: { pointerType: 'mouse' }, actions }] });
        pointer = [x, y];
      },
      go: (to) => wd('POST', `/session/${id}/url`, { url: to }),
      size: (width, height) => wd('POST', `/session/${id}/window/rect`, { width, height }),
      shot: () => wd('GET', `/session/${id}/screenshot`),
      close,
    };
  } catch (e) {
    await close();
    throw e;
  }
}
