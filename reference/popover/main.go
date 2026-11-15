// Command popover: BROWSER=safari|ios|chrome|chromium|firefox go run ./reference/popover
// MEASURED.md holds its claims and method.
package main

import (
	"encoding/json"
	"fmt"
	"math"
	"net"
	"net/http"
	"os"
	"path/filepath"

	"github.com/at-least/ciderui/internal/ojson"
	"github.com/at-least/ciderui/internal/probe"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	browser := probe.Browser()
	html, err := os.ReadFile(filepath.Join(probe.Dir("popover"), "page.html"))
	if err != nil {
		return err
	}
	l, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		return err
	}
	server := &http.Server{Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "text/html; charset=utf-8")
		w.Write(html)
	})}
	go server.Serve(l)
	defer server.Close()
	src, err := os.ReadFile(filepath.Join(probe.Dir("popover"), "popover.js"))
	if err != nil {
		return err
	}
	call := func(fn string) string { return "(async () => {\n" + string(src) + "\nreturn " + fn + ";\n})()" }
	ios := browser == "ios"
	keyboard := browser == "safari"
	page, err := probe.Open(browser, fmt.Sprintf("http://127.0.0.1:%d/", l.Addr().(*net.TCPAddr).Port), nil)
	if err != nil {
		return err
	}
	got := ojson.Object{}
	err = steps(page, &got, call, ios, keyboard)
	label := page.Label()
	if cerr := page.Close(); err == nil {
		err = cerr
	}
	if err != nil {
		return err
	}
	if err := probe.Record("popover", browser+".json", ojson.Object{{Key: "browser", Value: ojson.String(label)}, {Key: "date", Value: probe.Date()}}.With(got...)); err != nil {
		return err
	}

	g := js(got)
	near := func(a, b any) bool { return math.Abs(num(a)-num(b)) <= 0.05 }
	X := [][]float64{{3.5, 3.5, 11.5, 11.5}, {3.5, 3.5, 11.5, 11.5}}
	open1, close1 := at(g, "samples", 0), at(g, "samples", 1)
	opened := at(g, "opened")
	growing := false
	for _, f := range list(at(open1, "frames")) {
		if h := num(at(f, 1)); h > 48.5 && h < num(at(opened, "windowHeight"))-0.5 {
			growing = true
		}
	}
	underOne := false
	for _, f := range list(at(open1, "frames")) {
		if at(f, 3) != nil && num(at(f, 3)) < 1 {
			underOne = true
		}
	}
	anX := true
	for i, line := range list(at(opened, "lines")) {
		for j, v := range list(line) {
			anX = anX && near(v, X[i][j])
		}
	}
	closedFlat := true
	for _, f := range list(at(close1, "frames")) {
		closedFlat = closedFlat && num(at(f, 1)) == 48
	}
	closed, again, tabbed, escaped := at(g, "closed"), at(g, "closedAgain"), at(g, "tabbed"), at(g, "afterEscape")
	type claim struct {
		name string
		ok   bool
	}
	claims := []claim{
		{"closed, the page shows the list as its row", at(closed, "open") == false && at(closed, "display") == "flex" && at(closed, "position") == "static" && num(at(closed, "height")) == 48},
		{"its button opens it over the window", at(opened, "open") == true && at(opened, "position") == "fixed" && num(at(opened, "top")) == 0 && num(at(opened, "height")) == num(at(opened, "windowHeight"))},
		{"it grows from the bar: a frame between 48px and the window", at(open1, "to") == "open" && growing},
		{"its rows come in from transparent", num(at(open1, "frames", 0, 2)) == 0 && num(at(opened, "rowOpacity")) == 1},
		{"the page is held while it is open", at(opened, "rootOverflow") == "hidden" && at(again, "rootOverflow") != "hidden"},
		{"the focus is on its close button", at(opened, "focus") == "close"},
		{"the close button's lines turn from two lines (a frame with a line under 1px tall) into an X, each 11.5 by 11.5 from 3.5, 3.5", underOne && anX},
	}
	if !ios {
		key := "Tab"
		if keyboard {
			key = "Option-Tab"
		}
		claims = append(claims, claim{key + " goes from the close button to the first link", at(tabbed, "focus") == "one"})
	}
	claims = append(claims, claim{"the close button closes it at once", at(again, "open") == false && at(again, "position") == "static" && num(at(again, "height")) == 48 && at(close1, "to") == "closed" && closedFlat})
	if !ios {
		claims = append(claims,
			claim{"the focus is back on the button", at(again, "focus") == "button"},
			claim{"Escape closes it, the focus back on the button", at(escaped, "open") == false && at(escaped, "focus") == "button"})
	}
	how := ""
	if ios {
		how = " (clicked by script; no Tab, no Escape)"
	} else if keyboard {
		how = " (by its keyboard: Return, Option-Tab)"
	}
	fmt.Println(label + how)
	pass := true
	for _, c := range claims {
		pass = pass && c.ok
		mark := "ok  "
		if !c.ok {
			mark = "FAIL"
		}
		fmt.Printf("%s %s\n", mark, c.name)
	}
	if !pass {
		samples, _ := got.Get("samples").(ojson.Array)
		short := ojson.Array{}
		for _, s := range samples {
			o, _ := s.(ojson.Object)
			frames, _ := o.Get("frames").(ojson.Array)
			short = append(short, ojson.Object{{Key: "to", Value: o.Get("to")}, {Key: "frames", Value: frames[:min(12, len(frames))]}})
		}
		fmt.Println(string(ojson.Marshal(got.With(ojson.Member{Key: "samples", Value: short}))))
		fmt.Println("FAIL")
		os.Exit(1)
	}
	fmt.Println("PASS")
	return nil
}

