package site

import (
	"io"

	"github.com/at-least/ciderui/go/internal/class"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

type Posts struct {
	// Small: the related stories under an article; else the Newsroom archive.
	Small   bool
	Heading string
	Class   string
	Items   []Post
}

func (p Posts) Render(w io.Writer) error {
	if p.Heading == "" {
		panic("site.Posts needs a Heading")
	}
	return h.Section(class.Root("cui-posts", p.Class), g.If(p.Small, h.Data("size", "sm")),
		h.H2(h.Class("cui-posts-heading"), g.Text(p.Heading)),
		h.Ul(h.Class("cui-posts-list"), g.Map(p.Items, func(post Post) g.Node { return h.Li(post.node()) })),
	).Render(w)
}

type Post struct {
	Href string
	// Media: an <img> carrying cui-post-media.
	Media    g.Node
	Category string
	Date     *Date
	Headline string
	Class    string
}

func (p Post) node() g.Node {
	if p.Href == "" {
		panic("site.Post needs an Href")
	}
	if p.Headline == "" {
		panic("site.Post needs a Headline")
	}
	return h.A(class.Root("cui-post", p.Class), h.Href(p.Href),
		p.Media,
		h.Div(h.Class("cui-post-content"),
			g.If(p.Category != "", h.P(h.Class("cui-post-category"), g.Text(p.Category))),
			h.H3(h.Class("cui-post-headline"), g.Text(p.Headline)),
			p.Date.node("cui-post-date"),
		),
	)
}
