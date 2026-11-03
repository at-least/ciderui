package site

import (
	"io"

	"github.com/at-least/ciderui/go/internal/class"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// PostSize's zero value is the Newsroom archive.
type PostSize struct{ value string }

// PostSizeSmall: the related stories under an article.
var PostSizeSmall = PostSize{"sm"}

type Posts struct {
	Size    PostSize
	Heading string
	Class   string
	// Children: <li>s, each holding a [Post].
	Children g.Node
}

func (p Posts) Render(w io.Writer) error {
	if p.Heading == "" {
		panic("site.Posts needs a Heading")
	}
	return h.Section(class.Root("cui-posts", p.Class), g.If(p.Size.value != "", h.Data("size", p.Size.value)),
		h.H2(h.Class("cui-posts-heading"), g.Text(p.Heading)),
		h.Ul(h.Class("cui-posts-list"), p.Children),
	).Render(w)
}

type Post struct {
	Href string
	// Media: an <img> carrying cui-post-media.
	Media    g.Node
	Category string
	// Datetime: machine-readable (2026-10-19).
	Datetime string
	Date     string
	Headline string
	Class    string
}

func (p Post) Render(w io.Writer) error {
	if p.Href == "" {
		panic("site.Post needs an Href")
	}
	if p.Headline == "" {
		panic("site.Post needs a Headline")
	}
	if p.Date != "" && p.Datetime == "" {
		panic("site.Post's Date needs its Datetime")
	}
	return h.A(class.Root("cui-post", p.Class), h.Href(p.Href),
		p.Media,
		h.Div(h.Class("cui-post-content"),
			g.If(p.Category != "", h.P(h.Class("cui-post-category"), g.Text(p.Category))),
			h.H3(h.Class("cui-post-headline"), g.Text(p.Headline)),
			g.If(p.Date != "", h.Time(h.Class("cui-post-date"), h.DateTime(p.Datetime), g.Text(p.Date))),
		),
	).Render(w)
}
