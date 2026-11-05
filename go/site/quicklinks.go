package site

import (
	"io"

	"github.com/at-least/ciderui/go/internal/class"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

type QuickLinks struct {
	ID    string
	Title string
	Links []Link
	Class string
}

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
			return h.Li(h.A(h.Class("cui-btn"), h.Data("variant", "secondary-neutral"), h.Href(link.Href), g.Text(link.Label)))
		})),
	).Render(w)
}
