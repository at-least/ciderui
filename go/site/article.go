package site

import (
	"io"

	"github.com/at-least/ciderui/go/internal/class"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// Article is apple.com's Newsroom story.
type Article struct {
	Category string
	// Datetime: machine-readable (2026-10-19).
	Datetime string
	Date     string
	Headline string
	Subhead  string
	// Body: the plain elements markdown writes; footnotes go last, in a
	// cui-article-footnotes section.
	Body  g.Node
	Class string
}

func (a Article) Render(w io.Writer) error {
	if a.Headline == "" {
		panic("site.Article needs a Headline")
	}
	if a.Date != "" && a.Datetime == "" {
		panic("site.Article's Date needs its Datetime")
	}
	return h.Article(class.Root("cui-article", a.Class),
		h.Header(h.Class("cui-article-header"),
			g.If(a.Category != "" || a.Date != "", h.P(h.Class("cui-article-eyebrow"),
				g.If(a.Category != "", h.Span(h.Class("cui-article-category"), g.Text(a.Category))),
				g.If(a.Date != "", h.Time(h.Class("cui-article-date"), h.DateTime(a.Datetime), g.Text(a.Date))),
			)),
			h.H1(h.Class("cui-article-headline"), g.Text(a.Headline)),
			g.If(a.Subhead != "", h.P(h.Class("cui-article-subhead"), g.Text(a.Subhead))),
		),
		h.Div(h.Class("cui-article-body"), a.Body),
	).Render(w)
}
