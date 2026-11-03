package docs

import (
	"io"

	"github.com/at-least/ciderui/go/icon"
	"github.com/at-least/ciderui/go/internal/class"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// DocsBar sits under the site's global nav.
type DocsBar struct {
	// Title is the documentation's title.
	Title string
	// Href is the documentation's first page.
	Href string
	// Navigator is the id of the navigator popover the bar's button opens
	// below 1024px; "" for no navigator and no button.
	Navigator string
	Class     string
}

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
