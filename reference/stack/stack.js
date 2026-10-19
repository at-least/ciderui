// Runs in page.html, the kit's fonts.css and ciderui's stack: which fonts the
// page fetched by the time it had drawn the stack, Inter's faces' states then,
// and the stack's width against system-ui's and Inter's own (Inter is loaded
// on purpose only after the first two are read).
async function stack() {
  await new Promise(r => requestAnimationFrame(() => requestAnimationFrame(r)));
  await new Promise(r => setTimeout(r, 1000));
  const fetched = () => performance.getEntriesByType('resource').map(e => e.name)
    .filter(n => n.endsWith('.woff2')).map(n => n.slice(n.lastIndexOf('/') + 1));
  const before = {
    fetched: fetched(),
    inter: [...document.fonts].filter(f => f.family.replace(/["']/g, '') === 'Inter').map(f => f.status),
  };
  const sample = document.getElementById('stack').textContent;
  const width = (family) => {
    const span = document.createElement('span');
    span.style.fontFamily = family;
    span.textContent = sample;
    document.body.append(span);
    const w = span.getBoundingClientRect().width;
    span.remove();
    return w;
  };
  const stackWidth = document.getElementById('stack').getBoundingClientRect().width;
  const systemWidth = width('system-ui');
  const loaded = await document.fonts.load('400 17px Inter', sample);
  if (!loaded.length || loaded.some(f => f.status !== 'loaded')) throw new Error(`Inter not loaded: ${loaded.length} faces`);
  const interWidth = width('Inter');
  return { ua: navigator.userAgent, before, stackWidth, systemWidth, interWidth, after: { fetched: fetched() } };
}
