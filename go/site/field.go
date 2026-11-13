package site

import (
	"io"
	"strings"

	"github.com/at-least/ciderui/go/internal/attr"
	"github.com/at-least/ciderui/go/internal/glyph"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

type FieldType struct{ value string }

var (
	FieldTypeText  = FieldType{value: ""}
	FieldTypeEmail = FieldType{"email"}
	FieldTypeTel   = FieldType{"tel"}
)

// Field is apple.com's text field: its input's id is its Name.
type Field struct {
	Type         FieldType
	Name         string
	Label        string
	Autocomplete string
	Value        string
	// Error: the message under the field, wired to its input; implies Invalid.
	Error string
	// Invalid: the red edge with no message to show (a value too short to
	// judge).
	Invalid bool
	// Help: a line under the field that says what to write, wired to its
	// input.
	Help string
	// Prefix: a currency or a country code ($, +1).
	Prefix string
	// Attrs: nodes appended to the <input> (disabled, hx-post, inputmode, pattern, required).
	Attrs g.Node
}

func (f Field) Render(w io.Writer) error {
	if f.Name == "" {
		panic("site.Field needs a Name")
	}
	if f.Label == "" {
		panic("site.Field needs a Label")
	}
	errorID := f.Name + "-error"
	helpID := f.Name + "-help"
	prefixID := f.Name + "-prefix"
	kind := f.Type.value
	if kind == "" {
		kind = "text"
	}
	var ids []string
	if f.Prefix != "" {
		ids = append(ids, prefixID)
	}
	if f.Help != "" {
		ids = append(ids, helpID)
	}
	if f.Error != "" {
		ids = append(ids, errorID)
	}
	described := strings.Join(ids, " ")
	return g.Group{
		h.Div(h.Class("cui-field"),
			// @trap: placeholder " ": the label reads :placeholder-shown to tell the field empty
			h.Input(h.Class("cui-field-input"), h.ID(f.Name), h.Type(kind), h.Name(f.Name), h.Placeholder(" "),
				attr.Optional("autocomplete", f.Autocomplete),
				attr.Optional("value", f.Value),
				g.If(f.Error != "" || f.Invalid, h.Aria("invalid", "true")),
				attr.Optional("aria-describedby", described),
				f.Attrs,
			),
			h.Label(h.Class("cui-field-label"), h.For(f.Name), g.Text(f.Label)),
			g.If(f.Prefix != "", h.Span(h.Class("cui-field-prefix"), h.ID(prefixID), g.Text(f.Prefix))),
		),
		g.If(f.Help != "", h.P(h.Class("cui-field-help"), h.ID(helpID), g.Text(f.Help))),
		g.If(f.Error != "", h.P(h.Class("cui-field-error"), h.ID(errorID), glyph.Alert, g.Text(f.Error))),
	}.Render(w)
}
