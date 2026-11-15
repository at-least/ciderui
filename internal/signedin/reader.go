package signedin

// Reader is the in-page reader of a spec's probes, its source as the page's
// reader.toString() gives it. It finds each probe's element (`select` in the
// document, or in `shadow`'s shadow root, among those with a box unless `unseen`,
// as a gradient's stop; with `text`, the one whose text is that, the page's own
// label; `nth` of them; `up` that many parents up), and reads `props` of its
// computed style (`pseudo`'s with it) and, with `box`, its box. An element is
// named by its tag, classes and role alone, never its text, link or label; the
// path's app and team ids are left out. What it hands back is the result and the
// SHA-256 of the result written with its keys sorted (as the tool prints it), its
// bytes as numbers (the tool hides a hex digest as encoded data), so --check fails
// on a value copied wrong; the result holds the digests of the probes it was given
// and of its own source, so --check fails on a script pasted wrong too. A probe
// that matches nothing is null and fails --check.
//
// @invariant: byte for byte the text the records' readerDigest is the SHA-256
// of: a raw string holds no backtick, so each is spliced in
const Reader = `async function reader(probes) {
  const canon = (x) => Array.isArray(x) ? ` + "`" + `[${x.map(canon).join(',')}]` + "`" + `
    : x && typeof x === 'object' ? ` + "`" + `{${Object.keys(x).sort().map((k) => ` + "`" + `${JSON.stringify(k)}:${canon(x[k])}` + "`" + `).join(',')}}` + "`" + `
    : JSON.stringify(x);
  const sha = async (s) => [...new Uint8Array(await crypto.subtle.digest('SHA-256', new TextEncoder().encode(s)))];
  const name = (e) => {
    const cls = [...e.classList].slice(0, 3).join(' ');
    const role = e.getAttribute('role');
    return ` + "`" + `${e.tagName.toLowerCase()}${cls ? ` + "`" + ` {${cls}}` + "`" + ` : ''}${role ? ` + "`" + `[role=${role}]` + "`" + ` : ''}` + "`" + `;
  };
  const find = (p) => {
    const root = p.shadow ? document.querySelector(p.shadow)?.shadowRoot : document;
    if (!root) return null;
    let all = [...root.querySelectorAll(p.select)].filter((e) => p.unseen || e.getClientRects().length);
    if (p.text !== undefined) all = all.filter((e) => e.textContent.trim() === p.text);
    let el = all[p.nth ?? 0] ?? null;
    for (let i = 0; el && i < (p.up ?? 0); i++) el = el.parentElement;
    return el;
  };
  const read = {};
  for (const p of probes) {
    const el = find(p);
    if (!el) { read[p.name] = null; continue; }
    const cs = getComputedStyle(el, p.pseudo ?? null);
    const styles = {};
    for (const k of p.props) styles[k] = cs.getPropertyValue(k);
    const r = el.getBoundingClientRect();
    read[p.name] = { html: name(el), styles, ...(p.box ? { box: { width: +r.width.toFixed(3), height: +r.height.toFixed(3) } } : {}) };
  }
  const result = {
    path: location.pathname.replace(/\d{6,}/g, '<app>').replace(/[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}/g, '<team>'),
    window: [innerWidth, innerHeight, devicePixelRatio],
    browser: navigator.userAgent.match(/Chrome\/[\d.]+/)[0],
    date: new Date().toISOString(),
    probes: read,
    probesDigest: await sha(canon(probes)),
    readerDigest: await sha(reader.toString()),
  };
  return { result, sha256: await sha(canon(result)) };
}`

// Tokens reads App Store Connect's toolkit's tokens (the custom properties its newer
// look declares on :root, --tb-…) a spec names by Tokens: those whose name begins
// with one of its prefixes and matches none of its excludes, each as the page
// computes it (getComputedStyle on the root: every var() in it substituted), never
// the declaration as written. Too long for the tool's output, the result is handed
// back as a file the page downloads, <spec>.json, which --save checks and writes as
// results/<spec>.go; its digests are Reader's.
//
// @invariant: byte for byte the text the tokens record's readerDigest is the SHA-256
// of, its backticks spliced in as Reader's
const Tokens = `async function tokens(spec, file) {
  const canon = (x) => Array.isArray(x) ? ` + "`" + `[${x.map(canon).join(',')}]` + "`" + `
    : x && typeof x === 'object' ? ` + "`" + `{${Object.keys(x).sort().map((k) => ` + "`" + `${JSON.stringify(k)}:${canon(x[k])}` + "`" + `).join(',')}}` + "`" + `
    : JSON.stringify(x);
  const sha = async (s) => [...new Uint8Array(await crypto.subtle.digest('SHA-256', new TextEncoder().encode(s)))];
  const names = new Set();
  const walk = (rules) => {
    for (const r of rules) {
      if (r.cssRules) walk(r.cssRules);
      if (r.selectorText === ':root') for (const p of r.style) if (p.startsWith('--tb-')) names.add(p);
    }
  };
  for (const s of document.styleSheets) walk(s.cssRules);
  const excludes = spec.exclude.map((e) => new RegExp(e));
  const root = getComputedStyle(document.documentElement);
  const read = {};
  for (const n of [...names].sort()) {
    if (!spec.prefixes.some((p) => n.startsWith(p)) || excludes.some((e) => e.test(n))) continue;
    read[n] = root.getPropertyValue(n).trim();
  }
  const result = {
    path: location.pathname.replace(/\d{6,}/g, '<app>').replace(/[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}/g, '<team>'),
    window: [innerWidth, innerHeight, devicePixelRatio],
    browser: navigator.userAgent.match(/Chrome\/[\d.]+/)[0],
    date: new Date().toISOString(),
    tokens: read,
    probesDigest: await sha(canon(spec)),
    readerDigest: await sha(tokens.toString()),
  };
  const text = JSON.stringify({ ...result, sha256: await sha(canon(result)) }, null, 1) + '\n';
  const a = document.createElement('a');
  a.href = URL.createObjectURL(new Blob([text], { type: 'application/json' }));
  a.download = file;
  document.body.append(a);
  a.click();
  a.remove();
  return { count: Object.keys(read).length, bytes: text.length };
}`
