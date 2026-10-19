// The footer against apple.com's home page (home:footer*): its ground and
// type, its content's width and inset, its footnotes and their rule, its
// directory (a row of columns, a title and its list, the gaps) from 834px;
// to 833px its sections, rows over rules, one opened (home-footer:*); the
// shop line, the legal row and its links' rules, the locale; on a dark page
// light (product:footer*). Golden, with their why: the letter-spacing, a
// section's toggle, its motion, the chevron turned.

import { examples } from '../../test/harness.mjs';

const ex = examples(import.meta.url);
const TYPE = ['font-size', 'line-height', 'font-weight'];
const WIDE = [1600, 1200, 900];
const PHONE = [390];
const SECTION = '.cui-footer-column:first-child .cui-footer-section:first-child';

export const cases = [
  {
    name: 'footer',
    html: ex('footer'),
    checks: [
      { target: '.cui-footer', props: [...TYPE, 'color', 'background-color'], apple: ['home', 'footer'] },
      { target: '.cui-footer-content', props: ['max-width', 'padding-left', 'padding-right'], apple: ['home', 'footer.content'] },
      { target: '.cui-footer-notes', props: [...TYPE, 'color', 'padding-top', 'padding-bottom', 'border-bottom-width', 'border-bottom-style', 'border-bottom-color'], apple: ['home', 'footer.sosumi'] },
      { target: '.cui-footer-notes a', props: [...TYPE, 'color', 'text-decoration-line'], apple: ['home', 'footer.sosumi-link'] },
      { target: '.cui-footer-notes a', state: 'hover', props: ['color', 'text-decoration-line'], apple: ['home', 'footer.sosumi-link', 'hover'] },
      { target: '.cui-footer-directory', props: ['display', 'flex-direction', 'flex-wrap', 'padding-top'], apple: ['home', 'footer.directory'] },
      { target: '.cui-footer-shop', props: [...TYPE, 'color'], apple: ['home', 'footer.shop'] },
      { target: '.cui-footer-legal', props: ['padding-top', 'padding-bottom'], apple: ['home', 'footer.footer'] },
      { target: '.cui-footer-row', props: ['display'], apple: ['home', 'footer.end'] },
      { target: '.cui-footer-copyright', props: [...TYPE, 'color', 'margin-top'], apple: ['home', 'footer.copyright'] },
      { target: '.cui-footer-links > li:first-child', props: ['display', 'margin-right'], apple: ['home', 'footer.legal-item'] },
      { target: '.cui-footer-links > li:first-child a', props: ['padding-left', 'padding-right', 'margin-left'], apple: ['home', 'footer.legal-link-box'] },
      { target: '.cui-footer-links > li:first-child a', props: [...TYPE, 'color', 'text-decoration-line'], apple: ['home', 'footer.legal-link'] },
      { target: '.cui-footer-links > li:first-child', pseudo: '::after', props: ['content', 'height', 'background-color', 'margin-top', 'position', 'right'], apple: ['home', 'footer.legal-rule'] },
      { target: '.cui-footer-links > li:first-child', pseudo: '::after', props: ['width'], golden: 'apple.com\'s 1px rule (home:footer.legal-rule)' },
      { target: '.cui-footer-locale', props: [...TYPE, 'color', 'text-decoration-line'], apple: ['home', 'footer.locale-link'] },
      { target: '.cui-footer-locale', state: 'hover', props: ['text-decoration-line'], apple: ['home', 'footer.locale-link', 'hover'] },
      // apple.com's two locales (a phone's, a desktop's) are 5px down alike
      { target: '.cui-footer-locale', props: ['margin-top'], apple: ['home', 'footer.locale'] },
      { target: '.cui-footer', props: ['letter-spacing'], golden: 'the system font\'s fitted letter-spacing (tokens.css), not apple.com\'s' },
    ],
  },
  {
    name: 'footer-wide',
    html: ex('footer'),
    widths: WIDE,
    checks: [
      { target: '.cui-footer-column', props: ['flex-grow', 'flex-shrink', 'flex-basis'], apple: ['home', 'footer.column'] },
      { target: `${SECTION} + .cui-footer-section`, props: ['padding-top'], apple: ['home', 'footer.section-second'] },
      { target: `${SECTION} .cui-footer-title`, props: [...TYPE, 'color'], apple: ['home', 'footer.title'] },
      { target: `${SECTION} .cui-footer-title`, props: ['margin-right', 'margin-bottom'], apple: ['home', 'footer.title-text'] },
      { target: `${SECTION} .cui-footer-list`, props: ['margin-right', 'padding-top', 'padding-bottom'], apple: ['home', 'footer.list'] },
      { target: `${SECTION} .cui-footer-list li:first-child`, props: ['margin-bottom'], apple: ['home', 'footer.item'] },
      { target: `${SECTION} .cui-footer-list li:last-child`, props: ['margin-bottom'], apple: ['home', 'footer.item-last'] },
      { target: `${SECTION} .cui-footer-list a`, props: [...TYPE, 'color', 'text-decoration-line'], apple: ['home', 'footer.link'] },
      { target: `${SECTION} .cui-footer-list a`, state: 'hover', props: ['color', 'text-decoration-line'], apple: ['home', 'footer.link', 'hover'] },
      { target: '.cui-footer-shop', props: ['padding-bottom', 'margin-bottom', 'border-bottom-width', 'border-bottom-color'], apple: ['home', 'footer.shop'] },
      { target: `${SECTION} .cui-footer-toggle`, props: ['display'], golden: 'its toggle gone from 834px: every section shown' },
    ],
  },
  {
    name: 'footer-phone',
    html: ex('footer'),
    widths: PHONE,
    checks: [
      { target: SECTION, props: ['padding-left', 'padding-right', 'padding-bottom', 'margin-left', 'position', 'box-height'], apple: ['home', 'footer.section-phone'] },
      { target: SECTION, pseudo: '::after', props: ['content', 'position', 'bottom', 'left', 'right', 'height', 'background-color'], apple: ['home', 'footer.section-rule'] },
      { target: `${SECTION} .cui-footer-title label`, props: [...TYPE, 'color', 'padding-top', 'padding-bottom', 'cursor', 'box-height'], apple: ['home', 'footer.title-button'] },
      { target: `${SECTION} .cui-footer-title svg`, props: ['color', 'margin-top', 'box-height'], apple: ['home', 'footer.title-icon'] },
      { target: '.cui-footer-shop', props: ['padding-bottom', 'margin-bottom', 'border-bottom-width'], apple: ['home', 'footer.shop'] },
      { target: '.cui-footer-locale', props: ['margin-top', 'display'], apple: ['home', 'footer.locale'] },
      { target: `${SECTION} .cui-footer-list`, props: ['display'], golden: 'closed, its links out of sight and out of the focus order (apple.com\'s hidden: home:footer.list-closed)' },
    ],
  },
  {
    // a phone's section opened by its title
    name: 'footer-open',
    html: ex('footer'),
    widths: PHONE,
    click: `${SECTION} .cui-footer-title label`,
    checks: [
      { target: `${SECTION} .cui-footer-list`, props: ['padding-top', 'padding-bottom'], apple: ['home-footer', 'list.open'] },
      { target: `${SECTION} .cui-footer-list li:first-child`, props: ['box-height'], apple: ['home-footer', 'item.open'] },
      { target: `${SECTION} .cui-footer-list a`, props: [...TYPE, 'color', 'display', 'padding-top', 'padding-right', 'padding-bottom', 'padding-left'], apple: ['home-footer', 'link.open'] },
      { target: `${SECTION} .cui-footer-title svg`, props: ['color'], apple: ['home-footer', 'icon.open'] },
      { target: `${SECTION} .cui-footer-title svg`, props: ['transform'], golden: 'opened, its chevron turned up: apple.com\'s line moved to the same points' },
    ],
  },
  {
    name: 'dark',
    html: `<div class="cui-dark">${ex('footer')}</div>`,
    widths: WIDE,
    checks: [
      { target: '.cui-footer', props: ['color', 'background-color'], apple: ['product', 'footer'] },
      { target: `${SECTION} .cui-footer-title`, props: ['color'], apple: ['product', 'footer.title'] },
      { target: `${SECTION} .cui-footer-list a`, props: ['color'], apple: ['product', 'footer.link'] },
      { target: '.cui-footer-shop', props: ['border-bottom-color'], apple: ['product', 'footer.shop'] },
    ],
  },
];
