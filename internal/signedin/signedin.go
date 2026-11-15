// Package signedin is the command line of Apple's signed-in sites (App Store
// Connect's, and Apple Account's, the developer account's and Feedback Assistant's,
// each a reference/<site> command calling CLI with its own specs): their pages read
// in a Chrome the owner signed in to (no headless browser can sign in), which hold
// the account's data.
//
//	go run ./reference/<site> --login         opens Google Chrome on the site with
//	                                          Profile, the one profile every
//	                                          signed-in site shares, for the owner
//	                                          to sign in; it ends when Chrome quits
//	go run ./reference/<site> --run <spec>    reads the spec in that profile
//	                                          through Playwright, then saves it
//	go run ./reference/<site> <spec>          prints the script to run in a page
//	                                          by hand (Claude in Chrome's
//	                                          javascript tool)
//	go run ./reference/<site> --save <spec> <file>
//	                                          checks <file>, what the page handed
//	                                          back (its result and sha256, one
//	                                          object; the file TOKENS downloads),
//	                                          and, if it passes, writes it as
//	                                          results/<spec>.go and the registry,
//	                                          results/pages.go
//	go run ./reference/<site> --check <spec>  holds results/<spec>.go to what the
//	                                          page handed back and to the
//	                                          account's privacy
//
// Profile is outside the repository: it holds the owner's session. ids.json, beside
// it, holds each site's ids a path's <app> and <team> stand for (written by the
// owner or by an agent from the signed-in page), so no id reaches the repository.
// A spec's Hover (a selector) rests the pointer on that element and Click clicks it
// before the read, by Playwright's pointer. A read is at 1920 by 907 at 2 device
// pixels, the window the head spec was read in.
package signedin

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/at-least/ciderui/internal/ojson"
	"github.com/at-least/ciderui/internal/probe"
	"github.com/at-least/ciderui/internal/reference"
	"github.com/mxschmitt/playwright-go"
)

// Probe is one element a spec reads; its JSON (the page's argument, and what the
// probes digest is of) has the members set, in this order.
type Probe struct {
	Name   string
	Shadow string
	Select string
	Text   string
	Nth    int
	Up     int
	Unseen bool
	Pseudo string
	Props  []string
	Box    bool
}

// TokenSpec names the tokens a spec reads: those beginning with one of Prefixes and
// matching none of Exclude (regular expressions, the page's).
type TokenSpec struct {
	Prefixes []string
	Exclude  []string
}

// Spec is one page's read: Path (its <app> and <team> filled from ids.json), and
// either Probes (read by Reader) or Tokens (read by the Tokens reader). Wait is the
// milliseconds after the page's text appears, 3000 if 0.
type Spec struct {
	Path   string
	Wait   int
	Click  string
	Hover  string
	Probes []Probe
	Tokens *TokenSpec
}

// The property lists the specs read, apple.com's specs' own.
var (
	Type = []string{"font-family", "font-size", "line-height", "font-weight", "letter-spacing", "color"}
	Box  = []string{"padding-top", "padding-right", "padding-bottom", "padding-left", "margin-top", "margin-bottom",
		"border-top-width", "border-top-style", "border-top-color", "border-bottom-width", "border-bottom-style", "border-bottom-color",
		"border-top-left-radius", "background-color", "box-shadow", "min-width", "min-height", "max-width"}
	Link = []string{"text-decoration-line", "text-decoration-color", "text-underline-offset"}
	// Page is the page's own: its text, its ground, how its text is drawn.
	Page = slices.Concat(Type, []string{"background-color", "-webkit-font-smoothing", "font-synthesis", "text-rendering"})
	// Text is a control's text: how it sits, wraps and shows a link.
	Text = []string{"display", "text-align", "white-space", "text-decoration-line", "cursor", "vertical-align", "opacity"}
	// Flex is a flex or grid box's own layout.
	Flex = []string{"display", "flex-direction", "flex-wrap", "align-items", "justify-content", "row-gap", "column-gap"}
)

// JSON is the probe as the spec's object literal writes it.
func (p Probe) JSON() ojson.Object {
	o := ojson.Object{{Key: "name", Value: ojson.String(p.Name)}}
	add := func(key string, v ojson.Value) { o = append(o, ojson.Member{Key: key, Value: v}) }
	if p.Shadow != "" {
		add("shadow", ojson.String(p.Shadow))
	}
	add("select", ojson.String(p.Select))
	if p.Text != "" {
		add("text", ojson.String(p.Text))
	}
	if p.Nth != 0 {
		add("nth", ojson.Float(float64(p.Nth)))
	}
	if p.Up != 0 {
		add("up", ojson.Float(float64(p.Up)))
	}
	if p.Unseen {
		add("unseen", ojson.Bool(true))
	}
	if p.Pseudo != "" {
		add("pseudo", ojson.String(p.Pseudo))
	}
	add("props", stringArray(p.Props))
	if p.Box {
		add("box", ojson.Bool(true))
	}
	return o
}

