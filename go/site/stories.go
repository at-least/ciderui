package site

import (
	"io"

	"github.com/at-least/ciderui/go/internal/class"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// StorySize is a [Story]'s size (data-size): how many of it go across the
// grid. Its zero value is three across, four of the grid's twelve columns.
type StorySize struct{ value string }

var (
	// StorySizeMedium is two across (data-size="md"), six columns.
	StorySizeMedium = StorySize{"md"}
	// StorySizeLarge is one across (data-size="lg"), its picture beside its
	// text from 1069px.
	StorySizeLarge = StorySize{"lg"}
)

// Stories is cui-stories (components/site/stories): apple.com's Newsroom
// index, a band on the alternate ground, its heading over a grid of [Story]
// tiles.
type Stories struct {
	// Heading is its heading, an <h2>. Required: an empty one panics.
	Heading string
	// Class is the page's utilities on the band; "" for none.
	Class string
	// Children are its stories, each an <li> with a [Story] in it.
	Children g.Node
}

// Render writes the band.
func (s Stories) Render(w io.Writer) error {
	if s.Heading == "" {
		panic("site.Stories needs a Heading")
	}
	return h.Section(class.Root("cui-stories", s.Class),
		h.Div(h.Class("cui-stories-content"),
			h.H2(h.Class("cui-stories-heading"), g.Text(s.Heading)),
			h.Ul(h.Class("cui-stories-list"), s.Children),
		),
	).Render(w)
}

// Story is cui-story (components/site/stories): a story in a [Stories] grid,
// a white rounded tile, the whole of it its link: its picture, its category,
// its headline, its date.
type Story struct {
	// Size is how many go across: [StorySize].
	Size StorySize
	// Href is where it goes, the whole tile's link. Required: an empty one
	// panics.
	Href string
	// Media is its picture (cui-story-media on an <img>); nil for none.
	Media g.Node
	// Category is its category, over its headline; "" for none.
	Category string
	// Datetime is its date as a machine reads it, the <time>'s datetime; ""
	// with no date.
	Datetime string
	// Date is its date as the page shows it, at the tile's foot; "" for
	// none. A Date with no Datetime panics.
	Date string
	// Headline is its headline, an <h3>. Required: an empty one panics.
	Headline string
	// Class is the page's utilities on its root; "" for none.
	Class string
}

// Render writes the story's link.
func (s Story) Render(w io.Writer) error {
	if s.Href == "" {
		panic("site.Story needs an Href")
	}
	if s.Headline == "" {
		panic("site.Story needs a Headline")
	}
	if s.Date != "" && s.Datetime == "" {
		panic("site.Story's Date needs its Datetime")
	}
	return h.A(class.Root("cui-story", s.Class), g.If(s.Size.value != "", h.Data("size", s.Size.value)), h.Href(s.Href),
		s.Media,
		h.Div(h.Class("cui-story-content"),
			h.Div(h.Class("cui-story-head"),
				g.If(s.Category != "", h.P(h.Class("cui-story-category"), g.Text(s.Category))),
				h.H3(h.Class("cui-story-headline"), g.Text(s.Headline)),
			),
			g.If(s.Date != "", h.Time(h.Class("cui-story-date"), h.DateTime(s.Datetime), g.Text(s.Date))),
		),
	).Render(w)
}
