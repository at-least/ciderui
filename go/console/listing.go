package console

import (
	"io"
	"strconv"

	"github.com/at-least/ciderui/go/internal/class"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// ConsoleListing is cui-console-listing (components/console/console-listing) with
// its lines numbered (cui-console-listing-line): a frame's source, one line
// highlighted, as Apple's developer documentation's tutorials mark the lines
// that changed. A listing without numbers is one class on the page's own
// <pre>.
type ConsoleListing struct {
	// FirstLine is the first line's number. At least 1: 0 is Highlighted's
	// none, so a 0 panics.
	FirstLine uint
	// Lines are the lines, each without its newline. Required: none panics.
	Lines []string
	// Highlighted is the highlighted line's number; 0 for none.
	Highlighted uint
	// Class is the page's utilities on the <pre>; "" for none.
	Class string
}

// Render writes the listing.
func (l ConsoleListing) Render(w io.Writer) error {
	if len(l.Lines) == 0 {
		panic("console.ConsoleListing needs a line")
	}
	if l.FirstLine == 0 {
		panic("console.ConsoleListing's FirstLine is at least 1")
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
