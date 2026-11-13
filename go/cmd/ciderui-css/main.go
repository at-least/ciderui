// Command ciderui-css writes what a site's own utilities' stylesheet imports:
//
//	go run github.com/at-least/ciderui/go/cmd/ciderui-css target/ciderui
package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/at-least/ciderui"
)

const usage = "usage: ciderui-css <dir>: writes utilities.css and theme.css into <dir>"

func main() {
	if len(os.Args) != 2 || strings.HasPrefix(os.Args[1], "-") {
		fmt.Fprintln(os.Stderr, usage)
		os.Exit(2)
	}
	if err := ciderui.WriteCSS(os.Args[1]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
