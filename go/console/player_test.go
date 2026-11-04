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

func player(state console.ConsolePlayerState, hidden bool, title string) console.ConsolePlayer {
	return console.ConsolePlayer{
		Video:               h.Video(h.Poster(playerPoster), h.Preload("none"), g.Attr("playsinline")),
		State:               state,
		ControlsHidden:      hidden,
		Title:               title,
		PlayAriaLabel:       "Play",
		PauseAriaLabel:      "Pause",
		Volume:              50,
		VolumeAriaLabel:     "Volume",
		MuteAriaLabel:       "Mute",
		Position:            5,
		Duration:            62,
		ScrubAriaLabel:      "Playback position",
		SubtitlesMenuID:     "subtitles",
		SubtitlesAriaLabel:  "Subtitles",
		AudioMenuID:         "voices",
		AudioAriaLabel:      "Voices",
		FullScreenAriaLabel: "Full Screen",
		Class:               "aspect-video",
	}
}

var playerExamples = map[string]console.ConsolePlayer{
	"paused":  player(console.ConsolePlayerStatePaused, false, "The Wedding"),
	"playing": player(console.ConsolePlayerStatePlaying, false, "The Wedding"),
	"hidden":  player(console.ConsolePlayerStatePlaying, true, "The Wedding"),
}

func TestPlayerRendersTheExamples(t *testing.T) {
	for id, p := range playerExamples {
		path := "console/console-player/examples/" + id + ".html"
		if os.Getenv("WRITE_PLAYER_EXAMPLES") == "1" {
			html := example.Render(p)
			for _, tag := range []string{"<video", "<div", "<input", "<button", "<p ", "<time"} {
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
		{"needs a SubtitlesAriaLabel with a SubtitlesMenuID", func(p *console.ConsolePlayer) { p.SubtitlesAriaLabel = "" }},
		{"needs an AudioAriaLabel with an AudioMenuID", func(p *console.ConsolePlayer) { p.AudioAriaLabel = "" }},
		{"Volume is 0 to 100", func(p *console.ConsolePlayer) { p.Volume = 101 }},
		{"Position is 0 to its Duration", func(p *console.ConsolePlayer) { p.Position = 63 }},
	} {
		p := player(console.ConsolePlayerStatePaused, false, "")
		c.change(&p)
		mustPanic(t, c.want, func() { example.Render(p) })
	}
}

func TestPlayerWritesNoMenuButtonWithoutItsMenu(t *testing.T) {
	p := player(console.ConsolePlayerStatePaused, false, "")
	p.SubtitlesMenuID, p.SubtitlesAriaLabel, p.AudioMenuID, p.AudioAriaLabel = "", "", "", ""
	got := example.Render(p)
	if strings.Contains(got, "popovertarget") || strings.Contains(got, "cui-console-player-title") {
		t.Errorf("a player with no menus and no title wrote one: %s", got)
	}
}
