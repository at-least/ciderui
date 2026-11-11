package ciderui

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path"
	"regexp"
	"strings"
	"testing"
)

var decisionTag = regexp.MustCompile(`^@(invariant|trap|external): `)

var cssComment = regexp.MustCompile(`(?s)/\*.*?\*/`)

// TestCommentsCarryDecisionTags holds every comment in the repository to
// DESIGN.md's Comments section. Excepted: go/lucide/lucide.go and
// scripts/build-lucide.mjs, which writes that file's per-icon godoc as text of
// its own, and dist/, the build's output.
func TestCommentsCarryDecisionTags(t *testing.T) {
	for _, file := range trackedFiles(t) {
		if file == "go/lucide/lucide.go" || file == "scripts/build-lucide.mjs" || strings.HasPrefix(file, "dist/") {
			continue
		}
		switch {
		case path.Ext(file) == ".go":
			goCommentTags(t, file)
		case path.Ext(file) == ".css":
			text := readOrFatal(t, file)
			for _, comment := range cssComment.FindAllString(text, -1) {
				first := strings.TrimSpace(strings.TrimPrefix(comment, "/*"))
				if !decisionTag.MatchString(first) {
					t.Errorf("%s: %q carries no @invariant/@trap/@external tag", file, first)
				}
			}
		case path.Ext(file) == ".mjs" || path.Ext(file) == ".js":
			report(t, lineCommentTags(file, readOrFatal(t, file), "//"))
		case strings.HasSuffix(file, "Dockerfile"):
			report(t, lineCommentTags(file, readOrFatal(t, file), "#"))
		}
	}
}

func report(t *testing.T, problems []string) {
	t.Helper()
	for _, problem := range problems {
		t.Error(problem)
	}
}

func readOrFatal(t *testing.T, file string) string {
	t.Helper()
	raw, err := os.ReadFile(file)
	if err != nil {
		t.Fatal(err)
	}
	return string(raw)
}

// @invariant: a doc above an unexported declaration is a decision and carries
// its tag itself
func goCommentTags(t *testing.T, name string) {
	t.Helper()
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, name, nil, parser.ParseComments)
	if err != nil {
		t.Fatal(err)
	}
	free := map[*ast.CommentGroup]bool{file.Doc: true}
	names := func(spec ast.Spec) (out []*ast.Ident) {
		switch s := spec.(type) {
		case *ast.TypeSpec:
			out = append(out, s.Name)
		case *ast.ValueSpec:
			out = append(out, s.Names...)
		}
		return out
	}
	isExported := func(idents []*ast.Ident) bool {
		for _, name := range idents {
			if name.IsExported() {
				return true
			}
		}
		return false
	}
	for _, decl := range file.Decls {
		d, ok := decl.(*ast.GenDecl)
		if !ok || d.Doc == nil {
			continue
		}
		var idents []*ast.Ident
		for _, spec := range d.Specs {
			idents = append(idents, names(spec)...)
		}
		if isExported(idents) {
			free[d.Doc] = true
		}
	}
	ast.Inspect(file, func(n ast.Node) bool {
		switch s := n.(type) {
		case *ast.FuncDecl:
			if s.Doc != nil && s.Name.IsExported() {
				free[s.Doc] = true
			}
		case *ast.TypeSpec:
			if s.Doc != nil && s.Name.IsExported() {
				free[s.Doc] = true
			}
		case *ast.ValueSpec:
			if s.Doc != nil && isExported(s.Names) {
				free[s.Doc] = true
			}
		case *ast.Field:
			if s.Doc != nil && isExported(s.Names) {
				free[s.Doc] = true
			}
		}
		return true
	})
	for _, group := range file.Comments {
		if free[group] {
			continue
		}
		first := strings.TrimSpace(group.List[0].Text[2:])
		// @external: an Example's Output comment is the test framework's
		// assertion, not a comment
		if strings.HasPrefix(first, "Output:") {
			continue
		}
		if !decisionTag.MatchString(first) {
			t.Errorf("%s: %q carries no @invariant/@trap/@external tag", fset.Position(group.Pos()), first)
		}
	}
}

// @invariant: a run of whole-line comments is one comment; its first line
// names its kind, the lines after it say the rest. A line inside a template
// literal or a block comment is payload (page JavaScript a probe writes as
// text), not a comment; a file whose template or block comment the scanner
// leaves open is misread, and fails instead.
func lineCommentTags(name, text, mark string) []string {
	var problems []string
	fail := func(format string, args ...any) {
		problems = append(problems, fmt.Sprintf(name+": "+format, args...))
	}
	s, inComment, tagged := scan{}, false, false
	for i, line := range strings.Split(text, "\n") {
		wasPayload := s.inTemplate || s.inBlockComment
		s = s.line(line)
		if wasPayload {
			continue
		}
		trimmed := strings.TrimSpace(line)
		if strings.HasPrefix(trimmed, mark) {
			if !inComment {
				inComment, tagged = true, decisionTag.MatchString(strings.TrimSpace(strings.TrimPrefix(trimmed, mark)))
			}
			continue
		}
		if trimmed == "" {
			continue
		}
		if inComment && !tagged {
			fail("line %d: a comment carries no @invariant/@trap/@external tag", i+1)
		}
		inComment = false
	}
	if s.inTemplate || s.inBlockComment {
		fail("a template literal or block comment left open, or backticks the scanner misread")
	}
	if inComment && !tagged {
		fail("a comment at the end carries no @invariant/@trap/@external tag")
	}
	return problems
}

