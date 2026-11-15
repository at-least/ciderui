// Command console: BROWSER=safari|ios|chrome|chromium|firefox SITE=<dir> go run ./reference/console
// Do the console parts their tests (internal/harness) hold in Chromium do the same in
// this browser? It builds the pages into SITE (build.go), served at
// http://127.0.0.1:8738 (the browsers: internal/probe). Each part in a frame of a
// set width (index.html), so its media queries read the frame's width in every
// browser, a phone's below the window's least width in Safari:
//  1. the tabs' phone row (320: six tabs overflow it in every font) is one
//     row that scrolls sideways, none of its height a scrollbar's;
//  2. not read: the tabs' edge fade, removed by decision 40 (its runs are
//     MEASURED.md's);
//  3. the row's last tab, its start a sliver inside the row (Linux's Inter:
//     3px), focused by script: brought in whole (within 1px), its ring inside
//     the tab where :focus-visible matches;
//  4. the page column: one column at 900; at 1200 the sidebar's 225px
//     column 30px before the page, 40px in, its header and body with no room
//     at their sides; at 1920 the column 1440px wide, centered (240px in);
//
// The result is results/<browser>.json; exits 1 if a claim does not hold.
package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/at-least/ciderui/internal/ojson"
	"github.com/at-least/ciderui/internal/probe"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

const origin = "http://127.0.0.1:8738"

func frame(id string) string { return "document.getElementById(" + quoted(id) + ").contentWindow" }

