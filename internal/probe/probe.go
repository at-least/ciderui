// Package probe is the browsers the reference probes run in, behind what Open gives:
//
//	safari    Safari through safaridriver, raw WebDriver over HTTP (Safari's "Allow
//	          remote automation" on), in a Safari the run launches and quits when
//	          it closes: Safari 27 takes no session in one it did not launch
//	          ("Safari was not launched for automation"), so quit Safari before the
//	          run
//	ios       the same on an iPhone simulator (Xcode) of iOS 26.5 (IOS names
//	          another), whichever of its safaridriver finds: with no version it
//	          chose iOS 27.1's iPhone Duo too, whose Safari it could not reach
//	          ("session not created", 2026-10-22)
//	chrome    the installed Google Chrome through Playwright
//	chromium  Playwright's own Chromium, with --font-render-hinting=none: on Linux
//	          it hints every glyph to a whole-pixel advance otherwise, and no text
//	          width measured there means anything
//	firefox   Playwright's own Firefox (`go run github.com/mxschmitt/playwright-go/cmd/playwright install firefox`):
//	          Gecko with Playwright's automation patched in, not Mozilla's release
//	          build; with its window, on a virtual X display of the run's own
//	          (Xvfb), its pictures that display's pixels (ImageMagick's import):
//	          Playwright's picture of a Firefox page leaves backdrop-filter out,
//	          headless or not, which Firefox paints on the display (2026-10-22)
//
// Move is none on the simulator, which has no pointer. Size is Playwright's viewport,
// Safari's window, its frame included. Shot is at the screen's pixels, the window's
// width times its pixel ratio. In Safari, Behind("window") opens a second window in
// front of the page's, which is then not its key window, the commands still the
// page's; Behind("tab") a tab in the page's window, selected, which hides the page,
// the commands the tab's (WebDriver's New Window, its type a hint: it hands back the
// type Safari made; a tab it makes is not selected, a switch to it selects it).
package probe

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math"
	"net"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/at-least/ciderui/internal/build"
	"github.com/at-least/ciderui/internal/ojson"
	"github.com/mxschmitt/playwright-go"
)

// Page is a page open in one of the browsers.
type Page interface {
	// Label is the browser, its version and its platform.
	Label() string
	// Run is the value of a JavaScript expression that gives a promise, as the
	// page's JSON.stringify writes it; nil where it is undefined.
	Run(expr string) (ojson.Value, error)
	Click(selector string) error
	Dblclick(selector string) error
	// Press presses a key, or keys held together as Playwright writes them (Alt+Tab).
	Press(key string) error
	Type(text string) error
	Move(x, y, ms float64) error
	Go(url string) error
	Size(width, height int) error
	// Behind opens a window or a tab in front of the page ("window", "tab") and gives
	// what Safari made; Behind("") closes it.
	Behind(kind string) (string, error)
	// Shot is a picture of the page, base64 PNG.
	Shot() (string, error)
	Close() error
	// Playwright is Playwright's page, for a probe that resizes the viewport or forces
	// states; nil in Safari.
	Playwright() playwright.Page
}

// Script calls call() on the element selector names.
func Script(p Page, selector, call string) error {
	sel, err := json.Marshal(selector)
	if err != nil {
		return err
	}
	_, err = p.Run(fmt.Sprintf("Promise.resolve((document.querySelector(%s).%s(), true))", sel, call))
	return err
}

// Wait is an expression that waits ms milliseconds.
func Wait(ms int) string {
	return fmt.Sprintf("new Promise(r => setTimeout(() => r(true), %d))", ms)
}

// Open opens url in the browser named. options are Playwright's page options (a
// touch screen: HasTouch), for its browsers only.
func Open(name, url string, options *playwright.BrowserNewPageOptions) (Page, error) {
	if options != nil && (name == "safari" || name == "ios") {
		return nil, fmt.Errorf("no page options in %s", name)
	}
	switch name {
	case "chrome":
		return viaPlaywright("chromium", "chrome", url, options)
	case "chromium":
		return viaPlaywright("chromium", "", url, options)
	case "firefox":
		return viaPlaywright("firefox", "", url, options)
	case "safari":
		return viaSafariDriver(map[string]any{"browserName": "safari"}, url)
	case "ios":
		version := os.Getenv("IOS")
		if version == "" {
			version = "26.5"
		}
		return viaSafariDriver(map[string]any{"browserName": "safari", "platformName": "iOS", "safari:useSimulator": true, "safari:deviceType": "iPhone", "safari:platformVersion": version}, url)
	}
	return nil, fmt.Errorf("no browser %s: safari, ios, chrome, chromium or firefox", name)
}

