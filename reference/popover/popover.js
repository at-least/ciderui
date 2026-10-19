// Runs on page.html. state(): the list's place and look now, its first row's
// opacity, the root's overflow, each of the button's lines' box in its SVG.
// Each toggle of the list starts sample(): its height and its first row's
// opacity every frame for 700ms, into window.cuiSamples.
function state() {
  const list = document.querySelector('#list');
  const r = list.getBoundingClientRect();
  const svg = document.querySelector('#button svg').getBoundingClientRect();
  const lines = [...document.querySelectorAll('#button path')].map(p => {
    const b = p.getBoundingClientRect();
    return [b.left - svg.left, b.top - svg.top, b.width, b.height].map(v => +v.toFixed(3));
  });
  const cs = getComputedStyle(list);
  return {
    open: list.matches(':popover-open'), display: cs.display, position: cs.position,
    height: +r.height.toFixed(2), width: +r.width.toFixed(2), top: r.top, windowHeight: innerHeight,
    background: cs.backgroundColor, rowOpacity: +getComputedStyle(document.querySelector('#list > .row')).opacity,
    rootOverflow: getComputedStyle(document.documentElement).overflowY, lines,
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
      run.frames.push([Math.round(performance.now() - t0), +list.getBoundingClientRect().height.toFixed(1), +getComputedStyle(document.querySelector('#list > .row')).opacity]);
      if (performance.now() - t0 < 700) requestAnimationFrame(frame);
    };
    frame();
  });
}
