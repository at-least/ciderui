package site

import (
	"io"

	"github.com/at-least/ciderui/go/internal/class"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// OptionTile is apple.com's buy-page choice. Not Option: that is a
// [Select]'s choice.
type OptionTile struct {
	ID      string
	Name    string
	Value   string
	Title   string
	Detail  string
	Checked bool
	// Attrs: nodes appended to the <input> (disabled, required); nil for none.
	Attrs g.Node
	Class string
}

func (o OptionTile) Render(w io.Writer) error {
	if o.ID == "" {
		panic("site.OptionTile needs an ID")
	}
	if o.Name == "" {
		panic("site.OptionTile needs a Name")
	}
	if o.Title == "" {
		panic("site.OptionTile needs a Title")
	}
	return h.Div(class.Root("cui-option", o.Class),
		h.Input(h.Class("cui-option-input"), h.Type("radio"), h.ID(o.ID), h.Name(o.Name), h.Value(o.Value), g.If(o.Checked, h.Checked()), o.Attrs),
		h.Label(h.Class("cui-option-label"), h.For(o.ID),
			h.Span(h.Class("cui-option-title"),
				g.Text(o.Title),
				g.If(o.Detail != "", h.Span(h.Class("cui-option-detail"), g.Text(o.Detail))),
			),
		),
	).Render(w)
}
