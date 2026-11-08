// Package example holds a component's output to its example
// (components/<set>/<name>/examples/<id>.html), the markup the kit's tests
// measure, so the two cannot part.
package example

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"runtime"
	"slices"
	"strings"
	"testing"

	"github.com/at-least/ciderui/go/internal/document"
	g "maragu.dev/gomponents"
)

var root = func() string {
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		panic("no caller")
	}
	return filepath.Join(filepath.Dir(file), "../../..")
}()

// Read is the example at path under components/, its lines trimmed and
// joined: the markup as a component writes it.
func Read(path string) string {
	b, err := os.ReadFile(filepath.Join(root, "components", path))
	if err != nil {
		panic(err)
	}
	return JoinTrimmed(string(b))
}

func JoinTrimmed(text string) string {
	var out strings.Builder
	for _, line := range strings.Split(text, "\n") {
		out.WriteString(strings.TrimSpace(line))
	}
	return out.String()
}

// Render is node's HTML, with gomponents' two escapes the examples write
// otherwise put as they do: ' as itself (gomponents writes &#39;), " as
// &quot; (&#34;). Each pair means the same in text and in a double-quoted
// attribute. It panics on what ciderui.CheckMarkup refuses: a branch no example
// draws is held to the stylesheets too.
func Render(node g.Node) string {
	var b strings.Builder
	if err := node.Render(&b); err != nil {
		panic(err)
	}
	html := strings.NewReplacer("&#39;", "'", "&#34;", "&quot;").Replace(b.String())
	document.Check("example.Render", html)
	return html
}

// Check fails t where node does not write the example at path.
func Check(t testing.TB, node g.Node, path string) {
	t.Helper()
	if got, want := Render(node), Read(path); got != want {
		t.Errorf("%s\n got: %s\nwant: %s", path, got, want)
	}
}

// Break is one thing wrong with a part, and the panic it draws.
type Break[T g.Node] struct {
	Want   string
	Change func(*T)
}

// Panics fails t unless valid() renders, and each Break of a new valid() panics
// with exactly its Want: valid is a function since a Change may write through a
// slice, and a panic from anything but its Change has another message.
func Panics[T g.Node](t testing.TB, valid func() T, breaks []Break[T]) {
	t.Helper()
	if got := panicOf(valid()); got != "" {
		t.Errorf("the valid part panicked with %q", got)
	}
	for i, b := range breaks {
		node := valid()
		b.Change(&node)
		if got := panicOf(node); got != b.Want {
			t.Errorf("break %d: panicked with %q, want %q", i, got, b.Want)
		}
	}
}

func panicOf(node g.Node) (message string) {
	defer func() {
		if r := recover(); r != nil {
			message = fmt.Sprint(r)
		}
	}()
	Render(node)
	return ""
}

// PageClass is the Class a ClassRow's Node is given.
const PageClass = "page-class"

// ClassRow is a part given PageClass, and the classes of its own before it.
type ClassRow struct {
	Root string
	Node g.Node
}

// PagesClass fails t where a row's Node does not write PageClass after its Root,
// and for each exported struct of the package in dir with a Class field that no
// row's Node is and nested does not name. A part that drops its Class draws fine,
// in the wrong place: this stops a new part going without a row. nested: a part
// only inside another part (that part's row), or one whose Class is its whole
// class attribute.
func PagesClass(t testing.TB, dir string, nested []string, rows []ClassRow) {
	t.Helper()
	held := slices.Clone(nested)
	for _, row := range rows {
		_, name, _ := strings.Cut(fmt.Sprintf("%T", row.Node), ".")
		held = append(held, name)
		want := `class="` + row.Root + " " + PageClass + `"`
		if got := Render(row.Node); !strings.Contains(got, want) {
			t.Errorf("%s: no %s in %s", name, want, got)
		}
	}
	for _, name := range structsWithClass(dir) {
		if !slices.Contains(held, name) {
			t.Errorf("%s has a Class and no ClassRow", name)
		}
	}
}

func structsWithClass(dir string) []string {
	paths, err := filepath.Glob(filepath.Join(dir, "*.go"))
	if err != nil {
		panic(err)
	}
	var names []string
	for _, path := range paths {
		if strings.HasSuffix(path, "_test.go") {
			continue
		}
		file, err := parser.ParseFile(token.NewFileSet(), path, nil, parser.SkipObjectResolution)
		if err != nil {
			panic(err)
		}
		ast.Inspect(file, func(n ast.Node) bool {
			spec, ok := n.(*ast.TypeSpec)
			if !ok || !spec.Name.IsExported() {
				return true
			}
			if fields, ok := spec.Type.(*ast.StructType); ok {
				for _, field := range fields.Fields.List {
					if slices.ContainsFunc(field.Names, func(f *ast.Ident) bool { return f.Name == "Class" }) {
						names = append(names, spec.Name.Name)
					}
				}
			}
			return true
		})
	}
	return names
}
