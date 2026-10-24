// A section's See All link under the pointer.
import { LINK } from '../../apple.com/pages/common.mjs';
export default {
  path: '/apps/<app>/analytics/overview',
  probes: [
    { name: 'section.link', select: 'a[class*="rt-Link--variant--inline"]:hover', props: ['color', ...LINK, 'opacity', 'cursor'], box: true },
  ],
};
