package console

import (
	"io"

	"github.com/at-least/ciderui/go/internal/class"
	"github.com/at-least/ciderui/go/lucide"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// ConsolePager follows Apple Developer Forums' pager: a chevron alone for
// Previous and Next, a "First" button before them, and no button at all for
// a direction with no page.
type ConsolePager struct {
	// Label is the <nav>'s aria-label ("Issue pages"), so a page with two
	// lists names each.
	Label string
	// Summary is what the list shows ("51–100 of 230"); "" for a list paged
	// by a cursor past its first page, which knows no count.
	Summary string
	// First: "" where the previous page is the first. A list paged by a
	// cursor, which cannot step back one page, gives First and no Previous.
	First    string
	Previous string
	Next     string
	Class    string
}

func (p ConsolePager) Render(w io.Writer) error {
	if p.Label == "" {
		panic("console.ConsolePager needs a Label")
	}
	if p.Summary == "" && p.First == "" && p.Previous == "" && p.Next == "" {
		panic("console.ConsolePager needs a Summary or a page to go to")
	}
	return h.Nav(class.Root("cui-console-pager", p.Class), h.Aria("label", p.Label),
		g.If(p.First != "", h.A(h.Class("cui-console-btn"), h.Href(p.First), h.Aria("label", "First page"), g.Text("First"))),
		g.If(p.Previous != "", h.A(h.Class("cui-console-btn cui-console-btn-icon"), h.Href(p.Previous), h.Aria("label", "Previous page"), lucide.ChevronLeft)),
		g.If(p.Summary != "", h.P(h.Class("cui-console-pager-summary"), g.Text(p.Summary))),
		g.If(p.Next != "", h.A(h.Class("cui-console-btn cui-console-btn-icon"), h.Href(p.Next), h.Aria("label", "Next page"), lucide.ChevronRight)),
	).Render(w)
}
