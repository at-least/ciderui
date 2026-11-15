// Package harness runs a component's cases: its test file
// (components/<set>/<name>/<name>_test.go, build tag browser) lists them and calls
// Run, which opens each case's markup with every part's stylesheet, reads its
// checks at each width and holds each value to its source.
package harness

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strconv"
	"strings"
	"sync"
	"testing"
	"time"
	"unicode/utf8"

	"github.com/at-least/ciderui/internal/browser"
	"github.com/at-least/ciderui/internal/build"
	"github.com/at-least/ciderui/internal/check"
	"github.com/at-least/ciderui/internal/reference"
	asc "github.com/at-least/ciderui/reference/asc/results"
	"github.com/mxschmitt/playwright-go"
)

// Case is one page of a component's markup and the checks read off it.
type Case struct {
	// Name is what the case is; `go test -run 'TestCases/<regexp>'` picks cases by it.
	Name string
	// HTML is ciderui's markup: an example (Example), or markup written here.
	HTML string
	// Widths are the viewport widths it is read at (default reference.Widths);
	// Height the window's (default 800).
	Widths []int
	Height int
	// Touch, ReducedMotion and ForcedColors are the page's.
	Touch         bool
	ReducedMotion bool
	ForcedColors  bool
	// Click are selectors clicked in turn once, at the first width, before the
	// checks (a menu's button), then the keys in Press pressed in turn, the motion
	// then run to its end (browser.Page.Settle), or held as Wait and FirstFrame
	// say. A Wait or FirstFrame alone waits so once the page shows (a motion that
	// runs as it loads).
	Click []string
	Press []string
	// Wait is how long the case waits after its clicks and presses, and after each
	// hover, in place of running the motion to its end: a motion held, a case's own
	// long delay.
	Wait time.Duration
	// FirstFrame reads every motion at its first frame after the clicks and presses.
	FirstFrame bool
	// Scroll scrolls an element to X at every width before the checks (a scroller,
	// snapping where it snaps).
	Scroll *Scroll
	// Hover are selectors the pointer moves onto in turn, in steps, at every width,
	// its motion run to its end (a flyout the pointer shows).
	Hover  []string
	Checks []Check
}

// Scroll is an element scrolled to X.
type Scroll struct {
	Selector string
	X        int
}

// Check reads Props of Target (of its Pseudo element, ::after, where one is named),
// with State's pseudo-classes forced on it, or on On, and holds each to its one
// source: Apple, ASC, TB or Golden. Props are browser.Page.Read's.
type Check struct {
	Target string
	Props  []string
	Pseudo string
	State  []string
	On     string
	// Apple holds each value to apple.com's measured one.
	Apple *Apple
	// ASC holds each value to App Store Connect's measured one (a record in
	// reference/asc/results/, read at 1920px), at every width; never a width (its
	// text was SF).
	ASC *ASC
	// TB names a token of App Store Connect's toolkit for each of Props, in its order,
	// by its name without --tb- (reference/asc/results/tokens.go): the value equals
	// the token as the page computes it beside the target (a box read, box-height,
	// box-left, …, as the height or width it sets).
	TB []string
	// Golden is why the value is ciderui's own: it equals golden.json's, which
	// UPDATE_GOLDEN=1 rewrites.
	Golden string
}

// Apple is apple.com's probe (Page, a record in reference/apple.com/results/), in
// State if one is named, read at the same width, or at At[width] for a width the
// reference was not read at; never a box's width or a `width` (apple.com's text was
// SF, ciderui's here is Inter). As maps a property to apple.com's that does the same
// (an icon's margin-left to its glyph's padding-left).
type Apple struct {
	Page         reference.Page
	Probe, State string
	At           map[int]int
	As           map[string]string
}

// ASC is App Store Connect's probe.
type ASC struct {
	Page  reference.SignedInPage
	Probe string
}

// AppleProbe is apple.com's probe at rest, or in the one state named.
func AppleProbe(page reference.Page, probe string, state ...string) *Apple {
	if len(state) > 1 {
		panic(fmt.Sprintf("harness: %s:%s in %d states: a probe was read in one", page.Name, probe, len(state)))
	}
	return &Apple{Page: page, Probe: probe, State: strings.Join(state, ""), At: nil, As: nil}
}

// ASCProbe is App Store Connect's probe.
func ASCProbe(page reference.SignedInPage, probe string) *ASC {
	return &ASC{Page: page, Probe: probe}
}

