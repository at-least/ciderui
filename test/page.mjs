import { compile, ownSheet, shipped } from './compile.mjs';
import { open, quit } from './browser.mjs';
import { ref } from './reference.mjs';
import { check, done } from './check.mjs';

const css = compile({ css: '@import "site.css";', classes: 'dark:underline outline-none md:underline lg:underline xl:underline text-secondary bg-secondary text-alt bg-alt' });
const html = `<p id="text">Text <strong id="strong">bold</strong></p><button id="button">Buy</button>
<section id="dark" class="cui-dark"><p id="in-dark" class="dark:underline">in a dark block</p>
  <div id="light" class="cui-light"><p id="in-light" class="dark:underline">light inside dark</p></div><p id="secondary-dark" class="text-secondary">secondary</p></section>
<p id="outside" class="dark:underline">outside</p><p id="moving" style="transition: color 1s 0.5s">moving</p><button id="plain" class="outline-none">Plain</button>
<p id="md" class="md:underline">md</p><p id="lg" class="lg:underline">lg</p><p id="xl" class="xl:underline">xl</p><p id="secondary" class="text-secondary">secondary</p>
<style>@keyframes cui-test { to { opacity: 0.5; } }</style><p id="animated" style="animation: cui-test 1s infinite">animated</p><div id="smooth" style="scroll-behavior: smooth">smooth</div>`;
const LIGHT = { color: 'rgb(29, 29, 31)', 'background-color': 'rgb(255, 255, 255)' };
const DARK = { color: 'rgb(245, 245, 247)', 'background-color': 'rgb(0, 0, 0)' };
const W = 1200;

let b = await open(css, html);
try {
  const body = await b.read('body', ['font-size', 'line-height', 'font-weight', 'color', 'background-color', '-webkit-font-smoothing', 'font-synthesis', 'text-rendering', 'letter-spacing', 'font-family']);
  for (const prop of ['font-size', 'line-height', 'font-weight', 'color', 'background-color', '-webkit-font-smoothing', 'font-synthesis', 'text-rendering']) {
    check(`body ${prop} is apple.com's home:page.body`, body[prop], ref('home', 'page.body', W, prop));
  }
  // @external: Firefox's, on macOS: Chromium drops it, so the stylesheet is read
  check('body declares -moz-osx-font-smoothing grayscale, as apple.com\'s stylesheets do', /(^|[}\s])body\{[^}]*-moz-osx-font-smoothing:grayscale/.test(css), true);
  check('body letter-spacing is the 17px fitted value', body['letter-spacing'], `${17 * 0.0029}px`, { tolerance: 0.001 });
  check('body font-family is the stack', body['font-family'], '-apple-system, BlinkMacSystemFont, Inter, "Helvetica Neue", Helvetica, Arial, "PingFang TC", "Noto Sans TC", "Noto Sans CJK TC", "Microsoft JhengHei", sans-serif');
  check('strong is apple.com\'s store:page.strong weight', (await b.read('#strong', ['font-weight']))['font-weight'], ref('store', 'page.strong', W, 'font-weight'));

  const ring = await b.read('#button', ['outline-style', 'outline-color', 'outline-width', 'outline-offset'], { state: 'focus-visible' });
  for (const prop of Object.keys(ring)) check(`the focus ring's ${prop} is apple.com's home:button:focus-visible`, ring[prop], ref('home', 'button', W, prop, 'focus-visible'));
  check('no ring at rest', (await b.read('#button', ['outline-style']))['outline-style'], 'none');
  check('a utility on the element wins over base.css: outline-none drops the ring', (await b.read('#plain', ['outline-style'], { state: 'focus-visible' }))['outline-style'], 'none');

  for (const [id, want] of [['text', LIGHT], ['dark', DARK], ['in-dark', DARK], ['light', LIGHT], ['in-light', LIGHT]]) {
    const got = await b.read(`#${id}`, ['color', 'background-color']);
    check(`#${id} color`, got.color, want.color);
    // @external: a paragraph paints no ground of its own: only the blocks do
    if (!id.startsWith('in-') && id !== 'text') check(`#${id} background-color`, got['background-color'], want['background-color']);
  }
  check('dark: applies in a dark block', (await b.read('#in-dark', ['text-decoration-line']))['text-decoration-line'], 'underline');
  check('dark: does not apply outside', (await b.read('#outside', ['text-decoration-line']))['text-decoration-line'], 'none');
  check('dark: applies in a light block inside a dark one too: for what no token carries', (await b.read('#in-light', ['text-decoration-line']))['text-decoration-line'], 'underline');
  check('text-secondary is --cui-c-text-2, light', (await b.read('#secondary', ['color'])).color, 'rgb(110, 110, 115)');
  check('text-secondary follows a dark block', (await b.read('#secondary-dark', ['color'])).color, 'rgb(134, 134, 139)');
  check('text- and bg- names apart: of bg-alt, bg-secondary and text-alt, bg-alt alone a utility', ['bg-alt', 'bg-secondary', 'text-alt'].filter(u => css.includes(`.${u}{`)).join(' '), 'bg-alt');
  check('a transition moves without the preference', (await b.read('#moving', ['transition-duration']))['transition-duration'], '1s');

  await b.page.evaluate(() => document.documentElement.classList.add('cui-dark'));
  const dark = await b.read('body', ['color', 'background-color']);
  check('a dark page\'s body color', dark.color, DARK.color);
  check('a dark page\'s body background-color', dark['background-color'], DARK['background-color']);

  for (const [id, below, at] of [['md', 734, 735], ['lg', 1068, 1069], ['xl', 1440, 1441]]) {
    await b.resize(below);
    check(`${id}: not at ${below}px`, (await b.read(`#${id}`, ['text-decoration-line']))['text-decoration-line'], 'none');
    await b.resize(at);
    check(`${id}: from ${at}px, apple.com's`, (await b.read(`#${id}`, ['text-decoration-line']))['text-decoration-line'], 'underline');
  }
} finally {
  await b.close();
}

