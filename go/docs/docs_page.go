package docs

import (
	"io"

	"github.com/at-least/ciderui/go/icon"
	"github.com/at-least/ciderui/go/internal/class"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// DocsNavigatorID is the navigator's id: the bar's button and the sheet's
// close button name it (popovertarget). One docs page to a document.
const DocsNavigatorID = "docs-navigator"

// DocsPage is cui-docs-page (components/docs/docs-page): a page of Apple's
// developer documentation, under the site's global nav: the documentation's
// bar ([DocsBar]) across it, the navigator ([DocsNavigator]) beside the
// content from 1024px and a sheet the bar's button opens below that, with its
// close button, and the page's <main>, whose children are the page's: its
// columns (cui-docs-page-column), each holding a hero, the content or the
// topics.
type DocsPage struct {
	// Title is the documentation's title, the bar's and the navigator's.
	// Required: an empty one panics.
	Title string
	// Href is where the title links: the documentation's first page.
	Href string
	// Navigator is the navigator's rows.
	Navigator []DocsNavigatorItem
	// Class is the page's utilities on the page; "" for none.
	Class string
	// Children is the page's <main>: its columns.
	Children g.Node
}

// Render writes the page.
func (p DocsPage) Render(w io.Writer) error {
	return h.Div(class.Root("cui-docs-page", p.Class),
		DocsBar{Title: p.Title, Href: p.Href, Navigator: DocsNavigatorID, Class: ""},
		h.Div(h.Class("cui-docs-page-sidebar"), h.ID(DocsNavigatorID), h.Popover(),
			h.Button(h.Class("cui-docs-page-close"), h.Type("button"), h.PopoverTarget(DocsNavigatorID), h.PopoverTargetAction("hide"), h.Aria("label", "Close Navigator"), icon.DocsClose),
			DocsNavigator{Title: p.Title, Href: p.Href, Items: p.Navigator, Class: ""},
		),
		h.Main(h.Class("cui-docs-page-main"), p.Children),
	).Render(w)
}
