// Command search: BROWSER=safari|ios|chrome|chromium|firefox go run ./reference/search
// MEASURED.md holds its claims and method.
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

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
	raw, err := os.ReadFile(filepath.Join(probe.Dir("search"), "page.html"))
	if err != nil {
		return err
	}
	html := string(raw)
	l, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		return err
	}
	server := &http.Server{Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "text/html; charset=utf-8")
		if strings.Contains(r.RequestURI, "sheet") {
			w.Write([]byte(strings.Replace(html, `<html lang="en">`, `<html lang="en" class="sheet">`, 1)))
			return
		}
		w.Write(raw)
	})}
	go server.Serve(l)
	defer server.Close()
	origin := fmt.Sprintf("http://127.0.0.1:%d", l.Addr().(*net.TCPAddr).Port)
	src, err := os.ReadFile(filepath.Join(probe.Dir("search"), "search.js"))
	if err != nil {
		return err
	}
	ios := browser == "ios"
	keyboard := browser == "safari"
	page, err := probe.Open(browser, origin+"/", nil)
	if err != nil {
		return err
	}
	s := &steps{page: page, src: string(src), ios: ios, keyboard: keyboard, got: ojson.Object{}}
	err = s.take(origin)
	label := page.Label()
	if cerr := page.Close(); err == nil {
		err = cerr
	}
	if err != nil {
		return err
	}
	got := s.got
	if err := probe.Record("search", browser+".json", ojson.Object{{Key: "browser", Value: ojson.String(label)}, {Key: "date", Value: probe.Date()}}.With(got...)); err != nil {
		return err
	}

	g := js(got)
	o, p, sh := at(g, "opened"), at(g, "openedPixels"), at(g, "sheet")
	closed, closedPixels, scrolled := at(g, "closed"), at(g, "closedPixels"), at(g, "scrolled")
	near := func(a, b any) bool { return math.Abs(num(a)-num(b)) <= 0.5 }
	color := func(c any, want []float64) bool {
		ok := true
		for i, v := range list(c) {
			ok = ok && math.Abs(num(v)-want[i]) <= 2
		}
		return ok
	}
	type claim struct {
		name string
		ok   bool
	}
	by := func(cond bool, text string) string {
		if cond {
			return text
		}
		return ""
	}
	claims := []claim{
		{"the browser has anchor positioning", at(closed, "anchors") == true},
		{"closed, the bar's list is its row and the flyout is not shown", at(closed, "menuDisplay") == "flex" && at(closed, "open") == false && at(closed, "display") == "none"},
		{"the search button, in the bar's closed list, opens the flyout, with no script", at(o, "open") == true},
		{"the focus is on its field (autofocus)", at(o, "focus") == "input"},
		{"it hangs from the bar's bottom, under the banner, across the window", near(at(o, "flyout", "top"), at(o, "bar", "bottom")) && near(at(o, "bar", "bottom"), 114.0) && num(at(o, "flyout", "left")) == 0 && num(at(o, "flyout", "width")) == num(at(o, "pageWidth"))},
		{"its curtain fills the window under its panel", near(at(o, "curtain", "top"), at(o, "panel", "bottom")) && near(at(o, "curtain", "bottom"), at(o, "windowHeight"))},
		{"the page scrolled, the bar stuck at the top, the flyout follows it", near(at(scrolled, "bar", "top"), 0.0) && near(at(scrolled, "flyout", "top"), at(scrolled, "bar", "bottom")) && at(scrolled, "open") == true},
		{"in the picture the bar is its own color, not under the curtain, and the panel #fafafc", color(at(p, "bar"), []float64{200, 220, 240}) && color(at(closedPixels, "bar"), []float64{200, 220, 240}) && color(at(p, "panel"), []float64{250, 250, 252})},
		{"in the picture the curtain blurs the page under it (the stripes' spread under a fifth of theirs closed)", num(at(closedPixels, "spread")) > 60 && num(at(p, "spread")) < num(at(closedPixels, "spread"))/5},
		{"the curtain's middle hits the curtain, not the page; the bar hits its link", at(o, "curtainHit") == "curtain" && at(o, "barHit") == "link"},
		{"a popover's ::backdrop takes no click (pointer-events: none), so the curtain is a button of its own", at(o, "backdropPointer") == "none"},
	}
	if !ios {
		claims = append(claims, claim{"a click on the bar's link closes it and reaches the link", at(g, "afterLink", "open") == false && num(at(g, "afterLink", "clicks", "link")) == 1})
	}
	claims = append(claims, claim{"a click on the curtain closes it" + by(ios, " (by script)") + " and reaches nothing under it", at(g, "afterCurtain", "open") == false && num(at(g, "afterCurtain", "clicks", "under")) == 0})
	if !ios {
		claims = append(claims, claim{"Escape closes it, the focus back on the search button", at(g, "afterEscape", "open") == false && at(g, "afterEscape", "focus") == "search"})
	}
	claims = append(claims, claim{"its button closes it again", at(g, "afterToggle", "open") == false})
	if !ios {
		claims = append(claims, claim{"Return in its field sends the form to the site's search page (GET, ?q=)", at(g, "sent", "at") == "/search?q=iphone"})
	}
	claims = append(claims,
		claim{"as a sheet its button opens it over the window and the bar", at(sh, "open") == true && num(at(sh, "flyout", "top")) == 0 && num(at(sh, "flyout", "left")) == 0 && num(at(sh, "flyout", "width")) == num(at(sh, "pageWidth")) && near(at(sh, "flyout", "height"), at(sh, "windowHeight")) && at(sh, "barHit") != "link"},
		claim{"as a sheet the focus is on its field", at(sh, "focus") == "input"},
		claim{"as a sheet the page is held while it is open, and not after", at(sh, "rootOverflow") == "hidden" && at(g, "sheetClosed", "rootOverflow") != "hidden" && at(g, "sheetAfterClose", "rootOverflow") != "hidden"},
		claim{"as a sheet its close button closes it", at(g, "sheetAfterClose", "open") == false},
	)
	if !ios {
		claims = append(claims, claim{"as a sheet the focus is back on the search button", at(g, "sheetAfterClose", "focus") == "search"})
	}
	claims = append(claims,
		claim{"empty, its field's clear button is hidden", at(g, "empty", "reset") == "hidden" && at(g, "empty", "value") == ""},
		claim{"with text in the field, its clear button shows, with no script", at(g, "typed", "reset") == "visible"},
		claim{"its clear button (a reset button) clicked" + by(ios, " (by script)") + " empties the field and hides itself, the flyout still open", at(g, "cleared", "value") == "" && at(g, "cleared", "reset") == "hidden" && at(g, "cleared", "open") == true},
	)
	if !ios {
		claims = append(claims,
			claim{"then Shift+Tab takes the focus back to the field", at(g, "clearedBack", "focus") == "input"},
			claim{"and Escape still closes it", at(g, "clearedEscape", "open") == false},
		)
	}
	how := ""
	if ios {
		how = " (clicked by script; no key, no click outside)"
	} else if keyboard {
		how = " (its buttons by its keyboard: Return)"
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
		fmt.Println(string(ojson.Marshal(got)))
		fmt.Println("FAIL")
		os.Exit(1)
	}
	fmt.Println("PASS")
	return nil
}

