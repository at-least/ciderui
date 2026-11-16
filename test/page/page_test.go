//go:build browser

package page

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"testing"

	"github.com/at-least/ciderui/internal/browser"
	"github.com/at-least/ciderui/internal/build"
	"github.com/at-least/ciderui/internal/check"
	"github.com/at-least/ciderui/internal/reference"
	apple "github.com/at-least/ciderui/reference/apple.com/results"
	"github.com/mxschmitt/playwright-go"
)

const (
	html = `<p id="text">Text <strong id="strong">bold</strong> <b id="b">bold</b></p><p id="sans" class="font-sans">sans</p><button id="button">Buy</button>
<section id="dark" class="cui-dark dark:underline"><p id="in-dark" class="dark:underline">in a dark block</p>
  <div id="light" class="cui-light"><p id="in-light" class="dark:underline">light inside dark</p></div></section>
<p id="outside" class="dark:underline">outside</p><p id="moving" style="transition: color 1s 0.5s">moving</p><button id="plain" class="outline-none">Plain</button>
<p id="md" class="md:underline">md</p><p id="lg" class="lg:underline">lg</p><p id="xl" class="xl:underline">xl</p>
<style>@keyframes cui-test { to { opacity: 0.5; } }</style><p id="animated" style="animation: cui-test 1s infinite">animated</p><div id="smooth" style="scroll-behavior: smooth">smooth</div>
<div id="fixed" style="background-attachment: fixed">fixed</div><style>#pseudo::before { content: ""; transition: color 1s 0.5s; animation: cui-test 1s infinite; }</style><p id="pseudo">pseudo</p>`
	stack = `-apple-system, BlinkMacSystemFont, Inter, "Helvetica Neue", Helvetica, Arial, "PingFang TC", "Noto Sans TC", "Noto Sans CJK TC", "Microsoft JhengHei", sans-serif`
	w     = 1200
)

var (
	light      = map[string]string{"color": "rgb(29, 29, 31)", "background-color": "rgb(255, 255, 255)"}
	dark       = map[string]string{"color": "rgb(245, 245, 247)", "background-color": "rgb(0, 0, 0)"}
	mozSmooth  = regexp.MustCompile(`(^|[}\s])body\{[^}]*-moz-osx-font-smoothing:grayscale`)
	minWidth   = regexp.MustCompile(`min-width:[^)]+`)
	divider    = map[string]string{"site": "--cui-c-divider", "console": "--cui-c-console-neutral-border-tertiary"}
	ownClasses = "text-body text-secondary bg-page bg-alt border border-t border-border border-divider text-link animate-spin"
)

type setColor struct {
	id, classes, prop string
	token             map[string]string
}

var setColors = []setColor{
	{"body", "text-body", "color", map[string]string{"site": "--cui-c-text", "console": "--cui-c-console-neutral-label-primary"}},
	{"secondary", "text-secondary", "color", map[string]string{"site": "--cui-c-text-2", "console": "--cui-c-console-neutral-label-secondary"}},
	{"page", "bg-page", "background-color", map[string]string{"site": "--cui-c-bg", "console": "--cui-c-console-neutral-background-primary"}},
	{"alt", "bg-alt", "background-color", map[string]string{"site": "--cui-c-bg-alt", "console": "--cui-c-console-card"}},
	{"edge", "border border-border", "border-top-color", map[string]string{"site": "--cui-c-border", "console": "--cui-c-console-field-border"}},
	{"bare", "border", "border-top-color", divider},
	{"top", "border-t", "border-top-color", divider},
	{"divider", "border border-divider", "border-top-color", divider},
	{"link", "text-link", "color", map[string]string{"site": "--cui-c-link", "console": "--cui-c-console-brand-label-primary"}},
}

type page struct {
	t *testing.T
	*browser.Page
}

func open(t *testing.T, b *browser.Browser, css, html string, o browser.Options) page {
	t.Helper()
	p, err := b.Open(css, html, o)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { p.Close() })
	return page{t, p}
}

func (p page) read(selector string, props []string, o browser.ReadOptions) map[string]string {
	p.t.Helper()
	got, err := p.Read(selector, props, o)
	if err != nil {
		p.t.Fatal(err)
	}
	return got
}

func (p page) one(selector, prop string) string {
	p.t.Helper()
	return p.read(selector, []string{prop}, browser.ReadOptions{})[prop]
}

func (p page) resize(width int) {
	p.t.Helper()
	if err := p.Resize(width); err != nil {
		p.t.Fatal(err)
	}
}

func appleValue(t *testing.T, page reference.Page, probe, prop, state string) string {
	t.Helper()
	v, err := page.Value(probe, w, prop, state)
	if err != nil {
		t.Fatal(err)
	}
	return v
}

