package site_test

import (
	"testing"

	"github.com/at-least/ciderui/go/internal/example"
	"github.com/at-least/ciderui/go/site"
)

func TestQuickLinksRendersTheExample(t *testing.T) {
	links := []site.Link{
		{Label: "Visit a Store", Href: "/retail/"},
		{Label: "Track an Order", Href: "/shop/order/list"},
		{Label: "Help with Buying", Href: "/shop/help"},
	}
	example.Check(t, site.QuickLinks{ID: "quicklinks", Title: "Shortcuts", Links: links, Class: ""}, "site/quicklinks/examples/quicklinks.html")
}

func TestQuickLinksAnEmptyTitlePanics(t *testing.T) {
	mustPanic(t, "needs a Title", func() {
		example.Render(site.QuickLinks{ID: "q", Title: "", Links: nil, Class: ""})
	})
}

func TestQuickLinksAnEmptyIDPanics(t *testing.T) {
	mustPanic(t, "QuickLinks needs an ID", func() {
		example.Render(site.QuickLinks{ID: "", Title: "Shortcuts", Links: nil, Class: ""})
	})
}
