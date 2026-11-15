//go:build browser

// @invariant: Every component's example at 390px on an emulated touch screen
// (browser.Options.Touch): each control a finger can press, its box and its target
// (hit-tested from its center, below), held to the HIG's control sizes (Accessibility,
// Mobility: 44x44pt the default on iOS, 28x28pt the minimum). A control under the
// default fails; -v prints every control.
package touch

import (
	"fmt"
	"math"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/at-least/ciderui/internal/browser"
	"github.com/at-least/ciderui/internal/build"
	"github.com/at-least/ciderui/internal/check"
)

const controls = `() => {
  const sel = 'a[href], button, input:not([type=hidden]), select, textarea, summary, [role=button], [tabindex]:not([tabindex="-1"])';
  return [...document.querySelectorAll(sel)]
    // @external: what a finger can tap now: shown, not disabled, and taking the
    // pointer itself or through a label
    .filter((e) => e.checkVisibility({ visibilityProperty: true }) && !e.matches(':disabled') && (e.labels?.length || getComputedStyle(e).pointerEvents !== 'none'))
    // @external: a scroll box takes tabindex for the keyboard; a finger drags it and presses
    // nothing
    .filter((e) => e.matches('a[href], button, input, select, textarea, summary, [role=button]') || !/auto|scroll/.test(getComputedStyle(e).overflow)).map((e) => {
    const r = e.getBoundingClientRect();
    // @external: a link inside a sentence: WCAG 2.5.8's inline exception, listed,
    // not counted (its line's height is the text's)
    let block = e.parentElement;
    while (getComputedStyle(block).display === 'inline') block = block.parentElement;
    const inText = e.tagName === 'A' && getComputedStyle(e).display === 'inline' && block.textContent.replace(e.textContent, '').trim() !== '';
    // @external: its labels (wrapping it or naming it by for=) are its target too: a
    // tap on one acts on it (a floating label)
    const labels = [...(e.labels ?? [])].filter((l) => l.checkVisibility());
    const tiny = r.width <= 1 || r.height <= 1;
    const label = labels.length && (tiny || e.type === 'checkbox' || e.type === 'radio')
      ? labels.reduce((big, l) => (l.getBoundingClientRect().width * l.getBoundingClientRect().height > big.getBoundingClientRect().width * big.getBoundingClientRect().height ? l : big))
      : null;
    const anchor = label || e;
    // @external: a control hidden until it is focused (the skip link): read focused,
    // then blurred, as shown it covers the bar after it (the global nav's wordmark)
    const focused = !label && tiny;
    if (focused) e.focus();
    anchor.scrollIntoView({ block: 'center', inline: 'center' });
    const a = anchor.getBoundingClientRect();
    const x = a.left + a.width / 2, y = a.top + a.height / 2;
    // @trap: the target as a finger meets it: from the anchor's center, each
    // way, as far as a tap still lands on the control (a ::after drawn
    // round it counts; one clipped by a scroll box or under a
    // neighbour's does not), every 0.5px up to 40px: each way reads up
    // to 0.5px short, so 43 holds a 44px target
    const ours = (px, py) => {
      if (px < 0 || py < 0 || px >= innerWidth || py >= innerHeight) return false;
      const hit = document.elementFromPoint(px, py);
      return Boolean(hit) && (e.contains(hit) || labels.some((l) => l.contains(hit)));
    };
    const reach = (dx, dy) => {
      let d = 0;
      while (d < 40 && ours(x + dx * (d + 0.5), y + dy * (d + 0.5))) d += 0.5;
      return d;
    };
    const w = ours(x, y) ? reach(-1, 0) + reach(1, 0) : 0;
    const h = ours(x, y) ? reach(0, -1) + reach(0, 1) : 0;
    if (focused) e.blur();
    const name = e.tagName.toLowerCase() + (e.className && typeof e.className === 'string' ? '.' + e.className.trim().split(/\s+/)[0] : '') + (e.type && e.tagName === 'INPUT' ? '[' + e.type + ']' : '');
    return { name, inText, box: [+r.width.toFixed(1), +r.height.toFixed(1)], target: [w, h] };
  });
}`

type row struct {
	Name   string     `json:"name"`
	InText bool       `json:"inText"`
	Box    [2]float64 `json:"box"`
	Target [2]float64 `json:"target"`
}

func TestTouch(t *testing.T) {
	parts, err := build.Parts()
	if err != nil {
		t.Fatal(err)
	}
	b := browser.Launch(t)
	short, count := 0, 0
	for _, part := range parts {
		css, err := build.Shipped(part.Set)
		if err != nil {
			t.Fatal(err)
		}
		files, err := build.ExampleFiles(part)
		if err != nil {
			t.Fatal(err)
		}
		for _, file := range files {
			example, err := os.ReadFile(filepath.Join(build.Root, part.Dir, "examples", file))
			if err != nil {
				t.Fatal(err)
			}
			// @external: room round the example, so a target at its edge is not read short
			// where it runs out of the window (a tap there lands nowhere)
			page, err := b.Open(css, `<div style="padding: 48px 24px">`+string(example)+`</div>`, browser.Options{Width: 390, Touch: true})
			if err != nil {
				t.Fatal(err)
			}
			v, err := page.Page.Evaluate(controls)
			if err != nil {
				t.Fatal(err)
			}
			if err := page.Close(); err != nil {
				t.Fatal(err)
			}
			var rows []row
			if err := browser.Decode(v, &rows); err != nil {
				t.Fatal(err)
			}
			for _, r := range rows {
				least := math.Min(r.Target[0], r.Target[1])
				verdict := "ok   "
				switch {
				case r.InText:
					verdict = "text "
				case least < 28:
					verdict = "SHORT"
				case least < 43:
					verdict = "MIN  "
				}
				line := fmt.Sprintf("%s %s/%s %s box %s target %s", verdict, part.Name, file, r.Name, pair(r.Box), pair(r.Target))
				count++
				if least < 43 && !r.InText {
					short++
					t.Error(line)
				} else {
					t.Log(line)
				}
			}
		}
	}
	t.Logf("touch: %d controls, %d under 44pt", count, short)
	// @trap: no control read would pass having checked nothing
	if count == 0 {
		t.Fatal("touch: no control read")
	}
}

func pair(v [2]float64) string {
	return strings.Join([]string{check.Number(v[0]), check.Number(v[1])}, "x")
}
