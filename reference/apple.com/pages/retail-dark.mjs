// The retail page's field in a dark color scheme (prefers-color-scheme:
// dark): apple.com's forms keep their light ground and text.
import { probes } from './retail.mjs';
export default {
  url: 'https://www.apple.com/retail/',
  colorScheme: 'dark',
  probes: probes.filter(p => ['field.input', 'field.label'].includes(p.name)),
};
