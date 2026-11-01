package site_test

import (
	"testing"

	"github.com/at-least/ciderui/go/internal/example"
	"github.com/at-least/ciderui/go/site"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

func TestRibbonRendersTheExamples(t *testing.T) {
	for announce, file := range map[bool]string{
		false: "site/ribbon/examples/ribbon.html",
		true:  "site/ribbon/examples/announce.html",
	} {
		example.Check(t, site.Ribbon{Announce: announce, Class: "", Children: h.P(
			g.Text("Pre-order Cider Phone Fold now. Already have one saved for pre-order? "),
			site.MoreLink{Href: "#checkout", Text: "Check out now", Class: ""},
		)}, file)
	}
}
