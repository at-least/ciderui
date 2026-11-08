// @invariant: node page.mjs: page.html, the accordion's example on the kit's CSS,
// compiled as a site compiles it (../../test/compile.mjs), and ring.html,
// for the focus ring in an answer: 24px in from the window's edge, the example
// as an answer of itself, and the example with a link beginning its first
// answer.
import { readFileSync, writeFileSync } from 'node:fs';
import { compile } from '../../test/compile.mjs';
const css = compile({ css: '@import "index.css";\n@import "components/site/accordion/accordion.css";', sources: ['components'] });
const example = (id) => readFileSync(new URL(`../../components/site/accordion/examples/${id}.html`, import.meta.url), 'utf8').trim();
const page = (body, style = '') => `<!doctype html><html lang="en"><head><meta charset="utf-8"><meta name="viewport" content="width=device-width"><style>${css}</style>${style && `<style>${style}</style>`}</head><body>${body}</body></html>\n`;
const html = example('accordion');
writeFileSync(new URL('./page.html', import.meta.url), page(html));
const FIRST = '<p>Add your address at checkout to see a delivery date.</p>';
// @invariant: the example's items render closed; the page opens the first, as a page wanting one open writes it
const opened = html.replace('<details class="cui-accordion-item">', '<details class="cui-accordion-item" open>');
const nested = opened.replace(FIRST, html);
const linked = opened.replace(FIRST, '<p><a href="#store">Visit a store</a> to collect your order.</p>');
if (opened === html || nested === opened || linked === opened) throw new Error('the examples changed: ring.html is not what it says');
writeFileSync(new URL('./ring.html', import.meta.url), page(`<main style="padding: 24px"><div id="nested">${nested}</div><div id="linked">${linked}</div></main>`));
