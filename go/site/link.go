package site

import (
	"io"

	"github.com/at-least/ciderui/go/internal/class"
	"github.com/at-least/ciderui/go/internal/glyph"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

type Link struct {
	Text string
	Href string
}

type MoreLink struct {
	Href  string
	Text  string
	Class string
}

func (l MoreLink) Render(w io.Writer) error {
	if l.Href == "" {
		panic("site.MoreLink needs an Href")
	}
	if l.Text == "" {
		panic("site.MoreLink needs a Text")
	}
	return h.A(class.Root("cui-link", l.Class), h.Data("variant", "more"), h.Href(l.Href),
		// no space between the text and the icon: it would widen the gap
		h.Span(h.Class("cui-link-text"), g.Text(l.Text)),
		glyph.More.Class("cui-link-icon"),
	).Render(w)
}
