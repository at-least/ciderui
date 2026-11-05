package console

import (
	"io"

	"github.com/at-least/ciderui/go/internal/class"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// ConsoleColor is a badge's, a status's or a banner's data-color, the toolkit's color roles.
type ConsoleColor struct{ value string }

var (
	ConsoleColorNeutral = ConsoleColor{value: ""}
	ConsoleColorInfo    = ConsoleColor{"info"}
	ConsoleColorSuccess = ConsoleColor{"success"}
	ConsoleColorWarning = ConsoleColor{"warning"}
	ConsoleColorError   = ConsoleColor{"error"}
)

// ConsoleBadge: a colored mark before a word is [ConsoleStatus], not a badge
// (decision 27).
type ConsoleBadge struct {
	Label string
	Color ConsoleColor
	Class string
}

func (b ConsoleBadge) Render(w io.Writer) error {
	if b.Label == "" {
		panic("console.ConsoleBadge needs a Label")
	}
	return h.Span(class.Root("cui-console-badge", b.Class),
		optionalAttr("data-color", b.Color.value),
		g.Text(b.Label),
	).Render(w)
}
