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
		{"a misspelt class of a part names the nearest class", `<li class="cui-list-detial">a</li>`, []string{`class "cui-list-detial" on <li class="cui-list-detial">: ciderui has no such class; did you mean cui-list-detail`}},
		{"a class of a part with none near names its part's classes", `<li class="cui-list-zzzzzz">a</li>`, []string{`class "cui-list-zzzzzz" on <li class="cui-list-zzzzzz">: ciderui has no such class; the cui-list classes are cui-list, cui-list-aside, cui-list-detail, cui-list-row and cui-list-title`}},
		{"a class of a part of one class with none near names that class", `<span class="cui-console-tab-zzzzzz">a</span>`, []string{`class "cui-console-tab-zzzzzz" on <span class="cui-console-tab-zzzzzz">: ciderui has no such class; the only cui-console-tab class is cui-console-tab`}},
		{"the prefix alone", `<p class="cui-">a</p>`, []string{`class "cui-" on <p class="cui-">: ciderui has no such class; did you mean cui-dl`}},
		{"a misspelt class among others, lines between them", "<p>a</p><b class=\"cui-console-badge\n\tcui-console-badg mt-1\">a</b>", []string{"class \"cui-console-badg\" on <b class=\"cui-console-badge\\n\\tcui-console-badg mt-1\">: ciderui has no such class; did you mean cui-console-badge"}},
	})
}

func TestCheckMarkupNamesTheNearestClass(t *testing.T) {
	checkEach(t, []markupCase{
		{"three edits away", `<span class="cui-console-bnagx">a</span>`, []string{`class "cui-console-bnagx" on <span class="cui-console-bnagx">: ciderui has no such class`}},
		{"two letters replaced", `<span class="cui-console-bxxge">a</span>`, []string{`class "cui-console-bxxge" on <span class="cui-console-bxxge">: ciderui has no such class; did you mean cui-console-badge`}},
		{"two letters too many", `<span class="cui-console-baddgee">a</span>`, []string{`class "cui-console-baddgee" on <span class="cui-console-baddgee">: ciderui has no such class; did you mean cui-console-badge`}},
		{"two letters missing", `<span class="cui-console-bge">a</span>`, []string{`class "cui-console-bge" on <span class="cui-console-bge">: ciderui has no such class; did you mean cui-console-badge`}},
		{"a class after other letters", `<a class="cui-cui-btn">a</a>`, []string{`class "cui-cui-btn" on <a class="cui-cui-btn">: ciderui has no such class`}},
		{"on a page, not the other set's nearest", page(ciderui.ConsoleStylesheet(), `<a class="cui-bnt">a</a>`), []string{`class "cui-bnt" on <a class="cui-bnt">: the console stylesheet this page links has no such class`}},
		{"on a page, not the other set's part", page(ciderui.ConsoleStylesheet(), `<li class="cui-list-zzzzzz">a</li>`), []string{`class "cui-list-zzzzzz" on <li class="cui-list-zzzzzz">: the console stylesheet this page links has no such class`}},
	})
}

// @invariant: cui-console-tab and cui-console-tabs are one edit from it; a
// map's order differs between calls, so one call alone passes half the time
func TestCheckMarkupNamesTheFirstByNameOfTwoClassesAsNear(t *testing.T) {
	want := []string{`class "cui-console-tabz" on <b class="cui-console-tabz">: ciderui has no such class; did you mean cui-console-tab`}
	for range 20 {
		if got := ciderui.CheckMarkup(`<b class="cui-console-tabz">a</b>`); !slices.Equal(got, want) {
			t.Fatalf("CheckMarkup reported %q\nwant     %q", got, want)
		}
	}
}

func TestCheckMarkupNamesASlotsUtility(t *testing.T) {
	var cases []markupCase
	for _, prefix := range []string{"text", "bg", "border", "outline", "decoration", "fill", "stroke", "ring", "divide", "accent", "caret"} {
		class := prefix + "-(--cui-c-console-brand-label-primary)"
		cases = append(cases, markupCase{"a --color-* slot under " + prefix, `<p class="` + class + `">a</p>`, []string{`class "` + class + `" on <p class="` + class + `">: write ` + prefix + `-link, the set's name for this color`}})
	}
	checkEach(t, append(cases, []markupCase{
		{"a --color-* slot, under two variants and with an opacity", `<hr class="dark:md:border-(--cui-c-console-neutral-border-tertiary)/50">`, []string{`class "dark:md:border-(--cui-c-console-neutral-border-tertiary)/50" on <hr class="dark:md:border-(--cui-c-console-neutral-border-tertiary)/50">: write dark:md:border-divider/50, the set's name for this color`}},
		{"a site page's slot", page(ciderui.SiteStylesheet()+`"><link rel="stylesheet" href="/own.css`, `<p class="text-(--cui-c-text-2)">a</p>`), []string{`class "text-(--cui-c-text-2)" on <p class="text-(--cui-c-text-2)">: write text-secondary, the set's name for this color`}},
		{"a ground's slot", `<div class="bg-(--cui-c-console-card)">a</div>`, []string{`class "bg-(--cui-c-console-card)" on <div class="bg-(--cui-c-console-card)">: write bg-alt, the set's name for this color`}},
		{"a fragment's site slot", `<p class="text-(--cui-c-text-2)">a</p>`, []string{`class "text-(--cui-c-text-2)" on <p class="text-(--cui-c-text-2)">: write text-secondary, the set's name for this color`}},
		{"a slot left open", `<p class="text-(--cui-c-text-2">a</p>`, nil},
		{"a slot with no utility name, under a variant and with an opacity", `<p class="md:text-(--cui-c-console-brand-label-secondary)/50">a</p>`, nil},
		{"a text color's slot as a ground, which no utility names", `<p class="bg-(--cui-c-console-neutral-label-secondary)">a</p>`, nil},
		{"the other set's slot on a page", page(ciderui.SiteStylesheet()+`"><link rel="stylesheet" href="/own.css`, `<p class="text-(--cui-c-console-neutral-label-secondary)">a</p>`), nil},
	}...))
}

