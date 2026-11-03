package docs

import (
	"io"

	"github.com/at-least/ciderui/go/icon"
	"github.com/at-least/ciderui/go/internal/class"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// DocsCopy's button stays hidden until the page's script unhides it; the
// script copies the listing's code and sets the button's data-state="copied"
// for a second.
type DocsCopy struct {
	Class string
	// Children: the page's <pre>, inside cui-docs-content.
	Children g.Node
}

func (c DocsCopy) Render(w io.Writer) error {
	return h.Div(class.Root("cui-docs-copy", c.Class),
		c.Children,
		h.Button(h.Class("cui-docs-copy-button"), h.Type("button"), h.Aria("label", "Copy"), g.Attr("hidden"),
			icon.DocsCopy.Class("cui-docs-copy-idle"),
			icon.DocsCopied.Class("cui-docs-copy-done"),
		),
	).Render(w)
}
