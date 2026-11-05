package site_test

import (
	"testing"

	"github.com/at-least/ciderui/go/internal/example"
	"github.com/at-least/ciderui/go/site"
)

func TestSkipLinkRendersTheExample(t *testing.T) {
	example.Check(t, site.SkipLink{Href: "#main", Label: "Skip to Content", Class: ""}, "site/skip/examples/skip.html")
}

func TestSkipLinkPanics(t *testing.T) {
	example.Panics(t, func() site.SkipLink {
		return site.SkipLink{Href: "#main", Label: "x", Class: ""}
	}, []example.Break[site.SkipLink]{
		{Want: "site.SkipLink needs an Href", Change: func(p *site.SkipLink) { p.Href = "" }},
		{Want: "site.SkipLink needs a Label", Change: func(p *site.SkipLink) { p.Label = "" }},
	})
}
