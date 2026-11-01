package site

import (
	"io"

	"github.com/at-least/ciderui/go/internal/class"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// Article is cui-article (components/site/article): apple.com's Newsroom
// story, its header (an eyebrow of its category and its date, the headline, a
// subhead) over its copy (Body, the plain elements markdown writes, its
// footnotes a cui-article-footnotes section at its end).
type Article struct {
	// Category is its category, over its date (cui-article-category); ""
	// for none.
	Category string
	// Datetime is its date as a machine reads it (2026-10-19), the <time>'s
	// datetime; "" with no date.
	Datetime string
	// Date is its date as the page shows it (October 19, 2026); "" for none.
	// A Date with no Datetime panics.
	Date string
	// Headline is the headline, the page's <h1>. Required: an empty one
	// panics.
	Headline string
	// Subhead is the subhead under it; "" for none.
	Subhead string
	// Body is the copy: paragraphs, subheads, lists, figures, its footnotes.
	Body g.Node
	// Class is the page's utilities on the article; "" for none.
	Class string
}

// Render writes the article.
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
