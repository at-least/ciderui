package console

import (
	"io"

	"github.com/at-least/ciderui/go/internal/class"
	"github.com/at-least/ciderui/go/lucide"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

type ConsoleEmptySize struct{ value string }

var (
	// ConsoleEmptySizeSmall is for a small place, a card.
	ConsoleEmptySizeSmall = ConsoleEmptySize{"sm"}
	// ConsoleEmptySizeLarge is for a whole page.
	ConsoleEmptySizeLarge = ConsoleEmptySize{"lg"}
)

// ConsoleEmpty is the empty state where a list or a chart has nothing to show.
type ConsoleEmpty struct {
	Size   ConsoleEmptySize
	Backed bool
	Icon   lucide.Icon
	Title  string
	// Body says what to do next.
	Body  string
	Class string
	// Children: its actions, in a <p class="cui-console-empty-actions">.
	Children g.Node
}

func (e ConsoleEmpty) Render(w io.Writer) error {
	if e.Size.value == "" {
		panic("console.ConsoleEmpty needs a Size")
	}
	if e.Title == "" {
		panic("console.ConsoleEmpty needs a Title")
	}
	return h.Div(class.Root("cui-console-empty", e.Class), h.Data("size", e.Size.value), g.If(e.Backed, h.Data("variant", "backed")),
		iconPart(e.Icon, "cui-console-empty-icon"),
		h.P(h.Class("cui-console-empty-title"), g.Text(e.Title)),
		g.If(e.Body != "", h.P(h.Class("cui-console-empty-body"), g.Text(e.Body))),
		e.Children,
	).Render(w)
}
