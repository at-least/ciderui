package site

import (
	"io"

	"github.com/at-least/ciderui/go/internal/class"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

type Footer struct {
	// Copyright: text, or a link to the company.
	Copyright g.Node
	Class     string
}

func (f Footer) Render(w io.Writer) error {
	if f.Copyright == nil {
		panic("site.Footer needs a Copyright")
	}
	return h.Footer(class.Root("cui-footer", f.Class),
		h.Div(h.Class("cui-footer-content"),
			h.Div(h.Class("cui-footer-legal"),
				h.P(h.Class("cui-footer-copyright"), f.Copyright),
			),
		),
	).Render(w)
}
