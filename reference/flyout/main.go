// Command flyout: BROWSER=safari|chrome|chromium|firefox|ipad|iphone go run ./reference/flyout:
// MEASURED.md holds its claims and method.
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"slices"
	"strconv"
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

func quoted(s string) string {
	b, err := json.Marshal(s)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func run() error {
	browser := probe.Browser()
	html, err := os.ReadFile(filepath.Join(probe.Dir("flyout"), "page.html"))
	if err != nil {
		return err
	}
	var mu sync.Mutex
	var reports ojson.Array
	l, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		return err
	}
	server := &http.Server{Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" && r.URL.RequestURI() == "/report" {
			body, err := io.ReadAll(r.Body)
			if err != nil {
				panic(err)
			}
			v, err := ojson.Parse(body)
			if err != nil {
				panic(err)
			}
			mu.Lock()
			reports = append(reports, v)
			mu.Unlock()
			w.WriteHeader(http.StatusNoContent)
			return
		}
		w.Header().Set("content-type", "text/html; charset=utf-8")
		w.Write(html)
	})}
	go server.Serve(l)
	url := fmt.Sprintf("http://127.0.0.1:%d/", l.Addr().(*net.TCPAddr).Port)

	var label, tip string
	var got ojson.Object
	switch browser {
	case "ipad":
		var res xcuitest.Result
		res, err = xcuitest.Taps(url, []string{"Store menu", "Store curtain"}, xcuitest.Device{Type: "iPad-Pro-13-inch-M5-12GB", Name: "iPad Pro 13-inch (M5)"})
		label, tip = res.Label, res.Tip
	case "iphone":
		var res xcuitest.Result
		res, err = xcuitest.Taps(url, []string{"Menu", "Mac menu", "Back", "Mac menu", "Close all"}, xcuitest.IPhone17)
		label, tip = res.Label, res.Tip
	default:
		label, got, err = desktop(browser, url)
	}
	server.Close()
	if err != nil {
		return err
	}
	mu.Lock()
	defer mu.Unlock()
	record := ojson.Object{{Key: "browser", Value: ojson.String(label)}, {Key: "date", Value: probe.Date()}}
	if tip != "" {
		record = record.With(ojson.Member{Key: "tip", Value: ojson.String(tip)})
	}
	record = record.With(got...)
	if len(reports) > 0 && got.Get("closed") == nil {
		record = record.With(ojson.Member{Key: "reports", Value: reports})
	}
	if err := probe.Record("flyout", browser+".json", record); err != nil {
		return err
	}

	near := func(a, b float64) bool { return math.Abs(a-b) <= 0.5 }
	color := func(c ojson.Value, want []float64) bool {
		a, _ := c.(ojson.Array)
		if a == nil {
			return false
		}
		for i, v := range a {
			if i >= len(want) || !(math.Abs(num(v)-want[i]) <= 2) {
				return false
			}
		}
		return true
	}
	// @external: the first frame a flyout shows, from the pointer's entering its item
	shownAfter := func(w ojson.Value, id string) ojson.Value {
		var entered, shown ojson.Value
		for _, e := range array(get(w, "entered")) {
			if is(get(e, 0), id) {
				entered = get(e, 1)
				break
			}
		}
		for _, f := range array(get(w, "frames")) {
			if slices.ContainsFunc(array(get(f, 1)), func(v ojson.Value) bool { return is(v, id) }) {
				shown = get(f, 0)
				break
			}
		}
		if entered == nil || shown == nil {
			return ojson.Null{}
		}
		return ojson.Float(num(shown) - num(entered))
	}
	type claim struct {
		text string
		ok   bool
	}
	var claims []claim
	switch browser {
	case "ipad":
		if len(reports) < 2 {
			return fmt.Errorf("%d reports, not 2", len(reports))
		}
		opened, curtained := reports[0], reports[1]
		noHover := true
		for _, r := range reports {
			noHover = noHover && !truthy(get(r, "media", "hover"))
		}
		claims = []claim{
			{"the iPad's screen has no hover: (hover: hover) does not match", noHover},
			{"a tap on Store's link (its button over it) opens its flyout as a popover, from the bar's bottom", truthy(get(opened, "store", "open")) && near(num(get(opened, "store", "box", "top")), num(get(opened, "bar", "bottom")))},
			{"the link is not followed", is(get(opened, "hash"), "") && num(get(opened, "clicks", "followed")) == 0},
			{"a tap on its curtain closes it and reaches nothing under it", !truthy(get(curtained, "store", "open")) && num(get(curtained, "clicks", "under")) == 0 && is(get(curtained, "hash"), "")},
		}
	case "iphone":
		if len(reports) < 5 {
			return fmt.Errorf("%d reports, not 5", len(reports))
		}
		menu, flyout, back, closed := reports[0], reports[1], reports[2], reports[4]
		claims = []claim{
			{"the menu's button opens the menu", truthy(get(menu, "menu")) && !truthy(get(menu, "mac", "open"))},
			{"a tap on Mac's row (its button over it) opens Mac's flyout over the open menu, across the window", truthy(get(flyout, "menu")) && truthy(get(flyout, "mac", "open")) && num(get(flyout, "mac", "box", "top")) == 0 && near(num(get(flyout, "mac", "box", "height")), num(get(flyout, "window", "height"))) && is(get(flyout, "hash"), "")},
			{"its back button closes the flyout, the menu still open", truthy(get(back, "menu")) && !truthy(get(back, "mac", "open"))},
			{"its close button closes both", !truthy(get(closed, "menu")) && !truthy(get(closed, "mac", "open"))},
		}
	default:
		closed, hovered, switched, left, swept, popover, tabbed, escaped, curtained := got.Get("closed"), got.Get("hovered"), got.Get("switched"), got.Get("left"), got.Get("swept"), got.Get("popover"), got.Get("tabbed"), got.Get("escaped"), got.Get("curtained")
		p := got.Get("hoveredPixels")
		closedPixels := got.Get("closedPixels")
		after := shownAfter(got.Get("hoveredWatch"), "store")
		afterMs := num(after)
		switchedWatch := array(get(got, "switchedWatch", "frames"))
		everyFull := len(switchedWatch) > 0
		for _, f := range switchedWatch {
			everyFull = everyFull && len(array(get(f, 1))) > 0 && num(get(f, 2)) >= num(get(switched, "mac", "panel", "bottom"))-0.5
		}
		sweptEntered := array(get(got, "sweptWatch", "entered"))
		var enteredIDs []string
		for _, e := range sweptEntered {
			s, _ := get(e, 0).(ojson.String)
			enteredIDs = append(enteredIDs, string(s))
		}
		sweptFrames := array(get(got, "sweptWatch", "frames"))
		noneShown := len(sweptEntered) >= 2 && len(sweptFrames) > 0
		for _, f := range sweptFrames {
			noneShown = noneShown && len(array(get(f, 1))) == 0
		}
		phoneMenu, phoneFlyout, phoneBack, phoneClosed := got.Get("phoneMenu"), got.Get("phoneFlyout"), got.Get("phoneBack"), got.Get("phoneClosed")
		claims = []claim{
			{"the browser has anchor positioning and a pointer that hovers", truthy(get(closed, "media", "anchors")) && truthy(get(closed, "media", "hover"))},
			{"closed, no flyout shows, the curtain is hidden, the bar blurred", !truthy(get(closed, "store", "shown")) && !truthy(get(closed, "mac", "shown")) && is(get(closed, "curtain", "visibility"), "hidden") && !is(get(closed, "barFilter"), "none")},
			{"closed, each flyout is the bar's height, nothing of it laid out below the bar", near(num(get(closed, "store", "box", "height")), num(get(closed, "bar", "height"))) && near(num(get(closed, "mac", "box", "height")), num(get(closed, "bar", "height")))},
			{fmt.Sprintf("the pointer resting on Store's link shows its flyout, no popover opened, 100 to 250ms after it reached the link (%s)", string(ojson.Marshal(after))+"ms"), truthy(get(hovered, "store", "shown")) && !truthy(get(hovered, "store", "open")) && afterMs >= 100 && afterMs <= 250},
			{"it hangs from the bar's top, across the window", near(num(get(hovered, "store", "box", "top")), num(get(hovered, "bar", "top"))) && num(get(hovered, "store", "box", "left")) == 0 && near(num(get(hovered, "store", "box", "width")), num(get(hovered, "window", "page")))},
			{"the bar's links are over it", is(get(hovered, "macHit"), "mac")},
			{"the curtain shows, fixed, and the bar's blur is off", is(get(hovered, "curtain", "visibility"), "visible") && is(get(hovered, "curtain", "opacity"), "1") && is(get(hovered, "curtain", "position"), "fixed") && is(get(hovered, "barFilter"), "none")},
			{"in the picture the curtain blurs the page under it (the stripes' spread under a fifth of theirs closed), and the bar shows the flyout's ground", num(get(closedPixels, "spread")) > 60 && num(get(p, "spread")) < num(get(closedPixels, "spread"))/5 && color(get(p, "bar"), []float64{250, 250, 252})},
			{"the pointer moved on to Mac's link shows Mac's flyout and not Store's", truthy(get(switched, "mac", "shown")) && !truthy(get(switched, "store", "shown"))},
			{"and in every frame between, a flyout shows to its full height (the one left kept until the next is shown)", everyFull},
			{"the pointer off the bar to the page hides it, the curtain hidden, the bar blurred again", !truthy(get(left, "store", "shown")) && !truthy(get(left, "mac", "shown")) && is(get(left, "curtain", "visibility"), "hidden") && !is(get(left, "barFilter"), "none")},
			{"and it is the bar's height again", near(num(get(left, "mac", "box", "height")), num(get(left, "bar", "height")))},
			{fmt.Sprintf("swept along the bar past its links (%s entered), no flyout shows in any frame", strings.Join(enteredIDs, ", ")), noneShown && !truthy(get(swept, "store", "shown")) && !truthy(get(swept, "mac", "shown"))},
			{"Mac's button opens its flyout as a popover, from the bar's bottom, its curtain to the window's bottom", truthy(get(popover, "mac", "open")) && near(num(get(popover, "mac", "box", "top")), num(get(popover, "bar", "bottom"))) && near(num(get(popover, "mac", "curtain", "bottom")), num(get(popover, "window", "height"))) && is(get(popover, "mac", "curtainOpacity"), "1")},
			{"Tab takes the focus to its first link", is(get(tabbed, "focus"), "mac-first")},
			{"Escape closes it, the focus back on its button", !truthy(get(escaped, "mac", "open")) && is(get(escaped, "focus"), "mac-button")},
			{"a click on its curtain closes it and reaches nothing under it", !truthy(get(curtained, "mac", "open")) && num(get(curtained, "clicks", "under")) == 0},
			{"at a phone's width the menu's button opens the menu", truthy(get(phoneMenu, "menu")) && !truthy(get(phoneMenu, "mac", "open"))},
			{"Mac's row's button opens Mac's flyout over the open menu, across the window", truthy(get(phoneFlyout, "menu")) && truthy(get(phoneFlyout, "mac", "open")) && num(get(phoneFlyout, "mac", "box", "top")) == 0 && near(num(get(phoneFlyout, "mac", "box", "height")), num(get(phoneFlyout, "window", "height")))},
			{"its back button closes the flyout, the menu still open", truthy(get(phoneBack, "menu")) && !truthy(get(phoneBack, "mac", "open"))},
			{"its close button closes both", !truthy(get(phoneClosed, "menu")) && !truthy(get(phoneClosed, "mac", "open"))},
		}
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
		fmt.Println("FAIL")
		os.Exit(1)
	}
	fmt.Println("PASS")
	return nil
}

