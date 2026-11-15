// Package browser opens the tests' pages in Playwright's Chromium: a stylesheet
// and markup on a page of their own, read through the DevTools protocol.
package browser

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"

	"github.com/at-least/ciderui/internal/build"
	"github.com/mxschmitt/playwright-go"
)

// Browser is one Chromium, its pages each in a context of their own.
type Browser struct {
	pw      *playwright.Playwright
	browser playwright.Browser
}

// Launch starts Chromium for t's test, stopped when it ends.
//
// @external: only in test/Dockerfile's image: its Chromium and its fallback fonts
// (Liberation Mono for code) make a value or a picture the same on every machine
func Launch(t testing.TB) *Browser {
	t.Helper()
	if os.Getenv("CIDERUI_CONTAINER") != "1" {
		t.Fatal("the browser runs only in test/Dockerfile's image: node scripts/docker.mjs <command>")
	}
	pw, err := playwright.Run(&playwright.RunOptions{SkipInstallBrowsers: true})
	if err != nil {
		t.Fatalf("browser: playwright: %v", err)
	}
	// @external: on Linux Chromium hints every glyph to a whole-pixel advance without
	// --font-render-hinting=none (reference/tracking/)
	b, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{Args: []string{"--font-render-hinting=none"}})
	if err != nil {
		t.Fatalf("browser: launch: %v", err)
	}
	t.Cleanup(func() {
		if err := b.Close(); err != nil {
			t.Errorf("browser: close: %v", err)
		}
		if err := pw.Stop(); err != nil {
			t.Errorf("browser: stop: %v", err)
		}
	})
	return &Browser{pw: pw, browser: b}
}

// Options are a page's window and settings: Width and Height default to 1200 by
// 800.
type Options struct {
	Width, Height int
	Touch         bool
	ReducedMotion bool
	ForcedColors  bool
	// Dark is the system's appearance dark; light otherwise.
	Dark bool
}

// Page is one open page.
type Page struct {
	Page    playwright.Page
	cdp     playwright.CDPSession
	context playwright.BrowserContext
	height  int
}

// @trap: The stylesheet's url(fonts/…) rewritten to an origin whose fonts are served
// from memory: not inlined as data: URLs, which put every font into every
// case's page. The page stays about:blank, set by setContent: navigated to that
// origin, the pointer's place (0, 0) hovered what the page drew there; a
// <base> for it took #fragment links off the page.
const fontsOrigin = "http://ciderui.test/fonts/"

var unrewrittenFont = regexp.MustCompile(`url\(["']?(\.\./)*fonts/`)

// Open opens a page of css and html.
func (b *Browser) Open(css, html string, o Options) (*Page, error) {
	if o.Width == 0 {
		o.Width = 1200
	}
	if o.Height == 0 {
		o.Height = 800
	}
	options := playwright.BrowserNewContextOptions{Viewport: &playwright.Size{Width: o.Width, Height: o.Height}}
	if o.ReducedMotion {
		options.ReducedMotion = playwright.ReducedMotionReduce
	}
	if o.ForcedColors {
		options.ForcedColors = playwright.ForcedColorsActive
	}
	if o.Dark {
		options.ColorScheme = playwright.ColorSchemeDark
	}
	context, err := b.browser.NewContext(options)
	if err != nil {
		return nil, fmt.Errorf("browser: context: %w", err)
	}
	p, err := open(context, css, html, o)
	if err != nil {
		context.Close()
		return nil, err
	}
	return p, nil
}

