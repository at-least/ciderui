// node page.mjs: page.html, the accordion's example on the kit's CSS,
// compiled as a site compiles it (../../test/compile.mjs).
import { readFileSync, writeFileSync } from 'node:fs';
import { compile } from '../../test/compile.mjs';
const css = compile({ css: '@import "index.css";\n@import "components/accordion/accordion.css";', sources: ['components'] });
const html = readFileSync(new URL('../../components/accordion/examples/accordion.html', import.meta.url), 'utf8');
writeFileSync(new URL('./page.html', import.meta.url), `<!doctype html><html lang="en"><head><meta charset="utf-8"><meta name="viewport" content="width=device-width"><style>${css}</style></head><body>${html}</body></html>\n`);
