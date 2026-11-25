// Command tap: BROWSER=chromium|chrome|firefox|ios go run ./reference/tap
// MEASURED.md holds its claims and method.
package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"slices"
	"sync"
	"time"

	"github.com/at-least/ciderui/internal/ojson"
	"github.com/at-least/ciderui/internal/probe"
	"github.com/at-least/ciderui/internal/xcuitest"
	"github.com/mxschmitt/playwright-go"
)

// @invariant: what is tapped, in order: its id (the DevTools protocol's taps) and its
// label (XCUITest's)
var steps = []struct{ id, label string }{{"plain", "Plain"}, {"gated", "Gated"}, {"elsewhere", "Elsewhere"}, {"menu-button", "Menu"}, {"menu-close", "Close"}}

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// @invariant: what the page sent to /report, in the order it came
type reports struct {
	mu   sync.Mutex
	list ojson.Array
}

func (r *reports) all() ojson.Array {
	r.mu.Lock()
	defer r.mu.Unlock()
	return append(ojson.Array{}, r.list...)
}

// @invariant: the page's report of the i-th tap
func (r *reports) reported(i int, id string) error {
	for t := 0; len(r.all()) < i+1; t++ {
		if t == 100 {
			return fmt.Errorf("no report 10s after tapping %s", id)
		}
		time.Sleep(100 * time.Millisecond)
	}
	return nil
}

func run() error {
	browser := probe.Browser()
	html, err := os.ReadFile(filepath.Join(probe.Dir("tap"), "page.html"))
	if err != nil {
		return err
	}
	sent := &reports{mu: sync.Mutex{}, list: nil}
	l, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		return err
	}
	server := &http.Server{Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost && r.RequestURI == "/report" {
			body, err := io.ReadAll(r.Body)
			if err != nil {
				panic(err)
			}
			v, err := ojson.Parse(body)
			if err != nil {
				panic(err)
			}
			sent.mu.Lock()
			sent.list = append(sent.list, v)
			sent.mu.Unlock()
			w.WriteHeader(http.StatusNoContent)
			return
		}
		w.Header().Set("content-type", "text/html; charset=utf-8")
		w.Write(html)
	})}
	go func() {
		if err := server.Serve(l); !errors.Is(err, http.ErrServerClosed) {
			fmt.Fprintln(os.Stderr, "serve:", err)
		}
	}()
	defer server.Close()
	url := fmt.Sprintf("http://127.0.0.1:%d/", l.Addr().(*net.TCPAddr).Port)

	var label, tip string
	switch browser {
	case "ios":
		var labels []string
		for _, s := range steps {
			labels = append(labels, s.label)
		}
		result, err := xcuitest.Taps(url, labels, xcuitest.IPhone17)
		if err != nil {
			return err
		}
		label, tip = result.Label, result.Tip
	case "firefox":
		label, err = viaTouchscreen(browser, url, sent)
	default:
		label, err = viaDevTools(browser, url, sent)
	}
	if err != nil {
		return err
	}
	all := sent.all()
	record := ojson.Object{{Key: "browser", Value: ojson.String(label)}, {Key: "date", Value: probe.Date()}}
	if tip != "" {
		record = record.With(ojson.Member{Key: "tip", Value: ojson.String(tip)})
	}
	if err := probe.Record("tap", browser+".json", record.With(ojson.Member{Key: "reports", Value: all})); err != nil {
		return err
	}

	r := list(js(all))
	report := func(i int) any {
		if i < len(r) {
			return r[i]
		}
		return nil
	}
	plain, gated, elsewhere, menu, closed := report(0), report(1), report(2), report(3), report(4)
	on := func(report any, id string) []any {
		var events []any
		for _, e := range list(at(report, "events")) {
			if at(e, "on") == id {
				events = append(events, e)
			}
		}
		return events
	}
	has := func(events []any, typ string, pointer ...string) bool {
		return slices.ContainsFunc(events, func(e any) bool {
			return at(e, "type") == typ && at(e, "trusted") == true && (len(pointer) == 0 || at(e, "pointer") == pointer[0])
		})
	}
	tapped := on(plain, "plain")
	mouse := true
	for _, t := range []string{"mouseover", "mousedown", "mouseup", "click"} {
		mouse = mouse && has(tapped, t)
	}
	noHover := true
	for _, rep := range r {
		noHover = noHover && at(rep, "media", "hover") != true && at(rep, "media", "coarse") == true
	}
	claims := []struct {
		name string
		ok   bool
	}{
		{"a tap is a touch: its pointerdown and pointerup `touch` and trusted, then mouseover, mousedown, mouseup and a click on the button", has(tapped, "pointerdown", "touch") && has(tapped, "pointerup", "touch") && mouse},
		{"the screen has no hover: (hover: hover) does not match, (pointer: coarse) does", noHover},
		{"the tapped button is left hovered (:hover 800ms on), and a hover rule of its own shows there (apple.com's way: #0076df)", truthy(at(plain, "plain", "hover")) && at(plain, "plain", "background") == "rgb(0, 118, 223)"},
		{"a hover behind (hover: hover) does not: the tapped button hovered, at its rest #0071e3 (ciderui's way)", truthy(at(gated, "gated", "hover")) && at(gated, "gated", "background") == "rgb(0, 113, 227)"},
		{"no focus ring after a tap: :focus-visible on neither button", !truthy(at(plain, "plain", "focusVisible")) && !truthy(at(gated, "gated", "focusVisible"))},
		{"the menu's button opens its popover by a tap", at(menu, "after") == "menu-button" && at(menu, "menuOpen") == true},
		{"its close button, tapped, closes it", at(closed, "after") == "menu-close" && at(closed, "menuOpen") == false},
	}
	fmt.Println(label)
	pass := true
	for _, c := range claims {
		pass = pass && c.ok
		mark := "ok  "
		if !c.ok {
			mark = "FAIL"
		}
		fmt.Printf("%s %s\n", mark, c.name)
	}
	// @invariant: what is recorded beside the claims: where the hover goes with the next tap,
	// and the focus a tap gives
	fmt.Printf("seen: the plain button tapped, focused %s; the gated one tapped, the plain one still hovered %s, the gated one focused %s; somewhere else tapped, the gated one still hovered %s\n",
		text(at(plain, "plain", "focus")), text(at(gated, "plain", "hover")), text(at(gated, "gated", "focus")), text(at(elsewhere, "gated", "hover")))
	if !pass {
		fmt.Println(string(ojson.Marshal(all)))
		fmt.Println("FAIL")
		os.Exit(1)
	}
	fmt.Println("PASS")
	return nil
}