func open(context playwright.BrowserContext, css, html string, o Options) (*Page, error) {
	page, err := context.NewPage()
	if err != nil {
		return nil, fmt.Errorf("browser: page: %w", err)
	}
	cdp, err := context.NewCDPSession(page)
	if err != nil {
		return nil, fmt.Errorf("browser: cdp: %w", err)
	}
	// @external: the DevTools protocol's touch emulation is what makes (hover: none)
	// match: Chromium ignores an emulated hover media feature
	if o.Touch {
		if _, err := cdp.Send("Emulation.setTouchEmulationEnabled", map[string]any{"enabled": true, "maxTouchPoints": 5}); err != nil {
			return nil, fmt.Errorf("browser: touch: %w", err)
		}
	}
	// @external: about:blank's origin is opaque: a font from another needs CORS
	err = context.Route(fontsOrigin+"*", func(r playwright.Route) {
		name := strings.TrimPrefix(r.Request().URL(), fontsOrigin)
		body, err := os.ReadFile(filepath.Join(build.Root, "fonts", name))
		if err != nil {
			panic(fmt.Sprintf("browser: font %s: %v", name, err))
		}
		if err := r.Fulfill(playwright.RouteFulfillOptions{ContentType: playwright.String("font/woff2"), Headers: map[string]string{"access-control-allow-origin": "*"}, Body: body}); err != nil {
			panic(fmt.Sprintf("browser: font %s: %v", name, err))
		}
	})
	if err != nil {
		return nil, fmt.Errorf("browser: route: %w", err)
	}
	served := strings.ReplaceAll(css, "url(fonts/", "url("+fontsOrigin)
	// @trap: a form not rewritten would load from about:blank, which fails silently: its
	// text in a fallback face, every check read off that
	if unrewrittenFont.MatchString(served) {
		return nil, fmt.Errorf("browser: the stylesheet names its fonts in a form not rewritten here")
	}
	if err := page.SetContent(`<!doctype html><html lang="en"><head><meta charset="utf-8"><style>` + served + `</style></head><body>` + html + `</body></html>`); err != nil {
		return nil, fmt.Errorf("browser: content: %w", err)
	}
	p := &Page{Page: page, cdp: cdp, context: context, height: o.Height}
	if _, err := page.Evaluate(`() => document.fonts.ready`); err != nil {
		return nil, fmt.Errorf("browser: fonts: %w", err)
	}
	for query, want := range map[string]bool{"(hover: hover)": !o.Touch, "(forced-colors: active)": o.ForcedColors} {
		got, err := p.Matches(query)
		if err != nil {
			return nil, err
		}
		if got != want {
			return nil, fmt.Errorf("browser: %s is %v on a page of %+v", query, got, o)
		}
	}
	for _, domain := range []string{"DOM.enable", "CSS.enable"} {
		if _, err := cdp.Send(domain, nil); err != nil {
			return nil, fmt.Errorf("browser: %s: %w", domain, err)
		}
	}
	return p, nil
}

// Matches is whether the page matches the media query.
func (p *Page) Matches(query string) (bool, error) {
	v, err := p.Page.Evaluate(`(q) => matchMedia(q).matches`, query)
	if err != nil {
		return false, fmt.Errorf("browser: matchMedia %s: %w", query, err)
	}
	matches, ok := v.(bool)
	if !ok {
		return false, fmt.Errorf("browser: matchMedia %s gave %v", query, v)
	}
	return matches, nil
}

// Close closes the page's context.
func (p *Page) Close() error { return p.context.Close() }

// ReadOptions is what Page.Read reads beside the target's own values: State, the
// pseudo-classes forced on it, or on On where On names another element; Pseudo,
// a pseudo-element (::after) whose values are read in place of the element's.
type ReadOptions struct {
	State  []string
	On     string
	Pseudo string
}

const readValues = `([selector, props, pseudo]) => {
  const el = document.querySelector(selector);
  if (!el) throw new Error('no element for ' + selector);
  const cs = getComputedStyle(el, pseudo || null);
  const box = el.getBoundingClientRect();
  const parent = () => el.parentElement.getBoundingClientRect();
  const BOX = { 'box-height': () => box.height, 'box-width': () => box.width, 'box-top': () => box.top, 'box-left': () => box.left, 'box-right': () => box.right,
    'box-x': () => box.left - parent().left, 'box-y': () => box.top - parent().top };
  const hit = () => [[0.1, 0.1], [0.9, 0.1], [0.1, 0.9], [0.9, 0.9], [0.5, 0.5]].map(([x, y]) => {
    const e = document.elementFromPoint(box.left + box.width * x, box.top + box.height * y);
    return e ? [e.tagName.toLowerCase(), ...e.classList].join('.') : 'none';
  }).join(' ');
  return Object.fromEntries(props.map(p => [p, p in BOX ? BOX[p]() + 'px' : p === 'focused' ? String(document.activeElement === el) : p === 'hit' ? hit() : cs.getPropertyValue(p)]));
}`

