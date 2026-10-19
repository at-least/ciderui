// Runs inside https://www.apple.com/: apple.com's webfont cut at apple.com's
// letter-spacing against each system-font name at ciderui's letter-spacing,
// per size and weight. Equal widths mean ciderui's values land on apple.com's text.
async function verify(scale, tracking, weights, sample, families) {
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
      const fams = {};
      for (const family of families) fams[family] = width(family, size, weight, tracking[size]);
      rows.push({ size: Number(size), weight, web: width(`"SF Pro ${cut}"`, size, weight, ls), fams });
    }
  }
  box.remove();
  return { ua: navigator.userAgent, rows };
}
