// @invariant: node build.mjs <dir>
// Writes the pages run.mjs serves: ciderui compiled as test/components.mjs
// compiles it (every component's sheet), its fonts beside it, and a page per
// console part from its example, each loaded in a frame of a set width by
// index.html (run.mjs reads the frames).
import { cpSync, mkdirSync, readFileSync, writeFileSync } from 'node:fs';
import { join } from 'node:path';
import { COMPONENTS, ROOT, compile, componentDir } from '../../test/compile.mjs';

const OUT = process.argv[2];
if (!OUT) throw new Error('node build.mjs <dir>');
mkdirSync(OUT, { recursive: true });
const sheets = COMPONENTS.map((c) => `@import "${c.dir}/${c.name}.css";`).join('\n');
writeFileSync(join(OUT, 'style.css'), compile({ css: `@import "fonts.css";\n@import "index.css";\n${sheets}`, sources: ['components'] }));
cpSync(join(ROOT, 'fonts'), join(OUT, 'fonts'), { recursive: true });

const ex = (component, name) => readFileSync(join(ROOT, componentDir(component), 'examples', `${name}.html`), 'utf8');
const page = (body, head = '', scheme = '') => `<!doctype html><html lang="en"${scheme ? ` class="${scheme}"` : ''}><head><meta charset="utf-8"><meta name="viewport" content="width=device-width, initial-scale=1"><link rel="stylesheet" href="style.css">${head}</head><body>${body}</body></html>\n`;

// @external: crashcart's syncTabs (its app.js): the bar's overflow state, marked on load
// and on the bar's scroll; sync() is the probe's, after it scrolls the bar.
// The bar's ground a flat blue, the probe's: the fade is read off its pixels
// (run.mjs), a mask fading the ground as it fades the tabs. A page per scheme,
// cui-light and cui-dark on <html>: the page behind the bar white or black
// whatever the system's appearance (the console's root follows it).
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
const TABS_STYLE = '<style>.cui-console-tabs { background-color: rgb(0, 0, 255); }</style>';
writeFileSync(join(OUT, 'tabs.html'), page(`${ex('console-tabs', 'many')}${TABS_SCRIPT}`, TABS_STYLE, 'cui-light'));
writeFileSync(join(OUT, 'tabs-dark.html'), page(`${ex('console-tabs', 'many')}${TABS_SCRIPT}`, TABS_STYLE, 'cui-dark'));

writeFileSync(join(OUT, 'content.html'), page(ex('console-content', 'content')));

// @invariant: the frames, top left first: tabs and tabs-dark are read off the window's pixels
const FRAMES = [
  ['tabs', 320, 60],
  ['tabs-dark', 320, 60],
  ['content', 900, 200],
  ['content', 1200, 200],
  ['content', 1920, 200],
];
writeFileSync(join(OUT, 'index.html'), `<!doctype html><html lang="en"><head><meta charset="utf-8"><meta name="viewport" content="width=device-width, initial-scale=1"><style>body { margin: 0; background: white; } iframe { display: block; border: 0; }</style></head><body>${FRAMES.map(([name, w, h]) => `<iframe id="${name}-${w}" src="${name}.html" width="${w}" height="${h}"></iframe>`).join('')}</body></html>\n`);
