package ciderui

import (
	"maps"
	"os"
	"os/exec"
	"path"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"testing"
)

const pathExtension = `\.(?:md|css|go|mjs|js|json|html|swift|yml)`

var (
	backtickedPath = regexp.MustCompile("`([\\w./@-]+(?:" + pathExtension + `|/))` + "`")
	repositoryPath = regexp.MustCompile(`(?:^|[^\w./-])((?:components|reference|scripts|test|go)/[\w./-]*` + pathExtension + `)\b`)
	designCite     = regexp.MustCompile(`\b([Dd]ecision|[Rr]ule)s? (\d{1,2}(?:(?:, | and )\d{1,2}\b)*)`)
	namedClass     = regexp.MustCompile(`(?:^|[^\w/-])(cui-[a-z][\w-]*)(\*?)`)
	citedNumber    = regexp.MustCompile(`\d+`)
	numberedItem   = regexp.MustCompile(`(?m)^(\d+)\. \*\*(Withdrawn)?`)
)

func number(digits string) int {
	n, err := strconv.Atoi(digits)
	if err != nil {
		panic(err)
	}
	return n
}

// @trap: not a walk of the disk: .claude/worktrees holds whole checkouts of older commits
func trackedFiles(t *testing.T) []string {
	t.Helper()
	out, err := exec.Command("git", "ls-files").Output()
	if err != nil {
		t.Fatal(err)
	}
	return strings.Split(strings.TrimSpace(string(out)), "\n")
}

// @invariant: a path in a document's backticks is a claim that it exists: a
// removed file's name is written without them.
//
// stalePaths returns the paths text names that exist neither beside file nor
// from the repository's top, nor, a bare name, anywhere in names: a document
// names a file of the part it is about by its name alone.
func stalePaths(file, text string, names, exists map[string]bool) []string {
	var stale []string
	check := func(p string) {
		if strings.HasPrefix(p, "/") || strings.Contains(p, "@") {
			return
		}
		clean := strings.TrimSuffix(p, "/")
		if exists[path.Join(path.Dir(file), clean)] || exists[path.Clean(clean)] {
			return
		}
		if !strings.Contains(clean, "/") && names[clean] {
			return
		}
		if !slices.Contains(stale, p) {
			stale = append(stale, p)
		}
	}
	if strings.HasSuffix(file, ".md") {
		for _, m := range backtickedPath.FindAllStringSubmatch(text, -1) {
			check(m[1])
		}
	}
	for _, m := range repositoryPath.FindAllStringSubmatch(text, -1) {
		check(m[1])
	}
	return stale
}

// @invariant: staleCites returns each cite of a decision or rule DESIGN.md does not hold,
// or of a withdrawn decision on a line that does not say it is withdrawn.
func staleCites(text string, decisions, rules int, withdrawn []int) []string {
	var stale []string
	for line := range strings.Lines(text) {
		saysWithdrawn := strings.Contains(strings.ToLower(line), "withdrawn")
		for _, m := range designCite.FindAllStringSubmatch(line, -1) {
			isDecision := strings.EqualFold(m[1], "decision")
			limit := rules
			if isDecision {
				limit = decisions
			}
			for _, digits := range citedNumber.FindAllString(m[2], -1) {
				n := number(digits)
				gone := isDecision && slices.Contains(withdrawn, n) && !saysWithdrawn
				if n < 1 || n > limit || gone {
					stale = append(stale, m[0])
				}
			}
		}
	}
	return stale
}

// @invariant: staleClasses returns each cui-* class text names that classes lacks; `cui-x-*`
// names a family, which at least one class must begin. A line that says it is
// withdrawn is about what is gone.
func staleClasses(text string, classes map[string]bool) []string {
	var stale []string
	for line := range strings.Lines(text) {
		if strings.Contains(strings.ToLower(line), "withdrawn") {
			continue
		}
		for _, m := range namedClass.FindAllStringSubmatch(line, -1) {
			held := classes[m[1]]
			if m[2] != "" {
				held = slices.ContainsFunc(slices.Collect(maps.Keys(classes)), func(class string) bool { return strings.HasPrefix(class, m[1]) })
			}
			if !held {
				stale = append(stale, m[1]+m[2])
			}
		}
	}
	return stale
}

