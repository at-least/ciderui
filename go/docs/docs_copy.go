package docs

import (
	"io"

	"github.com/at-least/ciderui/go/icon"
	"github.com/at-least/ciderui/go/internal/class"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// DocsCopy is cui-docs-copy (components/docs/docs-copy): a code listing with
// Apple's developer documentation's copy button at its corner, hidden until
// the page's script shows it: the script copies the listing's code and sets
// the button's data-state="copied" for a second. The listing is the page's (a
// <pre> the docs content draws).
type DocsCopy struct {
	// Class is the page's utilities on the listing's box; "" for none.
	Class string
	// Children is the listing: the page's <pre>.
	Children g.Node
}

// Render writes the listing's box, the listing and the hidden button.
func (c DocsCopy) Render(w io.Writer) error {
	return h.Div(class.Root("cui-docs-copy", c.Class),
		c.Children,
		h.Button(h.Class("cui-docs-copy-button"), h.Type("button"), h.Aria("label", "Copy"), g.Attr("hidden"),
			icon.DocsCopy.Class("cui-docs-copy-idle"),
			icon.DocsCopied.Class("cui-docs-copy-done"),
		),
	).Render(w)
}
