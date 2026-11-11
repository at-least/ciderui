package site

import (
	"io"

	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// List is Apple Developer Forums' post list on a site's page: rows ruled
// apart, each a linked title, a line or two of its text, and an aside (its
// counts, its date) at the row's end, under the text on a phone.
type List struct {
	Rows []ListRow
}

type ListRow struct {
	Title string
	Href  string
	// Text: two lines at most; the rest is cut with an ellipsis.
	Text  string
	Aside g.Node
}

func (l List) Render(w io.Writer) error {
	if len(l.Rows) == 0 {
		panic("site.List needs Rows: a page with none says so in its own words")
	}
	return h.Ul(h.Class("cui-list"), g.Map(l.Rows, ListRow.node)).Render(w)
}

func (r ListRow) node() g.Node {
	if r.Title == "" || r.Href == "" {
		panic("site.ListRow needs a Title and an Href")
	}
	return h.Li(h.Class("cui-list-row"),
		h.Div(
			h.A(h.Class("cui-list-title"), h.Href(r.Href), g.Text(r.Title)),
			g.If(r.Text != "", h.P(h.Class("cui-list-text"), g.Text(r.Text))),
		),
		g.If(r.Aside != nil, h.Div(h.Class("cui-list-aside"), r.Aside)),
	)
}
