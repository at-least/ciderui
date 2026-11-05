package console_test

import (
	"os"
	"path/filepath"
	"strings"
)

// not example.Read: it trims each line
func readRaw(path string) string {
	b, err := os.ReadFile(filepath.Join("../../components", path))
	if err != nil {
		panic(err)
	}
	return string(b)
}

// the src of the image of alt in the example at path, a data: URL too long to copy
func exampleSrc(path, alt string) string {
	raw := readRaw(path)
	at := strings.Index(raw, `" alt="`+alt+`"`)
	start := strings.LastIndex(raw[:at], `src="`) + len(`src="`)
	return raw[start:at]
}
