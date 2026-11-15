//go:build browser

package globalnav

import (
	"slices"
	"strings"
	"testing"

	h "github.com/at-least/ciderui/internal/harness"
	"github.com/at-least/ciderui/internal/reference"
	apple "github.com/at-least/ciderui/reference/apple.com/results"
)

const (
	list     = ".cui-globalnav-list"
	wordmark = ".cui-globalnav-wordmark"
	button   = ".cui-globalnav-open"
	closeX   = ".cui-globalnav-close"
	link     = list + " > li:nth-child(3) > .cui-globalnav-link"
	link2    = list + " > li:nth-child(4) > .cui-globalnav-link"
	name     = "the site's name in the kit's type, cui-body-reduced's at 600: apple.com's bar draws its logo, no name (the one site's bar wrote this type)"
	clear    = "an in-page target (the skip link's #main) lands below the bar: the page's scroll padding is the bar's height"
	rows     = "shown with the sheet, where apple.com's fade in 240ms from 200ms (home:menu.row, home:menu.open's animations): a phone's menu is read at once (DESIGN.md decision 25)"
	post     = list + " > li:nth-child(6) > .cui-globalnav-link"
	asLink   = "a POST item (sign out) is a <button>, drawn as the link: each value equal to the link's beside it"
)

var (
	nav     = h.Example("globalnav")
	dark    = strings.Replace(nav, `class="cui-globalnav"`, `class="cui-globalnav cui-dark"`, 1)
	desktop = []int{1600, 1200, 900}
	flex    = []string{"display", "flex-direction", "flex-wrap", "align-items", "justify-content"}
	inline  = []string{"padding-left", "padding-right"}
	bars    = []string{button + " path:first-child", button + " path:last-child"}
	x       = []string{closeX + " path:first-child", closeX + " path:last-child"}
	// @external: apple.com reads its menu's button at 900 and 390px alone
	triggerAt = map[int]int{1600: 900, 1200: 900}
	// @external: apple.com's phone bar draws its ground on its content, ciderui's on the
	// nav at every width: the same ground, apple.com's at 900px
	groundAt = map[int]int{390: 900}
	typeP    = []string{"font-size", "line-height", "font-weight", "letter-spacing"}
	// @trap: a page longer than the window, so the bar has somewhere to stay
	scrolled = nav + `<div style="height: 2000px"></div><p id="target">Target</p><div style="height: 2000px"></div><a class="to" href="#target">To</a>`
	stays    = h.Check{Target: ".cui-globalnav", Props: []string{"position", "top", "z-index", "box-top"}, Golden: "the bar stays at the top as the page scrolls, over it (DESIGN.md rule 6's exception)"}
	postNav  = h.Example("post")
	// @external: what a <button> brings of its own, a link not
	buttonsOwn = []string{"font-family", "font-style", "text-align", "vertical-align", "border-top-width", "border-left-width", "background-color", "cursor", "margin-top", "margin-left", "box-sizing", "box-top", "opacity"}
	menuOwn    = h.Props(slices.DeleteFunc(slices.Clone(buttonsOwn), func(p string) bool { return p == "box-top" }), []string{"box-left", "box-width"})
)

func appleAt(page reference.Page, probe string, at map[int]int) *h.Apple {
	return &h.Apple{Page: page, Probe: probe, State: "", At: at, As: nil}
}

func appleState(page reference.Page, probe, state string, at map[int]int) *h.Apple {
	return &h.Apple{Page: page, Probe: probe, State: state, At: at, As: nil}
}

