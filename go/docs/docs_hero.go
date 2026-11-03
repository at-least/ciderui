package docs

import (
	"io"

	"github.com/at-least/ciderui/go/internal/class"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

type DocsCrumb struct {
	Label string
	Href  string
}

// DocsHero is the top of a page.
type DocsHero struct {
	// Crumbs are the pages above this one, from the top; the page's own
	// title is the current crumb, written from Title.
	Crumbs []DocsCrumb
	// Eyebrow is the kind of page ("Article", "Component").
	Eyebrow  string
	Title    string
	Abstract string
	Class    string
}

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
