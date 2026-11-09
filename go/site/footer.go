package site

import (
	"io"

	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

type Footer struct {
	// Copyright: text, or a link to the company.
	Copyright g.Node
}

func (f Footer) Render(w io.Writer) error {
	if f.Copyright == nil {
		panic("site.Footer needs a Copyright")
	}
	return h.Footer(h.Class("cui-footer"),
		h.Div(h.Class("cui-footer-content"),
			h.Div(h.Class("cui-footer-legal"),
				h.P(h.Class("cui-footer-copyright"), f.Copyright),
			),
		),
	).Render(w)
}
