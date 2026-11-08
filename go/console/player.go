package console

import (
	"io"

	"github.com/at-least/ciderui/go/internal/class"
	"github.com/at-least/ciderui/go/lucide"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// ConsolePlayerIntro is what the player shows over its picture before it
// first plays, in place of its controls: the title, a line of detail and a
// summary, a button that plays and a smaller one that starts over. The page's
// script labels the buttons (a saved position: "Resume from 3:12", and the
// smaller one shown) and drops the player's data-intro as it plays, the
// layer fading out.
type ConsolePlayerIntro struct {
	Title string
	// Detail: the length ("8 min 33 s"); "" none.
	Detail string
	// Summary: a sentence of what it is about; "" none.
	Summary   string
	PlayLabel string
	// RestartLabel is the smaller button's words, hidden until the page's
	// script shows it.
	RestartLabel string
}

// ConsolePlayerEnd is what the player shows over its last picture once it
// has played to the end, in place of its controls: a title, the page's own
// next step under it, and a smaller button that plays it again. The page's
// script sets the player's data-end as the video ends and drops it as it
// plays again, the layer fading in and out.
type ConsolePlayerEnd struct {
	Title string
	// Body is the page's: its words, buttons and links, one under another;
	// nil for none.
	Body        g.Node
	ReplayLabel string
}

// ConsolePlayer is a video's controls over it, tv.apple.com's player's,
// written at rest: paused, the controls shown, not muted, the volume full,
// nothing played, no cue or caption, not full screen, not ended. The page's
// script plays the video and keeps the markup current: data-state="playing"
// (and the centre button's aria-label), data-controls="hidden" (while playing
// and the pointer is still), data-muted (the mute button's glyph),
// data-fullscreen (while it fills the screen: the bar's FullScreenControls
// show and the full-screen button shows its exit glyph), data-end (the End
// layer), each range's value, max and --cui-console-player-value, the times,
// the cue and its lang, the caption. It finds each control by its
// data-control (TestPlayerMarksEachControl holds the names). The scrubber's
// preview stays hidden until the page fills its picture and line and shows it
// over the time under the pointer.
type ConsolePlayer struct {
	// Video is the picture, the player's first child: the page's <video>
	// (its source and tracks; no controls attribute), or an element the
	// page draws its stills in, which the player sizes as it would a video.
	Video g.Node
	// Title: "" for none.
	Title           string
	PlayAriaLabel   string
	VolumeAriaLabel string
	MuteAriaLabel   string
	ScrubAriaLabel  string
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
	// Intro: the layer before it first plays (data-intro); nil for none.
	Intro *ConsolePlayerIntro
	// End: the layer once it has played to the end (data-end, the page's
	// script's to set); nil for none.
	End *ConsolePlayerEnd
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
	for name, label := range map[string]string{"PlayAriaLabel": p.PlayAriaLabel, "VolumeAriaLabel": p.VolumeAriaLabel, "MuteAriaLabel": p.MuteAriaLabel, "ScrubAriaLabel": p.ScrubAriaLabel, "FullScreenAriaLabel": p.FullScreenAriaLabel} {
		if label == "" {
			panic("console.ConsolePlayer needs a " + name)
		}
	}
	if p.SubtitlesMenu.ID != "" && p.SubtitlesAriaLabel == "" {
		panic("console.ConsolePlayer needs a SubtitlesAriaLabel with a SubtitlesMenu")
	}
	if p.Intro != nil && (p.Intro.Title == "" || p.Intro.PlayLabel == "" || p.Intro.RestartLabel == "") {
		panic("console.ConsolePlayerIntro needs a Title, a PlayLabel and a RestartLabel")
	}
	if p.End != nil && (p.End.Title == "" || p.End.ReplayLabel == "") {
		panic("console.ConsolePlayerEnd needs a Title and a ReplayLabel")
	}
	button := func(control, label, menu string, glyph g.Node) g.Node {
		return h.Button(h.Class("cui-console-player-btn"), h.Type("button"), h.Data("control", control), h.Aria("label", label), optionalAttr("popovertarget", menu), glyph)
	}
	return h.Div(class.Root("cui-console-player", p.Class), g.If(p.Intro != nil, g.Attr("data-intro")),
		p.Video,
		h.Div(h.Class("cui-console-player-overlay"),
			h.Div(h.Class("cui-console-player-header"),
				h.P(h.Class("cui-console-player-caption"), h.Data("control", "caption")),
				// @invariant: the range opens beside its button while the pointer is over either
				h.Div(h.Class("cui-console-player-volume"),
					rangeInput("volume", "100", p.VolumeAriaLabel, "", "100%"),
					button("mute", p.MuteAriaLabel, "", g.Group{lucide.Volume2.Class("cui-console-player-glyph-volume"), lucide.VolumeX.Class("cui-console-player-glyph-muted")}),
				),
			),
			h.Button(h.Class("cui-console-player-play"), h.Type("button"), h.Data("control", "play"), h.Aria("label", p.PlayAriaLabel),
				lucide.Play.Class("cui-console-player-glyph-play"), lucide.Pause.Class("cui-console-player-glyph-pause")),
			h.Div(h.Class("cui-console-player-subtitles"),
				h.P(h.Span(h.Class("cui-console-player-cue"))),
			),
			h.Div(h.Class("cui-console-player-footer"),
				g.If(p.Title != "", h.P(h.Class("cui-console-player-title"), g.Text(p.Title))),
				h.Div(h.Class("cui-console-player-buttons"),
					g.If(p.FullScreenControls != nil, h.Div(h.Class("cui-console-player-full-screen-controls"), p.FullScreenControls)),
					g.If(p.SubtitlesMenu.ID != "", button("subtitles", p.SubtitlesAriaLabel, p.SubtitlesMenu.ID, lucide.Captions)),
					button("full-screen", p.FullScreenAriaLabel, "", g.Group{lucide.Maximize2.Class("cui-console-player-glyph-full-screen"), lucide.Minimize2.Class("cui-console-player-glyph-exit-full-screen")}),
				),
				h.Div(h.Class("cui-console-player-progress"),
					rangeInput("scrubber", "0", p.ScrubAriaLabel, "0:00", "0%"),
					// @invariant: the range's aria-valuetext says the time: the preview is for the eye
					h.Div(h.Class("cui-console-player-preview"), h.Data("control", "preview"), h.Aria("hidden", "true"), g.Attr("hidden"),
						h.Img(h.Class("cui-console-player-preview-picture"), h.Alt("")),
						h.P(h.Class("cui-console-player-preview-label")),
					),
					h.Time(h.Class("cui-console-player-time"), h.Data("control", "elapsed"), g.Text("0:00")),
					h.Time(h.Class("cui-console-player-time"), h.Data("control", "remaining"), g.Text("-0:00")),
				),
			),
		),
		g.Iff(p.Intro != nil, func() g.Node { return intro(*p.Intro) }),
		g.Iff(p.End != nil, func() g.Node {
			return h.Div(h.Class("cui-console-player-end"),
				h.P(h.Class("cui-console-player-end-title"), g.Text(p.End.Title)),
				h.Div(h.Class("cui-console-player-end-body"), p.End.Body),
				h.Button(h.Class("cui-console-player-text-btn"), h.Type("button"), h.Data("control", "replay"), g.Text(p.End.ReplayLabel)),
			)
		}),
		g.If(p.SubtitlesMenu.ID != "", p.SubtitlesMenu),
		p.Popovers,
	).Render(w)
}

func intro(i ConsolePlayerIntro) g.Node {
	return h.Div(h.Class("cui-console-player-intro"),
		h.P(h.Class("cui-console-player-intro-title"), g.Text(i.Title)),
		g.If(i.Detail != "", h.P(h.Class("cui-console-player-intro-detail"), g.Text(i.Detail))),
		g.If(i.Summary != "", h.P(h.Class("cui-console-player-intro-summary"), g.Text(i.Summary))),
		h.Div(h.Class("cui-console-player-intro-actions"),
			h.Button(h.Class("cui-console-player-intro-play"), h.Type("button"), h.Data("control", "start"), lucide.Play, h.Span(g.Text(i.PlayLabel))),
			h.Button(h.Class("cui-console-player-text-btn"), h.Type("button"), h.Data("control", "restart"), g.Attr("hidden"), g.Text(i.RestartLabel)),
		),
	)
}

// @invariant: at rest each range is at its max: the volume full, the scrubber 0 of 0
func rangeInput(control, max, label, valueText, percent string) g.Node {
	return h.Input(h.Class("cui-console-player-range"), h.Type("range"), h.Data("control", control), h.Min("0"), h.Max(max), h.Value(max),
		h.Aria("label", label), optionalAttr("aria-valuetext", valueText),
		h.Style("--cui-console-player-value: "+percent))
}