func center(page playwright.Page, id string) (float64, float64, error) {
	r, err := page.Locator("#" + id).BoundingBox()
	if err != nil {
		return 0, 0, err
	}
	return r.X + r.Width/2, r.Y + r.Height/2, nil
}

func viaDevTools(browser, url string, sent *reports) (label string, err error) {
	page, err := probe.Open(browser, url, nil)
	if err != nil {
		return "", err
	}
	defer func() {
		if cerr := page.Close(); err == nil {
			err = cerr
		}
	}()
	pw := page.Playwright()
	if err := pw.SetViewportSize(390, 844); err != nil {
		return "", err
	}
	cdp, err := pw.Context().NewCDPSession(pw)
	if err != nil {
		return "", err
	}
	if _, err := cdp.Send("Emulation.setTouchEmulationEnabled", map[string]any{"enabled": true, "maxTouchPoints": 5}); err != nil {
		return "", err
	}
	if _, err := pw.Reload(playwright.PageReloadOptions{WaitUntil: playwright.WaitUntilStateLoad}); err != nil {
		return "", err
	}
	for i, s := range steps {
		x, y, err := center(pw, s.id)
		if err != nil {
			return "", err
		}
		// @trap: plain maps and []any down to the leaves: playwright-go's argument
		// serializer panics on any other map
		if _, err := cdp.Send("Input.dispatchTouchEvent", map[string]any{"type": "touchStart", "touchPoints": []any{map[string]any{"x": x, "y": y}}}); err != nil {
			return "", err
		}
		if _, err := cdp.Send("Input.dispatchTouchEvent", map[string]any{"type": "touchEnd", "touchPoints": []any{}}); err != nil {
			return "", err
		}
		if err := sent.reported(i, s.id); err != nil {
			return "", err
		}
	}
	return page.Label() + ", the DevTools protocol's touch events", nil
}

func viaTouchscreen(browser, url string, sent *reports) (label string, err error) {
	page, err := probe.Open(browser, url, &playwright.BrowserNewPageOptions{Viewport: &playwright.Size{Width: 390, Height: 844}, HasTouch: playwright.Bool(true)})
	if err != nil {
		return "", err
	}
	defer func() {
		if cerr := page.Close(); err == nil {
			err = cerr
		}
	}()
	pw := page.Playwright()
	for i, s := range steps {
		x, y, err := center(pw, s.id)
		if err != nil {
			return "", err
		}
		// @trap: playwright-go's Touchscreen.Tap takes whole pixels where Node's took the
		// box's exact middle: rounded to the nearest
		if err := pw.Touchscreen().Tap(int(math.Round(x)), int(math.Round(y))); err != nil {
			return "", err
		}
		if err := sent.reported(i, s.id); err != nil {
			return "", err
		}
	}
	return page.Label() + ", Playwright's touchscreen taps", nil
}

// @invariant: a value as a JavaScript template prints it: undefined where it is none
func text(v any) string {
	switch v := v.(type) {
	case nil:
		return "undefined"
	case string:
		return v
	}
	b, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return string(b)
}

// @invariant: JavaScript's truthiness, for the values a report holds
func truthy(v any) bool {
	switch v := v.(type) {
	case nil:
		return false
	case bool:
		return v
	case string:
		return v != ""
	case float64:
		return v != 0 && !math.IsNaN(v)
	}
	return true
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

func list(v any) []any {
	a, _ := v.([]any)
	return a
}
