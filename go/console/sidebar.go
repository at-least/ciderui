package console

import (
	"io"

	"github.com/at-least/ciderui/go/lucide"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// @invariant: a page has one sidebar, so its id is the kit's, not the page's
const sidebarID = "sidebar"

// ConsoleSidebar is a popover: below 1024px it is a sheet that a
// [ConsoleSidebarOpen] opens.
type ConsoleSidebar struct {
	AriaLabel string
	// Home is the wordmark's href.
	Home           string
	Wordmark       string
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
	if s.CloseAriaLabel == "" {
		panic("console.ConsoleSidebar needs a CloseAriaLabel")
	}
	return h.Nav(h.Class("cui-console-sidebar"), h.ID(sidebarID), h.Popover(), h.Aria("label", s.AriaLabel),
		h.Div(h.Class("cui-console-sidebar-header"),
			h.A(h.Class("cui-console-sidebar-wordmark"), h.Href(s.Home), g.Text(s.Wordmark)),
			ConsoleBtn{Href: "", Label: "", AriaLabel: s.CloseAriaLabel, Icon: lucide.X, Submit: false, Variant: ConsoleBtnVariantPlain, Size: ConsoleBtnSizeMedium, Destructive: false, Attrs: g.Group{h.PopoverTarget(sidebarID), h.PopoverTargetAction("hide"), h.AutoFocus()}, Class: "cui-console-sidebar-close"},
		),
		h.Div(h.Class("cui-console-sidebar-body cui-console-menu"), h.Data("variant", "sidebar"), s.Children),
		h.Div(h.Class("cui-console-sidebar-footer"), s.Footer),
	).Render(w)
}

// ConsoleSidebarOpen is the button that opens a [ConsoleSidebar] below 1024px,
// placed in the page's own column.
type ConsoleSidebarOpen struct {
	AriaLabel string
	Class     string
}

func (o ConsoleSidebarOpen) Render(w io.Writer) error {
	if o.AriaLabel == "" {
		panic("console.ConsoleSidebarOpen needs an AriaLabel")
	}
	root := "cui-console-sidebar-open"
	if o.Class != "" {
		root += " " + o.Class
	}
	return ConsoleBtn{Href: "", Label: "", AriaLabel: o.AriaLabel, Icon: lucide.Menu, Submit: false, Variant: ConsoleBtnVariantPlain, Size: ConsoleBtnSizeMedium, Destructive: false, Attrs: h.PopoverTarget(sidebarID), Class: root}.Render(w)
}
