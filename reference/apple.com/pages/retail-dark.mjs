// The retail page's field in a dark color scheme (prefers-color-scheme:
// dark): apple.com's forms keep their light ground and text; a field in
// error and a disabled one, as copies the probe adds.
import { probes } from './retail.mjs';
export default {
  url: 'https://www.apple.com/retail/',
  colorScheme: 'dark',
  probes: probes.filter(p => ['field.input', 'field.label', 'field.error', 'field.error-label', 'field.error-message', 'field.error-icon', 'field.disabled', 'field.disabled-label'].includes(p.name)),
};
