package site_test

import (
	"testing"

	"github.com/at-least/ciderui/go/internal/example"
	"github.com/at-least/ciderui/go/site"
)

func TestMoreLinkRendersTheExample(t *testing.T) {
	example.Check(t, site.MoreLink{Href: "#learn", Label: "Learn more", Class: ""}, "site/link/examples/more.html")
}

func TestMoreLinkPanics(t *testing.T) {
	example.Panics(t, func() site.MoreLink {
		return site.MoreLink{Href: "/x", Label: "x", Class: ""}
	}, []example.Break[site.MoreLink]{
		{Want: "site.MoreLink needs an Href", Change: func(p *site.MoreLink) { p.Href = "" }},
		{Want: "site.MoreLink needs a Label", Change: func(p *site.MoreLink) { p.Label = "" }},
	})
}
