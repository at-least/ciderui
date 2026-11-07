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
	// Error: the message under the field, wired to its input; implies invalid.
	Error string
	// Invalid: the red edge with no message to show (a value too short to
	// judge); the input writes aria-invalid when set or Error is.
	Invalid bool
	// Help: a line under the field that says what to write, wired to its
	// input; "" for none.
	Help string
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
	help := f.ID + "-help"
	prefix := f.ID + "-prefix"
	var ids []string
	if f.Prefix != "" {
		ids = append(ids, prefix)
	}
	if f.Help != "" {
		ids = append(ids, help)
	}
	if f.Error != "" {
		ids = append(ids, message)
	}
	described := strings.Join(ids, " ")
	return g.Group{
		h.Div(class.Root("cui-field", f.Class),
			// @trap: placeholder " ": the label reads :placeholder-shown to tell the field empty
			h.Input(h.Class("cui-field-input"), h.ID(f.ID), h.Type(f.Type), h.Name(f.Name), h.Placeholder(" "),
				g.If(f.Autocomplete != "", h.AutoComplete(f.Autocomplete)),
				g.If(f.Value != "", h.Value(f.Value)),
				g.If(f.Error != "" || f.Invalid, h.Aria("invalid", "true")),
				g.If(described != "", h.Aria("describedby", described)),
				f.Attrs,
			),
			h.Label(h.Class("cui-field-label"), h.For(f.ID), g.Text(f.Label)),
			g.If(f.Prefix != "", h.Span(h.Class("cui-field-prefix"), h.ID(prefix), g.Text(f.Prefix))),
		),
		g.If(f.Help != "", h.P(h.Class("cui-field-help"), h.ID(help), g.Text(f.Help))),
		g.If(f.Error != "", h.P(h.Class("cui-field-message"), h.ID(message), glyph.Alert, g.Text(f.Error))),
	}.Render(w)
}
