// ring(selector, png): the element's box and its accent's pixels beside it
// in the window's picture (a PNG in base64, at the screen's pixels): in the
// 8px left of its box and the 8px right of it, down the window's whole
// height, where its ring's sides are drawn (2px of the accent 3px out:
// tokens.css) and nothing else of the page is the accent. visible is
// :focus-visible, hasFocus document.hasFocus(): false in a window that is not
// in front, where Safari matches no :focus (../docs/).
async function ring(selector, png) {
  const el = document.querySelector(selector);
  const box = el.getBoundingClientRect();
  const img = new Image();
  img.src = `data:image/png;base64,${png}`;
  await img.decode();
  const scale = img.width / innerWidth;
  const c = document.createElement('canvas');
  c.width = img.width;
  c.height = img.height;
  const g = c.getContext('2d', { willReadFrequently: true });
  g.drawImage(img, 0, 0);
  // #0071e3, as the screen draws it: blue, with little red
  const count = (x0, x1) => {
    const from = Math.max(0, Math.round(x0 * scale)), to = Math.min(img.width, Math.round(x1 * scale));
    if (to <= from) return 0;
    const d = g.getImageData(from, 0, to - from, img.height).data;
    let n = 0;
    for (let i = 0; i < d.length; i += 4) if (d[i + 2] > 170 && d[i] < 80 && d[i + 1] > 60 && d[i + 1] < 170) n++;
    return n;
  };
  const outline = getComputedStyle(el);
  return {
    focused: document.activeElement === el,
    visible: el.matches(':focus-visible'),
    hasFocus: document.hasFocus(),
    outline: `${outline.outlineWidth} ${outline.outlineStyle} ${outline.outlineColor} ${outline.outlineOffset}`,
    box: { left: +box.left.toFixed(2), right: +box.right.toFixed(2), top: +box.top.toFixed(2), bottom: +box.bottom.toFixed(2) },
    height: +box.height.toFixed(2),
    scale,
    left: count(box.left - 8, box.left),
    right: count(box.right, box.right + 8),
  };
}

// held(png): an answer held short, in the window's picture: its details'
// content's height and overflow, and the ink (pixels darker than mid-gray)
// in its paragraph's box, below what is shown of it
async function held(png) {
  const item = document.querySelector('.cui-accordion-item[open]');
  const box = item.querySelector('.cui-accordion-answer p').getBoundingClientRect();
  const content = getComputedStyle(item, '::details-content');
  const img = new Image();
  img.src = `data:image/png;base64,${png}`;
  await img.decode();
  const scale = img.width / innerWidth;
  const c = document.createElement('canvas');
  c.width = img.width;
  c.height = img.height;
  const g = c.getContext('2d', { willReadFrequently: true });
  g.drawImage(img, 0, 0);
  const x = Math.round(box.left * scale), y = Math.round(box.top * scale);
  const d = g.getImageData(x, y, Math.round(box.width * scale), Math.round(box.height * scale)).data;
  let ink = 0;
  for (let i = 0; i < d.length; i += 4) if ((d[i] + d[i + 1] + d[i + 2]) / 3 < 128) ink++;
  // the content's height as laid out: the item's less its question's and
  // its rule, where a browser's style of the pseudo-element is not its own
  const laid = item.getBoundingClientRect().height - item.querySelector('summary').getBoundingClientRect().height - parseFloat(getComputedStyle(item).borderTopWidth);
  return { height: content.height, laid: +laid.toFixed(2), overflow: `${content.overflowX} ${content.overflowY}`, box: { top: +box.top.toFixed(2), height: +box.height.toFixed(2) }, scale, ink };
}
