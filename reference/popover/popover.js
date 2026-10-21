// Runs on page.html. state(): the list's place and look now, its first row's
// opacity, the root's overflow, each of the close button's lines' box in its
// SVG, the focused element. Each toggle of the list starts sample(): its
// height, its first row's opacity and the close button's first line's box
// every frame for 700ms, into window.cuiSamples.
const linesOf = () => {
  // the line's geometry box (getBBox: the same in every browser, the stroke
  // excluded) mapped through its CSS transform about its origin. The probe
  // reads no getBoundingClientRect: that box includes the stroke and, on a
  // stroked <path>, Firefox pads it by about twice the stroke-width a side
  // where Chromium gives the bare geometry, so it reads differently per
  // browser (the svg's user space is 1:1 with its px: viewBox 0 0 18 18 at
  // 18 by 18, as the claim's numbers are).
  return [...document.querySelectorAll('#close path')].map(p => {
    const cs = getComputedStyle(p);
    const m = cs.transform === 'none' ? new DOMMatrixReadOnly() : new DOMMatrixReadOnly(cs.transform);
    const o = cs.transformOrigin.split(' ').map(parseFloat);
    const b = p.getBBox();
    const xs = [], ys = [];
    for (const [x, y] of [[b.x, b.y], [b.x + b.width, b.y], [b.x, b.y + b.height], [b.x + b.width, b.y + b.height]]) {
      const qx = x - o[0], qy = y - o[1];
      xs.push(m.a * qx + m.c * qy + m.e + o[0]);
      ys.push(m.b * qx + m.d * qy + m.f + o[1]);
    }
    return [Math.min(...xs), Math.min(...ys), Math.max(...xs) - Math.min(...xs), Math.max(...ys) - Math.min(...ys)].map(v => +v.toFixed(3));
  });
};
function state() {
  const list = document.querySelector('#list');
  const r = list.getBoundingClientRect();
  const lines = list.matches(':popover-open') ? linesOf() : null;
  const cs = getComputedStyle(list);
  return {
    open: list.matches(':popover-open'), display: cs.display, position: cs.position,
    height: +r.height.toFixed(2), width: +r.width.toFixed(2), top: r.top, windowHeight: innerHeight,
    background: cs.backgroundColor, rowOpacity: +getComputedStyle(document.querySelector('#list > .row')).opacity,
    rootOverflow: getComputedStyle(document.documentElement).overflowY, lines,
    focus: document.activeElement.id || document.activeElement.tagName,
  };
}
function arm() {
  window.cuiSamples = [];
  document.querySelector('#list').addEventListener('toggle', (e) => {
    const run = { to: e.newState, frames: [] };
    window.cuiSamples.push(run);
    const t0 = performance.now();
    const frame = () => {
      const list = document.querySelector('#list');
      const open = list.matches(':popover-open');
      run.frames.push([Math.round(performance.now() - t0), +list.getBoundingClientRect().height.toFixed(1), +getComputedStyle(document.querySelector('#list > .row')).opacity, open ? linesOf()[0][3] : null]);
      if (performance.now() - t0 < 700) requestAnimationFrame(frame);
    };
    frame();
  });
}
