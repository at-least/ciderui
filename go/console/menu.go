package console

import (
	"io"

	"github.com/at-least/ciderui/go/icon"
	"github.com/at-least/ciderui/go/internal/class"
	"github.com/at-least/ciderui/go/lucide"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// ConsoleMenu is a section's pages down its left side.
type ConsoleMenu struct {
	AriaLabel string
	Class     string
	// Children: [ConsoleMenuItem]s and [ConsoleMenuGroup]s, with a
	// <p class="cui-console-menu-heading"> or an
	// <hr class="cui-console-menu-rule"> between runs of them.
	Children g.Node
}

func (m ConsoleMenu) Render(w io.Writer) error {
	if m.AriaLabel == "" {
		panic("console.ConsoleMenu needs an AriaLabel")
	}
	return h.Nav(class.Root("cui-console-menu", m.Class), h.Aria("label", m.AriaLabel), m.Children).Render(w)
}

// ConsoleMenuItem is a page in a [ConsoleMenu], a [ConsoleMenuGroup] or a
// [ConsoleSidebar].
type ConsoleMenuItem struct {
	Href     string
	Label    string
	Icon     icon.Icon
	Current  bool
	Disabled bool
	// Secondary is a count at the item's end.
	Secondary string
	Class     string
}

func (m ConsoleMenuItem) Render(w io.Writer) error {
	if m.Label == "" {
		panic("console.ConsoleMenuItem needs a Label")
	}
	if m.Disabled != (m.Href == "") {
		panic("console.ConsoleMenuItem needs an Href, unless it is disabled")
	}
	return h.A(class.Root("cui-console-menu-item", m.Class), g.If(!m.Disabled, h.Href(m.Href)), g.If(m.Current, h.Aria("current", "page")), g.If(m.Disabled, h.Aria("disabled", "true")),
		iconPart(m.Icon, "cui-console-menu-icon"),
		h.Span(h.Class("cui-console-menu-label"), g.Text(m.Label)),
		g.If(m.Secondary != "", h.Span(h.Class("cui-console-menu-secondary"), g.Text(m.Secondary))),
	).Render(w)
}

type ConsoleMenuGroup struct {
	Label string
	Icon  icon.Icon
	Open  bool
	Class string
	// Children: [ConsoleMenuItem]s.
	Children g.Node
}

func (m ConsoleMenuGroup) Render(w io.Writer) error {
	if m.Label == "" {
		panic("console.ConsoleMenuGroup needs a Label")
	}
	return h.Details(class.Root("cui-console-menu-group", m.Class), g.If(m.Open, h.Open()),
		h.Summary(h.Class("cui-console-menu-item"),
			iconPart(m.Icon, "cui-console-menu-icon"),
			h.Span(h.Class("cui-console-menu-label"), g.Text(m.Label)),
			lucide.ChevronRight.Class("cui-console-menu-chevron"),
		),
		m.Children,
	).Render(w)
}
