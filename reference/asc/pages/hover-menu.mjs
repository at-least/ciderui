// The analytics menu's item under the pointer (the page marks it
// rt-Item--effect--hover): read with the pointer resting on a menu item that
// is not the current one.
const ITEM = ['height', 'background-color', 'border-top-left-radius', 'padding-left', 'padding-right', 'color', 'cursor'];
export default {
  path: '/apps/<app>/analytics/overview',
  probes: [
    { name: 'menu.item', select: 'nav[class*="rt-DisplayList"] a[class*="rt-Item--effect--hover"]', props: ITEM, box: true },
    { name: 'menu.item.label', select: 'nav[class*="rt-DisplayList"] a[class*="rt-Item--effect--hover"] > div > div', props: ['color', 'font-weight'], box: true },
  ],
};
