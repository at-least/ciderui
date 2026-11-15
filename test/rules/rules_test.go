//go:build browser

// @invariant: The kit's rules that hold for every part, read off every example, not listed
// per part: a new part is held to them with no line of its own.
package rules

import (
	"encoding/json"
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
	"github.com/at-least/ciderui/internal/harness"
	"github.com/at-least/ciderui/internal/reference"
)

// @invariant: rule 5: what (hover: hover) gates, a touch screen never matches
// (internal/browser fails a touch page where it does)
func hoverOutsideItsQuery(text string) []string {
	var enclosing, found []string
	start := 0
	for i := 0; i < len(text); i++ {
		switch text[i] {
		case '{':
			prelude := strings.TrimSpace(text[start:i])
			enclosing = append(enclosing, prelude)
			if !strings.HasPrefix(prelude, "@") && strings.Contains(prelude, ":hover") && !slices.ContainsFunc(enclosing, func(p string) bool { return strings.Contains(p, "(hover:hover)") }) {
				found = append(found, prelude)
			}
			start = i + 1
		case '}':
			if len(enclosing) > 0 {
				enclosing = enclosing[:len(enclosing)-1]
			}
			start = i + 1
		case ';':
			start = i + 1
		}
	}
	return found
}

// @invariant: rule 9's exceptions, each read in its part's test: a text field draws its
// edge or halo in the ring's place
var noRing = []string{"cui-console-field-input", "cui-field-input"}

// @invariant: rule 10 is a pointer's color change; these change with a panel opening, a
// field's label rising or its text typed, each read in its part's test
var ownMotion = []string{".cui-field-label"}

const keyboardReachable = `a[href], button, input:not([type=hidden]), select, textarea, summary, [tabindex]:not([tabindex="-1"])`

const readTextAndMotion = `([set, tracking, ownMotion]) => {
  const label = (e) => e.tagName.toLowerCase() + (typeof e.className === 'string' && e.className ? '.' + e.className.trim().split(/\s+/).join('.') : '');
  const motion = [];
  const motionUsed = [];
  const spacing = [];
  let texts = 0;
  for (const e of document.body.querySelectorAll('*')) {
    const cs = getComputedStyle(e);
    const durations = cs.transitionDuration.split(', ');
    cs.transitionProperty.split(', ').forEach((prop, i) => {
      if (!(parseFloat(durations[i % durations.length]) > 0 && /(^|-)color$|^background$|^all$|^fill$|^stroke$/.test(prop))) return;
      const own = ownMotion.filter((s) => e.matches(s));
      if (own.length) motionUsed.push(...own);
      else motion.push(label(e) + ': ' + prop);
    });
    if (!e.matches('input:not([type=checkbox], [type=radio], [type=range]), textarea, select') && ![...e.childNodes].some((n) => n.nodeType === Node.TEXT_NODE && n.textContent.trim())) continue;
    texts++;
    if (set === 'console') {
      if (cs.letterSpacing !== 'normal') spacing.push(label(e) + ': ' + cs.letterSpacing);
    } else if (/monospace\s*$/.test(cs.fontFamily)) {
      // @external: the mono font is not the system font, whose tracking tokens.css fits:
      // code keeps its own spacing, as the console's does
      if (cs.letterSpacing !== 'normal') spacing.push(label(e) + ': mono at ' + cs.letterSpacing);
    } else {
      const size = parseFloat(cs.fontSize);
      // @trap: not "more than off": normal parses to NaN, which compares false either way
      if (!(Math.abs(parseFloat(cs.letterSpacing) - size * tracking[size]) <= 0.001)) spacing.push(label(e) + ': ' + cs.fontSize + ' at ' + cs.letterSpacing);
    }
  }
  return { motion, motionUsed, spacing, texts };
}`

type found struct {
	Motion     []string `json:"motion"`
	MotionUsed []string `json:"motionUsed"`
	Spacing    []string `json:"spacing"`
	Texts      int      `json:"texts"`
}

