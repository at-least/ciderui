//go:build browser

package consoleplaybutton

import (
	"strings"
	"testing"

	h "github.com/at-least/ciderui/internal/harness"
)

const (
	b               = "component-Button-variant"
	paused          = ".cui-console-play-button:not([data-state])"
	playing         = `.cui-console-play-button[data-state="playing"]`
	ring            = "playing, a ring round the glyph fills with how much has played (the owner, 2026-11-05: a progress ring round the icon, as Apple's Voice Memos and Podcasts mini buttons; no reference measured, so golden): 0.1875em out from the glyph, its line 0.125em, in the label's color on a track of it at a quarter (color-mix), from the top round clockwise (a conic gradient to --cui-console-play-button-progress, the page's script's)"
	labelledPaused  = ".cui-console-play-button:not([data-state])"
	labelledPlaying = `.cui-console-play-button[data-state="playing"]`
)

// @invariant: the page's script's progress on the playing one
func halfway(html string) string {
	return strings.Replace(html, `data-state="playing" type="button">`, `data-state="playing" type="button" style="--cui-console-play-button-progress: 40%">`, 1)
}

func TestCases(t *testing.T) {
	h.Run(t, []h.Case{
		{
			Name:   "play",
			HTML:   h.Example("play"),
			Widths: []int{1600, 1200, 390},
			Checks: []h.Check{
				{Target: paused, Props: []string{"box-height", "box-width", "border-top-left-radius", "font-size"}, TB: []string{b + "-controlSize-lg-blockSize", b + "-controlSize-lg-minWidth", b + "-controlShape-capsule-controlSize-lg-borderRadius", "component-Button-anatomy-artStart-withoutLabel-controlSize-lg-fontSize"}},
				{Target: paused, Props: []string{"background-color"}, TB: []string{b + "-colorRole-neutral-effect-default-variant-gray-backgroundColor"}},
				{Target: paused, Props: []string{"background-color"}, State: []string{"hover"}, TB: []string{b + "-colorRole-neutral-effect-hover-variant-gray-backgroundColor"}},
				{Target: paused, Props: []string{"background-color"}, State: []string{"active"}, TB: []string{b + "-colorRole-neutral-effect-press-variant-gray-backgroundColor"}},
				{Target: paused, Props: []string{"color"}, Golden: "the gray button's label color, as console-btn's icon buttons beside it: the toolkit's icon-only gray (its artOnlyColor, 60% gray) is not taken here"},
				{Target: paused + " .cui-console-play-button-glyph > .cui-console-play-button-paused", Props: []string{"box-height", "box-width"}, TB: []string{"component-Button-anatomy-artStart-withoutLabel-controlSize-lg-fontSize", "component-Button-anatomy-artStart-withoutLabel-controlSize-lg-fontSize"}},
				{Target: paused + " .cui-console-play-button-glyph > .cui-console-play-button-paused", Props: []string{"display", "fill", "stroke"}, Golden: "paused: Lucide's play shown, filled in the label's color and not stroked (in a translucent color the stroke over the fill drew a darker outline) (the toolkit has no play glyph; tv.apple.com's are SF Symbols, Apple's)"},
				{Target: paused + " .cui-console-play-button-glyph > .cui-console-play-button-playing", Props: []string{"display"}, Golden: "paused: the stop glyph hidden"},
				{Target: paused + " .cui-console-play-button-glyph", Pseudo: "::before", Props: []string{"content"}, Golden: "paused: no ring"},
				{Target: playing + " .cui-console-play-button-glyph > .cui-console-play-button-paused", Props: []string{"display"}, Golden: "playing (data-state, the page's script's): the play glyph hidden"},
				{Target: playing + " .cui-console-play-button-glyph > .cui-console-play-button-playing", Props: []string{"display", "fill", "stroke", "box-width"}, Golden: "playing: Lucide's square, stop (the owner, 2026-11-05: 「■ 停止」, where it was pause: a line replays from its start), filled, not stroked, at 0.625em, inside its ring"},
				{Target: playing + " .cui-console-play-button-glyph", Pseudo: "::before", Props: []string{"content", "position", "top", "left", "width", "height", "border-top-left-radius", "background-image", "mask-image"}, Golden: ring},
			},
		},
		{
			Name:   "play labelled",
			HTML:   halfway(h.Example("labelled")),
			Widths: []int{1200},
			Checks: []h.Check{
				{Target: labelledPaused, Props: []string{"box-height", "padding-left", "column-gap", "font-size", "border-top-left-radius"}, Golden: "muse's 「▶ 重聽這句」 (the owner, 2026-11-05): a small capsule console button, its glyph and its words, console-btn's small size and gap"},
				{Target: labelledPaused + " > span.cui-console-play-button-paused", Props: []string{"display"}, Golden: "paused, its words: what a press does"},
				{Target: labelledPaused + " > span.cui-console-play-button-playing", Props: []string{"display"}, Golden: "and not the playing ones"},
				{Target: labelledPaused + " .cui-console-play-button-glyph", Props: []string{"box-width", "box-height"}, Golden: "the glyph the label's size, 1em"},
				{Target: labelledPlaying + " > span.cui-console-play-button-playing", Props: []string{"display"}, Golden: "playing, its playing words (「■ 停止」), switched by data-state with the glyph"},
				{Target: labelledPlaying + " > span.cui-console-play-button-paused", Props: []string{"display"}, Golden: "and not the paused ones"},
				{Target: labelledPlaying + " .cui-console-play-button-glyph", Pseudo: "::before", Props: []string{"background-image", "width", "height"}, Golden: ring + "; 40% here"},
			},
		},
		{
			Name:         "play labelled forced colors",
			HTML:         halfway(h.Example("labelled")),
			Widths:       []int{1200},
			ForcedColors: true,
			Checks: []h.Check{
				{Target: labelledPlaying + " .cui-console-play-button-glyph", Pseudo: "::before", Props: []string{"background-image"}, Golden: "the ring in the theme's ButtonText on GrayText"},
			},
		},
		{
			Name:   "play dark",
			HTML:   `<div class="cui-dark">` + h.Example("play") + `</div>`,
			Widths: []int{1200},
			Checks: []h.Check{
				{Target: paused, Props: []string{"background-color", "color"}, Golden: "the dark side derived (tokens.css, decision 18)"},
				{Target: paused + " .cui-console-play-button-glyph > .cui-console-play-button-paused", Props: []string{"fill"}, Golden: "filled in the dark label's color"},
			},
		},
	})
}
