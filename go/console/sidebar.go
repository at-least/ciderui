package console

import (
	"io"

	"github.com/at-least/ciderui/go/internal/need"
	"github.com/at-least/ciderui/go/lucide"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// @invariant: a page has one sidebar, so its id is the kit's, not the page's
const sidebarID = "sidebar"

// ConsoleSidebar is a popover: below 1024px it is a sheet that its open
// button, which [ConsoleContent] writes first in the page's <main>, opens.
type ConsoleSidebar struct {
	AriaLabel string
	// Home is the wordmark's href.
	Home     string
	Wordmark string
	// SkipLabel: the words of the skip link written right before the sidebar,
	// the first stop the keyboard reaches; it lands on [ConsoleContent]'s
	// <main id="main">.
	SkipLabel      string
	OpenAriaLabel  string
	CloseAriaLabel string
	// Footer: at the sidebar's foot (a sign-out form).
	Footer g.Node
	// Sections: each one's Title above its items; a section after the first
	// with no Title is set apart by a rule.
	Sections []ConsoleMenuSection
}

type ConsoleMenuSection struct {
	Title string
	Items []ConsoleMenuItem
}

func menuSections(sections []ConsoleMenuSection) g.Node {
	var nodes g.Group
	for i, s := range sections {
		section := need.At("console.ConsoleSidebar.Sections", i)
		need.Some(section, "Items", s.Items)
		switch {
		case s.Title != "":
			nodes = append(nodes, h.P(h.Class("cui-console-menu-title"), g.Text(s.Title)))
		case i > 0:
			nodes = append(nodes, h.Hr(h.Class("cui-console-menu-rule")))
		}
		for j, item := range s.Items {
			nodes = append(nodes, item.node(need.At(section+".Items", j)))
		}
	}
	return nodes
}

func (s ConsoleSidebar) Render(w io.Writer) error {
	need.Field("console.ConsoleSidebar", "an AriaLabel", s.AriaLabel)
	need.Field("console.ConsoleSidebar", "a Home", s.Home)
	need.Field("console.ConsoleSidebar", "a Wordmark", s.Wordmark)
	need.Field("console.ConsoleSidebar", "a SkipLabel", s.SkipLabel)
	need.Field("console.ConsoleSidebar", "an OpenAriaLabel", s.OpenAriaLabel)
	need.Field("console.ConsoleSidebar", "a CloseAriaLabel", s.CloseAriaLabel)
	need.Some("console.ConsoleSidebar", "Sections", s.Sections)
	return g.Group{
		ConsoleBtn{Href: "#main", Label: s.SkipLabel, Submit: false, Variant: ConsoleBtnVariantFilled, Size: ConsoleBtnSizeMedium, Destructive: false, Attrs: nil, Class: "cui-console-skip"},
		h.Nav(h.Class("cui-console-sidebar"), h.ID(sidebarID), h.Popover(), h.Aria("label", s.AriaLabel),
			h.Div(h.Class("cui-console-sidebar-header"),
				h.A(h.Class("cui-console-sidebar-wordmark"), h.Href(s.Home), g.Text(s.Wordmark)),
				iconButton(lucide.X, s.CloseAriaLabel, g.Group{h.PopoverTarget(sidebarID), h.PopoverTargetAction("hide"), h.AutoFocus()}, "cui-console-sidebar-close"),
			),
			h.Div(h.Class("cui-console-sidebar-body cui-console-menu"), h.Data("variant", "sidebar"), menuSections(s.Sections)),
			h.Div(h.Class("cui-console-sidebar-footer"), s.Footer),
		),
	}.Render(w)
}

func sidebarOpen(ariaLabel string) g.Node {
	return iconButton(lucide.Menu, ariaLabel, h.PopoverTarget(sidebarID), "cui-console-sidebar-open")
}
