// Command details: BROWSER=safari|ios|chrome|chromium|firefox go run ./reference/details:
// MEASURED.md holds its claims and method. It writes page.html and ring.html first
// (page.go), then runs.
package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/at-least/ciderui/internal/ojson"
	"github.com/at-least/ciderui/internal/probe"
)

func main() {
	if err := writePages(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

const item = ".cui-accordion-item:nth-child(2)"

// @invariant: ring.html's: the first question in an answer, and the link beginning the
// example's first answer
var rings = []struct{ name, target string }{
	{"question", "#nested .cui-accordion-answer > .cui-accordion-item:first-child > .cui-accordion-question"},
	{"link", "#linked .cui-accordion-answer a"},
}

const read = "Promise.resolve((() => { const d = document.querySelector('" + item + "'); return { open: d.open, item: d.getBoundingClientRect().height, answer: d.querySelector('.cui-accordion-answer').getBoundingClientRect().height, chevron: getComputedStyle(d.querySelector('svg')).transform }; })())"

func quoted(s string) string {
	b, err := json.Marshal(s)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func run() error {
	browser := probe.Browser()
	dir := probe.Dir("details")
	html, err := os.ReadFile(filepath.Join(dir, "page.html"))
	if err != nil {
		return err
	}
	ring, err := os.ReadFile(filepath.Join(dir, "ring.html"))
	if err != nil {
		return err
	}
	l, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		return err
	}
	server := &http.Server{Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "text/html; charset=utf-8")
		if r.URL.RequestURI() == "/ring.html" {
			w.Write(ring)
			return
		}
		w.Write(html)
	})}
	go func() {
		if err := server.Serve(l); !errors.Is(err, http.ErrServerClosed) {
			fmt.Fprintln(os.Stderr, "serve:", err)
		}
	}()
	defer server.Close()
	origin := fmt.Sprintf("http://127.0.0.1:%d", l.Addr().(*net.TCPAddr).Port)
	src, err := os.ReadFile(filepath.Join(dir, "ring.js"))
	if err != nil {
		return err
	}
	page, err := probe.Open(browser, origin+"/", nil)
	if err != nil {
		return err
	}
	got, err := steps(page, browser, origin, string(src))
	label := page.Label()
	if err := closeAfter(page, err); err != nil {
		return err
	}
	if err := probe.Record("details", browser+".json", ojson.Object{{Key: "browser", Value: ojson.String(label)}, {Key: "date", Value: probe.Date()}}.With(got...)); err != nil {
		return err
	}
	closed, opening, opened := got.Get("closed"), got.Get("opening"), got.Get("opened")
	type claim struct {
		text string
		ok   bool
	}
	claims := []claim{
		{"it opens with no script, its answer shown", !truthy(get(closed, "open")) && truthy(get(opened, "open")) && num(get(opened, "item")) > num(get(closed, "item"))+30 && num(get(opened, "answer")) > 30},
		{"its chevron turned up", is(get(opened, "chevron"), "matrix(1, 0, 0, -1, 0, 0)")},
		{"it opens at once (its full height 120ms on)", num(get(opening, "item")) == num(get(opened, "item"))},
	}
	// @invariant: a side shown: at least a column of the screen's pixels the element's
	// height in the accent, beside it
	shown := func(r ojson.Value, side string) bool {
		return truthy(get(r, "focused")) && truthy(get(r, "visible")) && num(get(r, side)) >= num(get(r, "height"))*num(get(r, "scale"))
	}
	ringsGot, _ := got.Get("rings").(ojson.Object)
	for _, m := range ringsGot {
		parts := strings.Split(m.Key, " ")
		name, width := parts[0], parts[2]
		what, right := "link beginning an answer", ""
		if name == "question" {
			what, right = "question in an answer", ", and its right"
		}
		claims = append(claims, claim{
			fmt.Sprintf("the %s, focused by the keyboard at %spx: its ring's left side shown%s", what, width, right),
			shown(m.Value, "left") && (name != "question" || shown(m.Value, "right")),
		})
	}
	fmt.Println(label)
	pass := true
	for _, c := range claims {
		pass = pass && c.ok
		mark := "ok  "
		if !c.ok {
			mark = "FAIL"
		}
		fmt.Printf("%s %s\n", mark, c.text)
	}
	if !pass {
		os.Exit(1)
	}
	return nil
}

