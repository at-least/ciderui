// The marketing pages' modal on a touch screen, opened: its close button's
// circle, hovered. Read only.
import ricmodal from './ricmodal.mjs';
import { touched } from './common.mjs';

export default touched(ricmodal, ['modal', 'close.icon-states']);
