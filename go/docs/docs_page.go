package docs

import (
	"io"

	"github.com/at-least/ciderui/go/icon"
	"github.com/at-least/ciderui/go/internal/class"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// DocsNavigatorID is fixed, so a document holds one DocsPage.
const DocsNavigatorID = "docs-navigator"

// DocsPage goes under the site's global nav; the navigator is beside the
// content from 1024px and a sheet below that.
type DocsPage struct {
	// Title and Href: as DocsBar's.
	Title     string
	Href      string
	Navigator []DocsNavigatorItem
	Class     string
	// Children: <div class="cui-docs-page-column">s, each holding a
	// [DocsHero], the cui-docs-content or a [DocsTopics].
	Children g.Node
}

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
