package console

import (
	"io"

	"github.com/at-least/ciderui/go/internal/class"
	"github.com/at-least/ciderui/go/lucide"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// ConsolePager is cui-console-pager (components/console/console-pager): under a
// long list, the buttons to the pages either side of a line saying what it
// shows, as Apple Developer Forums' pager draws them: a chevron alone, named
// "Previous page" or "Next page", a "First" button before them, and no
// button at all for a direction with no page.
type ConsolePager struct {
	// Label is the <nav>'s aria-label, what the list's pages are called
	// ("Issue pages"), so a page with two lists names each. Required: an
	// empty one panics.
	Label string
	// Summary is what the list shows ("51–100 of 230"); "" for none (a list
	// paged by a cursor knows nothing to say past its first page).
	Summary string
	// First is the first page's URL, a "First" button; "" where there is
	// none or where the previous page is the first. A list paged by a
	// cursor, which cannot step back one page, gives this and no Previous.
	First string
	// Previous is the previous page's URL; "" where there is none.
	Previous string
	// Next is the next page's URL; "" where there is none.
	Next string
	// Class is the page's utilities on the <nav>; "" for none.
	Class string
}

// Render writes the pager.
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
