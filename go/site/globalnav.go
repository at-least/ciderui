package site

import (
	"io"
	"strconv"

	"github.com/at-least/ciderui/go/icon"
	"github.com/at-least/ciderui/go/internal/class"
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
	Text   string
	Href   string
	Flyout []FlyoutGroup
	// FlyoutAriaLabel: the button that opens Flyout ("Store menu"); "" with no Flyout.
	FlyoutAriaLabel string
}

type FlyoutGroup struct {
	Heading string
	// Large: apple.com's elevated links.
	Large []Link
	Links []Link
}

type Search interface{ search() }

type LinkedSearch struct {
	Href string
}

type FlyoutSearch struct {
	Action  string
	Heading string
	Links   []Link
}

func (LinkedSearch) search() {}
func (FlyoutSearch) search() {}

type Bag interface{ bag() }

type LinkedBag struct {
	Href string
}

// FlyoutBag is apple.com's empty-bag flyout.
type FlyoutBag struct {
	Message string
	Line    *BagLine
	Heading string
	Links   []BagLink
}

func (LinkedBag) bag() {}
func (FlyoutBag) bag() {}

type BagLine struct {
	Text string
	Href string
	// Rest follows the link with no space: start it with one
	// (" for items you saved earlier").
	Rest string
}

// BagLink's icons on apple.com: [icon.Orders], [icon.Saves],
// [icon.Account], [icon.SignIn], each with its Compact one.
type BagLink struct {
	Icon        icon.Icon
	IconCompact icon.Icon
	Text        string
	Href        string
}

// GlobalNavText is the nav's own words, in the page's language.
type GlobalNavText struct {
	AriaLabel      string
	MenuAriaLabel  string
	CloseAriaLabel string
	// MainMenuAriaLabel: a phone's flyout's back button.
	MainMenuAriaLabel    string
	SearchAriaLabel      string
	SearchPlaceholder    string
	ClearSearchAriaLabel string
	BagAriaLabel         string
}

