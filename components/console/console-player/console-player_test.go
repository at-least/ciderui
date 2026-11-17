//go:build browser

package consoleplayer

import (
	"regexp"
	"slices"
	"strconv"
	"strings"
	"testing"
	"time"

	h "github.com/at-least/ciderui/internal/harness"
	apple "github.com/at-least/ciderui/reference/apple.com/results"
)

const (
	p                = ".cui-console-player"
	btn              = `.cui-console-player-btn[aria-label="Subtitles"]`
	scrub            = ".cui-console-player-progress > .cui-console-player-range"
	volume           = ".cui-console-player-volume > .cui-console-player-range"
	openWhy          = "the owner's (2026-11-05: only the speaker until hovered, the range opening beside it): tv.apple.com's range is 100px open beside its button with the pointer elsewhere (tv-player:volume.range, the pointer at 800,200), so none of the closing is measured; it opens over 0.3s on the header's own fade curve (tv-player-playing:hidden.header)"
	layout           = "tv.apple.com's player's room: 30px at the sides, 20px at the top and foot (tv-player:scrim's grid, 20px 36px 699px 125px 20px by 30px 1540px 30px); its rows the header, the centre and the footer, whatever the video's size"
	stage            = ".cui-console-player > :first-child"
	line             = ".cui-console-player-subtitles > p"
	cue              = ".cui-console-player-cue"
	muse             = "muse's console's subtitles, not tv.apple.com's, whose cues are its <video>'s own text track (shown, none active in the first 22s of its first trailer, 2026-11-05; ::cue unreadable): 2.4% of the width at the least 15px, so it grows with the player (full screen too), 1.6 lines, 500, white under a soft shadow (--cui-video-cue-shadow) and no box (the owner, 2026-11-05, from Apple's view), 42em at its widest, balanced"
	last             = `#cui-console-subtitles [value="vi"]`
	preview          = ".cui-console-player-preview"
	previewWhy       = "muse's, tv.apple.com's trailers scrubbing with no preview (the owner, 2026-11-05: the shot's thumbnail, its scene and the time over the pointer): the storyboard's 320px thumbnail at 160px, sharp at 2x; 12px above the bar, clear of the thumb; the line under it 12px semibold in the times' tabular figures, white under the cue's shadow, on one line however long the scene's name"
	textBtn          = ".cui-console-player-text-btn"
	full             = "muse's, for its 回報問題 and 下一個要看 while the player fills the screen (the owner, 2026-11-05); tv.apple.com's bar has no control of words: a 24px capsule, the bar's buttons' height (tv-player:audio), white on a 16% white fill, 24% hovered or pressed, 13px medium, 8px apart; half opacity disabled"
	heading          = ".cui-console-player-heading"
	paused           = "muse's (the owner, 2026-11-05: paused, a small line at the picture's top left saying the story, the scene and the shot; gone on play): in tv.apple.com's header's first column, 720px and empty on its trailers (tv-player:header, grid-template-columns 720px 820px), on the header's row; 15px semibold white under the cue's shadow (--cui-video-cue-shadow) so it reads on a bright picture with no box, cut short with an ellipsis before the volume; tv.apple.com shows no such line"
	intro            = "muse's pre-play layer (the owner, 2026-11-05: over the cover the title large, the demo's length, its description's first sentence, one clear Play; tv.apple.com's pre-play page was not measured, so golden): at the picture's foot in the overlay's room (30px at the sides, tv-player:scrim), on a soft gradient in that corner alone (--cui-video-title-scrim; the owner, 2026-11-08: tv.apple.com's, --cui-video-scrim, darkened the cover's lower half to black), the lines white under the cue's shadow; the length 15px medium and the sentence 17px on 24px in the controls' 90% white, two lines at most, 36em wide"
	end              = "muse's end layer (the owner, 2026-11-05: on end a 「看完了」 layer fades in over the last picture, the user's next step under it and a small 「重新播放」; tv.apple.com's end of a film was not measured, so golden): the intro's layer, its place, gradient and title, in the other direction; the next step in the length's type (15px medium, the controls' 90% white), its buttons and links one under another 8px apart, 14px under the title as the intro's Play is"
	subtitlesNoSlide = "at its other place at once as the controls show and hide: a slide between the two only draws the eye (DESIGN.md decision 25)"
	noSlide          = "faded where they are: tv.apple.com's header goes 20px up and its footer 20px down as they hide (tv-player-playing:hidden.*), a slide that only draws the eye (DESIGN.md decision 25); the fades' durations and curves its own"
)

var (
	tv       = []int{1600, 1200}
	rangeAsA = map[string]string{"--cui-console-player-track": "--progress-track-height", "--cui-console-player-thumb": "--progress-thumb-height", "--cui-console-player-thumb-width": "--progress-thumb-width"}
)

func hidden(html string) string {
	return strings.Replace(html, `data-state="playing"`, `data-state="playing" data-controls="hidden"`, 1)
}

// @external: muse's player: its subtitles' Off and eleven languages, and no voices (one)
// or two, the player in a column beside a 288px sidebar under a 50px bar, as
// wide as the window's height leaves (muse's player_page.go)
var languageNames = [][2]string{{"zh-Hant", "繁體中文"}, {"zh-Hans", "简体中文"}, {"en", "English"}, {"ja", "日本語"}, {"ko", "한국어"}, {"es", "Español"}, {"pt-BR", "Português (Brasil)"}, {"de", "Deutsch"}, {"fr", "Français"}, {"th", "ภาษาไทย"}, {"vi", "Tiếng Việt"}}

