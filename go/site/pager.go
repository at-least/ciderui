package site

import (
	"io"

	"github.com/at-least/ciderui/go/internal/glyph"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// Pager is Apple Developer Forums' pager: the page's own number between
// buttons to the pages around it. A link with no Href is not drawn, as the
// forums draw no previous button on the first page.
type Pager struct {
	AriaLabel string
	// Page: the page's own, in words ("Page 3").
	Page string
	// First and Last show their Label; Previous and Next a chevron, their
	// Label its name.
	First    PagerLink
	Previous PagerLink
	Next     PagerLink
	Last     PagerLink
}

type PagerLink struct {
	Href  string
	Label string
}

func (p Pager) Render(w io.Writer) error {
	if p.AriaLabel == "" {
		panic("site.Pager needs an AriaLabel")
	}
	if p.Page == "" {
		panic("site.Pager needs a Page")
	}
	button := func(l PagerLink, content g.Node, named bool) g.Node {
		if l.Href == "" {
			return nil
		}
		if l.Label == "" {
			panic("site.Pager's links need their Label")
		}
		return h.A(h.Class("cui-pager-button"), h.Href(l.Href), g.If(named, h.Aria("label", l.Label)), content)
	}
	return h.Nav(h.Class("cui-pager"), h.Aria("label", p.AriaLabel),
		button(p.First, g.Text(p.First.Label), false),
		button(p.Previous, glyph.ChevronLeft, true),
		h.Span(h.Class("cui-pager-page"), h.Aria("current", "page"), g.Text(p.Page)),
		button(p.Next, glyph.ChevronRight, true),
		button(p.Last, g.Text(p.Last.Label), false),
	).Render(w)
}
