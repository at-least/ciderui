package console

import (
	"io"

	"github.com/at-least/ciderui/go/lucide"
	h "maragu.dev/gomponents/html"
)

// ConsoleCheckboxBox goes right after the page's own
// <input class="cui-console-checkbox-input" type="checkbox">, both in a
// <label class="cui-console-checkbox"> (or a <span>, the input then carrying
// an aria-label). The input stays the page's so it keeps its own attributes.
type ConsoleCheckboxBox struct{}

func (ConsoleCheckboxBox) Render(w io.Writer) error {
	return h.Span(h.Class("cui-console-checkbox-box"), h.Aria("hidden", "true"), lucide.Check, lucide.Minus).Render(w)
}
