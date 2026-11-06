// node page.mjs: page.html, the kit's parts that draw a state or a color in
// a color of their own, on the kit's CSS compiled as a site compiles it
// (../../test/compile.mjs): an option unchecked and checked (the control: its
// state drawn by an edge) and the empty select four times: as it is (f1), its
// empty choice's text "" (f2), its label hidden (f3), both (f4), so that
// "Select"'s ink is f3 against f4 and with the label shown f1 against f2.
import { readFileSync, writeFileSync } from 'node:fs';
import { compile } from '../../test/compile.mjs';

const NAMES = ['option', 'field'];
const css = compile({ css: `@import "index.css";\n${NAMES.map((n) => `@import "components/site/${n}/${n}.css";`).join('\n')}`, sources: ['components'] });
const ex = (n, id) => readFileSync(new URL(`../../components/site/${n}/examples/${id}.html`, import.meta.url), 'utf8').trim();
const swap = (s, a, b) => { if (!s.includes(a)) throw new Error(`not in the example: ${a}`); return s.replaceAll(a, b); };
const select = ex('field', 'select');
const blank = (s) => swap(s, '<option value="" selected>Select</option>', '<option value="" selected></option>');
const unlabeled = (s) => swap(s, 'class="cui-field-label"', 'class="cui-field-label" style="visibility: hidden"');
const named = (s, n) => swap(s, '"model"', `"model-${n}"`);
const body = `<div style="padding: 24px">
<div id="opt" style="display: flex; gap: 16px; width: 700px">${ex('option', 'option')}${ex('option', 'checked')}</div>
<div style="width: 400px; margin-top: 24px"><div id="f1">${select}</div><div id="f2" style="margin-top: 8px">${named(blank(select), 2)}</div><div id="f3" style="margin-top: 8px">${named(unlabeled(select), 3)}</div><div id="f4" style="margin-top: 8px">${named(unlabeled(blank(select)), 4)}</div></div>
</div>`;
writeFileSync(new URL('./page.html', import.meta.url), `<!doctype html><html lang="en"><head><meta charset="utf-8"><style>${css}</style></head><body>${body}</body></html>\n`);
