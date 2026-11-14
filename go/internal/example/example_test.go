package example_test

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
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
	message  string
}

func (r *recorder) Helper() {}

func (r *recorder) Errorf(format string, args ...any) {
	r.failures++
	r.message = fmt.Sprintf(format, args...)
}

func TestCheckComparesTheExamplesLinesTrimmedAndJoined(t *testing.T) {
	const path = "../go/internal/example/testdata/indented.html"
	for _, c := range []struct {
		name     string
		node     g.Node
		failures int
	}{
		{"the example's lines trimmed and joined", g.Raw("<p><b>a b</b></p>"), 0},
		{"the example's lines as written", g.Raw("<p>\n  <b>a b</b>\t\n</p>\n"), 1},
		{"markup the example does not have", g.Raw("<p><b>a b</b></p>x"), 1},
		{"markup as long as the example's, not it", g.Raw("<p><b>a c</b></p>"), 1},
	} {
		var r recorder
		example.Check(&r, c.node, path)
		if r.failures != c.failures {
			t.Errorf("Check failed %d times on %s, want %d", r.failures, c.name, c.failures)
		}
	}
}

// @trap: an example of the test's own, by its path from components/: indented.html's
// 17 bytes reach no difference with 40 bytes before it
func exampleOf(t *testing.T, markup string) string {
	components, err := filepath.Abs("../../../components")
	if err != nil {
		t.Fatal(err)
	}
	file := filepath.Join(t.TempDir(), "example.html")
	if err := os.WriteFile(file, []byte(markup), 0o644); err != nil {
		t.Fatal(err)
	}
	path, err := filepath.Rel(components, file)
	if err != nil {
		t.Fatal(err)
	}
	return path
}

func TestCheckNamesTheExampleBothMarkupsAndWhereTheyFirstDiffer(t *testing.T) {
	const short = "<p><b>a b</b></p>"
	digits := strings.Repeat("0123456789", 10)
	for _, c := range []struct {
		name, example, got string
		at                 int
		gotAround          string
		exampleAround      string
	}{
		{"a byte of another value", short, "<p><b>a c</b></p>", 8, "<p><b>a c</b></p>", short},
		{"the first byte", short, "a", 0, "a", short},
		{"markup that ends where the example goes on", short, "<p><b>a b</b>", 13, "<p><b>a b</b>", short},
		{"long markup, 40 bytes each side of the difference", digits, digits[:50] + "x" + digits[51:], 50, digits[10:50] + "x" + digits[51:90], digits[10:90]},
	} {
		path := exampleOf(t, c.example)
		want := fmt.Sprintf("%s\n got: %s\nwant: %s\nfirst differs at byte %d:\n got: %q\nwant: %q", path, c.got, c.example, c.at, c.gotAround, c.exampleAround)
		var r recorder
		example.Check(&r, g.Raw(c.got), path)
		if r.message != want {
			t.Errorf("%s: Check's failure is\n%s\nwant\n%s", c.name, r.message, want)
		}
	}
}

func TestCheckPanicsOnAnExampleThatIsNotThere(t *testing.T) {
	var r recorder
	if got := panicOf(func() { example.Check(&r, g.Raw(""), "no-such/example.html") }); !strings.Contains(got, "no-such/example.html") {
		t.Errorf("Check of an example that is not there panicked with %q, want the missing file named", got)
	}
}

func TestRenderWritesQuotesAsTheExamplesDo(t *testing.T) {
	got := example.Render(h.P(h.Title(`it's "a"`), g.Text(`it's "a"`)))
	if want := `<p title="it's &quot;a&quot;">it's &quot;a&quot;</p>`; got != want {
		t.Errorf("Render wrote %s, want %s", got, want)
	}
}

func TestRenderPanicsWithEachProblemCheckMarkupFinds(t *testing.T) {
	for html, count := range map[string]int{
		`<p class="cui-no-such"></p>`:                             1,
		`<p class="cui-no-such"></p><p class="cui-nor-this"></p>`: 2,
	} {
		problems := ciderui.CheckMarkup(html)
		if len(problems) != count {
			t.Fatalf("CheckMarkup found %d problems in %s, want %d: %q", len(problems), html, count, problems)
		}
		want := "example.Render: " + strings.Join(problems, "\n")
		if got := panicOf(func() { example.Render(g.Raw(html)) }); got != want {
			t.Errorf("Render of %s panicked with %q, want %q", html, got, want)
		}
	}
}

type failing struct{}

func (failing) Render(io.Writer) error { return errors.New("disk full") }

func TestRenderPanicsWithTheNodesError(t *testing.T) {
	if got := panicOf(func() { example.Render(failing{}) }); got != "disk full" {
		t.Errorf("Render of a node that cannot write panicked with %q, want its error", got)
	}
}

func panicOf(f func()) (message string) {
	defer func() {
		if r := recover(); r != nil {
			message = fmt.Sprint(r)
		}
	}()
	f()
	return ""
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
	panics := example.Break[part]{Want: "part needs a name", Change: blank}
	stays := example.Break[part]{Want: "part needs a name", Change: func(*part) {}}
	for _, c := range []struct {
		name     string
		valid    func() part
		breaks   []example.Break[part]
		failures int
	}{
		{"a break that panics with its want", valid, []example.Break[part]{panics}, 0},
		{"a break that panics with another message", valid, []example.Break[part]{{Want: "part needs", Change: blank}}, 1},
		{"a second break that does not panic", valid, []example.Break[part]{panics, stays}, 1},
		{"a valid part that panics", func() part { return part{name: ""} }, []example.Break[part]{panics}, 1},
	} {
		var r recorder
		example.Panics(&r, c.valid, c.breaks)
		if r.failures != c.failures {
			t.Errorf("%s: Panics failed %d times, want %d", c.name, r.failures, c.failures)
		}
	}
}

// @invariant: named as p.go's Classed, since PagesClass holds a row by its Node's type name
type Classed struct{ Class string }

func (c Classed) Render(w io.Writer) error {
	return h.P(g.Attr("class", "root large "+c.Class)).Render(w)
}

func TestPagesClassFailsOnAPartWithoutThePagesClassAfterItsRootOrARow(t *testing.T) {
	dir := t.TempDir()
	t.Chdir(dir)
	for name, source := range map[string]string{
		"p.go":      "package p\n\ntype Classed struct{ Class string }\n\ntype hidden struct{ Class string }\n\ntype Plain struct{ Name string }\n",
		"p_test.go": "package p\n\ntype TestOnly struct{ Class string }\n",
	} {
		if err := os.WriteFile(filepath.Join(dir, name), []byte(source), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	writes := example.ClassRow{Root: "root large", Node: Classed{Class: example.PageClass}}
	row := func(class string) []example.ClassRow {
		return []example.ClassRow{{Root: writes.Root, Node: Classed{Class: class}}}
	}
	for _, c := range []struct {
		name     string
		rows     []example.ClassRow
		failures int
	}{
		{"each exported struct of the package a row's part", []example.ClassRow{writes}, 0},
		{"a part that writes a class of its own where the page's goes", row("small"), 1},
		{"a part that writes a class of its own after the page's", row(example.PageClass + " small"), 1},
		{"a part that writes the page's class after more classes than its Root", []example.ClassRow{{Root: "large", Node: writes.Node}}, 1},
		{"a struct with a Class and no row", nil, 1},
	} {
		var r recorder
		example.PagesClass(&r, c.rows)
		if r.failures != c.failures {
			t.Errorf("%s: PagesClass failed %d times, want %d", c.name, r.failures, c.failures)
		}
	}
}
