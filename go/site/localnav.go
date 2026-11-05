package site

import (
	"io"
	"slices"

	"github.com/at-least/ciderui/go/internal/class"
	"github.com/at-least/ciderui/go/internal/glyph"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// a fixed id: a page has one local nav
const localNavToggleID = "cui-localnav-menu"

type LocalNav struct {
	Title string
	Href  string
	Links []Link
	// Current: the Href of the link of this page; "" where none is.
	Current string
	Button  *Link
	// AriaLabel: names the nav among the page's others.
	AriaLabel string
	// MenuAriaLabel: a phone's toggle of its Links; "" with no Links.
	MenuAriaLabel string
	Class         string
}

func (n LocalNav) Render(w io.Writer) error {
	if n.Title == "" {
		panic("site.LocalNav needs a Title")
	}
	if n.Href == "" {
		panic("site.LocalNav needs an Href")
	}
	if n.AriaLabel == "" {
		panic("site.LocalNav needs an AriaLabel")
	}
	menu := len(n.Links) > 0
	if menu && n.MenuAriaLabel == "" {
		panic("site.LocalNav's Links need a MenuAriaLabel")
	}
	for _, link := range n.Links {
		if link.Label == "" || link.Href == "" {
			panic("site.LocalNav's Links need their Label and Href")
		}
	}
	if n.Current != "" && !slices.ContainsFunc(n.Links, func(l Link) bool { return l.Href == n.Current }) {
		panic("site.LocalNav's Current is none of its Links'")
	}
	if n.Button != nil && (n.Button.Label == "" || n.Button.Href == "") {
		panic("site.LocalNav's Button needs its Label and Href")
	}
	return h.Nav(class.Root("cui-localnav", n.Class), h.Aria("label", n.AriaLabel),
		h.Div(h.Class("cui-localnav-content"),
			h.A(h.Class("cui-localnav-title"), h.Href(n.Href), g.Text(n.Title)),
			g.If(menu, g.Group{
				h.Input(h.Class("cui-localnav-toggle"), h.Type("checkbox"), h.ID(localNavToggleID), h.Aria("label", n.MenuAriaLabel)),
				h.Div(h.Class("cui-localnav-menu"),
					h.Ul(h.Class("cui-localnav-list"), g.Map(n.Links, func(link Link) g.Node {
						return h.Li(h.A(h.Class("cui-localnav-link"), h.Href(link.Href), g.If(link.Href == n.Current, h.Aria("current", "page")), g.Text(link.Label)))
					})),
				),
			}),
			h.Div(h.Class("cui-localnav-actions"),
				g.If(menu, h.Label(h.Class("cui-localnav-menubutton"), h.For(localNavToggleID), h.Aria("hidden", "true"), glyph.Chevron)),
				g.Iff(n.Button != nil, func() g.Node {
					return h.A(h.Class("cui-btn"), h.Data("size", "sm"), h.Href(n.Button.Href), g.Text(n.Button.Label))
				}),
			),
		),
	).Render(w)
}
