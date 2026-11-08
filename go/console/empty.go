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
	// ConsoleEmptySizeSmall is for a card: small, on its own gray inset.
	ConsoleEmptySizeSmall = ConsoleEmptySize{"sm"}
	// ConsoleEmptySizeLarge is for a whole page: large, on the page's ground.
	ConsoleEmptySizeLarge = ConsoleEmptySize{"lg"}
)

// ConsoleEmpty is the empty state where a list or a chart has nothing to show.
type ConsoleEmpty struct {
	Size  ConsoleEmptySize
	Icon  lucide.Icon
	Title string
	// Body says what to do next.
	Body  string
	Class string
}

func (e ConsoleEmpty) Render(w io.Writer) error {
	if e.Size.value == "" {
		panic("console.ConsoleEmpty needs a Size: ConsoleEmptySizeSmall (a card) or ConsoleEmptySizeLarge (a whole page)")
	}
	if e.Title == "" {
		panic("console.ConsoleEmpty needs a Title")
	}
	return h.Div(class.Root("cui-console-empty", e.Class), h.Data("size", e.Size.value),
		iconPart(e.Icon, "cui-console-empty-icon"),
		h.P(h.Class("cui-console-empty-title"), g.Text(e.Title)),
		g.If(e.Body != "", h.P(h.Class("cui-console-empty-body"), g.Text(e.Body))),
	).Render(w)
}
