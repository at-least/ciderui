// Command build-css writes dist/<set>.css, the stylesheet a site links, from the
// sources.
package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/at-least/ciderui/internal/build"
)

const usage = "usage: go run ./scripts/build-css"

func main() {
	if len(os.Args) != 1 {
		fmt.Fprintln(os.Stderr, usage)
		os.Exit(2)
	}
	for _, set := range build.Sets {
		css, err := build.Dist(set)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		if err := os.WriteFile(filepath.Join(build.Root, "dist", set+".css"), []byte(css), 0o644); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
}
