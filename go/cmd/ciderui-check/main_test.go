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

const runMain = "CIDERUI_CHECK_RUN_MAIN"

func TestMain(m *testing.M) {
	if os.Getenv(runMain) == "1" {
		main()
		os.Exit(0)
	}
	os.Exit(m.Run())
}

const complete = `Href: "", Label: "Buy", Icon: nil, Submit: false, Variant: site.BtnVariantPrimary, Size: site.BtnSizeMedium, Block: false, Attrs: nil, Class: ""`

// siteModule is a site's module that requires this checkout, its one file a site.Btn
// literal of fields.
func siteModule(t *testing.T, fields string) string {
	t.Helper()
	checkout, err := filepath.Abs("../../..")
	if err != nil {
		t.Fatal(err)
	}
	dir := t.TempDir()
	for name, content := range map[string]string{
		"go.mod":  "module example.com/store\n\ngo 1.26.0\n\nrequire github.com/at-least/ciderui v0.0.0\n\nreplace github.com/at-least/ciderui => " + checkout + "\n",
		"main.go": "package main\n\nimport (\n\t\"os\"\n\n\t\"github.com/at-least/ciderui/go/site\"\n)\n\nfunc main() {\n\tsite.Btn{" + fields + "}.Render(os.Stdout)\n}\n",
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
	// -mod=mod: the site's go.mod names ciderui alone and has no go.sum; without it
	// the command says "./... matched no packages".
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

func TestCommandPassesALiteralOfEveryFieldSilently(t *testing.T) {
	if code, stdout, stderr := run(t, siteModule(t, complete), "./..."); code != 0 || stdout != "" || stderr != "" {
		t.Errorf("ciderui-check ./... exited %d, wrote %q and %q to stderr, want 0 and nothing", code, stdout, stderr)
	}
}

func TestCommandFailsALiteralWithAFieldLeftOutNamingTheTypeAndTheField(t *testing.T) {
	dir := siteModule(t, strings.Replace(complete, " Block: false,", "", 1))
	const want = "main.go:10:2: site.Btn is missing field Block\n"
	if code, stdout, stderr := run(t, dir, "./..."); code == 0 || stdout != "" || !strings.HasSuffix(stderr, want) {
		t.Errorf("ciderui-check ./... exited %d, wrote %q and %q to stderr, want non-zero and a line ending %q", code, stdout, stderr, want)
	}
}

func TestCommandWithNoPackagesOrAFlagExitsTwoWithItsUsage(t *testing.T) {
	dir := siteModule(t, strings.Replace(complete, " Block: false,", "", 1))
	for _, args := range [][]string{nil, {"-h"}, {"-e", ".*", "./..."}, {"-test=false", "./..."}} {
		if code, stdout, stderr := run(t, dir, args...); code != 2 || stdout != "" || stderr != usage+"\n" {
			t.Errorf("ciderui-check %v exited %d, wrote %q and %q to stderr, want 2 and its usage", args, code, stdout, stderr)
		}
	}
}
