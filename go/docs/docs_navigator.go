package docs

import (
	"io"
	"strings"

	"github.com/at-least/ciderui/go/icon"
	"github.com/at-least/ciderui/go/internal/class"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

type DocsNavigatorItem interface {
	docsNavigatorItem()
}

// DocsNavigatorGroup is a title over the pages after it, up to the next
// group.
type DocsNavigatorGroup struct {
	Title string
}

type DocsNavigatorPage struct {
	Label string
	// Href: its path's words must be unique among the navigator's pages, or
	// two chevrons share one checkbox id (chevronID); nothing checks it.
	Href    string
	Current bool
	// Open: set it on the current page's ancestors, as DocC does.
	Open     bool
	Children []DocsNavigatorItem
}

func (DocsNavigatorGroup) docsNavigatorItem() {}
func (DocsNavigatorPage) docsNavigatorItem()  {}

// DocsNavigator's filter and its no-results line start hidden: the page's
// script unhides the filter, filters the rows and writes what it found into
// the status line.
type DocsNavigator struct {
	Title string
	Href  string
	Items []DocsNavigatorItem
	Class string
}

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
