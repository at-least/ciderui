package ciderui_test

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"reflect"
	"regexp"
	"strings"
	"testing"

	"github.com/at-least/ciderui/go/console"
	"github.com/at-least/ciderui/go/site"
)

// @invariant: README's Go example is a literal a call could write: every field
// of every kit type it names is set and none the type lacks, or a
// reader copies markup that does not compile
func TestReadmesGoExampleCompiles(t *testing.T) {
	types := map[string]any{
		"console.ConsolePage":    console.ConsolePage{Lang: "", Head: nil, Scheme: console.ConsoleSchemeSystem, Content: console.ConsoleContent{Sidebar: nil, Back: nil, Title: "", Detail: "", Actions: nil, Children: nil}, Overlays: nil},
		"console.ConsoleContent": console.ConsoleContent{Sidebar: nil, Back: nil, Title: "", Detail: "", Actions: nil, Children: nil},
		"site.Page":              site.Page{Lang: "", Title: "", Head: nil, Children: nil},
	}
	raw, err := os.ReadFile("README.md")
	if err != nil {
		t.Fatal(err)
	}
	code := string(raw)
	start := strings.Index(code, "```go\n")
	end := strings.Index(code[start+6:], "```")
	if start < 0 || end < 0 {
		t.Fatal("README.md holds no Go example")
	}
	code = code[start+6 : start+6+end]
	file, err := parser.ParseFile(token.NewFileSet(), "readme.go", "package p\nfunc _() {\n"+code+"\n}", 0)
	if err != nil {
		t.Fatalf("README's Go example does not parse: %v", err)
	}
	ast.Inspect(file, func(n ast.Node) bool {
		lit, ok := n.(*ast.CompositeLit)
		if !ok {
			return true
		}
		selector, ok := lit.Type.(*ast.SelectorExpr)
		if !ok {
			return true
		}
		name := selector.X.(*ast.Ident).Name + "." + selector.Sel.Name
		strct, known := types[name]
		if !known {
			t.Errorf("README's example writes %s, which this test holds no type for: add it", name)
			return true
		}
		fields := map[string]bool{}
		typ := reflect.TypeOf(strct)
		for i := range typ.NumField() {
			fields[typ.Field(i).Name] = true
		}
		written := map[string]bool{}
		for _, elt := range lit.Elts {
			pair, ok := elt.(*ast.KeyValueExpr)
			if !ok {
				t.Errorf("%s in README's example writes no field names", name)
				continue
			}
			field := pair.Key.(*ast.Ident).Name
			if !fields[field] {
				t.Errorf("README's example sets %s.%s, which the type does not hold", name, field)
			}
			written[field] = true
		}
		for field := range fields {
			if !written[field] {
				t.Errorf("README's example leaves %s.%s out", name, field)
			}
		}
		return true
	})
}

var (
	readmeTool    = regexp.MustCompile(`(?m)^go get -tool github\.com/at-least/ciderui/(\S+)$`)
	readmeCommand = regexp.MustCompile(`(?m)^go tool ciderui (\w+) `)
	usageVerb     = regexp.MustCompile(`(?m)^\s*(?:const usage = .)?(?:usage: )?go tool ciderui (\w+) `)
)

// @invariant: README's commands are the kit command's: the one `tool` line names
// go/cmd/ciderui, and each `go tool ciderui <verb>` is a verb its usage lists, or a
// reader copies a command that prints its usage
func TestReadmesCommandsAreTheKitCommands(t *testing.T) {
	readme, err := os.ReadFile("README.md")
	if err != nil {
		t.Fatal(err)
	}
	tools := readmeTool.FindAllSubmatch(readme, -1)
	if len(tools) != 1 || string(tools[0][1]) != "go/cmd/ciderui" {
		t.Errorf("README.md's go get -tool lines name %q, want go/cmd/ciderui alone", tools)
	}
	main, err := os.ReadFile("go/cmd/ciderui/main.go")
	if err != nil {
		t.Fatal(err)
	}
	verbs := map[string]bool{}
	for _, m := range usageVerb.FindAllSubmatch(main, -1) {
		verbs[string(m[1])] = true
	}
	if !verbs["check"] || !verbs["css"] {
		t.Fatalf("go/cmd/ciderui's usage lists %v, want check and css: this test reads it wrong", verbs)
	}
	written := map[string]bool{}
	for _, m := range readmeCommand.FindAllSubmatch(readme, -1) {
		written[string(m[1])] = true
		if !verbs[string(m[1])] {
			t.Errorf("README.md runs go tool ciderui %s, which go/cmd/ciderui's usage does not list", m[1])
		}
	}
	for verb := range verbs {
		if !written[verb] {
			t.Errorf("README.md never runs go tool ciderui %s", verb)
		}
	}
}
