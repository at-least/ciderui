package console

import (
	"io"

	"github.com/at-least/ciderui/go/internal/need"
	"github.com/at-least/ciderui/go/lucide"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// ConsoleSidebarID is the [ConsoleSidebar]'s id, for a page's popovertarget
// and script.
//
// @invariant: a page has one sidebar, so its id is the kit's, not the page's
const ConsoleSidebarID = "cui-console-sidebar"

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
	// Current: the Href of the item for the page shown; "" for none.
	Current string
}

type ConsoleMenuSection struct {
	Title string
	Items []ConsoleMenuItem
}

// ConsoleMenuItem is a page in a [ConsoleSidebar].
type ConsoleMenuItem struct {
	Href  string
	Label string
	Icon  lucide.Icon
}

func (s ConsoleSidebar) Render(w io.Writer) error {
	need.Field("console.ConsoleSidebar", "an AriaLabel", s.AriaLabel)
	need.Field("console.ConsoleSidebar", "a Home", s.Home)
	need.Field("console.ConsoleSidebar", "a Wordmark", s.Wordmark)
	need.Field("console.ConsoleSidebar", "a SkipLabel", s.SkipLabel)
	need.Field("console.ConsoleSidebar", "an OpenAriaLabel", s.OpenAriaLabel)
	need.Field("console.ConsoleSidebar", "a CloseAriaLabel", s.CloseAriaLabel)
	need.Some("console.ConsoleSidebar", "Sections", s.Sections)
	var menu g.Group
	var hrefs []string
	for i, section := range s.Sections {
		path := need.At("console.ConsoleSidebar.Sections", i)
		need.Some(path, "Items", section.Items)
		switch {
		case section.Title != "":
			menu = append(menu, h.P(h.Class("cui-console-menu-title"), g.Text(section.Title)))
		case i > 0:
			menu = append(menu, h.Hr(h.Class("cui-console-menu-rule")))
		}
		for j, item := range section.Items {
			menu = append(menu, menuItem(need.At(path+".Items", j), item, item.Href == s.Current))
			hrefs = append(hrefs, item.Href)
		}
	}
	needCurrent("console.ConsoleSidebar", s.Current, hrefs)
	return g.Group{
		ConsoleBtn{Href: "#main", Label: s.SkipLabel, Submit: false, Variant: ConsoleBtnVariantFilled, Size: ConsoleBtnSizeMedium, Destructive: false, Attrs: nil, Class: "cui-console-skip"},
		h.Nav(h.Class("cui-console-sidebar"), h.ID(ConsoleSidebarID), h.Popover(), h.Aria("label", s.AriaLabel),
			h.Div(h.Class("cui-console-sidebar-header"),
				h.A(h.Class("cui-console-sidebar-wordmark"), h.Href(s.Home), g.Text(s.Wordmark)),
				iconButton(lucide.X, s.CloseAriaLabel, g.Group{h.PopoverTarget(ConsoleSidebarID), h.PopoverTargetAction("hide"), h.AutoFocus()}, "cui-console-sidebar-close"),
			),
			h.Div(h.Class("cui-console-sidebar-body cui-console-menu"), h.Data("variant", "sidebar"), menu),
			h.Div(h.Class("cui-console-sidebar-footer"), s.Footer),
		),
	}.Render(w)
}

func menuItem(path string, item ConsoleMenuItem, current bool) g.Node {
	need.Field(path, "an Href", item.Href)
	need.Field(path, "a Label", item.Label)
	return pageLink("cui-console-menu-item", item.Href, current, g.Group{
		g.If(item.Icon != (lucide.Icon{}), item.Icon.Class("cui-console-menu-icon")),
		h.Span(h.Class("cui-console-menu-label"), g.Text(item.Label)),
	})
}

func sidebarOpen(ariaLabel string) g.Node {
	return iconButton(lucide.Menu, ariaLabel, h.PopoverTarget(ConsoleSidebarID), "cui-console-sidebar-open")
}
