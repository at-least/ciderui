// Command build-css writes dist/<set>.css, the stylesheet a site links, from the
// sources; with --check it fails where dist/ is stale.
package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/at-least/ciderui/internal/build"
)

const usage = "usage: go run ./scripts/build-css [--check]"

func main() {
	args := os.Args[1:]
	if len(args) > 1 || (len(args) == 1 && args[0] != "--check") {
		fmt.Fprintln(os.Stderr, usage)
		os.Exit(2)
	}
	checking := len(args) == 1
	stale := false
	for _, set := range build.Sets {
		path := filepath.Join(build.Root, "dist", set+".css")
		css, err := build.Dist(set)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		if !checking {
			if err := os.WriteFile(path, []byte(css), 0o644); err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
			continue
		}
		have, err := os.ReadFile(path)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		if string(have) != css {
			fmt.Printf("dist/%s.css is stale: go run ./scripts/build-css\n", set)
			stale = true
		}
	}
	if stale {
		os.Exit(1)
	}
}
