import { TYPE, BOX, TEXT } from '../../apple.com/pages/common.mjs';

export default {
  path: '/apps/<app>/analytics/acquisition/sources',
  probes: [
    { name: 'button', select: 'button[class*="rt-Button--variant--gray"]', props: [...TYPE, ...BOX, ...TEXT, 'height', 'column-gap', 'border-top-right-radius', 'border-bottom-left-radius', 'outline-style', 'transition-property', 'transition-duration'], box: true },
  ],
};
