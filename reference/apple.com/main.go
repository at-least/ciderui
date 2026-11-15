// Command apple.com: PAGE=<name> go run ./reference/apple.com: pages/<name>.go's url at
// each of its widths, each probe's computed styles into results/<name>.go, and the
// registry, results/pages.go. A probe's keys are what this file reads. A probe that
// matches nothing is recorded as Missing and fails the run.
// go run ./reference/apple.com report <name> [probe-prefix]: each probe's values across
// the widths of results/<name>.go, a line per value that is read.
// What reading it cannot show:
//   - browser `chrome`: Google Chrome plays what needs Widevine (chrome.Dockerfile's
//     image, CIDERUI_CHROME=1).
//   - select's ` >>> ` steps into a shadow root: tv.apple.com's controls keep their
//     parts behind theirs. inject: for a class the page defines but may not use.
//   - states are forced through the DevTools protocol, transitions off, animations left
//     running: ended, a filled animation's value read as its base's, the state's own as
//     it seemed, and it ran again once the still style went (ribbon:ribbon.link,
//     2026-10-22). A frame's element is found in the page's own session, so only a
//     frame in the page's process can have its states forced; another fails the run.
//   - set's `value` sets an input's value with no event: the page's script hears nothing.
//   - click leaves the page so for later probes and widths, so a spec puts such probes
//     last or in a spec of their own. The pointer then leaves for the window's bottom
//     corner, or `pointer` inside what opened, where leaving it closes it.
//   - tap is on a touch spec alone, through the DevTools protocol's touch events, which
//     Chromium's gesture recognizer makes a tap of; a page scrolled after it keeps what
//     it hovered.
//   - hover moves in twelve steps, as a hand does: apple.com's script opens a flyout on
//     its own mouseover, 120ms on.
//   - image never records the picture, the page's asset: what it draws (viewBox, shapes,
//     paint); any other value than a base64 data: SVG fails the run.
//   - touch: the DevTools protocol's emulation before a reload makes (hover: none) match;
//     the run fails if (hover: hover) still matches.
package main

import (
	"fmt"
	"os"
)

func main() {
	var err error
	if len(os.Args) > 1 && os.Args[1] == "report" {
		if len(os.Args) < 3 || len(os.Args) > 4 {
			fmt.Fprintln(os.Stderr, "usage: go run ./reference/apple.com report <name> [probe-prefix]")
			os.Exit(2)
		}
		prefix := ""
		if len(os.Args) == 4 {
			prefix = os.Args[3]
		}
		err = report(os.Args[2], prefix)
	} else if len(os.Args) > 1 {
		fmt.Fprintln(os.Stderr, "usage: PAGE=<name> go run ./reference/apple.com, or go run ./reference/apple.com report <name> [probe-prefix]")
		os.Exit(2)
	} else {
		err = measure(os.Getenv("PAGE"))
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
