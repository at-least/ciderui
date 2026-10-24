// The date control's period button under the pointer.
export default {
  path: '/apps/<app>/analytics/overview',
  probes: [
    { name: 'date.period', select: 'button:hover:has(> span[class*="rt-Text--trim--cap-alphabetic"])', props: ['background-color', 'color', 'opacity', 'cursor', 'filter'], box: true },
  ],
};
