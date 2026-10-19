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
