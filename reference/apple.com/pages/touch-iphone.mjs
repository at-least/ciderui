// The iPhone page on a touch screen: its more link, a link in text, the
// gallery's paddle and dots, hovered. Read only.
import iphone from './iphone.mjs';
import { touched } from './common.mjs';

export default touched(iphone, ['link.more-copy', 'link.inline', 'paddle', 'gallery.dot']);
