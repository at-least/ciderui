package site

import (
	"io"

	"github.com/at-least/ciderui/go/internal/class"
	"github.com/at-least/ciderui/go/internal/glyph"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

type Link struct {
	Label string
	Href  string
}

type MoreLink struct {
	Href  string
	Label string
	Class string
}

func (l MoreLink) Render(w io.Writer) error {
	if l.Href == "" {
		panic("site.MoreLink needs an Href")
	}
	if l.Label == "" {
		panic("site.MoreLink needs a Label")
	}
	return h.A(class.Root("cui-link", l.Class), h.Data("variant", "more"), h.Href(l.Href),
		// no space between the text and the icon: it would widen the gap
		h.Span(h.Class("cui-link-text"), g.Text(l.Label)),
		glyph.More.Class("cui-link-icon"),
	).Render(w)
}
