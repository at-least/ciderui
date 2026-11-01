package ciderui_test

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/at-least/ciderui"
)

func ExampleCheckMarkup() {
	page := `<a class="cui-console-btn" data-variant="plian" href="/">Open</a>`
	fmt.Println(ciderui.CheckMarkup(page))
	// Output: [data-variant="plian" on <a class="cui-console-btn">: no rule of ciderui's selects it (cui-console-btn's: filled, plain)]
}

func ExampleWriteCSS() {
	dir, err := os.MkdirTemp("", "site")
	if err != nil {
		panic(err)
	}
	defer os.RemoveAll(dir)

	if err := ciderui.WriteCSS(filepath.Join(dir, "target/ciderui")); err != nil {
		panic(err)
	}
	_, err = os.Stat(filepath.Join(dir, "target/ciderui/index.css"))
	fmt.Println(err == nil)
	// Output: true
}
