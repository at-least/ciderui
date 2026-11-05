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

func TestQuickLinksPanics(t *testing.T) {
	example.Panics(t, func() site.QuickLinks {
		return site.QuickLinks{ID: "q", Title: "x", Links: nil, Class: ""}
	}, []example.Break[site.QuickLinks]{
		{Want: "site.QuickLinks needs an ID", Change: func(p *site.QuickLinks) { p.ID = "" }},
		{Want: "site.QuickLinks needs a Title", Change: func(p *site.QuickLinks) { p.Title = "" }},
	})
}
