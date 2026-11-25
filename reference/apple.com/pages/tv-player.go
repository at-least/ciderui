package pages

import (
	"slices"

	"github.com/at-least/ciderui/internal/pagespec"
)

const tvPlayerURL = "https://tv.apple.com/us/collection/trailers/uts.col.Trailers.umc.cmc.vtoh0mn0xn7t3c643xqonfzy"

// @invariant: opening the player, showing its controls and pausing it: each once, in turn
func tvPlayerOpen() []pagespec.Probe {
	return []pagespec.Probe{
		{Name: "dialog", Click: "a.lockup", Clicked: 6000, Select: "dialog.playback-view", Props: []string{"background-color", "color", "padding-top", "padding-left"}, Box: true},
		{Name: "scrim", ClickAt: pagespec.At(700, 400), Clicked: 1000, Select: ".video-player-scrim", Props: slices.Concat(FLEX, []string{"grid-template-rows", "grid-template-columns", "padding-top", "padding-left", "padding-right", "padding-bottom", "opacity"}, MOTION), Box: true},
	}
}

// @external: tv.apple.com's video player, a trailer paused with its controls shown (the
// owner, 2026-11-05: 「播放器照 tv.apple.com」): the first trailer opened from Ted Lasso's
// trailers, the video clicked (its controls shown), the centre pause clicked; paused,
// the controls stay. The video is clicked at a point: over it lies the trailers'
// carousel, which takes a click on an element. The overlay's gradient, its rows (the
// header's volume, the centre play, the footer's scrubber, times and buttons), each
// control at rest, hovered, pressed and focused. Read in Google Chrome (`browser`):
// Playwright's Chromium opens the player and never plays (its <video> never made).
// tv-player-playing reads the playing state, tv-player-menu the subtitles menu,
// tv-player-playlist the trailers' row.
func tvPlayer() pagespec.Spec {
	STATES := []string{"hover", "active", "focus-visible"}
	CONTROL := slices.Concat(BOX, []string{"color", "opacity", "cursor", "width", "height", "transform"}, FOCUS, MOTION)
	// @external: the custom properties its range inputs are drawn by (a track, the part
	// played, the thumb), read on the input: a thumb's computed style is not
	// read (getComputedStyle with ::-webkit-slider-thumb gave the input's own)
	RANGE := []string{
		"--progress-track-height", "--progress-track-max-height", "--progress-track-color", "--progress-track-color-elapsed", "--progress-track-color-elapsed-active",
		"--progress-thumb-height", "--progress-thumb-width", "--progress-thumb-color", "--progress-thumb-color-active", "--progress-thumb-multiplier", "--progress-thumb-multiplier-active", "--progress-thumb-border-radius", "--progress-thumb-box-shadow", "--progress-thumb-outline-active", "--volume-thumb-box-shadow",
	}

	return pagespec.Spec{
		URL:     tvPlayerURL,
		Browser: "chrome",
		Wait:    2000,
		Widths:  []int{1600, 1200},
		Probes: append(tvPlayerOpen(), []pagespec.Probe{
			{Name: "play", Click: "button.playback-play__pause", Clicked: 1500, Pointer: pagespec.At(800, 200), Select: "button.playback-play__play", Props: CONTROL, Box: true, States: STATES},
			{Name: "play.icon", Select: "button.playback-play__play amp-icon[slot=\"play\"]", Props: []string{"color", "width", "height"}, Box: true},
			{Name: "pause.hidden", Select: "button.playback-play__pause", Props: slices.Concat([]string{"opacity", "visibility", "pointer-events"}, MOTION)},
			{Name: "scrim.gradient", Select: ".scrim__gradient", Props: slices.Concat([]string{"background-image", "background-color", "opacity"}, MOTION), Box: true, Position: true},
			{Name: "scrim.header", Select: ".scrim__header", Props: slices.Concat([]string{"opacity"}, MOTION), Box: true, Position: true},
			{Name: "scrim.footer", Select: ".scrim__footer", Props: slices.Concat([]string{"opacity", "padding-left", "padding-right", "padding-bottom"}, MOTION), Box: true, Position: true},
			{Name: "header", Select: ".video-player__header", Props: slices.Concat(FLEX, []string{"grid-template-columns", "padding-left", "padding-right", "padding-top"}), Box: true, Position: true},
			{Name: "footer", Select: ".video-player__footer", Props: slices.Concat(FLEX, []string{"grid-template-columns", "grid-template-rows", "grid-template-areas", "padding-left", "padding-right", "padding-bottom"}), Box: true, Position: true},
			{Name: "title", Select: ".video-player__metadata .title", Props: TYPE, Box: true, Position: true},
			{Name: "progress", Select: ".video-player__progress", Props: []string{"padding-top", "padding-bottom", "margin-top"}, Box: true, Position: true},
			{Name: "progress.bar", Select: "amp-playback-controls-progress", Props: slices.Concat(FLEX, []string{"grid-template-columns", "grid-template-rows", "grid-template-areas"}), Box: true, Position: true},
			{Name: "elapsed", Select: "amp-playback-controls-progress >>> time.elapsed", Props: slices.Concat(TYPE, []string{"font-variant-numeric", "font-feature-settings", "text-align", "grid-area"}), Box: true, Position: true},
			{Name: "remaining", Select: "amp-playback-controls-progress >>> time.remaining", Props: slices.Concat(TYPE, []string{"font-variant-numeric", "font-feature-settings", "text-align", "grid-area"}), Box: true, Position: true},
			{Name: "scrubber", Select: "amp-playback-controls-progress >>> input[type=\"range\"]", Props: slices.Concat([]string{"height", "cursor", "background-color", "color"}, FOCUS), Vars: RANGE, Box: true, Position: true, States: STATES},
			{Name: "volume", Select: ".video-player__volume", Props: slices.Concat(FLEX, []string{"grid-template-columns"}), Box: true, Position: true},
			{Name: "volume.range", Select: "amp-volume-control >>> input[type=\"range\"]", Props: slices.Concat([]string{"height", "width", "cursor", "color"}, FOCUS), Vars: RANGE, Box: true, Position: true, States: STATES},
			{Name: "volume.button", Select: "button.volume-unified__indicator", Props: CONTROL, Box: true, Position: true, States: STATES},
			{Name: "volume.icon", Select: "button.volume-unified__indicator svg", Props: []string{"color", "fill", "width", "height"}, Box: true},
			{Name: "controls", Select: ".video-player__controls", Props: slices.Concat(FLEX, []string{"grid-area"}), Box: true, Position: true},
			{Name: "subtitles", Select: "amp-captions-control >>> button.contextual-menu__trigger", Props: slices.Concat(CONTROL, []string{"backdrop-filter"}), Box: true, Position: true, States: STATES},
			{Name: "subtitles.icon", Select: "amp-captions-control >>> amp-icon", Props: []string{"color", "width", "height"}, Box: true},
			{Name: "audio", Select: "amp-audio-tracks-control >>> label", Props: CONTROL, Box: true, Position: true, States: []string{"hover"}},
			{Name: "audio.select", Select: "amp-audio-tracks-control >>> select", Props: slices.Concat([]string{"opacity", "color", "appearance", "cursor", "position", "width", "height"}, FOCUS), Box: true, States: []string{"hover", "focus-visible"}},
			{Name: "fullscreen", Select: "amp-playback-controls-full-screen >>> button", Props: CONTROL, Box: true, Position: true, States: STATES},
			{Name: "fullscreen.icon", Select: "amp-playback-controls-full-screen >>> amp-icon", Props: []string{"color", "width", "height"}, Box: true},
		}...),
	}
}
