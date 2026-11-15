// Command icons: ciderui's icons drawn from apple.com's numbers, each held to
// apple.com's own, read live (never stored), by the share of ink each pair has in
// common (intersection over union, rasterized by svg.js). One probe a run, each in
// Playwright's Chromium, its result results/<probe>.json:
//
//	go run ./reference/icons form|lines|nav|pager|search
package main

import (
	"encoding/base64"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/at-least/ciderui/internal/build"
	"github.com/at-least/ciderui/internal/ojson"
	"github.com/at-least/ciderui/internal/probe"
	"github.com/mxschmitt/playwright-go"
)

// @invariant: a probe's own pass mark, set before the run: each pair shares at least
// this of its ink
const mark = 0.85

func main() {
	probes := map[string]func() (bool, error){"form": form, "lines": lines, "nav": nav, "pager": pager, "search": search}
	if len(os.Args) < 2 || probes[os.Args[1]] == nil {
		fmt.Fprintln(os.Stderr, "usage: go run ./reference/icons form|lines|nav|pager|search")
		os.Exit(2)
	}
	pass, err := probes[os.Args[1]]()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if !pass {
		os.Exit(1)
	}
}

func svgJS() string {
	src, err := os.ReadFile(filepath.Join(probe.Dir("icons"), "svg.js"))
	if err != nil {
		panic(err)
	}
	return string(src)
}

func quote(v ojson.Value) string { return string(ojson.Marshal(v)) }

func str(s string) string { return quote(ojson.String(s)) }

func num(v ojson.Value) float64 {
	n, ok := v.(ojson.Number)
	if !ok {
		panic(fmt.Sprintf("icons: %v is no number", v))
	}
	f, err := strconv.ParseFloat(string(n), 64)
	if err != nil {
		panic(err)
	}
	return f
}

func iou(v ojson.Value) float64 {
	o, _ := v.(ojson.Object)
	return num(o.Get("iou"))
}

func object(v ojson.Value, what string) (ojson.Object, error) {
	o, ok := v.(ojson.Object)
	if !ok {
		return nil, fmt.Errorf("%s gave %T", what, v)
	}
	return o, nil
}

func closeAfter(page probe.Page, err error) error {
	if cerr := page.Close(); err == nil {
		err = cerr
	}
	return err
}

func record(file, label string, ours ojson.Object, result ojson.Object) error {
	return probe.Record("icons", file, ojson.Object{
		{Key: "browser", Value: ojson.String(label)},
		{Key: "date", Value: probe.Date()},
		{Key: "ours", Value: ours},
	}.With(result...))
}

func printResult(label string, result ojson.Object) {
	fmt.Println(label)
	fmt.Println(string(ojson.Marshal(result)))
}

func readExample(path string) string {
	raw, err := os.ReadFile(filepath.Join(build.Root, path))
	if err != nil {
		panic(err)
	}
	return string(raw)
}

var svgWrapper = regexp.MustCompile(`^<svg[^>]*>|</svg>$`)

// @invariant: a drawing as the kit's example holds it: its paths without the svg
// around them, in currentColor
func inner(svg string) string {
	return strings.ReplaceAll(svgWrapper.ReplaceAllString(svg, ""), `stroke="#000"`, `stroke="currentColor"`)
}

func passLine(pass bool, failing string) {
	if pass {
		fmt.Println("PASS: each shares at least 0.85 of its ink with apple.com's")
	} else {
		fmt.Println(failing)
	}
}

