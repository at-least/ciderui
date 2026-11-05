package site

import (
	"io"

	"github.com/at-least/ciderui/go/internal/class"
	"github.com/at-least/ciderui/go/internal/glyph"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// a fixed id: a page has one local nav
const localNavToggleID = "cui-localnav-menu"

type LocalNav struct {
	Title   string
	Href    string
	Links   []Link
	Current string
	Button  *Link
	Dark    bool
	Class   string
}

func (n LocalNav) Render(w io.Writer) error {
	if n.Title == "" {
		panic("site.LocalNav needs a Title")
	}
	if n.Href == "" {
		panic("site.LocalNav needs an Href")
	}
	block := "cui-localnav"
	if n.Dark {
		block = "cui-localnav cui-dark"
	}
	menu := len(n.Links) > 0
	return h.Nav(class.Root(block, n.Class), h.Aria("label", "Local"),
		h.Div(h.Class("cui-localnav-content"),
			h.A(h.Class("cui-localnav-title"), h.Href(n.Href), g.Text(n.Title)),
			g.If(menu, g.Group{
				h.Input(h.Class("cui-localnav-toggle"), h.Type("checkbox"), h.ID(localNavToggleID), h.Aria("label", "Menu")),
				h.Div(h.Class("cui-localnav-menu"),
					h.Ul(h.Class("cui-localnav-list"), g.Map(n.Links, func(link Link) g.Node {
						return h.Li(h.A(h.Class("cui-localnav-link"), h.Href(link.Href), g.If(link.Href == n.Current, h.Aria("current", "page")), g.Text(link.Text)))
					})),
				),
			}),
			h.Div(h.Class("cui-localnav-actions"),
				g.If(menu, h.Label(h.Class("cui-localnav-menubutton"), h.For(localNavToggleID), h.Aria("hidden", "true"), glyph.Chevron)),
				g.Iff(n.Button != nil, func() g.Node {
					return h.A(h.Class("cui-btn"), h.Data("size", "sm"), h.Href(n.Button.Href), g.Text(n.Button.Text))
				}),
			),
		),
	).Render(w)
}
