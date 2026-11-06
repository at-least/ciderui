// What a site links: dist/<set>.css as it ships, and its own utilities' stylesheet
// after it.

import { mkdtempSync, rmSync, writeFileSync } from 'node:fs';
import { join } from 'node:path';
import { ROOT, UTILITIES_ENTRY, shipped, tailwind } from './compile.mjs';
import { open, quit } from './browser.mjs';
import { check, done } from './check.mjs';

const dir = mkdtempSync(join(ROOT, 'assets-'));
let own;
try {
  writeFileSync(join(dir, 'entry.css'), `${UTILITIES_ENTRY('../utilities.css')}\n@source inline("px-8 md:px-8");\n`);
  own = tailwind(join(dir, 'entry.css'));
} finally {
  rmSync(dir, { recursive: true });
}

const site = await open(shipped('site'), `<h1 id="headline" class="cui-headline">a</h1><a id="plain" class="cui-btn">a</a><a id="own" class="cui-btn px-8">a</a><a id="fixed" class="cui-btn md:px-8">a</a>`);
await site.page.addStyleTag({ content: own });
check('every class of the type ships: cui-headline', (await site.read('#headline', ['font-size']))['font-size'], '48px');
check('a part ships: cui-btn', (await site.read('#plain', ['padding-left']))['padding-left'], '15px');
check('a utility of the second stylesheet beats a part', (await site.read('#own', ['padding-left']))['padding-left'], '32px');
check('a breakpoint of the second stylesheet is the kit\'s', (await site.read('#fixed', ['padding-left']))['padding-left'], '32px');
await site.resize(734);
check('and steps at the kit\'s width', (await site.read('#fixed', ['padding-left']))['padding-left'], '15px');
await site.close();

const consolePage = await open(shipped('console'), `<h2 id="h" class="cui-console-h2">a</h2><a id="btn" class="cui-console-btn">a</a>`);
check('a part\'s own type ships: cui-console-h2', (await consolePage.read('#h', ['font-size']))['font-size'], '26px');
check('a part ships: cui-console-btn', (await consolePage.read('#btn', ['display']))['display'] !== 'inline', true);
await consolePage.close();
await quit();
done('assets');
