// Command fetch fetches what the checks run beside the module's own code: the
// pinned tools (internal/pinned) and Playwright's driver, the Node.js it runs
// on with it. The test image's build runs it, so a run in the image reaches no
// network for them.
package main

import (
	"fmt"
	"os"

	"github.com/at-least/ciderui/internal/pinned"
	"github.com/mxschmitt/playwright-go"
)

func main() {
	if len(os.Args) != 1 {
		fmt.Fprintln(os.Stderr, "usage: go run ./scripts/fetch")
		os.Exit(2)
	}
	if err := pinned.Fetch(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	// @external: the browsers are the image's own, Playwright's at the driver's version
	if err := playwright.Install(&playwright.RunOptions{SkipInstallBrowsers: true}); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
