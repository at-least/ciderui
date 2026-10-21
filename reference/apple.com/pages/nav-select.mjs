// What of the home page can be selected, for the global nav's: apple.com's
// stylesheet sets user-select: none on #globalnav and on #globalnav
// .globalnav-list (globalheader.css, its text read 2026-10-22), nowhere else
// in the nav, its search's field not excepted. Read here as each element
// computes it: the nav, its list, the logo, a link and its text, the bag,
// the menu's button, a closed flyout; the page's text and its footer beside
// them; then the search opened (its button clicked once, at each of the
// nav's layouts' widths read) and its field, last. Once it has opened the
// pointer moves into it, where it hovers none (search's). At 1600px "ma" is
// typed into the field (the owner's leave to type on apple.com, 2026-10-21;
// nothing sent), its placeholder gone if the field took the keys.
const SELECT = ['user-select'];
const FIELD = '#globalnav-submenu-search .globalnav-searchfield-input';

export default {
  url: 'https://www.apple.com/',
  widths: [1600, 390],
  wait: 2500,
  probes: [
    { name: 'nav', select: '#globalnav', props: SELECT },
    { name: 'list', select: '#globalnav .globalnav-list', props: SELECT },
    { name: 'logo', select: '#globalnav .globalnav-link-apple', props: SELECT },
    { name: 'link', select: '#globalnav .globalnav-item-mac .globalnav-link', props: SELECT },
    { name: 'link-text', select: '#globalnav .globalnav-item-mac .globalnav-link-text', props: SELECT },
    { name: 'bag', select: '#globalnav .globalnav-link-bag', props: SELECT },
    { name: 'menutrigger', select: '#globalnav .globalnav-menutrigger-button', props: SELECT, widths: [390] },
    { name: 'flyout', select: '#globalnav-submenu-link-mac', props: SELECT },
    { name: 'page', select: 'main', props: SELECT },
    { name: 'footer', select: '#ac-globalfooter', props: SELECT },
    { name: 'search.field', click: '#globalnav-menubutton-link-search', pointer: [200, 430], select: FIELD, props: SELECT, widths: [1600] },
    { name: 'search.typed', type: ['m', 'a'], typed: 1500, select: FIELD, props: SELECT, matches: [':placeholder-shown'], widths: [1600] },
    { name: 'search.field-phone', click: '#globalnav-menubutton-link-search', pointer: [200, 850], select: FIELD, props: SELECT, widths: [390] },
  ],
};
