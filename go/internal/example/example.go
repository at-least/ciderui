// Package example holds a component's output to its example
// (components/<set>/<name>/examples/<id>.html), the markup the kit's tests
// measure, so the two cannot part.
package example

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"github.com/at-least/ciderui"
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
	return joinTrimmed(string(b))
}

func joinTrimmed(text string) string {
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
	if problems := ciderui.CheckMarkup(html); problems != nil {
		panic(fmt.Sprintf("%q", problems))
	}
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
	for _, b := range breaks {
		node := valid()
		b.Change(&node)
		if got := panicOf(node); got != b.Want {
			t.Errorf("panicked with %q, want %q", got, b.Want)
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
