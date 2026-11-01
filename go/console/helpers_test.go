package console_test

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func mustPanic(t *testing.T, want string, render func()) {
	t.Helper()
	defer func() {
		r := recover()
		if r == nil || !strings.Contains(fmt.Sprint(r), want) {
			t.Errorf("panicked with %v, want %q", r, want)
		}
	}()
	render()
}

// the example at path, under components/, as its file holds it: not trimmed
// as example.Read trims it (go test runs in the package's directory)
func readRaw(path string) string {
	b, err := os.ReadFile(filepath.Join("../../components", path))
	if err != nil {
		panic(err)
	}
	return string(b)
}
