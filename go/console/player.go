package console

import (
	"fmt"
	"io"
	"strconv"

	"github.com/at-least/ciderui/go/internal/class"
	"github.com/at-least/ciderui/go/lucide"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// ConsolePlayerState is what the video is doing; the page's script sets
// data-state and the centre button's aria-label as it changes.
type ConsolePlayerState struct{ value string }

var (
	ConsolePlayerStatePaused  = ConsolePlayerState{""}
	ConsolePlayerStatePlaying = ConsolePlayerState{"playing"}
)

// ConsolePlayer is a video's controls over it, tv.apple.com's player's. The
// page's script plays the video and keeps the markup current: data-state,
// data-controls="hidden" (while playing and the pointer is still), each
// range's value and --cui-console-player-value, the times, the cue. It finds
// each control by its data-control: play, volume, mute, subtitles, voices,
// full-screen, scrubber, elapsed, remaining.
type ConsolePlayer struct {
	// Video is the picture, the player's first child: the page's <video>
	// (its source and tracks; no controls attribute), or an element the
	// page draws its stills in, which the player sizes as it would a video.
	Video          g.Node
	State          ConsolePlayerState
	ControlsHidden bool
	// Title: "" for none.
	Title          string
	PlayAriaLabel  string
	PauseAriaLabel string
	// Volume is 0 to 100.
	Volume          int
	VolumeAriaLabel string
	MuteAriaLabel   string
	// Position and Duration are in seconds.
	Position       int
	Duration       int
	ScrubAriaLabel string
	// Cue is the subtitle shown, written over the picture; "" for none.
	Cue string
	// CueLang is the cue's language, a BCP 47 tag; "" for the page's.
	CueLang string
	// SubtitlesMenu is the popover the subtitles button opens, written
	// inside the player, so the controls stay while the focus is in it; no
	// ID, no button.
	SubtitlesMenu      ConsoleChoiceMenu
	SubtitlesAriaLabel string
	// VoicesMenu: the same for the voices.
	VoicesMenu          ConsoleChoiceMenu
	VoicesAriaLabel     string
	FullScreenAriaLabel string
	Class               string
}

func (p ConsolePlayer) Render(w io.Writer) error {
	if p.Video == nil {
		panic("console.ConsolePlayer needs a Video")
	}
	for name, label := range map[string]string{"PlayAriaLabel": p.PlayAriaLabel, "PauseAriaLabel": p.PauseAriaLabel, "VolumeAriaLabel": p.VolumeAriaLabel, "MuteAriaLabel": p.MuteAriaLabel, "ScrubAriaLabel": p.ScrubAriaLabel, "FullScreenAriaLabel": p.FullScreenAriaLabel} {
		if label == "" {
			panic("console.ConsolePlayer needs a " + name)
		}
	}
	if p.SubtitlesMenu.ID != "" && p.SubtitlesAriaLabel == "" {
		panic("console.ConsolePlayer needs a SubtitlesAriaLabel with a SubtitlesMenu")
	}
	if p.VoicesMenu.ID != "" && p.VoicesAriaLabel == "" {
		panic("console.ConsolePlayer needs a VoicesAriaLabel with a VoicesMenu")
	}
	if p.CueLang != "" && p.Cue == "" {
		panic("console.ConsolePlayer needs a Cue with a CueLang")
	}
	if p.Volume < 0 || p.Volume > 100 {
		panic("console.ConsolePlayer's Volume is 0 to 100")
	}
	if p.Duration < 0 || p.Position < 0 || p.Position > p.Duration {
		panic("console.ConsolePlayer's Position is 0 to its Duration")
	}
	playLabel := p.PlayAriaLabel
	if p.State == ConsolePlayerStatePlaying {
		playLabel = p.PauseAriaLabel
	}
	played := 0.0
	if p.Duration > 0 {
		played = float64(p.Position) * 100 / float64(p.Duration)
	}
	button := func(control, label, menu string, glyph g.Node) g.Node {
		return h.Button(h.Class("cui-console-player-btn"), h.Type("button"), h.Data("control", control), h.Aria("label", label), optionalAttr("popovertarget", menu), glyph)
	}
	return h.Div(class.Root("cui-console-player", p.Class), optionalAttr("data-state", p.State.value), g.If(p.ControlsHidden, g.Attr("data-controls", "hidden")),
		p.Video,
		h.Div(h.Class("cui-console-player-overlay"),
			h.Div(h.Class("cui-console-player-header"),
				rangeInput("volume", p.Volume, 100, p.VolumeAriaLabel, "", float64(p.Volume)),
				button("mute", p.MuteAriaLabel, "", lucide.Volume2),
			),
			h.Button(h.Class("cui-console-player-play"), h.Type("button"), h.Data("control", "play"), h.Aria("label", playLabel),
				lucide.Play.Class("cui-console-player-glyph-play"), lucide.Pause.Class("cui-console-player-glyph-pause")),
			h.Div(h.Class("cui-console-player-subtitles"),
				h.P(h.Span(h.Class("cui-console-player-cue"), optionalAttr("lang", p.CueLang), g.Text(p.Cue))),
			),
			h.Div(h.Class("cui-console-player-footer"),
				g.If(p.Title != "", h.P(h.Class("cui-console-player-title"), g.Text(p.Title))),
				h.Div(h.Class("cui-console-player-buttons"),
					g.If(p.SubtitlesMenu.ID != "", button("subtitles", p.SubtitlesAriaLabel, p.SubtitlesMenu.ID, lucide.Captions)),
					g.If(p.VoicesMenu.ID != "", button("voices", p.VoicesAriaLabel, p.VoicesMenu.ID, lucide.AudioLines)),
					button("full-screen", p.FullScreenAriaLabel, "", lucide.Maximize2),
				),
				h.Div(h.Class("cui-console-player-progress"),
					rangeInput("scrubber", p.Position, p.Duration, p.ScrubAriaLabel, clock(p.Position), played),
					h.Time(h.Class("cui-console-player-time"), h.Data("control", "elapsed"), g.Text(clock(p.Position))),
					h.Time(h.Class("cui-console-player-time"), h.Data("control", "remaining"), g.Text("-"+clock(p.Duration-p.Position))),
				),
			),
		),
		g.If(p.SubtitlesMenu.ID != "", p.SubtitlesMenu),
		g.If(p.VoicesMenu.ID != "", p.VoicesMenu),
	).Render(w)
}

func rangeInput(control string, value, max int, label, valueText string, percent float64) g.Node {
	return h.Input(h.Class("cui-console-player-range"), h.Type("range"), h.Data("control", control), h.Min("0"), h.Max(strconv.Itoa(max)), h.Value(strconv.Itoa(value)),
		h.Aria("label", label), optionalAttr("aria-valuetext", valueText),
		h.Style("--cui-console-player-value: "+strconv.FormatFloat(percent, 'f', -1, 64)+"%"))
}

// clock is m:ss.
func clock(seconds int) string {
	return fmt.Sprintf("%d:%02d", seconds/60, seconds%60)
}