// Serve serves the built site in dir on 127.0.0.1:port, a directory's index.html for
// the directory; it returns once it listens.
func Serve(dir string, port int) (*http.Server, error) {
	types := map[string]string{".html": "text/html; charset=utf-8", ".css": "text/css", ".js": "text/javascript", ".woff2": "font/woff2", ".png": "image/png", ".svg": "image/svg+xml", ".xml": "application/xml", ".txt": "text/plain", ".json": "application/json"}
	server := &http.Server{Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		f := filepath.Join(dir, filepath.FromSlash(r.URL.Path))
		if st, err := os.Stat(f); err == nil && st.IsDir() {
			f = filepath.Join(f, "index.html")
		}
		b, err := os.ReadFile(f)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		t, ok := types[filepath.Ext(f)]
		if !ok {
			t = "application/octet-stream"
		}
		w.Header().Set("content-type", t)
		w.Write(b)
	})}
	l, err := net.Listen("tcp", "127.0.0.1:"+strconv.Itoa(port))
	if err != nil {
		return nil, fmt.Errorf("probe: serve: %w", err)
	}
	go func() {
		if err := server.Serve(l); !errors.Is(err, http.ErrServerClosed) {
			fmt.Fprintln(os.Stderr, "probe: serve:", err)
		}
	}()
	return server, nil
}

// @invariant: the expression's value crosses as the page's own JSON.stringify text,
// so an object keeps its keys' order; undefined crosses as itself
const asJSON = `(async () => { const v = await (%s); return v === undefined ? undefined : JSON.stringify(v); })()`

func decoded(v any) (ojson.Value, error) {
	if v == nil {
		return nil, nil
	}
	s, ok := v.(string)
	if !ok {
		return nil, fmt.Errorf("probe: the page gave %T, no JSON text", v)
	}
	return ojson.Parse([]byte(s))
}

type playwrightPage struct {
	label   string
	pw      *playwright.Playwright
	browser playwright.Browser
	page    playwright.Page
	screen  *xvfb
}

func viaPlaywright(engine, channel, url string, options *playwright.BrowserNewPageOptions) (Page, error) {
	pw, err := playwright.Run(&playwright.RunOptions{SkipInstallBrowsers: true})
	if err != nil {
		return nil, fmt.Errorf("probe: playwright: %w", err)
	}
	var screen *xvfb
	launch := playwright.BrowserTypeLaunchOptions{}
	switch {
	case engine == "firefox":
		if screen, err = startXvfb(); err != nil {
			return nil, errors.Join(err, pw.Stop())
		}
		launch.Headless = playwright.Bool(false)
		launch.Env = screen.env
	case channel != "":
		launch.Channel = playwright.String(channel)
	default:
		launch.Args = []string{"--font-render-hinting=none"}
	}
	browserType := pw.Chromium
	if engine == "firefox" {
		browserType = pw.Firefox
	}
	browser, err := browserType.Launch(launch)
	if err != nil {
		screenErr := screen.close()
		return nil, errors.Join(fmt.Errorf("probe: launch %s: %w", engine, err), screenErr, pw.Stop())
	}
	p := &playwrightPage{pw: pw, browser: browser, screen: screen}
	opts := playwright.BrowserNewPageOptions{}
	if options != nil {
		opts = *options
	}
	if opts.Viewport == nil {
		opts.Viewport = &playwright.Size{Width: 1440, Height: 900}
	}
	if p.page, err = browser.NewPage(opts); err != nil {
		return nil, errors.Join(fmt.Errorf("probe: page: %w", err), p.Close())
	}
	if err := p.Go(url); err != nil {
		return nil, errors.Join(err, p.Close())
	}
	name := engine
	if channel != "" {
		name = channel
	}
	p.label = fmt.Sprintf("%s %s on %s", name, browser.Version(), platform())
	return p, nil
}

// @external: Node's process.platform, which the records name: darwin, linux
func platform() string { return runtime.GOOS }

func (p *playwrightPage) Label() string { return p.label }

func (p *playwrightPage) Run(expr string) (ojson.Value, error) {
	v, err := p.page.Evaluate(fmt.Sprintf(asJSON, expr))
	if err != nil {
		return nil, fmt.Errorf("probe: run: %w", err)
	}
	return decoded(v)
}