func TestCases(t *testing.T) {
	h.Run(t, []h.Case{
		{
			Name:   "post-bar",
			HTML:   postNav,
			Widths: desktop,
			Checks: []h.Check{
				{Target: post, Props: []string{"font-size", "line-height", "font-weight", "color"}, Apple: h.AppleProbe(apple.Home, "nav.link")},
				{Target: post, State: []string{"hover"}, Props: []string{"color"}, Apple: h.AppleProbe(apple.Home, "nav.link", "hover")},
				{Target: post, Props: h.Props(flex, inline, []string{"box-height", "text-decoration-line"}), Apple: h.AppleProbe(apple.Home, "nav.link-box")},
				{Target: link, Props: buttonsOwn, Golden: asLink},
				{Target: post, Props: buttonsOwn, Golden: asLink},
			},
		},
		{
			Name:          "post-menu",
			HTML:          postNav,
			Click:         []string{button},
			Widths:        []int{390},
			Height:        900,
			ReducedMotion: true,
			Checks: []h.Check{
				{Target: post, Props: h.Props([]string{"font-size", "line-height", "font-weight", "color", "display"}, h.Padding, []string{"box-height"}), Apple: h.AppleProbe(apple.Home, "menu.link")},
				{Target: post, State: []string{"hover"}, Props: []string{"color"}, Apple: h.AppleProbe(apple.Home, "menu.link", "hover")},
				{Target: link2, Props: menuOwn, Golden: asLink},
				{Target: post, Props: menuOwn, Golden: asLink},
			},
		},
		{
			Name:   "bar",
			HTML:   nav,
			Widths: desktop,
			Checks: []h.Check{
				{Target: ".cui-globalnav", Props: []string{"box-height", "background-color", "backdrop-filter"}, Apple: h.AppleProbe(apple.Home, "nav")},
				{Target: ".cui-globalnav-content", Props: h.Props([]string{"max-width"}, inline), Apple: h.AppleProbe(apple.Home, "nav.content")},
				{Target: list, Props: h.Props(flex, []string{"margin-left", "margin-right", "box-height"}), Apple: h.AppleProbe(apple.Home, "nav.list")},
				{Target: link, Props: []string{"font-size", "line-height", "font-weight", "color"}, Apple: h.AppleProbe(apple.Home, "nav.link")},
				{Target: link, State: []string{"hover"}, Props: []string{"color"}, Apple: h.AppleProbe(apple.Home, "nav.link", "hover")},
				{Target: link, Props: h.Props(flex, inline, []string{"box-height", "text-decoration-line"}), Apple: h.AppleProbe(apple.Home, "nav.link-box")},
				{Target: wordmark, Props: h.Props(flex, inline, []string{"box-height", "box-left", "color"}), Apple: h.AppleProbe(apple.Home, "nav.logo")},
				{Target: wordmark, State: []string{"hover"}, Props: []string{"color"}, Apple: h.AppleProbe(apple.Home, "nav.logo", "hover")},
				{Target: wordmark, Props: typeP, Golden: name},
				{Target: button, Props: []string{"box-height"}, Apple: appleAt(apple.Home, "nav.menutrigger", triggerAt)},
			},
		},
		{
			Name:   "phone",
			HTML:   nav,
			Widths: []int{390},
			Checks: []h.Check{
				{Target: ".cui-globalnav", Props: []string{"box-height"}, Apple: h.AppleProbe(apple.Home, "nav")},
				{Target: ".cui-globalnav", Props: []string{"background-color", "backdrop-filter"}, Apple: appleAt(apple.Home, "nav", groundAt)},
				{Target: ".cui-globalnav-content", Props: h.Props([]string{"max-width"}, inline), Apple: h.AppleProbe(apple.Home, "nav.content")},
				{Target: list, Props: []string{"box-height"}, Apple: h.AppleProbe(apple.Home, "nav.list")},
				{Target: wordmark, Props: h.Props(flex, inline, []string{"box-height", "box-left", "color"}), Apple: h.AppleProbe(apple.Home, "nav.logo")},
				{Target: wordmark, Props: typeP, Golden: name},
				{Target: button, Props: h.Props(flex, []string{"box-height", "box-left", "color", "opacity", "cursor"}), Apple: h.AppleProbe(apple.Home, "nav.menutrigger")},
				{Target: button, State: []string{"hover"}, Props: []string{"opacity"}, Apple: h.AppleProbe(apple.Home, "nav.menutrigger", "hover")},
				{Target: bars[0], Props: []string{"stroke-width"}, Apple: h.AppleProbe(apple.Home, "nav.menutrigger-bread")},
				{Target: list + " > li:nth-child(2)", Props: []string{"box-height"}, Golden: "the close button is the opened menu's"},
				{Target: link, Props: []string{"box-height"}, Golden: "to 833px the links are the menu's rows, not the bar's: apple.com's bar holds the logo, the search, the bag and the button"},
			},
		},
		{
			Name:   "select",
			HTML:   nav,
			Widths: []int{1600, 390},
			Checks: []h.Check{
				{Target: ".cui-globalnav", Props: []string{"user-select"}, Apple: h.AppleProbe(apple.NavSelect, "nav")},
				{Target: list, Props: []string{"user-select"}, Apple: h.AppleProbe(apple.NavSelect, "list")},
				{Target: wordmark, Props: []string{"user-select"}, Apple: h.AppleProbe(apple.NavSelect, "logo")},
				{Target: button, Props: []string{"user-select"}, Apple: appleAt(apple.NavSelect, "menutrigger", map[int]int{1600: 390})},
			},
		},
		{
			Name:          "menu-zh-hant",
			HTML:          `<div lang="zh-Hant">` + nav + `</div>`,
			Click:         []string{button},
			Widths:        []int{390},
			Height:        900,
			ReducedMotion: true,
			Checks:        []h.Check{{Target: link2, Props: []string{"line-height", "letter-spacing"}, Apple: h.AppleProbe(apple.HomeTw, "menu.link")}},
		},
		{
			Name:          "nav boundary",
			HTML:          nav,
			Widths:        []int{833, 834},
			ReducedMotion: true,
			Checks: []h.Check{
				{Target: link, Props: []string{"box-height"}, Golden: "the nav's own step at 834px, apple.com's measured width (DESIGN.md): to 833px a link is the menu's row, hidden in the bar; from 834px the bar's row, 44px"},
			},
		},
		{
			Name:          "menu",
			HTML:          nav,
			Click:         []string{button},
			Widths:        []int{390},
			Height:        900,
			ReducedMotion: true,
			Checks: slices.Concat([]h.Check{
				{Target: ".cui-globalnav", Props: []string{"box-height"}, Apple: h.AppleProbe(apple.Home, "menu.open")},
				{Target: list, Props: []string{"background-color", "box-top", "box-left", "box-height"}, Apple: h.AppleProbe(apple.Home, "menu.content")},
				{Target: list, Props: []string{"padding-bottom"}, Apple: h.AppleProbe(apple.Home, "menu.list")},
				{Target: link, Props: []string{"padding-top", "box-top", "box-height"}, Apple: h.AppleProbe(apple.Home, "menu.link-first")},
				{Target: link2, Props: h.Props([]string{"font-size", "line-height", "font-weight", "color", "display"}, h.Padding, []string{"box-top", "box-height"}), Apple: h.AppleProbe(apple.Home, "menu.link")},
				{Target: link2, State: []string{"hover"}, Props: []string{"color"}, Apple: h.AppleProbe(apple.Home, "menu.link", "hover")},
				// @external: apple.com fades the links themselves; ciderui their items
				{Target: list + " > li:nth-child(1)", Props: []string{"opacity", "visibility"}, Apple: h.AppleProbe(apple.Home, "menu.logo")},
				{Target: closeX, Props: []string{"color", "opacity", "box-height", "box-left", "box-top"}, Apple: h.AppleProbe(apple.Home, "menu.trigger")},
				{Target: closeX, Props: []string{"box-width"}, Golden: "apple.com's 48px square, where the menu's button is (home:nav.menutrigger at 390px)"},
			}, xLines(), []h.Check{
				{Target: "html", Props: []string{"overflow-y"}, Apple: h.AppleProbe(apple.Home, "menu.root")},
			}),
		},
		{
			Name:   "menu-motion",
			HTML:   nav,
			Click:  []string{button},
			Widths: []int{390},
			Height: 900,
			Checks: []h.Check{
				{Target: list + " > li:nth-child(1)", Props: h.Props(h.Motion, []string{"transition-delay"}), Apple: h.AppleProbe(apple.Home, "menu.logo")},
				{Target: list + " > li:nth-child(3)", Props: []string{"transition-duration", "animation-name", "animation-duration", "animation-delay"}, Golden: rows},
				{Target: list, Props: h.Motion, Golden: "apple.com's ground and height in 450ms (home:menu.open's animations; its computed transition names its ground alone, its height's declared in a state between)"},
				{Target: x[0], Props: h.Motion, Golden: "apple.com's lines turn in 0.24s (their SVG animation, reference/icons/)"},
			},
		},
		{
			Name:          "menu-closed",
			HTML:          nav,
			Click:         []string{button, closeX},
			Widths:        []int{390},
			Height:        900,
			ReducedMotion: true,
			Checks: []h.Check{
				{Target: list, Props: []string{"box-height"}, Apple: h.AppleProbe(apple.Home, "nav.list")},
				{Target: list, Props: []string{"position", "background-color"}, Golden: "closed, the bar's row again (reference/popover/)"},
				{Target: "html", Props: []string{"overflow-y"}, Golden: "the page scrolls again"},
				{Target: link, Props: []string{"box-height"}, Golden: "the links are the menu's again"},
				{Target: button, Props: []string{"focused"}, Golden: "the X closes it, the focus back on the menu's button (reference/popover/)"},
			},
		},
		{
			Name:       "menu-closed-at-once",
			HTML:       nav,
			Click:      []string{button, closeX},
			Widths:     []int{390},
			Height:     900,
			FirstFrame: true,
			Checks: []h.Check{
				{Target: list, Props: []string{"position", "box-height"}, Golden: "closed at once, out of the top layer (apple.com's runs its motion back)"},
			},
		},
		{
			Name:          "menu-tab-out",
			HTML:          nav + `<a class="after" href="#after">After</a>`,
			Click:         []string{button},
			Widths:        []int{390},
			Height:        900,
			Press:         []string{"Tab", "Tab", "Tab", "Tab", "Tab"},
			ReducedMotion: true,
			Checks: []h.Check{
				{Target: ".after", Props: []string{"focused"}, Golden: "Tab past its last link leaves it for the page: no script holds the focus in it (apple.com's does)"},
			},
		},
		{
			Name:          "menu-widened",
			HTML:          nav,
			Click:         []string{button},
			Widths:        []int{390, 900, 1600},
			ReducedMotion: true,
			Checks: []h.Check{
				{Target: list, Props: []string{"width", "left", "right"}, Golden: "the menu fills the window: the browser places a popover at its content's width"},
				{Target: link2, Props: []string{"font-size", "box-top"}, Golden: "opened, the menu is a sheet at any width: its rows"},
				{Target: closeX, Props: []string{"box-height", "opacity"}, Golden: "opened, the menu is a sheet at any width: its close button"},
			},
		},
		{
			Name:       "menu-start",
			HTML:       nav,
			Click:      []string{button},
			Widths:     []int{390},
			Height:     900,
			FirstFrame: true,
			Checks: []h.Check{
				{Target: list + " > li:nth-child(3)", Props: []string{"opacity", "transform"}, Golden: rows},
				{Target: list, Props: []string{"box-height", "background-color"}, Golden: "a departure at the reveal's start: apple.com's sheet grows from the bar's 48px and its white ground (home:menu.open's animations); ciderui's grows from its content's height, transparent — this Chromium ignores a popover's @starting-style where the element renders closed (the bar's logo lives in its list), and the bar is not restructured for the first frames of a reveal (measured in the browser, 2026-11-07)"},
			},
		},
		{
			Name:          "menu-focus",
			HTML:          nav,
			Click:         []string{button},
			Widths:        []int{390},
			Height:        900,
			ReducedMotion: true,
			Checks: []h.Check{
				{Target: closeX, Props: []string{"focused"}, Golden: "opened, the focus is on the X (autofocus; reference/popover/): the menu's button beside the list, as a popover opened by a button inside it holds the focus on that button"},
			},
		},
		{
			Name:          "menu-tab",
			HTML:          nav,
			Click:         []string{button},
			Widths:        []int{390},
			Height:        900,
			Press:         []string{"Tab"},
			ReducedMotion: true,
			Checks: []h.Check{
				{Target: link, Props: []string{"focused"}, Golden: "Tab goes from the X to the first link, not back to a button (reference/popover/): the X second in the list, after the wordmark"},
			},
		},
		{
			Name:          "menu-escape",
			HTML:          nav,
			Click:         []string{button},
			Widths:        []int{390},
			Height:        900,
			Press:         []string{"Escape"},
			ReducedMotion: true,
			Checks: []h.Check{
				{Target: list, Props: []string{"position"}, Golden: "Escape closes it"},
				{Target: button, Props: []string{"focused"}, Golden: "the focus back on the menu's button (reference/popover/)"},
			},
		},
		{
			Name: "dark",
			HTML: dark,
			Checks: []h.Check{
				{Target: ".cui-globalnav", Props: []string{"background-color", "backdrop-filter"}, Apple: appleAt(apple.Newsroom, "nav", groundAt)},
				{Target: button, Props: []string{"color", "opacity"}, Apple: appleAt(apple.Newsroom, "nav.menutrigger", triggerAt)},
				{Target: button, State: []string{"hover"}, Props: []string{"opacity"}, Apple: appleState(apple.Newsroom, "nav.menutrigger", "hover", triggerAt)},
			},
		},
		{
			Name:   "dark-links",
			HTML:   dark,
			Widths: desktop,
			Checks: []h.Check{
				{Target: link, Props: []string{"color"}, Apple: h.AppleProbe(apple.Newsroom, "nav.link")},
				{Target: link, State: []string{"hover"}, Props: []string{"color"}, Apple: h.AppleProbe(apple.Newsroom, "nav.link", "hover")},
			},
		},
		{
			Name:          "dark-menu",
			HTML:          dark,
			Click:         []string{button},
			Widths:        []int{390},
			Height:        900,
			ReducedMotion: true,
			Checks: []h.Check{
				{Target: list, Props: []string{"background-color"}, Apple: h.AppleProbe(apple.Newsroom, "menu.content")},
				{Target: link2, Props: []string{"color"}, Apple: h.AppleProbe(apple.Newsroom, "menu.link")},
				{Target: link2, State: []string{"hover"}, Props: []string{"color"}, Apple: h.AppleProbe(apple.Newsroom, "menu.link", "hover")},
				{Target: closeX, Props: []string{"color", "opacity"}, Apple: h.AppleProbe(apple.Newsroom, "menu.trigger")},
			},
		},
		{
			Name:   "sticky-phone",
			HTML:   scrolled,
			Widths: []int{390},
			Click:  []string{".to"},
			Checks: []h.Check{stays, {Target: "#target", Props: []string{"box-top"}, Golden: clear}},
		},
		{
			Name:   "sticky-bar",
			HTML:   scrolled,
			Widths: []int{1200},
			Click:  []string{".to"},
			Checks: []h.Check{stays, {Target: "#target", Props: []string{"box-top"}, Golden: clear}},
		},
	})
}

func xLines() []h.Check {
	var checks []h.Check
	for _, target := range x {
		checks = append(checks, h.Check{Target: target, Props: []string{"box-x", "box-y", "box-width", "box-height"}, Golden: "apple.com's X: its lines from (3.5, 3.5) to (15, 15) and from (3.5, 15) to (15, 3.5), the end of its animation (reference/icons/)"})
	}
	return checks
}
