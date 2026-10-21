// The home page on a touch screen: its global nav's links, logo, search,
// bag and menu button, its buttons, a gallery's dots, its footer's links, and
// the menu's links (opened at 390px), hovered, each against itself at rest. Read only.
import home from './home.mjs';
import { touched } from './common.mjs';

export default touched(home, ['nav.link', 'nav.logo', 'nav.search', 'nav.bag', 'nav.menutrigger', 'button', 'button.secondary', 'gallery.dot', 'footer.sosumi-link', 'footer.link', 'footer.legal-link', 'footer.locale-link', 'menu.open', 'menu.link']);