// JSON is the token spec as its object literal writes it.
func (t TokenSpec) JSON() ojson.Object {
	return ojson.Object{{Key: "prefixes", Value: stringArray(t.Prefixes)}, {Key: "exclude", Value: stringArray(t.Exclude)}}
}

// JSON is the spec as its module's default export writes it.
func (s Spec) JSON() ojson.Object {
	o := ojson.Object{{Key: "path", Value: ojson.String(s.Path)}}
	if s.Wait != 0 {
		o = append(o, ojson.Member{Key: "wait", Value: ojson.Float(float64(s.Wait))})
	}
	if s.Click != "" {
		o = append(o, ojson.Member{Key: "click", Value: ojson.String(s.Click)})
	}
	if s.Hover != "" {
		o = append(o, ojson.Member{Key: "hover", Value: ojson.String(s.Hover)})
	}
	if s.Tokens != nil {
		return append(o, ojson.Member{Key: "tokens", Value: s.Tokens.JSON()})
	}
	probes := ojson.Array{}
	for _, p := range s.Probes {
		probes = append(probes, p.JSON())
	}
	return append(o, ojson.Member{Key: "probes", Value: probes})
}

// @invariant: the reader's argument and what its digest is of: the probes, or the
// token spec
func (s Spec) given() ojson.Value {
	if s.Tokens != nil {
		return s.Tokens.JSON()
	}
	return s.JSON().Get("probes")
}

func stringArray(list []string) ojson.Array {
	a := ojson.Array{}
	for _, s := range list {
		a = append(a, ojson.String(s))
	}
	return a
}

// Stringify is JSON.stringify(v) with no indent; with sorted, each object's keys
// sorted, as the reader's canon writes it (the text its digests are of).
func Stringify(v ojson.Value, sorted bool) string {
	var b bytes.Buffer
	stringify(&b, v, sorted)
	return b.String()
}

func stringify(b *bytes.Buffer, v ojson.Value, sorted bool) {
	switch v := v.(type) {
	case ojson.Object:
		members := slices.Clone(v)
		if sorted {
			slices.SortStableFunc(members, func(a, b ojson.Member) int { return cmpString(a.Key, b.Key) })
		}
		b.WriteByte('{')
		first := true
		for _, m := range members {
			if m.Value == nil {
				continue
			}
			if !first {
				b.WriteByte(',')
			}
			first = false
			b.Write(ojson.Marshal(ojson.String(m.Key)))
			b.WriteByte(':')
			stringify(b, m.Value, sorted)
		}
		b.WriteByte('}')
	case ojson.Array:
		b.WriteByte('[')
		for i, e := range v {
			if i > 0 {
				b.WriteByte(',')
			}
			stringify(b, e, sorted)
		}
		b.WriteByte(']')
	case ojson.Number:
		// @invariant: a number as JavaScript writes the one it parsed, whatever the
		// file's literal
		f, err := strconv.ParseFloat(string(v), 64)
		if err != nil {
			panic(fmt.Sprintf("signedin: %q is no number", v))
		}
		b.WriteString(string(ojson.Float(f)))
	default:
		b.Write(ojson.Marshal(v))
	}
}

// @external: Array.prototype.sort's default order: UTF-16 code units, which differs
// from Go's byte order only past U+FFFF
func cmpString(a, b string) int {
	ua, ub := utf16Units(a), utf16Units(b)
	return slices.Compare(ua, ub)
}

func utf16Units(s string) []uint16 {
	var u []uint16
	for _, r := range s {
		if r >= 0x10000 {
			r -= 0x10000
			u = append(u, uint16(0xD800+(r>>10)), uint16(0xDC00+(r&0x3FF)))
			continue
		}
		u = append(u, uint16(r))
	}
	return u
}

// Digest is the reader's sha(): the SHA-256 of s, its bytes as numbers.
func Digest(s string) ojson.Array {
	sum := sha256.Sum256([]byte(s))
	a := ojson.Array{}
	for _, b := range sum {
		a = append(a, ojson.Float(float64(b)))
	}
	return a
}

// Profile is the Chrome profile every signed-in site shares, outside the
// repository: it holds the owner's session.
var Profile = filepath.Join(home(), "Library", "Application Support", "ciderui-signed-in")

