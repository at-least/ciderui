package console

import (
	"io"

	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// ConsoleStatus is an amber mark before a word, its detail line under it when
// it has one. One color, the warning's: a status marks what needs a look.
type ConsoleStatus struct {
	Label string
	// Detail: the line under the word ("Timed out · 10 s").
	Detail string
}

func (s ConsoleStatus) Render(w io.Writer) error {
	need("console.ConsoleStatus", "a Label", s.Label)
	return g.Group{
		h.Span(h.Class("cui-console-status"),
			h.Span(h.Class("cui-console-status-mark")),
			g.Text(s.Label),
		),
		optionalText(h.Span, "cui-console-status-detail", s.Detail),
	}.Render(w)
}
