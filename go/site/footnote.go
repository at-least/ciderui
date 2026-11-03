package site

import (
	"io"

	"github.com/at-least/ciderui/go/internal/class"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

type Footnote struct {
	Href string
	// Mark: a number or a symbol (◊, ※).
	Mark  string
	Class string
}

func (f Footnote) Render(w io.Writer) error {
	if f.Mark == "" {
		panic("site.Footnote needs a Mark")
	}
	return h.Sup(class.Root("cui-footnote", f.Class), h.A(h.Href(f.Href), h.Aria("label", "Footnote "+f.Mark), g.Text(f.Mark))).Render(w)
}
