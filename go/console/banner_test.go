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
	warning := console.ConsoleBanner{Color: console.ConsoleColorWarning, Icon: lucide.TriangleAlert, Title: "Symbols are missing", Dismiss: "Close", Class: "", Children: g.Group{
		h.P(h.Class("cui-console-banner-body"), g.Text("Upload a dSYM to see function names in this crash.")),
		h.P(h.Class("cui-console-banner-actions"), h.A(h.Class("cui-console-btn"), h.Data("size", "sm"), h.Href("#upload"), g.Text("Upload dSYM"))),
	}}
	example.Check(t, warning, "console/console-banner/examples/banner.html")

	rows := []struct {
		color console.ConsoleColor
		icon  icon.Icon
		text  string
	}{
		{console.ConsoleColorNeutral, lucide.Info, "Data is up to date as of 9:41."},
		{console.ConsoleColorInfo, lucide.Info, "A new version of the SDK is out."},
		{console.ConsoleColorSuccess, lucide.CircleCheck, "The alert channel works."},
		{console.ConsoleColorError, lucide.CircleX, "The upload failed. Try again."},
	}
	var banners g.Group
	for _, r := range rows {
		banners = append(banners, console.ConsoleBanner{Color: r.color, Icon: r.icon, Title: "", Dismiss: "", Class: "", Children: h.P(h.Class("cui-console-banner-body"), g.Text(r.text))})
	}
	example.Check(t, h.Div(h.Class("grid gap-3"), banners), "console/console-banner/examples/colors.html")
}

func TestBannerABrandBannerPanics(t *testing.T) {
	mustPanic(t, "has no brand color", func() {
		example.Render(console.ConsoleBanner{Color: console.ConsoleColorBrand, Icon: icon.Icon{}, Title: "", Dismiss: "", Class: "", Children: nil})
	})
}
