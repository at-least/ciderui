package console

import (
	"io"

	"github.com/at-least/ciderui/go/lucide"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

type ConsoleSectionHeaderLevel struct{ value string }

var (
	// ConsoleSectionHeaderLevelPage heads a whole page: <h1 class="cui-console-h2">.
	ConsoleSectionHeaderLevelPage = ConsoleSectionHeaderLevel{"page"}
	// ConsoleSectionHeaderLevelSection heads one of a page's sections:
	// <h2 class="cui-console-h4-emph">.
	ConsoleSectionHeaderLevelSection = ConsoleSectionHeaderLevel{"section"}
)

// ConsoleSectionHeaderBack is the way back to the page a page is under, over
// its title, named by that page (iOS's back button).
type ConsoleSectionHeaderBack struct {
	Href  string
	Label string
}

// ConsoleSectionHeader is a page's or a section's title, the line under it,
// and its controls at the row's end.
type ConsoleSectionHeader struct {
	Level ConsoleSectionHeaderLevel
	// Back: the zero value for none; a page's alone.
	Back  ConsoleSectionHeaderBack
	Title string
	// Detail: the line under the title, what the page or section is for; ""
	// for none.
	Detail string
	// Children: the controls; nil for none.
	Children g.Node
}

func (s ConsoleSectionHeader) Render(w io.Writer) error {
	if s.Title == "" {
		panic("console.ConsoleSectionHeader needs a Title")
	}
	if (s.Back.Href == "") != (s.Back.Label == "") {
		panic("console.ConsoleSectionHeader needs a Back with both an Href and a Label, or neither")
	}
	if s.Back.Href != "" && s.Level != ConsoleSectionHeaderLevelPage {
		panic("console.ConsoleSectionHeader's Back is a page's: a section is on the page it belongs to")
	}
	var title g.Node
	switch s.Level {
	case ConsoleSectionHeaderLevelPage:
		title = h.H1(h.Class("cui-console-h2"), g.Text(s.Title))
	case ConsoleSectionHeaderLevelSection:
		title = h.H2(h.Class("cui-console-h4-emph"), g.Text(s.Title))
	default:
		panic("console.ConsoleSectionHeader needs a Level: ConsoleSectionHeaderLevelPage (a whole page) or ConsoleSectionHeaderLevelSection (one of its sections)")
	}
	return h.Header(h.Class("cui-console-section-header"),
		h.Div(
			g.If(s.Back.Href != "", h.A(h.Class("cui-console-link cui-console-section-header-back"), h.Href(s.Back.Href), lucide.ChevronLeft, g.Text(s.Back.Label))),
			title, g.If(s.Detail != "", h.P(h.Class("cui-console-section-header-detail"), g.Text(s.Detail)))),
		s.Children,
	).Render(w)
}
