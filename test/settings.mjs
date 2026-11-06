// The system's settings a page must follow, read off every example as it ships, so
// a new part is held with no line of its own.

import { readFileSync } from 'node:fs';
import { createRequire } from 'node:module';
import { COMPONENTS, exampleFiles, exampleHtml, shipped } from './compile.mjs';
import { open, quit } from './browser.mjs';
import { check, done } from './check.mjs';

const AXE = readFileSync(createRequire(import.meta.url).resolve('axe-core/axe.min.js'), 'utf8');

const emulate = async (page, name, value) => {
  const cdp = await page.page.context().newCDPSession(page.page);
  await cdp.send('Emulation.setEmulatedMedia', { features: [{ name, value }] });
  check(`${name}: ${value} is on (the control)`, await page.page.evaluate((q) => matchMedia(q).matches, `(${name}: ${value})`), true);
};

const show = (page, part, file) => page.page.evaluate((html) => { document.body.innerHTML = html; }, exampleHtml(part, file));

// a color as this page draws it
const colorOf = (page, value) => page.page.evaluate((v) => {
  const probe = document.createElement('i');
  probe.style.color = v;
  document.body.append(probe);
  const color = getComputedStyle(probe).color;
  probe.remove();
  return color;
}, value);

// Reduce Transparency: nothing blurs what is behind it, no ground lets it show.
// A tint or a dim is not a ground: under this alpha it colors what is behind on
// purpose. A mark too small to hold text is not one either. A field is near white
// on a page it is the same color as.
const TINT = 0.5;
const MARK = 24;
const FIELD_GROUND = ['--cui-c-field'];
for (const set of ['site', 'console']) {
  const page = await open(shipped(set), '');
  await emulate(page, 'prefers-reduced-transparency', 'reduce');
  const allowed = await Promise.all(FIELD_GROUND.map((name) => colorOf(page, `var(${name})`)));
  let read = 0;
  for (const width of [1200, 390]) {
    await page.resize(width);
    for (const part of COMPONENTS.filter((c) => c.set === set)) {
      for (const file of exampleFiles(part)) {
        await show(page, part, file);
        const found = await page.page.evaluate(([tint, mark, allowed]) => [...document.body.querySelectorAll('*')].flatMap((e) => [null, '::before', '::after'].flatMap((pseudo) => {
          const cs = getComputedStyle(e, pseudo);
          const name = `${e.tagName.toLowerCase()}.${e.className}${pseudo ?? ''}`;
          const alpha = Number(cs.backgroundColor.match(/, ([\d.]+)\)$/)?.[1] ?? cs.backgroundColor.match(/\/ ([\d.]+)\)$/)?.[1] ?? 1);
          return [
            ...(cs.backdropFilter !== 'none' ? [`${name} blurs`] : []),
            ...(alpha > tint && alpha < 1 && e.getBoundingClientRect().height >= mark && !allowed.includes(cs.backgroundColor) ? [`${name} on ${cs.backgroundColor}`] : []),
          ];
        })), [TINT, MARK, allowed]);
        read++;
        check(`${part.name}/${file} at ${width}px: nothing shows through under Reduce Transparency`, found.join(' | '), '');
      }
    }
  }
  check(`${set}: the examples read (${read})`, read > 0, true);
  await page.close();
}

// Increase Contrast: secondary text and edges at the text's color...
const CONTRAST = {
  site: { '--cui-c-text': ['--cui-c-text-2', '--cui-c-text-3', '--cui-c-border', '--cui-c-divider'], '--cui-c-on-bar-hover': ['--cui-c-on-bar'], '--cui-c-on-footer-title': ['--cui-c-on-footer', '--cui-c-on-footer-link'] },
  console: {
    '--cui-c-console-neutral-label-primary': ['--cui-c-console-neutral-label-secondary', '--cui-c-console-field-border', '--cui-c-console-neutral-border-secondary', '--cui-c-console-dialog-stroke'],
    '--cui-c-on-video': ['--cui-c-on-video-time'],
  },
};
for (const [set, slots] of Object.entries(CONTRAST)) {
  const pairs = Object.entries(slots).flatMap(([strong, names]) => names.map((name) => [name, strong]));
  const page = await open(shipped(set), '');
  for (const [name, strong] of pairs) check(`${name} is not ${strong} at rest (the control)`, await colorOf(page, `var(${name})`) !== await colorOf(page, `var(${strong})`), true);
  await emulate(page, 'prefers-contrast', 'more');
  for (const [name, strong] of pairs) check(`${name} is ${strong} under Increase Contrast`, await colorOf(page, `var(${name})`), await colorOf(page, `var(${strong})`));
  await page.close();
}

// ...and so no neutral text is under AAA's 7:1 in any example; a hue, on the text
// or under it, is an accent's, a link's or a status's, held to 4.5:1 (tokens.css)
const channels = (hex) => [1, 3, 5].map((i) => parseInt(hex.slice(i, i + 2), 16));
const hue = (hex) => Math.max(...channels(hex)) - Math.min(...channels(hex)) > 16;
for (const part of COMPONENTS) {
  for (const file of exampleFiles(part)) {
    const page = await open(shipped(part.set), exampleHtml(part, file));
    await emulate(page, 'prefers-contrast', 'more');
    await page.page.addScriptTag({ content: AXE });
    const found = await page.page.evaluate(async () => {
      const r = await axe.run(document, { runOnly: { type: 'rule', values: ['color-contrast-enhanced'] }, resultTypes: ['violations'] });
      return r.violations.flatMap((v) => v.nodes.map((n) => ({ target: n.target.join(' '), fg: n.any[0].data.fgColor, bg: n.any[0].data.bgColor })));
    });
    await page.close();
    check(`${part.name}/${file}: its neutral text 7:1 under Increase Contrast`, found.filter((f) => !hue(f.fg) && !hue(f.bg)).map((f) => `${f.target} ${f.fg} on ${f.bg}`).join(' | '), '');
  }
}

// Appearance: the console follows the system, the site set stays light; a class holds it
const ground = async (set, rootClass, colorScheme) => {
  const page = await open(shipped(set), '<i></i>', { colorScheme });
  await page.page.evaluate((c) => { if (c) document.documentElement.classList.add(c); }, rootClass);
  const color = (await page.read('body', ['background-color']))['background-color'];
  await page.close();
  return color;
};
const WHITE = 'rgb(255, 255, 255)';
const BLACK = 'rgb(0, 0, 0)';
check('the console, the system dark: dark', await ground('console', '', 'dark'), BLACK);
check('the console, the system light: light', await ground('console', '', 'light'), WHITE);
check('the console with cui-light, the system dark: light', await ground('console', 'cui-light', 'dark'), WHITE);
check('the console with cui-dark, the system light: dark', await ground('console', 'cui-dark', 'light'), BLACK);
check('the site set, the system dark: light', await ground('site', '', 'dark'), WHITE);
check('the site set with cui-dark, the system light: dark', await ground('site', 'cui-dark', 'light'), BLACK);

await quit();
done('settings');
