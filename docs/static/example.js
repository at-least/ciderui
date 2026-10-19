/* ciderui's documentation: what each example frame runs
 * (templates/cui-example.html). The frame is a page of its own inside a page
 * of the documentation, and it is as tall as its content, so it never
 * scrolls: its flow, and what is laid out of it and visible (a local nav's
 * menu below its bar), whenever any of it changes size or shows; while a
 * global nav's menu is open, a sheet over the frame's window, as tall as
 * the page's window below the page's own bar, and scrolled there. A srcdoc
 * frame shares the page's origin, which lets it size its own element
 * there. */
const root = document.documentElement;
let sheet = 0;
const size = () => {
  const bottoms = [...document.querySelectorAll('body *')].filter((e) => getComputedStyle(e).visibility !== 'hidden').map((e) => e.getBoundingClientRect().bottom);
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
// the page's bars that stay at the top of its window as it scrolls (its
// local nav), which the sheet keeps below
const bars = () => [...parent.document.querySelectorAll('.cui-globalnav, .cui-localnav')]
  .filter((e) => ['sticky', 'fixed'].includes(parent.getComputedStyle(e).position))
  .reduce((sum, e) => sum + e.getBoundingClientRect().height, 0);
document.addEventListener('toggle', (e) => {
  const covered = bars();
  sheet = e.newState === 'open' ? parent.innerHeight - covered : 0;
  size();
  if (sheet) {
    frameElement.style.scrollMarginTop = `${covered}px`;
    frameElement.scrollIntoView({ block: 'start' });
  }
}, true);

// an example's links are the example's, not pages of the documentation: a
// srcdoc frame's base URL is the page's, so a click would load the page (a
// link to #buy) or another one (/iphone/) into the frame. A link to an
// element of the frame (a footnote's mark) scrolls to it; a gallery's dot is
// the theme's script's, which has already handled it.
addEventListener('click', (e) => {
  const link = e.target.closest('a[href]');
  if (!link || e.defaultPrevented) return;
  e.preventDefault();
  const href = link.getAttribute('href');
  const target = href.startsWith('#') && document.getElementById(decodeURIComponent(href.slice(1)));
  if (target) target.scrollIntoView();
});