var (
	offRe          = regexp.MustCompile(`<button class="cui-console-menu-item"[^>]*value="off"[\s\S]*?</button>`)
	subtitlesRe    = regexp.MustCompile(`(<div role="menu" aria-labelledby="cui-console-subtitles-subtitles-title">)[\s\S]*?(</div></div></div></div></div>$)`)
	voicesColumnRe = regexp.MustCompile(`<div class="cui-console-menu" data-column="voices">[\s\S]*?</div></div>\n`)
	voicedRe       = regexp.MustCompile(`(<div role="menu" aria-labelledby="cui-console-subtitles-voices-title">)[\s\S]*?(</div></div>\n)$`)
	pictureRe      = regexp.MustCompile(`<img class="absolute[^>]* src="([^"]+)"`)
	endBodyRe      = regexp.MustCompile(`<div class="cui-console-player-end-body">[\s\S]*?</div>\n?<button class="cui-console-player-text-btn" type="button" data-control="replay">`)
)

func replaceFirst(re *regexp.Regexp, s string, replace func(groups []string) string) string {
	at := re.FindStringSubmatchIndex(s)
	if at == nil {
		return s
	}
	groups := make([]string, len(at)/2)
	for i := range groups {
		groups[i] = s[at[2*i]:at[2*i+1]]
	}
	return s[:at[0]] + replace(groups) + s[at[1]:]
}

func languages(html string, voices int) string {
	off := offRe.FindString(html)
	if off == "" {
		panic("the stills example has no Off choice where languages() looks")
	}
	// @invariant: Off is checked at rest, the column's first; a voices column checks its first language
	choice := func(checked bool, language [2]string) string {
		c := strings.Replace(off, `aria-checked="true"`, `aria-checked="`+strconv.FormatBool(checked)+`"`, 1)
		c = strings.Replace(c, `value="off"`, `value="`+language[0]+`"`, 1)
		return strings.Replace(c, `<span class="cui-console-menu-label">Off</span>`, `<span class="cui-console-menu-label" lang="`+language[0]+`">`+language[1]+`</span>`, 1)
	}
	choices := func(n int, checkFirst bool) []string {
		var out []string
		for i, language := range languageNames[:n] {
			out = append(out, choice(checkFirst && i == 0, language))
		}
		return out
	}
	if !subtitlesRe.MatchString(html) {
		panic("the stills example has no subtitles column where languages() looks")
	}
	if !voicesColumnRe.MatchString(html) {
		panic("the stills example has no voices column where languages() looks")
	}
	voiced := func(column string) string {
		return replaceFirst(voicedRe, column, func(g []string) string {
			return g[1] + "\n" + strings.Join(choices(voices, true), "\n") + g[2]
		})
	}
	player := replaceFirst(subtitlesRe, html, func(g []string) string {
		return g[1] + "\n" + strings.Join(append([]string{off}, choices(len(languageNames), false)...), "\n") + g[2]
	})
	player = replaceFirst(voicesColumnRe, player, func(g []string) string {
		if voices > 1 {
			return voiced(g[0])
		}
		return ""
	})
	return `<div style="height:50px"></div><div style="display:flex"><div style="flex:1;min-width:0;padding:0 24px"><div style="margin:16px auto 0;max-width:calc((100vh - 14rem) * 16 / 9)">` + player + `</div></div><div style="flex:none;width:288px"></div></div>`
}

// @invariant: the preview as the page's script shows it: its picture, its line and the pointer's x
func previewing(html string) string {
	src := pictureRe.FindStringSubmatch(html)
	if src == nil {
		panic("the example has no picture where previewing() looks")
	}
	html2 := strings.Replace(html, `aria-hidden="true" hidden>`, `aria-hidden="true" style="--cui-console-player-preview-x: 300px">`, 1)
	html2 = strings.Replace(html2, `<img class="cui-console-player-preview-picture" alt="">`, `<img class="cui-console-player-preview-picture" alt="" src="`+src[1]+`">`, 1)
	return strings.Replace(html2, `<p class="cui-console-player-preview-label"></p>`, `<p class="cui-console-player-preview-label">場景 4　大野狼到外婆家 · 2:31</p>`, 1)
}

// @invariant: the page's script's heading, written into the player at rest
func withHeading(html string) string {
	return strings.Replace(html, `data-control="heading"></p>`, `data-control="heading">The Wedding · Scene 2　The Vows · S02-03</p>`, 1)
}

