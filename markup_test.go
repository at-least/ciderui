package ciderui_test

import (
	"os"
	"path/filepath"
	"slices"
	"testing"

	"github.com/at-least/ciderui"
)

type markupCase struct {
	name, html string
	want       []string
}

func checkEach(t *testing.T, cases []markupCase) {
	t.Helper()
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := ciderui.CheckMarkup(c.html); !slices.Equal(got, c.want) {
				t.Errorf("CheckMarkup(%s)\nreported %q\nwant     %q", c.html, got, c.want)
			}
		})
	}
}

func page(stylesheet, body string) string {
	return `<!doctype html><html><head><link rel="stylesheet" href="` + stylesheet + `"></head><body>` + body + `</body></html>`
}

func TestEveryExamplePasses(t *testing.T) {
	paths, err := filepath.Glob("components/*/*/examples/*.html")
	if err != nil {
		t.Fatal(err)
	}
	if len(paths) <= 50 {
		t.Fatalf("%d examples read", len(paths))
	}
	for _, path := range paths {
		html, err := os.ReadFile(path)
		if err != nil {
			t.Fatal(err)
		}
		if problems := ciderui.CheckMarkup(string(html)); problems != nil {
			t.Errorf("%s: %q", path, problems)
		}
	}
}

func TestCheckMarkupRefusesAClassTheKitLacks(t *testing.T) {
	checkEach(t, []markupCase{
		{"a misspelt class", `<p>a</p><span class="cui-console-badg mt-1">b</span>`, []string{`class "cui-console-badg" on <span class="cui-console-badg mt-1">: ciderui has no such class`}},
		{"lines between classes", "<b class=\"cui-console-badge\n\tcui-console-badg\">a</b>", []string{"class \"cui-console-badg\" on <b class=\"cui-console-badge\\n\\tcui-console-badg\">: ciderui has no such class"}},
	})
}

func TestCheckMarkupRefusesADataValueNoRuleSelects(t *testing.T) {
	checkEach(t, []markupCase{
		{"a value no rule selects", `<span class="cui-console-badge" data-color="purple">b</span>`, []string{`data-color="purple" on <span class="cui-console-badge">: no rule of ciderui selects it (cui-console-badge: error, info, success, warning; default: no data-color)`}},
		{"a value is listed for each cui-* class of its element", `<a class="cui-console-badge cui-console-link" data-color="purple">b</a>`, []string{`data-color="purple" on <a class="cui-console-badge cui-console-link">: no rule of ciderui selects it (cui-console-badge: error, info, success, warning; cui-console-link: none; default: no data-color)`}},
		{"each value is listed once", `<button class="cui-console-btn" data-size="xl">a</button>`, []string{`data-size="xl" on <button class="cui-console-btn">: no rule of ciderui selects it (cui-console-btn: lg, sm; default: no data-size)`}},

		// @invariant: [data-stack] selects any value, on its own class alone
		{"a value of a bare selected attribute", `<table class="cui-console-table" data-stack="true"></table>`, nil},
		{"a bare selected attribute on another class", `<div class="cui-console-card" data-stack></div>`, []string{`data-stack="" on <div class="cui-console-card">: no rule of ciderui selects it (cui-console-card: none; default: no data-stack)`}},

		// @invariant: .cui-console-player-end > [data-control="replay"] names no class of the element
		{"an attribute only a rule for another element selects", `<p class="cui-console-player-caption" data-control="caption">x</p>`, nil},
		// @invariant: cui-console-btn-icon's sizes are its own
		{"a value of another class on the element", `<button class="cui-console-btn cui-console-btn-icon" data-size="sm">x</button>`, nil},
		{"a data-* no rule selects", `<a class="cui-console-btn" data-testid="save">x</a>`, nil},
	})
}

func TestCheckMarkupRefusesAnAriaStateValueARIALacks(t *testing.T) {
	checkEach(t, []markupCase{
		{"a misspelt aria state", `<li class="cui-console-menu-item" aria-current="pgae">a</li>`, []string{`aria-current="pgae" on <li class="cui-console-menu-item">: not a value of aria-current (date, false, location, page, step, time, true)`}},
		{"an aria state's other legal values", `<li class="cui-console-menu-item" aria-current="location">a</li>`, nil},
		{"an aria state with no class of ciderui's", `<td aria-sort="desc">a</td>`, nil},
	})
}

// @external: a site's whole page, as any writer's HTML has it
func TestCheckMarkupReadsHTMLAsABrowserDoes(t *testing.T) {
	noClass := func(tag string) []string {
		return []string{`class "cui-console-badg" on <` + tag + ` class="cui-console-badg">: ciderui has no such class`}
	}
	checkEach(t, []markupCase{
		{"a comment", `<!-- <b class="cui-console-badg"> --><b class="cui-console-badge">a</b>`, nil},
		{"after a comment", `<!-- a --><b class="cui-console-badg">a</b>`, noClass("b")},
		{"a comment left open", `<!-- <b class="cui-console-badg">`, nil},
		{"a script", `<script>let s = '<a class="cui-console-badg">';</script>`, nil},
		{"a style", `<style>/* <a class="cui-console-badg"> */</style>`, nil},
		{"a script left open", `<SCRIPT>let s = '<a class="cui-console-badg">';`, nil},
		{"after a script", `<script>a</script><b class="cui-console-badg">a</b>`, noClass("b")},
		{"a < in text", `1 < 2 <3 <b class="cui-console-badg">a</b>`, noClass("b")},
		{"a < at the end", `<b class="cui-console-badg">a</b><`, noClass("b")},
		{"an uppercase tag", `<B class="cui-console-badg">a</B>`, noClass("B")},
		{"single quotes", `<b class='cui-console-badg'>a</b>`, noClass("b")},
		{"no quotes, then an attribute", `<b class=cui-console-badg id=x>a</b>`, noClass("b")},
		{"a quote left open", `<b class="cui-console-badg`, noClass("b")},
		{"self-closed", `<img class="cui-console-badg"/><br/>`, noClass("img")},
		{"lines and tabs between attributes", "<b\r\tid=\"x\"\n\fclass=\"cui-console-badg\">a</b>", noClass("b")},
	})
}