func home() string {
	h, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	return h
}

var (
	idsFile     = filepath.Join(Profile, "ids.json")
	cookiesFile = filepath.Join(Profile, "cookies.json")
)

// CLI is the command line over reference/<site>'s specs and its records (results/
// pages.go's Pages), so every signed-in site is read and checked by the one Reader.
// origin is the site's.
func CLI(site, origin string, specs map[string]Spec, records map[string]reference.SignedInPage) {
	status, err := cli(site, origin, specs, records, os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	os.Exit(status)
}

func cli(site, origin string, specs map[string]Spec, records map[string]reference.SignedInPage, args []string) (int, error) {
	var flag, name, file string
	switch {
	case len(args) > 0 && slices.Contains([]string{"--check", "--run", "--login", "--save"}, args[0]):
		flag = args[0]
		if len(args) > 1 {
			name = args[1]
		}
		if flag == "--save" {
			if len(args) != 3 {
				return 0, fmt.Errorf("usage: go run ./reference/%s --save <spec> <file>", site)
			}
			file = args[2]
		}
	case len(args) > 0:
		name = args[0]
	}
	if flag == "--login" {
		return 0, login(origin)
	}
	spec, ok := specs[name]
	if !ok {
		return 0, fmt.Errorf("no spec %q in reference/%s/pages", name, site)
	}
	switch flag {
	case "--run":
		raw, err := run(site, origin, name, spec)
		if err != nil {
			return 0, err
		}
		return save(site, name, spec, raw)
	case "--save":
		raw, err := os.ReadFile(file)
		if err != nil {
			return 0, err
		}
		return save(site, name, spec, raw)
	case "":
		fmt.Println(script(name, spec))
		return 0, nil
	}
	record, ok := records[name]
	if !ok {
		return 0, fmt.Errorf("no record %q in reference/%s/results", name, site)
	}
	return check(name, spec, record.JSON())
}

// @invariant: a record is written only once it passes check: what fails it (a probe
// that matched nothing, an id) never reaches the repository
func save(site, name string, spec Spec, raw []byte) (int, error) {
	parsed, err := ojson.Parse(raw)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", name, err)
	}
	file, ok := parsed.(ojson.Object)
	if !ok {
		return 0, fmt.Errorf("%s holds no object", name)
	}
	status, err := check(name, spec, file)
	if err != nil || status != 0 {
		return status, err
	}
	page, err := reference.SignedInFromJSON(name, file)
	if err != nil {
		return 0, err
	}
	return 0, reference.WriteSignedIn(filepath.Join(probe.Dir(site), "results"), site, page)
}

// @invariant: the script the page runs: its reader called with the spec's own
func script(name string, spec Spec) string {
	if spec.Tokens != nil {
		return fmt.Sprintf("await (%s)(%s, %s)", Tokens, Stringify(spec.given(), false), ojson.Marshal(ojson.String(name+".json")))
	}
	return fmt.Sprintf("await (%s)(%s)", Reader, Stringify(spec.given(), false))
}

// @invariant: nothing of the account in what was read: no id (six digits or more, not
// a fraction's; a UUID), no address
var (
	uuidOrAddress = regexp.MustCompile(`[0-9a-f]{8}-[0-9a-f]{4}-|@`)
	digits        = regexp.MustCompile(`[0-9]+`)
)

// @trap: /(?<![\d.])\d{6,}/ in JavaScript: Go's regexp has no lookbehind; a run of six
// digits or more not after a digit or a point is a maximal run not after a point
func holdsID(s string) bool {
	for _, loc := range digits.FindAllStringIndex(s, -1) {
		if loc[1]-loc[0] >= 6 && (loc[0] == 0 || s[loc[0]-1] != '.') {
			return true
		}
	}
	return false
}

func private(s string) bool { return holdsID(s) || uuidOrAddress.MatchString(s) }

// @invariant: a JavaScript value as a template literal writes it
func text(v ojson.Value) string {
	switch v := v.(type) {
	case ojson.String:
		return string(v)
	case nil:
		return "undefined"
	case ojson.Array:
		parts := make([]string, len(v))
		for i, e := range v {
			parts[i] = text(e)
		}
		return strings.Join(parts, ",")
	default:
		return Stringify(v, false)
	}
}

