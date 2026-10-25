// The one Chromium every check and picture is read in: Playwright's, in the
// image test/Dockerfile builds (its Chromium, its fonts: what a page falls
// back to, Liberation Mono for code, is the image's, not the machine's), so
// a value or a picture is the same on every machine. Outside that image it
// refuses to start: node scripts/docker.mjs runs a step in it.
//
// --font-render-hinting=none: on Linux Chromium hints every glyph to a
// whole-pixel advance otherwise (reference/tracking/).

import { chromium } from 'playwright-chromium';

export const launch = () => {
  if (process.env.CIDERUI_CONTAINER !== '1') {
    throw new Error('the browser runs only in test/Dockerfile\'s image: node scripts/docker.mjs <command>');
  }
  return chromium.launch({ args: ['--font-render-hinting=none'] });
};
