// Command dialog: BROWSER=safari|ios|chrome|chromium|firefox go run ./reference/dialog:
// MEASURED.md holds its claims and method.
package main

import (
	"fmt"
	"math"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

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
	html, err := os.ReadFile(filepath.Join(probe.Dir("dialog"), "page.html"))
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
	src, err := os.ReadFile(filepath.Join(probe.Dir("dialog"), "dialog.js"))
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
	click := func(selector string) error {
		switch {
		case ios:
			return probe.Script(page, selector, "click")
		case keyboard:
			if err := probe.Script(page, selector, "focus"); err != nil {
				return err
			}
			return page.Press("Enter")
		}
		return page.Click(selector)
	}
	got, err := steps(page, ios, click, call)
	label := page.Label()
	if err := closeAfter(page, err); err != nil {
		return err
	}
	if err := probe.Record("dialog", browser+".json", ojson.Object{{Key: "browser", Value: ojson.String(label)}, {Key: "date", Value: probe.Date()}}.With(got...)); err != nil {
		return err
	}

	o := got.Get("opened")
	var frames []ojson.Value
	samples, _ := got.Get("samples").(ojson.Array)
	if len(samples) == 0 {
		return fmt.Errorf("no samples: %v", got.Get("samples"))
	}
	all, _ := get(samples[0], "frames").(ojson.Array)
	for _, f := range all {
		if truthy(get(f, 3)) {
			frames = append(frames, f)
		}
	}
	someUnder1 := false
	for _, f := range frames {
		if num(get(f, 1)) < 1 {
			someUnder1 = true
		}
	}
	closed, closedAgain, afterEscape := got.Get("closed"), got.Get("closedAgain"), got.Get("afterEscape")
	formButton, formSubmit := got.Get("formButton"), got.Get("formSubmit")
	type claim struct {
		text string
		ok   bool
	}
	claims := []claim{
		{"the browser has the command attribute", get(closed, "command") == ojson.Bool(true)},
		{"the page's button opens it as a modal dialog, with no script", truthy(get(o, "open")) && truthy(get(o, "modal"))},
		{"it is the layer over the window", is(get(o, "position"), "fixed") && num(get(o, "top")) == 0 && num(get(o, "left")) == 0 && num(get(o, "width")) == num(get(o, "windowWidth")) && num(get(o, "height")) == num(get(o, "windowHeight"))},
		{"the page is dimmed under it", is(get(o, "backdrop"), "rgba(0, 0, 0, 0.48)")},
		{"the page is held while it is open", is(get(o, "rootOverflow"), "hidden") && !is(get(closed, "rootOverflow"), "hidden") && !is(get(closedAgain, "rootOverflow"), "hidden")},
		{"the layer scrolls the panel taller than the window", get(o, "scrolls") == ojson.Bool(true)},
		{"a point above the panel hits the layer, not the page", is(get(o, "hit"), "modal")},
		{"the focus is on its close button, the dialog's first control", is(get(o, "focus"), "close")},
		{"it fades in: an open frame under 1, then 1", someUnder1 && truthy(get(o, "open")) && num(get(frames[len(frames)-1], 1)) == 1},
		{"its close button closes it", get(closedAgain, "open") == ojson.Bool(false)},
	}
	if !ios {
		claims = append(claims,
			claim{"the focus is back on the button that opened it", is(get(closedAgain, "focus"), "button")},
			claim{"Escape closes it, the focus back on the button", get(afterEscape, "open") == ojson.Bool(false) && is(get(afterEscape, "focus"), "button")},
		)
	}
	claims = append(claims,
		claim{"a button in a form, its type button, opens it, the form not sent", get(formButton, "open") == ojson.Bool(true) && is(get(formButton, "hash"), "")},
		claim{"a button in a form with no type does nothing, the dialog closed and the form not sent (the control)", get(formSubmit, "open") == ojson.Bool(false) && is(get(formSubmit, "hash"), "")},
	)
	suffix := ""
	if ios {
		suffix = " (clicked by script; no Escape)"
	} else if keyboard {
		suffix = " (by its keyboard: Return)"
	}
	fmt.Println(label + suffix)
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
		var first10 ojson.Array
		for _, s := range samples {
			f, _ := get(s, "frames").(ojson.Array)
			first10 = append(first10, f[:min(10, len(f))])
		}
		fmt.Println(string(ojson.Marshal(got.With(ojson.Member{Key: "samples", Value: first10}))))
		fmt.Println("FAIL")
		os.Exit(1)
	}
	fmt.Println("PASS")
	return nil
}

// @invariant: got's keys in the order the steps set them, as the record keeps them
func steps(page probe.Page, ios bool, click func(string) error, call func(string) string) (ojson.Object, error) {
	var got ojson.Object
	set := func(key, expr string) error {
		v, err := page.Run(expr)
		got = got.With(ojson.Member{Key: key, Value: v})
		return err
	}
	wait := func(ms int) error {
		_, err := page.Run(probe.Wait(ms))
		return err
	}
	if !ios {
		if err := page.Size(1440, 900); err != nil {
			return nil, err
		}
	}
	if err := wait(300); err != nil {
		return nil, err
	}
	if err := set("closed", call("state()")); err != nil {
		return nil, err
	}
	if _, err := page.Run(call("watch()")); err != nil {
		return nil, err
	}
	if err := click("#button"); err != nil {
		return nil, err
	}
	if err := wait(700); err != nil {
		return nil, err
	}
	if err := set("opened", call("state()")); err != nil {
		return nil, err
	}
	if err := click("#close"); err != nil {
		return nil, err
	}
	if err := wait(300); err != nil {
		return nil, err
	}
	if err := set("closedAgain", call("state()")); err != nil {
		return nil, err
	}
	if !ios {
		if err := click("#button"); err != nil {
			return nil, err
		}
		if err := wait(500); err != nil {
			return nil, err
		}
		if err := page.Press("Escape"); err != nil {
			return nil, err
		}
		if err := wait(300); err != nil {
			return nil, err
		}
		if err := set("afterEscape", call("state()")); err != nil {
			return nil, err
		}
	}
	if err := set("samples", "Promise.resolve(window.cuiSamples)"); err != nil {
		return nil, err
	}
	const opened = "Promise.resolve({ open: document.querySelector('#modal').open, hash: location.hash })"
	if err := click("#form-button"); err != nil {
		return nil, err
	}
	if err := wait(500); err != nil {
		return nil, err
	}
	if err := set("formButton", opened); err != nil {
		return nil, err
	}
	if err := click("#close"); err != nil {
		return nil, err
	}
	if err := wait(300); err != nil {
		return nil, err
	}
	if err := click("#form-submit"); err != nil {
		return nil, err
	}
	if err := wait(1000); err != nil {
		return nil, err
	}
	if err := set("formSubmit", opened); err != nil {
		return nil, err
	}
	return got, nil
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
