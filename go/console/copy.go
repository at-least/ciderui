package console

import (
	"io"

	"github.com/at-least/ciderui/go/lucide"
	g "maragu.dev/gomponents"
)

// ConsoleCopyIcons is cui-console-copy (components/console/console-copy): the two
// icons of a copy button, Lucide's copy and, while data-state="copied", its
// check, which go in the page's own
// <button type="button" class="cui-console-copy-button"> (or
// cui-console-copy-inline), so the button keeps whatever attributes the
// page's script reads.
type ConsoleCopyIcons struct{}

// Render writes the two icons.
func (ConsoleCopyIcons) Render(w io.Writer) error {
	return g.Group{lucide.Copy.Class("cui-console-copy-idle"), lucide.Check.Class("cui-console-copy-done")}.Render(w)
}
