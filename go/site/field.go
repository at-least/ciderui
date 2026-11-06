package site

import (
	"io"
	"strings"

	"github.com/at-least/ciderui/go/internal/class"
	"github.com/at-least/ciderui/go/internal/glyph"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

type Field struct {
	ID           string
	Type         string
	Name         string
	Label        string
	Autocomplete string
	Value        string
	Error        string
	// Prefix: a currency or a country code ($, +1).
	Prefix string
	// Attrs: nodes appended to the <input> (disabled, hx-post, inputmode, pattern, required); nil for none.
	Attrs g.Node
	Class string
}

func (f Field) Render(w io.Writer) error {
	if f.ID == "" {
		panic("site.Field needs an ID")
	}
	if f.Name == "" {
		panic("site.Field needs a Name")
	}
	if f.Label == "" {
		panic("site.Field needs a Label")
	}
	message := f.ID + "-message"
	prefix := f.ID + "-prefix"
	var ids []string
	if f.Prefix != "" {
		ids = append(ids, prefix)
	}
	if f.Error != "" {
		ids = append(ids, message)
	}
	described := strings.Join(ids, " ")
	return g.Group{
		h.Div(class.Root("cui-field", f.Class),
			// placeholder " ": the label reads :placeholder-shown to tell the field empty
			h.Input(h.Class("cui-field-input"), h.ID(f.ID), h.Type(f.Type), h.Name(f.Name), h.Placeholder(" "),
				g.If(f.Autocomplete != "", h.AutoComplete(f.Autocomplete)),
				g.If(f.Value != "", h.Value(f.Value)),
				g.If(f.Error != "", h.Aria("invalid", "true")),
				g.If(described != "", h.Aria("describedby", described)),
				f.Attrs,
			),
			h.Label(h.Class("cui-field-label"), h.For(f.ID), g.Text(f.Label)),
			g.If(f.Prefix != "", h.Span(h.Class("cui-field-prefix"), h.ID(prefix), g.Text(f.Prefix))),
		),
		g.If(f.Error != "", h.P(h.Class("cui-field-message"), h.ID(message), glyph.Alert, g.Text(f.Error))),
	}.Render(w)
}
