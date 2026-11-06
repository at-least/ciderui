package ciderui

import (
	"fmt"
	"maps"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"testing"
)

func TestEveryExamplePasses(t *testing.T) {
	paths, err := filepath.Glob("components/*/*/examples/*.html")
	if err != nil {
		t.Fatal(err)
	}
	if len(paths) <= 100 {
		t.Fatalf("%d examples read", len(paths))
	}
	for _, path := range paths {
		html, err := os.ReadFile(path)
		if err != nil {
			t.Fatal(err)
		}
		if problems := CheckMarkup(string(html)); problems != nil {
			t.Errorf("%s: %q", path, problems)
		}
	}
}

func TestCheckMarkup(t *testing.T) {
	noClass := func(tag string) []string {
		return []string{fmt.Sprintf(`class "cui-console-badg" on <%s class="cui-console-badg">: ciderui has no such class`, tag)}
	}
	for _, c := range []struct {
		name, html string
		want       []string
	}{
		{"a misspelt class", `<p>a</p><span class="cui-console-badg mt-1">b</span>`, []string{`class "cui-console-badg" on <span class="cui-console-badg mt-1">: ciderui has no such class`}},
		{"a keyframes name is no class", `<li class="cui-globalnav-row-in">a</li>`, []string{`class "cui-globalnav-row-in" on <li class="cui-globalnav-row-in">: ciderui has no such class`}},
		{"a value no rule selects", `<span class="cui-console-badge" data-color="purple">b</span>`, []string{`data-color="purple" on <span class="cui-console-badge">: no rule of ciderui's selects it (cui-console-badge's: error, info, success, warning)`}},
		{"a value is listed for each cui-* class of its element", `<a class="cui-console-badge cui-console-link" data-color="purple">b</a>`, []string{`data-color="purple" on <a class="cui-console-badge cui-console-link">: no rule of ciderui's selects it (cui-console-badge's: error, info, success, warning; cui-console-link's: none)`}},
		{"a typo in any selected attribute", `<dl class="cui-console-dl" data-orientation="vertcal"></dl>`, []string{`data-orientation="vertcal" on <dl class="cui-console-dl">: no rule of ciderui's selects it (cui-console-dl's: vertical)`}},
		{"the value itself", `<dl class="cui-console-dl" data-orientation="vertical"></dl>`, nil},

		// [data-sticky] selects any value, on its own class alone
		{"a bare selected attribute", `<table class="cui-console-table" data-sticky></table>`, nil},
		{"a value of a bare selected attribute", `<table class="cui-console-table" data-sticky="true"></table>`, nil},
		{"a bare selected attribute on another class", `<div class="cui-console-card" data-sticky></div>`, []string{`data-sticky="" on <div class="cui-console-card">: no rule of ciderui's selects it (cui-console-card's: none)`}},

		// :not([data-x]) alone: the bare attribute leaves the default, a misspelt
		// value would as well, with no rule of its own
		{"the bare attribute a :not selects", `<ul class="cui-console-tabs" data-overflow-end></ul>`, nil},
		{"a value only a :not selects", `<ul class="cui-console-tabs" data-overflow-end="yes"></ul>`, []string{`data-overflow-end="yes" on <ul class="cui-console-tabs">: no rule of ciderui's selects it (cui-console-tabs's: "")`}},

		// :is(.cui-console-copy-button, .cui-console-copy-inline)[data-state="copied"]
		{"a value for one class of an :is", `<button class="cui-console-copy-button" data-state="copied">x</button>`, nil},
		{"a value for the other class of an :is", `<button class="cui-console-copy-inline" data-state="copied">x</button>`, nil},
		// :where(.cui-btn, .cui-link[data-variant="more"], .cui-console-link, …)
		{"a value is not another alternative's", `<a class="cui-console-link" data-variant="more">x</a>`, []string{`data-variant="more" on <a class="cui-console-link">: no rule of ciderui's selects it (cui-console-link's: label, text)`}},
		// .cui-console-player-end > [data-control="replay"] names no class of the element
		{"an attribute only a rule for another element selects", `<p class="cui-console-player-caption" data-control="caption">x</p>`, nil},
		// cui-console-btn-icon's sizes are its own
		{"a value of another class on the element", `<button class="cui-console-btn cui-console-btn-icon" data-size="sm">x</button>`, nil},

		{"a page's class and data-*", `<div class="mt-2" data-variant="whatever">a</div>`, nil},
		{"a data-* no rule selects", `<a class="cui-console-btn" data-testid="save">x</a>`, nil},

		// a site's whole page, as any writer's HTML has it
		{"a comment", `<!-- <b class="cui-console-badg"> --><b class="cui-console-badge">a</b>`, nil},
		{"after a comment", `<!-- a --><b class="cui-console-badg">a</b>`, noClass("b")},
		{"a comment left open", `<!-- <b class="cui-console-badg">`, nil},
		{"a script", `<script>let s = '<a class="cui-console-badg">';</script>`, nil},
		{"a style", `<style>/* <a class="cui-console-badg"> */</style>`, nil},
		{"a script left open", `<SCRIPT>let s = '<a class="cui-console-badg">';`, nil},
		{"after a script", `<script>a</script><b class="cui-console-badg">a</b>`, noClass("b")},
		{"after a style", `<style>a</style><b class="cui-console-badg">a</b>`, noClass("b")},
		{"inside an element", `<div><b class="cui-console-badg">a</b></div>`, noClass("b")},
		{"a < in text", `1 < 2 <3 <b class="cui-console-badg">a</b>`, noClass("b")},
		{"a < at the end", `<b class="cui-console-badg">a</b><`, noClass("b")},
		{"an uppercase tag", `<B class="cui-console-badg">a</B>`, noClass("B")},
		{"single quotes", `<b class='cui-console-badg'>a</b>`, noClass("b")},
		{"no quotes", `<b class=cui-console-badg>a</b>`, noClass("b")},
		{"no quotes, then an attribute", `<b class=cui-console-badg id=x>a</b>`, noClass("b")},
		{"a quote left open", `<b class="cui-console-badg`, noClass("b")},
		{"a tag left open", `<b class="cui-console-badg"`, noClass("b")},
		{"an attribute before, with no value", `<b hidden class="cui-console-badg">a</b>`, noClass("b")},
		{"self-closed", `<img class="cui-console-badg"/><br/>`, noClass("img")},
		{"lines and tabs between attributes", "<b\n\tid=\"x\"\r\n\fclass=\"cui-console-badg\">a</b>", noClass("b")},
		{"lines between classes", "<b class=\"cui-console-badge\n\tcui-console-badg\">a</b>", []string{"class \"cui-console-badg\" on <b class=\"cui-console-badge\\n\\tcui-console-badg\">: ciderui has no such class"}},
	} {
		t.Run(c.name, func(t *testing.T) {
			if got := CheckMarkup(c.html); !slices.Equal(got, c.want) {
				t.Errorf("got  %q\nwant %q", got, c.want)
			}
		})
	}
}

