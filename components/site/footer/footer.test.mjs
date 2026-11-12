import { examples, OUTLINE, SIZE } from '../../../test/harness.mjs';

const ex = examples(import.meta.url);
const WIDE = [1600, 1200, 900];
const PHONE = [390];
const SECTION = '.cui-footer-section:first-child';

export const cases = [
  {
    name: 'footer',
    html: ex('footer'),
    checks: [
      { target: '.cui-footer', props: [...SIZE, 'color', 'background-color'], apple: ['home', 'footer'] },
      { target: '.cui-footer-content', props: ['max-width', 'padding-left', 'padding-right'], apple: ['home', 'footer.content'] },
      { target: '.cui-footer-legal', props: ['padding-top', 'padding-bottom'], apple: ['home', 'footer.footer'] },
      { target: '.cui-footer-copyright', props: [...SIZE, 'color', 'margin-top'], apple: ['home', 'footer.copyright'] },
      { target: '.cui-footer-copyright a', props: ['color', 'text-decoration-line'], apple: ['home', 'footer.legal-link'] },
      { target: '.cui-footer', props: ['box-top'], golden: 'the body\'s last child at the window\'s foot on a page shorter than the window, not under the last line of a short page' },
      { target: '.cui-footer-copyright a', state: 'hover', props: ['text-decoration-line'], golden: 'a link in the copyright underlined under the pointer, as apple.com\'s footer links' },
    ],
  },
  {
    name: 'touch target',
    html: ex('footer'),
    touch: true,
    widths: [390],
    checks: [
      { target: '.cui-footer-copyright a', pseudo: '::after', props: ['height'], golden: 'the company link\'s target 44pt on a touch screen, the HIG\'s default, past its 16px line' },
    ],
  },
  {
    name: 'dark',
    html: `<div class="cui-dark">${ex('directory')}</div>`,
    widths: WIDE,
    checks: [
      { target: '.cui-footer', props: ['color', 'background-color'], apple: ['product', 'footer'] },
      { target: `${SECTION} .cui-footer-title`, props: ['color'], apple: ['product', 'footer.title'] },
      { target: `${SECTION} .cui-footer-list a`, props: ['color'], apple: ['product', 'footer.link'] },
    ],
  },
  {
    name: 'directory',
    html: ex('directory'),
    checks: [
      { target: '.cui-footer-directory', props: ['display', 'flex-direction', 'flex-wrap', 'padding-top'], apple: ['home', 'footer.directory'] },
    ],
  },
  {
    // @trap: apple.com's column holds sections; here a section is the column, read against its column
    name: 'directory-wide',
    html: ex('directory'),
    widths: WIDE,
    checks: [
      { target: SECTION, props: ['flex-grow', 'flex-shrink', 'flex-basis'], apple: ['home', 'footer.column'] },
      { target: `${SECTION} .cui-footer-title`, props: [...SIZE, 'color'], apple: ['home', 'footer.title'] },
      { target: `${SECTION} .cui-footer-title`, props: ['margin-right', 'margin-bottom'], apple: ['home', 'footer.title-text'] },
      { target: `${SECTION} .cui-footer-list`, props: ['margin-right', 'padding-top', 'padding-bottom'], apple: ['home', 'footer.list'] },
      { target: `${SECTION} .cui-footer-list li:first-child`, props: ['margin-bottom'], apple: ['home', 'footer.item'] },
      { target: `${SECTION} .cui-footer-list li:last-child`, props: ['margin-bottom'], apple: ['home', 'footer.item-last'] },
      { target: `${SECTION} .cui-footer-list a`, props: [...SIZE, 'color', 'text-decoration-line'], apple: ['home', 'footer.link'] },
      { target: `${SECTION} .cui-footer-list a`, state: 'hover', props: ['color', 'text-decoration-line'], apple: ['home', 'footer.link', 'hover'] },
      { target: `${SECTION} .cui-footer-toggle`, props: ['display'], golden: 'its toggle gone from 834px: every section shown' },
    ],
  },
  {
    name: 'directory-phone',
    html: ex('directory'),
    widths: PHONE,
    checks: [
      { target: SECTION, props: ['padding-left', 'padding-right', 'padding-bottom', 'margin-left', 'position', 'box-height'], apple: ['home', 'footer.section-phone'] },
      { target: SECTION, pseudo: '::after', props: ['content', 'position', 'bottom', 'left', 'right', 'height', 'background-color'], apple: ['home', 'footer.section-rule'] },
      { target: `${SECTION} .cui-footer-title label`, props: [...SIZE, 'color', 'padding-top', 'padding-bottom', 'cursor', 'box-height'], apple: ['home', 'footer.title-button'] },
      { target: `${SECTION} .cui-footer-title svg`, props: ['color', 'margin-top', 'box-height'], apple: ['home', 'footer.title-icon'] },
      { target: `${SECTION} .cui-footer-title svg`, props: ['transition-property', 'transition-duration', 'transition-timing-function'], golden: 'its turn the kit\'s one chevron turn, 0.32s cubic-bezier(.4,0,.6,1), the accordion\'s, where apple.com\'s reads 0.3s ease here' },
      { target: `${SECTION} .cui-footer-list`, props: ['display'], golden: 'closed, its links out of sight and out of the focus order (apple.com\'s hidden: home:footer.list-closed)' },
    ],
  },
  {
    name: 'directory-open',
    html: ex('directory'),
    widths: PHONE,
    click: `${SECTION} .cui-footer-title label`,
    checks: [
      { target: `${SECTION} .cui-footer-list`, props: ['padding-top', 'padding-bottom'], apple: ['home-footer', 'list.open'] },
      { target: `${SECTION} .cui-footer-list li:first-child`, props: ['box-height'], apple: ['home-footer', 'item.open'] },
      { target: `${SECTION} .cui-footer-list a`, props: [...SIZE, 'color', 'display', 'padding-top', 'padding-right', 'padding-bottom', 'padding-left'], apple: ['home-footer', 'link.open'] },
      { target: `${SECTION} .cui-footer-title svg`, props: ['color'], apple: ['home-footer', 'icon.open'] },
      { target: `${SECTION} .cui-footer-title svg`, props: ['transform'], golden: 'opened, its chevron turned up: apple.com\'s line moved to the same points' },
    ],
  },
  {
    name: 'directory step at 834px',
    html: ex('directory'),
    widths: [834, 833],
    checks: [
      { target: '.cui-footer-content', props: ['padding-left', 'padding-right'], golden: '22px in from 834px, 16px to 833px: apple.com\'s footer steps there, with the global nav (reference/apple.com/MEASURED.md, not read at these widths)' },
      { target: '.cui-footer-directory', props: ['flex-direction'], golden: 'a row of columns from 834px, a column of sections to 833px, as apple.com\'s' },
      { target: `${SECTION} .cui-footer-toggle`, props: ['display'], golden: 'its toggle from 833px down alone' },
    ],
  },
  {
    name: 'directory open at once',
    html: ex('directory'),
    widths: PHONE,
    click: `${SECTION} .cui-footer-title label`,
    wait: 0,
    checks: [
      { target: `${SECTION} .cui-footer-list`, props: ['display', 'box-height'], golden: 'opened at once, its whole list at the click: apple.com\'s moves in 0.3s' },
    ],
  },
  {
    name: 'directory toggle focus',
    html: ex('directory'),
    widths: PHONE,
    press: ['Tab'],
    checks: [
      { target: `${SECTION} .cui-footer-toggle`, props: ['focused'], golden: 'the toggle, out of sight, the page\'s first stop' },
      { target: `${SECTION} .cui-footer-title label`, props: OUTLINE, golden: 'its focus ring on its title' },
    ],
  },
  {
    // @trap: from 834px: below it a closed section's links have no box
    name: 'directory touch targets',
    html: ex('directory'),
    touch: true,
    widths: [1200],
    checks: [
      { target: `${SECTION} .cui-footer-title label`, props: ['box-height'], golden: 'a section\'s title 44px tall on a touch screen' },
      { target: `${SECTION} .cui-footer-list li:first-child a`, props: ['box-height'], golden: 'a link 44px tall on a touch screen, the HIG\'s 44pt (Accessibility, Mobility), a tablet\'s columns too' },
      { target: `${SECTION} .cui-footer-list li:first-child`, props: ['margin-bottom'], golden: 'the list\'s gaps inside the links\' 44px' },
    ],
  },
];