type steps struct {
	page          probe.Page
	src           string
	ios, keyboard bool
	got           ojson.Object
}

func (s *steps) call(fn string) string {
	return "(async () => {\n" + s.src + "\nreturn " + fn + ";\n})()"
}

func (s *steps) record(key string, v ojson.Value, err error) error {
	if err == nil {
		s.got = s.got.With(ojson.Member{Key: key, Value: v})
	}
	return err
}

func (s *steps) state() (ojson.Value, error) { return s.page.Run(s.call("state()")) }

func (s *steps) wait(ms int) error { _, err := s.page.Run(probe.Wait(ms)); return err }

func (s *steps) run(expr string) error { _, err := s.page.Run(expr); return err }

// @external: a button: by script on the simulator, by Return in Safari, else clicked
func (s *steps) press(selector string) error {
	if s.ios {
		return probe.Script(s.page, selector, "click")
	}
	if s.keyboard {
		if err := probe.Script(s.page, selector, "focus"); err != nil {
			return err
		}
		return s.page.Press("Enter")
	}
	return s.page.Click(selector)
}

// @external: a click anywhere else: the browser's own, but on the simulator
func (s *steps) click(selector string) error {
	if s.ios {
		return probe.Script(s.page, selector, "click")
	}
	return s.page.Click(selector)
}

// @invariant: the offset is spliced in as JavaScript text: undefined on the first
// picture, then the first one's value
func (s *steps) pixels(offset ojson.Value) (ojson.Value, error) {
	shot, err := s.page.Shot()
	if err != nil {
		return nil, err
	}
	png, err := json.Marshal(shot)
	if err != nil {
		return nil, err
	}
	arg := "undefined"
	if offset != nil {
		arg = compact(offset)
	}
	return s.page.Run(s.call(fmt.Sprintf("pixels(%s, %s)", png, arg)))
}

func (s *steps) opened() error {
	if err := s.press("#search"); err != nil {
		return err
	}
	return s.wait(900)
}

