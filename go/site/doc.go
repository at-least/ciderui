// Package site is ciderui's site set (components/site/) as gomponents nodes:
// the markup the cui-* classes expect. Only the components whose markup is
// more than a class on one element are here: a button is a class, written on
// the element with the attributes the page gives it
// (<a class="cui-btn" data-size="cta">).
//
// Each component is a struct whose every field a call writes: the module's
// exhaustruct check (go tool exhaustruct) fails a call that leaves one out,
// as a site's own gate should, so a field added here stops every call that
// does not set it.
package site
