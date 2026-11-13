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
		{"a misspelt class names its part's classes", `<li class="cui-list-detial">a</li>`, []string{`class "cui-list-detial" on <li class="cui-list-detial">: ciderui has no such class; the cui-list classes are cui-list, cui-list-aside, cui-list-detail, cui-list-row and cui-list-title`}},
		{"a misspelt console class names its part's classes", `<span class="cui-console-badge-lg">a</span>`, []string{`class "cui-console-badge-lg" on <span class="cui-console-badge-lg">: ciderui has no such class; the only cui-console-badge class is cui-console-badge`}},
		{"lines between classes", "<b class=\"cui-console-badge\n\tcui-console-badg\">a</b>", []string{"class \"cui-console-badg\" on <b class=\"cui-console-badge\\n\\tcui-console-badg\">: ciderui has no such class"}},
	})
}

func TestCheckMarkupRefusesADataValueNoRuleSelects(t *testing.T) {
	checkEach(t, []markupCase{
		{"a value no rule selects", `<span class="cui-console-badge" data-color="purple">b</span>`, []string{`data-color="purple" on <span class="cui-console-badge">: cui-console-badge takes data-color "error", "info", "success" or "warning", or none for its default`}},
		{"a class that takes none of the attribute is left out", `<a class="cui-console-badge cui-console-link" data-color="purple">b</a>`, []string{`data-color="purple" on <a class="cui-console-badge cui-console-link">: cui-console-badge takes data-color "error", "info", "success" or "warning", or none for its default`}},
		{"each class that takes the attribute", `<div class="cui-console-btn cui-console-empty" data-size="xl">a</div>`, []string{`data-size="xl" on <div class="cui-console-btn cui-console-empty">: cui-console-btn takes data-size "lg" or "sm"; cui-console-empty takes data-size "lg" or "sm"; or none for their defaults`}},
		{"each value is listed once", `<button class="cui-console-btn cui-console-play" data-size="xl">a</button>`, []string{`data-size="xl" on <button class="cui-console-btn cui-console-play">: cui-console-btn takes data-size "lg" or "sm", or none for its default`}},

		{"a value selected next to a :not", `<div class="cui-console-player" data-controls="hiden"></div>`, []string{`data-controls="hiden" on <div class="cui-console-player">: cui-console-player takes data-controls "hidden", or none for its default`}},

		// @invariant: [data-stack] selects any value, on its own class alone
		{"a value of a bare selected attribute", `<table class="cui-console-table" data-stack="true"></table>`, nil},
		{"an attribute none of its classes takes names those they take", `<span class="cui-console-badge" data-size="sm">b</span>`, []string{`data-size="sm" on <span class="cui-console-badge">: none of its classes takes data-size; cui-console-badge takes data-color "error", "info", "success" or "warning"`}},
		{"a bare selected attribute among those it takes", `<table class="cui-console-table" data-size="sm"></table>`, []string{`data-size="sm" on <table class="cui-console-table">: none of its classes takes data-size; cui-console-table takes data-stack`}},
		{"a bare selected attribute on another class", `<div class="cui-console-card" data-stack></div>`, []string{`data-stack="" on <div class="cui-console-card">: none of its classes takes data-stack`}},

		// @invariant: .cui-console-player-end > [data-control="replay"] names no class of the element
		{"an attribute only a rule for another element selects", `<p class="cui-console-player-heading" data-control="heading">x</p>`, nil},
		// @invariant: cui-console-btn-icon's sizes are its own
		{"a value of another class on the element", `<button class="cui-console-btn cui-console-btn-icon" data-size="sm">x</button>`, nil},
		{"a data-* no rule selects", `<a class="cui-console-btn" data-testid="save">x</a>`, nil},
	})
}

