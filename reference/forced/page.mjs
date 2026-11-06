// node page.mjs: page.html, the kit's parts that draw a state or a color in
// a color of their own, on the kit's CSS compiled as a site compiles it
// (../../test/compile.mjs): an option unchecked and checked (the control: its
// state drawn by an edge).
import { readFileSync, writeFileSync } from 'node:fs';
import { compile } from '../../test/compile.mjs';

const NAMES = ['option'];
const css = compile({ css: `@import "index.css";\n${NAMES.map((n) => `@import "components/site/${n}/${n}.css";`).join('\n')}`, sources: ['components'] });
const ex = (n, id) => readFileSync(new URL(`../../components/site/${n}/examples/${id}.html`, import.meta.url), 'utf8').trim();
const body = `<div style="padding: 24px">
<div id="opt" style="display: flex; gap: 16px; width: 700px">${ex('option', 'option')}${ex('option', 'checked')}</div>
</div>`;
writeFileSync(new URL('./page.html', import.meta.url), `<!doctype html><html lang="en"><head><meta charset="utf-8"><style>${css}</style></head><body>${body}</body></html>\n`);