var GlobalNavTextEnglish = GlobalNavText{
	AriaLabel:            "Global",
	MenuAriaLabel:        "Menu",
	CloseAriaLabel:       "Close",
	MainMenuAriaLabel:    "Main menu",
	SearchAriaLabel:      "Search",
	SearchPlaceholder:    "Search",
	ClearSearchAriaLabel: "Clear search",
	BagAriaLabel:         "Bag",
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
	Dark          bool
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
		{"MainMenuAriaLabel", n.Text.MainMenuAriaLabel},
		{"SearchAriaLabel", n.Text.SearchAriaLabel},
		{"SearchPlaceholder", n.Text.SearchPlaceholder},
		{"ClearSearchAriaLabel", n.Text.ClearSearchAriaLabel},
		{"BagAriaLabel", n.Text.BagAriaLabel},
	} {
		if text.value == "" {
			panic("site.GlobalNav's Text needs its " + text.field)
		}
	}
	for _, link := range n.Links {
		if link.Text == "" || link.Href == "" {
			panic("site.GlobalNav's Links need their Text and Href")
		}
		if len(link.Flyout) > 0 && link.FlyoutAriaLabel == "" {
			panic("site.GlobalNav's Links with a Flyout need their FlyoutAriaLabel")
		}
		for i, group := range link.Flyout {
			if group.Heading == "" {
				panic("site.GlobalNav's FlyoutGroups need their Heading")
			}
			if i > 0 && len(group.Large) > 0 {
				panic("site.GlobalNav's Large links are only in a flyout's first group")
			}
		}
	}
	if search, ok := n.Search.(FlyoutSearch); ok {
		if search.Action == "" {
			panic("site.GlobalNav's search needs its page (Action)")
		}
		if len(search.Links) > 0 && search.Heading == "" {
			panic("site.GlobalNav's search's Links need their Heading")
		}
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
			links = append(links, h.Li(globalNavFlyout(link, i+1, n.Text)))
		}
	}
	var searchItem, searchPopover g.Node
	switch search := n.Search.(type) {
	case nil:
	case LinkedSearch:
		searchItem = h.Li(h.A(h.Class("cui-globalnav-icon"), h.Href(search.Href), h.Aria("label", n.Text.SearchAriaLabel), globalNavImages(icon.Search, icon.SearchCompact)))
	case FlyoutSearch:
		searchItem = h.Li(h.Button(h.Class("cui-globalnav-icon"), h.PopoverTarget(globalNavSearchID), h.Aria("label", n.Text.SearchAriaLabel), globalNavImages(icon.Search, icon.SearchCompact)))
		searchPopover = h.Div(h.Class("cui-globalnav-search"), h.ID(globalNavSearchID), h.Popover(),
			h.Div(h.Class("cui-globalnav-searchpanel"),
				h.Form(h.Class("cui-globalnav-searchfield"), h.Action(search.Action), h.Method("get"), h.Role("search"),
					globalNavImages(icon.Searchfield, icon.SearchfieldCompact),
					h.Input(h.Class("cui-globalnav-searchinput"), h.Type("text"), h.Name("q"), h.Placeholder(n.Text.SearchPlaceholder), h.Aria("label", n.Text.SearchAriaLabel), h.AutoComplete("off"), h.AutoCapitalize("off"), h.AutoFocus()),
					h.Button(h.Class("cui-globalnav-searchreset"), h.Type("reset"), h.Aria("label", n.Text.ClearSearchAriaLabel), globalNavImages(icon.Clear, icon.ClearCompact)),
				),
				g.If(len(search.Links) > 0, g.Group{
					h.H2(h.Class("cui-globalnav-searchheading"), h.ID(globalNavHeadingID), g.Text(search.Heading)),
					h.Ul(h.Class("cui-globalnav-searchlinks"), h.Aria("labelledby", globalNavHeadingID), g.Map(search.Links, func(link Link) g.Node {
						return h.Li(h.A(h.Class("cui-globalnav-searchlink"), h.Href(link.Href), globalNavImages(icon.Arrow, icon.ArrowCompact), g.Text(link.Text)))
					})),
				}),
				h.Button(h.Class("cui-globalnav-closebutton"), h.PopoverTarget(globalNavSearchID), h.PopoverTargetAction("hide"), h.Aria("label", n.Text.CloseAriaLabel), icon.Menu),
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
		bagItem = h.Li(h.A(h.Class("cui-globalnav-icon"), h.Href(bag.Href), h.Aria("label", n.Text.BagAriaLabel), globalNavImages(icon.Bag, icon.BagCompact)))
	case FlyoutBag:
		bagItem = h.Li(h.Button(h.Class("cui-globalnav-icon"), h.PopoverTarget(globalNavBagID), h.Aria("label", n.Text.BagAriaLabel), globalNavImages(icon.Bag, icon.BagCompact)))
		bagPopover = h.Div(h.Class("cui-globalnav-bag"), h.ID(globalNavBagID), h.Popover(),
			h.Div(h.Class("cui-globalnav-bagpanel"),
				h.Button(h.Class("cui-globalnav-closebutton"), h.PopoverTarget(globalNavBagID), h.PopoverTargetAction("hide"), h.Aria("label", n.Text.CloseAriaLabel), h.AutoFocus(), icon.Menu),
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
	return h.Nav(class.Root(block, n.Class), g.If(n.Clear, h.Data("variant", "clear")), h.Aria("label", n.Text.AriaLabel),
		h.Div(h.Class("cui-globalnav-content"),
			h.Ul(h.Class("cui-globalnav-list"), h.ID(globalNavMenuID), h.Popover(),
				h.Li(h.A(h.Class("cui-globalnav-logo"), h.Href(n.Home), h.Aria("label", n.LogoAriaLabel), n.Logo)),
				h.Li(h.Button(h.Class("cui-globalnav-closebutton"), h.PopoverTarget(globalNavMenuID), h.PopoverTargetAction("hide"), h.Aria("label", n.Text.CloseAriaLabel), h.AutoFocus(), icon.Menu)),
				links,
				searchItem,
				bagItem,
			),
			h.Button(h.Class("cui-globalnav-menubutton"), h.PopoverTarget(globalNavMenuID), h.Aria("label", n.Text.MenuAriaLabel), icon.Menu),
			searchPopover,
			bagPopover,
		),
	).Render(w)
}

func globalNavFlyout(link NavLink, n int, text GlobalNavText) g.Node {
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
		h.Button(h.Class("cui-globalnav-flyoutbutton"), h.PopoverTarget(id), h.Aria("label", link.FlyoutAriaLabel), globalNavImages(icon.Flyout, icon.FlyoutCompact)),
		h.Div(h.Class("cui-globalnav-flyout"), h.ID(id), h.Popover(),
			h.Div(h.Class("cui-globalnav-flyoutpanel"),
				h.Button(h.Class("cui-globalnav-backbutton"), h.PopoverTarget(id), h.PopoverTargetAction("hide"), h.Aria("label", text.MainMenuAriaLabel), icon.FlyoutBack),
				h.Button(h.Class("cui-globalnav-closebutton"), h.PopoverTarget(globalNavMenuID), h.PopoverTargetAction("hide"), h.Aria("label", text.CloseAriaLabel), icon.Menu),
				groups,
			),
			globalNavCurtain(id),
		),
	}
}
