package docs

import (
	"io"
	"strings"

	"github.com/at-least/ciderui/go/icon"
	"github.com/at-least/ciderui/go/internal/class"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// DocsNavigatorItem is a row of a [DocsNavigator]: a [DocsNavigatorGroup], a
// group's title over the pages after it, or a [DocsNavigatorPage], a page.
// The two are its only types (its method is unexported).
type DocsNavigatorItem interface {
	docsNavigatorItem()
}

// DocsNavigatorGroup is a group's title in a [DocsNavigator], the pages after
// it its own until the next.
type DocsNavigatorGroup struct {
	// Title is the group's title.
	Title string
}

// DocsNavigatorPage is a page in a [DocsNavigator], and the pages under it: a
// collection's glyph where it has pages under it, which its chevron shows
// and hides, an article's where not.
type DocsNavigatorPage struct {
	// Label is its name. Required: an empty one panics.
	Label string
	// Href is its address. Its path unique among the navigator's pages: a
	// page with pages under it names its chevron's checkbox after its path's
	// words.
	Href string
	// Current is whether it is the page shown (aria-current="page").
	Current bool
	// Open is whether the pages under it are shown at first (the current
	// page's ancestors' are, as DocC's); false for a page with none: an open
	// one with none panics.
	Open bool
	// Children are the pages under it, and their groups' titles; nil for
	// none.
	Children []DocsNavigatorItem
}

func (DocsNavigatorGroup) docsNavigatorItem() {}
func (DocsNavigatorPage) docsNavigatorItem()  {}

// DocsNavigator is cui-docs-navigator (components/docs/docs-navigator): the
// navigator of Apple's developer documentation, its pages as a tree, a <nav>
// named Navigator (the bar's button's name): the documentation's title, a
// link to its first page, over its rows. A row with pages under it opens and
// closes by its chevron, a checkbox and its label (no script); the current
// page's row is marked. Under the rows, the line saying the filter matches
// nothing, the filter at the navigator's foot and a status read aloud, the
// first two hidden: the page's script shows the filter, filters the rows and
// says what it found.
type DocsNavigator struct {
	// Title is the documentation's title, at the navigator's top. Required:
	// an empty one panics.
	Title string
	// Href is where the title links: the documentation's first page.
	Href string
	// Items are its rows, from the top.
	Items []DocsNavigatorItem
	// Class is the page's utilities on the navigator; "" for none.
	Class string
}

// Render writes the navigator.
func (n DocsNavigator) Render(w io.Writer) error {
	if n.Title == "" {
		panic("docs.DocsNavigator needs a Title")
	}
	return h.Nav(class.Root("cui-docs-navigator", n.Class), h.Aria("label", "Navigator"),
		h.A(h.Class("cui-docs-navigator-title"), h.Href(n.Href), g.Text(n.Title)),
		navigatorRows(n.Items),
		h.P(h.Class("cui-docs-navigator-empty"), g.Attr("hidden"), g.Text("No results found.")),
		h.Div(h.Class("cui-docs-navigator-filter"), h.Role("search"), g.Attr("hidden"),
			h.Label(h.Class("cui-docs-navigator-field"),
				icon.DocsFilter.Class("cui-docs-navigator-filter-icon"),
				h.Input(h.Class("cui-docs-navigator-input"), h.Type("search"), h.Placeholder("Filter"), h.Aria("label", "Filter"), h.AutoComplete("off"), h.SpellCheck("false")),
			),
		),
		h.P(h.Class("cui-docs-navigator-status"), h.Aria("live", "polite")),
	).Render(w)
}

// the rows of one level, as a list
func navigatorRows(items []DocsNavigatorItem) g.Node {
	return h.Ul(h.Class("cui-docs-navigator-list"), g.Map(items, navigatorRow))
}

func navigatorRow(item DocsNavigatorItem) g.Node {
	switch item := item.(type) {
	case DocsNavigatorGroup:
		return h.Li(h.Class("cui-docs-navigator-group"), g.Text(item.Title))
	case DocsNavigatorPage:
		return navigatorPage(item)
	}
	panic("docs.DocsNavigator: a nil item")
}

func navigatorPage(page DocsNavigatorPage) g.Node {
	if page.Label == "" {
		panic("docs.DocsNavigator: a page needs a Label")
	}
	if page.Open && len(page.Children) == 0 {
		panic("docs.DocsNavigator: " + page.Label + " is open with no pages under it")
	}
	if len(page.Children) == 0 {
		return h.Li(h.Class("cui-docs-navigator-item"),
			h.Div(h.Class("cui-docs-navigator-row"),
				icon.DocsArticle.Class("cui-docs-navigator-icon"),
				navigatorLink(page),
			),
		)
	}
	id := chevronID(page.Href)
	return h.Li(h.Class("cui-docs-navigator-item"),
		h.Input(h.Class("cui-docs-navigator-toggle"), h.Type("checkbox"), h.ID(id), h.Aria("label", page.Label), g.If(page.Open, h.Checked())),
		h.Div(h.Class("cui-docs-navigator-row"),
			h.Label(h.Class("cui-docs-navigator-chevron"), h.For(id), icon.DocsChevron),
			icon.DocsCollection.Class("cui-docs-navigator-icon"),
			navigatorLink(page),
		),
		navigatorRows(page.Children),
	)
}

func navigatorLink(page DocsNavigatorPage) g.Node {
	return h.A(h.Class("cui-docs-navigator-link"), h.Href(page.Href), g.If(page.Current, h.Aria("current", "page")), g.Text(page.Label))
}

// chevronID is the chevron's checkbox id, named after the page's address, its
// path's words alone (/components/ and https://example.test/components/
// docs-navigator-components).
func chevronID(href string) string {
	path := href
	if _, rest, ok := strings.Cut(href, "://"); ok {
		path = ""
		if i := strings.IndexByte(rest, '/'); i >= 0 {
			path = rest[i:]
		}
	}
	words := strings.FieldsFunc(path, func(r rune) bool {
		return !(r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z' || r >= '0' && r <= '9')
	})
	id := "docs-navigator"
	for _, word := range words {
		id += "-" + strings.ToLower(word)
	}
	return id
}