func steps(page probe.Page, got *ojson.Object, call func(string) string, ios, keyboard bool) error {
	click := func(selector string) error {
		if ios {
			return probe.Script(page, selector, "click")
		}
		if keyboard {
			if err := probe.Script(page, selector, "focus"); err != nil {
				return err
			}
			return page.Press("Enter")
		}
		return page.Click(selector)
	}
	record := func(key, expr string) error {
		v, err := page.Run(expr)
		*got = got.With(ojson.Member{Key: key, Value: v})
		return err
	}
	wait := func(ms int) error { _, err := page.Run(probe.Wait(ms)); return err }
	if _, err := page.Run(call("arm()")); err != nil {
		return err
	}
	if err := record("closed", call("state()")); err != nil {
		return err
	}
	if err := click("#button"); err != nil {
		return err
	}
	if err := wait(900); err != nil {
		return err
	}
	if err := record("opened", call("state()")); err != nil {
		return err
	}
	if !ios {
		key := "Tab"
		if keyboard {
			key = "Alt+Tab"
		}
		if err := page.Press(key); err != nil {
			return err
		}
		if err := record("tabbed", call("state()")); err != nil {
			return err
		}
		if err := probe.Script(page, "#close", "focus"); err != nil {
			return err
		}
	}
	if err := click("#close"); err != nil {
		return err
	}
	if err := record("closedAgain", call("state()")); err != nil {
		return err
	}
	if !ios {
		if err := wait(800); err != nil {
			return err
		}
		if err := click("#button"); err != nil {
			return err
		}
		if err := wait(600); err != nil {
			return err
		}
		if err := page.Press("Escape"); err != nil {
			return err
		}
		if err := record("afterEscape", call("state()")); err != nil {
			return err
		}
	}
	return record("samples", "Promise.resolve(window.cuiSamples)")
}

// @invariant: the record as plain Go values for the claims: objects maps, numbers
// float64, null and undefined nil
func js(v ojson.Value) any {
	var x any
	if v != nil {
		if err := json.Unmarshal(ojson.Marshal(v), &x); err != nil {
			panic(err)
		}
	}
	return x
}

// @invariant: a member or an element down path, nil where there is none (JavaScript's
// undefined, which compares unequal to every value but undefined)
func at(v any, path ...any) any {
	for _, k := range path {
		switch k := k.(type) {
		case string:
			m, _ := v.(map[string]any)
			v = m[k]
		case int:
			a, _ := v.([]any)
			if k >= len(a) {
				return nil
			}
			v = a[k]
		}
	}
	return v
}

// @invariant: NaN where v is no number, so every comparison with it is false, as
// with undefined in JavaScript
func num(v any) float64 {
	f, ok := v.(float64)
	if !ok {
		return math.NaN()
	}
	return f
}

func list(v any) []any {
	a, _ := v.([]any)
	return a
}