// @external: ARIA's values of each stateful attribute, in the order the problem lists them
func TestCheckMarkupRefusesAnAriaStateValueARIALacks(t *testing.T) {
	cases := []markupCase{
		{"an aria state's other legal values", `<li class="cui-console-menu-item" aria-current="location">a</li>`, nil},
		{"an aria state with no class of ciderui's", `<td aria-sort="desc">a</td>`, nil},
	}
	for name, values := range map[string]string{
		"aria-checked":  "false, mixed, true",
		"aria-current":  "date, false, location, page, step, time, true",
		"aria-disabled": "false, true",
		"aria-expanded": "false, true",
		"aria-invalid":  "false, grammar, spelling, true",
		"aria-pressed":  "false, mixed, true",
		"aria-selected": "false, true",
		"aria-sort":     "ascending, descending, none, other",
	} {
		cases = append(cases, markupCase{"a typo in " + name, `<li class="cui-console-menu-item" ` + name + `="typo">a</li>`, []string{name + `="typo" on <li class="cui-console-menu-item">: not a value of ` + name + " (" + values + ")"}})
	}
	checkEach(t, cases)
}

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
		{"slashes between attributes", `<img/hidden/class="cui-console-badg">`, noClass("img")},
		{"tabs between attributes", "<b\tid=\"x\"\tclass=\"cui-console-badg\">a</b>", noClass("b")},
		{"lines between attributes", "<b\nid=\"x\"\nclass=\"cui-console-badg\">a</b>", noClass("b")},
		{"carriage returns between attributes", "<b\rid=\"x\"\rclass=\"cui-console-badg\">a</b>", noClass("b")},
		{"form feeds between attributes", "<b\fid=\"x\"\fclass=\"cui-console-badg\">a</b>", noClass("b")},
	})
}

func TestCheckMarkupReadsAnAttributeNameInAnyCase(t *testing.T) {
	checkEach(t, []markupCase{
		{"a class in capitals", `<b CLASS="cui-console-badg">a</b>`, []string{`class "cui-console-badg" on <b class="cui-console-badg">: ciderui has no such class`}},
		{"an aria state in capitals", `<li class="cui-console-menu-item" ARIA-CURRENT="pgae">a</li>`, []string{`ARIA-CURRENT="pgae" on <li class="cui-console-menu-item">: not a value of aria-current (date, false, location, page, step, time, true)`}},
		{"a data-* in capitals", `<span class="cui-console-badge" DATA-COLOR="purple">b</span>`, []string{`DATA-COLOR="purple" on <span class="cui-console-badge">: cui-console-badge takes data-color "error", "info", "success" or "warning", or none for its default`}},
	})
}

func TestCheckMarkupReadsThePagesStylesheet(t *testing.T) {
	checkEach(t, []markupCase{
		{"a site part on a console page", page(ciderui.ConsoleStylesheet(), `<button class="cui-btn">a</button>`), []string{`class "cui-btn" on <button class="cui-btn">: a site class; this page links the console stylesheet`}},
		{"a console part on a site page", page(ciderui.SiteStylesheet(), `<span class="cui-console-badge">a</span>`), []string{`class "cui-console-badge" on <span class="cui-console-badge">: a console class; this page links the site stylesheet`}},
		{"a value no rule selects on a console page", page(ciderui.ConsoleStylesheet(), `<span class="cui-console-badge" data-color="purple">a</span>`), []string{`data-color="purple" on <span class="cui-console-badge">: cui-console-badge takes data-color "error", "info", "success" or "warning", or none for its default`}},
		{"a console part on a console page", page(ciderui.ConsoleStylesheet(), `<span class="cui-console-badge" data-color="error">a</span>`), nil},
		{"a misspelt class on a console page", page(ciderui.ConsoleStylesheet(), `<span class="cui-console-badg">a</span>`), []string{`class "cui-console-badg" on <span class="cui-console-badg">: the console stylesheet this page links has no such class`}},
		{"a data-* only the other set selects is the page's own", page(ciderui.SiteStylesheet(), `<a class="cui-btn" data-color="purple">a</a>`), nil},
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
		{"a skip link to no main", sitePage(`<a class="cui-btn cui-skip" href="#main">Skip</a><main>a</main>`), []string{`href="#main" on <a class="cui-btn cui-skip">: this page has no id "main"`}},
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
