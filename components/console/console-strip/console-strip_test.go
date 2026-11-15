//go:build browser

package consolestrip

import (
	"fmt"
	"testing"

	h "github.com/at-least/ciderui/internal/harness"
)

const (
	strip          = ".cui-console-strip"
	group          = ".cui-console-strip-group:first-child"
	header         = group + " > .cui-console-strip-header"
	title          = header + " > .cui-console-strip-title"
	current        = `.cui-console-strip-item[aria-current="true"]`
	waiting        = `.cui-console-strip-mark[data-state="waiting"]`
	working        = `.cui-console-strip-mark[data-state="working"]`
	currentWorking = current + " > " + working
	marks          = "two marks only (the owner, 2026-11-05, for muse's storyboard and assets strip): who the shot waits on, the reader or someone at work on it; nothing, no mark"
)

func item(group, n int) string {
	return fmt.Sprintf(".cui-console-strip-group:nth-child(%d) > .cui-console-strip-item:nth-of-type(%d)", group, n)
}

var (
	first   = item(1, 1)
	img     = first + " > img"
	caption = first + " > .cui-console-strip-caption"
	tip     = first + " > .cui-console-strip-tip"
	// @external: tv.apple.com's player's row was read at 1600 and 1200px
	at = map[int]int{390: 1200}
)

// @invariant: one scene of one shot, its name longer than the shot is wide, and no caption
func lone(title string) string {
	return `<div class="cui-console-page"><nav class="cui-console-strip" aria-label="Shots"><div class="cui-console-strip-group"><p class="cui-console-strip-header"><span class="cui-console-strip-title">` + title + `</span></p><a class="cui-console-strip-item" href="#a"><img src="data:image/gif;base64,R0lGODlhAQABAAAAACw=" alt="Shot 1"></a></div></nav></div>`
}

