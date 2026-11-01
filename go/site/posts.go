package site

import (
	"io"

	"github.com/at-least/ciderui/go/internal/class"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// PostSize is a [Posts] list's size (data-size). Its zero value is the
// archive's: 980px, a month's heading, rows ruled above.
type PostSize struct{ value string }

// PostSizeSmall is the related stories under an article (data-size="sm"):
// the article's text column, rows ruled below.
var PostSizeSmall = PostSize{"sm"}

// Posts is cui-posts (components/site/posts): apple.com's Newsroom archive, a
// heading over ruled rows, each a [Post]; [PostSizeSmall], the related
// stories under an article.
type Posts struct {
	// Size is the archive's, or an article's related stories: [PostSize].
	Size PostSize
	// Heading is its heading, an <h2>. Required: an empty one panics.
	Heading string
	// Class is the page's utilities on it; "" for none.
	Class string
	// Children are its rows, each an <li> with a [Post] in it.
	Children g.Node
}

// Render writes the list.
func (p Posts) Render(w io.Writer) error {
	if p.Heading == "" {
		panic("site.Posts needs a Heading")
	}
	return h.Section(class.Root("cui-posts", p.Class), g.If(p.Size.value != "", h.Data("size", p.Size.value)),
		h.H2(h.Class("cui-posts-heading"), g.Text(p.Heading)),
		h.Ul(h.Class("cui-posts-list"), p.Children),
	).Render(w)
}

// Post is cui-post (components/site/posts): a row of a [Posts] list, a link:
// its picture before its category, its headline and its date.
type Post struct {
	// Href is where it goes, the row's link. Required: an empty one panics.
	Href string
	// Media is its picture (cui-post-media on an <img>); nil for none.
	Media g.Node
	// Category is its category, over its headline; "" for none.
	Category string
	// Datetime is its date as a machine reads it, the <time>'s datetime; ""
	// with no date.
	Datetime string
	// Date is its date as the page shows it, under its headline; "" for
	// none. A Date with no Datetime panics.
	Date string
	// Headline is its headline, an <h3>. Required: an empty one panics.
	Headline string
	// Class is the page's utilities on its root; "" for none.
	Class string
}

// Render writes the row's link.
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
