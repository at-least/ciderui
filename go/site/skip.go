package site

import (
	"io"

	"github.com/at-least/ciderui/go/internal/class"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// SkipLink is unseen until the keyboard reaches it, the page's first stop.
type SkipLink struct {
	// Href: the page's main content (#main).
	Href  string
	Label string
	Class string
}

func (l SkipLink) Render(w io.Writer) error {
	if l.Href == "" {
		panic("site.SkipLink needs an Href")
	}
	if l.Label == "" {
		panic("site.SkipLink needs a Label")
	}
	return h.A(class.Root("cui-skip cui-btn", l.Class), h.Href(l.Href), g.Text(l.Label)).Render(w)
}