// Read reads props of the element selector names: a computed property; the box,
// box-height, box-width, box-top, box-left, box-right, box-x, box-y (from its
// parent's); `focused`; or `hit`, what the pointer finds at five points of it.
func (p *Page) Read(selector string, props []string, r ReadOptions) (map[string]string, error) {
	var node any
	if len(r.State) > 0 {
		on := selector
		if r.On != "" {
			on = r.On
		}
		doc, err := p.cdp.Send("DOM.getDocument", map[string]any{"depth": 0})
		if err != nil {
			return nil, fmt.Errorf("browser: DOM.getDocument: %w", err)
		}
		root := doc.(map[string]any)["root"].(map[string]any)["nodeId"]
		found, err := p.cdp.Send("DOM.querySelector", map[string]any{"nodeId": root, "selector": on})
		if err != nil {
			return nil, fmt.Errorf("browser: DOM.querySelector %s: %w", on, err)
		}
		node = found.(map[string]any)["nodeId"]
		if node == float64(0) {
			return nil, fmt.Errorf("browser: no element for %s", on)
		}
		// @invariant: transitions off until styles are flushed after the forced state is
		// dropped: a state's last change must land while they are off (reference/apple.com/)
		still, err := p.Page.AddStyleTag(playwright.PageAddStyleTagOptions{Content: playwright.String("*, *::before, *::after { transition: none !important; animation: none !important; }")})
		if err != nil {
			return nil, fmt.Errorf("browser: still: %w", err)
		}
		if _, err := still.Evaluate(`e => e.setAttribute('data-cui-still', '')`); err != nil {
			return nil, fmt.Errorf("browser: still: %w", err)
		}
		if _, err := p.cdp.Send("CSS.forcePseudoState", map[string]any{"nodeId": node, "forcedPseudoClasses": r.State}); err != nil {
			return nil, fmt.Errorf("browser: force %v: %w", r.State, err)
		}
	}
	v, err := p.Page.Evaluate(readValues, []any{selector, props, r.Pseudo})
	if err != nil {
		return nil, fmt.Errorf("browser: read %s: %w", selector, err)
	}
	if len(r.State) > 0 {
		if _, err := p.cdp.Send("CSS.forcePseudoState", map[string]any{"nodeId": node, "forcedPseudoClasses": []string{}}); err != nil {
			return nil, fmt.Errorf("browser: unforce %v: %w", r.State, err)
		}
		if _, err := p.Page.Evaluate(`() => { void document.body.offsetHeight; document.querySelector('[data-cui-still]').remove(); }`); err != nil {
			return nil, fmt.Errorf("browser: still: %w", err)
		}
	}
	values := map[string]string{}
	for k, x := range v.(map[string]any) {
		s, ok := x.(string)
		if !ok {
			return nil, fmt.Errorf("browser: read %s %s gave %v", selector, k, x)
		}
		values[k] = s
	}
	return values, nil
}

// Settle waits for every motion the page runs to end.
func (p *Page) Settle() error {
	_, err := p.Page.Evaluate(`async () => {
  const frames = () => new Promise((r) => requestAnimationFrame(() => requestAnimationFrame(r)));
  await frames();
  const running = document.getAnimations().filter((a) => a.playState === 'running' && a.effect.getComputedTiming().endTime !== Infinity);
  await Promise.all(running.map((a) => a.finished.catch((e) => { if (e.name !== 'AbortError') throw e; })));
  await frames();
}`)
	if err != nil {
		return fmt.Errorf("browser: settle: %w", err)
	}
	return nil
}

// Resize sets the window's width, its motion then run to its end.
//
// @trap: a resize starts the transitions a width drives (the field label's `all` runs
// its max-width): a read or a shot before they end missed by one level at five pixels
// in one run of four
func (p *Page) Resize(width int) error {
	if err := p.Page.SetViewportSize(width, p.height); err != nil {
		return fmt.Errorf("browser: resize %d: %w", width, err)
	}
	return p.Settle()
}
