package site

import (
	"io"
	"strconv"

	"github.com/at-least/ciderui/go/internal/glyph"
	"github.com/at-least/ciderui/go/internal/need"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// GlobalNavMenuID is the id of the global nav's menu, for a page's script or
// test. Fixed: a page has one global nav.
const GlobalNavMenuID = "cui-globalnav-menu"

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
	OpenAriaLabel  string
	CloseAriaLabel string
}

func (n GlobalNav) Render(w io.Writer) error {
	need.Field("site.GlobalNav", "a Home", n.Home)
	need.Field("site.GlobalNav", "a Wordmark", n.Wordmark)
	need.Some("site.GlobalNav", "Links", n.Links)
	need.Field("site.GlobalNav", "an AriaLabel", n.AriaLabel)
	need.Field("site.GlobalNav", "a SkipLabel", n.SkipLabel)
	need.Field("site.GlobalNav", "an OpenAriaLabel", n.OpenAriaLabel)
	need.Field("site.GlobalNav", "a CloseAriaLabel", n.CloseAriaLabel)
	var links, postForms g.Group
	for i, link := range n.Links {
		path := need.At("site.GlobalNav.Links", i)
		need.Field(path, "a Label", link.Label)
		need.Field(path, "an Href", link.Href)
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
		Btn{Href: "#main", Label: n.SkipLabel, Icon: nil, Submit: false, Variant: BtnVariantPrimary, Size: BtnSizeMedium, FullWidth: false, Attrs: nil, Class: "cui-skip"},
		h.Nav(h.Class("cui-globalnav"), h.Aria("label", n.AriaLabel),
			h.Div(h.Class("cui-globalnav-content"),
				h.Ul(h.Class("cui-globalnav-list"), h.ID(GlobalNavMenuID), h.Popover(),
					h.Li(h.A(h.Class("cui-globalnav-wordmark"), h.Href(n.Home), g.Text(n.Wordmark))),
					h.Li(h.Button(h.Class("cui-globalnav-close"), h.PopoverTarget(GlobalNavMenuID), h.PopoverTargetAction("hide"), h.Aria("label", n.CloseAriaLabel), h.AutoFocus(), glyph.Menu)),
					links,
				),
				h.Button(h.Class("cui-globalnav-open"), h.PopoverTarget(GlobalNavMenuID), h.Aria("label", n.OpenAriaLabel), glyph.Menu),
				postForms,
			),
		),
	}.Render(w)
}
