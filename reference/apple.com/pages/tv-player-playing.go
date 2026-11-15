package pages

import (
	"slices"

	"github.com/at-least/ciderui/internal/pagespec"
)

// @external: tv.apple.com's video player playing (tv-player's page and opening): its
// controls shown by a click on the video, the centre pause in the play's
// place; then, the pointer still, the controls hidden by its script
// (`video-player-scrim--hidden`, 4s on), and what that leaves. At 1600px
// alone: a resize, and the time it takes, would move both.
func tvPlayerPlaying() pagespec.Spec {
	SHOWN := slices.Concat([]string{"opacity", "visibility", "pointer-events"}, MOTION)
	return pagespec.Spec{
		URL:     tvPlayerURL,
		Browser: "chrome",
		Wait:    2000,
		Widths:  []int{1600},
		Probes: append(tvPlayerOpen(), []pagespec.Probe{
			{Name: "pause", Select: "button.playback-play__pause", Props: slices.Concat(SHOWN, []string{"color", "width", "height"}), Box: true},
			{Name: "pause.icon", Select: "button.playback-play__pause amp-icon[slot=\"pause\"]", Props: []string{"color", "width", "height"}, Box: true},
			{Name: "play.hidden", Select: "button.playback-play__play", Props: SHOWN},
			{Name: "hidden", Select: ".scrim__gradient", Props: SHOWN, Box: true, HoverAt: pagespec.At(1599, 200), Hovered: 6000},
			{Name: "hidden.scrim", Select: ".video-player-scrim", Props: slices.Concat(SHOWN, []string{"cursor"}), Matches: []string{".video-player-scrim--hidden"}},
			{Name: "hidden.header", Select: ".scrim__header", Props: slices.Concat(SHOWN, []string{"transform"})},
			{Name: "hidden.body", Select: ".scrim__body", Props: slices.Concat(SHOWN, []string{"transform"})},
			{Name: "hidden.footer", Select: ".scrim__footer", Props: slices.Concat(SHOWN, []string{"transform"})},
		}...),
	}
}
