package console

import (
	"io"

	"github.com/at-least/ciderui/go/internal/class"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// ConsoleStatus is a colored mark before a word. A line under it,
// <span class="cui-console-status-detail">, is the page's own.
type ConsoleStatus struct {
	Label string
	Color ConsoleColor
	Class string
}

func (s ConsoleStatus) Render(w io.Writer) error {
	if s.Label == "" {
		panic("console.ConsoleStatus needs a Label")
	}
	return h.Span(class.Root("cui-console-status", s.Class), optionalAttr("data-color", s.Color.value),
		h.Span(h.Class("cui-console-status-mark")),
		g.Text(s.Label),
	).Render(w)
}
