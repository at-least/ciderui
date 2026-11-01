package site

import (
	"io"
	"strconv"

	"github.com/at-least/ciderui/go/icon"
	"github.com/at-least/ciderui/go/internal/class"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// the list's id, which its button opens, and the search's: a page has one
// global nav; a flyout's is cui-globalnav-flyout-<n>, n its link's place
const (
	globalNavMenuID       = "cui-globalnav-menu"
	globalNavSearchID     = "cui-globalnav-search"
	globalNavHeadingID    = "cui-globalnav-searchheading"
	globalNavBagID        = "cui-globalnav-bag"
	globalNavBagHeadingID = "cui-globalnav-bagheading"
)

// NavLink is a link of the global nav ([GlobalNav]'s Links): its text, where
// it goes, and the groups of the flyout it opens, if it opens one: from 834px
// a menu under the bar, which the pointer resting on the link shows, its
// button (the keyboard's, and on a screen with no hover a tap on the link)
// opens; to 833px the menu's row opens it over the menu.
type NavLink struct {
	// Text is its text. Required: an empty one panics.
	Text string
	// Href is where it goes. Required: an empty one panics.
	Href string
	// Flyout is its flyout's groups, in order; nil for no flyout.
	Flyout []FlyoutGroup
}

// FlyoutGroup is a group of a flyout ([NavLink]'s Flyout): a heading over its
// links; its large links first (apple.com's elevated group, the flyout's
// first: 24px from 834px, 28px to 833px, where its heading is not shown),
// then its small ones.
type FlyoutGroup struct {
	// Heading is its heading, which names its list. Required: an empty one
	// panics.
	Heading string
	// Large are its large links; nil for none.
	Large []Link
	// Links are its small links, after the large ones; nil for none.
	Links []Link
}

// Search is the global nav's search ([GlobalNav]'s Search): a
// [LinkedSearch] or a [FlyoutSearch]; nil for none.
type Search interface{ search() }

// LinkedSearch is a link to the site's search page.
type LinkedSearch struct {
	// Href is the site's search page.
	Href string
}

// FlyoutSearch is a button opening the search with no script: its field,
// focused, sent as q (GET) to Action, the site's search page, its clear
// button shown while it holds text, and under it the quick links under the
// heading Shortcuts (none, no heading); from 834px a flyout under the bar, to
// 833px a sheet over the window.
type FlyoutSearch struct {
	// Action is the site's search page. Required: an empty one panics.
	Action string
	// Links are the quick links, in order; nil for none.
	Links []Link
}

func (LinkedSearch) search() {}
func (FlyoutSearch) search() {}

// Bag is the global nav's bag ([GlobalNav]'s Bag): a [LinkedBag] or a
// [FlyoutBag]; nil for none.
type Bag interface{ bag() }

// LinkedBag is a link to the site's bag.
type LinkedBag struct {
	// Href is the site's bag.
	Href string
}

// FlyoutBag is a button opening the bag with no script, an empty bag's, as
// apple.com's: its message, a line under it (a link, then the rest of its
// sentence), and its links under their heading, each after its icons
// ([BagLink]); from 834px a flyout under the bar, to 833px a sheet over the
// window.
type FlyoutBag struct {
	// Message is its message (Nothing here yet.). Required: an empty one
	// panics.
	Message string
	// Line is the line under it; nil for none.
	Line *BagLine
	// Heading is the heading over its links (Account), which names their
	// list. Required where there are links.
	Heading string
	// Links are its links, in order; nil for none, and no heading.
	Links []BagLink
}

func (LinkedBag) bag() {}
func (FlyoutBag) bag() {}

// BagLine is the line under a [FlyoutBag]'s message: a link, then the rest
// of its sentence (Sign in, /signin, " for items you saved earlier").
type BagLine struct {
	// Text is the link's text. Required: an empty one panics.
	Text string
	// Href is where the link goes. Required: an empty one panics.
	Href string
	// Rest is the rest of the sentence, after the link.
	Rest string
}

// BagLink is a link of the bag's flyout ([FlyoutBag]'s Links): its icons,
// the one from 834px and the one to 833px (apple.com's: [icon.Orders],
// [icon.Saves], [icon.Account], [icon.SignIn] and each one's Compact), its
// text and where it goes.
type BagLink struct {
	// Icon is its icon from 834px.
	Icon icon.Icon
	// IconCompact is its icon to 833px.
	IconCompact icon.Icon
	// Text is its text. Required: an empty one panics.
	Text string
	// Href is where it goes. Required: an empty one panics.
	Href string
}

// GlobalNav is cui-globalnav (components/site/globalnav): the bar across the
// top of the page, the site's logo, its links and its search and bag; to
// 833px the links are a menu its button opens over the window, with no script
// (the list a popover beside the button, with a close button of its own that
// takes the focus); a link may open a flyout ([NavLink]). Not placed on the
// page: a site adds sticky top-0 (its Class) where the nav stays; its
// z-index, apple.com's 9999, is the nav's own, over the page and a local nav
// (a z-* utility in its Class would replace it).
type GlobalNav struct {
	// Logo is the site's logo, its image (an <svg> in currentColor takes the
	// bar's color), in the logo's link.
	Logo g.Node
	// Home is where the logo goes, the site's home. Required: an empty one
	// panics.
	Home string
	// Name is the site's name, the logo link's accessible name. Required: an
	// empty one panics.
	Name string
	// Links are the links, in order, each perhaps opening a flyout. At least
	// one: a menu of nothing panics.
	Links []NavLink
	// Search is the site's search, if it has one ([icon.Search]): a link to
	// its page, or the search opened over the page; nil for none.
	Search Search
	// Bag is the site's bag, if it has one ([icon.Bag]): a link to its page,
	// or the bag opened over the page; nil for none.
	Bag Bag
	// Dark is whether the nav is dark (cui-dark): the dark bar and menu.
	Dark bool
	// Clear is whether the bar is clear (data-variant="clear"): no ground and
	// no blur, over a hero.
	Clear bool
	// Class is the page's utilities on the nav; "" for none.
	Class string
}

// an icon's two sizes, from 834px and to 833px
func globalNavImages(regular, compact icon.Icon) g.Node {
	return g.Group{regular.Class("cui-globalnav-image"), compact.Class("cui-globalnav-image-compact")}
}

// a popover's curtain, the click outside it that closes it
func globalNavCurtain(id string) g.Node {
	return h.Button(h.Class("cui-globalnav-curtain"), h.PopoverTarget(id), h.PopoverTargetAction("hide"), h.TabIndex("-1"), h.Aria("hidden", "true"))
}

// Render writes the nav.
func (n GlobalNav) Render(w io.Writer) error {
	if n.Home == "" {
		panic("site.GlobalNav needs a Home")
	}
	if n.Name == "" {
		panic("site.GlobalNav needs a Name")
	}
	if len(n.Links) == 0 {
		panic("site.GlobalNav needs Links")
	}
	for _, link := range n.Links {
		if link.Text == "" || link.Href == "" {
			panic("site.GlobalNav's Links need their Text and Href")
		}
		for _, group := range link.Flyout {
			if group.Heading == "" {
				panic("site.GlobalNav's FlyoutGroups need their Heading")
			}
		}
	}
	if search, ok := n.Search.(FlyoutSearch); ok && search.Action == "" {
		panic("site.GlobalNav's search needs its page (Action)")
	}
	if bag, ok := n.Bag.(FlyoutBag); ok {
		if bag.Message == "" {
			panic("site.GlobalNav's bag needs its Message")
		}
		if bag.Line != nil && (bag.Line.Text == "" || bag.Line.Href == "") {
			panic("site.GlobalNav's bag's Line needs its link's Text and Href")
		}
		if len(bag.Links) > 0 && bag.Heading == "" {
			panic("site.GlobalNav's bag's Links need their Heading")
		}
		for _, link := range bag.Links {
			if link.Text == "" || link.Href == "" {
				panic("site.GlobalNav's bag's Links need their Text and Href")
			}
		}
	}
	block := "cui-globalnav"
	if n.Dark {
		block = "cui-globalnav cui-dark"
	}
	var links g.Group
	for i, link := range n.Links {
		if len(link.Flyout) == 0 {
			links = append(links, h.Li(h.A(h.Class("cui-globalnav-link"), h.Href(link.Href), g.Text(link.Text))))
		} else {
			links = append(links, h.Li(globalNavFlyout(link, i+1)))
		}
	}
	var searchItem, searchPopover g.Node
	switch search := n.Search.(type) {
	case nil:
	case LinkedSearch:
		searchItem = h.Li(h.A(h.Class("cui-globalnav-icon"), h.Href(search.Href), h.Aria("label", "Search"), globalNavImages(icon.Search, icon.SearchCompact)))
	case FlyoutSearch:
		searchItem = h.Li(h.Button(h.Class("cui-globalnav-icon"), h.PopoverTarget(globalNavSearchID), h.Aria("label", "Search"), globalNavImages(icon.Search, icon.SearchCompact)))
		searchPopover = h.Div(h.Class("cui-globalnav-search"), h.ID(globalNavSearchID), h.Popover(),
			h.Div(h.Class("cui-globalnav-searchpanel"),
				h.Form(h.Class("cui-globalnav-searchfield"), h.Action(search.Action), h.Method("get"), h.Role("search"),
					globalNavImages(icon.Searchfield, icon.SearchfieldCompact),
					h.Input(h.Class("cui-globalnav-searchinput"), h.Type("text"), h.Name("q"), h.Placeholder("Search"), h.Aria("label", "Search"), h.AutoComplete("off"), h.AutoCapitalize("off"), h.AutoFocus()),
					h.Button(h.Class("cui-globalnav-searchreset"), h.Type("reset"), h.Aria("label", "Clear search"), globalNavImages(icon.Clear, icon.ClearCompact)),
				),
				g.If(len(search.Links) > 0, g.Group{
					h.H2(h.Class("cui-globalnav-searchheading"), h.ID(globalNavHeadingID), g.Text("Shortcuts")),
					h.Ul(h.Class("cui-globalnav-searchlinks"), h.Aria("labelledby", globalNavHeadingID), g.Map(search.Links, func(link Link) g.Node {
						return h.Li(h.A(h.Class("cui-globalnav-searchlink"), h.Href(link.Href), globalNavImages(icon.Arrow, icon.ArrowCompact), g.Text(link.Text)))
					})),
				}),
				h.Button(h.Class("cui-globalnav-closebutton"), h.PopoverTarget(globalNavSearchID), h.PopoverTargetAction("hide"), h.Aria("label", "Close"), icon.Menu),
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
		bagItem = h.Li(h.A(h.Class("cui-globalnav-icon"), h.Href(bag.Href), h.Aria("label", "Bag"), globalNavImages(icon.Bag, icon.BagCompact)))
	case FlyoutBag:
		bagItem = h.Li(h.Button(h.Class("cui-globalnav-icon"), h.PopoverTarget(globalNavBagID), h.Aria("label", "Bag"), globalNavImages(icon.Bag, icon.BagCompact)))
		bagPopover = h.Div(h.Class("cui-globalnav-bag"), h.ID(globalNavBagID), h.Popover(),
			h.Div(h.Class("cui-globalnav-bagpanel"),
				h.Button(h.Class("cui-globalnav-closebutton"), h.PopoverTarget(globalNavBagID), h.PopoverTargetAction("hide"), h.Aria("label", "Close"), h.AutoFocus(), icon.Menu),
				h.H2(h.Class("cui-globalnav-bagmessage"), g.Text(bag.Message)),
				g.Iff(bag.Line != nil, func() g.Node {
					return h.P(h.Class("cui-globalnav-bagtext"), h.A(h.Href(bag.Line.Href), g.Text(bag.Line.Text)), g.Text(bag.Line.Rest))
				}),
				g.If(len(bag.Links) > 0, g.Group{
					h.H3(h.Class("cui-globalnav-bagheading"), h.ID(globalNavBagHeadingID), g.Text(bag.Heading)),
					h.Ul(h.Class("cui-globalnav-baglinks"), h.Aria("labelledby", globalNavBagHeadingID), g.Map(bag.Links, func(link BagLink) g.Node {
						return h.Li(h.A(h.Class("cui-globalnav-baglink"), h.Href(link.Href), globalNavImages(link.Icon, link.IconCompact), g.Text(link.Text)))
					})),
				}),
			),
			globalNavCurtain(globalNavBagID),
		)
	default:
		panic("site.GlobalNav's Bag is neither a LinkedBag nor a FlyoutBag")
	}
	return h.Nav(class.Root(block, n.Class), g.If(n.Clear, h.Data("variant", "clear")), h.Aria("label", "Global"),
		h.Div(h.Class("cui-globalnav-content"),
			h.Ul(h.Class("cui-globalnav-list"), h.ID(globalNavMenuID), h.Popover(),
				h.Li(h.A(h.Class("cui-globalnav-logo"), h.Href(n.Home), h.Aria("label", n.Name), n.Logo)),
				h.Li(h.Button(h.Class("cui-globalnav-closebutton"), h.PopoverTarget(globalNavMenuID), h.PopoverTargetAction("hide"), h.Aria("label", "Close"), h.AutoFocus(), icon.Menu)),
				links,
				searchItem,
				bagItem,
			),
			h.Button(h.Class("cui-globalnav-menubutton"), h.PopoverTarget(globalNavMenuID), h.Aria("label", "Menu"), icon.Menu),
			searchPopover,
			bagPopover,
		),
	).Render(w)
}

// a link with its flyout, the nth link: the link, the flyout's button, the
// flyout (a popover) with its panel (the sheet's back button and X, its
// groups) and its curtain
func globalNavFlyout(link NavLink, n int) g.Node {
	id := "cui-globalnav-flyout-" + strconv.Itoa(n)
	var groups g.Group
	for i, group := range link.Flyout {
		heading := id + "-" + strconv.Itoa(i+1)
		groups = append(groups, h.Div(h.Class("cui-globalnav-flyoutgroup"),
			h.H2(h.Class("cui-globalnav-flyoutheading"), h.ID(heading), g.Text(group.Heading)),
			h.Ul(h.Class("cui-globalnav-flyoutlinks"), h.Aria("labelledby", heading),
				g.Map(group.Large, func(l Link) g.Node {
					return h.Li(h.A(h.Class("cui-globalnav-flyoutlink"), h.Data("size", "lg"), h.Href(l.Href), g.Text(l.Text)))
				}),
				g.Map(group.Links, func(l Link) g.Node {
					return h.Li(h.A(h.Class("cui-globalnav-flyoutlink"), h.Href(l.Href), g.Text(l.Text)))
				}),
			),
		))
	}
	return g.Group{
		h.A(h.Class("cui-globalnav-link"), h.Href(link.Href), g.Text(link.Text)),
		h.Button(h.Class("cui-globalnav-flyoutbutton"), h.PopoverTarget(id), h.Aria("label", link.Text+" menu"), globalNavImages(icon.Flyout, icon.FlyoutCompact)),
		h.Div(h.Class("cui-globalnav-flyout"), h.ID(id), h.Popover(),
			h.Div(h.Class("cui-globalnav-flyoutpanel"),
				h.Button(h.Class("cui-globalnav-backbutton"), h.PopoverTarget(id), h.PopoverTargetAction("hide"), h.Aria("label", "Main menu"), icon.FlyoutBack),
				h.Button(h.Class("cui-globalnav-closebutton"), h.PopoverTarget(globalNavMenuID), h.PopoverTargetAction("hide"), h.Aria("label", "Close"), icon.Menu),
				groups,
			),
			globalNavCurtain(id),
		),
	}
}
