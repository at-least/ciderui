/* The frame never scrolls: as tall as its content, counting what is laid
 * out of its flow and visible (a local nav's menu below its bar); while a
 * popover or a modal is open, as tall as the page's window below the page's
 * sticky bars, and scrolled there. frameElement is reachable: a srcdoc frame
 * shares the page's origin. */
const root = document.documentElement;
let sheet = 0;
// a closed <details>' content is laid out though not shown
const size = () => {
  const bottoms = [...document.querySelectorAll('body *')].filter((e) => getComputedStyle(e).visibility !== 'hidden' && !e.closest('details:not([open]) > :not(summary)')).map((e) => e.getBoundingClientRect().bottom);
  frameElement.style.height = `${Math.ceil(Math.max(sheet, root.getBoundingClientRect().height, ...bottoms))}px`;
};
const observer = new ResizeObserver(size);
observer.observe(root);
addEventListener('DOMContentLoaded', () => {
  for (const e of document.querySelectorAll('body *')) observer.observe(e);
});
// a menu shown by a checkbox (the local nav's) changes what is visible, not
// its boxes' sizes
for (const type of ['change', 'transitionend']) document.addEventListener(type, size, true);
const bars = () => [...parent.document.querySelectorAll('.cui-globalnav, .cui-localnav, .cui-docs-bar')]
  .filter((e) => ['sticky', 'fixed'].includes(parent.getComputedStyle(e).position))
  .reduce((sum, e) => sum + e.getBoundingClientRect().height, 0);
const cover = (open) => {
  const covered = bars();
  sheet = open ? parent.innerHeight - covered : 0;
  size();
  if (sheet) {
    frameElement.style.scrollMarginTop = `${covered}px`;
    frameElement.scrollIntoView({ block: 'start' });
  }
};
document.addEventListener('toggle', (e) => {
  if (e.target.matches('[popover]')) cover(e.newState === 'open');
}, true);
new MutationObserver(() => cover(Boolean(document.querySelector('dialog:modal')))).observe(document, { subtree: true, attributeFilter: ['open'] });

// a srcdoc frame's base URL is the page's: a click on an example's link
// would load the page (#buy) or another one (/phone/) into the frame. A
// gallery's dot is the theme's script's, which has already handled it.
addEventListener('click', (e) => {
  const link = e.target.closest('a[href]');
  if (!link || e.defaultPrevented) return;
  e.preventDefault();
  const href = link.getAttribute('href');
  const target = href.startsWith('#') && document.getElementById(decodeURIComponent(href.slice(1)));
  if (target) {
    target.style.scrollMarginTop = `${bars()}px`;
    target.scrollIntoView();
  }
});
// sent, a form would load the page it names into the frame
addEventListener('submit', (e) => {
  if (e.target.method !== 'dialog') e.preventDefault();
});
