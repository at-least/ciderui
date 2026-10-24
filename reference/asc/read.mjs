// App Store Connect's pages, read in the owner's own Chrome, signed in by
// the owner (no headless browser can sign in): its pages hold the account's
// data, so the read runs in that browser through Claude in Chrome's
// javascript tool, and its result comes back through the tool and is written
// here by hand.
//
//   node read.mjs <spec>          prints the script to run in the page
//                                 (pages/<spec>.mjs's probes in READER)
//   node read.mjs --check <spec>  holds results/<spec>.json to what the page
//                                 handed back and to the account's privacy
//
// READER finds each probe's element (`select` in the document, or in
// `shadow`'s shadow root, among those with a box unless `unseen`, as a
// gradient's stop; with `text`, the one whose text is that, the page's own
// label; `nth` of them; `up` that many parents up), and reads
// `props` of its computed style (`pseudo`'s with it) and, with `box`, its
// box. An element is named by its tag, classes and role alone, never its
// text, link or label; the path's app and team ids are left out. What it
// hands back is the result and the SHA-256 of the result written with its
// keys sorted (as the tool prints it), its bytes as numbers (the tool hides
// a hex digest as encoded data), so --check fails on a value copied wrong;
// the result holds the digests of the probes it was given and of its own
// source, so --check fails on a script pasted wrong too. A probe that
// matches nothing is null and fails --check.
import { readFileSync } from 'node:fs';
import { createHash } from 'node:crypto';
import { fileURLToPath } from 'node:url';

export const READER = async function reader(probes) {
  const canon = (x) => Array.isArray(x) ? `[${x.map(canon).join(',')}]`
    : x && typeof x === 'object' ? `{${Object.keys(x).sort().map((k) => `${JSON.stringify(k)}:${canon(x[k])}`).join(',')}}`
    : JSON.stringify(x);
  const sha = async (s) => [...new Uint8Array(await crypto.subtle.digest('SHA-256', new TextEncoder().encode(s)))];
  const name = (e) => {
    const cls = [...e.classList].slice(0, 3).join(' ');
    const role = e.getAttribute('role');
    return `${e.tagName.toLowerCase()}${cls ? ` {${cls}}` : ''}${role ? `[role=${role}]` : ''}`;
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
};

// TOKENS reads App Store Connect's toolkit's tokens (the custom properties
// its newer look declares on :root, --tb-…) a spec names by `tokens`: those
// whose name begins with one of its prefixes and matches none of its
// excludes, each as the page computes it (getComputedStyle on the root: every
// var() in it substituted), never the declaration as written. Too long for
// the tool's output, the result is handed back as a file the page downloads,
// results/<spec>.json as it is to be written; its digests are READER's.
export const TOKENS = async function tokens(spec, file) {
  const canon = (x) => Array.isArray(x) ? `[${x.map(canon).join(',')}]`
    : x && typeof x === 'object' ? `{${Object.keys(x).sort().map((k) => `${JSON.stringify(k)}:${canon(x[k])}`).join(',')}}`
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
};

const canon = (x) => Array.isArray(x) ? `[${x.map(canon).join(',')}]`
  : x && typeof x === 'object' ? `{${Object.keys(x).sort().map((k) => `${JSON.stringify(k)}:${canon(x[k])}`).join(',')}}`
  : JSON.stringify(x);
const sha = (s) => [...createHash('sha256').update(s).digest()];
const same = (a, b) => JSON.stringify(a) === JSON.stringify(b);

// The command line over root's pages/ and results/: read.mjs's own folder
// when it is run, another site's when that site's folder calls cli with its
// own, so every signed-in site is read and checked by the one READER.
export async function cli(root) {
  const [flag, name] = process.argv[2] === '--check' ? ['--check', process.argv[3]] : [null, process.argv[2]];
  const spec = (await import(new URL(`pages/${name}.mjs`, root))).default;
  if (!flag) {
    if (spec.tokens) console.log(`await (${TOKENS.toString()})(${JSON.stringify(spec.tokens)}, ${JSON.stringify(`${name}.json`)})`);
    else console.log(`await (${READER.toString()})(${JSON.stringify(spec.probes)})`);
    process.exit(0);
  }
  const text = readFileSync(new URL(`results/${name}.json`, root), 'utf8');
  const { sha256, ...result } = JSON.parse(text);
  const fails = [];
  if (!same(sha(canon(result)), sha256)) fails.push('the result is not what the page handed back (its digest differs)');
  const [given, reader, read] = spec.tokens ? [spec.tokens, TOKENS, result.tokens] : [spec.probes, READER, result.probes];
  if (!same(result.probesDigest, sha(canon(given)))) fails.push(`the page was given other ${spec.tokens ? 'tokens' : 'probes'} than pages/${name}.mjs's`);
  if (!same(result.readerDigest, sha(reader.toString()))) fails.push(`the page ran another reader than ${spec.tokens ? 'TOKENS' : 'READER'}`);
  if (result.path !== spec.path) fails.push(`read on ${result.path}, not ${spec.path}`);
  for (const [k, v] of Object.entries(read)) if (v === null || v === '') fails.push(`${k} matched nothing`);
  if (spec.tokens) for (const p of spec.tokens.prefixes) if (!Object.keys(read).some((k) => k.startsWith(p))) fails.push(`no token begins ${p}`);
  // nothing of the account in what was read: no id (six digits or more, not
  // a fraction's; a UUID), no address
  if (/(?<![\d.])\d{6,}|[0-9a-f]{8}-[0-9a-f]{4}-|@/.test(result.path + JSON.stringify(read))) fails.push('it holds what may be the account\'s: an id or an address');
  for (const f of fails) console.log(`FAIL ${name}: ${f}`);
  if (!fails.length) console.log(`ok ${name}: ${Object.keys(read).length} ${spec.tokens ? 'tokens' : 'probes'}, ${result.path}, ${result.browser}, ${result.window.join(' x ')}, ${result.date}`);
  process.exit(fails.length ? 1 : 0);
}

if (process.argv[1] === fileURLToPath(import.meta.url)) await cli(new URL('.', import.meta.url));
