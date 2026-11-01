package site

import (
	"io"

	"github.com/at-least/ciderui/go/icon"
	"github.com/at-least/ciderui/go/internal/class"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// Link is a link a component writes in a list of them: its text and where it
// goes.
type Link struct {
	// Text is the link's text.
	Text string
	// Href is where it goes.
	Href string
}

// MoreLink is cui-link's more link (components/site/link,
// data-variant="more"): a standalone link, its text, and the chevron after it
// ([icon.More]), not underlined until hovered, then its text alone. One line
// of markup: a space between the text and the chevron would widen the gap.
type MoreLink struct {
	// Href is where it goes. Required: a link nowhere is not a link, so an
	// empty one panics.
	Href string
	// Text is its text (cui-link-text). Required, so an empty one panics.
	Text string
	// Class is the page's utilities on the link; "" for none.
	Class string
}

// Render writes the link.
func (l MoreLink) Render(w io.Writer) error {
	if l.Href == "" {
		panic("site.MoreLink needs an Href")
	}
	if l.Text == "" {
		panic("site.MoreLink needs a Text")
	}
	return h.A(class.Root("cui-link", l.Class), h.Data("variant", "more"), h.Href(l.Href),
		h.Span(h.Class("cui-link-text"), g.Text(l.Text)),
		icon.More.Class("cui-link-icon"),
	).Render(w)
}