func TestCheckMarkupReadsThePagesStylesheet(t *testing.T) {
	checkEach(t, []markupCase{
		{"a site part on a console page", page(ciderui.ConsoleStylesheet(), `<button class="cui-btn">a</button>`), []string{`class "cui-btn" on <button class="cui-btn">: a site class; this page links the console stylesheet`}},
		{"a console part on a site page", page(ciderui.SiteStylesheet(), `<span class="cui-console-badge">a</span>`), []string{`class "cui-console-badge" on <span class="cui-console-badge">: a console class; this page links the site stylesheet`}},
		{"a value no rule selects on a console page", page(ciderui.ConsoleStylesheet(), `<span class="cui-console-badge" data-color="purple">a</span>`), []string{`data-color="purple" on <span class="cui-console-badge">: no rule of the console stylesheet this page links selects it (cui-console-badge: error, info, success, warning; default: no data-color)`}},
		{"a console part on a console page", page(ciderui.ConsoleStylesheet(), `<span class="cui-console-badge" data-color="error">a</span>`), nil},
		{"the type on a console page", page(ciderui.ConsoleStylesheet(), `<p class="cui-body">a</p>`), nil},
		{"another stylesheet is the site's own", page("/own.css", `<button class="cui-btn">a</button>`), nil},
		{"a page in capitals", `<!doctype html><HTML><HEAD><LINK rel="stylesheet" href="` + ciderui.ConsoleStylesheet() + `"></HEAD><BODY><button class="cui-btn">a</button></BODY></HTML>`, []string{`class "cui-btn" on <button class="cui-btn">: a site class; this page links the console stylesheet`}},
	})
}

// @external: the browser keeps an attribute's first copy and drops the rest
// with no error: a part's own aria-invalid and the page's through Attrs
func TestCheckMarkupRefusesAnAttributeWrittenTwice(t *testing.T) {
	checkEach(t, []markupCase{
		{"twice", `<input aria-invalid="true" id="a" aria-invalid="true">`, []string{`aria-invalid on <input>: written twice, and the browser keeps the first`}},
		{"in another case", `<b class="cui-body" CLASS="x">a</b>`, []string{`CLASS on <b class="cui-body">: written twice, and the browser keeps the first`}},
	})
}

func TestCheckMarkupRefusesAReferenceToNoIDOnAPage(t *testing.T) {
	sitePage := func(body string) string { return page(ciderui.SiteStylesheet(), body) }
	checkEach(t, []markupCase{
		{"a skip link to no main", sitePage(`<a class="cui-btn cui-skip" href="#main">Skip</a><main>a</main>`), []string{`href="#main" on <a>: this page has no id "main"`}},
		{"a link to the top", sitePage(`<a href="#">Top</a>`), nil},
		{"a link to another page", sitePage(`<a href="/a#b">a</a>`), nil},
		{"a value that begins with #", sitePage(`<input value="#ffffff">`), nil},
		{"each id of a list", sitePage(`<input id="a" aria-describedby="a-help a-error"><p id="a-help">h</p>`), []string{`aria-describedby="a-help a-error" on <input>: this page has no id "a-error"`}},
		{"a label for no input", sitePage(`<label for="name">Name</label>`), []string{`for="name" on <label>: this page has no id "name"`}},
		{"a popover's button with no popover", sitePage(`<button popovertarget="menu">a</button>`), []string{`popovertarget="menu" on <button>: this page has no id "menu"`}},
		{"a command's button", sitePage(`<button commandfor="dialog" command="show-modal">a</button>`), []string{`commandfor="dialog" on <button>: this page has no id "dialog"`}},
		{"a section labelled by no heading", sitePage(`<section aria-labelledby="t"><h2>a</h2></section>`), []string{`aria-labelledby="t" on <section>: this page has no id "t"`}},
		{"a controlled element", sitePage(`<button aria-controls="p">a</button>`), []string{`aria-controls="p" on <button>: this page has no id "p"`}},
		{"a page in capitals", `<!DOCTYPE html><HTML><BODY><LABEL FOR="name">Name</LABEL></BODY></HTML>`, []string{`FOR="name" on <LABEL>: this page has no id "name"`}},
		{"an ID in capitals", sitePage(`<label for="name">Name</label><input ID="name">`), nil},
		{"a fragment's may be on the page around it", `<button popovertarget="menu">a</button><label for="name">Name</label>`, nil},
	})
}
