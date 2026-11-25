// Command select: BROWSER=safari|chrome|chromium|firefox|ios go run ./reference/select
// MEASURED.md holds its claims and method.
package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/at-least/ciderui/internal/ojson"
	"github.com/at-least/ciderui/internal/probe"
	"github.com/at-least/ciderui/internal/xcuitest"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// @invariant: what the page sent to /report, in the order it came
type reports struct {
	mu   sync.Mutex
	list []ojson.Value
}

func (r *reports) since(from int) ojson.Array {
	r.mu.Lock()
	defer r.mu.Unlock()
	return append(ojson.Array{}, r.list[from:]...)
}

func (r *reports) len() int {
	r.mu.Lock()
	defer r.mu.Unlock()
	return len(r.list)
}

func run() error {
	browser := probe.Browser()
	ios := browser == "ios"
	raw, err := os.ReadFile(filepath.Join(probe.Dir("select"), "page.html"))
	if err != nil {
		return err
	}
	html := string(raw)
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
		if strings.Contains(r.RequestURI, "plain") {
			w.Write([]byte(strings.Replace(html, `<html lang="en">`, `<html lang="en" class="plain">`, 1)))
			return
		}
		w.Write(raw)
	})}
	go func() {
		if err := server.Serve(l); !errors.Is(err, http.ErrServerClosed) {
			fmt.Fprintln(os.Stderr, "serve:", err)
		}
	}()
	defer server.Close()
	origin := fmt.Sprintf("http://127.0.0.1:%d", l.Addr().(*net.TCPAddr).Port)

	got := ojson.Object{}
	put := func(keys []string, values ojson.Array) {
		for i, k := range keys {
			if i < len(values) {
				got = got.With(ojson.Member{Key: k, Value: values[i]})
			}
		}
	}
	var label, tip string
	if ios {
		// @external: the same steps on each page, each on a simulator of its own
		steps := []string{"Search", "type:mac", "More", "press:Explore"}
		onIOS := func(path string) (xcuitest.Result, ojson.Array, error) {
			from := sent.len()
			result, err := xcuitest.Taps(origin+path, steps, xcuitest.IPhone17)
			if err != nil {
				fmt.Printf("reported before the failure: %s\n", ojson.Marshal(sent.since(from)))
				return xcuitest.Result{}, nil, err
			}
			return result, sent.since(from), nil
		}
		_, plain, err := onIOS("/?plain")
		if err != nil {
			return err
		}
		put([]string{"plainOpened", "plainTyped", "plainMore", "plainExplore"}, plain)
		result, bar, err := onIOS("/")
		if err != nil {
			return err
		}
		label, tip = result.Label, result.Tip
		put([]string{"opened", "typed", "more", "explore"}, bar)
	} else {
		page, err := probe.Open(browser, origin+"/?plain", nil)
		if err != nil {
			return err
		}
		label = page.Label()
		err = inBrowser(page, browser, origin, &got)
		if cerr := page.Close(); err == nil {
			err = cerr
		}
		if err != nil {
			return err
		}
	}
	record := ojson.Object{{Key: "browser", Value: ojson.String(label)}, {Key: "date", Value: probe.Date()}}
	if tip != "" {
		record = record.With(ojson.Member{Key: "tip", Value: ojson.String(tip)})
	}
	if err := probe.Record("select", browser+".json", record.With(got...)); err != nil {
		return err
	}

	g := js(got)
	type claim struct {
		name string
		ok   bool
	}
	var claims []claim
	if ios {
		moved := true
		for _, r := range []any{at(g, "plainMore"), at(g, "more")} {
			moved = moved && at(r, "open") == false && at(r, "moreOpen") == true && at(r, "focus") != "input"
		}
		claims = []claim{
			{"the search button, tapped, opens its flyout, the focus on its field", at(g, "opened", "open") == true && at(g, "opened", "focus") == "input"},
			{"XCUITest's keyboard types into the field (\"mac\"), the bar's user-select: none over it", at(g, "typed", "value") == "mac"},
			{"the link's flyout's button, tapped, closes the search and opens its flyout, the field left", moved},
			{"where the bar can be selected, a long press on the flyout's heading selects its text", trim(at(g, "plainExplore", "selection")) == "Explore"},
			{"where it cannot, the same long press selects nothing", at(g, "explore", "selection") == ""},
		}
	} else {
		bar := at(g, "rest", "userSelect", "bar", 0)
		if s, _ := bar.(string); s == "" {
			bar = at(g, "rest", "userSelect", "bar", 1)
		}
		has := func(key, text string) bool {
			s, _ := at(g, key, "selection").(string)
			return strings.Contains(s, text)
		}
		claims = []claim{
			{"the bar computes user-select: none (-webkit-user-select where the browser has no other)", bar == "none"},
			{"where the bar can be selected, Select All on the page takes the bar's text with the page's", has("plainAll", "Store") && has("plainAll", "Selectable")},
			{"where it cannot, Select All takes the page's text alone", !has("all", "Store") && has("all", "Selectable")},
			{"where the bar can be selected, a double-click on a link's flyout's heading selects its text", trim(at(g, "plainExplore", "selection")) == "Explore"},
			{"where it cannot, the same double-click selects nothing", at(g, "more", "moreOpen") == true && at(g, "explore", "selection") == ""},
			{"the search button opens its flyout, the focus on its field", at(g, "opened", "open") == true && at(g, "opened", "moreOpen") == false && at(g, "opened", "focus") == "input"},
			{"where the bar can be selected, a double-click on the search's heading selects its text", trim(at(g, "plainHeading", "selection")) == "Shortcuts"},
			{"where it cannot, the same double-click selects nothing", at(g, "heading", "selection") == ""},
			{"the field clicked, keys typed go into it (\"mac\")", at(g, "focused", "focus") == "input" && at(g, "typed", "value") == "mac"},
			{"Select All (the command key and A) in the field selects its text whole", num(at(g, "selectedAll", "fieldSelection", 0)) == 0 && num(at(g, "selectedAll", "fieldSelection", 1)) == 3},
		}
	}
	how := ""
	if !ios && browser == "safari" {
		how = " (the search button by its keyboard: Return)"
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
	seen, _ := got.Get("rest").(ojson.Object)
	if seen == nil {
		seen, _ = got.Get("opened").(ojson.Object)
	}
	fmt.Printf("seen: [user-select, -webkit-user-select] %s; supported %s\n", compact(seen.Get("userSelect")), compact(seen.Get("supports")))
	if !pass {
		fmt.Println(string(ojson.Marshal(got)))
		fmt.Println("FAIL")
		os.Exit(1)
	}
	fmt.Println("PASS")
	return nil
}