// @invariant: file is the record as the page handed it back, its result's members and
// sha256
func check(name string, spec Spec, file ojson.Object) (int, error) {
	var result ojson.Object
	for _, m := range file {
		if m.Key != "sha256" {
			result = append(result, m)
		}
	}
	same := func(a, b ojson.Value) bool { return Stringify(a, false) == Stringify(b, false) }
	kind, readerName, readerSource, read := "probes", "READER", Reader, result.Get("probes")
	if spec.Tokens != nil {
		kind, readerName, readerSource, read = "tokens", "TOKENS", Tokens, result.Get("tokens")
	}
	var fails []string
	if !same(Digest(Stringify(result, true)), file.Get("sha256")) {
		fails = append(fails, "the result is not what the page handed back (its digest differs)")
	}
	if !same(result.Get("probesDigest"), Digest(Stringify(spec.given(), true))) {
		fails = append(fails, fmt.Sprintf("the page was given other %s than pages/%s's", kind, name))
	}
	if !same(result.Get("readerDigest"), Digest(readerSource)) {
		fails = append(fails, "the page ran another reader than "+readerName)
	}
	if path := text(result.Get("path")); path != spec.Path {
		fails = append(fails, fmt.Sprintf("read on %s, not %s", path, spec.Path))
	}
	entries, _ := read.(ojson.Object)
	for _, m := range entries {
		if _, null := m.Value.(ojson.Null); null || m.Value == ojson.String("") {
			fails = append(fails, m.Key+" matched nothing")
		}
	}
	if spec.Tokens != nil {
		for _, p := range spec.Tokens.Prefixes {
			if !slices.ContainsFunc(entries, func(m ojson.Member) bool { return strings.HasPrefix(m.Key, p) }) {
				fails = append(fails, "no token begins "+p)
			}
		}
	}
	if private(text(result.Get("path")) + Stringify(read, false)) {
		fails = append(fails, "it holds what may be the account's: an id or an address")
	}
	for _, f := range fails {
		fmt.Printf("FAIL %s: %s\n", name, f)
	}
	if len(fails) > 0 {
		return 1, nil
	}
	window, _ := result.Get("window").(ojson.Array)
	sizes := make([]string, len(window))
	for i, w := range window {
		sizes[i] = text(w)
	}
	fmt.Printf("ok %s: %d %s, %s, %s, %s, %s\n", name, len(entries), kind, text(result.Get("path")), text(result.Get("browser")), strings.Join(sizes, " x "), text(result.Get("date")))
	return 0, nil
}

// @external: Google Chrome with Profile, a window on the Mac: Apple's sign-in asks a
// person (a code on their device), so the profile is signed in by hand once and
// read by Playwright after. Apple's sign-in cookies are the session's (no expiry),
// which Chrome drops when it quits: --login keeps every cookie in cookies.json, and
// each launch puts them back.
func chrome() (*playwright.Playwright, playwright.BrowserContext, error) {
	pw, err := playwright.Run(&playwright.RunOptions{SkipInstallBrowsers: true})
	if err != nil {
		return nil, nil, fmt.Errorf("playwright: %w", err)
	}
	context, err := pw.Chromium.LaunchPersistentContext(Profile, playwright.BrowserTypeLaunchPersistentContextOptions{
		Channel:           playwright.String("chrome"),
		Headless:          playwright.Bool(false),
		Viewport:          &playwright.Size{Width: 1920, Height: 907},
		DeviceScaleFactor: playwright.Float(2),
	})
	if err != nil {
		pw.Stop()
		return nil, nil, fmt.Errorf("launch chrome: %w", err)
	}
	raw, err := os.ReadFile(cookiesFile)
	if errors.Is(err, os.ErrNotExist) {
		return pw, context, nil
	}
	if err == nil {
		var cookies []playwright.OptionalCookie
		if err = json.Unmarshal(raw, &cookies); err == nil {
			err = context.AddCookies(cookies)
		}
	}
	if err != nil {
		context.Close()
		pw.Stop()
		return nil, nil, fmt.Errorf("%s: %w", cookiesFile, err)
	}
	return pw, context, nil
}

func firstPage(context playwright.BrowserContext) (playwright.Page, error) {
	if pages := context.Pages(); len(pages) > 0 {
		return pages[0], nil
	}
	return context.NewPage()
}