func steps(page probe.Page, browser, origin, src string) (ojson.Object, error) {
	waited := ojson.Object{}
	// @external: the window in front, where Safari matches :focus and a click reaches the
	// page: an app of the owner's took the Mac's front by itself during a run,
	// back and forth with Safari every second or two, no input (2026-10-22), so
	// what needs it waits for it, up to 10s, and names how long (null: never)
	front := func() (ojson.Value, error) {
		for i := range 40 {
			v, err := page.Run("Promise.resolve(document.hasFocus())")
			if err != nil {
				return nil, err
			}
			if truthy(v) {
				return ojson.Float(float64(i * 250)), nil
			}
			if _, err := page.Run(probe.Wait(250)); err != nil {
				return nil, err
			}
		}
		return ojson.Null{}, nil
	}
	wait := func(ms int) error {
		_, err := page.Run(probe.Wait(ms))
		return err
	}
	closed, err := page.Run(read)
	if err != nil {
		return nil, err
	}
	w, err := front()
	if err != nil {
		return nil, err
	}
	waited = waited.With(ojson.Member{Key: "click", Value: w})
	if browser == "ios" {
		_, err = page.Run("Promise.resolve((document.querySelector('" + item + " > summary').click(), true))")
	} else {
		err = page.Click(item + " > summary")
	}
	if err != nil {
		return nil, err
	}
	if err := wait(120); err != nil {
		return nil, err
	}
	opening, err := page.Run(read)
	if err != nil {
		return nil, err
	}
	if err := wait(580); err != nil {
		return nil, err
	}
	opened, err := page.Run(read)
	if err != nil {
		return nil, err
	}
	var ringsGot ojson.Value
	if browser != "ios" {
		tab := "Tab"
		if browser == "safari" {
			tab = "Alt+Tab"
		}
		got := ojson.Object{}
		for _, width := range []int{1200, 390} {
			if err := page.Size(width, 900); err != nil {
				return nil, err
			}
			for _, r := range rings {
				if err := page.Go(origin + "/ring.html"); err != nil {
					return nil, err
				}
				if err := wait(300); err != nil {
					return nil, err
				}
				focused := false
				for i := 0; i < 12 && !focused; i++ {
					if err := page.Press(tab); err != nil {
						return nil, err
					}
					v, err := page.Run("Promise.resolve(document.activeElement === document.querySelector(" + quoted(r.target) + "))")
					if err != nil {
						return nil, err
					}
					focused = truthy(v)
				}
				if err := wait(300); err != nil {
					return nil, err
				}
				// @external: read again, up to three times, if the window was behind as it was read
				var res ojson.Value = ojson.Object{{Key: "focused", Value: ojson.Bool(focused)}}
				for tries := 0; focused && tries < 3 && !truthy(get(res, "hasFocus")); tries++ {
					w, err := front()
					if err != nil {
						return nil, err
					}
					waited = waited.With(ojson.Member{Key: fmt.Sprintf("%s at %d, %d", r.name, width, tries+1), Value: w})
					shot, err := page.Shot()
					if err != nil {
						return nil, err
					}
					if res, err = page.Run("(async () => {\n" + src + "\nreturn ring(" + quoted(r.target) + ", " + quoted(shot) + ");\n})()"); err != nil {
						return nil, err
					}
				}
				got = got.With(ojson.Member{Key: fmt.Sprintf("%s at %d", r.name, width), Value: res})
			}
		}
		ringsGot = got
	}
	return ojson.Object{{Key: "waited", Value: waited}, {Key: "closed", Value: closed}, {Key: "opening", Value: opening}, {Key: "opened", Value: opened}, {Key: "rings", Value: ringsGot}}, nil
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