func inBrowser(page probe.Page, browser, origin string, got *ojson.Object) error {
	if err := page.Size(1440, 900); err != nil {
		return err
	}
	state := func() (ojson.Value, error) { return page.Run("Promise.resolve(window.cuiState())") }
	// @external: a button: by Return in Safari, else clicked
	press := func(selector string) func() error {
		return func() error {
			if browser == "safari" {
				if _, err := page.Run(fmt.Sprintf("Promise.resolve((document.querySelector('%s').focus(), true))", selector)); err != nil {
					return err
				}
				return page.Press("Enter")
			}
			return page.Click(selector)
		}
	}
	step := func(key string, act func() error, ms int) error {
		if err := act(); err != nil {
			return err
		}
		if _, err := page.Run(probe.Wait(ms)); err != nil {
			return err
		}
		v, err := state()
		if err != nil {
			return err
		}
		if key != "" {
			*got = got.With(ojson.Member{Key: key, Value: v})
		}
		return nil
	}
	key := func(k string) func() error { return func() error { return page.Press(k) } }
	dblclick := func(selector string) func() error { return func() error { return page.Dblclick(selector) } }
	clear := func() error {
		_, err := page.Run("Promise.resolve((getSelection().removeAllRanges(), true))")
		return err
	}
	if _, err := page.Run(probe.Wait(500)); err != nil {
		return err
	}
	// @external: the bar selectable
	if err := step("plainAll", key("ControlOrMeta+a"), 300); err != nil {
		return err
	}
	if err := clear(); err != nil {
		return err
	}
	if err := step("", press("#more"), 600); err != nil {
		return err
	}
	if err := step("plainExplore", dblclick("#explore"), 300); err != nil {
		return err
	}
	if err := step("", press("#search"), 600); err != nil {
		return err
	}
	if err := step("plainHeading", dblclick("#heading"), 300); err != nil {
		return err
	}
	// @external: the bar not selectable
	if err := page.Go(origin + "/"); err != nil {
		return err
	}
	if _, err := page.Run(probe.Wait(500)); err != nil {
		return err
	}
	v, err := state()
	if err != nil {
		return err
	}
	*got = got.With(ojson.Member{Key: "rest", Value: v})
	if err := step("all", key("ControlOrMeta+a"), 300); err != nil {
		return err
	}
	if err := clear(); err != nil {
		return err
	}
	for _, s := range []struct {
		key string
		act func() error
		ms  int
	}{
		{"more", press("#more"), 600},
		{"explore", dblclick("#explore"), 300},
		{"opened", press("#search"), 600},
		{"heading", dblclick("#heading"), 300},
		{"focused", func() error { return page.Click("#input") }, 300},
		{"typed", func() error { return page.Type("mac") }, 300},
		{"selectedAll", key("ControlOrMeta+a"), 300},
	} {
		if err := step(s.key, s.act, s.ms); err != nil {
			return err
		}
	}
	return nil
}

// @invariant: String.prototype.trim of a string; undefined (nil) where it is none,
// where JavaScript would have thrown
func trim(v any) any {
	s, ok := v.(string)
	if !ok {
		return nil
	}
	return strings.TrimSpace(s)
}

// @invariant: JSON.stringify(v), with no indent, its keys in order
func compact(v ojson.Value) string {
	var b bytes.Buffer
	if err := json.Compact(&b, ojson.Marshal(v)); err != nil {
		panic(err)
	}
	return b.String()
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