func quoted(s string) string {
	b, err := json.Marshal(s)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func run() error {
	browser := probe.Browser()
	site := os.Getenv("SITE")
	if site == "" {
		return errors.New("SITE=<the directory the pages are built into>")
	}
	if err := writeSite(site); err != nil {
		return err
	}
	if _, err := os.Stat(filepath.Join(site, "index.html")); err != nil {
		return err
	}
	server, err := probe.Serve(site, 8738)
	if err != nil {
		return err
	}
	defer server.Close()

	claims := ojson.Object{}
	var fails []string
	claim := func(name string, ok bool, value ojson.Value) {
		claims = claims.With(ojson.Member{Key: name, Value: ojson.Object{{Key: "ok", Value: ojson.Bool(ok)}, {Key: "value", Value: value}}})
		mark := "ok  "
		if !ok {
			mark = "FAIL"
		}
		fmt.Printf("%s %s %s\n", mark, name, compact(value))
		if !ok {
			fails = append(fails, name)
		}
	}

	page, err := probe.Open(browser, origin+"/index.html", nil)
	if err != nil {
		return err
	}
	label := page.Label()
	fmt.Println(label)
	err = claimsOf(page, claim)
	if err := closeAfter(page, err); err != nil {
		return err
	}
	got := ojson.Object{{Key: "browser", Value: ojson.String(browser)}, {Key: "claims", Value: claims}, {Key: "label", Value: ojson.String(label)}}
	if err := probe.Record("console", browser+".json", got); err != nil {
		return err
	}
	if len(fails) > 0 {
		fmt.Printf("%d failed: %s\n", len(fails), strings.Join(fails, "; "))
		os.Exit(1)
	}
	fmt.Println("all held")
	return nil
}

func claimsOf(page probe.Page, claim func(string, bool, ojson.Value)) error {
	// @external: every frame loaded, its fonts and images too
	if _, err := page.Run(`(async () => {
    for (const f of document.querySelectorAll('iframe')) {
      if (f.contentDocument?.readyState !== 'complete') await new Promise((r) => f.addEventListener('load', r, { once: true }));
      await f.contentDocument.fonts.ready;
      await Promise.all([...f.contentDocument.images].map((i) => i.decode()));
    }
    return true;
  })()`); err != nil {
		return err
	}

	row, err := page.Run(`(async () => {
    const w = ` + frame("tabs-320") + `;
    const bar = w.document.querySelector('.cui-console-tabs');
    return { wrap: w.getComputedStyle(bar).flexWrap, scrollHeight: bar.scrollHeight, clientHeight: bar.clientHeight, offsetHeight: bar.offsetHeight, scrollWidth: bar.scrollWidth, clientWidth: bar.clientWidth };
  })()`)
	if err != nil {
		return err
	}
	claim("1 tabs one row, scrolling sideways", is(get(row, "wrap"), "nowrap") && num(get(row, "scrollHeight")) == num(get(row, "clientHeight")) && num(get(row, "offsetHeight"))-num(get(row, "clientHeight")) <= 1 && num(get(row, "scrollWidth")) > num(get(row, "clientWidth")), row)

	focusValue, err := page.Run(`(async () => {
    const w = ` + frame("tabs-320") + `;
    const bar = w.document.querySelector('.cui-console-tabs');
    bar.scrollLeft = 0;
    const tab = bar.querySelector('[href="#settings"]');
    tab.focus();
    // @external: read once the row has not moved for 200ms: a focus's scroll may land
    // after the next frame
    for (let last = -1, still = performance.now(); performance.now() - still < 200;) {
      await new Promise((r) => requestAnimationFrame(r));
      if (bar.scrollLeft !== last) [last, still] = [bar.scrollLeft, performance.now()];
    }
    const t = tab.getBoundingClientRect();
    const b = bar.getBoundingClientRect();
    return { focused: w.document.activeElement === tab, focusVisible: tab.matches(':focus-visible'), outlineOffset: w.getComputedStyle(tab).outlineOffset, scrollLeft: bar.scrollLeft, tab: [t.left - b.left, t.right - b.left], row: b.width };
  })()`)
	if err != nil {
		return err
	}
	focus, ok := focusValue.(ojson.Object)
	if !ok {
		return fmt.Errorf("the focus read gave %T", focusValue)
	}
	tab0, tab1, rowWidth := num(get(focus, "tab", 0)), num(get(focus, "tab", 1)), num(get(focus, "row"))
	shown := math.Min(tab1, rowWidth) - math.Max(tab0, 0)
	focus = focus.With(ojson.Member{Key: "shown", Value: ojson.Float(shown)})
	claim("3 a tab focused by script: brought in whole, its ring inside where :focus-visible matches", truthy(get(focus, "focused")) && shown >= tab1-tab0-1 && (!truthy(get(focus, "focusVisible")) || is(get(focus, "outlineOffset"), "-2px")), focus)

	content := map[string]ojson.Value{}
	for _, id := range []string{"content-900", "content-1200", "content-1920"} {
		v, err := page.Run(`(async () => {
      const w = ` + frame(id) + `;
      const c = w.document.querySelector('.cui-console-content');
      const s = w.document.querySelector('.cui-console-sidebar');
      const m = w.document.querySelector('.cui-console-content > main');
      const cs = w.getComputedStyle(c);
      const r = (e) => e.getBoundingClientRect();
      return { display: cs.display, columns: cs.gridTemplateColumns, contentLeft: r(c).left, contentWidth: r(c).width, sidebarLeft: r(s).left, sidebarWidth: r(s).width, mainLeft: r(m).left,
        header: w.getComputedStyle(s.querySelector('.cui-console-sidebar-header')).paddingInlineStart, body: w.getComputedStyle(s.querySelector('.cui-console-sidebar-body')).paddingInlineStart };
    })()`)
		if err != nil {
			return err
		}
		content[id] = v
	}
	c9 := content["content-900"]
	claim("4 page column at 900: one column", !is(get(c9, "display"), "grid"), c9)
	c12 := content["content-1200"]
	claim("4 page column at 1200: the sidebar 225px, 40px in, the page 30px past it, no room at its sides", is(get(c12, "display"), "grid") && num(get(c12, "sidebarLeft")) == 40 && num(get(c12, "sidebarWidth")) == 225 && num(get(c12, "mainLeft")) == 40+225+30 && is(get(c12, "header"), "0px") && is(get(c12, "body"), "0px"), c12)
	c19 := content["content-1920"]
	claim("4 page column at 1920: 1440px, centered, the sidebar 240px in", num(get(c19, "contentLeft")) == 240 && num(get(c19, "contentWidth")) == 1440 && num(get(c19, "sidebarLeft")) == 280, c19)
	return nil
}

// @external: JSON.stringify(v) with no indent, as a claim's line prints its value
func compact(v ojson.Value) string {
	var b bytes.Buffer
	var write func(v ojson.Value)
	write = func(v ojson.Value) {
		switch v := v.(type) {
		case ojson.Object:
			b.WriteByte('{')
			first := true
			for _, m := range v {
				if m.Value == nil {
					continue
				}
				if !first {
					b.WriteByte(',')
				}
				first = false
				b.Write(ojson.Marshal(ojson.String(m.Key)))
				b.WriteByte(':')
				write(m.Value)
			}
			b.WriteByte('}')
		case ojson.Array:
			b.WriteByte('[')
			for i, e := range v {
				if i > 0 {
					b.WriteByte(',')
				}
				write(e)
			}
			b.WriteByte(']')
		default:
			b.Write(ojson.Marshal(v))
		}
	}
	write(v)
	return b.String()
}

// @invariant: the page closes whether or not its run failed, and the run's error comes
// first
func closeAfter(page probe.Page, err error) error {
	if cerr := page.Close(); err == nil {
		err = cerr
	}
	return err
}

// @invariant: v's member or element down path (a key, an index), nil where JavaScript's
// would be undefined
func get(v ojson.Value, path ...any) ojson.Value {
	for _, p := range path {
		switch p := p.(type) {
		case string:
			o, ok := v.(ojson.Object)
			if !ok {
				return nil
			}
			v = o.Get(p)
		case int:
			a, ok := v.(ojson.Array)
			if !ok || p >= len(a) {
				return nil
			}
			v = a[p]
		}
	}
	return v
}

// @invariant: a number's value; NaN for anything else, so a comparison with what is not
// a number fails as JavaScript's with undefined does
func num(v ojson.Value) float64 {
	n, ok := v.(ojson.Number)
	if !ok {
		return math.NaN()
	}
	f, err := strconv.ParseFloat(string(n), 64)
	if err != nil {
		panic(err)
	}
	return f
}

func is(v ojson.Value, s string) bool {
	have, ok := v.(ojson.String)
	return ok && string(have) == s
}

// @external: JavaScript's truthiness of a JSON value
func truthy(v ojson.Value) bool {
	switch v := v.(type) {
	case ojson.Bool:
		return bool(v)
	case ojson.Number:
		f := num(v)
		return f != 0 && !math.IsNaN(f)
	case ojson.String:
		return v != ""
	case ojson.Object, ojson.Array:
		return true
	}
	return false
}