// The property lists the checks share.
var (
	Padding     = []string{"padding-top", "padding-right", "padding-bottom", "padding-left"}
	Border      = []string{"border-top-width", "border-top-style", "border-bottom-width", "border-bottom-style"}
	borderColor = []string{"border-top-color", "border-bottom-color"}
	Colors      = append([]string{"color", "background-color"}, borderColor...)
	Size        = []string{"font-size", "line-height", "font-weight"}
	Radius      = []string{"border-top-left-radius"}
	Text        = []string{"display", "text-align", "white-space", "text-decoration-line", "cursor", "vertical-align", "opacity"}
	Outline     = []string{"outline-style", "outline-color", "outline-width", "outline-offset"}
	Motion      = []string{"transition-property", "transition-duration", "transition-timing-function"}
)

// With1600 is reference.Widths with 1600 first.
var With1600 = append([]int{1600}, reference.Widths...)

// Props joins property lists into one.
func Props(lists ...[]string) []string { return slices.Concat(lists...) }

// Example is the component's example examples/<id>.html, trimmed: the markup the Go
// module's tests render and the cases measure.
func Example(id string) string {
	b, err := os.ReadFile(filepath.Join("examples", id+".html"))
	if err != nil {
		panic(fmt.Sprintf("harness: %v", err))
	}
	return strings.TrimSpace(string(b))
}

// @invariant: every part's stylesheet with Tailwind's and the type's, as one page links
// them. fonts.css: without Inter the stack falls through to this Linux host's
// Arial (Liberation Sans). Tailwind scans the parts' files too: their cases write
// markup of their own (a type class around an example).
var sheet = sync.OnceValues(func() (string, error) {
	parts, err := build.Parts()
	if err != nil {
		return "", err
	}
	imports := []string{`@import "fonts.css";`, `@import "index.css";`}
	for _, p := range parts {
		imports = append(imports, fmt.Sprintf(`@import "%s/%s.css";`, p.Dir, p.Name))
	}
	return build.Compile(build.Sheet{CSS: strings.Join(imports, "\n"), Classes: "", Sources: []string{"components"}})
})

// Run runs the component's cases, the test's working directory its own.
func Run(t *testing.T, cases []Case) {
	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	part := filepath.Base(dir)
	update := os.Getenv("UPDATE_GOLDEN") == "1"
	picked := flag.Lookup("test.run").Value.String() != ""
	// @trap: rewriting golden.json from some cases would drop the others' values
	if update && picked {
		t.Fatal("UPDATE_GOLDEN=1 runs every case: drop -run")
	}
	for _, c := range cases {
		for _, k := range c.Checks {
			sources := 0
			for _, has := range []bool{k.Apple != nil, k.ASC != nil, k.TB != nil, k.Golden != ""} {
				if has {
					sources++
				}
			}
			if sources != 1 {
				t.Fatalf("%s %s: %s: a check is one of apple, asc, tb or golden", part, c.Name, k.Target)
			}
			if k.On != "" && len(k.State) == 0 {
				t.Fatalf("%s %s: %s: on needs a state (on is where the state is forced)", part, c.Name, k.Target)
			}
		}
	}
	css, err := sheet()
	if err != nil {
		t.Fatal(err)
	}
	golden := map[string]string{}
	goldenFile := filepath.Join(dir, "golden.json")
	if b, err := os.ReadFile(goldenFile); err == nil {
		if err := json.Unmarshal(b, &golden); err != nil {
			t.Fatalf("%s: %v", goldenFile, err)
		}
	} else if !os.IsNotExist(err) {
		t.Fatal(err)
	}
	r := &run{part: part, golden: golden, used: map[string]bool{}, update: update, written: nil}
	b := browser.Launch(t)
	ran := 0
	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			ran++
			r.run(t, b, css, c)
		})
	}
	if picked && ran == 0 {
		t.Fatalf("components: -run %s matches no case's name", flag.Lookup("test.run").Value.String())
	}
	switch {
	case update:
		if err := os.WriteFile(goldenFile, encode(r.written), 0o644); err != nil {
			t.Fatal(err)
		}
	case !picked:
		keys := make([]string, 0, len(golden))
		for k := range golden {
			keys = append(keys, k)
		}
		slices.Sort(keys)
		for _, key := range keys {
			check.Equal(t, fmt.Sprintf("%s golden.json's \"%s\" is read by a check", part, key), strconv.FormatBool(r.used[key]), "true")
		}
	}
}

type run struct {
	part    string
	golden  map[string]string
	used    map[string]bool
	update  bool
	written [][2]string
}

