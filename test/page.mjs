// base.css and theme.css on a page: the body apple.com's (reference/apple.com/,
// home:page.body) but for its letter-spacing and font, ciderui's; bold
// apple.com's 600; the focus ring apple.com's button's; a dark block dark, a
// light one inside it light again, a dark page dark; the dark variant on a
// dark block and never on the OS's appearance; reduced motion stilling a
// transition, and a <details>' content's (its ::details-content, a
// pseudo-element `*` does not reach: the accordion's).

import { compile } from './compile.mjs';
import { open, quit } from './browser.mjs';
import { ref } from './reference.mjs';
import { check, done } from './check.mjs';

const css = compile({ classes: 'dark:underline' });
const html = `<p id="text">Text <strong id="strong">bold</strong></p><button id="button">Buy</button>
<section id="dark" class="cui-dark"><p id="in-dark" class="dark:underline">in a dark block</p>
  <div id="light" class="cui-light"><p id="in-light">light inside dark</p></div></section>
<p id="outside" class="dark:underline">outside</p><p id="moving" style="transition: color 1s">moving</p>
<style>#opening::details-content { transition: height 1s; }</style><details id="opening"><summary>a question</summary><p>its answer</p></details>`;
const LIGHT = { color: 'rgb(29, 29, 31)', 'background-color': 'rgb(255, 255, 255)' };
const DARK = { color: 'rgb(245, 245, 247)', 'background-color': 'rgb(0, 0, 0)' };
const W = 1200;

let b = await open(css, html);
try {
  const body = await b.read('body', ['font-size', 'line-height', 'font-weight', 'color', 'background-color', '-webkit-font-smoothing', 'font-synthesis', 'text-rendering', 'letter-spacing', 'font-family']);
  for (const prop of ['font-size', 'line-height', 'font-weight', 'color', 'background-color', '-webkit-font-smoothing', 'font-synthesis', 'text-rendering']) {
    check(`body ${prop} is apple.com's home:page.body`, body[prop], ref('home', 'page.body', W, prop));
  }
  // Firefox's, on macOS: Chromium drops it, so the stylesheet is read
  check('body declares -moz-osx-font-smoothing grayscale, as apple.com\'s stylesheets do', /(^|[}\s])body\{[^}]*-moz-osx-font-smoothing:grayscale/.test(css), true);
  check('body letter-spacing is the 17px fitted value', body['letter-spacing'], `${17 * 0.0029}px`, { tolerance: 0.001 });
  check('body font-family is the stack', body['font-family'], '-apple-system, BlinkMacSystemFont, Inter, "Helvetica Neue", Helvetica, Arial, sans-serif');
  check('strong is apple.com\'s store:page.strong weight', (await b.read('#strong', ['font-weight']))['font-weight'], ref('store', 'page.strong', W, 'font-weight'));

  const ring = await b.read('#button', ['outline-style', 'outline-color', 'outline-width', 'outline-offset'], { state: 'focus-visible' });
  for (const prop of Object.keys(ring)) check(`the focus ring's ${prop} is apple.com's home:button:focus-visible`, ring[prop], ref('home', 'button', W, prop, 'focus-visible'));
  check('no ring at rest', (await b.read('#button', ['outline-style']))['outline-style'], 'none');

  for (const [id, want] of [['text', LIGHT], ['dark', DARK], ['in-dark', DARK], ['light', LIGHT], ['in-light', LIGHT]]) {
    const got = await b.read(`#${id}`, ['color', 'background-color']);
    check(`#${id} color`, got.color, want.color);
    // a paragraph paints no ground of its own: only the blocks do
    if (!id.startsWith('in-') && id !== 'text') check(`#${id} background-color`, got['background-color'], want['background-color']);
  }
  check('dark: applies in a dark block', (await b.read('#in-dark', ['text-decoration-line']))['text-decoration-line'], 'underline');
  check('dark: does not apply outside', (await b.read('#outside', ['text-decoration-line']))['text-decoration-line'], 'none');
  check('a transition moves without the preference', (await b.read('#moving', ['transition-duration']))['transition-duration'], '1s');
  check('a <details>\' content moves without the preference', (await b.read('#opening', ['transition-duration'], { pseudo: '::details-content' }))['transition-duration'], '1s');

  await b.page.evaluate(() => document.documentElement.classList.add('cui-dark'));
  const dark = await b.read('body', ['color', 'background-color']);
  check('a dark page\'s body color', dark.color, DARK.color);
  check('a dark page\'s body background-color', dark['background-color'], DARK['background-color']);
} finally {
  await b.close();
}

// the OS's appearance changes nothing; reduced motion stills a transition
b = await open(css, html, { colorScheme: 'dark', reducedMotion: 'reduce' });
try {
  check('the OS\'s dark appearance leaves the page light', (await b.read('body', ['color']))['color'], LIGHT.color);
  check('dark: does not follow the OS\'s appearance', (await b.read('#outside', ['text-decoration-line']))['text-decoration-line'], 'none');
  check('reduced motion stills a transition', (await b.read('#moving', ['transition-duration']))['transition-duration'], '0s');
  check('reduced motion stills a <details>\' content', (await b.read('#opening', ['transition-duration'], { pseudo: '::details-content' }))['transition-duration'], '0s');
} finally {
  await b.close();
}

await quit();
done('page');
