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

type Field struct {
	ID           string
	Type         string
	Name         string
	Label        string
	Autocomplete string
	Value        string
	Error        string
	Disabled     bool
	// Prefix: a currency or a country code ($, +1).
	Prefix string
	Class  string
}

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
			// placeholder " ": the label reads :placeholder-shown to tell the field empty
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

func errorMessage(id, text string) g.Node {
	return h.P(h.Class("cui-field-message"), h.ID(id), icon.Alert, g.Text(text))
}

type Option struct {
	Value string
	Text  string
}

type Select struct {
	ID       string
	Name     string
	Label    string
	Options  []Option
	Selected string
	Error    string
	Disabled bool
	Class    string
}

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

type TextArea struct {
	ID       string
	Name     string
	Label    string
	Value    string
	Error    string
	Disabled bool
	Class    string
}

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
				// placeholder " ": the label reads :placeholder-shown to tell it empty
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