func (r *run) run(t *testing.T, b *browser.Browser, css string, c Case) {
	height := c.Height
	if height == 0 {
		height = 800
	}
	page, err := b.Open(css, c.HTML, browser.Options{Width: 0, Height: c.Height, Touch: c.Touch, ReducedMotion: c.ReducedMotion, ForcedColors: c.ForcedColors, Dark: false})
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err := page.Close(); err != nil {
			t.Error(err)
		}
	}()
	p := page.Page
	widths := c.Widths
	if widths == nil {
		widths = reference.Widths
	}
	waits := c.Wait != 0 || c.FirstFrame
	for _, width := range widths {
		must(t, page.Resize(width))
		// @invariant: once, at the first width: a menu stays open across the widths after
		// it. The pointer then leaves: a click leaves it over what it clicked, hovered, a
		// state the checks force when they mean it.
		if (c.Click != nil || c.Press != nil || waits) && width == widths[0] {
			for _, selector := range c.Click {
				must(t, p.Click(selector))
			}
			if c.Click != nil {
				must(t, p.Mouse().Move(0, float64(height-1)))
			}
			for _, key := range c.Press {
				must(t, p.Keyboard().Press(key))
			}
			// @trap: the first frame, not a timeout of 0: the click's round trip let a
			// transition start before the read (globalnav's menu-start read
			// rgba(250, 250, 252, 0.004) for transparent in one run of three)
			switch {
			case c.FirstFrame:
				_, err := p.Evaluate(`() => { for (const a of document.getAnimations()) { a.pause(); a.currentTime = 0; } }`)
				must(t, err)
			case c.Wait != 0:
				p.WaitForTimeout(float64(c.Wait.Milliseconds()))
			default:
				must(t, page.Settle())
			}
		}
		// @external: at every width: what the pointer shows a resize may hide. In steps,
		// as a hand moves it.
		if c.Hover != nil {
			must(t, p.Mouse().Move(0, float64(height-1)))
			must(t, page.Settle())
			for _, to := range c.Hover {
				box, err := p.Locator(to).BoundingBox()
				must(t, err)
				if box == nil {
					t.Fatalf("%s %s: %s has no box to hover", r.part, c.Name, to)
				}
				must(t, p.Mouse().Move(box.X+box.Width/2, box.Y+box.Height/2, playwright.MouseMoveOptions{Steps: playwright.Int(8)}))
				if c.Wait != 0 {
					p.WaitForTimeout(float64(c.Wait.Milliseconds()))
				} else {
					must(t, page.Settle())
				}
			}
		}
		if c.Scroll != nil {
			_, err := p.Evaluate(`([selector, x]) => { document.querySelector(selector).scrollLeft = x; }`, []any{c.Scroll.Selector, c.Scroll.X})
			must(t, err)
			p.WaitForTimeout(100)
		}
		for _, k := range c.Checks {
			r.check(t, page, c, k, width)
		}
	}
}

func (r *run) check(t *testing.T, page *browser.Page, c Case, k Check, width int) {
	state := ""
	if len(k.State) > 0 {
		state = ":" + strings.Join(k.State, ",")
	}
	label := fmt.Sprintf("%s %s: %s%s%s at %dpx", r.part, c.Name, k.Target, k.Pseudo, state, width)
	got, err := page.Read(k.Target, k.Props, browser.ReadOptions{State: k.State, On: k.On, Pseudo: k.Pseudo})
	if err != nil {
		t.Fatal(err)
	}
	switch {
	case k.Apple != nil:
		at := width
		if w, ok := k.Apple.At[width]; ok {
			at = w
		}
		state := ""
		if k.Apple.State != "" {
			state = ":" + k.Apple.State
		}
		for _, prop := range k.Props {
			// @external: apple.com drew SF, ciderui draws Inter here: a width is the text's
			if prop == "width" || prop == "box-width" {
				t.Fatalf("%s: an apple.com check reads no width (%s)", label, prop)
			}
			theirs := prop
			if as, ok := k.Apple.As[prop]; ok {
				theirs = as
			}
			want, err := k.Apple.Page.Value(k.Apple.Probe, at, theirs, k.Apple.State)
			if err != nil {
				t.Fatal(err)
			}
			check.Equal(t, fmt.Sprintf("%s %s is apple.com's %s:%s%s %s at %dpx", label, prop, k.Apple.Page.Name, k.Apple.Probe, state, theirs, at), got[prop], want)
		}
	case k.ASC != nil:
		// @external: App Store Connect's, read at 1920px: a value that does not move with
		// the window's width (a console part's own sizes)
		for _, prop := range k.Props {
			if prop == "width" || prop == "box-width" {
				t.Fatalf("%s: an asc check reads no width (%s)", label, prop)
			}
			want, err := k.ASC.Page.Value(k.ASC.Probe, prop)
			if err != nil {
				t.Fatal(err)
			}
			check.Equal(t, fmt.Sprintf("%s %s is App Store Connect's %s:%s %s", label, prop, k.ASC.Page.Name, k.ASC.Probe, prop), got[prop], want)
		}
	case k.TB != nil:
		if len(k.TB) != len(k.Props) {
			t.Fatalf("%s: tb names a token for each of props", label)
		}
		pairs := make([]any, len(k.Props))
		for i, prop := range k.Props {
			token, err := asc.Tokens.Token(k.TB[i])
			if err != nil {
				t.Fatal(err)
			}
			pairs[i] = []any{prop, token}
		}
		want, err := tokens(page, k.Target, pairs)
		if err != nil {
			t.Fatal(err)
		}
		for i, prop := range k.Props {
			check.Equal(t, fmt.Sprintf("%s %s is App Store Connect's --tb-%s", label, prop, k.TB[i]), got[prop], want[i])
		}
	default:
		for _, prop := range k.Props {
			key := fmt.Sprintf("%s | %s%s%s | %d | %s", c.Name, k.Target, k.Pseudo, state, width, prop)
			r.used[key] = true
			want, ok := r.golden[key]
			switch {
			case r.update:
				// @invariant: a key read twice keeps its first place and its last value, as
				// a JavaScript object's did
				if i := slices.IndexFunc(r.written, func(kv [2]string) bool { return kv[0] == key }); i >= 0 {
					r.written[i][1] = got[prop]
				} else {
					r.written = append(r.written, [2]string{key, got[prop]})
				}
			case !ok:
				check.Equal(t, fmt.Sprintf("%s %s has a golden value (UPDATE_GOLDEN=1)", label, prop), "none", "recorded")
			default:
				check.Equal(t, fmt.Sprintf("%s %s is golden (%s)", label, prop, k.Golden), got[prop], want)
			}
		}
	}
}

