package console_test

import (
	"maps"
	"os"
	"slices"
	"strings"
	"testing"

	"github.com/at-least/ciderui/go/console"
	"github.com/at-least/ciderui/go/internal/example"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

const playerPoster = "data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHdpZHRoPSIxNjAwIiBoZWlnaHQ9IjkwMCI+PHJlY3Qgd2lkdGg9IjE2MDAiIGhlaWdodD0iOTAwIiBmaWxsPSIjNGE2YjhhIi8+PGNpcmNsZSBjeD0iMTE4MCIgY3k9IjMwMCIgcj0iMTQwIiBmaWxsPSIjZThjMjdhIi8+PHJlY3QgeT0iNjIwIiB3aWR0aD0iMTYwMCIgaGVpZ2h0PSIyODAiIGZpbGw9IiMyZjRhM2EiLz48L3N2Zz4="

// @external: a frame of stills, as muse draws them: each <img> over the last, the
// page's script showing one at a time
const playerStill = "data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHdpZHRoPSIxNjAwIiBoZWlnaHQ9IjkwMCI+PHJlY3Qgd2lkdGg9IjE2MDAiIGhlaWdodD0iOTAwIiBmaWxsPSIjNmE0YThhIi8+PGNpcmNsZSBjeD0iNDIwIiBjeT0iMzIwIiByPSIxNjAiIGZpbGw9IiNmMGQwYTAiLz48cmVjdCB5PSI2NDAiIHdpZHRoPSIxNjAwIiBoZWlnaHQ9IjI2MCIgZmlsbD0iIzNhMmE0YSIvPjwvc3ZnPg=="

// @external: the voices and the subtitles side by side, as the Apple TV app's
func subtitlesMenu() *console.ConsoleChoiceMenu {
	return &console.ConsoleChoiceMenu{ID: "subtitles", Columns: []console.ConsoleChoiceColumn{
		{Column: "voices", Title: "Voices", Options: []console.ConsoleChoiceOption{
			{Label: "English", Value: "en", Lang: "en", Checked: true},
			{Label: "中文", Value: "zh-Hant", Lang: "zh-Hant", Checked: false},
		}},
		{Column: "subtitles", Title: "Subtitles", Options: []console.ConsoleChoiceOption{
			{Label: "Off", Value: "off", Lang: "", Checked: false},
			{Label: "English", Value: "en", Lang: "en", Checked: true},
			{Label: "中文", Value: "zh-Hant", Lang: "zh-Hant", Checked: false},
		}},
	}}
}

func pausedPlayer() console.ConsolePlayer {
	return console.ConsolePlayer{
		Video:               h.Video(h.Poster(playerPoster), h.Preload("none"), g.Attr("playsinline")),
		PlayAriaLabel:       "Play",
		VolumeAriaLabel:     "Volume",
		MuteAriaLabel:       "Mute",
		ScrubberAriaLabel:   "Playback position",
		SubtitlesMenu:       subtitlesMenu(),
		SubtitlesAriaLabel:  "Subtitles",
		FullScreenAriaLabel: "Full Screen",
		FullScreenControls:  nil,
		Intro:               nil,
		End:                 nil,
		Popovers:            nil,
	}
}

func stills() console.ConsolePlayer {
	p := pausedPlayer()
	p.Video = h.Div(h.Class("relative"), h.Img(h.Class("absolute inset-0 size-full object-contain"), h.Src(playerStill), h.Alt("")))
	return p
}

func intro() console.ConsolePlayer {
	p := stills()
	p.Intro = &console.ConsolePlayerIntro{Title: "The Wedding", Detail: "8 min 33 s", Body: "The bride was late, and the old dog knew why.", StartLabel: "Play", RestartLabel: "Start Over"}
	return p
}

func end() console.ConsolePlayer {
	p := stills()
	p.End = &console.ConsolePlayerEnd{Title: "Finished", ReplayLabel: "Play Again"}
	return p
}

// @invariant: the player at rest, as ConsolePlayer writes it
var playerExamples = map[string]console.ConsolePlayer{
	"intro":  intro(),
	"paused": pausedPlayer(),
}

type scriptSets struct{ old, new string }

const (
	playerRoot      = `<div class="cui-console-player"`
	headingAtRest   = `data-control="heading"></p>`
	headingSet      = `data-control="heading">The Wedding · Scene 2　The Vows · S02-03</p>`
	endBodyAtRest   = `<div class="cui-console-player-end-body"></div>`
	endBodySet      = `<div class="cui-console-player-end-body"><p>Once approved, the other 9 languages are recorded.</p><button class="cui-console-btn" data-variant="filled" type="button">Approve Recordings</button><a class="cui-console-player-text-btn" href="#next">Next: The Wedding Feast ›</a></div>`
	volumeAtRest    = `value="100" aria-label="Volume" style="--cui-console-player-value: 100%"`
	volumeSet       = `value="50" aria-label="Volume" style="--cui-console-player-value: 50%"`
	playAtRest      = `data-control="play" aria-label="Play"`
	playSet         = `data-control="play" aria-label="Pause"`
	cueAtRest       = `<span class="cui-console-player-cue"></span>`
	cueSet          = `<span class="cui-console-player-cue" lang="en">Everyone stood when the bride came in, and the old dog barked once.</span>`
	scrubAtRest     = `max="0" value="0" aria-label="Playback position" aria-valuetext="0:00" style="--cui-console-player-value: 0%"`
	elapsedAtRest   = `data-control="elapsed">0:00</time>`
	remainingAtRest = `data-control="remaining">-0:00</time>`
)

var playedFiveSeconds = []scriptSets{
	{headingAtRest, headingSet},
	{volumeAtRest, volumeSet},
	{playAtRest, playSet},
	{scrubAtRest, `max="62" value="5" aria-label="Playback position" aria-valuetext="0:05" style="--cui-console-player-value: 8.064516129032258%"`},
	{elapsedAtRest, `data-control="elapsed">0:05</time>`},
	{remainingAtRest, `data-control="remaining">-0:57</time>`},
}

func fullScreen() console.ConsolePlayer {
	p := stills()
	p.FullScreenControls = g.Group{
		h.Button(h.Class("cui-console-player-text-btn"), h.Type("button"), g.Text("Report a Problem")),
		h.Button(h.Class("cui-console-player-text-btn"), h.Type("button"), g.Text("Next to Watch › 3")),
	}
	return p
}

// @invariant: the player as the page's script has set it (playing, the controls
// hidden, full screen, a cue, ended): ConsolePlayer's output at rest with the
// script's changes, each found exactly once, so the browser suite measures
// markup a page gets
var playerStateExamples = map[string]struct {
	rest console.ConsolePlayer
	sets []scriptSets
}{
	"playing":     {pausedPlayer(), append([]scriptSets{{playerRoot + ">", playerRoot + ` data-state="playing">`}}, playedFiveSeconds...)},
	"hidden":      {pausedPlayer(), append([]scriptSets{{playerRoot + ">", playerRoot + ` data-state="playing" data-controls="hidden">`}}, playedFiveSeconds...)},
	"stills":      {stills(), append([]scriptSets{{playerRoot + ">", playerRoot + ` data-state="playing">`}, {cueAtRest, cueSet}}, playedFiveSeconds...)},
	"full-screen": {fullScreen(), append([]scriptSets{{playerRoot + ">", playerRoot + ` data-state="playing" data-full-screen>`}, {cueAtRest, cueSet}}, playedFiveSeconds...)},
	"end": {end(), []scriptSets{
		{playerRoot + ">", playerRoot + ` data-end>`},
		{headingAtRest, headingSet},
		{endBodyAtRest, endBodySet},
		{volumeAtRest, volumeSet},
		{scrubAtRest, `max="62" value="62" aria-label="Playback position" aria-valuetext="1:02" style="--cui-console-player-value: 100%"`},
		{elapsedAtRest, `data-control="elapsed">1:02</time>`},
	}},
}

func TestPlayerStateExamplesAreItsOutputAsTheScriptSetsIt(t *testing.T) {
	for id, c := range playerStateExamples {
		html := example.Render(c.rest)
		for _, set := range c.sets {
			if n := strings.Count(html, set.old); n != 1 {
				t.Fatalf("%s: the script's change from %q found %d times in the player at rest, want 1", id, set.old, n)
			}
			html = strings.Replace(html, set.old, set.new, 1)
		}
		example.Check(t, g.Raw(html), "console/console-player/examples/"+id+".html")
	}
}

// each example is one ConsolePlayer writes or one the page's script makes, so
// a new one is held to the type or named as the script's
func TestPlayerExamplesAreRenderedOrState(t *testing.T) {
	entries, err := os.ReadDir("../../components/console/console-player/examples")
	if err != nil {
		t.Fatal(err)
	}
	var got []string
	for _, e := range entries {
		got = append(got, strings.TrimSuffix(e.Name(), ".html"))
	}
	want := slices.Collect(maps.Keys(playerExamples))
	want = slices.AppendSeq(want, maps.Keys(playerStateExamples))
	slices.Sort(want)
	if !slices.Equal(got, want) {
		t.Errorf("examples %q, want the rendered and the state ones %q", got, want)
	}
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

func TestPlayerWritesNoMenuButtonWithoutItsMenu(t *testing.T) {
	p := pausedPlayer()
	p.SubtitlesMenu, p.SubtitlesAriaLabel = nil, ""
	got := example.Render(p)
	if strings.Contains(got, "popover") || strings.Contains(got, `data-control="subtitles"`) {
		t.Errorf("a player with no menu wrote one: %s", got)
	}
}

func TestPlayerWritesItsFullScreenControls(t *testing.T) {
	p := pausedPlayer()
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

func TestPlayerMarksEachControl(t *testing.T) {
	for name, player := range map[string]console.ConsolePlayer{"paused": pausedPlayer(), "intro": intro(), "end": end()} {
		got := example.Render(player)
		controls := map[string][]string{
			"paused": {"play", "volume", "mute", "subtitles", "full-screen", "scrubber", "elapsed", "remaining", "heading", "preview"},
			"intro":  {"start", "restart"},
			"end":    {"replay"},
		}[name]
		for _, control := range controls {
			if n := strings.Count(got, `data-control="`+control+`"`); n != 1 {
				t.Errorf("%s: data-control=%q written %d times, want 1", name, control, n)
			}
		}
	}
}

func TestPlayerPanics(t *testing.T) {
	example.Panics(t, pausedPlayer, []example.Break[console.ConsolePlayer]{
		{Want: "console.ConsolePlayer needs a Video", Change: func(p *console.ConsolePlayer) { p.Video = nil }},
		{Want: "console.ConsolePlayer needs a PlayAriaLabel", Change: func(p *console.ConsolePlayer) { p.PlayAriaLabel = "" }},
		{Want: "console.ConsolePlayer needs a VolumeAriaLabel", Change: func(p *console.ConsolePlayer) { p.VolumeAriaLabel = "" }},
		{Want: "console.ConsolePlayer needs a MuteAriaLabel", Change: func(p *console.ConsolePlayer) { p.MuteAriaLabel = "" }},
		{Want: "console.ConsolePlayer needs a ScrubberAriaLabel", Change: func(p *console.ConsolePlayer) { p.ScrubberAriaLabel = "" }},
		{Want: "console.ConsolePlayer needs a FullScreenAriaLabel", Change: func(p *console.ConsolePlayer) { p.FullScreenAriaLabel = "" }},
		{Want: "console.ConsolePlayer needs a SubtitlesAriaLabel with a SubtitlesMenu", Change: func(p *console.ConsolePlayer) { p.SubtitlesAriaLabel = "" }},
	})
}

func TestPlayerIntroPanics(t *testing.T) {
	const want = "console.ConsolePlayerIntro needs a Title, a StartLabel and a RestartLabel"
	example.Panics(t, intro, []example.Break[console.ConsolePlayer]{
		{Want: want, Change: func(p *console.ConsolePlayer) { p.Intro.Title = "" }},
		{Want: want, Change: func(p *console.ConsolePlayer) { p.Intro.StartLabel = "" }},
		{Want: want, Change: func(p *console.ConsolePlayer) { p.Intro.RestartLabel = "" }},
	})
}

func TestPlayerEndPanics(t *testing.T) {
	const want = "console.ConsolePlayerEnd needs a Title and a ReplayLabel"
	example.Panics(t, end, []example.Break[console.ConsolePlayer]{
		{Want: want, Change: func(p *console.ConsolePlayer) { p.End.Title = "" }},
		{Want: want, Change: func(p *console.ConsolePlayer) { p.End.ReplayLabel = "" }},
	})
}

func TestPlayerIntroWritesNoDetailOrBodyItHasNone(t *testing.T) {
	p := intro()
	p.Intro.Detail, p.Intro.Body = "", ""
	want := `<p class="cui-console-player-intro-title">The Wedding</p><div class="cui-console-player-intro-actions">`
	if got := example.Render(p); !strings.Contains(got, want) {
		t.Errorf("the intro with no Detail or Body wrote a line for one, not its title then its actions: %s", got)
	}
}
