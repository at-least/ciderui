package console

import (
	"io"

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

// ConsoleEmpty is the empty state where a list or a table has nothing to
// show; an empty chart is its [ConsoleChart.Empty].
type ConsoleEmpty struct {
	// Size: a large one, a whole page's, draws the kit's inbox above its
	// title; a small one, a card's, no glyph.
	Size  ConsoleEmptySize
	Title string
	// Body says what to do next.
	Body string
	// Attrs: the root's other attributes, for the page's script (data-*, an
	// htmx attribute).
	Attrs g.Node
}

func (e ConsoleEmpty) Render(w io.Writer) error {
	if e.Size.value == "" {
		panic("console.ConsoleEmpty needs a Size: ConsoleEmptySizeSmall (a card) or ConsoleEmptySizeLarge (a whole page)")
	}
	need("console.ConsoleEmpty", "a Title", e.Title)
	return h.Div(h.Class("cui-console-empty"), h.Data("size", e.Size.value), e.Attrs,
		g.If(e.Size == ConsoleEmptySizeLarge, iconPart(lucide.Inbox, "cui-console-empty-icon")),
		h.P(h.Class("cui-console-empty-title"), g.Text(e.Title)),
		optionalText(h.P, "cui-console-empty-body", e.Body),
	).Render(w)
}
