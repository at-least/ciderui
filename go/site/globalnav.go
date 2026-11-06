package site

import (
	"io"
	"strconv"

	"github.com/at-least/ciderui/go/internal/class"
	"github.com/at-least/ciderui/go/internal/glyph"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// @invariant: fixed id: a page has one global nav
const globalNavMenuID = "cui-globalnav-menu"

type NavLink struct {
	Label string
	Href  string
	// Post submits a POST to the Href (sign out: a change a GET link must not make).
	Post bool
}

// GlobalNav does not position itself: put sticky top-0 in Class where the
// nav stays. Its z-index, apple.com's 9999, puts it over the page; a z-*
// utility in Class replaces it.
type GlobalNav struct {
	// Logo: an <svg> in currentColor takes the bar's color.
	Logo           g.Node
	Home           string
	LogoAriaLabel  string
	Links          []NavLink
	AriaLabel      string
	MenuAriaLabel  string
	CloseAriaLabel string
	Class          string
}

func (n GlobalNav) Render(w io.Writer) error {
	if n.Home == "" {
		panic("site.GlobalNav needs a Home")
	}
	if n.LogoAriaLabel == "" {
		panic("site.GlobalNav needs a LogoAriaLabel")
	}
	if len(n.Links) == 0 {
		panic("site.GlobalNav needs Links")
	}
	if n.AriaLabel == "" {
		panic("site.GlobalNav needs an AriaLabel")
	}
	if n.MenuAriaLabel == "" {
		panic("site.GlobalNav needs a MenuAriaLabel")
	}
	if n.CloseAriaLabel == "" {
		panic("site.GlobalNav needs a CloseAriaLabel")
	}
	var links, postForms g.Group
	for i, link := range n.Links {
		if link.Label == "" || link.Href == "" {
			panic("site.GlobalNav's Links need their Label and Href")
		}
		if !link.Post {
			links = append(links, h.Li(h.A(h.Class("cui-globalnav-link"), h.Href(link.Href), g.Text(link.Label))))
			continue
		}
		form := "cui-globalnav-post-" + strconv.Itoa(i+1)
		// @trap: not the <li> wrapping a <form>: the rows are li:has(> .cui-globalnav-link),
		// and a form between drew the item full size in a closed phone bar
		links = append(links, h.Li(h.Button(h.Class("cui-globalnav-link"), h.Type("submit"), h.FormAttr(form), g.Text(link.Label))))
		postForms = append(postForms, h.Form(h.ID(form), h.Action(link.Href), h.Method("post"), g.Attr("hidden")))
	}
	return h.Nav(class.Root("cui-globalnav", n.Class), h.Aria("label", n.AriaLabel),
		h.Div(h.Class("cui-globalnav-content"),
			h.Ul(h.Class("cui-globalnav-list"), h.ID(globalNavMenuID), h.Popover(),
				h.Li(h.A(h.Class("cui-globalnav-logo"), h.Href(n.Home), h.Aria("label", n.LogoAriaLabel), n.Logo)),
				h.Li(h.Button(h.Class("cui-globalnav-closebutton"), h.PopoverTarget(globalNavMenuID), h.PopoverTargetAction("hide"), h.Aria("label", n.CloseAriaLabel), h.AutoFocus(), glyph.Menu)),
				links,
			),
			h.Button(h.Class("cui-globalnav-menubutton"), h.PopoverTarget(globalNavMenuID), h.Aria("label", n.MenuAriaLabel), glyph.Menu),
			postForms,
		),
	).Render(w)
}