func TestCases(t *testing.T) {
	h.Run(t, []h.Case{
		{
			Name:   "strip",
			HTML:   h.Example("strip"),
			Widths: []int{1200, 390},
			Checks: []h.Check{
				{Target: tip, Props: []string{"font-size", "line-height"}, Golden: "the small type on a phone too, where the console page's own is larger"},
				{Target: strip, Props: []string{"display", "grid-auto-flow", "overflow-x", "scroll-snap-type", "padding-top", "padding-left", "margin-top", "margin-bottom"}, Apple: &h.Apple{Page: "tv-player-playlist", Probe: "row", At: at}},
				{Target: first, Props: []string{"scroll-snap-align"}, Apple: &h.Apple{Page: "tv-player-playlist", Probe: "item", At: at}},
				{Target: img, Props: []string{"border-top-left-radius", "border-top-width", "border-top-style", "box-shadow", "cursor"}, ASC: h.ASCProbe("media", "thumb.img")},
				{Target: img, Props: []string{"box-width", "box-height", "object-fit"}, Golden: "a 16:9 shot 160px wide, cropped to fill: a storyboard's frames are a film's shape (App Store Connect's are a device's), more of them in view than tv.apple.com's 288px trailers (tv-player-playlist lockup, five at 1600px)"},
				{Target: img, Props: []string{"background-color"}, Golden: "the console card's gray while it loads, as a console thumbnail's"},
				{Target: group, Props: []string{"column-gap"}, Golden: "a scene's shots 12px apart: App Store Connect's 6px either side of each (media.json thumb.item), where tv.apple.com's trailers are 20px apart"},
				{Target: strip, Props: []string{"column-gap", "overflow-y"}, Golden: "one row, scrolled sideways only as tv.apple.com's (its grid of columns); a scene's shots twice as close to each other as to the next scene's, so they read as its group: neither site groups its row"},
				{Target: title, Props: []string{"font-size", "line-height", "font-weight", "color", "text-overflow", "white-space"}, Golden: "the scene's name above its shots in 11px on 13px, the secondary label, cut to its group's width; neither site labels a group of its row"},
				{Target: group + " > .cui-console-strip-item:first-of-type > img", Props: []string{"box-top", "box-left"}, Golden: "the first picture 18px under its group's top, its name's 13px line and 5px (a thumbnail's caption's 5px from it), 4px in from the page's content edge, the items' room for their ring"},
				{Target: title, Props: []string{"box-top", "box-left", "padding-left", "position", "left"}, Golden: "the scene's name over its first picture's left edge, its text 4px into its box as the picture is into its item; sticky at the scroller's edge, so the name of the scene scrolled half out stays in view (muse's storyboard, the owner, 2026-11-05)"},
				{Target: header, Props: []string{"position", "box-top", "box-left", "box-width", "overflow-x"}, Golden: "the name's track: out of the flow so a long name widens no scene, the group's whole width (528px: three 168px shots, each its picture and 4px of ring room either side, 12px apart), not clipped (a clipping box would be the sticky name's scroller)"},
				{Target: caption, Props: []string{"font-size", "line-height", "font-weight", "color", "text-overflow", "white-space"}, Golden: "the shot's code under it in 11px on 13px, the secondary label, cut to the picture's width: muse's storyboard names each shot under it, not over it (the owner, 2026-11-05)"},
				{Target: caption, Props: []string{"box-top", "box-left", "box-width"}, Golden: "the caption 5px under its picture (the picture's foot at 108px: 18 and 90), as a console thumbnail's caption, outside the item's ring, at the picture's left and as wide"},
				{Target: group, Props: []string{"padding-bottom"}, Golden: "a scene with captions keeps their room under its shots: 1px past the ring and the 13px line, as its name's over them"},
				{Target: currentWorking, Props: []string{"box-width", "top", "right", "background-color"}, Golden: "the current shot's dot drawn as any shot's, inside its ring"},
				{Target: first, Props: []string{"border-top-width", "border-top-color", "padding-top", "border-top-left-radius"}, Golden: "a 2px edge 2px out from the picture, clear at rest, so the current shot's ring takes no room of its own"},
				{Target: current, Props: []string{"border-top-color", "border-top-width"}, Golden: "the current shot ringed in the brand accent, 2px out from its picture (Keynote's and Photos' chosen thumbnail): neither site marks one current (tv-player-playlist, none marked as playing); a ring, not a fill, so its picture shows whole"},
				{Target: waiting, Props: []string{"box-width", "box-height", "top", "right", "border-top-width", "border-top-color", "background-color", "border-top-left-radius"}, Golden: "waiting for the reader: a dot in console-status's warning mark's color (its 12px), ringed 2px in the page's ground so it shows on any picture, 8px in from the item (6px from the picture). " + marks},
				{Target: working, Props: []string{"box-width", "box-height", "top", "right", "border-top-width", "border-top-color", "background-color"}, Golden: "someone at work on it: the same dot in console-status's neutral mark's color (the secondary label). " + marks},
				{Target: tip, Props: []string{"visibility", "opacity", "position"}, Golden: "its line hidden until the pointer is on it"},
				{Target: current, Pseudo: "::after", Props: []string{"content"}, Golden: "the current shot draws no progress line, so its picture shows whole: the player's scrubber says how far it has played (muse's owner, 2026-11-07)"},
			},
		},
		{
			Name:   "strip hovered",
			HTML:   h.Example("strip"),
			Widths: []int{1200},
			Checks: []h.Check{
				{Target: tip, On: first, State: []string{"focus-visible"}, Props: []string{"visibility", "opacity"}, Golden: "the keyboard reaches what the pointer shows"},
				{Target: img, Props: []string{"border-top-left-radius", "box-shadow"}, State: []string{"hover"}, On: first, ASC: h.ASCProbe("media", "thumb.img")},
				{Target: first, Props: []string{"border-top-color"}, State: []string{"hover"}, Golden: "unchanged under the pointer, as App Store Connect's thumbnail (media: nothing changed with the pointer on one)"},
			},
		},
		{
			Name:   "strip tip",
			HTML:   `<div style="padding-top:120px">` + h.Example("strip") + `</div>`,
			Widths: []int{1200},
			Hover:  []string{first},
			Checks: []h.Check{
				{Target: tip, Props: []string{"visibility", "opacity", "position", "font-size", "line-height", "color", "background-color", "box-shadow", "padding-top", "padding-left", "border-top-left-radius", "max-width"}, Golden: "the shot's first spoken line over it while the pointer is on it (the owner, 2026-11-05, muse's storyboard; its assets strip, the asset's description), after half a second so a pointer passing along the row raises none: the console popover's ground, shadow and 13px on 18px (decision 27, its floating surface), 6px by 10px, its narrow 14px corners, 280px at most; neither site has a hover tip on its row"},
				{Target: tip, Props: []string{"box-left", "box-top", "box-width", "box-height"}, Golden: "centred over the item, 8px above it (the popover's gap), past the scroller's clip (fixed to the item by anchor positioning); at the window's edge it keeps inside"},
				{Target: first, Props: []string{"box-top", "box-left", "box-width"}, Golden: "the item it is over"},
				{Target: tip, Props: []string{"transition-property", "transition-duration", "transition-delay"}, Golden: "half a second's wait, then shown at once: the delay alone is the device (a pointer passing along the row raises nothing), a fade on top of it an entrance flourish decision 25 leaves out; leaving hides it at once"},
			},
		},
		{
			Name:   "strip dark",
			HTML:   `<div class="cui-dark">` + h.Example("strip") + `</div>`,
			Widths: []int{1200},
			Checks: []h.Check{
				{Target: title, Props: []string{"color"}, Golden: "the dark side derived (tokens.css, decision 18)"},
				{Target: caption, Props: []string{"color"}, Golden: "the dark side derived (tokens.css, decision 18)"},
				{Target: current, Props: []string{"border-top-color"}, Golden: "the dark side derived (tokens.css, decision 18)"},
				{Target: waiting, Props: []string{"background-color", "border-top-color"}, Golden: "the dark side derived (tokens.css, decision 18), the ring the dark page's ground"},
				{Target: working, Props: []string{"background-color"}, Golden: "the dark side derived (tokens.css, decision 18)"},
			},
		},
		{
			Name:         "strip forced colors",
			HTML:         h.Example("strip"),
			Widths:       []int{1200},
			ForcedColors: true,
			Checks: []h.Check{
				{Target: current, Props: []string{"border-top-color"}, Golden: "the current shot's ring the theme's Highlight, as a current tab's"},
				{Target: first, Props: []string{"border-top-color"}, Golden: "the others' clear edge the ground's color: forced colors drew it in the text's, every shot ringed as the current one"},
				{Target: waiting, Props: []string{"background-color", "border-top-color"}, Golden: "the waiting dot the forced text color on the ground's ring, as console-status's mark"},
				{Target: working, Props: []string{"background-color", "border-top-color"}, Golden: "the working dot the theme's GrayText, so the two keep apart where their colors are forced"},
			},
		},
		{
			Name:   "strip snap",
			HTML:   h.Example("strip"),
			Widths: []int{1200, 390},
			Scroll: &h.Scroll{Selector: strip, X: 300},
			Checks: []h.Check{
				{Target: item(1, 3), Props: []string{"box-left"}, Golden: "between two shots the row comes to rest with one at its first shot's place (scroll-snap-type x mandatory and each shot's start, tv.apple.com's row)"},
			},
		},
		{
			Name:   "strip end",
			HTML:   h.Example("strip"),
			Widths: []int{1200, 390},
			Scroll: &h.Scroll{Selector: strip, X: 99999},
			Checks: []h.Check{
				{Target: item(3, 3), Props: []string{"box-right"}, Golden: "at its end the last shot as far in from the right as the first from the left"},
			},
		},
		{
			Name:   "strip pinned",
			HTML:   h.Example("strip"),
			Widths: []int{1200, 390},
			Scroll: &h.Scroll{Selector: strip, X: 400},
			Checks: []h.Check{
				{Target: strip, Props: []string{"box-left"}, Golden: "the scroller 15px out of the content edge, its room for the rings (tv.apple.com's row)"},
				{Target: group, Props: []string{"box-left"}, Golden: "scrolled to the first scene's third shot (it snaps there from 400px, 2 × 180px): the scene's start 360px out of view, its end still in it"},
				{Target: title, Props: []string{"box-left"}, Golden: "the scene's name stays while its scene is still in view (muse's storyboard, the owner, 2026-11-05), at the content edge, 15px in from the scroller's (Chromium sticks it inside the scroller's padding, its room for the rings), so its text is where the first picture starts at rest; left in the group, it would have scrolled out with the scene's start"},
				{Target: ".cui-console-strip-group:nth-child(2) > .cui-console-strip-header > .cui-console-strip-title", Props: []string{"box-left"}, Golden: "the next scene's name over its first shot, where its scene starts (24px past the first scene's end)"},
			},
		},
		{
			Name:   "strip long title",
			HTML:   lone("A scene whose name is longer than its one shot is wide"),
			Widths: []int{1200},
			Checks: []h.Check{
				{Target: title, Props: []string{"box-width", "text-overflow"}, Golden: "a name longer than its scene cut to the scene's width (one shot, 168px), widening nothing"},
				{Target: ".cui-console-strip-group", Props: []string{"box-width", "padding-bottom"}, Golden: "the scene its shot's width, whatever its name; no caption, so no room under it"},
			},
		},
		{
			Name:   "strip assets",
			HTML:   h.Example("assets"),
			Widths: []int{1200, 390},
			Checks: []h.Check{
				{Target: first + " > .cui-console-strip-detail", Props: []string{"font-size", "line-height", "color", "text-overflow", "white-space", "box-top", "box-left", "box-width"}, Golden: "a second line under the caption, in its type, the next line down (126px: the caption's 113 and its 13px line), cut to the picture's width: muse's assets page says under each asset's name how many shots use it (the owner, 2026-11-05)"},
				{Target: group, Props: []string{"padding-bottom"}, Golden: "a strip with a detail line keeps both lines' room under every group's shots: 1px past the ring and two 13px lines"},
				{Target: first + " > img", Props: []string{"object-position", "object-fit", "box-width", "box-height"}, Golden: "a character sheet cropped to its left edge, where its face is (muse's assets page), at the row's 16:9 and not muse's square: a square would narrow its item, its caption and its marks' places in a row of 16:9 shots"},
				{Target: item(2, 1) + " > img", Props: []string{"object-position"}, Golden: "any other picture cropped round its middle"},
				{Target: item(1, 2) + " > img", Props: []string{"opacity"}, Golden: "an asset no shot uses: its picture dimmed to half, muse's; not its lines, which its words already say (dimmed, the caption fell to 2:1 on the page, the surfaces' contrast check)"},
				{Target: item(1, 2) + " > .cui-console-strip-caption", Props: []string{"opacity"}, Golden: "its caption not dimmed"},
				{Target: first + " > img", Props: []string{"opacity"}, Golden: "one in use not dimmed"},
			},
		},
		{
			Name:   "strip assets hovered",
			HTML:   h.Example("assets"),
			Widths: []int{1200},
			Checks: []h.Check{
				{Target: item(1, 2) + " > img", Props: []string{"opacity"}, State: []string{"hover"}, On: item(1, 2), Golden: "a dimmed asset under the pointer at 0.8, muse's: brighter, so it reads as a link, and still apart from those in use"},
			},
		},
	})
}
