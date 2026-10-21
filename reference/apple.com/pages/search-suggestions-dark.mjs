// The global nav's search showing its suggestions on a dark nav, the
// Newsroom's: search-suggestions.mjs's probes, at a width of each of the
// nav's layouts.
import suggestions from './search-suggestions.mjs';

export default { ...suggestions, url: 'https://www.apple.com/newsroom/', widths: [1600, 390] };