func (p *playwrightPage) Click(selector string) error    { return p.page.Click(selector) }
func (p *playwrightPage) Dblclick(selector string) error { return p.page.Dblclick(selector) }
func (p *playwrightPage) Press(key string) error         { return p.page.Keyboard().Press(key) }
func (p *playwrightPage) Type(text string) error         { return p.page.Keyboard().Type(text) }

func (p *playwrightPage) Move(x, y, ms float64) error {
	steps := int(math.Max(1, math.Floor(ms/16+0.5)))
	return p.page.Mouse().Move(x, y, playwright.MouseMoveOptions{Steps: playwright.Int(steps)})
}

func (p *playwrightPage) Go(url string) error {
	_, err := p.page.Goto(url, playwright.PageGotoOptions{WaitUntil: playwright.WaitUntilStateLoad, Timeout: playwright.Float(120000)})
	if err != nil {
		return fmt.Errorf("probe: go %s: %w", url, err)
	}
	return nil
}

func (p *playwrightPage) Size(width, height int) error { return p.page.SetViewportSize(width, height) }

func (p *playwrightPage) Behind(string) (string, error) {
	return "", fmt.Errorf("nothing in front of the page in %s: Safari's alone", p.label)
}

func (p *playwrightPage) Shot() (string, error) {
	if p.screen != nil {
		return p.screen.shot(p.page)
	}
	b, err := p.page.Screenshot()
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(b), nil
}

func (p *playwrightPage) Close() error {
	err := p.browser.Close()
	screenErr := p.screen.close()
	return errors.Join(err, screenErr, p.pw.Stop())
}

func (p *playwrightPage) Playwright() playwright.Page { return p.page }

// @external: a virtual X display of the run's own, on a number Xvfb finds free
// (-displayfd); shot crops the page out of the display's pixels, where Firefox says
// its page is (mozInnerScreenX, mozInnerScreenY)
type xvfb struct {
	server  *exec.Cmd
	display string
	env     map[string]string
}

func startXvfb() (*xvfb, error) {
	r, w, err := os.Pipe()
	if err != nil {
		return nil, err
	}
	server := exec.Command("Xvfb", "-displayfd", "3", "-screen", "0", "2560x1600x24", "-nolisten", "tcp")
	server.ExtraFiles = []*os.File{w}
	var log bytes.Buffer
	server.Stderr = &log
	if err := server.Start(); err != nil {
		return nil, fmt.Errorf("probe: Xvfb: %w", err)
	}
	w.Close()
	number := make(chan string, 1)
	go func() {
		b := make([]byte, 64)
		var out []byte
		for {
			n, err := r.Read(b)
			out = append(out, b[:n]...)
			if bytes.Contains(out, []byte("\n")) || err != nil {
				number <- strings.TrimSpace(string(out))
				return
			}
		}
	}()
	select {
	case n := <-number:
		if n == "" {
			return nil, errors.Join(fmt.Errorf("probe: Xvfb named no display: %s", log.String()), server.Process.Kill())
		}
		// @external: Firefox on X, not on the desktop's Wayland
		env := map[string]string{"DISPLAY": ":" + n, "GDK_BACKEND": "x11", "MOZ_ENABLE_WAYLAND": "0", "WAYLAND_DISPLAY": ""}
		for _, kv := range os.Environ() {
			k, v, _ := strings.Cut(kv, "=")
			if _, set := env[k]; !set {
				env[k] = v
			}
		}
		delete(env, "WAYLAND_DISPLAY")
		return &xvfb{server: server, display: ":" + n, env: env}, nil
	case <-time.After(10 * time.Second):
		return nil, errors.Join(fmt.Errorf("probe: Xvfb named no display in 10s: %s", log.String()), server.Process.Kill())
	}
}

func (x *xvfb) shot(page playwright.Page) (string, error) {
	v, err := page.Evaluate(`() => [mozInnerScreenX, mozInnerScreenY, innerWidth, innerHeight].map((v) => Math.round(v * devicePixelRatio))`)
	if err != nil {
		return "", err
	}
	var r [4]int
	b, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	if err := json.Unmarshal(b, &r); err != nil {
		return "", err
	}
	png, err := exec.Command("import", "-display", x.display, "-window", "root", "-crop", fmt.Sprintf("%dx%d+%d+%d", r[2], r[3], r[0], r[1]), "+repage", "png:-").Output()
	if err != nil {
		return "", fmt.Errorf("probe: import: %w", err)
	}
	return base64.StdEncoding.EncodeToString(png), nil
}

