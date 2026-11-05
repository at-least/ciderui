package site_test

import (
	"testing"

	"github.com/at-least/ciderui/go/internal/example"
	"github.com/at-least/ciderui/go/site"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

func TestRibbonRendersTheExamples(t *testing.T) {
	example.Check(t, site.Ribbon{Class: "", Children: h.P(
		g.Text("Pre-order Cider Phone Fold now. Already have one saved for pre-order? "),
		site.MoreLink{Href: "#checkout", Label: "Check out now", Class: ""},
	)}, "site/ribbon/examples/ribbon.html")
}
