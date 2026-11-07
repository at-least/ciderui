package console

import (
	"io"

	"github.com/at-least/ciderui/go/internal/class"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// ConsoleStatus is a colored mark before a word, its detail line under it when
// it has one.
type ConsoleStatus struct {
	Label string
	Color ConsoleColor
	// Detail: the line under the word ("200 · 410 ms"); "" for none.
	Detail string
	Class  string
}

func (s ConsoleStatus) Render(w io.Writer) error {
	if s.Label == "" {
		panic("console.ConsoleStatus needs a Label")
	}
	return g.Group{
		h.Span(class.Root("cui-console-status", s.Class), optionalAttr("data-color", s.Color.value),
			h.Span(h.Class("cui-console-status-mark")),
			g.Text(s.Label),
		),
		g.If(s.Detail != "", h.Span(h.Class("cui-console-status-detail"), g.Text(s.Detail))),
	}.Render(w)
}
