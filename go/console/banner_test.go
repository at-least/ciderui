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
		Body:    g.Group{g.Text("Upload a dSYM in "), textLink("#settings", "Settings"), g.Text(" to see function names in this crash.")},
		Actions: h.A(h.Class("cui-console-btn"), h.Data("size", "sm"), h.Href("#upload"), g.Text("Upload dSYM")),
	}
	example.Check(t, warning, "console/console-banner/examples/banner.html")

	rows := []struct {
		color console.ConsoleColor
		body  g.Group
	}{
		{console.ConsoleColorInfo, g.Group{g.Text("A new version of the SDK is out. Read the "), textLink("#release-notes", "release notes"), g.Text(".")}},
		{console.ConsoleColorSuccess, g.Group{g.Text("The alert channel works. See the "), textLink("#delivery-log", "delivery log"), g.Text(".")}},
		{console.ConsoleColorError, g.Group{g.Text("The upload failed. "), textLink("#retry", "Try again"), g.Text(" or check your connection.")}},
	}
	var banners g.Group
	for _, r := range rows {
		banners = append(banners, console.ConsoleBanner{Color: r.color, Title: "", Body: r.body, Actions: nil})
	}
	example.Check(t, h.Div(h.Class("grid gap-3"), banners), "console/console-banner/examples/colors.html")
}

func TestBannerWithNoBodyOrActionsWritesNoBlockForThem(t *testing.T) {
	got := example.Render(console.ConsoleBanner{Color: console.ConsoleColorInfo, Title: "Synced", Body: nil, Actions: nil})
	if want := `<div><p class="cui-console-banner-title">Synced</p></div></div>`; !strings.HasSuffix(got, want) {
		t.Errorf("a banner with no Body or Actions wrote a block for one after its title:\n got: %s\nwant: …%s", got, want)
	}
}

func TestBannerOfActionsAloneWritesThem(t *testing.T) {
	got := example.Render(console.ConsoleBanner{Color: console.ConsoleColorInfo, Title: "", Body: nil, Actions: g.Text("Try Again")})
	if want := `<div><p class="cui-console-banner-actions">Try Again</p></div></div>`; !strings.HasSuffix(got, want) {
		t.Errorf("a banner of Actions alone is not its actions after the icon:\n got: %s\nwant: …%s", got, want)
	}
}

func TestBannerPanics(t *testing.T) {
	example.Panics(t, func() console.ConsoleBanner {
		return console.ConsoleBanner{Color: console.ConsoleColorInfo, Title: "x", Body: nil, Actions: nil}
	}, []example.Break[console.ConsoleBanner]{
		{Want: "console.ConsoleBanner needs a Color: a banner's color is info, success, warning or error", Change: func(p *console.ConsoleBanner) { p.Color = console.ConsoleColorNeutral }},
		{Want: "console.ConsoleBanner needs a Title, a Body or Actions", Change: func(p *console.ConsoleBanner) { p.Title = "" }},
	})
}

func textLink(href, text string) g.Node {
	return h.A(h.Class("cui-console-link"), h.Href(href), g.Text(text))
}