func login(origin string) error {
	pw, context, err := chrome()
	if err != nil {
		return err
	}
	defer pw.Stop()
	page, err := firstPage(context)
	if err != nil {
		return err
	}
	if _, err := page.Goto(origin); err != nil {
		return err
	}
	fmt.Printf("sign in to %s in the window, then quit Chrome\n", origin)
	closed := make(chan struct{})
	context.OnClose(func(playwright.BrowserContext) { close(closed) })
	// @external: the cookies are kept every 5s while the window is open: closing a
	// window on the Mac leaves Chrome running with a page no one sees (2026-11-13),
	// so no close is waited for
	keep := func() {
		cookies, err := context.Cookies()
		if err != nil {
			return
		}
		raw, err := json.Marshal(cookies)
		if err != nil {
			return
		}
		os.WriteFile(cookiesFile, raw, 0o600)
	}
	tick := time.NewTicker(5 * time.Second)
	defer tick.Stop()
	for {
		select {
		case <-tick.C:
			keep()
		case <-closed:
			return nil
		}
	}
}

func run(site, origin, name string, spec Spec) ([]byte, error) {
	ids := map[string]string{}
	if raw, err := os.ReadFile(idsFile); err == nil {
		var all map[string]map[string]string
		if err := json.Unmarshal(raw, &all); err != nil {
			return nil, fmt.Errorf("%s: %w", idsFile, err)
		}
		if all[site] != nil {
			ids = all[site]
		}
	} else if !errors.Is(err, os.ErrNotExist) {
		return nil, err
	}
	var missing error
	path := regexp.MustCompile(`<(app|team)>`).ReplaceAllStringFunc(spec.Path, func(m string) string {
		k := m[1 : len(m)-1]
		if ids[k] == "" && missing == nil {
			missing = fmt.Errorf("%s has no %s.%s for %s", idsFile, site, k, spec.Path)
		}
		return ids[k]
	})
	if missing != nil {
		return nil, missing
	}
	pw, context, err := chrome()
	if err != nil {
		return nil, err
	}
	defer pw.Stop()
	defer context.Close()
	page, err := firstPage(context)
	if err != nil {
		return nil, err
	}
	if _, err := page.Goto(origin+path, playwright.PageGotoOptions{WaitUntil: playwright.WaitUntilStateLoad, Timeout: playwright.Float(120000)}); err != nil {
		return nil, err
	}
	// @external: App Store Connect's page is empty at load and drawn by its script
	// about 2s on (body text 0 at 2s, 841 characters at 10s, 2026-11-13)
	if _, err := page.WaitForFunction("() => document.body.innerText.trim().length > 0", nil, playwright.PageWaitForFunctionOptions{Timeout: playwright.Float(60000)}); err != nil {
		return nil, err
	}
	wait := spec.Wait
	if wait == 0 {
		wait = 3000
	}
	page.WaitForTimeout(float64(wait))
	// @external: the sign-in redirect can come after the load; nothing is written then.
	// The message names the spec's path, never the ids
	at, err := url.Parse(page.URL())
	if err != nil {
		return nil, err
	}
	if at.EscapedPath() != path {
		return nil, fmt.Errorf("%s: not on %s (sent to %s%s): sign in again (go run ./reference/%s --login)", name, spec.Path, at.Host, regexp.MustCompile(`\d{6,}`).ReplaceAllString(at.EscapedPath(), "<app>"), site)
	}
	if spec.Click != "" {
		if err := page.Click(spec.Click); err != nil {
			return nil, err
		}
		page.WaitForTimeout(1500)
	}
	if spec.Hover != "" {
		if err := page.Hover(spec.Hover); err != nil {
			return nil, err
		}
		page.WaitForTimeout(800)
	}
	var out []byte
	if spec.Tokens != nil {
		download, err := page.ExpectDownload(func() error {
			_, err := page.Evaluate(fmt.Sprintf("(%s)(%s, %s)", Tokens, Stringify(spec.given(), false), ojson.Marshal(ojson.String(name+".json"))))
			return err
		})
		if err != nil {
			return nil, err
		}
		file, err := download.Path()
		if err != nil {
			return nil, err
		}
		if out, err = os.ReadFile(file); err != nil {
			return nil, err
		}
	} else {
		// @invariant: the page's own JSON.stringify of what the reader handed back, so
		// the record keeps its keys' order
		v, err := page.Evaluate(fmt.Sprintf("(%s)(%s).then((v) => JSON.stringify(v))", Reader, Stringify(spec.given(), false)))
		if err != nil {
			return nil, err
		}
		s, ok := v.(string)
		if !ok {
			return nil, fmt.Errorf("%s: the reader gave %T", name, v)
		}
		got, err := ojson.Parse([]byte(s))
		if err != nil {
			return nil, err
		}
		handed, _ := got.(ojson.Object)
		result, _ := handed.Get("result").(ojson.Object)
		out = append(ojson.Marshal(result.With(ojson.Member{Key: "sha256", Value: handed.Get("sha256")})), '\n')
	}
	return out, nil
}
