package site

import (
	"io"

	"github.com/at-least/ciderui/go/internal/class"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// Ground's zero value: the page's ground.
type Ground struct{ value string }

var GroundAlt = Ground{"alt"}

// Rhythm's zero value: apple.com's iPhone page.
type Rhythm struct{ value string }

// RhythmLarge: apple.com's iPhone 18 Pro page.
var RhythmLarge = Rhythm{"lg"}

type Section struct {
	Ground   Ground
	Rhythm   Rhythm
	Dark     bool
	Class    string
	Children g.Node
}

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

type SectionHeader struct {
	Headline string
	Links    []Link
	Class    string
}

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