func TestRules(t *testing.T) {
	for _, set := range build.Sets {
		css, err := build.Shipped(set)
		if err != nil {
			t.Fatal(err)
		}
		check.Equal(t, set+": every :hover sits inside (hover: hover) (rule 5)", strings.Join(hoverOutsideItsQuery(css), " | "), "")
	}
	raw, err := os.ReadFile(filepath.Join(build.Root, "reference/tracking/cui-tracking.json"))
	if err != nil {
		t.Fatal(err)
	}
	// @external: map[string]any: playwright-go's argument serializer panics on any other
	// map type ("interface {} is map[string]float64, not map[string]interface {}")
	var tracking map[string]any
	if err := json.Unmarshal(raw, &tracking); err != nil {
		t.Fatal(err)
	}
	ring := map[string]string{}
	for _, prop := range harness.Outline {
		v, err := reference.Apple("home", "button", 1200, prop, "focus-visible")
		if err != nil {
			t.Fatal(err)
		}
		ring[prop] = v
	}

	noRingUsed := map[string]bool{}
	motionUsed := map[string]bool{}
	examples, controls, texts := 0, 0, 0
	b := browser.Launch(t)
	// @trap: one page a set, its body swapped: no example has a script, and a page's open is
	// the most of this step's time
	pages := map[string]*browser.Page{}
	for _, set := range build.Sets {
		css, err := build.Shipped(set)
		if err != nil {
			t.Fatal(err)
		}
		p, err := b.Open(css, "", browser.Options{})
		if err != nil {
			t.Fatal(err)
		}
		defer p.Close()
		pages[set] = p
	}
	parts, err := build.Parts()
	if err != nil {
		t.Fatal(err)
	}
	for _, part := range parts {
		page := pages[part.Set]
		files, err := build.ExampleFiles(part)
		if err != nil {
			t.Fatal(err)
		}
		for _, file := range files {
			name := part.Name + "/" + file
			html, err := build.ExampleHTML(part, file)
			if err != nil {
				t.Fatal(err)
			}
			if _, err := page.Page.Evaluate(`(html) => { document.body.innerHTML = html; }`, html); err != nil {
				t.Fatal(err)
			}
			examples++
			for _, width := range []int{390, 900, 1200} {
				if err := page.Resize(width); err != nil {
					t.Fatal(err)
				}
				v, err := page.Page.Evaluate(readTextAndMotion, []any{part.Set, tracking, ownMotion})
				if err != nil {
					t.Fatal(err)
				}
				var f found
				if err := browser.Decode(v, &f); err != nil {
					t.Fatal(err)
				}
				texts += f.Texts
				for _, s := range f.MotionUsed {
					motionUsed[s] = true
				}
				check.Equal(t, fmt.Sprintf("%s at %dpx: no color moves (rule 10)", name, width), strings.Join(f.Motion, " | "), "")
				what := "each size its fitted letter-spacing (tokens.css)"
				if part.Set == "console" {
					what = "the font tracks itself"
				}
				check.Equal(t, fmt.Sprintf("%s at %dpx: %s", name, width, what), strings.Join(f.Spacing, " | "), "")
			}
			v, err := page.Page.Evaluate(`(selector) => [...document.body.querySelectorAll(selector)].map((e, i) => {
  e.setAttribute('data-cui-control', i);
  return [...e.classList];
})`, keyboardReachable)
			if err != nil {
				t.Fatal(err)
			}
			var classes [][]string
			if err := browser.Decode(v, &classes); err != nil {
				t.Fatal(err)
			}
			for i, own := range classes {
				controls++
				var exempt []string
				for _, c := range noRing {
					if slices.Contains(own, c) {
						exempt = append(exempt, c)
					}
				}
				// @trap: both, as a keyboard's focus is: a field's :focus sets the ring aside
				got, err := page.Read(fmt.Sprintf(`[data-cui-control="%d"]`, i), harness.Outline, browser.ReadOptions{State: []string{"focus", "focus-visible"}})
				if err != nil {
					t.Fatal(err)
				}
				var off []string
				for _, prop := range harness.Outline {
					if got[prop] != ring[prop] {
						off = append(off, prop+" "+got[prop])
					}
				}
				if len(exempt) > 0 {
					if len(off) > 0 {
						for _, c := range exempt {
							noRingUsed[c] = true
						}
					}
					continue
				}
				control := "no class"
				if len(own) > 0 {
					control = "." + strings.Join(own, ".")
				}
				check.Equal(t, fmt.Sprintf("%s: control %d (%s)'s focus is base.css's one ring, apple.com's (rule 9)", name, i, control), strings.Join(off, " "), "")
			}
		}
	}
	// @invariant: an exception nothing needs is a part let off for nothing
	var idle []string
	for _, c := range noRing {
		if !noRingUsed[c] {
			idle = append(idle, c)
		}
	}
	check.Equal(t, "every part that sets the ring aside still does", strings.Join(idle, " "), "")
	idle = nil
	for _, s := range ownMotion {
		if !motionUsed[s] {
			idle = append(idle, s)
		}
	}
	check.Equal(t, "every part with its own color motion still has it", strings.Join(idle, " "), "")
	// @trap: a rule read off nothing would pass having held nothing
	check.Equal(t, fmt.Sprintf("the rules read every example (%d), its controls (%d) and its text (%d)", examples, controls, texts), strconv.FormatBool(examples > 0 && controls > 0 && texts > 0), "true")
}
