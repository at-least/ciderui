package example

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/at-least/ciderui"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// @trap: a comparator that has never failed proves nothing: every part's test rests on these
type recorder struct {
	testing.TB
	failures int
}

func (r *recorder) Helper() {}

func (r *recorder) Errorf(string, ...any) { r.failures++ }

func TestCheck(t *testing.T) {
	const path = "site/ribbon/examples/ribbon.html"
	same := g.Raw(Read(path))
	var pass, fail recorder
	Check(&pass, same, path)
	if pass.failures != 0 {
		t.Error("failed on its own example")
	}
	Check(&fail, g.Group{same, g.Text("x")}, path)
	if fail.failures != 1 {
		t.Errorf("%d failures on another page, want 1", fail.failures)
	}
}

func TestJoinTrimmed(t *testing.T) {
	if got := JoinTrimmed("<p>\n  <b>a b</b>\t\n</p>\n"); got != "<p><b>a b</b></p>" {
		t.Errorf("got %s", got)
	}
}

func TestRenderWritesQuotesAsTheExamplesDo(t *testing.T) {
	got := Render(h.P(h.Title(`it's "a"`), g.Text(`it's "a"`)))
	want := `<p title="it's &quot;a&quot;">it's &quot;a&quot;</p>`
	if got != want {
		t.Errorf("got %s", got)
	}
}

func TestRenderPanicsOnWhatCheckMarkupRefuses(t *testing.T) {
	const html = `<p class="cui-no-such"></p>`
	if got, want := panicOf(g.Raw(html)), fmt.Sprintf("%q", ciderui.CheckMarkup(html)); got != want || got == "[]" {
		t.Errorf("got %s", got)
	}
}

type part struct{ name string }

func (p part) Render(w io.Writer) error {
	if p.name == "" {
		panic("part needs a name")
	}
	_, err := io.WriteString(w, p.name)
	return err
}

func TestPanics(t *testing.T) {
	valid := func() part { return part{name: "a"} }
	blank := func(p *part) { p.name = "" }
	for _, c := range []struct {
		name     string
		brk      Break[part]
		failures int
	}{
		{"its message", Break[part]{"part needs a name", blank}, 0},
		{"another message", Break[part]{"part needs", blank}, 1},
		{"no panic", Break[part]{"part needs a name", func(*part) {}}, 1},
	} {
		var r recorder
		Panics(&r, valid, []Break[part]{c.brk})
		if r.failures != c.failures {
			t.Errorf("%s: %d failures, want %d", c.name, r.failures, c.failures)
		}
	}
	var r recorder
	Panics(&r, func() part { return part{name: ""} }, nil)
	if r.failures != 1 {
		t.Errorf("a valid part that panics: %d failures, want 1", r.failures)
	}
}

type classed struct{ Class string }

func (c classed) Render(w io.Writer) error {
	return h.P(g.If(c.Class != "drop", g.Attr("class", "root "+c.Class))).Render(w)
}

func TestPagesClass(t *testing.T) {
	dir := t.TempDir()
	source := "package p\n\ntype Classed struct{ Class string }\n\ntype hidden struct{ Class string }\n\ntype Plain struct{ Name string }\n"
	if err := os.WriteFile(filepath.Join(dir, "p.go"), []byte(source), 0o644); err != nil {
		t.Fatal(err)
	}
	// @external: %T names this package's type example.classed, not p.Classed: nested stands in
	writes := ClassRow{Root: "root", Node: classed{Class: PageClass}}
	drops := ClassRow{Root: "root", Node: classed{Class: "drop"}}
	for _, c := range []struct {
		name     string
		nested   []string
		rows     []ClassRow
		failures int
	}{
		{"each struct held", []string{"Classed"}, []ClassRow{writes}, 0},
		{"a part that drops it", []string{"Classed"}, []ClassRow{drops}, 1},
		{"a struct with no row", nil, []ClassRow{writes}, 1},
	} {
		var r recorder
		PagesClass(&r, dir, c.nested, c.rows)
		if r.failures != c.failures {
			t.Errorf("%s: %d failures, want %d", c.name, r.failures, c.failures)
		}
	}
}
