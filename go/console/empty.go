package console

import (
	"io"

	"github.com/at-least/ciderui/go/icon"
	"github.com/at-least/ciderui/go/internal/class"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// ConsoleEmptySize is an [ConsoleEmpty]'s size (data-size). Its zero value is
// [ConsoleEmptySizeMedium].
type ConsoleEmptySize struct{ value string }

var (
	// ConsoleEmptySizeSmall is in a small place (a card's): data-size="sm".
	ConsoleEmptySizeSmall = ConsoleEmptySize{"sm"}
	// ConsoleEmptySizeMedium is the default: no data-size.
	ConsoleEmptySizeMedium = ConsoleEmptySize{value: ""}
	// ConsoleEmptySizeLarge is a whole page's: data-size="lg".
	ConsoleEmptySizeLarge = ConsoleEmptySize{"lg"}
)

// ConsoleEmpty is cui-console-empty (components/console/console-empty): App Store
// Connect's newer look's empty state, centered where a list or a chart has
// nothing to show: its icon, heading, caption, and what the page gives it,
// its actions (<p class="cui-console-empty-actions">).
type ConsoleEmpty struct {
	// Size is its size.
	Size ConsoleEmptySize
	// Backed is whether it is on its own faint ground
	// (data-variant="backed").
	Backed bool
	// Icon is its icon, a package lucide icon; the zero icon.Icon for none.
	Icon icon.Icon
	// Heading is its heading. Required: an empty one panics.
	Heading string
	// Caption is the line under it, what to do next; "" for none.
	Caption string
	// Class is the page's utilities on it; "" for none.
	Class string
	// Children is its actions; nil for none.
	Children g.Node
}

// Render writes the empty state.
func (e ConsoleEmpty) Render(w io.Writer) error {
	if e.Heading == "" {
		panic("console.ConsoleEmpty needs a Heading")
	}
	return h.Div(class.Root("cui-console-empty", e.Class), optionalAttr("data-size", e.Size.value), g.If(e.Backed, h.Data("variant", "backed")),
		iconPart(e.Icon, "cui-console-empty-icon"),
		h.P(h.Class("cui-console-empty-heading"), g.Text(e.Heading)),
		g.If(e.Caption != "", h.P(h.Class("cui-console-empty-caption"), g.Text(e.Caption))),
		e.Children,
	).Render(w)
}
