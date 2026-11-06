package console

import (
	"io"

	"github.com/at-least/ciderui/go/icon"
	"github.com/at-least/ciderui/go/internal/class"
	"github.com/at-least/ciderui/go/lucide"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// ConsolePager follows Apple Developer Forums' pager: a chevron alone for
// Previous and Next, a first-page button before them, and no button at all
// for a direction with no page.
type ConsolePager struct {
	// AriaLabel: "Issue pages", so a page with two lists names each.
	AriaLabel string
	// Summary is what the list shows ("51–100 of 230"); "" for a list paged
	// by a cursor past its first page, which knows no count.
	Summary string
	// First: "" where the previous page is the first. A list paged by a
	// cursor, which cannot step back one page, gives First and no Previous.
	First string
	// FirstLabel, PreviousAriaLabel, NextAriaLabel: the page's own words,
	// required even where that link is not written, so no list's second page
	// finds one missing.
	FirstLabel        string
	Previous          string
	PreviousAriaLabel string
	Next              string
	NextAriaLabel     string
	Class             string
}

func (p ConsolePager) Render(w io.Writer) error {
	if p.AriaLabel == "" {
		panic("console.ConsolePager needs an AriaLabel")
	}
	if p.FirstLabel == "" {
		panic("console.ConsolePager needs a FirstLabel")
	}
	if p.PreviousAriaLabel == "" {
		panic("console.ConsolePager needs a PreviousAriaLabel")
	}
	if p.NextAriaLabel == "" {
		panic("console.ConsolePager needs a NextAriaLabel")
	}
	if p.Summary == "" && p.First == "" && p.Previous == "" && p.Next == "" {
		panic("console.ConsolePager needs a Summary or a page to go to")
	}
	return h.Nav(class.Root("cui-console-pager", p.Class), h.Aria("label", p.AriaLabel),
		g.If(p.First != "", pagerBtn(p.First, p.FirstLabel, "", icon.Icon{})),
		g.If(p.Previous != "", pagerBtn(p.Previous, "", p.PreviousAriaLabel, lucide.ChevronLeft)),
		g.If(p.Summary != "", h.P(h.Class("cui-console-pager-summary"), g.Text(p.Summary))),
		g.If(p.Next != "", pagerBtn(p.Next, "", p.NextAriaLabel, lucide.ChevronRight)),
	).Render(w)
}

func pagerBtn(href, label, ariaLabel string, glyph icon.Icon) ConsoleBtn {
	return ConsoleBtn{Href: href, Label: label, AriaLabel: ariaLabel, Icon: glyph, Submit: false, Variant: ConsoleBtnVariantGray, Size: ConsoleBtnSizeMedium, Error: false, Capsule: false, Attrs: nil, Class: ""}
}