func TestPage(t *testing.T) {
	css, err := build.Compile(build.Sheet{CSS: "@import \"index.css\";\n@import \"site.css\";", Classes: "dark:underline outline-none md:underline lg:underline xl:underline bg-secondary text-alt bg-alt font-sans", Sources: nil})
	if err != nil {
		t.Fatal(err)
	}
	b := browser.Launch(t)

	p := open(t, b, css, html, browser.Options{})
	body := p.read("body", []string{"font-size", "line-height", "font-weight", "color", "background-color", "-webkit-font-smoothing", "font-synthesis", "text-rendering", "letter-spacing", "font-family"}, browser.ReadOptions{})
	for _, prop := range []string{"font-size", "line-height", "font-weight", "color", "background-color", "-webkit-font-smoothing", "font-synthesis", "text-rendering"} {
		check.Equal(t, fmt.Sprintf("body %s is apple.com's home:page.body", prop), body[prop], appleValue(t, apple.Home, "page.body", prop, ""))
	}
	// @external: Firefox's, on macOS: Chromium drops it, so the stylesheet is read
	check.Equal(t, "body declares -moz-osx-font-smoothing grayscale, as apple.com's stylesheets do", strconv.FormatBool(mozSmooth.MatchString(css)), "true")
	fitted := 17.0
	check.Within(t, "body letter-spacing is the 17px fitted value", body["letter-spacing"], check.Number(fitted*0.0029)+"px", 0.001)
	check.Equal(t, "body font-family is the stack", body["font-family"], stack)
	check.Equal(t, "font-sans is the stack, not Tailwind's", p.one("#sans", "font-family"), stack)
	for _, id := range []string{"strong", "b"} {
		check.Equal(t, fmt.Sprintf("%s is apple.com's store:page.strong weight", id), p.one("#"+id, "font-weight"), func() string {
			v, err := apple.Store.Value("page.strong", w, "font-weight", "")
			if err != nil {
				t.Fatal(err)
			}
			return v
		}())
	}

	ringProps := []string{"outline-style", "outline-color", "outline-width", "outline-offset"}
	ring := p.read("#button", ringProps, browser.ReadOptions{State: []string{"focus-visible"}})
	for _, prop := range ringProps {
		check.Equal(t, fmt.Sprintf("the focus ring's %s is apple.com's home:button:focus-visible", prop), ring[prop], appleValue(t, apple.Home, "button", prop, "focus-visible"))
	}
	check.Equal(t, "no ring at rest", p.one("#button", "outline-style"), "none")
	check.Equal(t, "a utility on the element wins over base.css: outline-none drops the ring", p.read("#plain", []string{"outline-style"}, browser.ReadOptions{State: []string{"focus-visible"}})["outline-style"], "none")

	for _, s := range [][3]string{{"text", "text on the page", "light"}, {"dark", "a cui-dark block", "dark"}, {"in-dark", "text in a cui-dark block", "dark"}, {"light", "a cui-light block inside a dark one", "light"}, {"in-light", "text in a cui-light block inside a dark one", "light"}} {
		id, what, scheme := s[0], s[1], s[2]
		want := light
		if scheme == "dark" {
			want = dark
		}
		got := p.read("#"+id, []string{"color", "background-color"}, browser.ReadOptions{})
		check.Equal(t, fmt.Sprintf("%s: its color %s", what, scheme), got["color"], want["color"])
		// @external: a paragraph paints no ground of its own: only the blocks do
		if !strings.HasPrefix(id, "in-") && id != "text" {
			check.Equal(t, fmt.Sprintf("%s: its ground %s", what, scheme), got["background-color"], want["background-color"])
		}
	}
	for _, s := range [][3]string{{"dark", "applies to the dark block itself", "underline"}, {"in-dark", "applies in a dark block", "underline"}, {"outside", "does not apply outside", "none"}, {"in-light", "applies in a light block inside a dark one too: for what no token carries", "underline"}} {
		check.Equal(t, "dark: "+s[1], p.one("#"+s[0], "text-decoration-line"), s[2])
	}
	var utilities []string
	for _, u := range []string{"bg-alt", "bg-secondary", "text-alt"} {
		if strings.Contains(css, "."+u+"{") {
			utilities = append(utilities, u)
		}
	}
	check.Equal(t, "text- and bg- names apart: of bg-alt, bg-secondary and text-alt, bg-alt alone a utility", strings.Join(utilities, " "), "bg-alt")
	check.Equal(t, "a transition moves without the preference", p.one("#moving", "transition-duration"), "1s")

	breakpoints := []struct {
		id        string
		below, at int
	}{{"md", 734, 735}, {"lg", 1068, 1069}, {"xl", 1440, 1441}}
	kit, err := build.OwnSheet("sm:underline md:underline lg:underline xl:underline 2xl:underline")
	if err != nil {
		t.Fatal(err)
	}
	var want []string
	for _, bp := range breakpoints {
		want = append(want, fmt.Sprintf("min-width:%dpx", bp.at))
	}
	check.Equal(t, "the kit's breakpoints alone: no sm: or 2xl:", strings.Join(minWidth.FindAllString(kit, -1), " "), strings.Join(want, " "))
	for _, bp := range breakpoints {
		p.resize(bp.below)
		check.Equal(t, fmt.Sprintf("%s: not at %dpx", bp.id, bp.below), p.one("#"+bp.id, "text-decoration-line"), "none")
		p.resize(bp.at)
		check.Equal(t, fmt.Sprintf("%s: from %dpx, apple.com's", bp.id, bp.at), p.one("#"+bp.id, "text-decoration-line"), "underline")
	}

	p = open(t, b, css, html, browser.Options{Dark: true, ReducedMotion: true})
	check.Equal(t, "dark: does not follow the OS's appearance", p.one("#outside", "text-decoration-line"), "none")
	for _, s := range [][4]string{{"#moving", "#animated", "", "an element"}, {"#pseudo", "#pseudo", "::before", "a pseudo-element"}} {
		moving, animated, pseudo, what := s[0], s[1], s[2], s[3]
		transition := p.read(moving, []string{"transition-duration", "transition-delay"}, browser.ReadOptions{Pseudo: pseudo})
		check.Equal(t, fmt.Sprintf("reduced motion stills %s's transition", what), transition["transition-duration"], "0s")
		check.Equal(t, fmt.Sprintf("reduced motion drops %s's transition delay", what), transition["transition-delay"], "0s")
		animation := p.read(animated, []string{"animation-duration", "animation-delay", "animation-iteration-count"}, browser.ReadOptions{Pseudo: pseudo})
		for _, pw := range [][2]string{{"animation-duration", "0.001s"}, {"animation-delay", "-0.001s"}, {"animation-iteration-count", "1"}} {
			check.Equal(t, fmt.Sprintf("reduced motion ends %s's animation at once, once: %s", what, pw[0]), animation[pw[0]], pw[1])
		}
	}
	check.Equal(t, "reduced motion scrolls at once", p.one("#smooth", "scroll-behavior"), "auto")
	check.Equal(t, "reduced motion scrolls a fixed background with its page", p.one("#fixed", "background-attachment"), "scroll")

	// @invariant: a site's setup as README.md tells it: its own utilities' sheet, built once,
	// linked after the set's dist sheet
	own, err := build.OwnSheet(ownClasses)
	if err != nil {
		t.Fatal(err)
	}
	drawn := map[string]map[string]string{}
	for _, set := range build.Sets {
		var html strings.Builder
		for _, c := range setColors {
			fmt.Fprintf(&html, `<p id="%s" class="%s">a</p><p id="%s-want" style="%s: var(%s); border-style: solid">a</p>`, c.id, c.classes, c.id, c.prop, c.token[set])
		}
		html.WriteString(`<style>#bare-pseudo::before, #bare-pseudo::after { content: ""; display: block; border-top-width: 1px; border-top-style: solid; }</style><p id="bare-pseudo">a</p>`)
		html.WriteString(`<span id="spinner" class="animate-spin">a</span><span id="still" style="animation: spin 1s infinite">a</span>`)
		shipped, err := build.Shipped(set)
		if err != nil {
			t.Fatal(err)
		}
		p := open(t, b, shipped, html.String(), browser.Options{ReducedMotion: true})
		if _, err := p.Page.Page.AddStyleTag(playwright.PageAddStyleTagOptions{Content: playwright.String(own)}); err != nil {
			t.Fatal(err)
		}
		drawn[set] = map[string]string{}
		// @trap: dark too: the console's link and accent share a light side
		for _, scheme := range []string{"light", "dark"} {
			if scheme == "dark" {
				if _, err := p.Page.Page.Evaluate(`() => document.documentElement.classList.add('cui-dark')`); err != nil {
					t.Fatal(err)
				}
			}
			want := map[string]string{}
			for _, c := range setColors {
				got := p.one("#"+c.id, c.prop)
				want[c.id] = p.one("#"+c.id+"-want", c.prop)
				if _, ok := drawn[set][c.id]; !ok {
					drawn[set][c.id] = got
				}
				check.Equal(t, fmt.Sprintf("%s, %s: a site's own %s draws %s", set, scheme, c.classes, c.token[set]), got, want[c.id])
			}
			for _, pseudo := range []string{"::before", "::after"} {
				check.Equal(t, fmt.Sprintf("%s, %s: a bare border on %s draws %s too", set, scheme, pseudo, divider[set]), p.read("#bare-pseudo", []string{"border-top-color"}, browser.ReadOptions{Pseudo: pseudo})["border-top-color"], want["bare"])
			}
		}
		spinner := p.read("#spinner", []string{"animation-duration", "animation-iteration-count"}, browser.ReadOptions{})
		check.Equal(t, set+": under reduced motion animate-spin still turns, its own duration", spinner["animation-duration"], "1s")
		check.Equal(t, set+": under reduced motion animate-spin still turns, without end", spinner["animation-iteration-count"], "infinite")
		check.Equal(t, set+": under reduced motion another animation still ends at once", p.one("#still", "animation-duration"), "0.001s")
	}
	// @trap: a set whose names drew the other set's colors would pass the checks above were
	// the two the same
	check.Equal(t, "the sets draw text-secondary apart", strconv.FormatBool(drawn["site"]["secondary"] != drawn["console"]["secondary"]), "true")
	check.Equal(t, "the sets draw a bare border apart", strconv.FormatBool(drawn["site"]["bare"] != drawn["console"]["bare"]), "true")
}