func TestCases(t *testing.T) {
	h.Run(t, []h.Case{
		{
			Name:   "player paused",
			HTML:   h.Example("paused"),
			Widths: tv,
			Height: 900,
			Checks: []h.Check{
				{Target: p, Props: []string{"background-color"}, Apple: h.AppleProbe(apple.TvPlayer, "dialog")},
				{Target: p, Props: []string{"aspect-ratio"}, Golden: "its own 16:9, tv.apple.com's player's (tv-player:scrim, 1600px by 900px): a page's element of stills has no size of its own"},
				{Target: p, Props: []string{"position", "overflow-x"}, Golden: "the containing block its overlay is laid over its video in; not clipped: a window too small for its controls grows it to them (its aspect-ratio's automatic minimum, which a clip would drop)"},
				{Target: ".cui-console-player-overlay", Props: []string{"display", "grid-template-columns", "grid-template-rows"}, Golden: layout},
				{Target: ".cui-console-player-overlay", Pseudo: "::before", Props: []string{"background-image", "background-color", "opacity"}, Apple: h.AppleProbe(apple.TvPlayer, "scrim.gradient")},
				{Target: ".cui-console-player-overlay", Pseudo: "::before", Props: []string{"transition-property", "transition-duration"}, Apple: h.AppleProbe(apple.TvPlayer, "scrim.gradient")},
				{Target: ".cui-console-player-play", Props: []string{"color", "cursor", "opacity", "background-color", "border-top-width", "padding-top"}, Apple: h.AppleProbe(apple.TvPlayer, "play")},
				{Target: ".cui-console-player-play", Props: []string{"box-width", "box-height"}, Golden: "tv.apple.com's centre button, 66px (tv-player:play's box: an apple check reads no width)"},
				{Target: ".cui-console-player-play", Props: []string{"transform"}, State: []string{"active"}, Apple: h.AppleProbe(apple.TvPlayer, "play", "active")},
				{Target: ".cui-console-player-play", Props: []string{"transition-property", "transition-duration", "transition-timing-function"}, Golden: "its press's scale in tv.apple.com's 0.2s ease-out (tv-player:play), and its fade as the controls hide, the centre's 0.4s (tv-player-playing:hidden.body)"},
				{Target: ".cui-console-player-glyph-play", Props: []string{"box-width", "box-height", "display"}, Golden: "paused, the play glyph, tv.apple.com's 50px (tv-player:play.icon), Lucide's filled (decision 17: tv.apple.com's is an SF Symbol)"},
				{Target: ".cui-console-player-glyph-pause", Props: []string{"display"}, Golden: "paused, no pause glyph: tv.apple.com's pause is hidden then (tv-player:pause.hidden)"},
				{Target: ".cui-console-player-volume", Props: []string{"column-gap"}, Apple: h.AppleProbe(apple.TvPlayer, "volume")},
				{Target: ".cui-console-player-header", Props: []string{"display", "justify-content", "align-items"}, Golden: "the volume at the header's end, as tv.apple.com's (tv-player:volume at the right)"},
				{Target: ".cui-console-player-buttons", Props: []string{"column-gap", "display"}, Apple: h.AppleProbe(apple.TvPlayer, "controls")},
				{Target: ".cui-console-player-footer", Props: []string{"row-gap", "column-gap", "grid-template-areas"}, Golden: "the buttons at the row's end with nothing before them (no title: the heading at the top says what plays), the scrubber under them, 12px apart (tv-player:footer's 12px rows and gap)"},
				{Target: btn, Props: []string{"color", "cursor", "background-color", "border-top-width"}, Apple: h.AppleProbe(apple.TvPlayer, "subtitles")},
				{Target: `.cui-console-player-btn[aria-label="Full Screen"]`, Props: []string{"color"}, State: []string{"hover"}, Apple: h.AppleProbe(apple.TvPlayer, "fullscreen", "hover")},
				{Target: btn, Props: []string{"color"}, State: []string{"hover"}, Golden: "hovered white, as tv.apple.com's full screen button (tv-player:fullscreen:hover); its subtitles button reads no change, one look for the row"},
				{Target: btn, Props: []string{"box-width", "box-height"}, Golden: "tv.apple.com's 24px (tv-player:audio, its controls' row 24px)"},
				{Target: btn + " > svg", Props: []string{"box-width", "box-height"}, Golden: "its glyph 20px, tv.apple.com's (tv-player:subtitles.icon); Lucide's (decision 17)"},
				{Target: ".cui-console-player-volume > .cui-console-player-btn", Props: []string{"box-width", "color"}, Golden: "the mute button 26px in white, tv.apple.com's (tv-player:volume.button, its glyph filled white); its Lucide glyph 22px, a stroke glyph drawn smaller than a filled one"},
				{Target: ".cui-console-player-time", Props: []string{"font-size", "line-height", "font-weight", "color", "font-variant-numeric"}, Apple: h.AppleProbe(apple.TvPlayer, "elapsed")},
				{Target: ".cui-console-player-time:last-of-type", Props: []string{"color"}, Apple: h.AppleProbe(apple.TvPlayer, "remaining")},
				{Target: scrub, Props: []string{"--cui-console-player-track", "--cui-console-player-thumb", "--cui-console-player-thumb-width"}, Apple: &h.Apple{Page: apple.TvPlayer, Probe: "scrubber", As: rangeAsA}},
				{Target: scrub, Props: []string{"cursor"}, Apple: h.AppleProbe(apple.TvPlayer, "scrubber")},
				{Target: scrub, Props: []string{"box-height", "appearance", "--cui-console-player-value"}, Golden: "a 44px target for any pointer, the HIG's default (reference/touch/), where tv.apple.com's is 25px (tv-player:scrubber's box); the share played the page's, written as a percentage"},
				{Target: scrub, Props: []string{"position"}, Golden: "positioned, over the times' text its target reaches down into: unpositioned, a point on the time played hit the <time>, not the range"},
				{Target: scrub, Props: []string{"margin-top", "margin-bottom"}, Golden: "its margins give back all but the thumb's 13px, tv.apple.com's row (tv-player:progress.bar 13px 14px): the bar sits where it did, the line as thin"},
				{Target: ".cui-console-player-progress", Props: []string{"grid-template-rows"}, Apple: h.AppleProbe(apple.TvPlayer, "progress.bar")},
				{Target: scrub, Props: []string{"--cui-console-player-track"}, State: []string{"active"}, Golden: "held, the line 8px, thicker under the thumb (the owner, 2026-11-05: 「thicker while dragged」); tv.apple.com's forced :active reads 5px (tv-player:scrubber:active), its drag drawn by its script"},
				{Target: ".cui-console-player-buttons", Props: []string{"z-index", "position"}, Golden: "over the scrubber's target, which reaches 4.5px up into their row: a click on a button's foot is the button's"},
				{Target: preview, Props: []string{"display", "pointer-events"}, Golden: "the preview hidden until the page's script shows it, and never in the pointer's way"},
				{Target: volume, Props: []string{"--cui-console-player-track", "--cui-console-player-thumb", "--cui-console-player-thumb-width"}, Apple: &h.Apple{Page: apple.TvPlayer, Probe: "volume.range", As: rangeAsA}},
				{Target: volume, Props: []string{"box-width", "opacity", "margin-right"}, Golden: "closed, the speaker alone: no range and none of the 8px gap before the button; " + openWhy},
				{Target: volume, Props: []string{"pointer-events"}, Golden: "closed, it takes no pointer: its 24 by 13px thumb overflows its 0px over the button, and took 6px of the button's target (the touch audit read the button 38px wide)"},
				{Target: volume, Props: []string{"transition-property", "transition-duration", "transition-timing-function"}, Golden: openWhy},
				{Target: ".cui-console-player-glyph-volume", Props: []string{"display"}, Golden: "not muted, the speaker"},
				{Target: ".cui-console-player-glyph-muted", Props: []string{"display"}, Golden: "not muted, no struck speaker"},
				{Target: cue, Props: []string{"display"}, Golden: "no cue, no box: the slot is written empty until the page's script fills it"},
				{Target: "#cui-console-subtitles", Props: []string{"display"}, Golden: "its menu written inside it, closed until its button opens it"},
			},
		},
		{
			Name:   "player stills",
			HTML:   h.Example("stills"),
			Widths: tv,
			Height: 900,
			Checks: []h.Check{
				{Target: stage, Props: []string{"box-width", "box-height", "box-top", "box-left"}, Golden: "an element of stills in the video's place, the player's first child: the player's whole box, under the overlay"},
				{Target: stage + " > img", Props: []string{"box-width", "box-height"}, Golden: "the page's stills fill it (its own utilities: the player sizes the element, not what is in it)"},
				{Target: line, Props: []string{"font-size", "line-height", "font-weight", "color", "text-align", "text-wrap-style", "max-width"}, Golden: muse},
				{Target: cue, Props: []string{"text-shadow", "background-color", "padding-top", "padding-left", "border-top-left-radius"}, Golden: muse},
				{Target: line, Props: []string{"box-top", "box-height"}, Golden: "over the footer while the controls show: its foot 12px above the footer's top, the footer's own gap (tv-player:footer's 12px)"},
				{Target: ".cui-console-player-footer", Props: []string{"box-top"}, Golden: "the footer, under the subtitles, where it is with none"},
				{Target: ".cui-console-player-subtitles", Props: []string{"pointer-events"}, Golden: "the controls under it take the pointer"},
				{Target: ".cui-console-player-subtitles", Props: []string{"transition-property", "transition-duration", "transition-timing-function"}, Golden: subtitlesNoSlide},
				{Target: cue, Props: []string{"hit"}, Golden: "a click on the cue reaches the overlay under it, which the page's script hears (a click on the picture plays or pauses)"},
			},
		},
		{
			Name:   "player scrub preview",
			HTML:   previewing(h.Example("stills")),
			Widths: tv,
			Height: 900,
			Checks: []h.Check{
				{Target: preview, Props: []string{"box-width", "box-height", "box-left", "box-top", "display", "row-gap", "justify-items"}, Golden: previewWhy},
				{Target: ".cui-console-player-preview-picture", Props: []string{"box-width", "box-height", "border-top-left-radius", "object-fit", "background-color", "box-shadow"}, Golden: previewWhy},
				{Target: ".cui-console-player-preview-label", Props: []string{"box-width", "box-left"}, Golden: "a line longer than the picture centred under it, over both its sides"},
				{Target: ".cui-console-player-preview-label", Props: []string{"font-size", "line-height", "font-weight", "color", "white-space", "font-variant-numeric", "text-shadow"}, Golden: previewWhy},
				{Target: preview, Props: []string{"hit"}, Golden: "it takes no pointer: what is under it does (the overlay), so it never flickers under the pointer that shows it"},
			},
		},
		{
			Name:   "player scrub preview at the start",
			HTML:   strings.Replace(previewing(h.Example("stills")), "--cui-console-player-preview-x: 300px", "--cui-console-player-preview-x: 4px", 1),
			Widths: []int{1600},
			Height: 900,
			Checks: []h.Check{
				{Target: preview, Props: []string{"box-left"}, Golden: "the pointer near the bar's start: the preview stops at it, inside the bar (the bar at 30px)"},
			},
		},
		{
			Name:   "player volume open",
			HTML:   h.Example("paused"),
			Widths: tv,
			Height: 900,
			Hover:  []string{`[data-control="mute"]`},
			Checks: []h.Check{
				{Target: volume, Props: []string{"pointer-events"}, Golden: "open, it takes the pointer"},
				{Target: volume, Props: []string{"box-width", "opacity", "margin-right"}, Golden: "the pointer over its button, the range open: tv.apple.com's 100px (tv-player:volume.range), its thumb white as the scrubber's (tv-player:volume.range --progress-thumb-color); " + openWhy},
				{Target: ".cui-console-player-volume > .cui-console-player-btn", Props: []string{"box-left"}, Golden: "the button where it was: the range opens to its left (tv-player:volume.button at the header's end)"},
			},
		},
		{
			Name:   "player muted",
			HTML:   strings.Replace(h.Example("paused"), `<div class="cui-console-player">`, `<div class="cui-console-player" data-muted>`, 1),
			Widths: []int{1600},
			Height: 900,
			Checks: []h.Check{
				{Target: ".cui-console-player-glyph-muted", Props: []string{"display", "box-width"}, Golden: "muted, the speaker struck through, Lucide's volume-x at the speaker's 22px (tv.apple.com's muted glyph not read)"},
				{Target: ".cui-console-player-glyph-volume", Props: []string{"display"}, Golden: "muted, no plain speaker"},
			},
		},
		{
			Name:   "player heading paused",
			HTML:   withHeading(h.Example("paused")),
			Widths: tv,
			Height: 900,
			Checks: []h.Check{
				{Target: heading, Props: []string{"display", "opacity", "visibility", "font-size", "line-height", "font-weight", "color", "text-shadow", "white-space", "text-overflow", "box-left", "box-top", "box-height"}, Golden: paused},
				{Target: heading, Props: []string{"transition-property", "transition-duration", "transition-timing-function"}, Golden: "it comes and goes in the header's fade (tv-player-playing:hidden.header, 0.3s)"},
			},
		},
		{
			Name:   "player heading playing",
			HTML:   h.Example("playing"),
			Widths: []int{1600},
			Height: 900,
			Checks: []h.Check{
				{Target: volume, On: ".cui-console-player-volume > .cui-console-player-btn", State: []string{"active"}, Props: []string{"margin-right"}, Golden: "its range stays open while a control in it is pressed"},
				{Target: volume, On: ".cui-console-player-volume > .cui-console-player-btn", State: []string{"active"}, Props: []string{"pointer-events"}, Golden: "but takes no pointer while its button alone is pressed: opening from 0px, its thumb lay over the button, and on a screen with no hover the release landed on it and the speaker did not mute (muse, at 30%)"},
				{Target: btn, State: []string{"active"}, Props: []string{"color"}, Golden: "pressed, full white, as the pointer over it shows on a screen that has one"},
				{Target: heading, Props: []string{"opacity", "visibility"}, Golden: "gone while playing (the owner, 2026-11-05), the controls shown or not"},
			},
		},
		{
			Name:   "player heading empty",
			HTML:   h.Example("paused"),
			Widths: []int{1600},
			Height: 900,
			Checks: []h.Check{
				{Target: ".cui-console-player-volume > .cui-console-player-btn", Props: []string{"box-height"}, Golden: "the speaker a little taller than the bar's other buttons"},
				{Target: heading, Props: []string{"display"}, Golden: "empty, nothing drawn: the volume keeps the header's end"},
			},
		},
		{
			Name:   "player full screen glyph",
			HTML:   h.Example("full-screen"),
			Widths: []int{1600},
			Height: 900,
			Checks: []h.Check{
				{Target: `[data-control="full-screen"] .cui-console-player-glyph-exit-full-screen`, Props: []string{"display", "box-width"}, Golden: "filling the screen, the button shows the way out (Lucide's minimize-2 for its maximize-2), as tv.apple.com's full-screen button turns to exit (its glyph not read: a page cannot be put in full screen by the probe)"},
				{Target: `[data-control="full-screen"] .cui-console-player-glyph-full-screen`, Props: []string{"display"}, Golden: "and not the way in"},
			},
		},
		{
			Name:   "player full screen glyph not full screen",
			HTML:   h.Example("paused"),
			Widths: []int{1600},
			Height: 900,
			Checks: []h.Check{
				{Target: `[data-control="full-screen"] .cui-console-player-glyph-full-screen`, Props: []string{"display", "box-width"}, Golden: "not full screen, the way in"},
				{Target: `[data-control="full-screen"] .cui-console-player-glyph-exit-full-screen`, Props: []string{"display"}, Golden: "and not the way out"},
			},
		},
		{
			Name:   "player intro",
			HTML:   h.Example("intro"),
			Widths: tv,
			Height: 900,
			Checks: []h.Check{
				{Target: ".cui-console-player-intro-play", State: []string{"active"}, Props: []string{"transform", "background-color"}, Golden: "pressed, the hovered fill, the controls' 90% white (--cui-c-on-video-control), and no scale: the owner, 2026-11-09 (the centre play's scale is tv.apple.com's, measured; this button's had no measure)"},
				{Target: ".cui-console-player-intro", Props: []string{"display", "flex-direction", "justify-content", "align-items", "padding-left", "padding-bottom", "background-image", "opacity", "visibility"}, Golden: intro},
				{Target: ".cui-console-player-overlay", Props: []string{"visibility"}, Golden: "its controls hidden under it: one Play, nothing else to press (the owner, 2026-11-05)"},
				{Target: ".cui-console-player-intro-title", Props: []string{"font-size", "line-height", "font-weight", "color", "box-left"}, Golden: intro + "; the title as large as 4% of the player, at the least tv.apple.com's 26px bold title (tv-player:title)"},
				{Target: ".cui-console-player-intro-detail", Props: []string{"font-size", "line-height", "font-weight", "color"}, Golden: intro},
				{Target: ".cui-console-player-intro-body", Props: []string{"font-size", "line-height", "color", "max-width", "-webkit-line-clamp"}, Golden: intro},
				{Target: ".cui-console-player-intro-play", Props: []string{"box-height", "padding-left", "border-top-left-radius", "background-color", "color", "font-size", "font-weight", "column-gap"}, Golden: intro + "; Play a white 44px capsule in the picture's black, the HIG's 44pt target"},
				{Target: ".cui-console-player-intro-play", Props: []string{"background-color"}, State: []string{"hover"}, Golden: "hovered, the controls' 90% white (--cui-c-on-video-control)"},
				{Target: `.cui-console-player-intro [data-control="restart"]`, Props: []string{"display"}, Golden: "the smaller start-over hidden until the page's script finds a saved place"},
				{Target: ".cui-console-player-intro", Props: []string{"transition-property", "transition-duration"}, Golden: "it fades out in the footer's 0.4s curve (tv-player-playing:hidden.footer)"},
			},
		},
		{
			Name:   "player intro resume",
			HTML:   strings.Replace(h.Example("intro"), `data-control="restart" hidden`, `data-control="restart"`, 1),
			Widths: []int{1600},
			Height: 900,
			Checks: []h.Check{
				{Target: ".cui-console-player-intro-actions", Props: []string{"align-items", "column-gap"}, Golden: "its buttons on one center line, apart"},
				{Target: `.cui-console-player-intro [data-control="restart"]`, Props: []string{"display", "box-height", "background-color"}, Golden: "a saved place: start over beside Play, the bar's capsule of words, smaller"},
			},
		},
		{
			Name:   "player intro gone",
			HTML:   strings.Replace(h.Example("intro"), " data-intro", "", 1),
			Widths: []int{1600},
			Height: 900,
			Checks: []h.Check{
				{Target: ".cui-console-player-intro", Props: []string{"opacity", "visibility", "pointer-events"}, Golden: "playing (data-intro dropped by the page's script), the layer gone, taking no click while it fades (muse: a click on Play 0.1s after the intro went landed on the fading intro)"},
				{Target: ".cui-console-player-overlay", Props: []string{"visibility"}, Golden: "and the controls back"},
			},
		},
		{
			Name:   "player end",
			HTML:   h.Example("end"),
			Widths: tv,
			Height: 900,
			Checks: []h.Check{
				{Target: ".cui-console-player-end", Props: []string{"display", "flex-direction", "justify-content", "align-items", "padding-left", "padding-bottom", "background-image", "opacity", "visibility"}, Golden: end},
				{Target: ".cui-console-player-overlay", Props: []string{"visibility"}, Golden: "its controls hidden under it: the next step and Play Again, nothing else to press"},
				{Target: ".cui-console-player-end-title", Props: []string{"font-size", "line-height", "font-weight", "color", "box-left"}, Golden: end + "; the title the intro's, 4% of the player and 26px at the least"},
				{Target: ".cui-console-player-end-body", Props: []string{"display", "flex-direction", "align-items", "row-gap", "margin-top", "max-width", "font-size", "line-height", "font-weight", "color"}, Golden: end},
				{Target: ".cui-console-player-end-body > p", Props: []string{"margin-top", "margin-bottom"}, Golden: "its words without a paragraph's margin: the 8px gap spaces them"},
				{Target: `.cui-console-player-end [data-control="replay"]`, Props: []string{"display", "box-height", "background-color", "margin-top"}, Golden: "Play Again, the bar's capsule of words, smaller than the next step, 14px under it"},
				{Target: ".cui-console-player-end", Props: []string{"transition-property", "transition-duration"}, Golden: "it fades in in the footer's 0.4s curve, as the intro fades out"},
			},
		},
		{
			Name:   "player end not ended",
			HTML:   strings.Replace(h.Example("end"), " data-end", "", 1),
			Widths: []int{1600},
			Height: 900,
			Checks: []h.Check{
				{Target: ".cui-console-player-end", Props: []string{"opacity", "visibility", "pointer-events"}, Golden: "before the end (data-end not set, or dropped as it plays again), the layer gone, taking no click while it fades (muse: a click on Play 0.1s after the intro went landed on the fading intro)"},
				{Target: ".cui-console-player-overlay", Props: []string{"visibility"}, Golden: "and the controls there"},
			},
		},
		{
			Name: "player end empty",
			HTML: replaceFirst(endBodyRe, h.Example("end"), func([]string) string {
				return `<div class="cui-console-player-end-body"></div><button class="cui-console-player-text-btn" type="button" data-control="replay">`
			}),
			Widths: []int{1600},
			Height: 900,
			Checks: []h.Check{
				{Target: ".cui-console-player-end-body", Props: []string{"display"}, Golden: "no next step yet (the page fills it as the film ends): no room kept for it"},
			},
		},
		{
			Name:   "player full screen controls",
			HTML:   h.Example("full-screen"),
			Widths: tv,
			Height: 900,
			Checks: []h.Check{
				{Target: textBtn, State: []string{"active"}, Props: []string{"background-color"}, Golden: "pressed, the hovered fill"},
				{Target: ".cui-console-player-full-screen-controls", Props: []string{"display", "column-gap", "box-height"}, Golden: full},
				{Target: textBtn, Props: []string{"box-height", "padding-left", "border-top-left-radius", "background-color", "color", "font-size", "line-height", "font-weight", "white-space", "cursor"}, Golden: full},
				{Target: textBtn, Props: []string{"background-color"}, State: []string{"hover"}, Golden: full},
				{Target: ".cui-console-player-full-screen-controls", Props: []string{"box-left"}, Golden: "before the bar's buttons, in their row: the page's controls first, the player's own at the end"},
			},
		},
		{
			// @external: no example has a disabled control of words: the page disables one
			// (muse's 下一個要看 with nothing next)
			Name:   "player full screen control disabled",
			HTML:   strings.Replace(h.Example("full-screen"), `type="button">Next to Watch`, `type="button" disabled>Next to Watch`, 1),
			Widths: []int{1600},
			Height: 900,
			Checks: []h.Check{
				{Target: textBtn + ":disabled", Props: []string{"opacity", "cursor", "background-color"}, Golden: full},
				{Target: textBtn + ":disabled", State: []string{"hover"}, Props: []string{"background-color"}, Golden: "disabled, the fill at rest under the pointer"},
			},
		},
		{
			Name:   "player full screen controls not full screen",
			HTML:   strings.Replace(h.Example("full-screen"), " data-full-screen", "", 1),
			Widths: []int{1600},
			Height: 900,
			Checks: []h.Check{
				{Target: ".cui-console-player-full-screen-controls", Props: []string{"display"}, Golden: "not full screen, the page's controls hidden: the page shows them elsewhere (muse's under the player)"},
			},
		},
		{
			Name:   "player full screen controls touch",
			HTML:   h.Example("full-screen"),
			Touch:  true,
			Widths: []int{390},
			Checks: []h.Check{
				{Target: textBtn, Props: []string{"box-height"}, Golden: "a control of words 44px tall on a touch screen (the HIG's default, reference/touch/)"},
				{Target: ".cui-console-player-buttons", Props: []string{"box-right"}, Golden: "a narrow player's bar inside it: the controls of words give way, cut short with an ellipsis, the player's own buttons whole (its right edge the overlay's 30px in)"},
				{Target: textBtn, Props: []string{"text-overflow", "overflow-x"}, Golden: "cut short with an ellipsis where the bar runs out of room"},
			},
		},
		{
			Name:   "player stills hidden",
			HTML:   hidden(h.Example("stills")),
			Widths: []int{1600},
			Height: 900,
			Checks: []h.Check{
				{Target: line, Props: []string{"box-top", "box-height"}, Golden: "the controls hidden, the subtitles at the foot: 20px up, the overlay's room (the footer's row's foot)"},
				{Target: ".cui-console-player-subtitles", Props: []string{"opacity"}, Golden: "shown while the controls are hidden"},
				{Target: ".cui-console-player-footer", Props: []string{"opacity"}, Apple: h.AppleProbe(apple.TvPlayerPlaying, "hidden.footer")},
				{Target: p, Props: []string{"cursor"}, Apple: h.AppleProbe(apple.TvPlayerPlaying, "hidden.scrim")},
				{Target: ".cui-console-player-play", Props: []string{"cursor"}, Golden: "no pointer over a hidden button either, where its own pointer cursor would show (tv-player-playing:hidden.scrim, cursor none)"},
			},
		},
		{
			Name:   "player stills hidden focused",
			HTML:   hidden(h.Example("stills")),
			Widths: []int{1600},
			Height: 900,
			Press:  []string{"Tab"},
			Checks: []h.Check{
				{Target: line, Props: []string{"box-top"}, Golden: "the focus inside shows the controls, and the subtitles go back over the footer"},
				{Target: p, Props: []string{"cursor"}, Golden: "and the pointer with them"},
			},
		},
		{
			// @external: the probe puts the player's parent in full screen as muse's
			// script does, by a click (the Fullscreen API needs a user's gesture)
			Name:   "player full screen parent",
			HTML:   `<button type="button" onclick="document.getElementById('stage').requestFullscreen()">Enter</button><div id="stage">` + h.Example("stills") + `</div>`,
			Widths: []int{1600},
			Height: 576,
			Click:  []string{"button[onclick]"},
			Wait:   500 * time.Millisecond,
			Checks: []h.Check{
				{Target: "#stage", Props: []string{"position"}, Golden: "the parent in full screen, fixed by the browser's own :fullscreen rule: the check under it reads full screen"},
				{Target: p, Props: []string{"box-width", "box-height"}, Golden: "muse's (its [:fullscreen>&] width, moved into the kit): in full screen the player as large as the screen allows at 16:9, here min(100vw, 100vh * 16 / 9), 576px tall and 1024px wide in a 1600px by 576px screen"},
			},
		},
		{
			Name:   "player stills phone",
			HTML:   h.Example("stills"),
			Widths: []int{390},
			Checks: []h.Check{
				{Target: line, Props: []string{"font-size"}, Golden: "at the least 15px, where 2.4% of a narrow player would be smaller"},
			},
		},
		{
			Name:   "player menu open",
			HTML:   hidden(h.Example("stills")),
			Widths: []int{1600},
			Height: 900,
			Click:  []string{`[data-control="subtitles"]`},
			Checks: []h.Check{
				{Target: "#cui-console-subtitles", Props: []string{"display"}, Golden: "the subtitles button opens its menu, a console popover (decision 27)"},
				{Target: "#cui-console-subtitles", Props: []string{"box-top", "box-height"}, Golden: "above its button, its foot on the button's centre (the player's menus open upward, player menu languages)"},
				{Target: ".cui-console-player-footer", Props: []string{"opacity"}, Golden: "the controls hidden before, shown while its menu is open: a click leaves no keyboard focus, the open menu keeps them"},
			},
		},
		{
			Name:   "player menu languages",
			HTML:   languages(h.Example("stills"), 1),
			Widths: []int{1440},
			Height: 900,
			Click:  []string{`[data-control="subtitles"]`},
			Checks: []h.Check{
				{Target: ".cui-console-player-subtitles", Props: []string{"align-items"}, Golden: "a cue centered over the picture"},
				{Target: ".cui-console-player", Props: []string{"box-left", "box-top", "box-width", "box-height"}, Golden: "muse's player at 1440 by 900: 1104px wide at 24px, 66px down, beside its sidebar"},
				{Target: btn, Props: []string{"box-left", "box-top"}, Golden: "the subtitles button, 24px: its centre the menu's corner"},
				{Target: "#cui-console-subtitles", Props: []string{"box-x", "box-y", "box-width", "box-height"}, Golden: "inside the player (its parent: box-x and box-y from its corner): opening upward, its foot and its right edge on its button's centre, as tv.apple.com's (tv-player-menu:menu at 1258, 637, 185 by 142, its corner 1443, 779; tv-player:subtitles at 1434, 769, 20px: centred on 1444, 779); every choice in it, none scrolled"},
				{Target: "#cui-console-subtitles", Props: []string{"max-height"}, Golden: "as tall as the window leaves above its button, 8px from its top, not the toolkit's 435px: tv.apple.com's menu has none (tv-player-menu:menu, overflow visible)"},
			},
		},
		{
			Name:   "player menu voices and languages narrow",
			HTML:   languages(h.Example("stills"), 2),
			Widths: []int{735},
			Height: 900,
			Click:  []string{`[data-control="subtitles"]`},
			Checks: []h.Check{
				{Target: "#cui-console-subtitles", Props: []string{"width"}, Golden: "as wide as its columns, on a window too narrow to hold them in the browser's own width"},
			},
		},
		{
			Name:   "player menu voices and languages",
			HTML:   languages(h.Example("stills"), 2),
			Widths: []int{1440},
			Height: 900,
			Click:  []string{`[data-control="subtitles"]`},
			Checks: []h.Check{
				{Target: "#cui-console-subtitles", Props: []string{"box-x", "box-y", "box-width", "box-height"}, Golden: "one menu of two columns, muse's voices beside its subtitles (the owner, 2026-11-05: one button, one panel, as the Apple TV app's Audio and Subtitles; tv.apple.com's own menu is one column of groups, its audio a platform <select>, tv-player-menu): inside the player, opening upward, its foot and right edge on its button's centre as one column's"},
				{Target: `#cui-console-subtitles [data-column="voices"]`, Props: []string{"box-x", "box-width", "box-height"}, Golden: "the voices first, at the start, as wide as its longest name asks"},
				{Target: `#cui-console-subtitles [data-column="subtitles"]`, Props: []string{"box-x", "box-width", "box-height", "border-left-width", "border-left-style", "padding-left", "margin-left"}, Golden: "the subtitles beside them, after a rule in the popover's 8px gap, each column as tall as the menu"},
			},
		},
		{
			Name:   "player menu languages short",
			HTML:   languages(h.Example("stills"), 1),
			Widths: []int{1440},
			Height: 700,
			Click:  []string{`[data-control="subtitles"]`},
			Press:  slices.Repeat([]string{"Tab"}, 12),
			Checks: []h.Check{
				{Target: ".cui-console-player", Props: []string{"box-width", "box-height"}, Golden: "muse's player at 1440 by 700: 846px wide, too short for twelve choices above its button"},
				{Target: "#cui-console-subtitles", Props: []string{"box-x", "box-y", "box-width", "box-height", "overflow-y"}, Golden: "above the player's top where its choices need the room, its foot on its button's centre, every choice shown and none scrolled (the owner, 2026-11-14: 「A」; scrolled, WebKit 26.5 put it back to its top on every layout of the player)"},
				{Target: last, Props: []string{"focused", "hit"}, Golden: "the last choice reached: the focus Tabbed to it, the pointer finds it there"},
			},
		},
		{
			Name:   "player playing",
			HTML:   h.Example("playing"),
			Widths: []int{1600},
			Height: 900,
			Checks: []h.Check{
				{Target: ".cui-console-player-glyph-pause", Props: []string{"display", "box-width"}, Golden: "playing, the pause glyph in the play's place (tv-player-playing:pause, 50px)"},
				{Target: ".cui-console-player-glyph-play", Props: []string{"display"}, Golden: "playing, no play glyph (tv-player-playing:play.hidden)"},
				{Target: ".cui-console-player-play", Props: []string{"color", "opacity"}, Apple: h.AppleProbe(apple.TvPlayerPlaying, "pause")},
				{Target: ".cui-console-player-header", Props: []string{"opacity"}, Golden: "shown: data-controls is the page's to set"},
			},
		},
		{
			Name:   "player hidden",
			HTML:   h.Example("hidden"),
			Widths: []int{1600},
			Height: 900,
			Checks: []h.Check{
				{Target: ".cui-console-player-overlay", Pseudo: "::before", Props: []string{"opacity"}, Apple: h.AppleProbe(apple.TvPlayerPlaying, "hidden")},
				{Target: ".cui-console-player-header", Props: []string{"opacity"}, Apple: h.AppleProbe(apple.TvPlayerPlaying, "hidden.header")},
				{Target: ".cui-console-player-play", Props: []string{"opacity"}, Apple: h.AppleProbe(apple.TvPlayerPlaying, "hidden.body")},
				{Target: ".cui-console-player-footer", Props: []string{"opacity"}, Apple: h.AppleProbe(apple.TvPlayerPlaying, "hidden.footer")},
				{Target: ".cui-console-player-header", Props: []string{"transform", "transition-duration", "transition-timing-function"}, Golden: noSlide},
				{Target: ".cui-console-player-footer", Props: []string{"transform", "transition-duration", "transition-timing-function"}, Golden: noSlide},
			},
		},
		{
			Name:   "player hidden focused",
			HTML:   h.Example("hidden"),
			Widths: []int{1600},
			Height: 900,
			Press:  []string{"Tab"},
			Checks: []h.Check{
				{Target: volume, Props: []string{"height"}, Golden: "its 44px target, whatever the line it draws"},
				{Target: ".cui-console-player-volume", Props: []string{"align-items"}, Golden: "the speaker and its range on one center line"},
				{Target: volume, Props: []string{"focused"}, Golden: "Tab reaches the hidden controls"},
				{Target: volume, Props: []string{"box-width", "opacity"}, Golden: "focused from the keyboard, the range opens: a range that took the focus closed would be adjusted unseen"},
				{Target: ".cui-console-player-footer", Props: []string{"opacity"}, Golden: "shown while the focus is inside, for the keyboard (not read on tv.apple.com, whose script shows them on a key)"},
			},
		},
		{
			Name:   "player hidden clicked",
			HTML:   h.Example("hidden"),
			Widths: []int{1600},
			Height: 900,
			Click:  []string{".cui-console-player-play"},
			Checks: []h.Check{
				{Target: ".cui-console-player-play", Props: []string{"focused"}, Golden: "the click leaves the focus on Play, as Chrome does"},
				{Target: ".cui-console-player-overlay", Pseudo: "::before", Props: []string{"opacity"}, Golden: "a click's focus keeps no gradient: one click on Play kept it over the picture for the whole film (muse, the owner 2026-11-05: 「似乎太頻繁了，而不是只有滑鼠hover」)"},
				{Target: ".cui-console-player-footer", Props: []string{"opacity"}, Golden: "nor the controls: only the keyboard's focus keeps them (player hidden focused)"},
			},
		},
		{
			Name:   "player dark",
			HTML:   `<div class="cui-dark">` + h.Example("paused") + `</div>`,
			Widths: []int{1200},
			Height: 900,
			Checks: []h.Check{
				{Target: p, Props: []string{"background-color"}, Apple: h.AppleProbe(apple.TvPlayer, "dialog")},
				{Target: btn, Props: []string{"color"}, Apple: h.AppleProbe(apple.TvPlayer, "subtitles")},
				{Target: ".cui-console-player-time", Props: []string{"color"}, Apple: h.AppleProbe(apple.TvPlayer, "elapsed")},
			},
		},
		{
			Name:   "player touch",
			HTML:   h.Example("paused"),
			Touch:  true,
			Widths: []int{390},
			Checks: []h.Check{
				{Target: btn, Props: []string{"box-width", "box-height"}, Golden: "a button 44px each way at least on a touch screen (the HIG's default, reference/touch/)"},
				{Target: scrub, Props: []string{"box-height"}, Golden: "a range 44px tall on a touch screen as on a desktop, its track and thumb the same"},
				{Target: scrub, Props: []string{"margin-top", "margin-bottom"}, Golden: "on a touch screen its row its whole target: the buttons above are 44px too, and a target reaching up into theirs would leave each short (the touch audit read 39px)"},
				{Target: ".cui-console-player-progress", Props: []string{"grid-template-rows"}, Golden: "the range's row 44px on a touch screen, over the times' 14px"},
			},
		},
	})
}
