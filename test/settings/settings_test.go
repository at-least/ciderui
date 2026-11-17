//go:build browser

// @invariant: The system's settings a page must follow, read off every example as it
// ships, so a new part is held with no line of its own.
package settings

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strconv"
	"strings"
	"testing"

	"github.com/at-least/ciderui/internal/browser"
	"github.com/at-least/ciderui/internal/build"
	"github.com/at-least/ciderui/internal/check"
	"github.com/at-least/ciderui/internal/pinned"
	"github.com/mxschmitt/playwright-go"
)

type page struct {
	t *testing.T
	*browser.Page
}

func open(t *testing.T, b *browser.Browser, set, html string, o browser.Options) page {
	t.Helper()
	css, err := build.Shipped(set)
	if err != nil {
		t.Fatal(err)
	}
	p, err := b.Open(css, html, o)
	if err != nil {
		t.Fatal(err)
	}
	return page{t, p}
}

func (p page) close() {
	if err := p.Close(); err != nil {
		p.t.Error(err)
	}
}

func (p page) eval(js string, arg ...any) any {
	p.t.Helper()
	v, err := p.Page.Page.Evaluate(js, arg...)
	if err != nil {
		p.t.Fatal(err)
	}
	return v
}

func (p page) emulate(name, value string) {
	p.t.Helper()
	cdp, err := p.Page.Page.Context().NewCDPSession(p.Page.Page)
	if err != nil {
		p.t.Fatal(err)
	}
	if _, err := cdp.Send("Emulation.setEmulatedMedia", map[string]any{"features": []map[string]string{{"name": name, "value": value}}}); err != nil {
		p.t.Fatal(err)
	}
	on, err := p.Matches("(" + name + ": " + value + ")")
	if err != nil {
		p.t.Fatal(err)
	}
	check.Equal(p.t, fmt.Sprintf("%s: %s is on (the control)", name, value), strconv.FormatBool(on), "true")
}

func (p page) show(part build.Part, file string) {
	p.t.Helper()
	html, err := build.ExampleHTML(part, file)
	if err != nil {
		p.t.Fatal(err)
	}
	p.eval(`(html) => { document.body.innerHTML = html; }`, html)
}

func (p page) colorOf(value string) string {
	p.t.Helper()
	return p.eval(`(v) => {
  const probe = document.createElement('i');
  probe.style.color = v;
  document.body.append(probe);
  const color = getComputedStyle(probe).color;
  probe.remove();
  return color;
}`, value).(string)
}

// @invariant: Reduce Transparency: nothing blurs what is behind it, no ground lets it show.
// A tint or a dim is not a ground: under this alpha it colors what is behind on
// purpose. A mark too small to hold text is not one either. A field is near white
// on a page it is the same color as.
const (
	tint = 0.5
	mark = 24
)

// @trap: the field's ground allowed on a field's own box alone: the bar's translucent
// white is the same color, and its Reduce Transparency side read as a field's
const (
	fieldToken = "--cui-c-field"
	fieldOn    = ".cui-field-input, .cui-option-label"
)

const throughGround = `([tint, mark, field, fieldOn]) => [...document.body.querySelectorAll('*')].flatMap((e) => [null, '::before', '::after'].flatMap((pseudo) => {
  const cs = getComputedStyle(e, pseudo);
  const name = e.tagName.toLowerCase() + '.' + e.className + (pseudo ?? '');
  const alpha = Number(cs.backgroundColor.match(/, ([\d.]+)\)$/)?.[1] ?? cs.backgroundColor.match(/\/ ([\d.]+)\)$/)?.[1] ?? 1);
  return [
    ...(cs.backdropFilter !== 'none' ? [name + ' blurs'] : []),
    ...(alpha > tint && alpha < 1 && e.getBoundingClientRect().height >= mark && !(pseudo === null && e.matches(fieldOn) && cs.backgroundColor === field) ? [name + ' on ' + cs.backgroundColor] : []),
  ];
}))`

// @invariant: Increase Contrast: secondary text and edges at the text's color...
type slot struct {
	strong string
	names  []string
}

var contrast = []struct {
	set   string
	slots []slot
}{
	{"site", []slot{
		{"--cui-c-text", []string{"--cui-c-text-2", "--cui-c-text-3", "--cui-c-border", "--cui-c-divider"}},
		{"--cui-c-on-bar-hover", []string{"--cui-c-on-bar"}},
		{"--cui-c-on-footer-title", []string{"--cui-c-on-footer", "--cui-c-on-footer-link", "--cui-c-footer-separator"}},
	}},
	{"console", []slot{
		{"--cui-c-console-neutral-label-primary", []string{"--cui-c-console-neutral-label-secondary", "--cui-c-console-field-border", "--cui-c-console-neutral-border-secondary", "--cui-c-console-dialog-stroke"}},
		{"--cui-c-on-video", []string{"--cui-c-on-video-time"}},
	}},
}

