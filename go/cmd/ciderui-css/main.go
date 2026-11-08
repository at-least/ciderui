// Command ciderui-css writes what a site's own utilities' stylesheet imports:
//
//	go run github.com/at-least/ciderui/go/cmd/ciderui-css target/ciderui
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
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Printf("ciderui's utilities.css written under %s\n", os.Args[1])
}