// @external: only a backtick outside a quoted string, a regular expression
// and a comment opens or closes a template literal: test/probes.mjs names a
// path with three of them inside one regular expression, and a comment's
// paired ones name code. A slash opens a regular expression only where an
// expression can start, not after a value (division).
type scan struct{ inTemplate, inBlockComment bool }

func (s scan) line(text string) scan {
	const (
		code = iota
		single
		double
		template
		block
		regular
	)
	state := code
	switch {
	case s.inTemplate:
		state = template
	case s.inBlockComment:
		state = block
	}
	last := byte(' ')
	for i := 0; i < len(text); i++ {
		c := text[i]
		switch state {
		case block:
			if c == '*' && i+1 < len(text) && text[i+1] == '/' {
				state = code
				i++
			}
		case single:
			if c == '\\' {
				i++
			} else if c == '\'' {
				state = code
			}
		case double:
			if c == '\\' {
				i++
			} else if c == '"' {
				state = code
			}
		case template:
			if c == '\\' {
				i++
			} else if c == '`' {
				state = code
			}
		case regular:
			if c == '\\' {
				i++
			} else if c == '/' {
				state = code
				for i+1 < len(text) && text[i+1] >= 'a' && text[i+1] <= 'z' {
					i++
				}
			}
		case code:
			switch {
			case c == '\'':
				state = single
			case c == '"':
				state = double
			case c == '`':
				state = template
			case c == '/' && i+1 < len(text) && text[i+1] == '/':
				return scan{inTemplate: false, inBlockComment: false}
			case c == '/' && i+1 < len(text) && text[i+1] == '*':
				state = block
			case c == '/' && !(last == ')' || last == ']' || last == '}' || last == '.' ||
				last == '\'' || last == '"' || last == '`' || last >= '0' && last <= '9' ||
				last >= 'A' && last <= 'Z' || last == '_' || last >= 'a' && last <= 'z'):
				state = regular
			}
			if state == code && c != ' ' && c != '\t' {
				last = c
			}
		}
	}
	return scan{inTemplate: state == template, inBlockComment: state == block}
}

func TestLineCommentTagsFail(t *testing.T) {
	if got := lineCommentTags("a.mjs", "// @trap: a rejected way\nconst x = 1;\n", "//"); got != nil {
		t.Errorf("a tagged comment: %q", got)
	}
	if got := lineCommentTags("a.mjs", "// plain\nconst x = 1;\n", "//"); len(got) != 1 {
		t.Errorf("an untagged comment: %q", got)
	}
	payload := "const page = `<script>\n// the page's own comment, payload\n</script>`;\n// @external: after it\n"
	if got := lineCommentTags("probe.mjs", payload, "//"); got != nil {
		t.Errorf("a template literal's comment lines: %q", got)
	}
	if got := lineCommentTags("a.mjs", "// plain\nconst s = `open\n", "//"); len(got) != 2 {
		t.Errorf("an open template literal: %q", got)
	}
	if got := lineCommentTags("Dockerfile", "# @external: the image\nRUN x\n", "#"); got != nil {
		t.Errorf("a Dockerfile's tagged comment: %q", got)
	}
	regex := "// @invariant: before it\nconst paths = [...html.matchAll(/['`](\\.{1,2}\\/[^'`\\n]*)['`]/g)];\n"
	if got := lineCommentTags("probes.mjs", regex, "//"); got != nil {
		t.Errorf("backticks inside a regular expression: %q", got)
	}
	interpolated := "const run = `(async () => {\n  const w = ${frame('tabs')};\n  // payload\n})();`;\n"
	if got := lineCommentTags("run.mjs", interpolated, "//"); got != nil {
		t.Errorf("an interpolation inside a template literal: %q", got)
	}
	block := "/* a block comment with a ` inside */\n// @trap: after it\n"
	if got := lineCommentTags("a.mjs", block, "//"); got != nil {
		t.Errorf("a block comment on one line: %q", got)
	}
	division := "const half = width / 2 / 1;\n// @external: after it\n"
	if got := lineCommentTags("a.mjs", division, "//"); got != nil {
		t.Errorf("division read as a regular expression: %q", got)
	}
}
