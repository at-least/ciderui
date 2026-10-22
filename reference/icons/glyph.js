// Runs on an apple.com page: the "more" chevron, the character apple.com
// sets with its SF Pro Icons font in `.icon-after.more::after`, drawn on a
// canvas at a large size and read as numbers in em: its advance, its ink's
// box against the origin and the baseline, its strokes' thickness. Then a
// chevron of ciderui's own drawn from those numbers alone (two strokes from
// the ink's corners to its middle, round caps and joins), and the share of
// ink the two have in common (intersection over union); and, given
// `shipped` ({ points: [[x, y], …], stroke }, in em from the origin and above
// the baseline), the chevron the kit ships drawn the same way, and its share
// (shippedIou).
async function glyph(size, shipped) {
  const span = document.createElement('span');
  span.className = 'icon icon-after more';
  document.querySelector('main').append(span);
  const content = getComputedStyle(span, '::after').content;
  const family = getComputedStyle(span, '::after').fontFamily;
  span.remove();
  const ch = JSON.parse(content.split(' / ')[0]);
  const font = `${size}px ${family}`;
  const loaded = await document.fonts.load(font, ch);
  if (!loaded.length) throw new Error(`no face loaded for ${font}`);

  const W = size * 2, H = size * 2, X = size / 2, BASE = size * 1.4;
  const mask = (draw) => {
    const c = document.createElement('canvas');
    c.width = W; c.height = H;
    const ctx = c.getContext('2d');
    draw(ctx);
    const d = ctx.getImageData(0, 0, W, H).data;
    const m = new Uint8Array(W * H);
    for (let i = 0; i < W * H; i++) m[i] = d[i * 4 + 3] >= 128 ? 1 : 0;
    return m;
  };
  const apple = mask(ctx => { ctx.font = font; ctx.fillStyle = '#000'; ctx.fillText(ch, X, BASE); });
  let left = W, right = -1, top = H, bottom = -1;
  for (let y = 0; y < H; y++) for (let x = 0; x < W; x++) if (apple[y * W + x]) {
    left = Math.min(left, x); right = Math.max(right, x); top = Math.min(top, y); bottom = Math.max(bottom, y);
  }
  if (right < 0) throw new Error('the glyph drew no ink');
  const ctx = document.createElement('canvas').getContext('2d');
  ctx.font = font;
  const advance = ctx.measureText(ch).width;
  // thickness: the ink's horizontal run on the row a quarter down, through the
  // upper arm, times the arm's slope's sine
  const row = Math.round(top + (bottom - top) / 4);
  let run = 0;
  for (let x = left; x <= right; x++) run += apple[row * W + x];
  const midY = (top + bottom) / 2;
  const angle = Math.atan2(midY - top, right - left);
  const thickness = run * Math.sin(angle);
  // ciderui's chevron from those numbers: a polyline inset by half the
  // thickness, round caps and joins
  const t = thickness, inset = t / 2;
  const P = { x0: left + 0.5 + inset, y0: top + 0.5 + inset, x1: right + 0.5 - inset, y1: midY + 0.5, y2: bottom + 0.5 - inset };
  const ours = mask(ctx => {
    ctx.lineWidth = t; ctx.lineCap = 'round'; ctx.lineJoin = 'round'; ctx.strokeStyle = '#000';
    ctx.beginPath(); ctx.moveTo(P.x0, P.y0); ctx.lineTo(P.x1, P.y1); ctx.lineTo(P.x0, P.y2); ctx.stroke();
  });
  const share = (m) => {
    let both = 0, either = 0;
    for (let i = 0; i < W * H; i++) { both += apple[i] & m[i]; either += apple[i] | m[i]; }
    return +(both / either).toFixed(4);
  };
  const kit = shipped && mask(ctx => {
    ctx.lineWidth = shipped.stroke * size; ctx.lineCap = 'round'; ctx.lineJoin = 'round'; ctx.strokeStyle = '#000';
    ctx.beginPath();
    shipped.points.forEach(([x, y], i) => ctx[i ? 'lineTo' : 'moveTo'](X + x * size, BASE - y * size));
    ctx.stroke();
  });
  const em = v => +(v / size).toFixed(4);
  return {
    ua: navigator.userAgent, char: `U+${ch.codePointAt(0).toString(16).toUpperCase()}`, family, size,
    advance: em(advance),
    ink: { left: em(left - X), right: em(right + 1 - X), top: em(BASE - top), bottom: em(BASE - bottom - 1) },
    thickness: em(thickness),
    polyline: { x0: em(P.x0 - X), y0: em(BASE - P.y0), x1: em(P.x1 - X), y1: em(BASE - P.y1), y2: em(BASE - P.y2), stroke: em(t) },
    iou: share(ours),
    ...(kit && { shippedIou: share(kit) }),
  };
}
