package console

import (
	"io"

	"github.com/at-least/ciderui/go/lucide"
	g "maragu.dev/gomponents"
)

// ConsoleCopyIcons goes inside the page's own
// <button type="button" class="cui-console-copy-button"> (or
// cui-console-copy-inline); the page's script sets data-state="copied" on the
// button to show the check. The button stays the page's so it keeps the
// attributes its script reads.
type ConsoleCopyIcons struct{}

func (ConsoleCopyIcons) Render(w io.Writer) error {
	return g.Group{lucide.Copy.Class("cui-console-copy-idle"), lucide.Check.Class("cui-console-copy-done")}.Render(w)
}
