// Command ciderui is the kit's command a site runs, one `tool` line in its go.mod:
//
//	go get -tool github.com/at-least/ciderui/go/cmd/ciderui
//	go tool ciderui check ./...
//	go tool ciderui css public/utilities.css ./pages
package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"slices"
	"strings"

	"github.com/GaijinEntertainment/go-exhaustruct/v3/analyzer"
	"github.com/at-least/ciderui"
	"github.com/at-least/ciderui/internal/pinned"
	"golang.org/x/tools/go/analysis/singlechecker"
)

const usage = `usage: go tool ciderui check ./...: fails a ciderui struct literal that leaves a field out
       go tool ciderui css <out.css> <package dir>...: writes the utilities the directories' Go files use to <out.css>`

func main() {
	args := os.Args[1:]
	// @trap: no flags: singlechecker's (-e, -fix, -test=false) would each be a second
	// way to run the one check, and a flag given to css would be a file named -h
	flagged := slices.ContainsFunc(args, func(a string) bool { return strings.HasPrefix(a, "-") })
	switch {
	case !flagged && len(args) >= 2 && args[0] == "check":
		check(args[1:])
	case !flagged && len(args) >= 3 && args[0] == "css":
		if err := css(args[1], args[2:]); err != nil {
			fmt.Fprintln(os.Stderr, "css:", err)
			os.Exit(1)
		}
	default:
		fmt.Fprintln(os.Stderr, usage)
		os.Exit(2)
	}
}

func check(packages []string) {
	// @invariant: the repository's own test tooling (internal/) is no kit type a site
	// writes: a case leaves out what it does not use
	exhaustruct, err := analyzer.NewAnalyzer([]string{`github\.com/at-least/ciderui/.*`}, []string{`github\.com/at-least/ciderui/internal/.*`})
	if err != nil {
		panic(err)
	}
	// @external: singlechecker reads os.Args itself
	os.Args = append([]string{os.Args[0]}, packages...)
	singlechecker.Main(exhaustruct)
}

func css(out string, dirs []string) error {
	// @trap: each directory checked here: Tailwind scans a missing one, or one of no Go
	// files, as empty, and the sheet would lack every utility its pages use, silently
	var sources []string
	for _, dir := range dirs {
		abs, err := filepath.Abs(dir)
		if err != nil {
			return err
		}
		info, err := os.Stat(abs)
		if err != nil {
			return err
		}
		if !info.IsDir() {
			return fmt.Errorf("%s is not a directory", dir)
		}
		files, err := filepath.Glob(filepath.Join(abs, "*.go"))
		if err != nil {
			return err
		}
		if !slices.ContainsFunc(files, func(f string) bool { return !strings.HasSuffix(f, "_test.go") }) {
			return fmt.Errorf("%s holds no Go files", dir)
		}
		sources = append(sources, fmt.Sprintf("%q", filepath.Join(abs, "*.go")), fmt.Sprintf("not %q", filepath.Join(abs, "*_test.go")))
	}
	tailwind, err := pinned.Tailwind()
	if err != nil {
		return err
	}
	tmp, err := os.MkdirTemp("", "ciderui-")
	if err != nil {
		return err
	}
	defer os.RemoveAll(tmp)
	if err := ciderui.WriteCSS(tmp); err != nil {
		return err
	}
	entry := filepath.Join(tmp, "entry.css")
	if err := os.WriteFile(entry, pinned.UtilitiesEntry(filepath.Join(tmp, "utilities.css"), sources), 0o644); err != nil {
		return err
	}
	// @external: Tailwind reports its progress on stderr; printed only on a failure,
	// so a build that worked prints nothing
	if msg, err := exec.Command(tailwind, "-i", entry, "-o", out, "--minify").CombinedOutput(); err != nil {
		return fmt.Errorf("tailwindcss: %w\n%s", err, bytes.TrimSpace(msg))
	}
	return nil
}
