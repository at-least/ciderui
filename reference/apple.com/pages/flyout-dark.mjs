// The global nav's flyouts on a dark page, iPhone 18 Pro's (its bar clear,
// product:nav): flyout.mjs's probes, at 1600px. The Newsroom's dark nav
// opens none (`globalnav-submenus-disabled` on its nav, 2026-10-21).
import flyout from './flyout.mjs';

export default { ...flyout, url: 'https://www.apple.com/iphone-18-pro/', widths: [1600] };
