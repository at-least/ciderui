// The local nav on a touch screen: its links, its button, its chevron, and
// its tray's links (opened at 390px), hovered. Read only.
import localnav from './localnav.mjs';
import { touched } from './common.mjs';

export default touched(localnav, ['localnav.link', 'localnav.button', 'localnav.chevron-image', 'tray.open', 'tray.link']);
