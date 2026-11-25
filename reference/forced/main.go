// Command forced: go run ./reference/forced: MEASURED.md holds its claims and method.
// It writes page.html first (page.go), then runs in Playwright's Chromium and
// Firefox, whatever BROWSER names.
package main

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/at-least/ciderui/internal/ojson"
	"github.com/at-least/ciderui/internal/probe"
	"github.com/mxschmitt/playwright-go"
)

func main() {
	if err := writePage(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	pass, err := run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if !pass {
		os.Exit(1)
	}
}

type mode struct {
	name    string
	firefox bool
	forced  bool
	launch  playwright.BrowserTypeLaunchOptions
	page    playwright.BrowserNewPageOptions
}

var chromiumArgs = []string{"--font-render-hinting=none"}

var modes = []mode{
	{name: "chromium", firefox: false, forced: false, launch: playwright.BrowserTypeLaunchOptions{Args: chromiumArgs}, page: playwright.BrowserNewPageOptions{ForcedColors: playwright.ForcedColorsNone}},
	{name: "chromium-forced", firefox: false, forced: true, launch: playwright.BrowserTypeLaunchOptions{Args: chromiumArgs}, page: playwright.BrowserNewPageOptions{ForcedColors: playwright.ForcedColorsActive}},
	{name: "chromium-forced-dark", firefox: false, forced: true, launch: playwright.BrowserTypeLaunchOptions{Args: chromiumArgs}, page: playwright.BrowserNewPageOptions{ForcedColors: playwright.ForcedColorsActive, ColorScheme: playwright.ColorSchemeDark}},
	{name: "firefox", firefox: true, forced: false, launch: playwright.BrowserTypeLaunchOptions{}, page: playwright.BrowserNewPageOptions{}},
	// @trap: the prefs a map[string]any: playwright-go's argument serializer panics on
	// another map
	{name: "firefox-forced", firefox: true, forced: true, launch: playwright.BrowserTypeLaunchOptions{FirefoxUserPrefs: map[string]any{"browser.display.document_color_use": 2}}, page: playwright.BrowserNewPageOptions{}},
}

// @external: a PNG's pixels, decoded in an unforced Chromium page with no color
// conversion
const decodeFn = `async (b64) => {
  const bytes = Uint8Array.from(atob(b64), (c) => c.charCodeAt(0));
  const bmp = await createImageBitmap(new Blob([bytes], { type: 'image/png' }), { colorSpaceConversion: 'none', premultiplyAlpha: 'none' });
  const c = new OffscreenCanvas(bmp.width, bmp.height);
  const g = c.getContext('2d');
  g.drawImage(bmp, 0, 0);
  return { w: bmp.width, h: bmp.height, d: Array.from(g.getImageData(0, 0, bmp.width, bmp.height).data) };
}`

type picture struct {
	W int   `json:"w"`
	H int   `json:"h"`
	D []int `json:"d"`
}

func (im picture) px(x, y int) string {
	i := (y*im.W + x) * 4
	return fmt.Sprintf("rgb(%d,%d,%d)", im.D[i], im.D[i+1], im.D[i+2])
}

func run() (pass bool, err error) {
	html, err := os.ReadFile(filepath.Join(probe.Dir("forced"), "page.html"))
	if err != nil {
		return false, err
	}
	pw, err := playwright.Run(&playwright.RunOptions{SkipInstallBrowsers: true})
	if err != nil {
		return false, fmt.Errorf("playwright: %w", err)
	}
	defer func() { err = errors.Join(err, pw.Stop()) }()
	reader, err := pw.Chromium.Launch()
	if err != nil {
		return false, err
	}
	defer reader.Close()
	decoder, err := reader.NewPage()
	if err != nil {
		return false, err
	}
	decode := func(png []byte) (picture, error) {
		v, err := decoder.Evaluate(decodeFn, base64.StdEncoding.EncodeToString(png))
		if err != nil {
			return picture{}, err
		}
		b, err := json.Marshal(v)
		if err != nil {
			return picture{}, err
		}
		var im picture
		return im, json.Unmarshal(b, &im)
	}

	pass = true
	for _, m := range modes {
		engine := pw.Chromium
		if m.firefox {
			engine = pw.Firefox
		}
		browser, err := engine.Launch(m.launch)
		if err != nil {
			return false, err
		}
		got, err := read(browser, m, string(html), decode)
		if cerr := browser.Close(); err == nil {
			err = cerr
		}
		if err != nil {
			return false, err
		}
		if err := probe.Record("forced", m.name+".json", ojson.Object{{Key: "date", Value: probe.Date()}}.With(got...)); err != nil {
			return false, err
		}
		not := "not "
		if m.forced {
			not = ""
		}
		options := got.Get("options").(ojson.Array)
		claims := []struct {
			text string
			ok   bool
		}{
			{fmt.Sprintf("forced colors %smatched (the control)", not), got.Get("forced") == ojson.Bool(m.forced)},
			{"the checked option drawn unlike the unchecked one (the control)", options[0] != options[1]},
		}
		fmt.Printf("%s: %s\n", m.name, got.Get("browser").(ojson.String))
		for _, c := range claims {
			pass = pass && c.ok
			mark := "ok  "
			if !c.ok {
				mark = "FAIL"
			}
			fmt.Printf("  %s %s\n", mark, c.text)
		}
	}
	return pass, nil
}

func read(browser playwright.Browser, m mode, html string, decode func([]byte) (picture, error)) (ojson.Object, error) {
	opts := m.page
	opts.Viewport = &playwright.Size{Width: 1200, Height: 900}
	page, err := browser.NewPage(opts)
	if err != nil {
		return nil, err
	}
	if err := page.SetContent(html); err != nil {
		return nil, err
	}
	if _, err := page.Evaluate(`() => document.fonts.ready`); err != nil {
		return nil, err
	}
	shot := func(selector string, i int) (picture, error) {
		png, err := page.Locator(selector).Nth(i).Screenshot()
		if err != nil {
			return picture{}, err
		}
		return decode(png)
	}
	v, err := page.Evaluate(`() => ({
    forced: matchMedia('(forced-colors: active)').matches,
    browser: navigator.userAgent.match(/(Chrome|Firefox)\/[\d.]+/)[0],
  })`)
	if err != nil {
		return nil, err
	}
	media, ok := v.(map[string]any)
	if !ok {
		return nil, fmt.Errorf("the page gave %T", v)
	}
	forced, ok := media["forced"].(bool)
	if !ok {
		return nil, errors.New("no forced-colors match")
	}
	engine, ok := media["browser"].(string)
	if !ok {
		return nil, errors.New("no browser in the user agent")
	}
	// @external: an option's edge: the pixel 1px in from its left side at its middle
	o0, err := shot("#opt .cui-option-label", 0)
	if err != nil {
		return nil, err
	}
	o1, err := shot("#opt .cui-option-label", 1)
	if err != nil {
		return nil, err
	}
	options := ojson.Array{ojson.String(o0.px(1, o0.H/2)), ojson.String(o1.px(1, o1.H/2))}
	return ojson.Object{{Key: "forced", Value: ojson.Bool(forced)}, {Key: "browser", Value: ojson.String(engine)}, {Key: "options", Value: options}}, nil
}
