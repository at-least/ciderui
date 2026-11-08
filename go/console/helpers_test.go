package console_test

import (
	"errors"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/at-least/ciderui/go/console"
	g "maragu.dev/gomponents"
)

// @trap: not the reader example.Check compares with: it trims each line
func readRaw(path string) string {
	b, err := os.ReadFile(filepath.Join("../../components", path))
	if err != nil {
		panic(err)
	}
	return string(b)
}

// @invariant: read from the example: a data: URL too long to copy
func exampleSrc(path, alt string) string {
	raw := readRaw(path)
	at := strings.Index(raw, `" alt="`+alt+`"`)
	start := strings.LastIndex(raw[:at], `src="`) + len(`src="`)
	return raw[start:at]
}

func TestAPagesNodesErrorIsReturned(t *testing.T) {
	failure := errors.New("no")
	failing := g.NodeFunc(func(io.Writer) error { return failure })
	part := console.ConsoleListItem{Href: "", Label: "x", Detail: "", Class: "", Children: failing}
	if err := part.Render(io.Discard); !errors.Is(err, failure) {
		t.Errorf("the list item returned %v, not its Children's error", err)
	}
}
