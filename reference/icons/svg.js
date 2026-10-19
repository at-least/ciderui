// Runs on a blank page: SVG markup drawn on a canvas at `scale` pixels to the
// view box's unit, read as an ink mask; its ink's box in view box units, and
// the ink's runs along given rows and columns (a stroke's thickness where a
// row or column crosses it once). compare(a, b) is the share of ink the two
// have in common (intersection over union).
async function raster(markup, scale) {
  // an image needs the SVG namespace, which markup inline in HTML goes without
  if (!markup.includes('xmlns=')) markup = markup.replace('<svg', '<svg xmlns="http://www.w3.org/2000/svg"');
  const svg = new DOMParser().parseFromString(markup, 'image/svg+xml').documentElement;
  const [, , vw, vh] = svg.getAttribute('viewBox').split(/[\s,]+/).map(Number);
  svg.setAttribute('width', vw * scale);
  svg.setAttribute('height', vh * scale);
  svg.setAttribute('fill', '#000');
  svg.setAttribute('color', '#000');
  const url = URL.createObjectURL(new Blob([new XMLSerializer().serializeToString(svg)], { type: 'image/svg+xml' }));
  const img = new Image();
  img.src = url;
  await img.decode();
  const W = Math.round(vw * scale), H = Math.round(vh * scale);
  const c = document.createElement('canvas');
  c.width = W; c.height = H;
  const ctx = c.getContext('2d');
  ctx.drawImage(img, 0, 0, W, H);
  URL.revokeObjectURL(url);
  const d = ctx.getImageData(0, 0, W, H).data;
  const mask = new Uint8Array(W * H);
  for (let i = 0; i < W * H; i++) mask[i] = d[i * 4 + 3] >= 128 ? 1 : 0;
  let left = W, right = -1, top = H, bottom = -1;
  for (let y = 0; y < H; y++) for (let x = 0; x < W; x++) if (mask[y * W + x]) {
    left = Math.min(left, x); right = Math.max(right, x); top = Math.min(top, y); bottom = Math.max(bottom, y);
  }
  const u = v => +(v / scale).toFixed(3);
  // the runs of ink along a row (y) or a column (x), in view box units
  const runs = (axis, at) => {
    const out = [];
    let start = -1;
    const n = axis === 'row' ? W : H;
    for (let i = 0; i <= n; i++) {
      const on = i < n && (axis === 'row' ? mask[Math.round(at * scale) * W + i] : mask[i * W + Math.round(at * scale)]);
      if (on && start < 0) start = i;
      if (!on && start >= 0) { out.push([u(start), u(i)]); start = -1; }
    }
    return out;
  };
  return { mask, W, H, box: { left: u(left), right: u(right + 1), top: u(top), bottom: u(bottom + 1) }, runs };
}
function compare(a, b) {
  if (a.W !== b.W || a.H !== b.H) throw new Error(`the two are drawn at different sizes: ${a.W}x${a.H}, ${b.W}x${b.H}`);
  let both = 0, either = 0;
  for (let i = 0; i < a.mask.length; i++) { both += a.mask[i] & b.mask[i]; either += a.mask[i] | b.mask[i]; }
  return +(both / either).toFixed(4);
}
