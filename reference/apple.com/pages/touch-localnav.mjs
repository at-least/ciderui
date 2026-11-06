import localnav from './localnav.mjs';
import { touched } from './common.mjs';

export default touched(localnav, ['localnav.link', 'localnav.button', 'localnav.chevron-image', 'tray.open', 'tray.link']);
