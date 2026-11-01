package site

import (
	"io"

	"github.com/at-least/ciderui/go/internal/class"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// Footnote is cui-footnote (components/site/footnote): a footnote's mark, a
// link to its note at the text's size and on its baseline, its digits the
// font's numerators.
type Footnote struct {
	// Href is the note's link (#footnote-1).
	Href string
	// Mark is the mark: a number, or a symbol (◊, ※). Required: an empty one
	// panics.
	Mark string
	// Class is the page's utilities on its root; "" for none.
	Class string
}

// Render writes the footnote's mark.
func (f Footnote) Render(w io.Writer) error {
	if f.Mark == "" {
		panic("site.Footnote needs a Mark")
	}
	return h.Sup(class.Root("cui-footnote", f.Class), h.A(h.Href(f.Href), h.Aria("label", "Footnote "+f.Mark), g.Text(f.Mark))).Render(w)
}
