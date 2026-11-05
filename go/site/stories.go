package site

import (
	"io"

	"github.com/at-least/ciderui/go/internal/class"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

type StorySize struct{ value string }

var (
	// StorySizeSmall: three across.
	StorySizeSmall = StorySize{value: ""}
	// StorySizeMedium: two across.
	StorySizeMedium = StorySize{"md"}
	// StorySizeLarge: one across.
	StorySizeLarge = StorySize{"lg"}
)

type Stories struct {
	Heading string
	Class   string
	Items   []Story
}

func (s Stories) Render(w io.Writer) error {
	if s.Heading == "" {
		panic("site.Stories needs a Heading")
	}
	return h.Section(class.Root("cui-stories", s.Class),
		h.Div(h.Class("cui-stories-content"),
			h.H2(h.Class("cui-stories-heading"), g.Text(s.Heading)),
			h.Ul(h.Class("cui-stories-list"), g.Map(s.Items, func(story Story) g.Node { return h.Li(story.node()) })),
		),
	).Render(w)
}

type Story struct {
	Size StorySize
	Href string
	// Media: an <img> carrying cui-story-media.
	Media    g.Node
	Category string
	Date     *Date
	Headline string
	Class    string
}

func (s Story) node() g.Node {
	if s.Href == "" {
		panic("site.Story needs an Href")
	}
	if s.Headline == "" {
		panic("site.Story needs a Headline")
	}
	return h.A(class.Root("cui-story", s.Class), g.If(s.Size.value != "", h.Data("size", s.Size.value)), h.Href(s.Href),
		s.Media,
		h.Div(h.Class("cui-story-content"),
			h.Div(h.Class("cui-story-head"),
				g.If(s.Category != "", h.P(h.Class("cui-story-category"), g.Text(s.Category))),
				h.H3(h.Class("cui-story-headline"), g.Text(s.Headline)),
			),
			s.Date.node("cui-story-date"),
		),
	)
}