b = await open(css, html, { colorScheme: 'dark', reducedMotion: 'reduce' });
try {
  check('the OS\'s dark appearance leaves the page light', (await b.read('body', ['color']))['color'], LIGHT.color);
  check('dark: does not follow the OS\'s appearance', (await b.read('#outside', ['text-decoration-line']))['text-decoration-line'], 'none');
  const moving = await b.read('#moving', ['transition-duration', 'transition-delay']);
  check('reduced motion stills a transition', moving['transition-duration'], '0s');
  check('reduced motion drops a transition\'s delay', moving['transition-delay'], '0s');
  const animated = await b.read('#animated', ['animation-duration', 'animation-delay', 'animation-iteration-count']);
  for (const [prop, want] of [['animation-duration', '0.001s'], ['animation-delay', '-0.001s'], ['animation-iteration-count', '1']]) {
    check(`reduced motion ends an animation at once, once: ${prop}`, animated[prop], want);
  }
  check('reduced motion scrolls at once', (await b.read('#smooth', ['scroll-behavior']))['scroll-behavior'], 'auto');
} finally {
  await b.close();
}

// @invariant: what a site does (README.md): its own utilities' sheet, built once, linked
// after the set's dist sheet; each name draws the set's color, a spinner keeps
// turning under Reduce Motion
const OWN_CLASSES = 'text-body text-secondary bg-page bg-alt border border-t border-border border-divider text-link animate-spin';
const own = ownSheet('../utilities.css', OWN_CLASSES);
const SET_COLORS = [
  ['body', 'text-body', 'color', { site: '--cui-c-text', console: '--cui-c-console-neutral-label-primary' }],
  ['secondary', 'text-secondary', 'color', { site: '--cui-c-text-2', console: '--cui-c-console-neutral-label-secondary' }],
  ['page', 'bg-page', 'background-color', { site: '--cui-c-bg', console: '--cui-c-console-neutral-background-primary' }],
  ['alt', 'bg-alt', 'background-color', { site: '--cui-c-bg-alt', console: '--cui-c-console-card' }],
  ['edge', 'border border-border', 'border-top-color', { site: '--cui-c-border', console: '--cui-c-console-field-border' }],
  ['bare', 'border', 'border-top-color', { site: '--cui-c-divider', console: '--cui-c-console-neutral-border-tertiary' }],
  ['top', 'border-t', 'border-top-color', { site: '--cui-c-divider', console: '--cui-c-console-neutral-border-tertiary' }],
  ['divider', 'border border-divider', 'border-top-color', { site: '--cui-c-divider', console: '--cui-c-console-neutral-border-tertiary' }],
  ['link', 'text-link', 'color', { site: '--cui-c-link', console: '--cui-c-console-brand-label-primary' }],
];
const drawn = {};
for (const set of ['site', 'console']) {
  const html = SET_COLORS.map(([id, classes, prop, token]) => `<p id="${id}" class="${classes}">a</p><p id="${id}-want" style="${prop}: var(${token[set]}); border-style: solid">a</p>`).join('')
    + '<span id="spinner" class="animate-spin">a</span><span id="still" style="animation: spin 1s infinite">a</span>';
  b = await open(shipped(set), html, { reducedMotion: 'reduce' });
  try {
    await b.page.addStyleTag({ content: own });
    drawn[set] = {};
    for (const [id, classes, prop, token] of SET_COLORS) {
      const got = (await b.read(`#${id}`, [prop]))[prop];
      drawn[set][id] = got;
      check(`${set}: a site's own ${classes} draws ${token[set]}`, got, (await b.read(`#${id}-want`, [prop]))[prop]);
    }
    const spinner = await b.read('#spinner', ['animation-name', 'animation-duration', 'animation-iteration-count']);
    check(`${set}: under reduced motion animate-spin still turns, its own duration`, spinner['animation-duration'], '1s');
    check(`${set}: under reduced motion animate-spin still turns, without end`, spinner['animation-iteration-count'], 'infinite');
    check(`${set}: under reduced motion another animation still ends at once`, (await b.read('#still', ['animation-duration']))['animation-duration'], '0.001s');
  } finally {
    await b.close();
  }
}
// @trap: a set whose names drew the other set's colors would pass the checks above were
// the two the same
check('the sets draw text-secondary apart', drawn.site.secondary !== drawn.console.secondary, true);
check('the sets draw a bare border apart', drawn.site.bare !== drawn.console.bare, true);

await quit();
done('page');