func (x *xvfb) close() error {
	if x == nil {
		return nil
	}
	return kill(x.server)
}

// @external: Wait after Kill reports the exit the kill caused (an *exec.ExitError), and
// Kill of a process already gone os.ErrProcessDone: neither is a failure, and Wait reaps
// the process either way
func kill(cmd *exec.Cmd) error {
	var errs []error
	if err := cmd.Process.Kill(); err != nil && !errors.Is(err, os.ErrProcessDone) {
		errs = append(errs, fmt.Errorf("kill %s: %w", cmd.Path, err))
	}
	var exit *exec.ExitError
	if err := cmd.Wait(); err != nil && !errors.As(err, &exit) {
		errs = append(errs, fmt.Errorf("wait %s: %w", cmd.Path, err))
	}
	return errors.Join(errs...)
}

// @external: WebDriver's code points for the keys a probe presses (ControlOrMeta is
// Meta: safaridriver is the Mac's); a character is its own
var keys = map[string]string{"Tab": "", "Enter": "", "Escape": "", "Alt": "", "Shift": "", "Meta": "", "ControlOrMeta": ""}

type safariPage struct {
	label     string
	base, id  string
	simulator bool
	launched  bool
	driver    *exec.Cmd
	home      string
	other     string
	// @external: where the pointer is: WebDriver's starts at the window's corner
	pointer [2]float64
}

func safariRunning() bool { return exec.Command("pgrep", "-x", "Safari").Run() == nil }

func (s *safariPage) wd(method, path string, body any) (json.RawMessage, error) {
	var in io.Reader
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		in = bytes.NewReader(b)
	}
	req, err := http.NewRequest(method, s.base+path, in)
	if err != nil {
		return nil, err
	}
	req.Header.Set("content-type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("%s %s: %w", method, path, err)
	}
	defer res.Body.Close()
	var out struct {
		Value json.RawMessage `json:"value"`
	}
	if err := json.NewDecoder(res.Body).Decode(&out); err != nil {
		return nil, fmt.Errorf("%s %s: %w", method, path, err)
	}
	if res.StatusCode/100 != 2 {
		return nil, fmt.Errorf("%s %s: %d %s", method, path, res.StatusCode, out.Value)
	}
	return out.Value, nil
}

func viaSafariDriver(capabilities map[string]any, url string) (Page, error) {
	const port = 4473
	// @external: On the Mac a Safari the run launches is quit when it closes. Left with
	// no window, macOS quits it under memory pressure and still counts it as running
	// (automatic termination): safaridriver then waits for it and launches no other
	// ("waiting for its RWIApplication"). A Safari running before the run is left as
	// it is, and takes no session unless safaridriver launched it: one opened by
	// LaunchServices rejects it ("Safari was not launched for automation"), and the
	// one macOS reopens after quitting it does not even connect (2026-10-19).
	simulator := capabilities["safari:useSimulator"] == true
	s := &safariPage{base: "http://localhost:" + strconv.Itoa(port), simulator: simulator, launched: !simulator && !safariRunning()}
	s.driver = exec.Command("/usr/bin/safaridriver", "-p", strconv.Itoa(port))
	s.driver.Stdout, s.driver.Stderr = os.Stdout, os.Stderr
	if err := s.driver.Start(); err != nil {
		return nil, fmt.Errorf("probe: safaridriver: %w", err)
	}
	if err := s.open(capabilities, url); err != nil {
		return nil, errors.Join(err, s.Close())
	}
	return s, nil
}

func (s *safariPage) open(capabilities map[string]any, url string) error {
	for i := 0; ; i++ {
		res, err := http.Get(s.base + "/status")
		if err == nil {
			res.Body.Close()
			break
		}
		if i > 100 {
			return errors.New("safaridriver did not start")
		}
		time.Sleep(100 * time.Millisecond)
	}
	raw, err := s.wd("POST", "/session", map[string]any{"capabilities": map[string]any{"alwaysMatch": capabilities}})
	if err != nil {
		return err
	}
	var session struct {
		SessionID    string          `json:"sessionId"`
		Capabilities json.RawMessage `json:"capabilities"`
	}
	if err := json.Unmarshal(raw, &session); err != nil {
		return err
	}
	s.id = session.SessionID
	fmt.Println("capabilities:", string(session.Capabilities))
	var c struct {
		BrowserVersion  string `json:"browserVersion"`
		PlatformName    string `json:"platformName"`
		PlatformVersion string `json:"safari:platformVersion"`
	}
	if err := json.Unmarshal(session.Capabilities, &c); err != nil {
		return err
	}
	s.label = strings.TrimSpace(fmt.Sprintf("safari %s on %s %s", c.BrowserVersion, c.PlatformName, c.PlatformVersion))
	if _, err := s.wd("POST", s.at("/timeouts"), map[string]any{"script": 120000, "pageLoad": 180000}); err != nil {
		return err
	}
	if err := s.Go(url); err != nil {
		return err
	}
	home, err := s.wd("GET", s.at("/window"), nil)
	if err != nil {
		return err
	}
	return json.Unmarshal(home, &s.home)
}

