package console

import (
	"io"

	"github.com/at-least/ciderui/go/icon"
	"github.com/at-least/ciderui/go/internal/class"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

type ConsoleEmptySize struct{ value string }

var (
	// ConsoleEmptySizeSmall is for a small place, a card.
	ConsoleEmptySizeSmall  = ConsoleEmptySize{"sm"}
	ConsoleEmptySizeMedium = ConsoleEmptySize{value: ""}
	// ConsoleEmptySizeLarge is for a whole page.
	ConsoleEmptySizeLarge = ConsoleEmptySize{"lg"}
)

// ConsoleEmpty is the empty state where a list or a chart has nothing to show.
type ConsoleEmpty struct {
	Size    ConsoleEmptySize
	Backed  bool
	Icon    icon.Icon
	Heading string
	// Caption says what to do next.
	Caption string
	Class   string
	// Children: its actions, in a <p class="cui-console-empty-actions">.
	Children g.Node
}

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
