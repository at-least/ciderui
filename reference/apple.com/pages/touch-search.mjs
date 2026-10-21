// The global nav's search on a touch screen, opened: its quick links,
// hovered. Read only.
import search from './search.mjs';
import { touched } from './common.mjs';

export default touched(search, ['nav', 'nav.phone', 'results.link']);
