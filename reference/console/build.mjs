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
const page = (body) => `<!doctype html><html lang="en"><head><meta charset="utf-8"><meta name="viewport" content="width=device-width, initial-scale=1"><link rel="stylesheet" href="style.css"></head><body>${body}</body></html>\n`;

writeFileSync(join(OUT, 'tabs.html'), page(ex('console-tabs', 'many')));
writeFileSync(join(OUT, 'content.html'), page(ex('console-content', 'content')));

const FRAMES = [
  ['tabs', 320, 60],
  ['content', 900, 200],
  ['content', 1200, 200],
  ['content', 1920, 200],
];
writeFileSync(join(OUT, 'index.html'), `<!doctype html><html lang="en"><head><meta charset="utf-8"><meta name="viewport" content="width=device-width, initial-scale=1"><style>body { margin: 0; background: white; } iframe { display: block; border: 0; }</style></head><body>${FRAMES.map(([name, w, h]) => `<iframe id="${name}-${w}" src="${name}.html" width="${w}" height="${h}"></iframe>`).join('')}</body></html>\n`);
