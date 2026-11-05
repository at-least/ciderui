package console_test

import (
	"os"
	"strings"
	"testing"

	"github.com/at-least/ciderui/go/console"
	"github.com/at-least/ciderui/go/internal/example"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

const playerPoster = "data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHdpZHRoPSIxNjAwIiBoZWlnaHQ9IjkwMCI+PHJlY3Qgd2lkdGg9IjE2MDAiIGhlaWdodD0iOTAwIiBmaWxsPSIjNGE2YjhhIi8+PGNpcmNsZSBjeD0iMTE4MCIgY3k9IjMwMCIgcj0iMTQwIiBmaWxsPSIjZThjMjdhIi8+PHJlY3QgeT0iNjIwIiB3aWR0aD0iMTYwMCIgaGVpZ2h0PSIyODAiIGZpbGw9IiMyZjRhM2EiLz48L3N2Zz4="

// a frame of stills, as muse draws them: each <img> over the last, the
// page's script showing one at a time
const playerStill = "data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHdpZHRoPSIxNjAwIiBoZWlnaHQ9IjkwMCI+PHJlY3Qgd2lkdGg9IjE2MDAiIGhlaWdodD0iOTAwIiBmaWxsPSIjNmE0YThhIi8+PGNpcmNsZSBjeD0iNDIwIiBjeT0iMzIwIiByPSIxNjAiIGZpbGw9IiNmMGQwYTAiLz48cmVjdCB5PSI2NDAiIHdpZHRoPSIxNjAwIiBoZWlnaHQ9IjI2MCIgZmlsbD0iIzNhMmE0YSIvPjwvc3ZnPg=="

// the voices and the subtitles side by side, as the Apple TV app's
func subtitlesMenu() console.ConsoleChoiceMenu {
	return console.ConsoleChoiceMenu{ID: "subtitles", Class: "", Columns: []console.ConsoleChoiceColumn{
		{Name: "voices", Title: "Voices", Choices: []console.ConsoleChoice{
			{Label: "English", Value: "en", Lang: "en", Checked: true},
			{Label: "中文", Value: "zh-Hant", Lang: "zh-Hant", Checked: false},
		}},
		{Name: "subtitles", Title: "Subtitles", Choices: []console.ConsoleChoice{
			{Label: "Off", Value: "off", Lang: "", Checked: false},
			{Label: "English", Value: "en", Lang: "en", Checked: true},
			{Label: "中文", Value: "zh-Hant", Lang: "zh-Hant", Checked: false},
		}},
	}}
}

func player(state console.ConsolePlayerState, hidden bool, title string) console.ConsolePlayer {
	return console.ConsolePlayer{
		Video:               h.Video(h.Poster(playerPoster), h.Preload("none"), g.Attr("playsinline")),
		State:               state,
		ControlsHidden:      hidden,
		Title:               title,
		Caption:             "The Wedding · Scene 2　The Vows · S02-03",
		PlayAriaLabel:       "Play",
		PauseAriaLabel:      "Pause",
		Volume:              50,
		FullScreen:          false,
		Muted:               false,
		VolumeAriaLabel:     "Volume",
		MuteAriaLabel:       "Mute",
		Position:            5,
		Duration:            62,
		ScrubAriaLabel:      "Playback position",
		Cue:                 "",
		CueLang:             "",
		SubtitlesMenu:       subtitlesMenu(),
		SubtitlesAriaLabel:  "Subtitles",
		FullScreenAriaLabel: "Full Screen",
		FullScreenControls:  nil,
		Intro:               nil,
		End:                 nil,
		Ended:               false,
		Popovers:            nil,
		Class:               "aspect-video",
	}
}

func stills() console.ConsolePlayer {
	p := player(console.ConsolePlayerStatePlaying, false, "The Wedding")
	p.Video = h.Div(h.Class("relative"), h.Img(h.Class("absolute inset-0 size-full object-contain"), h.Src(playerStill), h.Alt("")))
	p.Cue = "Everyone stood when the bride came in, and the old dog barked once."
	p.CueLang = "en"
	return p
}

// a page's own controls in the bar while full screen, as muse's: a report and the next item to look at
func fullScreen() console.ConsolePlayer {
	p := stills()
	p.FullScreen = true
	p.FullScreenControls = g.Group{
		h.Button(h.Class("cui-console-player-text-btn"), h.Type("button"), g.Text("Report a Problem")),
		h.Button(h.Class("cui-console-player-text-btn"), h.Type("button"), g.Text("Next to Watch › 3")),
	}
	return p
}

// before it first plays: the title, its length, a sentence, Play, and Start Over hidden until the page's script shows it
func intro() console.ConsolePlayer {
	p := stills()
	p.State = console.ConsolePlayerStatePaused
	p.Cue, p.CueLang = "", ""
	p.Intro = &console.ConsolePlayerIntro{Title: "The Wedding", Detail: "8 min 33 s", Summary: "The bride was late, and the old dog knew why.", PlayLabel: "Play", RestartLabel: "Start Over"}
	return p
}

// played to the end, over its last picture: the page's next step, a sentence, a button and a link, and Play Again
func end() console.ConsolePlayer {
	p := stills()
	p.Ended = true
	p.State = console.ConsolePlayerStatePaused
	p.Position = p.Duration
	p.Cue, p.CueLang = "", ""
	p.End = &console.ConsolePlayerEnd{Title: "Finished", ReplayLabel: "Play Again", Body: g.Group{
		h.P(g.Text("Once approved, the other 9 languages are recorded.")),
		h.Button(h.Class("cui-console-btn"), h.Data("variant", "filled"), h.Type("button"), g.Text("Approve Recordings")),
		h.A(h.Class("cui-console-player-text-btn"), h.Href("#next"), g.Text("Next: The Wedding Feast ›")),
	}}
	return p
}

var playerExamples = map[string]console.ConsolePlayer{
	"end":         end(),
	"intro":       intro(),
	"full-screen": fullScreen(),
	"paused":      player(console.ConsolePlayerStatePaused, false, "The Wedding"),
	"playing":     player(console.ConsolePlayerStatePlaying, false, "The Wedding"),
	"hidden":      player(console.ConsolePlayerStatePlaying, true, "The Wedding"),
	"stills":      stills(),
}

func TestPlayerRendersTheExamples(t *testing.T) {
	for id, p := range playerExamples {
		path := "console/console-player/examples/" + id + ".html"
		if os.Getenv("WRITE_PLAYER_EXAMPLES") == "1" {
			html := example.Render(p)
			for _, tag := range []string{"<video", "<div", "<input", "<button", "<p ", "<p>", "<time", "<img"} {
				html = strings.ReplaceAll(html, "><"+tag[1:], ">\n"+tag)
			}
			if err := os.WriteFile("../../components/"+path, []byte(html+"\n"), 0o644); err != nil {
				t.Fatal(err)
			}
		}
		example.Check(t, p, path)
	}
}

func TestPlayerPanics(t *testing.T) {
	for _, c := range []struct {
		want   string
		change func(*console.ConsolePlayer)
	}{
		{"needs a Video", func(p *console.ConsolePlayer) { p.Video = nil }},
		{"needs a PlayAriaLabel", func(p *console.ConsolePlayer) { p.PlayAriaLabel = "" }},
		{"needs a PauseAriaLabel", func(p *console.ConsolePlayer) { p.PauseAriaLabel = "" }},
		{"needs a VolumeAriaLabel", func(p *console.ConsolePlayer) { p.VolumeAriaLabel = "" }},
		{"needs a MuteAriaLabel", func(p *console.ConsolePlayer) { p.MuteAriaLabel = "" }},
		{"needs a ScrubAriaLabel", func(p *console.ConsolePlayer) { p.ScrubAriaLabel = "" }},
		{"needs a FullScreenAriaLabel", func(p *console.ConsolePlayer) { p.FullScreenAriaLabel = "" }},
		{"needs a SubtitlesAriaLabel with a SubtitlesMenu", func(p *console.ConsolePlayer) { p.SubtitlesAriaLabel = "" }},
		{"needs a Cue with a CueLang", func(p *console.ConsolePlayer) { p.CueLang = "en" }},
		{"Volume is 0 to 100", func(p *console.ConsolePlayer) { p.Volume = 101 }},
		{"Position is 0 to its Duration", func(p *console.ConsolePlayer) { p.Position = 63 }},
		{"needs an End to be Ended", func(p *console.ConsolePlayer) { p.Ended = true }},
		{"needs a Title and a ReplayLabel", func(p *console.ConsolePlayer) {
			p.End = &console.ConsolePlayerEnd{Title: "Finished", Body: nil, ReplayLabel: ""}
		}},
		{"needs a Title, a PlayLabel and a RestartLabel", func(p *console.ConsolePlayer) {
			p.Intro = &console.ConsolePlayerIntro{Title: "", Detail: "", Summary: "", PlayLabel: "Play", RestartLabel: "Start Over"}
		}},
	} {
		p := player(console.ConsolePlayerStatePaused, false, "")
		c.change(&p)
		mustPanic(t, c.want, func() { example.Render(p) })
	}
}

func TestPlayerWritesNoMenuButtonWithoutItsMenu(t *testing.T) {
	p := player(console.ConsolePlayerStatePaused, false, "")
	p.SubtitlesMenu, p.SubtitlesAriaLabel = console.ConsoleChoiceMenu{ID: "", Columns: nil, Class: ""}, ""
	got := example.Render(p)
	if strings.Contains(got, "popover") || strings.Contains(got, "cui-console-player-title") {
		t.Errorf("a player with no menus and no title wrote one: %s", got)
	}
}

// muted, the player says so for the mute button's glyph; it holds both glyphs
func TestPlayerWritesMuted(t *testing.T) {
	p := player(console.ConsolePlayerStatePaused, false, "")
	if got := example.Render(p); strings.Contains(got, "data-muted") || !strings.Contains(got, "cui-console-player-glyph-muted") {
		t.Errorf("a player not muted wrote data-muted, or no muted glyph: %s", got)
	}
	p.Muted = true
	if got := example.Render(p); !strings.Contains(got, `<div class="cui-console-player aspect-video" data-muted>`) {
		t.Errorf("a muted player wrote no data-muted: %s", got)
	}
}

// the page's full-screen controls in the bar, before its buttons; none, no wrapper
func TestPlayerWritesItsFullScreenControls(t *testing.T) {
	p := player(console.ConsolePlayerStatePaused, false, "")
	if got := example.Render(p); strings.Contains(got, "cui-console-player-full-screen-controls") {
		t.Errorf("a player with no FullScreenControls wrote their row: %s", got)
	}
	p.FullScreenControls = h.Button(h.Class("cui-console-player-text-btn"), h.Type("button"), g.Text("Report"))
	want := `<div class="cui-console-player-buttons"><div class="cui-console-player-full-screen-controls"><button class="cui-console-player-text-btn" type="button">Report</button></div><button class="cui-console-player-btn"`
	if got := example.Render(p); !strings.Contains(got, want) {
		t.Errorf("the controls are not first in the buttons' row: %s", got)
	}
	p.Popovers = h.Div(h.Class("cui-console-popover"), h.ID("report"), g.Attr("popover"))
	if got := example.Render(p); !strings.HasSuffix(got, `<div class="cui-console-popover" id="report" popover></div></div>`) {
		t.Errorf("the page's popover is not the player's last child: %s", got)
	}
}

// the hooks a page's script finds the controls by: each once
func TestPlayerMarksEachControl(t *testing.T) {
	got := example.Render(player(console.ConsolePlayerStatePaused, false, ""))
	for _, control := range []string{"play", "volume", "mute", "subtitles", "full-screen", "scrubber", "elapsed", "remaining", "caption", "preview"} {
		if n := strings.Count(got, `data-control="`+control+`"`); n != 1 {
			t.Errorf("data-control=%q written %d times, want 1", control, n)
		}
	}
}
