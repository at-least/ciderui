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
	example.PagesClass(t, []example.ClassRow{
		{Root: "cui-console-btn", Node: console.ConsoleBtn{Href: "", Label: "x", AriaLabel: "", Icon: lucide.Icon{}, Submit: false, Variant: console.ConsoleBtnVariantGray, Size: console.ConsoleBtnSizeMedium, Destructive: false, Attrs: nil, Class: page}},
	})
}

func TestEachPartWritesThePagesWords(t *testing.T) {
	intro := intro()
	intro.Intro = &console.ConsolePlayerIntro{Title: "The Feast", Detail: "2 min", Body: "A short one.", StartLabel: "Watch", RestartLabel: "From the Top"}
	end := end()
	end.End = &console.ConsolePlayerEnd{Title: "All Done", ReplayLabel: "Watch Again"}
	tip := validStrip()
	tip.Groups[0].Items[0].Tip = "Hello there."
	menu := subtitlesMenu()
	menu.ID = "langs"
	player := pausedPlayer()
	player.PlayAriaLabel, player.VolumeAriaLabel, player.MuteAriaLabel, player.ScrubberAriaLabel = "Go", "Loudness", "Silence", "Position"
	player.SubtitlesMenu, player.SubtitlesAriaLabel, player.FullScreenAriaLabel = menu, "Captions", "Fill Screen"
	sidebar := console.ConsoleSidebar{AriaLabel: "Main", Home: "/home", Wordmark: "Muse", SkipLabel: "Jump to Main", OpenAriaLabel: "Show Menu", CloseAriaLabel: "Hide", Footer: nil, Sections: nil}
	for _, c := range []struct {
		part  string
		node  g.Node
		words []string
	}{
		{"ConsoleArtwork", console.ConsoleArtwork{Src: "/cover.png", Alt: "The cover", Placeholder: ""}, []string{`src="/cover.png" alt="The cover"`}},
		{"ConsoleArtwork", console.ConsoleArtwork{Src: "", Alt: "", Placeholder: "No Picture Yet"}, []string{">No Picture Yet<"}},
		{"ConsoleBanner", console.ConsoleBanner{Color: console.ConsoleColorInfo, Title: "Synced", Body: nil, Actions: nil}, []string{">Synced<"}},
		{"ConsoleChart", console.ConsoleChart{AriaLabel: "Events: 9", Scale: &console.ConsoleChartScale{Peak: "peak 9", Half: "4.5"}, Empty: "Nothing yet", Axis: &console.ConsoleChartAxis{Start: "Nov 1", End: "Nov 7"}, Children: nil},
			[]string{`aria-label="Events: 9"`, ">peak 9<", ">4.5<", ">Nothing yet<", ">Nov 1<", ">Nov 7<"}},
		{"ConsoleField", console.ConsoleField{ID: "name", Label: "Name", Error: "Too short.", Control: input()}, []string{">Name<", ">Too short.<"}},
		{"ConsoleChoiceMenu", *menu, []string{`id="langs" popover`, `id="langs-voices-title"`, `aria-labelledby="langs-voices-title"`, `popovertarget="langs"`}},
		{"ConsolePlayer", player, []string{`data-control="play" aria-label="Go"`, `data-control="volume" min="0" max="100" value="100" aria-label="Loudness"`, `data-control="mute" aria-label="Silence"`,
			`data-control="scrubber" min="0" max="0" value="0" aria-label="Position"`, `data-control="subtitles" aria-label="Captions" popovertarget="langs"`, `data-control="full-screen" aria-label="Fill Screen"`}},
		{"ConsolePlayButton", console.ConsolePlayButton{Attrs: nil, AriaLabel: "Listen", Label: "", PlayingLabel: "", Size: console.ConsoleBtnSizeMedium}, []string{`aria-label="Listen"`}},
		{"ConsolePlayButton", console.ConsolePlayButton{Attrs: nil, AriaLabel: "", Label: "Hear It", PlayingLabel: "Halt", Size: console.ConsoleBtnSizeMedium}, []string{">Hear It<", ">Halt<"}},
		{"ConsoleContent", console.ConsoleContent{Sidebar: &sidebar, Back: &console.ConsoleContentBack{Href: "/up", Label: "Up"}, Title: "Takes", Detail: "Every take.", Controls: nil, Children: nil},
			[]string{`aria-label="Show Menu"`, `href="/up"`, ">Up<", ">Takes<", ">Every take.<"}},
		{"ConsolePlayerIntro", intro, []string{">The Feast<", ">2 min<", ">A short one.<", ">Watch<", ">From the Top<"}},
		{"ConsolePlayerEnd", end, []string{">All Done<", ">Watch Again<"}},
		{"ConsoleSidebar", sidebar, []string{`aria-label="Main"`, `href="/home"`, ">Muse<", ">Jump to Main<", `aria-label="Hide"`}},
		{"ConsoleStatus", console.ConsoleStatus{Label: "Down", Detail: "500 · 2 s"}, []string{">Down<", ">500 · 2 s<"}},
		{"ConsoleStrip", tip, []string{">Hello there.<"}},
		{"ConsoleTabsChoice", console.ConsoleTabsChoice{AriaLabel: "Show", Name: "show", Options: tabsChoice().Options}, []string{`role="group" aria-label="Show"`, `name="show"`}},
		{"ConsoleViewer", console.ConsoleViewer{ID: "art", AriaLabel: "Artwork", CloseAriaLabel: "Shut"}, []string{`id="art"`, `aria-label="Artwork"`, `aria-label="Shut"`}},
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
		{"ConsoleEmpty", console.ConsoleEmpty{Size: console.ConsoleEmptySizeSmall, Title: "x", Body: "", Attrs: hook}},
		{"ConsolePlayButton", console.ConsolePlayButton{AriaLabel: "x", Label: "", PlayingLabel: "", Size: console.ConsoleBtnSizeMedium, Attrs: hook}},
	} {
		got := example.Render(c.node)
		root, _, _ := strings.Cut(got, ">")
		if !strings.Contains(root, ` data-hook="on"`) {
			t.Errorf("%s does not write the page's Attrs on its root: %s", c.part, got)
		}
	}
}
