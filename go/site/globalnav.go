package site

import (
	"io"
	"strconv"

	"github.com/at-least/ciderui/go/internal/glyph"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// @invariant: fixed id: a page has one global nav
const globalNavMenuID = "cui-globalnav-menu"

type GlobalNavLink struct {
	Label string
	Href  string
	// Post submits a POST to the Href (sign out: a change a GET link must not make).
	Post bool
}

// GlobalNav stays at the top of the page, over it (globalnav.css).
type GlobalNav struct {
	// Wordmark: the site's name, drawn in the kit's type; the home link's
	// text and so its accessible name.
	Wordmark  string
	Home      string
	Links     []GlobalNavLink
	AriaLabel string
	// SkipLabel: the words of the skip link written right before the nav, the
	// first stop the keyboard reaches; it lands on the page's own
	// <main id="main" tabindex="-1">.
	SkipLabel      string
	MenuAriaLabel  string
	CloseAriaLabel string
}

func (n GlobalNav) Render(w io.Writer) error {
	if n.Home == "" {
		panic("site.GlobalNav needs a Home")
	}
	if n.Wordmark == "" {
		panic("site.GlobalNav needs a Wordmark")
	}
	if len(n.Links) == 0 {
		panic("site.GlobalNav needs Links")
	}
	if n.AriaLabel == "" {
		panic("site.GlobalNav needs an AriaLabel")
	}
	if n.SkipLabel == "" {
		panic("site.GlobalNav needs a SkipLabel")
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
		formID := "cui-globalnav-post-" + strconv.Itoa(i+1)
		// @trap: not the <li> wrapping a <form>: the rows are li:has(> .cui-globalnav-link),
		// and a form between drew the item full size in a closed phone bar
		links = append(links, h.Li(h.Button(h.Class("cui-globalnav-link"), h.Type("submit"), h.FormAttr(formID), g.Text(link.Label))))
		postForms = append(postForms, h.Form(h.ID(formID), h.Action(link.Href), h.Method("post"), g.Attr("hidden")))
	}
	return g.Group{
		Btn{Href: "#main", Label: n.SkipLabel, Icon: nil, Submit: false, Variant: BtnVariantPrimary, Size: BtnSizeMedium, Block: false, Attrs: nil, Class: "cui-skip"},
		h.Nav(h.Class("cui-globalnav"), h.Aria("label", n.AriaLabel),
			h.Div(h.Class("cui-globalnav-content"),
				h.Ul(h.Class("cui-globalnav-list"), h.ID(globalNavMenuID), h.Popover(),
					h.Li(h.A(h.Class("cui-globalnav-wordmark"), h.Href(n.Home), g.Text(n.Wordmark))),
					h.Li(h.Button(h.Class("cui-globalnav-closebutton"), h.PopoverTarget(globalNavMenuID), h.PopoverTargetAction("hide"), h.Aria("label", n.CloseAriaLabel), h.AutoFocus(), glyph.Menu)),
					links,
				),
				h.Button(h.Class("cui-globalnav-menubutton"), h.PopoverTarget(globalNavMenuID), h.Aria("label", n.MenuAriaLabel), glyph.Menu),
				postForms,
			),
		),
	}.Render(w)
}