func (s *safariPage) at(path string) string { return "/session/" + s.id + path }

func (s *safariPage) Label() string { return s.label }

func (s *safariPage) Run(expr string) (ojson.Value, error) {
	script := "const done = arguments[arguments.length - 1];\n" + expr + ".then(r => done({ ok: true, undef: r === undefined, r: r === undefined ? '' : JSON.stringify(r) }), e => done({ ok: false, e: String((e && e.stack) || e) }));"
	raw, err := s.wd("POST", s.at("/execute/async"), map[string]any{"script": script, "args": []any{}})
	if err != nil {
		return nil, err
	}
	var out struct {
		OK    bool   `json:"ok"`
		Undef bool   `json:"undef"`
		R     string `json:"r"`
		E     string `json:"e"`
	}
	if err := json.Unmarshal(raw, &out); err != nil {
		return nil, err
	}
	if !out.OK {
		return nil, errors.New(out.E)
	}
	if out.Undef {
		return nil, nil
	}
	return ojson.Parse([]byte(out.R))
}

func (s *safariPage) Click(selector string) error {
	raw, err := s.wd("POST", s.at("/element"), map[string]any{"using": "css selector", "value": selector})
	if err != nil {
		return err
	}
	var el map[string]string
	if err := json.Unmarshal(raw, &el); err != nil {
		return err
	}
	for _, ref := range el {
		_, err = s.wd("POST", s.at("/element/"+ref+"/click"), map[string]any{})
		return err
	}
	return fmt.Errorf("no element for %s", selector)
}

func (s *safariPage) pointerActions(actions []map[string]any) error {
	_, err := s.wd("POST", s.at("/actions"), map[string]any{"actions": []any{map[string]any{"type": "pointer", "id": "mouse", "parameters": map[string]any{"pointerType": "mouse"}, "actions": actions}}})
	return err
}

// @external: two clicks at the element's middle, at once: the pointer jumps there
func (s *safariPage) Dblclick(selector string) error {
	raw, err := s.wd("POST", s.at("/execute/sync"), map[string]any{"script": "const r = document.querySelector(arguments[0]).getBoundingClientRect(); return [Math.round(r.left + r.width / 2), Math.round(r.top + r.height / 2)];", "args": []any{selector}})
	if err != nil {
		return err
	}
	var xy [2]float64
	if err := json.Unmarshal(raw, &xy); err != nil {
		return err
	}
	actions := []map[string]any{{"type": "pointerMove", "duration": 0, "origin": "viewport", "x": xy[0], "y": xy[1]}}
	for range 2 {
		actions = append(actions, map[string]any{"type": "pointerDown", "button": 0}, map[string]any{"type": "pointerUp", "button": 0})
	}
	if err := s.pointerActions(actions); err != nil {
		return err
	}
	s.pointer = xy
	return nil
}

func (s *safariPage) keyActions(actions []map[string]any) error {
	_, err := s.wd("POST", s.at("/actions"), map[string]any{"actions": []any{map[string]any{"type": "key", "id": "keyboard", "actions": actions}}})
	return err
}

func (s *safariPage) Press(key string) error {
	var values []string
	for _, k := range strings.Split(key, "+") {
		if len([]rune(k)) == 1 {
			values = append(values, k)
			continue
		}
		v, ok := keys[k]
		if !ok {
			return fmt.Errorf("no WebDriver key for %s", k)
		}
		values = append(values, v)
	}
	var actions []map[string]any
	for _, v := range values {
		actions = append(actions, map[string]any{"type": "keyDown", "value": v})
	}
	for _, v := range slices.Backward(values) {
		actions = append(actions, map[string]any{"type": "keyUp", "value": v})
	}
	return s.keyActions(actions)
}

func (s *safariPage) Type(text string) error {
	var actions []map[string]any
	for _, r := range text {
		actions = append(actions, map[string]any{"type": "keyDown", "value": string(r)}, map[string]any{"type": "keyUp", "value": string(r)})
	}
	return s.keyActions(actions)
}

