package docs

import (
	"io"

	"github.com/at-least/ciderui/go/icon"
	"github.com/at-least/ciderui/go/internal/class"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// DocsBar is cui-docs-bar (components/docs/docs-bar): the documentation's bar
// in Apple's developer documentation, under the global nav: the button that
// opens the navigator below 1024px (its popovertarget, the navigator's id),
// then the documentation's title, a link home.
type DocsBar struct {
	// Title is the documentation's title, its link's words and the bar's
	// name. Required: an empty one panics.
	Title string
	// Href is where the title links: the documentation's first page.
	Href string
	// Navigator is the navigator's id, which the button opens; "" for no
	// navigator, and no button.
	Navigator string
	// Class is the page's utilities on the bar; "" for none.
	Class string
}

// Render writes the bar.
func (b DocsBar) Render(w io.Writer) error {
	if b.Title == "" {
		panic("docs.DocsBar needs a Title")
	}
	return h.Nav(class.Root("cui-docs-bar", b.Class), h.Aria("label", b.Title),
		g.If(b.Navigator != "",
			h.Button(h.Class("cui-docs-bar-toggle"), h.Type("button"), h.PopoverTarget(b.Navigator), h.Aria("label", "Navigator"), icon.DocsSidebar),
		),
		h.A(h.Class("cui-docs-bar-title"), h.Href(b.Href), g.Text(b.Title)),
	).Render(w)
}
