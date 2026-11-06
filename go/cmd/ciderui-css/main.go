// Command ciderui-css writes what a site's own utilities' stylesheet imports:
//
//	go tool ciderui-css target/ciderui
package main

import (
	"fmt"
	"os"

	"github.com/at-least/ciderui"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "usage: ciderui-css <dir>")
		os.Exit(2)
	}
	if err := ciderui.WriteCSS(os.Args[1]); err != nil {
		fmt.Fprintf(os.Stderr, "writing %s: %v\n", os.Args[1], err)
		os.Exit(1)
	}
	fmt.Printf("ciderui's utilities.css written under %s\n", os.Args[1])
}
