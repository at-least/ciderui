package console

import (
	"io"

	"github.com/at-least/ciderui/go/internal/need"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// ConsoleSwitch is an on/off switch and its words: a checkbox with the
// switch role, drawn as App Store Connect's toggle. Control writes attrs on
// the page's own <input>, which keeps its name, value and state.
type ConsoleSwitch struct {
	Label   string
	Control func(attrs g.Node) g.Node
}

func (s ConsoleSwitch) Render(w io.Writer) error {
	need.Field("console.ConsoleSwitch", "a Label", s.Label)
	if s.Control == nil {
		panic("console.ConsoleSwitch needs a Control")
	}
	return h.Label(h.Class("cui-console-switch"),
		s.Control(g.Group{h.Class("cui-console-switch-input"), h.Type("checkbox"), h.Role("switch")}),
		h.Span(h.Class("cui-console-switch-track"), h.Aria("hidden", "true")),
		g.Text(s.Label),
	).Render(w)
}
