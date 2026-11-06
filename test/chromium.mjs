// @external: Only in test/Dockerfile's image: its Chromium and its fallback fonts
// (Liberation Mono for code) make a value or a picture the same on every
// machine.
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