func TestCheckMarkupReadsThePagesStylesheet(t *testing.T) {
	page := func(stylesheet, body string) string {
		return `<!doctype html><html><head><link rel="stylesheet" href="` + stylesheet + `"></head><body>` + body + `</body></html>`
	}
	for _, c := range []struct {
		name, html string
		want       []string
	}{
		{"a site part on a console page", page(ConsoleStylesheet(), `<button class="cui-btn">a</button>`), []string{`class "cui-btn" on <button class="cui-btn">: the console stylesheet this page links has no such class`}},
		{"a console part on a site page", page(SiteStylesheet(), `<span class="cui-console-badge">a</span>`), []string{`class "cui-console-badge" on <span class="cui-console-badge">: the site stylesheet this page links has no such class`}},
		{"a console part on a console page", page(ConsoleStylesheet(), `<span class="cui-console-badge" data-color="error">a</span>`), nil},
		{"a site part on a site page", page(SiteStylesheet(), `<button class="cui-btn">a</button>`), nil},
		{"the type on a console page", page(ConsoleStylesheet(), `<p class="cui-body">a</p>`), nil},
		{"the type on a site page", page(SiteStylesheet(), `<p class="cui-body">a</p>`), nil},
		{"a fragment links neither, so either set's parts pass", `<button class="cui-btn">a</button><span class="cui-console-badge">a</span>`, nil},
		{"another stylesheet is the site's own", page("/own.css", `<button class="cui-btn">a</button>`), nil},
	} {
		t.Run(c.name, func(t *testing.T) {
			if got := CheckMarkup(c.html); !slices.Equal(got, c.want) {
				t.Errorf("got  %q\nwant %q", got, c.want)
			}
		})
	}
}

