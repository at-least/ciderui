// Runs on page.html. state(): the list's place and look now, its first row's
// opacity, the root's overflow, each of the close button's lines' box in its
// SVG, the focused element. Each toggle of the list starts sample(): its
// height, its first row's opacity and the close button's first line's box
// every frame for 700ms, into window.cuiSamples.
const linesOf = () => {
  const svg = document.querySelector('#close svg').getBoundingClientRect();
  return [...document.querySelectorAll('#close path')].map(p => {
    const b = p.getBoundingClientRect();
    return [b.left - svg.left, b.top - svg.top, b.width, b.height].map(v => +v.toFixed(3));
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
