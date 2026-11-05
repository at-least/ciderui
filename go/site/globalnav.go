package site

import (
	"io"
	"strconv"

	"github.com/at-least/ciderui/go/icon"
	"github.com/at-least/ciderui/go/internal/class"
	"github.com/at-least/ciderui/go/internal/glyph"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// fixed ids: a page has one global nav
const (
	globalNavMenuID       = "cui-globalnav-menu"
	globalNavSearchID     = "cui-globalnav-search"
	globalNavHeadingID    = "cui-globalnav-searchheading"
	globalNavBagID        = "cui-globalnav-bag"
	globalNavBagHeadingID = "cui-globalnav-bagheading"
)

type NavLink struct {
	Label string
	Href  string
	// Action: nil for a plain link.
	Action NavAction
}

// NavAction is what a [NavLink] does beside going to its Href.
type NavAction interface{ navAction() }

// NavFlyout opens groups of links under the bar.
type NavFlyout struct {
	Groups []FlyoutGroup
	// AriaLabel: the button that opens it ("Store menu").
	AriaLabel string
}

// NavPost submits a POST to the link's Href (sign out: a change a GET link must not make).
type NavPost struct{}

func (NavFlyout) navAction() {}
func (NavPost) navAction()   {}

type FlyoutGroup struct {
	Heading string
	// Large: apple.com's elevated links.
	Large []Link
	Links []Link
}

type Search interface{ search() }

type LinkedSearch struct {
	Href      string
	AriaLabel string
}

type FlyoutSearch struct {
	Action string
	// AriaLabel: its button, and its field.
	AriaLabel      string
	Placeholder    string
	ClearAriaLabel string
	Heading        string
	Links          []Link
}

func (LinkedSearch) search() {}
func (FlyoutSearch) search() {}

type Bag interface{ bag() }

type LinkedBag struct {
	Href      string
	AriaLabel string
}

// FlyoutBag is apple.com's empty-bag flyout.
type FlyoutBag struct {
	AriaLabel string
	Message   string
	Line      *BagLine
	Heading   string
	Links     []BagLink
}

func (LinkedBag) bag() {}
func (FlyoutBag) bag() {}

type BagLine struct {
	Label string
	Href  string
	// Rest follows the link with no space: start it with one
	// (" for items you saved earlier").
	Rest string
}

// BagLinkIcon is the glyph before a bag's profile link; its zero value panics
// when rendered.
type BagLinkIcon struct{ regular, compact icon.Icon }

var (
	BagLinkIconOrders  = BagLinkIcon{regular: glyph.Orders, compact: glyph.OrdersCompact}
	BagLinkIconSaves   = BagLinkIcon{regular: glyph.Saves, compact: glyph.SavesCompact}
	BagLinkIconAccount = BagLinkIcon{regular: glyph.Account, compact: glyph.AccountCompact}
	BagLinkIconSignIn  = BagLinkIcon{regular: glyph.SignIn, compact: glyph.SignInCompact}
)

type BagLink struct {
	Icon  BagLinkIcon
	Label string
	Href  string
}

// GlobalNavText is the words of the nav's own parts, in the page's language.
type GlobalNavText struct {
	AriaLabel      string
	MenuAriaLabel  string
	CloseAriaLabel string
	// MainMenuAriaLabel: a phone's flyout's back button; "" with no [NavFlyout].
	MainMenuAriaLabel string
}

// GlobalNav does not position itself: put sticky top-0 in Class where the
// nav stays. Its z-index, apple.com's 9999, puts it over the page and a
// [LocalNav]; a z-* utility in Class replaces it.
type GlobalNav struct {
	// Logo: an <svg> in currentColor takes the bar's color.
	Logo          g.Node
	Home          string
	LogoAriaLabel string
	Links         []NavLink
	Search        Search
	Bag           Bag
	Text          GlobalNavText
	Clear         bool
	Class         string
}

func globalNavImages(regular, compact icon.Icon) g.Node {
	return g.Group{regular.Class("cui-globalnav-image"), compact.Class("cui-globalnav-image-compact")}
}

// the curtain takes the click outside a popover and closes it
func globalNavCurtain(id string) g.Node {
	return h.Button(h.Class("cui-globalnav-curtain"), h.PopoverTarget(id), h.PopoverTargetAction("hide"), h.TabIndex("-1"), h.Aria("hidden", "true"))
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
	for _, text := range []struct{ field, value string }{
		{"AriaLabel", n.Text.AriaLabel},
		{"MenuAriaLabel", n.Text.MenuAriaLabel},
		{"CloseAriaLabel", n.Text.CloseAriaLabel},
	} {
		if text.value == "" {
			panic("site.GlobalNav's Text needs its " + text.field)
		}
	}
	for _, link := range n.Links {
		if link.Label == "" || link.Href == "" {
			panic("site.GlobalNav's Links need their Label and Href")
		}
		flyout, ok := link.Action.(NavFlyout)
		if !ok {
			continue
		}
		if n.Text.MainMenuAriaLabel == "" {
			panic("site.GlobalNav's Text needs its MainMenuAriaLabel")
		}
		if flyout.AriaLabel == "" {
			panic("site.NavFlyout needs an AriaLabel")
		}
		if len(flyout.Groups) == 0 {
			panic("site.NavFlyout needs Groups")
		}
		for i, group := range flyout.Groups {
			if group.Heading == "" {
				panic("site.NavFlyout's Groups need their Heading")
			}
			if i > 0 && len(group.Large) > 0 {
				panic("site.NavFlyout's Large links are only in its first group")
			}
		}
	}
	switch search := n.Search.(type) {
	case LinkedSearch:
		if search.AriaLabel == "" {
			panic("site.LinkedSearch needs an AriaLabel")
		}
	case FlyoutSearch:
		if search.Action == "" {
			panic("site.FlyoutSearch needs an Action")
		}
		if search.AriaLabel == "" {
			panic("site.FlyoutSearch needs an AriaLabel")
		}
		if search.Placeholder == "" {
			panic("site.FlyoutSearch needs a Placeholder")
		}
		if search.ClearAriaLabel == "" {
			panic("site.FlyoutSearch needs a ClearAriaLabel")
		}
		if len(search.Links) > 0 && search.Heading == "" {
			panic("site.FlyoutSearch's Links need their Heading")
		}
	}
	switch bag := n.Bag.(type) {
	case LinkedBag:
		if bag.AriaLabel == "" {
			panic("site.LinkedBag needs an AriaLabel")
		}
	case FlyoutBag:
		if bag.AriaLabel == "" {
			panic("site.FlyoutBag needs an AriaLabel")
		}
		if bag.Message == "" {
			panic("site.FlyoutBag needs a Message")
		}
		if bag.Line != nil && (bag.Line.Label == "" || bag.Line.Href == "") {
			panic("site.FlyoutBag's Line needs its link's Label and Href")
		}
		if len(bag.Links) > 0 && bag.Heading == "" {
			panic("site.FlyoutBag's Links need their Heading")
		}
		for _, link := range bag.Links {
			if link.Label == "" || link.Href == "" {
				panic("site.FlyoutBag's Links need their Label and Href")
			}
			if link.Icon.regular == (icon.Icon{}) {
				panic("site.FlyoutBag's Links need their Icon")
			}
		}
	}
	var links, postForms g.Group
	for i, link := range n.Links {
		switch action := link.Action.(type) {
		case nil:
			links = append(links, h.Li(h.A(h.Class("cui-globalnav-link"), h.Href(link.Href), g.Text(link.Label))))
		case NavPost:
			form := "cui-globalnav-post-" + strconv.Itoa(i+1)
			// not the <li> wrapping a <form>: the rows are li:has(> .cui-globalnav-link),
			// and a form between drew the item full size in a closed phone bar
			links = append(links, h.Li(h.Button(h.Class("cui-globalnav-link"), h.Type("submit"), h.FormAttr(form), g.Text(link.Label))))
			postForms = append(postForms, h.Form(h.ID(form), h.Action(link.Href), h.Method("post"), g.Attr("hidden")))
		case NavFlyout:
			links = append(links, h.Li(globalNavFlyout(link, action, i+1, n.Text)))
		default:
			panic("site.NavLink's Action is neither a NavFlyout nor a NavPost")
		}
	}
	var searchItem, searchPopover g.Node
	switch search := n.Search.(type) {
	case nil:
	case LinkedSearch:
		searchItem = h.Li(h.A(h.Class("cui-globalnav-icon"), h.Href(search.Href), h.Aria("label", search.AriaLabel), globalNavImages(glyph.Search, glyph.SearchCompact)))
	case FlyoutSearch:
		searchItem = h.Li(h.Button(h.Class("cui-globalnav-icon"), h.PopoverTarget(globalNavSearchID), h.Aria("label", search.AriaLabel), globalNavImages(glyph.Search, glyph.SearchCompact)))
		searchPopover = h.Div(h.Class("cui-globalnav-search"), h.ID(globalNavSearchID), h.Popover(),
			h.Div(h.Class("cui-globalnav-searchpanel"),
				h.Form(h.Class("cui-globalnav-searchfield"), h.Action(search.Action), h.Method("get"), h.Role("search"),
					globalNavImages(glyph.Searchfield, glyph.SearchfieldCompact),
					h.Input(h.Class("cui-globalnav-searchinput"), h.Type("text"), h.Name("q"), h.Placeholder(search.Placeholder), h.Aria("label", search.AriaLabel), h.AutoComplete("off"), h.AutoCapitalize("off"), h.EnterKeyHint("search"), h.AutoFocus()),
					h.Button(h.Class("cui-globalnav-searchreset"), h.Type("reset"), h.Aria("label", search.ClearAriaLabel), globalNavImages(glyph.Clear, glyph.ClearCompact)),
				),
				g.If(len(search.Links) > 0, g.Group{
					h.H2(h.Class("cui-globalnav-searchheading"), h.ID(globalNavHeadingID), g.Text(search.Heading)),
					h.Ul(h.Class("cui-globalnav-searchlinks"), h.Aria("labelledby", globalNavHeadingID), g.Map(search.Links, func(link Link) g.Node {
						return h.Li(h.A(h.Class("cui-globalnav-searchlink"), h.Href(link.Href), globalNavImages(glyph.Arrow, glyph.ArrowCompact), g.Text(link.Label)))
					})),
				}),
				h.Button(h.Class("cui-globalnav-closebutton"), h.PopoverTarget(globalNavSearchID), h.PopoverTargetAction("hide"), h.Aria("label", n.Text.CloseAriaLabel), glyph.Menu),
			),
			globalNavCurtain(globalNavSearchID),
		)
	default:
		panic("site.GlobalNav's Search is neither a LinkedSearch nor a FlyoutSearch")
	}
	var bagItem, bagPopover g.Node
	switch bag := n.Bag.(type) {
	case nil:
	case LinkedBag:
		bagItem = h.Li(h.A(h.Class("cui-globalnav-icon"), h.Href(bag.Href), h.Aria("label", bag.AriaLabel), globalNavImages(glyph.Bag, glyph.BagCompact)))
	case FlyoutBag:
		bagItem = h.Li(h.Button(h.Class("cui-globalnav-icon"), h.PopoverTarget(globalNavBagID), h.Aria("label", bag.AriaLabel), globalNavImages(glyph.Bag, glyph.BagCompact)))
		bagPopover = h.Div(h.Class("cui-globalnav-bag"), h.ID(globalNavBagID), h.Popover(),
			h.Div(h.Class("cui-globalnav-bagpanel"),
				h.Button(h.Class("cui-globalnav-closebutton"), h.PopoverTarget(globalNavBagID), h.PopoverTargetAction("hide"), h.Aria("label", n.Text.CloseAriaLabel), h.AutoFocus(), glyph.Menu),
				h.H2(h.Class("cui-globalnav-bagmessage"), g.Text(bag.Message)),
				g.Iff(bag.Line != nil, func() g.Node {
					return h.P(h.Class("cui-globalnav-bagtext"), h.A(h.Href(bag.Line.Href), g.Text(bag.Line.Label)), g.Text(bag.Line.Rest))
				}),
				g.If(len(bag.Links) > 0, g.Group{
					h.H3(h.Class("cui-globalnav-bagheading"), h.ID(globalNavBagHeadingID), g.Text(bag.Heading)),
					h.Ul(h.Class("cui-globalnav-baglinks"), h.Aria("labelledby", globalNavBagHeadingID), g.Map(bag.Links, func(link BagLink) g.Node {
						return h.Li(h.A(h.Class("cui-globalnav-baglink"), h.Href(link.Href), globalNavImages(link.Icon.regular, link.Icon.compact), g.Text(link.Label)))
					})),
				}),
			),
			globalNavCurtain(globalNavBagID),
		)
	default:
		panic("site.GlobalNav's Bag is neither a LinkedBag nor a FlyoutBag")
	}
	return h.Nav(class.Root("cui-globalnav", n.Class), g.If(n.Clear, h.Data("variant", "clear")), h.Aria("label", n.Text.AriaLabel),
		h.Div(h.Class("cui-globalnav-content"),
			h.Ul(h.Class("cui-globalnav-list"), h.ID(globalNavMenuID), h.Popover(),
				h.Li(h.A(h.Class("cui-globalnav-logo"), h.Href(n.Home), h.Aria("label", n.LogoAriaLabel), n.Logo)),
				h.Li(h.Button(h.Class("cui-globalnav-closebutton"), h.PopoverTarget(globalNavMenuID), h.PopoverTargetAction("hide"), h.Aria("label", n.Text.CloseAriaLabel), h.AutoFocus(), glyph.Menu)),
				links,
				searchItem,
				bagItem,
			),
			h.Button(h.Class("cui-globalnav-menubutton"), h.PopoverTarget(globalNavMenuID), h.Aria("label", n.Text.MenuAriaLabel), glyph.Menu),
			searchPopover,
			bagPopover,
			postForms,
		),
	).Render(w)
}

func globalNavFlyout(link NavLink, flyout NavFlyout, n int, text GlobalNavText) g.Node {
	id := "cui-globalnav-flyout-" + strconv.Itoa(n)
	var groups g.Group
	for i, group := range flyout.Groups {
		heading := id + "-" + strconv.Itoa(i+1)
		groups = append(groups, h.Div(h.Class("cui-globalnav-flyoutgroup"),
			h.H2(h.Class("cui-globalnav-flyoutheading"), h.ID(heading), g.Text(group.Heading)),
			h.Ul(h.Class("cui-globalnav-flyoutlinks"), h.Aria("labelledby", heading),
				g.Map(group.Large, func(l Link) g.Node {
					return h.Li(h.A(h.Class("cui-globalnav-flyoutlink"), h.Data("size", "lg"), h.Href(l.Href), g.Text(l.Label)))
				}),
				g.Map(group.Links, func(l Link) g.Node {
					return h.Li(h.A(h.Class("cui-globalnav-flyoutlink"), h.Href(l.Href), g.Text(l.Label)))
				}),
			),
		))
	}
	return g.Group{
		h.A(h.Class("cui-globalnav-link"), h.Href(link.Href), g.Text(link.Label)),
		h.Button(h.Class("cui-globalnav-flyoutbutton"), h.PopoverTarget(id), h.Aria("label", flyout.AriaLabel), globalNavImages(glyph.Flyout, glyph.FlyoutCompact)),
		h.Div(h.Class("cui-globalnav-flyout"), h.ID(id), h.Popover(),
			h.Div(h.Class("cui-globalnav-flyoutpanel"),
				h.Button(h.Class("cui-globalnav-backbutton"), h.PopoverTarget(id), h.PopoverTargetAction("hide"), h.Aria("label", text.MainMenuAriaLabel), glyph.FlyoutBack),
				h.Button(h.Class("cui-globalnav-closebutton"), h.PopoverTarget(globalNavMenuID), h.PopoverTargetAction("hide"), h.Aria("label", text.CloseAriaLabel), glyph.Menu),
				groups,
			),
			globalNavCurtain(id),
		),
	}
}
