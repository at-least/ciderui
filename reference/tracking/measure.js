// Runs inside https://www.apple.com/, where apple.com's own "SF Pro Text" and
// "SF Pro Display" webfonts are already declared: the width of one sample
// string in each webfont and in system-ui, letter-spacing 0, per size and weight.
async function measure(sizes, weights, sample) {
  // not document.fonts.ready: on apple.com it never settles in Safari 27 (some
  // of its icon faces stay loading); each face measured is loaded explicitly
  const faces = {};
  for (const family of ['SF Pro Text', 'SF Pro Display']) {
    for (const weight of weights) {
      for (const size of sizes) {
        const loaded = await document.fonts.load(`${weight} ${size}px "${family}"`, sample);
        if (!loaded.length || loaded.some(f => f.status !== 'loaded')) {
          throw new Error(`webfont not loaded: ${weight} ${size}px "${family}" gave ${loaded.length} faces`);
        }
        faces[`${family} ${weight}`] = loaded.map(f => `${f.family} ${f.weight} ${f.status}`).join(', ');
      }
    }
  }
  const box = document.createElement('div');
  box.style.cssText = 'position:absolute;left:0;top:0;visibility:hidden';
  document.body.append(box);
  const rows = [];
  for (const size of sizes) {
    for (const weight of weights) {
      const row = { size, weight };
      for (const [key, family] of [['text', '"SF Pro Text"'], ['display', '"SF Pro Display"'], ['sys', 'system-ui']]) {
        const span = document.createElement('span');
        span.textContent = sample;
        span.style.cssText = `font-family:${family};font-size:${size}px;font-weight:${weight};font-style:normal;`
          + 'letter-spacing:0;word-spacing:0;white-space:nowrap;text-transform:none;font-variant:normal;'
          + 'font-optical-sizing:auto;font-kerning:auto;font-feature-settings:normal;'
          + 'font-variation-settings:normal;text-rendering:auto;font-synthesis:none';
        box.append(span);
        row[key] = span.getBoundingClientRect().width;
        span.remove();
      }
      rows.push(row);
    }
  }
  box.remove();
  return { ua: navigator.userAgent, faces, rows };
}