func TestClassesIn(t *testing.T) {
	classes := classesIn(`/* .cui-ghost { } */ .cui-real { color: red; } @media (width >= 1px) { .cui-inner { color: red; } } @utility cui-type { color: red; } @utility other { color: red; }`)
	want := map[string]bool{"cui-real": true, "cui-inner": true, "cui-type": true}
	if !maps.Equal(classes, want) {
		t.Errorf("got %v", classes)
	}
}

// each way a selector's compound names a class and a data-*, on one line
func compounds(css string) string {
	var lines []string
	for _, c := range compoundsIn(css) {
		line := "." + strings.Join(c.classes, ".")
		for _, a := range c.attributes {
			not := ""
			if a.negated {
				not = "!"
			}
			line += fmt.Sprintf("[%s%s=%s,%v]", not, a.name, a.value, a.hasValue)
		}
		lines = append(lines, line)
	}
	return strings.Join(lines, " ")
}

func TestCompoundsIn(t *testing.T) {
	for _, c := range []struct{ css, want string }{
		{`.a[data-x="1"] { }`, `.a[data-x=1,true]`},
		{`.a[data-x] { }`, `.a[data-x=,false]`},
		{`.a[title="1"] { }`, `.a`},
		{`.a[data-x="1"], .b[data-y="2"] { }`, `.a[data-x=1,true] .b[data-y=2,true]`},
		{`.a[data-x="1"].b { }`, `.a.b[data-x=1,true]`},
		{`.a:hover[data-x="1"]::before { }`, `.a[data-x=1,true]`},
		{`.a:nth-child(2n + 1)[data-x="1"] { }`, `.a[data-x=1,true]`},
		{`.a:not([data-x]) { }`, `.a[!data-x=,false]`},
		{`.a:not(.b)[data-x="1"] { }`, `.a[data-x=1,true]`},
		{`:is(.a, .b)[data-x="1"] { }`, `.a[data-x=1,true] .b[data-x=1,true]`},
		{`:where(.a, .b[data-x="1"]) { }`, `.a .b[data-x=1,true]`},
		// the subject's first: a compound of another element comes after
		{`.a[data-x="1"] .b[data-y="2"] { }`, `.b[data-y=2,true] .a[data-x=1,true]`},
		{`.a[data-x="1"] > .b { }`, `.b .a[data-x=1,true]`},
		{`.a[data-x="1"] + .b { }`, `.b .a[data-x=1,true]`},
		{`.a[data-x="1"] ~ .b { }`, `.b .a[data-x=1,true]`},
		{`.a:has(.b[data-x="1"]) { }`, `.a .b[data-x=1,true]`},
		{`.a { &[data-x="1"] { } }`, `.a .a[data-x=1,true]`},
		{`.a, .b { &[data-x="1"] { } }`, `.a .b .a[data-x=1,true] .b[data-x=1,true]`},
		{`.a { .b[data-x="1"] { } }`, `.a .b[data-x=1,true] .a`},
		{`@media (width >= 1px) { .a[data-x="1"] { color: red; } }`, `.a[data-x=1,true]`},
		{`.a { color: red; &[data-x="1"] { } }`, `.a .a[data-x=1,true]`},
		{`.a-b1C[data-x="1"] { }`, `.a-b1C[data-x=1,true]`},
	} {
		if got := compounds(c.css); got != c.want {
			t.Errorf("%s\n got %s\nwant %s", c.css, got, c.want)
		}
	}
}

// nothing else would say so: rules would go on reading the rest as selectors
func TestAStylesheetLeftOpenPanics(t *testing.T) {
	for _, c := range []struct{ css, want string }{
		{`/* .a { }`, "an unclosed comment"},
		{`.a:not(.b { }`, "an unclosed ( in .a:not(.b"},
	} {
		t.Run(c.want, func(t *testing.T) {
			defer func() {
				if got := fmt.Sprint(recover()); got != c.want {
					t.Errorf("panicked with %s", got)
				}
			}()
			compoundsIn(c.css)
		})
	}
}