// @invariant: each token as the page computes it where the target is: beside it, so a
// length in em is its parent's
func tokens(page *browser.Page, target string, pairs []any) ([]string, error) {
	v, err := page.Page.Evaluate(`([selector, pairs]) => {
  const el = document.querySelector(selector);
  // @trap: an element no component's selector matches (a div would be a banner's > div),
  // a block where nothing sets its display
  const probe = document.createElement('cui-tb-probe');
  probe.style.display = 'block';
  el.parentElement.append(probe);
  const out = pairs.map(([prop, value]) => {
    // @invariant: the harness's box reads (internal/browser) are lengths: a token for one
    // is read as a height or a width
    const css = prop === 'box-height' ? 'height' : ['box-width', 'box-top', 'box-left', 'box-right', 'box-x', 'box-y'].includes(prop) ? 'width' : prop;
    probe.style.setProperty(css, value);
    const v = getComputedStyle(probe).getPropertyValue(css);
    probe.style.removeProperty(css);
    return v;
  });
  probe.remove();
  return out;
}`, []any{target, pairs})
	if err != nil {
		return nil, fmt.Errorf("harness: tokens: %w", err)
	}
	list, ok := v.([]any)
	if !ok {
		return nil, fmt.Errorf("harness: tokens gave %v", v)
	}
	out := make([]string, len(list))
	for i, x := range list {
		s, ok := x.(string)
		if !ok {
			return nil, fmt.Errorf("harness: token %d gave %v", i, x)
		}
		out[i] = s
	}
	return out, nil
}

func must(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal(err)
	}
}

// @invariant: golden.json as the Node harness wrote it (JSON.stringify(values, null, 1)
// and a newline): the cases' order, one space of indent, nothing escaped that JSON
// lets stand, so a rewrite with no value changed leaves the file as it was
func encode(values [][2]string) []byte {
	if len(values) == 0 {
		return []byte("{}\n")
	}
	var b bytes.Buffer
	b.WriteString("{\n")
	for i, kv := range values {
		b.WriteString(" ")
		quote(&b, kv[0])
		b.WriteString(": ")
		quote(&b, kv[1])
		if i < len(values)-1 {
			b.WriteString(",")
		}
		b.WriteString("\n")
	}
	b.WriteString("}\n")
	return b.Bytes()
}

func quote(b *bytes.Buffer, s string) {
	if !utf8.ValidString(s) {
		panic(fmt.Sprintf("harness: %q is not UTF-8", s))
	}
	b.WriteByte('"')
	for _, r := range s {
		switch r {
		case '"':
			b.WriteString(`\"`)
		case '\\':
			b.WriteString(`\\`)
		case '\b':
			b.WriteString(`\b`)
		case '\f':
			b.WriteString(`\f`)
		case '\n':
			b.WriteString(`\n`)
		case '\r':
			b.WriteString(`\r`)
		case '\t':
			b.WriteString(`\t`)
		default:
			if r < 0x20 {
				fmt.Fprintf(b, `\u%04x`, r)
			} else {
				b.WriteRune(r)
			}
		}
	}
	b.WriteByte('"')
}
