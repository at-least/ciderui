import { compile, ownSheet, shipped } from './compile.mjs';
import { open, quit } from './browser.mjs';
import { ref } from './reference.mjs';
import { check, done } from './check.mjs';

const css = compile({ css: '@import "site.css";', classes: 'dark:underline outline-none md:underline lg:underline xl:underline bg-secondary text-alt bg-alt font-sans' });
const html = `<p id="text">Text <strong id="strong">bold</strong> <b id="b">bold</b></p><p id="sans" class="font-sans">sans</p><button id="button">Buy</button>
<section id="dark" class="cui-dark dark:underline"><p id="in-dark" class="dark:underline">in a dark block</p>
  <div id="light" class="cui-light"><p id="in-light" class="dark:underline">light inside dark</p></div></section>
<p id="outside" class="dark:underline">outside</p><p id="moving" style="transition: color 1s 0.5s">moving</p><button id="plain" class="outline-none">Plain</button>
<p id="md" class="md:underline">md</p><p id="lg" class="lg:underline">lg</p><p id="xl" class="xl:underline">xl</p>
<style>@keyframes cui-test { to { opacity: 0.5; } }</style><p id="animated" style="animation: cui-test 1s infinite">animated</p><div id="smooth" style="scroll-behavior: smooth">smooth</div>
<div id="fixed" style="background-attachment: fixed">fixed</div><style>#pseudo::before { content: ""; transition: color 1s 0.5s; animation: cui-test 1s infinite; }</style><p id="pseudo">pseudo</p>`;
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
  const STACK = '-apple-system, BlinkMacSystemFont, Inter, "Helvetica Neue", Helvetica, Arial, "PingFang TC", "Noto Sans TC", "Noto Sans CJK TC", "Microsoft JhengHei", sans-serif';
  check('body font-family is the stack', body['font-family'], STACK);
  check('font-sans is the stack, not Tailwind\'s', (await b.read('#sans', ['font-family']))['font-family'], STACK);
  for (const id of ['strong', 'b']) check(`${id} is apple.com's store:page.strong weight`, (await b.read(`#${id}`, ['font-weight']))['font-weight'], ref('store', 'page.strong', W, 'font-weight'));

  const ring = await b.read('#button', ['outline-style', 'outline-color', 'outline-width', 'outline-offset'], { state: 'focus-visible' });
  for (const prop of Object.keys(ring)) check(`the focus ring's ${prop} is apple.com's home:button:focus-visible`, ring[prop], ref('home', 'button', W, prop, 'focus-visible'));
  check('no ring at rest', (await b.read('#button', ['outline-style']))['outline-style'], 'none');
  check('a utility on the element wins over base.css: outline-none drops the ring', (await b.read('#plain', ['outline-style'], { state: 'focus-visible' }))['outline-style'], 'none');

  for (const [id, what, scheme] of [['text', 'text on the page', 'light'], ['dark', 'a cui-dark block', 'dark'], ['in-dark', 'text in a cui-dark block', 'dark'], ['light', 'a cui-light block inside a dark one', 'light'], ['in-light', 'text in a cui-light block inside a dark one', 'light']]) {
    const want = scheme === 'dark' ? DARK : LIGHT;
    const got = await b.read(`#${id}`, ['color', 'background-color']);
    check(`${what}: its color ${scheme}`, got.color, want.color);
    // @external: a paragraph paints no ground of its own: only the blocks do
    if (!id.startsWith('in-') && id !== 'text') check(`${what}: its ground ${scheme}`, got['background-color'], want['background-color']);
  }
  for (const [id, what, want] of [['dark', 'applies to the dark block itself', 'underline'], ['in-dark', 'applies in a dark block', 'underline'], ['outside', 'does not apply outside', 'none'], ['in-light', 'applies in a light block inside a dark one too: for what no token carries', 'underline']]) {
    check(`dark: ${what}`, (await b.read(`#${id}`, ['text-decoration-line']))['text-decoration-line'], want);
  }
  check('text- and bg- names apart: of bg-alt, bg-secondary and text-alt, bg-alt alone a utility', ['bg-alt', 'bg-secondary', 'text-alt'].filter(u => css.includes(`.${u}{`)).join(' '), 'bg-alt');
  check('a transition moves without the preference', (await b.read('#moving', ['transition-duration']))['transition-duration'], '1s');

  const BREAKPOINTS = [['md', 734, 735], ['lg', 1068, 1069], ['xl', 1440, 1441]];
  check('the kit\'s breakpoints alone: no sm: or 2xl:', ownSheet('utilities.css', 'sm:underline md:underline lg:underline xl:underline 2xl:underline').match(/min-width:[^)]+/g).join(' '), BREAKPOINTS.map(([, , at]) => `min-width:${at}px`).join(' '));
  for (const [id, below, at] of BREAKPOINTS) {
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
  check('dark: does not follow the OS\'s appearance', (await b.read('#outside', ['text-decoration-line']))['text-decoration-line'], 'none');
  for (const [moving, animated, pseudo, what] of [['#moving', '#animated', undefined, 'an element'], ['#pseudo', '#pseudo', '::before', 'a pseudo-element']]) {
    const transition = await b.read(moving, ['transition-duration', 'transition-delay'], { pseudo });
    check(`reduced motion stills ${what}'s transition`, transition['transition-duration'], '0s');
    check(`reduced motion drops ${what}'s transition delay`, transition['transition-delay'], '0s');
    const animation = await b.read(animated, ['animation-duration', 'animation-delay', 'animation-iteration-count'], { pseudo });
    for (const [prop, want] of [['animation-duration', '0.001s'], ['animation-delay', '-0.001s'], ['animation-iteration-count', '1']]) {
      check(`reduced motion ends ${what}'s animation at once, once: ${prop}`, animation[prop], want);
    }
  }
  check('reduced motion scrolls at once', (await b.read('#smooth', ['scroll-behavior']))['scroll-behavior'], 'auto');
  check('reduced motion scrolls a fixed background with its page', (await b.read('#fixed', ['background-attachment']))['background-attachment'], 'scroll');
} finally {
  await b.close();
}

