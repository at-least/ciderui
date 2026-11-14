package console

import (
	"io"

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
		if len(s.Items) == 0 {
			panic("console.ConsoleMenuSection needs Items")
		}
		switch {
		case s.Title != "":
			nodes = append(nodes, h.P(h.Class("cui-console-menu-title"), g.Text(s.Title)))
		case i > 0:
			nodes = append(nodes, h.Hr(h.Class("cui-console-menu-rule")))
		}
		for _, item := range s.Items {
			nodes = append(nodes, item.node())
		}
	}
	return nodes
}

func (s ConsoleSidebar) Render(w io.Writer) error {
	need("console.ConsoleSidebar", "an AriaLabel", s.AriaLabel)
	need("console.ConsoleSidebar", "a Home", s.Home)
	need("console.ConsoleSidebar", "a Wordmark", s.Wordmark)
	need("console.ConsoleSidebar", "a SkipLabel", s.SkipLabel)
	need("console.ConsoleSidebar", "an OpenAriaLabel", s.OpenAriaLabel)
	need("console.ConsoleSidebar", "a CloseAriaLabel", s.CloseAriaLabel)
	if len(s.Sections) == 0 {
		panic("console.ConsoleSidebar needs Sections")
	}
	return g.Group{
		ConsoleBtn{Href: "#main", Label: s.SkipLabel, AriaLabel: "", Icon: lucide.Icon{}, Submit: false, Variant: ConsoleBtnVariantFilled, Size: ConsoleBtnSizeMedium, Destructive: false, Attrs: nil, Class: "cui-console-skip"},
		h.Nav(h.Class("cui-console-sidebar"), h.ID(sidebarID), h.Popover(), h.Aria("label", s.AriaLabel),
			h.Div(h.Class("cui-console-sidebar-header"),
				h.A(h.Class("cui-console-sidebar-wordmark"), h.Href(s.Home), g.Text(s.Wordmark)),
				ConsoleBtn{Href: "", Label: "", AriaLabel: s.CloseAriaLabel, Icon: lucide.X, Submit: false, Variant: ConsoleBtnVariantPlain, Size: ConsoleBtnSizeMedium, Destructive: false, Attrs: g.Group{h.PopoverTarget(sidebarID), h.PopoverTargetAction("hide"), h.AutoFocus()}, Class: "cui-console-sidebar-close"},
			),
			h.Div(h.Class("cui-console-sidebar-body cui-console-menu"), h.Data("variant", "sidebar"), menuSections(s.Sections)),
			h.Div(h.Class("cui-console-sidebar-footer"), s.Footer),
		),
	}.Render(w)
}

func sidebarOpen(ariaLabel string) g.Node {
	return ConsoleBtn{Href: "", Label: "", AriaLabel: ariaLabel, Icon: lucide.Menu, Submit: false, Variant: ConsoleBtnVariantPlain, Size: ConsoleBtnSizeMedium, Destructive: false, Attrs: h.PopoverTarget(sidebarID), Class: "cui-console-sidebar-open"}
}
