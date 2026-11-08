package glyph_test

import (
	"os"
	"regexp"
	"strconv"
	"strings"
	"testing"

	"github.com/at-least/ciderui/go/internal/glyph"
)

func TestClassEscapesItsClassOnTheSVG(t *testing.T) {
	var b strings.Builder
	if err := glyph.Menu.Class(`a"b<c&d`).Render(&b); err != nil {
		t.Fatal(err)
	}
	if got, want := b.String(), `<svg class="a&#34;b&lt;c&amp;d"`; !strings.HasPrefix(got, want) {
		t.Errorf("the icon starts %s, want its class escaped: %s", got, want)
	}
}

// globalnav.css turns Menu's two lines into the close button's X, each about its
// own middle, onto one point: a move of a line must move those transforms too
func TestGlobalnavCSSTurnsMenusLinesIntoAnX(t *testing.T) {
	stylesheet, err := os.ReadFile("../../../components/site/globalnav/globalnav.css")
	if err != nil {
		t.Fatal(err)
	}
	var menu strings.Builder
	if err := glyph.Menu.Render(&menu); err != nil {
		t.Fatal(err)
	}
	lines := regexp.MustCompile(`M(\d+) (\d+)H(\d+)`).FindAllStringSubmatch(menu.String(), -1)
	turns := regexp.MustCompile(`transform-origin: (\d+)px (\d+)px;\s+transform: translate\((-?[\d.]+)px, (-?[\d.]+)px\) rotate`).FindAllStringSubmatch(string(stylesheet), -1)
	if len(lines) != 2 || len(turns) != 2 {
		t.Fatalf("%d lines in Menu, %d transforms in globalnav.css", len(lines), len(turns))
	}
	number := func(s string) float64 {
		n, err := strconv.ParseFloat(s, 64)
		if err != nil {
			t.Fatal(err)
		}
		return n
	}
	var crossings [2][2]float64
	for i := range lines {
		middle := (number(lines[i][1]) + number(lines[i][3])) / 2
		y := number(lines[i][2])
		if number(turns[i][1]) != middle || number(turns[i][2]) != y {
			t.Errorf("line %d's middle is %v %v, its transform-origin %s %s", i, middle, y, turns[i][1], turns[i][2])
		}
		crossings[i] = [2]float64{middle + number(turns[i][3]), y + number(turns[i][4])}
	}
	if crossings[0] != crossings[1] {
		t.Errorf("the X's lines cross at no one point: %v", crossings)
	}
}