// @invariant: a site's setup as README.md tells it: its own utilities' sheet, built once,
// linked after the set's dist sheet
const OWN_CLASSES = 'text-body text-secondary bg-page bg-alt border border-t border-border border-divider text-link animate-spin';
const own = ownSheet('utilities.css', OWN_CLASSES);
const DIVIDER = { site: '--cui-c-divider', console: '--cui-c-console-neutral-border-tertiary' };
const SET_COLORS = [
  ['body', 'text-body', 'color', { site: '--cui-c-text', console: '--cui-c-console-neutral-label-primary' }],
  ['secondary', 'text-secondary', 'color', { site: '--cui-c-text-2', console: '--cui-c-console-neutral-label-secondary' }],
  ['page', 'bg-page', 'background-color', { site: '--cui-c-bg', console: '--cui-c-console-neutral-background-primary' }],
  ['alt', 'bg-alt', 'background-color', { site: '--cui-c-bg-alt', console: '--cui-c-console-card' }],
  ['edge', 'border border-border', 'border-top-color', { site: '--cui-c-border', console: '--cui-c-console-field-border' }],
  ['bare', 'border', 'border-top-color', DIVIDER],
  ['top', 'border-t', 'border-top-color', DIVIDER],
  ['divider', 'border border-divider', 'border-top-color', DIVIDER],
  ['link', 'text-link', 'color', { site: '--cui-c-link', console: '--cui-c-console-brand-label-primary' }],
];
const drawn = {};
for (const set of ['site', 'console']) {
  const html = SET_COLORS.map(([id, classes, prop, token]) => `<p id="${id}" class="${classes}">a</p><p id="${id}-want" style="${prop}: var(${token[set]}); border-style: solid">a</p>`).join('')
    + '<style>#bare-pseudo::before, #bare-pseudo::after { content: ""; display: block; border-top-width: 1px; border-top-style: solid; }</style><p id="bare-pseudo">a</p>'
    + '<span id="spinner" class="animate-spin">a</span><span id="still" style="animation: spin 1s infinite">a</span>';
  b = await open(shipped(set), html, { reducedMotion: 'reduce' });
  try {
    await b.page.addStyleTag({ content: own });
    drawn[set] = {};
    // @trap: dark too: the console's link and accent share a light side
    for (const scheme of ['light', 'dark']) {
      if (scheme === 'dark') await b.page.evaluate(() => document.documentElement.classList.add('cui-dark'));
      const want = {};
      for (const [id, classes, prop, token] of SET_COLORS) {
        const got = (await b.read(`#${id}`, [prop]))[prop];
        want[id] = (await b.read(`#${id}-want`, [prop]))[prop];
        drawn[set][id] ??= got;
        check(`${set}, ${scheme}: a site's own ${classes} draws ${token[set]}`, got, want[id]);
      }
      for (const pseudo of ['::before', '::after']) {
        check(`${set}, ${scheme}: a bare border on ${pseudo} draws ${DIVIDER[set]} too`, (await b.read('#bare-pseudo', ['border-top-color'], { pseudo }))['border-top-color'], want.bare);
      }
    }
    const spinner = await b.read('#spinner', ['animation-duration', 'animation-iteration-count']);
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
