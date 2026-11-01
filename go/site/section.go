package site

import (
	"io"

	"github.com/at-least/ciderui/go/internal/class"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// Ground is a [Section]'s ground (data-variant). Its zero value is no ground
// of its own: the page's.
type Ground struct{ value string }

// GroundAlt is the alternate ground (data-variant="alt"), dark in a dark
// block.
var GroundAlt = Ground{"alt"}

// Rhythm is a [Section]'s rhythm, its space above and below (data-size). Its
// zero value is the iPhone page's: 56px, 96px from 735px, 112px from 1069px.
type Rhythm struct{ value string }

// RhythmLarge is iPhone 18 Pro's (data-size="lg"): 112px, 120px, 144px.
var RhythmLarge = Rhythm{"lg"}

// Section is cui-section (components/site/section): a band of the page, its
// ground, its rhythm and, with Dark, a dark block (cui-dark), its content
// inside cui-section-content (87.5% of the window, centered).
type Section struct {
	// Ground is its ground.
	Ground Ground
	// Rhythm is its rhythm.
	Rhythm Rhythm
	// Dark is whether it is a dark block.
	Dark bool
	// Class is the page's utilities on the section; "" for none.
	Class string
	// Children is the content.
	Children g.Node
}

// Render writes the section.
func (s Section) Render(w io.Writer) error {
	block := "cui-section"
	if s.Dark {
		block = "cui-section cui-dark"
	}
	return h.Section(class.Root(block, s.Class),
		g.If(s.Ground.value != "", h.Data("variant", s.Ground.value)),
		g.If(s.Rhythm.value != "", h.Data("size", s.Rhythm.value)),
		h.Div(h.Class("cui-section-content"), s.Children),
	).Render(w)
}

// SectionHeader is cui-section-header: a section's headline
// (cui-section-headline, an <h2>) and the more links beside it
// (cui-section-links), each a [MoreLink].
type SectionHeader struct {
	// Headline is the headline. Required: an empty one panics.
	Headline string
	// Links are the links beside it; none, no list.
	Links []Link
	// Class is the page's utilities on its root; "" for none.
	Class string
}

// Render writes the header.
func (s SectionHeader) Render(w io.Writer) error {
	if s.Headline == "" {
		panic("site.SectionHeader needs a Headline")
	}
	return h.Header(class.Root("cui-section-header", s.Class),
		h.H2(h.Class("cui-section-headline"), g.Text(s.Headline)),
		g.If(len(s.Links) > 0, h.Ul(h.Class("cui-section-links"), g.Map(s.Links, func(link Link) g.Node {
			return h.Li(MoreLink{Href: link.Href, Text: link.Text, Class: ""})
		}))),
	).Render(w)
}