// @external: form: the store's form icons, apple.com's read live (never stored) and ciderui's
// own drawn from their numbers: the share of ink each pair has in common. The error's
// mark, SF Pro Icons U+F312 before a field's message (the gift card's page,
// giftcard:error.icon), drawn at 400px against ciderui's ALERT drawn on the same
// canvas, and where its ink sits in the message's row (the page at four device pixels
// to the CSS pixel, an empty required field the focus has left: nothing typed); the
// dropdown's chevron, the SVG its background draws (iPhone 18 Pro's buy page, its
// trade-in, tradein:dropdown.chevron), rasterized at 20 pixels to the unit (svg.js)
// against ciderui's DROPDOWN, recorded alone: no part draws it (decision 36). Fails if
// the mark shares below 0.85 (intersection over union), a mark set before the run, or
// if the field's error example draws another. The result is results/form.json.
//
// @invariant: ciderui's, drawn from MEASURED.md's numbers: alert in thousandths of an
// em, its box the glyph's advance (1.171em) by its ink's height (0.995em, from
// 0.85em above the baseline); dropdown in apple.com's 21 by 21.
const (
	alert    = `<circle cx="585" cy="497.5" r="456.25" fill="none" stroke="#000" stroke-width="82.5"/><path d="M584 280V550" fill="none" stroke="#000" stroke-width="70" stroke-linecap="round"/><circle cx="585" cy="692.5" r="50" fill="#000"/>`
	dropdown = `<svg viewBox="0 0 21 21"><path d="M6.075 9.275 10.5 13.7 14.925 9.275" fill="none" stroke="#000" stroke-width="1.45" stroke-linecap="round" stroke-linejoin="round"/></svg>`
)

// @external: the ink left of the text: its box in CSS pixels from the row's corner
// (the function's source as Playwright's evaluate was handed it, a function, not
// a template's text)
const alertPlace = `async ([png, textLeft]) => {
      const img = new Image(); img.src = ` + "`" + `data:image/png;base64,${png}` + "`" + `; await img.decode();
      const c = document.createElement('canvas'); c.width = img.width; c.height = img.height;
      const g = c.getContext('2d'); g.drawImage(img, 0, 0);
      const d = g.getImageData(0, 0, img.width, img.height).data, k = img.width / textLeft.row;
      let l = img.width, r = -1, t = img.height, b = -1;
      for (let y = 0; y < img.height; y++) for (let x = 0; x < Math.floor(textLeft.at * k); x++) {
        const i = (y * img.width + x) * 4;
        if (d[i] - d[i + 1] > 60) { l = Math.min(l, x); r = Math.max(r, x); t = Math.min(t, y); b = Math.max(b, y); }
      }
      // @trap: no red ink left of the text is no mark found, not a box: a box read
      // so would be inside out (its left past its right), which the check
      // below would take as a place
      if (r < l || b < t) throw new Error('no red ink left of the message\'s text: the mark is not where it was read');
      return { scale: k, ink: { left: l / k, right: (r + 1) / k, top: t / k, bottom: (b + 1) / k }, textLeft: textLeft.at };
    }`

