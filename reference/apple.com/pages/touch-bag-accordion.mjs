// The bag's accordion on a touch screen, opened: its compact question's
// chevron, its button hovered. Read only.
import bag_accordion from './bag-accordion.mjs';
import { touched } from './common.mjs';

export default touched(bag_accordion, ['accordion.open', 'compact.icon-hover']);
