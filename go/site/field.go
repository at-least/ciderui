package site

import (
	"io"
	"slices"
	"strings"

	"github.com/at-least/ciderui/go/icon"
	"github.com/at-least/ciderui/go/internal/class"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// Field is cui-field (components/site/field): apple.com's 56px text field,
// its label the placeholder until the field is focused or filled, then small
// at its top (read off the input: its placeholder is a space, so the label can
// tell it empty); in error with its message under it, disabled, or with a
// currency or a country code before it.
type Field struct {
	// ID is the input's id, its label's for. Required: an empty one panics.
	ID string
	// Type is the input's type (text, email, tel, password, …).
	Type string
	// Name is the input's name.
	Name string
	// Label is the label. Required: an empty one panics.
	Label string
	// Autocomplete is the input's autocomplete; "" for none.
	Autocomplete string
	// Value is its value; "" for an empty field.
	Value string
	// Error is its error, the message under it (aria-invalid,
	// aria-describedby, its id the field's with -message); "" for none.
	Error string
	// Disabled is whether it is disabled.
	Disabled bool
	// Prefix is a currency or a country code before it ($, +1), which
	// describes it (its id the field's with -prefix); "" for none.
	Prefix string
	// Class is the page's utilities on the field; "" for none.
	Class string
}

// Render writes the field.
func (f Field) Render(w io.Writer) error {
	if f.ID == "" {
		panic("site.Field needs an ID")
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
			h.Input(h.Class("cui-field-input"), h.ID(f.ID), h.Type(f.Type), h.Name(f.Name), h.Placeholder(" "),
				g.If(f.Autocomplete != "", h.AutoComplete(f.Autocomplete)),
				g.If(f.Value != "", h.Value(f.Value)),
				g.If(f.Error != "", h.Aria("invalid", "true")),
				g.If(described != "", h.Aria("describedby", described)),
				g.If(f.Disabled, h.Disabled()),
			),
			h.Label(h.Class("cui-field-label"), h.For(f.ID), g.Text(f.Label)),
			g.If(f.Prefix != "", h.Span(h.Class("cui-field-prefix"), h.ID(prefix), g.Text(f.Prefix))),
		),
		g.If(f.Error != "", errorMessage(message, f.Error)),
	}.Render(w)
}

// a field's error, under it
func errorMessage(id, text string) g.Node {
	return h.P(h.Class("cui-field-message"), h.ID(id), icon.Alert, g.Text(text))
}

// Option is one of a [Select]'s choices.
type Option struct {
	// Value is the option's value. Required: the empty one is the select's
	// empty choice's.
	Value string
	// Text is what the option shows.
	Text string
}

// Select is cui-field on a select (components/site/field): the store's
// dropdown, the field's look with a chevron at its end; its empty choice
// first, chosen while nothing is, its label the placeholder until a choice is
// made.
type Select struct {
	// ID is the select's id, its label's for. Required: an empty one panics.
	ID string
	// Name is the select's name.
	Name string
	// Label is the label. Required: an empty one panics.
	Label string
	// Options are the choices after the empty one. At least one.
	Options []Option
	// Selected is the chosen option's value; "" for none (the empty choice).
	// One the options do not have panics.
	Selected string
	// Error is its error, the message under it; "" for none.
	Error string
	// Disabled is whether it is disabled.
	Disabled bool
	// Class is the page's utilities on the field; "" for none.
	Class string
}

// Render writes the select.
func (s Select) Render(w io.Writer) error {
	if s.ID == "" {
		panic("site.Select needs an ID")
	}
	if s.Label == "" {
		panic("site.Select needs a Label")
	}
	if len(s.Options) == 0 {
		panic("site.Select needs Options")
	}
	if slices.ContainsFunc(s.Options, func(o Option) bool { return o.Value == "" }) {
		panic("site.Select's Options need values: the empty one is its empty choice's")
	}
	if s.Selected != "" && !slices.ContainsFunc(s.Options, func(o Option) bool { return o.Value == s.Selected }) {
		panic("site.Select's Selected value is none of its Options'")
	}
	message := s.ID + "-message"
	return g.Group{
		h.Div(class.Root("cui-field", s.Class),
			h.Select(h.Class("cui-field-input"), h.ID(s.ID), h.Name(s.Name),
				g.If(s.Error != "", h.Aria("invalid", "true")),
				g.If(s.Error != "", h.Aria("describedby", message)),
				g.If(s.Disabled, h.Disabled()),
				h.Option(h.Value(""), g.If(s.Selected == "", h.Selected()), g.Text("Select")),
				g.Map(s.Options, func(o Option) g.Node {
					return h.Option(h.Value(o.Value), g.If(o.Value == s.Selected, h.Selected()), g.Text(o.Text))
				}),
			),
			h.Label(h.Class("cui-field-label"), h.For(s.ID), g.Text(s.Label)),
			icon.Dropdown.Class("cui-field-chevron"),
		),
		g.If(s.Error != "", errorMessage(message, s.Error)),
	}.Render(w)
}

// TextArea is cui-field on a textarea (components/site/field): a message of
// several lines, its textarea in the field's box (cui-field-box), 3 lines
// tall and taller by its handle, its label the placeholder until it is
// focused or filled.
type TextArea struct {
	// ID is the textarea's id, its label's for. Required: an empty one
	// panics.
	ID string
	// Name is the textarea's name.
	Name string
	// Label is the label. Required: an empty one panics.
	Label string
	// Value is its text; "" for an empty one.
	Value string
	// Error is its error, the message under it; "" for none.
	Error string
	// Disabled is whether it is disabled.
	Disabled bool
	// Class is the page's utilities on the field; "" for none.
	Class string
}

// Render writes the textarea.
func (a TextArea) Render(w io.Writer) error {
	if a.ID == "" {
		panic("site.TextArea needs an ID")
	}
	if a.Label == "" {
		panic("site.TextArea needs a Label")
	}
	message := a.ID + "-message"
	return g.Group{
		h.Div(class.Root("cui-field", a.Class),
			h.Div(h.Class("cui-field-box"),
				h.Textarea(h.Class("cui-field-textarea"), h.ID(a.ID), h.Name(a.Name), h.Placeholder(" "),
					g.If(a.Error != "", h.Aria("invalid", "true")),
					g.If(a.Error != "", h.Aria("describedby", message)),
					g.If(a.Disabled, h.Disabled()),
					// the parser drops one newline after <textarea>: a text
					// starting with one keeps it only behind a second
					g.If(strings.HasPrefix(a.Value, "\n"), g.Text("\n")),
					g.Text(a.Value),
				),
			),
			h.Label(h.Class("cui-field-label"), h.For(a.ID), g.Text(a.Label)),
		),
		g.If(a.Error != "", errorMessage(message, a.Error)),
	}.Render(w)
}