func TestCheckMarkupRefusesADataValueNoRuleSelects(t *testing.T) {
	checkEach(t, []markupCase{
		{"each class that takes the attribute", `<div class="cui-console-btn cui-console-empty" data-size="xl">a</div>`, []string{`data-size="xl" on <div class="cui-console-btn cui-console-empty">: cui-console-btn takes data-size "lg" or "sm", or none for its default; cui-console-empty takes data-size "lg" or "sm"`}},

		// @invariant: [data-intro] selects any value, on its own class alone
		{"a value of an attribute selected bare in a list of selectors", `<div class="cui-console-player" data-intro="x"></div>`, nil},
		{"an attribute none of its classes takes names those they take, bare and with values", `<div class="cui-console-player" data-color="red"></div>`, []string{`data-color="red" on <div class="cui-console-player">: none of its classes takes data-color; cui-console-player takes data-controls "hidden"; cui-console-player takes data-end; cui-console-player takes data-full-screen; cui-console-player takes data-intro; cui-console-player takes data-muted; cui-console-player takes data-state "playing"`}},
		{"a bare selected attribute on another class", `<div class="cui-console-card" data-stack></div>`, []string{`data-stack="" on <div class="cui-console-card">: none of its classes takes data-stack`}},

		// @invariant: .cui-console-player-end > [data-control="replay"] names no class of the element
		{"an attribute only a rule for another element selects", `<p class="cui-console-player-heading" data-control="heading">x</p>`, nil},
	})
}

func TestCheckMarkupRefusesAClassWithNoDefaultWithoutItsAttribute(t *testing.T) {
	checkEach(t, []markupCase{
		{"a banner of no color", `<div class="cui-console-banner">a</div>`, []string{`class "cui-console-banner" on <div class="cui-console-banner">: write data-color "error", "info", "success" or "warning"; it has no default`}},
		{"an empty state of no size", `<div class="cui-console-empty">a</div>`, []string{`class "cui-console-empty" on <div class="cui-console-empty">: write data-size "lg" or "sm"; it has no default`}},
		{"a mark of no state", `<span class="cui-console-strip-mark">a</span>`, []string{`class "cui-console-strip-mark" on <span class="cui-console-strip-mark">: write data-state "waiting" or "working"; it has no default`}},
		{"the attribute in capitals", `<div class="cui-console-empty" DATA-SIZE="sm">a</div>`, nil},
		{"on a page of the other set, the class alone", page(ciderui.SiteStylesheet(), `<div class="cui-console-banner">a</div>`), []string{`class "cui-console-banner" on <div class="cui-console-banner">: a console class; this page links the site stylesheet`}},
	})
}

// @external: ARIA's values of each stateful attribute, in the order the problem lists them
func TestCheckMarkupRefusesAnAriaStateValueARIALacks(t *testing.T) {
	cases := []markupCase{
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
		return []string{`class "cui-console-badg" on <` + tag + ` class="cui-console-badg">: ciderui has no such class; did you mean cui-console-badge`}
	}
	checkEach(t, []markupCase{
		{"after a comment", `<!-- a --><b class="cui-console-badg">a</b>`, noClass("b")},
		{"a comment left open", `<!-- <b class="cui-console-badg">`, nil},
		{"a > in a comment", `<!-- a > <b class="cui-console-badg"> -->`, nil},
		{"a script that writes a closing tag", `<script>let s = '</b><a class="cui-console-badg">';</script>`, nil},
		{"a style in capitals", `<STYLE>/* <a class="cui-console-badg"> */</STYLE>`, nil},
		{"a script left open", `<SCRIPT>let s = '<a class="cui-console-badg">';`, nil},
		{"after a script", `<script>a</script><b class="cui-console-badg">a</b>`, noClass("b")},
		{"a < in text", `1 < 2 <3 <b class="cui-console-badg">a</b>`, noClass("b")},
		{"a < at the end", `<b class="cui-console-badg">a</b><`, noClass("b")},
		{"single quotes", `<b class='cui-console-badg'>a</b>`, noClass("b")},
		{"no quotes, between attributes and at the tag's end", `<b id=x class=cui-console-badg>a</b>`, noClass("b")},
		{"a quote left open", `<b class="cui-console-badg`, noClass("b")},
		{"an = at the end", `<b class=`, nil},
		{"an attribute of no value before the class", `<button disabled class="cui-console-badg">a</button>`, noClass("button")},
		{"slashes between attributes", `<img/hidden/class="cui-console-badg">`, noClass("img")},
		{"tabs between attributes", "<b\tid=\"x\"\tclass=\"cui-console-badg\">a</b>", noClass("b")},
		{"lines between attributes", "<b\nid=\"x\"\nclass=\"cui-console-badg\">a</b>", noClass("b")},
		{"carriage returns between attributes", "<b\rid=\"x\"\rclass=\"cui-console-badg\">a</b>", noClass("b")},
		{"form feeds between attributes", "<b\fid=\"x\"\fclass=\"cui-console-badg\">a</b>", noClass("b")},
	})
}

