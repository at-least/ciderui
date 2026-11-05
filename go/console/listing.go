package console

import (
	"io"
	"strconv"

	"github.com/at-least/ciderui/go/internal/class"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// ConsoleListing is a listing with numbered lines. A listing without numbers
// is the page's own <pre class="cui-console-listing" tabindex="0">.
type ConsoleListing struct {
	FirstLine uint
	Lines     []ConsoleListingLine
	Class     string
}

type ConsoleListingLine struct {
	// Label: without its newline.
	Label       string
	Highlighted bool
}

func (l ConsoleListing) Render(w io.Writer) error {
	if len(l.Lines) == 0 {
		panic("console.ConsoleListing needs a line")
	}
	var lines g.Group
	for i, line := range l.Lines {
		number := l.FirstLine + uint(i)
		if i > 0 {
			lines = append(lines, g.Text("\n"))
		}
		lines = append(lines, h.Span(h.Class("cui-console-listing-line"), h.Data("line", strconv.FormatUint(uint64(number), 10)), g.If(line.Highlighted, g.Attr("data-highlighted")), g.Text(line.Label)))
	}
	// tabindex: the keyboard scrolls a long line, as axe's scrollable-region-focusable asks
	return h.Pre(class.Root("cui-console-listing", l.Class), h.TabIndex("0"), h.Code(lines)).Render(w)
}
