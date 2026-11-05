package site

import (
	"io"

	"github.com/at-least/ciderui/go/internal/class"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

type Section struct {
	// Alt: the alternate ground.
	Alt bool
	// Large: apple.com's iPhone 18 Pro page's rhythm; else its iPhone page's.
	Large    bool
	Class    string
	Children g.Node
}

func (s Section) Render(w io.Writer) error {
	return h.Section(class.Root("cui-section", s.Class),
		g.If(s.Alt, h.Data("variant", "alt")),
		g.If(s.Large, h.Data("size", "lg")),
		h.Div(h.Class("cui-section-content"), s.Children),
	).Render(w)
}

type SectionHeader struct {
	Headline string
	Links    []Link
	Class    string
}

func (s SectionHeader) Render(w io.Writer) error {
	if s.Headline == "" {
		panic("site.SectionHeader needs a Headline")
	}
	return h.Header(class.Root("cui-section-header", s.Class),
		h.H2(h.Class("cui-section-headline"), g.Text(s.Headline)),
		g.If(len(s.Links) > 0, h.Ul(h.Class("cui-section-links"), g.Map(s.Links, func(link Link) g.Node {
			return h.Li(MoreLink{Href: link.Href, Label: link.Label, Class: ""})
		}))),
	).Render(w)
}
