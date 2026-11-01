package console

import (
	"io"

	"github.com/at-least/ciderui/go/lucide"
	h "maragu.dev/gomponents/html"
)

// ConsoleCheckboxBox is cui-console-checkbox-box (components/console/console-checkbox):
// the box of App Store Connect's newer look's checkbox, its check and its
// dash, which goes after the page's own
// <input class="cui-console-checkbox-input" type="checkbox"> in a
// <label class="cui-console-checkbox"> (or a <span>, the input then carrying
// an aria-label), so the input keeps whatever attributes the page gives it.
type ConsoleCheckboxBox struct{}

// Render writes the box.
func (ConsoleCheckboxBox) Render(w io.Writer) error {
	return h.Span(h.Class("cui-console-checkbox-box"), h.Aria("hidden", "true"), lucide.Check, lucide.Minus).Render(w)
}