func desktop(browser, url string) (string, ojson.Object, error) {
	page, err := probe.Open(browser, url, nil)
	if err != nil {
		return "", nil, err
	}
	got, err := desktopSteps(page, browser)
	label := page.Label()
	if cerr := page.Close(); err == nil {
		err = cerr
	}
	return label, got, err
}

func desktopSteps(page probe.Page, browser string) (ojson.Object, error) {
	// @external: a button by the keyboard in Safari (a click there focuses none), else clicked
	keyboard := browser == "safari"
	var got ojson.Object
	var errs []error
	// @invariant: the first error stops every step after it, as a throw does; steps is
	// the error
	run := func(expr string) ojson.Value {
		if len(errs) > 0 {
			return nil
		}
		v, err := page.Run(expr)
		if err != nil {
			errs = append(errs, err)
		}
		return v
	}
	do := func(err error) {
		if len(errs) == 0 && err != nil {
			errs = append(errs, err)
		}
	}
	set := func(key string, v ojson.Value) {
		if len(errs) == 0 {
			got = got.With(ojson.Member{Key: key, Value: v})
		}
	}
	wait := func(ms int) { run(fmt.Sprintf("new Promise(r => setTimeout(() => r(true), %d))", ms)) }
	state := func() ojson.Value { return run("Promise.resolve(window.probe.state())") }
	watch := func() ojson.Value { return run("Promise.resolve(window.probe.watch())") }
	stop := func() ojson.Value { return run("Promise.resolve(window.probe.stop())") }
	pixels := func(offset ojson.Value, given bool) ojson.Value {
		if len(errs) > 0 {
			return nil
		}
		shot, err := page.Shot()
		if err != nil {
			errs = append(errs, err)
			return nil
		}
		arg := ""
		if given {
			arg = ", " + interpolated(offset)
		}
		return run("window.probe.pixels(" + quoted(shot) + arg + ")")
	}
	mid := func(selector string) (float64, float64) {
		v := run("Promise.resolve((r => [r.left + r.width / 2, r.top + r.height / 2])(document.querySelector(" + quoted(selector) + ").getBoundingClientRect()))")
		return num(get(v, 0)), num(get(v, 1))
	}
	move := func(x, y, ms float64) {
		if len(errs) == 0 {
			do(page.Move(x, y, ms))
		}
	}
	press := func(selector string) {
		if len(errs) > 0 {
			return
		}
		if !keyboard {
			do(page.Click(selector))
			return
		}
		run("Promise.resolve((document.querySelector(" + quoted(selector) + ").focus(), true))")
		if len(errs) == 0 {
			do(page.Press("Enter"))
		}
	}
	key := func(k string) {
		if len(errs) == 0 {
			do(page.Press(k))
		}
	}

	do(page.Size(1440, 900))
	wait(800)
	// @external: a point on the page near the window's foot (Safari's window, its frame
	// included, is 900px tall: its page is less)
	foot := num(run("Promise.resolve(innerHeight - 50)"))
	move(700, foot, 100)
	wait(500)
	set("closed", state())
	set("closedPixels", pixels(nil, false))
	offset := get(got, "closedPixels", "offset")
	// @external: the pointer to Store's link, resting there
	watch()
	x, y := mid("#store")
	move(x, y, 150)
	wait(800)
	set("hovered", state())
	set("hoveredWatch", stop())
	set("hoveredPixels", pixels(offset, true))
	// @external: on along the bar to Mac's
	watch()
	x, y = mid("#mac")
	move(x, y, 150)
	wait(900)
	set("switched", state())
	set("switchedWatch", stop())
	// @external: off the bar, through Mac's flyout, down to the page
	watch()
	move(700, foot, 150)
	wait(900)
	set("left", state())
	set("leftWatch", stop())
	// @external: swept along the bar past each link in 90ms, from before the first to
	// after the last, then down to the page
	x0, y0 := mid("#store")
	x1, _ := mid("#search")
	move(x0-40, y0, 100)
	wait(500)
	watch()
	move(x1+60, y0, 90)
	move(x1+60, foot, 50)
	wait(800)
	set("swept", state())
	set("sweptWatch", stop())
	// @external: Mac's button: its flyout a popover
	press("#mac-button")
	wait(600)
	set("popover", state())
	if keyboard {
		key("Alt+Tab")
	} else {
		key("Tab")
	}
	wait(200)
	set("tabbed", state())
	key("Escape")
	wait(400)
	set("escaped", state())
	press("#mac-button")
	wait(600)
	if len(errs) == 0 {
		do(page.Click("#mac-curtain"))
	}
	wait(400)
	set("curtained", state())
	// @external: a phone's layout
	move(10, foot, 50)
	if len(errs) == 0 {
		do(page.Size(700, 900))
	}
	wait(600)
	press("#menu-button")
	wait(600)
	set("phoneMenu", state())
	press("#mac-button")
	wait(600)
	set("phoneFlyout", state())
	press("#mac-back")
	wait(600)
	set("phoneBack", state())
	press("#mac-button")
	wait(600)
	press("#mac-close")
	wait(600)
	set("phoneClosed", state())
	if len(errs) > 0 {
		return nil, errs[0]
	}
	return got, nil
}

// @external: a value as a JavaScript template literal writes it: a string bare, a
// number, a boolean or null as its text
func interpolated(v ojson.Value) string {
	switch v := v.(type) {
	case ojson.String:
		return string(v)
	case ojson.Number:
		return string(v)
	case nil:
		return "undefined"
	}
	return string(ojson.Marshal(v))
}

func array(v ojson.Value) ojson.Array {
	a, _ := v.(ojson.Array)
	return a
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
