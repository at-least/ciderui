// node build.mjs <dir>
// Writes the pages run.mjs serves: ciderui compiled as test/components.mjs
// compiles it (every component's sheet), its fonts beside it, and a page per
// console part from its example, each loaded in a frame of a set width by
// index.html (run.mjs reads the frames), the toast a page of its own (it is
// fixed to the window).
import { cpSync, existsSync, mkdirSync, readFileSync, readdirSync, writeFileSync } from 'node:fs';
import { join } from 'node:path';
import { ROOT, compile } from '../../test/compile.mjs';

const OUT = process.argv[2];
if (!OUT) throw new Error('node build.mjs <dir>');
mkdirSync(OUT, { recursive: true });
const names = readdirSync(join(ROOT, 'components')).filter((n) => existsSync(join(ROOT, 'components', n, `${n}.css`)));
const sheets = names.map((n) => `@import "components/${n}/${n}.css";`).join('\n');
writeFileSync(join(OUT, 'style.css'), compile({ css: `@import "fonts.css";\n@import "index.css";\n${sheets}`, sources: ['components'] }));
cpSync(join(ROOT, 'fonts'), join(OUT, 'fonts'), { recursive: true });

const ex = (component, name) => readFileSync(join(ROOT, 'components', component, 'examples', `${name}.html`), 'utf8');
const page = (body, head = '') => `<!doctype html><html lang="en"><head><meta charset="utf-8"><meta name="viewport" content="width=device-width, initial-scale=1"><link rel="stylesheet" href="style.css">${head}</head><body>${body}</body></html>\n`;

// crashcart's syncTabs (its app.js): the bar's overflow state, marked on load
// and on the bar's scroll; sync() is the probe's, after it scrolls the bar.
// The bar's ground a flat blue, the probe's: the fade is read off its pixels
// (run.mjs), a mask fading the ground as it fades the tabs.
const TABS_SCRIPT = `<script>
function sync() {
  document.querySelectorAll('.cui-console-tabs').forEach((bar) => {
    var over = bar.scrollWidth > bar.clientWidth + 1;
    bar.toggleAttribute('data-overflow', over);
    bar.toggleAttribute('data-overflow-end', over && bar.scrollLeft + bar.clientWidth >= bar.scrollWidth - 1);
    bar.toggleAttribute('data-overflow-start', over && bar.scrollLeft <= 1);
  });
}
sync();
document.addEventListener('scroll', sync, true);
</script>`;
writeFileSync(join(OUT, 'tabs.html'), page(`${ex('console-tabs', 'many')}${TABS_SCRIPT}`, '<style>.cui-console-tabs { background-color: rgb(0, 0, 255); }</style>'));

// a third thumbnail 600px wide, wider than a phone's column
const WIDE = `data:image/svg+xml;base64,${Buffer.from('<svg xmlns="http://www.w3.org/2000/svg" width="600" height="240"><rect width="600" height="240" fill="#5b8def"/></svg>').toString('base64')}`;
const thumbs = ex('console-thumb', 'thumbs');
const wide = thumbs.replace(/(\s*<\/div>\s*<\/div>\s*)$/, `\n    <figure class="cui-console-thumb"><img src="${WIDE}" alt=""><figcaption class="cui-console-thumb-caption">wide.png</figcaption></figure>$1`);
if (wide === thumbs) throw new Error('the thumbs example does not end as build.mjs expects');
writeFileSync(join(OUT, 'thumbs.html'), page(wide));
writeFileSync(join(OUT, 'pager.html'), page(ex('console-pager', 'pager')));
writeFileSync(join(OUT, 'content.html'), page(ex('console-content', 'content')));

// the toast in the page itself, at rest: the example's wrapper is the docs
// frame's (a transform holding the fixed toast in its box)
const toast = ex('console-toast', 'toast').match(/<span class="cui-console-toast"[^>]*>[^<]*<\/span>/)[0].replace(' data-shown', '');
writeFileSync(join(OUT, 'toast.html'), page(toast));

// the frames, top left first: tabs is read off the window's pixels
const FRAMES = [
  ['tabs', 320, 60],
  ['thumbs', 375, 300],
  ['thumbs', 1200, 300],
  ['pager', 375, 80],
  ['pager', 1200, 80],
  ['content', 900, 200],
  ['content', 1200, 200],
  ['content', 1920, 200],
];
writeFileSync(join(OUT, 'index.html'), `<!doctype html><html lang="en"><head><meta charset="utf-8"><meta name="viewport" content="width=device-width, initial-scale=1"><style>body { margin: 0; background: white; } iframe { display: block; border: 0; }</style></head><body>${FRAMES.map(([name, w, h]) => `<iframe id="${name}-${w}" src="${name}.html" width="${w}" height="${h}"></iframe>`).join('')}</body></html>\n`);