func TestCheckMarkupReadsAnAttributeNameInAnyCase(t *testing.T) {
	checkEach(t, []markupCase{
		{"an aria state in capitals", `<li class="cui-console-menu-item" ARIA-CURRENT="pgae">a</li>`, []string{`ARIA-CURRENT="pgae" on <li class="cui-console-menu-item">: not a value of aria-current (date, false, location, page, step, time, true)`}},
		{"a data-* in capitals", `<span class="cui-console-badge" DATA-COLOR="purple">b</span>`, []string{`DATA-COLOR="purple" on <span class="cui-console-badge">: cui-console-badge takes data-color "error", "info", "success" or "warning", or none for its default`}},
	})
}

func TestCheckMarkupReadsThePagesStylesheet(t *testing.T) {
	checkEach(t, []markupCase{
		{"a data-* only the other set selects is the page's own", page(ciderui.SiteStylesheet(), `<a class="cui-btn" data-color="purple">a</a>`), nil},
		{"a link in the text to a stylesheet is not the page's", `<a href="` + ciderui.ConsoleStylesheet() + `">a</a><button class="cui-btn">a</button>`, nil},
		{"a site part on a console page in capitals", `<!doctype html><HTML><HEAD><LINK rel="stylesheet" href="` + ciderui.ConsoleStylesheet() + `"></HEAD><BODY><button class="cui-btn">a</button></BODY></HTML>`, []string{`class "cui-btn" on <button class="cui-btn">: a site class; this page links the console stylesheet`}},
	})
}

// @external: the browser keeps an attribute's first copy and drops the rest
// with no error: a part's own aria-invalid and the page's through Attrs
func TestCheckMarkupRefusesAnAttributeWrittenTwice(t *testing.T) {
	checkEach(t, []markupCase{
		{"each copy in another case", `<b CLASS="cui-body" id="a" Class="x">a</b>`, []string{`Class on <b class="cui-body">: written twice, and the browser keeps the first`}},
	})
}

func TestCheckMarkupRefusesAReferenceToNoIDOnAPage(t *testing.T) {
	sitePage := func(body string) string { return page(ciderui.SiteStylesheet(), body) }
	checkEach(t, []markupCase{
		{"a link to the top", sitePage(`<a href="#">Top</a>`), nil},
		{"a link to another page", sitePage(`<a href="/a#b">a</a>`), nil},
		{"a value that begins with #", sitePage(`<input value="#ffffff">`), nil},
		{"each id of a list", sitePage(`<input id="a" aria-describedby="a-help a-error"><p id="a-help">h</p>`), []string{`aria-describedby="a-help a-error" on <input>: this page has no id "a-error"`}},
		{"a popover's button with no popover", sitePage(`<button popovertarget="menu">a</button>`), []string{`popovertarget="menu" on <button>: this page has no id "menu"`}},
		{"a command's button", sitePage(`<button commandfor="dialog" command="show-modal">a</button>`), []string{`commandfor="dialog" on <button>: this page has no id "dialog"`}},
		{"a section labelled by no heading", sitePage(`<section aria-labelledby="t"><h2>a</h2></section>`), []string{`aria-labelledby="t" on <section>: this page has no id "t"`}},
		{"a controlled element", sitePage(`<button aria-controls="p">a</button>`), []string{`aria-controls="p" on <button>: this page has no id "p"`}},
		{"a skip link in capitals", sitePage(`<A HREF="#main">Skip</A>`), []string{`HREF="#main" on <A>: this page has no id "main"`}},
		{"a page in capitals", `<!DOCTYPE html><HTML><BODY><LABEL FOR="name">Name</LABEL></BODY></HTML>`, []string{`FOR="name" on <LABEL>: this page has no id "name"`}},
		{"an ID in capitals", sitePage(`<label for="name">Name</label><input ID="name">`), nil},
		{"a fragment's may be on the page around it", `<button popovertarget="menu">a</button><label for="name">Name</label>`, nil},
	})
}
