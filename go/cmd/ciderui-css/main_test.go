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

const runMain = "CIDERUI_CSS_RUN_MAIN"

func TestMain(m *testing.M) {
	if os.Getenv(runMain) == "1" {
		main()
		os.Exit(0)
	}
	os.Exit(m.Run())
}

func run(t *testing.T, args ...string) (code int, stdout, stderr string) {
	t.Helper()
	cmd := exec.Command(os.Args[0], args...)
	cmd.Dir = t.TempDir()
	cmd.Env = append(os.Environ(), runMain+"=1")
	var out, errOut bytes.Buffer
	cmd.Stdout, cmd.Stderr = &out, &errOut
	err := cmd.Run()
	var exit *exec.ExitError
	if err != nil && !errors.As(err, &exit) {
		t.Fatal(err)
	}
	return cmd.ProcessState.ExitCode(), out.String(), errOut.String()
}

func TestCommandWritesTheUtilitiesUnderItsDir(t *testing.T) {
	dir := filepath.Join(t.TempDir(), "target", "ciderui")
	code, stdout, stderr := run(t, dir)
	if want := "ciderui's utilities.css written under " + dir + "\n"; code != 0 || stdout != want || stderr != "" {
		t.Errorf("ciderui-css %s exited %d, wrote %q and %q to stderr, want 0 and %q", dir, code, stdout, stderr, want)
	}
	if _, err := os.Stat(filepath.Join(dir, "utilities.css")); err != nil {
		t.Errorf("ciderui-css wrote no utilities.css: %v", err)
	}
}

func TestCommandWithoutOneDirExitsTwoWithItsUsage(t *testing.T) {
	for _, args := range [][]string{nil, {"a", "b"}} {
		if code, stdout, stderr := run(t, args...); code != 2 || stdout != "" || stderr != "usage: ciderui-css <dir>\n" {
			t.Errorf("ciderui-css %v exited %d, wrote %q and %q to stderr, want 2 and its usage", args, code, stdout, stderr)
		}
	}
}

func TestCommandThatCannotWriteExitsOneNamingTheDir(t *testing.T) {
	blocked := filepath.Join(t.TempDir(), "not-a-dir")
	if err := os.WriteFile(blocked, nil, 0o644); err != nil {
		t.Fatal(err)
	}
	dir := filepath.Join(blocked, "under")
	if code, stdout, stderr := run(t, dir); code != 1 || stdout != "" || !strings.Contains(stderr, dir) {
		t.Errorf("ciderui-css under a file exited %d, wrote %q and %q to stderr, want 1 and an error naming %s", code, stdout, stderr, dir)
	}
}
