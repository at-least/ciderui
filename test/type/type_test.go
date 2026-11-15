//go:build browser

package typetest

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"testing"

	"github.com/at-least/ciderui/internal/browser"
	"github.com/at-least/ciderui/internal/build"
	"github.com/at-least/ciderui/internal/check"
	"github.com/at-least/ciderui/internal/reference"
	apple "github.com/at-least/ciderui/reference/apple.com/results"
	"github.com/at-least/ciderui/reference/tracking/table"
)

const stack = `-apple-system, BlinkMacSystemFont, Inter, "Helvetica Neue", Helvetica, Arial, "PingFang TC", "Noto Sans TC", "Noto Sans CJK TC", "Microsoft JhengHei", sans-serif`

var utility = regexp.MustCompile(`@utility cui-([a-z-]+)`)

func TestType(t *testing.T) {
	typeCSS, err := os.ReadFile(filepath.Join(build.Root, "type.css"))
	if err != nil {
		t.Fatal(err)
	}
	var names []string
	for _, m := range utility.FindAllStringSubmatch(string(typeCSS), -1) {
		names = append(names, m[1])
	}
	tracking := map[string]float64{}
	for _, r := range table.CUITracking {
		tracking[strconv.Itoa(r.Size)] = r.Em
	}
	var classes, html strings.Builder
	for _, n := range names {
		classes.WriteString("cui-" + n + " ")
		fmt.Fprintf(&html, `<p id="%s" class="cui-%s">Sample text</p>`, n, n)
	}
	classes.WriteString("lg:cui-headline-elevated text-[20px]")
	html.WriteString(`<p id="variant" class="cui-headline lg:cui-headline-elevated">v</p><p id="plain" class="cui-headline text-[20px]">p</p>`)
	css, err := build.Compile(build.Sheet{CSS: `@import "index.css";`, Classes: classes.String(), Sources: nil})
	if err != nil {
		t.Fatal(err)
	}
	b, err := browser.Launch(t).Open(css, html.String(), browser.Options{})
	if err != nil {
		t.Fatal(err)
	}
	defer b.Close()

	// @external: the store defines no t-body: cui-body is the page's body text
	reading := func(page reference.Page, name string, width int, prop string) string {
		probe := "type.t-" + name
		if name == "body" {
			probe = "page.body"
		}
		v, err := page.Value(probe, width, prop, "")
		if err != nil {
			t.Fatal(err)
		}
		return v
	}
	storeValue := func(name string, width int, prop string) string { return reading(apple.Store, name, width, prop) }
	read := func(selector string, props ...string) map[string]string {
		got, err := b.Read(selector, props, browser.ReadOptions{})
		if err != nil {
			t.Fatal(err)
		}
		return got
	}
	resize := func(width int) {
		if err := b.Resize(width); err != nil {
			t.Fatal(err)
		}
	}

	for _, width := range reference.Widths {
		resize(width)
		for _, name := range names {
			got := read("#"+name, "font-size", "line-height", "font-weight", "letter-spacing", "font-family")
			for _, prop := range []string{"font-size", "line-height", "font-weight"} {
				check.Equal(t, fmt.Sprintf("cui-%s %s at %dpx", name, prop, width), got[prop], storeValue(name, width, prop))
			}
			// @trap: for apple.com's size, not the read one: a wrong size is its own failure
			size, err := strconv.ParseFloat(strings.TrimSuffix(storeValue(name, width, "font-size"), "px"), 64)
			if err != nil {
				t.Fatal(err)
			}
			fitted, ok := tracking[check.Number(size)]
			if !ok {
				t.Fatalf("reference/tracking/table has no %vpx", size)
			}
			check.Within(t, fmt.Sprintf("cui-%s letter-spacing at %dpx", name, width), got["letter-spacing"], check.Number(size*fitted)+"px", 0.001)
			check.Equal(t, fmt.Sprintf("cui-%s font-family at %dpx", name, width), got["font-family"], stack)
		}
	}
	for _, w := range [][2]int{{734, 390}, {735, 900}} {
		resize(w[0])
		for _, name := range names {
			check.Equal(t, fmt.Sprintf("cui-%s font-size at %dpx is its %dpx size", name, w[0], w[1]), read("#"+name, "font-size")["font-size"], storeValue(name, w[1], "font-size"))
		}
	}
	for _, v := range []struct {
		width   int
		variant string
	}{{390, "32px"}, {900, "48px"}, {1068, "48px"}, {1069, "64px"}, {1600, "64px"}} {
		resize(v.width)
		check.Equal(t, fmt.Sprintf("lg:cui-headline-elevated over cui-headline at %dpx", v.width), read("#variant", "font-size")["font-size"], v.variant)
		check.Equal(t, fmt.Sprintf("text-[20px] over cui-headline at %dpx", v.width), read("#plain", "font-size")["font-size"], "20px")
	}
	// @external: Traditional Chinese: apple.com/tw's store, line height and letter-spacing
	// as it reads them
	if _, err := b.Page.Evaluate(`() => { document.documentElement.lang = 'zh-Hant'; }`); err != nil {
		t.Fatal(err)
	}
	zh := []string{"font-size", "line-height", "letter-spacing"}
	for _, width := range reference.Widths {
		resize(width)
		for _, name := range names {
			got := read("#"+name, zh...)
			for _, prop := range zh {
				check.Equal(t, fmt.Sprintf("zh-Hant cui-%s %s at %dpx is apple.com/tw's", name, prop, width), got[prop], reading(apple.StoreTw, name, width, prop))
			}
		}
		body := read("body", "line-height", "letter-spacing")
		for _, prop := range []string{"line-height", "letter-spacing"} {
			check.Equal(t, fmt.Sprintf("zh-Hant body %s at %dpx is apple.com/tw's", prop, width), body[prop], reading(apple.StoreTw, "body", width, prop))
		}
	}
}