func form() (bool, error) {
	src := svgJS()
	result := ojson.Object{}
	page, err := probe.Open("chromium", "https://www.apple.com/shop/buy-giftcard/giftcard", nil)
	if err != nil {
		return false, err
	}
	// @external: the mark, drawn on a canvas at 400px, and ciderui's on the same one
	v, err := page.Run(`(async () => {
    const size = 400, ch = String.fromCodePoint(0xF312), font = size + 'px "SF Pro Icons"';
    if (!(await document.fonts.load(font, ch)).length) throw new Error('no SF Pro Icons face on the page');
    const W = size * 2, H = size * 2, X = size / 2, BASE = size * 1.4;
    const mask = async (draw) => {
      const c = document.createElement('canvas'); c.width = W; c.height = H;
      const ctx = c.getContext('2d'); await draw(ctx);
      const d = ctx.getImageData(0, 0, W, H).data, m = new Uint8Array(W * H);
      for (let i = 0; i < W * H; i++) m[i] = d[i * 4 + 3] >= 128 ? 1 : 0;
      return m;
    };
    const apple = await mask((ctx) => { ctx.font = font; ctx.fillStyle = '#000'; ctx.fillText(ch, X, BASE); });
    const ours = await mask(async (ctx) => {
      const img = new Image();
      img.src = 'data:image/svg+xml,' + encodeURIComponent('<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 1171 995" width="' + 1.171 * size + '" height="' + 0.995 * size + '">' + ` + str(alert) + ` + '</svg>');
      await img.decode();
      ctx.drawImage(img, X, BASE - 0.85 * size, 1.171 * size, 0.995 * size);
    });
    let l = W, r = -1, t = H, b = -1, both = 0, either = 0;
    for (let y = 0; y < H; y++) for (let x = 0; x < W; x++) if (apple[y * W + x]) { l = Math.min(l, x); r = Math.max(r, x); t = Math.min(t, y); b = Math.max(b, y); }
    for (let i = 0; i < W * H; i++) { both += apple[i] & ours[i]; either += apple[i] | ours[i]; }
    const em = v => +(v / size).toFixed(4);
    const ctx = document.createElement('canvas').getContext('2d'); ctx.font = font;
    return { char: 'U+F312', size, advance: em(ctx.measureText(ch).width), ink: { left: em(l - X), right: em(r + 1 - X), top: em(BASE - t), bottom: em(BASE - b - 1) }, iou: +(both / either).toFixed(4) };
  })()`)
	if err := closeAfter(page, err); err != nil {
		return false, err
	}
	result = result.With(ojson.Member{Key: "alert", Value: v})

	// @external: where the mark's ink sits in its message's row: the page at four device
	// pixels to the CSS pixel, an empty required field the focus left
	scaled, err := probe.Open("chromium", "https://www.apple.com/shop/buy-giftcard/giftcard", &playwright.BrowserNewPageOptions{
		Viewport: &playwright.Size{Width: 1200, Height: 900}, DeviceScaleFactor: playwright.Float(4),
	})
	if err != nil {
		return false, err
	}
	place, err := func() (ojson.Value, error) {
		p := scaled.Playwright()
		p.WaitForTimeout(2500)
		for _, s := range []string{"input[name=dimensionGiftCardType][value=electronic]", "input[name=dimensionGiftCardTheme]", ".rf-choose-amount-dimensions > div:nth-child(2) .form-selector-input"} {
			if err := p.Click(s); err != nil {
				return nil, err
			}
			p.WaitForTimeout(1200)
		}
		field := ".rf-giftcard-delivery-section-fields .form-textbox:nth-child(2)"
		if err := p.Click(field + " .form-textbox-input"); err != nil {
			return nil, err
		}
		if _, err := p.Evaluate("() => document.activeElement.blur()"); err != nil {
			return nil, err
		}
		p.WaitForTimeout(800)
		row := p.Locator(field + " .form-message-wrapper").First()
		box, err := row.BoundingBox()
		if err != nil {
			return nil, err
		}
		text, err := p.Locator(field + " .form-message").First().BoundingBox()
		if err != nil {
			return nil, err
		}
		png, err := row.Screenshot()
		if err != nil {
			return nil, err
		}
		args := ojson.Array{ojson.String(base64.StdEncoding.EncodeToString(png)), ojson.Object{
			{Key: "row", Value: ojson.Float(box.Width)}, {Key: "at", Value: ojson.Float(text.X - box.X)},
		}}
		v, err := scaled.Run("(" + alertPlace + ")(" + quote(args) + ")")
		if err != nil {
			return nil, err
		}
		o, err := object(v, "the mark's place")
		if err != nil {
			return nil, err
		}
		return o.With(ojson.Member{Key: "rowHeight", Value: ojson.Float(box.Height)}), nil
	}()
	if err := closeAfter(scaled, err); err != nil {
		return false, err
	}
	result = result.With(ojson.Member{Key: "alertPlace", Value: place})

	// @external: the dropdown's chevron, its background's SVG
	page, err = probe.Open("chromium", "https://www.apple.com/shop/buy-iphone/iphone-18-pro", nil)
	if err != nil {
		return false, err
	}
	v, err = func() (ojson.Value, error) {
		page.Playwright().WaitForTimeout(2500)
		apple, err := page.Run(`Promise.resolve((() => { const u = getComputedStyle(document.querySelector('.form-dropdown-chevron')).backgroundImage; return decodeURIComponent(u.slice(u.indexOf(',') + 1, -2)); })())`)
		if err != nil {
			return nil, err
		}
		if err := page.Playwright().SetContent("<!doctype html><body></body>"); err != nil {
			return nil, err
		}
		return page.Run("(async () => {\n" + src + "\nconst a = await raster(" + quote(apple) + ", 20), o = await raster(" + str(dropdown) + ", 20);\nreturn { appleBox: a.box, oursBox: o.box, iou: compare(a, o), rows: [10, 12, 13.5].map(y => [y, a.runs('row', y), o.runs('row', y)]) };\n})()")
	}()
	label := page.Label()
	if err := closeAfter(page, err); err != nil {
		return false, err
	}
	result = result.With(ojson.Member{Key: "dropdown", Value: v})

	ours := ojson.Object{{Key: "alert", Value: ojson.String(alert)}, {Key: "dropdown", Value: ojson.String(dropdown)}}
	if err := record("form.json", label, ours, result); err != nil {
		return false, err
	}
	printResult(label, result)
	// @invariant: the icons compared are the ones the kit ships, in the field's error
	// example, in currentColor
	examples := readExample("components/site/field/examples/error.html")
	ok := strings.Contains(examples, strings.ReplaceAll(alert, `"#000"`, `"currentColor"`))
	if ok {
		fmt.Println("ok   alert is the examples'")
	} else {
		fmt.Println("FAIL alert is the examples'")
	}
	pass := iou(result.Get("alert")) >= mark && ok
	passLine(pass, "FAIL")
	return pass, nil
}

