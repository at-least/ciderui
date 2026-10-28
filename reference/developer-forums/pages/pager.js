// Run in the page (Claude in Chrome's javascript tool) on
// https://developer.apple.com/forums/latest?page=2; it leaves the record in
// window.__pager, read back in pieces (the tool's output is cut at about
// 1000 characters) and written by hand to ../results/pager.json, the
// values equal to a property's default left out.
const P = ['display', 'justify-content', 'align-items', 'gap', 'font-family', 'font-size', 'line-height', 'font-weight', 'letter-spacing', 'color', 'background-color', 'border-top-width', 'border-top-left-radius', 'padding-top', 'padding-right', 'padding-bottom', 'padding-left', 'margin-top', 'margin-left', 'margin-right', 'width', 'height'];
const read = (e) => {
  if (!e) return null;
  const cs = getComputedStyle(e);
  const r = e.getBoundingClientRect();
  const o = { box: { x: +r.left.toFixed(2), y: +r.top.toFixed(2), w: +r.width.toFixed(2), h: +r.height.toFixed(2) } };
  P.forEach((k) => (o[k] = cs.getPropertyValue(k)));
  return o;
};
const p = document.querySelector('nav.pagination');
const prev = p.querySelector('.previous-page');
const next = p.querySelector('.next-page');
window.__pager = JSON.stringify({
  read: new Date().toISOString(),
  window: [innerWidth, innerHeight, devicePixelRatio],
  browser: navigator.userAgent.match(/Chrome\/[\d.]+/)[0] + ' ' + navigator.platform,
  'pager.nav': Object.assign(read(p), { role: p.getAttribute('role'), label: p.getAttribute('aria-label') }),
  'pager.list': read(p.querySelector('ol')),
  'pager.previous': Object.assign(read(prev), { label: prev.getAttribute('aria-label') }),
  'pager.page': Object.assign(read(p.querySelector('.page-number')), { text: p.querySelector('.page-number').innerText.trim() }),
  'pager.next': Object.assign(read(next), { label: next.getAttribute('aria-label') }),
  'pager.next.icon': read(next.querySelector('svg,span,i')),
  'pager.first': read(p.querySelector('.first-page')),
});
