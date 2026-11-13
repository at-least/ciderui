package ciderui_test

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"reflect"
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
		"console.ConsolePage": console.ConsolePage{Lang: "", Title: "", Head: nil, Scheme: console.ConsoleSchemeSystem, Children: nil},
		"site.Page":           site.Page{Lang: "", Title: "", Head: nil, Children: nil},
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
