package site

import (
	"io"

	"github.com/at-least/ciderui/go/icon"
	"github.com/at-least/ciderui/go/internal/class"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// the menu's checkbox: a page has one local nav
const localNavToggleID = "cui-localnav-menu"

// LocalNav is cui-localnav (components/site/localnav): the bar under the
// global nav on a product's pages, stuck to the top as the page scrolls: the
// product's name, its pages' links (the current one marked) and a button; to
// 833px the links are a menu its chevron opens below the bar, with no script
// (a checkbox, the chevron its label). No links, no menu.
type LocalNav struct {
	// Title is the product's name. Required: an empty one panics.
	Title string
	// Href is where the name goes, the product's main page. Required: an
	// empty one panics.
	Href string
	// Links are the product's pages, in order; none, no menu.
	Links []Link
	// Current is the current page's href: its link is marked
	// (aria-current="page"); "" for none.
	Current string
	// Button is the button after the links: a cui-btn, data-size="sm"; nil
	// for none.
	Button *Link
	// Dark is whether the nav is dark (cui-dark).
	Dark bool
	// Class is the page's utilities on the nav; "" for none.
	Class string
}

// Render writes the local nav.
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
				g.If(menu, h.Label(h.Class("cui-localnav-menubutton"), h.For(localNavToggleID), h.Aria("hidden", "true"), icon.Chevron)),
				g.Iff(n.Button != nil, func() g.Node {
					return h.A(h.Class("cui-btn"), h.Data("size", "sm"), h.Href(n.Button.Href), g.Text(n.Button.Text))
				}),
			),
		),
	).Render(w)
}
