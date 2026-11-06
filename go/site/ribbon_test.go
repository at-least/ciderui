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
		h.A(h.Class("cui-link"), h.Href("#checkout"), g.Text("Check out now")),
	)}, "site/ribbon/examples/ribbon.html")
}
