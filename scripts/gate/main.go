// Command gate is the repository's gate: the checks that open no browser, on this
// machine, then the browser tests in the test image (scripts/docker).
//
//	go run ./scripts/gate
//
// @invariant: the browser half only once the other passes (the owner, 2026-11-14: a
// failure a fast check finds wastes no browser run). On success it prints one line a
// half; a failure prints its output and the step that failed.
package main

import (
	"bufio"
	"bytes"
	"fmt"
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
	fmt.Println("gate: lucide.go, dist, gofmt, vet, ciderui-check and tests clean")
	os.Exit(browser())
}

func run(name string, args ...string) (string, error) {
	cmd := exec.Command(name, args...)
	cmd.Dir = build.Root
	var out bytes.Buffer
	cmd.Stdout, cmd.Stderr = &out, &out
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("%s %s failed (%w)\n%s", name, strings.Join(args, " "), err, out.String())
	}
	return out.String(), nil
}

// @trap: a part's button class is a compile error, not a string: btn.go alone writes it.
// play_button.go's two glyphs and labels are no <button> the type can hold.
var (
	ownButton    = regexp.MustCompile(`(^|/)(btn|play_button)\.go$`)
	buttonClass  = regexp.MustCompile(`"cui-(console-)?btn[" ]`)
	statementsOf = []string{"go/site", "go/console"}
)

func host() error {
	for _, step := range [][]string{
		{"go", "run", "./scripts/build-lucide", "--check"},
		{"go", "run", "./scripts/build-css", "--check"},
	} {
		if _, err := run(step[0], step[1:]...); err != nil {
			return err
		}
	}
	// @trap: the module's Go files, -tags browser with them, not `.`: gofmt walks a
	// directory whole, and the browser tests' files are no package's without the tag
	each := func(files string) string { return "{{range ." + files + "}}{{$.Dir}}/{{.}}{{\"\\n\"}}{{end}}" }
	listed, err := run("go", "list", "-tags", "browser", "-f", each("GoFiles")+each("TestGoFiles")+each("XTestGoFiles"), "./...")
	if err != nil {
		return err
	}
	unformatted, err := run("gofmt", append([]string{"-l"}, strings.Fields(listed)...)...)
	if err != nil {
		return err
	}
	if strings.TrimSpace(unformatted) != "" {
		return fmt.Errorf("not gofmt'd (gofmt -w <file>):\n%s", unformatted)
	}
	if _, err := run("go", "vet", "-tags", "browser", "./..."); err != nil {
		return err
	}
	// @invariant: a component's field left out of a struct literal is a default nobody chose
	if _, err := run("go", "run", "./go/cmd/ciderui-check", "./..."); err != nil {
		return err
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
	if err := ownSheet(); err != nil {
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
	return nil
}

// @invariant: what a site with utilities of its own does: its Go tool writes
// utilities.css, Tailwind builds its second stylesheet from there. Everything else
// reads the checkout, so a stylesheet left out of go:embed fails here and nowhere else.
func ownSheet() error {
	written, err := os.MkdirTemp("", "ciderui-written-kit-")
	if err != nil {
		return err
	}
	defer os.RemoveAll(written)
	if _, err := run("go", "run", "./go/cmd/ciderui-css", written); err != nil {
		return err
	}
	built, err := build.OwnSheet(filepath.Join(written, "utilities.css"), "px-8 md:px-8 cui-btn cui-headline")
	if err != nil {
		return err
	}
	for _, rule := range []string{".px-8", `.md\:px-8`} {
		if !strings.Contains(built, rule) {
			return fmt.Errorf("utilities.css written by ciderui-css built with no %s: %.200s", rule, built)
		}
	}
	for _, kit := range []string{".cui-btn", ".cui-headline"} {
		if strings.Contains(built, kit) {
			return fmt.Errorf("utilities.css built the kit's %s, which the kit's own sheet ships", kit)
		}
	}
	return nil
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
		io.Copy(io.Discard, pipe)
		cmd.Wait()
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
