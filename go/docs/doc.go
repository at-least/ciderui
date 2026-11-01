// Package docs is ciderui's docs set (components/docs/) as gomponents nodes:
// Apple's developer documentation's parts, the markup the cui-docs-* classes
// expect. The docs content (cui-docs-content) is a class on the page's own
// markup and has no node here.
//
// Each component is a struct whose every field a call writes: the module's
// exhaustruct check (go tool exhaustruct) fails a call that leaves one out,
// as a site's own gate should, so a field added here stops every call that
// does not set it.
package docs
