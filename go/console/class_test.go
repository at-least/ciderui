package console_test

import (
	"strings"
	"testing"

	"github.com/at-least/ciderui/go/console"
	"github.com/at-least/ciderui/go/internal/example"
	"github.com/at-least/ciderui/go/lucide"
	g "maragu.dev/gomponents"
)

func TestEachPartWritesThePagesClass(t *testing.T) {
	const page = example.PageClass
	example.PagesClass(t, nil, []example.ClassRow{
		{Root: "cui-console-btn", Node: console.ConsoleBtn{Href: "", Label: "x", AriaLabel: "", Icon: lucide.Icon{}, Submit: false, Variant: console.ConsoleBtnVariantGray, Size: console.ConsoleBtnSizeMedium, Destructive: false, Attrs: nil, Class: page}},
		{Root: "cui-console-btn cui-console-btn-icon cui-console-play", Node: console.ConsolePlay{Attrs: nil, AriaLabel: "x", Label: "", PlayingLabel: "", Size: console.ConsoleBtnSizeMedium, Class: page}},
	})
}

func TestEachPartWritesThePagesWords(t *testing.T) {
	intro := intro()
	intro.Intro = &console.ConsolePlayerIntro{Title: "The Feast", Detail: "2 min", Summary: "A short one.", StartLabel: "Watch", RestartLabel: "From the Top"}
	end := end()
	end.End = &console.ConsolePlayerEnd{Title: "All Done", ReplayLabel: "Watch Again"}
	tip := validStrip()
	tip.Groups[0].Items[0].Tip = "Hello there."
	sidebar := console.ConsoleSidebar{AriaLabel: "Main", Home: "/home", Wordmark: "Muse", SkipLabel: "Jump to Main", OpenAriaLabel: "Show Menu", CloseAriaLabel: "Hide", Footer: nil, Children: nil}
	for _, c := range []struct {
		part  string
		node  g.Node
		words []string
	}{
		{"ConsoleBanner", console.ConsoleBanner{Color: console.ConsoleColorInfo, Title: "Synced", Body: nil, Children: nil}, []string{">Synced<"}},
		{"ConsoleChart", console.ConsoleChart{AriaLabel: "Events: 9", Scale: &console.ConsoleChartScale{Peak: "peak 9", Half: "4.5"}, Empty: "Nothing yet", Axis: &console.ConsoleChartAxis{Start: "Nov 1", End: "Nov 7"}, Children: nil},
			[]string{`aria-label="Events: 9"`, ">peak 9<", ">4.5<", ">Nothing yet<", ">Nov 1<", ">Nov 7<"}},
		{"ConsoleField", console.ConsoleField{ID: "name", Label: "Name", Error: "Too short.", Control: input()}, []string{">Name<", ">Too short.<"}},
		{"ConsolePlay", console.ConsolePlay{Attrs: nil, AriaLabel: "Listen", Label: "", PlayingLabel: "", Size: console.ConsoleBtnSizeMedium, Class: ""}, []string{`aria-label="Listen"`}},
		{"ConsolePlay", console.ConsolePlay{Attrs: nil, AriaLabel: "", Label: "Hear It", PlayingLabel: "Halt", Size: console.ConsoleBtnSizeMedium, Class: ""}, []string{">Hear It<", ">Halt<"}},
		{"ConsoleContent", console.ConsoleContent{Sidebar: &sidebar, Back: console.ConsoleContentBack{Href: "/up", Label: "Up"}, Title: "Takes", Detail: "Every take.", Controls: nil, Children: nil},
			[]string{`aria-label="Show Menu"`, `href="/up"`, ">Up<", ">Takes<", ">Every take.<"}},
		{"ConsolePlayerIntro", intro, []string{">The Feast<", ">2 min<", ">A short one.<", ">Watch<", ">From the Top<"}},
		{"ConsolePlayerEnd", end, []string{">All Done<", ">Watch Again<"}},
		{"ConsoleSidebar", sidebar, []string{`aria-label="Main"`, `href="/home"`, ">Muse<", ">Jump to Main<", `aria-label="Hide"`}},
		{"ConsoleStatus", console.ConsoleStatus{Label: "Down", Color: console.ConsoleColorError, Detail: "500 · 2 s"}, []string{">Down<", ">500 · 2 s<"}},
		{"ConsoleStrip", tip, []string{">Hello there.<"}},
		{"ConsoleViewer", console.ConsoleViewer{ID: "art", AriaLabel: "Artwork", CloseAriaLabel: "Shut"}, []string{`aria-label="Artwork"`, `aria-label="Shut"`}},
	} {
		got := example.Render(c.node)
		for _, word := range c.words {
			if !strings.Contains(got, word) {
				t.Errorf("%s does not write the page's %s: %s", c.part, word, got)
			}
		}
	}
}

func TestEachPartWritesThePagesAttrs(t *testing.T) {
	hook := g.Attr("data-hook", "on")
	for _, c := range []struct {
		part string
		node g.Node
	}{
		{"ConsoleDisclosure", console.ConsoleDisclosure{Summary: "x", Attrs: hook, Children: nil}},
		{"ConsoleEmpty", console.ConsoleEmpty{Size: console.ConsoleEmptySizeSmall, Icon: lucide.Icon{}, Title: "x", Body: "", Attrs: hook}},
		{"ConsolePlay", console.ConsolePlay{AriaLabel: "x", Label: "", PlayingLabel: "", Size: console.ConsoleBtnSizeMedium, Attrs: hook, Class: ""}},
	} {
		got := example.Render(c.node)
		root, _, _ := strings.Cut(got, ">")
		if !strings.Contains(root, ` data-hook="on"`) {
			t.Errorf("%s does not write the page's Attrs on its root: %s", c.part, got)
		}
	}
}
