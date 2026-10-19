// Runs on page.html. arm(): counts the clicks the bar's link and the link
// under the curtain get, into window.cuiClicks. state(): the flyout's state
// and place now, the bar's and the curtain's, what a point in the bar and one
// in the curtain's middle hit, the focused element, the root's overflow, the
// page's address, whether the browser has anchor positioning, and open, the
// pointer-events of the flyout's ::backdrop. pixels(png, offset): from a
// picture of the window (base64), where the page starts in it (found from
// the bar's color, or given), the bar's color and the panel's at a point
// each, and the spread of the brightness along a row 100px above the
// window's bottom (the page's stripes, under the curtain when it is open).
function arm() {
  window.cuiClicks = { link: 0, under: 0 };
  for (const id of ['link', 'under']) document.getElementById(id).addEventListener('click', (e) => { e.preventDefault(); window.cuiClicks[id]++; });
  return true;
}
const box = (e) => { const r = e.getBoundingClientRect(); return { top: r.top, left: r.left, width: r.width, height: r.height, bottom: r.bottom }; };
const hit = (x, y) => { const e = document.elementFromPoint(x, y); return e ? e.id || e.className || e.tagName : null; };
function state() {
  const f = document.getElementById('flyout');
  const bar = box(document.getElementById('bar'));
  const link = box(document.getElementById('link'));
  const open = f.matches(':popover-open');
  const curtain = box(document.getElementById('curtain'));
  return {
    anchors: CSS.supports('top: anchor(bottom)') && CSS.supports('anchor-name: --a'),
    open, display: getComputedStyle(f).display, flyout: box(f), bar, curtain,
    panel: box(document.querySelector('.panel')),
    // the window's width less a scrollbar's (Safari on the Mac draws one)
    windowWidth: innerWidth, pageWidth: document.documentElement.clientWidth, windowHeight: innerHeight, scrollY,
    barHit: hit(link.left + link.width / 2, link.top + link.height / 2),
    curtainHit: open && curtain.height ? hit(curtain.left + curtain.width / 2, curtain.top + curtain.height / 2) : null,
    focus: document.activeElement.id || document.activeElement.tagName,
    rootOverflow: getComputedStyle(document.documentElement).overflowY,
    menuDisplay: getComputedStyle(document.getElementById('menu')).display,
    backdropPointer: open ? getComputedStyle(f, '::backdrop').pointerEvents : null,
    clicks: window.cuiClicks ?? null,
    at: location.pathname + location.search,
  };
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
  const bar = document.getElementById('bar').getBoundingClientRect();
  const x = innerWidth - 40;
  // where the page starts in the picture (the simulator's has the phone's
  // status bar above it): from the bar's color at its first row, looked for
  // down that column in a picture with the flyout closed, or given
  if (offset === undefined) {
    offset = null;
    for (let y = 0; y < img.height; y++) {
      const [r, gr, b] = g.getImageData(Math.round(x * scale), y, 1, 1).data;
      if (Math.abs(r - 200) <= 2 && Math.abs(gr - 220) <= 2 && Math.abs(b - 240) <= 2) {
        offset = y / scale - bar.top;
        break;
      }
    }
    if (offset === null) return { scale, offset };
  }
  const at = (x, y) => [...g.getImageData(Math.round(x * scale), Math.round((y + offset) * scale), 1, 1).data.slice(0, 3)];
  const panel = document.querySelector('.panel').getBoundingClientRect();
  const row = g.getImageData(0, Math.round((innerHeight - 100 + offset) * scale), img.width, 1).data;
  const light = [];
  for (let i = 0; i < row.length; i += 4) light.push((row[i] + row[i + 1] + row[i + 2]) / 3);
  const mean = light.reduce((a, b) => a + b, 0) / light.length;
  const spread = Math.sqrt(light.reduce((a, b) => a + (b - mean) ** 2, 0) / light.length);
  return { scale, offset, bar: at(x, bar.top + bar.height / 2), panel: at(innerWidth / 2, panel.top + panel.height - 20), spread: +spread.toFixed(2), mean: +mean.toFixed(1) };
}