// @external: a move of a duration is one jump at its end in Safari 27 (one pointermove,
// 2026-10-21): it is moved a step a frame, from where it was
func (s *safariPage) Move(x, y, ms float64) error {
	if s.simulator {
		return errors.New("the simulator has no pointer to move")
	}
	steps := int(math.Max(1, math.Floor(ms/16+0.5)))
	x0, y0 := s.pointer[0], s.pointer[1]
	var actions []map[string]any
	for i := range steps {
		f := float64(i+1) / float64(steps)
		actions = append(actions, map[string]any{
			"type": "pointerMove", "duration": math.Floor(ms/float64(steps) + 0.5), "origin": "viewport",
			"x": math.Floor(x0 + (x-x0)*f + 0.5), "y": math.Floor(y0 + (y-y0)*f + 0.5),
		})
	}
	if err := s.pointerActions(actions); err != nil {
		return err
	}
	s.pointer = [2]float64{x, y}
	return nil
}

func (s *safariPage) Go(url string) error {
	_, err := s.wd("POST", s.at("/url"), map[string]any{"url": url})
	return err
}

func (s *safariPage) Size(width, height int) error {
	_, err := s.wd("POST", s.at("/window/rect"), map[string]any{"width": width, "height": height})
	return err
}

func (s *safariPage) Behind(kind string) (string, error) {
	if kind == "" {
		if _, err := s.wd("POST", s.at("/window"), map[string]any{"handle": s.other}); err != nil {
			return "", err
		}
		if _, err := s.wd("DELETE", s.at("/window"), nil); err != nil {
			return "", err
		}
		if _, err := s.wd("POST", s.at("/window"), map[string]any{"handle": s.home}); err != nil {
			return "", err
		}
		s.other = ""
		return "", nil
	}
	if s.other != "" {
		return "", errors.New("a window or tab is in front already")
	}
	raw, err := s.wd("POST", s.at("/window/new"), map[string]any{"type": kind})
	if err != nil {
		return "", err
	}
	var made struct {
		Handle string `json:"handle"`
		Type   string `json:"type"`
	}
	if err := json.Unmarshal(raw, &made); err != nil {
		return "", err
	}
	s.other = made.Handle
	if made.Type == "tab" {
		if _, err := s.wd("POST", s.at("/window"), map[string]any{"handle": s.other}); err != nil {
			return "", err
		}
	}
	return made.Type, nil
}

func (s *safariPage) Shot() (string, error) {
	raw, err := s.wd("GET", s.at("/screenshot"), nil)
	if err != nil {
		return "", err
	}
	var png string
	return png, json.Unmarshal(raw, &png)
}

func (s *safariPage) Close() error {
	var errs []error
	if s.id != "" {
		if _, err := s.wd("DELETE", "/session/"+s.id, nil); err != nil {
			fmt.Fprintln(os.Stderr, "delete session:", err)
		}
	}
	if err := kill(s.driver); err != nil {
		errs = append(errs, err)
	}
	if s.launched && safariRunning() {
		if err := exec.Command("osascript", "-e", `quit app "Safari"`).Run(); err != nil {
			errs = append(errs, err)
		}
		for i := 0; safariRunning(); i++ {
			if i > 100 {
				return errors.Join(append(errs, errors.New("Safari did not quit"))...)
			}
			time.Sleep(100 * time.Millisecond)
		}
	}
	return errors.Join(errs...)
}

func (s *safariPage) Playwright() playwright.Page { return nil }

// Date is now as JavaScript's toISOString writes it, the records' date.
func Date() ojson.String {
	return ojson.String(time.Now().UTC().Format("2006-01-02T15:04:05.000Z"))
}

// Dir is reference/<name>, the probe's own directory.
func Dir(name string) string { return filepath.Join(build.Root, "reference", name) }

// Record writes reference/<name>/results/<file> as JSON.stringify(v, null, 1) wrote it.
func Record(name, file string, v ojson.Value) error {
	dir := filepath.Join(Dir(name), "results")
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(dir, file), ojson.Marshal(v), 0o644)
}

// Browser is the browser BROWSER names, which a probe requires.
func Browser() string {
	b := os.Getenv("BROWSER")
	if b == "" {
		fmt.Fprintln(os.Stderr, "BROWSER=safari|ios|chrome|chromium|firefox names the browser")
		os.Exit(2)
	}
	return b
}
