package docs

import (
	"io"

	"github.com/at-least/ciderui/go/icon"
	"github.com/at-least/ciderui/go/internal/class"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// DocsTopic is a page a [DocsTopics] links to.
type DocsTopic struct {
	// Label is its name, the link's words. Required: an empty one panics.
	Label string
	// Href is its address.
	Href string
	// Abstract is what it is about, under its link; "" for none.
	Abstract string
	// Collection is whether it is a page of pages: a collection's glyph
	// before it, an article's where not.
	Collection bool
}

// DocsTopicsGroup is a group of a [DocsTopics]: its title over its pages.
type DocsTopicsGroup struct {
	// Heading is its title; "" for none.
	Heading string
	// Topics are its pages.
	Topics []DocsTopic
}

// DocsTopics is cui-docs-topics (components/docs/docs-topics): Apple's
// developer documentation's Topics and See Also, after the content: its title
// over groups of links, each after its glyph, each with its abstract under it.
type DocsTopics struct {
	// Title is its title (Topics, See Also). Required: an empty one panics.
	Title string
	// Groups are its groups.
	Groups []DocsTopicsGroup
	// Class is the page's utilities on the section; "" for none.
	Class string
}

// Render writes the topics.
func (t DocsTopics) Render(w io.Writer) error {
	if t.Title == "" {
		panic("docs.DocsTopics needs a Title")
	}
	for _, group := range t.Groups {
		for _, topic := range group.Topics {
			if topic.Label == "" {
				panic("docs.DocsTopics: a topic needs a Label")
			}
		}
	}
	return h.Section(class.Root("cui-docs-topics", t.Class),
		h.H2(h.Class("cui-docs-topics-title"), g.Text(t.Title)),
		g.Map(t.Groups, func(group DocsTopicsGroup) g.Node {
			return h.Div(h.Class("cui-docs-topics-group"),
				g.If(group.Heading != "", h.H3(h.Class("cui-docs-topics-heading"), g.Text(group.Heading))),
				g.Map(group.Topics, func(topic DocsTopic) g.Node {
					glyph := icon.DocsArticle
					if topic.Collection {
						glyph = icon.DocsCollection
					}
					return h.Div(h.Class("cui-docs-topics-item"),
						h.A(h.Class("cui-docs-topics-link"), h.Href(topic.Href),
							glyph.Class("cui-docs-topics-icon"),
							h.Span(g.Text(topic.Label)),
						),
						g.If(topic.Abstract != "", h.P(h.Class("cui-docs-topics-abstract"), g.Text(topic.Abstract))),
					)
				}),
			)
		}),
	).Render(w)
}