// @external: lines: the chevron ciderui draws as apple.com's own line, read live (never
// stored): the bag's accordion's (its polyline, 17 by 8.85, the bag's compact
// items'), rasterized at 20 pixels to the unit (svg.js) against ciderui's: the share
// of ink in common. Fails below 0.85 (intersection over union), a mark set before the
// run, and if the drawing is not the accordion's example's. The result is
// results/lines.json.
func lines() (bool, error) {
	line := func(w string) string {
		return `fill="none" stroke="#000" stroke-width="` + w + `" stroke-linecap="round" stroke-linejoin="round"`
	}
	ours := ojson.Object{
		{Key: "disclosure", Value: ojson.String(`<svg viewBox="0 0 17 8.85"><path d="M15 1.13 8.5 7.72 2 1.13" ` + line("2.25") + `/></svg>`)},
	}
	examples := map[string]string{"disclosure": "components/site/accordion/examples/accordion.html"}
	// @external: apple.com's: the polyline with its stroke as its page computes it
	read := func(selector string) string {
		return `(() => {
  const svg = document.querySelector(` + str(selector) + `).cloneNode(true);
  for (const l of svg.querySelectorAll('polyline')) {
    l.querySelectorAll('animate').forEach(a => a.remove());
    l.setAttribute('stroke-width', getComputedStyle(document.querySelector(` + str(selector) + ` + ' polyline')).strokeWidth);
  }
  return svg.outerHTML;
})()`
	}
	src := svgJS()
	apple := ojson.Object{}
	for _, at := range []struct {
		key, url, selector string
		width              int
	}{
		{"disclosure", "https://www.apple.com/shop/bag", ".rc-accordion-item .accordion-icon-svg", 1200},
	} {
		page, err := probe.Open("chromium", at.url, nil)
		if err != nil {
			return false, err
		}
		v, err := func() (ojson.Value, error) {
			if err := page.Playwright().SetViewportSize(at.width, 900); err != nil {
				return nil, err
			}
			page.Playwright().WaitForTimeout(2500)
			return page.Run("Promise.resolve(" + read(at.selector) + ")")
		}()
		if err := closeAfter(page, err); err != nil {
			return false, err
		}
		apple = apple.With(ojson.Member{Key: at.key, Value: v})
	}
	page, err := probe.Open("chromium", "about:blank", nil)
	if err != nil {
		return false, err
	}
	v, err := page.Run("(async () => {\n" + src + "\nconst apple = " + quote(apple) + ", ours = " + quote(ours) + `;
const out = {};
for (const k of Object.keys(ours)) {
  const a = await raster(apple[k], 20), o = await raster(ours[k], 20);
  out[k] = { appleBox: a.box, oursBox: o.box, iou: compare(a, o) };
}
return out;
})()`)
	label := page.Label()
	if err := closeAfter(page, err); err != nil {
		return false, err
	}
	result, err := object(v, "the rasters")
	if err != nil {
		return false, err
	}
	if err := record("lines.json", label, ours, result); err != nil {
		return false, err
	}
	printResult(label, result)
	pass := true
	for _, m := range ours {
		pass = pass && iou(result.Get(m.Key)) >= mark
	}
	for _, m := range ours {
		ok := strings.Contains(readExample(examples[m.Key]), inner(string(m.Value.(ojson.String))))
		if ok {
			fmt.Printf("ok   %s is its component's example's\n", m.Key)
		} else {
			fmt.Printf("FAIL %s is its component's example's\n", m.Key)
		}
		pass = pass && ok
	}
	passLine(pass, "FAIL: below 0.85")
	return pass, nil
}

