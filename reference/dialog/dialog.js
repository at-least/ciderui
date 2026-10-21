// Runs on page.html. state(): the dialog's state and place now, the root's
// overflow, what a point above the panel hits (in its 40px of room), the focused element, and
// whether the browser has the command attribute. watch(): from now, the
// dialog's opacity, its backdrop's and whether it is open, every frame for
// 600ms, into window.cuiSamples (started just before a click).
function state() {
  const d = document.querySelector('#modal');
  const r = d.getBoundingClientRect();
  const cs = getComputedStyle(d);
  const hit = d.open ? document.elementFromPoint(Math.round(innerWidth / 2), 10) : null;
  return {
    open: d.open, modal: d.matches(':modal'), position: cs.position,
    top: r.top, left: r.left, width: r.width, height: r.height, windowWidth: innerWidth, windowHeight: innerHeight,
    backdrop: getComputedStyle(d, '::backdrop').backgroundColor,
    scrolls: d.open ? d.scrollHeight > d.clientHeight && cs.overflowY === 'auto' : null,
    rootOverflow: getComputedStyle(document.documentElement).overflowY,
    hit: hit ? hit.id || hit.className || hit.tagName : null,
    focus: document.activeElement.id || document.activeElement.tagName,
    command: 'command' in HTMLButtonElement.prototype,
  };
}
function watch() {
  const run = { frames: [] };
  (window.cuiSamples ??= []).push(run);
  const t0 = performance.now();
  const frame = () => {
    const d = document.querySelector('#modal');
    run.frames.push([Math.round(performance.now() - t0), +getComputedStyle(d).opacity, +getComputedStyle(d, '::backdrop').opacity, d.open]);
    if (performance.now() - t0 < 600) requestAnimationFrame(frame);
  };
  frame();
  return true;
}

// pageState(): the marketing modal's state, its backdrop's filter, its close
// button's place in the window, the root's overflow. watchPage(): from now,
// its backdrop's filter and color every frame for 1s, into window.cuiPage.
// pixels(png, offset): from a picture of the window (base64), where the page
// starts in it (found from the band's color, or given), the band's color at
// a point, and the spread and mean of the brightness along a row 100px above
// the window's bottom (the page's stripes, under the scrim when it is open).
function pageState() {
  const d = document.querySelector('#page');
  const b = getComputedStyle(d, '::backdrop');
  return {
    open: d.open, modal: d.matches(':modal'), filter: b.backdropFilter, backdrop: b.backgroundColor,
    close: Math.round(document.querySelector('#page-close').getBoundingClientRect().top),
    panel: Math.round(document.querySelector('#page .panel').getBoundingClientRect().top),
    below: Math.round(innerHeight - document.querySelector('#page .panel').getBoundingClientRect().bottom),
    rootOverflow: getComputedStyle(document.documentElement).overflowY,
    focus: document.activeElement.id || document.activeElement.tagName,
  };
}
function watchPage() {
  window.cuiPage = [];
  const t0 = performance.now();
  const frame = () => {
    const b = getComputedStyle(document.querySelector('#page'), '::backdrop');
    window.cuiPage.push([Math.round(performance.now() - t0), b.backdropFilter, b.backgroundColor]);
    if (performance.now() - t0 < 1000) requestAnimationFrame(frame);
  };
  frame();
  return true;
}
async function pixels(png, offset) {
  const img = new Image();
  img.src = `data:image/png;base64,${png}`;
  await img.decode();
  const scale = img.width / innerWidth;
  const c = document.createElement('canvas');
  c.width = img.width;
  c.height = img.height;
  const g = c.getContext('2d', { willReadFrequently: true });
  g.drawImage(img, 0, 0);
  const band = document.getElementById('band').getBoundingClientRect();
  const x = innerWidth / 2;
  if (offset === undefined) {
    offset = null;
    for (let y = 0; y < img.height; y++) {
      const [r, gr, b] = g.getImageData(Math.round(x * scale), y, 1, 1).data;
      if (Math.abs(r - 200) <= 2 && Math.abs(gr - 220) <= 2 && Math.abs(b - 240) <= 2) {
        offset = y / scale - band.top;
        break;
      }
    }
    if (offset === null) return { scale, offset };
  }
  const row = g.getImageData(0, Math.round((innerHeight - 100 + offset) * scale), img.width, 1).data;
  const light = [];
  for (let i = 0; i < row.length; i += 4) light.push((row[i] + row[i + 1] + row[i + 2]) / 3);
  const mean = light.reduce((a, b) => a + b, 0) / light.length;
  const spread = Math.sqrt(light.reduce((a, b) => a + (b - mean) ** 2, 0) / light.length);
  return { scale, offset, spread: +spread.toFixed(2), mean: +mean.toFixed(1) };
}
