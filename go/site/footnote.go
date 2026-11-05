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
	Mark string
	// AriaLabel: what it reads as ("Footnote 1"), in the page's language.
	AriaLabel string
	Class     string
}

func (f Footnote) Render(w io.Writer) error {
	if f.Mark == "" {
		panic("site.Footnote needs a Mark")
	}
	if f.AriaLabel == "" {
		panic("site.Footnote needs an AriaLabel")
	}
	return h.Sup(class.Root("cui-footnote", f.Class), h.A(h.Href(f.Href), h.Aria("label", f.AriaLabel), g.Text(f.Mark))).Render(w)
}
