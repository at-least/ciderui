package docs

import (
	"io"

	"github.com/at-least/ciderui/go/icon"
	"github.com/at-least/ciderui/go/internal/class"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

type DocsTopic struct {
	Label    string
	Href     string
	Abstract string
	// Collection: the page is a page of pages.
	Collection bool
}

type DocsTopicsGroup struct {
	Heading string
	Topics  []DocsTopic
}

// DocsTopics goes after the content.
type DocsTopics struct {
	// Title: "Topics" or "See Also".
	Title  string
	Groups []DocsTopicsGroup
	Class  string
}

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
