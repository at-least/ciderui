// Command gate is the repository's gate: the checks that open no browser, on this
// machine, then the browser tests in the test image (scripts/docker).
//
//	go run ./scripts/gate
//
// @invariant: the browser half only once the other passes (the owner's decision: a
// failure a fast check finds wastes no browser run).
package main

import (
	"bufio"
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/at-least/ciderui/internal/build"
)

func main() {
	if len(os.Args) != 1 {
		fmt.Fprintln(os.Stderr, "usage: go run ./scripts/gate")
		os.Exit(2)
	}
	if err := host(); err != nil {
		fmt.Println("gate:", err)
		os.Exit(1)
	}
	fmt.Println("gate: dist written; lucide.go, gofumpt, vet, errcheck, ciderui check and tests clean")
	os.Exit(browser())
}

func run(name string, args ...string) (string, error) {
	cmd := exec.Command(name, args...)
	cmd.Dir = build.Root
	var stdout, both bytes.Buffer
	cmd.Stdout, cmd.Stderr = io.MultiWriter(&stdout, &both), &both
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("%s %s failed (%w)\n%s", name, strings.Join(args, " "), err, both.String())
	}
	return stdout.String(), nil
}

// @trap: a part's button class is a compile error, not a string: btn.go alone writes it.
// play_button.go's two glyphs and labels are no <button> the type can hold.
var (
	ownButton    = regexp.MustCompile(`(^|/)(btn|play_button)\.go$`)
	buttonClass  = regexp.MustCompile(`"cui-(console-)?btn[" ]`)
	statementsOf = []string{"go/site", "go/console"}
)

func host() error {
	// @invariant: dist/ is written, not checked (the owner's decision): a CSS edit is
	// one gate run, and the commit carries the dist/ the gate built
	for _, step := range [][]string{
		{"go", "run", "./scripts/build-lucide", "--check"},
		{"go", "run", "./scripts/build-css"},
	} {
		if _, err := run(step[0], step[1:]...); err != nil {
			return err
		}
	}
	// @trap: the module's Go files, -tags browser with them, not `.`: gofumpt walks a
	// directory whole, and the browser tests' files are no package's without the tag
	each := func(files string) string { return "{{range ." + files + "}}{{$.Dir}}/{{.}}{{\"\\n\"}}{{end}}" }
	listed, err := run("go", "list", "-tags", "browser", "-f", each("GoFiles")+each("TestGoFiles")+each("XTestGoFiles"), "./...")
	if err != nil {
		return err
	}
	written, err := handWritten(strings.Fields(listed))
	if err != nil {
		return err
	}
	unformatted, err := run("go", append([]string{"tool", "gofumpt", "-l"}, written...)...)
	if err != nil {
		return err
	}
	if strings.TrimSpace(unformatted) != "" {
		return fmt.Errorf("not gofumpt'd (go tool gofumpt -w <file>):\n%s", unformatted)
	}
	if _, err := run("go", "vet", "-tags", "browser", "./..."); err != nil {
		return err
	}
	if _, err := run("go", "tool", "errcheck", "-blank", "-tags", "browser", "-exclude", "errcheck_excludes.txt", "./..."); err != nil {
		return err
	}
	// @invariant: a component's field left out of a struct literal is a default nobody chose
	if _, err := run("go", "run", "./go/cmd/ciderui", "check", "./..."); err != nil {
		return err
	}
	for _, dir := range statementsOf {
		entries, err := os.ReadDir(filepath.Join(build.Root, dir))
		if err != nil {
			return err
		}
		for _, e := range entries {
			file := dir + "/" + e.Name()
			if !strings.HasSuffix(file, ".go") || strings.HasSuffix(file, "_test.go") || ownButton.MatchString(file) {
				continue
			}
			b, err := os.ReadFile(filepath.Join(build.Root, file))
			if err != nil {
				return err
			}
			if buttonClass.Match(b) {
				return fmt.Errorf("%s writes a button's class: use its Btn type", file)
			}
		}
	}
	// @invariant: a part's every branch has a test: a new panic with no Break, or a new
	// branch no example draws, stops here
	tested, err := run("go", "test", "-count=1", "-cover", "./...")
	if err != nil {
		return err
	}
	for _, set := range statementsOf {
		var line string
		for l := range strings.SplitSeq(tested, "\n") {
			if f := strings.Split(l, "\t"); len(f) > 1 && f[1] == "github.com/at-least/ciderui/"+set {
				line = l
			}
		}
		if !strings.Contains(line, "coverage: 100.0% of statements") {
			return fmt.Errorf("a statement of %s no test runs (go test -coverprofile): %s", set, line)
		}
	}
	return nil
}

// @external: gofumpt applies its rules to a file named on its command line, generated
// or not (it skips a generated file only when it walks a directory), and the generators
// write gofmt's form: handWritten leaves the generated files out
func handWritten(files []string) ([]string, error) {
	var written []string
	for _, f := range files {
		parsed, err := parser.ParseFile(token.NewFileSet(), f, nil, parser.PackageClauseOnly|parser.ParseComments)
		if err != nil {
			return nil, err
		}
		if !ast.IsGenerated(parsed) {
			written = append(written, f)
		}
	}
	return written, nil
}

// @invariant: a passed package's line is dropped and a failed one's output printed whole:
// the last line names every failed package, the one a reader of the tail sees
func browser() int {
	started := time.Now()
	cmd := exec.Command("go", "run", "./scripts/docker")
	cmd.Dir = build.Root
	cmd.Stderr = os.Stderr
	pipe, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("gate:", err)
		return 1
	}
	if err := cmd.Start(); err != nil {
		fmt.Println("gate:", err)
		return 1
	}
	passed, failed := 0, []string{}
	lines := bufio.NewScanner(pipe)
	lines.Buffer(make([]byte, 1<<20), 1<<24)
	for lines.Scan() {
		line := lines.Text()
		switch {
		case strings.HasPrefix(line, "ok  \t"):
			passed++
			continue
		case strings.HasPrefix(line, "?   \t"):
			continue
		case strings.HasPrefix(line, "FAIL\t"):
			failed = append(failed, strings.Fields(line)[1])
		}
		fmt.Println(line)
	}
	if err := lines.Err(); err != nil {
		fmt.Println("gate:", err)
		if _, err := io.Copy(io.Discard, pipe); err != nil {
			fmt.Println("gate:", err)
		}
		if err := cmd.Wait(); err != nil {
			fmt.Println("gate:", err)
		}
		return 1
	}
	err = cmd.Wait()
	seconds := int(time.Since(started).Seconds())
	if err != nil || len(failed) > 0 {
		fmt.Printf("gate: %d browser packages passed, failed: %s, in %ds\n", passed, strings.Join(failed, " "), seconds)
		return 1
	}
	fmt.Printf("gate: %d browser packages passed in %ds\n", passed, seconds)
	return 0
}
