package main

import (
	"strings"
	"testing"
)

// @trap: go writes "go: downloading …" to stderr on a machine whose module cache lacks
// a module; merged into stdout it became gofmt's file list (mini, 2026-11-15).
func TestRunReturnsStdoutAlone(t *testing.T) {
	got, err := run("sh", "-c", "echo noise >&2; echo data")
	if err != nil {
		t.Fatal(err)
	}
	if got != "data\n" {
		t.Fatalf("run returned %q, want %q", got, "data\n")
	}
}

func TestRunErrorCarriesStderr(t *testing.T) {
	_, err := run("sh", "-c", "echo why >&2; exit 3")
	if err == nil || !strings.Contains(err.Error(), "why") {
		t.Fatalf("run's error %v does not carry stderr", err)
	}
}
