// Runs inside https://www.apple.com/: Inter (ciderui's font where the system
// font is not SF), added from its bytes, at ciderui's letter-spacing and at 0,
// against apple.com's cut at apple.com's letter-spacing and at 0.
async function inter(scale, tracking, weights, sample, fontBase64) {
  const bytes = Uint8Array.from(atob(fontBase64), c => c.charCodeAt(0));
  const face = new FontFace('cui-probe-inter', bytes.buffer, { weight: '100 900', style: 'normal' });
  await face.load();
  if (face.status !== 'loaded') throw new Error(`Inter not loaded: ${face.status}`);
  document.fonts.add(face);
  for (const family of ['SF Pro Text', 'SF Pro Display']) {
    for (const weight of weights) {
      for (const size of Object.keys(scale)) {
        const loaded = await document.fonts.load(`${weight} ${size}px "${family}"`, sample);
        if (!loaded.length || loaded.some(f => f.status !== 'loaded')) {
          throw new Error(`webfont not loaded: ${weight} ${size}px "${family}" gave ${loaded.length} faces`);
        }
      }
    }
  }
  const box = document.createElement('div');
  box.style.cssText = 'position:absolute;left:0;top:0;visibility:hidden';
  document.body.append(box);
  const width = (family, size, weight, ls) => {
    const span = document.createElement('span');
    span.textContent = sample;
    span.style.cssText = `font-family:${family};font-size:${size}px;font-weight:${weight};font-style:normal;`
      + `letter-spacing:${ls}em;word-spacing:0;white-space:nowrap;text-transform:none;font-variant:normal;`
      + 'font-optical-sizing:auto;font-kerning:auto;font-feature-settings:normal;'
      + 'font-variation-settings:normal;text-rendering:auto;font-synthesis:none';
    box.append(span);
    const w = span.getBoundingClientRect().width;
    span.remove();
    return w;
  };
  const rows = [];
  for (const [size, [cut, ls]] of Object.entries(scale)) {
    for (const weight of weights) {
      rows.push({
        size: Number(size), weight,
        web: width(`"SF Pro ${cut}"`, size, weight, ls),
        web0: width(`"SF Pro ${cut}"`, size, weight, 0),
        inter: width('cui-probe-inter', size, weight, tracking[size]),
        inter0: width('cui-probe-inter', size, weight, 0),
      });
    }
  }
  box.remove();
  return { ua: navigator.userAgent, rows };
}
