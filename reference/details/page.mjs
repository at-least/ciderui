// node page.mjs: page.html, the accordion's example on the kit's CSS,
// compiled as a site compiles it (../../test/compile.mjs), and ring.html,
// for the focus ring in an answer: 24px in from the window's edge, the bag's
// structure (the compact example as the large question's answer, opened, as
// the components' nested case) and the compact example with a link
// beginning its first answer; held.html, its first item alone, 24px in, its
// open answer held 10px tall (its text 13.6px down, below what is shown),
// and held-visible.html, the same with nothing clipped (overflow: visible),
// the control.
import { readFileSync, writeFileSync } from 'node:fs';
import { compile } from '../../test/compile.mjs';
const css = compile({ css: '@import "index.css";\n@import "components/accordion/accordion.css";', sources: ['components'] });
const example = (id) => readFileSync(new URL(`../../components/accordion/examples/${id}.html`, import.meta.url), 'utf8').trim();
const page = (body, style = '') => `<!doctype html><html lang="en"><head><meta charset="utf-8"><meta name="viewport" content="width=device-width"><style>${css}</style>${style && `<style>${style}</style>`}</head><body>${body}</body></html>\n`;
const html = example('accordion');
writeFileSync(new URL('./page.html', import.meta.url), page(html));
const nested = example('large').replace('<details class="cui-accordion-item">', '<details class="cui-accordion-item" open>').replace('<p>What people ask most, answered.</p>', html);
const linked = html.replace('<p>Add your address at checkout to see a delivery date.</p>', '<p><a href="#store">Visit a store</a> to collect your order.</p>');
if (nested === example('large') || linked === html) throw new Error('the examples changed: ring.html is not what it says');
writeFileSync(new URL('./ring.html', import.meta.url), page(`<main style="padding: 24px">${nested}${linked}</main>`));
const first = html.match(/<details class="cui-accordion-item" open>[\s\S]*?<\/details>/)[0];
const hold = '.cui-accordion-item[open]::details-content { height: 10px; transition: none; }';
writeFileSync(new URL('./held.html', import.meta.url), page(`<main style="padding: 24px"><div class="cui-accordion">${first}</div></main>`, hold));
writeFileSync(new URL('./held-visible.html', import.meta.url), page(`<main style="padding: 24px"><div class="cui-accordion">${first}</div></main>`, `${hold} .cui-accordion-item::details-content { overflow: visible; }`));
// held-clip.html: the same clipping both its axes (overflow: clip, the kit's
// rule before it clipped its height alone), to tell a browser that clips
// neither form from one that clips one
writeFileSync(new URL('./held-clip.html', import.meta.url), page(`<main style="padding: 24px"><div class="cui-accordion">${first}</div></main>`, `${hold} .cui-accordion-item::details-content { overflow: clip; }`));
