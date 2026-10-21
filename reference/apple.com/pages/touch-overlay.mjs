// The store's overlay on a touch screen, opened: its close button's circle,
// hovered. Read only.
import overlay from './overlay.mjs';
import { touched } from './common.mjs';

export default touched(overlay, ['overlay', 'close.circle-hover']);