func TestSettings(t *testing.T) {
	b := browser.Launch(t)
	parts, err := build.Parts()
	if err != nil {
		t.Fatal(err)
	}
	for _, set := range build.Sets {
		p := open(t, b, set, "", browser.Options{})
		p.emulate("prefers-reduced-transparency", "reduce")
		field := p.colorOf("var(" + fieldToken + ")")
		read := 0
		for _, width := range []int{1200, 390} {
			if err := p.Resize(width); err != nil {
				t.Fatal(err)
			}
			for _, part := range parts {
				if part.Set != set {
					continue
				}
				files, err := build.ExampleFiles(part)
				if err != nil {
					t.Fatal(err)
				}
				for _, file := range files {
					p.show(part, file)
					var found []string
					if err := browser.Decode(p.eval(throughGround, []any{tint, mark, field, fieldOn}), &found); err != nil {
						t.Fatal(err)
					}
					read++
					check.Equal(t, fmt.Sprintf("%s/%s at %dpx: nothing shows through under Reduce Transparency", part.Name, file, width), strings.Join(found, " | "), "")
				}
			}
		}
		check.Equal(t, fmt.Sprintf("%s: the examples read (%d)", set, read), strconv.FormatBool(read > 0), "true")
		p.close()
	}

	for _, c := range contrast {
		var pairs [][2]string
		for _, s := range c.slots {
			for _, name := range s.names {
				pairs = append(pairs, [2]string{name, s.strong})
			}
		}
		p := open(t, b, c.set, "", browser.Options{})
		for _, pair := range pairs {
			check.Equal(t, fmt.Sprintf("%s is not %s at rest (the control)", pair[0], pair[1]), strconv.FormatBool(p.colorOf("var("+pair[0]+")") != p.colorOf("var("+pair[1]+")")), "true")
		}
		p.emulate("prefers-contrast", "more")
		for _, pair := range pairs {
			check.Equal(t, fmt.Sprintf("%s is %s under Increase Contrast", pair[0], pair[1]), p.colorOf("var("+pair[0]+")"), p.colorOf("var("+pair[1]+")"))
		}
		p.close()
	}

	// @invariant: ...and so no neutral text is under AAA's 7:1 in any example; a hue, on the
	// text or under it, is an accent's, a link's or a status's, held to 4.5:1 (tokens.css)
	axeDir, err := pinned.Axe()
	if err != nil {
		t.Fatal(err)
	}
	axe, err := os.ReadFile(filepath.Join(axeDir, "axe.min.js"))
	if err != nil {
		t.Fatal(err)
	}
	for _, part := range parts {
		files, err := build.ExampleFiles(part)
		if err != nil {
			t.Fatal(err)
		}
		for _, file := range files {
			html, err := build.ExampleHTML(part, file)
			if err != nil {
				t.Fatal(err)
			}
			p := open(t, b, part.Set, html, browser.Options{})
			p.emulate("prefers-contrast", "more")
			if _, err := p.Page.Page.AddScriptTag(playwright.PageAddScriptTagOptions{Content: playwright.String(string(axe))}); err != nil {
				t.Fatal(err)
			}
			var found []struct {
				Target string `json:"target"`
				FG     string `json:"fg"`
				BG     string `json:"bg"`
			}
			// @external: a transition the switch to more contrast starts is read mid-way (a
			// field label under `transition: all` read #5d5d61 between #6e6e73 and its contrast
			// color); axe reads them done.
			// Transitions alone: an infinite animation (a busy spinner) never finishes, and
			// waiting on every animation hung the test until go test's timeout
			if err := browser.Decode(p.eval(`async () => {
  await Promise.all(document.getAnimations().filter((a) => a instanceof CSSTransition).map((a) => a.finished));
  const r = await axe.run(document, { runOnly: { type: 'rule', values: ['color-contrast-enhanced'] }, resultTypes: ['violations'] });
  return r.violations.flatMap((v) => v.nodes.map((n) => ({ target: n.target.join(' '), fg: n.any[0].data.fgColor, bg: n.any[0].data.bgColor })));
}`), &found); err != nil {
				t.Fatal(err)
			}
			p.close()
			var neutral []string
			for _, f := range found {
				if !hue(t, f.FG) && !hue(t, f.BG) {
					neutral = append(neutral, f.Target+" "+f.FG+" on "+f.BG)
				}
			}
			check.Equal(t, fmt.Sprintf("%s/%s: its neutral text 7:1 under Increase Contrast", part.Name, file), strings.Join(neutral, " | "), "")
		}
	}

	// @invariant: Appearance: the console follows the system, the site set stays light; a
	// class holds it
	ground := func(set, rootClass string, dark bool) string {
		p := open(t, b, set, "<i></i>", browser.Options{Dark: dark})
		defer p.close()
		p.eval(`(c) => { if (c) document.documentElement.classList.add(c); }`, rootClass)
		got, err := p.Read("body", []string{"background-color"}, browser.ReadOptions{})
		if err != nil {
			t.Fatal(err)
		}
		return got["background-color"]
	}
	const white, black = "rgb(255, 255, 255)", "rgb(0, 0, 0)"
	check.Equal(t, "the console, the system dark: dark", ground("console", "", true), black)
	check.Equal(t, "the console, the system light: light", ground("console", "", false), white)
	check.Equal(t, "the console with cui-light, the system dark: light", ground("console", "cui-light", true), white)
	check.Equal(t, "the console with cui-dark, the system light: dark", ground("console", "cui-dark", false), black)
	check.Equal(t, "the site set, the system dark: light", ground("site", "", true), white)
	check.Equal(t, "the site set with cui-dark, the system light: dark", ground("site", "cui-dark", false), black)
}

func hue(t *testing.T, hex string) bool {
	t.Helper()
	var channels []int64
	for _, i := range []int{1, 3, 5} {
		v, err := strconv.ParseInt(hex[i:i+2], 16, 0)
		if err != nil {
			t.Fatalf("axe gave %q, no #rrggbb", hex)
		}
		channels = append(channels, v)
	}
	return slices.Max(channels)-slices.Min(channels) > 16
}