func TestStalePathsAndCitesFail(t *testing.T) {
	exists := map[string]bool{"a/here.css": true, "here.css": true}
	names := map[string]bool{"elsewhere.html": true}
	got := stalePaths("a/doc.md", "`here.css` `elsewhere.html` `/v/url.css` `pkg@1/x.js` `gone.css` `a/gone/` test/gone.mjs", names, exists)
	if want := []string{"gone.css", "a/gone/", "test/gone.mjs"}; !slices.Equal(got, want) {
		t.Errorf("stalePaths: %q, want %q", got, want)
	}
	if got := stalePaths("a/code.go", "`gone.css`", names, exists); got != nil {
		t.Errorf("stalePaths outside a document reads a backticked name: %q", got)
	}
	got = staleCites("decision 3\nrule 2\ndecision 25, 2026-11-06\nrule 15\nsee decision 7\ndecisions 7 and 8 withdrawn\n", 25, 14, []int{7})
	if want := []string{"rule 15", "decision 7"}; !slices.Equal(got, want) {
		t.Errorf("staleCites: %q, want %q", got, want)
	}
	got = staleClasses("`cui-btn` `cui-btn-*` `--cui-gone` `cui-tile` `cui-docs-*`\n`cui-gone` withdrawn\n", map[string]bool{"cui-btn": true, "cui-btn-icon": true})
	if want := []string{"cui-tile", "cui-docs-*"}; !slices.Equal(got, want) {
		t.Errorf("staleClasses: %q, want %q", got, want)
	}
}

// the root's alone is written for people; a second would read as another
func TestOneReadme(t *testing.T) {
	for _, file := range trackedFiles(t) {
		if path.Base(file) == "README.md" && file != "README.md" {
			t.Errorf("%s: an agent's document takes another name", file)
		}
	}
}

func TestPathsAndCitesResolve(t *testing.T) {
	raw, err := os.ReadFile("DESIGN.md")
	if err != nil {
		t.Fatal(err)
	}
	design := string(raw)
	decisionsAt := strings.Index(design, "## Decisions")
	rulesAt := strings.Index(design, "## Rules")
	tokensAt := strings.Index(design, "## The token contract")
	if decisionsAt < 0 || rulesAt < 0 || tokensAt < 0 {
		t.Fatal("DESIGN.md lost a heading this test finds its numbers by")
	}
	items := numberedItem.FindAllStringSubmatch(design[decisionsAt:], -1)
	decisions := len(items)
	rules := len(numberedItem.FindAllString(design[rulesAt:tokensAt], -1))
	var withdrawn []int
	for _, m := range items {
		if m[2] != "" {
			withdrawn = append(withdrawn, number(m[1]))
		}
	}
	if decisions == 0 || rules == 0 || len(withdrawn) == 0 {
		t.Fatalf("read %d decisions, %d rules, %d withdrawn from DESIGN.md", decisions, rules, len(withdrawn))
	}

	files := trackedFiles(t)
	names := map[string]bool{}
	exists := map[string]bool{}
	for _, file := range files {
		names[path.Base(file)] = true
		for at := file; at != "."; at = path.Dir(at) {
			exists[at] = true
		}
	}
	classes := vocabularies()[""].classes
	for _, file := range files {
		switch {
		case file == "pointers_test.go", strings.HasPrefix(file, "dist/"), strings.Contains(file, "/results/"):
			continue
		case !slices.Contains([]string{".md", ".go", ".css", ".mjs", ".js"}, path.Ext(file)):
			continue
		}
		raw, err := os.ReadFile(file)
		if err != nil {
			t.Fatal(err)
		}
		text := string(raw)
		for _, p := range stalePaths(file, text, names, exists) {
			t.Errorf("%s names %s, which does not exist", file, p)
		}
		for _, cite := range staleCites(text, decisions, rules, withdrawn) {
			t.Errorf("%s cites %s: not in DESIGN.md, or withdrawn", file, cite)
		}
		if file == "DESIGN.md" || file == "README.md" {
			for _, class := range staleClasses(text, classes) {
				t.Errorf("%s names %s, which the kit does not hold", file, class)
			}
		}
	}
}
