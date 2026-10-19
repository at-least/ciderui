// Runs in page.html: each element's computed color and ground.
async function scheme() {
  const read = id => { const cs = getComputedStyle(document.getElementById(id)); return [cs.color, cs.backgroundColor]; };
  return { ua: navigator.userAgent, got: Object.fromEntries(['page', 'dark', 'in-dark', 'back'].map(id => [id, read(id)])) };
}
