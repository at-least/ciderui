// Command ciderui-check fails a struct literal of this module's types that leaves a
// field out, in a site's packages:
//
//	go get -tool github.com/at-least/ciderui/go/cmd/ciderui-check
//	go tool ciderui-check ./...
package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/GaijinEntertainment/go-exhaustruct/v3/analyzer"
	"golang.org/x/tools/go/analysis/singlechecker"
)

const usage = "usage: ciderui-check <packages>: fails a ciderui struct literal that leaves a field out"

// @trap: singlechecker reads os.Args itself, and its flags (-e, -fix, -test=false) would
// each be a second way to run the one check, so a flag stops here.
func main() {
	if len(os.Args) < 2 || strings.HasPrefix(os.Args[1], "-") {
		fmt.Fprintln(os.Stderr, usage)
		os.Exit(2)
	}
	check, err := analyzer.NewAnalyzer([]string{`github\.com/at-least/ciderui/.*`}, nil)
	if err != nil {
		panic(err)
	}
	singlechecker.Main(check)
}
