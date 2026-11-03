package console

import (
	"io"
	"strconv"

	"github.com/at-least/ciderui/go/internal/class"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// ConsoleListing is a listing with numbered lines. A listing without numbers
// is the page's own <pre class="cui-console-listing">.
type ConsoleListing struct {
	FirstLine uint
	// Lines: each without its newline.
	Lines []string
	// Highlighted: 0 for none.
	Highlighted uint
	Class       string
}

func (l ConsoleListing) Render(w io.Writer) error {
	if len(l.Lines) == 0 {
		panic("console.ConsoleListing needs a line")
	}
	if l.FirstLine == 0 {
		panic("console.ConsoleListing's FirstLine is at least 1")
	}
	if l.Highlighted != 0 && (l.Highlighted < l.FirstLine || l.Highlighted >= l.FirstLine+uint(len(l.Lines))) {
		panic("console.ConsoleListing's Highlighted is one of its lines")
	}
	var lines g.Group
	for i, line := range l.Lines {
		number := l.FirstLine + uint(i)
		if i > 0 {
			lines = append(lines, g.Text("\n"))
		}
		lines = append(lines, h.Span(h.Class("cui-console-listing-line"), h.Data("line", strconv.FormatUint(uint64(number), 10)), g.If(l.Highlighted == number, g.Attr("data-highlighted")), g.Text(line)))
	}
	return h.Pre(class.Root("cui-console-listing", l.Class), h.Code(lines)).Render(w)
}
