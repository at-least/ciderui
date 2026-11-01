package site_test

import (
	"testing"

	"github.com/at-least/ciderui/go/internal/example"
	"github.com/at-least/ciderui/go/site"
)

func TestQuickLinksRendersTheExample(t *testing.T) {
	links := []site.Link{
		{Text: "Visit a Store", Href: "/retail/"},
		{Text: "Track an Order", Href: "/shop/order/list"},
		{Text: "Help with Buying", Href: "/shop/help"},
	}
	example.Check(t, site.QuickLinks{ID: "quicklinks", Title: "Shortcuts", Links: links, Class: ""}, "site/quicklinks/examples/quicklinks.html")
}

func TestQuickLinksAnEmptyTitlePanics(t *testing.T) {
	mustPanic(t, "needs a Title", func() {
		example.Render(site.QuickLinks{ID: "q", Title: "", Links: nil, Class: ""})
	})
}