// @external: nav: the global nav's icons, apple.com's SVG read live off its home page (never
// stored), rasterized at 20 pixels to the unit on a blank page (svg.js), and
// ciderui's own drawn from their numbers: the share of ink each pair has in common.
// The menu's button closed and open (its two lines' points, open, are the end of the
// animation apple.com runs on them). Fails below 0.85 (intersection over union), a
// mark set before the run, or if the nav's example draws others. The result is
// results/nav.json.
func nav() (bool, error) {
	// @invariant: ciderui's, in apple.com's view boxes: the numbers are MEASURED.md's. The
	// menu's open lines are its closed ones turned about their middles, as
	// globalnav.css turns them (the same numbers)
	bar := `fill="none" stroke="#000" stroke-width="1.2" stroke-linecap="round"`
	ours := ojson.Object{
		{Key: "menu", Value: ojson.String(`<svg viewBox="0 0 18 18"><path d="M2 5H16" ` + bar + `/><path d="M2 12H16" ` + bar + `/></svg>`)},
		{Key: "menu-open", Value: ojson.String(`<svg viewBox="0 0 18 18"><path d="M2 5H16" ` + bar + ` transform="translate(9.25 9.25) rotate(45) scale(1.16168 1) translate(-9 -5)"/><path d="M2 12H16" ` + bar + ` transform="translate(9.25 9.25) rotate(-45) scale(1.16168 1) translate(-9 -12)"/></svg>`)},
	}
	src := svgJS()
	page, err := probe.Open("chromium", "https://www.apple.com/", nil)
	if err != nil {
		return false, err
	}
	v, err := func() (ojson.Value, error) {
		apple, err := page.Run(`(async () => {
    const svg = (sel) => document.querySelector('#globalnav ' + sel).outerHTML;
    // the menu's lines open: each polyline's points at the end of the animation
    // that starts from its closed points
    const menu = document.querySelector('#globalnav-menutrigger-button svg').cloneNode(true);
    for (const line of menu.querySelectorAll('polyline')) {
      const values = (a) => a.getAttribute('values').split(';').map(v => v.trim());
      const opening = [...line.querySelectorAll('animate')].find(a => values(a)[0] === line.getAttribute('points'));
      line.dataset.open = values(opening).at(-1);
      line.dataset.dur = opening.getAttribute('dur');
      line.querySelectorAll('animate').forEach(a => a.remove());
    }
    const closed = menu.outerHTML;
    for (const line of menu.querySelectorAll('polyline')) line.setAttribute('points', line.dataset.open);
    return {
      menu: closed, 'menu-open': menu.outerHTML,
      openPoints: [...menu.querySelectorAll('polyline')].map(l => l.id + ': ' + l.getAttribute('points') + ' in ' + l.dataset.dur),
    };
  })()`)
		if err != nil {
			return nil, err
		}
		a, err := object(apple, "apple.com's menu")
		if err != nil {
			return nil, err
		}
		points, _ := a.Get("openPoints").(ojson.Array)
		var open []string
		for _, p := range points {
			open = append(open, string(p.(ojson.String)))
		}
		fmt.Println("apple.com's menu lines, open:", strings.Join(open, "; "))
		if err := page.Playwright().SetContent("<!doctype html><body></body>"); err != nil {
			return nil, err
		}
		return page.Run("(async () => {\n" + src + "\nconst apple = " + quote(apple) + ", ours = " + quote(ours) + `;
const out = {};
for (const k of Object.keys(ours)) {
  const a = await raster(apple[k], 20), o = await raster(ours[k], 20);
  out[k] = { appleBox: a.box, oursBox: o.box, iou: compare(a, o) };
}
out.openPoints = apple.openPoints;
return out;
})()`)
	}()
	label := page.Label()
	if err := closeAfter(page, err); err != nil {
		return false, err
	}
	result, err := object(v, "the rasters")
	if err != nil {
		return false, err
	}
	if err := record("nav.json", label, ours, result); err != nil {
		return false, err
	}
	printResult(label, result)
	// @invariant: the icons compared are the ones the kit ships: each drawing (but the open
	// X, which globalnav.css turns) is in the nav's example, in currentColor
	example := readExample("components/site/globalnav/examples/globalnav.html")
	shipped := true
	for _, m := range ours {
		if m.Key == "menu-open" {
			continue
		}
		ok := strings.Contains(example, inner(string(m.Value.(ojson.String))))
		if ok {
			fmt.Printf("ok   %s is the example's\n", m.Key)
		} else {
			fmt.Printf("FAIL %s is the example's\n", m.Key)
		}
		shipped = shipped && ok
	}
	pass := true
	for _, m := range ours {
		pass = pass && iou(result.Get(m.Key)) >= mark
	}
	pass = pass && shipped
	passLine(pass, "FAIL: below 0.85")
	return pass, nil
}

