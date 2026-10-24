// An app's analytics sources page, the newer look: a gray neutral button of
// the toolkit's (rt-Button--variant--gray, controlSize md, controlShape
// roundedRect): its box, its type and its colors at rest.
import { TYPE, BOX, TEXT } from '../../apple.com/pages/common.mjs';

export default {
  path: '/apps/<app>/analytics/acquisition/sources',
  probes: [
    { name: 'button', select: 'button[class*="rt-Button--variant--gray"]', props: [...TYPE, ...BOX, ...TEXT, 'height', 'column-gap', 'border-top-right-radius', 'border-bottom-left-radius', 'outline-style', 'transition-property', 'transition-duration'], box: true },
  ],
};
