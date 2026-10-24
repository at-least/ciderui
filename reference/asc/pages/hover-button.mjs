// The sources page's gray button with the pointer resting on it.
export default {
  path: '/apps/<app>/analytics/acquisition/sources',
  probes: [
    { name: 'button', select: 'button[class*="rt-Button--variant--gray"]:hover', props: ['background-color', 'color', 'cursor', 'opacity'], box: true },
  ],
};