// @invariant: each pair's iou at the mark or over it, else the first under it named on
// stderr
func underMark(result ojson.Object, whose string) bool {
	for _, m := range result {
		if v := iou(m.Value); v < mark {
			fmt.Fprintf(os.Stderr, "%s shares %s of its ink with %s, under %s\n", m.Key, ojson.Float(v), whose, ojson.Float(mark))
			return true
		}
	}
	return false
}

// @external: pager: Apple Developer Forums' pager's chevrons, read live (never stored): SF
// Symbols' chevron.left and chevron.right at medium weight, each a path in a 74.561
// by 70.459 view box drawn past it (overflow visible), 0.74561em wide at the button's
// 15px (../apple.com/'s forums-pager:previous.icon). ciderui's own, drawn from their
// numbers, rasterized beside them over the ink's whole box (0 −10 74.561 90) at 4
// pixels to the unit (svg.js): the share of ink each pair has in common. Fails if
// either shares below 0.85 (intersection over union), a mark set before the run. The
// result is results/pager.json.
func pager() (bool, error) {
	// @invariant: ciderui's, in SF Symbols' units: two strokes 11.5 thick, round, meeting at
	// the point 11.78 from the advance's side, 35.2 down, their ends 49.9 in and 37.5
	// above and below it; the right one the left mirrored in the 74.561 advance
	const (
		left  = `<svg viewBox="0 -10 74.561 90"><path d="M49.9 -2.3 11.78 35.2 49.9 72.7" fill="none" stroke="#000" stroke-width="11.5" stroke-linecap="round" stroke-linejoin="round"/></svg>`
		right = `<svg viewBox="0 -10 74.561 90"><path d="M24.661 -2.3 62.781 35.2 24.661 72.7" fill="none" stroke="#000" stroke-width="11.5" stroke-linecap="round" stroke-linejoin="round"/></svg>`
	)
	src := svgJS()
	page, err := probe.Open("chromium", "https://developer.apple.com/forums/latest?page=3", nil)
	if err != nil {
		return false, err
	}
	v, err := func() (ojson.Value, error) {
		page.Playwright().WaitForTimeout(3000)
		apple, err := page.Run(`Promise.resolve(['previous-page', 'next-page'].map((c) => {
    const d = document.querySelector('nav.pagination a.' + c + ' sf-symbol svg path').getAttribute('d');
    return '<svg viewBox="0 -10 74.561 90"><path d="' + d + '"/></svg>';
  }))`)
		if err != nil {
			return nil, err
		}
		pair, _ := apple.(ojson.Array)
		if len(pair) != 2 {
			return nil, fmt.Errorf("the forums' chevrons gave %v", apple)
		}
		// @external: about:blank, not setContent: the forums' Content-Security-Policy stays with
		// the page's origin and refuses the blob: images svg.js draws
		if _, err := page.Playwright().Goto("about:blank"); err != nil {
			return nil, err
		}
		return page.Run("(async () => {\n" + src + "\nconst pair = async (a, o) => { const A = await raster(a, 4), O = await raster(o, 4); return { appleBox: A.box, oursBox: O.box, iou: compare(A, O) }; };\nreturn { left: await pair(" + quote(pair[0]) + ", " + str(left) + "), right: await pair(" + quote(pair[1]) + ", " + str(right) + ") };\n})()")
	}()
	label := page.Label()
	if err := closeAfter(page, err); err != nil {
		return false, err
	}
	result, err := object(v, "the rasters")
	if err != nil {
		return false, err
	}
	ours := ojson.Object{{Key: "left", Value: ojson.String(left)}, {Key: "right", Value: ojson.String(right)}}
	if err := record("pager.json", label, ours, result); err != nil {
		return false, err
	}
	printResult(label, result)
	return !underMark(result, "Apple's"), nil
}

