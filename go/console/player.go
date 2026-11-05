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
// data-controls="hidden" (while playing and the pointer is still),
// data-muted (the mute button's glyph), data-fullscreen (while it fills
// the screen: the bar's FullScreenControls show and the full-screen button
// shows its exit glyph), each range's value and --cui-console-player-value,
// the times, the cue, the caption. It finds each control by its
// data-control: play, volume, mute, subtitles, full-screen, scrubber,
// elapsed, remaining, caption, and the scrubber's preview
// (preview), hidden until the page fills its picture and line and shows it
// over the time under the pointer.
type ConsolePlayer struct {
	// Video is the picture, the player's first child: the page's <video>
	// (its source and tracks; no controls attribute), or an element the
	// page draws its stills in, which the player sizes as it would a video.
	Video          g.Node
	State          ConsolePlayerState
	ControlsHidden bool
	// Title: "" for none.
	Title string
	// Caption is a line at the header's start shown while paused, what is
	// on screen; "" for none (the page's script fills it).
	Caption        string
	PlayAriaLabel  string
	PauseAriaLabel string
	// Volume is 0 to 100.
	Volume int
	// FullScreen shows the bar's FullScreenControls (data-fullscreen).
	FullScreen bool
	// Muted shows the mute button's muted glyph.
	Muted           bool
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
	// SubtitlesMenu is the popover the subtitles button opens, its columns
	// such as the voices and the subtitles side by side, written
	// inside the player, so the controls stay while the focus is in it; no
	// ID, no button.
	SubtitlesMenu       ConsoleChoiceMenu
	SubtitlesAriaLabel  string
	FullScreenAriaLabel string
	// FullScreenControls are the page's own controls the bar shows before
	// its buttons while the player fills the screen (data-fullscreen), and
	// hides otherwise; nil for none.
	FullScreenControls g.Node
	// Popovers are the page's own (cui-console-popover, a field its
	// FullScreenControls open), written inside the player after its menus,
	// so they open over it in full screen and are placed as its menus are;
	// nil for none.
	Popovers g.Node
	Class    string
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
	return h.Div(class.Root("cui-console-player", p.Class), optionalAttr("data-state", p.State.value), g.If(p.ControlsHidden, g.Attr("data-controls", "hidden")), g.If(p.Muted, g.Attr("data-muted")), g.If(p.FullScreen, g.Attr("data-fullscreen")),
		p.Video,
		h.Div(h.Class("cui-console-player-overlay"),
			h.Div(h.Class("cui-console-player-header"),
				h.P(h.Class("cui-console-player-caption"), h.Data("control", "caption"), g.Text(p.Caption)),
				// the range opens beside its button while the pointer is over either
				h.Div(h.Class("cui-console-player-volume"),
					rangeInput("volume", p.Volume, 100, p.VolumeAriaLabel, "", float64(p.Volume)),
					button("mute", p.MuteAriaLabel, "", g.Group{lucide.Volume2.Class("cui-console-player-glyph-volume"), lucide.VolumeX.Class("cui-console-player-glyph-muted")}),
				),
			),
			h.Button(h.Class("cui-console-player-play"), h.Type("button"), h.Data("control", "play"), h.Aria("label", playLabel),
				lucide.Play.Class("cui-console-player-glyph-play"), lucide.Pause.Class("cui-console-player-glyph-pause")),
			h.Div(h.Class("cui-console-player-subtitles"),
				h.P(h.Span(h.Class("cui-console-player-cue"), optionalAttr("lang", p.CueLang), g.Text(p.Cue))),
			),
			h.Div(h.Class("cui-console-player-footer"),
				g.If(p.Title != "", h.P(h.Class("cui-console-player-title"), g.Text(p.Title))),
				h.Div(h.Class("cui-console-player-buttons"),
					g.If(p.FullScreenControls != nil, h.Div(h.Class("cui-console-player-full-screen-controls"), p.FullScreenControls)),
					g.If(p.SubtitlesMenu.ID != "", button("subtitles", p.SubtitlesAriaLabel, p.SubtitlesMenu.ID, lucide.Captions)),
					button("full-screen", p.FullScreenAriaLabel, "", g.Group{lucide.Maximize2.Class("cui-console-player-glyph-full-screen"), lucide.Minimize2.Class("cui-console-player-glyph-exit-full-screen")}),
				),
				h.Div(h.Class("cui-console-player-progress"),
					rangeInput("scrubber", p.Position, p.Duration, p.ScrubAriaLabel, clock(p.Position), played),
					// the range's aria-valuetext says the time: the preview is for the eye
					h.Div(h.Class("cui-console-player-preview"), h.Data("control", "preview"), h.Aria("hidden", "true"), g.Attr("hidden"),
						h.Img(h.Class("cui-console-player-preview-picture"), h.Alt("")),
						h.P(h.Class("cui-console-player-preview-label")),
					),
					h.Time(h.Class("cui-console-player-time"), h.Data("control", "elapsed"), g.Text(clock(p.Position))),
					h.Time(h.Class("cui-console-player-time"), h.Data("control", "remaining"), g.Text("-"+clock(p.Duration-p.Position))),
				),
			),
		),
		g.If(p.SubtitlesMenu.ID != "", p.SubtitlesMenu),
		p.Popovers,
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
