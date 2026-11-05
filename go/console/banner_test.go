package console_test

import (
	"testing"

	"github.com/at-least/ciderui/go/console"
	"github.com/at-least/ciderui/go/icon"
	"github.com/at-least/ciderui/go/internal/example"
	"github.com/at-least/ciderui/go/lucide"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

func TestBannerRendersTheExamples(t *testing.T) {
	warning := console.ConsoleBanner{Color: console.ConsoleColorWarning, Icon: lucide.TriangleAlert, Title: "Symbols are missing", Class: "", Children: g.Group{
		h.P(h.Class("cui-console-banner-body"), g.Text("Upload a dSYM in "), textLink("#settings", "Settings"), g.Text(" to see function names in this crash.")),
		h.P(h.Class("cui-console-banner-actions"), h.A(h.Class("cui-console-btn"), h.Data("size", "sm"), h.Href("#upload"), g.Text("Upload dSYM"))),
	}}
	example.Check(t, warning, "console/console-banner/examples/banner.html")

	rows := []struct {
		color console.ConsoleColor
		icon  icon.Icon
		body  g.Group
	}{
		{console.ConsoleColorNeutral, lucide.Info, g.Group{g.Text("Data is up to date as of 9:41. See the "), textLink("#sync-log", "sync log"), g.Text(".")}},
		{console.ConsoleColorInfo, lucide.Info, g.Group{g.Text("A new version of the SDK is out. Read the "), textLink("#release-notes", "release notes"), g.Text(".")}},
		{console.ConsoleColorSuccess, lucide.CircleCheck, g.Group{g.Text("The alert channel works. See the "), textLink("#delivery-log", "delivery log"), g.Text(".")}},
		{console.ConsoleColorError, lucide.CircleX, g.Group{g.Text("The upload failed. "), textLink("#retry", "Try again"), g.Text(" or check your connection.")}},
	}
	var banners g.Group
	for _, r := range rows {
		banners = append(banners, console.ConsoleBanner{Color: r.color, Icon: r.icon, Title: "", Class: "", Children: h.P(h.Class("cui-console-banner-body"), r.body)})
	}
	example.Check(t, h.Div(h.Class("grid gap-3"), banners), "console/console-banner/examples/colors.html")
}

func textLink(href, text string) g.Node {
	return h.A(h.Class("cui-console-link"), h.Data("variant", "text"), h.Href(href), g.Text(text))
}
