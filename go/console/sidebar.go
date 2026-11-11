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
	// Children: [ConsoleMenuItem]s, with a <p class="cui-console-menu-title">
	// or an <hr class="cui-console-menu-rule"> between runs of them.
	Children g.Node
}

func (s ConsoleSidebar) Render(w io.Writer) error {
	if s.AriaLabel == "" {
		panic("console.ConsoleSidebar needs an AriaLabel")
	}
	if s.Wordmark == "" {
		panic("console.ConsoleSidebar needs a Wordmark")
	}
	if s.SkipLabel == "" {
		panic("console.ConsoleSidebar needs a SkipLabel")
	}
	if s.OpenAriaLabel == "" {
		panic("console.ConsoleSidebar needs an OpenAriaLabel")
	}
	if s.CloseAriaLabel == "" {
		panic("console.ConsoleSidebar needs a CloseAriaLabel")
	}
	return g.Group{
		ConsoleBtn{Href: "#main", Label: s.SkipLabel, AriaLabel: "", Icon: lucide.Icon{}, Submit: false, Variant: ConsoleBtnVariantFilled, Size: ConsoleBtnSizeMedium, Destructive: false, Attrs: nil, Class: "cui-console-skip"},
		h.Nav(h.Class("cui-console-sidebar"), h.ID(sidebarID), h.Popover(), h.Aria("label", s.AriaLabel),
			h.Div(h.Class("cui-console-sidebar-header"),
				h.A(h.Class("cui-console-sidebar-wordmark"), h.Href(s.Home), g.Text(s.Wordmark)),
				ConsoleBtn{Href: "", Label: "", AriaLabel: s.CloseAriaLabel, Icon: lucide.X, Submit: false, Variant: ConsoleBtnVariantPlain, Size: ConsoleBtnSizeMedium, Destructive: false, Attrs: g.Group{h.PopoverTarget(sidebarID), h.PopoverTargetAction("hide"), h.AutoFocus()}, Class: "cui-console-sidebar-close"},
			),
			h.Div(h.Class("cui-console-sidebar-body cui-console-menu"), h.Data("variant", "sidebar"), s.Children),
			h.Div(h.Class("cui-console-sidebar-footer"), s.Footer),
		),
	}.Render(w)
}

func sidebarOpen(ariaLabel string) g.Node {
	return ConsoleBtn{Href: "", Label: "", AriaLabel: ariaLabel, Icon: lucide.Menu, Submit: false, Variant: ConsoleBtnVariantPlain, Size: ConsoleBtnSizeMedium, Destructive: false, Attrs: h.PopoverTarget(sidebarID), Class: "cui-console-sidebar-open"}
}
