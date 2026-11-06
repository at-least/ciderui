import home from './home.mjs';
import { touched } from './common.mjs';

export default touched(home, ['nav.link', 'nav.logo', 'nav.search', 'nav.bag', 'nav.menutrigger', 'button', 'button.secondary', 'gallery.dot', 'footer.sosumi-link', 'footer.link', 'footer.legal-link', 'footer.locale-link', 'menu.open', 'menu.link']);
