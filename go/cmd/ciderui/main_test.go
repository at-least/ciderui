package main

import (
	"bytes"
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

const runMain = "CIDERUI_RUN_MAIN"

func TestMain(m *testing.M) {
	if os.Getenv(runMain) == "1" {
		main()
		os.Exit(0)
	}
	os.Exit(m.Run())
}

const (
	completeBtn    = `site.Btn{Href: "", Label: "Buy", Icon: nil, Submit: false, Variant: site.BtnVariantPrimary, Size: site.BtnSizeMedium, Block: false, Attrs: nil, Class: ""}.Render(os.Stdout)`
	completeBanner = `console.ConsoleBanner{Color: console.ConsoleColorInfo, Title: "Paid", Body: nil, Actions: nil}.Render(os.Stdout)`
)

func siteModule(t *testing.T, statement string) string {
	t.Helper()
	checkout, err := filepath.Abs("../../..")
	if err != nil {
		t.Fatal(err)
	}
	dir := t.TempDir()
	for name, content := range map[string]string{
		"go.mod":  "module example.com/store\n\ngo 1.26.0\n\nrequire github.com/at-least/ciderui v0.0.0\n\nreplace github.com/at-least/ciderui => " + checkout + "\n",
		"main.go": "package main\n\nimport (\n\t\"os\"\n\n\t\"github.com/at-least/ciderui/go/console\"\n\t\"github.com/at-least/ciderui/go/site\"\n)\n\nvar _ = []any{os.Stdout, console.ConsoleColorInfo, site.BtnSizeMedium}\n\ntype order struct{ ID, Note string }\n\nfunc main() {\n\t" + statement + "\n}\n",
	} {
		if err := os.WriteFile(filepath.Join(dir, name), []byte(content), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	return dir
}

func run(t *testing.T, dir string, args ...string) (code int, stdout, stderr string) {
	t.Helper()
	cmd := exec.Command(os.Args[0], args...)
	cmd.Dir = dir
	// @trap: -mod=mod: the site's go.mod names ciderui alone and has no go.sum; without
	// it the command says "./... matched no packages".
	cmd.Env = append(os.Environ(), runMain+"=1", "GOFLAGS=-mod=mod")
	var out, errOut bytes.Buffer
	cmd.Stdout, cmd.Stderr = &out, &errOut
	err := cmd.Run()
	var exit *exec.ExitError
	if err != nil && !errors.As(err, &exit) {
		t.Fatal(err)
	}
	return cmd.ProcessState.ExitCode(), out.String(), errOut.String()
}

func TestCommandPassesACompleteLiteralAndAnotherModulesIncompleteOneSilently(t *testing.T) {
	dir := siteModule(t, completeBtn+"; "+completeBanner+`; _ = order{ID: "1"}`)
	if code, stdout, stderr := run(t, dir, "check", "./..."); code != 0 || stdout != "" || stderr != "" {
		t.Errorf("ciderui check ./... exited %d, wrote %q and %q to stderr, want 0 and nothing", code, stdout, stderr)
	}
}

func TestCommandFailsALiteralOfEitherSetWithAFieldLeftOutNamingTheTypeAndTheField(t *testing.T) {
	for statement, want := range map[string]string{
		strings.Replace(completeBtn, " Block: false,", "", 1):    "main.go:15:2: site.Btn is missing field Block\n",
		strings.Replace(completeBanner, ", Actions: nil", "", 1): "main.go:15:2: console.ConsoleBanner is missing field Actions\n",
	} {
		if code, stdout, stderr := run(t, siteModule(t, statement), "check", "./..."); code == 0 || stdout != "" || !strings.HasSuffix(stderr, want) {
			t.Errorf("ciderui check ./... exited %d, wrote %q and %q to stderr, want non-zero and a line ending %q", code, stdout, stderr, want)
		}
	}
}

func TestCommandWithoutAVerbItsArgumentsOrWithAFlagExitsTwoWithItsUsageAndWritesNothing(t *testing.T) {
	const want = "usage: ciderui check <packages>: fails a ciderui struct literal that leaves a field out\n" +
		"       ciderui css <out.css> <package dir>...: writes the utilities the directories' Go files use to <out.css>\n"
	dir := siteModule(t, strings.Replace(completeBtn, " Block: false,", "", 1))
	for _, args := range [][]string{nil, {"./..."}, {"check"}, {"check", "-h"}, {"check", "-e", ".*", "./..."}, {"css"}, {"css", "out.css"}, {"css", "-h", "."}, {"css", "out.css", "-h"}, {"lint", "./..."}} {
		if code, stdout, stderr := run(t, dir, args...); code != 2 || stdout != "" || stderr != want {
			t.Errorf("ciderui %v exited %d, wrote %q and %q to stderr, want 2 and %q", args, code, stdout, stderr, want)
		}
	}
	for _, name := range []string{"out.css", "-h"} {
		if _, err := os.Stat(filepath.Join(dir, name)); !errors.Is(err, os.ErrNotExist) {
			t.Errorf("ciderui with its usage wrote %s: %v", name, err)
		}
	}
}

func TestCSSOfAMissingDirectoryExitsOneNamingItAndWritesNothing(t *testing.T) {
	// @external: macOS's temporary directory is under a symlink, which the command's
	// working directory resolves
	dir, err := filepath.EvalSymlinks(t.TempDir())
	if err != nil {
		t.Fatal(err)
	}
	code, stdout, stderr := run(t, dir, "css", "out.css", "pages")
	if want := "css: stat " + filepath.Join(dir, "pages") + ": no such file or directory\n"; code != 1 || stdout != "" || stderr != want {
		t.Errorf("ciderui css out.css pages exited %d, wrote %q and %q to stderr, want 1 and %q", code, stdout, stderr, want)
	}
	if _, err := os.Stat(filepath.Join(dir, "out.css")); !errors.Is(err, os.ErrNotExist) {
		t.Errorf("ciderui css of a missing directory wrote out.css: %v", err)
	}
}

func TestCSSOfADirectoryWithNoGoFileButTestsExitsOneNamingItAndWritesNothing(t *testing.T) {
	for _, file := range []string{"page_test.go", "page.html"} {
		dir := t.TempDir()
		pages := filepath.Join(dir, "pages")
		if err := os.Mkdir(pages, 0o755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(filepath.Join(pages, file), []byte("package pages\n"), 0o644); err != nil {
			t.Fatal(err)
		}
		code, stdout, stderr := run(t, dir, "css", "out.css", "pages")
		if want := "css: pages holds no Go files\n"; code != 1 || stdout != "" || stderr != want {
			t.Errorf("ciderui css out.css pages, holding only %s, exited %d, wrote %q and %q to stderr, want 1 and %q", file, code, stdout, stderr, want)
		}
		if _, err := os.Stat(filepath.Join(dir, "out.css")); !errors.Is(err, os.ErrNotExist) {
			t.Errorf("ciderui css of a directory holding only %s wrote out.css: %v", file, err)
		}
	}
}

func TestCSSOfAFileExitsOneNamingIt(t *testing.T) {
	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "page.go"), []byte("package pages\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	code, stdout, stderr := run(t, dir, "css", "out.css", "page.go")
	if want := "css: page.go is not a directory\n"; code != 1 || stdout != "" || stderr != want {
		t.Errorf("ciderui css out.css page.go exited %d, wrote %q and %q to stderr, want 1 and %q", code, stdout, stderr, want)
	}
}

// @invariant: the command builds from the kit's embedded utilities.css and theme.css,
// where internal/build reads the checkout's: this test is the build that sees go:embed
func TestCSSWritesTheUtilitiesOfEachDirectorysGoFilesMinifiedLeavingOutTestFilesAndTheKitsOwn(t *testing.T) {
	dir := t.TempDir()
	commandsTMPDIR := filepath.Join(dir, "tmp")
	if err := os.Mkdir(commandsTMPDIR, 0o755); err != nil {
		t.Fatal(err)
	}
	t.Setenv("TMPDIR", commandsTMPDIR)
	for file, class := range map[string]string{"pages/page.go": "px-8 md:px-8 cui-btn cui-headline", "pages/page_test.go": "px-9", "admin/admin.go": "py-7"} {
		path := filepath.Join(dir, file)
		if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(path, []byte("package p\n\nconst class = \""+class+"\"\n"), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	if code, stdout, stderr := run(t, dir, "css", "out.css", "pages", "admin"); code != 0 || stdout != "" || stderr != "" {
		t.Fatalf("ciderui css out.css pages admin exited %d, wrote %q and %q to stderr, want 0 and nothing", code, stdout, stderr)
	}
	b, err := os.ReadFile(filepath.Join(dir, "out.css"))
	if err != nil {
		t.Fatal(err)
	}
	built := string(b)
	for _, rule := range []string{".px-8{", `.md\:px-8{`, ".py-7{"} {
		if !strings.Contains(built, rule) {
			t.Errorf("ciderui css built no %s, which a page's Go writes: %.300s", rule, built)
		}
	}
	for _, rule := range []string{".px-9", ".cui-btn", ".cui-headline"} {
		if strings.Contains(built, rule) {
			t.Errorf("ciderui css built %s, which only a _test.go or the kit's own sheet holds", rule)
		}
	}
	if strings.Contains(built, "{\n") {
		t.Errorf("ciderui css wrote an unminified sheet: %.300s", built)
	}
	if left, err := os.ReadDir(commandsTMPDIR); err != nil || len(left) != 0 {
		t.Errorf("ciderui css left %v in its temporary directory (%v)", left, err)
	}
}

func TestCSSThatTailwindFailsToWriteExitsOneWithTailwindsMessage(t *testing.T) {
	dir := t.TempDir()
	for _, sub := range []string{"pages", "out.css"} {
		if err := os.Mkdir(filepath.Join(dir, sub), 0o755); err != nil {
			t.Fatal(err)
		}
	}
	if err := os.WriteFile(filepath.Join(dir, "pages", "page.go"), []byte("package pages\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	code, stdout, stderr := run(t, dir, "css", "out.css", "pages")
	if code != 1 || stdout != "" || !strings.HasPrefix(stderr, "css: tailwindcss: exit status ") || !strings.Contains(stderr, "EISDIR") {
		t.Errorf("ciderui css onto a directory exited %d, wrote %q and %q to stderr, want 1 and Tailwind's EISDIR", code, stdout, stderr)
	}
}
