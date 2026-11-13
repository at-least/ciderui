package site

import (
	"io"

	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// Option is apple.com's buy-page choice, a radio drawn as a tile: its
// input's id is Name-Value.
type Option struct {
	Name    string
	Value   string
	Title   string
	Detail  string
	Checked bool
	// Attrs: nodes appended to the <input> (disabled, required).
	Attrs g.Node
}

func (o Option) Render(w io.Writer) error {
	if o.Name == "" {
		panic("site.Option needs a Name")
	}
	if o.Value == "" {
		panic("site.Option needs a Value")
	}
	if o.Title == "" {
		panic("site.Option needs a Title")
	}
	id := o.Name + "-" + o.Value
	return h.Div(h.Class("cui-option"),
		h.Input(h.Class("cui-option-input"), h.Type("radio"), h.ID(id), h.Name(o.Name), h.Value(o.Value), g.If(o.Checked, h.Checked()), o.Attrs),
		h.Label(h.Class("cui-option-label"), h.For(id),
			h.Span(h.Class("cui-option-title"),
				g.Text(o.Title),
				g.If(o.Detail != "", h.Span(h.Class("cui-option-detail"), g.Text(o.Detail))),
			),
		),
	).Render(w)
}
