package docs

import (
	"io"

	"github.com/at-least/ciderui/go/internal/class"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// DocsCrumb is a page above this one in the documentation, as a crumb of a
// [DocsHero]: its name and its address.
type DocsCrumb struct {
	// Label is its name, the link's words.
	Label string
	// Href is its address.
	Href string
}

// DocsHero is cui-docs-hero (components/docs/docs-hero): the top of a page in
// Apple's developer documentation: its breadcrumbs (the pages above it, then
// its own title, the current crumb), its eyebrow, its title and its abstract.
type DocsHero struct {
	// Crumbs are the pages above this one, from the top; nil for none, a page
	// at the top, which shows no breadcrumbs.
	Crumbs []DocsCrumb
	// Eyebrow is what kind of page it is (Article, Component); "" for none.
	Eyebrow string
	// Title is its title, the page's <h1> and its current crumb. Required: an
	// empty one panics.
	Title string
	// Abstract is what it is about, under the title; "" for none.
	Abstract string
	// Class is the page's utilities on the hero; "" for none.
	Class string
}

// Render writes the hero.
func (x DocsHero) Render(w io.Writer) error {
	if x.Title == "" {
		panic("docs.DocsHero needs a Title")
	}
	return h.Header(class.Root("cui-docs-hero", x.Class),
		g.If(len(x.Crumbs) > 0,
			h.Nav(h.Aria("label", "Breadcrumbs"),
				h.Ol(h.Class("cui-docs-hero-crumbs"),
					g.Map(x.Crumbs, func(c DocsCrumb) g.Node {
						return h.Li(h.A(h.Href(c.Href), g.Text(c.Label)))
					}),
					h.Li(h.Aria("current", "page"), g.Text(x.Title)),
				),
			),
		),
		g.If(x.Eyebrow != "", h.P(h.Class("cui-docs-hero-eyebrow"), g.Text(x.Eyebrow))),
		h.H1(h.Class("cui-docs-hero-title"), g.Text(x.Title)),
		g.If(x.Abstract != "", h.P(h.Class("cui-docs-hero-abstract"), g.Text(x.Abstract))),
	).Render(w)
}
