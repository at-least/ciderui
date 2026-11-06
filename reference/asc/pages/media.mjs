const THUMB = ['border-top-left-radius', 'border-top-width', 'border-top-style', 'background-color', 'object-fit', 'cursor', 'box-shadow'];
export default {
  path: '/apps/<app>/distribution/ios/version/inflight',
  probes: [
    { name: 'thumb.img', select: 'img[class*="Img-sc"]', props: THUMB, box: true },
    { name: 'thumb.frame', select: 'img[class*="Img-sc"]', up: 1, props: THUMB, box: true },
    { name: 'thumb.item', select: 'img[class*="Img-sc"]', up: 3, props: ['margin-left', 'margin-right', 'cursor'], box: true },
    { name: 'thumb.row', select: 'img[class*="Img-sc"]', up: 5, props: ['display', 'margin-top', 'gap'] },
  ],
};
