package site_test

import (
	"testing"

	"github.com/at-least/ciderui/go/internal/example"
	"github.com/at-least/ciderui/go/site"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

func company() g.Node { return h.A(h.Href("/company/"), g.Text("Example Inc.")) }

func TestFooterRendersTheExample(t *testing.T) {
	example.Check(t, site.Footer{Copyright: company()}, "site/footer/examples/footer.html")
}

func TestFooterPanics(t *testing.T) {
	example.Panics(t, func() site.Footer {
		return site.Footer{Copyright: g.Text("©")}
	}, []example.Break[site.Footer]{
		{Want: "site.Footer needs a Copyright", Change: func(p *site.Footer) { p.Copyright = nil }},
	})
}
