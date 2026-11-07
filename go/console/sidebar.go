package console

import (
	"io"

	"github.com/at-least/ciderui/go/internal/class"
	"github.com/at-least/ciderui/go/lucide"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// ConsoleSidebar is a popover: below 1024px it is a sheet that the page's own
// <button class="cui-console-sidebar-open" popovertarget="ID"> opens.
type ConsoleSidebar struct {
	ID        string
	AriaLabel string
	// Home is the wordmark's href.
	Home           string
	Wordmark       string
	CloseAriaLabel string
	Class          string
	// Footer: an account.
	Footer g.Node
	// Children: [ConsoleMenuItem]s, with a <p class="cui-console-menu-heading">
	// or an <hr class="cui-console-menu-rule"> between runs of them.
	Children g.Node
}

func (s ConsoleSidebar) Render(w io.Writer) error {
	if s.ID == "" {
		panic("console.ConsoleSidebar needs an ID")
	}
	if s.AriaLabel == "" {
		panic("console.ConsoleSidebar needs an AriaLabel")
	}
	if s.Wordmark == "" {
		panic("console.ConsoleSidebar needs a Wordmark")
	}
	if s.CloseAriaLabel == "" {
		panic("console.ConsoleSidebar needs a CloseAriaLabel")
	}
	return h.Nav(class.Root("cui-console-sidebar", s.Class), h.ID(s.ID), h.Popover(), h.Aria("label", s.AriaLabel),
		h.Div(h.Class("cui-console-sidebar-header"),
			h.A(h.Class("cui-console-sidebar-wordmark"), h.Href(s.Home), g.Text(s.Wordmark)),
			ConsoleBtn{Href: "", Label: "", AriaLabel: s.CloseAriaLabel, Icon: lucide.X, Submit: false, Variant: ConsoleBtnVariantPlain, Size: ConsoleBtnSizeMedium, Error: false, Attrs: g.Group{h.PopoverTarget(s.ID), h.PopoverTargetAction("hide"), h.AutoFocus()}, Class: "cui-console-sidebar-close"},
		),
		h.Div(h.Class("cui-console-sidebar-body cui-console-menu"), h.Data("variant", "sidebar"), s.Children),
		h.Div(h.Class("cui-console-sidebar-footer"), s.Footer),
	).Render(w)
}

// ConsoleSidebarOpen is the button that opens a [ConsoleSidebar] below 1024px,
// placed in the page's own column.
type ConsoleSidebarOpen struct {
	// Target: the [ConsoleSidebar]'s ID.
	Target    string
	AriaLabel string
	Class     string
}

func (o ConsoleSidebarOpen) Render(w io.Writer) error {
	if o.Target == "" {
		panic("console.ConsoleSidebarOpen needs a Target")
	}
	if o.AriaLabel == "" {
		panic("console.ConsoleSidebarOpen needs an AriaLabel")
	}
	root := "cui-console-sidebar-open"
	if o.Class != "" {
		root += " " + o.Class
	}
	return ConsoleBtn{Href: "", Label: "", AriaLabel: o.AriaLabel, Icon: lucide.Menu, Submit: false, Variant: ConsoleBtnVariantPlain, Size: ConsoleBtnSizeMedium, Error: false, Attrs: h.PopoverTarget(o.Target), Class: root}.Render(w)
}
