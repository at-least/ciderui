package console

import (
	"io"

	"github.com/at-least/ciderui/go/internal/attr"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// ConsoleStatus is a colored mark before a word, its detail line under it when
// it has one.
type ConsoleStatus struct {
	Label string
	Color ConsoleColor
	// Detail: the line under the word ("200 · 410 ms").
	Detail string
}

func (s ConsoleStatus) Render(w io.Writer) error {
	need("console.ConsoleStatus", "a Label", s.Label)
	return g.Group{
		h.Span(h.Class("cui-console-status"), attr.Optional("data-color", s.Color.value),
			h.Span(h.Class("cui-console-status-mark")),
			g.Text(s.Label),
		),
		optionalText(h.Span, "cui-console-status-detail", s.Detail),
	}.Render(w)
}
