package example_test

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/at-least/ciderui"
	"github.com/at-least/ciderui/go/internal/example"
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

func TestCheckFailsOnlyWhereTheNodeDiffersFromItsExample(t *testing.T) {
	const path = "site/ribbon/examples/ribbon.html"
	same := g.Raw(example.Read(path))
	var pass, fail recorder
	example.Check(&pass, same, path)
	if pass.failures != 0 {
		t.Errorf("Check failed %d times on the example's own markup, want 0", pass.failures)
	}
	example.Check(&fail, g.Group{same, g.Text("x")}, path)
	if fail.failures != 1 {
		t.Errorf("Check failed %d times on markup the example does not have, want 1", fail.failures)
	}
}

func TestJoinTrimmedTrimsEachLineAndJoinsThem(t *testing.T) {
	if got, want := example.JoinTrimmed("<p>\n  <b>a b</b>\t\n</p>\n"), "<p><b>a b</b></p>"; got != want {
		t.Errorf("JoinTrimmed wrote %s, want %s", got, want)
	}
}

func TestRenderWritesQuotesAsTheExamplesDo(t *testing.T) {
	got := example.Render(h.P(h.Title(`it's "a"`), g.Text(`it's "a"`)))
	if want := `<p title="it's &quot;a&quot;">it's &quot;a&quot;</p>`; got != want {
		t.Errorf("Render wrote %s, want %s", got, want)
	}
}

func TestRenderPanicsWithEachProblemCheckMarkupFinds(t *testing.T) {
	const html = `<p class="cui-no-such"></p><p class="cui-nor-this"></p>`
	problems := ciderui.CheckMarkup(html)
	if len(problems) != 2 {
		t.Fatalf("CheckMarkup found %d problems in two unknown classes, want 2: %q", len(problems), problems)
	}
	want := "example.Render: " + problems[0] + "; " + problems[1]
	defer func() {
		if got := fmt.Sprint(recover()); got != want {
			t.Errorf("Render of a refused page panicked with %q, want %q", got, want)
		}
	}()
	example.Render(g.Raw(html))
}

type part struct{ name string }

func (p part) Render(w io.Writer) error {
	if p.name == "" {
		panic("part needs a name")
	}
	_, err := io.WriteString(w, p.name)
	return err
}

func TestPanicsFailsUnlessEachBreakPanicsWithItsWant(t *testing.T) {
	valid := func() part { return part{name: "a"} }
	blank := func(p *part) { p.name = "" }
	for _, c := range []struct {
		name     string
		valid    func() part
		brk      example.Break[part]
		failures int
	}{
		{"a break that panics with its want", valid, example.Break[part]{Want: "part needs a name", Change: blank}, 0},
		{"a break that panics with another message", valid, example.Break[part]{Want: "part needs", Change: blank}, 1},
		{"a break that does not panic", valid, example.Break[part]{Want: "part needs a name", Change: func(*part) {}}, 1},
		{"a valid part that panics", func() part { return part{name: ""} }, example.Break[part]{Want: "part needs a name", Change: blank}, 1},
	} {
		var r recorder
		example.Panics(&r, c.valid, []example.Break[part]{c.brk})
		if r.failures != c.failures {
			t.Errorf("%s: Panics failed %d times, want %d", c.name, r.failures, c.failures)
		}
	}
}

type classed struct{ Class string }

func (c classed) Render(w io.Writer) error {
	return h.P(g.If(c.Class != "drop", g.Attr("class", "root "+c.Class))).Render(w)
}

func TestPagesClassFailsOnAPartThatDropsItsClassOrHasNoRow(t *testing.T) {
	dir := t.TempDir()
	for name, source := range map[string]string{
		"p.go":      "package p\n\ntype Classed struct{ Class string }\n\ntype hidden struct{ Class string }\n\ntype Plain struct{ Name string }\n",
		"p_test.go": "package p\n\ntype TestOnly struct{ Class string }\n",
	} {
		if err := os.WriteFile(filepath.Join(dir, name), []byte(source), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	// @external: %T names this package's type example_test.classed, not p.Classed: nested stands in
	writes := example.ClassRow{Root: "root", Node: classed{Class: example.PageClass}}
	drops := example.ClassRow{Root: "root", Node: classed{Class: "drop"}}
	for _, c := range []struct {
		name     string
		nested   []string
		rows     []example.ClassRow
		failures int
	}{
		{"each exported struct of the package held", []string{"Classed"}, []example.ClassRow{writes}, 0},
		{"a part that drops its Class", []string{"Classed"}, []example.ClassRow{drops}, 1},
		{"a struct with a Class and no row", nil, []example.ClassRow{writes}, 1},
	} {
		var r recorder
		example.PagesClass(&r, dir, c.nested, c.rows)
		if r.failures != c.failures {
			t.Errorf("%s: PagesClass failed %d times, want %d", c.name, r.failures, c.failures)
		}
	}
}
