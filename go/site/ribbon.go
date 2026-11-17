package site

import (
	"io"

	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

type Ribbon struct {
	// Children: a <p>.
	Children g.Node
}

func (r Ribbon) Render(w io.Writer) error {
	return h.Div(h.Class("cui-ribbon"),
		h.Div(h.Class("cui-ribbon-content"), r.Children),
	).Render(w)
}