// @external: search: apple.com's search page's two field icons, read live (never stored), each
// an SVG its button's background draws (search-serp:icon.left.button, :icon.right):
// the magnifier before the text and the reset after it, both 25 by 25. ciderui's
// own, drawn from their numbers, rasterized beside them at 20 pixels to the unit
// (svg.js): the share of ink each pair has in common, and the magnifier's ring read
// along its middle row. Fails if either shares below 0.85 (intersection over union),
// a mark set before the run. The result is results/search.json.
func search() (bool, error) {
	// @invariant: ciderui's, in apple.com's 25 by 25: the magnifier a ring about (10.69, 10.87)
	// of radius 6.28, 1.1 thick, its handle from the ring to (20.15, 20.05), 1.7 thick,
	// round; the reset a disc of radius 9.5 about the middle, an X cut from it to 3.5 each
	// way of the middle, 1.1 thick, round
	const (
		magnifier = `<svg viewBox="0 0 25 25"><circle cx="10.69" cy="10.87" r="6.28" fill="none" stroke="#000" stroke-width="1.1"/><path d="M15.25 15.35 20.15 20.05" fill="none" stroke="#000" stroke-width="1.7" stroke-linecap="round"/></svg>`
		reset     = `<svg viewBox="0 0 25 25"><path d="M12.5 3A9.5 9.5 0 1 0 12.5 22 9.5 9.5 0 1 0 12.5 3ZM13.2778 12.5 16.3889 15.6111A.55.55 0 0 1 15.6111 16.3889L12.5 13.2778 9.3889 16.3889A.55.55 0 0 1 8.6111 15.6111L11.7222 12.5 8.6111 9.3889A.55.55 0 0 1 9.3889 8.6111L12.5 11.7222 15.6111 8.6111A.55.55 0 0 1 16.3889 9.3889Z" fill="#000" fill-rule="evenodd"/></svg>`
	)
	src := svgJS()
	page, err := probe.Open("chromium", "https://www.apple.com/us/search/airpods?src=globalnav", nil)
	if err != nil {
		return false, err
	}
	v, err := func() (ojson.Value, error) {
		page.Playwright().WaitForTimeout(2500)
		apple, err := page.Run(`Promise.resolve(['.form-icons-search15', '.form-icons-reset'].map((s) => {
    const u = getComputedStyle(document.querySelector('.rf-serp-searchbox ' + s)).backgroundImage;
    return decodeURIComponent(u.slice(u.indexOf(',') + 1, -2));
  }))`)
		if err != nil {
			return nil, err
		}
		pair, _ := apple.(ojson.Array)
		if len(pair) != 2 {
			return nil, fmt.Errorf("the search field's icons gave %v", apple)
		}
		if err := page.Playwright().SetContent("<!doctype html><body></body>"); err != nil {
			return nil, err
		}
		return page.Run("(async () => {\n" + src + "\nconst pair = async (a, o) => { const A = await raster(a, 20), O = await raster(o, 20); return { appleBox: A.box, oursBox: O.box, iou: compare(A, O), rows: [10.87].map(y => [y, A.runs('row', y), O.runs('row', y)]) }; };\nreturn { magnifier: await pair(" + quote(pair[0]) + ", " + str(magnifier) + "), reset: await pair(" + quote(pair[1]) + ", " + str(reset) + ") };\n})()")
	}()
	label := page.Label()
	if err := closeAfter(page, err); err != nil {
		return false, err
	}
	result, err := object(v, "the rasters")
	if err != nil {
		return false, err
	}
	ours := ojson.Object{{Key: "magnifier", Value: ojson.String(magnifier)}, {Key: "reset", Value: ojson.String(reset)}}
	if err := record("search.json", label, ours, result); err != nil {
		return false, err
	}
	printResult(label, result)
	return !underMark(result, "apple.com's"), nil
}