func (s *steps) take(origin string) error {
	if !s.ios {
		if err := s.page.Size(1440, 900); err != nil {
			return err
		}
	}
	if err := s.wait(500); err != nil {
		return err
	}
	if err := s.run(s.call("arm()")); err != nil {
		return err
	}
	if err := s.record(s.state1("closed")); err != nil {
		return err
	}
	v, err := s.pixels(nil)
	if err := s.record("closedPixels", v, err); err != nil {
		return err
	}
	if err := s.opened(); err != nil {
		return err
	}
	if err := s.record(s.state1("opened")); err != nil {
		return err
	}
	closedPixels, _ := s.got.Get("closedPixels").(ojson.Object)
	v, err = s.pixels(closedPixels.Get("offset"))
	if err := s.record("openedPixels", v, err); err != nil {
		return err
	}
	if err := s.run(`Promise.resolve((scrollTo(0, 200), true))`); err != nil {
		return err
	}
	if err := s.wait(400); err != nil {
		return err
	}
	if err := s.record(s.state1("scrolled")); err != nil {
		return err
	}
	if err := s.run(`Promise.resolve((scrollTo(0, 0), true))`); err != nil {
		return err
	}
	if err := s.wait(400); err != nil {
		return err
	}
	if !s.ios {
		if err := s.click("#link"); err != nil {
			return err
		}
		if err := s.wait(400); err != nil {
			return err
		}
		if err := s.record(s.state1("afterLink")); err != nil {
			return err
		}
		if err := s.opened(); err != nil {
			return err
		}
	}
	if err := s.click("#curtain"); err != nil {
		return err
	}
	if err := s.wait(400); err != nil {
		return err
	}
	if err := s.record(s.state1("afterCurtain")); err != nil {
		return err
	}
	if !s.ios {
		if err := s.opened(); err != nil {
			return err
		}
		if err := s.page.Press("Escape"); err != nil {
			return err
		}
		if err := s.wait(400); err != nil {
			return err
		}
		if err := s.record(s.state1("afterEscape")); err != nil {
			return err
		}
	}
	if err := s.opened(); err != nil {
		return err
	}
	if err := s.press("#search"); err != nil {
		return err
	}
	if err := s.wait(400); err != nil {
		return err
	}
	if err := s.record(s.state1("afterToggle")); err != nil {
		return err
	}
	// @external: text in its field (set by the page's script: what the CSS reads is the
	// value), then its clear button clicked
	if err := s.opened(); err != nil {
		return err
	}
	if err := s.record(s.state1("empty")); err != nil {
		return err
	}
	if err := s.run(`Promise.resolve((document.querySelector('#input').value = 'iphone', true))`); err != nil {
		return err
	}
	if err := s.wait(400); err != nil {
		return err
	}
	if err := s.record(s.state1("typed")); err != nil {
		return err
	}
	if err := s.click("#reset"); err != nil {
		return err
	}
	if err := s.wait(400); err != nil {
		return err
	}
	if err := s.record(s.state1("cleared")); err != nil {
		return err
	}
	if !s.ios {
		// @external: the focus's way back to the field: Shift+Tab, from where it was
		if err := s.page.Press("Shift+Tab"); err != nil {
			return err
		}
		if err := s.wait(200); err != nil {
			return err
		}
		if err := s.record(s.state1("clearedBack")); err != nil {
			return err
		}
		if err := s.page.Press("Escape"); err != nil {
			return err
		}
		if err := s.wait(400); err != nil {
			return err
		}
		if err := s.record(s.state1("clearedEscape")); err != nil {
			return err
		}
	}
	if !s.ios {
		if err := s.opened(); err != nil {
			return err
		}
		if err := s.run(`Promise.resolve((document.querySelector('#input').value = 'iphone', true))`); err != nil {
			return err
		}
		if err := s.page.Press("Enter"); err != nil {
			return err
		}
		// @external: the page it loads: waited for here, not in the page it leaves
		time.Sleep(1500 * time.Millisecond)
		if err := s.record(s.state1("sent")); err != nil {
			return err
		}
	}
	if err := s.page.Go(origin + "/?sheet"); err != nil {
		return err
	}
	if err := s.wait(500); err != nil {
		return err
	}
	if err := s.record(s.state1("sheetClosed")); err != nil {
		return err
	}
	if err := s.opened(); err != nil {
		return err
	}
	if err := s.record(s.state1("sheet")); err != nil {
		return err
	}
	if err := s.press("#close"); err != nil {
		return err
	}
	if err := s.wait(400); err != nil {
		return err
	}
	return s.record(s.state1("sheetAfterClose"))
}

func (s *steps) state1(key string) (string, ojson.Value, error) {
	v, err := s.state()
	return key, v, err
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

func list(v any) []any {
	a, _ := v.([]any)
	return a
}
