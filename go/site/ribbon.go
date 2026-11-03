package site

import (
	"io"

	"github.com/at-least/ciderui/go/internal/class"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

type Ribbon struct {
	Class string
	// Children: a <p>, a [MoreLink] in it.
	Children g.Node
}

func (r Ribbon) Render(w io.Writer) error {
	return h.Div(class.Root("cui-ribbon", r.Class),
		h.Div(h.Class("cui-ribbon-content"), r.Children),
	).Render(w)
}
