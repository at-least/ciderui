package site

import (
	"io"

	"github.com/at-least/ciderui/go/internal/class"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// QuickLinks is cui-quicklinks (components/site/quicklinks): the store's row
// of pill links (cui-btn, data-variant="secondary-neutral") under a title, as
// far in as a shelf's header.
type QuickLinks struct {
	// ID makes the title's id, the nav's aria-labelledby (<id>-title).
	// Required: an empty one panics.
	ID string
	// Title is the title. Required: an empty one panics.
	Title string
	// Links are the links.
	Links []Link
	// Class is the page's utilities on the nav; "" for none.
	Class string
}

// Render writes the nav.
func (q QuickLinks) Render(w io.Writer) error {
	if q.ID == "" {
		panic("site.QuickLinks needs an ID")
	}
	if q.Title == "" {
		panic("site.QuickLinks needs a Title")
	}
	titleID := q.ID + "-title"
	return h.Nav(class.Root("cui-quicklinks", q.Class), h.Aria("labelledby", titleID),
		h.H2(h.Class("cui-quicklinks-title"), h.ID(titleID), g.Text(q.Title)),
		h.Ul(h.Class("cui-quicklinks-list"), g.Map(q.Links, func(link Link) g.Node {
			return h.Li(h.A(h.Class("cui-btn"), h.Data("variant", "secondary-neutral"), h.Href(link.Href), g.Text(link.Text)))
		})),
	).Render(w)
}
