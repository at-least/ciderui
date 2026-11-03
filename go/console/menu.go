package console

import (
	"io"

	"github.com/at-least/ciderui/go/icon"
	"github.com/at-least/ciderui/go/internal/class"
	"github.com/at-least/ciderui/go/lucide"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// ConsoleMenu is cui-console-menu (components/console/console-menu): App Store
// Connect's newer look's list of a section's pages down its left side; its
// [ConsoleMenuItem]s, [ConsoleMenuGroup]s, and a heading
// (<p class="cui-console-menu-heading">) or a rule
// (<hr class="cui-console-menu-rule">) between runs of them.
type ConsoleMenu struct {
	// Label is what the list is, for a screen reader (aria-label).
	// Required: an empty one panics.
	Label string
	// Class is the page's utilities on the list; "" for none.
	Class string
	// Children is its items, groups, headings and rules.
	Children g.Node
}

// Render writes the menu.
func (m ConsoleMenu) Render(w io.Writer) error {
	if m.Label == "" {
		panic("console.ConsoleMenu needs a Label")
	}
	return h.Nav(class.Root("cui-console-menu", m.Class), h.Aria("label", m.Label), m.Children).Render(w)
}

// ConsoleMenuItem is a page in a [ConsoleMenu], a [ConsoleMenuGroup] or a
// [ConsoleSidebar], a link.
type ConsoleMenuItem struct {
	// Href is where it goes. Required unless it is disabled, when it is
	// left out ("").
	Href string
	// Label is its label. Required: an empty one panics.
	Label string
	// Icon is its icon before the label, a package lucide icon; the zero
	// icon.Icon for none.
	Icon icon.Icon
	// Current is whether it is the current page (aria-current="page").
	Current bool
	// Disabled is whether it goes nowhere now (aria-disabled, no href).
	Disabled bool
	// Secondary is what follows it at its end (a count); "" for nothing.
	Secondary string
	// Class is the page's utilities on the item; "" for none.
	Class string
}

// Render writes the item.
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

// ConsoleMenuGroup is a group of a [ConsoleMenu] (<details>): its label opening its
// items, set in from the left, with no script.
type ConsoleMenuGroup struct {
	// Label is its label. Required: an empty one panics.
	Label string
	// Icon is its icon before the label; the zero icon.Icon for none.
	Icon icon.Icon
	// Open is whether it is open.
	Open bool
	// Class is the page's utilities on the group; "" for none.
	Class string
	// Children is its items, [ConsoleMenuItem]s.
	Children g.Node
}

// Render writes the group.
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
