package console_test

import (
	"strings"
	"testing"

	"github.com/at-least/ciderui/go/console"
	"github.com/at-least/ciderui/go/internal/example"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

func TestBannerRendersTheExamples(t *testing.T) {
	warning := console.ConsoleBanner{Color: console.ConsoleColorWarning, Title: "Symbols are missing",
		Body:     g.Group{g.Text("Upload a dSYM in "), textLink("#settings", "Settings"), g.Text(" to see function names in this crash.")},
		Children: h.P(h.Class("cui-console-banner-actions"), h.A(h.Class("cui-console-btn"), h.Data("size", "sm"), h.Href("#upload"), g.Text("Upload dSYM"))),
	}
	example.Check(t, warning, "console/console-banner/examples/banner.html")

	rows := []struct {
		color console.ConsoleColor
		body  g.Group
	}{
		{console.ConsoleColorNeutral, g.Group{g.Text("Data is up to date as of 9:41. See the "), textLink("#sync-log", "sync log"), g.Text(".")}},
		{console.ConsoleColorInfo, g.Group{g.Text("A new version of the SDK is out. Read the "), textLink("#release-notes", "release notes"), g.Text(".")}},
		{console.ConsoleColorSuccess, g.Group{g.Text("The alert channel works. See the "), textLink("#delivery-log", "delivery log"), g.Text(".")}},
		{console.ConsoleColorError, g.Group{g.Text("The upload failed. "), textLink("#retry", "Try again"), g.Text(" or check your connection.")}},
	}
	var banners g.Group
	for _, r := range rows {
		banners = append(banners, console.ConsoleBanner{Color: r.color, Title: "", Body: r.body, Children: nil})
	}
	example.Check(t, h.Div(h.Class("grid gap-3"), banners), "console/console-banner/examples/colors.html")
}

func TestBannerWithNoBodyWritesNoBodyBlock(t *testing.T) {
	got := example.Render(console.ConsoleBanner{Color: console.ConsoleColorInfo, Title: "Synced", Body: nil, Children: nil})
	if want := `<div><p class="cui-console-banner-title">Synced</p></div></div>`; !strings.HasSuffix(got, want) {
		t.Errorf("a banner with no Body wrote a body block after its title:\n got: %s\nwant: …%s", got, want)
	}
}

func textLink(href, text string) g.Node {
	return h.A(h.Class("cui-console-link"), h.Href(href), g.Text(text))
}
